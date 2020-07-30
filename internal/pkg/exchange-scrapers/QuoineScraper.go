package scrapers

import (
	"errors"
	"strconv"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	"github.com/jjjjpppp/quoinex-go-client/v2"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

const (
	API_DELAY       = 1*time.Second + 500*time.Millisecond
	EXECUTION_LIMIT = 200
)

type QuoineScraper struct {
	client       *quoinex.Client
	exchangeName string

	// channels to signal events
	run          bool
	initDone     chan nothing
	shutdown     chan nothing
	shutdownDone chan nothing

	errorLock sync.RWMutex
	error     error
	closed    bool

	pairScrapers   map[string]*QuoinePairScraper
	productPairIds map[string]int
	chanTrades     chan *dia.Trade
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
		chanTrades:     make(chan *dia.Trade),
	}

	err = scraper.readProductIds()
	if err != nil {
		log.Error("Couldn't obtain Quoine product ids:", err)
	}

	go scraper.mainLoop()
	return scraper
}

func (scraper *QuoineScraper) mainLoop() {
	scraper.run = true
	for scraper.run {
		if len(scraper.pairScrapers) == 0 {
			scraper.error = errors.New("QuoineScraper: No pairs to scrap provided")
			log.Error(scraper.error.Error())
			break
		}
		for pair, pairScraper := range scraper.pairScrapers {
			time.Sleep(API_DELAY)

			productId, present := scraper.productPairIds[pair]
			if !present {
				log.Error("Unknown product id for pair", pair)
			}

			ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

			executions, err := scraper.client.GetExecutions(ctx, productId, EXECUTION_LIMIT, 1)
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

			if executions.Models[0].TakerSide == "sell" {
				volume = -volume
			}

			trade := &dia.Trade{
				Symbol:         pairScraper.pair.Symbol,
				Pair:           pair,
				Price:          price,
				Volume:         volume,
				Time:           time.Unix(int64(executions.Models[0].CreatedAt), 0),
				ForeignTradeID: strconv.Itoa(int(executions.Models[0].ID)),
				Source:         scraper.exchangeName,
			}

			pairScraper.parent.chanTrades <- trade
		}
	}
	if scraper.error == nil {
		scraper.error = errors.New("Main loop terminated by Close()")
	}
	scraper.cleanup(nil)
}

func (scraper *QuoineScraper) normalizeSymbol(foreignName string, params ...string) (symbol string, err error) {
	symbol = params[0]
	if helpers.NameForSymbol(symbol) == symbol {
		if !helpers.SymbolIsName(symbol) {
			return symbol, errors.New("Foreign name can not be normalized:" + foreignName + " symbol:" + symbol)
		}
	}
	if helpers.SymbolIsBlackListed(symbol) {
		return symbol, errors.New("Symbol is black listed:" + symbol)
	}
	return symbol, nil
}

func (scraper *QuoineScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	products, err := scraper.client.GetProducts(ctx)
	if err != nil {
		return
	}

	pairs = make([]dia.Pair, len(products))

	for _, prod := range products {
		symbol, serr := scraper.normalizeSymbol(prod.CurrencyPairCode, prod.BaseCurrency)
		if serr == nil {
			pairs = append(pairs, dia.Pair{
				Symbol:      symbol,
				ForeignName: prod.CurrencyPairCode,
				Exchange:    scraper.exchangeName,
			})
		} else {
			log.Error(serr)
		}
	}
	return
}

func (scraper *QuoineScraper) readProductIds() error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	products, err := scraper.client.GetProducts(ctx)
	if err != nil {
		return err
	}

	for _, prod := range products {
		// create a pair -> id mapping
		scraper.productPairIds[prod.CurrencyPairCode] = int(prod.ID)
	}

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
		parent: scraper,
		pair:   pair,
	}

	scraper.pairScrapers[pair.ForeignName] = pairScraper

	return pairScraper, nil
}
func (s *QuoineScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()
	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.shutdownDone)
}

func (scraper *QuoineScraper) Close() error {
	// close the pair scraper channels
	scraper.run = false
	for _, pairScraper := range scraper.pairScrapers {
		pairScraper.closed = true
	}

	close(scraper.shutdown)
	<-scraper.shutdownDone
	return nil
}

type QuoinePairScraper struct {
	parent *QuoineScraper
	pair   dia.Pair
	closed bool
}

func (pairScraper *QuoinePairScraper) Pair() dia.Pair {
	return pairScraper.pair
}

func (pairScraper *QuoineScraper) Channel() chan *dia.Trade {
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
	pairScraper.closed = true
	return nil
}
