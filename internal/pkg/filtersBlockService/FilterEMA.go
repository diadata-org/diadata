package filters

import (
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

// FilterEMA is the struct for a moving average filter implementing
// the Filter interface.
type FilterEMA struct {
	pair            dia.Pair
	exchange        string
	currentTime     time.Time
	previousPrices  []float64
	previousVolumes []float64
	lastTrade       *dia.FilterPoint
	param           int
	value           float64
	modified        bool
	filterName      string
	multiplier      float64
}

// NewFilterEMA returns a moving average filter.
func NewFilterEMA(pair dia.Pair, exchange string, currentTime time.Time, blockSize int) *FilterEMA {

	s := &FilterEMA{
		pair:            pair,
		exchange:        exchange,
		previousPrices:  []float64{},
		previousVolumes: []float64{},
		currentTime:     currentTime,
		filterName:      "EMA" + strconv.Itoa(blockSize),
		multiplier:      2 / (float64(blockSize) + 1),
		param:           blockSize,
	}
	return s
}

func (s *FilterEMA) Compute(trade dia.FilterPoint) {
	s.compute(trade)
}

func (s *FilterEMA) compute(trade dia.FilterPoint) {
	s.modified = true
	if s.lastTrade != nil {
		if trade.Time.After(s.currentTime) {
			log.Errorln("FilterMA: Ignoring Trade out of order ", s.currentTime, trade.Time)
			return
		}
		s.fill(trade.Time, *s.lastTrade)
	}
	s.fill(trade.Time, trade)
	log.Println("FilterEMA compute: filled order ", s.currentTime, trade)

	s.lastTrade = &trade
}

func (s *FilterEMA) fill(t time.Time, trade dia.FilterPoint) {
	log.Println("FilterEMA fill ", trade)
	log.Println("FilterEMA e.multiplier ", s.multiplier)
	log.Println("FilterEMA e.value ", s.value)
	s.currentTime = trade.Time
	if s.value == 0 { // this is a proxy for "uninitialized"
		s.value = trade.Value
	} else {
		s.value = (trade.Value * float64(s.multiplier)) + (s.value * (1 - float64(s.multiplier)))
		log.Println("FilterEMA e.value and multiplier ", s.value, s.multiplier)

	}
	log.Println("FilterEMA e.value ", s.value)

}

func (s *FilterEMA) FinalCompute(t time.Time) float64 {
	return s.finalCompute(t)
}

func (s *FilterEMA) finalCompute(t time.Time) float64 {
	return s.value
}

func (s *FilterEMA) FilterPointForBlock() *dia.FilterPoint {
	log.Println("FilterPointForBlock", s.value, s.currentTime, s.pair)

	return &dia.FilterPoint{
		Pair:  s.pair,
		Value: s.value,
		Name:  "EMA" + strconv.Itoa(s.param),
		Time:  s.currentTime,
	}

}

func (s *FilterEMA) filterPointForBlock() *dia.FilterPoint {
	if s.exchange != "" || s.filterName != dia.FilterKing {
		return nil
	}
	return &dia.FilterPoint{
		Pair:  s.pair,
		Value: s.value,
		Name:  "EMA" + strconv.Itoa(s.param),
		Time:  s.currentTime,
	}
}

func (s *FilterEMA) save(ds models.Datastore) error {

	if s.modified {
		s.modified = false
		err := ds.SetPairFilter(s.filterName, s.pair, s.exchange, s.value, s.currentTime)
		if err != nil {
			log.Errorln("FilterMA: Error:", err)
		}
		return err
	}
	return nil
}
