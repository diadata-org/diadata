// FilterMAIR implements a trimmed moving average. Outliers are eliminated using interquartile range
// see: https://en.wikipedia.org/wiki/Interquartile_range
package filters

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
	"math"
	"sort"
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
}

//NewFilterMAIR creates a FilterMAIR
func NewFilterMAIR(symbol string, exchange string, currentTime time.Time, memory int) *FilterMAIR {
	s := &FilterMAIR{
		symbol:         symbol,
		exchange:       exchange,
		previousPrices: []float64{},
		currentTime:    currentTime,
		memory:         memory,
	}
	return s
}
func computeMean(samples []float64) (mean float64) {
	var total float64
	length := float64(len(samples))
	if length == 0 {
		return
	}
	for _, s := range samples {
		total += s
	}
	mean = total / length
	return
}
func computeMedian(samples []float64) (median float64) {
	var length = len(samples)
	if length > 0 {
		sort.Float64s(samples)
		if length%2 == 0 {
			median = (samples[length/2-1] + samples[length/2]) / 2
		} else {
			median = samples[(length+1)/2-1]
		}
	}
	return
}
func computeQuartiles(samples []float64) (Q1 float64, Q3 float64) {
	sort.Float64s(samples)
	var length = len(samples)
	if length > 0 {
		if length%2 == 0 {
			Q1 = computeMedian(samples[0 : length/2])
			Q3 = computeMedian(samples[length/2 : length])
		} else {
			Q1 = computeMedian(samples[0:int(math.Floor(float64(length/2)))])
			Q3 = computeMedian(samples[int(math.Floor(float64(length/2)))+1 : length])
		}
	}
	return
}

// RemoveOutliers Cleans a data set it accordance to the acceptable range within interquartile range.
func removeOutliers(samples []float64) []float64 {
	if len(samples) == 0 || len(samples) == 1 {
		return samples
	}
	Q1, Q3 := computeQuartiles(samples)
	IQR := Q3 - Q1
	lowerBound := Q1 - 1.5*IQR
	upperBound := Q3 + 1.5*IQR
	lowerIndex := 0
	upperIndex := len(samples)
	for index, value := range samples {
		if value < lowerBound {
			lowerIndex = index + 1
		} else if value > upperBound {
			upperIndex = index
			break
		}
	}
	return samples[lowerIndex:upperIndex]
}
func (s *FilterMAIR) processDataPoint(price float64) {
	/// first remove extra value from buffer if already full
	if len(s.previousPrices) >= s.memory {
		s.previousPrices = s.previousPrices[0 : s.memory-1]
	}
	s.previousPrices = append([]float64{price}, s.previousPrices...)
}
func (s *FilterMAIR) finalComputeEndOfBlock(t time.Time) {
	if s.lastTrade == nil {
		return
	}
	// Add the last trade again to compensate for the delay since measurement to EOB
	// adopted behaviour from FilterMA
	s.processDataPoint(s.lastTrade.EstimatedUSDPrice)
	cleanPrices := removeOutliers(s.previousPrices)
	s.value = computeMean(cleanPrices)
}
func (s *FilterMAIR) filterPointForBlock() *dia.FilterPoint {
	if s.exchange != "" {
		return nil
	}
	return &dia.FilterPoint{
		Symbol: s.symbol,
		Value:  s.value,
		Name:   "MAIR" + strconv.Itoa(s.memory),
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
	err := ds.SetPriceZSET(s.symbol, s.exchange, s.value)
	if err != nil {
		log.Errorln("FilterMAIR: Error:", err)
	}
	if s.exchange == "" {
		err = ds.SetPriceUSD(s.symbol, s.value)
		if err != nil {
			log.Errorln("FilterMAIR: Error:", err)
		}
	}
	return err
}
