package filters

import (
	"github.com/VividCortex/ewma"
	"github.com/diadata-org/api-golang/dia"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

const (
	tradeWeirdPriceRatioWarning = 0.2
)

type Filter struct {
	Symbol            string
	currentTime       time.Time
	currentPrice      float64
	previousValues    []float64
	calculationValues []int
	Results           map[string]float64
	previousTrade     *dia.Trade
}

func NewFilter(calculationValues []int, symbol string, beginTime time.Time, beginPrice float64) *Filter {
	s := &Filter{
		Symbol:            symbol,
		currentTime:       beginTime,
		currentPrice:      beginPrice,
		previousValues:    []float64{},
		calculationValues: calculationValues,
		Results:           make(map[string]float64),
	}
	return s
}

func (s *Filter) calculateEWMAs(current time.Time) {
	for _, v := range s.calculationValues {
		if len(s.previousValues) >= v {
			e := ewma.NewMovingAverage() //=> Returns a SimpleEWMA if called without params
			for i := v; i > 0; i-- {
				e.Add(s.previousValues[i-1])
			}
			s.Results["EWMA"+strconv.Itoa(v)] = e.Value()
		}
	}
}

func (s *Filter) calculateMAs(current time.Time) {
	var total float64 = 0
	for index, v := range s.previousValues {
		total += v
		for _, v := range s.calculationValues {
			if index+1 == v {
				var result = total / float64(v)
				s.Results["MA"+strconv.Itoa(v)] = result
			}
		}
	}
}

func (s *Filter) appendBogusValue() {
	s.previousValues = append([]float64{s.currentPrice}, s.previousValues...)
	maxSize := s.calculationValues[len(s.calculationValues)-1]
	sizeSliceMax := len(s.previousValues)
	if sizeSliceMax > maxSize {
		sizeSliceMax = maxSize
	}
	s.previousValues = s.previousValues[0:sizeSliceMax]
}

func (s *Filter) checkWeirdTrades(trade *dia.Trade) bool {
	if s.previousTrade != nil {

		if trade.Symbol != s.Symbol {
			log.Warning("trade received on wrong filter:", s.Symbol, "/", trade.Symbol, " ", trade)
			return false
		}

		s1 := s.previousTrade.Symbol
		s2 := trade.Symbol

		if s1 != s2 {
			log.Warning("Weird received two differents symbols on the same filter:", s.Symbol, " ", s1, "/", s2, " ", trade, s.previousTrade)
		}

		p1 := s.previousTrade.EstimatedUSDPrice
		p2 := trade.EstimatedUSDPrice

		if ((p1 / p2) > 1.0+tradeWeirdPriceRatioWarning) ||
			((p1 / p2) < 1.0-tradeWeirdPriceRatioWarning) {
			log.Warning("Weird price differences on two trades:", p1, " USD vs ", p2, " USD ", trade, s.previousTrade)
		}
	}

	type Trade struct {
		Symbol            string
		Pair              string
		Price             float64
		Volume            float64 // negative if result of Market order Sell
		Time              time.Time
		ForeignTradeID    string
		EstimatedUSDPrice float64
		Source            string
	}

	s.previousTrade = &dia.Trade{
		Symbol:            trade.Symbol,
		Pair:              trade.Pair,
		Price:             trade.Price,
		Volume:            trade.Volume,
		Time:              trade.Time,
		ForeignTradeID:    trade.ForeignTradeID,
		EstimatedUSDPrice: trade.EstimatedUSDPrice,
		Source:            trade.Source,
	}
	return true

}

func (s *Filter) ProcessTrade(trade *dia.Trade) {

	var price = trade.EstimatedUSDPrice
	var t = trade.Time

	if s.checkWeirdTrades(trade) {
		diff := int(t.Sub(s.currentTime).Seconds())

		for diff > 1 {
			s.appendBogusValue()
			nextTime := s.currentTime.Add(time.Second)
			s.currentTime = nextTime
			diff--
		}

		if t.Before(s.currentTime) {
			log.Println("ignoring Trade out of order ", s.currentTime, t)
		} else {
			s.currentTime = t
			s.currentPrice = price
		}
	}
}
