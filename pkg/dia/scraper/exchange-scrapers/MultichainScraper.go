package scrapers

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/anyerc20"
	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/uniswap"

	"math/big"

	"sync"
)

type BridgeSwapToken struct {
	Address  common.Address
	Symbol   string
	Decimals uint8
	Name     string
}

type BridgeSwapPair struct {
	Token0      BridgeSwapToken
	Token1      BridgeSwapToken
	ForeignName string
	Address     common.Address
}

type BridgeSwapSwap struct {
	TransactionHash string
	TokenAddress    common.Address
	ToAddress       common.Address
	Amount          *big.Int
	FromChainId     *big.Int
	ToChainId       *big.Int
}

type BridgeSwapScraper struct {
	WsClient   *ethclient.Client
	RestClient *ethclient.Client

	// signaling channels for session initialization and finishing
	//initDone     chan nothing
	//run          bool
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	pairScrapers map[string]*BridgeSwapPairScraper
	//exchangeName string
	chanTrades chan *dia.Trade
	// waitTime     int
	// If true, only pairs given in config file are scraped. Default is false.
	//listenByAddress bool
	relDB *models.RelDB
}

type MultiChainConfig struct {
	restURL                string
	wsURL                  string
	contractAddress        string
	contratDeployedAtBlock int64
}

var (
	restClients       map[string]*ethclient.Client
	wsClients         map[string]*ethclient.Client
	multichainconfigs map[string]MultiChainConfig
)

const (
	ftmHTTP   = "https://rpc.ftm.tools/"
	ethHTTP   = "https://eth-mainnet.alchemyapi.io/v2/UpWALFqrTh5m8bojhDcgtBIif-Ug5UUE"
	abiString = `[{
		"anonymous": false,
		"inputs": [{
			"indexed": true,
			"internalType": "bytes32",
			"name": "txhash",
			"type": "bytes32"
		}, {
			"indexed": true,
			"internalType": "address",
			"name": "token",
			"type": "address"
		}, {
			"indexed": true,
			"internalType": "address",
			"name": "to",
			"type": "address"
		}, {
			"indexed": false,
			"internalType": "uint256",
			"name": "amount",
			"type": "uint256"
		}, {
			"indexed": false,
			"internalType": "uint256",
			"name": "fromChainID",
			"type": "uint256"
		}, {
			"indexed": false,
			"internalType": "uint256",
			"name": "toChainID",
			"type": "uint256"
		}],
		"name": "LogAnySwapIn",
		"type": "event"
	}, {
		"anonymous": false,
		"inputs": [{
			"indexed": true,
			"internalType": "address",
			"name": "token",
			"type": "address"
		}, {
			"indexed": true,
			"internalType": "address",
			"name": "from",
			"type": "address"
		}, {
			"indexed": true,
			"internalType": "address",
			"name": "to",
			"type": "address"
		}, {
			"indexed": false,
			"internalType": "uint256",
			"name": "amount",
			"type": "uint256"
		}, {
			"indexed": false,
			"internalType": "uint256",
			"name": "fromChainID",
			"type": "uint256"
		}, {
			"indexed": false,
			"internalType": "uint256",
			"name": "toChainID",
			"type": "uint256"
		}],
		"name": "LogAnySwapOut",
		"type": "event"
	}]`
)

