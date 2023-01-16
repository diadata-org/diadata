package pairfilters

import (
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

type FilterCOUNT struct {
	pair        dia.Pair
	exchange    string
	currentTime time.Time
	endTime     time.Time
	numTrades   int64
	value       int64
	filterName  string
	memory      int
	modified    bool
}

func NewFilterCOUNT(quotetoken dia.Asset, basetoken dia.Asset, exchange string, endTime time.Time, memory int) *FilterCOUNT {
	filter := &FilterCOUNT{
		pair:       dia.Pair{QuoteToken: quotetoken, BaseToken: basetoken},
		exchange:   exchange,
		numTrades:  int64(0),
		filterName: "COUNT" + strconv.Itoa(memory),
		endTime:    endTime,
		memory:     memory,
	}
	return filter
}

func (filter *FilterCOUNT) Compute(trade dia.Trade) {
	filter.compute(trade)
}
func (filter *FilterCOUNT) FinalCompute(t time.Time) {
	filter.finalCompute(t)
}

func (filter *FilterCOUNT) compute(trade dia.Trade) {
	filter.modified = true
	filter.numTrades += 1
	filter.currentTime = trade.Time
}

func (filter *FilterCOUNT) finalCompute(t time.Time) float64 {
	filter.value = filter.numTrades
	filter.numTrades = int64(0)
	return float64(filter.value)
}

func (filter *FilterCOUNT) filterPointForBlock() *dia.PairFilterPoint {
	return &dia.PairFilterPoint{
		Pair:   filter.pair,
		Source: filter.exchange,
		Value:  float64(filter.value),
		Name:   filter.filterName,
		Time:   filter.currentTime,
	}
}

func (filter *FilterCOUNT) FilterPointForBlock() *dia.PairFilterPoint {
	return filter.filterPointForBlock()
}

func (filter *FilterCOUNT) save(ds models.Datastore) error {
	if filter.modified {
		filter.modified = false
		err := ds.SetPairFilter(filter.filterName, filter.pair, filter.exchange, float64(filter.value), filter.endTime)
		if err != nil {
			log.Errorln("FilterCOUNT Error:", err)
		}
		return err
	}
	return nil
}
