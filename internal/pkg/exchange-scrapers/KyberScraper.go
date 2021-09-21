package scrapers

import (
	"context"
	"errors"
	"fmt"
	"github.com/diadata-org/diadata/internal/pkg/exchange-scrapers/kyber"
	"github.com/diadata-org/diadata/internal/pkg/exchange-scrapers/kyber/token"
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
	kyberContract       = "0x9AAb3f75489902f3a48495025729a0AF77d4b11e"
	kyberWsDial         = "ws://159.69.120.42:8546/"
	kyberRestDial       = "http://159.69.120.42:8545/"
	kyberLookBackBlocks = 6 * 60 * 24
)

type KyberToken struct {
	Symbol   string
	Decimals *big.Int
}

type KyberScraper struct {
	exchangeName string

	// channels to signal events
	run          bool
	initDone     chan nothing
	shutdown     chan nothing
	shutdownDone chan nothing

	errorLock sync.RWMutex
	error     error
	closed    bool

	pairScrapers   map[string]*KyberPairScraper
	productPairIds map[string]int
	chanTrades     chan *dia.Trade

	WsClient    *ethclient.Client
	RestClient  *ethclient.Client
	resubscribe chan nothing
	tokens      map[string]dia.Asset
}

func NewKyberScraper(exchange dia.Exchange, scrape bool) *KyberScraper {
	scraper := &KyberScraper{
		exchangeName:   exchange.Name,
		initDone:       make(chan nothing),
		shutdown:       make(chan nothing),
		shutdownDone:   make(chan nothing),
		productPairIds: make(map[string]int),
		pairScrapers:   make(map[string]*KyberPairScraper),
		chanTrades:     make(chan *dia.Trade),
		resubscribe:    make(chan nothing),
		tokens:         make(map[string]dia.Asset),
	}
	wsClient, err := ethclient.Dial(kyberWsDial)
	if err != nil {
		log.Fatal(err)
	}
	scraper.WsClient = wsClient
	restClient, err := ethclient.Dial(kyberRestDial)
	if err != nil {
		log.Fatal(err)
	}
	scraper.RestClient = restClient

	scraper.loadTokens()
	time.Sleep(5 * time.Second)

	if scrape {
		go scraper.mainLoop()
	}
	return scraper
}

func (scraper *KyberScraper) loadTokens() {

	scraper.tokens["0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"] = dia.Asset{
		Symbol:   "ETH",
		Decimals: uint8(18),
	}

	// added by hand because the symbol method returns a bytes32 instead of string
	scraper.tokens["0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2"] = dia.Asset{
		Symbol:   "MKR",
		Decimals: uint8(18),
	}

	filterer, err := kyber.NewKyberFilterer(common.HexToAddress(kyberContract), scraper.WsClient)
	if err != nil {
		log.Error(err)

	}
	header, err := scraper.RestClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	startblock := header.Number.Uint64() - uint64(kyberLookBackBlocks)

	it, err := filterer.FilterExecuteTrade(&bind.FilterOpts{Start: startblock}, nil)
	if err != nil {
		log.Error(err)
	}

	for it.Next() {

		i, _ := scraper.loadTokenData(it.Event.Src)
		o, _ := scraper.loadTokenData(it.Event.Dest)
		log.Printf("\n %v  -%v- %v -%v- %v %v",
			it.Event.Src.Hex(),
			i.Symbol, i.Decimals,
			o.Symbol, o.Decimals,
			it.Event.Dest.Hex())
	}

}

