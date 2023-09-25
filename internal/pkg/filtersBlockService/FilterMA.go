package filters

import (
	"math"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

// FilterMA is the struct for a moving average filter implementing
// the Filter interface.
type FilterMA struct {
	asset              dia.Asset
	exchange           string
	currentTime        time.Time
	prices             []float64
	volumes            []float64
	lastTrade          dia.Trade
	memory             int
	value              float64
	modified           bool
	filterName         string
	nativeDenomination bool
	//max         float64
	min float64
}

// NewFilterMA returns a moving average filter.
// @currentTime is the begin time of the filtersBlock.
func NewFilterMA(asset dia.Asset, exchange string, currentTime time.Time, memory int, nativeDenomination bool) *FilterMA {
	filter := &FilterMA{
		asset:              asset,
		exchange:           exchange,
		prices:             []float64{},
		volumes:            []float64{},
		currentTime:        currentTime,
		memory:             memory,
		filterName:         "MA" + strconv.Itoa(memory),
		min:                -1,
		nativeDenomination: nativeDenomination,
	}
	return filter
}

func (filter *FilterMA) Compute(trade dia.Trade) {
	filter.compute(trade)
}

func (filter *FilterMA) compute(trade dia.Trade) {
	filter.modified = true
	if filter.lastTrade != (dia.Trade{}) {
		if trade.Time.Before(filter.currentTime) {
			log.Errorln("FilterMA: Ignoring Trade out of order ", filter.currentTime, trade.Time)
			return
		}
	}
	filter.fill(trade)
	filter.lastTrade = trade
}

// fill fills up the 120 seconds slots with trades.
func (filter *FilterMA) fill(trade dia.Trade) {
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

func (filter *FilterMA) processDataPoint(trade dia.Trade) {
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

func (filter *FilterMA) FinalCompute(t time.Time) float64 {
	return filter.finalCompute(t)
}

func (filter *FilterMA) finalCompute(t time.Time) float64 {
	if filter.lastTrade == (dia.Trade{}) {
		return 0.0
	}
	filter.fill(filter.lastTrade)
	var totalVolume float64
	var totalPrice float64
	for priceIndex, price := range filter.prices {
		totalPrice += price * math.Abs(filter.volumes[priceIndex])
		totalVolume += math.Abs(filter.volumes[priceIndex])
	}
	if filter.asset.Symbol == "USDT" || filter.asset.Symbol == "USDC" {
		var nonweightedPrice float64
		for _, price := range filter.prices {
			nonweightedPrice += price
		}
	}
	if totalVolume > 0 {
		filter.value = totalPrice / totalVolume

	}
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

func (filter *FilterMA) FilterPointForBlock() *dia.FilterPoint {
	return &dia.FilterPoint{
		Asset: filter.asset,
		Value: filter.value,
		Name:  "MA" + strconv.Itoa(filter.memory),
		Time:  filter.currentTime,
	}
}

func (filter *FilterMA) filterPointForBlock() *dia.FilterPoint {
	if filter.exchange != "" || filter.filterName != dia.FilterKing {
		return nil
	}
	return &dia.FilterPoint{
		Asset: filter.asset,
		Value: filter.value,
		Name:  "MA" + strconv.Itoa(filter.memory),
		Time:  filter.currentTime,
	}
}

func (filter *FilterMA) save(ds models.Datastore) error {
	if filter.modified {
		filter.modified = false
		err := ds.SetFilter(filter.filterName, filter.asset, filter.exchange, filter.value, filter.currentTime)
		if err != nil {
			log.Errorln("FilterMA: Error:", err)
		}
		return err
	}
	return nil
}
