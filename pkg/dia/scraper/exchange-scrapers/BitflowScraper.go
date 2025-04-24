package scrapers

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	bitflow "github.com/diadata-org/diadata/pkg/dia/helpers/bitflowhelper"
	stacks "github.com/diadata-org/diadata/pkg/dia/helpers/stackshelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/zekroTJA/timedmap"
)

type bitflowSwapEvent struct {
	txID        string
	action      string
	amountIn    string
	amountOut   string
	poolAddress string
	blockTime   int
}

type BitflowScraper struct {
	logger             *logrus.Entry
	pairScrapers       map[string]*BitflowPairScraper // pc.ExchangePair -> pairScraperSet
	swapContracts      map[string]nothing
	shutdown           chan nothing
	shutdownDone       chan nothing
	errorLock          sync.RWMutex
	error              error
	closed             bool
	ticker             *time.Ticker
	exchangeName       string
	blockchain         string
	chanTrades         chan *dia.Trade
	api                *stacks.StacksClient
	db                 *models.RelDB
	currentHeight      int
	initialBlockHeight int
}

var (
	singleDirectionPoolsBitflow *[]string
)

// NewBitflowScraper returns a new BitflowScraper initialized with default values.
// The instance is asynchronously scraping as soon as it is created.
// ENV values:
//
//	 	BITFLOW_SLEEP_TIMEOUT - (optional, millisecond), make timeout between API calls, default "stackshelper.DefaultSleepBetweenCalls" value
//		BITFLOW_REFRESH_DELAY - (optional, millisecond) refresh data after each poll, default "stackshelper.DefaultRefreshDelay" value
//		BITFLOW_HIRO_API_KEY - (optional, string), Hiro Stacks API key, improves scraping performance, default = ""
//		BITFLOW_INITIAL_BLOCK_HEIGHT (optional, int), useful for debug, default = 0
//		BITFLOW_DEBUG - (optional, bool), make stdout output with alephium client http call, default = false
func NewBitflowScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *BitflowScraper {
	envPrefix := strings.ToUpper(exchange.Name)

	sleepBetweenCalls := utils.GetTimeDurationFromIntAsMilliseconds(
		utils.GetenvInt(
			envPrefix+"_SLEEP_TIMEOUT",
			stacks.DefaultSleepBetweenCalls,
		),
	)
	refreshDelay := utils.GetTimeDurationFromIntAsMilliseconds(
		utils.GetenvInt(envPrefix+"_REFRESH_DELAY", stacks.DefaultRefreshDelay),
	)
	hiroAPIKey := utils.Getenv(envPrefix+"_HIRO_API_KEY", "")
	initialBlockHeight := utils.GetenvInt(envPrefix+"_INITIAL_BLOCK_HEIGHT", 0)
	isDebug := utils.GetenvBool(envPrefix+"_DEBUG", false)

	stacksClient := stacks.NewStacksClient(
		log.WithContext(context.Background()).WithField("context", "StacksClient"),
		sleepBetweenCalls,
		hiroAPIKey,
		isDebug,
	)

	swapContracts := make(map[string]nothing, len(bitflow.SwapContracts))

	for _, contract := range bitflow.SwapContracts {
		contractId := fmt.Sprintf("%s.%s", contract.DeployerAddress, contract.ContractRegistry)
		swapContracts[contractId] = nothing{}
	}

	s := &BitflowScraper{
		shutdown:           make(chan nothing),
		shutdownDone:       make(chan nothing),
		pairScrapers:       make(map[string]*BitflowPairScraper),
		swapContracts:      swapContracts,
		ticker:             time.NewTicker(refreshDelay),
		chanTrades:         make(chan *dia.Trade),
		api:                stacksClient,
		db:                 relDB,
		exchangeName:       exchange.Name,
		blockchain:         exchange.BlockChain.Name,
		initialBlockHeight: initialBlockHeight,
	}

	s.logger = logrus.
		New().
		WithContext(context.Background()).
		WithField("context", "BitflowDEXScraper")

	if scrape {
		go s.mainLoop()
	}
	return s
}

func (s *BitflowScraper) mainLoop() {

	var err error
	singleDirectionPoolsBitflow, err = getReverseTokensFromConfig("bitflow/singleDirectionPools/" + s.exchangeName + "SingleDirectionPools")
	if err != nil {
		log.Error("error getting fullPools for which pairs should be reversed: ", err)
	}
	log.Info("singleDiration: ", *singleDirectionPoolsBitflow)

	if s.initialBlockHeight <= 0 {
		latestBlock, err := s.api.GetLatestBlock()
		if err != nil {
			s.logger.WithError(err).Error("failed to GetLatestBlock")
			s.cleanup(err)
			return
		}
		s.currentHeight = latestBlock.Height
	} else {
		s.currentHeight = s.initialBlockHeight
	}

	for {
		select {
		case <-s.ticker.C:
			err := s.Update()
			if err != nil {
				s.logger.Error(err)
			}
		case <-s.shutdown:
			s.logger.Println("shutting down")
			s.cleanup(nil)
			return
		}
	}
}

