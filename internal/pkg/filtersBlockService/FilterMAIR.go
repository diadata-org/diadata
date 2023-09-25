package filters

import (
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

// FilterMAIR implements a trimmed moving average.
// Outliers are eliminated using interquartile range.
// see: https://en.wikipedia.org/wiki/Interquartile_range
type FilterMAIR struct {
	asset              dia.Asset
	exchange           string
	currentTime        time.Time
	prices             []float64
	volumes            []float64
	lastTrade          dia.Trade
	memory             int
	value              float64
	filterName         string
	nativeDenomination bool
	modified           bool
}

// NewFilterMAIR returns a FilterMAIR
func NewFilterMAIR(asset dia.Asset, exchange string, currentTime time.Time, memory int, nativeDenomination bool) *FilterMAIR {
	filter := &FilterMAIR{
		asset:              asset,
		exchange:           exchange,
		prices:             []float64{},
		volumes:            []float64{},
		currentTime:        currentTime,
		memory:             memory,
		filterName:         "MAIR" + strconv.Itoa(memory),
		nativeDenomination: nativeDenomination,
	}
	return filter
}

func (filter *FilterMAIR) Compute(trade dia.Trade) {
	filter.compute(trade)
}

func (filter *FilterMAIR) compute(trade dia.Trade) {
	filter.modified = true
	if filter.lastTrade != (dia.Trade{}) {
		if trade.Time.Before(filter.currentTime) {
			log.Errorln("FilterMAIR: Ignoring Trade out of order ", filter.currentTime, trade.Time)
			return
		}
	}
	filter.fill(trade)
	filter.lastTrade = trade
}

// fill fills up the 120 seconds slots with trades.
func (filter *FilterMAIR) fill(trade dia.Trade) {
	// filter.currentTime is the timestamp of the last filled trade.
	// It is initialized with begin time of tradesblock upon creation of the filter.
	diff := int(trade.Time.Sub(filter.currentTime).Seconds())
	if diff > 1 {
		for diff > 1 {
			filter.processDataPoint(trade)
			diff--
		}
	} else {
		if diff == 0.0 {
			if len(filter.prices) >= 1 {
				/// Remove latest data point and update with newer
				filter.prices = filter.prices[1:]
				filter.volumes = filter.volumes[1:]
			}
		}
		filter.processDataPoint(trade)
	}
	filter.currentTime = trade.Time
}

func (filter *FilterMAIR) processDataPoint(trade dia.Trade) {
	/// first remove extra value from buffer if already full
	if len(filter.prices) >= filter.memory {
		filter.prices = filter.prices[0 : filter.memory-1]
		filter.volumes = filter.volumes[0 : filter.memory-1]
	}
	if !filter.nativeDenomination {
		filter.prices = append([]float64{trade.EstimatedUSDPrice}, filter.prices...)
	} else {
		filter.prices = append([]float64{trade.Price}, filter.prices...)
	}

	filter.volumes = append([]float64{trade.Volume}, filter.volumes...)
}

func (filter *FilterMAIR) FinalCompute(t time.Time) float64 {
	return filter.finalCompute(t)
}

func (filter *FilterMAIR) finalCompute(t time.Time) float64 {
	if filter.lastTrade == (dia.Trade{}) {
		return 0.0
	}

	if len(filter.prices) < 2 {
		filter.value = filter.prices[0]
		return filter.prices[0]
	}

	// Add the last trade again to compensate for the delay since measurement to EOB
	// adopted behaviour from FilterMA
	filter.processDataPoint(filter.lastTrade)
	cleanPrices, bounds := removeOutliers(filter.prices)
	mean, err := computeMean(cleanPrices, filter.volumes[bounds[0]:bounds[1]])
	if err != nil {
		return 0.0
	}
	filter.value = mean
	// Reduce the filter values to the last recorded value for the next tradesblock.
	if len(filter.prices) > 0 && len(filter.volumes) > 0 {
		if !filter.nativeDenomination {
			filter.prices = []float64{filter.lastTrade.EstimatedUSDPrice}
		} else {
			filter.prices = []float64{filter.lastTrade.Price}
		}
		filter.volumes = []float64{filter.lastTrade.Volume}
	}
	return filter.value
}

func (filter *FilterMAIR) FilterPointForBlock() *dia.FilterPoint {
	return &dia.FilterPoint{
		Asset: filter.asset,
		Value: filter.value,
		Name:  filter.filterName,
		Time:  filter.currentTime,
	}
}

func (filter *FilterMAIR) filterPointForBlock() *dia.FilterPoint {
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

func (filter *FilterMAIR) save(ds models.Datastore) error {
	if filter.modified {
		filter.modified = false
		err := ds.SetFilter(filter.filterName, filter.asset, filter.exchange, filter.value, filter.currentTime)
		if err != nil {
			log.Errorln("FilterMAIR: Error:", err)
		}

		// Additionally, the price across exchanges is saved in influx as a quotation.
		// This price is used for the estimation of quote tokens' prices in the tradesBlockService.
		if filter.exchange == "" {
			err = ds.SetAssetPriceUSD(filter.asset, filter.value, filter.currentTime)
			if err != nil {
				log.Errorln("FilterMAIR: Error:", err)
			}
		}
		return err
	}
	return nil
}
