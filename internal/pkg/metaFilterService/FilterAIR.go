package metafilters

import (
	"math"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

type FilterAIR struct {
	asset           dia.Asset
	source          string
	values          []float64
	volumes         []float64
	currentTime     time.Time
	lastFilterValue dia.PairFilterPoint
	value           float64
	name            string
	childName       string
	modified        bool
}

// NewFilterAIR returns a FilterAIR
func NewFilterAIR(asset dia.Asset, source string, childFilter string, currentTime time.Time, memory int) *FilterAIR {
	filter := &FilterAIR{
		asset:       asset,
		source:      source,
		values:      []float64{},
		volumes:     []float64{},
		currentTime: currentTime,
		name:        dia.AIR_META_FILTER,
		childName:   childFilter,
	}
	return filter
}

func (filter *FilterAIR) Collect(filterPoint dia.PairFilterPoint, starttime time.Time, endtime time.Time) {
	filter.collect(filterPoint, starttime, endtime)
}

func (filter *FilterAIR) collect(filterPoint dia.PairFilterPoint, starttime time.Time, endtime time.Time) {
	if filterPoint.Name != filter.childName {
		// Additional safety check. Child filter method does not match metafilter's name.
		log.Warn("filter point does not come from the correct child filter.")
		return
	}
	if filter.lastFilterValue != (dia.PairFilterPoint{}) {
		if filterPoint.Time.Before(starttime) {
			// log.Errorln("FilterAIR: Ignoring filterPoint out of order ", starttime, filterPoint.Time)
			// log.Errorf("filter asset and source: %s -- %s", filter.asset, filter.source)
			return
		}
	}
	filter.modified = true
	filter.processDataPoint(filterPoint)
	filter.lastFilterValue = filterPoint
	filter.currentTime = endtime

}

func (filter *FilterAIR) processDataPoint(filterPoint dia.PairFilterPoint) {
	filter.values = append([]float64{filterPoint.Value}, filter.values...)
	filter.volumes = append([]float64{filterPoint.BlockVolume}, filter.volumes...)
}

func (filter *FilterAIR) FinalCompute(t time.Time) float64 {
	return filter.finalCompute(t)
}

func (filter *FilterAIR) finalCompute(t time.Time) float64 {
	if filter.lastFilterValue == (dia.PairFilterPoint{}) {
		log.Info("last filter point empty.")
		return 0.0
	}
	// log.Infof("filter values modified - %v for asset %s on %s: %v", filter.modified, filter.asset.Blockchain+"-"+filter.asset.Address, filter.source, filter.values)

	// For now rely on the outlier analysis from the pair filters and only take into account
	// volume weighting here.
	var err error
	filter.value, err = computeMean(filter.values, filter.volumes)
	if err != nil {
		return 0.0
	}
	if filter.asset.Blockchain == "BitcoinCash" {
		if math.IsNaN(filter.value) {
			log.Warn("values: ", filter.values)
			log.Warn("volumes: ", filter.volumes)
		}

	}

	// Reset the filter values for the next filtersblock.
	if len(filter.values) > 0 && len(filter.volumes) > 0 {
		filter.values = []float64{}
		filter.volumes = []float64{}
	}
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
		// log.Infof("modified false for asset %s on %s.", filter.asset.Blockchain+"-"+filter.asset.Address, filter.source)
		err := ds.SetFilter(getFilterName(filter.name, filter.childName), filter.asset, filter.source, filter.value, filter.currentTime)
		if err != nil {
			log.Errorln("FilterAIR: Error:", err)
			log.Infof("info %s and value %v", filter.asset.Blockchain+"-"+filter.asset.Address+"--"+filter.source, filter.value)
		}

		if filter.childName == dia.FilterKing+strconv.Itoa(dia.BlockSizeSeconds) {
			// Additionally, the price across exchanges is saved in influx as a quotation.
			// This price is used for the estimation of quote tokens' prices in the tradesBlockService.
			err = ds.SetAssetPriceUSD(filter.asset, filter.value, filter.currentTime)
			if err != nil {
				log.Errorln("FilterAIR: Error:", err)
			}
		}

		return err
	}
	return nil
}