func (s *BitflowScraper) Update() error {
	txs, err := s.api.GetAllBlockTransactions(s.currentHeight)
	if err != nil {
		s.logger.WithError(err).Error("failed to GetBlockTransactions")
		return err
	}

	if len(txs) == 0 {
		return nil
	}
	s.currentHeight += 1

	swapEvents, err := s.fetchSwapEvents(txs)
	if err != nil {
		s.logger.WithError(err).Error("failed to fetchSwapEvents")
		return err
	}

	if len(swapEvents) == 0 {
		return nil
	}

	pools, err := s.getPools()
	if err != nil {
		s.logger.WithError(err).Error("failed to GetAllPoolsExchange")
		return err
	}

	for _, pool := range pools {
		tmFalseDuplicateTrades := timedmap.New(duplicateTradesScanFrequency)
		tmDuplicateTrades := timedmap.New(duplicateTradesScanFrequency)
		if len(pool.Assetvolumes) != 2 {
			s.logger.WithField("poolAddress", pool.Address).Error("pool is missing required asset volumes")
			continue
		}

		for _, e := range swapEvents {
			if e.poolAddress != pool.Address {
				continue
			}

			diaTrade := s.handleTrade(&pool, &e)
			log.Infof("got trade at height %v: %v -- %s -- %v --%v -- %s", s.currentHeight-1, diaTrade.Time, diaTrade.Pair, diaTrade.Price, diaTrade.Volume, diaTrade.ForeignTradeID)
			discardTrade := diaTrade.IdentifyDuplicateFull(tmFalseDuplicateTrades, duplicateTradesMemory)
			if discardTrade {
				log.Warn("Identical trade already scraped: ", diaTrade)
				continue
			} else {
				diaTrade.IdentifyDuplicateTagset(tmDuplicateTrades, duplicateTradesMemory)
				s.chanTrades <- diaTrade
				if !utils.Contains(singleDirectionPoolsBitflow, pool.Address) {
					tSwapped, err := dia.SwapTrade(*diaTrade)
					if err == nil {
						if tSwapped.Price > 0 {
							log.Infof("got trade at height %v: %v -- %s -- %v --%v -- %s", s.currentHeight-1, tSwapped.Time, tSwapped.Pair, tSwapped.Price, tSwapped.Volume, tSwapped.ForeignTradeID)
							s.chanTrades <- &tSwapped
						}
					}
				}
			}
		}
	}

	return nil
}

func (s *BitflowScraper) getPools() ([]dia.Pool, error) {
	return s.db.GetAllPoolsExchange(s.exchangeName, 0)
}

func (s *BitflowScraper) fetchSwapEvents(transactions []stacks.Transaction) ([]bitflowSwapEvent, error) {
	swapEvents := make([]bitflowSwapEvent, 0)

	for _, tx := range transactions {
		if tx.TxStatus != "success" || tx.TxType != "contract_call" {
			continue
		}

		// This is a temporary workaround introduced due to a bug in hiro stacks API.
		// Results returned from /blocks/{block_height}/transactions route have empty
		// `name` field in `contract_call.function_args` list.
		// TODO: remove this as soon as the issue is fixed.
		normalizedTx, err := s.api.GetTransactionAt(tx.TxID)
		if err != nil {
			return nil, err
		}

		_, contractFound := s.swapContracts[tx.ContractCall.ContractID]
		isStableSwapTransaction := contractFound &&
			strings.Contains(tx.ContractCall.ContractID, "stableswap") &&
			strings.HasPrefix(tx.ContractCall.FunctionName, "swap")

		if isStableSwapTransaction {
			event := bitflowSwapEvent{
				txID:      tx.TxID,
				action:    tx.ContractCall.FunctionName,
				amountOut: normalizedTx.TxResult.Repr[5 : len(normalizedTx.TxResult.Repr)-1],
				blockTime: tx.BlockTime,
			}

			for _, arg := range normalizedTx.ContractCall.FunctionArgs {
				value := arg.Repr[1:]
				switch arg.Name {
				case "x-amount":
					event.amountIn = value
				case "y-amount":
					event.amountIn = value
				case "lp-token":
					event.poolAddress = value
				}
			}

			swapEvents = append(swapEvents, event)
		} else {
			for _, e := range normalizedTx.Events {
				log := &e.ContractLog

				isBitflowSwap := e.Type == "smart_contract_log" &&
					log.Topic == "print" &&
					(strings.HasPrefix(log.ContractID, bitflow.StableSwapDeployer) || strings.HasPrefix(log.ContractID, bitflow.XykDeployer)) &&
					(strings.Contains(log.Value.Repr, "swap-x-for-y") || strings.Contains(log.Value.Repr, "swap-y-for-x"))

				if !isBitflowSwap {
					continue
				}

				bytes, err := hex.DecodeString(e.ContractLog.Value.Hex[2:])
				if err != nil {
					s.logger.WithError(err).Error("failed to decode contract log")
					return nil, err
				}

				event, err := s.decodeXykSwapEvent(tx.TxID, tx.BlockTime, bytes)
				if err != nil {
					return nil, err
				}
				swapEvents = append(swapEvents, event)
			}
		}
	}

	return swapEvents, nil
}

