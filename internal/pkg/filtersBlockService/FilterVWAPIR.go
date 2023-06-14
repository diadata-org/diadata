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
	exchange    string
	currentTime time.Time
	prices      []float64
	volumes     []float64
	lastTrade   dia.Trade
	param       int
	value       float64
	modified    bool
	filterName  string
	asset       dia.Asset
}

// NewFilterVWAP ...
func NewFilterVWAPIR(asset dia.Asset, exchange string, currentTime time.Time, param int) *FilterVWAPIR {
	s := &FilterVWAPIR{
		asset:       asset,
		exchange:    exchange,
		prices:      []float64{},
		volumes:     []float64{},
		currentTime: currentTime,
		param:       param,
		filterName:  "VWAPIR" + strconv.Itoa(param),
	}
	return s
}

// Compute ...
func (s *FilterVWAPIR) Compute(trade dia.Trade) {
	s.compute(trade)
}

func (filter *FilterVWAPIR) compute(trade dia.Trade) {
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
func (filter *FilterVWAPIR) fill(trade dia.Trade) {
	// filter.currentTime is the timestamp of the last filled trade.
	filter.processDataPoint(trade)
	filter.currentTime = trade.Time
}

func (filter *FilterVWAPIR) processDataPoint(trade dia.Trade) {
	filter.prices = append([]float64{trade.EstimatedUSDPrice}, filter.prices...)
	filter.volumes = append([]float64{trade.Volume}, filter.volumes...)
}

// FinalCompute ...
func (s *FilterVWAPIR) FinalCompute(t time.Time) float64 {
	log.Info("final compute of time ", t)
	return s.finalCompute()
}

func (s *FilterVWAPIR) finalCompute() float64 {
	if s.lastTrade == (dia.Trade{}) {
		return 0.0
	}

	// s.processDataPoint(*s.lastTrade)
	cleanPrices, bounds := removeOutliers(s.prices)

	priceVolume := []float64{}

	// TODO handle bounds
	if len(bounds) == 0 {
		return 0.0
	}

	cleanedVolumes := s.volumes[bounds[0]:bounds[1]]

	for index, price := range cleanPrices {
		priceVolume = append(priceVolume, price*math.Abs(cleanedVolumes[index]))
	}

	var total float64 = 0
	var totalVolume float64 = 0

	for _, v := range cleanedVolumes {
		totalVolume += math.Abs(v)
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
			Name:  s.filterName,
			Time:  s.currentTime,
			Asset: s.asset,
		}
	} else {
		return &dia.FilterPoint{
			Value: s.value,
			Name:  s.filterName,
			Time:  s.currentTime,
			Asset: s.asset,
		}
	}
}
