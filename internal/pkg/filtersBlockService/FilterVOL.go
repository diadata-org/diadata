package filters

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

func (filter *FilterVOL) compute(trade dia.Trade) {
	filter.volumeUSD += trade.EstimatedUSDPrice * math.Abs(trade.Volume)
	if filter.asset.Address == "0x249e38Ea4102D0cf8264d3701f1a0E39C4f2DC3B" || filter.asset.Address == "0xFca59Cd816aB1eaD66534D82bc21E7515cE441CF" {
		log.Infof("volumeUSD for %s: %v", filter.asset.Address, filter.volumeUSD)
	}
	filter.currentTime = trade.Time
}

func (filter *FilterVOL) finalCompute(time time.Time) float64 {
	filter.value = filter.volumeUSD
	filter.volumeUSD = 0.0
	return filter.value
}

func (filter *FilterVOL) filterPointForBlock() *dia.FilterPoint {
	return nil
}

func (filter *FilterVOL) save(ds models.Datastore) error {
	err := ds.SetFilter(filter.filterName, filter.asset, filter.exchange, filter.value, filter.currentTime)
	if err != nil {
		log.Errorln("FilterVOL Error:", err)
	}
	return err
}