// NewBridgeSwapScraper returns a new BridgeSwapScraper for the given pair
func NewBridgeSwapScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *BridgeSwapScraper {
	var s *BridgeSwapScraper
	// var waitgroup sync.WaitGroup
	multichainconfigs = make(map[string]MultiChainConfig)

	multichainconfigs["1"] = MultiChainConfig{restURL: chainConfigs["1"].RestURL, wsURL: chainConfigs["1"].WSURL, contratDeployedAtBlock: 12242619, contractAddress: "0x765277eebeca2e31912c9946eae1021199b39c61"}
	multichainconfigs["56"] = MultiChainConfig{restURL: chainConfigs["56"].RestURL, wsURL: chainConfigs["56"].WSURL, contratDeployedAtBlock: 7910338, contractAddress: "0xd1c5966f9f5ee6881ff6b261bbeda45972b1b5f3"}
	multichainconfigs["137"] = MultiChainConfig{restURL: chainConfigs["137"].RestURL, wsURL: chainConfigs["137"].WSURL, contratDeployedAtBlock: 17355461, contractAddress: "0x6ff0609046a38d76bd40c5863b4d1a2dce687f73"}
	multichainconfigs["250"] = MultiChainConfig{restURL: chainConfigs["250"].RestURL, wsURL: chainConfigs["250"].WSURL, contratDeployedAtBlock: 8475644, contractAddress: "0x1ccca1ce62c62f7be95d4a67722a8fdbed6eecb4"}
	multichainconfigs["42161"] = MultiChainConfig{restURL: chainConfigs["42161"].RestURL, wsURL: chainConfigs["42161"].WSURL, contratDeployedAtBlock: 15315466, contractAddress: "0x650af55d5877f289837c30b94af91538a7504b76"}
	multichainconfigs["43114"] = MultiChainConfig{restURL: chainConfigs["43114"].RestURL, wsURL: chainConfigs["43114"].WSURL, contratDeployedAtBlock: 3397229, contractAddress: "0xB0731d50C681C45856BFc3f7539D5f61d4bE81D8"}

	log.Info("NewBridgeSwapScraper: ", exchange.Name)
	log.Infof("Init rest and ws client for %s.", exchange.BlockChain.Name)

	s = &BridgeSwapScraper{
		chanTrades:   make(chan *dia.Trade),
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*BridgeSwapPairScraper),
		relDB:        relDB,
	}

	if scrape {
		s.loop()
	}

	return s
}

func (s *BridgeSwapScraper) loop() {

	InitialiseRestClientsMap()
	InitialiseWsClientsMap()

	events := getFilteredEvents()
	s.checkTransactionOnChain(events)

}

