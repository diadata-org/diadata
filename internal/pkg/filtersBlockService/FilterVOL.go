package filters

import (
	"math"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

type FilterVOL struct {
	asset       dia.Asset
	exchange    string
	currentTime time.Time
	volumeUSD   float64
	lastTrade   *dia.Trade
	value       float64
	filterName  string
	memory      int
}

func NewFilterVOL(asset dia.Asset, exchange string, memory int) *FilterVOL {
	s := &FilterVOL{
		asset:      asset,
		exchange:   exchange,
		volumeUSD:  0.0,
		filterName: "VOL" + strconv.Itoa(memory),
		memory:     memory,
	}
	return s
}

func (s *FilterVOL) compute(trade dia.Trade) {
	s.volumeUSD += trade.EstimatedUSDPrice * math.Abs(trade.Volume)
	s.currentTime = trade.Time
}

func (s *FilterVOL) finalCompute(time time.Time) float64 {
	s.value = s.volumeUSD
	s.volumeUSD = 0.0
	return s.value
}

func (s *FilterVOL) filterPointForBlock() *dia.FilterPoint {
	return nil
}

func (s *FilterVOL) save(ds models.Datastore) error {
	err := ds.SetFilter(s.filterName, s.asset.Symbol, s.exchange, s.value, s.currentTime)
	if err != nil {
		log.Errorln("FilterVOL Error:", err)
	}
	return err
}
