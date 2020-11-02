package scrapers

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/diadata-org/diadata/internal/pkg/exchange-scrapers/kyber"
	"github.com/diadata-org/diadata/internal/pkg/exchange-scrapers/kyber/token"

	"github.com/diadata-org/diadata/pkg/dia"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	kyberContract              = "0x9AAb3f75489902f3a48495025729a0AF77d4b11e"
	kyberStartBlock            = uint64(11062549 - 5250)
	kyberStartBlockToFindPairs = uint64(11062549 - 5250)

	kyberWsDial   = "ws://159.69.120.42:8546/"
	kyberRestDial = "http://159.69.120.42:8545/"
)

type KyberToken struct {
	Symbol   string
	Decimals uint8
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
	tokens      map[string]*KyberToken
}

func NewKyberScraper(exchangeName string) *KyberScraper {
	scraper := &KyberScraper{
		exchangeName:   exchangeName,
		initDone:       make(chan nothing),
		shutdown:       make(chan nothing),
		shutdownDone:   make(chan nothing),
		productPairIds: make(map[string]int),
		pairScrapers:   make(map[string]*KyberPairScraper),
		chanTrades:     make(chan *dia.Trade),
		resubscribe:    make(chan nothing),
		tokens:         make(map[string]*KyberToken),
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

	go scraper.mainLoop()
	return scraper
}

func (scraper *KyberScraper) loadTokens() {

	scraper.tokens["0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"] = &KyberToken{
		Symbol:   "ETH",
		Decimals: 18,
	}

	// added by hand because the symbol method returns a bytes32 instead of string
	scraper.tokens["0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2"] = &KyberToken{
		Symbol:   "MKR",
		Decimals: 18,
	}

	filterer, err := kyber.NewKyberFilterer(common.HexToAddress(kyberContract), scraper.WsClient)
	if err != nil {
		log.Error(err)

	}

	it, err := filterer.FilterExecuteTrade(&bind.FilterOpts{Start: kyberStartBlockToFindPairs}, nil)
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

func (scraper *KyberScraper) loadTokenData(tokenAddress common.Address) (*KyberToken, error) {
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
		decimals, err := tokenCaller.Decimals(&bind.CallOpts{})
		if err != nil {
			log.Error(err)
		}
		dfToken := &KyberToken{
			Symbol:   symbol,
			Decimals: decimals,
		}
		dfToken.normalizeETH()
		scraper.tokens[tokenStr] = dfToken
		return dfToken, err
	}
}

func (scraper *KyberScraper) subscribeToTrades() error {

	filterer, err := kyber.NewKyberFilterer(common.HexToAddress(kyberContract), scraper.WsClient)
	if err != nil {
		log.Error(err)
		return err
	}
	start := startBlock
	sink := make(chan *kyber.KyberExecuteTrade)
	sub, err := filterer.WatchExecuteTrade(&bind.WatchOpts{Start: &start}, sink, nil)
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

func (scraper *KyberScraper) processTrade(trade *kyber.KyberExecuteTrade) {
	symbol, foreignName, volume, price, err := scraper.getTradeDataKyber(trade)
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

func (scraper *KyberScraper) mainLoop() {

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
			scraper.error = errors.New("Kyber: No pairs to scrape provided")
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

// getTradeData returns the foreign name, volume and price of a swap
func (scraper *KyberScraper) getTradeDataKyber(s *kyber.KyberExecuteTrade) (symbol string, foreignName string, volume float64, price float64, err error) {
	buyToken, err := scraper.loadTokenData(s.Dest)
	if err != nil {
		log.Error(err)
	}
	sellToken, err := scraper.loadTokenData(s.Src)
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

func (scraper *KyberScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {

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

func (t *KyberToken) normalizeETH() {
	if t.Symbol == "WETH" {
		t.Symbol = "ETH"
	}
}

func (scraper *KyberScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {
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
	pair   dia.Pair
	closed bool
}

func (pairScraper *KyberPairScraper) Pair() dia.Pair {
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
