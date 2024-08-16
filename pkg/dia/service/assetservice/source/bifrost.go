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
	println("assetService::source::bifrost exchange: ", exchange.Name)
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

	println("assetService::source::bifrost connecting to Bifrost RPC")
	bifrostClient := bifrosthelper.NewBifrostClient(
		log.WithContext(context.Background()).WithField("context", "BifrostClient"),
		sleepBetweenContractCalls,
		isDebug,
	)

	println("assetService::source::bifrost connected to Bifrost RPC with success!!")
	
	logger := log.
		WithContext(context.Background()).
		WithField("service", "assetCollector").
		WithField("network", exchange.BlockChain.Name)

	scraper := &BifrostAssetSource{
		bifrostClient:      bifrostClient,
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

func (s *BifrostAssetSource) getSubcontracts() ([]string, error) {
	if s.targetSwapContract != "" {
		return []string{s.targetSwapContract}, nil
	}

	s.bifrostClient.ShowPalletes()
	// s.bifrostClient.GetStorageKeys("VtokenMinting")
	// s.bifrostClient.ScrapAssets()

	// contractAddresses, err := s.bifrostClient.GetSwapPairsContractAddresses(s.swapContractsLimit)
	// if err != nil {
	// 	return nil, err
	// }
	// s.logger.Info(contractAddresses)
	return []string{}, nil
}

func (s *BifrostAssetSource) fetchAssets() {
	s.logger.Info("started fetchAssets")

	s.getSubcontracts()
	// contractAddresses, err := s.getSubcontracts()
	// if err != nil {
	// 	s.logger.WithError(err).Error("failed to get contract addresses")
	// 	return
	// }

	// pools := make(map[string]dia.Pool)

	// for _, contractAddress := range contractAddresses {
	// 	tokenPairs, err := s.bifrostClient.GetTokenPairAddresses(contractAddress)

	// 	if err != nil {
	// 		s.logger.
	// 			WithField("contractAddress", contractAddress).
	// 			WithError(err).
	// 			Error("failed to get GetTokenPairAddresses")
	// 		continue
	// 	}

	// 	pools[contractAddress] = dia.Pool{
	// 		Exchange:     dia.Exchange{Name: s.exchangeName},
	// 		Blockchain:   dia.BlockChain{Name: s.blockchain},
	// 		Address:      contractAddress,
	// 		Assetvolumes: make([]dia.AssetVolume, 2),
	// 	}

	// 	for i := 0; i < 2; i++ {

	// 		// native ALPH token has no contract - use data from variable
	// 		if tokenPairs[i] == bifrosthelper.ALPHNativeToken.Address {
	// 			pools[contractAddress].Assetvolumes[i].Asset = bifrosthelper.ALPHNativeToken
	// 			pools[contractAddress].Assetvolumes[i].Asset.Blockchain = s.blockchain

	// 			s.assetChannel <- pools[contractAddress].Assetvolumes[i].Asset

	// 			continue
	// 		}

	// 		asset, err := s.bifrostClient.GetTokenInfoForContractDecoded(tokenPairs[i], s.blockchain)
	// 		if err != nil {
	// 			s.logger.
	// 				WithField("contractAddress", contractAddress).
	// 				WithField("tokenPairs[i]", tokenPairs[i]).
	// 				WithError(err).
	// 				Error("failed to get GetTokenInfoForContract")

	// 			delete(pools, contractAddress)
	// 			continue
	// 		}
	// 		pools[contractAddress].Assetvolumes[i].Asset = *asset

	// 		s.assetChannel <- *asset
	// 	}
	// }

	// s.doneChannel <- true
}

func (s *BifrostAssetSource) Asset() chan dia.Asset {
	return s.assetChannel
}

func (s *BifrostAssetSource) Done() chan bool {
	return s.doneChannel
}
