package scrapers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
)

const (
	BalancerApiDelay   = 20
	BalancerBatchDelay = 60 * 1
)

type BalancerSwapList struct {
	Data struct {
		Swaps []BalancerSwap `json:"swaps"`
	} `json:"data"`
}

type BalancerSwap struct {
	SellToken  string `json:"tokenInSym"`
	BuyToken   string `json:"tokenOutSym"`
	SellVolume string `json:"tokenAmountIn"`
	BuyVolume  string `json:"tokenAmountOut"`
	ID         string `json:"id"`
	Timestamp  int64  `json:"timestamp"`
}

type BalancerScraper struct {
	exchangeName string

	// channels to signal events
	run          bool
	initDone     chan nothing
	shutdown     chan nothing
	shutdownDone chan nothing

	errorLock sync.RWMutex
	error     error
	closed    bool

	pairScrapers   map[string]*BalancerPairScraper
	productPairIds map[string]int
	chanTrades     chan *dia.Trade
}

func NewBalancerScraper(exchangeName string) *BalancerScraper {
	scraper := &BalancerScraper{
		exchangeName:   exchangeName,
		initDone:       make(chan nothing),
		shutdown:       make(chan nothing),
		shutdownDone:   make(chan nothing),
		productPairIds: make(map[string]int),
		pairScrapers:   make(map[string]*BalancerPairScraper),
		chanTrades:     make(chan *dia.Trade),
	}

	go scraper.mainLoop()
	return scraper
}

func (scraper *BalancerScraper) mainLoop() {

	scraper.run = true

	// The scraping is organized in batches, bounded by time range and trade id
	starttime := time.Now().Add(time.Duration(-60*2) * time.Second).Unix()
	startTradeID := ""

	for scraper.run {
		if len(scraper.pairScrapers) == 0 {
			scraper.error = errors.New("Balancer: No pairs to scrape provided")
			log.Error(scraper.error.Error())
			break
		}

		// Determine the next swap batch
		swapsPre, lastTradeTime, lastTradeID, err := scraper.GetNewSwaps(starttime, time.Now().Unix())
		if err != nil {
			log.Error("error getting trades: ", err)
		}
		swaps := cutSwapsBalancer(swapsPre, startTradeID)
		starttime = lastTradeTime - int64(1)
		startTradeID = lastTradeID

		swapsMap := make(map[string]BalancerSwap)
		for _, s := range swaps {
			swapsMap[s.BuyToken+"-"+s.SellToken] = s
		}

		for pair, pairScraper := range scraper.pairScrapers {

			// Check whether swap has an admissible pair
			swap, ok := swapsMap[pair]
			if !ok {
				continue
			}

			// Get trading data from swap in "classic" format
			_, volume, price, err := getSwapDataBalancer(swap)
			timestamp := swap.Timestamp

			if err != nil {
				log.Error("error parsing time: ", err)
			}

			trade := &dia.Trade{
				Symbol:         pairScraper.pair.Symbol,
				Pair:           pair,
				Price:          price,
				Volume:         volume,
				Time:           time.Unix(timestamp, 0),
				ForeignTradeID: swap.ID,
				Source:         scraper.exchangeName,
			}
			pairScraper.parent.chanTrades <- trade
			fmt.Println("got trade: ", trade)
		}
		time.Sleep(time.Duration(BalancerBatchDelay) * time.Second)
	}
	if scraper.error == nil {
		scraper.error = errors.New("main loop terminated by Close()")
	}
	scraper.cleanup(nil)
}

// cutSwapList receives a list of swaps ordered by timestamps in descending order.
// It returns all newest swaps up to the one with id.
func cutSwapsBalancer(swaps []BalancerSwap, id string) (cutSwaps []BalancerSwap) {
	if id == "" {
		return swaps
	}
	// Append new swaps until the swap with id is hit
	for _, swap := range swaps {
		if swap.ID == id {
			return
		}
		cutSwaps = append(cutSwaps, swap)
	}
	return
}

// getSwapData returns the foreign name, volume and price of a swap
func getSwapDataBalancer(s BalancerSwap) (foreignName string, volume float64, price float64, err error) {

	amountOut, err := strconv.ParseFloat(s.BuyVolume, 64)
	if err != nil {
		return
	}
	amountIn, err := strconv.ParseFloat(s.SellVolume, 64)
	if err != nil {
		return
	}
	volume = amountOut
	price = amountIn / amountOut
	foreignName = s.BuyToken + "-" + s.SellToken
	return
}

