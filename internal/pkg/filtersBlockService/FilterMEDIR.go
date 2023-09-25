package filters

import (
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

// FilterMEDIR contains the configuration parameters of the filter.
// It implements a trimmed median. Outliers are eliminated using interquartile range
// see: https://en.wikipedia.org/wiki/Interquartile_range
type FilterMEDIR struct {
	asset              dia.Asset
	exchange           string
	currentTime        time.Time
	prices             []float64
	lastTrade          dia.Trade
	memory             int
	value              float64
	filterName         string
	modified           bool
	nativeDenomination bool
}

// NewFilterMEDIR creates a FilterMEDIR
func NewFilterMEDIR(asset dia.Asset, exchange string, currentTime time.Time, memory int, nativeDenomination bool) *FilterMEDIR {
	filter := &FilterMEDIR{
		asset:              asset,
		exchange:           exchange,
		prices:             []float64{},
		currentTime:        currentTime,
		memory:             memory,
		filterName:         "MEDIR" + strconv.Itoa(memory),
		nativeDenomination: nativeDenomination,
	}
	return filter
}

func (filter *FilterMEDIR) compute(trade dia.Trade) {
	filter.modified = true
	if filter.lastTrade != (dia.Trade{}) {
		if trade.Time.Before(filter.currentTime) {
			log.Errorln("FilterMEDIR: Ignoring Trade out of order ", filter.currentTime, trade.Time)
			return
		}
	}
	filter.processDataPoint(trade)
	filter.lastTrade = trade
	filter.currentTime = trade.Time
}

func (filter *FilterMEDIR) Compute(trade dia.Trade) {
	filter.compute(trade)
}
func (filter *FilterMEDIR) FinalCompute(t time.Time) {
	filter.finalCompute(t)
}

func (filter *FilterMEDIR) processDataPoint(trade dia.Trade) {
	/// first remove extra value from buffer if already full
	if len(filter.prices) >= filter.memory {
		filter.prices = filter.prices[0 : filter.memory-1]
	}
	if !filter.nativeDenomination {
		filter.prices = append([]float64{trade.EstimatedUSDPrice}, filter.prices...)
	} else {
		filter.prices = append([]float64{trade.Price}, filter.prices...)
	}
}

func (filter *FilterMEDIR) finalCompute(t time.Time) float64 {
	if filter.lastTrade == (dia.Trade{}) {
		log.Info("last trade emtpy")
		return 0.0
	}
	cleanPrices, _ := removeOutliers(filter.prices)
	filter.value = computeMedian(cleanPrices)
	filter.prices = []float64{filter.lastTrade.EstimatedUSDPrice}
	return filter.value
}

func (filter *FilterMEDIR) filterPointForBlock() *dia.FilterPoint {
	if filter.exchange != "" || filter.filterName != dia.FilterKing {
		return nil
	}
	return &dia.FilterPoint{
		Asset: filter.asset,
		Value: filter.value,
		Name:  filter.filterName,
		Time:  filter.currentTime,
	}
}

func (filter *FilterMEDIR) FilterPointForBlock() *dia.FilterPoint {
	return &dia.FilterPoint{
		Asset: filter.asset,
		Value: filter.value,
		Name:  filter.filterName,
		Time:  filter.currentTime,
	}
}
func (filter *FilterMEDIR) save(ds models.Datastore) error {
	if filter.modified {
		filter.modified = false
		err := ds.SetFilter(filter.filterName, filter.asset, filter.exchange, filter.value, filter.currentTime)
		if err != nil {
			log.Errorln("FilterMEDIR: Error:", err)
		}
		return err
	} else {
		return nil
	}
}
