package scrapers

import (
	"context"
	"errors"
	"fmt"
	"github.com/diadata-org/diadata/internal/pkg/exchange-scrapers/zerox"
	"github.com/diadata-org/diadata/internal/pkg/exchange-scrapers/zerox/token"
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	zeroxContract = "0x61935CbDd02287B511119DDb11Aeb42F1593b7Ef"
	zeroxWsDial   = "ws://159.69.120.42:8546/"
	zeroxRestDial = "http://159.69.120.42:8545/"
	// zeroxRestDial       = "https://mainnet.infura.io/v3/251a25bd10b8460fa040bb7202e22571"
	// zeroxWsDial         = "wss://mainnet.infura.io/ws/v3/251a25bd10b8460fa040bb7202e22571"
	zeroxLookBackBlocks = 6 * 60 * 24
)

type ZeroxToken struct {
	Symbol     string
	Name       string
	Address    string
	Blockchain dia.BlockChain
	Decimals   uint8
}

type ZeroxScraper struct {
	exchangeName string

	// channels to signal events
	run          bool
	initDone     chan nothing
	shutdown     chan nothing
	shutdownDone chan nothing

	errorLock sync.RWMutex
	error     error
	closed    bool

	pairScrapers   map[string]*ZeroxPairScraper
	productPairIds map[string]int
	chanTrades     chan *dia.Trade

	WsClient    *ethclient.Client
	RestClient  *ethclient.Client
	resubscribe chan nothing
	tokens      map[string]dia.Asset
}

func NewZeroxScraper(exchange dia.Exchange, scrape bool) *ZeroxScraper {
	scraper := &ZeroxScraper{
		exchangeName:   exchange.Name,
		initDone:       make(chan nothing),
		shutdown:       make(chan nothing),
		shutdownDone:   make(chan nothing),
		productPairIds: make(map[string]int),
		pairScrapers:   make(map[string]*ZeroxPairScraper),
		chanTrades:     make(chan *dia.Trade),
		resubscribe:    make(chan nothing),
		tokens:         make(map[string]dia.Asset),
	}
	wsClient, err := ethclient.Dial(zeroxWsDial)
	if err != nil {
		log.Fatal(err)
	}
	scraper.WsClient = wsClient
	restClient, err := ethclient.Dial(zeroxRestDial)
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

func (scraper *ZeroxScraper) loadTokens() {

	// added by hand because the symbol method returns a bytes32 instead of string
	scraper.tokens["0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2"] = dia.Asset{
		Symbol:     "MKR",
		Name:       "Maker",
		Decimals:   18,
		Address:    "0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2",
		Blockchain: dia.ETHEREUM,
	}

	filterer, err := zerox.NewZeroxFilterer(common.HexToAddress(zeroxContract), scraper.WsClient)
	if err != nil {
		log.Error(err)

	}

	header, err := scraper.RestClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	startblock := header.Number.Uint64() - uint64(zeroxLookBackBlocks)

	it, err := filterer.FilterFill(&bind.FilterOpts{Start: startblock}, nil, nil, nil)
	if err != nil {
		log.Error(err)
	}

	for it.Next() {
		i, err := scraper.loadTokenData(common.BytesToAddress(it.Event.TakerAssetData))
		if err != nil {
			// skip non-existing token data
			continue
		}
		o, err := scraper.loadTokenData(common.BytesToAddress(it.Event.MakerAssetData))
		if err != nil {
			// skip non-existing token data
			continue
		}
		log.Printf("\n %v  -%v- %v -%v- %v %v ",
			common.BytesToAddress(it.Event.TakerAssetData).Hex(),
			i.Symbol, i.Decimals,
			o.Symbol, o.Decimals,
			common.BytesToAddress(it.Event.MakerAssetData).Hex(),
		)
	}

}

// FillSymbolData is not used by DEX scrapers.
func (scraper *ZeroxScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{}, nil
}

func (scraper *ZeroxScraper) loadTokenData(tokenAddress common.Address) (dia.Asset, error) {

	tokenStr := tokenAddress.Hex()
	if foundToken, ok := (scraper.tokens[tokenStr]); ok {
		return foundToken, nil
	}

	tokenCaller, err := token.NewTokenCaller(tokenAddress, scraper.RestClient)
	if err != nil {
		return dia.Asset{}, err
	}
	symbol, err := tokenCaller.Symbol(&bind.CallOpts{})
	if err != nil {
		return dia.Asset{}, err
	}
	name, err := tokenCaller.Name(&bind.CallOpts{})
	if err != nil {
		log.Error("get token name: ", err)
	}
	decimals, err := tokenCaller.Decimals(&bind.CallOpts{})
	if err != nil {
		return dia.Asset{}, err
	}
	dfToken := dia.Asset{
		Symbol:     symbol,
		Name:       name,
		Address:    tokenAddress.Hex(),
		Blockchain: dia.ETHEREUM,
		Decimals:   uint8(decimals.Int64()),
	}
	scraper.tokens[tokenStr] = normalizeETH(dfToken)
	return dfToken, err

}

func (scraper *ZeroxScraper) subscribeToTrades() error {

	filterer, err := zerox.NewZeroxFilterer(common.HexToAddress(zeroxContract), scraper.WsClient)
	if err != nil {
		log.Error(err)
		return err
	}
	header, err := scraper.RestClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	startblock := header.Number.Uint64() - uint64(15250)

	sink := make(chan *zerox.ZeroxFill)
	sub, err := filterer.WatchFill(&bind.WatchOpts{Start: &startblock}, sink, nil, nil, nil)
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

func (scraper *ZeroxScraper) processTrade(trade *zerox.ZeroxFill) {
	token0, token1, symbol, foreignName, volume, price, err := scraper.getFillDataZerox(trade)
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
				BaseToken:      token1,
				QuoteToken:     token0,
				VerifiedPair:   true,
			}
			pairScraper.parent.chanTrades <- trade
			fmt.Println("got trade: ", trade)
		}
	}

}

