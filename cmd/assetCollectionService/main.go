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
func NewAssetScraper(exchange string, secret string, relDB *models.RelDB) source.AssetSource {
	switch exchange {
	case dia.UniswapExchange:
		return source.NewUniswapAssetSource(exchanges[dia.UniswapExchange], relDB)
	case dia.UniswapExchangeBase:
		return source.NewUniswapAssetSource(exchanges[dia.UniswapExchangeBase], relDB)
	case dia.PanCakeSwap:
		return source.NewUniswapAssetSource(exchanges[dia.PanCakeSwap], relDB)
	case dia.PanCakeSwapExchangeV3:
		return source.NewUniswapV3AssetSource(exchanges[dia.PanCakeSwapExchangeV3], relDB)
	case dia.ApeswapExchange:
		return source.NewUniswapAssetSource(exchanges[dia.ApeswapExchange], relDB)
	case dia.SushiSwapExchange:
		return source.NewUniswapAssetSource(exchanges[dia.SushiSwapExchange], relDB)
	case dia.SushiSwapExchangeArbitrum:
		return source.NewUniswapAssetSource(exchanges[dia.SushiSwapExchangeArbitrum], relDB)
	case dia.SushiSwapExchangeFantom:
		return source.NewUniswapAssetSource(exchanges[dia.SushiSwapExchangeFantom], relDB)
	case dia.SushiSwapExchangePolygon:
		return source.NewUniswapAssetSource(exchanges[dia.SushiSwapExchangePolygon], relDB)
	case dia.CamelotExchange:
		return source.NewUniswapAssetSource(exchanges[dia.CamelotExchange], relDB)
	case dia.CamelotExchangeV3:
		return source.NewCamelotV3AssetSource(exchanges[dia.CamelotExchangeV3], relDB)
	case dia.ThenaV3Exchange:
		return source.NewCamelotV3AssetSource(exchanges[dia.ThenaV3Exchange], relDB)
	case dia.DfynNetwork:
		return source.NewUniswapAssetSource(exchanges[dia.DfynNetwork], relDB)
	case dia.QuickswapExchange:
		return source.NewUniswapAssetSource(exchanges[dia.QuickswapExchange], relDB)
	case dia.UbeswapExchange:
		return source.NewUniswapAssetSource(exchanges[dia.UbeswapExchange], relDB)
	case dia.SpookyswapExchange:
		return source.NewUniswapAssetSource(exchanges[dia.SpookyswapExchange], relDB)
	case dia.SpiritswapExchange:
		return source.NewUniswapAssetSource(exchanges[dia.SpiritswapExchange], relDB)
	case dia.SolarbeamExchange:
		return source.NewUniswapAssetSource(exchanges[dia.SolarbeamExchange], relDB)
	case dia.RamsesV1Exchange:
		return source.NewUniswapAssetSource(exchanges[dia.RamsesV1Exchange], relDB)
	case dia.SerumExchange:
		return source.NewSerumAssetSource(exchanges[dia.SerumExchange])
	case dia.TrisolarisExchange:
		return source.NewUniswapAssetSource(exchanges[dia.TrisolarisExchange], relDB)
	case dia.AnyswapExchange:
		return source.NewAnyswapAssetSource(exchanges[dia.AnyswapExchange])
	case dia.NetswapExchange:
		return source.NewUniswapAssetSource(exchanges[dia.NetswapExchange], relDB)
	case dia.BalancerV2Exchange:
		return source.NewBalancerV2AssetSource(exchanges[dia.BalancerV2Exchange])
	case dia.BalancerV2ExchangeArbitrum:
		return source.NewBalancerV2AssetSource(exchanges[dia.BalancerV2ExchangeArbitrum])
	case dia.BeetsExchange:
		return source.NewBalancerV2AssetSource(exchanges[dia.BeetsExchange])
	case dia.BeetsV2ExchangeSonic:
		return source.NewBalancerV2AssetSource(exchanges[dia.BeetsV2ExchangeSonic])
	case dia.BeetsV3ExchangeSonic:
		return source.NewBalancerV3AssetSource(exchanges[dia.BeetsV3ExchangeSonic])
	case dia.BalancerV2ExchangePolygon:
		return source.NewBalancerV2AssetSource(exchanges[dia.BalancerV2ExchangePolygon])
	case dia.BalancerV3Exchange:
		return source.NewBalancerV3AssetSource(exchanges[dia.BalancerV3Exchange])
	case dia.BalancerV3ExchangeBase:
		return source.NewBalancerV3AssetSource(exchanges[dia.BalancerV3ExchangeBase])
	case dia.HuckleberryExchange:
		return source.NewUniswapAssetSource(exchanges[dia.HuckleberryExchange], relDB)
	case dia.TraderJoeExchange:
		return source.NewUniswapAssetSource(exchanges[dia.TraderJoeExchange], relDB)
	case dia.PangolinExchange:
		return source.NewUniswapAssetSource(exchanges[dia.PangolinExchange], relDB)
	case dia.TethysExchange:
		return source.NewUniswapAssetSource(exchanges[dia.TethysExchange], relDB)
	case dia.HermesExchange:
		return source.NewUniswapAssetSource(exchanges[dia.HermesExchange], relDB)
	case dia.OmniDexExchange:
		return source.NewUniswapAssetSource(exchanges[dia.OmniDexExchange], relDB)
	case dia.DiffusionExchange:
		return source.NewUniswapAssetSource(exchanges[dia.DiffusionExchange], relDB)
	case dia.ArthswapExchange:
		return source.NewUniswapAssetSource(exchanges[dia.ArthswapExchange], relDB)
	case dia.NileV1Exchange:
		return source.NewUniswapAssetSource(exchanges[dia.NileV1Exchange], relDB)
	case dia.UniswapExchangeV3:
		return source.NewUniswapV3AssetSource(exchanges[dia.UniswapExchangeV3], relDB)
	case dia.UniswapExchangeV3Base:
		return source.NewUniswapV3AssetSource(exchanges[dia.UniswapExchangeV3Base], relDB)
	case dia.UniswapExchangeV3Polygon:
		return source.NewUniswapV3AssetSource(exchanges[dia.UniswapExchangeV3Polygon], relDB)
	case dia.UniswapExchangeV3Arbitrum:
		return source.NewUniswapV3AssetSource(exchanges[dia.UniswapExchangeV3Arbitrum], relDB)
	case dia.UniswapExchangeV3Celo:
		return source.NewUniswapV3AssetSource(exchanges[dia.UniswapExchangeV3Celo], relDB)
	case dia.UniswapExchangeV4:
		return source.NewUniswapV4AssetSource(exchanges[dia.UniswapExchangeV4], relDB)
	case dia.RamsesV2Exchange:
		return source.NewUniswapV3AssetSource(exchanges[dia.RamsesV2Exchange], relDB)
	case dia.NileV2Exchange:
		return source.NewUniswapV3AssetSource(exchanges[dia.NileV2Exchange], relDB)
	case dia.StellaswapExchange:
		return source.NewUniswapAssetSource(exchanges[dia.StellaswapExchange], relDB)
	case dia.WanswapExchange:
		return source.NewUniswapAssetSource(exchanges[dia.WanswapExchange], relDB)
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
	case dia.VelodromeExchange:
		return source.NewVelodromeAssetSource(exchanges[dia.VelodromeExchange], relDB)
	case dia.VelodromeExchangeSwellchain:
		return source.NewVelodromeAssetSource(exchanges[dia.VelodromeExchangeSwellchain], relDB)
	case dia.VelodromeSlipstreamExchange:
		return source.NewVelodromeAssetSource(exchanges[dia.VelodromeSlipstreamExchange], relDB)
	case dia.TraderJoeExchangeV2_1:
		return source.NewTraderJoeAssetSource(exchanges[dia.TraderJoeExchangeV2_1], relDB)
	case dia.TraderJoeExchangeV2_1Arbitrum:
		return source.NewTraderJoeAssetSource(exchanges[dia.TraderJoeExchangeV2_1Arbitrum], relDB)
	case dia.TraderJoeExchangeV2_1Avalanche:
		return source.NewTraderJoeAssetSource(exchanges[dia.TraderJoeExchangeV2_1Avalanche], relDB)
	case dia.TraderJoeExchangeV2_1BNB:
		return source.NewTraderJoeAssetSource(exchanges[dia.TraderJoeExchangeV2_1BNB], relDB)
	case dia.TraderJoeExchangeV2_2Avalanche:
		return source.NewTraderJoeAssetSource(exchanges[dia.TraderJoeExchangeV2_2Avalanche], relDB)
	case dia.PearlfiExchangeTestnet:
		return source.NewUniswapAssetSource(exchanges[dia.PearlfiExchangeTestnet], relDB)
	case dia.PearlfiExchange:
		return source.NewUniswapAssetSource(exchanges[dia.PearlfiExchange], relDB)
	case dia.PearlfiStableswapExchange:
		return source.NewUniswapAssetSource(exchanges[dia.PearlfiStableswapExchange], relDB)
	case dia.ThenaExchange:
		return source.NewUniswapAssetSource(exchanges[dia.ThenaExchange], relDB)
	case dia.ShadowV2Exchange:
		return source.NewUniswapAssetSource(exchanges[dia.ShadowV2Exchange], relDB)
	case dia.ShadowV3Exchange:
		return source.NewUniswapV3AssetSource(exchanges[dia.ShadowV3Exchange], relDB)
	case "assetlists":
		return source.NewJSONReader(exchange, secret)
	case dia.AyinExchange:
		return source.NewAyinAssetSource(exchanges[dia.AyinExchange], relDB)
	case dia.BitflowExchange:
		return source.NewBitflowAssetSource(exchanges[dia.BitflowExchange], relDB)
	case dia.VelarExchange:
		return source.NewVelarAssetSource(exchanges[dia.VelarExchange], relDB)
	case dia.AerodromeSlipstreamExchange:
		return source.NewVelodromeAssetSource(exchanges[dia.AerodromeSlipstreamExchange], relDB)
	case dia.AerodromeV1Exchange:
		return source.NewVelodromeAssetSource(exchanges[dia.AerodromeV1Exchange], relDB)
	case dia.HydrationExchange:
		return source.NewHydrationAssetSource(exchanges[dia.HydrationExchange], relDB)
	case dia.BifrostExchange:
		return source.NewBifrostAssetSource(exchanges[dia.BifrostExchange], relDB)
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
	asset := NewAssetScraper(source, secret, relDB)

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
