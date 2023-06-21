package source

import (
	"strconv"
	"strings"

	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/itering/substrate-api-rpc/metadata"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/diadata-org/diadata/pkg/dia"
)

type AcalaAssetSource struct {
	RestClient   *gsrpc.SubstrateAPI
	WsClient     *gsrpc.SubstrateAPI
	SpecVersion  int
	Metadata2    *metadata.Instant
	Metadata     *types.Metadata
	assetChannel chan dia.Asset
	doneChannel  chan bool
	blockchain   string
	registries   []platypusRegistry
	waitTime     int
}

const (
	acalaHTTP  = "https://acala-polkadot.api.onfinality.io/public-rpc"
	karuraHTTP = "https://karura.api.onfinality.io/public-rpc"
	acalaWSS   = "wss://acala-rpc-2.aca-api.network/ws"
	karuraWSS  = "wss://karura-rpc-3.aca-api.network/ws"
)

func NewAcalaAssetScraper(exchange dia.Exchange) *AcalaAssetSource {

	var (
		assetChannel = make(chan dia.Asset)
		doneChannel  = make(chan bool)
		pas          *AcalaAssetSource
	)

	switch exchange.Name {

	case dia.AcalaswapExchange:
		pas = makeAcalaswapScraper(exchange, false, acalaHTTP, acalaWSS, "200", uint64(1102858))
	case dia.AcalaswapExchangeKarura:
		pas = makeAcalaswapScraper(exchange, false, karuraHTTP, karuraWSS, "200", uint64(1826919))
	}
	log.Println("dia.AcalaswapExchange", exchange.Name)

	log.Println("pas", pas)

	pas.assetChannel = assetChannel
	pas.doneChannel = doneChannel
	pas.blockchain = exchange.BlockChain.Name
	// pas.waitTime = waitTime

	go func() {
		pas.fetchAssets()
	}()

	return pas
}

func makeAcalaswapScraper(exchange dia.Exchange, listenByAddress bool, restDial string, wsDial string, waitMilliseconds string, startBlock uint64) *AcalaAssetSource {
	log.Println("making")

	var restClient, wsClient *gsrpc.SubstrateAPI
	var err error

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

	s := &AcalaAssetSource{
		WsClient:    wsClient,
		RestClient:  restClient,
		Metadata2:   metadata2,
		Metadata:    metadata,
		SpecVersion: int(runtimeVersionLatest.SpecVersion),
		waitTime:    waitTime,
		// listenByAddress: listenByAddress,
		// startBlock:      startBlock,
	}

	return s
}

func (pas *AcalaAssetSource) fetchAssets() {
	log.Println("fetchAssets")

	// for _, registry := range pas.registries {
	// 	err := pas.fetchPools(registry)
	// 	if err != nil {
	// 		log.Error("loadPoolsAndCoins: ", err)
	// 	}
	// }

	keyOut, err := types.CreateStorageKey(pas.Metadata, "AssetRegistry", "AssetRegistered", nil)
	if err != nil {
		log.Errorln("CreateStorageKey", err)

		return
	}

	dataRaw, err := pas.RestClient.RPC.State.GetKeysLatest(keyOut)

	// dataRaw, err := pas.RestClient.RPC.State.
	if err != nil {
		panic(err)
	}
	log.Infof("Found: %v", dataRaw)
}

// // Load pools and coins metadata from master registry
// func (scraper *AcalaAssetSource) fetchPools(registry platypusRegistry) error {
// 	log.Infof("loading master contract %s version %d and querying registry", registry.Address.Hex(), registry.Version)
// 	contractMaster, err := platypusfinance.NewBaseMasterPlatypusCaller(registry.Address, scraper.RestClient)
// 	if err != nil {
// 		log.Error("NewBaseMasterPlatypusCaller: ", err)
// 		return err
// 	}

// 	poolCount, err := contractMaster.PoolLength(&bind.CallOpts{})
// 	if err != nil {
// 		log.Error("PoolLength: ", err)
// 		return err
// 	}

// 	for i := 0; i < int(poolCount.Int64()); i++ {
// 		asset, err := contractMaster.PoolInfo(&bind.CallOpts{}, big.NewInt(int64(i)))
// 		if err != nil {
// 			log.Error("PoolInfo: ", err)
// 			return err
// 		}

// 		err = scraper.loadPoolData(asset.LpToken.Hex())
// 		if err != nil {
// 			log.Errorf("loadPoolData error at asset %s: %s", asset.LpToken.Hex(), err)
// 			return err
// 		}

// 	}

// 	return err
// }

// func (scraper *AcalaAssetSource) loadPoolData(poolAddress string) (err error) {

// 	contractAsset, err := platypusAssetABI.NewAssetCaller(common.HexToAddress(poolAddress), scraper.RestClient)
// 	if err != nil {
// 		log.Error("NewAssetCaller: ", err)
// 	}

// 	poolContract, err := contractAsset.Pool(&bind.CallOpts{})
// 	if err != nil {
// 		log.Error("Pool: ", err)
// 	}

// 	poolCaller, err := platypusPoolABI.NewPoolCaller(poolContract, scraper.RestClient)
// 	if err != nil {
// 		log.Error("NewPoolCaller: ", err)
// 	}

// 	poolTokenAddresses, errGetTokens := poolCaller.GetTokenAddresses(&bind.CallOpts{})
// 	if errGetTokens != nil {
// 		symbol, err := contractAsset.Symbol(&bind.CallOpts{})
// 		if err != nil {
// 			log.Error("contractAsset.Symbol: ", err)
// 		}
// 		log.Warnf("error calling GetTokenAddresses for %s %s asset: %s", symbol, poolAddress, errGetTokens)
// 	}

// 	for _, c := range poolTokenAddresses {
// 		var (
// 			symbol      string
// 			decimalsBig *big.Int
// 			decimals    uint8
// 			name        string
// 		)

// 		if c == common.HexToAddress("0x0000000000000000000000000000000000000000") {
// 			continue
// 		} else {
// 			contractToken, err := platypusTokenABI.NewTokenCaller(c, scraper.RestClient)
// 			if err != nil {
// 				log.Error("NewTokenCaller: ", err)
// 				continue
// 			}

// 			symbol, err = contractToken.Symbol(&bind.CallOpts{})
// 			if err != nil {
// 				log.Error("Symbol: ", err, c.Hex())
// 				continue
// 			}

// 			decimalsBig, err = contractToken.Decimals(&bind.CallOpts{})
// 			if err != nil {
// 				log.Error("Decimals: ", err)
// 				continue
// 			}
// 			decimals = uint8(decimalsBig.Uint64())

// 			name, err = contractToken.Name(&bind.CallOpts{})
// 			if err != nil {
// 				log.Error("Name: ", err)
// 				continue
// 			}

// 		}

// 		poolAsset := dia.Asset{
// 			Address:    c.Hex(),
// 			Blockchain: scraper.blockchain,
// 			Decimals:   decimals,
// 			Symbol:     symbol,
// 			Name:       name,
// 		}

// 		scraper.Asset() <- poolAsset

// 	}
// 	return
// }

func (pas *AcalaAssetSource) Asset() chan dia.Asset {
	return pas.assetChannel
}

func (pas *AcalaAssetSource) Done() chan bool {
	return pas.doneChannel
}
