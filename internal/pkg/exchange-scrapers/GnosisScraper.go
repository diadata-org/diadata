package scrapers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
)

const (
	GnosisApiDelay   = 20
	GnosisBatchDelay = 60 * 1
)

type GnosisPair struct {
	SellToken GnosisToken `json:"sellToken"`
	BuyToken  GnosisToken `json:"buyToken"`
}

type GnosisTokenList struct {
	Data struct {
		Tokens []GnosisToken `json:"tokens"`
	} `json:"data"`
}

type GnosisToken struct {
	Symbol   string `json:"symbol"`
	Decimals string `json:"decimals"`
}

type GnosisSwapList struct {
	Data struct {
		Swaps []GnosisSwap `json:"trades"`
	} `json:"data"`
}

type GnosisSwap struct {
	Timestamp  string `json:"createEpoch"`
	BuyToken   GnosisToken
	SellToken  GnosisToken
	SellVolume string `json:"sellVolume"`
	BuyVolume  string `json:"buyVolume"`
	ID         string `json:"id"`
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
}

func NewGnosisScraper(exchangeName string) *GnosisScraper {
	scraper := &GnosisScraper{
		exchangeName:   exchangeName,
		initDone:       make(chan nothing),
		shutdown:       make(chan nothing),
		shutdownDone:   make(chan nothing),
		productPairIds: make(map[string]int),
		pairScrapers:   make(map[string]*GnosisPairScraper),
		chanTrades:     make(chan *dia.Trade),
	}

	go scraper.mainLoop()
	return scraper
}

func (scraper *GnosisScraper) mainLoop() {

	scraper.run = true

	// The scraping is organized in batches, bounded by time range and trade id
	starttime := time.Now().Add(time.Duration(-60*60*24) * time.Second).Unix()
	startTradeID := ""

	for scraper.run {
		if len(scraper.pairScrapers) == 0 {
			scraper.error = errors.New("Gnosis: No pairs to scrape provided")
			log.Error(scraper.error.Error())
			break
		}

		// Determine the next swap batch
		swapsPre, lastTradeTime, lastTradeID, err := scraper.GetNewSwaps(starttime, time.Now().Unix())
		if err != nil {
			log.Error("error getting trades: ", err)
		}
		swaps := cutSwapsGnosis(swapsPre, startTradeID)
		starttime = lastTradeTime - int64(1)
		startTradeID = lastTradeID

		swapsMap := make(map[string]GnosisSwap)
		for _, s := range swaps {
			foreignName, _, _, _ := getSwapDataGnosis(s)
			swapsMap[foreignName] = s
		}

		for pair, pairScraper := range scraper.pairScrapers {

			// Check whether swap has an admissible pair
			swap, ok := swapsMap[pair]
			if !ok {
				continue
			}

			// Get trading data from swap in "classic" format
			_, volume, price, err := getSwapDataGnosis(swap)
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
			fmt.Println("got trade: ", trade)
		}
		time.Sleep(time.Duration(GnosisBatchDelay) * time.Second)
	}
	if scraper.error == nil {
		scraper.error = errors.New("main loop terminated by Close()")
	}
	scraper.cleanup(nil)
}

