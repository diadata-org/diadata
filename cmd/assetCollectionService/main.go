package main

import (
	"flag"

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
	case dia.ApeswapExchange:
		return source.NewUniswapAssetSource(exchanges[dia.ApeswapExchange])
	case dia.SushiSwapExchange:
		return source.NewUniswapAssetSource(exchanges[dia.SushiSwapExchange])
	case dia.SushiSwapExchangeArbitrum:
		return source.NewUniswapAssetSource(exchanges[dia.SushiSwapExchangeArbitrum])
	case dia.SushiSwapExchangeFantom:
		return source.NewUniswapAssetSource(exchanges[dia.SushiSwapExchangeFantom])
	case dia.CamelotExchange:
		return source.NewUniswapAssetSource(exchanges[dia.CamelotExchange])
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
	case dia.AnyswapExchange:
		return source.NewAnyswapAssetSource(exchanges[dia.AnyswapExchange])
	case dia.NetswapExchange:
		return source.NewUniswapAssetSource(exchanges[dia.NetswapExchange])
	case dia.BalancerV2Exchange:
		return source.NewBalancerV2AssetSource(exchanges[dia.BalancerV2Exchange])
	case dia.BalancerV2ExchangeArbitrum:
		return source.NewBalancerV2AssetSource(exchanges[dia.BalancerV2ExchangeArbitrum])
	case dia.BeetsExchange:
		return source.NewBalancerV2AssetSource(exchanges[dia.BeetsExchange])
	case dia.BalancerV2ExchangePolygon:
		return source.NewBalancerV2AssetSource(exchanges[dia.BalancerV2ExchangePolygon])
	case dia.HuckleberryExchange:
		return source.NewUniswapAssetSource(exchanges[dia.HuckleberryExchange])
	case dia.TraderJoeExchange:
		return source.NewUniswapAssetSource(exchanges[dia.TraderJoeExchange])
	case dia.PangolinExchange:
		return source.NewUniswapAssetSource(exchanges[dia.PangolinExchange])
	case dia.TethysExchange:
		return source.NewUniswapAssetSource(exchanges[dia.TethysExchange])
	case dia.HermesExchange:
		return source.NewUniswapAssetSource(exchanges[dia.HermesExchange])
	case dia.OmniDexExchange:
		return source.NewUniswapAssetSource(exchanges[dia.OmniDexExchange])
	case dia.DiffusionExchange:
		return source.NewUniswapAssetSource(exchanges[dia.DiffusionExchange])
	case dia.ArthswapExchange:
		return source.NewUniswapAssetSource(exchanges[dia.ArthswapExchange])
	case dia.UniswapExchangeV3:
		return source.NewUniswapV3AssetSource(exchanges[dia.UniswapExchangeV3])
	case dia.UniswapExchangeV3Polygon:
		return source.NewUniswapV3AssetSource(exchanges[dia.UniswapExchangeV3Polygon])
	case dia.UniswapExchangeV3Arbitrum:
		return source.NewUniswapV3AssetSource(exchanges[dia.UniswapExchangeV3Arbitrum])
	case dia.StellaswapExchange:
		return source.NewUniswapAssetSource(exchanges[dia.StellaswapExchange])
	case dia.WanswapExchange:
		return source.NewUniswapAssetSource(exchanges[dia.WanswapExchange])
	case dia.CurveFIExchange:
		return source.NewCurvefiAssetSource(exchanges[dia.CurveFIExchange])
	case dia.CurveFIExchangeFantom:
		return source.NewCurvefiAssetSource(exchanges[dia.CurveFIExchangeFantom])
	case dia.CurveFIExchangeMoonbeam:
		return source.NewCurvefiAssetSource(exchanges[dia.CurveFIExchangeMoonbeam])
	case dia.CurveFIExchangePolygon:
		return source.NewCurvefiAssetSource(exchanges[dia.CurveFIExchangePolygon])
	case dia.CurveFIExchangeArbitrum:
		return source.NewCurvefiAssetSource(exchanges[dia.CurveFIExchangeArbitrum])
	case dia.PlatypusExchange:
		return source.NewPlatypusScraper(exchanges[dia.PlatypusExchange])
	case dia.OrcaExchange:
		return source.NewOrcaScraper(exchanges[dia.OrcaExchange])
	case dia.OsmosisExchange:
		return source.NewOsmosisScraper(exchanges[dia.OsmosisExchange])
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
	log.Println("Fetching asset from ", source)
	asset := NewAssetScraper(source, secret)

	for {
		select {
		case receivedAsset := <-asset.Asset():
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
		case <-asset.Done():
			return
		}
	}

}
