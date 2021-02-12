package main

import (
	"bufio"
	"flag"
	"os"
	"time"

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
	caching     *bool
)

const (
	postgresKey        = "postgres_key.txt"
	fetchPeriodMinutes = 24 * 60
)

func init() {
	assetSource = flag.String("source", "Uniswap", "Data source for asset collection")
	secret = flag.String("secret", "", "secret for asset source")
	caching = flag.Bool("caching", true, "caching assets in redis")

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

	assetDB, err := database.NewAssetDB("postgresql://localhost/postgres?user=postgres&password=" + getPostgresKeyFromSecrets())
	if err != nil {
		log.Errorln("Error connecting to asset DB: ", err)
		return
	}
	// Initial run:
	runAssetSource(assetDB, *assetSource, *caching, *secret)
	// Afterwards, run every @fetchPeriodMinutes
	ticker := time.NewTicker(fetchPeriodMinutes * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				runAssetSource(assetDB, *assetSource, *caching, *secret)
			}
		}
	}()
	select {}

}

func runAssetSource(assetDB database.AssetStore, source string, caching bool, secret string) {

	assetCache, err := database.NewAssetCache()
	if err != nil {
		log.Errorln("Error connecting to asset cache: ", err)
		return
	}

	log.Println("Fetching asset from ", source)
	asset := NewAssetScraper(source, secret)
	for {
		select {
		case receivedAsset := <-asset.Asset():
			// Set to persistent DB
			err := assetDB.SetAsset(receivedAsset)
			if err != nil {
				log.Errorf("Error saving asset %v: %v", receivedAsset, err)
			} else {
				log.Info("successfully set asset ", receivedAsset)
			}

			// Set to cache
			if caching {
				err := assetCache.SetAsset(receivedAsset)
				if err != nil {
					log.Error("Error caching asset: ", err)
				}
			}
		}
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

func feedAssetToRedis(assetsaver database.AssetStore) {

	assetCache, err := database.NewAssetCache()
	if err != nil {
		log.Errorln("Error connecting to asset cache: ", err)
		return
	}
	totalAssets, err := assetsaver.Count()
	if err != nil {
		log.Errorln("Error Getting Asset counts", err)
	}
	log.Infoln("Total Assets: ", totalAssets)

	pageIndex := uint32(0)
	hasNextPage := true
	for hasNextPage {
		assets := []dia.Asset{}
		assets, hasNextPage, err = assetsaver.GetPage(pageIndex)
		if err != nil {
			log.Errorln("Error getting assets for page number", pageIndex, err)
		}
		for _, asset := range assets {
			assetCache.SetAsset(asset)
		}
		if hasNextPage {
			pageIndex++
		}
	}
}
