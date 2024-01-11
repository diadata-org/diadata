package scrapers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/kr/pretty"
	"github.com/sirupsen/logrus"
	"github.com/stellar/go/clients/horizonclient"
	hProtocol "github.com/stellar/go/protocols/horizon"
)

const (
	// TODO check me
	stellarRefreshDelay  = time.Second * 20 * 60
	stellarBaseAssetType = "native"
	stellarBaseAssetName = "XLM"
	stellarDecimals      = 7
)

var stellarDefaultTradeCursor = "" // get data from now()

/*
	Helpful tools
https://laboratory.stellar.org/#explorer?resource=trades&endpoint=all&network=test
https://stellar.expert/explorer
https://developers.stellar.org/api/horizon/resources/trades/object
https://developers.stellar.org/api/horizon

GET https://horizon-testnet.stellar.org/trades

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
	cursor        *string
	chanTrades    chan *dia.Trade
	db            *models.RelDB
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

// NewStellarScraper returns a new NewStellarScraper initialized with default values.
// The instance is asynchronously scraping as soon as it is created.
func NewStellarScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *StellarScraper {
	cursor := utils.Getenv(strings.ToUpper(exchange.Name)+"_CURSOR", "")
	if cursor == "" {
		cursor = stellarDefaultTradeCursor
	}
	s := &StellarScraper{
		logger: log.WithFields(logrus.Fields{
			"context": "StellarScraper",
			"cursor":  cursor,
		}),
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*StellarPairScraper),
		exchangeName: exchange.Name,
		error:        nil,
		cursor:       &cursor,
		// can be DefaultTestNetClient too
		horizonClient: horizonclient.DefaultPublicNetClient,
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
	baseSymbol := stellarTrade.BaseAssetCode
	if stellarTrade.BaseAssetType == stellarBaseAssetType {
		baseSymbol = stellarBaseAssetName
	}
	targetSymbol := stellarTrade.CounterAssetCode
	if stellarTrade.CounterAssetType == stellarBaseAssetType {
		targetSymbol = stellarBaseAssetName
	}

	// TODO get asset info from stellar
	token0 := dia.Asset{
		Address:  stellarTrade.BaseAccount, // TODO check me
		Symbol:   baseSymbol,
		Name:     baseSymbol, // TODO
		Decimals: stellarDecimals,
		// Blockchain: Exchanges[s.exchangeName].BlockChain.Name, // TODO
	}
	token1 := dia.Asset{
		Address: stellarTrade.CounterAccount, // TODO check me
		Symbol:  targetSymbol,
		Name:    targetSymbol,
		//Name:       pair.Token1.Name, // TODO
		Decimals: stellarDecimals,
		// Blockchain: Exchanges[s.exchangeName].BlockChain.Name, // TODO
	}
	volume, _ := strconv.ParseFloat(stellarTrade.CounterAmount, 64)
	symbolPair := getSymbolPairs(baseSymbol, targetSymbol)
	diaTrade := &dia.Trade{
		Symbol:         symbolPair,
		Pair:           symbolPair,
		ForeignTradeID: stellarTrade.ID,
		Source:         s.exchangeName,
		Time:           stellarTrade.LedgerCloseTime,
		Price:          float64(stellarTrade.Price.N),
		Volume:         volume,
		VerifiedPair:   true,
		BaseToken:      token0,
		QuoteToken:     token1,
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
		logger: log.WithFields(logrus.Fields{
			"context": "HorizonPairScraper",
		}),
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
	response, err := http.Get("https://api.stellar.expert/explorer/public/asset/?order=desc&sort=rating")
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

type StellarPairScraper struct {
	logger     *logrus.Entry
	parent     *StellarScraper
	pair       dia.ExchangePair
	closed     bool
	lastRecord int64
}

func (ps *StellarPairScraper) Pair() dia.ExchangePair {
	// ps.logger.Infof("StellarPairScraper.Pair %# v", pretty.Formatter(ps.pair))
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

func getSymbolPairs(baseSymbol, quotedSymbol string) string {
	result := fmt.Sprintf("%s-%s", baseSymbol, quotedSymbol)
	return result
}
