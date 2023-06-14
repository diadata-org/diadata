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

var (
	log       *logrus.Logger
	exchanges map[string]dia.Exchange
)

func init() {
	log = logrus.New()
	exchanges = scrapers.Exchanges

}

// NewLiquidityScraper returns a liquidity scraper for @source.
func NewLiquidityScraper(source string, relDB *models.RelDB) LiquidityScraper {
	switch source {
	case dia.UniswapExchange:
		return NewUniswapScraper(exchanges[dia.UniswapExchange], relDB)
	case dia.SushiSwapExchange:
		return NewUniswapScraper(exchanges[dia.SushiSwapExchange], relDB)
	case dia.SushiSwapExchangePolygon:
		return NewUniswapScraper(exchanges[dia.SushiSwapExchangePolygon], relDB)
	case dia.SushiSwapExchangeFantom:
		return NewUniswapScraper(exchanges[dia.SushiSwapExchangeFantom], relDB)
	case dia.SushiSwapExchangeArbitrum:
		return NewUniswapScraper(exchanges[dia.SushiSwapExchangeArbitrum], relDB)
	case dia.CamelotExchange:
		return NewUniswapScraper(exchanges[dia.CamelotExchange], relDB)
	case dia.PanCakeSwap:
		return NewUniswapScraper(exchanges[dia.PanCakeSwap], relDB)
	case dia.DfynNetwork:
		return NewUniswapScraper(exchanges[dia.DfynNetwork], relDB)
	case dia.QuickswapExchange:
		return NewUniswapScraper(exchanges[dia.QuickswapExchange], relDB)
	case dia.UbeswapExchange:
		return NewUniswapScraper(exchanges[dia.UbeswapExchange], relDB)
	case dia.SpookyswapExchange:
		return NewUniswapScraper(exchanges[dia.SpookyswapExchange], relDB)
	case dia.SpiritswapExchange:
		return NewUniswapScraper(exchanges[dia.SpiritswapExchange], relDB)
	case dia.SolarbeamExchange:
		return NewUniswapScraper(exchanges[dia.SolarbeamExchange], relDB)
	case dia.TrisolarisExchange:
		return NewUniswapScraper(exchanges[dia.TrisolarisExchange], relDB)
	case dia.NetswapExchange:
		return NewUniswapScraper(exchanges[dia.NetswapExchange], relDB)
	case dia.HuckleberryExchange:
		return NewUniswapScraper(exchanges[dia.HuckleberryExchange], relDB)
	case dia.TraderJoeExchange:
		return NewUniswapScraper(exchanges[dia.TraderJoeExchange], relDB)
	case dia.PangolinExchange:
		return NewUniswapScraper(exchanges[dia.PangolinExchange], relDB)
	case dia.TethysExchange:
		return NewUniswapScraper(exchanges[dia.TethysExchange], relDB)
	case dia.HermesExchange:
		return NewUniswapScraper(exchanges[dia.HermesExchange], relDB)
	case dia.OmniDexExchange:
		return NewUniswapScraper(exchanges[dia.OmniDexExchange], relDB)
	case dia.DiffusionExchange:
		return NewUniswapScraper(exchanges[dia.DiffusionExchange], relDB)
	case dia.ApeswapExchange:
		return NewUniswapScraper(exchanges[dia.ApeswapExchange], relDB)
	case dia.BiswapExchange:
		return NewUniswapScraper(exchanges[dia.BiswapExchange], relDB)
	case dia.ArthswapExchange:
		return NewUniswapScraper(exchanges[dia.ArthswapExchange], relDB)
	case dia.StellaswapExchange:
		return NewUniswapScraper(exchanges[dia.StellaswapExchange], relDB)
	case dia.WanswapExchange:
		return NewUniswapScraper(exchanges[dia.WanswapExchange], relDB)

	case dia.CurveFIExchange:
		return NewCurveFIScraper(exchanges[dia.CurveFIExchange])
	case dia.CurveFIExchangePolygon:
		return NewCurveFIScraper(exchanges[dia.CurveFIExchangePolygon])
	case dia.CurveFIExchangeFantom:
		return NewCurveFIScraper(exchanges[dia.CurveFIExchangeFantom])
	case dia.CurveFIExchangeMoonbeam:
		return NewCurveFIScraper(exchanges[dia.CurveFIExchangeMoonbeam])
	case dia.CurveFIExchangeArbitrum:
		return NewCurveFIScraper(exchanges[dia.CurveFIExchangeArbitrum])

	case dia.BalancerV2Exchange:
		return NewBalancerV2Scraper(exchanges[dia.BalancerV2Exchange], relDB)
	case dia.BalancerV2ExchangeArbitrum:
		return NewBalancerV2Scraper(exchanges[dia.BalancerV2ExchangeArbitrum], relDB)
	case dia.BeetsExchange:
		return NewBalancerV2Scraper(exchanges[dia.BeetsExchange], relDB)
	case dia.BalancerV2ExchangePolygon:
		return NewBalancerV2Scraper(exchanges[dia.BalancerV2ExchangePolygon], relDB)

	case dia.PlatypusExchange:
		return NewPlatypusScraper(exchanges[dia.PlatypusExchange])

	case dia.UniswapExchangeV3:
		return NewUniswapV3Scraper(exchanges[dia.UniswapExchangeV3], relDB)
	case dia.UniswapExchangeV3Polygon:
		return NewUniswapV3Scraper(exchanges[dia.UniswapExchangeV3Polygon], relDB)
	case dia.UniswapExchangeV3Arbitrum:
		return NewUniswapV3Scraper(exchanges[dia.UniswapExchangeV3Arbitrum], relDB)
	case dia.BancorExchange:
		return NewBancorPoolScraper(exchanges[dia.BancorExchange])
	case dia.OrcaExchange:
		return NewOrcaScraper(exchanges[dia.OrcaExchange])
	case dia.OsmosisExchange:
		return NewOsmosisScraper(exchanges[dia.OsmosisExchange], relDB)

	default:
		return nil
	}

}
