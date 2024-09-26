package source

import (
	"context"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	bifrosthelper "github.com/diadata-org/diadata/pkg/dia/helpers/bifrost-helper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
)

// BifrostAssetSource asset collector object - which serves assetCollector command
type BifrostAssetSource struct {
	// client - interaction with bifrost REST API services
	bifrostClient *bifrosthelper.BifrostClient
	// channel to store received asset info
	assetChannel chan dia.Asset
	// channel which informs about work is finished
	doneChannel chan bool
	// blockchain name
	blockchain string
	// DB connector to interact with databases
	relDB *models.RelDB
	// logs all events here
	logger *logrus.Entry
	// swap contracts count limitation in bifrost REST API
	swapContractsLimit int

	sleepTimeout       time.Duration
	exchangeName       string
	targetSwapContract string
}

// NewBifrostAssetSource creates object to get bifrost assets
// ENV values:
//
//	 	BIFROST_ASSETS_SLEEP_TIMEOUT - (optional,millisecond), make timeout between API calls, default "bifrosthelper.DefaultSleepBetweenContractCalls" value
//		BIFROST_SWAP_CONTRACTS_LIMIT - (optional, int), limit to get swap contact addresses, default "bifrosthelper.DefaultSwapContractsLimit" value
//		BIFROST_TARGET_SWAP_CONTRACT - (optional, string), useful for debug, default = ""
//		BIFROST_DEBUG - (optional, bool), make stdout output with bifrost client http call, default = false
func NewBifrostAssetSource(exchange dia.Exchange, relDB *models.RelDB) *BifrostAssetSource {
	sleepBetweenContractCalls := utils.GetTimeDurationFromIntAsMilliseconds(
		utils.GetenvInt(strings.ToUpper(exchange.Name)+"_SLEEP_TIMEOUT", bifrosthelper.DefaultSleepBetweenContractCalls),
	)
	swapContractsLimit := utils.GetenvInt(
		strings.ToUpper(exchange.Name)+"_SWAP_CONTRACTS_LIMIT",
		bifrosthelper.DefaultSwapContractsLimit,
	)
	targetSwapContract := utils.Getenv(strings.ToUpper(exchange.Name)+"_TARGET_SWAP_CONTRACT", "")
	isDebug := utils.GetenvBool(strings.ToUpper(exchange.Name)+"_DEBUG", false)

	var (
		assetChannel = make(chan dia.Asset)
		doneChannel  = make(chan bool)
	)
	bifrostClient := bifrosthelper.NewBifrostClient(
		log.WithContext(context.Background()).WithField("context", "BifrostClient"),
		sleepBetweenContractCalls,
		isDebug,
	)
	
	logger := log.
		WithContext(context.Background()).
		WithField("service", "assetCollector").
		WithField("network", "Bifrost")

	scraper := &BifrostAssetSource{
		bifrostClient:      bifrostClient,
		assetChannel:       assetChannel,
		doneChannel:        doneChannel,
		blockchain:         exchange.BlockChain.Name,
		relDB:              relDB,
		logger:             logger,
		swapContractsLimit: swapContractsLimit,
		exchangeName:       "Bifrost",
		sleepTimeout:       sleepBetweenContractCalls,
		targetSwapContract: targetSwapContract,
	}

	go scraper.fetchAssets()

	return scraper
}

func (s *BifrostAssetSource) fetchAssets() {
	s.logger.Info("Scraping assets...")

	assets, err := s.bifrostClient.ScrapAssets()
	if err != nil {
		s.logger.Error("Error when scraping assets: ", err)
		return
	}

	for _, asset := range assets {
		s.assetChannel <- *asset
	}
	s.doneChannel <- true
}

func (s *BifrostAssetSource) Asset() chan dia.Asset {
	return s.assetChannel
}

func (s *BifrostAssetSource) Done() chan bool {
	return s.doneChannel
}
