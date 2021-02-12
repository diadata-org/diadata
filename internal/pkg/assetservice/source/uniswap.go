package source

import (
	"fmt"
	"math/big"
	"time"

	uniswapcontract "github.com/diadata-org/diadata/internal/pkg/exchange-scrapers/uniswap"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type UniswapPair struct {
	Token0      dia.Asset
	Token1      dia.Asset
	ForeignName string
	Address     common.Address
}

const (
	wsDial   = "wss://eth-mainnet.ws.alchemyapi.io/v2/CP4k5FRH3BZdqr_ANmGJFr0iI076CxR8"
	restDial = "https://eth-mainnet.alchemyapi.io/v2/CP4k5FRH3BZdqr_ANmGJFr0iI076CxR8"

	wsDialBSC   = "wss://bsc-ws-node.nariox.org:443"
	restDialBSC = "https://bsc-dataseed2.defibit.io/"
)

type UniswapAssetSource struct {
	WsClient   *ethclient.Client
	RestClient *ethclient.Client
	asset      chan dia.Asset
}

var exchangeFactoryContractAddress string

func NewUniswapAssetSource(exchange dia.Exchange) *UniswapAssetSource {

	var wsClient, restClient *ethclient.Client
	var err error
	var asset = make(chan dia.Asset)

	switch exchange.Name {
	case dia.UniswapExchange:
		exchangeFactoryContractAddress = exchange.Contract.String()
		wsClient, err = ethclient.Dial(wsDial)
		if err != nil {
			log.Fatal(err)
		}

		restClient, err = ethclient.Dial(restDial)
		if err != nil {
			log.Fatal(err)
		}
		break
	case dia.SushiSwapExchange:
		exchangeFactoryContractAddress = exchange.Contract.String()
		wsClient, err = ethclient.Dial(wsDial)
		if err != nil {
			log.Fatal(err)
		}

		restClient, err = ethclient.Dial(restDial)
		if err != nil {
			log.Fatal(err)
		}

		break
	case dia.PanCakeSwap:
		log.Infoln("Init ws and rest client for BSC chain")
		wsClient, err = ethclient.Dial(wsDialBSC)
		if err != nil {
			log.Fatal(err)
		}
		restClient, err = ethclient.Dial(restDialBSC)
		if err != nil {
			log.Fatal(err)
		}
		exchangeFactoryContractAddress = exchange.Contract.String()
	}

	uas := &UniswapAssetSource{WsClient: wsClient, RestClient: restClient, asset: asset}
	go func() {
		uas.fetchAssets()

	}()
	return uas

}

func (uas *UniswapAssetSource) Asset() chan dia.Asset {
	return uas.asset
}

func (uas *UniswapAssetSource) getNumPairs() (int, error) {
	var contract *uniswapcontract.IUniswapV2FactoryCaller
	contract, err := uniswapcontract.NewIUniswapV2FactoryCaller(common.HexToAddress(exchangeFactoryContractAddress), uas.RestClient)
	if err != nil {
		log.Error(err)
	}

	numPairs, err := contract.AllPairsLength(&bind.CallOpts{})
	return int(numPairs.Int64()), err
}

func (uas *UniswapAssetSource) fetchAssets() {
	var blockchain dia.BlockChain
	blockchain.Name = dia.Ethereum
	genesisDate, err := time.Parse("2006-01-02", "2015-07-30")
	if err != nil {
		fmt.Println(err)
	}
	blockchain.GenesisDate = genesisDate
	blockchain.NativeToken = "ETH"
	blockchain.VerificationMechanism = dia.PROOF_OF_WORK

	numPairs, err := uas.getNumPairs()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Found ", numPairs, " pairs")
	checkMap := make(map[string]struct{})
	for i := 0; i < numPairs; i++ {
		pair, err := uas.GetPairByID(int64(i))
		if err != nil {
			log.Errorln("Error getting pair with ID ", i)
		}
		asset0 := pair.Token0
		asset1 := pair.Token1
		asset0.Blockchain = blockchain
		asset1.Blockchain = blockchain
		// Don't repeat sending already sent assets
		if _, ok := checkMap[asset0.Address]; !ok {
			if asset0.Symbol != "" {
				checkMap[asset0.Address] = struct{}{}
				uas.asset <- asset0
			}
		}
		if _, ok := checkMap[asset1.Address]; !ok {
			if asset1.Symbol != "" {
				checkMap[asset1.Address] = struct{}{}
				uas.asset <- asset1
			}
		}
	}
}

// GetPairByID returns the UniswapPair with the integer id @num
func (uas *UniswapAssetSource) GetPairByID(num int64) (UniswapPair, error) {
	var contract *uniswapcontract.IUniswapV2FactoryCaller
	contract, err := uniswapcontract.NewIUniswapV2FactoryCaller(common.HexToAddress(exchangeFactoryContractAddress), uas.RestClient)
	if err != nil {
		log.Error(err)
		return UniswapPair{}, err
	}
	numToken := big.NewInt(num)
	pairAddress, err := contract.AllPairs(&bind.CallOpts{}, numToken)
	if err != nil {
		log.Error(err)
		return UniswapPair{}, err
	}

	pair, err := uas.GetPairByAddress(pairAddress)
	if err != nil {
		log.Error(err)
		return UniswapPair{}, err
	}
	// log.Infof("Get pair with ID %v: %v", num, pair)
	return pair, err
}

func (uas *UniswapAssetSource) GetPairByAddress(pairAddress common.Address) (pair UniswapPair, err error) {
	connection := uas.RestClient
	var pairContract *uniswapcontract.IUniswapV2PairCaller
	pairContract, err = uniswapcontract.NewIUniswapV2PairCaller(pairAddress, connection)
	if err != nil {
		log.Error(err)
		return UniswapPair{}, err
	}

	// Getting tokens from pair ---------------------
	address0, _ := pairContract.Token0(&bind.CallOpts{})
	address1, _ := pairContract.Token1(&bind.CallOpts{})
	var token0Contract *uniswapcontract.IERC20Caller
	var token1Contract *uniswapcontract.IERC20Caller
	token0Contract, err = uniswapcontract.NewIERC20Caller(address0, connection)
	if err != nil {
		log.Error(err)
	}
	token1Contract, err = uniswapcontract.NewIERC20Caller(address1, connection)
	if err != nil {
		log.Error(err)
	}
	symbol0, err := token0Contract.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Error(err)
	}
	symbol1, err := token1Contract.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Error(err)
	}
	decimals0, err := token0Contract.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Error(err)
	}
	decimals1, err := token1Contract.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Error(err)
	}

	name0, err := uas.GetName(address0)
	if err != nil {
		log.Error(err)
		return UniswapPair{}, err
	}
	name1, err := uas.GetName(address1)
	if err != nil {
		log.Error(err)
		return UniswapPair{}, err
	}
	token0 := dia.Asset{
		Address:  address0.String(),
		Symbol:   symbol0,
		Name:     name0,
		Decimals: decimals0,
	}
	token1 := dia.Asset{
		Address:  address1.String(),
		Symbol:   symbol1,
		Name:     name1,
		Decimals: decimals1,
	}
	pair.Token0 = token0
	pair.Token1 = token1

	return pair, nil
}

// GetDecimals returns the decimals of the token with address @tokenAddress
func (uas *UniswapAssetSource) GetDecimals(tokenAddress common.Address) (decimals uint8, err error) {

	var contract *uniswapcontract.IERC20Caller
	contract, err = uniswapcontract.NewIERC20Caller(tokenAddress, uas.RestClient)
	if err != nil {
		log.Error(err)
		return
	}
	decimals, err = contract.Decimals(&bind.CallOpts{})

	return
}

func (uas *UniswapAssetSource) GetName(tokenAddress common.Address) (name string, err error) {

	var contract *uniswapcontract.IERC20Caller
	contract, err = uniswapcontract.NewIERC20Caller(tokenAddress, uas.RestClient)
	if err != nil {
		log.Error(err)
		return
	}
	name, err = contract.Name(&bind.CallOpts{})

	return
}
