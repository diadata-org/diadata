package source

import (
	"context"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	alephiumhelper "github.com/diadata-org/diadata/pkg/dia/helpers/alephium-helper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
)

// Default values for asset collector
const (
	defaultSleepTimeout       = 200 // millisecond
	defaultSwapContractsLimit = 100
)

// AlphiumAssetSource aset collector object - which serves assetCollector command
type AlphiumAssetSource struct {
	// client - interaction with alephium REST API services
	alephiumClient *alephiumhelper.AlephiumClient
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
	// sleeps between alephium API calls - to avoid "Too many requests" exception from alephium REST API
	sleepTimeout int
	// swap contracts count limitation in alephium REST API
	swapContractsLimit int
}

// NewAlephiumAssetSource creates object to get alephium assets
// ENV values:
//
//	 	AYIN_ASSETS_SLEEP_TIMEOUT - (optional,millisecond), make timeout between API calls, default "defaultSleepTimeout" value
//		AYIN_SWAP_CONTRACTS_LIMIT - (optional, int), limit to get swap contact addresses, default "defaultSwapContractsLimit" value
//		AYIN_DEBUG - (optional, bool), make stdout output with alephium client http call, default = false
func NewAlephiumAssetSource(exchange dia.Exchange, relDB *models.RelDB) *AlphiumAssetSource {
	sleepTimeout := utils.GetenvInt(strings.ToUpper(exchange.Name)+"_ASSETS_SLEEP_TIMEOUT", defaultSleepTimeout)
	swapContractsLimit := utils.GetenvInt(strings.ToUpper(exchange.Name)+"_SWAP_CONTRACTS_LIMIT", defaultSwapContractsLimit)
	isDebug := utils.GetenvBool(strings.ToUpper(exchange.Name)+"_DEBUG", false)

	var (
		assetChannel = make(chan dia.Asset)
		doneChannel  = make(chan bool)
	)

	alephiumClient := alephiumhelper.NewAlephiumClient(
		log.WithContext(context.Background()).WithField("context", "AlephiumClient"),
		isDebug,
	)

	logger := log.
		WithContext(context.Background()).
		WithField("service", "assetCollector").
		WithField("network", exchange.BlockChain.Name)

	scraper := &AlphiumAssetSource{
		alephiumClient:     alephiumClient,
		assetChannel:       assetChannel,
		doneChannel:        doneChannel,
		blockchain:         exchange.BlockChain.Name,
		relDB:              relDB,
		sleepTimeout:       sleepTimeout,
		logger:             logger,
		swapContractsLimit: swapContractsLimit,
	}

	go func() {
		scraper.fetchAssets()
	}()
	return scraper
}

func (s *AlphiumAssetSource) fetchAssets() {
	s.logger.Info("started")

	contractAddresses, err := s.alephiumClient.GetSwapPairsContractAddresses(s.swapContractsLimit)
	if err != nil {
		s.logger.WithError(err).Error("failed to get contract addresses")
		return
	}

	uniqueAddressesMap := make(map[string]int)

	for _, contractAddress := range contractAddresses.SubContracts {
		tokenPairs, err := s.alephiumClient.GetTokenPairAddresses(contractAddress)

		if err != nil {
			s.logger.
				WithField("contractAddress", contractAddress).
				WithError(err).
				Error("failed to get GetTokenPairAddresses")
			continue
		}

		uniqueAddressesMap[tokenPairs[0]]++
		uniqueAddressesMap[tokenPairs[1]]++

		swapRelation := dia.SwapRelation{
			Blockchain:    s.blockchain,
			ParentAddress: contractAddress,
			AssetAddress0: tokenPairs[0],
			AssetAddress1: tokenPairs[1],
		}

		err = s.relDB.SetSwapRelation(swapRelation)
		if err != nil {
			s.logger.
				WithField("contractAddress", contractAddress).
				WithError(err).
				Error("failed to call SetSwapRelation")
			time.Sleep(time.Duration(s.sleepTimeout) * time.Millisecond)
			continue
		}

		s.logger.
			WithField("contractAddress", contractAddress).
			WithField("tokenPairs", tokenPairs).
			Info("token pairs")

		time.Sleep(time.Duration(s.sleepTimeout) * time.Millisecond)
	}

	uniqueAddresses := make([]string, 0)

	for k := range uniqueAddressesMap {
		uniqueAddresses = append(uniqueAddresses, k)
	}

	for _, contractAddress := range uniqueAddresses {
		if contractAddress == alephiumhelper.ALPHNativeToken.Address {
			alephiumhelper.ALPHNativeToken.Blockchain = s.blockchain
			s.assetChannel <- alephiumhelper.ALPHNativeToken
			continue
		}

		asset, err := s.alephiumClient.GetTokenInfoForContractDecoded(contractAddress, s.blockchain)
		if err != nil {
			s.logger.
				WithField("contractAddress", contractAddress).
				WithError(err).
				Error("failed to get GetTokenInfoForContract")
			continue
		}

		s.assetChannel <- *asset

		time.Sleep(time.Duration(s.sleepTimeout) * time.Millisecond)
	}
	s.doneChannel <- true
}

func (s *AlphiumAssetSource) Asset() chan dia.Asset {
	return s.assetChannel
}

func (s *AlphiumAssetSource) Done() chan bool {
	return s.doneChannel
}
