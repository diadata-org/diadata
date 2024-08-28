package liquidityscrapers

import (
	"context"
	"strconv"
	"strings"
	"time"
	"unicode"

	// "unicode"

	"github.com/diadata-org/diadata/pkg/dia"
	bifrosthelper "github.com/diadata-org/diadata/pkg/dia/helpers/bifrost-helper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
)

type BifrostLiquidityScraper struct {
	logger                    *logrus.Entry
	api                       *bifrosthelper.BifrostClient
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

// NewBifrostLiquidityScraper returns a new BifrostLiquidityScraper initialized with default values.
// The instance is asynchronously scraping as soon as it is created.
// ENV values:
//
//	 	BIFROST_SLEEP_TIMEOUT - (optional,millisecond), make timeout between API calls, default "bifrosthelper.DefaultSleepBetweenContractCalls" value
//		BIFROST_TARGET_SWAP_CONTRACT - (optional, string), useful for debug, default = ""
//		BIFROST_DEBUG - (optional, bool), make stdout output with bifrost client http call, default = false
func NewBifrostLiquidityScraper(exchange dia.Exchange, relDB *models.RelDB, datastore *models.DB) *BifrostLiquidityScraper {
	targetSwapContract := utils.Getenv(
		strings.ToUpper(exchange.Name)+"_TARGET_SWAP_CONTRACT",
		"",
	)
	isDebug := utils.GetenvBool(strings.ToUpper(exchange.Name)+"_DEBUG", false)
	sleepBetweenContractCalls := utils.GetTimeDurationFromIntAsMilliseconds(
		utils.GetenvInt(
			strings.ToUpper(exchange.Name)+"_SLEEP_TIMEOUT",
			bifrosthelper.DefaultSleepBetweenContractCalls,
		),
	)
	swapContractsLimit := utils.GetenvInt(
		strings.ToUpper(exchange.Name)+"_SWAP_CONTRACTS_LIMIT",
		bifrosthelper.DefaultSwapContractsLimit,
	)

	var (
		poolChannel = make(chan dia.Pool)
		doneChannel = make(chan bool)
		scraper     *BifrostLiquidityScraper
	)

	bifrostClient := bifrosthelper.NewBifrostClient(
		log.WithContext(context.Background()).WithField("context", "BifrostClient"),
		sleepBetweenContractCalls,
		isDebug,
	)

	scraper = &BifrostLiquidityScraper{
		api:                       bifrostClient,
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
		WithField("context", "BifrostLiquidityScraper")

	go scraper.fetchPools()

	return scraper
}

// func detectPrecision(input string) (uint8, error) {
// 	normalized := normalizeNumber(input)
// 	decimalIndex := strings.Index(normalized, ".")
// 	if decimalIndex == -1 {
// 		return 0, nil
// 	}

// 	precision := len(normalized) - decimalIndex - 1
// 	return uint8(precision), nil
// }

func normalizeNumber(input string) string {
	if strings.Contains(input, ",") && strings.Contains(input, ".") {
		input = strings.ReplaceAll(input, ",", "")
	} else if strings.Count(input, ",") > 1 {
		input = strings.ReplaceAll(input, ",", "")
	} else if strings.Count(input, ".") > 1 {
		input = strings.ReplaceAll(input, ".", "")
		input = strings.Replace(input, ",", ".", 1)
	} else if strings.Contains(input, ",") {
		input = strings.Replace(input, ",", ".", 1)
	}

	// Remove any remaining non-numeric characters (e.g., spaces)
	return strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) || r == '.' {
			return r
		}
		return -1
	}, input)
}

func (s *BifrostLiquidityScraper) fetchPools() {
	logger := s.logger.WithFields(logrus.Fields{
		"function": "fetchPools",
	})

	// Fetch all pair tokens pool entries from api
	bifrostPoolAssets, err := s.api.GetAllPoolAssets()
	if err != nil {
		s.logger.WithError(err).Error("failed to GetAllPoolAssets")
	}

	s.logger.Infof("Found %d pools", len(bifrostPoolAssets))

	// Iterate all over the pool assets creating pool objects and sending them to the pool channel
	for _, bPool := range bifrostPoolAssets {
		dbAssets := make([]dia.Asset, 0)
		for _, assetId := range bPool.Assets {
			assetKey := "Bifrost:Asset:" + assetId
			dbTokenInfo, err := s.relDB.GetAsset(assetKey, s.blockchain)
			if err != nil {
				logger.WithError(err).Error("Failed to GetAsset with key: ", assetKey)
			}

			dbAssets = append(dbAssets, dbTokenInfo)
		}

		if len(dbAssets) != 2 {
			logger.Error("found more than 2 asset types for the pool pair")
			continue
		}

		if len(bPool.Balances) != 2 {
			logger.Error("found more than 2 balances for the pool pair")
			continue
		}

		tokenABalance, err := strconv.ParseFloat(normalizeNumber(bPool.Balances[0]), 64)
		if err != nil {
			logger.WithError(err).Error("failed to parse tokenA balance")
			continue
		}

		tokenBBalance, err := strconv.ParseFloat(normalizeNumber(bPool.Balances[1]), 64)
		if err != nil {
			logger.WithError(err).Error("failed to parse tokenB balance")
			continue
		}

		s.logger.WithFields(logrus.Fields{
			"BalanceA": tokenABalance,
			"BalanceB": tokenBBalance,
		}).Info("Found balances")

		tokenA := dia.AssetVolume{
			Index:  0,
			Asset:  dbAssets[0],
			Volume: tokenABalance,
		}

		tokenB := dia.AssetVolume{
			Index:  1,
			Asset:  dbAssets[1],
			Volume: tokenBBalance,
		}

		pool := dia.Pool{
			Exchange:     dia.Exchange{Name: s.exchangeName},
			Blockchain:   dia.BlockChain{Name: s.blockchain},
			Address:      "Polkadot:Bifrost:" + bPool.Assets[0] + ":" + bPool.Assets[1],
			Time:         time.Now(),
			Assetvolumes: []dia.AssetVolume{tokenA, tokenB},
		}

		// Determine USD liquidity.
		if pool.SufficientNativeBalance(GLOBAL_NATIVE_LIQUIDITY_THRESHOLD) {
			s.datastore.GetPoolLiquiditiesUSD(&pool, priceCache)
		}

		s.logger.WithFields(logrus.Fields{
			"Address": pool.Address,
		}).Info("sending pool to poolChannel")

		s.poolChannel <- pool
	}

	s.doneChannel <- true
}

func (s *BifrostLiquidityScraper) Pool() chan dia.Pool {
	return s.poolChannel
}

func (s *BifrostLiquidityScraper) Done() chan bool {
	return s.doneChannel
}
