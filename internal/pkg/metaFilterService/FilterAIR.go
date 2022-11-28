package metafilters

import (
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

// FilterMAIR implements a trimmed moving average.
// Outliers are eliminated using interquartile range.
// see: https://en.wikipedia.org/wiki/Interquartile_range
type FilterAIR struct {
	asset           dia.Asset
	source          string
	currentTime     time.Time
	values          []float64
	volumes         []float64
	lastFilterValue dia.FilterPoint
	value           float64
	name            string
	childName       string
	modified        bool
}

// NewFilterAIR returns a FilterAIR
func NewFilterAIR(asset dia.Asset, currentTime time.Time, memory int) *FilterAIR {
	filter := &FilterAIR{
		asset:       asset,
		values:      []float64{},
		volumes:     []float64{},
		currentTime: currentTime,
		name:        dia.AIR_META_FILTER,
		childName:   dia.MAIR_FILTER + strconv.Itoa(memory),
	}
	return filter
}

func (filter *FilterAIR) Collect(filterPoint dia.FilterPoint) {
	filter.collect(filterPoint)
}

func (filter *FilterAIR) collect(filterPoint dia.FilterPoint) {
	filter.modified = true
	if filterPoint.Name != filter.name {
		// Child filter method does not match metafilter's name.
		return
	}
	if filter.lastFilterValue != (dia.FilterPoint{}) {
		if filterPoint.Time.Before(filter.currentTime) {
			log.Errorln("FilterMAIR: Ignoring filterPoint out of order ", filter.currentTime, filterPoint.Time)
			return
		}
	}
	filter.processDataPoint(filterPoint)
	filter.lastFilterValue = filterPoint
}

func (filter *FilterAIR) processDataPoint(filterPoint dia.FilterPoint) {
	filter.values = append([]float64{filterPoint.Value}, filter.values...)
	filter.volumes = append([]float64{filterPoint.BlockVolume}, filter.volumes...)
}

func (filter *FilterAIR) FinalCompute(t time.Time) float64 {
	return filter.finalCompute(t)
}

func (filter *FilterAIR) finalCompute(t time.Time) float64 {
	if filter.lastFilterValue == (dia.FilterPoint{}) {
		return 0.0
	}

	var err error
	cleanPrices, bounds := removeOutliers(filter.values)
	if len(bounds) > 1 {
		filter.value, err = computeMean(cleanPrices, filter.volumes[bounds[0]:bounds[1]])
		if err != nil {
			return 0.0
		}
	} else {
		filter.value = cleanPrices[0]
	}

	// Reduce the filter values to the last recorded value for the next filtersblock.
	// if len(filter.values) > 0 && len(filter.volumes) > 0 {
	// 	filter.values = []float64{filter.lastFilterValue.Value}
	// 	filter.volumes = []float64{filter.lastFilterValue.BlockVolume}
	// }
	return filter.value
}

func (filter *FilterAIR) FilterPointForBlock() *dia.MetaFilterPoint {
	return &dia.MetaFilterPoint{
		Asset:  filter.asset,
		Source: filter.source,
		Value:  filter.value,
		Name:   filter.name,
		Time:   filter.currentTime,
	}
}

func (filter *FilterAIR) filterPointForBlock() *dia.MetaFilterPoint {
	if filter.name != dia.FilterKing {
		return nil
	}
	return &dia.MetaFilterPoint{
		Asset:  filter.asset,
		Source: filter.source,
		Value:  filter.value,
		Name:   filter.name,
		Time:   filter.currentTime,
	}
}

func (filter *FilterAIR) save(ds models.Datastore) error {
	if filter.modified {
		filter.modified = false
		err := ds.SetFilter(getFilterName(filter.name, filter.childName), filter.asset, filter.source, filter.value, filter.currentTime)
		if err != nil {
			log.Errorln("FilterMAIR: Error:", err)
		}

		// Additionally, the price across exchanges is saved in influx as a quotation.
		// This price is used for the estimation of quote tokens' prices in the tradesBlockService.
		err = ds.SetAssetPriceUSD(filter.asset, filter.value, filter.currentTime)
		if err != nil {
			log.Errorln("FilterMAIR: Error:", err)
		}

		return err
	}
	return nil
}
