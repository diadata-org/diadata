package tradesBlockService

import (
	"errors"
	"math"
	"sort"
	"sync"
	"time"

	"github.com/cnf/structhash"
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

type TradesBlockService struct {
	shutdown        chan nothing
	shutdownDone    chan nothing
	chanTrades      chan *dia.Trade
	chanTradesBlock chan *dia.TradesBlock
	errorLock       sync.RWMutex
	error           error
	closed          bool
	started         bool
	BlockDuration   int64
	currentBlock    *dia.TradesBlock
	datastore       models.Datastore
}

func NewTradesBlockService(datastore models.Datastore, blockDuration int64) *TradesBlockService {
	s := &TradesBlockService{
		shutdown:        make(chan nothing),
		shutdownDone:    make(chan nothing),
		chanTrades:      make(chan *dia.Trade),
		chanTradesBlock: make(chan *dia.TradesBlock),
		error:           nil,
		started:         false,
		currentBlock:    nil,
		BlockDuration:   blockDuration,
		datastore:       datastore,
	}
	go s.mainLoop()
	return s
}

// runs in a goroutine until s is closed
func (s *TradesBlockService) mainLoop() {
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

func (s *TradesBlockService) process(t dia.Trade) {

	var verifiedTrade bool
	// baseTokenSymbol := t.GetBaseToken()

	// Price estimation can only be done for verified pairs.
	// Trades with unverified pairs are still saved, but not sent to the filtersBlockService.
	if t.VerifiedPair {
		if t.BaseToken.Symbol == "USD" && t.BaseToken.Blockchain == "fiat" {
			// All prices are measured in US-Dollar, so just price for base token == USD
			t.EstimatedUSDPrice = t.Price
			verifiedTrade = true
		} else {
			// Get price of base token.
			// This can be switched to GetAssetPriceUSD(asset, timestamp) when switching to historical scrapers.
			val, err := s.datastore.GetAssetPriceUSDCache(t.BaseToken)
			if err != nil {
				log.Errorf("Cannot use trade %s. Can't find quotation for base token", t.Pair)
			} else {
				if val > 0.0 {
					log.Infof("price of trade %s: %v", t.Pair, t.Price)
					log.Info("price of base token: ", val)
					t.EstimatedUSDPrice = t.Price * val
					verifiedTrade = true
				}
			}
		}
	}

	// // If estimated price for stablecoin diverges too much ignore trade
	if _, ok := stablecoins[t.Symbol]; ok {
		if math.Abs(t.EstimatedUSDPrice-1) > tol {
			log.Errorf("price for stablecoin %s diverges by %v", t.Symbol, math.Abs(t.EstimatedUSDPrice-1))
			verifiedTrade = false
		}
	}
	// Comment Philipp: We could make another check here. Store CG and/or CMC quotation in redis cache
	// and compare with estimatedUSDPrice. If deviation is too large ignore trade. If we do so,
	// we should already think about how to do it best with regards to historic values, as these are coming up.

	err := s.datastore.SaveTradeInflux(&t)
	if err != nil {
		log.Error(err)
	}

	if s.currentBlock != nil && s.currentBlock.TradesBlockData.BeginTime.After(t.Time) {
		log.Debugf("ignore trade should be in previous block %v", t)
		verifiedTrade = false
	}

	// Only verified trades of verified pairs are added to the tradesBlock
	if verifiedTrade {
		if s.currentBlock == nil || s.currentBlock.TradesBlockData.EndTime.Before(t.Time) {
			if s.currentBlock != nil {
				s.finaliseCurrentBlock()
			}

			b := &dia.TradesBlock{
				TradesBlockData: dia.TradesBlockData{
					Trades:    []dia.Trade{},
					EndTime:   time.Unix((t.Time.Unix()/s.BlockDuration)*s.BlockDuration+s.BlockDuration, 0),
					BeginTime: time.Unix((t.Time.Unix()/s.BlockDuration)*s.BlockDuration, 0),
				},
			}
			if s.currentBlock != nil {
				log.Info("created new block beginTime:", b.TradesBlockData.BeginTime, "previous block nb trades:", len(s.currentBlock.TradesBlockData.Trades))
			}
			s.currentBlock = b
			err = s.datastore.Flush()
			if err != nil {
				log.Error(err)
			}
		}
		s.currentBlock.TradesBlockData.Trades = append(s.currentBlock.TradesBlockData.Trades, t)
	} else {
		log.Debugf("ignore trade  %v", t)
	}
}

func (s *TradesBlockService) finaliseCurrentBlock() {

	sort.Slice(s.currentBlock.TradesBlockData.Trades, func(i, j int) bool {
		return s.currentBlock.TradesBlockData.Trades[i].Time.Before(s.currentBlock.TradesBlockData.Trades[j].Time)
	})

	hash, err := structhash.Hash(s.currentBlock.TradesBlockData, 1)
	if err != nil {
		log.Printf("error on hash")
		hash = "hashError"
	}
	s.currentBlock.BlockHash = hash
	s.currentBlock.TradesBlockData.TradesNumber = len(s.currentBlock.TradesBlockData.Trades)
	s.chanTradesBlock <- s.currentBlock
}

func (s *TradesBlockService) ProcessTrade(trade *dia.Trade) {
	s.chanTrades <- trade
}

func (s *TradesBlockService) Close() error {
	if s.closed {
		return errors.New("TradesBlockService: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	return s.error
}

// must only be called from mainLoop
func (s *TradesBlockService) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()
	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.shutdownDone) // signal that shutdown is complete
}

func (s *TradesBlockService) Channel() chan *dia.TradesBlock {
	return s.chanTradesBlock
}
