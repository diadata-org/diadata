package scrapers

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/config"
	"github.com/centrifuge/go-substrate-rpc-client/v4/rpc/state"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/acalaswap"
	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/dia"
)

type AcalaswapToken struct {
	Address  string
	Symbol   string
	Decimals uint8
	Name     string
}

type AcalaswapPair struct {
	Token0      AcalaswapToken
	Token1      AcalaswapToken
	ForeignName string
	Address     string
}

type AcalaswapSwap struct {
	ID        string
	Timestamp int64
	Pair      *AcalaswapPair
	Amount0   float64
	Amount1   float64
}

type AcalaswapScraper struct {
	WsClient   *gsrpc.SubstrateAPI
	RestClient *gsrpc.SubstrateAPI
	Metadata   *types.Metadata
	// signaling channels for session initialization and finishing
	//initDone     chan nothing
	run          bool
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	pairScrapers map[string]*AcalaswapPairScraper
	pairRecieved chan *AcalaswapPair

	exchangeName    string
	startBlock      uint64
	waitTime        int
	listenByAddress bool
	chanTrades      chan *dia.Trade
	resubscribe     chan nothing
}

type AcalaConfig struct {
	restURL                string
	wsURL                  string
	contractAddress        string
	contratDeployedAtBlock int64
}

const (
	acalaHTTP  = "https://acala-polkadot.api.onfinality.io/public-rpc"
	karuraHTTP = "https://karura.api.onfinality.io/public-rpc"
	acalaWSS   = "wss://acala-rpc-2.aca-api.network/ws"
	karuraWSS  = "wss://karura-rpc-2.aca-api.network/ws"
)

// NewAcalaswapScraper returns a new AcalaswapScraper
func NewAcalaswapScraper(exchange dia.Exchange) *AcalaswapScraper {
	log.Info("NewAcalaswapScraper ", exchange.Name)
	log.Info("NewAcalaswapScraper Address ", exchange.Contract)

	var s *AcalaswapScraper
	switch exchange.Name {
	case dia.AcalaswapExchange:
		s = makeAcalaswapScraper(exchange, false, acalaHTTP, acalaWSS, "200", uint64(1102858))
	case dia.AcalaswapExchangeKarura:
		s = makeAcalaswapScraper(exchange, false, karuraHTTP, karuraWSS, "200", uint64(1826919))
	}

	go s.mainLoop()
	return s
}