func (scraper *KyberScraper) loadTokenData(tokenAddress common.Address) (dia.Asset, error) {
	tokenStr := tokenAddress.Hex()
	if foundToken, ok := (scraper.tokens[tokenStr]); ok {
		return foundToken, nil
	} else {
		tokenCaller, err := token.NewTokenCaller(tokenAddress, scraper.RestClient)
		if err != nil {
			log.Error(err)
		}
		symbol, err := tokenCaller.Symbol(&bind.CallOpts{})
		if err != nil {
			log.Error(err)
		}
		name, err := tokenCaller.Name(&bind.CallOpts{})
		if err != nil {
			log.Error(err)
		}
		decimals, err := tokenCaller.Decimals(&bind.CallOpts{})
		if err != nil {
			log.Error(err)
		}
		dfToken := dia.Asset{
			Symbol:     symbol,
			Name:       name,
			Decimals:   uint8(decimals.Int64()),
			Address:    tokenAddress.Hex(),
			Blockchain: dia.ETHEREUM,
		}
		scraper.tokens[tokenStr] = normalizeETH(dfToken)
		return dfToken, err
	}
}

func (scraper *KyberScraper) subscribeToTrades() error {

	filterer, err := kyber.NewKyberFilterer(common.HexToAddress(kyberContract), scraper.WsClient)
	if err != nil {
		log.Error(err)
		return err
	}
	header, err := scraper.RestClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	startblock := header.Number.Uint64() - uint64(5250)

	sink := make(chan *kyber.KyberExecuteTrade)
	sub, err := filterer.WatchExecuteTrade(&bind.WatchOpts{Start: &startblock}, sink, nil)
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

func (scraper *KyberScraper) processTrade(trade *kyber.KyberExecuteTrade) {
	symbol, foreignName, volume, price, token0, token1, err := scraper.getTradeDataKyber(trade)
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

func (scraper *KyberScraper) mainLoop() {

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

// getTradeData returns the foreign name, volume and price of a swap
func (scraper *KyberScraper) getTradeDataKyber(s *kyber.KyberExecuteTrade) (symbol string, foreignName string, volume float64, price float64, buyToken, sellToken dia.Asset, err error) {
	buyToken, err = scraper.loadTokenData(s.Dest)
	if err != nil {
		log.Error(err)
	}
	sellToken, err = scraper.loadTokenData(s.Src)
	if err != nil {
		log.Error(err)
	}
	buyDecimals := buyToken.Decimals
	sellDecimals := sellToken.Decimals

	amountOut, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(s.ActualDestAmount), new(big.Float).SetFloat64(math.Pow10(int(buyDecimals)))).Float64()

	amountIn, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(s.ActualSrcAmount), new(big.Float).SetFloat64(math.Pow10(int(sellDecimals)))).Float64()

	volume = amountOut
	price = amountIn / amountOut
	foreignName = buyToken.Symbol + "-" + sellToken.Symbol
	symbol = buyToken.Symbol
	return
}

func (scraper *KyberScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return dia.ExchangePair{}, nil
}

func (scraper *KyberScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {

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

func (scraper *KyberScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{}, nil
}

func (scraper *KyberScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	scraper.errorLock.RLock()
	defer scraper.errorLock.RUnlock()

	if scraper.error != nil {
		return nil, scraper.error
	}

	if scraper.closed {
		return nil, errors.New("KyberScraper is closed")
	}

	pairScraper := &KyberPairScraper{
		parent: scraper,
		pair:   pair,
	}

	scraper.pairScrapers[pair.ForeignName] = pairScraper

	return pairScraper, nil
}
func (scraper *KyberScraper) cleanup(err error) {
	scraper.errorLock.Lock()
	defer scraper.errorLock.Unlock()
	if err != nil {
		scraper.error = err
	}
	scraper.closed = true
	close(scraper.shutdownDone)
}

func (scraper *KyberScraper) Close() error {
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

type KyberPairScraper struct {
	parent *KyberScraper
	pair   dia.ExchangePair
	closed bool
}

func (pairScraper *KyberPairScraper) Pair() dia.ExchangePair {
	return pairScraper.pair
}

func (scraper *KyberScraper) Channel() chan *dia.Trade {
	return scraper.chanTrades
}

func (pairScraper *KyberPairScraper) Error() error {
	s := pairScraper.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

func (pairScraper *KyberPairScraper) Close() error {
	pairScraper.parent.errorLock.RLock()
	defer pairScraper.parent.errorLock.RUnlock()
	pairScraper.closed = true
	return nil
}
