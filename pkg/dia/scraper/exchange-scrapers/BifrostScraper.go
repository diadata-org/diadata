package scrapers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
)

const (
	Blockchain = "Polkadot"
)

type BifrostScraper struct {
	logger        *logrus.Entry
	pairScrapers  map[string]*BifrostPairScraper // pc.ExchangePair -> pairScraperSet
	shutdown      chan nothing
	shutdownDone  chan nothing
	errorLock     sync.RWMutex
	error         error
	closed        bool
	ticker        *time.Ticker
	chanTrades    chan *dia.Trade
	db            *models.RelDB
	api           string
	currentHeight string
	exchangeName  string
	blockchain    string
}

// NewBifrostScraper initializes a new BifrostScraper.
func NewBifrostScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *BifrostScraper {
	refreshDelay := utils.GetTimeDurationFromIntAsMilliseconds(
		utils.GetenvInt(strings.ToUpper(exchange.Name)+"_REFRESH_DELAY", 10000),
	)

	apiURL := "http://localhost:3000/bifrost/v1"

	s := &BifrostScraper{
		shutdown:      make(chan nothing),
		shutdownDone:  make(chan nothing),
		ticker:        time.NewTicker(refreshDelay),
		chanTrades:    make(chan *dia.Trade),
		db:            relDB,
		api:           apiURL,
		currentHeight: "0x5d184786631f977395636a98b231358b9a82b63c235c801c170fb7e532cb839b",
		exchangeName:  exchange.Name,
		blockchain:    Blockchain,
	}

	s.logger = logrus.
		New().
		WithContext(context.Background()).
		WithField("context", "BifrostScraper")

	s.logger.Info("Initialized BifrostScraper")

	if scrape {
		go s.mainLoop()
	}
	return s
}

// mainLoop continuously fetches swap events from Bifrost API and processes them.
func (s *BifrostScraper) mainLoop() {
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

// Update fetches the latest swap events from Bifrost API and processes them into dia.Trade objects.
func (s *BifrostScraper) Update() error {
	endpoint := fmt.Sprintf("%s/events/stable/swap/%s", s.api, s.currentHeight)

	resp, err := s.fetchWithRetry(endpoint, "application/json", 3)
	if err != nil {
		return fmt.Errorf("failed to fetch swap events after retries: %v", err)
	}

	s.logger.WithField("status", resp.Status).Info("Fetched swap events")
	defer resp.Body.Close()

	var events []StableSwapEvent
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return fmt.Errorf("failed to decode swap events: %v", err)
	}

	pools, err := s.getPools()
	if err != nil {
		s.logger.WithError(err).Error("failed to GetAllPoolsExchange")
		return err
	}

	s.logger.WithField("pools", pools).Info("pools")

	for _, pool := range pools {
		if len(pool.Assetvolumes) != 2 {
			s.logger.WithField("poolAddress", pool.Address).Error("pool is missing required asset volumes")
			continue
		}

		poolEvents := make([]StableSwapEvent, 0)

		for _, event := range events {
			tokenAName := strings.ToLower(event.InputAsset)
			tokenBName := strings.ToLower(event.OutputAsset)
			eventAddressA := fmt.Sprintf("%s:%s:%s:%s", s.blockchain, s.exchangeName, tokenAName, tokenBName)
			eventAddressB := fmt.Sprintf("%s:%s:%s:%s", s.blockchain, s.exchangeName, tokenBName, tokenAName)

			
			if eventAddressA == pool.Address {
				s.logger.WithField("PoolAddress:   ", pool.Address).Info("PoolAddress")
				s.logger.WithField("EventAddressA: ", eventAddressA).Info("EventAddress")
				poolEvents = append(poolEvents, event)
			}
			if eventAddressB == pool.Address {
				s.logger.WithField("PoolAddress:   ", pool.Address).Info("PoolAddress")
				s.logger.WithField("EventAddressB: ", eventAddressB).Info("EventAddress")
			}
		}

		if len(poolEvents) == 0 {
			s.logger.WithField("poolAddress", pool.Address).Warn("no events match for pool")
			continue
		}

		for _, event := range poolEvents {
			s.logger.WithField("event", event).WithField("currentHeight", s.currentHeight).Info("event")
			timestamp := time.Now()

			diaTrade := s.handleTrade(pool, event, timestamp)
			s.logger.WithField("diaTrade", diaTrade).Info("diaTrade")

			s.chanTrades <- diaTrade
		}
	}

	return nil
}

