package scrapers

import (
	"context"
	"sync"

	"github.com/diadata-org/diadata/pkg/dia"

	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

const (
	orcaRestDialSolana       = ""
	orcaWsDialSolana         = ""
	orcaProgramAddrWhirlpool = "whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc"
)

type OrcaScraper struct {
	exchangeName string

	RestClient *rpc.Client
	WsClient   *ws.Client

	// state variables to signal events
	run          bool
	shutdown     chan nothing
	shutdownDone chan nothing

	errorLock sync.RWMutex
	error     error

	pairScrapers map[string]*OrcaPairScraper
	chanTrades   chan *dia.Trade
}

func (s *OrcaScraper) mainLoop() {
	s.run = true
}

func NewOrcaScraper(exchange dia.Exchange, scrape bool) *OrcaScraper {

	log.Infof("init rest and ws client for %s", exchange.BlockChain.Name)
	wsClient, err := ws.Connect(context.Background(), orcaWsDialSolana)
	if err != nil {
		log.Fatal("init ws client: ", err)
	}

	scraper := &OrcaScraper{
		exchangeName: exchange.Name,
		RestClient:   rpc.New(orcaRestDialSolana),
		WsClient:     wsClient,
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*OrcaPairScraper),
		chanTrades:   make(chan *dia.Trade),
	}

	if scrape {
		go scraper.mainLoop()
	}
	return scraper
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from the scraper
func (s *OrcaScraper) ScrapePair(pair dia.ExchangePair) (ps PairScraper, err error) {
	return
}

// Returns the list of all available trade pairs
func (s *OrcaScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	return pairs, nil
}

// Channel returns a channel that can be used to receive trades
func (s *OrcaScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

// FillSymbolData adds the name to the asset underlying @symbol on scraper
func (s *OrcaScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

// NormalizePair accounts for the pair
func (s *OrcaScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

type OrcaPairScraper struct {
	parent *OrcaScraper
	pair   dia.ExchangePair
	closed bool
}

// Close stops listening for trades of the pair associated with the scraper
func (ps *OrcaPairScraper) Close() error {
	ps.parent.errorLock.RLock()
	defer ps.parent.errorLock.RUnlock()
	ps.closed = true
	return nil
}

// Error returns an error when the channel Channel() is closed and nil otherwise
func (ps *OrcaPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *OrcaPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
