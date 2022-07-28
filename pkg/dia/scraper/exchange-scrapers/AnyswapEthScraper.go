package scrapers

import (
	"encoding/json"
	"errors"
	"math"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"

	anyswap "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/anyswap"
	models "github.com/diadata-org/diadata/pkg/model"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	// anyswapEthereumContractAddress = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
	anyswapAPIURL = "https://bridgeapi.anyswap.exchange/v3/serverinfoV3?chainId=all&version=STABLEV3"
	// maps chainID to chain name.
	// TO DO: import from postgres.
	chainMap map[string]string
	// maps chainID+address to the corresponding dia.Asset.
	assetMap map[string]dia.Asset
)

const (
	anyswapWaitMilliseconds = "200"
)

type AnyswapToken struct {
	Address  common.Address
	Symbol   string
	Decimals uint8
	Name     string
}

type AnyswapPair struct {
	Token0      UniswapToken
	Token1      UniswapToken
	ForeignName string
	Address     common.Address
}

type AnyswapSwap struct {
	ID         string
	Timestamp  int64
	Pair       UniswapPair
	Amount0In  float64
	Amount0Out float64
	Amount1In  float64
	Amount1Out float64
}

type AnyswapScraper struct {
	WsClientMap   map[string]*ethclient.Client
	RestClientMap map[string]*ethclient.Client
	db            *models.RelDB
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
	pairScrapers     map[string]*AnyswapPairScraper
	exchangeName     string
	chanTrades       chan *dia.Trade
	waitTime         int
	anyswapAssetInfo map[string]map[string]interface{}
}

// NewUniswapScraper returns a new UniswapScraper for the given pair
func NewAnyswapScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *AnyswapScraper {
	log.Info("NewUniswapScraper: ", exchange.Name)
	var wsClientMap, restClientMap map[string]*ethclient.Client
	var waitTime int
	var err error
	assetMap = make(map[string]dia.Asset)

	switch exchange.Name {
	case dia.AnyswapExchange:
		exchangeFactoryContractAddress = exchange.Contract
		waitTimeString := utils.Getenv("UNISWAP_WAIT_TIME", anyswapWaitMilliseconds)
		waitTime, err = strconv.Atoi(waitTimeString)
		if err != nil {
			log.Error("could not parse wait time: ", err)
			waitTime = 100
		}

	}

	chainMap = make(map[string]string)
	chainMap["1"] = dia.ETHEREUM
	chainMap["56"] = dia.BINANCESMARTCHAIN
	chainMap["137"] = dia.POLYGON
	chainMap["250"] = dia.FANTOM
	chainMap["1284"] = dia.MOONBEAM
	chainMap["1285"] = dia.MOONRIVER
	chainMap["42161"] = dia.ARBITRUM
	chainMap["43114"] = dia.AVALANCHE

	restClientMap, wsClientMap, err = getClientMaps()
	if err != nil {
		log.Fatal("get client map: ", err)
	}

	allAssetsAllChains, err := fetchEndpoint(anyswapAPIURL)
	if err != nil {
		log.Fatal("fetch asset info from anyswap API endpoint: ", err)
	}

	s := &AnyswapScraper{
		RestClientMap:    restClientMap,
		WsClientMap:      wsClientMap,
		db:               relDB,
		shutdown:         make(chan nothing),
		shutdownDone:     make(chan nothing),
		pairScrapers:     make(map[string]*AnyswapPairScraper),
		exchangeName:     exchange.Name,
		error:            nil,
		chanTrades:       make(chan *dia.Trade),
		waitTime:         waitTime,
		anyswapAssetInfo: allAssetsAllChains,
	}

	if scrape {
		go s.mainLoop()
	}
	return s
}

// runs in a goroutine until s is closed
func (s *AnyswapScraper) mainLoop() {

	// wait for all pairs have added into s.PairScrapers
	time.Sleep(4 * time.Second)
	s.run = true

	var wg sync.WaitGroup
	for key := range chainMap {
		time.Sleep(time.Duration(s.waitTime) * time.Millisecond)
		wg.Add(1)
		go func(chainID string, w *sync.WaitGroup) {
			defer w.Done()
			s.ListenToChainOut(chainID)
		}(key, &wg)
	}
	wg.Wait()

}

