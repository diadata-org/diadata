package nftsource

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/sirupsen/logrus"
)

// AssetSource is the interface that must be implemented by an asset scraper.
type NFTClassSource interface {
	NFTClass() chan dia.NFTClass
	Close() chan bool
}

var log *logrus.Logger

func init() {
	log = logrus.New()
}
