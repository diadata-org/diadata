package filters

import (
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

type FilterMA struct {
	symbol         string
	exchange       string
	currentTime    time.Time
	previousPrices []float64
	previousVolumes []float64
	lastTrade      *dia.Trade
	param          int
	value          float64
	modified       bool
}

func NewFilterMA(symbol string, exchange string, currentTime time.Time, param int) *FilterMA {
	s := &FilterMA{
		symbol:         symbol,
		exchange:       exchange,
		previousPrices: []float64{},
		currentTime:    currentTime,
		param:          param,
	}
	return s
}
func (s *FilterMA) finalCompute(t time.Time) float64 {
	if s.lastTrade == nil {
		return 0.0
	} else {
		s.fill(t, *s.lastTrade)
	}

	var totalVolume float64 = 0
	var total float64 = 0
	for priceIndex, price := range s.previousPrices {
		total += price * s.previousVolumes[priceIndex]
		totalVolume += s.previousVolumes[priceIndex]
	}
	s.value = total / totalVolume
	return s.value
}

func (s *FilterMA) filterPointForBlock() *dia.FilterPoint {
	if s.exchange != "" {
		return nil
	} else {
		return &dia.FilterPoint{
			Symbol: s.symbol,
			Value:  s.value,
			Name:   "MA" + strconv.Itoa(s.param),
			Time:   s.currentTime,
		}
	}
}

func (s *FilterMA) fill(t time.Time, trade dia.Trade) {
	diff := int(t.Sub(s.currentTime).Seconds())
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

func (s *FilterMA) compute(trade dia.Trade) {
	s.modified = true
	if s.lastTrade != nil {
		if trade.Time.Before(s.currentTime) {
			log.Errorln("FilterMA: Ignoring Trade out of order ", s.currentTime, trade.Time)
			return
		} else {
			s.fill(trade.Time, *s.lastTrade)
		}
	}
	s.fill(trade.Time, trade)
	s.lastTrade = &trade
}

func (s *FilterMA) save(ds models.Datastore) error {
	log.Infof("save called on symbol %s on exchange %s", s.symbol, s.exchange)
	if s.modified {
		s.modified = false
		err := ds.SetPriceZSET(s.symbol, s.exchange, s.value, s.currentTime)
		if err != nil {
			log.Errorln("FilterMA: Error:", err)
		}
		if s.exchange == "" {
			err = ds.SetPriceUSD(s.symbol, s.value)
			if err != nil {
				log.Errorln("FilterMA: Error:", err)
			}
		}
		return err
	} else {
		return nil
	}
}
