package filters

import (
	"errors"
	"github.com/cnf/structhash"
	"github.com/diadata-org/api-golang/dia"
	"github.com/diadata-org/api-golang/dia/helpers/configCollectors"
	"github.com/diadata-org/api-golang/services/model"
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
	filters              map[string]Filter
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
		filters:              make(map[string]Filter),
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
	log.Info("previousBlockFilters", previousBlockFilters)
	log.Info("newFilters:", newFilters)
	c := configCollectors.NewConfigCollectors()
	missingPoints := 0
	result := newFilters
	newFiltersMap := make(map[string]*dia.FilterPoint)
	for _, filter := range newFilters {
		newFiltersMap[filter.Name+filter.Symbol] = &filter
	}

	for _, filter := range previousBlockFilters {
		if c.IsSymbolInConfig(filter.Symbol) {
			_, ok := newFiltersMap[filter.Name+filter.Symbol]
			if !ok {
				result = append(result, filter)
				log.Println("Adding", filter.Name+filter.Symbol)
				missingPoints++
			}
		}
	}
	if missingPoints != 0 {
		log.Printf("Added %v missing point from previous block", missingPoints)
	}
	log.Debug("result:", result)
	return result
}

func (s *FiltersBlockService) processTradesBlock(tb *dia.TradesBlock) {

	for _, trade := range tb.TradesBlockData.Trades {
		f, ok := s.filters[trade.Symbol]
		if !ok {
			f = NewFilterMA(trade.Symbol, tb.TradesBlockData.BeginTime, filtersParam)
			s.filters[trade.Symbol] = f
			f2 := NewFilterVolume(trade.Symbol, tb.TradesBlockData.BeginTime, filtersParam)
			s.filters[trade.Symbol] = f2
		}
		f.compute(&trade)
	}

	resultFilters := []dia.FilterPoint{}
	for _, filter := range s.filters {
		if filter.copyToFilterBlock() {
			resultFilters = append(resultFilters, filter.filterPoint(tb.TradesBlockData.EndTime))
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

	if len(resultFilters) != 0 {
		s.chanFiltersBlock <- fb
		for _, filter := range s.filters {
			filter.save(s.datastore)
		}
	}
}

// runs in a goroutine until s is closed
func (s *FiltersBlockService) mainLoop() {
	for {
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
