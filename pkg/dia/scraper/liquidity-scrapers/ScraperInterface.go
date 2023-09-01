package liquidityscrapers

import (
	"github.com/diadata-org/diadata/pkg/dia"
	scrapers "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/sirupsen/logrus"
)

type LiquidityScraper interface {
	Pool() chan dia.Pool
	Done() chan bool
}

const (
	GLOBAL_NATIVE_LIQUIDITY_THRESHOLD = 0.5
)

var (
	log        *logrus.Logger
	exchanges  map[string]dia.Exchange
	priceCache = make(map[string]float64)
)

func init() {
	log = logrus.New()
	exchanges = scrapers.Exchanges

}

// NewLiquidityScraper returns a liquidity scraper for @source.
func NewLiquidityScraper(source string, relDB *models.RelDB, datastore *models.DB) LiquidityScraper {
	switch source {
	case dia.UniswapExchange:
		return NewUniswapScraper(exchanges[dia.UniswapExchange], relDB, datastore)
	case dia.SushiSwapExchange:
		return NewUniswapScraper(exchanges[dia.SushiSwapExchange], relDB, datastore)
	case dia.SushiSwapExchangePolygon:
		return NewUniswapScraper(exchanges[dia.SushiSwapExchangePolygon], relDB, datastore)
	case dia.SushiSwapExchangeFantom:
		return NewUniswapScraper(exchanges[dia.SushiSwapExchangeFantom], relDB, datastore)
	case dia.SushiSwapExchangeArbitrum:
		return NewUniswapScraper(exchanges[dia.SushiSwapExchangeArbitrum], relDB, datastore)
	case dia.CamelotExchange:
		return NewUniswapScraper(exchanges[dia.CamelotExchange], relDB, datastore)
	case dia.PanCakeSwap:
		return NewUniswapScraper(exchanges[dia.PanCakeSwap], relDB, datastore)
	case dia.PanCakeSwapExchangeV3:
		return NewUniswapV3Scraper(exchanges[dia.PanCakeSwapExchangeV3], relDB, datastore)
	case dia.DfynNetwork:
		return NewUniswapScraper(exchanges[dia.DfynNetwork], relDB, datastore)
	case dia.QuickswapExchange:
		return NewUniswapScraper(exchanges[dia.QuickswapExchange], relDB, datastore)
	case dia.UbeswapExchange:
		return NewUniswapScraper(exchanges[dia.UbeswapExchange], relDB, datastore)
	case dia.SpookyswapExchange:
		return NewUniswapScraper(exchanges[dia.SpookyswapExchange], relDB, datastore)
	case dia.SpiritswapExchange:
		return NewUniswapScraper(exchanges[dia.SpiritswapExchange], relDB, datastore)
	case dia.SolarbeamExchange:
		return NewUniswapScraper(exchanges[dia.SolarbeamExchange], relDB, datastore)
	case dia.TrisolarisExchange:
		return NewUniswapScraper(exchanges[dia.TrisolarisExchange], relDB, datastore)
	case dia.NetswapExchange:
		return NewUniswapScraper(exchanges[dia.NetswapExchange], relDB, datastore)
	case dia.HuckleberryExchange:
		return NewUniswapScraper(exchanges[dia.HuckleberryExchange], relDB, datastore)
	case dia.TraderJoeExchange:
		return NewUniswapScraper(exchanges[dia.TraderJoeExchange], relDB, datastore)
	case dia.TraderJoeExchangeV2_1:
		return NewUniswapScraper(exchanges[dia.TraderJoeExchangeV2_1], relDB, datastore)
	case dia.TraderJoeExchangeV2_1Arbitrum:
		return NewUniswapScraper(exchanges[dia.TraderJoeExchangeV2_1Arbitrum], relDB, datastore)
	case dia.TraderJoeExchangeV2_1Avalanche:
		return NewUniswapScraper(exchanges[dia.TraderJoeExchangeV2_1Avalanche], relDB, datastore)
	case dia.TraderJoeExchangeV2_1BNB:
		return NewUniswapScraper(exchanges[dia.TraderJoeExchangeV2_1BNB], relDB, datastore)
	case dia.PangolinExchange:
		return NewUniswapScraper(exchanges[dia.PangolinExchange], relDB, datastore)
	case dia.TethysExchange:
		return NewUniswapScraper(exchanges[dia.TethysExchange], relDB, datastore)
	case dia.HermesExchange:
		return NewUniswapScraper(exchanges[dia.HermesExchange], relDB, datastore)
	case dia.OmniDexExchange:
		return NewUniswapScraper(exchanges[dia.OmniDexExchange], relDB, datastore)
	case dia.DiffusionExchange:
		return NewUniswapScraper(exchanges[dia.DiffusionExchange], relDB, datastore)
	case dia.ApeswapExchange:
		return NewUniswapScraper(exchanges[dia.ApeswapExchange], relDB, datastore)
	case dia.BiswapExchange:
		return NewUniswapScraper(exchanges[dia.BiswapExchange], relDB, datastore)
	case dia.ArthswapExchange:
		return NewUniswapScraper(exchanges[dia.ArthswapExchange], relDB, datastore)
	case dia.StellaswapExchange:
		return NewUniswapScraper(exchanges[dia.StellaswapExchange], relDB, datastore)
	case dia.WanswapExchange:
		return NewUniswapScraper(exchanges[dia.WanswapExchange], relDB, datastore)

	case dia.CurveFIExchange:
		return NewCurveFIScraper(exchanges[dia.CurveFIExchange], relDB, datastore)
	case dia.CurveFIExchangePolygon:
		return NewCurveFIScraper(exchanges[dia.CurveFIExchangePolygon], relDB, datastore)
	case dia.CurveFIExchangeFantom:
		return NewCurveFIScraper(exchanges[dia.CurveFIExchangeFantom], relDB, datastore)
	case dia.CurveFIExchangeMoonbeam:
		return NewCurveFIScraper(exchanges[dia.CurveFIExchangeMoonbeam], relDB, datastore)
	case dia.CurveFIExchangeArbitrum:
		return NewCurveFIScraper(exchanges[dia.CurveFIExchangeArbitrum], relDB, datastore)

	case dia.BalancerV2Exchange:
		return NewBalancerV2Scraper(exchanges[dia.BalancerV2Exchange], relDB, datastore)
	case dia.BalancerV2ExchangeArbitrum:
		return NewBalancerV2Scraper(exchanges[dia.BalancerV2ExchangeArbitrum], relDB, datastore)
	case dia.BeetsExchange:
		return NewBalancerV2Scraper(exchanges[dia.BeetsExchange], relDB, datastore)
	case dia.BalancerV2ExchangePolygon:
		return NewBalancerV2Scraper(exchanges[dia.BalancerV2ExchangePolygon], relDB, datastore)

	case dia.UniswapExchangeV3:
		return NewUniswapV3Scraper(exchanges[dia.UniswapExchangeV3], relDB, datastore)
	case dia.UniswapExchangeV3Polygon:
		return NewUniswapV3Scraper(exchanges[dia.UniswapExchangeV3Polygon], relDB, datastore)
	case dia.UniswapExchangeV3Arbitrum:
		return NewUniswapV3Scraper(exchanges[dia.UniswapExchangeV3Arbitrum], relDB, datastore)

	case dia.BancorExchange:
		return NewBancorPoolScraper(exchanges[dia.BancorExchange], datastore)
	case dia.OrcaExchange:
		return NewOrcaScraper(exchanges[dia.OrcaExchange], datastore)
	case dia.OsmosisExchange:
		return NewOsmosisScraper(exchanges[dia.OsmosisExchange], relDB, datastore)
	case dia.PlatypusExchange:
		return NewPlatypusScraper(exchanges[dia.PlatypusExchange], datastore)
	case dia.VelodromeExchange:
		return NewVelodromePoolScraper(exchanges[dia.VelodromeExchange], relDB, datastore)
	case dia.MaverickExchange:
		return NewMaverickScraper(exchanges[dia.MaverickExchange], relDB, datastore)

	default:
		return nil
	}

}
