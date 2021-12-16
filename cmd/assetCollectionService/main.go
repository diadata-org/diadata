package main

import (
	"flag"
	"time"

	scrapers "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/service/assetservice/source"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/sirupsen/logrus"
)

/*
fetch assets from various exchanges and save them in postgresql
*/

// var blockchains map[string]dia.BlockChain

var (
	log         = logrus.New()
	assetSource *string
	// key         *string
	secret  *string
	caching *bool
)

const (
	fetchPeriodMinutes = 24 * 60
)

var exchanges map[string]dia.Exchange

func init() {
	assetSource = flag.String("source", "Uniswap", "Data source for asset collection")
	secret = flag.String("secret", "", "secret for asset source")
	caching = flag.Bool("caching", true, "caching assets in redis")
	flag.Parse()

	// source, err := datasource.InitSource()
	// if err != nil {
	// 	panic(err)
	// }
	exchanges = scrapers.Exchanges
}

// NewAssetScraper returns a scraper for assets on @exchange.
// For NewJSONReader @exchange is the folder in the config folder and @secret the filename.
func NewAssetScraper(exchange string, secret string) source.AssetSource {
	switch exchange {
	case dia.UniswapExchange:
		return source.NewUniswapAssetSource(exchanges[dia.UniswapExchange])
	case dia.PanCakeSwap:
		return source.NewUniswapAssetSource(exchanges[dia.PanCakeSwap])
	case dia.SushiSwapExchange:
		return source.NewUniswapAssetSource(exchanges[dia.SushiSwapExchange])
	case dia.DfynNetwork:
		return source.NewUniswapAssetSource(exchanges[dia.DfynNetwork])
	case dia.QuickswapExchange:
		return source.NewUniswapAssetSource(exchanges[dia.QuickswapExchange])
	case dia.UbeswapExchange:
		return source.NewUniswapAssetSource(exchanges[dia.UbeswapExchange])
	case dia.SpookyswapExchange:
		return source.NewUniswapAssetSource(exchanges[dia.SpookyswapExchange])
	case dia.SpiritswapExchange:
		return source.NewUniswapAssetSource(exchanges[dia.SpiritswapExchange])
	case dia.SolarbeamExchange:
		return source.NewUniswapAssetSource(exchanges[dia.SolarbeamExchange])
	case dia.SerumExchange:
		return source.NewSerumAssetSource(exchanges[dia.SerumExchange])
	case dia.TrisolarisExchange:
		return source.NewUniswapAssetSource(exchanges[dia.TrisolarisExchange])
	case "assetlists":
		return source.NewJSONReader(exchange, secret)
	default:
		return nil
	}
}

func main() {

	relDB, err := models.NewRelDataStore()
	if err != nil {
		log.Errorln("Error connecting to asset DB: ", err)
		return
	}
	runAssetSource(relDB, *assetSource, *caching, *secret)
	log.Infof("Successfully ran asset collector for %s", *assetSource)
}

func runAssetSource(relDB *models.RelDB, source string, caching bool, secret string) {
	// TO DO: check for duplicate key error and return if error is different
	log.Println("Fetching asset from ", source)
	asset := NewAssetScraper(source, secret)
	for receivedAsset := range asset.Asset() {
		// Set to persistent DB
		err := relDB.SetAsset(receivedAsset)
		if err != nil {
			log.Errorf("Error saving asset %v: %v", receivedAsset, err)
		} else {
			log.Info("successfully set asset ", receivedAsset)
		}

		// Set to cache
		if caching {
			err := relDB.SetAssetCache(receivedAsset)
			if err != nil {
				log.Error("Error caching asset: ", err)
			}
		}

	}
}
