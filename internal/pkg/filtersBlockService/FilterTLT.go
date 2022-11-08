// TLT Time last trade

package filters

import (
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

type FilterTLT struct {
	pair          dia.Pair
	exchange      string
	lastTradeTime time.Time
}

func NewFilterTLT(pair dia.Pair, exchange string) *FilterTLT {
	s := &FilterTLT{
		pair:     pair,
		exchange: exchange,
	}
	return s
}

func (s *FilterTLT) filterPointForBlock() *dia.FilterPoint {
	return nil
}

func (s *FilterTLT) compute(trade dia.Trade) {
	s.lastTradeTime = trade.Time
}

func (s *FilterTLT) save(ds models.Datastore) error {
	err := ds.SetLastTradeTimeForExchange(s.pair, s.exchange, s.lastTradeTime)
	if err != nil {
		log.Errorln("FilterTLT Error:", err)
	}
	return err
}

func (s *FilterTLT) finalCompute(time time.Time) float64 {
	return 0.0
}
