package scrapers

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/rpc/state"
	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/acalaswap"
	"github.com/diadata-org/diadata/pkg/utils"
	scalecodec "github.com/itering/scale.go"
	"github.com/itering/substrate-api-rpc"
	"github.com/itering/substrate-api-rpc/metadata"
)

type AcalaswapToken struct {
	Symbol   string
	Decimals uint8
	Name     string
	Address  string
}

type AcalaswapPair struct {
	Token0      AcalaswapToken
	Token1      AcalaswapToken
	ForeignName string
}

type AcalaswapSwap struct {
	Timestamp int64
	Pair      *AcalaswapPair
	Amount0   float64
	Amount1   float64
}

type AcalaswapScraper struct {
	WsClient    *gsrpc.SubstrateAPI
	RestClient  *gsrpc.SubstrateAPI
	Metadata2   *metadata.Instant
	Metadata    *types.Metadata
	SpecVersion int
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
	karuraWSS  = "wss://karura-rpc-3.aca-api.network/ws"
)

// NewAcalaswapScraper returns a new AcalaswapScraper
func NewAcalaswapScraper(exchange dia.Exchange, scrape bool) *AcalaswapScraper {
	log.Info("NewAcalaswapScraper ", exchange.Name)
	log.Info("NewAcalaswapScraper Address ", exchange.Contract)

	var s *AcalaswapScraper
	switch exchange.Name {
	case dia.AcalaswapExchange:
		s = makeAcalaswapScraper(exchange, false, acalaHTTP, acalaWSS, "200", uint64(1102858))
	case dia.AcalaswapExchangeKarura:
		s = makeAcalaswapScraper(exchange, false, karuraHTTP, karuraWSS, "200", uint64(1826919))
	}
	if scrape {
		go s.mainLoop()
	}
	return s
}

// makeAcalaswapScraper returns a uniswap scraper as used in NewAcalaswapScraper.
func makeAcalaswapScraper(exchange dia.Exchange, listenByAddress bool, restDial string, wsDial string, waitMilliseconds string, startBlock uint64) *AcalaswapScraper {
	var restClient, wsClient *gsrpc.SubstrateAPI
	var err error
	var s *AcalaswapScraper

	log.Infof("Init rest and ws client for %s.", exchange.BlockChain.Name)

	restClient, err = gsrpc.NewSubstrateAPI((utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", restDial)))
	if err != nil {
		log.Fatal("init rest client: ", err)
	}
	wsClient, err = gsrpc.NewSubstrateAPI((utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_WS", wsDial)))
	if err != nil {
		log.Fatal("init ws client: ", err)
	}

	runtimeVersionLatest, err := restClient.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		log.Fatal("get the substrate runtime version: ", err)
	}
	var res string
	err = wsClient.Client.Call(&res, "state_getMetadata")
	if err != nil {
		log.Fatal("get the substrate raw metadata: ", err)
	}
	rr := metadata.RuntimeRaw{Spec: int(runtimeVersionLatest.SpecVersion), Raw: res}
	metadata2 := metadata.Process(&rr)
	metadata, err := wsClient.RPC.State.GetMetadataLatest()
	if err != nil {
		log.Fatal("get the substrate metadata: ", err)
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
		Metadata2:       metadata2,
		Metadata:        metadata,
		SpecVersion:     int(runtimeVersionLatest.SpecVersion),
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
		default:
		}
	}

	time.Sleep(10 * time.Second)
	log.Info("####################")
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
		defer log.Info("Unsubscribed from the system events")
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
				blockId := set.Block.Hex()
				// Uncomment for testing the decoders.
				// hexHash, err := types.NewHashFromHexString(hexHash, err := types.NewHashFromHexString("0x015fdc88c003b75d2871978a50238a4c7ac0adf34731a7e5007d3c20e2fae60d"))
				// if err != nil {
				// 	panic(err)
				// }
				// dataRaw, err := s.RestClient.RPC.State.GetStorageRaw(key, hexHash)
				// if err != nil {
				// 	panic(err)
				// }
				// log.Infof("Found: %v", dataRaw)
				for _, chng := range set.Changes {
					if !types.Eq(chng.StorageKey, key) || !chng.HasStorageData {
						// skip, we are only interested in events with content
						continue
					}
					out, err := substrate.DecodeEvent(chng.StorageData.Hex(), s.Metadata2, s.SpecVersion)
					if err != nil {
						log.Error(err)
					}
					outLs := out.([]interface{})
					for _, evt := range outLs {
						evtMap := evt.(map[string]interface{})

						if evtMap["event_id"] == "Swap" && evtMap["module_id"] == "Dex" {
							id := fmt.Sprintf("%s-%v", blockId, evtMap["event_idx"])
							log.Debugf("Event found: %s", id)
							log.Println("event map-------", evtMap)
							log.Println("event params-------", evtMap["params"])

							s.processSwap(id, evtMap)
						}
					}
				}
			}
		}
	}()
	return nil
}

