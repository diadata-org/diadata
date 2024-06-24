package scrapers

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	alephiumhelper "github.com/diadata-org/diadata/pkg/dia/helpers/alephium-helper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
)

type AyinScraper struct {
	logger *logrus.Entry
	// signaling channels
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock                 sync.RWMutex
	error                     error
	closed                    bool
	pairScrapers              map[string]*AyinPairScraper // pc.ExchangePair -> pairScraperSet
	api                       *alephiumhelper.AlephiumClient
	ticker                    *time.Ticker
	exchangeName              string
	blockchain                string
	currentHeight             int
	eventsLimit               int
	swapContractsLimit        int
	targetSwapContract        string
	chanTrades                chan *dia.Trade
	db                        *models.RelDB
	refreshDelay              time.Duration
	sleepBetweenContractCalls time.Duration
	reverseBasetokens         []string
	reverseQuotetokens        []string
}

// NewAyinScraper returns a new AyinScraper initialized with default values.
// The instance is asynchronously scraping as soon as it is created.
// ENV values:
//
//		AYIN_REFRESH_DELAY - (optional,millisecond) refresh data after each poll, default "alephiumhelper.DefaultRefreshDelay" value
//	 	AYIN_SLEEP_TIMEOUT - (optional,millisecond), make timeout between API calls, default "alephiumhelper.DefaultSleepBetweenContractCalls" value
//		AYIN_SWAP_CONTRACTS_LIMIT - (optional, int), limit to get swap contact addresses, default "alephiumhelper.DefaultSwapContractsLimit" value
//		AYIN_TARGET_SWAP_CONTRACT - (optional, string), useful for debug, default = ""
//		AYIN_DEBUG - (optional, bool), make stdout output with alephium client http call, default = false
func NewAyinScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *AyinScraper {
	ayinRefreshDelay := utils.GetTimeDurationFromIntAsMilliseconds(
		utils.GetenvInt(strings.ToUpper(exchange.Name)+"_REFRESH_DELAY", alephiumhelper.DefaultRefreshDelay),
	)
	sleepBetweenContractCalls := utils.GetTimeDurationFromIntAsMilliseconds(
		utils.GetenvInt(strings.ToUpper(exchange.Name)+"_SLEEP_TIMEOUT", alephiumhelper.DefaultSleepBetweenContractCalls),
	)
	isDebug := utils.GetenvBool(strings.ToUpper(exchange.Name)+"_DEBUG", false)
	eventsLimit := utils.GetenvInt(strings.ToUpper(exchange.Name)+"_REFRESH_DELAY", alephiumhelper.DefaultEventsLimit)
	swapContractsLimit := utils.GetenvInt(strings.ToUpper(exchange.Name)+"_SWAP_CONTRACTS_LIMIT", alephiumhelper.DefaultSwapContractsLimit)
	targetSwapContract := utils.Getenv(strings.ToUpper(exchange.Name)+"_TARGET_SWAP_CONTRACT", "")

	alephiumClient := alephiumhelper.NewAlephiumClient(
		log.WithContext(context.Background()).WithField("context", "AlephiumClient"),
		sleepBetweenContractCalls,
		isDebug,
	)
	s := &AyinScraper{
		shutdown:                  make(chan nothing),
		shutdownDone:              make(chan nothing),
		pairScrapers:              make(map[string]*AyinPairScraper),
		api:                       alephiumClient,
		ticker:                    time.NewTicker(ayinRefreshDelay),
		exchangeName:              exchange.Name,
		blockchain:                exchange.BlockChain.Name,
		currentHeight:             0,
		error:                     nil,
		chanTrades:                make(chan *dia.Trade),
		db:                        relDB,
		refreshDelay:              ayinRefreshDelay,
		sleepBetweenContractCalls: sleepBetweenContractCalls,
		eventsLimit:               eventsLimit,
		swapContractsLimit:        swapContractsLimit,
		targetSwapContract:        targetSwapContract,
	}

	// Import tokens which appear as base token and we need a quotation for
	var err error
	reverseBasetokens, err = getReverseTokensFromConfig("ayin/reverse_tokens/" + s.exchangeName + "Basetoken")
	if err != nil {
		log.Error("error getting tokens for which pairs should be reversed: ", err)
	}
	log.Info("reverse basetokens: ", reverseBasetokens)
	reverseQuotetokens, err = getReverseTokensFromConfig("ayin/reverse_tokens/" + s.exchangeName + "Quotetoken")
	if err != nil {
		log.Error("error getting tokens for which pairs should be reversed: ", err)
	}
	log.Info("reverse quotetokens: ", reverseQuotetokens)
	s.reverseBasetokens = *reverseBasetokens
	s.reverseQuotetokens = *reverseQuotetokens

	s.logger = logrus.
		New().
		WithContext(context.Background()).
		WithField("context", "AyinDEXScraper")

	if scrape {
		go s.mainLoop()
	}
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (s *AyinScraper) mainLoop() {
	currentHeight, err := s.api.GetCurrentHeight()
	if err != nil {
		s.logger.WithError(err).Error("failed to GetCurrentHeight")
		s.cleanup(err)
		return
	}
	s.currentHeight = currentHeight

	err = s.Update()
	if err != nil {
		s.logger.Error(err)
	}
	for {
		select {
		case <-s.ticker.C:
			err := s.Update()
			if err != nil {
				s.logger.Error(err)
			}
		case <-s.shutdown: // user requested shutdown
			s.logger.Println("shutting down")
			s.cleanup(nil)
			return
		}
	}
}

func (s *AyinScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

func (s *AyinScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

func (s *AyinScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("AlephiumScraper: Call ScrapePair on closed scraper")
	}
	ps := &AyinPairScraper{
		parent:     s,
		pair:       pair,
		lastRecord: 0,
	}

	s.pairScrapers[pair.Symbol] = ps

	return ps, nil
}

func (s *AyinScraper) getPools() ([]dia.Pool, error) {
	if s.targetSwapContract != "" {
		result := make([]dia.Pool, 1)
		pool, err := s.db.GetPoolByAddress(s.blockchain, s.targetSwapContract)
		result[0] = pool
		return result, err
	}
	return s.db.GetAllPoolsExchange(s.exchangeName, 0)
}

func (s *AyinScraper) Update() error {
	logger := s.logger.WithFields(logrus.Fields{
		"function": "Update",
	})

	blockHashes, err := s.api.GetBlockHashes(s.currentHeight)
	if err != nil {
		s.logger.WithError(err).Error("failed to GetBlockHashes")
		return err
	}
	if len(blockHashes) == 0 {
		// logger.Info("no new blocks in the network, waiting...")
		return nil
	}

	allEvents, err := s.fetchEvents(blockHashes)
	if err != nil {
		logger.WithError(err).Error("failed to fetch events")
		return err
	}
	s.currentHeight += 1

	if len(allEvents) == 0 {
		logger.WithField("currentHeight", s.currentHeight).Info("no events, skipping to the next block...")
		return nil
	}

	pools, err := s.getPools()
	if err != nil {
		logger.
			WithError(err).
			Error("failed to GetAllPoolsExchange")
		return err
	}

	for _, pool := range pools {
		if len(pool.Assetvolumes) != 2 {
			logger.WithField("poolAddress", pool.Address).Error("pool is missing required asset volumes")
			continue
		}

		poolEvents := make([]alephiumhelper.ContractEvent, 0)

		for _, event := range allEvents {
			if event.ContractAddress == pool.Address {
				poolEvents = append(poolEvents, event)
			}
		}

		if len(poolEvents) == 0 {
			continue
		}

		for _, event := range poolEvents {
			logger.WithField("event", event).WithField("currentHeight", s.currentHeight).Info("event")
			transactionDetails, err := s.api.GetTransactionDetails(event.TxID)
			if err != nil {
				logger.
					WithError(err).
					Error("failed to GetTransactionDetails")
				continue
			}

			var timestamp time.Time
			if transactionDetails.Timestamp > 0 {
				timestamp = time.UnixMilli(transactionDetails.Timestamp)
			} else {
				timestamp = time.Now()
			}

			diaTrade := s.handleTrade(&pool, &event, timestamp)
			logger.
				WithField("parentAddress", pool.Address).
				WithField("diaTrade", diaTrade).
				Info("diaTrade")
			s.chanTrades <- diaTrade
		}
	}

	return nil
}

func (s *AyinScraper) fetchEvents(blockHashes []string) ([]alephiumhelper.ContractEvent, error) {
	allEvents := make([]alephiumhelper.ContractEvent, 0)

	for _, hash := range blockHashes {
		events, err := s.api.GetBlockEvents(hash)
		if err != nil {
			return allEvents, err
		}

		filtered := s.api.FilterEvents(events, alephiumhelper.SwapEventIndex)
		allEvents = append(allEvents, filtered...)
	}

	return allEvents, nil
}

func (s *AyinScraper) handleTrade(pool *dia.Pool, event *alephiumhelper.ContractEvent, time time.Time) *dia.Trade {
	var volume, price float64

	decimals0 := int64(pool.Assetvolumes[0].Asset.Decimals)
	decimals1 := int64(pool.Assetvolumes[1].Asset.Decimals)

	if event.Fields[1].Value != "0" {
		// if we are swapping from ALPH(asset0) to USDT(asset1), - default behaviour
		//	then amount0In ((fields[1]) will be the amount for ALPH
		//	and amount1Out (fields[4]) will be the amount for USDT.
		amount0In, _ := utils.StringToFloat64(event.Fields[1].Value, decimals0)
		amount1Out, _ := utils.StringToFloat64(event.Fields[4].Value, decimals1)
		volume = amount0In
		price = amount1Out / amount0In
	} else {
		// If we are swapping from USDT(asset1) to ALPH(asset0),
		//	then amount1In ((fields[2]) will be the amount for USDT
		//  and amount0Out (fields[3]) will be the amount for ALPH.
		amount1In, _ := utils.StringToFloat64(event.Fields[2].Value, decimals1)
		amount0Out, _ := utils.StringToFloat64(event.Fields[3].Value, decimals0)
		volume = -amount0Out
		price = amount1In / amount0Out
	}

	symbolPair := fmt.Sprintf("%s-%s", pool.Assetvolumes[0].Asset.Symbol, pool.Assetvolumes[1].Asset.Symbol)

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

	// Reverse the order of the trade for selected assets.
	switch {
	case utils.Contains(reverseQuotetokens, diaTrade.QuoteToken.Address):
		// If we don't need quotation of quote token, reverse pair.
		tSwapped, err := dia.SwapTrade(*diaTrade)
		if err == nil {
			diaTrade = &tSwapped
		}
	case utils.Contains(reverseBasetokens, diaTrade.BaseToken.Address):
		// If we need quotation of a base token, reverse pair
		tSwapped, err := dia.SwapTrade(*diaTrade)
		if err == nil {
			diaTrade = &tSwapped
		}
	}

	return diaTrade
}

// closes all connected PairScrapers
// must only be called from mainLoop
func (s *AyinScraper) cleanup(err error) {

	s.errorLock.Lock()
	defer s.errorLock.Unlock()

	s.ticker.Stop()

	if err != nil {
		s.error = err
	}
	s.closed = true

	close(s.shutdownDone) // signal that shutdown is complete
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *AyinScraper) Close() error {
	if s.closed {
		return errors.New("AlephiumScraper: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Channel returns a channel that can be used to receive trades/pricing information
func (s *AyinScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

// FetchAvailablePairs returns a list with all available trade pairs as dia.ExchangePair for the pairDiscorvery service
func (s *AyinScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	logger := s.logger.WithFields(logrus.Fields{
		"function": "FetchAvailablePairs",
	})
	contractAddresses, err := s.api.GetSwapPairsContractAddresses(s.swapContractsLimit)
	if err != nil {
		logger.WithError(err).Error("failed to get swap contract addresses")
		return
	}
	for _, contractAddress := range contractAddresses.SubContracts {
		tokenPairs, err := s.api.GetTokenPairAddresses(contractAddress)

		if err != nil {
			logger.
				WithField("contractAddress", contractAddress).
				WithError(err).
				Error("failed to get GetTokenPairAddresses")
			continue
		}

		token0, err := s.api.GetTokenInfoForContractDecoded(tokenPairs[0], s.blockchain)
		if err != nil {
			logger.
				WithField("contractAddress", contractAddress).
				WithError(err).
				Error("failed to get GetTokenInfoForContractDecoded for token0")
			continue
		}

		token1, err := s.api.GetTokenInfoForContractDecoded(tokenPairs[1], s.blockchain)
		if err != nil {
			logger.
				WithField("contractAddress", contractAddress).
				WithError(err).
				Error("failed to get GetTokenInfoForContractDecoded for token1")
			continue
		}
		pair := dia.ExchangePair{
			Symbol:      token0.Symbol,
			ForeignName: fmt.Sprintf("%s-%s", token0.Symbol, token1.Symbol),
			Exchange:    s.exchangeName,
		}

		pairs = append(pairs, pair)

		time.Sleep(s.sleepBetweenContractCalls)
	}
	return pairs, nil
}

type AyinPairScraper struct {
	parent     *AyinScraper
	pair       dia.ExchangePair
	closed     bool
	lastRecord int64
}

func (ps *AyinPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}

func (ps *AyinPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *AyinPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}