func (s *BridgeSwapScraper) checkTransactionOnChain(events chan types.Log) {
	log.Infoln("Listening swaps")

	//get hex value for swap function name for the transaction
	LogAnySwapInbyteHex := crypto.Keccak256Hash([]byte("LogAnySwapIn(bytes32,address,address,uint256,uint256,uint256)"))

	for {
		msg := <-events
		switch msg.Topics[0].Hex() {

		case LogAnySwapInbyteHex.Hex():
			log.Infoln("msg TxHash", msg.TxHash)

			event, contractAbi := getEventDetailsAbi("LogAnySwapIn", msg)

			outAmount, _ := event[0].(*big.Int)
			fromChainIdValue, _ := event[1].(*big.Int)
			toChainIdValue, _ := event[2].(*big.Int)

			tokenAddress := getMultichainUnderlyingToken(common.HexToAddress(msg.Topics[2].Hex()), toChainIdValue.String())

			if tokenAddress == common.HexToAddress("0x0000000000000000000000000000000000000000") {
				tokenAddress = common.HexToAddress(msg.Topics[2].Hex())
			}
			bs := BridgeSwapSwap{
				TransactionHash: msg.Topics[1].Hex(),
				TokenAddress:    tokenAddress,
				ToAddress:       common.HexToAddress(msg.Topics[3].Hex()),
				Amount:          outAmount,
				FromChainId:     fromChainIdValue,
				ToChainId:       toChainIdValue,
			}
			log.Debugln("BridgeSwap", bs)
			tokenbridged, inAmount, err := getDetailsFromTransactionHash(msg, fromChainIdValue, contractAbi)
			if err != nil {
				continue
			}

			quoteTokenName, err := GetName(tokenAddress, toChainIdValue.String())
			if err != nil {
				log.Warnf("Error getting GetName token %s of chain id %s", tokenAddress, toChainIdValue.String())
			}
			quoteTokenSymbol, err := GetSymbol(tokenAddress, toChainIdValue.String())
			if err != nil {
				log.Warnf("Error getting GetSymbol token %s of chain id %s", tokenAddress, toChainIdValue.String())
			}
			quoteTokenDecimal, err := GetDecimals(tokenAddress, toChainIdValue.String())
			if err != nil {
				log.Warnf("Error getting GetDecimals token %s of chain id %s", tokenAddress, toChainIdValue.String())
				continue
			}

			baseBlockchain, isAvailable := evmID[fromChainIdValue.String()]
			if !isAvailable {
				log.Warn("Blockchain configs not available for chain with ID ", fromChainIdValue)
				continue

			}

			quoteBlockchain, isAvailable := evmID[toChainIdValue.String()]
			if !isAvailable {
				log.Warn("Blockchain configs not available for chain with ID ", toChainIdValue)
				continue

			}

			quoteToken := dia.Asset{
				Address:    tokenAddress.Hex(),
				Symbol:     quoteTokenSymbol,
				Name:       quoteTokenName,
				Decimals:   quoteTokenDecimal,
				Blockchain: Blockchains[quoteBlockchain].Name,
			}

			baseTokenName, err := GetName(tokenbridged, fromChainIdValue.String())
			if err != nil {
				log.Warnf("Error getting GetName token %s of chain id %s", tokenbridged, fromChainIdValue.String())
			}
			baseTokenSymbol, err := GetSymbol(tokenbridged, fromChainIdValue.String())
			if err != nil {
				log.Warnf("Error getting GetSymbol token %s of chain id %s", tokenbridged, fromChainIdValue.String())
			}
			baseTokenDecimal, err := GetDecimals(tokenbridged, fromChainIdValue.String())
			if err != nil {
				log.Warnf("Error getting GetDecimals token %s of chain id %s", tokenbridged, fromChainIdValue.String())
				continue
			}

			baseToken := dia.Asset{
				Address:    tokenbridged.Hex(),
				Symbol:     baseTokenName,
				Name:       baseTokenSymbol,
				Decimals:   baseTokenDecimal,
				Blockchain: Blockchains[baseBlockchain].Name,
			}

			inAmountt := inAmount.Quo(inAmount, inAmount.Exp(big.NewInt(10), big.NewInt(int64(baseTokenDecimal)), nil))

			outAmountt := outAmount.Div(outAmount, big.NewInt(int64(quoteTokenDecimal)))

			priceamt := inAmountt.Div(inAmountt, outAmountt)

			priceamtfloat, _ := new(big.Float).SetInt(priceamt).Float64()
			inAmountfloat, _ := new(big.Float).SetInt(inAmountt).Float64()

			t := &dia.Trade{
				Symbol:         baseTokenName + "" + quoteTokenName,
				Pair:           baseTokenName + "" + quoteTokenName,
				Price:          priceamtfloat,
				Volume:         inAmountfloat,
				BaseToken:      baseToken,
				QuoteToken:     quoteToken,
				Time:           time.Now(),
				ForeignTradeID: msg.TxHash.Hex(),
				// Source:         s.exchangeName,
				VerifiedPair: true,
			}
			log.Println("trade", t)
			s.mapasset(*t)
			// s.chanTrades <- t

		}
	}

}

func (s *BridgeSwapScraper) mapasset(t dia.Trade) {
	//check if quote token exists

	quoteToken_id, err := s.relDB.GetAssetID(t.QuoteToken)
	if err != nil {
		log.Errorln("Error getting quotetoken asset id", err)
	}
	baseToken_id, err := s.relDB.GetAssetID(t.BaseToken)
	if err != nil {
		log.Errorln("Error getting basetoken asset id", err)
	}

	quote_group_id, err := s.relDB.GetAssetMap(quoteToken_id)
	if err != nil {
		log.Errorln("quotetoken not exists", quoteToken_id)
	} else if quote_group_id != "" {
		log.Errorln("InsertAssetMap1 ", quote_group_id, baseToken_id)
		errInsertAssetMap := s.relDB.InsertAssetMap(quote_group_id, baseToken_id)
		log.Errorln("err InsertAssetMap1", errInsertAssetMap)
		return
	}

	base_group_id, err := s.relDB.GetAssetMap(baseToken_id)
	if err != nil {
		log.Errorln("base does not exists ")
	} else if quote_group_id != "" {
		log.Errorln("InsertAssetMap2 ", quote_group_id, baseToken_id)
		errInsertAssetMap := s.relDB.InsertAssetMap(base_group_id, quoteToken_id)
		log.Errorln("err InsertAssetMap2", errInsertAssetMap)

		return
	}
	log.Errorln("InsertAssetMap3 ", quoteToken_id)

	err = s.relDB.InsertNewAssetMap(quoteToken_id)
	log.Errorln("err InsertAssetMap3", err)

	gpid, err := s.relDB.GetAssetMap(quoteToken_id)
	if err != nil {
		log.Errorln("gpid generated err ", err)
	}
	s.relDB.InsertAssetMap(gpid, baseToken_id)
	log.Infoln("quote_group_id, base_group_id", baseToken_id, gpid)
}

