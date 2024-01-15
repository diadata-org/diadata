package scrapers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/kr/pretty"
	"github.com/sirupsen/logrus"
	"github.com/stellar/go/clients/horizonclient"
	hProtocol "github.com/stellar/go/protocols/horizon"
)

const (
	stellarTestChainName   = "test"
	stellarPublicChainName = "public"
	stellarExpertUrl       = "https://api.stellar.expert/explorer/public/asset/?order=desc&sort=rating"
)

/*
	Helpful urls
https://laboratory.stellar.org/#explorer?resource=trades&endpoint=all&network=test
https://stellar.expert/explorer
https://developers.stellar.org/api/horizon/resources/trades/object
https://developers.stellar.org/api/horizon

https://horizon-testnet.stellar.org/trades
https://horizon-testnet.stellar.org/assets

https://horizon.stellar.org/trades
https://horizon.stellar.org/assets

ENV variables:
	STELLAREXCHANGE_CURSOR - number, pagination, default "" - means - stream latest trades from now
	STELLAREXCHANGE_CHAIN - string, enum(public, test), default public
*/

type StellarScraper struct {
	logger *logrus.Entry
	// signaling channels
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock     sync.RWMutex
	error         error
	closed        bool
	pairScrapers  map[string]*StellarPairScraper // pc.ExchangePair -> pairScraperSet
	horizonClient *horizonclient.Client
	exchangeName  string
	chainName     string
	cursor        *string
	chanTrades    chan *dia.Trade
	db            *models.RelDB
	cachedAssets  sync.Map // map[string]dia.Asset
}

type StellarExpertAsset struct {
	Asset    string      `json:"asset"`
	Domain   string      `json:"domain"`
	Price7D  [][]float64 `json:"price7d"`
	TomlInfo struct {
		Code            string `json:"code"`
		Issuer          string `json:"issuer"`
		Name            string `json:"name"`
		Image           string `json:"image"`
		AnchorAssetType string `json:"anchorAssetType"`
		AnchorAsset     string `json:"anchorAsset"`
		OrgName         string `json:"orgName"`
		OrgLogo         string `json:"orgLogo"`
	} `json:"tomlInfo"`
}

type StellarExpertAssets struct {
	Embedded struct {
		Records []StellarExpertAsset `json:"records"`
	} `json:"_embedded"`
}

// NewStellarScraper returns a new StellarScraper initialized with default values.
// The instance is asynchronously scraping as soon as it is created.
func NewStellarScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *StellarScraper {
	cursor := utils.Getenv(strings.ToUpper(exchange.Name)+"_CURSOR", "")
	chainName := utils.Getenv(strings.ToUpper(exchange.Name)+"_CHAIN", "")

	horizonClient := horizonclient.DefaultPublicNetClient
	if chainName == "" {
		chainName = stellarPublicChainName
	}
	if chainName == stellarTestChainName {
		horizonClient = horizonclient.DefaultTestNetClient
	}
	s := &StellarScraper{
		logger: log.WithFields(logrus.Fields{
			"context": "StellarScraper",
			"chain":   chainName,
			"cursor":  cursor,
		}),
		shutdown:      make(chan nothing),
		shutdownDone:  make(chan nothing),
		pairScrapers:  make(map[string]*StellarPairScraper),
		exchangeName:  exchange.Name,
		chainName:     chainName,
		error:         nil,
		cursor:        &cursor,
		horizonClient: horizonClient,
		chanTrades:    make(chan *dia.Trade),
		db:            relDB,
	}

	if scrape {
		go s.mainLoop()
	}
	return s
}

func (s *StellarScraper) mainLoop() {
	tradeRequest := horizonclient.TradeRequest{
		Cursor:    *s.cursor,
		TradeType: "orderbook",
	}
	ctx, _ := context.WithCancel(context.Background())
	//go func() {
	//	// Stop streaming after 60 seconds.
	//	s.logger.Print("Waiting data before sleep")
	//	time.Sleep(horizonRefreshDelay)
	//	cancel()
	//}()

	err := s.horizonClient.StreamTrades(ctx, tradeRequest, s.tradeHandler)
	if err != nil {
		s.logger.WithError(err).Error("failed to stream trades")
	}
}

