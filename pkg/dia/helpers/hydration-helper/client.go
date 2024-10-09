package hydrationhelper

import (
	"github.com/sirupsen/logrus"
	"time"
)

const (
	AssetAddressURI = "AssetRegistry:Assets"
	Blockchain      = "polkadot"
	GetAssetsPath   = "assets"
)

type HydrationClient struct {
	logger            *logrus.Entry
	sleepBetweenCalls time.Duration
	debug             bool
}

func NewHydrationClient(logger *logrus.Entry, sleepBetweenCalls time.Duration, isDebug bool) *HydrationClient {
	return &HydrationClient{
		logger:            logger,
		sleepBetweenCalls: sleepBetweenCalls,
		debug:             isDebug,
	}
}