func GetDecimals(tokenAddress common.Address, chainid string) (decimals uint8, err error) {

	var contract *uniswap.IERC20Caller
	contract, err = uniswap.NewIERC20Caller(tokenAddress, restClients[chainid])
	if err != nil {
		log.Error(err)
		return
	}
	decimals, err = contract.Decimals(&bind.CallOpts{})

	return
}

func GetName(tokenAddress common.Address, chainid string) (name string, err error) {

	var contract *uniswap.IERC20Caller
	contract, err = uniswap.NewIERC20Caller(tokenAddress, restClients[chainid])
	if err != nil {
		log.Error(err)
		return
	}
	name, err = contract.Name(&bind.CallOpts{})

	return
}

func GetSymbol(tokenAddress common.Address, chainid string) (name string, err error) {

	var contract *uniswap.IERC20Caller
	contract, err = uniswap.NewIERC20Caller(tokenAddress, restClients[chainid])
	if err != nil {
		log.Error(err)
		return
	}
	name, err = contract.Symbol(&bind.CallOpts{})

	return
}

func getEventDetailsAbi(funcName string, msg types.Log) ([]interface{}, abi.ABI) {
	var event []interface{}

	contractAbi, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		log.Fatal("read contract abi: ", err)
	}
	event, err = contractAbi.Unpack(funcName, msg.Data)
	if err != nil {
		log.Fatal("unpack event: ", err)
	}
	return event, contractAbi
}

func getDetailsFromTransactionHash(msg types.Log, fromChainIdValue *big.Int, contractAbi abi.ABI) (tokenmoved common.Address, inAmount *big.Int, err error) {
	//check if chain address value is present in map
	restClient, ok := restClients[fromChainIdValue.String()]
	//LogAnySwapOut (index_topic_1 address token, index_topic_2 address from, index_topic_3 address to, uint256 amount, uint256 fromChainID, uint256 toChainID)
	LogAnySwapOutbyteHex := crypto.Keccak256Hash([]byte("LogAnySwapOut(address,address,address,uint256,uint256,uint256)"))
	if ok {
		//get transaction receipt from transaction hash
		var receipt *types.Receipt
		receipt, err = restClient.TransactionReceipt(context.Background(), common.HexToHash(msg.Topics[1].Hex()))
		if err != nil {
			log.Errorf("fetch transaction receipt of tx %s on chain with ID %v: %v", msg.Topics[1].Hex(), fromChainIdValue.Int64(), err)
			return
		}

		for _, txlog := range receipt.Logs {
			switch txlog.Topics[0].Hex() {

			case LogAnySwapOutbyteHex.Hex():
				fmt.Println("token swapped between chains ", common.HexToAddress(txlog.Topics[1].Hex()))
				tokenmoved = getMultichainUnderlyingToken(common.HexToAddress(txlog.Topics[1].Hex()), fromChainIdValue.String())
				// fmt.Println("underlyingtoken", underlyingtoken)
				event, errUnpack := contractAbi.Unpack("LogAnySwapOut", txlog.Data)
				if errUnpack != nil {
					log.Fatal("unpack event LogAnySwapOut: ", errUnpack)
				}
				fmt.Println("------", event[0].(*big.Int))

				inAmount = event[0].(*big.Int)
			}
		}
		//putDetailsInChanel()
	} else {
		fmt.Println("-------Client not available for this chain---------", fromChainIdValue)
		err = errors.New("client not available for chain" + fromChainIdValue.String())
		return
	}

	return
}

func getMultichainUnderlyingToken(multichainTokenAddress common.Address, chainid string) (tokenAddress common.Address) {

	anyerc20caller, err := anyerc20.NewAnyerc20Caller(multichainTokenAddress, restClients[chainid])
	if err != nil {
		log.Errorln(err)
		return
	}

	tokenAddress, _ = anyerc20caller.Underlying(&bind.CallOpts{})

	return

}

