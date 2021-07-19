package scrapers

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"

	gnosis2 "github.com/diadata-org/diadata/internal/exchange-scrapers/gnosis"
	token2 "github.com/diadata-org/diadata/internal/exchange-scrapers/gnosis/token"

	"github.com/diadata-org/diadata/pkg/dia"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	gnosisWsDial         = "ws://159.69.120.42:8546/"
	gnosisRestDial       = "http://159.69.120.42:8545/"
	gnosisLookBackBlocks = 6 * 60 * 24 * 7
)

type GnosisToken struct {
	Symbol   string
	Decimals uint8
	Address  string
	Name     string
}

type GnosisScraper struct {
	exchangeName string

	// channels to signal events
	run          bool
	initDone     chan nothing
	shutdown     chan nothing
	shutdownDone chan nothing

	errorLock sync.RWMutex
	error     error
	closed    bool

	pairScrapers   map[string]*GnosisPairScraper
	productPairIds map[string]int
	chanTrades     chan *dia.Trade

	WsClient    *ethclient.Client
	RestClient  *ethclient.Client
	resubscribe chan nothing
	tokens      map[uint16]*GnosisToken
	contract    common.Address
}

func NewGnosisScraper(exchange dia.Exchange, scrape bool) *GnosisScraper {
	scraper := &GnosisScraper{
		exchangeName:   exchange.Name,
		contract:       exchange.Contract,
		initDone:       make(chan nothing),
		shutdown:       make(chan nothing),
		shutdownDone:   make(chan nothing),
		productPairIds: make(map[string]int),
		pairScrapers:   make(map[string]*GnosisPairScraper),
		chanTrades:     make(chan *dia.Trade),
		resubscribe:    make(chan nothing),
		tokens:         make(map[uint16]*GnosisToken),
	}

	wsClient, err := ethclient.Dial(gnosisWsDial)
	if err != nil {
		log.Fatal(err)
	}
	scraper.WsClient = wsClient
	restClient, err := ethclient.Dial(gnosisRestDial)
	if err != nil {
		log.Fatal(err)
	}
	scraper.RestClient = restClient

	scraper.loadTokens()

	if scrape {
		go scraper.mainLoop()
	}
	return scraper
}
func (scraper *GnosisScraper) loadTokens() {
	fmt.Println("contract address: ", scraper.contract.String())
	contract, err := gnosis2.NewGnosisCaller(scraper.contract, scraper.RestClient)
	if err != nil {
		log.Error(err)
	}
	numTokens, err := contract.NumTokens(&bind.CallOpts{})
	if err != nil {
		log.Error(err)
	}
	count := uint16(0)
	for i := uint16(0); i < numTokens; i++ {
		tokenAddress, err := contract.TokenIdToAddressMap(&bind.CallOpts{}, i)
		if err != nil {
			log.Error(err)
			continue
		}
		tokenCaller, err := token2.NewTokenCaller(tokenAddress, scraper.RestClient)
		if err != nil {
			log.Error(err)
			continue
		}
		symbol, err := tokenCaller.Symbol(&bind.CallOpts{})
		if err != nil {
			log.Error(err)
			continue
		}
		decimals, err := tokenCaller.Decimals(&bind.CallOpts{})
		if err != nil {
			log.Error(err)
		}
		name, err := tokenCaller.Name(&bind.CallOpts{})
		if err != nil {
			log.Error(err)
		}

		scraper.tokens[i] = &GnosisToken{
			Symbol:   symbol,
			Decimals: uint8(decimals.Int64()),
			Address:  tokenAddress.String(),
			Name:     name,
		}
		scraper.tokens[count].normalizeETH()
		fmt.Println(count, tokenAddress.Hex(), symbol, decimals)
		count++
	}
	fmt.Println("i, tokenAddress.Hex(), symbol, decimals")
}

// FillSymbolData is not used by DEX scrapers.
func (s *GnosisScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{}, nil
}

func (scraper *GnosisScraper) subscribeToTrades() error {
	filterer, err := gnosis2.NewGnosisFilterer(scraper.contract, scraper.WsClient)
	if err != nil {
		log.Error(err)
		return err
	}

	header, err := scraper.RestClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	startblock := header.Number.Uint64() - uint64(gnosisLookBackBlocks)

	sink := make(chan *gnosis2.GnosisTrade)
	sub, err := filterer.WatchTrade(&bind.WatchOpts{Start: &startblock}, sink, nil, nil, nil)
	if err != nil {
		log.Error(err)
		return err
	}

	go func() {
		fmt.Println("Subscribed to trades")
		defer fmt.Println("Unsubscribed to trades")
		defer sub.Unsubscribe()
		subscribed := true

		for scraper.run && subscribed {

			select {
			case err = <-sub.Err():
				if err != nil {
					log.Error(err)
				}
				subscribed = false
				if scraper.run {
					scraper.resubscribe <- nothing{}
				}
			case trade := <-sink:
				scraper.processTrade(trade)
			}
		}
	}()

	return err
}

