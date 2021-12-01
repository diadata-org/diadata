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
	asset           dia.Asset
	exchange        string
	currentTime     time.Time
	previousPrices  []float64
	previousVolumes []float64
	lastTrade       *dia.Trade
	param           int
	value           float64
	modified        bool
	filterName      string
}

// NewFilterMA returns a moving average filter.
// @currentTime is the begin time of the filtersBlock.
func NewFilterMA(asset dia.Asset, exchange string, currentTime time.Time, param int) *FilterMA {
	s := &FilterMA{
		asset:           asset,
		exchange:        exchange,
		previousPrices:  []float64{},
		previousVolumes: []float64{},
		currentTime:     currentTime,
		param:           param,
		filterName:      "MA" + strconv.Itoa(param),
	}
	return s
}

func (s *FilterMA) Compute(trade dia.Trade) {
	s.compute(trade)
}

func (s *FilterMA) compute(trade dia.Trade) {
	s.modified = true
	if s.lastTrade != nil {
		if trade.Time.Before(s.currentTime) {
			log.Errorln("FilterMA: Ignoring Trade out of order ", s.currentTime, trade.Time)
			return
		}
		s.fill(trade.Time, *s.lastTrade)
	}
	s.fill(trade.Time, trade)
	s.lastTrade = &trade
}

func (s *FilterMA) fill(t time.Time, trade dia.Trade) {
	// Time difference of trade time and @currentTime in seconds.
	// Initially, @currentTime is the begin time of the filtersBlock.
	diff := int(t.Sub(s.currentTime).Seconds())
	// log.Infof("diff for asset %s: %v ", s.asset.Symbol, diff)
	currPrice := trade.EstimatedUSDPrice
	currVolume := trade.Volume
	if diff > 1 {

		for diff > 1 {
			s.previousPrices = append([]float64{currPrice}, s.previousPrices...)
			s.previousVolumes = append([]float64{currVolume}, s.previousVolumes...)
			diff--
		}
	} else {
		// If timestamp @t is the same as @trade.Time, shift slices and append @trade.
		if diff == 0.0 {
			if len(s.previousPrices) >= 1 {
				s.previousPrices = s.previousPrices[1:]
				s.previousVolumes = s.previousVolumes[1:]
			}
		}
		s.previousPrices = append([]float64{currPrice}, s.previousPrices...)
		s.previousVolumes = append([]float64{currVolume}, s.previousVolumes...)
	}

	if len(s.previousPrices) > s.param {
		s.previousPrices = s.previousPrices[0:s.param]
	}
	if len(s.previousVolumes) > s.param {
		s.previousVolumes = s.previousVolumes[0:s.param]
	}
	s.currentTime = t
}

func (s *FilterMA) FinalCompute(t time.Time) float64 {
	return s.finalCompute(t)
}

func (s *FilterMA) finalCompute(t time.Time) float64 {
	if s.lastTrade == nil {
		return 0.0
	}
	s.fill(t, *s.lastTrade)
	var totalVolume float64
	var totalPrice float64
	for priceIndex, price := range s.previousPrices {
		totalPrice += price * math.Abs(s.previousVolumes[priceIndex])
		totalVolume += math.Abs(s.previousVolumes[priceIndex])
	}
	if s.asset.Symbol == "USDT" || s.asset.Symbol == "USDC" {
		// log.Infof("total price for %s on %s: %v", s.asset.Symbol, s.exchange, totalPrice)
		// log.Infof("total Volume for %s on %s: %v", s.asset.Symbol, s.exchange, totalVolume)
		// log.Infof("resulting price for %s on %s: %v", s.asset.Symbol, s.exchange, totalPrice/totalVolume)
		var nonweightedPrice float64
		for _, price := range s.previousPrices {
			nonweightedPrice += price
		}
		// log.Infof("average on non-volume-weighted prices for %s on %s: %v", s.asset.Symbol, s.exchange, nonweightedPrice/float64(len(s.previousPrices)))
		// log.Info("prices in filtersblock: ", s.previousPrices)
		// log.Info("volumes in filtersblock: ", s.previousVolumes)
		// log.Info("-------------------------------------------------------------------------")
	}
	s.value = totalPrice / totalVolume
	return s.value
}

func (s *FilterMA) FilterPointForBlock() *dia.FilterPoint {
	return s.filterPointForBlock()
}

func (s *FilterMA) filterPointForBlock() *dia.FilterPoint {
	if s.exchange != "" || s.filterName != dia.FilterKing {
		return nil
	}
	return &dia.FilterPoint{
		Asset: s.asset,
		Value: s.value,
		Name:  "MA" + strconv.Itoa(s.param),
		Time:  s.currentTime,
	}
}

func (s *FilterMA) save(ds models.Datastore) error {

	if s.modified {
		s.modified = false
		err := ds.SetFilter(s.filterName, s.asset, s.exchange, s.value, s.currentTime)
		if err != nil {
			log.Errorln("FilterMA: Error:", err)
		}
		return err
	}
	return nil
}
