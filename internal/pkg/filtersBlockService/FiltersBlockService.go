package filters

import (
	"errors"
	"github.com/cnf/structhash"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

const (
	filtersParam = dia.BlockSizeSeconds
)

type nothing struct{}

type FiltersBlockService struct {
	shutdown             chan nothing
	shutdownDone         chan nothing
	chanTradesBlock      chan *dia.TradesBlock
	chanFiltersBlock     chan *dia.FiltersBlock
	errorLock            sync.RWMutex
	error                error
	closed               bool
	started              bool
	currentTime          time.Time
	filters              map[string][]Filter
	lastLog              time.Time
	calculationValues    []int
	previousBlockFilters []dia.FilterPoint
	datastore            models.Datastore
}

func NewFiltersBlockService(previousBlockFilters []dia.FilterPoint, datastore models.Datastore, chanFiltersBlock chan *dia.FiltersBlock) *FiltersBlockService {
	s := &FiltersBlockService{
		shutdown:             make(chan nothing),
		shutdownDone:         make(chan nothing),
		chanTradesBlock:      make(chan *dia.TradesBlock),
		chanFiltersBlock:     chanFiltersBlock,
		error:                nil,
		started:              false,
		filters:              make(map[string][]Filter),
		lastLog:              time.Now(),
		calculationValues:    make([]int, 0),
		previousBlockFilters: previousBlockFilters,
		datastore:            datastore,
	}
	s.calculationValues = append(s.calculationValues, dia.BlockSizeSeconds)

	go s.mainLoop()
	return s
}

func (s *FiltersBlockService) ProcessTradesBlock(tradesBlock *dia.TradesBlock) {
	s.chanTradesBlock <- tradesBlock
	log.Info("ProcessTradesBlock finito")
}

func (s *FiltersBlockService) Close() error {
	if s.closed {
		return errors.New("Filters: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	return s.error
}

// must only be called from mainLoop
func (s *FiltersBlockService) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()
	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.shutdownDone) // signal that shutdown is complete
}

func addMissingPoints(previousBlockFilters []dia.FilterPoint, newFilters []dia.FilterPoint) []dia.FilterPoint {
	log.Debug("previousBlockFilters", previousBlockFilters)
	log.Debug("newFilters:", newFilters)
	missingPoints := 0
	result := newFilters
	newFiltersMap := make(map[string]*dia.FilterPoint)
	for _, filter := range newFilters {
		newFiltersMap[filter.Name+filter.Symbol] = &filter
	}

	for _, filter := range previousBlockFilters {

		d := time.Now().Sub(filter.Time)
		log.Info("filter:", filter, " age:", d)

		if d > time.Hour*24 {
			_, ok := newFiltersMap[filter.Name+filter.Symbol]
			if !ok {
				result = append(result, filter)
				log.Debug("Adding", filter.Name+filter.Symbol)
				missingPoints++
			}
		} else {
			log.Info("ignoring old filter", filter.Symbol)
		}
	}
	if missingPoints != 0 {
		log.Printf("Added %v missing point from previous block", missingPoints)
	}
	log.Debug("result:", result)
	return result
}

func (s *FiltersBlockService) createFilters(symbol string, exchange string, BeginTime time.Time) {
	_, ok := s.filters[symbol+exchange]
	if !ok {
		s.filters[symbol+exchange] = []Filter{
			NewFilterMA(symbol, exchange, BeginTime, dia.BlockSizeSeconds),
			NewFilterTLT(symbol, exchange),
			NewFilterVOL(symbol, exchange, dia.BlockSizeSeconds),
			NewFilterMAIR(symbol, exchange, BeginTime, dia.BlockSizeSeconds),
			NewFilterMEDIR(symbol, exchange, BeginTime, dia.BlockSizeSeconds),
		}
	}
}

func (s *FiltersBlockService) computeFilters(t dia.Trade, key string) {
	for _, f := range s.filters[key] {
		f.compute(t)
	}
}

func (s *FiltersBlockService) processTradesBlock(tb *dia.TradesBlock) {

	log.Infoln("processTradesBlock starting")

	for _, trade := range tb.TradesBlockData.Trades {
		s.createFilters(trade.Symbol, "", tb.TradesBlockData.BeginTime)
		s.createFilters(trade.Symbol, trade.Source, tb.TradesBlockData.BeginTime)
		s.computeFilters(trade, trade.Symbol)
		s.computeFilters(trade, trade.Symbol+trade.Source)
	}

	resultFilters := []dia.FilterPoint{}
	for _, filters := range s.filters {
		for _, f := range filters {
			f.finalCompute(tb.TradesBlockData.EndTime)
			fp := f.filterPointForBlock()
			if fp != nil {
				resultFilters = append(resultFilters, *fp)
			}
		}
	}

	resultFilters = addMissingPoints(s.previousBlockFilters, resultFilters)

	s.previousBlockFilters = resultFilters

	fb := &dia.FiltersBlock{
		FiltersBlockData: dia.FiltersBlockData{
			FilterPoints:    resultFilters,
			FiltersNumber:   len(resultFilters),
			EndTime:         tb.TradesBlockData.EndTime,
			BeginTime:       tb.TradesBlockData.BeginTime,
			TradesBlockHash: tb.BlockHash,
		},
	}

	hash, err := structhash.Hash(fb.FiltersBlockData, 1)
	if err != nil {
		log.Printf("error on hash")
		hash = "hashError"
	}
	fb.BlockHash = hash
	log.Printf("Generating Filters block %v (size:%v)", hash, fb.FiltersBlockData.FiltersNumber)

	if len(resultFilters) != 0 && s.chanFiltersBlock != nil {
		s.chanFiltersBlock <- fb
	}
	for _, filters := range s.filters {
		for _, f := range filters {
			f.save(s.datastore)
		}
	}
	s.datastore.Flush()
	c, err := s.datastore.GetCoins()
	if err == nil {
		for i, v := range c.Coins {
			log.Info("UpdateSymbolDetails on ", v.Symbol)
			s.datastore.UpdateSymbolDetails(v.Symbol, i+1)
		}
	}
}

// runs in a goroutine until s is closed
func (s *FiltersBlockService) mainLoop() {
	for {
		log.Info("x FiltersBlockService mainloop")
		select {
		case <-s.shutdown:
			log.Println("Filters shutting down")
			s.cleanup(nil)
			return
		case tb, _ := <-s.chanTradesBlock:
			s.processTradesBlock(tb)
		}
	}
}
