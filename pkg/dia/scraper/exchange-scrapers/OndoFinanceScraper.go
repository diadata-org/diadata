package scrapers

import (
	"context"
	"errors"
	"math"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"

	ondo "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/ondofinance"
	models "github.com/diadata-org/diadata/pkg/model"

	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
)

const PRECISION = 18

type OndoFinanceTrade struct {
	ID        string
	Timestamp time.Time
	Pair      dia.Pair
	Price     float64
	Amount    float64
}

type OndoFinanceScraper struct {
	WsClient   *ethclient.Client
	RestClient *ethclient.Client
	relDB      *models.RelDB
	baseToken  dia.Asset
	// signaling channels for session initialization and finishing
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	pairScrapers map[string]*OndoFinancePairScraper

	exchangeName           string
	chanTrades             chan *dia.Trade
	factoryContractAddress common.Address
}

// NewOndoFinanceScraper returns a new OndoFinanceScraper
func NewOndoFinanceScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) (s *OndoFinanceScraper) {
	log.Infof("New %s scraper with address %s", exchange.Name, exchange.Contract)

	s = makeOndoFinanceScraper(exchange, "", "")

	s.relDB = relDB
	s.baseToken = dia.Asset{
		Symbol:     "USDC",
		Name:       "USD Coin",
		Address:    "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
		Blockchain: dia.ETHEREUM,
		Decimals:   uint8(6),
	}

	pingNodeInterval, err := strconv.ParseInt(utils.Getenv("PING_SERVER", "0"), 10, 64)
	if err != nil {
		log.Error("parse PING_SERVER: ", err)
	}
	if pingNodeInterval > 0 {
		s.pingNode(pingNodeInterval)
	}

	if scrape {
		go s.mainLoop()
	}
	return s
}

// makeOndoFinanceScraper returns an OndoFinance scraper as used in NewOndoFinanceScraper.
func makeOndoFinanceScraper(exchange dia.Exchange, restDial string, wsDial string) *OndoFinanceScraper {
	var restClient, wsClient *ethclient.Client
	var err error
	var s *OndoFinanceScraper

	log.Infof("Init rest and ws client for %s.", exchange.BlockChain.Name)
	restClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", restDial))
	if err != nil {
		log.Fatal("init rest client: ", err)
	}
	wsClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_WS", wsDial))
	if err != nil {
		log.Fatal("init ws client: ", err)
	}

	s = &OndoFinanceScraper{
		WsClient:               wsClient,
		RestClient:             restClient,
		shutdown:               make(chan nothing),
		shutdownDone:           make(chan nothing),
		pairScrapers:           make(map[string]*OndoFinancePairScraper),
		exchangeName:           exchange.Name,
		error:                  nil,
		chanTrades:             make(chan *dia.Trade),
		factoryContractAddress: common.HexToAddress(exchange.Contract),
	}

	return s
}

// runs in a goroutine until s is closed
func (s *OndoFinanceScraper) mainLoop() {

	time.Sleep(4 * time.Second)

	sink, subscription, err := s.getTradesChannel()
	if err != nil {
		log.Error("error fetching trades channel: ", err)
	}

	go func() {
		for {
			select {
			case rawTrade, ok := <-sink:
				if ok {
					trade, err := s.normalizeRawTrade(rawTrade)
					if err != nil {
						log.Error("normalizeRawTrade: ", err)
						continue
					}
					if trade.Price > 0 {
						log.Infof("Got trade on pair %s: %v", trade.Pair, trade)
						s.chanTrades <- &trade
					}
				}
			case err := <-subscription.Err():
				log.Error("subscription: ", err)
				log.Info("wait and resubscribe...")
				time.Sleep(1 * time.Second)
				subscription.Unsubscribe()
				time.Sleep(5 * time.Second)

				sink, subscription, err = s.getTradesChannel()
				if err != nil {
					log.Error("error fetching trades channel: ", err)
				}
			}
		}
	}()

}

