package metafilters

import (
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

type FilterVOL struct {
	asset       dia.Asset
	source      string
	currentTime time.Time
	volumeUSD   float64
	value       float64
	name        string
	childName   string
	modified    bool
}

func NewFilterVOL(asset dia.Asset, exchange string, childFilter string, memory int) *FilterVOL {
	filter := &FilterVOL{
		asset:     asset,
		volumeUSD: 0.0,
		name:      dia.VOL_META_FILTER,
		childName: childFilter,
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
		Asset:  filter.asset,
		Source: filter.source,
		Value:  filter.value,
		Name:   filter.name,
		Time:   filter.currentTime,
	}
}

func (filter *FilterVOL) save(ds models.Datastore) error {
	if filter.modified {
		filter.modified = false
		err := ds.SetFilter(getFilterName(filter.name, filter.childName), filter.asset, filter.source, filter.value, filter.currentTime)
		if err != nil {
			log.Errorln("FilterVOL Error:", err)
		}
		return err
	}
	return nil
}
