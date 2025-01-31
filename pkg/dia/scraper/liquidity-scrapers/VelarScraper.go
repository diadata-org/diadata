package liquidityscrapers

import (
	"context"
	"encoding/hex"
	"errors"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/stackshelper"
	"github.com/diadata-org/diadata/pkg/dia/helpers/velarhelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
)

type VelarLiquidityScraper struct {
	logger                *logrus.Entry
	api                   *stackshelper.StacksClient
	velarClient           *velarhelper.VelarClient
	poolChannel           chan dia.Pool
	doneChannel           chan bool
	blockchain            string
	exchangeName          string
	sleepTimeMilliseconds int
	relDB                 *models.RelDB
	datastore             *models.DB
	handlerType           string
}

// NewVelarLiquidityScraper returns a new VelarLiquidityScraper initialized with default values.
// The instance is asynchronously scraping as soon as it is created.
// ENV values:
//
//	 	VELAR_SLEEP_TIMEOUT - (optional, millisecond), make timeout between API calls, default "stackshelper.DefaultSleepBetweenCalls" value
//		VELAR_HIRO_API_KEY - (optional, string), Hiro Stacks API key, improves scraping performance, default = ""
//		VELAR_DEBUG - (optional, bool), make stdout output with velar client http call, default = false
func NewVelarLiquidityScraper(exchange dia.Exchange, relDB *models.RelDB, datastore *models.DB) *VelarLiquidityScraper {
	envPrefix := strings.ToUpper(exchange.Name)

	sleepBetweenCalls := utils.GetenvInt(envPrefix+"_SLEEP_TIMEOUT", stackshelper.DefaultSleepBetweenCalls)
	hiroAPIKey := utils.Getenv(envPrefix+"_HIRO_API_KEY", "")
	isDebug := utils.GetenvBool(envPrefix+"_DEBUG", false)

	stacksClient := stackshelper.NewStacksClient(
		log.WithContext(context.Background()).WithField("context", "StacksClient"),
		utils.GetTimeDurationFromIntAsMilliseconds(sleepBetweenCalls),
		hiroAPIKey,
		isDebug,
	)

	velarClient := velarhelper.NewVelarClient(
		log.WithContext(context.Background()).WithField("context", "VelarClient"),
		isDebug,
	)

	sleepTime, err := strconv.Atoi(utils.Getenv("SLEEP_TIME_MILLISECONDS", "1000"))
	if err != nil {
		log.Error("parse SLEEP_TIME_MILLISECONDS: ", err)
	}
	s := &VelarLiquidityScraper{
		poolChannel:           make(chan dia.Pool),
		doneChannel:           make(chan bool),
		exchangeName:          exchange.Name,
		blockchain:            exchange.BlockChain.Name,
		sleepTimeMilliseconds: sleepTime,
		api:                   stacksClient,
		velarClient:           velarClient,
		relDB:                 relDB,
		datastore:             datastore,
		handlerType:           "liquidity",
	}

	s.logger = logrus.
		New().
		WithContext(context.Background()).
		WithField("handlerType", s.handlerType).
		WithField("context", "VelarLiquidityScraper")

	go s.fetchPools()

	return s
}

func (s *VelarLiquidityScraper) fetchPools() {
	tickers, err := s.velarClient.GetAllTickers()
	if err != nil {
		s.logger.WithError(err).Error("failed to fetch velar tickers")
		return
	}

	for _, t := range tickers {
		time.Sleep(time.Duration(s.sleepTimeMilliseconds) * time.Millisecond)

		balances, err := s.fetchPoolBalances(t.PoolID)
		if err != nil {
			s.logger.WithError(err).Error("failed to fetch velar pool balances")
			continue
		}

		tokens := [...]string{t.BaseCurrency, t.TargetCurrency}
		dbAssets := make([]dia.Asset, 0, len(tokens))

		for _, address := range tokens {
			assset, err := s.relDB.GetAsset(address, s.blockchain)
			if err != nil {
				s.logger.WithError(err).Errorf("failed to GetAsset with key: %s", address)
				continue
			}
			dbAssets = append(dbAssets, assset)
		}

		if len(dbAssets) != len(tokens) {
			s.logger.Error("found less than 2 assets for the pool pair")
			continue
		}

		assetVolumes := make([]dia.AssetVolume, len(balances))

		for i, b := range balances {
			asset := dbAssets[i]
			volume, _ := utils.StringToFloat64(b.String(), int64(asset.Decimals))

			assetVolumes[i] = dia.AssetVolume{
				Index:  uint8(i),
				Asset:  asset,
				Volume: volume,
			}
		}

		pool := dia.Pool{
			Exchange:     dia.Exchange{Name: s.exchangeName},
			Blockchain:   dia.BlockChain{Name: s.blockchain},
			Address:      t.ID,
			Time:         time.Now(),
			Assetvolumes: assetVolumes,
		}

		if pool.SufficientNativeBalance(GLOBAL_NATIVE_LIQUIDITY_THRESHOLD) {
			s.datastore.GetPoolLiquiditiesUSD(&pool, priceCache)
		}

		s.logger.WithField("pool", pool).Info("sending pool to poolChannel")
		s.poolChannel <- pool
	}

	s.doneChannel <- true
}

func (s *VelarLiquidityScraper) fetchPoolBalances(poolID string) ([]*big.Int, error) {
	var data []byte
	var err error

	if strings.HasPrefix(poolID, "21000") {
		// Handle univ2 liquidity pools
		args := stackshelper.ContractCallArgs{Sender: velarhelper.DeployerAddressV2}
		poolContract := "univ2-pool-v1_0_0-" + poolID[4:]

		var resp []byte
		resp, err = s.api.CallContractFunction(velarhelper.DeployerAddressV2, poolContract, "get-pool", args)
		if err != nil {
			return nil, err
		}

		if entry, ok := stackshelper.DeserializeCVResponse(resp); ok {
			data = entry
		} else {
			err = errors.New("failed to call get-pool: runtime error")
		}
	} else {
		value := new(big.Int)
		value.SetString(poolID, 10)
		key := hex.EncodeToString(stackshelper.SerializeCVUint(value))

		data, err = s.api.GetDataMapEntry(velarhelper.VelarCoreAddress, "pools", key)
	}

	if err != nil {
		s.logger.WithError(err).Error("failed to fetch velar pool information")
		return nil, err
	}

	pool, err := stackshelper.DeserializeCVTuple(data)
	if err != nil {
		s.logger.WithError(err).Error("failed to deserialize cv tuple")
		return nil, err
	}

	balance0, err := stackshelper.DeserializeCVUint(pool["reserve0"])
	if err != nil {
		return nil, err
	}

	balance1, err := stackshelper.DeserializeCVUint(pool["reserve1"])
	if err != nil {
		return nil, err
	}

	return []*big.Int{balance0, balance1}, nil
}

func (s *VelarLiquidityScraper) Pool() chan dia.Pool {
	return s.poolChannel
}

func (s *VelarLiquidityScraper) Done() chan bool {
	return s.doneChannel
}
