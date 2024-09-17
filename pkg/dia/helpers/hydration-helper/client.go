package hydrationhelper

import (
	"github.com/sirupsen/logrus"
	"time"
)

const (
	AssetAddressURI = "AssetRegistry:Assets"
	Blockchain      = "polkadot"
	DiaPolkadotApi  = "http://localhost:3000/hydration/v1"
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