// GetNewSwaps returns all swaps with timestamps ranging from @starttime until now
func (scraper *BalancerScraper) GetNewSwaps(starttime, endtime int64) (swaps []BalancerSwap, lastTradeTime int64, lastTradeID string, err error) {
	swaplist := BalancerSwapList{}
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{
			swaps(orderBy:timestamp, orderDirection:desc,
						where: {timestamp_gt:%d, timestamp_lt:%d})
			{
				tokenInSym
				tokenOutSym
				tokenAmountIn
				tokenAmountOut
				id
				timestamp
			}
		  }		  
			`, starttime, endtime),
	}
	jsonValue, _ := json.Marshal(jsonData)
	jsondata, err := utils.PostRequest("https://api.thegraph.com/subgraphs/name/balancer-labs/balancer", bytes.NewBuffer(jsonValue))
	err = json.Unmarshal(jsondata, &swaplist)
	swaps = swaplist.Data.Swaps

	if len(swaps) == 0 {
		return
	}
	for i := 0; i < len(swaps); i++ {
		swaps[i].normalizeETH()
	}

	lastTradeTime = swaps[0].Timestamp
	if err != nil {
		log.Error("error parsing unix time: ", err)
		return
	}
	lastTradeID = swaps[0].ID
	return
}

// FetchAvailablePairs fetches pairs by getting trading dates from the last 7 days and exctracting the pairs
// Unfortunately, pairs are not available on thegraph API and the available tokenlist is incomplete.
func (scraper *BalancerScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {

	endtime := time.Now().Unix()
	starttime := time.Now().Unix() - 60*60*24*14
	periodstamps := makeBalancerPeriods(starttime, endtime)

	set := make(map[string]struct{})
	for i := 1; i < len(periodstamps); i++ {
		swaps, _, _, err := scraper.GetNewSwaps(periodstamps[i-1], periodstamps[i])
		if err != nil {
			log.Error("error fetching swaps: ", err)
		}
		for _, swap := range swaps {
			if swap.BuyToken == "" || swap.SellToken == "" {
				continue
			}
			swap.normalizeETH()

			foreignName := swap.BuyToken + "-" + swap.SellToken
			if _, ok := set[foreignName]; !ok {
				pairs = append(pairs, dia.Pair{
					Symbol:      swap.BuyToken,
					ForeignName: foreignName,
					Exchange:    scraper.exchangeName,
					Ignore:      false,
				})
				set[foreignName] = struct{}{}
			}
			foreignName = swap.SellToken + "-" + swap.BuyToken
			if _, ok := set[foreignName]; !ok {
				pairs = append(pairs, dia.Pair{
					Symbol:      swap.SellToken,
					ForeignName: foreignName,
					Exchange:    scraper.exchangeName,
					Ignore:      false,
				})
				set[foreignName] = struct{}{}
			}
		}
	}
	return
}

// makeBalancerPeriods returns unix timestamps between timeInit and timeFinal such that
// there is Period between two consecutive timestamps
func makeBalancerPeriods(timeInit, timeFinal int64) (periodstamps []int64) {
	for timeInit < timeFinal {
		periodstamps = append(periodstamps, timeInit)
		auxTime := timeInit + 60*60*12
		timeInit = auxTime
	}
	return
}

func (bs BalancerSwap) normalizeETH() {
	if bs.SellToken == "WETH" {
		bs.SellToken = "ETH"
	}
	if bs.BuyToken == "WETH" {
		bs.BuyToken = "ETH"
	}

}

func (scraper *BalancerScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {
	scraper.errorLock.RLock()
	defer scraper.errorLock.RUnlock()

	if scraper.error != nil {
		return nil, scraper.error
	}

	if scraper.closed {
		return nil, errors.New("BalancerScraper is closed")
	}

	pairScraper := &BalancerPairScraper{
		parent: scraper,
		pair:   pair,
	}

	scraper.pairScrapers[pair.ForeignName] = pairScraper

	return pairScraper, nil
}
func (scraper *BalancerScraper) cleanup(err error) {
	scraper.errorLock.Lock()
	defer scraper.errorLock.Unlock()
	if err != nil {
		scraper.error = err
	}
	scraper.closed = true
	close(scraper.shutdownDone)
}

func (scraper *BalancerScraper) Close() error {
	// close the pair scraper channels
	scraper.run = false
	for _, pairScraper := range scraper.pairScrapers {
		pairScraper.closed = true
	}

	close(scraper.shutdown)
	<-scraper.shutdownDone
	return nil
}

type BalancerPairScraper struct {
	parent *BalancerScraper
	pair   dia.Pair
	closed bool
}

func (pairScraper *BalancerPairScraper) Pair() dia.Pair {
	return pairScraper.pair
}

func (scraper *BalancerScraper) Channel() chan *dia.Trade {
	return scraper.chanTrades
}

func (pairScraper *BalancerPairScraper) Error() error {
	s := pairScraper.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

func (pairScraper *BalancerPairScraper) Close() error {
	pairScraper.parent.errorLock.RLock()
	defer pairScraper.parent.errorLock.RUnlock()
	pairScraper.closed = true
	return nil
}
