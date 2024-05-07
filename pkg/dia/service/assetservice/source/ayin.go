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
type AyinAssetSource struct {
	// client - interaction with alephium REST API services
	alephiumClient *alephiumhelper.AlephiumClient
	// channel to store received asset info
	assetChannel chan dia.Asset
	// channel which informs about work is finished
	doneChannel chan bool
	// channel to set data to Pool
	poolChannel chan dia.Pool
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

	exchangeName string
}

// NewAlephiumAssetSource creates object to get alephium assets
// ENV values:
//
//	 	AYIN_ASSETS_SLEEP_TIMEOUT - (optional,millisecond), make timeout between API calls, default "defaultSleepTimeout" value
//		AYIN_SWAP_CONTRACTS_LIMIT - (optional, int), limit to get swap contact addresses, default "defaultSwapContractsLimit" value
//		AYIN_DEBUG - (optional, bool), make stdout output with alephium client http call, default = false
func NewAyinAssetSource(exchange dia.Exchange, relDB *models.RelDB) *AyinAssetSource {
	sleepTimeout := utils.GetenvInt(strings.ToUpper(exchange.Name)+"_ASSETS_SLEEP_TIMEOUT", defaultSleepTimeout)
	swapContractsLimit := utils.GetenvInt(strings.ToUpper(exchange.Name)+"_SWAP_CONTRACTS_LIMIT", defaultSwapContractsLimit)
	isDebug := utils.GetenvBool(strings.ToUpper(exchange.Name)+"_DEBUG", false)

	var (
		assetChannel = make(chan dia.Asset)
		doneChannel  = make(chan bool)
		poolChannel  = make(chan dia.Pool)
	)

	alephiumClient := alephiumhelper.NewAlephiumClient(
		log.WithContext(context.Background()).WithField("context", "AlephiumClient"),
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
		sleepTimeout:       sleepTimeout,
		logger:             logger,
		swapContractsLimit: swapContractsLimit,
		exchangeName:       exchange.Name,
		poolChannel:        poolChannel,
	}

	go scraper.fetchAssets()

	go scraper.poolListener()

	return scraper
}

func (s *AyinAssetSource) fetchAssets() {
	s.logger.Info("started")

	contractAddresses, err := s.alephiumClient.GetSwapPairsContractAddresses(s.swapContractsLimit)
	if err != nil {
		s.logger.WithError(err).Error("failed to get contract addresses")
		return
	}

	pools := make(map[string]dia.Pool)

	for _, contractAddress := range contractAddresses.SubContracts {
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
			pools[contractAddress].Assetvolumes[i] = dia.AssetVolume{
				Asset: dia.Asset{
					Address: tokenPairs[i],
				},
			}
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

			s.waiting()
		} // END for i := 0; i < 2; i++
		if _, ok := pools[contractAddress]; ok {
			s.poolChannel <- pools[contractAddress]
		}

		s.waiting()
	} // END for _, contractAddress := range contractAddresses.SubContracts {

	s.doneChannel <- true
}

func (s *AyinAssetSource) Asset() chan dia.Asset {
	return s.assetChannel
}

func (s *AyinAssetSource) Done() chan bool {
	return s.doneChannel
}

func (s *AyinAssetSource) waiting() {
	time.Sleep(time.Duration(s.sleepTimeout) * time.Millisecond)
}

func (s *AyinAssetSource) poolListener() {
	logger := s.logger.WithField("function", "poolListener")
	for {
		select {
		case receivedPool := <-s.poolChannel:
			// Set to persistent DB.
			err := s.relDB.SetPool(receivedPool)
			if err != nil {
				logger.
					WithField("receivedPool", receivedPool).
					Errorf("Error saving pool %v", err)
			} else {
				logger.
					WithField("receivedPool", receivedPool).
					Info("successfully set pool")
			}

		case <-s.Done():
			return
		}
	}
}