func (s *OndoFinanceScraper) getTradesChannel() (tradesSink chan *ondo.OndotokenmanagerTradeExecuted, subscription event.Subscription, err error) {

	var contract *ondo.OndotokenmanagerFilterer
	contract, err = ondo.NewOndotokenmanagerFilterer(s.factoryContractAddress, s.WsClient)
	if err != nil {
		return
	}
	tradesSink = make(chan *ondo.OndotokenmanagerTradeExecuted)
	subscription, err = contract.WatchTradeExecuted(&bind.WatchOpts{}, tradesSink)

	return
}

// normalizeOndoFinanceTrade takes a raw trade as returned by the trade contract's channel and converts it to a OndoFinanceTrade type
func (s *OndoFinanceScraper) normalizeRawTrade(trade *ondo.OndotokenmanagerTradeExecuted) (normalizedTrade dia.Trade, err error) {

	quoteToken := dia.Asset{Address: trade.Asset.Hex(), Blockchain: dia.ETHEREUM}
	quoteToken, err = s.relDB.GetAsset(trade.Asset.Hex(), dia.ETHEREUM)
	if err != nil {
		log.Errorf("GetAsset by address %s: %v", trade.Asset.Hex(), err)
	}

	amount := new(big.Float).Quo(big.NewFloat(0).SetInt(trade.Quantity), new(big.Float).SetFloat64(math.Pow10(PRECISION)))
	volume, _ := amount.Float64()
	if trade.Side == 1 {
		volume *= -1
	}

	price, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(trade.Price), new(big.Float).SetFloat64(math.Pow10(PRECISION))).Float64()

	block, err := s.RestClient.HeaderByNumber(context.Background(), big.NewInt(int64(trade.Raw.BlockNumber)))
	if err != nil {
		log.Errorf("HeaderByNumber %v: %v", int64(trade.Raw.BlockNumber), err)
	}
	timestamp := time.Unix(int64(block.Time), 0)

	t := dia.Trade{
		Symbol:         quoteToken.Symbol,
		Pair:           quoteToken.Symbol + "-" + s.baseToken.Symbol,
		Price:          price,
		Volume:         volume,
		BaseToken:      s.baseToken,
		QuoteToken:     quoteToken,
		Time:           timestamp,
		ForeignTradeID: trade.Raw.TxHash.Hex(),
		Source:         s.exchangeName,
		VerifiedPair:   true,
	}

	return t, nil
}

// FetchAvailablePairs returns a list with all available trade pairs as dia.Pair for the pairDiscorvery service
func (s *OndoFinanceScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	return
}

func (s *OndoFinanceScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

func (s *OndoFinanceScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *OndoFinanceScraper) Close() error {

	if s.closed {
		return errors.New("OndoFinance: Already closed")
	}
	s.WsClient.Close()
	s.RestClient.Close()
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *OndoFinanceScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {

	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("OndoFinanceScraper: Call ScrapePair on closed scraper")
	}
	ps := &OndoFinancePairScraper{
		parent: s,
		pair:   pair,
	}
	s.pairScrapers[pair.ForeignName] = ps
	return ps, nil
}

func (s *OndoFinanceScraper) pingNode(pingNodeInterval int64) {
	ticker := time.NewTicker(time.Duration(pingNodeInterval) * time.Second)
	go func() {
		for range ticker.C {
			blockNumber, err := s.WsClient.BlockNumber(context.Background())
			if err != nil {
				log.Error("pingNode: ", err)
			} else {
				log.Infof("ping websocket: %v -- blockNumber: %d", time.Now(), blockNumber)
			}
		}
	}()
}

// OndoFinanceScraper implements PairScraper for OndoFinance
type OndoFinancePairScraper struct {
	parent *OndoFinanceScraper
	pair   dia.ExchangePair
	//closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *OndoFinancePairScraper) Close() error {
	return nil
}

// Channel returns a channel that can be used to receive trades
func (s *OndoFinanceScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *OndoFinancePairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *OndoFinancePairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
