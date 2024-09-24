package liquidityscrapers

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/bitflowhelper"
	"github.com/diadata-org/diadata/pkg/dia/helpers/stackshelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
)

type BitflowLiquidityScraper struct {
	logger             *logrus.Entry
	api                *stackshelper.StacksClient
	poolChannel        chan dia.Pool
	doneChannel        chan bool
	blockchain         string
	exchangeName       string
	relDB              *models.RelDB
	datastore          *models.DB
	handlerType        string
	targetSwapContract string
}

// NewBitflowLiquidityScraper returns a new BitflowLiquidityScraper initialized with default values.
// The instance is asynchronously scraping as soon as it is created.
// ENV values:
//
//	 	BITFLOW_SLEEP_TIMEOUT - (optional, millisecond), make timeout between API calls, default "stackshelper.DefaultSleepBetweenCalls" value
//		BITFLOW_TARGET_SWAP_CONTRACT - (optional, string), useful for debug, default = ""
//		BITFLOW_HIRO_API_KEY - (optional, string), Hiro Stacks API key, improves scraping performance, default = ""
//		BITFLOW_DEBUG - (optional, bool), make stdout output with bitflow client http call, default = false
func NewBitflowLiquidityScraper(exchange dia.Exchange, relDB *models.RelDB, datastore *models.DB) *BitflowLiquidityScraper {
	envPrefix := strings.ToUpper(exchange.Name)

	sleepBetweenCalls := utils.GetenvInt(envPrefix+"_SLEEP_TIMEOUT", stackshelper.DefaultSleepBetweenCalls)
	targetSwapContract := utils.Getenv(envPrefix+"_TARGET_SWAP_CONTRACT", "")
	hiroAPIKey := utils.Getenv(envPrefix+"_HIRO_API_KEY", "")
	isDebug := utils.GetenvBool(envPrefix+"_DEBUG", false)

	stacksClient := stackshelper.NewStacksClient(
		log.WithContext(context.Background()).WithField("context", "StacksClient"),
		utils.GetTimeDurationFromIntAsMilliseconds(sleepBetweenCalls),
		hiroAPIKey,
		isDebug,
	)

	s := &BitflowLiquidityScraper{
		poolChannel:        make(chan dia.Pool),
		doneChannel:        make(chan bool),
		exchangeName:       exchange.Name,
		blockchain:         exchange.BlockChain.Name,
		api:                stacksClient,
		relDB:              relDB,
		datastore:          datastore,
		handlerType:        "liquidity",
		targetSwapContract: targetSwapContract,
	}

	s.logger = logrus.
		New().
		WithContext(context.Background()).
		WithField("handlerType", s.handlerType).
		WithField("context", "BitflowLiquidityScraper")

	go s.fetchPools()

	return s
}

func (s *BitflowLiquidityScraper) fetchPools() {
	poolTxs := make([]stackshelper.Transaction, 0)

	swapContracts := bitflowhelper.StableSwapContracts[:]
	if s.targetSwapContract != "" {
		swapContracts = []string{s.targetSwapContract}
	}

	for _, contractName := range swapContracts {
		s.logger.Infof("Fetching pools of %s", contractName)
		contractId := fmt.Sprintf("%s.%s", bitflowhelper.DeployerAddress, contractName)

		total := stackshelper.MaxPageLimit

		for offset := 0; offset < total; offset += stackshelper.MaxPageLimit {
			resp, err := s.api.GetAddressTransactions(contractId, stackshelper.MaxPageLimit, offset)
			if err != nil {
				s.logger.WithError(err).Error("failed to GetAddressTransactions")
				continue
			}

			total = resp.Total
			filtered := s.filterPoolTransactions(resp.Results)
			poolTxs = append(poolTxs, filtered...)
		}
	}

	for _, tx := range poolTxs {
		args := make(map[string]stackshelper.FunctionArg, len(tx.ContractCall.FunctionArgs))
		for _, item := range tx.ContractCall.FunctionArgs {
			args[item.Name] = item
		}

		tokens := [...]string{args["x-token"].Repr[1:], args["y-token"].Repr[1:]}
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

		contractId := tx.ContractCall.ContractID
		balances, err := s.fetchPoolBalances(
			contractId,
			args["x-token"].Hex,
			args["y-token"].Hex,
			args["lp-token"].Hex,
		)

		if err != nil {
			s.logger.WithError(err).Error("failed to fetch bitflow pool balances")
			continue
		}

		assetVolumes := make([]dia.AssetVolume, len(balances))

		for i, balance := range balances {
			assetVolumes[i] = dia.AssetVolume{
				Index:  uint8(i),
				Asset:  dbAssets[i],
				Volume: balance,
			}
		}

		pool := dia.Pool{
			Exchange:     dia.Exchange{Name: s.exchangeName},
			Blockchain:   dia.BlockChain{Name: s.blockchain},
			Address:      args["lp-token"].Repr[1:],
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

func (s *BitflowLiquidityScraper) fetchPoolBalances(stableSwapContract, xToken, yToken, lpToken string) ([]float64, error) {
	xTokenBytes, _ := hex.DecodeString(xToken[2:])
	yTokenBytes, _ := hex.DecodeString(yToken[2:])
	lpTokenBytes, _ := hex.DecodeString(lpToken[2:])

	pairKey := stackshelper.CVTuple{
		"lp-token": lpTokenBytes,
		"x-token":  xTokenBytes,
		"y-token":  yTokenBytes,
	}
	encodedKey := "0x" + hex.EncodeToString(stackshelper.SerializeCVTuple(pairKey))

	entry, err := s.api.GetDataMapEntry(stableSwapContract, "PairsDataMap", encodedKey)
	if err != nil {
		s.logger.WithError(err).Error("failed to GetDataMapEntry")
		return nil, err
	}

	tuple, err := stackshelper.DeserializeCVTuple(entry)
	if err != nil {
		s.logger.WithError(err).Error("failed to deserialize cv tuple")
		return nil, err
	}

	balanceX, _ := stackshelper.DeserializeCVUint(tuple["balance-x"])
	decimalsX, _ := stackshelper.DeserializeCVUint(tuple["x-decimals"])

	balanceY, _ := stackshelper.DeserializeCVUint(tuple["balance-y"])
	decimalsY, _ := stackshelper.DeserializeCVUint(tuple["y-decimals"])

	balances := make([]float64, 2)
	balances[0], _ = utils.StringToFloat64(balanceX.String(), decimalsX.Int64())
	balances[1], _ = utils.StringToFloat64(balanceY.String(), decimalsY.Int64())

	return balances, nil
}

func (s *BitflowLiquidityScraper) filterPoolTransactions(txs []stackshelper.AddressTransaction) []stackshelper.Transaction {
	poolTxs := make([]stackshelper.Transaction, 0)

	for _, item := range txs {
		isCreatePairCall := item.Tx.TxType == "contract_call" &&
			item.Tx.ContractCall.FunctionName == "create-pair"

		if isCreatePairCall && item.Tx.TxStatus == "success" {
			poolTxs = append(poolTxs, item.Tx)
		}
	}

	return poolTxs
}

func (s *BitflowLiquidityScraper) Pool() chan dia.Pool {
	return s.poolChannel
}

func (s *BitflowLiquidityScraper) Done() chan bool {
	return s.doneChannel
}
