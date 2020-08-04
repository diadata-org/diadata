package scrapers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	"github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
)

const (
	UniswapApiDelay   = 20
	UniswapBatchDelay = 60 * 1
)

type UniSwapPairs struct {
	Data struct {
		Pairs []UniSwapPair `json:"pairs"`
	} `json:"data"`
}

type UniSwapPair struct {
	Token0 struct {
		Symbol string `json:"symbol"`
	} `json:"token0"`
	Token1 struct {
		Symbol string `json:"symbol"`
	} `json:"token1"`
}

type SwapList struct {
	Data struct {
		Swaps []Swap `json:"swaps"`
	} `json:"data"`
}

type Swap struct {
	ID         string `json:"id"`
	Timestamp  string `json:"timestamp"`
	Pair       UniSwapPair
	Amount0In  string `json:"amount0In"`
	Amount0Out string `json:"amount0Out"`
	Amount1In  string `json:"amount1In"`
	Amount1Out string `json:"amount1Out"`
}

type UniswapScraper struct {
	exchangeName string

	// channels to signal events
	run          bool
	initDone     chan nothing
	shutdown     chan nothing
	shutdownDone chan nothing

	errorLock sync.RWMutex
	error     error
	closed    bool

	pairScrapers   map[string]*UniswapPairScraper
	productPairIds map[string]int
	chanTrades     chan *dia.Trade
}

func NewUniswapScraper(exchangeName string) *UniswapScraper {
	scraper := &UniswapScraper{
		exchangeName:   exchangeName,
		initDone:       make(chan nothing),
		shutdown:       make(chan nothing),
		shutdownDone:   make(chan nothing),
		productPairIds: make(map[string]int),
		pairScrapers:   make(map[string]*UniswapPairScraper),
		chanTrades:     make(chan *dia.Trade),
	}

	go scraper.mainLoop()
	return scraper
}

func (scraper *UniswapScraper) mainLoop() {

	scraper.run = true

	// The scraping is organized in batches, bounded by time range and trade id
	starttime := time.Now().Add(time.Duration(-UniswapBatchDelay) * time.Second).Unix()
	startTradeID := ""

	for scraper.run {
		if len(scraper.pairScrapers) == 0 {
			scraper.error = errors.New("Uniswap: No pairs to scrape provided")
			log.Error(scraper.error.Error())
			break
		}

		// Determine the next swap batch
		swapsPre, lastTradeTime, lastTradeID, err := scraper.GetNewSwaps(starttime)
		if err != nil {
			log.Error("error getting trades: ", err)
		}
		swaps := cutSwapList(swapsPre, startTradeID)
		starttime = lastTradeTime - int64(1)
		startTradeID = lastTradeID

		swapsMap := make(map[string]Swap)
		for _, s := range swaps.Data.Swaps {
			foreignName, _, _, _ := getSwapData(s)
			swapsMap[foreignName] = s
		}

		for pair, pairScraper := range scraper.pairScrapers {

			// Check, whether pair of swap is in list of available pairs
			swap, ok := swapsMap[pair]
			if !ok {
				continue
			}

			// Get trading data from swap in "classic" format
			_, volume, price, err := getSwapData(swap)
			timestamp, err := strconv.ParseInt(swap.Timestamp, 10, 64)
			if err != nil {
				log.Error("error parsing time: ", err)
			}

			trade := &dia.Trade{
				Symbol:         pairScraper.pair.Symbol,
				Pair:           pair,
				Price:          price,
				Volume:         volume,
				Time:           time.Unix(timestamp, 0),
				ForeignTradeID: "",
				Source:         scraper.exchangeName,
			}
			pairScraper.parent.chanTrades <- trade
		}
		time.Sleep(time.Duration(UniswapBatchDelay) * time.Second)
	}
	if scraper.error == nil {
		scraper.error = errors.New("main loop terminated by Close()")
	}
	scraper.cleanup(nil)
}

// cutSwapList receives a list of swaps ordered by timestamps in descending order.
// It returns all newest swaps up to the one with id.
func cutSwapList(swaps SwapList, id string) (cutSwaps SwapList) {
	if id == "" {
		return swaps
	}
	// Append new swaps until the swap with id is hit
	for _, swap := range swaps.Data.Swaps {
		if swap.ID == id {
			return
		}
		cutSwaps.Data.Swaps = append(cutSwaps.Data.Swaps, swap)
	}
	return
}

// getSwapData returns the foreign name, volume and price of a swap
func getSwapData(s Swap) (foreignName string, volume float64, price float64, err error) {
	amount0In, err := strconv.ParseFloat(s.Amount0In, 64)
	if err != nil {
		return
	}
	amount1In, err := strconv.ParseFloat(s.Amount1In, 64)
	if err != nil {
		return
	}
	amount0Out, err := strconv.ParseFloat(s.Amount0Out, 64)
	if err != nil {
		return
	}
	amount1Out, err := strconv.ParseFloat(s.Amount1Out, 64)
	if err != nil {
		return
	}
	if amount0In == float64(0) {
		volume = amount0Out
		price = amount1In / amount0Out
		foreignName = s.Pair.Token0.Symbol + "-" + s.Pair.Token1.Symbol
		return
	}
	volume = amount1Out
	price = amount0In / amount1Out
	foreignName = s.Pair.Token1.Symbol + "-" + s.Pair.Token0.Symbol
	return
}