func (scraper *ZeroxScraper) mainLoop() {

	scraper.run = true

	err := scraper.subscribeToTrades()
	if err != nil {
		log.Error(err)
	}

	go func() {
		for scraper.run {
			<-scraper.resubscribe
			if scraper.run {
				fmt.Println("resubscribe...")
				err = scraper.subscribeToTrades()
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

// getfillData returns the foreign name, volume and price of a order fill
func (scraper *ZeroxScraper) getFillDataZerox(s *zerox.ZeroxFill) (buyToken, sellToken dia.Asset, symbol string, foreignName string, volume float64, price float64, err error) {

	buyToken, err = scraper.loadTokenData(common.BytesToAddress(s.MakerAssetData))
	if err != nil {
		log.Error(err)
	}
	sellToken, err = scraper.loadTokenData(common.BytesToAddress(s.TakerAssetData))
	if err != nil {
		log.Error(err)
	}
	buyDecimals := buyToken.Decimals
	sellDecimals := sellToken.Decimals

	amountOut, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(s.MakerAssetFilledAmount), new(big.Float).SetFloat64(math.Pow10(int(buyDecimals)))).Float64()

	amountIn, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(s.TakerAssetFilledAmount), new(big.Float).SetFloat64(math.Pow10(int(sellDecimals)))).Float64()

	volume = amountOut
	price = amountIn / amountOut
	foreignName = buyToken.Symbol + "-" + sellToken.Symbol
	symbol = buyToken.Symbol
	return
}

func (scraper *ZeroxScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {

	pairSet := make(map[string]struct{})
	for _, p1 := range scraper.tokens {
		if p1.Symbol == "" || p1.Symbol == "BPT" {
			continue
		}
		for _, p2 := range scraper.tokens {
			if p2.Symbol == "" || p2.Symbol == "BPT" {
				continue
			}
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

func (scraper *ZeroxScraper) FetchTickerData(symbol string) (dia.Asset, error) {
	return dia.Asset{}, nil
}

func normalizeETH(t dia.Asset) dia.Asset {
	if t.Symbol == "WETH" {
		t.Symbol = "ETH"
	}
	return t
}

func (scraper *ZeroxScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	scraper.errorLock.RLock()
	defer scraper.errorLock.RUnlock()

	if scraper.error != nil {
		return nil, scraper.error
	}

	if scraper.closed {
		return nil, errors.New("ZeroxScraper is closed")
	}

	pairScraper := &ZeroxPairScraper{
		parent: scraper,
		pair:   pair,
	}

	scraper.pairScrapers[pair.ForeignName] = pairScraper

	return pairScraper, nil
}
func (scraper *ZeroxScraper) cleanup(err error) {
	scraper.errorLock.Lock()
	defer scraper.errorLock.Unlock()
	if err != nil {
		scraper.error = err
	}
	scraper.closed = true
	close(scraper.shutdownDone)
}

func (scraper *ZeroxScraper) Close() error {
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
func (s *ZeroxScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return dia.ExchangePair{}, nil
}

type ZeroxPairScraper struct {
	parent *ZeroxScraper
	pair   dia.ExchangePair
	closed bool
}

func (pairScraper *ZeroxPairScraper) Pair() dia.ExchangePair {
	return pairScraper.pair
}

func (scraper *ZeroxScraper) Channel() chan *dia.Trade {
	return scraper.chanTrades
}

func (pairScraper *ZeroxPairScraper) Error() error {
	s := pairScraper.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

func (pairScraper *ZeroxPairScraper) Close() error {
	pairScraper.parent.errorLock.RLock()
	defer pairScraper.parent.errorLock.RUnlock()
	pairScraper.closed = true
	return nil
}
