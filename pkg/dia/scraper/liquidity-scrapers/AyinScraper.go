package liquidityscrapers

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

type AyinLiquidityScraper struct {
	logger                    *logrus.Entry
	api                       *alephiumhelper.AlephiumClient
	poolChannel               chan dia.Pool
	doneChannel               chan bool
	blockchain                string
	exchangeName              string
	relDB                     *models.RelDB
	datastore                 *models.DB
	targetSwapContract        string
	swapContractsLimit        int
	handlerType               string
	sleepBetweenContractCalls time.Duration
}

// NewAyinLiquidityScraper returns a new AyinLiquidityScraper initialized with default values.
// The instance is asynchronously scraping as soon as it is created.
// ENV values:
//
//	 	AYIN_SLEEP_TIMEOUT - (optional,millisecond), make timeout between API calls, default "alephiumhelper.DefaultSleepBetweenContractCalls" value
//		AYIN_TARGET_SWAP_CONTRACT - (optional, string), useful for debug, default = ""
//		AYIN_DEBUG - (optional, bool), make stdout output with alephium client http call, default = false
func NewAyinLiquidityScraper(exchange dia.Exchange, relDB *models.RelDB, datastore *models.DB) *AyinLiquidityScraper {
	targetSwapContract := utils.Getenv(
		strings.ToUpper(exchange.Name)+"_TARGET_SWAP_CONTRACT",
		"",
	)
	isDebug := utils.GetenvBool(strings.ToUpper(exchange.Name)+"_DEBUG", false)
	sleepBetweenContractCalls := utils.GetTimeDurationFromIntAsMilliseconds(
		utils.GetenvInt(
			strings.ToUpper(exchange.Name)+"_SLEEP_TIMEOUT",
			alephiumhelper.DefaultSleepBetweenContractCalls,
		),
	)
	swapContractsLimit := utils.GetenvInt(
		strings.ToUpper(exchange.Name)+"_SWAP_CONTRACTS_LIMIT",
		alephiumhelper.DefaultSwapContractsLimit,
	)

	var (
		poolChannel = make(chan dia.Pool)
		doneChannel = make(chan bool)
		scraper     *AyinLiquidityScraper
	)

	alephiumClient := alephiumhelper.NewAlephiumClient(
		log.WithContext(context.Background()).WithField("context", "AlephiumClient"),
		sleepBetweenContractCalls,
		isDebug,
	)

	scraper = &AyinLiquidityScraper{
		api:                       alephiumClient,
		poolChannel:               poolChannel,
		doneChannel:               doneChannel,
		exchangeName:              exchange.Name,
		blockchain:                exchange.BlockChain.Name,
		relDB:                     relDB,
		datastore:                 datastore,
		targetSwapContract:        targetSwapContract,
		swapContractsLimit:        swapContractsLimit,
		handlerType:               "liquidity",
		sleepBetweenContractCalls: sleepBetweenContractCalls,
	}
	scraper.logger = logrus.
		New().
		WithContext(context.Background()).
		WithField("handlerType", scraper.handlerType).
		WithField("context", "AyinLiquidityScraper")

	go scraper.fetchPools()

	return scraper
}

func (s *AyinLiquidityScraper) fetchPools() {
	logger := s.logger.WithFields(logrus.Fields{
		"function": "fetchPools",
	})

	contractAddresses, err := s.getSubcontracts()
	if err != nil {
		s.logger.WithError(err).Error("failed to get contract addresses")
		return
	}
	log.Infof("fetched %v contract addresses.", len(contractAddresses))

	// Fetch all pool addresses from on-chain
	for _, contractAddress := range contractAddresses {
		var decimals []int64

		tokenPairs, err := s.api.GetTokenPairAddresses(contractAddress)
		if err != nil {
			s.logger.
				WithField("contractAddress", contractAddress).
				WithError(err).
				Error("failed to get GetTokenPairAddresses")
			continue
		}

		pool := dia.Pool{
			Exchange:     dia.Exchange{Name: s.exchangeName},
			Blockchain:   dia.BlockChain{Name: s.blockchain},
			Address:      contractAddress,
			Time:         time.Now(),
			Assetvolumes: make([]dia.AssetVolume, 2),
		}

		for i := 0; i < 2; i++ {
			asset, err := s.relDB.GetAsset(tokenPairs[i], dia.ALEPHIUM)
			if err != nil {
				log.Error("GetAsset from DB: ", err)
			}
			pool.Assetvolumes[i] = dia.AssetVolume{
				Index: uint8(i),
				Asset: asset,
			}
			decimals = append(decimals, int64(asset.Decimals))
		}

		contractState, err := s.api.GetContractState(contractAddress)
		if err != nil {
			logger.
				WithError(err).
				Error("failed to GetContractState")
			continue
		}

		pool.Assetvolumes[0].Volume, _ = utils.StringToFloat64(
			contractState.MutFields[0].Value,
			decimals[0],
		)
		pool.Assetvolumes[1].Volume, _ = utils.StringToFloat64(
			contractState.MutFields[1].Value,
			decimals[1],
		)

		// Determine USD liquidity.
		if pool.SufficientNativeBalance(GLOBAL_NATIVE_LIQUIDITY_THRESHOLD) {
			s.datastore.GetPoolLiquiditiesUSD(&pool, priceCache)
		}
		s.poolChannel <- pool

	}

	s.doneChannel <- true
}

func (s *AyinLiquidityScraper) getSubcontracts() ([]string, error) {
	if s.targetSwapContract != "" {
		return []string{s.targetSwapContract}, nil
	}
	contractAddresses, err := s.api.GetSwapPairsContractAddresses(s.swapContractsLimit)
	return contractAddresses.SubContracts, err
}

func (s *AyinLiquidityScraper) Pool() chan dia.Pool {
	return s.poolChannel
}

func (s *AyinLiquidityScraper) Done() chan bool {
	return s.doneChannel
}
