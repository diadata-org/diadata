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
	asset           dia.Asset
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
func NewFilterEMA(asset dia.Asset, exchange string, currentTime time.Time, blockSize int) *FilterEMA {

	s := &FilterEMA{
		asset:           asset,
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

func (e *FilterEMA) fill(t time.Time, trade dia.FilterPoint) {
	log.Println("FilterEMA fill ", trade)
	log.Println("FilterEMA e.multiplier ", e.multiplier)
	log.Println("FilterEMA e.value ", e.value)
	e.currentTime = trade.Time
	if e.value == 0 { // this is a proxy for "uninitialized"
		e.value = trade.Value
	} else {
		e.value = (trade.Value * float64(e.multiplier)) + (e.value * (1 - float64(e.multiplier)))
		log.Println("FilterEMA e.value and multiplier ", e.value, e.multiplier)

	}
	log.Println("FilterEMA e.value ", e.value)

}

func (s *FilterEMA) FinalCompute(t time.Time) float64 {
	return s.finalCompute(t)
}

func (e *FilterEMA) finalCompute(t time.Time) float64 {
	return e.value
}

func (e *FilterEMA) FilterPointForBlock() *dia.FilterPoint {
	log.Println("FilterPointForBlock", e.value, e.currentTime, e.asset)

	return &dia.FilterPoint{
		Asset: e.asset,
		Value: e.value,
		Name:  "EMA" + strconv.Itoa(e.param),
		Time:  e.currentTime,
	}

}

func (s *FilterEMA) filterPointForBlock() *dia.FilterPoint {
	if s.exchange != "" || s.filterName != dia.FilterKing {
		return nil
	}
	return &dia.FilterPoint{
		Asset: s.asset,
		Value: s.value,
		Name:  "EMA" + strconv.Itoa(s.param),
		Time:  s.currentTime,
	}
}

func (s *FilterEMA) save(ds models.Datastore) error {

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
