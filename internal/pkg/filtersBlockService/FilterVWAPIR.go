package filters

import (
	"math"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	log "github.com/sirupsen/logrus"
)

// FilterVWAP ...
type FilterVWAPIR struct {
	exchange        string
	currentTime     time.Time
	previousPrices  []float64
	previousVolumes []float64
	lastTrade       *dia.Trade
	param           int
	value           float64
	modified        bool
	asset           dia.Asset
}

// NewFilterVWAP ...
func NewFilterVWAPIR(asset dia.Asset, exchange string, currentTime time.Time, param int) *FilterVWAPIR {
	s := &FilterVWAPIR{
		asset:          asset,
		exchange:       exchange,
		previousPrices: []float64{},
		currentTime:    currentTime,
		param:          param,
	}
	return s
}

// FinalCompute ...
func (s *FilterVWAPIR) FinalCompute(t time.Time) float64 {
	return s.finalCompute(t)
}

func (s *FilterVWAPIR) finalCompute(t time.Time) float64 {
	if s.lastTrade == nil {
		return 0.0
	}

	//else {
	// 	s.fill(t, s.lastTrade.EstimatedUSDPrice, s.lastTrade.Volume)
	// }

	// s.processDataPoint(*s.lastTrade)
	cleanPrices, bounds := removeOutliers(s.previousPrices)

	priceVolume := []float64{}

	// TODO handle bounds
	if len(bounds) == 0 {
		return 0.0
	}

	cleanedVolumes := s.previousVolumes[bounds[0]:bounds[1]]

	for index, price := range cleanPrices {
		priceVolume = append(priceVolume, price*math.Abs(cleanedVolumes[index]))
	}

	var total float64 = 0
	var totalVolume float64 = 0

	for _, v := range cleanedVolumes {
		totalVolume += v
	}

	for _, v := range priceVolume {
		total += v
	}

	s.value = total / totalVolume

	return s.value
}

// FilterPointForBlock ...
func (s *FilterVWAPIR) FilterPointForBlock() *dia.FilterPoint {
	return s.filterPointForBlock()
}
func (s *FilterVWAPIR) filterPointForBlock() *dia.FilterPoint {
	if s.exchange != "" {
		return &dia.FilterPoint{
			Value: s.value,
			Name:  "FilterVWAPIR" + strconv.Itoa(s.param),
			Time:  s.currentTime,
		}
	} else {
		return &dia.FilterPoint{
			Value: s.value,
			Name:  "FilterVWAPIR" + strconv.Itoa(s.param),
			Time:  s.currentTime,
		}
	}
}

func (s *FilterVWAPIR) fill(t time.Time, price float64, volume float64) {
	volume = math.Abs(volume)
	diff := int(t.Sub(s.currentTime).Seconds())
	if diff > 1 {
		for diff > 1 {
			s.previousPrices = append([]float64{price}, s.previousPrices...)
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
		s.previousPrices = append([]float64{price}, s.previousPrices...)
		s.previousVolumes = append([]float64{volume}, s.previousVolumes...)

	}

	if len(s.previousPrices) > s.param {
		s.previousPrices = s.previousPrices[0:s.param]
	}
	s.currentTime = t
}

// Compute ...
func (s *FilterVWAPIR) Compute(trade dia.Trade) {
	s.compute(trade)
}

func (s *FilterVWAPIR) compute(trade dia.Trade) {
	s.modified = true
	if s.lastTrade != nil {
		if trade.Time.Before(s.currentTime) {
			log.Errorln("FilterVWAPIR: Ignoring Trade out of order ", s.currentTime, trade.Time)
			return
		} else {
			s.fill(trade.Time, s.lastTrade.EstimatedUSDPrice, s.lastTrade.Volume)
		}
	}

	s.fill(trade.Time, trade.EstimatedUSDPrice, trade.Volume)
	s.lastTrade = &trade

}