// cutSwapList receives a list of swaps ordered by timestamps in descending order.
// It returns all newest swaps up to the one with id.
func cutSwapsGnosis(swaps []GnosisSwap, id string) (cutSwaps []GnosisSwap) {
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
func getSwapDataGnosis(s GnosisSwap) (foreignName string, volume float64, price float64, err error) {
	buyDecimals, err := strconv.Atoi(s.BuyToken.Decimals)
	if err != nil {
		log.Error("error parsing decimals: ", err)
	}
	sellDecimals, err := strconv.Atoi(s.SellToken.Decimals)
	if err != nil {
		log.Error("error parsing decimals: ", err)
	}
	amountOut, err := strconv.ParseFloat(s.BuyVolume, 64)
	if err != nil {
		return
	}
	amountIn, err := strconv.ParseFloat(s.SellVolume, 64)
	if err != nil {
		return
	}
	amountOut /= math.Pow10(buyDecimals)
	amountIn /= math.Pow10(sellDecimals)
	volume = amountOut
	price = amountIn / amountOut
	foreignName = s.BuyToken.Symbol + "-" + s.SellToken.Symbol
	return
}

// GetNewSwaps returns all swaps with timestamps ranging from @starttime until now
func (scraper *GnosisScraper) GetNewSwaps(starttime, endtime int64) (swaps []GnosisSwap, lastTradeTime int64, lastTradeID string, err error) {
	swaplist := GnosisSwapList{}
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{
			trades(orderBy:createEpoch, orderDirection:desc,
						where: {createEpoch_gt:%d, createEpoch_lt:%d})
			{
			  sellToken {
				symbol
				decimals
			  }
			  buyToken{
				symbol
				decimals
			  }
			  sellVolume
			  buyVolume
			  createEpoch
			  id
			}
		  }		  
			`, starttime, endtime),
	}
	jsonValue, _ := json.Marshal(jsonData)
	jsondata, err := utils.PostRequest("https://api.thegraph.com/subgraphs/name/gnosis/protocol", bytes.NewBuffer(jsonValue))
	err = json.Unmarshal(jsondata, &swaplist)
	swaps = swaplist.Data.Swaps
	if len(swaps) == 0 {
		return
	}
	for i := 0; i < len(swaps); i++ {
		swaps[i].BuyToken.normalizeETH()
		swaps[i].SellToken.normalizeETH()
	}
	lastTradeTime, err = strconv.ParseInt(swaps[0].Timestamp, 10, 64)
	if err != nil {
		log.Error("error parsing unix time: ", err)
		return
	}
	lastTradeID = swaps[0].ID
	return
}

// FetchAvailablePairs fetches pairs by getting trading dates from the last 7 days and exctracting the pairs
// Unfortunately, pairs are not available on thegraph API and the available tokenlist is incomplete.
func (scraper *GnosisScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {

	endtime := time.Now().Unix()
	starttime := time.Now().Unix() - 60*60*24*14
	periodstamps := makePeriods(starttime, endtime)

	set := make(map[string]struct{})
	for i := 1; i < len(periodstamps); i++ {
		swaps, _, _, err := scraper.GetNewSwaps(periodstamps[i-1], periodstamps[i])
		if err != nil {
			log.Error("error fetching swaps: ", err)
		}
		for _, swap := range swaps {
			if swap.BuyToken.Symbol == "" || swap.SellToken.Symbol == "" {
				continue
			}
			swap.BuyToken.normalizeETH()
			swap.SellToken.normalizeETH()

			foreignName := swap.BuyToken.Symbol + "-" + swap.SellToken.Symbol
			if _, ok := set[foreignName]; !ok {
				pairs = append(pairs, dia.Pair{
					Symbol:      swap.BuyToken.Symbol,
					ForeignName: foreignName,
					Exchange:    scraper.exchangeName,
				})
				set[foreignName] = struct{}{}
			}
			foreignName = swap.SellToken.Symbol + "-" + swap.BuyToken.Symbol
			if _, ok := set[foreignName]; !ok {
				pairs = append(pairs, dia.Pair{
					Symbol:      swap.SellToken.Symbol,
					ForeignName: foreignName,
					Exchange:    scraper.exchangeName,
				})
				set[foreignName] = struct{}{}
			}
		}
	}
	return
}

// makePeriods returns unix timestamps between timeInit and timeFinal such that
// there is Period between two consecutive timestamps
func makePeriods(timeInit, timeFinal int64) (periodstamps []int64) {
	for timeInit < timeFinal {
		periodstamps = append(periodstamps, timeInit)
		auxTime := timeInit + 60*60*12
		timeInit = auxTime
	}
	return
}

func (t *GnosisToken) normalizeETH() {
	if t.Symbol == "WETH" {
		t.Symbol = "ETH"
	}
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
