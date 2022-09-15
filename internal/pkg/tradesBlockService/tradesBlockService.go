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
	batchTimeString := utils.Getenv("BATCH_TIME_SECONDS", "30")
	var err error
	batchTimeSeconds, err = strconv.Atoi(batchTimeString)
	if err != nil {
		log.Error("parse batchTimeString: ", err)
	}

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
	USDT             = dia.Asset{Address: "0xdAC17F958D2ee523a2206206994597C13D831ec7", Blockchain: dia.ETHEREUM}
	tol              = float64(0.04)
	log              *logrus.Logger
	batchTimeSeconds int
	// volumeUpdateSeconds = 60 * 60 * 6
	volumeUpdateSeconds = 60 * 10
	volumeThreshold     = float64(100000)
	blueChipThreshold   = float64(50000000)
	smallX              = float64(10)
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
	s.volumeCache = s.loadVolumes()
	log.Info("...done loading volumes.")

	go s.mainLoop()
	go s.loadVolumesLoop()
	return s
}

// runs in a goroutine until s is closed
func (s *TradesBlockService) mainLoop() {
	var (
		acceptCount        int
		acceptCountCEX     int
		acceptCountDEX     int
		acceptCountSwapCEX int
		acceptCountSwapDEX int
		totalCount         int
		logTicker          = *time.NewTicker(120 * time.Second)
	)
	for {
		select {
		case <-s.shutdown:
			log.Println("TradesBlockService shutting down")
			s.cleanup(nil)
			return
		case t := <-s.chanTrades:
			tSwapped, err := dia.SwapTrade(*t)
			if err != nil {
				log.Error("swap trade: ", err)
			}
			tradeOk := s.checkTrade(*t)
			swapppedTradeOk := s.checkTrade(tSwapped)
			if tradeOk {
				s.process(*t)
				log.Info("source: ", (*t).Source)
				log.Info("Exchange: ", scrapers.Exchanges[(*t).Source])
				if scrapers.Exchanges[(*t).Source].Centralized {
					acceptCountCEX++
				} else {
					acceptCountDEX++
				}
			}
			if swapppedTradeOk {
				s.process(tSwapped)
				if scrapers.Exchanges[(*t).Source].Centralized {
					acceptCountSwapCEX++
				} else {
					acceptCountSwapDEX++
				}
			}

			if tradeOk || swapppedTradeOk {
				acceptCount++
			}
			totalCount++

		case <-s.batchTicker.C:
			err := s.datastore.Flush()
			if err != nil {
				log.Error("flush influx batch: ", err)
			}
		case <-logTicker.C:
			log.Info("accepted trades CEX: ", acceptCountCEX)
			log.Info("accepted swapped trades CEX: ", acceptCountSwapCEX)
			log.Info("accepted trades DEX: ", acceptCountDEX)
			log.Info("accepted swapped trades DEX: ", acceptCountSwapDEX)
			log.Info("discarded trades: ", totalCount-acceptCount)
		}
	}
}

// checkTrade determines whether a trade should be taken into account for price determination.
func (s *TradesBlockService) checkTrade(t dia.Trade) bool {
	basetoken := buildBridge(t)

	if basetoken.Blockchain == dia.FIAT {
		return true
	}
	if t.QuoteToken.Blockchain == dia.FIAT {
		return false
	}
	if t.QuoteToken.Address == "0xdAC17F958D2ee523a2206206994597C13D831ec7" && t.QuoteToken.Blockchain == dia.ETHEREUM {
		if basetoken.Blockchain != dia.FIAT {
			return false
		}
	}

	if baseVolume, ok := s.volumeCache[assetIdentifier(basetoken)]; ok {
		if baseVolume > blueChipThreshold {
			return true
		}
		if quoteVolume, ok := s.volumeCache[assetIdentifier(t.QuoteToken)]; ok {
			if baseVolume < volumeThreshold {
				// For small volume assets, quote asset must be a small volume asset too.
				return quoteVolume < smallX*baseVolume
			}
		}
		// Base asset has enough volume or quote asset has no volume yet.
		return true
	}
	return false
}

func (s *TradesBlockService) process(t dia.Trade) {

	var verifiedTrade bool

	// Price estimation can only be done for verified pairs.
	// Trades with unverified pairs are still saved, but not sent to the filtersBlockService.
	if t.VerifiedPair {
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
			if !s.historical {
				// Get latest price from cache.
				// price, err = s.datastore.GetAssetPriceUSDLatest(t.BaseToken)

				basetoken := buildBridge(t)

				if _, ok = s.priceCache[assetIdentifier(basetoken)]; ok {
					price = s.priceCache[assetIdentifier(basetoken)]
				} else {
					quotation, err = s.datastore.GetAssetQuotationCache(basetoken)
					price = quotation.Price
					s.priceCache[assetIdentifier(basetoken)] = price
					// log.Infof("quotation for %s from redis cache: %v", basetoken.Symbol, price)
				}

			} else {

				basetoken := t.BaseToken
				// Look for historic price of base token at trade time.
				if _, ok = s.priceCache[assetIdentifier(basetoken)]; ok {
					price = s.priceCache[assetIdentifier(basetoken)]
				} else {
					price, err = s.datastore.GetAssetPriceUSD(t.BaseToken, t.Time)
					s.priceCache[assetIdentifier(basetoken)] = price
					if t.BaseToken.Address == "0x0000000000000000000000000000000000000000" {
						if basetoken.Blockchain == "Bitcoin" {
							log.Infof("quotation for BTC from influx: %v", price)
						}
						if basetoken.Blockchain == "Ethereum" {
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
		s.currentBlock.TradesBlockData.Trades = append(s.currentBlock.TradesBlockData.Trades, t)
	} else {
		log.Debugf("ignore trade  %v", t)
	}
}

func (s *TradesBlockService) loadVolumes() map[string]float64 {
	// Clean asset volumes
	volumeCache := make(map[string]float64)
	endtime := time.Now()
	assets, err := s.relDB.GetAssetsWithVOLRange(endtime.AddDate(0, 0, -7), endtime)
	if err != nil {
		log.Error("could not load asset with volume: ", err)
	}
	for _, asset := range assets {
		volumeCache[assetIdentifier(asset.Asset)] = asset.Volume
	}
	return volumeCache
}

func (s *TradesBlockService) loadVolumesLoop() {
	for range s.volumeTicker.C {
		s.volumeCache = s.loadVolumes()
	}
}

// assetIdentifier returns the unique identifier blockchain-address for an asset.
func assetIdentifier(asset dia.Asset) string {
	return asset.Blockchain + "-" + asset.Address
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

func buildBridge(t dia.Trade) dia.Asset {

	basetoken := t.BaseToken

	if basetoken.Blockchain == dia.SOLANA && t.Source == dia.SerumExchange && basetoken.Address == "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v" {
		basetoken = dia.Asset{
			Symbol:     "USDC",
			Address:    "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			Blockchain: dia.ETHEREUM,
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
	if basetoken.Blockchain == dia.EVMOS && t.Source == dia.DiffusionExchange && basetoken.Address == common.HexToAddress("0x51e44FfaD5C2B122C8b635671FCC8139dc636E82").Hex() {
		basetoken = dia.Asset{
			Symbol:     "USDC",
			Address:    "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			Blockchain: dia.ETHEREUM,
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
	return basetoken
}
