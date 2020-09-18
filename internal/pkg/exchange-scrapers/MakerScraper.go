package scrapers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/utils"

	"sync"

	"github.com/diadata-org/diadata/pkg/dia"
)

const (
	MakerBatchDelay = 60 * 1
)

type MakerPairResponse struct {
	Data    map[string]MakerPair `json:"data"`
	Time    time.Time            `json:"time"`
	Message string               `json:"message"`
}

type MakerPair struct {
	Base           string `json:"base"`
	Quote          string `json:"quote"`
	BasePrecision  int    `json:"basePrecision"`
	QuotePrecision int    `json:"quotePrecision"`
	Active         bool   `json:"active"`
}

type MakerScraper struct {
	exchangeName string

	// channels to signal events
	run          bool
	initDone     chan nothing
	shutdown     chan nothing
	shutdownDone chan nothing

	errorLock sync.RWMutex
	error     error
	closed    bool

	pairScrapers   map[string]*MakerPairScraper
	productPairIds map[string]int
	chanTrades     chan *dia.Trade
}

type MakerTradeResponse struct {
	MakerTrades []MakerTrade `json:"data"`
	Time        time.Time    `json:"time"`
	Message     string       `json:"message"`
}

type MakerTrade struct {
	ID     int       `json:"id"`
	Price  string    `json:"price"`
	Volume string    `json:"volume"`
	Time   time.Time `json:"time"`
}

func NewMakerScraper(exchangeName string) *MakerScraper {
	scraper := &MakerScraper{
		exchangeName:   exchangeName,
		initDone:       make(chan nothing),
		shutdown:       make(chan nothing),
		shutdownDone:   make(chan nothing),
		productPairIds: make(map[string]int),
		pairScrapers:   make(map[string]*MakerPairScraper),
		chanTrades:     make(chan *dia.Trade),
	}

	go scraper.mainLoop()
	return scraper
}

func (scraper *MakerScraper) GetNewTrades(pair string, startTradeID string) ([]MakerTrade, error) {
	var (
		makertraderesponse MakerTradeResponse
		err                error
		bytes              []byte
		url                string
	)
	auxPair := strings.Split(pair, "-")
	pair = auxPair[0] + "/" + auxPair[1]
	if startTradeID == "" {
		url = "https://api.oasisdex.com/v2/trades/" + pair
	} else {
		tradeId, _ := strconv.Atoi(startTradeID)
		next := tradeId + 100
		url = "https://api.oasisdex.com/v2/trades/" + pair + "?limit=100?fromId+" + strconv.Itoa(next)
	}

	bytes, err = utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &makertraderesponse)
	return makertraderesponse.MakerTrades, nil
}

func (scraper *MakerScraper) mainLoop() {
	scraper.run = true
	startTradeID := make(map[string]string)

	for scraper.run {
		if len(scraper.pairScrapers) == 0 {
			scraper.error = errors.New("Maker: No pairs to scrape provided")
			log.Error(scraper.error.Error())
			break
		}

		for pair := range scraper.pairScrapers {
			trades, _ := scraper.GetNewTrades(pair, startTradeID[pair])
			if len(trades) > 0 {
				startTradeID[pair] = strconv.Itoa(trades[0].ID)

			}
			for _, v := range trades {

				price, err := strconv.ParseFloat(v.Price, 64)
				if err != nil {
					return
				}
				VolumeIn, err := strconv.ParseFloat(v.Volume, 64)
				if err != nil {
					return
				}

				trade := &dia.Trade{
					Symbol:         strings.Split(pair, "-")[1],
					Pair:           pair,
					Price:          price,
					Volume:         VolumeIn,
					Time:           v.Time,
					ForeignTradeID: strconv.Itoa(v.ID),
					Source:         scraper.exchangeName,
				}
				log.Infoln("Got Trade  ", trade)
				scraper.chanTrades <- trade

			}

		}
		time.Sleep(time.Duration(MakerBatchDelay) * time.Second)

	}

	if scraper.error == nil {
		scraper.error = errors.New("Main loop terminated by Close().")
	}
	scraper.cleanup(nil)
}

func (scraper *MakerScraper) getPairs() (pairs []dia.Pair, err error) {
	var response MakerPairResponse
	byte, err := utils.GetRequest("https://api.oasisdex.com/v2/pairs")
	err = json.Unmarshal(byte, &response)
	for i, v := range response.Data {
		pair := strings.Split(i, "/")
		pairs = append(pairs, dia.Pair{
			Symbol:      v.Base,
			ForeignName: pair[0] + "-" + pair[1],
			Exchange:    scraper.exchangeName,
			Ignore:      false,
		})

	}
	return
}

func (scraper *MakerScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {
	return scraper.getPairs()
}

func (scraper *MakerScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {
	scraper.errorLock.RLock()
	defer scraper.errorLock.RUnlock()

	if scraper.error != nil {
		return nil, scraper.error
	}

	if scraper.closed {
		return nil, errors.New("uniswapScraper is closed")
	}

	pairScraper := &MakerPairScraper{
		parent: scraper,
		pair:   pair,
	}

	scraper.pairScrapers[pair.ForeignName] = pairScraper

	return pairScraper, nil
}
func (scraper *MakerScraper) cleanup(err error) {
	scraper.errorLock.Lock()
	defer scraper.errorLock.Unlock()
	if err != nil {
		scraper.error = err
	}
	scraper.closed = true
	close(scraper.shutdownDone)
}

func (scraper *MakerScraper) Close() error {
	// close the pair scraper channels
	scraper.run = false
	for _, pairScraper := range scraper.pairScrapers {
		pairScraper.closed = true
	}

	close(scraper.shutdown)
	<-scraper.shutdownDone
	return nil
}

type MakerPairScraper struct {
	parent *MakerScraper
	pair   dia.Pair
	closed bool
}

func (pairScraper *MakerPairScraper) Pair() dia.Pair {
	return pairScraper.pair
}

func (scraper *MakerScraper) Channel() chan *dia.Trade {
	return scraper.chanTrades
}

func (pairScraper *MakerPairScraper) Error() error {
	s := pairScraper.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

func (pairScraper *MakerPairScraper) Close() error {
	pairScraper.parent.errorLock.RLock()
	defer pairScraper.parent.errorLock.RUnlock()
	pairScraper.closed = true
	return nil
}
