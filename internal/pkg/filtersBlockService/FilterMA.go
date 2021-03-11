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
	log.Infof("diff for asset %s: %v ", s.asset.Symbol, diff)
	currPrice := trade.EstimatedUSDPrice
	currVolume := trade.Volume
	if diff > 1 {
		for diff > 1 {
			s.previousPrices = append([]float64{currPrice}, s.previousPrices...)
			s.previousVolumes = append([]float64{currVolume}, s.previousVolumes...)
			diff--
		}
	} else {
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
	s.value = totalPrice / totalVolume
	return s.value
}

func (s *FilterMA) filterPointForBlock() *dia.FilterPoint {
	if s.exchange != "" {
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
		log.Infof("set price for %s: %v", s.asset.Symbol, s.value)
		if err != nil {
			log.Errorln("FilterMA: Error:", err)
		}
		// Additionally, the price across exchanges is saved in influx as a quotation.
		// This price is used for the estimation of quote tokens' prices in the tradesBlockService.
		if s.exchange == "" {
			err = ds.SetAssetPriceUSD(s.asset, s.value, s.currentTime)
			if err != nil {
				log.Errorln("FilterMA: Error:", err)
			}
		}
		return err
	}
	return nil
}
