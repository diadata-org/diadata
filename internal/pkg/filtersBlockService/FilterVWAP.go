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
	asset              dia.Asset
	exchange           string
	currentTime        time.Time
	prices             []float64
	volumes            []float64
	lastTrade          dia.Trade
	param              int
	value              float64
	filterName         string
	nativeDenomination bool
	modified           bool
}

// NewFilterVWAP ...
func NewFilterVWAP(asset dia.Asset, exchange string, currentTime time.Time, param int, nativeDenomination bool) *FilterVWAP {
	s := &FilterVWAP{
		asset:              asset,
		exchange:           exchange,
		prices:             []float64{},
		volumes:            []float64{},
		currentTime:        currentTime,
		param:              param,
		filterName:         "VWAP" + strconv.Itoa(param),
		nativeDenomination: nativeDenomination,
	}
	return s
}

// Compute ...
func (s *FilterVWAP) Compute(trade dia.Trade) {
	s.compute(trade)
}

func (filter *FilterVWAP) compute(trade dia.Trade) {
	filter.modified = true
	if filter.lastTrade != (dia.Trade{}) {
		if trade.Time.Before(filter.currentTime) {
			log.Errorln("FilterVWAP: Ignoring Trade out of order ", filter.currentTime, trade.Time)
			return
		}
	}
	filter.fill(trade)
	filter.lastTrade = trade
}

// fill just adds a trade to the prices and volumes slices.
func (filter *FilterVWAP) fill(trade dia.Trade) {
	// filter.currentTime is the timestamp of the last filled trade.
	filter.processDataPoint(trade)
	filter.currentTime = trade.Time
}

func (filter *FilterVWAP) processDataPoint(trade dia.Trade) {
	if !filter.nativeDenomination {
		filter.prices = append([]float64{trade.EstimatedUSDPrice}, filter.prices...)
	} else {
		filter.prices = append([]float64{trade.Price}, filter.prices...)
	}
	filter.volumes = append([]float64{trade.Volume}, filter.volumes...)
}

// FinalCompute ...
func (s *FilterVWAP) FinalCompute(t time.Time) float64 {
	return s.finalCompute(t)
}

func (s *FilterVWAP) finalCompute(t time.Time) float64 {
	log.Infof("computed value %v at time %v ", s.value, t)
	if s.lastTrade == (dia.Trade{}) {
		return 0.0
	}

	var totalVolume float64 = 0
	var priceVolume []float64
	var totalPriceVolume float64 = 0

	for index, price := range s.prices {
		priceVolume = append(priceVolume, price*math.Abs(s.volumes[index]))
	}

	for _, v := range s.volumes {
		totalVolume += math.Abs(v)
	}

	for _, v := range priceVolume {
		totalPriceVolume += v
	}

	if totalVolume > 0 {
		s.value = totalPriceVolume / totalVolume
	}

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
