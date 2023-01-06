package assetfilters

import (
	"math"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

type FilterVOL struct {
	asset       dia.Asset
	exchange    string
	currentTime time.Time
	volumeUSD   float64
	value       float64
	filterName  string
	memory      int
	modified    bool
}

func NewFilterVOL(asset dia.Asset, exchange string, memory int) *FilterVOL {
	filter := &FilterVOL{
		asset:      asset,
		exchange:   exchange,
		volumeUSD:  0.0,
		filterName: "VOL" + strconv.Itoa(memory),
		memory:     memory,
	}
	return filter
}

func (filter *FilterVOL) Compute(trade dia.Trade) {
	filter.compute(trade)
}
func (filter *FilterVOL) FinalCompute(t time.Time) {
	filter.finalCompute(t)
}

func (filter *FilterVOL) compute(trade dia.Trade) {
	filter.modified = true
	filter.volumeUSD += trade.EstimatedUSDPrice * math.Abs(trade.Volume)
	filter.currentTime = trade.Time
}

func (filter *FilterVOL) finalCompute(time time.Time) float64 {
	filter.value = filter.volumeUSD
	filter.volumeUSD = 0.0
	return filter.value
}

func (filter *FilterVOL) filterPointForBlock() *dia.AssetFilterPoint {
	return nil
}

func (filter *FilterVOL) FilterPointForBlock() *dia.AssetFilterPoint {
	return &dia.AssetFilterPoint{
		Asset: filter.asset,
		Value: filter.value,
		Name:  filter.filterName,
		Time:  filter.currentTime,
	}
}

func (filter *FilterVOL) save(ds models.Datastore) error {
	if filter.modified {
		filter.modified = false
		err := ds.SetFilter(filter.filterName, filter.asset, filter.exchange, filter.value, filter.currentTime)
		if err != nil {
			log.Errorln("FilterVOL Error:", err)
		}
		return err
	}
	return nil
}
