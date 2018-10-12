package filters

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
	"math"
	"time"
)

type FilterVOL struct {
	symbol      string
	exchange    string
	currentTime time.Time
	volumeUSD   float64
	lastTrade   *dia.Trade
	value       float64
}

func NewFilterVOL(symbol string, exchange string) *FilterVOL {
	s := &FilterVOL{
		symbol:    symbol,
		exchange:  exchange,
		volumeUSD: 0.0,
	}
	return s
}

func (s *FilterVOL) finalComputeEndOfBlock(time time.Time) {
	s.value = s.volumeUSD
	s.volumeUSD = 0.0
}

func (s *FilterVOL) filterPointForBlock() *dia.FilterPoint {
	return nil
}

func (s *FilterVOL) compute(trade dia.Trade) {
	s.volumeUSD += trade.EstimatedUSDPrice * math.Abs(trade.Volume)
}

func (s *FilterVOL) save(ds models.Datastore) error {
	err := ds.SetVolume(s.symbol, s.exchange, s.value)
	if err != nil {
		log.Errorln("FilterVOL Error:", err)
	}
	return err
}
