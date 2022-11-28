package metafilters

import (
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

type FilterTLT struct {
	asset         dia.Asset
	source        string
	lastTradeTime time.Time
}

func NewFilterTLT(asset dia.Asset, source string) *FilterTLT {
	s := &FilterTLT{
		asset:  asset,
		source: source,
	}
	return s
}

func (s *FilterTLT) filterPointForBlock() *dia.MetaFilterPoint {
	return nil
}

func (s *FilterTLT) collect(filterPoint dia.FilterPoint) {
	s.lastTradeTime = filterPoint.LastTrade.Time
}

func (s *FilterTLT) save(ds models.Datastore) error {
	// TO DO: Write method SetLastTradeTimeForExchangeAsset
	err := ds.SetLastTradeTimeForExchange(s.asset, s.source, s.lastTradeTime)
	if err != nil {
		log.Errorln("FilterTLT Error:", err)
	}
	return err
}

func (s *FilterTLT) finalCompute(time time.Time) float64 {
	return 0.0
}
