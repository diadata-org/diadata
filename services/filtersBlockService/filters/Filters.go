package filters

import (
	"errors"
	"github.com/cnf/structhash"
	"github.com/diadata-org/api-golang/dia"
	"github.com/diadata-org/api-golang/dia/helpers/configCollectors"
	"github.com/diadata-org/api-golang/services/model"
	log "github.com/sirupsen/logrus"
	"math"
	"sync"
	"time"
)

type nothing struct{}

type Filters struct {
	shutdown             chan nothing
	shutdownDone         chan nothing
	chanTradesBlock      chan *dia.TradesBlock
	chanFiltersBlock     chan *dia.FiltersBlock
	errorLock            sync.RWMutex
	error                error
	closed               bool
	started              bool
	currentTime          time.Time
	filters              map[string]*Filter
	lastLog              time.Time
	calculationValues    []int
	previousBlockFilters []dia.FilterPoint
	datastore            models.Datastore
}

func NewFilters(previousBlockFilters []dia.FilterPoint, datastore models.Datastore, chanFiltersBlock chan *dia.FiltersBlock) *Filters {
	s := &Filters{
		shutdown:             make(chan nothing),
		shutdownDone:         make(chan nothing),
		chanTradesBlock:      make(chan *dia.TradesBlock),
		chanFiltersBlock:     chanFiltersBlock,
		error:                nil,
		started:              false,
		filters:              make(map[string]*Filter),
		lastLog:              time.Now(),
		calculationValues:    make([]int, 0),
		previousBlockFilters: previousBlockFilters,
		datastore:            datastore,
	}
	s.calculationValues = append(s.calculationValues, dia.BlockSizeSeconds)
	go s.mainLoop()
	return s
}

func (s *Filters) ProcessTradesBlock(tradesBlock *dia.TradesBlock) {
	s.chanTradesBlock <- tradesBlock
}

func (s *Filters) Close() error {
	if s.closed {
		return errors.New("Filters: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	return s.error
}

// must only be called from mainLoop
func (s *Filters) cleanup(err error) {
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
				log.Printf("Adding", filter.Name+filter.Symbol)
				missingPoints++
			}
		}
	}
	if missingPoints != 0 {
		log.Printf("Added %v missing point from previous block", missingPoints)
	}
	log.Info("result:", result)
	return result
}

func (s *Filters) processTradesBlockForVolumes(tb *dia.TradesBlock) {
	var volumes = make(map[string]float64)
	log.Info("processTradesBlock")
	for _, trade := range tb.TradesBlockData.Trades {
		f, ok := volumes[trade.Symbol]
		if !ok {
			volumes[trade.Symbol] = 0.0
		}
		volumes[trade.Symbol] = f + trade.EstimatedUSDPrice*math.Abs(trade.Volume)
	}

	for symbol, value := range volumes {
		err := s.datastore.SetVolume(symbol, value)
		if err != nil {
			log.Printf("Error: %v\n", err)
		}
	}

}

func (s *Filters) processTradesBlock(tb *dia.TradesBlock) {

	for _, trade := range tb.TradesBlockData.Trades {
		f, ok := s.filters[trade.Symbol]
		if !ok {
			f = NewFilter(s.calculationValues, trade.Symbol, tb.TradesBlockData.BeginTime, trade.EstimatedUSDPrice)
			log.Info("create filter ", trade.Symbol)
			s.filters[trade.Symbol] = f
		}

		f.ProcessTrade(&trade)
	}
	resultFilters := []dia.FilterPoint{}
	for _, filter := range s.filters {
		filter.calculateMAs(tb.TradesBlockData.EndTime)
		for key, result := range filter.Results {
			resultFilters = append(resultFilters, dia.FilterPoint{
				Symbol: filter.Symbol,
				Value:  result,
				Name:   key,
				Time:   tb.TradesBlockData.EndTime,
			})
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

	log.Printf("Generating Filters block %v (size:%v)", fb.FiltersBlockData.FiltersBlockNumber, fb.FiltersBlockData.FiltersNumber)

	hash, err := structhash.Hash(fb.FiltersBlockData, 1)
	if err != nil {
		log.Printf("error on hash")
		hash = "hashError"
	}
	fb.BlockHash = hash

	//maForDia := "MA120"

	if len(resultFilters) != 0 {
		s.chanFiltersBlock <- fb
		for _, f := range fb.FiltersBlockData.FilterPoints {
			//if f.Name == maForDia {
			err := s.datastore.SetPriceUSD(f.Symbol, f.Value)
			if err != nil {
				log.Error("Error: %v\n", err)
			}

			err = s.datastore.SetPriceZSET(f.Symbol, f.Value)
			if err != nil {
				log.Error("Error: %v\n", err)
			}
			//	}
		}
	}
}

// runs in a goroutine until s is closed
func (s *Filters) mainLoop() {
	for {
		select {
		case <-s.shutdown:
			log.Println("Filters shutting down")
			s.cleanup(nil)
			return
		case tb, _ := <-s.chanTradesBlock:
			s.processTradesBlock(tb)
			s.processTradesBlockForVolumes(tb)
		}
	}
}
