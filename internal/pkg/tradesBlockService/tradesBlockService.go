package tradesBlockService

import (
	"errors"
	"github.com/cnf/structhash"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
	"sort"
	"sync"
	"time"
)

type nothing struct{}

type TradesBlockService struct {
	pair            string
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

func (ps *TradesBlockService) Channel() chan *dia.TradesBlock {
	return ps.chanTradesBlock
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

func (s *TradesBlockService) process(t dia.Trade) {

	var ignoreTrade bool
	secondPair := t.SecondPair()
	if secondPair != "USD" {
		val, err := s.datastore.GetPriceUSD(secondPair)
		if err != nil {
			log.Error("redisClient error ", err, " ignoring ", t)
			ignoreTrade = true
		} else {
			t.EstimatedUSDPrice = t.Price * val
		}
	} else {
		t.EstimatedUSDPrice = t.Price
	}

	if ignoreTrade == false {
		s.datastore.SaveTradeInflux(&t)
	}

	if s.currentBlock != nil &&
		s.currentBlock.TradesBlockData.BeginTime.After(t.Time) {
		log.Debug("ignore trade should be in previous block %v", t)
		ignoreTrade = true
	}

	if ignoreTrade == false {

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

			log.Printf("created new block beginTime: %v", b.TradesBlockData.BeginTime)

			s.currentBlock = b
			s.datastore.Flush()
		}
		s.currentBlock.TradesBlockData.Trades = append(s.currentBlock.TradesBlockData.Trades, t)
	} else {
		log.Debug("ignore trade  %v", t)
	}
}

// runs in a goroutine until s is closed
func (s *TradesBlockService) mainLoop() {
	for {
		select {
		case <-s.shutdown:
			log.Println("TradesBlockService shutting down")
			s.cleanup(nil)
			return
		case t, _ := <-s.chanTrades:
			s.process(*t)
		}
	}
}
