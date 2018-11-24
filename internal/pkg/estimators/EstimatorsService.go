package estimators

import (
	"errors"
	"github.com/diadata-org/diadata/pkg/dia"
	log "github.com/sirupsen/logrus"
	"strings"
	"sync"
	"time"
)

type nothing struct{}

//EstimatorsService Process trades and estimates PDF's
type EstimatorsService struct {
	shutdown         chan nothing
	shutdownDone     chan nothing
	chanTradesBlock  chan *dia.TradesBlock
	chanPDFEstimates chan PDF
	errorLock        sync.RWMutex
	error            error
	closed           bool
	started          bool
	currentTime      time.Time
	estimators       map[string][]PDFEstimator
	lastLog          time.Time
}

//NewEstimatorsService Open channels and start main lopp
func NewEstimatorsService(chanPDFEstimates chan PDF) *EstimatorsService {
	s := &EstimatorsService{
		shutdown:         make(chan nothing),
		shutdownDone:     make(chan nothing),
		chanTradesBlock:  make(chan *dia.TradesBlock),
		chanPDFEstimates: chanPDFEstimates,
		error:            nil,
		started:          false,
		estimators:       make(map[string][]PDFEstimator),
		lastLog:          time.Now(),
	}
	go s.mainLoop()
	return s
}

//ProcessTradesBlock publish trade to internal channel
func (s *EstimatorsService) ProcessTradesBlock(tradesBlock *dia.TradesBlock) {
	s.chanTradesBlock <- tradesBlock
	log.Info("completed ProcessTradesBlock")
}

//Close channel and terminate main loop
func (s *EstimatorsService) Close() error {
	if s.closed {
		return errors.New("Estimators: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	return s.error
}

// must only be called from mainLoop
func (s *EstimatorsService) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()
	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.shutdownDone) // signal that shutdown is complete
}

//CreateEstimator creates only selected estimators since some are really CPU intense process
func (s *EstimatorsService) CreateEstimator(estimator string, symbol string, exchange string) {
	//check if exist then append
	for _, e := range PDFEstimators() {
		if strings.EqualFold(e, estimator) {
			s.estimators[symbol+exchange] = append(s.estimators[symbol+exchange], NewPDFEstimator(estimator))
			return
		}
	}
}

//computeEstimators will run the estimators which can be cpu intensive
func (s *EstimatorsService) computeEstimators(key string) {
	for _, e := range s.estimators[key] {
		e.Compute()
	}
}

func (s *EstimatorsService) processTradesBlock(tb *dia.TradesBlock) {
	log.Infoln("processTradesBlock starting")
	for _, trade := range tb.TradesBlockData.Trades {
		if ea, ok := s.estimators[trade.Symbol+trade.Source]; ok == true {
			for _, e := range ea {
				e.AddSamples([]float64{trade.EstimatedUSDPrice})
			}
		}
		if ea, ok := s.estimators[trade.Symbol]; ok == true {
			for _, e := range ea {
				e.AddSamples([]float64{trade.EstimatedUSDPrice})
			}
		}
	}
	for _, v := range s.estimators {
		for _, e := range v {
			if err := e.Compute(); err == nil {
				s.chanPDFEstimates <- e.GetPDF()
			}
		}
	}
}

// runs in a goroutine until s is closed
func (s *EstimatorsService) mainLoop() {
	for {
		log.Info("Estimators Service mainloop")
		select {
		case <-s.shutdown:
			log.Println("Estimators Service shutting down")
			s.cleanup(nil)
			return
		case tb, _ := <-s.chanTradesBlock:
			s.processTradesBlock(tb)
		}
	}
}
