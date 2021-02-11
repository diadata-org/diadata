package main

import (
	"bufio"
	"flag"
	"os"

	"github.com/diadata-org/diadata/internal/pkg/assetservice/database"
	"github.com/diadata-org/diadata/internal/pkg/assetservice/source"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
)

/*
fetch assets from various exchanges and save them in postgresql
*/

var blockchains map[string]dia.BlockChain

var (
	log         = logrus.New()
	assetSource *string
	key         *string
	secret      *string
	postgresKey = "postgres_key.txt"
)

func init() {
	assetSource = flag.String("source", "Uniswap", "Data source for asset collection")
	secret = flag.String("secret", "", "secret for asset source")

	flag.Parse()
	blockchains = make(map[string]dia.BlockChain)
	blockchains[dia.Bitcoin] = dia.BlockChain{Name: dia.BinanceExchange, NativeToken: "BTC", VerificationMechanism: dia.PROOF_OF_WORK}
	blockchains[dia.Ethereum] = dia.BlockChain{Name: dia.BinanceExchange, NativeToken: "ETH", VerificationMechanism: dia.PROOF_OF_WORK}
}

// NewAssetScraper returns a scraper for assets on @exchange
func NewAssetScraper(exchange string, secret string) source.AssetSource {
	switch exchange {
	case dia.UniswapExchange:
		return source.NewUniswapAssetSource(dia.Exchange{Name: dia.UniswapExchange, Centralized: false, BlockChain: blockchains[dia.Ethereum], Contract: common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")})
	default:
		return nil
	}
}

func main() {

	// data, err := database.NewPostgres("postgres://postgres:example@localhost/postgres")
	data, err := database.NewPostgres("postgresql://localhost/postgres?user=postgres&password=" + getPostgresKeyFromSecrets())
	if err != nil {
		log.Errorln("Error connecting to data source: ", err)
		return
	}
	runAssetSource(data, *assetSource, *secret)

}

func runAssetSource(data database.AssetStore, source, secret string) {

	log.Println("Fetching asset from ", source)

	asset := NewAssetScraper(source, secret)
	for {
		select {
		case receivedAsset := <-asset.Asset():
			log.Infoln("Received asset", receivedAsset)
			// asset, err := data.GetByName(receivedAsset.Name)
			// if err != nil {
			// 	log.Errorf("error getting asset %s: %v\n", receivedAsset.Name, err)
			// }
			// log.Infof("asset %s already in DB \n", asset.Name)
			err := data.Save(receivedAsset)
			if err != nil {
				log.Error("Error saving asset: ", err)
			} else {
				log.Infof("successfully stored %s \n", receivedAsset.Symbol)
			}
		}

	}

}

func feedAssetToRedis(assetsaver database.AssetStore) {

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

// getPostgresKeyFromSecrets returns the password for postgres db
func getPostgresKeyFromSecrets() string {
	var lines []string
	executionMode := os.Getenv("EXEC_MODE")
	var file *os.File
	var err error
	if executionMode == "production" {
		file, err = os.Open("/run/secrets/" + postgresKey)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		file, err = os.Open("../../../secrets/" + postgresKey)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if len(lines) != 1 {
		log.Fatal("Secrets file should have exactly one line")
	}
	return lines[0]
}
