package main

import (
	"flag"
	"github.com/diadata-org/diadata/cmd/services/assets/source"
	"github.com/diadata-org/diadata/internal/pkg/database"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"sync"
)

/*


fetch assets from various exchange ans save it in mongo
*/
var blockchains map[string]dia.BlockChain

var log = logrus.New()
var feedRedis bool

func init() {

	flag.BoolVar(&feedRedis, "feedRedis", false, "Feed Asset to redis")
	flag.Parse()

	blockchains = make(map[string]dia.BlockChain)
	blockchains[dia.Bitcoin] = dia.BlockChain{Name: dia.BinanceExchange, NativeToken: "BTC", VerificationMechanism: dia.PROOF_OF_STAKE}
	blockchains[dia.Ethereum] = dia.BlockChain{Name: dia.BinanceExchange, NativeToken: "ETH", VerificationMechanism: dia.PROOF_OF_STAKE}

}

func NewAssetScraper(exchange string, key string, secret string) source.AssetSource {
	switch exchange {
	case dia.UniswapExchange:
		return source.NewUniswapAssetSource(dia.Exchange{Name: dia.UniswapExchange, Centralized: false, BlockChain: blockchains[dia.Ethereum], Contract: common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")})

	default:
		return nil
	}

}

func main() {

	data, err := database.NewPostgres("postgres://postgres:example@localhost/postgres")
	if err != nil {
		log.Errorln("Error connecting to data source", err)
		return
	}

	if feedRedis {
		feedAssetToRedis(data)
	} else {
		fetchAssetFromSource(data)

	}

}

func feedAssetToRedis(assetsaver database.AssetSaver) {

	assetCache := database.NewAssetcache()

	totalAssets, err := assetsaver.Count()
	if err != nil {
		log.Errorln("Error Getting Asset counts", err)
	}
	log.Infoln("Total Assets", totalAssets)

	// page size is 100
	totalPage := totalAssets / 100

	log.Infoln("Total Page", totalPage)
	var i int64
	for i = 0; i <= totalPage; i++ {
		assets, err := assetsaver.GetPage(i)
		if err != nil {
			log.Errorln("Error getting assets for page number", i, err)
		}

		for _, asset := range assets {
			log.Println(asset)
			assetCache.SetAsset(asset)
		}
		log.Infoln("updated cache with all assets")
	}

}
func fetchAssetFromSource(data database.AssetSaver) {
	var wg sync.WaitGroup

	log.Println("Fetching asset from ", dia.UniswapExchange)
	wg.Add(1)
	asset := NewAssetScraper(dia.UniswapExchange, "", "")
	for {
		select {
		case recievedAsset := <-asset.Asset():
			log.Infoln("Received asset", recievedAsset)
			err := data.Save(recievedAsset)
			if err != nil {
				log.Error("Error saving asset ", err)
			}
		}

	}
	wg.Wait()
}
