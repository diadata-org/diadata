// TLT Time last trade

package filters

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
	"time"
)

type FilterTLT struct {
	symbol        string
	exchange      string
	lastTradeTime time.Time
}

func NewFilterTLT(symbol string, exchange string) *FilterTLT {
	s := &FilterTLT{
		symbol:   symbol,
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
	err := ds.SetLastTradeTimeForExchange(s.symbol, s.exchange, s.lastTradeTime)
	if err != nil {
		log.Errorln("FilterTLT Error:", err)
	}
	return err
}

func (s *FilterTLT) finalCompute(time time.Time) float64 {
	return 0.0
}
