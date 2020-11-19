package scrapers

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/diadata-org/diadata/internal/pkg/exchange-scrapers/gnosis"
	"github.com/diadata-org/diadata/internal/pkg/exchange-scrapers/gnosis/token"

	"github.com/diadata-org/diadata/pkg/dia"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	gnosisStartBlock = uint64(10780772 - 5250)

	gnosisWsDial   = "ws://159.69.120.42:8546/"
	gnosisRestDial = "http://159.69.120.42:8545/"
)

type GnosisToken struct {
	Symbol   string
	Decimals uint8
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
	contract common.Address
}

func NewGnosisScraper(exchange dia.Exchange) *GnosisScraper {
	scraper := &GnosisScraper{
		exchangeName:   exchange.Name,
		contract:   exchange.Contract,
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

	go scraper.mainLoop()
	return scraper
}
func (scraper *GnosisScraper) loadTokens() {

	contract, err := gnosis.NewGnosisCaller(scraper.contract, scraper.RestClient)
	if err != nil {
		log.Error(err)
	}
	numTokens, err := contract.NumTokens(&bind.CallOpts{})
	if err != nil {
		log.Error(err)
	}

	for i := uint16(0); i < numTokens; i++ {
		tokenAddress, err := contract.TokenIdToAddressMap(&bind.CallOpts{}, i)

		if err != nil {
			log.Error(err)
		}
		tokenCaller, err := token.NewTokenCaller(tokenAddress, scraper.RestClient)
		if err != nil {
			log.Error(err)
		}
		symbol, err := tokenCaller.Symbol(&bind.CallOpts{})
		if err != nil {
			log.Error(err)
		}
		decimals, err := tokenCaller.Decimals(&bind.CallOpts{})
		if err != nil {
			log.Error(err)
		}
		scraper.tokens[i] = &GnosisToken{
			Symbol:   symbol,
			Decimals: decimals,
		}

		scraper.tokens[i].normalizeETH()

		fmt.Println(i, tokenAddress.Hex(), symbol, decimals)
	}
	fmt.Println("i, tokenAddress.Hex(), symbol, decimals")
}

func (scraper *GnosisScraper) subscribeToTrades() error {
	filterer, err := gnosis.NewGnosisFilterer(scraper.contract, scraper.WsClient)
	if err != nil {
		log.Error(err)
		return err
	}
	start := startBlock
	sink := make(chan *gnosis.GnosisTrade)
	sub, err := filterer.WatchTrade(&bind.WatchOpts{Start: &start}, sink, nil, nil, nil)
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
			case err := <-sub.Err():
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

func (scraper *GnosisScraper) processTrade(trade *gnosis.GnosisTrade) {
	symbol, foreignName, volume, price, err := scraper.getSwapDataGnosis(trade)
	timestamp := time.Now().Unix()
	if err != nil {
		log.Error(err)
	} else {
		if pairScraper, ok := scraper.pairScrapers[foreignName]; ok {

			trade := &dia.Trade{
				Symbol:         symbol,
				Pair:           pairScraper.pair.ForeignName,
				Price:          price,
				Volume:         volume,
				Time:           time.Unix(timestamp, 0),
				ForeignTradeID: "",
				Source:         scraper.exchangeName,
			}
			pairScraper.parent.chanTrades <- trade
			fmt.Println("got trade: ", trade)
		}
	}

}

func (scraper *GnosisScraper) mainLoop() {

	scraper.run = true

	scraper.subscribeToTrades()
	go func() {
		for scraper.run {
			_ = <-scraper.resubscribe
			if scraper.run {
				fmt.Println("resubscribe...")
				scraper.subscribeToTrades()
			}
		}
	}()

	for scraper.run {
		if len(scraper.pairScrapers) == 0 {
			scraper.error = errors.New("Gnosis: No pairs to scrape provided")
			log.Error(scraper.error.Error())
			break
		}

	}
	time.Sleep(time.Duration(10) * time.Second)
	if scraper.error == nil {
		scraper.error = errors.New("main loop terminated by Close()")
	}
	scraper.cleanup(nil)
}

// getSwapData returns the foreign name, volume and price of a swap
func (scraper *GnosisScraper) getSwapDataGnosis(s *gnosis.GnosisTrade) (symbol string, foreignName string, volume float64, price float64, err error) {
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

func (scraper *GnosisScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {

	pairSet := make(map[string]struct{})
	for _, p1 := range scraper.tokens {
		for _, p2 := range scraper.tokens {
			token1 := p1
			token2 := p2
			if token1 != token2 {

				foreignName := token1.Symbol + "-" + token2.Symbol
				if _, ok := pairSet[foreignName]; !ok {
					pairs = append(pairs, dia.Pair{
						Symbol:      token1.Symbol,
						ForeignName: foreignName,
						Exchange:    scraper.exchangeName,
						Ignore:      false,
					})
					pairSet[foreignName] = struct{}{}
				}

				foreignName = token2.Symbol + "-" + token1.Symbol
				if _, ok := pairSet[foreignName]; !ok {
					pairs = append(pairs, dia.Pair{
						Symbol:      token2.Symbol,
						ForeignName: foreignName,
						Exchange:    scraper.exchangeName,
						Ignore:      false,
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

func (scraper *GnosisScraper) NormalizePair(pair dia.Pair) (dia.Pair, error) {
	return dia.Pair{}, nil
}


func (scraper *GnosisScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {
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
	pair   dia.Pair
	closed bool
}

func (pairScraper *GnosisPairScraper) Pair() dia.Pair {
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
