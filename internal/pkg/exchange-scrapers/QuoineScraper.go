package scrapers

import (
	"encoding/json"
	"errors"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/jjjjpppp/quoinex-go-client/v2"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"io/ioutil"
	"strconv"
	"sync"
	"time"
)

const (
	PAIRS_JSON      = "../../../config/Quoine.json"
	API_DELAY       = 1*time.Second + 500*time.Millisecond
	EXECUTION_LIMIT = 200
)

type QuoineScraper struct {
	client       *quoinex.Client
	exchangeName string

	// channels to signal events
	initDone     chan nothing
	shutdown     chan nothing
	shutdownDone chan nothing

	errorLock sync.RWMutex
	error     error
	closed    bool

	pairScrapers   map[string]*QuoinePairScraper
	productPairIds map[string]int

	ticker *time.Ticker
}

func NewQuoineScraper(exchangeName string) *QuoineScraper {
	qClient, err := quoinex.NewClient("x", "x", nil)
	if err != nil {
		log.Error("Couldn't create Quoine client:", err)
		return nil
	}

	scraper := &QuoineScraper{
		client:         qClient,
		exchangeName:   exchangeName,
		initDone:       make(chan nothing),
		shutdown:       make(chan nothing),
		shutdownDone:   make(chan nothing),
		productPairIds: make(map[string]int),
		pairScrapers:   make(map[string]*QuoinePairScraper),
		ticker:         time.NewTicker(API_DELAY),
	}

	// generate JSON file with currency pairs
	err = scraper.createJSONPairs()
	if err != nil {
		log.Error("Couldn't obtain currency pairs:", err)
	}

	go scraper.mainLoop()
	return scraper
}

func (scraper *QuoineScraper) mainLoop() {
	for {
		for pair, id := range scraper.productPairIds {
			ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

			executions, err := scraper.client.GetExecutions(ctx, id, EXECUTION_LIMIT, 1)
			if err != nil {
				log.Error("Couldn't get executions:", err)
				continue
			}

			if len(executions.Models) <= 0 {
				continue
			}

			price, err := strconv.ParseFloat(executions.Models[0].Price, 64)
			if err != nil {
				log.Error("Price isn't parseable float")
				continue
			}

			volume, err := strconv.ParseFloat(executions.Models[0].Quantity, 64)
			if err != nil {
				log.Error("Volume isn't parseable float")
				continue
			}

			pairScraper := scraper.pairScrapers[pair]
			trade := &dia.Trade{
				Symbol:         pairScraper.pair.Symbol,
				Pair:           pair,
				Price:          price,
				Volume:         volume,
				Time:           time.Unix(int64(executions.Models[0].CreatedAt), 0),
				ForeignTradeID: strconv.Itoa(id),
				Source:         dia.QuoineExchange,
			}

			scraper.pairScrapers[pair].chanTrades <- trade
			time.Sleep(API_DELAY)
		}
	}
}

type CoinPairs struct {
	Coins []Pair
}

type Pair struct {
	Symbol      string
	ForeignName string
	Exchange    string
}

func (scraper *QuoineScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	products, err := scraper.client.GetProducts(ctx)
	if err != nil {
		return
	}

	pairs = make([]dia.Pair, len(products))

	for i, prod := range products {
		pairs[i].ForeignName = prod.CurrencyPairCode
		pairs[i].Symbol = prod.BaseCurrency
		pairs[i].Exchange = "Quoine"
	}

	return
}

func (scraper *QuoineScraper) createJSONPairs() error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	products, err := scraper.client.GetProducts(ctx)
	if err != nil {
		return err
	}

	pairs := CoinPairs{
		Coins: make([]Pair, len(products)),
	}

	for i, prod := range products {
		pair := Pair{
			Symbol:      prod.BaseCurrency,
			ForeignName: prod.CurrencyPairCode,
			Exchange:    "Quoine",
		}

		pairs.Coins[i] = pair

		// create a pair -> id mapping
		intId, err := strconv.Atoi(prod.ID)
		if err != nil {
			continue
		}
		scraper.productPairIds[prod.CurrencyPairCode] = intId
	}

	// create a JSON file with currency pairs
	jsonCoins, err := json.MarshalIndent(pairs, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(PAIRS_JSON, jsonCoins, 0644)
	return nil
}

func (scraper *QuoineScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {
	scraper.errorLock.RLock()
	defer scraper.errorLock.RUnlock()

	if scraper.error != nil {
		return nil, scraper.error
	}

	if scraper.closed {
		return nil, errors.New("Quoine scraper is closed")
	}

	pairScraper := &QuoinePairScraper{
		parent:     scraper,
		pair:       pair,
		chanTrades: make(chan *dia.Trade),
	}

	scraper.pairScrapers[pair.ForeignName] = pairScraper

	return pairScraper, nil
}

func (scraper *QuoineScraper) Close() error {
	scraper.errorLock.Lock()
	defer scraper.errorLock.Unlock()

	// close the pair scraper channels
	for _, pairScraper := range scraper.pairScrapers {
		close(pairScraper.chanTrades)
		pairScraper.closed = true
	}

	close(scraper.shutdown)
	<-scraper.shutdownDone
	scraper.closed = true
	return nil
}

type QuoinePairScraper struct {
	parent     *QuoineScraper
	pair       dia.Pair
	chanTrades chan *dia.Trade
	closed     bool
}

func (pairScraper *QuoinePairScraper) Pair() dia.Pair {
	return pairScraper.pair
}

func (pairScraper *QuoinePairScraper) Channel() chan *dia.Trade {
	return pairScraper.chanTrades
}

func (pairScraper *QuoinePairScraper) Error() error {
	s := pairScraper.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

func (pairScraper *QuoinePairScraper) Close() error {
	pairScraper.parent.errorLock.RLock()
	defer pairScraper.parent.errorLock.RUnlock()
	close(pairScraper.chanTrades)
	pairScraper.closed = true
	return nil
}
