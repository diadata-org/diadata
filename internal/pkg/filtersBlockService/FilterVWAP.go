package filters

import (
	"math"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	log "github.com/sirupsen/logrus"
)

// FilterVWAP ...
type FilterVWAP struct {
	asset           dia.Asset
	exchange        string
	currentTime     time.Time
	previousPrices  []float64
	previousVolumes []float64
	lastTrade       *dia.Trade
	param           int
	value           float64
	modified        bool
}

// NewFilterVWAP ...
func NewFilterVWAP(asset dia.Asset, exchange string, currentTime time.Time, param int) *FilterVWAP {
	s := &FilterVWAP{
		asset:          asset,
		exchange:       exchange,
		previousPrices: []float64{},
		currentTime:    currentTime,
		param:          param,
	}
	return s
}

// FinalCompute ...
func (s *FilterVWAP) FinalCompute(t time.Time) float64 {
	return s.finalCompute(t)
}

func (s *FilterVWAP) finalCompute(t time.Time) float64 {
	// if s.lastTrade == nil {
	// 	return 0.0
	// } else {
	// 	s.fill(t, s.lastTrade.EstimatedUSDPrice, s.lastTrade.Volume)
	// }

	var total float64 = 0
	var totalVolume float64 = 0

	for _, v := range s.previousVolumes {
		totalVolume += v
	}

	for _, v := range s.previousPrices {
		total += v
	}

	s.value = total / totalVolume

	return s.value
}

// FilterPointForBlock ...
func (s *FilterVWAP) FilterPointForBlock() *dia.FilterPoint {
	return s.filterPointForBlock()
}
func (s *FilterVWAP) filterPointForBlock() *dia.FilterPoint {
	if s.exchange != "" {
		return &dia.FilterPoint{
			Value: s.value,
			Name:  "VWAP" + strconv.Itoa(s.param),
			Time:  s.currentTime,
			Asset: s.asset,
		}
	} else {
		return &dia.FilterPoint{
			Value: s.value,
			Name:  "VWAP" + strconv.Itoa(s.param),
			Time:  s.currentTime,
			Asset: s.asset,
		}
	}
}

func (s *FilterVWAP) fill(t time.Time, price float64, volume float64) {
	volume = math.Abs(volume)
	diff := int(t.Sub(s.currentTime).Seconds())
	if diff > 1 {
		for diff > 1 {
			s.previousPrices = append([]float64{price * volume}, s.previousPrices...)
			s.previousVolumes = append([]float64{volume}, s.previousVolumes...)

			diff--
		}
	} else {

		if diff == 0.0 {
			if len(s.previousPrices) >= 1 {
				s.previousPrices = s.previousPrices[1:]
				s.previousVolumes = s.previousVolumes[1:]
			}
		}
		s.previousPrices = append([]float64{price * volume}, s.previousPrices...)
		s.previousVolumes = append([]float64{volume}, s.previousVolumes...)

	}

	if len(s.previousPrices) > s.param {
		s.previousPrices = s.previousPrices[0:s.param]
	}
	s.currentTime = t
}

// Compute ...
func (s *FilterVWAP) Compute(trade dia.Trade) {
	s.compute(trade)
}

func (s *FilterVWAP) compute(trade dia.Trade) {
	s.modified = true
	if s.lastTrade != nil {
		if trade.Time.Before(s.currentTime) {
			log.Errorln("FilterVWAP: Ignoring Trade out of order ", s.currentTime, trade.Time)
			return
		} else {
			s.fill(trade.Time, s.lastTrade.EstimatedUSDPrice, s.lastTrade.Volume)
		}
	}

	s.fill(trade.Time, trade.EstimatedUSDPrice, trade.Volume)
	s.lastTrade = &trade

}