func InitialiseRestClientsMap() {
	var err error
	restClients = make(map[string]*ethclient.Client)

	restClients["250"], err = ethclient.Dial(ftmHTTP)
	if err != nil {
		log.Fatal("init Fantom rest client: ", err)
	}

	restClients["1"], err = ethclient.Dial(ethHTTP)
	if err != nil {
		log.Fatal("init Ethereum rest client: ", err)
	}

	restClients["137"], err = ethclient.Dial(multichainconfigs["137"].restURL)
	if err != nil {
		log.Fatal("init Polygon rest client: ", err)
	}
	restClients["56"], err = ethclient.Dial(multichainconfigs["56"].restURL)
	if err != nil {
		log.Fatal("init Binance Smart Chain rest client: ", err)
	}

	restClients["25"], err = ethclient.Dial("https://cronosrpc-1.xstaking.sg")
	if err != nil {
		log.Fatal("init Cronos rest client: ", err)
	}

	restClients["43114"], err = ethclient.Dial("https://rpc.ankr.com/avalanche")
	if err != nil {
		log.Fatal("init Avalanche rest client: ", err)
	}

	restClients["10"], err = ethclient.Dial("https://mainnet.optimism.io")
	if err != nil {
		log.Fatal("init Optimism rest client: ", err)
	}

	restClients["1285"], err = ethclient.Dial("https://rpc.api.moonriver.moonbeam.network")
	if err != nil {
		log.Fatal("init Moonbeam rest client: ", err)
	}

	restClients["66"], err = ethclient.Dial("https://exchainrpc.okex.org")
	if err != nil {
		log.Fatal("init OKXChain rest client: ", err)
	}

	restClients["42161"], err = ethclient.Dial("https://arb1.arbitrum.io/rpc")
	if err != nil {
		log.Fatal("init Arbitrum rest client: ", err)
	}

}

func InitialiseWsClientsMap() {
	var err error

	wsClients = make(map[string]*ethclient.Client)

	for chainID, chainconfig := range multichainconfigs {
		wsClients[chainID], err = ethclient.Dial(chainconfig.wsURL)
		if err != nil {
			log.Errorf("init ws client on chain with id %s: %v", chainID, err)
		}

	}

}

func getFilteredEvents() chan types.Log {

	var channels []chan types.Log
	out := make(chan types.Log)

	for chainID, config := range multichainconfigs {

		//filter query by block number
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(config.contratDeployedAtBlock),
			Addresses: []common.Address{
				common.HexToAddress(config.contractAddress),
			},
		}

		events := make(chan types.Log)
		_, err := wsClients[chainID].SubscribeFilterLogs(context.Background(), query, events)
		if err != nil {
			log.Fatal("error connecting to wsclient", err)
		}

		channels = append(channels, events)

	}

	// merge all events in single channel

	for _, c := range channels {
		go func(c <-chan types.Log) {
			for v := range c {
				out <- v
			}
		}(c)
	}

	return out
}

// BridgeSwapPairScraper implements PairScraper for Uniswap
type BridgeSwapPairScraper struct {
	parent *BridgeSwapScraper
	pair   dia.ExchangePair
	closed bool
}

func (s *BridgeSwapScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {

	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("UniswapScraper: Call ScrapePair on closed scraper")
	}
	ps := &BridgeSwapPairScraper{
		parent: s,
		pair:   pair,
	}
	s.pairScrapers[pair.ForeignName] = ps
	return ps, nil

}

func (b *BridgeSwapScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {

	return []dia.ExchangePair{}, nil

}

func (ps *BridgeSwapScraper) Channel() chan *dia.Trade {
	return ps.chanTrades

}

func (s *BridgeSwapScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

func (up *BridgeSwapScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *BridgeSwapScraper) Close() error {
	if s.closed {
		return errors.New("BridgeSwapScraper: Already closed")
	}
	s.WsClient.Close()
	s.RestClient.Close()
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Close stops listening for trades of the pair associated with s
func (ps *BridgeSwapPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *BridgeSwapPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *BridgeSwapPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