// ListenToChainOut screens swaps out of the chain with @chainID to any other chain
// offered by Anyswap.
func (s *AnyswapScraper) ListenToChainOut(chainID string) {
	log.Info("listen to chain: ", chainMap[chainID])
	var err error

	// Fetch all addresses that can be bridged on chain with @chainID.
	addresses, err := getAddressesByChain(chainID)
	if err != nil {
		log.Error("")
	}

	// Switch from anyToken to underlying token.
	anyTokenMap, err := getAnyTokenMap()
	if err != nil {
		log.Error("get anyToken map: ", err)
	}

	// Listen to swaps out of current chain.
	sink, err := s.GetSwapOutChannel(addresses, chainID)
	if err != nil {
		log.Error("error fetching swaps channel: ", err)
	}

	go func() {
		for {
			rawSwap, ok := <-sink
			if ok {

				swap, err := s.processSwap(*rawSwap, anyTokenMap)
				if err != nil {
					log.Error("process swap: ", err)
				} else {
					log.Infof("got swap -- %v", swap)
					s.chanTrades <- &swap
				}
			}
		}
	}()
}

// processSwap returns a dia.Trade object from a rawSwap as emitted in LogAnySwapOut.
func (s *AnyswapScraper) processSwap(rawSwap anyswap.AnyswapV4RouterLogAnySwapOut, anyTokenMap map[string]string) (trade dia.Trade, err error) {

	// Get Basetoken
	fromChainID := rawSwap.FromChainID.String()
	basetokenaddress := rawSwap.Token.Hex()

	// If outToken is an anyToken, switch to the underlying asset.
	if underlyingToken, ok := anyTokenMap[fromChainID+"-"+basetokenaddress]; ok {
		basetokenaddress = underlyingToken
	}

	if basetoken, ok := assetMap[fromChainID+basetokenaddress]; ok {
		trade.BaseToken = basetoken
	} else {
		basetoken, err = s.db.GetAsset(common.HexToAddress(basetokenaddress).Hex(), chainMap[fromChainID])
		if err != nil {
			log.Errorf("get base asset %s on chainID %v: %v", basetokenaddress, rawSwap.FromChainID, err)
			return
		}
		trade.BaseToken = basetoken
		assetMap[fromChainID+basetokenaddress] = basetoken
	}

	// Get Quotetoken
	toChainID := rawSwap.ToChainID.String()
	toAddress := s.anyswapAssetInfo[fromChainID][strings.ToLower(basetokenaddress)].(map[string]interface{})["destChains"].(map[string]interface{})[toChainID].(map[string]interface{})["address"].(string)
	if quotetoken, ok := assetMap[toChainID+toAddress]; ok {
		trade.QuoteToken = quotetoken
	} else {
		quotetoken, err = s.db.GetAsset(common.HexToAddress(toAddress).Hex(), chainMap[toChainID])
		if err != nil {
			log.Errorf("get quote asset %s on chainID %s: %v", toAddress, rawSwap.ToChainID, err)
			return
		}
		trade.QuoteToken = quotetoken
		assetMap[fromChainID+toAddress] = quotetoken
	}

	trade.Symbol = trade.QuoteToken.Symbol
	trade.Pair = trade.QuoteToken.Symbol + "-" + trade.BaseToken.Symbol
	trade.Price = float64(1)
	trade.Volume, _ = new(big.Float).Quo(big.NewFloat(0).SetInt(rawSwap.Amount), new(big.Float).SetFloat64(math.Pow10(int(trade.BaseToken.Decimals)))).Float64()
	trade.ForeignTradeID = rawSwap.Raw.TxHash.String()
	trade.Time = time.Now()
	trade.Source = dia.AnyswapExchange
	trade.VerifiedPair = true
	return
}

// GetSwapOutChannel returns the channel @sink delivering the events LogAnySwapOut.
func (s *AnyswapScraper) GetSwapOutChannel(tokens []common.Address, chainID string) (chan *anyswap.AnyswapV4RouterLogAnySwapOut, error) {
	sink := make(chan *anyswap.AnyswapV4RouterLogAnySwapOut)
	anyswapRouterContractAddress := s.anyswapAssetInfo[chainID][strings.ToLower(tokens[0].Hex())].(map[string]interface{})["router"].(string)
	outFiltererContract, err := anyswap.NewAnyswapV4RouterFilterer(common.HexToAddress(anyswapRouterContractAddress), s.WsClientMap[chainID])
	if err != nil {
		log.Fatal(err)
	}
	_, err = outFiltererContract.WatchLogAnySwapOut(&bind.WatchOpts{}, sink, tokens, []common.Address{}, []common.Address{})
	if err != nil {
		return sink, err
	}
	return sink, nil
}

