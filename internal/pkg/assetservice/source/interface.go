package source

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/sirupsen/logrus"
)

type AssetSource interface {
	Asset() chan dia.Asset
}

var log *logrus.Logger

func init() {
	log = logrus.New()
}
