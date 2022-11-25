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
type FilterMAIR struct {
	asset           dia.Asset
	currentTime     time.Time
	values          []float64
	volumes         []float64
	lastFilterValue dia.FilterPoint
	value           float64
	name            string
	modified        bool
}

// NewFilterMAIR returns a FilterMAIR
func NewFilterMAIR(asset dia.Asset, currentTime time.Time, memory int) *FilterMAIR {
	filter := &FilterMAIR{
		asset:       asset,
		values:      []float64{},
		volumes:     []float64{},
		currentTime: currentTime,
		name:        "MAIR" + strconv.Itoa(memory),
	}
	return filter
}

func (filter *FilterMAIR) Collect(filterPoint dia.FilterPoint) {
	filter.collect(filterPoint)
}

func (filter *FilterMAIR) collect(filterPoint dia.FilterPoint) {
	filter.modified = true
	if filterPoint.Name != filter.name {
		// Filter method does not match metafilter's name.
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

func (filter *FilterMAIR) processDataPoint(filterPoint dia.FilterPoint) {
	filter.values = append([]float64{filterPoint.Value}, filter.values...)
	filter.volumes = append([]float64{filterPoint.BlockVolume}, filter.volumes...)
}

func (filter *FilterMAIR) FinalCompute(t time.Time) float64 {
	return filter.finalCompute(t)
}

func (filter *FilterMAIR) finalCompute(t time.Time) float64 {
	if filter.lastFilterValue == (dia.FilterPoint{}) {
		return 0.0
	}
	cleanPrices, bounds := removeOutliers(filter.values)
	mean, err := computeMean(cleanPrices, filter.volumes[bounds[0]:bounds[1]])
	if err != nil {
		return 0.0
	}
	filter.value = mean
	// Reduce the filter values to the last recorded value for the next tradesblock.
	if len(filter.values) > 0 && len(filter.volumes) > 0 {
		filter.values = []float64{filter.lastFilterValue.Value}
		filter.volumes = []float64{filter.lastFilterValue.BlockVolume}
	}
	return filter.value
}

func (filter *FilterMAIR) FilterPointForBlock() *dia.MetaFilterPoint {
	return &dia.MetaFilterPoint{
		Asset: filter.asset,
		Value: filter.value,
		Name:  filter.name,
		Time:  filter.currentTime,
	}
}

func (filter *FilterMAIR) filterPointForBlock() *dia.MetaFilterPoint {
	if filter.name != dia.FilterKing {
		return nil
	}
	return &dia.MetaFilterPoint{
		Asset: filter.asset,
		Value: filter.value,
		Name:  filter.name,
		Time:  filter.currentTime,
	}
}

func (filter *FilterMAIR) save(ds models.Datastore) error {
	if filter.modified {
		filter.modified = false
		err := ds.SetMetaFilter(filter.name, filter.pair, filter.value, filter.currentTime)
		if err != nil {
			log.Errorln("FilterMAIR: Error:", err)
		}

		// Additionally, the price across exchanges is saved in influx as a quotation.
		// This price is used for the estimation of quote tokens' prices in the tradesBlockService.
		err = ds.SetAssetPriceUSD(filter.pair, filter.value, filter.currentTime)
		if err != nil {
			log.Errorln("FilterMAIR: Error:", err)
		}

		return err
	}
	return nil
}
