package volumeService

import (
	"errors"
	"github.com/diadata-org/api-golang/dia"
	log "github.com/sirupsen/logrus"
	"math"
	"sync"
)

type nothing struct{}

type VolumeService struct {
	shutdown        chan nothing
	shutdownDone    chan nothing
	chanTradesBlock chan *dia.TradesBlock
	errorLock       sync.RWMutex
	error           error
	closed          bool
}

func NewVolumeService() *VolumeService {
	s := &VolumeService{
		shutdown:        make(chan nothing),
		shutdownDone:    make(chan nothing),
		chanTradesBlock: make(chan *dia.TradesBlock),
		error:           nil,
	}
	go s.mainLoop()
	return s
}

func (s *VolumeService) ProcessTradesBlock(tradesBlock *dia.TradesBlock) {
	s.chanTradesBlock <- tradesBlock
}

func (s *VolumeService) Close() error {
	if s.closed {
		return errors.New("VolumeService: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	return s.error
}

// must only be called from mainLoop
func (s *VolumeService) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()
	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.shutdownDone) // signal that shutdown is complete
}

func (s *VolumeService) processTradesBlock(tb *dia.TradesBlock) {
	var volumes = make(map[string]float64)
	log.Info("processTradesBlock")
	for _, trade := range tb.TradesBlockData.Trades {
		f, ok := volumes[trade.Symbol]
		if !ok {
			volumes[trade.Symbol] = 0.0
		}
		volumes[trade.Symbol] = f + trade.EstimatedUSDPrice*math.Abs(trade.Volume)
	}
}

// runs in a goroutine until s is closed
func (s *VolumeService) mainLoop() {
	for {
		select {
		case <-s.shutdown:
			log.Println("VolumeService shutting down")
			s.cleanup(nil)
			return
		case tb, _ := <-s.chanTradesBlock:
			s.processTradesBlock(tb)
		}
	}
}
