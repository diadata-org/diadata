package tradesBlockService

import (
	"errors"
	"math"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/cnf/structhash"
	"github.com/diadata-org/diadata/pkg/dia"
	scrapers "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
)

type nothing struct{}

func init() {
	log = logrus.New()
	var err error
	batchTimeSeconds, err = strconv.Atoi(utils.Getenv("BATCH_TIME_SECONDS", "30"))
	if err != nil {
		log.Error("parse BATCH_TIME_SECONDS: ", err)
	}
	volumeThreshold, err = strconv.ParseFloat(utils.Getenv("VOLUME_THRESHOLD", "100000"), 64)
	if err != nil {
		log.Error("parse env var VOLUME_THRESHOLD: ", err)
	}
	blueChipThreshold, err = strconv.ParseFloat(utils.Getenv("BLUECHIP_THRESHOLD", "50000000"), 64)
	if err != nil {
		log.Error("parse env var BLUECHIP_THRESHOLD: ", err)
	}
	smallX, err = strconv.ParseFloat(utils.Getenv("SMALL_X", "10"), 64)
	if err != nil {
		log.Error("parse env var SMALL_X: ", err)
	}
	normalX, err = strconv.ParseFloat(utils.Getenv("NORMAL_X", "10"), 64)
	if err != nil {
		log.Error("parse env var NORMAL_X: ", err)
	}
	tradeVolumeThresholdExponent, err := strconv.ParseFloat(utils.Getenv("TRADE_VOLUME_THRESHOLD_EXPONENT", ""), 64)
	if err != nil {
		log.Error("Parse TRADE_VOLUME_THRESHOLD_EXPONENT: ", err)
	}
	tradeVolumeThreshold = math.Pow(10, -tradeVolumeThresholdExponent)
	tradeVolumeThresholdUSDExponent, err := strconv.ParseFloat(utils.Getenv("TRADE_VOLUME_THRESHOLD_USD_EXPONENT", ""), 64)
	if err != nil {
		log.Error("Parse TRADE_VOLUME_THRESHOLD_USD_EXPONENT: ", err)
	}
	tradeVolumeThresholdUSD = math.Pow(10, -tradeVolumeThresholdUSDExponent)
}

