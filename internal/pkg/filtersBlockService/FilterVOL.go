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

func (filter *FilterVOL) Compute(trade dia.Trade) {
	filter.compute(trade)
}
func (filter *FilterVOL) FinalCompute(t time.Time) {
	filter.finalCompute(t)
}

func (filter *FilterVOL) compute(trade dia.Trade) {
	filter.volumeUSD += trade.EstimatedUSDPrice * math.Abs(trade.Volume)
	if (filter.asset.Address == "0x249e38Ea4102D0cf8264d3701f1a0E39C4f2DC3B" && filter.asset.Blockchain == dia.ETHEREUM) || (filter.asset.Address == "0xFca59Cd816aB1eaD66534D82bc21E7515cE441CF" && filter.asset.Blockchain == dia.ETHEREUM) {
		log.Infof("volumeUSD for %s on %s: %v", filter.asset.Address, filter.exchange, filter.volumeUSD)
	}
	filter.currentTime = trade.Time
}

func (filter *FilterVOL) finalCompute(time time.Time) float64 {
	filter.value = filter.volumeUSD
	filter.volumeUSD = 0.0
	if (filter.asset.Address == "0x249e38Ea4102D0cf8264d3701f1a0E39C4f2DC3B" && filter.asset.Blockchain == dia.ETHEREUM) || (filter.asset.Address == "0xFca59Cd816aB1eaD66534D82bc21E7515cE441CF" && filter.asset.Blockchain == dia.ETHEREUM) {
		log.Infof("filter.value in final compute for %s on %s: %v", filter.asset.Address, filter.exchange, filter.value)
	}
	return filter.value
}

func (filter *FilterVOL) filterPointForBlock() *dia.FilterPoint {
	return nil
}

func (filter *FilterVOL) FilterPointForBlock() *dia.FilterPoint {
	return &dia.FilterPoint{
		Asset: filter.asset,
		Value: filter.value,
		Name:  filter.filterName,
		Time:  filter.currentTime,
	}
}

func (filter *FilterVOL) save(ds models.Datastore) error {
	if (filter.asset.Address == "0x249e38Ea4102D0cf8264d3701f1a0E39C4f2DC3B" && filter.asset.Blockchain == dia.ETHEREUM) || (filter.asset.Address == "0xFca59Cd816aB1eaD66534D82bc21E7515cE441CF" && filter.asset.Blockchain == dia.ETHEREUM) {
		log.Infof("set filter.value in save() for %s on %s: %v", filter.asset.Address, filter.exchange, filter.value)
	}
	err := ds.SetFilter(filter.filterName, filter.asset, filter.exchange, filter.value, filter.currentTime)
	if err != nil {
		log.Errorln("FilterVOL Error:", err)
	}
	return err
}