// makeAcalaswapScraper returns a uniswap scraper as used in NewAcalaswapScraper.
func makeAcalaswapScraper(exchange dia.Exchange, listenByAddress bool, restDial string, wsDial string, waitMilliseconds string, startBlock uint64) *AcalaswapScraper {
	var restClient, wsClient *gsrpc.SubstrateAPI
	var err error
	var s *AcalaswapScraper

	log.Infof("Init rest and ws client for %s.", exchange.BlockChain.Name)
	// Create our API with a default connection to the local node
	api, err := gsrpc.NewSubstrateAPI(config.Default().RPCURL)
	if err != nil {
		panic(err)
	}

	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		panic(err)
	}
	restClient, err = gsrpc.NewSubstrateAPI((utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", restDial)))
	if err != nil {
		log.Fatal("init rest client: ", err)
	}
	wsClient, err = gsrpc.NewSubstrateAPI((utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_WS", restDial)))
	if err != nil {
		log.Fatal("init ws client: ", err)
	}

	var waitTime int
	waitTimeString := utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_WAIT_TIME", waitMilliseconds)
	waitTime, err = strconv.Atoi(waitTimeString)
	if err != nil {
		log.Error("could not parse wait time: ", err)
		waitTime = 500
	}

	s = &AcalaswapScraper{
		WsClient:        wsClient,
		RestClient:      restClient,
		Metadata:        meta,
		shutdown:        make(chan nothing),
		shutdownDone:    make(chan nothing),
		pairScrapers:    make(map[string]*AcalaswapPairScraper),
		exchangeName:    exchange.Name,
		pairRecieved:    make(chan *AcalaswapPair),
		error:           nil,
		chanTrades:      make(chan *dia.Trade),
		resubscribe:     make(chan nothing),
		waitTime:        waitTime,
		listenByAddress: listenByAddress,
		startBlock:      startBlock,
	}
	return s
}

// runs in a goroutine until s is closed
func (s *AcalaswapScraper) mainLoop() {
	time.Sleep(4 * time.Second)
	s.run = true

	err := s.subscribeToSystemEvents()
	if err != nil {
		log.Error("error fetching swaps channel: ", err)
	}

	go func() {
		for s.run {
			select {
			case <-s.resubscribe:
				if s.run {
					log.Info("resubscribe to system events")
					err = s.subscribeToSystemEvents()
					if err != nil {
						log.Error(err)
					}

				}
			}
		}
	}()

	time.Sleep(10 * time.Second)

	if s.error == nil {
		s.error = errors.New("main loop terminated by Close()")
	}
	s.cleanup(nil)
}

func (s *AcalaswapScraper) subscribeToSystemEvents() error {
	sub, key, err := s.GetEventsChannel()
	if err != nil {
		return err
	}
	go func() {
		// outer for loop for subscription notifications
		fmt.Println("subscribed to NewPools")
		defer fmt.Println("Unsubscribed to NewPools")
		defer sub.Unsubscribe()
		subscribed := true
		for s.run && subscribed {
			select {
			case err := <-sub.Err():
				if err != nil {
					log.Error(err)
				}
				subscribed = false
				if s.run {
					s.resubscribe <- nothing{}
				}

			case set := <-sub.Chan():
				// inner loop for the changes within one of those notifications
				for _, chng := range set.Changes {
					if !types.Eq(chng.StorageKey, key) || !chng.HasStorageData {
						// skip, we are only interested in events with content
						continue
					}

					// Decode the event records
					events := acalaswap.Events{}
					err = types.EventRecordsRaw(chng.StorageData).DecodeEventRecords(s.Metadata, &events)
					if err != nil {
						log.Error(err)
						continue
					}

					// Show what we are busy with
					for _, e := range events.Dex_Swapped {
						log.Debugf("\tDex:Swapped:: (phase=%#v)\n", e.Phase)
						log.Debugf("\t\t%#x, %v\n", e.Trader, e.Path)
						// TODO: maybe need more checks here.
						if e.Phase.IsApplyExtrinsic {
							s.processSwap(e)
						}
					}
				}
			}
		}
	}()
	return nil
}

func (s *AcalaswapScraper) processSwap(e acalaswap.EventDexSwapped) error {
	swap, err := s.normalizeAcalaswapSwap(e)
	if err != nil {
		log.Error("error normalizing swap: ", err)
		return err
	}
	pair := swap.Pair
	price, volume := s.getSwapData(swap)
	token0 := dia.Asset{
		Address:    pair.Token0.Address,
		Symbol:     pair.Token0.Symbol,
		Name:       pair.Token0.Name,
		Decimals:   pair.Token0.Decimals,
		Blockchain: Exchanges[s.exchangeName].BlockChain.Name,
	}
	token1 := dia.Asset{
		Address:    pair.Token1.Address,
		Symbol:     pair.Token1.Symbol,
		Name:       pair.Token1.Name,
		Decimals:   pair.Token1.Decimals,
		Blockchain: Exchanges[s.exchangeName].BlockChain.Name,
	}

	t := &dia.Trade{
		Symbol:         pair.Token0.Symbol,
		Pair:           pair.ForeignName,
		Price:          price,
		Volume:         volume,
		BaseToken:      token1,
		QuoteToken:     token0,
		Time:           time.Unix(swap.Timestamp, 0),
		ForeignTradeID: swap.ID,
		Source:         s.exchangeName,
		VerifiedPair:   true,
	}

	switch {
	case utils.Contains(reverseBasetokens, pair.Token1.Address):
		// If we need quotation of a base token, reverse pair
		tSwapped, err := dia.SwapTrade(*t)
		if err == nil {
			t = &tSwapped
		}
	case utils.Contains(reverseQuotetokens, pair.Token0.Address):
		// If we need quotation of a base token, reverse pair
		tSwapped, err := dia.SwapTrade(*t)
		if err == nil {
			t = &tSwapped
		}
	}
	if price > 0 {
		log.Info("Got trade: ", t)
		s.chanTrades <- t
	}
	return nil
}

// GetEventsChannel returns a channel for events
func (s *AcalaswapScraper) GetEventsChannel() (*state.StorageSubscription, types.StorageKey, error) {
	// Subscribe to system events via storage
	key, err := types.CreateStorageKey(s.Metadata, "System", "Events", nil)
	if err != nil {
		return nil, nil, err
	}

	sub, err := s.RestClient.RPC.State.SubscribeStorageRaw([]types.StorageKey{key})
	if err != nil {
		return nil, nil, err
	}

	return sub, key, nil

}

func (s *AcalaswapScraper) getSwapData(swap AcalaswapSwap) (price float64, volume float64) {
	volume = swap.Amount0
	price = math.Abs(swap.Amount1 / swap.Amount0)
	return
}

// normalizeAcalaswapSwap takes a swap as returned by the swap contract's channel and converts it to a AcalaswapSwap type
func (s *AcalaswapScraper) normalizeAcalaswapSwap(swap acalaswap.EventDexSwapped) (normalizedSwap AcalaswapSwap, err error) {

	pair, err := s.GetPairByToken(swap.Path[0], swap.Path[len(swap.Path)-1])
	if err != nil {
		log.Error("error getting pair by address: ", err)
		return
	}
	decimals0 := int(pair.Token0.Decimals)
	decimals1 := int(pair.Token1.Decimals)
	amount0, _ := new(big.Float).Quo(big.NewFloat(0).SetInt((swap.LiquidityChanges[0]).Int), new(big.Float).SetFloat64(math.Pow10(decimals0))).Float64()
	amount1, _ := new(big.Float).Quo(big.NewFloat(0).SetInt((swap.LiquidityChanges[1]).Int), new(big.Float).SetFloat64(math.Pow10(decimals1))).Float64()

	normalizedSwap = AcalaswapSwap{
		ID:        fmt.Sprintf("%d", swap.Phase.AsApplyExtrinsic),
		Timestamp: time.Now().Unix(),
		Pair:      pair,
		Amount0:   amount0,
		Amount1:   amount1,
	}
	return
}

// FetchAvailablePairs returns a list with all available trade pairs as dia.Pair for the pairDiscorvery service
func (s *AcalaswapScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	return
}

func (s *AcalaswapScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

func (s *AcalaswapScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	if pair.ForeignName == "WETH" {
		pair.Symbol = "ETH"
	}
	return pair, nil
}

func (s *AcalaswapScraper) GetPairByToken(currencyIdIn acalaswap.CurrencyID, currencyIdOut acalaswap.CurrencyID) (pair *AcalaswapPair, err error) {
	argIn, err := types.Encode(currencyIdIn)
	if err != nil {
		return nil, err
	}
	keyIn, err := types.CreateStorageKey(s.Metadata, "AssetRegistry", "AssetMetadatas", argIn)
	if err != nil {
		return nil, err
	}

	resultIn, err := s.RestClient.RPC.State.QueryStorageAtLatest([]types.StorageKey{keyIn})
	if err != nil {
		return nil, err
	}

	// Decode the event records
	assetMetadataIn := acalaswap.AcalaAssetMetadata{}
	err = types.EventRecordsRaw(resultIn[0].Changes[0].StorageData).DecodeEventRecords(s.Metadata, &assetMetadataIn)
	if err != nil {
		return nil, err
	}

	argOut, err := types.Encode(currencyIdIn)
	if err != nil {
		return nil, err
	}
	keyOut, err := types.CreateStorageKey(s.Metadata, "AssetRegistry", "AssetMetadatas", argOut)
	if err != nil {
		return nil, err
	}

	resultOut, err := s.RestClient.RPC.State.QueryStorageAtLatest([]types.StorageKey{keyOut})
	if err != nil {
		return nil, err
	}

	// Decode the event records
	assetMetadataOut := acalaswap.AcalaAssetMetadata{}
	err = types.EventRecordsRaw(resultOut[0].Changes[0].StorageData).DecodeEventRecords(s.Metadata, &assetMetadataOut)
	if err != nil {
		return nil, err
	}
	token0 := AcalaswapToken{
		Address:  string(assetMetadataIn.Name),
		Symbol:   string(assetMetadataIn.Symbol),
		Decimals: uint8(assetMetadataIn.Decimals),
	}
	token1 := AcalaswapToken{
		Address:  string(assetMetadataOut.Name),
		Symbol:   string(assetMetadataOut.Symbol),
		Decimals: uint8(assetMetadataOut.Decimals),
	}
	foreignName := token0.Symbol + "-" + token1.Symbol
	pair = &AcalaswapPair{
		ForeignName: foreignName,
		Token0:      token0,
		Token1:      token1,
	}
	return pair, nil
}

func (s *AcalaswapScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()

	if err != nil {
		s.error = err
	}
	s.closed = true

	close(s.shutdownDone)
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *AcalaswapScraper) Close() error {

	if s.closed {
		return errors.New("AcalaswapScraper: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// AcalaswapPairScraper implements PairScraper for Acalaswap
type AcalaswapPairScraper struct {
	parent *UniswapScraper
	pair   dia.ExchangePair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *AcalaswapPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Channel returns a channel that can be used to receive trades
func (ps *AcalaswapScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *AcalaswapPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *AcalaswapPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