var (
	stablecoins = map[string]interface{}{
		"USDT": "",
		"TUSD": "",
		// "DAI":  "",
		// "USDC": "",
		"PAX":  "",
		"BUSD": "",
	}

	// These should be loaded from postgres once we have a list.
	USDT             = dia.Asset{Address: "0xdAC17F958D2ee523a2206206994597C13D831ec7", Blockchain: dia.ETHEREUM}
	USDT_BNB_CHAIN   = dia.Asset{Address: "0x55d398326f99059fF775485246999027B3197955", Blockchain: dia.BINANCESMARTCHAIN}
	USDC             = dia.Asset{Address: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", Blockchain: dia.ETHEREUM}
	BUSD             = dia.Asset{Address: "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56", Blockchain: dia.BINANCESMARTCHAIN}
	DAI              = dia.Asset{Address: "0x6B175474E89094C44Da98b954EedeAC495271d0F", Blockchain: dia.ETHEREUM}
	TUSD             = dia.Asset{Address: "0x0000000000085d4780B73119b644AE5ecd22b376", Blockchain: dia.ETHEREUM}
	stablecoinAssets = map[string]interface{}{
		USDT.Identifier(): "",
		USDC.Identifier(): "",
		DAI.Identifier():  "",
		TUSD.Identifier(): "",
	}

	tol                     = float64(0.04)
	log                     *logrus.Logger
	batchTimeSeconds        int
	tradeVolumeThreshold    float64
	tradeVolumeThresholdUSD float64
	volumeUpdateSeconds     = 60 * 10
	volumeThreshold         float64
	blueChipThreshold       float64
	smallX                  float64
	normalX                 float64
	checkTradesDuplicate    = make(map[string]struct{})
)

type TradesBlockService struct {
	shutdown         chan nothing
	shutdownDone     chan nothing
	chanTrades       chan *dia.Trade
	chanTradesBlock  chan *dia.TradesBlock
	errorLock        sync.RWMutex
	error            error
	closed           bool
	started          bool
	BlockDuration    int64
	currentBlock     *dia.TradesBlock
	priceCache       map[string]float64
	volumeCache      map[string]float64
	datastore        models.Datastore
	relDB            models.RelDatastore
	historical       bool
	writeMeasurement string
	batchTicker      *time.Ticker
	volumeTicker     *time.Ticker
}

func NewTradesBlockService(datastore models.Datastore, relDB models.RelDatastore, blockDuration int64, historical bool) *TradesBlockService {
	s := &TradesBlockService{
		shutdown:        make(chan nothing),
		shutdownDone:    make(chan nothing),
		chanTrades:      make(chan *dia.Trade),
		chanTradesBlock: make(chan *dia.TradesBlock),
		error:           nil,
		started:         false,
		currentBlock:    nil,
		BlockDuration:   blockDuration,
		priceCache:      make(map[string]float64),
		volumeCache:     make(map[string]float64),
		datastore:       datastore,
		relDB:           relDB,
		historical:      historical,
		batchTicker:     time.NewTicker(time.Duration(batchTimeSeconds) * time.Second),
		volumeTicker:    time.NewTicker(time.Duration(volumeUpdateSeconds) * time.Second),
	}
	if historical {
		s.writeMeasurement = utils.Getenv("INFLUX_MEASUREMENT_WRITE", "tradesTmp")
	}
	log.Info("write measurement: ", s.writeMeasurement)
	log.Info("historical: ", s.historical)
	log.Info("batch ticker time: ", batchTimeSeconds)
	log.Info("volume threshold: ", volumeThreshold)
	log.Info("bluechip threshold: ", blueChipThreshold)
	log.Info("smallX: ", smallX)
	log.Info("normalX: ", normalX)
	log.Info("tradeVolumeThreshold: ", tradeVolumeThreshold)
	log.Info("tradeVolumeThresholdUSD: ", tradeVolumeThresholdUSD)

	s.volumeCache = s.loadVolumes()
	log.Info("...done loading volumes.")

	go s.mainLoop()
	go s.loadVolumesLoop()
	return s
}

// runs in a goroutine until s is closed
func (s *TradesBlockService) mainLoop() {
	var (
		acceptCountDEX int
		// acceptCountSwapDEX int
		totalCount int
		logTicker  = *time.NewTicker(120 * time.Second)
	)
	for {
		select {
		case <-s.shutdown:
			log.Println("TradesBlockService shutting down")
			s.cleanup(nil)
			return
		case t := <-s.chanTrades:

			// Only take into account original order for CEX trade.
			if scrapers.Exchanges[(*t).Source].Centralized {
				s.process(*t)
			} else {

				// tSwapped, err := dia.SwapTrade(*t)
				// if err != nil {
				// 	log.Error("swap trade: ", err)
				// }

				// Collect booleans for stats.
				tradeOk := s.checkTrade(*t)
				// swapppedTradeOk := s.checkTrade(tSwapped)
				if tradeOk {
					acceptCountDEX++
				}
				// if swapppedTradeOk {
				// 	acceptCountSwapDEX++
				// }
				// if tradeOk || swapppedTradeOk {
				// 	acceptCount++
				// }
				totalCount++

				// Process (possibly) both trades.
				if tradeOk {
					s.process(*t)
				}
				// s.process(tSwapped)
			}

		case <-s.batchTicker.C:
			err := s.datastore.Flush()
			if err != nil {
				log.Error("flush influx batch: ", err)
			}
		case <-logTicker.C:
			log.Info("accepted trades DEX: ", acceptCountDEX)
			// log.Info("accepted swapped trades DEX: ", acceptCountSwapDEX)
			log.Info("discarded trades: ", totalCount-acceptCountDEX)
			acceptCountDEX = 0
			// acceptCountSwapDEX = 0
			totalCount = 0
		}
	}
}

// checkTrade determines whether a (DEX-)trade should be taken into account for price determination.
func (s *TradesBlockService) checkTrade(t dia.Trade) bool {

	// Discard (very) low volume trade.
	if math.Abs(t.Volume) < tradeVolumeThreshold {
		log.Info("low volume trade: ", t)
		return false
	}

	// Replace basetoken with bridged asset for pricing if necessary.
	// The basetoken in the stored trade will remain unchanged.
	basetoken := buildBridge(t)

	// Allow trade where basetoken is stablecoin.
	if _, ok := stablecoinAssets[basetoken.Identifier()]; ok {
		return true
	}

	// Only take into account stablecoin trade if basetoken is stable coin as well.
	if _, ok := stablecoinAssets[t.QuoteToken.Identifier()]; ok {
		if _, ok := stablecoinAssets[basetoken.Identifier()]; !ok {
			return false
		}
	}

	if baseVolume, ok := s.volumeCache[basetoken.Identifier()]; ok {
		if baseVolume > blueChipThreshold {
			return true
		}
		if quoteVolume, ok := s.volumeCache[t.QuoteToken.Identifier()]; ok {
			if baseVolume < volumeThreshold {
				// For small volume basetoken, quotetoken must be a small volume asset too.
				return quoteVolume < smallX*baseVolume
			}
			// Discard trade if base volume is too small compared to quote volume.
			return quoteVolume < normalX*baseVolume
		}
		// Base asset has enough volume or quotetoken has no volume yet.
		return true
	}
	return false
}

func (s *TradesBlockService) process(t dia.Trade) {

	var (
		verifiedTrade bool
		tradeOk       bool
	)

	if scrapers.Exchanges[t.Source].Centralized {
		tradeOk = true
	} else {
		tradeOk = s.checkTrade(t)
	}

	// Price estimation can only be done for verified pairs.
	// Trades with unverified pairs are still saved, but not sent to the filtersBlockService.
	if t.VerifiedPair && tradeOk {
		if t.BaseToken.Address == "840" && t.BaseToken.Blockchain == dia.FIAT {
			// All prices are measured in US-Dollar, so just price for base token == USD
			t.EstimatedUSDPrice = t.Price
			verifiedTrade = true
		} else {
			var (
				quotation *models.AssetQuotation
				price     float64
				ok        bool
				err       error
			)

			// Bridge basetoken if necessary.
			basetoken := buildBridge(t)

			// Get latest price from cache.
			if _, ok = s.priceCache[basetoken.Identifier()]; ok {
				price = s.priceCache[basetoken.Identifier()]
			} else {
				quotation, err = s.datastore.GetAssetQuotationCache(basetoken)
				price = quotation.Price
				s.priceCache[basetoken.Identifier()] = price
				// log.Infof("quotation for %s from redis cache: %v", basetoken.Symbol, price)
			}

			if err != nil {
				log.Errorf("Can't find quotation for base token in trade %s: %v.\n Basetoken address -- blockchain:  %s --- %s",
					t.Pair,
					err,
					t.BaseToken.Address,
					t.BaseToken.Blockchain,
				)
			} else {
				if price > 0.0 {
					t.EstimatedUSDPrice = t.Price * price
					if t.VolumeUSD() > tradeVolumeThresholdUSD {
						verifiedTrade = true
					} else {
						log.Warn("low $ volume on trade: ", t)
					}
				}
			}
		}
	}

	// // If estimated price for stablecoin diverges too much ignore trade
	if _, ok := stablecoins[t.Symbol]; ok {
		if math.Abs(t.EstimatedUSDPrice-1) > tol && t.EstimatedUSDPrice > 0 {
			log.Errorf("%s on %s. price for %s diverges by %v", t.Pair, t.Source, t.Symbol, math.Abs(t.EstimatedUSDPrice-1))
			verifiedTrade = false
		}
	}
	// Comment Philipp: We could make another check here. Store CG and/or CMC quotation in redis cache
	// and compare with estimatedUSDPrice. If deviation is too large ignore trade.
	var err error
	if !s.historical {
		err = s.datastore.SaveTradeInflux(&t)
		if err != nil {
			log.Error(err)
		}
	} else {
		err = s.datastore.SaveTradeInfluxToTable(&t, s.writeMeasurement)
		if err != nil {
			log.Error(err)
		}
	}

	if s.currentBlock != nil && s.currentBlock.TradesBlockData.BeginTime.After(t.Time) {
		log.Debugf("ignore trade should be in previous block %v", t)
		verifiedTrade = false
	}

	// Only verified trades of verified pairs with nonzero price are added to the tradesBlock
	if verifiedTrade && t.EstimatedUSDPrice > 0 {
		if s.currentBlock == nil || s.currentBlock.TradesBlockData.EndTime.Before(t.Time) {
			if s.currentBlock != nil {
				s.finaliseCurrentBlock()
				s.priceCache = make(map[string]float64)
			}

			b := &dia.TradesBlock{
				TradesBlockData: dia.TradesBlockData{
					Trades:    []dia.Trade{},
					EndTime:   time.Unix((t.Time.Unix()/s.BlockDuration)*s.BlockDuration+s.BlockDuration, 0),
					BeginTime: time.Unix((t.Time.Unix()/s.BlockDuration)*s.BlockDuration, 0),
				},
			}
			if s.currentBlock != nil {
				log.Info("created new block beginTime:", b.TradesBlockData.BeginTime, "previous block nb trades:", len(s.currentBlock.TradesBlockData.Trades))
			}
			s.currentBlock = b
			err = s.datastore.Flush()
			if err != nil {
				log.Error(err)
			}
		}
		// For centralized exchanges check if trade is not in the block yet
		// (we have observed ws APIs sending identical trades).
		if scrapers.Exchanges[t.Source].Centralized {
			if _, ok := checkTradesDuplicate[t.TradeIdentifierFull()]; !ok {
				s.currentBlock.TradesBlockData.Trades = append(s.currentBlock.TradesBlockData.Trades, t)
				checkTradesDuplicate[t.TradeIdentifierFull()] = struct{}{}
			} else {
				if scrapers.Exchanges[t.Source].Name != dia.BitforexExchange {
					log.Warn("duplicate trade within one tradesblock: ", t)
				}
			}
		} else {
			s.currentBlock.TradesBlockData.Trades = append(s.currentBlock.TradesBlockData.Trades, t)
		}
	} else {
		log.Debugf("ignore trade  %v", t)
	}
}

func (s *TradesBlockService) loadVolumes() map[string]float64 {
	// Clean asset volumes
	volumeCache := make(map[string]float64)
	endtime := time.Now()
	assets, err := s.relDB.GetAssetsWithVolByBlockchain(endtime.AddDate(0, 0, -7), endtime, "")
	if err != nil {
		log.Error("could not load asset with volume: ", err)
	}
	for _, asset := range assets {
		volumeCache[asset.Asset.Identifier()] = asset.Volume
	}
	return volumeCache
}

func (s *TradesBlockService) loadVolumesLoop() {
	for range s.volumeTicker.C {
		s.volumeCache = s.loadVolumes()
	}
}

func (s *TradesBlockService) finaliseCurrentBlock() {

	sort.Slice(s.currentBlock.TradesBlockData.Trades, func(i, j int) bool {
		return s.currentBlock.TradesBlockData.Trades[i].Time.Before(s.currentBlock.TradesBlockData.Trades[j].Time)
	})

	hash, err := structhash.Hash(s.currentBlock.TradesBlockData, 1)
	if err != nil {
		log.Printf("error on hash")
		hash = "hashError"
	}
	s.currentBlock.BlockHash = hash
	s.currentBlock.TradesBlockData.TradesNumber = len(s.currentBlock.TradesBlockData.Trades)
	// Reset duplicate trades identifier.
	checkTradesDuplicate = make(map[string]struct{})
	s.chanTradesBlock <- s.currentBlock
}

func (s *TradesBlockService) ProcessTrade(trade *dia.Trade) {
	s.chanTrades <- trade
}

func (s *TradesBlockService) Close() error {
	if s.closed {
		return errors.New("TradesBlockService: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	return s.error
}

// must only be called from mainLoop
func (s *TradesBlockService) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()
	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.shutdownDone) // signal that shutdown is complete
}

func (s *TradesBlockService) Channel() chan *dia.TradesBlock {
	return s.chanTradesBlock
}

func buildBridge(t dia.Trade) dia.Asset {

	basetoken := t.BaseToken

	// if basetoken.Blockchain == dia.ETHEREUM && basetoken.Address == "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" {
	// 	basetoken = dia.Asset{
	// 		Symbol:     "ETH",
	// 		Address:    "0x0000000000000000000000000000000000000000",
	// 		Blockchain: dia.ETHEREUM,
	// 	}
	// }
	if basetoken.Blockchain == dia.SOLANA && t.Source == dia.OrcaExchange {
		if basetoken.Address == "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v" {
			basetoken = dia.Asset{
				Symbol:     "USDC",
				Address:    "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
				Blockchain: dia.ETHEREUM,
			}
		}
		if basetoken.Address == "Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB" {
			basetoken = dia.Asset{
				Symbol:     "USDT",
				Address:    "0xdAC17F958D2ee523a2206206994597C13D831ec7",
				Blockchain: dia.ETHEREUM,
			}
		}
	}
	if basetoken.Blockchain == dia.METIS && (t.Source == dia.NetswapExchange || t.Source == dia.TethysExchange || t.Source == dia.HermesExchange) && basetoken.Address == "0xEA32A96608495e54156Ae48931A7c20f0dcc1a21" {
		basetoken = dia.Asset{
			Symbol:     "USDC",
			Address:    "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			Blockchain: dia.ETHEREUM,
		}
	}
	if basetoken.Blockchain == dia.FANTOM && (t.Source == dia.SpookyswapExchange || t.Source == dia.SpiritswapExchange || t.Source == dia.BeetsExchange || t.Source == dia.SushiSwapExchangeFantom) {
		if basetoken.Address == "0x21be370D5312f44cB42ce377BC9b8a0cEF1A4C83" {
			basetoken = dia.Asset{
				Symbol:     "FTM",
				Address:    "0x0000000000000000000000000000000000000000",
				Blockchain: dia.FANTOM,
			}
		}
	}
	if basetoken.Blockchain == dia.TELOS && (t.Source == dia.OmniDexExchange) && basetoken.Address == common.HexToAddress("0xd102ce6a4db07d247fcc28f366a623df0938ca9e").Hex() {
		basetoken = dia.Asset{
			Symbol:     "TLOS",
			Address:    "0x0000000000000000000000000000000000000000",
			Blockchain: dia.TELOS,
		}
	}
	if basetoken.Blockchain == dia.EVMOS && t.Source == dia.DiffusionExchange {
		if basetoken.Address == common.HexToAddress("0xD4949664cD82660AaE99bEdc034a0deA8A0bd517").Hex() {
			basetoken = dia.Asset{
				Symbol:     "EVMOS",
				Address:    "0x0000000000000000000000000000000000000000",
				Blockchain: dia.EVMOS,
			}
		}
	}
	if t.Source == dia.StellaswapExchange && basetoken.Blockchain == dia.MOONBEAM {
		if basetoken.Address == common.HexToAddress("0xAcc15dC74880C9944775448304B263D191c6077F").Hex() {
			basetoken = dia.Asset{
				Symbol:     "GLMR",
				Address:    "0x0000000000000000000000000000000000000000",
				Blockchain: dia.MOONBEAM,
			}
		}
	}
	if t.Source == dia.CurveFIExchangeMoonbeam && basetoken.Blockchain == dia.MOONBEAM {
		if basetoken.Address == common.HexToAddress("0xFFFFFFfFea09FB06d082fd1275CD48b191cbCD1d").Hex() {
			basetoken = dia.Asset{
				Symbol:     "USDT",
				Address:    "0xdAC17F958D2ee523a2206206994597C13D831ec7",
				Blockchain: dia.ETHEREUM,
			}
		}
	}
	if (t.Source == dia.UniswapExchangeV3Polygon || t.Source == dia.QuickswapExchange || t.Source == dia.SushiSwapExchangePolygon || t.Source == dia.DfynNetwork) && basetoken.Blockchain == dia.POLYGON {
		if basetoken.Address == common.HexToAddress("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174").Hex() {
			basetoken = dia.Asset{
				Symbol:     "USDC",
				Address:    "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
				Blockchain: dia.ETHEREUM,
			}
		}
		if basetoken.Address == common.HexToAddress("0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270").Hex() {
			basetoken = dia.Asset{
				Symbol:     "MATIC",
				Address:    "0x0000000000000000000000000000000000001010",
				Blockchain: dia.POLYGON,
			}
		}
	}
	if t.Source == dia.ArthswapExchange && basetoken.Blockchain == dia.ASTAR {
		if basetoken.Address == common.HexToAddress("0x6a2d262D56735DbA19Dd70682B39F6bE9a931D98").Hex() {
			basetoken = dia.Asset{
				Symbol:     "USDC",
				Address:    "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
				Blockchain: dia.ETHEREUM,
			}
		}
		if basetoken.Address == common.HexToAddress("0xAeaaf0e2c81Af264101B9129C00F4440cCF0F720").Hex() {
			basetoken = dia.Asset{
				Symbol:     "ASTR",
				Address:    "0x0000000000000000000000000000000000000000",
				Blockchain: dia.ASTAR,
			}
		}
	}
	if basetoken.Blockchain == dia.AVALANCHE && (t.Source == dia.TraderJoeExchange || t.Source == dia.PangolinExchange) {
		if basetoken.Address == common.HexToAddress("0xB31f66AA3C1e785363F0875A1B74E27b85FD66c7").Hex() {
			basetoken = dia.Asset{
				Symbol:     "AVAX",
				Address:    "0x0000000000000000000000000000000000000000",
				Blockchain: dia.AVALANCHE,
			}
		}
	}
	if basetoken.Blockchain == dia.WANCHAIN && t.Source == dia.WanswapExchange {
		if basetoken.Address == common.HexToAddress("0xdabD997aE5E4799BE47d6E69D9431615CBa28f48").Hex() {
			basetoken = dia.Asset{
				Symbol:     "WAN",
				Address:    "0x0000000000000000000000000000000000000000",
				Blockchain: dia.WANCHAIN,
			}
		}
	}
	if basetoken.Blockchain == dia.ARBITRUM && (t.Source == dia.UniswapExchangeV3Arbitrum || t.Source == dia.SushiSwapExchangeArbitrum || t.Source == dia.CamelotExchange || t.Source == dia.TraderJoeExchangeV2_1Arbitrum) {
		if basetoken.Address == common.HexToAddress("0x82aF49447D8a07e3bd95BD0d56f35241523fBab1").Hex() {
			basetoken = dia.Asset{
				Symbol:     "ETH",
				Address:    "0x0000000000000000000000000000000000000000",
				Blockchain: dia.ETHEREUM,
			}
		}
	}
	if basetoken.Blockchain == dia.OSMOSIS && t.Source == dia.OsmosisExchange {
		if basetoken.Address == "ibc/D189335C6E4A68B513C10AB227BF1C1D38C746766278BA3EEB4FB14124F1D858" {
			basetoken = dia.Asset{
				Symbol:     "USDC",
				Address:    "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
				Blockchain: dia.ETHEREUM,
			}
		}
	}
	if basetoken.Blockchain == dia.BIFROST && t.Source == dia.ZenlinkswapExchange {
		if basetoken.Address == "516" {
			basetoken = dia.Asset{
				Symbol:     "KSM",
				Address:    "0x0000000000000000000000000000000000000000",
				Blockchain: dia.KUSAMA,
			}
		}
	}
	if basetoken.Blockchain == dia.BIFROST_POLKADOT && t.Source == dia.ZenlinkswapExchangeBifrostPolkadot {
		if basetoken.Address == "2048" {
			basetoken = dia.Asset{
				Symbol:     "DOT",
				Address:    "0x0000000000000000000000000000000000000000",
				Blockchain: dia.POLKADOT,
			}
		}
	}
	if basetoken.Blockchain == dia.UNREAL_TESTNET && t.Source == dia.PearlfiExchangeTestnet {
		if basetoken.Address == "0x0C68a3C11FB3550e50a4ed8403e873D367A8E361" {
			basetoken = dia.Asset{
				Symbol:     "WETH",
				Address:    "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
				Blockchain: dia.ETHEREUM,
			}
		}
		if basetoken.Address == "0x3F93beBAd7BA4d7A5129eA8159A5829Eacb06497" {
			basetoken = dia.Asset{
				Symbol:     "DAI",
				Address:    "0x6B175474E89094C44Da98b954EedeAC495271d0F",
				Blockchain: dia.ETHEREUM,
			}
		}
	}
	if basetoken.Blockchain == dia.UNREAL && t.Source == dia.PearlfiExchange {
		if basetoken.Address == "0x75d0cBF342060b14c2fC756fd6E717dFeb5B1B70" {
			basetoken = dia.Asset{
				Symbol:     "DAI",
				Address:    "0x6B175474E89094C44Da98b954EedeAC495271d0F",
				Blockchain: dia.ETHEREUM,
			}
		}
	}
	if basetoken.Blockchain == dia.BINANCESMARTCHAIN && t.Source == dia.PanCakeSwapExchangeV3 {
		if basetoken.Address == "0x55d398326f99059fF775485246999027B3197955" {
			basetoken = dia.Asset{
				Symbol:     "USDT",
				Address:    "0xdAC17F958D2ee523a2206206994597C13D831ec7",
				Blockchain: dia.ETHEREUM,
			}
		}
	}
	if basetoken.Blockchain == dia.LINEA && (t.Source == dia.NileV1Exchange || t.Source == dia.NileV2Exchange) {
		if basetoken.Address == "0x176211869cA2b568f2A7D4EE941E073a821EE1ff" {
			basetoken = dia.Asset{
				Symbol:     "USDC",
				Address:    "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
				Blockchain: dia.ETHEREUM,
			}
		}
	}
	if basetoken.Blockchain == dia.OPTIMISM && t.Source == dia.VelodromeExchange {
		if basetoken.Address == "0x7F5c764cBc14f9669B88837ca1490cCa17c31607" {
			basetoken = dia.Asset{
				Symbol:     "USDC",
				Address:    "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
				Blockchain: dia.ETHEREUM,
			}
		}
	}

	return basetoken
}