// fetchWithRetry fetches data from the API with retry logic
func (s *BifrostScraper) fetchWithRetry(endpoint string, contentType string, retries int) (*http.Response, error) {
	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	var resp *http.Response
	var err error
	for i := 0; i < retries; i++ {
		s.logger.WithField("attempt", i+1).Info("Fetching data from API")
		s.logger.WithField("endpoint", endpoint).Info("Requesting data")

		req, err := http.NewRequest("GET", endpoint, bytes.NewBuffer([]byte("")))
		if err != nil {
			s.logger.WithError(err).Error("Error creating request")
			return nil, err
		}
		req.Header.Set("Content-Type", contentType)

		resp, err = client.Do(req)
		if err == nil && resp != nil {
			if resp.StatusCode >= 200 && resp.StatusCode < 300 {
				return resp, nil
			}
			resp.Body.Close()

			if resp.StatusCode >= 500 && resp.StatusCode < 600 {
				s.logger.WithField("status", resp.StatusCode).Warn("Server error. Retrying...")
			} else {
				// Client error or other non-retryable response
				return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
			}

		} else if err != nil {
			s.logger.WithError(err).Warn("Failed to fetch data from API. Retrying...")
		}

		time.Sleep(time.Second * 2)
	}
	return nil, err
}

// handleTrade processes a BifrostSwapEvent into a dia.Trade object.
func (s *BifrostScraper) handleTrade(pool dia.Pool, event StableSwapEvent, time time.Time) *dia.Trade {
	var volume, price float64

	decimals0 := int64(pool.Assetvolumes[0].Asset.Decimals)
	s.logger.WithField("decimals0", decimals0).Info("decimals0")

	decimals1 := int64(pool.Assetvolumes[1].Asset.Decimals)
	s.logger.WithField("decimals1", decimals1).Info("decimals1")

	amount0, _ := utils.StringToFloat64(event.InputAmount, decimals0)
	s.logger.WithField("amount0", amount0).Info("amount0")

	amount1, _ := utils.StringToFloat64(event.OutputAmount, decimals1)
	s.logger.WithField("amount1", amount1).Info("amount1")

	volume = -amount0
	s.logger.WithField("volume", volume).Info("volume")

	price = amount0 / amount1
	s.logger.WithField("price", price).Info("price")

	symbolPair := fmt.Sprintf("%s-%s", pool.Assetvolumes[0].Asset.Symbol, pool.Assetvolumes[1].Asset.Symbol)
	s.logger.WithField("symbolPair", symbolPair).Info("symbolPair")

	diaTrade := &dia.Trade{
		Time:           time,
		Symbol:         pool.Assetvolumes[0].Asset.Symbol,
		Pair:           symbolPair,
		ForeignTradeID: event.TxID,
		Source:         s.exchangeName,
		Price:          price,
		Volume:         volume,
		VerifiedPair:   true,
		QuoteToken:     pool.Assetvolumes[0].Asset,
		BaseToken:      pool.Assetvolumes[1].Asset,
		PoolAddress:    pool.Address,
	}

	s.logger.WithField("diaTrade", diaTrade).Info("diaTrade")

	return diaTrade
}

// FetchAvailablePairs returns a list with all trading pairs available on
// the exchange associated to the APIScraper. The format is such that it can
// be used by the corr. pairScraper in order to fetch trades.
func (s *BifrostScraper) FetchAvailablePairs() ([]dia.ExchangePair, error) {
	return []dia.ExchangePair{}, nil
}

func (s *BifrostScraper) getPools() ([]dia.Pool, error) {
	return s.db.GetAllPoolsExchange(s.exchangeName, 0)
}

func (s *BifrostScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

func (s *BifrostScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

func (s *BifrostScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("AlephiumScraper: Call ScrapePair on closed scraper")
	}
	ps := &BifrostPairScraper{
		parent:     s,
		pair:       pair,
		lastRecord: 0,
	}

	s.pairScrapers[pair.Symbol] = ps

	return ps, nil
}

// cleanup handles the shutdown procedure.
func (s *BifrostScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()

	s.ticker.Stop()

	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.shutdownDone)
}

// Close gracefully shuts down the BifrostScraper.
func (s *BifrostScraper) Close() error {
	if s.closed {
		return errors.New("BifrostScraper: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Channel returns the channel used to receive trades/pricing information.
func (s *BifrostScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

type BifrostPairScraper struct {
	parent     *BifrostScraper
	pair       dia.ExchangePair
	closed     bool
	lastRecord int64
}

func (ps *BifrostPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}

func (ps *BifrostPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *BifrostPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

type StableSwapEvent struct {
	TxID            string   `json:"txId"`
	Timestamp       int64    `json:"timestamp"`
	BlockHash       string   `json:"blockHash"`
	Swapper         string   `json:"swapper"`
	PoolID          string   `json:"poolId"`
	A               string   `json:"a"`
	InputAsset      string   `json:"inputAsset"`
	OutputAsset     string   `json:"outputAsset"`
	InputAmount     string   `json:"inputAmount"`
	MinOutputAmount string   `json:"minOutputAmount"`
	Balances        []string `json:"balances"`
	TotalSupply     string   `json:"totalSupply"`
	OutputAmount    string   `json:"outputAmount"`
}
