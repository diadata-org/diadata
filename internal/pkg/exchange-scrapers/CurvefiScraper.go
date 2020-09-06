package scrapers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
)

const (
	CurveBatchDelay = 60 * 1
)

type CurveTokensResponse struct {
	Data struct {
		Tokens CurveTokens `json:"tokens"`
	} `json:"data"`
}

type CurveTokens []struct {
	Decimals string `json:"decimals"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
}

const (
	CurveFIApiDelay = 60 * 60 * 24
)

type Token struct {
	Decimals string `json:"decimals"`
	ID       string `json:"id"`
	Symbol   string `json:"symbol"`
}

type SwapListCurve struct {
	Data struct {
		Swaps []SwapCurve `json:"swaps"`
	} `json:"data"`
}

type SwapCurve struct {
	Exchange struct {
		ID                           string `json:"id"`
		TotalUnderlyingVolumeDecimal string `"json:totalUnderlyingVolumeDecimal"`
	} `json:"exchange"`
	FromToken       Token
	FromTokenAmount string `json:"fromTokenAmount"`
	ToToken         Token
	ToTokenAmount   string `json:"toTokenAmount"`
	Transaction     struct {
		Hash string `json:"hash"`
		ID   string `json:"id"`
	} `json:"transaction"`
	Timestamp       string `json:"timestamp"`
	UnderlyingPrice string `json:"underlyingPrice"`
}

type CurveFIAssetPairs struct {
	Data struct {
		Pairs []CurveFIPair `json:"pairs"`
	} `json:"data"`
}

type CurveFIPair struct {
	CreatedAtTimestamp string `json:"createdAtTimestamp"`
	ID                 string `json:"id"`
	Token0             struct {
		Symbol string `json:"symbol"`
	} `json:"token0"`
	Token0Price string `json:"token0Price"`
	Token1      struct {
		Symbol string `json:"symbol"`
	} `json:"token1"`
	Token1Price string `json:"token1Price"`
	VolumeUSD   string `json:"volumeUSD"`
}

type CurveFIScraper struct {
	exchangeName string

	// channels to signal events
	run          bool
	initDone     chan nothing
	shutdown     chan nothing
	shutdownDone chan nothing

	errorLock sync.RWMutex
	error     error
	closed    bool

	pairScrapers   map[string]*CurveFIPairScraper
	productPairIds map[string]int
	chanTrades     chan *dia.Trade
}

func NewCurveFIScraper(exchangeName string) *CurveFIScraper {
	scraper := &CurveFIScraper{
		exchangeName:   exchangeName,
		initDone:       make(chan nothing),
		shutdown:       make(chan nothing),
		shutdownDone:   make(chan nothing),
		productPairIds: make(map[string]int),
		pairScrapers:   make(map[string]*CurveFIPairScraper),
		chanTrades:     make(chan *dia.Trade),
	}

	go scraper.mainLoop()
	return scraper
}

func (scraper *CurveFIScraper) mainLoop() {
	scraper.run = true

	// The scraping is organized in batches, bounded by time range and trade id
	starttime := time.Now().Add(time.Duration(-1000*CurveBatchDelay) * time.Second).Unix()
	startTradeID := ""

	for scraper.run {
		if len(scraper.pairScrapers) == 0 {
			scraper.error = errors.New("Curvefi: No pairs to scrape provided")
			log.Error(scraper.error.Error())
			break
		}

		// Determine the next swap batch
		swapsPre, lastTradeTime, lastTradeID, err := scraper.GetNewSwaps(starttime)

		fmt.Println("starttime: ", starttime)
		fmt.Println("last trade time: ", lastTradeTime)
		swaps := cutSwapListCurve(swapsPre, startTradeID)
		if err != nil {
			log.Error("error getting trades: ", err)
		}

		starttime = lastTradeTime - int64(1)
		startTradeID = lastTradeID

		swapsMap := make(map[string]SwapCurve)
		for _, s := range swaps.Data.Swaps {
			foreignName, _, _, _ := getSwapDataCurve(s)
			swapsMap[foreignName] = s
		}

		for pair, pairScraper := range scraper.pairScrapers {

			// Check, whether pair of swap is in list of available pairs
			swap, ok := swapsMap[pair]
			if !ok {
				continue
			}

			// Get trading data from swap in "classic" format
			_, volume, price, err := getSwapDataCurve(swap)
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
				ForeignTradeID: swap.Transaction.ID,
				Source:         scraper.exchangeName,
			}
			log.Infoln("Got Trade  ", trade)

			scraper.chanTrades <- trade
		}
		time.Sleep(time.Duration(CurveBatchDelay) * time.Second)

	}

	if scraper.error == nil {
		scraper.error = errors.New("Main loop terminated by Close().")
	}
	scraper.cleanup(nil)
}

// cutSwapList receives a list of swaps ordered by timestamps in descending order.
// It returns all newest swaps up to the one with id.
func cutSwapListCurve(swaps SwapListCurve, id string) (cutSwaps SwapListCurve) {
	if id == "" {
		return swaps
	}
	// Append new swaps until the swap with id is hit
	for _, swap := range swaps.Data.Swaps {
		if swap.Transaction.ID == id {
			return
		}
		cutSwaps.Data.Swaps = append(cutSwaps.Data.Swaps, swap)
	}
	return
}

// getSwapDataCurve returns the foreign name, volume and price of a swap
func getSwapDataCurve(s SwapCurve) (foreignName string, volume float64, price float64, err error) {
	amountIn, err := strconv.ParseFloat(s.FromTokenAmount, 64)
	if err != nil {
		return
	}
	decimalsIn, err := strconv.ParseInt(s.FromToken.Decimals, 10, 64)
	if err != nil {
		log.Error("error parsing decimals of TokenIn: ", err)
		return
	}
	amountIn /= math.Pow10(int(decimalsIn))

	amountOut, err := strconv.ParseFloat(s.ToTokenAmount, 64)
	if err != nil {
		return
	}
	decimalsOut, err := strconv.ParseInt(s.ToToken.Decimals, 10, 64)
	if err != nil {
		log.Error("error parsing decimals of TokenOut: ", err)
		return
	}
	amountOut /= math.Pow10(int(decimalsOut))

	volume = amountOut
	price = amountIn / amountOut
	foreignName = s.ToToken.Symbol + "-" + s.FromToken.Symbol

	return
}

// GetNewSwaps returns swaps with timestamps later than @starttime.
// Furthermore, last trade Time and ID are returned
func (scraper *CurveFIScraper) GetNewSwaps(starttime int64) (swapresponse SwapListCurve, lastTradeTime int64, lastTradeID string, err error) {

	// Get available assets
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
    			fromToken {
						id	
						symbol
						decimals
    			}
    			fromTokenAmount
    			toToken {
  				    id
  				    symbol
    				decimals
    			}
				toTokenAmount
				exchange{
					id
					totalUnderlyingVolumeDecimal
			  	}
				transaction {
					id
					hash
				  }
    			underlyingPrice
   				timestamp
   				}
}
`, starttime, time.Now().Unix()),
	}

	jsonValue, _ := json.Marshal(jsonData)
	jsondata, err := utils.PostRequest("https://api.thegraph.com/subgraphs/name/blocklytics/curve", bytes.NewBuffer(jsonValue))
	err = json.Unmarshal(jsondata, &swapresponse)
	if len(swapresponse.Data.Swaps) == 0 {
		return
	}
	for i := 0; i < len(swapresponse.Data.Swaps); i++ {
		swapresponse.Data.Swaps[i].FromToken.normalizeETH()
		swapresponse.Data.Swaps[i].ToToken.normalizeETH()
	}
	lastTradeTime, err = strconv.ParseInt(swapresponse.Data.Swaps[0].Timestamp, 10, 64)
	if err != nil {
		log.Error("error parsing unix time: ", err)
		return
	}
	lastTradeID = swapresponse.Data.Swaps[0].Transaction.ID

	return

}

