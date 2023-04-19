package dscrapers

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/sirupsen/logrus"
)

func init() {
	var log *logrus.Logger
	log.Info("init done.")
}

type CEXScraper interface {
	ReadPairs() ([]dia.ExchangePair, error)
	TradeChannel() chan *dia.Trade
}

type DEXScraper interface {
	ReadPools() ([]dia.Pool, error)
	TradeChannel() chan *dia.Trade
}
