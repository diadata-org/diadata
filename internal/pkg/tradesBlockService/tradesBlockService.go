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
	tradeVolumeThresholdExponent, err := strconv.ParseFloat(utils.Getenv("TRADE_VOLUME_THRESHOLD_EXPONENT", ""), 64)
	if err != nil {
		log.Error("Parse TRADE_VOLUME_THRESHOLD_EXPONENT: ", err)
	}
	tradeVolumeThreshold = math.Pow(10, -tradeVolumeThresholdExponent)
}

var (
	stablecoins = map[string]interface{}{
		"USDC": "",
		"USDT": "",
		"TUSD": "",
		"DAI":  "",
		"PAX":  "",
		"BUSD": "",
	}
	tol                  = float64(0.04)
	log                  *logrus.Logger
	batchTimeSeconds     int
	tradeVolumeThreshold float64
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
	priceCache       map[dia.Asset]float64
	datastore        models.Datastore
	historical       bool
	writeMeasurement string
	batchTicker      *time.Ticker
}

func NewTradesBlockService(datastore models.Datastore, blockDuration int64, historical bool) *TradesBlockService {
	s := &TradesBlockService{
		shutdown:        make(chan nothing),
		shutdownDone:    make(chan nothing),
		chanTrades:      make(chan *dia.Trade),
		chanTradesBlock: make(chan *dia.TradesBlock),
		error:           nil,
		started:         false,
		currentBlock:    nil,
		BlockDuration:   blockDuration,
		priceCache:      make(map[dia.Asset]float64),
		datastore:       datastore,
		historical:      historical,
		batchTicker:     time.NewTicker(time.Duration(batchTimeSeconds) * time.Second),
	}
	if historical {
		s.writeMeasurement = utils.Getenv("INFLUX_MEASUREMENT_WRITE", "tradesTmp")
	}
	log.Info("write measurement: ", s.writeMeasurement)
	log.Info("historical: ", s.historical)
	log.Info("batch ticker time: ", batchTimeSeconds)
	go s.mainLoop()
	return s
}

// runs in a goroutine until s is closed
func (s *TradesBlockService) mainLoop() {
	for {
		select {
		case <-s.shutdown:
			log.Println("TradesBlockService shutting down")
			s.cleanup(nil)
			return
		case t := <-s.chanTrades:
			s.process(*t)
		case <-s.batchTicker.C:
			err := s.datastore.Flush()
			if err != nil {
				log.Error("flush influx batch: ", err)
			}
		}
	}
}

func (s *TradesBlockService) process(t dia.Trade) {

	var verifiedTrade bool

	// Price estimation can only be done for verified pairs.
	// Trades with unverified pairs are still saved, but not sent to the filtersBlockService.
	if t.VerifiedPair && s.checkTrade(t) {
		if t.BaseToken.Address == "840" && t.BaseToken.Blockchain == dia.FIAT {
			// All prices are measured in US-Dollar, so just price for base token == USD
			t.EstimatedUSDPrice = t.Price
			verifiedTrade = true
		} else {
			// Get price of base token.
			var quotation *models.AssetQuotation
			var price float64
			var ok bool
			var err error
			if !s.historical {

				// Bridge basetoken if necessary.
				basetoken := buildBridge(t)

				// Get latest price from cache.
				if _, ok = s.priceCache[basetoken]; ok {
					price = s.priceCache[basetoken]
				} else {
					quotation, err = s.datastore.GetAssetQuotationCache(basetoken)
					price = quotation.Price
					s.priceCache[basetoken] = price
					log.Infof("quotation for %s from redis cache: %v", basetoken.Symbol, price)
				}

			} else {

				// Look for historic price of base token at trade time.
				if _, ok = s.priceCache[t.BaseToken]; ok {
					price = s.priceCache[t.BaseToken]
				} else {
					price, err = s.datastore.GetAssetPriceUSD(t.BaseToken, t.Time)
					s.priceCache[t.BaseToken] = price
					if t.BaseToken.Address == "0x0000000000000000000000000000000000000000" {
						if t.BaseToken.Blockchain == "Bitcoin" {
							log.Infof("quotation for BTC from influx: %v", price)
						}
						if t.BaseToken.Blockchain == "Ethereum" {
							log.Infof("quotation for ETH from influx: %v", price)
						}
					}
				}

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
					if t.EstimatedUSDPrice > 0 {
						verifiedTrade = true
					}
				}
			}
		}
	}

	// // If estimated price for stablecoin diverges too much ignore trade
	if _, ok := stablecoins[t.Symbol]; ok {
		if math.Abs(t.EstimatedUSDPrice-1) > tol {
			log.Errorf("price for stablecoin %s diverges by %v", t.Symbol, math.Abs(t.EstimatedUSDPrice-1))
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
				s.priceCache = make(map[dia.Asset]float64)
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
		s.currentBlock.TradesBlockData.Trades = append(s.currentBlock.TradesBlockData.Trades, t)
	} else {
		log.Debugf("ignore trade  %v", t)
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

func (s *TradesBlockService) checkTrade(t dia.Trade) bool {
	if math.Abs(t.Volume) < tradeVolumeThreshold {
		log.Info("low volume trade: ", t)
		return false
	}
	return true
}

func buildBridge(t dia.Trade) dia.Asset {

	basetoken := t.BaseToken

	if basetoken.Blockchain == dia.ETHEREUM && basetoken.Address == "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" {
		basetoken = dia.Asset{
			Symbol:     "ETH",
			Address:    "0x0000000000000000000000000000000000000000",
			Blockchain: dia.ETHEREUM,
		}
	}
	if basetoken.Blockchain == dia.SOLANA && t.Source == dia.SerumExchange {
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
	if basetoken.Blockchain == dia.FANTOM && (t.Source == dia.SpookyswapExchange || t.Source == dia.SpiritswapExchange || t.Source == dia.BeetsExchange) {
		if basetoken.Address == "0x21be370D5312f44cB42ce377BC9b8a0cEF1A4C83" {
			basetoken = dia.Asset{
				Symbol:     "FTM",
				Address:    "0x0000000000000000000000000000000000000000",
				Blockchain: dia.FANTOM,
			}
		}
		if basetoken.Address == "0x04068DA6C83AFCFA0e13ba15A6696662335D5B75" {
			basetoken = dia.Asset{
				Symbol:     "USDC",
				Address:    "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
				Blockchain: dia.ETHEREUM,
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
	if t.Source == dia.StellaswapExchange && basetoken.Blockchain == dia.MOONBEAM && basetoken.Address == common.HexToAddress("0xAcc15dC74880C9944775448304B263D191c6077F").Hex() {
		basetoken = dia.Asset{
			Symbol:     "GLMR",
			Address:    "0x0000000000000000000000000000000000000000",
			Blockchain: dia.MOONBEAM,
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
	if basetoken.Blockchain == dia.ARBITRUM && t.Source == dia.UniswapExchangeV3Arbitrum {
		if basetoken.Address == common.HexToAddress("0x82aF49447D8a07e3bd95BD0d56f35241523fBab1").Hex() {
			basetoken = dia.Asset{
				Symbol:     "ETH",
				Address:    "0x0000000000000000000000000000000000000000",
				Blockchain: dia.ETHEREUM,
			}
		}
	}
	return basetoken
}
