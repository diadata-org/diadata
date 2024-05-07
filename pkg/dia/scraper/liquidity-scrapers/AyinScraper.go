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

func (s *AyinLiquidityScraper) getPools() ([]dia.Pool, error) {
	if s.targetSwapContract != "" {
		result := make([]dia.Pool, 1)
		pool, err := s.relDB.GetPoolByAddress(s.blockchain, s.targetSwapContract)
		result[0] = pool
		return result, err
	}
	return s.relDB.GetAllPoolsExchange(s.exchangeName, 0)
}

func (s *AyinLiquidityScraper) fetchPools() {
	logger := s.logger.WithFields(logrus.Fields{
		"function": "fetchPools",
	})
	pools, err := s.getPools()
	if err != nil {
		logger.
			WithError(err).
			Error("failed to GetAllPoolsExchange")
		return
	}
	for _, pool := range pools {
		decimals0 := int64(pool.Assetvolumes[0].Asset.Decimals)
		decimals1 := int64(pool.Assetvolumes[1].Asset.Decimals)

		contractState, err := s.api.GetContractState(pool.Address)
		if err != nil {
			logger.
				WithError(err).
				Error("failed to GetContractState")
			continue
		}

		pool.Assetvolumes[0].Volume, _ = utils.StringToFloat64(
			contractState.MutFields[0].Value,
			decimals0,
		)
		pool.Assetvolumes[1].Volume, _ = utils.StringToFloat64(
			contractState.MutFields[1].Value,
			decimals1,
		)
		// Determine USD liquidity.
		if pool.SufficientNativeBalance(GLOBAL_NATIVE_LIQUIDITY_THRESHOLD) {
			s.datastore.GetPoolLiquiditiesUSD(&pool, priceCache)
		}

		s.poolChannel <- pool
	} // END for _, pool := range pools

	s.doneChannel <- true
}

func (s *AyinLiquidityScraper) Pool() chan dia.Pool {
	return s.poolChannel
}

func (s *AyinLiquidityScraper) Done() chan bool {
	return s.doneChannel
}
