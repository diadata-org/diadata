package scrapers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/horizonhelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/kr/pretty"
	"github.com/sirupsen/logrus"
	"github.com/stellar/go/clients/horizonclient"
	hProtocol "github.com/stellar/go/protocols/horizon"
)

const (
	stellarExpertUrl           = "https://api.stellar.expert/explorer/public/asset/?order=desc&sort=rating"
	defaultTradesRequestCursor = ""
	defaultTradesRequestLimit  = 200
)

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
	chanTrades    chan *dia.Trade
	blockchain    string
	exchangeName  string
	relDB         *models.RelDB
	cursor        *string
	limit         uint
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
	cursor := utils.Getenv(strings.ToUpper(exchange.Name)+"_TRADES_CURSOR", defaultTradesRequestCursor)
	limit := utils.GetenvUint(strings.ToUpper(exchange.Name)+"_TRADES_LIMIT", defaultTradesRequestLimit)

	var (
		shutdown     = make(chan nothing)
		shutdownDone = make(chan nothing)
		pairScrapers = make(map[string]*StellarPairScraper)
		chanTrades   = make(chan *dia.Trade)
		logger       = logrus.New().WithFields(logrus.Fields{
			"context": "StellarExchangeScraper",
			"cursor":  cursor,
		})
	)

	horizonClient := horizonclient.DefaultPublicNetClient

	s := &StellarScraper{
		logger:        logger,
		shutdown:      shutdown,
		shutdownDone:  shutdownDone,
		error:         nil,
		pairScrapers:  pairScrapers,
		horizonClient: horizonClient,
		chanTrades:    chanTrades,
		blockchain:    exchange.BlockChain.Name,
		exchangeName:  exchange.Name,
		relDB:         relDB,
		cursor:        &cursor,
		limit:         limit,
	}

	if scrape {
		go s.mainLoop()
	}
	return s
}

func (s *StellarScraper) mainLoop() {
	tradeRequest := horizonclient.TradeRequest{
		Cursor:    *s.cursor,
		Limit:     s.limit,
		TradeType: "orderbook",
	}
	ctx, _ := context.WithCancel(context.Background())

	err := s.horizonClient.StreamTrades(ctx, tradeRequest, s.tradeHandler)
	if err != nil {
		s.logger.WithError(err).Error("failed to stream trades")
	}
}

func (s *StellarScraper) tradeHandler(stellarTrade hProtocol.Trade) {
	log := s.logger.WithFields(logrus.Fields{
		"context": "StellarExchangeScraper.tradeHandler",
		"tradeID": stellarTrade.ID,
	})

	baseToken, err := s.getDIAAsset(stellarTrade.BaseAssetType, stellarTrade.BaseAssetCode, stellarTrade.BaseAssetIssuer)
	if err != nil {
		log.
			WithError(err).
			Error("failed to get base asset of a trade")
		return
	}

	quoteToken, err := s.getDIAAsset(stellarTrade.CounterAssetType, stellarTrade.CounterAssetCode, stellarTrade.CounterAssetIssuer)
	if err != nil {
		log.
			WithError(err).
			Error("failed to get counter asset of a trade")
		return
	}

	volumeOut := stellarTrade.CounterAmount
	if stellarTrade.BaseIsSeller {
		baseToken, quoteToken = quoteToken, baseToken
		volumeOut = stellarTrade.BaseAmount
	}
	symbolPair := horizonhelper.GetAssetSymbolPair(baseToken.Symbol, quoteToken.Symbol)

	price := float64(stellarTrade.Price.N) / float64(stellarTrade.Price.D)
	volume, err := strconv.ParseFloat(volumeOut, 64)
	if err != nil {
		log.
			WithError(err).
			Error("failed to parse volume of a trade")
		return
	}

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

	log.Infof("stellarTrade %# v", pretty.Formatter(stellarTrade))
	log.Infof("diaTrade %# v", pretty.Formatter(diaTrade))

	s.chanTrades <- diaTrade
}

func (s *StellarScraper) getDIAAsset(assetType, assetCode, assetIssuer string) (asset dia.Asset, err error) {
	s.logger.
		WithFields(logrus.Fields{
			"assetType":   assetType,
			"assetCode":   assetCode,
			"assetIssuer": assetIssuer,
		}).Info("fetching asset info")

	if assetType == "native" {
		asset = horizonhelper.GetStellarNativeAssetInfo(s.blockchain)
		return
	}

	if assetCode == "" || assetIssuer == "" {
		err = errors.New("assetCode or assetIssuer is empty")
		return
	}

	assetID := horizonhelper.GetAssetAddress(assetCode, assetIssuer)
	asset, err = s.relDB.GetAsset(assetID, s.blockchain)
	if err == nil {
		return
	}
	err = nil

	asset, err = horizonhelper.GetStellarAssetInfo(s.horizonClient, assetCode, assetIssuer, s.blockchain)
	if err != nil {
		log.Error("cannot fetch asset info: ", err)
	}
	return
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
		return nil, errors.New("StellarExchangeScraper: Call ScrapePair on closed scraper")
	}
	ps := &StellarPairScraper{
		parent:     s,
		pair:       pair,
		lastRecord: 0,
	}

	s.pairScrapers[pair.Symbol] = ps

	return ps, nil
}

func (s *StellarScraper) Close() error {
	if s.closed {
		return errors.New("HorizonScraper: Already closed")
	}
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

	body, err := io.ReadAll(response.Body)

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
	return dia.Asset{Symbol: symbol}, nil
}

func (s *StellarScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	// as for now we do not need to normalize the pair
	return pair, nil
}

// Channel returns a channel that can be used to receive trades
func (s *StellarScraper) Channel() chan *dia.Trade {
	return s.chanTrades
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