func (s *StellarScraper) tradeHandler(stellarTrade hProtocol.Trade) {
	token1 := dia.Asset{}
	token2 := dia.Asset{}

	skipBase := false
	skipCounter := false

	var cachedErr1 error
	var cachedErr2 error

	if stellarTrade.BaseAssetIssuer == "" || stellarTrade.BaseAssetCode == "" {
		skipBase = true
	}
	if stellarTrade.CounterAssetIssuer == "" || stellarTrade.CounterAssetCode == "" {
		skipCounter = true
	}
	if !skipBase {
		token1, cachedErr1 = s.getTokenInfoAndCache(stellarTrade.BaseAssetCode, stellarTrade.BaseAssetIssuer)
	}
	if !skipCounter {
		token2, cachedErr2 = s.getTokenInfoAndCache(stellarTrade.CounterAssetCode, stellarTrade.CounterAssetIssuer)
	}

	if skipBase {
		s.logger.
			WithField("BaseAssetCode", stellarTrade.BaseAssetCode).
			WithField("BaseAssetIssuer", stellarTrade.BaseAssetIssuer).
			WithField("ID", stellarTrade.ID).
			Warn("BaseAssetIssuer.Impossible to get base asset address.Skip")
		return
	}
	if skipCounter {
		s.logger.
			WithField("CounterAssetCode", stellarTrade.CounterAssetCode).
			WithField("CounterAssetIssuer", stellarTrade.CounterAssetIssuer).
			WithField("ID", stellarTrade.ID).
			Warn("CounterAssetIssuer.Impossible to get counter asset address.Skip")
		return
	}

	if cachedErr1 != nil {
		s.logger.
			WithField("ID", stellarTrade.ID).
			WithError(cachedErr1).
			Error("failed to get and cache baseToken.")
		return
	}
	if cachedErr2 != nil {
		s.logger.
			WithField("ID", stellarTrade.ID).
			WithError(cachedErr2).
			Error("failed to get and cache counterToken.")
		return
	}
	baseToken, quoteToken := s.getAssetPair(stellarTrade, token1, token2)

	price, volume := s.getPriceAndVolumeFromStellarTradeData(stellarTrade)
	symbolPair := s.getSymbolPairs(baseToken.Symbol, quoteToken.Symbol)

	diaTrade := &dia.Trade{
		Symbol:         symbolPair,
		Pair:           symbolPair,
		ForeignTradeID: stellarTrade.ID,
		Source:         s.exchangeName,
		Time:           stellarTrade.LedgerCloseTime,
		Price:          price,
		Volume:         volume,
		VerifiedPair:   true,
		BaseToken:      baseToken,
		QuoteToken:     quoteToken,
	}
	s.logger.Infof("StellarScraper.tradeHandler.stellarTrade %# v", pretty.Formatter(stellarTrade))
	s.logger.Infof("StellarScraper.tradeHandler.diaTrade %# v", pretty.Formatter(diaTrade))
	s.chanTrades <- diaTrade
}

func (s *StellarScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()
	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.shutdownDone) // signal that shutdown is complete
}

func (s *StellarScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("StellarScraper: Call ScrapePair on closed scraper")
	}
	ps := &StellarPairScraper{
		parent:     s,
		pair:       pair,
		lastRecord: 0, //TODO FIX to figure out the last we got...
	}

	s.pairScrapers[pair.Symbol] = ps

	// s.logger.Infof("ScrapePair %# v", pretty.Formatter(pair))

	return ps, nil
}

