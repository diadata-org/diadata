package scrapers

import (
	"time"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

const (
	duplicateTradesMemory        = 2 * time.Second
	duplicateTradesScanFrequency = 1 * time.Second
)

func init() {
	log = logrus.New()
}
