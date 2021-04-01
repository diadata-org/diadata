package source

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/sirupsen/logrus"
)

// AssetSource is the interface that must be implemented by an asset scraper.
type AssetSource interface {
	Asset() chan dia.Asset
	Close() chan bool
}

var log *logrus.Logger

func init() {
	log = logrus.New()
}
