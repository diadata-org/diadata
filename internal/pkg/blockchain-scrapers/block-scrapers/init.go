package blockscrapers

import (
	"github.com/sirupsen/logrus"
)

// Block-scrapers fetch information on blocks for various blockchains.
// For instance block numbers, timestamps, gas fees,...

var log *logrus.Logger

func init() {
	log = logrus.New()
}
