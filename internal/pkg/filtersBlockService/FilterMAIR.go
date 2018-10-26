// FilterMAIR implements a trimmed moving average. Outliers are eliminated using interquartile range
// see: https://en.wikipedia.org/wiki/Interquartile_range
package filters

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

// FilterMAIR contains the configuration parameters of the filter
type FilterMAIR struct {
	symbol         string
	exchange       string
	currentTime    time.Time
	previousPrices []float64
	lastTrade      *dia.Trade
	memory         int
	value          float64
	filterName     string
	modified       bool
}

//NewFilterMAIR creates a FilterMAIR
func NewFilterMAIR(symbol string, exchange string, currentTime time.Time, memory int) *FilterMAIR {
	s := &FilterMAIR{
		symbol:         symbol,
		exchange:       exchange,
		previousPrices: []float64{},
		currentTime:    currentTime,
		memory:         memory,
		filterName:     "MAIR" + strconv.Itoa(memory),
	}
	return s
}

func (s *FilterMAIR) processDataPoint(price float64) {
	/// first remove extra value from buffer if already full
	if len(s.previousPrices) >= s.memory {
		s.previousPrices = s.previousPrices[0 : s.memory-1]
	}
	s.previousPrices = append([]float64{price}, s.previousPrices...)
}
func (s *FilterMAIR) finalCompute(t time.Time) float64 {
	if s.lastTrade == nil {
		return 0.0
	}
	// Add the last trade again to compensate for the delay since measurement to EOB
	// adopted behaviour from FilterMA
	s.processDataPoint(s.lastTrade.EstimatedUSDPrice)
	cleanPrices := removeOutliers(s.previousPrices)
	s.value = computeMean(cleanPrices)
	return s.value
}
func (s *FilterMAIR) filterPointForBlock() *dia.FilterPoint {
	if s.exchange != "" || s.filterName != dia.FilterKing {
		return nil
	}
	return &dia.FilterPoint{
		Symbol: s.symbol,
		Value:  s.value,
		Name:   s.filterName,
		Time:   s.currentTime,
	}
}
func (s *FilterMAIR) fill(t time.Time, price float64) {
	diff := int(t.Sub(s.currentTime).Seconds())
	if diff > 1 {
		for diff > 1 {
			s.processDataPoint(price)
			diff--
		}
	} else {
		if diff == 0.0 {
			if len(s.previousPrices) >= 1 {
				/// Remove latest data point and update with newer
				s.previousPrices = s.previousPrices[1:]
			}
		}
		s.processDataPoint(price)
	}
	s.currentTime = t
}
func (s *FilterMAIR) compute(trade dia.Trade) {
	s.modified = true
	if s.lastTrade != nil {
		if trade.Time.Before(s.currentTime) {
			log.Errorln("FilterMAIR: Ignoring Trade out of order ", s.currentTime, trade.Time)
			return
		}
	}
	s.fill(trade.Time, trade.EstimatedUSDPrice)
	s.lastTrade = &trade
}

func (s *FilterMAIR) save(ds models.Datastore) error {
	if s.modified {
		s.modified = false
		err := ds.SetFilter(s.filterName, s.symbol, s.exchange, s.value, s.currentTime)
		if err != nil {
			log.Errorln("FilterMAIR: Error:", err)
		}
		return err
	} else {
		return nil
	}
}
