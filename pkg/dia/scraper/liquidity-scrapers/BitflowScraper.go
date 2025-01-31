package liquidityscrapers

import (
	"context"
	"encoding/hex"
	"errors"
	"math/big"
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
	swapContracts := bitflowhelper.SwapContracts[:]

	if s.targetSwapContract != "" {
		address := strings.Split(s.targetSwapContract, ".")

		contractType := 0
		if strings.HasPrefix(address[1], "xyk") {
			contractType = 1
		}

		swapContracts = []bitflowhelper.SwapContract{
			{
				ContractType:     contractType,
				DeployerAddress:  address[0],
				ContractRegistry: address[1],
			},
		}
	}

	for _, contract := range swapContracts {
		s.logger.Infof("Fetching pools of %s", contract.ContractRegistry)
		contractID := contract.DeployerAddress + "." + contract.ContractRegistry

		switch contract.ContractType {
		case 0:
			total := stackshelper.MaxPageLimit

			for offset := 0; offset < total; offset += stackshelper.MaxPageLimit {
				resp, err := s.api.GetAddressTransactions(contractID, stackshelper.MaxPageLimit, offset)
				if err != nil {
					s.logger.WithError(err).Error("failed to GetAddressTransactions")
					continue
				}

				total = resp.Total
				filtered := s.fetchPoolTransactions(resp.Results, contract.ContractType)
				for _, tx := range filtered {
					pool, err := s.parseTx(tx)
					if err != nil {
						continue
					}
					// s.logger.WithField("pool", pool).Info("sending pool to poolChannel")
					s.poolChannel <- pool
				}
			}
		case 1:
			lastID, err := s.api.GetDataVar(contract.DeployerAddress, contract.ContractRegistry, "last-pool-id")
			if err != nil {
				s.logger.WithError(err).Error("failed to get last pool ID")
				continue
			}

			total, err := stackshelper.DeserializeCVUint(lastID)
			if err != nil {
				s.logger.WithError(err).Error("failed to deserialize CV uint")
				continue
			}

			xykContract := contract.DeployerAddress + "." + contract.ContractRegistry
			one := big.NewInt(1)

			for id := new(big.Int).Set(one); id.Cmp(total) <= 0; id.Add(id, one) {
				poolContract, err := s.getXykPoolContractAddress(xykContract, id)
				if err != nil {
					s.logger.Error("failed to get xyk pool contract address")
					continue
				}

				pool, err := s.fetchXykPool(poolContract)
				if err != nil {
					continue
				}

				s.logger.WithField("pool", pool).Info("sending pool to poolChannel")
				s.poolChannel <- pool
			}

		}
	}

	s.doneChannel <- true
}

func (s *BitflowLiquidityScraper) parseTx(tx stackshelper.Transaction) (dia.Pool, error) {
	args := make(map[string]stackshelper.FunctionArg, len(tx.ContractCall.FunctionArgs))
	for _, item := range tx.ContractCall.FunctionArgs {
		args[item.Name] = item
	}

	tokens := [...]string{"", args["y-token"].Repr[1:]}
	if xToken, ok := args["x-token"]; ok {
		tokens[0] = xToken.Repr[1:]
	}

	dbAssets, err := s.fetchAssets(tokens[:])
	if err != nil {
		return dia.Pool{}, err
	}

	balances, err := s.fetchStableswapPoolBalances(
		tx.ContractCall.ContractID,
		args["x-token"].Hex,
		args["y-token"].Hex,
		args["lp-token"].Hex,
	)

	if err != nil {
		return dia.Pool{}, errors.New("failed to fetch bitflow pool balances")
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
		Time:         time.Now(),
		Assetvolumes: assetVolumes,
		Address:      args["lp-token"].Repr[1:],
	}

	if pool.SufficientNativeBalance(GLOBAL_NATIVE_LIQUIDITY_THRESHOLD) {
		s.datastore.GetPoolLiquiditiesUSD(&pool, priceCache)
	}
	return pool, nil
}

