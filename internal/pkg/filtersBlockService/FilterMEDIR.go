// FilterMEDIR implements a trimmed median. Outliers are eliminated using interquartile range
// see: https://en.wikipedia.org/wiki/Interquartile_range
package filters

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

// FilterMEDIR contains the configuration parameters of the filter
type FilterMEDIR struct {
	symbol         string
	exchange       string
	currentTime    time.Time
	previousPrices []float64
	lastTrade      *dia.Trade
	memory         int
	value          float64
	filterName     string
}

//NewFilterMEDIR creates a FilterMEDIR
func NewFilterMEDIR(symbol string, exchange string, currentTime time.Time, memory int) *FilterMEDIR {
	s := &FilterMEDIR{
		symbol:         symbol,
		exchange:       exchange,
		previousPrices: []float64{},
		currentTime:    currentTime,
		memory:         memory,
		filterName:     "MEDIR" + strconv.Itoa(memory),
	}
	return s
}

func (s *FilterMEDIR) processDataPoint(price float64) {
	/// first remove extra value from buffer if already full
	if len(s.previousPrices) >= s.memory {
		s.previousPrices = s.previousPrices[0 : s.memory-1]
	}
	s.previousPrices = append([]float64{price}, s.previousPrices...)
}
func (s *FilterMEDIR) finalCompute(t time.Time) float64 {
	if s.lastTrade == nil {
		return 0.0
	}
	cleanPrices := removeOutliers(s.previousPrices)
	s.value = computeMedian(cleanPrices)
	s.previousPrices = []float64{}
	return s.value
}
func (s *FilterMEDIR) filterPointForBlock() *dia.FilterPoint {
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

func (s *FilterMEDIR) compute(trade dia.Trade) {
	if s.lastTrade != nil {
		if trade.Time.Before(s.currentTime) {
			log.Errorln("FilterMEDIR: Ignoring Trade out of order ", s.currentTime, trade.Time)
			return
		}
	}
	s.processDataPoint(trade.EstimatedUSDPrice)
	s.currentTime = trade.Time
	s.lastTrade = &trade
}

func (s *FilterMEDIR) save(ds models.Datastore) error {
	err := ds.SetFilter(s.filterName, s.symbol, s.exchange, s.value)
	if err != nil {
		log.Errorln("FilterMAIR: Error:", err)
	}
	return err
}