func (s *StellarScraper) Close() error {
	if s.closed {
		return errors.New("HorizonScraper: Already closed")
	}
	// close all related connections
	//err := s.horizonClient.HTTP.Close()
	//if err != nil {
	//	s.logger.Error(err)
	//}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

func (s *StellarScraper) FetchAvailablePairs() ([]dia.ExchangePair, error) {
	var stellarExpertAssets StellarExpertAssets
	response, err := http.Get(stellarExpertUrl)
	if err != nil {
		s.logger.WithError(err).Error("failed to get symbols: ")
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		errClose := Body.Close()
		if errClose != nil {
			s.logger.WithError(errClose).Error("failed to body close got error ")
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		s.logger.WithError(err).Error("failed to read symbols")
		return nil, err
	}

	err = json.Unmarshal(body, &stellarExpertAssets)

	if err != nil {
		s.logger.WithError(err).Error("failed to unmarshal symbols")
		return nil, err
	}

	pairs := make([]dia.ExchangePair, len(stellarExpertAssets.Embedded.Records))
	for index, record := range stellarExpertAssets.Embedded.Records {
		pairToNormalize := dia.ExchangePair{
			ForeignName: "USD",
			Exchange:    s.exchangeName,
		}
		symbol := strings.Split(record.Asset, `-`)
		pairToNormalize.Symbol = symbol[0]
		pairToNormalize.ForeignName = fmt.Sprintf("%s-USD", pairToNormalize.Symbol)

		pair, serr := s.NormalizePair(pairToNormalize)
		if serr == nil {
			pairs[index] = pair
		} else {
			s.logger.WithError(serr).Error("failed to NormalizePair")
		}
	}

	return pairs, nil
}

func (s *StellarScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	// s.logger.Infof("StellarScraper.FillSymbolData %# v", pretty.Formatter(symbol))
	return dia.Asset{Symbol: symbol}, nil
}

func (s *StellarScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	// s.logger.Infof("StellarScraper.NormalizePair %# v", pretty.Formatter(pair))
	return pair, nil
}

// Channel returns a channel that can be used to receive trades
func (s *StellarScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

func (s *StellarScraper) getTokenInfoAndCache(assetCode, assetIssuer string) (assetToken dia.Asset, err error) {
	assetAddress := s.getSymbolPairs(assetCode, assetIssuer)
	cached, ok := s.cachedAssets.Load(assetAddress)
	if !ok {
		assetInfoReader := &ethhelper.StellarAssetInfo{
			Logger: log.WithFields(logrus.Fields{
				"context": "StellarTomlReader",
			}),
		}
		asset, err := assetInfoReader.GetStellarAssetInfo(s.horizonClient, assetCode, assetIssuer, s.chainName)
		if err != nil {
			return dia.Asset{}, err
		}

		s.cachedAssets.Store(assetAddress, asset)

		s.logger.
			WithField("action", "cached").
			Infof("assetFromToken.asset.cached %v", pretty.Formatter(asset))

		return asset, nil
	}

	asset := cached.(dia.Asset)

	s.logger.
		WithField("action", "fromCache").
		Infof("assetFromToken.asset.fromCache %v", pretty.Formatter(asset))

	return asset, nil
}

func (s *StellarScraper) getAssetPair(stellarTrade hProtocol.Trade, token1, token2 dia.Asset) (bToken, qToken dia.Asset) {
	if stellarTrade.BaseIsSeller {
		return token2, token1
	}
	return token1, token2
}

// getPriceAndVolumeFromStellarData returns price, volume
func (s *StellarScraper) getPriceAndVolumeFromStellarTradeData(stellarTrade hProtocol.Trade) (price, volume float64) {
	//// if baseIsSeller=true
	////		price = stellarTrade.Price.D / stellarTrade.Price.N
	//// else
	////		price = stellarTrade.Price.N / stellarTrade.Price.D

	// TODO check the commented code
	//decimalsInt := big.NewInt(stellarDecimals)
	//decimalsDivisorInt := decimalsInt.Exp(big.NewInt(10), decimalsInt, nil)
	//decimalsDivisorFloat := big.NewFloat(0).SetInt(decimalsDivisorInt)
	//
	//priceNFloat := big.NewFloat(0).SetInt64(stellarTrade.Price.N)
	//priceDFloat := big.NewFloat(0).SetInt64(stellarTrade.Price.D)
	//amount0In := new(big.Float).Quo(priceNFloat, decimalsDivisorFloat)
	//amount1Out := new(big.Float).Quo(priceDFloat, decimalsDivisorFloat)
	//if stellarTrade.BaseIsSeller {
	//	amount0In = new(big.Float).Quo(priceDFloat, decimalsDivisorFloat)
	//	amount1Out = new(big.Float).Quo(priceNFloat, decimalsDivisorFloat)
	//}
	//// _resultVolume := amount1Out.Quo(amount0In, decimalsDivisorFloat)
	//resultVolume, _ := amount1Out.Float64()
	//resultPrice, _ := amount0In.Float64()

	//return resultPrice, resultVolume
	amount0In := stellarTrade.Price.N
	amount1Out := stellarTrade.Price.D
	if stellarTrade.BaseIsSeller {
		amount0In = stellarTrade.Price.D
		amount1Out = stellarTrade.Price.N
	}
	resultPrice := float64(amount1Out) / float64(amount0In)
	resultVolume := float64(amount1Out)
	return resultPrice, resultVolume
}

func (s *StellarScraper) getSymbolPairs(bs, cs string) string {
	result := fmt.Sprintf("%s-%s", bs, cs)
	return result
}

type StellarPairScraper struct {
	parent     *StellarScraper
	pair       dia.ExchangePair
	closed     bool
	lastRecord int64
}

func (ps *StellarPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}

func (ps *StellarPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *StellarPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}
