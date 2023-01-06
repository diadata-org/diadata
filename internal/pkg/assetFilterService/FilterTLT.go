// TLT Time last trade

package assetfilters

import (
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
)

type FilterTLT struct {
	asset         dia.Asset
	exchange      string
	lastTradeTime time.Time
}

func NewFilterTLT(asset dia.Asset, exchange string) *FilterTLT {
	s := &FilterTLT{
		asset:    asset,
		exchange: exchange,
	}
	return s
}

func (s *FilterTLT) filterPointForBlock() *dia.AssetFilterPoint {
	return nil
}

func (s *FilterTLT) compute(trade dia.Trade) {
	s.lastTradeTime = trade.Time
}

func (s *FilterTLT) save(ds models.Datastore) error {
	// TO DO: Implement setter for asset.

	// err := ds.SetLastTradeTimeForExchange(s.asset, s.exchange, s.lastTradeTime)
	// if err != nil {
	// 	log.Errorln("FilterTLT Error:", err)
	// }
	return nil
}

func (s *FilterTLT) finalCompute(time time.Time) float64 {
	return 0.0
}