func (scraper *GnosisScraper) processTrade(trade *gnosis2.GnosisTrade) {
	symbol, foreignName, volume, price, txHash := scraper.getSwapDataGnosis(trade)
	timestamp := time.Now().Unix()

	if pairScraper, ok := scraper.pairScrapers[foreignName]; ok {

		buyToken := scraper.tokens[trade.BuyToken]
		sellToken := scraper.tokens[trade.SellToken]

		token0 := dia.Asset{
			Address: buyToken.Address,
			Symbol:  buyToken.Symbol,
			Name:    buyToken.Name,
		}
		token1 := dia.Asset{
			Address: sellToken.Address,
			Symbol:  sellToken.Symbol,
			Name:    sellToken.Name,
		}

		trade := &dia.Trade{
			Symbol:         symbol,
			Pair:           pairScraper.pair.ForeignName,
			Price:          price,
			BaseToken:      token0,
			QuoteToken:     token1,
			Volume:         volume,
			Time:           time.Unix(timestamp, 0),
			ForeignTradeID: txHash,
			Source:         scraper.exchangeName,
			VerifiedPair:   true,
		}
		pairScraper.parent.chanTrades <- trade
		fmt.Println("got trade: ", trade)
	}

}

func (scraper *GnosisScraper) mainLoop() {

	scraper.run = true
	log.Info("subscribe to trades...")
	err := scraper.subscribeToTrades()
	if err != nil {
		log.Error(err)
	}

	go func() {
		for scraper.run {
			<-scraper.resubscribe
			if scraper.run {
				fmt.Println("resubscribe...")
				err := scraper.subscribeToTrades()
				if err != nil {
					log.Error(err)
				}
			}
		}
	}()

	if scraper.run {
		if len(scraper.pairScrapers) == 0 {
			scraper.error = errors.New("no pairs to scrape provided")
			log.Error(scraper.error.Error())
		}
	}

	time.Sleep(10 * time.Second)

	if scraper.error == nil {
		scraper.error = errors.New("main loop terminated by Close()")
	}
	scraper.cleanup(nil)
}

// getSwapData returns the foreign name, volume and price of a swap
func (scraper *GnosisScraper) getSwapDataGnosis(s *gnosis2.GnosisTrade) (symbol string, foreignName string, volume float64, price float64, txHash string) {
	txHash = s.Raw.TxHash.Hex()
	buyToken := scraper.tokens[s.BuyToken]
	sellToken := scraper.tokens[s.SellToken]
	buyDecimals := buyToken.Decimals
	sellDecimals := sellToken.Decimals

	amountOut, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(s.ExecutedBuyAmount), new(big.Float).SetFloat64(math.Pow10(int(buyDecimals)))).Float64()

	amountIn, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(s.ExecutedSellAmount), new(big.Float).SetFloat64(math.Pow10(int(sellDecimals)))).Float64()

	volume = amountOut
	price = amountIn / amountOut
	foreignName = buyToken.Symbol + "-" + sellToken.Symbol
	symbol = buyToken.Symbol
	return
}

func (scraper *GnosisScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {

	pairSet := make(map[string]struct{})
	for _, p1 := range scraper.tokens {
		for _, p2 := range scraper.tokens {
			token1 := p1
			token2 := p2
			if token1 != token2 {

				foreignName := token1.Symbol + "-" + token2.Symbol
				if _, ok := pairSet[foreignName]; !ok {
					pairs = append(pairs, dia.ExchangePair{
						Symbol:      token1.Symbol,
						ForeignName: foreignName,
						Exchange:    scraper.exchangeName,
					})
					pairSet[foreignName] = struct{}{}
				}

				foreignName = token2.Symbol + "-" + token1.Symbol
				if _, ok := pairSet[foreignName]; !ok {
					pairs = append(pairs, dia.ExchangePair{
						Symbol:      token2.Symbol,
						ForeignName: foreignName,
						Exchange:    scraper.exchangeName,
					})
					pairSet[foreignName] = struct{}{}
				}

			}
		}
	}
	return
}

func (t *GnosisToken) normalizeETH() {
	if t.Symbol == "WETH" {
		t.Symbol = "ETH"
	}
}

func (scraper *GnosisScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return dia.ExchangePair{}, nil
}

func (scraper *GnosisScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	scraper.errorLock.RLock()
	defer scraper.errorLock.RUnlock()

	if scraper.error != nil {
		return nil, scraper.error
	}

	if scraper.closed {
		return nil, errors.New("GnosisScraper is closed")
	}

	pairScraper := &GnosisPairScraper{
		parent: scraper,
		pair:   pair,
	}

	scraper.pairScrapers[pair.ForeignName] = pairScraper

	return pairScraper, nil
}
func (scraper *GnosisScraper) cleanup(err error) {
	scraper.errorLock.Lock()
	defer scraper.errorLock.Unlock()
	if err != nil {
		scraper.error = err
	}
	scraper.closed = true
	close(scraper.shutdownDone)
}

func (scraper *GnosisScraper) Close() error {
	// close the pair scraper channels
	scraper.run = false
	for _, pairScraper := range scraper.pairScrapers {
		pairScraper.closed = true
	}
	scraper.WsClient.Close()
	scraper.RestClient.Close()

	close(scraper.shutdown)
	<-scraper.shutdownDone
	return nil
}

type GnosisPairScraper struct {
	parent *GnosisScraper
	pair   dia.ExchangePair
	closed bool
}

func (pairScraper *GnosisPairScraper) Pair() dia.ExchangePair {
	return pairScraper.pair
}

func (scraper *GnosisScraper) Channel() chan *dia.Trade {
	return scraper.chanTrades
}

func (pairScraper *GnosisPairScraper) Error() error {
	s := pairScraper.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

func (pairScraper *GnosisPairScraper) Close() error {
	pairScraper.parent.errorLock.RLock()
	defer pairScraper.parent.errorLock.RUnlock()
	pairScraper.closed = true
	return nil
}
