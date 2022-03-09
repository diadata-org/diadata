package filters

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
	exchange        string
	currentTime     time.Time
	previousPrices  []float64
	previousVolumes []float64
	lastTrade       *dia.Trade
	memory          int
	value           float64
	filterName      string
	modified        bool
}

//NewFilterMAIR returns a FilterMAIR
func NewFilterMAIR(asset dia.Asset, exchange string, currentTime time.Time, memory int) *FilterMAIR {
	s := &FilterMAIR{
		asset:           asset,
		exchange:        exchange,
		previousPrices:  []float64{},
		previousVolumes: []float64{},
		currentTime:     currentTime,
		memory:          memory,
		filterName:      "MAIR" + strconv.Itoa(memory),
	}
	return s
}

func (s *FilterMAIR) Compute(trade dia.Trade) {
	s.compute(trade)
}

func (s *FilterMAIR) compute(trade dia.Trade) {
	s.modified = true
	if s.lastTrade != nil {
		if trade.Time.Before(s.currentTime) {
			log.Errorln("FilterMAIR: Ignoring Trade out of order ", s.currentTime, trade.Time)
			return
		}
	}
	s.fill(trade.Time, trade)
	s.lastTrade = &trade
}

func (s *FilterMAIR) fill(t time.Time, trade dia.Trade) {
	diff := int(t.Sub(s.currentTime).Seconds())
	if diff > 1 {
		for diff > 1 {
			s.processDataPoint(trade)
			diff--
		}
	} else {
		if diff == 0.0 {
			if len(s.previousPrices) >= 1 {
				/// Remove latest data point and update with newer
				s.previousPrices = s.previousPrices[1:]
				s.previousVolumes = s.previousVolumes[1:]
			}
		}
		s.processDataPoint(trade)
	}
	s.currentTime = t
}

func (s *FilterMAIR) processDataPoint(trade dia.Trade) {
	/// first remove extra value from buffer if already full
	if len(s.previousPrices) >= s.memory {
		s.previousPrices = s.previousPrices[0 : s.memory-1]
		s.previousVolumes = s.previousVolumes[0 : s.memory-1]
	}
	s.previousPrices = append([]float64{trade.EstimatedUSDPrice}, s.previousPrices...)
	s.previousVolumes = append([]float64{trade.Volume}, s.previousVolumes...)
}

func (s *FilterMAIR) FinalCompute(t time.Time) float64 {
	return s.finalCompute(t)
}

func (s *FilterMAIR) finalCompute(t time.Time) float64 {
	if s.lastTrade == nil {
		return 0.0
	}
	if s.asset.Address == "0xdAC17F958D2ee523a2206206994597C13D831ec7" && s.asset.Blockchain == dia.ETHEREUM {
		log.Info("estimatedUSDPrices in finalCompute for USDT: ", s.previousPrices)
		log.Info("volumes in finalCompute for USDT: ", s.previousPrices)
	}
	// Add the last trade again to compensate for the delay since measurement to EOB
	// adopted behaviour from FilterMA
	s.processDataPoint(*s.lastTrade)
	cleanPrices, bounds := removeOutliers(s.previousPrices)
	mean, err := computeMean(cleanPrices, s.previousVolumes[bounds[0]:bounds[1]])
	if err != nil {
		return 0.0
	}
	s.value = mean
	return s.value
}

func (s *FilterMAIR) FilterPointForBlock() *dia.FilterPoint {
	return &dia.FilterPoint{
		Asset: s.asset,
		Value: s.value,
		Name:  s.filterName,
		Time:  s.currentTime,
	}
}

func (s *FilterMAIR) filterPointForBlock() *dia.FilterPoint {
	if s.exchange != "" || s.filterName != dia.FilterKing {
		return nil
	}
	return &dia.FilterPoint{
		Asset: s.asset,
		Value: s.value,
		Name:  s.filterName,
		Time:  s.currentTime,
	}
}

func (s *FilterMAIR) save(ds models.Datastore) error {
	if s.modified {
		s.modified = false
		err := ds.SetFilter(s.filterName, s.asset, s.exchange, s.value, s.currentTime)
		if err != nil {
			log.Errorln("FilterMAIR: Error:", err)
		}
		// log.Infof("set price for %s: %v", s.asset.Symbol, s.value)

		// Additionally, the price across exchanges is saved in influx as a quotation.
		// This price is used for the estimation of quote tokens' prices in the tradesBlockService.
		if s.exchange == "" {
			err = ds.SetAssetPriceUSD(s.asset, s.value, s.currentTime)
			if err != nil {
				log.Errorln("FilterMAIR: Error:", err)
			}
		}
		return err
	}
	return nil
}
