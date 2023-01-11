package pairfilters

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

// getIdentifier returns the unique identifier for pair @p.
// It is used as a part of a filter's map's key.
func getIdentifier(asset dia.Asset) string {
	return asset.Blockchain + "-" + asset.Address
}

// filtersPair is only used in the filters package. It is the auxilliary
// structure enabling us to compute prices for both, a pair on one exchange
// and a pair across all exchanges.
// @Identifier is the asset's unique identifier blockchain+Address
type filtersPair struct {
	IdentifierQuotetoken string
	IdentifierBasetoken  string
	Source               string
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
	filters           map[filtersPair][]Filter
	lastLog           time.Time
	calculationValues []int
	datastore         models.Datastore
}

// NewFiltersBlockService returns a new FiltersBlockService and
// runs mainLoop() in a go routine.
func NewFiltersBlockService(datastore models.Datastore, chanFiltersBlock chan *dia.FiltersBlock) *FiltersBlockService {
	s := &FiltersBlockService{
		shutdown:          make(chan nothing),
		shutdownDone:      make(chan nothing),
		chanTradesBlock:   make(chan *dia.TradesBlock),
		chanFiltersBlock:  chanFiltersBlock,
		error:             nil,
		started:           false,
		filters:           make(map[filtersPair][]Filter),
		lastLog:           time.Now(),
		calculationValues: make([]int, 0),
		datastore:         datastore,
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
		s.createFilters(trade, "", tb.TradesBlockData.BeginTime)
		s.createFilters(trade, trade.Source, tb.TradesBlockData.BeginTime)
		s.computeFilters(trade, "")
		s.computeFilters(trade, trade.Source)
	}

	log.Info("time spent for create and compute filters: ", time.Since(t0))
	log.Info("filter begin time: ", tb.TradesBlockData.BeginTime)
	resultFilters := []dia.PairFilterPoint{}

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

func (s *FiltersBlockService) createFilters(trade dia.Trade, exchange string, BeginTime time.Time) {
	fa := filtersPair{
		IdentifierQuotetoken: getIdentifier(trade.QuoteToken),
		IdentifierBasetoken:  getIdentifier(trade.BaseToken),
		Source:               exchange,
	}
	pair := dia.Pair{QuoteToken: trade.QuoteToken, BaseToken: trade.BaseToken}
	_, ok := s.filters[fa]
	if !ok {
		s.filters[fa] = []Filter{
			NewFilterMA(pair, exchange, BeginTime, dia.BlockSizeSeconds),
			NewFilterMAIR(pair, exchange, BeginTime, dia.BlockSizeSeconds),
			NewFilterMEDIR(pair, exchange, BeginTime, dia.BlockSizeSeconds),
			NewFilterVOL(pair, exchange, dia.BlockSizeSeconds),
			NewFilterCOUNT(trade.QuoteToken, trade.BaseToken, exchange, dia.BlockSizeSeconds),
			NewFilterTLT(pair, exchange),
		}
	}
}

func (s *FiltersBlockService) computeFilters(trade dia.Trade, exchange string) {
	fa := filtersPair{
		IdentifierQuotetoken: getIdentifier(trade.QuoteToken),
		IdentifierBasetoken:  getIdentifier(trade.BaseToken),
		Source:               exchange,
	}
	for _, f := range s.filters[fa] {
		f.compute(trade)
	}
}

// ProcessTradesBlock sends a filled tradesBlock into the filtersBlock channel.
func (s *FiltersBlockService) ProcessTradesBlock(tradesBlock *dia.TradesBlock) {
	s.chanTradesBlock <- tradesBlock
	log.Info("Processing TradesBlock in PairFilterService done.")
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
