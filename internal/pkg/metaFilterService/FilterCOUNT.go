package metafilters

import (
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

type FilterCOUNT struct {
	asset       dia.Asset
	source      string
	currentTime time.Time
	numTrades   int64
	value       int64
	name        string
	childName   string
	memory      int
	modified    bool
}

func NewFilterCOUNT(asset dia.Asset, source string, childFilter string, memory int) *FilterCOUNT {
	filter := &FilterCOUNT{
		asset:     asset,
		source:    source,
		numTrades: int64(0),
		name:      dia.COUNT_META_FILTER,
		childName: childFilter,
		memory:    memory,
	}
	return filter
}

func (filter *FilterCOUNT) Collect(filterPoint dia.PairFilterPoint, starttime time.Time, endtime time.Time) {
	filter.collect(filterPoint, starttime, endtime)
}

func (filter *FilterCOUNT) FinalCompute(t time.Time) {
	filter.finalCompute(t)
}

func (filter *FilterCOUNT) collect(filterPoint dia.PairFilterPoint, starttime time.Time, endtime time.Time) {
	if filterPoint.Name != filter.childName {
		// Additional safety check. Child filter method does not match metafilter's name.
		log.Warn("filter point does not come from the correct child filter.")
		return
	}
	filter.numTrades += int64(filterPoint.Value)
	filter.currentTime = endtime
	filter.modified = true
}

func (filter *FilterCOUNT) finalCompute(t time.Time) float64 {
	filter.value = filter.numTrades
	filter.numTrades = int64(0)
	return float64(filter.value)
}

func (filter *FilterCOUNT) filterPointForBlock() *dia.MetaFilterPoint {
	return nil
}

func (filter *FilterCOUNT) FilterPointForBlock() *dia.MetaFilterPoint {
	return &dia.MetaFilterPoint{
		Asset:  filter.asset,
		Source: filter.source,
		Value:  float64(filter.value),
		Name:   filter.name,
		Time:   filter.currentTime,
	}
}

func (filter *FilterCOUNT) save(ds models.Datastore) error {
	if filter.modified {
		filter.modified = false
		err := ds.SetFilter(getFilterName(filter.name, filter.childName), filter.asset, filter.source, float64(filter.value), filter.currentTime)
		if err != nil {
			log.Errorln("FilterCOUNT Error:", err)
		}
		return err
	}
	return nil
}
