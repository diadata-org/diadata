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

// AlphiumAssetSource aset collector object - which serves assetCollector command
type AyinAssetSource struct {
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
	// swap contracts count limitation in alephium REST API
	swapContractsLimit int

	sleepTimeout       time.Duration
	exchangeName       string
	targetSwapContract string
}

// NewAlephiumAssetSource creates object to get alephium assets
// ENV values:
//
//	 	AYIN_ASSETS_SLEEP_TIMEOUT - (optional,millisecond), make timeout between API calls, default "alephiumhelper.DefaultSleepBetweenContractCalls" value
//		AYIN_SWAP_CONTRACTS_LIMIT - (optional, int), limit to get swap contact addresses, default "alephiumhelper.DefaultSwapContractsLimit" value
//		AYIN_TARGET_SWAP_CONTRACT - (optional, string), useful for debug, default = ""
//		AYIN_DEBUG - (optional, bool), make stdout output with alephium client http call, default = false
func NewAyinAssetSource(exchange dia.Exchange, relDB *models.RelDB) *AyinAssetSource {
	sleepBetweenContractCalls := utils.GetTimeDurationFromIntAsMilliseconds(
		utils.GetenvInt(strings.ToUpper(exchange.Name)+"_SLEEP_TIMEOUT", alephiumhelper.DefaultSleepBetweenContractCalls),
	)
	swapContractsLimit := utils.GetenvInt(
		strings.ToUpper(exchange.Name)+"_SWAP_CONTRACTS_LIMIT",
		alephiumhelper.DefaultSwapContractsLimit,
	)
	targetSwapContract := utils.Getenv(strings.ToUpper(exchange.Name)+"_TARGET_SWAP_CONTRACT", "")
	isDebug := utils.GetenvBool(strings.ToUpper(exchange.Name)+"_DEBUG", false)

	var (
		assetChannel = make(chan dia.Asset)
		doneChannel  = make(chan bool)
	)

	alephiumClient := alephiumhelper.NewAlephiumClient(
		log.WithContext(context.Background()).WithField("context", "AlephiumClient"),
		sleepBetweenContractCalls,
		isDebug,
	)

	logger := log.
		WithContext(context.Background()).
		WithField("service", "assetCollector").
		WithField("network", exchange.BlockChain.Name)

	scraper := &AyinAssetSource{
		alephiumClient:     alephiumClient,
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

func (s *AyinAssetSource) getSubcontracts() ([]string, error) {
	if s.targetSwapContract != "" {
		return []string{s.targetSwapContract}, nil
	}
	contractAddresses, err := s.alephiumClient.GetSwapPairsContractAddresses(s.swapContractsLimit)
	return contractAddresses.SubContracts, err
}

func (s *AyinAssetSource) fetchAssets() {
	s.logger.Info("started")

	contractAddresses, err := s.getSubcontracts()
	if err != nil {
		s.logger.WithError(err).Error("failed to get contract addresses")
		return
	}

	pools := make(map[string]dia.Pool)

	for _, contractAddress := range contractAddresses {
		tokenPairs, err := s.alephiumClient.GetTokenPairAddresses(contractAddress)

		if err != nil {
			s.logger.
				WithField("contractAddress", contractAddress).
				WithError(err).
				Error("failed to get GetTokenPairAddresses")
			continue
		}

		pools[contractAddress] = dia.Pool{
			Exchange:     dia.Exchange{Name: s.exchangeName},
			Blockchain:   dia.BlockChain{Name: s.blockchain},
			Address:      contractAddress,
			Assetvolumes: make([]dia.AssetVolume, 2),
		}

		for i := 0; i < 2; i++ {

			// native ALPH token has no contract - use data from variable
			if tokenPairs[i] == alephiumhelper.ALPHNativeToken.Address {
				pools[contractAddress].Assetvolumes[i].Asset = alephiumhelper.ALPHNativeToken
				pools[contractAddress].Assetvolumes[i].Asset.Blockchain = s.blockchain

				s.assetChannel <- pools[contractAddress].Assetvolumes[i].Asset

				continue
			}

			asset, err := s.alephiumClient.GetTokenInfoForContractDecoded(tokenPairs[i], s.blockchain)
			if err != nil {
				s.logger.
					WithField("contractAddress", contractAddress).
					WithField("tokenPairs[i]", tokenPairs[i]).
					WithError(err).
					Error("failed to get GetTokenInfoForContract")

				delete(pools, contractAddress)
				continue
			}
			pools[contractAddress].Assetvolumes[i].Asset = *asset

			s.assetChannel <- *asset
		}
	}

	s.doneChannel <- true
}

func (s *AyinAssetSource) Asset() chan dia.Asset {
	return s.assetChannel
}

func (s *AyinAssetSource) Done() chan bool {
	return s.doneChannel
}