// GetNewSwaps returns all swaps with timestamps ranging from @starttime until now
func (scraper *UniswapScraper) GetNewSwaps(starttime int64) (swaps SwapList, lastTradeTime int64, lastTradeID string, err error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
			{
				swaps(
					orderBy:timestamp, orderDirection:desc
					where: {
						timestamp_gt:%d,
						timestamp_lt: %d
					})
				{
				id
				pair{
					token0 {
						symbol
					}
					token1 {
						symbol
					}
				}
				amount0In
				amount1In
				amount0Out
				amount1Out
				timestamp

				}
			}
		
			`, starttime, time.Now().Unix()),
	}
	jsonValue, _ := json.Marshal(jsonData)
	jsondata, err := utils.PostRequest("https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2", bytes.NewBuffer(jsonValue))
	err = json.Unmarshal(jsondata, &swaps)
	if len(swaps.Data.Swaps) == 0 {
		return
	}
	for i := 0; i < len(swaps.Data.Swaps); i++ {
		swaps.Data.Swaps[i].Pair.normalizeETH()
	}
	lastTradeTime, err = strconv.ParseInt(swaps.Data.Swaps[0].Timestamp, 10, 64)
	if err != nil {
		log.Error("error parsing unix time: ", err)
		return
	}
	lastTradeID = swaps.Data.Swaps[0].ID
	return
}

// FetchAvailablePairs is the method used by the pairDiscoveryService, returning all available swap pairs
func (scraper *UniswapScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {
	assets, err := scraper.readAssets()
	if err != nil {
		log.Error("Couldn't obtain Uniswap product ids:", err)
		return
	}

	set := make(map[string]struct{})
	for _, v := range assets.Data.Pairs {
		if !helpers.SymbolIsBlackListed(v.Token0.Symbol) && !helpers.SymbolIsBlackListed(v.Token1.Symbol) {
			pairs = append(pairs, dia.Pair{
				Symbol:      v.Token1.Symbol,
				ForeignName: v.Token1.Symbol + "-" + v.Token0.Symbol,
				Exchange:    scraper.exchangeName,
			})
			set[v.Token1.Symbol+"-"+v.Token0.Symbol] = struct{}{}
		}
	}
	// Add swapped pairs
	for key := range set {
		symbols := strings.Split(key, "-")
		if _, ok := set[symbols[1]+"-"+symbols[0]]; !ok {
			pairs = append(pairs, dia.Pair{
				Symbol:      symbols[1],
				ForeignName: symbols[1] + "-" + symbols[0],
				Exchange:    scraper.exchangeName,
			})
		}
	}
	return
}

// This method is not used any more
func (scraper *UniswapScraper) fetchAllTokens() (tokens []string, err error) {
	assets, err := scraper.readAssets()
	if err != nil {
		log.Error("Couldn't obtain Uniswap product ids:", err)
		return
	}

	set := make(map[string]struct{})
	for _, v := range assets.Data.Pairs {
		if !helpers.SymbolIsBlackListed(v.Token0.Symbol) && !helpers.SymbolIsBlackListed(v.Token1.Symbol) {
			// Add token symbol to map if not present yet
			if _, ok := set[v.Token0.Symbol]; !ok {
				set[v.Token0.Symbol] = struct{}{}
			}
			if _, ok := set[v.Token1.Symbol]; !ok {
				set[v.Token1.Symbol] = struct{}{}
			}
		}
	}
	for key := range set {
		tokens = append(tokens, key)
	}
	return
}

// readAssets returns a list of available pairs.
func (scraper *UniswapScraper) readAssets() (pair UniSwapPairs, err error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
         {
  pairs(first: %d, orderBy:volumeUSD, orderDirection:desc) {
    token0{
      symbol
    }
    token1{
      symbol
    }
    volumeUSD
  }
}
`, 1000),
	}

	jsonValue, _ := json.Marshal(jsonData)
	jsondata, err := utils.PostRequest("https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2", bytes.NewBuffer(jsonValue))
	err = json.Unmarshal(jsondata, &pair)

	// substitue WETH with ETH as they can be identified
	for i := 0; i < len(pair.Data.Pairs); i++ {
		pair.Data.Pairs[i].normalizeETH()
	}
	return
}

func (p *UniSwapPair) normalizeETH() {
	if p.Token0.Symbol == "WETH" {
		p.Token0.Symbol = "ETH"
	}
	if p.Token1.Symbol == "WETH" {
		p.Token1.Symbol = "ETH"
	}
}

func (scraper *UniswapScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {
	scraper.errorLock.RLock()
	defer scraper.errorLock.RUnlock()

	if scraper.error != nil {
		return nil, scraper.error
	}

	if scraper.closed {
		return nil, errors.New("uniswapScraper is closed")
	}

	pairScraper := &UniswapPairScraper{
		parent: scraper,
		pair:   pair,
	}

	scraper.pairScrapers[pair.ForeignName] = pairScraper

	return pairScraper, nil
}
func (scraper *UniswapScraper) cleanup(err error) {
	scraper.errorLock.Lock()
	defer scraper.errorLock.Unlock()
	if err != nil {
		scraper.error = err
	}
	scraper.closed = true
	close(scraper.shutdownDone)
}

func (scraper *UniswapScraper) Close() error {
	// close the pair scraper channels
	scraper.run = false
	for _, pairScraper := range scraper.pairScrapers {
		pairScraper.closed = true
	}

	close(scraper.shutdown)
	<-scraper.shutdownDone
	return nil
}

type UniswapPairScraper struct {
	parent *UniswapScraper
	pair   dia.Pair
	closed bool
}

func (pairScraper *UniswapPairScraper) Pair() dia.Pair {
	return pairScraper.pair
}

func (scraper *UniswapScraper) Channel() chan *dia.Trade {
	return scraper.chanTrades
}

func (pairScraper *UniswapPairScraper) Error() error {
	s := pairScraper.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

func (pairScraper *UniswapPairScraper) Close() error {
	pairScraper.parent.errorLock.RLock()
	defer pairScraper.parent.errorLock.RUnlock()
	pairScraper.closed = true
	return nil
}
