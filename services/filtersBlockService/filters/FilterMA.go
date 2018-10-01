package filters

import (
	"github.com/diadata-org/api-golang/dia"
	"github.com/diadata-org/api-golang/services/model"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type FilterMA struct {
	Symbol         string
	Key            string
	currentTime    time.Time
	previousPrices []float64
	param          int
	Value          float64
	lastTrade      *dia.Trade
}

func NewFilterMA(symbol string, key string, currentTime time.Time, param int) *FilterMA {
	s := &FilterMA{
		Symbol:         symbol,
		Key:            "MA" + strconv.Itoa(param) + "_" + key,
		param:          param,
		previousPrices: []float64{},
		currentTime:    currentTime,
	}
	return s
}

func (s *FilterMA) filterPoint(time time.Time) dia.FilterPoint {
	var total float64 = 0
	for _, v := range s.previousPrices {
		total += v
	}

	div := s.param

	if len(s.previousPrices) > 0 && len(s.previousPrices) < s.param {
		div = len(s.previousPrices)
	}

	s.Value = total / float64(div)

	log.Info(s.Symbol, s.previousPrices, s.Value)

	return dia.FilterPoint{
		Symbol: s.Symbol,
		Value:  s.Value,
		Name:   "MA" + strconv.Itoa(s.param),
		Time:   time,
	}
}

func (s *FilterMA) fill(t time.Time, price float64) {
	diff := int(t.Sub(s.currentTime).Seconds())
	if diff >= 1 {
		for diff >= 1 {
			s.previousPrices = append([]float64{price}, s.previousPrices...)
			if len(s.previousPrices) > s.param {
				s.previousPrices = s.previousPrices[0:s.param]
			}
			diff--
		}
		s.currentTime = t
	}
}

func (s *FilterMA) compute(trade *dia.Trade) {
	if s.lastTrade != nil {
		if trade.Time.Before(s.currentTime) {
			log.Error("Ignoring Trade out of order ", s.currentTime, trade.Time)
		} else {
			s.fill(trade.Time, s.lastTrade.EstimatedUSDPrice)
		}
	}
	s.lastTrade = trade
}

func (s *FilterMA) save(ds models.Datastore) error {

	err := ds.SetPriceZSET(s.Key, s.Value)
	if err != nil {
		log.Error("Error: %v\n", err)
	}

	// TOFIX
	if s.Key != "MA"+strconv.Itoa(s.param)+"_"+s.Symbol {
		return nil
	}
	//

	err = ds.SetPriceUSD(s.Symbol, s.Value)
	if err != nil {
		log.Error("Error: %v\n", err)
	}

	return err
}

func (s *FilterMA) copyToFilterBlock() bool {

	// TOFIX
	if s.Key != "MA"+strconv.Itoa(s.param)+"_"+s.Symbol {
		return false
	}
	//

	return true
}
