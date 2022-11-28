package metafilters

import (
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

type FilterVOL struct {
	asset       dia.Asset
	currentTime time.Time
	volumeUSD   float64
	value       float64
	name        string
	modified    bool
}

func NewFilterVOL(asset dia.Asset, exchange string, memory int) *FilterVOL {
	filter := &FilterVOL{
		asset:     asset,
		volumeUSD: 0.0,
		name:      "VOL" + strconv.Itoa(memory),
	}
	return filter
}

func (filter *FilterVOL) Collect(filterPoint dia.FilterPoint) {
	filter.collect(filterPoint)
}
func (filter *FilterVOL) FinalCompute(t time.Time) {
	filter.finalCompute(t)
}

func (filter *FilterVOL) collect(filterPoint dia.FilterPoint) {
	filter.modified = true
	if filterPoint.Name != filter.name {
		// Child filter method does not match metafilter's name.
		return
	}
	filter.volumeUSD += filterPoint.Value
	filter.currentTime = filterPoint.Time
}

func (filter *FilterVOL) finalCompute(time time.Time) float64 {
	filter.value = filter.volumeUSD
	filter.volumeUSD = 0.0
	return filter.value
}

func (filter *FilterVOL) filterPointForBlock() *dia.MetaFilterPoint {
	return nil
}

func (filter *FilterVOL) FilterPointForBlock() *dia.MetaFilterPoint {
	return &dia.MetaFilterPoint{
		Asset: filter.asset,
		Value: filter.value,
		Name:  filter.name,
		Time:  filter.currentTime,
	}
}

func (filter *FilterVOL) save(ds models.Datastore) error {
	if filter.modified {
		filter.modified = false
		err := ds.SetMetaFilter(filter.name, filter.asset, filter.value, filter.currentTime)
		if err != nil {
			log.Errorln("FilterVOL Error:", err)
		}
		return err
	}
	return nil
}
