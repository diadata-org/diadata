package source

import (
	"context"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	hydrationhelper "github.com/diadata-org/diadata/pkg/dia/helpers/hydration-helper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
)

// HydrationAssetSource asset collector object - which serves assetCollector command
type HydrationAssetSource struct {
	// client - interaction with hydration REST API services
	hydrationClient *hydrationhelper.HydrationClient
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
	// swap contracts count limitation in hydration REST API
	swapContractsLimit int

	sleepTimeout       time.Duration
	exchangeName       string
	targetSwapContract string
}

// NewHydrationAssetSource creates object to get hydration assets
// ENV values:
//
//	 	BIFROST_ASSETS_SLEEP_TIMEOUT - (optional,millisecond), make timeout between API calls, default "hydrationhelper.DefaultSleepBetweenContractCalls" value
//		BIFROST_SWAP_CONTRACTS_LIMIT - (optional, int), limit to get swap contact addresses, default "hydrationhelper.DefaultSwapContractsLimit" value
//		BIFROST_TARGET_SWAP_CONTRACT - (optional, string), useful for debug, default = ""
//		BIFROST_DEBUG - (optional, bool), make stdout output with hydration client http call, default = false
func NewHydrationAssetSource(exchange dia.Exchange, relDB *models.RelDB) *HydrationAssetSource {
	println("assetService::source::hydration exchange: ", exchange.Name)
	sleepBetweenContractCalls := utils.GetTimeDurationFromIntAsMilliseconds(
		utils.GetenvInt(strings.ToUpper(exchange.Name)+"_SLEEP_TIMEOUT", hydrationhelper.DefaultSleepBetweenContractCalls),
	)
	swapContractsLimit := utils.GetenvInt(
		strings.ToUpper(exchange.Name)+"_SWAP_CONTRACTS_LIMIT",
		hydrationhelper.DefaultSwapContractsLimit,
	)
	targetSwapContract := utils.Getenv(strings.ToUpper(exchange.Name)+"_TARGET_SWAP_CONTRACT", "")
	isDebug := utils.GetenvBool(strings.ToUpper(exchange.Name)+"_DEBUG", false)

	var (
		assetChannel = make(chan dia.Asset)
		doneChannel  = make(chan bool)
	)

	println("assetService::source::hydration connecting to Hydration RPC")
	hydrationClient := hydrationhelper.NewHydrationClient(
		log.WithContext(context.Background()).WithField("context", "HydrationClient"),
		sleepBetweenContractCalls,
		isDebug,
	)

	println("assetService::source::hydration connected to Hydration RPC with success!!")
	
	logger := log.
		WithContext(context.Background()).
		WithField("service", "assetCollector").
		WithField("network", exchange.BlockChain.Name)

	scraper := &HydrationAssetSource{
		hydrationClient:      hydrationClient,
		assetChannel:       assetChannel,
		doneChannel:        doneChannel,
		blockchain:         exchange.BlockChain.Name,
		relDB:              relDB,
		logger:             logger,
		swapContractsLimit: swapContractsLimit,
		exchangeName:       exchange.Name,
		sleepTimeout:       sleepBetweenContractCalls,
		targetSwapContract: targetSwapContract,
	}

	go scraper.fetchAssets()

	return scraper
}

func (s *HydrationAssetSource) fetchAssets() {
	s.logger.Info("Scraping assets...")

	// fetch all assets to send in the channel
	assets, err := s.hydrationClient.ScrapAssets()
	if err != nil {
		s.logger.Error("Error when scraping assets: ", err)
		return
	}

	for _, asset := range assets {
		s.assetChannel <- *asset
	}
	s.doneChannel <- true
}

func (s *HydrationAssetSource) Asset() chan dia.Asset {
	return s.assetChannel
}

func (s *HydrationAssetSource) Done() chan bool {
	return s.doneChannel
}