func (s *BitflowScraper) decodeXykSwapEvent(txID string, blockTime int, src []byte) (bitflowSwapEvent, error) {
	empty := bitflowSwapEvent{}

	tuple, err := stacks.DeserializeCVTuple(src)
	if err != nil {
		s.logger.WithError(err).Error("failed to deserialize cv tuple")
		return empty, err
	}

	action, err := stacks.DeserializeCVString(tuple["action"])
	if err != nil {
		s.logger.WithError(err).Error("failed to deserialize event action")
		return empty, err
	}

	data, err := stacks.DeserializeCVTuple(tuple["data"])
	if err != nil {
		s.logger.WithError(err).Error("failed to deserialize cv tuple")
		return empty, err
	}

	var keyIn, keyOut string
	if action == "swap-x-for-y" {
		keyIn = "x-amount"
		keyOut = "dy"
	} else {
		keyIn = "y-amount"
		keyOut = "dx"
	}

	amountIn, err := stacks.DeserializeCVUint(data[keyIn])
	if err != nil {
		s.logger.WithError(err).Error("failed to deserialize input amount")
		return empty, err
	}

	amountOut, err := stacks.DeserializeCVUint(data[keyOut])
	if err != nil {
		s.logger.WithError(err).Error("failed to deserialize output amount")
		return empty, err
	}

	poolContract, err := stacks.DeserializeCVPrincipal(data["pool-contract"])
	if err != nil {
		s.logger.WithError(err).Error("failed to deserialize pool contract address")
		return empty, err
	}

	event := bitflowSwapEvent{
		txID:        txID,
		action:      action,
		amountIn:    amountIn.String(),
		amountOut:   amountOut.String(),
		poolAddress: poolContract,
		blockTime:   blockTime,
	}

	return event, nil
}

func (s *BitflowScraper) handleTrade(pool *dia.Pool, event *bitflowSwapEvent) *dia.Trade {
	var volume, price float64

	decimals0 := int64(pool.Assetvolumes[0].Asset.Decimals)
	decimals1 := int64(pool.Assetvolumes[1].Asset.Decimals)

	if event.action == "swap-x-for-y" {
		amount0In, _ := utils.StringToFloat64(event.amountIn, decimals0)
		amount1Out, _ := utils.StringToFloat64(event.amountOut, decimals1)
		volume = amount0In
		price = amount1Out / amount0In
	} else {
		amount1In, _ := utils.StringToFloat64(event.amountIn, decimals1)
		amount0Out, _ := utils.StringToFloat64(event.amountOut, decimals0)
		volume = -amount0Out
		price = amount1In / amount0Out
	}

	symbolPair := fmt.Sprintf("%s-%s", pool.Assetvolumes[0].Asset.Symbol, pool.Assetvolumes[1].Asset.Symbol)

	return &dia.Trade{
		Time:           time.Unix(int64(event.blockTime), 0),
		Symbol:         pool.Assetvolumes[0].Asset.Symbol,
		Pair:           symbolPair,
		ForeignTradeID: event.txID,
		Source:         s.exchangeName,
		Price:          price,
		Volume:         volume,
		VerifiedPair:   true,
		BaseToken:      pool.Assetvolumes[1].Asset,
		QuoteToken:     pool.Assetvolumes[0].Asset,
		PoolAddress:    pool.Address,
	}
}

func (s *BitflowScraper) FetchAvailablePairs() ([]dia.ExchangePair, error) {
	return []dia.ExchangePair{}, nil
}

func (s *BitflowScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

func (s *BitflowScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

func (s *BitflowScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("BitflowScraper: Call ScrapePair on closed scraper")
	}
	ps := &BitflowPairScraper{
		parent:     s,
		pair:       pair,
		lastRecord: 0,
	}

	s.pairScrapers[pair.Symbol] = ps

	return ps, nil
}

func (s *BitflowScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()

	s.ticker.Stop()

	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.shutdownDone)
}

// Close gracefully shuts down the BitflowScraper.
func (s *BitflowScraper) Close() error {
	if s.closed {
		return errors.New("BitflowScraper: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Channel returns the channel used to receive trades/pricing information.
func (s *BitflowScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

type BitflowPairScraper struct {
	parent     *BitflowScraper
	pair       dia.ExchangePair
	closed     bool
	lastRecord int64
}

func (ps *BitflowPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}

func (ps *BitflowPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *BitflowPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}