func (s *BitflowLiquidityScraper) fetchStableswapPoolBalances(contract, xToken, yToken, lpToken string) ([]float64, error) {
	yTokenBytes, _ := hex.DecodeString(yToken[2:])
	lpTokenBytes, _ := hex.DecodeString(lpToken[2:])
	pairKey := stackshelper.CVTuple{"lp-token": lpTokenBytes, "y-token": yTokenBytes}

	if xToken != "" {
		xTokenBytes, _ := hex.DecodeString(xToken[2:])
		pairKey["x-token"] = xTokenBytes
	}

	encodedKey := "0x" + hex.EncodeToString(stackshelper.SerializeCVTuple(pairKey))
	entry, err := s.api.GetDataMapEntry(contract, "PairsDataMap", encodedKey)
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

func (s *BitflowLiquidityScraper) getXykPoolContractAddress(contractID string, poolID *big.Int) (string, error) {
	encodedPoolID := hex.EncodeToString(stackshelper.SerializeCVUint(poolID))

	result, err := s.api.GetDataMapEntry(contractID, "pools", encodedPoolID)
	if err != nil {
		log.WithError(err).Error("failed to get pool by ID")
		return "", err
	}

	tuple, err := stackshelper.DeserializeCVTuple(result)
	if err != nil {
		log.WithError(err).Error("failed to deserialize cv tuple")
		return "", err
	}

	return stackshelper.DeserializeCVPrincipal(tuple["pool-contract"])
}

func (s *BitflowLiquidityScraper) fetchXykPool(poolContract string) (dia.Pool, error) {
	address := strings.Split(poolContract, ".")
	args := stackshelper.ContractCallArgs{Sender: address[0]}

	result, err := s.api.CallContractFunction(address[0], address[1], "get-pool", args)
	if err != nil {
		return dia.Pool{}, err
	}

	data, ok := stackshelper.DeserializeCVResponse(result)
	if !ok {
		return dia.Pool{}, errors.New("failed to deserialize CV response")
	}
	poolInfo, err := stackshelper.DeserializeCVTuple(data)
	if err != nil {
		return dia.Pool{}, err
	}

	xToken, _ := stackshelper.DeserializeCVPrincipal(poolInfo["x-token"])
	yToken, _ := stackshelper.DeserializeCVPrincipal(poolInfo["y-token"])

	xDecimals, err := s.fetchTokenDecimals(xToken)
	if err != nil {
		return dia.Pool{}, err
	}
	yDecimals, err := s.fetchTokenDecimals(yToken)
	if err != nil {
		return dia.Pool{}, err
	}

	dbAssets, err := s.fetchAssets([]string{xToken, yToken})
	if err != nil {
		return dia.Pool{}, err
	}

	xBalance, _ := stackshelper.DeserializeCVUint(poolInfo["x-balance"])
	yBalance, _ := stackshelper.DeserializeCVUint(poolInfo["y-balance"])

	balances := make([]float64, 2)
	balances[0], _ = utils.StringToFloat64(xBalance.String(), xDecimals)
	balances[1], _ = utils.StringToFloat64(yBalance.String(), yDecimals)

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
		Time:         time.Now(),
		Assetvolumes: assetVolumes,
		Address:      poolContract,
	}

	if pool.SufficientNativeBalance(GLOBAL_NATIVE_LIQUIDITY_THRESHOLD) {
		s.datastore.GetPoolLiquiditiesUSD(&pool, priceCache)
	}
	return pool, nil
}

func (s *BitflowLiquidityScraper) fetchTokenDecimals(tokenContract string) (int64, error) {
	address := strings.Split(tokenContract, ".")
	args := stackshelper.ContractCallArgs{Sender: address[0]}

	result, err := s.api.CallContractFunction(address[0], address[1], "get-decimals", args)
	if err != nil {
		return 0, err
	}

	data, ok := stackshelper.DeserializeCVResponse(result)
	if !ok {
		return 0, errors.New("failed to deserialize CV response")
	}

	decimals, err := stackshelper.DeserializeCVUint(data)
	if err != nil {
		return 0, err
	}
	return decimals.Int64(), nil
}

func (s *BitflowLiquidityScraper) fetchAssets(tokens []string) ([]dia.Asset, error) {
	dbAssets := make([]dia.Asset, 0, len(tokens))

	for _, address := range tokens {
		// Workaround to fetch the native STX token data from DB
		key := address
		if address == "" {
			key = "0x0000000000000000000000000000000000000000"
		}

		assset, err := s.relDB.GetAsset(key, s.blockchain)
		if err != nil {
			s.logger.WithError(err).Errorf("failed to GetAsset with key: %s", key)
			continue
		}
		dbAssets = append(dbAssets, assset)
	}

	if len(dbAssets) != len(tokens) {
		return nil, errors.New("found less than 2 assets for the pool pair")
	}
	return dbAssets, nil
}

func (s *BitflowLiquidityScraper) fetchPoolTransactions(txs []stackshelper.AddressTransaction, poolType int) []stackshelper.Transaction {
	poolTxs := make([]stackshelper.Transaction, 0)

	for _, item := range txs {
		var isCreatePairCall bool
		if poolType == 0 {
			isCreatePairCall = item.Tx.TxType == "contract_call" &&
				item.Tx.ContractCall.FunctionName == "create-pair"
		} else if poolType == 1 {
			isCreatePairCall = item.Tx.TxType == "contract_call" &&
				item.Tx.ContractCall.FunctionName == "create-pool"
		}

		if isCreatePairCall && item.Tx.TxStatus == "success" {
			// This is a temporary workaround introduced due to a bug in hiro stacks API.
			// Results returned from /addresses/{address}/transactions route have empty
			// `name` field in `contract_call.function_args` list.
			// TODO: remove this as soon as the issue is fixed.
			normalizedTx, err := s.api.GetTransactionAt(item.Tx.TxID)
			if err != nil {
				s.logger.WithError(err).Error("failed to GetTransactionAt")
				continue
			}
			poolTxs = append(poolTxs, normalizedTx)
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
