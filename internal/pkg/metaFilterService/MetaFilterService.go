package metafilters

import (
	"errors"
	"sync"
	"time"

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

// getIdentifier returns the unique identifier for pair @p.
// It is used as a part of a filter's map's key.
func getIdentifier(asset dia.Asset) string {
	return asset.Blockchain + "-" + asset.Address
}

// filtersPair is only used in the filters package. It is the auxilliary
// structure enabling us to compute prices for both, a pair on one exchange
// and a pair across all exchanges.
// @Identifier is the asset's unique identifier blockchain+Address
type metaFilterIdentifier struct {
	IdentifierQuotetoken string
	filterName           string
}

// FiltersBlockService is the data structure containing all objects
// necessary for the processing of a tradesBlock.
type MetaFilterService struct {
	shutdown         chan nothing
	shutdownDone     chan nothing
	chanFiltersBlock chan *dia.FiltersBlock
	// chanPairFiltersBlock chan *dia.PairFiltersBlock
	errorLock sync.RWMutex
	error     error
	closed    bool
	started   bool
	// currentTime          time.Time
	metaFilters       map[metaFilterIdentifier][]MetaFilter
	lastLog           time.Time
	calculationValues []int
	datastore         models.Datastore
}

// NewFiltersBlockService returns a new FiltersBlockService and
// runs mainLoop() in a go routine.
func NewMetaFilterService(datastore models.Datastore, chanFiltersBlock chan *dia.FiltersBlock) *MetaFilterService {
	s := &MetaFilterService{
		shutdown:          make(chan nothing),
		shutdownDone:      make(chan nothing),
		chanFiltersBlock:  chanFiltersBlock,
		error:             nil,
		started:           false,
		metaFilters:       make(map[metaFilterIdentifier][]MetaFilter),
		lastLog:           time.Now(),
		calculationValues: make([]int, 0),
		datastore:         datastore,
	}
	s.calculationValues = append(s.calculationValues, dia.BlockSizeSeconds)

	go s.mainLoop()
	return s
}

// mainLoop runs processTradesBlock until FiltersBlockService @s is shut down.
func (s *MetaFilterService) mainLoop() {
	for {
		log.Info("x FiltersBlockService mainloop")
		select {
		case <-s.shutdown:
			log.Println("Filters shutting down")
			s.cleanup(nil)
			return
		case fb, ok := <-s.chanFiltersBlock:
			log.Info("receive tradesBlock for further processing ok: ", ok)
			s.processFiltersBlock(fb)
		}
	}
}

// processFiltersBlock is the 'main' function in the sense that all mathematical
// computations are done here.
func (s *MetaFilterService) processFiltersBlock(fb *dia.FiltersBlock) {

	log.Infoln("processFiltersBlock starting")
	t0 := time.Now()

	for _, filterPoint := range fb.FiltersBlockData.FilterPoints {
		s.createMetaFilters(filterPoint, fb.FiltersBlockData.BeginTime)
		s.computeMetaFilters(filterPoint)
	}

	resultFilters := []dia.MetaFilterPoint{}

	t0 = time.Now()

	for _, filters := range s.metaFilters {
		for _, f := range filters {
			f.finalCompute(fb.FiltersBlockData.EndTime)
			mfp := f.filterPointForBlock()
			if mfp != nil {
				resultFilters = append(resultFilters, *mfp)
			}
		}
	}
	log.Info("time spent for final compute: ", time.Since(t0))

	t0 = time.Now()
	var err error
	for _, filters := range s.metaFilters {
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

func (s *MetaFilterService) createMetaFilters(filterPoint dia.FilterPoint, BeginTime time.Time) {
	mfi := metaFilterIdentifier{
		IdentifierQuotetoken: getIdentifier(filterPoint.Pair.QuoteToken),
		filterName:           filterPoint.Name,
	}
	_, ok := s.metaFilters[mfi]
	if !ok {
		s.metaFilters[mfi] = []MetaFilter{
			// NewFilterMA(filterPoint.Pair.QuoteToken, BeginTime, dia.BlockSizeSeconds),
			NewFilterAIR(filterPoint.Pair.QuoteToken, BeginTime),
			// NewFilterMEDIR(pair, exchange, BeginTime, dia.BlockSizeSeconds),
			NewFilterVOL(filterPoint.Pair.QuoteToken, filterPoint.Source, dia.BlockSizeSeconds),
			NewFilterCOUNT(filterPoint.Pair.QuoteToken, filterPoint.Source, dia.BlockSizeSeconds),
			NewFilterTLT(filterPoint.Pair.QuoteToken, filterPoint.Source),
		}
	}
}

func (s *MetaFilterService) computeMetaFilters(filterPoint dia.FilterPoint) {
	mfi := metaFilterIdentifier{
		IdentifierQuotetoken: getIdentifier(filterPoint.Pair.QuoteToken),
		filterName:           filterPoint.Name,
	}
	for _, f := range s.metaFilters[mfi] {
		f.collect(filterPoint)
	}
}

// ProcessFiltersBlock sends a filled fitlersBlock into the filtersBlock channel.
func (s *MetaFilterService) ProcessFiltersBlock(filtersBlock *dia.FiltersBlock) {
	s.chanFiltersBlock <- filtersBlock
	log.Info("Processing TradesBlock done.")
}

// Close gracefully closes the Filtersblockservice
func (s *MetaFilterService) Close() error {
	if s.closed {
		return errors.New("filters: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	return s.error
}

// cleanup must only be called from mainLoop
func (s *MetaFilterService) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()
	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.shutdownDone) // signal that shutdown is complete
}