// getClientMaps returns maps for rest and ws clients. Keys are the corresponding chain IDs.
func getClientMaps() (map[string]*ethclient.Client, map[string]*ethclient.Client, error) {

	restClientMap := make(map[string]*ethclient.Client)
	wsClientMap := make(map[string]*ethclient.Client)
	for key := range chainMap {
		restClient, err := ethclient.Dial(utils.Getenv("ETH_URI_REST_"+key, ""))
		if err != nil {
			return restClientMap, wsClientMap, err
		}
		restClientMap[key] = restClient
		wsClient, err := ethclient.Dial(utils.Getenv("ETH_URI_WS_"+key, ""))
		if err != nil {
			return restClientMap, wsClientMap, err
		}
		wsClientMap[key] = wsClient

	}

	return restClientMap, wsClientMap, nil
}

// getAddressesByChain returns all addresses of assets which can be bridged away from the chain with @chainID.
// This includes anyTokens.
func getAddressesByChain(chainID string) (addresses []common.Address, err error) {
	allAssetsAllChains, err := fetchEndpoint(anyswapAPIURL)
	if err != nil {
		return
	}
	for address := range allAssetsAllChains[chainID] {
		addresses = append(addresses, common.HexToAddress(address))
		anyAddress := allAssetsAllChains[chainID][address].(map[string]interface{})["anyToken"].(map[string]interface{})["address"].(string)
		addresses = append(addresses, common.HexToAddress(anyAddress))
	}

	return
}

// getAnyTokenMap maps an anyToken to its underlying asset.
func getAnyTokenMap() (map[string]string, error) {
	anyTokenMap := make(map[string]string)
	allAssetsAllChains, err := fetchEndpoint(anyswapAPIURL)
	if err != nil {
		return anyTokenMap, err
	}
	for chainID := range allAssetsAllChains {
		for address := range allAssetsAllChains[chainID] {
			anyAddress := allAssetsAllChains[chainID][address].(map[string]interface{})["anyToken"].(map[string]interface{})["address"].(string)
			anyTokenMap[chainID+"-"+common.HexToAddress(anyAddress).Hex()] = common.HexToAddress(address).Hex()
		}
	}
	return anyTokenMap, nil
}

// fetchEndpoint returns all assets available in the Anyswap bridge obtained through an API endpoint.
func fetchEndpoint(url string) (response map[string]map[string]interface{}, err error) {
	// @response is of type map[chainID]map[assetAddress]interface{}
	data, _, err := utils.GetRequest(url)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return
	}
	return
}

// FetchAvailablePairs returns a list with all available trade pairs as dia.ExchangePair for the pairDiscorvery service
func (s *AnyswapScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	// TO DO: Use API in order to fetch available pairs
	// https://bridgeapi.anyswap.exchange/v3/serverinfoV3?chainId=all&version=STABLEV3

	return
}

// FillSymbolData is not used by DEX scrapers.
func (s *AnyswapScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

func (up *AnyswapScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *AnyswapScraper) Close() error {
	if s.closed {
		return errors.New("UniswapScraper: Already closed")
	}
	for i := range s.RestClientMap {
		s.WsClientMap[i].Close()
		s.RestClientMap[i].Close()
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *AnyswapScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("UniswapScraper: Call ScrapePair on closed scraper")
	}
	ps := &AnyswapPairScraper{
		parent: s,
		pair:   pair,
	}
	s.pairScrapers[pair.ForeignName] = ps
	return ps, nil
}

// UniswapPairScraper implements PairScraper for Uniswap
type AnyswapPairScraper struct {
	parent *AnyswapScraper
	pair   dia.ExchangePair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *AnyswapPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Channel returns a channel that can be used to receive trades
func (ps *AnyswapScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *AnyswapPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *AnyswapPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
