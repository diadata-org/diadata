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

const (
	UniswapApiDelay = 60 * 60
)

type UniSwapTicker struct {
	Data struct {
		Name            string  `json:"name"`
		Symbol          string  `json:"symbol"`
		Code            string  `json:"code"`
		Decimals        int     `json:"decimals"`
		TotalSupply     string  `json:"totalSupply"`
		TotalSupplyD    string  `json:"totalSupplyD"`
		Volume24H       string  `json:"volume24h"`
		Volume24HD      float64 `json:"volume24hD"`
		DisplayCurrency string  `json:"displayCurrency"`
		Price24H        string  `json:"price24h"`
		Price           float64 `json:"price"`
	} `json:"data"`
}

type UniSwapAssetPairs struct {
	Data struct {
		Pairs []UniSwapPair `json:"pairs"`
	} `json:"data"`
}

type UniSwapPair struct {
	CreatedAtTimestamp string `json:"createdAtTimestamp"`
	ID                 string `json:"id"`
	Token0             struct {
		Symbol string `json:"symbol"`
	} `json:"token0"`
	Token0Price string `json:"token0Price"`
	Token1      struct {
		Symbol string `json:"symbol"`
	} `json:"token1"`
	VolumeUSD string `json:"volumeUSD"`
}

type UniSwapScraper struct {
	exchangeName string

	// channels to signal events
	run          bool
	initDone     chan nothing
	shutdown     chan nothing
	shutdownDone chan nothing

	errorLock sync.RWMutex
	error     error
	closed    bool

	pairScrapers   map[string]*UniSwapPairScraper
	productPairIds map[string]int
	chanTrades     chan *dia.Trade
}

func NewUniSwapScraper(exchangeName string) *UniSwapScraper {
	scraper := &UniSwapScraper{
		exchangeName:   exchangeName,
		initDone:       make(chan nothing),
		shutdown:       make(chan nothing),
		shutdownDone:   make(chan nothing),
		productPairIds: make(map[string]int),
		pairScrapers:   make(map[string]*UniSwapPairScraper),
		chanTrades:     make(chan *dia.Trade),
	}

	go scraper.mainLoop()
	return scraper
}

func (scraper *UniSwapScraper) mainLoop() {
	scraper.run = true
	assetMap := make(map[string]UniSwapPair)
	assets, err := scraper.readAssets()
	if err != nil {
		log.Error("Couldn't obtain Uniswap product ids:", err)
	}

	for _, v := range assets.Data.Pairs {
		assetMap[v.Token0.Symbol+"-"+v.Token1.Symbol] = v
	}

	for scraper.run {
		if len(scraper.pairScrapers) == 0 {
			scraper.error = errors.New("Uniswap: No pairs to scrap provided")
			log.Error(scraper.error.Error())
			break
		}
		for pair, pairScraper := range scraper.pairScrapers {
			// Sleep duration such that each pair is scraped once per hour
			time.Sleep(time.Duration(UniswapApiDelay/len(scraper.pairScrapers)) * time.Second)

			log.Println(pairScraper.pair.ForeignName)
			ticker := assetMap[pairScraper.pair.ForeignName]

			volume, err := strconv.ParseFloat(ticker.VolumeUSD, 64)
			if err != nil {
				log.Errorln("Volume isn't parseable float", ticker.VolumeUSD)
				continue
			}

			price, err := strconv.ParseFloat(ticker.Token0Price, 64)
			if err != nil {
				log.Error("Token0Price isn't parseable float")
				continue
			}

			trade := &dia.Trade{
				Symbol:         pairScraper.pair.Symbol,
				Pair:           pair,
				Price:          price,
				Volume:         volume,
				Time:           time.Now(),
				ForeignTradeID: "",
				Source:         scraper.exchangeName,
			}
			log.Infoln("Got Trade  ", trade)

			pairScraper.parent.chanTrades <- trade
		}
	}
	if scraper.error == nil {
		scraper.error = errors.New("Main loop terminated by Close().")
	}
	scraper.cleanup(nil)
}

func (scraper *UniSwapScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {
	assets, err := scraper.readAssets()
	if err != nil {
		log.Error("Couldn't obtain Uniswap product ids:", err)
	}

	for _, v := range assets.Data.Pairs {
		pairs = append(pairs, dia.Pair{
			Symbol:      v.Token0.Symbol,
			ForeignName: v.Token0.Symbol + "-" + v.Token1.Symbol,
			Exchange:    scraper.exchangeName,
		})
	}
	return
}

func (scraper *UniSwapScraper) readAssets() (pair UniSwapAssetPairs, err error) {
	jsonData := map[string]string{
		"query": `
         {
  pairs(first: 1000) {
    id
    token0{
      symbol
    }
    token1{
      symbol
    }
    token0Price
    volumeUSD
    createdAtTimestamp
 
  }
 
}
`,
	}

	jsonValue, _ := json.Marshal(jsonData)
	jsondata, err := utils.PostRequest("https://api.thegraph.com/subgraphs/name/ianlapham/uniswapv2", bytes.NewBuffer(jsonValue))
	err = json.Unmarshal(jsondata, &pair)

	return

}

func (scraper *UniSwapScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {
	scraper.errorLock.RLock()
	defer scraper.errorLock.RUnlock()

	if scraper.error != nil {
		return nil, scraper.error
	}

	if scraper.closed {
		return nil, errors.New("uniswapScraper is closed")
	}

	pairScraper := &UniSwapPairScraper{
		parent: scraper,
		pair:   pair,
	}

	scraper.pairScrapers[pair.ForeignName] = pairScraper

	return pairScraper, nil
}
func (scraper *UniSwapScraper) cleanup(err error) {
	scraper.errorLock.Lock()
	defer scraper.errorLock.Unlock()
	if err != nil {
		scraper.error = err
	}
	scraper.closed = true
	close(scraper.shutdownDone)
}

func (scraper *UniSwapScraper) Close() error {
	// close the pair scraper channels
	scraper.run = false
	for _, pairScraper := range scraper.pairScrapers {
		pairScraper.closed = true
	}

	close(scraper.shutdown)
	<-scraper.shutdownDone
	return nil
}

type UniSwapPairScraper struct {
	parent *UniSwapScraper
	pair   dia.Pair
	closed bool
}

func (pairScraper *UniSwapPairScraper) Pair() dia.Pair {
	return pairScraper.pair
}

func (scraper *UniSwapScraper) Channel() chan *dia.Trade {
	return scraper.chanTrades
}

func (pairScraper *UniSwapPairScraper) Error() error {
	s := pairScraper.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

func (pairScraper *UniSwapPairScraper) Close() error {
	pairScraper.parent.errorLock.RLock()
	defer pairScraper.parent.errorLock.RUnlock()
	pairScraper.closed = true
	return nil
}
