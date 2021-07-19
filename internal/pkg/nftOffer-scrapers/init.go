package nftofferscrapers

import (
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

const (
	blockDelayEthereum = 8
)

func init() {
	log = logrus.New()
}
