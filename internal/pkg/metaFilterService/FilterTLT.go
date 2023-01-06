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
	name          string
	childName     string
	lastTradeTime time.Time
}

func NewFilterTLT(asset dia.Asset, source string, childFilter string, memory int) *FilterTLT {
	s := &FilterTLT{
		asset:     asset,
		source:    source,
		name:      dia.TLT_META_FILTER,
		childName: childFilter,
	}
	return s
}

func (s *FilterTLT) filterPointForBlock() *dia.MetaFilterPoint {
	return nil
}

func (s *FilterTLT) collect(filterPoint dia.PairFilterPoint) {

	s.lastTradeTime = filterPoint.LastTrade.Time
}

func (s *FilterTLT) save(ds models.Datastore) error {
	err := ds.SetLastTradeTimeForExchangeAsset(s.asset, s.source, s.lastTradeTime)
	if err != nil {
		log.Errorln("FilterTLT Error:", err)
	}
	return err
}

func (s *FilterTLT) finalCompute(time time.Time) float64 {
	return 0.0
}
