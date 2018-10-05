package filters

import (
	"github.com/diadata-org/api-golang/dia"
	"github.com/diadata-org/api-golang/pkg/model"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type FilterMA struct {
	symbol         string
	exchange       string
	currentTime    time.Time
	previousPrices []float64
	lastTrade      *dia.Trade
	param          int
	value          float64
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
func (s *FilterMA) finalComputeEndOfBlock(t time.Time) {
	if s.lastTrade == nil {
		return
	} else {
		s.fill(t, s.lastTrade.EstimatedUSDPrice)
	}

	var total float64 = 0
	for _, v := range s.previousPrices {
		total += v
	}
	div := s.param
	if len(s.previousPrices) > 0 && len(s.previousPrices) < s.param {
		div = len(s.previousPrices)
	}
	s.value = total / float64(div)
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

func (s *FilterMA) fill(t time.Time, price float64) {
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

func (s *FilterMA) compute(trade dia.Trade) {
	if s.lastTrade != nil {
		if trade.Time.Before(s.currentTime) {
			log.Errorln("FilterMA: Ignoring Trade out of order ", s.currentTime, trade.Time)
		} else {
			s.fill(trade.Time, s.lastTrade.EstimatedUSDPrice)
		}
	}
	s.fill(trade.Time, trade.EstimatedUSDPrice)
	s.lastTrade = &trade
}

func (s *FilterMA) save(ds models.Datastore) error {

	err := ds.SetPriceZSET(s.symbol, s.exchange, s.value)
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
}