func (t *Token) normalizeETH() {
	if t.Symbol == "WETH" {
		t.Symbol = "ETH"
	}
}

func (scraper *CurveFIScraper) readCurvefiTokens() (pair CurveTokens, err error) {

	var curveTokenResponse CurveTokensResponse
	//we dont have list of pairs, to get poairs we will get all aavailable assets and create pairs
	// Get available assets
	jsonData := map[string]string{
		"query": `
       {
  tokens(first: 25) {
    id
    name
    symbol
    decimals
  }
}
`,
	}

	jsonValue, _ := json.Marshal(jsonData)
	jsondata, err := utils.PostRequest("https://api.thegraph.com/subgraphs/name/blocklytics/curve", bytes.NewBuffer(jsonValue))
	err = json.Unmarshal(jsondata, &curveTokenResponse)
	// normalize with respect to ETH/WETH

	return curveTokenResponse.Data.Tokens, nil

}

func (scraper *CurveFIScraper) createPairs() (pairs map[string]string) {
	pairs = make(map[string]string)
	tokens, _ := scraper.readCurvefiTokens()
	for _, token1 := range tokens {
		for _, token2 := range tokens {
			pair := token1.Symbol + "-" + token2.Symbol
			pairs[pair] = token1.Symbol + "-" + token2.Symbol
		}
	}
	return
}

func (scraper *CurveFIScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {
	pairmap := scraper.createPairs()
	for _, v := range pairmap {
		symbols := strings.Split(v, "-")
		pairs = append(pairs, dia.Pair{
			Symbol:      symbols[0],
			ForeignName: v,
			Exchange:    scraper.exchangeName,
		})

	}
	return
}

func (scraper *CurveFIScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {
	scraper.errorLock.RLock()
	defer scraper.errorLock.RUnlock()

	if scraper.error != nil {
		return nil, scraper.error
	}

	if scraper.closed {
		return nil, errors.New("uniswapScraper is closed")
	}

	pairScraper := &CurveFIPairScraper{
		parent: scraper,
		pair:   pair,
	}

	scraper.pairScrapers[pair.ForeignName] = pairScraper

	return pairScraper, nil
}
func (scraper *CurveFIScraper) cleanup(err error) {
	scraper.errorLock.Lock()
	defer scraper.errorLock.Unlock()
	if err != nil {
		scraper.error = err
	}
	scraper.closed = true
	close(scraper.shutdownDone)
}

func (scraper *CurveFIScraper) Close() error {
	// close the pair scraper channels
	scraper.run = false
	for _, pairScraper := range scraper.pairScrapers {
		pairScraper.closed = true
	}

	close(scraper.shutdown)
	<-scraper.shutdownDone
	return nil
}

type CurveFIPairScraper struct {
	parent *CurveFIScraper
	pair   dia.Pair
	closed bool
}

func (pairScraper *CurveFIPairScraper) Pair() dia.Pair {
	return pairScraper.pair
}

func (scraper *CurveFIScraper) Channel() chan *dia.Trade {
	return scraper.chanTrades
}

func (pairScraper *CurveFIPairScraper) Error() error {
	s := pairScraper.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

func (pairScraper *CurveFIPairScraper) Close() error {
	pairScraper.parent.errorLock.RLock()
	defer pairScraper.parent.errorLock.RUnlock()
	pairScraper.closed = true
	return nil
}
