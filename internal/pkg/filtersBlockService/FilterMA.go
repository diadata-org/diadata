package filters

import (
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

// FilterMA is the struct for a moving average filter implementing
// the Filter interface.
type FilterMA struct {
	asset          dia.Asset
	exchange       string
	currentTime    time.Time
	previousPrices []float64
	lastTrade      *dia.Trade
	param          int
	value          float64
	modified       bool
}

// NewFilterMA returns a moving average filter.
// @currentTime is the begin time of the filtersBlock.
func NewFilterMA(asset dia.Asset, exchange string, currentTime time.Time, param int) *FilterMA {
	s := &FilterMA{
		asset:          asset,
		exchange:       exchange,
		previousPrices: []float64{},
		currentTime:    currentTime,
		param:          param,
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
		s.fill(trade.Time, s.lastTrade.EstimatedUSDPrice)
	}
	s.fill(trade.Time, trade.EstimatedUSDPrice)
	s.lastTrade = &trade
}

func (s *FilterMA) fill(t time.Time, price float64) {
	// Time difference of trade time and @currentTime in seconds.
	// Initially, @currentTime is the begin time of the filtersBlock.
	diff := int(t.Sub(s.currentTime).Seconds())
	if diff > 1 {
		for diff > 1 {
			s.previousPrices = append([]float64{price}, s.previousPrices...)
			diff--
		}
	} else {
		if diff == 0.0 {
			if len(s.previousPrices) >= 1 {
				s.previousPrices = s.previousPrices[1:]
			}
		}
		s.previousPrices = append([]float64{price}, s.previousPrices...)
	}

	if len(s.previousPrices) > s.param {
		s.previousPrices = s.previousPrices[0:s.param]
	}
	s.currentTime = t
}

func (s *FilterMA) finalCompute(t time.Time) float64 {
	if s.lastTrade == nil {
		return 0.0
	}
	s.fill(t, s.lastTrade.EstimatedUSDPrice)

	var total float64 = 0
	for _, v := range s.previousPrices {
		total += v
	}
	div := s.param
	if len(s.previousPrices) > 0 && len(s.previousPrices) < s.param {
		div = len(s.previousPrices)
	}
	s.value = total / float64(div)
	return s.value
}

func (s *FilterMA) filterPointForBlock() *dia.FilterPoint {
	if s.exchange != "" {
		return nil
	}
	return &dia.FilterPoint{
		Symbol: s.asset.Symbol,
		Value:  s.value,
		Name:   "MA" + strconv.Itoa(s.param),
		Time:   s.currentTime,
	}
}

func (s *FilterMA) save(ds models.Datastore) error {
	log.Infof("save called on symbol %s on exchange %s", s.asset.Symbol, s.exchange)
	if s.modified {
		s.modified = false
		err := ds.SetPriceZSET(s.asset.Symbol, s.exchange, s.value, s.currentTime)
		if err != nil {
			log.Errorln("FilterMA: Error:", err)
		}
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
