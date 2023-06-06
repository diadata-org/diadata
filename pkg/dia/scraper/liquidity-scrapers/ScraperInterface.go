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
		return NewUniswapScraper(exchanges[dia.UniswapExchange])
	case dia.SushiSwapExchange:
		return NewUniswapScraper(exchanges[dia.SushiSwapExchange])
	case dia.SushiSwapExchangePolygon:
		return NewUniswapScraper(exchanges[dia.SushiSwapExchangePolygon])
	case dia.SushiSwapExchangeFantom:
		return NewUniswapScraper(exchanges[dia.SushiSwapExchangeFantom])
	case dia.SushiSwapExchangeArbitrum:
		return NewUniswapScraper(exchanges[dia.SushiSwapExchangeArbitrum])
	case dia.CamelotExchange:
		return NewUniswapScraper(exchanges[dia.CamelotExchange])
	case dia.PanCakeSwap:
		return NewUniswapScraper(exchanges[dia.PanCakeSwap])
	case dia.DfynNetwork:
		return NewUniswapScraper(exchanges[dia.DfynNetwork])
	case dia.QuickswapExchange:
		return NewUniswapScraper(exchanges[dia.QuickswapExchange])
	case dia.UbeswapExchange:
		return NewUniswapScraper(exchanges[dia.UbeswapExchange])
	case dia.SpookyswapExchange:
		return NewUniswapScraper(exchanges[dia.SpookyswapExchange])
	case dia.SpiritswapExchange:
		return NewUniswapScraper(exchanges[dia.SpiritswapExchange])
	case dia.SolarbeamExchange:
		return NewUniswapScraper(exchanges[dia.SolarbeamExchange])
	case dia.TrisolarisExchange:
		return NewUniswapScraper(exchanges[dia.TrisolarisExchange])
	case dia.NetswapExchange:
		return NewUniswapScraper(exchanges[dia.NetswapExchange])
	case dia.HuckleberryExchange:
		return NewUniswapScraper(exchanges[dia.HuckleberryExchange])
	case dia.TraderJoeExchange:
		return NewUniswapScraper(exchanges[dia.TraderJoeExchange])
	case dia.PangolinExchange:
		return NewUniswapScraper(exchanges[dia.PangolinExchange])
	case dia.TethysExchange:
		return NewUniswapScraper(exchanges[dia.TethysExchange])
	case dia.HermesExchange:
		return NewUniswapScraper(exchanges[dia.HermesExchange])
	case dia.OmniDexExchange:
		return NewUniswapScraper(exchanges[dia.OmniDexExchange])
	case dia.DiffusionExchange:
		return NewUniswapScraper(exchanges[dia.DiffusionExchange])
	case dia.ApeswapExchange:
		return NewUniswapScraper(exchanges[dia.ApeswapExchange])
	case dia.BiswapExchange:
		return NewUniswapScraper(exchanges[dia.BiswapExchange])
	case dia.ArthswapExchange:
		return NewUniswapScraper(exchanges[dia.ArthswapExchange])
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
		return NewBalancerV2Scraper(exchanges[dia.BalancerV2Exchange])
	case dia.BalancerV2ExchangeArbitrum:
		return NewBalancerV2Scraper(exchanges[dia.BalancerV2ExchangeArbitrum])
	case dia.BeetsExchange:
		return NewBalancerV2Scraper(exchanges[dia.BeetsExchange])
	case dia.BalancerV2ExchangePolygon:
		return NewBalancerV2Scraper(exchanges[dia.BalancerV2ExchangePolygon])
	case dia.PlatypusExchange:
		return NewPlatypusScraper(exchanges[dia.PlatypusExchange])
	case dia.UniswapExchangeV3:
		return NewUniswapV3Scraper(exchanges[dia.UniswapExchangeV3])
	case dia.UniswapExchangeV3Polygon:
		return NewUniswapV3Scraper(exchanges[dia.UniswapExchangeV3Polygon])
	case dia.StellaswapExchange:
		return NewUniswapScraper(exchanges[dia.StellaswapExchange])
	case dia.WanswapExchange:
		return NewUniswapScraper(exchanges[dia.WanswapExchange])
	case dia.UniswapExchangeV3Arbitrum:
		return NewUniswapV3Scraper(exchanges[dia.UniswapExchangeV3Arbitrum])
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