func (s *AcalaswapScraper) processSwap(id string, e map[string]interface{}) error {
	swaps, err := s.normalizeAcalaswapSwaps(e)
	if err != nil {
		log.Error("error normalizing swap: ", err)
		return err
	}
	for _, swap := range swaps {
		pair := swap.Pair
		price, volume := s.getSwapData(swap)
		token0 := dia.Asset{
			Symbol:     pair.Token0.Symbol,
			Name:       pair.Token0.Name,
			Decimals:   pair.Token0.Decimals,
			Address:    pair.Token0.Address,
			Blockchain: Exchanges[s.exchangeName].BlockChain.Name,
		}
		token1 := dia.Asset{
			Symbol:   pair.Token1.Symbol,
			Name:     pair.Token1.Name,
			Decimals: pair.Token1.Decimals,
			Address:  pair.Token1.Address,

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
			ForeignTradeID: id,
			Source:         s.exchangeName,
			VerifiedPair:   true,
		}
		if price > 0 {
			log.Info("Got trade: ", t)
			s.chanTrades <- t
		}
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

	sub, err := s.WsClient.RPC.State.SubscribeStorageRaw([]types.StorageKey{key})
	if err != nil {
		return nil, nil, err
	}

	return sub, key, nil

}

func (s *AcalaswapScraper) getSwapData(swap *AcalaswapSwap) (price float64, volume float64) {
	volume = swap.Amount0
	price = math.Abs(swap.Amount1 / swap.Amount0)
	return
}

// normalizeAcalaswapSwap takes a swap as returned by the swap contract's channel and converts it to a AcalaswapSwap type
func (s *AcalaswapScraper) normalizeAcalaswapSwaps(swap map[string]interface{}) ([]*AcalaswapSwap, error) {
	trades := make([]*AcalaswapSwap, 0)
	var params, ok = swap["params"].([]scalecodec.EventParam)
	if !ok {
		return nil, errors.New("unknown params type")
	}
	// acountId := params[0].Value.(string)
	path, ok := params[1].Value.([]interface{})
	if !ok {
		return nil, errors.New("unknown path array type")
	}
	balances, ok := params[2].Value.([]interface{})
	if !ok {
		return nil, errors.New("unknown balance array type")
	}
	n := len(path)
	for i := 0; i < n-1; i++ {
		path0, ok := path[i].(map[string]interface{})
		if !ok {
			log.Error("unknown path type")
			continue
		}
		path1, ok := path[i+1].(map[string]interface{})
		if !ok {
			log.Error("unknown path type")
			continue
		}
		pair, err := s.GetPairByToken(path0, path1)
		if err != nil {
			log.Error("error getting pair by address: ", err)
			continue
		}
		decimals0 := int(pair.Token0.Decimals)
		decimals1 := int(pair.Token1.Decimals)
		amount0, ok := big.NewFloat(0).SetString(balances[i].(string))
		if !ok {
			continue
		}
		amount1, ok := big.NewFloat(0).SetString(balances[i+1].(string))
		if !ok {
			continue
		}
		amount0Normalized, _ := new(big.Float).Quo(amount0, new(big.Float).SetFloat64(math.Pow10(decimals0))).Float64()
		amount1Normalized, _ := new(big.Float).Quo(amount1, new(big.Float).SetFloat64(math.Pow10(decimals1))).Float64()

		trades = append(trades, &AcalaswapSwap{
			Timestamp: time.Now().Unix(),
			Pair:      pair,
			Amount0:   amount0Normalized,
			Amount1:   amount1Normalized,
		})
	}
	return trades, nil
}

// FetchAvailablePairs returns a list with all available trade pairs as dia.Pair for the pairDiscorvery service
func (s *AcalaswapScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	return
}

func (s *AcalaswapScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}
func (s *AcalaswapScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()

	if s.error != nil {
		return nil, s.error
	}

	if s.closed {
		return nil, errors.New("BalancerScraper is closed")
	}

	pairScraper := &AcalaswapPairScraper{
		parent: s,
		pair:   pair,
	}

	s.pairScrapers[pair.ForeignName] = pairScraper

	return pairScraper, nil
}
func (s *AcalaswapScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	if pair.ForeignName == "WETH" {
		pair.Symbol = "ETH"
	}
	return pair, nil
}

func (s *AcalaswapScraper) EncodeAssetId(currencyId map[string]interface{}) ([]byte, error) {
	if len(currencyId) > 1 {
		return nil, errors.New("unknown currency id")
	}

	var key string
	var rawValue interface{}

	for k, v := range currencyId {
		key = k
		rawValue = v
		break
	}
	var value []byte
	var enumIndex []byte
	var err error
	switch key {
	case "Erc20":
		enumIndex = []byte{0x00}
		value, err = types.Encode(types.MustHexDecodeString(rawValue.(string)))
	case "StableAssetPoolToken", "StableAssetId":
		enumIndex = []byte{0x01}
		value, err = types.Encode(rawValue.(uint32))
	// TODO: not working remotely, using same logic as StableAssetId for now!
	// case "StableAssetPoolToken":
	// 	enumIndex = []byte{0x03, 0x03}
	// 	value, err = types.Encode(rawValue.(uint32))
	case "ForeignAssetId", "ForeignAsset":
		enumIndex = []byte{0x02}
		value, err = types.Encode(rawValue.(uint16))
	// TODO: not working remotely, using same logic as ForeignAssetId for now!
	// case "ForeignAsset":
	// 	enumIndex = []byte{0x03, 0x05}
	// 	value, err = types.Encode(rawValue.(uint16))
	case "Token":
		enumIndex = []byte{0x03, 0x00}
		value, err = types.Encode(acalaswap.TokenSymbolMap[rawValue.(string)])

	case "LiquidCrowdloan":
		enumIndex = []byte{0x03, 0x04}
		value, err = types.Encode(rawValue.(uint32))

	default:
		return nil, errors.New("unknown currency type")
	}

	return append(enumIndex, value...), err
}

func (s *AcalaswapScraper) GetPairByToken(currencyIdIn map[string]interface{}, currencyIdOut map[string]interface{}) (pair *AcalaswapPair, err error) {
	var addressIn, addressOut string
	argsIn, err := s.EncodeAssetId(currencyIdIn)
	if err != nil {
		return nil, err
	}
	log.Debugln("currencyIn: %v", currencyIdIn)
	for k, v := range currencyIdIn {
		log.Infoln("token address  ", k, v)
		addressIn = fmt.Sprintf("%v:%v", k, v)
	}
	keyIn, err := types.CreateStorageKey(s.Metadata, "AssetRegistry", "AssetMetadatas", argsIn)
	if err != nil {
		return nil, err
	}

	resultIn, err := s.WsClient.RPC.State.GetStorageRawLatest(keyIn)
	if err != nil {
		return nil, err
	}
	log.Debugln("resultIn: 0x%x", *resultIn)
	// Decode the event records
	targetIn := acalaswap.AcalaAssetMetadata{}
	err = scale.NewDecoder(bytes.NewBuffer(*resultIn)).Decode(&targetIn)
	if err != nil {
		return nil, err
	}

	argsOut, err := s.EncodeAssetId(currencyIdOut)
	if err != nil {
		return nil, err
	}
	log.Debugln("currencyOut: %v", currencyIdOut)
	for k, v := range currencyIdOut {
		log.Debugln("token address  ", k, v)
		addressOut = fmt.Sprintf("%v:%v", k, v)
	}
	keyOut, err := types.CreateStorageKey(s.Metadata, "AssetRegistry", "AssetMetadatas", argsOut)
	if err != nil {
		return nil, err
	}

	resultOut, err := s.WsClient.RPC.State.GetStorageRawLatest(keyOut)
	if err != nil {
		return nil, err
	}
	log.Debugln("resultOut: 0x%x", *resultOut)
	// Decode the event records
	targetOut := acalaswap.AcalaAssetMetadata{}
	err = scale.NewDecoder(bytes.NewBuffer(*resultOut)).Decode(&targetOut)
	if err != nil {
		return nil, err
	}
	token0 := AcalaswapToken{
		Name:     string(targetIn.Name),
		Symbol:   string(targetIn.Symbol),
		Decimals: uint8(targetIn.Decimals),
		Address:  addressIn,
	}
	token1 := AcalaswapToken{
		Name:     string(targetOut.Name),
		Symbol:   string(targetOut.Symbol),
		Decimals: uint8(targetOut.Decimals),
		Address:  addressOut,
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
	parent *AcalaswapScraper
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
