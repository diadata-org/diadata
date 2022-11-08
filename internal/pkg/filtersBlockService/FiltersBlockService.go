package filters

import (
	"errors"
	"sync"
	"time"

	"github.com/cnf/structhash"
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

/*
const (
	filtersParam = dia.BlockSizeSeconds
)
*/

type nothing struct{}

// getIdentifier returns the unique identifier for asset @a.
// It is used as a part of a filter's map's key.
func getIdentifier(a dia.Asset) string {
	return a.Blockchain + "-" + a.Address
}

// filtersAsset is only used in the filters package. It is the auxilliary
// structure enabling us to compute prices for both, an asset on one exchange
// and an asset across all exchanges.
// @Identifier is the asset's unique identifier blockchain+Address
type filtersAsset struct {
	Identifier string
	Source     string
}

// FiltersBlockService is the data structure containing all objects
// necessary for the processing of a tradesBlock.
type FiltersBlockService struct {
	shutdown         chan nothing
	shutdownDone     chan nothing
	chanTradesBlock  chan *dia.TradesBlock
	chanFiltersBlock chan *dia.FiltersBlock
	errorLock        sync.RWMutex
	error            error
	closed           bool
	started          bool
	// currentTime          time.Time
	filters              map[filtersAsset][]Filter
	lastLog              time.Time
	calculationValues    []int
	previousBlockFilters []dia.FilterPoint
	datastore            models.Datastore
}

// NewFiltersBlockService returns a new FiltersBlockService and
// runs mainLoop() in a go routine.
func NewFiltersBlockService(previousBlockFilters []dia.FilterPoint, datastore models.Datastore, chanFiltersBlock chan *dia.FiltersBlock) *FiltersBlockService {
	s := &FiltersBlockService{
		shutdown:             make(chan nothing),
		shutdownDone:         make(chan nothing),
		chanTradesBlock:      make(chan *dia.TradesBlock),
		chanFiltersBlock:     chanFiltersBlock,
		error:                nil,
		started:              false,
		filters:              make(map[filtersAsset][]Filter),
		lastLog:              time.Now(),
		calculationValues:    make([]int, 0),
		previousBlockFilters: previousBlockFilters,
		datastore:            datastore,
	}
	s.calculationValues = append(s.calculationValues, dia.BlockSizeSeconds)

	go s.mainLoop()
	return s
}

// mainLoop runs processTradesBlock until FiltersBlockService @s is shut down.
func (s *FiltersBlockService) mainLoop() {
	for {
		log.Info("x FiltersBlockService mainloop")
		select {
		case <-s.shutdown:
			log.Println("Filters shutting down")
			s.cleanup(nil)
			return
		case tb, ok := <-s.chanTradesBlock:
			log.Info("receive tradesBlock for further processing ok: ", ok)
			s.processTradesBlock(tb)
		}
	}
}

// processTradesBlock is the 'main' function in the sense that all mathematical
// computations are done here.
func (s *FiltersBlockService) processTradesBlock(tb *dia.TradesBlock) {

	log.Infoln("processTradesBlock starting")
	t0 := time.Now()

	for _, trade := range tb.TradesBlockData.Trades {
		s.createFilters(trade.QuoteToken, "", tb.TradesBlockData.BeginTime)
		s.createFilters(trade.QuoteToken, trade.Source, tb.TradesBlockData.BeginTime)
		s.computeFilters(trade, "")
		s.computeFilters(trade, trade.Source)
	}

	log.Info("time spent for create and compute filters: ", time.Since(t0))
	log.Info("filter begin time: ", tb.TradesBlockData.BeginTime)
	resultFilters := []dia.FilterPoint{}

	t0 = time.Now()

	for _, filters := range s.filters {
		for _, f := range filters {
			f.finalCompute(tb.TradesBlockData.EndTime)
			fp := f.filterPointForBlock()
			if fp != nil {
				resultFilters = append(resultFilters, *fp)
			}
		}
	}
	log.Info("time spent for final compute: ", time.Since(t0))

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

	t0 = time.Now()
	for _, filters := range s.filters {
		for _, f := range filters {
			err = f.save(s.datastore)
			if err != nil {
				log.Error(err)
			}
		}
	}
	log.Info("time spent for save filters: ", time.Since(t0))

	err = s.datastore.ExecuteRedisPipe()
	if err != nil {
		log.Error("execute redis pipe: ", err)
	}

	err = s.datastore.FlushRedisPipe()
	if err != nil {
		log.Error("flush redis pipe: ", err)
	}

	err = s.datastore.Flush()
	if err != nil {
		log.Error("flush influx batch: ", err)
	}

}

func (s *FiltersBlockService) createFilters(asset dia.Asset, exchange string, BeginTime time.Time) {
	fa := filtersAsset{
		Identifier: getIdentifier(asset),
		Source:     exchange,
	}
	_, ok := s.filters[fa]
	if !ok {
		s.filters[fa] = []Filter{
			NewFilterMA(asset, exchange, BeginTime, dia.BlockSizeSeconds),
			NewFilterMAIR(asset, exchange, BeginTime, dia.BlockSizeSeconds),
			NewFilterMEDIR(asset, exchange, BeginTime, dia.BlockSizeSeconds),
			NewFilterVOL(asset, exchange, dia.BlockSizeSeconds),
			NewFilterCOUNT(asset, exchange, dia.BlockSizeSeconds),
			NewFilterTLT(asset, exchange),
		}
	}
}

func (s *FiltersBlockService) computeFilters(t dia.Trade, exchange string) {
	fa := filtersAsset{
		Identifier: getIdentifier(t.QuoteToken),
		Source:     exchange,
	}
	for _, f := range s.filters[fa] {
		f.compute(t)
	}
}

func addMissingPoints(previousBlockFilters []dia.FilterPoint, newFilters []dia.FilterPoint) []dia.FilterPoint {
	log.Debug("previousBlockFilters", previousBlockFilters)
	log.Debug("newFilters:", newFilters)
	missingPoints := 0
	result := newFilters
	newFiltersMap := make(map[filtersAsset]*dia.FilterPoint)
	for i, filter := range newFilters {
		fa := filtersAsset{
			Identifier: getIdentifier(filter.Asset),
			Source:     filter.Name,
		}
		newFiltersMap[fa] = &newFilters[i]
	}

	for _, filter := range previousBlockFilters {

		d := time.Since(filter.Time)
		// log.Info("filter:", filter, " age:", d)
		fa := filtersAsset{
			Identifier: getIdentifier(filter.Asset),
			Source:     filter.Name,
		}
		if d > time.Hour*24 {
			_, ok := newFiltersMap[fa]
			if !ok {
				result = append(result, filter)
				log.Debug("Adding", filter.Name+filter.Asset.Symbol)
				missingPoints++
			}
		} else {
			// log.Warn("ignoring old filter", filter.Asset.Symbol)
		}
	}
	if missingPoints != 0 {
		log.Printf("Added %v missing point from previous block", missingPoints)
	}
	log.Debug("result:", result)
	return result
}

// ProcessTradesBlock sends a filled tradesBlock into the filtersBlock channel.
func (s *FiltersBlockService) ProcessTradesBlock(tradesBlock *dia.TradesBlock) {
	s.chanTradesBlock <- tradesBlock
	log.Info("Processing TradesBlock done.")
}

// Close gracefully closes the Filtersblockservice
func (s *FiltersBlockService) Close() error {
	if s.closed {
		return errors.New("filters: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	return s.error
}

// cleanup must only be called from mainLoop
func (s *FiltersBlockService) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()
	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.shutdownDone) // signal that shutdown is complete
}
