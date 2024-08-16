package scrapers

import (
	"errors"
	"strings"
	"sync"
	"time"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/diadata-org/diadata/pkg/dia"
	alephiumhelper "github.com/diadata-org/diadata/pkg/dia/helpers/alephium-helper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
)

type BifrostScraper struct {
	logger *logrus.Entry
	// signaling channels
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock                 sync.RWMutex
	error                     error
	closed                    bool
	pairScrapers              map[string]*BifrostPairScraper // pc.ExchangePair -> pairScraperSet
	gsrpcApi                  *gsrpc.SubstrateAPI
	ticker                    *time.Ticker
	exchangeName              string
	blockchain                string
	currentHeight             int
	eventsLimit               int
	swapContractsLimit        int
	targetSwapContract        string
	chanTrades                chan *dia.Trade
	db                        *models.RelDB
	refreshDelay              time.Duration
	sleepBetweenContractCalls time.Duration
	reverseBasetokens         []string
	reverseQuotetokens        []string
}

type BifrostPairScraper struct {
	parent     *BifrostScraper
	pair       dia.ExchangePair
	closed     bool
	lastRecord int64
}

func (ps *BifrostPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}

func (ps *BifrostPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *BifrostPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

func NewBifrostScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *BifrostScraper {
	refreshDelay := utils.GetTimeDurationFromIntAsMilliseconds(
		utils.GetenvInt(strings.ToUpper(exchange.Name)+"_REFRESH_DELAY", alephiumhelper.DefaultRefreshDelay),
	)
	sleepBetweenContractCalls := utils.GetTimeDurationFromIntAsMilliseconds(
		utils.GetenvInt(strings.ToUpper(exchange.Name)+"_SLEEP_TIMEOUT", alephiumhelper.DefaultSleepBetweenContractCalls),
	)
	eventsLimit := utils.GetenvInt(strings.ToUpper(exchange.Name)+"_REFRESH_DELAY", alephiumhelper.DefaultEventsLimit)
	swapContractsLimit := utils.GetenvInt(strings.ToUpper(exchange.Name)+"_SWAP_CONTRACTS_LIMIT", alephiumhelper.DefaultSwapContractsLimit)
	targetSwapContract := utils.Getenv(strings.ToUpper(exchange.Name)+"_TARGET_SWAP_CONTRACT", "")

	logrus.Info("[BifrostScraper] Connecting to Bifrost RPC")
	client, err := gsrpc.NewSubstrateAPI("wss://bifrost-rpc.liebi.com/ws")
	if err != nil {
		logrus.WithError(err).Error("failed to create new SubstrateAPI")
		return nil
	}
	logrus.Info("[BifrostScraper] Connected to Bifrost RPC with success!!")

	s := &BifrostScraper{
		shutdown:                  make(chan nothing),
		shutdownDone:              make(chan nothing),
		pairScrapers:              make(map[string]*BifrostPairScraper),
		gsrpcApi:                  client,
		ticker:                    time.NewTicker(refreshDelay),
		exchangeName:              exchange.Name,
		blockchain:                exchange.BlockChain.Name,
		currentHeight:             0,
		error:                     nil,
		chanTrades:                make(chan *dia.Trade),
		db:                        relDB,
		refreshDelay:              refreshDelay,
		sleepBetweenContractCalls: sleepBetweenContractCalls,
		eventsLimit:               eventsLimit,
		swapContractsLimit:        swapContractsLimit,
		targetSwapContract:        targetSwapContract,
		logger:                    logrus.WithField("exchange", exchange.Name),
	}

	if scrape {
		go s.mainLoop()
	}
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (s *BifrostScraper) mainLoop() {
	for {
		println("BifrostScraper running...")
	}
}

// closes all connected PairScrapers
// must only be called from mainLoop
func (s *BifrostScraper) cleanup(err error) {

	s.errorLock.Lock()
	defer s.errorLock.Unlock()

	s.ticker.Stop()

	if err != nil {
		s.error = err
	}
	s.closed = true

	close(s.shutdownDone) // signal that shutdown is complete
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *BifrostScraper) Close() error {
	if s.closed {
		return errors.New("AlephiumScraper: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

func (s *BifrostScraper) Update() error {
	logger := s.logger.WithFields(logrus.Fields{
		"function": "Update",
	})

	// TODO implement it
	logger.Info("updating...")
	return nil
}

func (s *BifrostScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	logger := s.logger.WithFields(logrus.Fields{
		"function": "ScrapePair",
	})

	// TODO implement it
	logger.Info("scraping pair...")
	return nil, nil
}

// FetchAvailablePairs returns a list with all available trade pairs as dia.ExchangePair for the pairDiscorvery service
func (s *BifrostScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	logger := s.logger.WithFields(logrus.Fields{
		"function": "FetchAvailablePairs",
	})

	var mexcExchangeInfo *MEXCExchangeInfo
	mexcExchangeInfo.Symbols = make([]MEXCExchangeSymbol, 0)
	mexcExchangeInfo.Symbols = append(mexcExchangeInfo.Symbols, MEXCExchangeSymbol{
		BaseAsset:  "BTC",
		QuoteAsset: "USDT",
	})
	mexcExchangeInfo.Symbols = append(mexcExchangeInfo.Symbols, MEXCExchangeSymbol{
		BaseAsset:  "ETH",
		QuoteAsset: "USDT",
	})

	logger.Info("fetching available pairs...")
	for _, p := range mexcExchangeInfo.Symbols {
		pairToNormalized := dia.ExchangePair{
			Symbol:      p.BaseAsset,
			ForeignName: p.BaseAsset + p.QuoteAsset,
			Exchange:    s.exchangeName,
		}
		pairs = append(pairs, pairToNormalized)
	}
	return
}

// Channel returns a channel that can be used to receive trades/pricing information
func (s *BifrostScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

func (s *BifrostScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

func (s *BifrostScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}
