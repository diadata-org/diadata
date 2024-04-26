package source

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	alephiumhelper "github.com/diadata-org/diadata/pkg/dia/helpers/alephium-helper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
)

const (
	defaultAssetRequestLimit = 100
	multicallPageSize        = 20
	defaultSleepTimeout      = 200 // millisecond
	defaultContractsLimit    = 100
)

type AlphiumAssetSource struct {
	alephiumClient *alephiumhelper.AlephiumClient
	assetChannel   chan dia.Asset
	doneChannel    chan bool
	blockchain     string
	limit          uint
	relDB          *models.RelDB
	logger         *logrus.Entry
	sleepTimeout   int
}

func NewAlphiumAssetSource(exchange dia.Exchange, relDB *models.RelDB) *AlphiumAssetSource {
	limit := utils.GetenvUint(strings.ToUpper(exchange.Name)+"_ASSETS_LIMIT", defaultAssetRequestLimit)
	sleepTimeout := utils.GetenvInt(strings.ToUpper(exchange.Name)+"_ASSETS_SLEEP_TIMEOUT", defaultSleepTimeout)
	contractsLimit := utils.GetenvInt(strings.ToUpper(exchange.Name)+"_ASSETS_CONTRACTS_LIMIT", defaultContractsLimit)
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
		alephiumClient: alephiumClient,
		assetChannel:   assetChannel,
		doneChannel:    doneChannel,
		blockchain:     exchange.BlockChain.Name,
		limit:          limit,
		relDB:          relDB,
		sleepTimeout:   sleepTimeout,
		logger:         logger,
	}

	go func() {
		scraper.fetchAssets(contractsLimit)
	}()
	return scraper
}

func (s *AlphiumAssetSource) fetchAssets(contractsLimit int) {
	s.logger.Info("started")

	contractAddresses, err := s.alephiumClient.GetSwapPairsContractAddresses(contractsLimit)
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
		tokens, err := s.alephiumClient.GetTokenInfoForContract(contractAddress)
		if err != nil {
			s.logger.
				WithField("contractAddress", contractAddress).
				WithError(err).
				Error("failed to get GetTokenInfoForContract")
			continue
		}

		asset, err := s.decodeMulticallRequestToAssets(contractAddress, tokens)

		s.assetChannel <- asset

		time.Sleep(time.Duration(s.sleepTimeout) * time.Millisecond)
	}
	s.doneChannel <- true
}

func (s *AlphiumAssetSource) decodeMulticallRequestToAssets(contractAddress string, resp *alephiumhelper.OutputResult) (dia.Asset, error) {
	asset := dia.Asset{}

	symbol, err := alephiumhelper.DecodeHex(resp.Results[alephiumhelper.SymbolMethod].Value)
	if err != nil {
		s.logger.
			WithField("row", resp).
			WithError(err).
			Error("failed to decode symbol")
		return asset, err
	}
	asset.Symbol = symbol

	name, err := alephiumhelper.DecodeHex(resp.Results[alephiumhelper.NameMethod].Value)
	if err != nil {
		s.logger.
			WithField("row", resp).
			WithError(err).
			Error("failed to decode name")
		return asset, err
	}
	asset.Name = name

	decimals, err := strconv.ParseUint(resp.Results[alephiumhelper.DecimalsMethod].Value, 10, 32)
	if err != nil {
		s.logger.
			WithField("row", resp).
			WithError(err).
			Error("failed to decode decimals")
		return asset, err
	}
	asset.Decimals = uint8(decimals)
	asset.Address = contractAddress
	asset.Blockchain = s.blockchain

	return asset, nil
}

func (s *AlphiumAssetSource) Asset() chan dia.Asset {
	return s.assetChannel
}

func (s *AlphiumAssetSource) Done() chan bool {
	return s.doneChannel
}
