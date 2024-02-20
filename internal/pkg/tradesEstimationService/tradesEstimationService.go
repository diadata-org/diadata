package tradesEstimationService

import (
	"errors"
	"math"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/sirupsen/logrus"
)

type nothing struct{}

var log *logrus.Logger

func init() {
	log = logrus.New()
}

var (
	stablecoins = map[string]interface{}{
		"USDC": "",
		"USDT": "",
		"TUSD": "",
		"DAI":  "",
		"PAX":  "",
		"BUSD": "",
	}
	tol = float64(0.1)
)

const (
	priceFrame = 1000 * 120
)

type pricetime struct {
	Price     float64
	Timestamp time.Time
}

type TradesEstimationService struct {
	shutdown     chan nothing
	shutdownDone chan nothing
	chanTrades   chan *dia.Trade
	errorLock    sync.RWMutex
	error        error
	closed       bool
	started      bool
	priceCache   map[dia.Asset]pricetime
	datastore    models.Datastore
}

func NewTradesEstimationService(datastore models.Datastore) *TradesEstimationService {
	s := &TradesEstimationService{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		chanTrades:   make(chan *dia.Trade),
		error:        nil,
		started:      false,
		priceCache:   make(map[dia.Asset]pricetime),
		datastore:    datastore,
	}
	go s.mainLoop()
	return s
}

// runs in a goroutine until s is closed
func (s *TradesEstimationService) mainLoop() {
	for {
		select {
		case <-s.shutdown:
			log.Println("TradesBlockService shutting down")
			s.cleanup(nil)
			return
		case t := <-s.chanTrades:
			s.process(*t)
		}
	}
}

func (s *TradesEstimationService) process(t dia.Trade) {

	var verifiedTrade bool
	var price float64
	var err error

	// Price estimation can only be done for verified pairs.
	if t.VerifiedPair {
		if t.BaseToken.Address == "840" && t.BaseToken.Blockchain == dia.FIAT {
			// All prices are measured in US-Dollar, so just price for base token == USD
			t.EstimatedUSDPrice = t.Price
			verifiedTrade = true
		} else {
			// Check if price cache is still valid:
			_, ok := s.priceCache[t.BaseToken]
			if ok && t.Time.Sub(s.priceCache[t.BaseToken].Timestamp) < time.Duration(priceFrame*time.Millisecond) {
				price = s.priceCache[t.BaseToken].Price
			} else {
				// Look for historic price of base token at trade time...
				price, err = s.datastore.GetAssetPriceUSD(t.BaseToken, time.Time{}, t.Time)
				s.priceCache[t.BaseToken] = pricetime{
					Price:     price,
					Timestamp: t.Time,
				}
			}
			if err != nil {
				log.Errorf("Cannot use trade %s. Can't find quotation for base token.", t.Pair)
			} else {
				if price > 0.0 {
					t.EstimatedUSDPrice = t.Price * price
					if t.EstimatedUSDPrice > 0 {
						verifiedTrade = true
					}
				}
			}
		}
	}

	// If estimated price for stablecoin diverges too much ignore trade.
	if _, ok := stablecoins[t.Symbol]; ok {
		if math.Abs(t.EstimatedUSDPrice-1) > tol {
			log.Errorf("price for stablecoin %s diverges by %v", t.Symbol, math.Abs(t.EstimatedUSDPrice-1))
			verifiedTrade = false
		}
	}

	if verifiedTrade {
		err := s.datastore.SaveTradeInflux(&t)
		if err != nil {
			log.Error(err)
		}
	}
}

func (s *TradesEstimationService) ProcessTrade(trade *dia.Trade) {
	s.chanTrades <- trade
}

func (s *TradesEstimationService) Close() error {
	if s.closed {
		return errors.New("TradesBlockService: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	return s.error
}

// must only be called from mainLoop
func (s *TradesEstimationService) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()
	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.shutdownDone) // signal that shutdown is complete
}
