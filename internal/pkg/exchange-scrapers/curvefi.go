package scrapers

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
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

type CurveFISwapsResponse struct {
	Data struct {
		Swaps []struct {
			Exchange struct {
				ID string `json:"id"`
				TotalUnderlyingVolumeDecimal string `"json:totalUnderlyingVolumeDecimal"`
			} `json:"exchange"`
			FromToken struct {
				Symbol string `json:"symbol"`
			} `json:"fromToken"`
			FromTokenAmount  string `json:"fromTokenAmount"`
			IsUnderlyingSwap bool   `json:"isUnderlyingSwap"`
			Timestamp        string `json:"timestamp"`
			ToToken          struct {
				Decimals string `json:"decimals"`
				ID       string `json:"id"`
				Symbol   string `json:"symbol"`
			} `json:"toToken"`
			ToTokenAmount string `json:"toTokenAmount"`
			Transaction   struct {
				Block    string `json:"block"`
				Hash     string `json:"hash"`
				LogIndex string `json:"logIndex"`
			} `json:"transaction"`
			UnderlyingPrice string `json:"underlyingPrice"`
			User            struct {
				ID string `json:"id"`
			} `json:"user"`
		} `json:"swaps"`
	} `json:"data"`
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
	for scraper.run {

		swapresponse, err := scraper.readSwaps()
		if err != nil {
			log.Error("Couldn't obtain curvefi swaps product ids:", err)
		}

		for _, trade := range swapresponse.Data.Swaps {
			price, err := strconv.ParseFloat(trade.UnderlyingPrice, 64)
			if err != nil {
				log.Error("UnderlyingPrice isn't parseable float")
				continue
			}

			volume, err := strconv.ParseFloat(trade.Exchange.TotalUnderlyingVolumeDecimal, 64)
			if err != nil {
				log.Error("UnderlyingVolume isn't parseable float")
				continue
			}
			timestamp, err := strconv.ParseInt(trade.Timestamp, 10, 64)
			if err != nil {
				logger.Println("Error Parsing time", err)
			}

			trade := &dia.Trade{
				Symbol:         trade.FromToken.Symbol+"-"+trade.ToToken.Symbol,
				Pair:           trade.FromToken.Symbol+"-"+trade.ToToken.Symbol,
				Price:          price,
				Volume:         volume,
				Time:           time.Unix(timestamp, 0),
				ForeignTradeID: trade.Transaction.Hash,
				Source:         scraper.exchangeName,
			}
			log.Infoln("Got Trade  ", trade)

			scraper.chanTrades  <- trade
		}
		time.Sleep(time.Duration(CurveFIApiDelay/len(scraper.pairScrapers)) * time.Second)

	}

	if scraper.error == nil {
		scraper.error = errors.New("Main loop terminated by Close().")
	}
	scraper.cleanup(nil)
}

func (scraper *CurveFIScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {
	pairmap := scraper.createPairs()
	for _, v := range pairmap {
		pairs = append(pairs, dia.Pair{
			Symbol:      v,
			ForeignName: v,
			Exchange:    scraper.exchangeName,
		})

	}
	return
}

func (scraper *CurveFIScraper) readSwaps() (pair CurveFISwapsResponse, err error) {

	var swapresponse CurveFISwapsResponse
	//we dont have list of pairs, to get poairs we will get all aavailable assets and create pairs
	// Get available assets
	jsonData := map[string]string{
		"query": `
      {
  swaps(first: 100, orderBy: timestamp) {
    isUnderlyingSwap
    fromToken {
      symbol
    }
    exchange{
      id
      totalUnderlyingVolumeDecimal
    }
    fromTokenAmount
    toToken {
      id
      symbol
      decimals
    }
    toTokenAmount
    underlyingPrice
    timestamp
    transaction {
      hash
      block
      logIndex
    }
    user {
      id
    }
  }
}
`,
	}

	jsonValue, _ := json.Marshal(jsonData)
	jsondata, err := utils.PostRequest("https://api.thegraph.com/subgraphs/name/blocklytics/curve", bytes.NewBuffer(jsonValue))
	err = json.Unmarshal(jsondata, &swapresponse)

	return swapresponse, nil

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
