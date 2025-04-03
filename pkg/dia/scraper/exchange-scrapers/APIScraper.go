package scrapers

import (
	"io"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
)

// The collector kills a scraper after @watchdogDelayXXX seconds of inactivity
const (
// TODO use this with test
// watchdogDelay3Mins = 3 * 60
// watchdogDelayShort = 10 * 60
// watchdogDelay      = 20 * 60
// watchdogDelayLong  = 120 * 60
)

// empty type used for signaling
type nothing struct{}

var (
	Exchanges          = make(map[string]dia.Exchange)
	ExchangeDuplicates = make(map[string]dia.Exchange)
	Blockchains        map[string]dia.BlockChain
	chainConfigs       map[string]dia.ChainConfig
)

var evmID map[string]string

func init() {

	relDB, err := models.NewRelDataStore()
	if err != nil {
		log.Fatal("get rel datastore: ", err)
	}

	exchanges, err := relDB.GetAllExchanges()
	if err != nil {
		log.Fatal("get all exchanges: ", err)
	}
	for _, exchange := range exchanges {
		Exchanges[exchange.Name] = exchange
	}

	evmID = make(map[string]string)
	evmID["137"] = dia.POLYGON
	evmID["1"] = dia.ETHEREUM
	evmID["250"] = dia.FANTOM
	evmID["56"] = dia.BINANCESMARTCHAIN
	evmID["43114"] = dia.BINANCESMARTCHAIN
	evmID["1284"] = dia.MOONBEAM
	evmID["1285"] = dia.MOONRIVER
	evmID["42161"] = dia.ARBITRUM
	evmID["43114"] = dia.AVALANCHE

	chains, err := relDB.GetAllBlockchains(false)
	if err != nil {
		log.Fatal("get all chains: ", err)
	}
	Blockchains = make(map[string]dia.BlockChain)
	for _, chain := range chains {
		Blockchains[chain.Name] = chain
	}

	chainconfigurations, err := relDB.GetAllChainConfig()
	if err != nil {
		log.Fatal("get all chains: ", err)
	}
	chainConfigs = make(map[string]dia.ChainConfig)
	for _, chainconfig := range chainconfigurations {
		chainConfigs[chainconfig.ChainID] = chainconfig
	}

	ExchangeDuplicates[dia.Binance2Exchange] = dia.Exchange{Name: "Binance2", Centralized: true, WatchdogDelay: 300}
	ExchangeDuplicates[dia.BKEX2Exchange] = dia.Exchange{Name: "BKEX2", Centralized: true, WatchdogDelay: 1200}

}

// APIScraper provides common methods needed to get Trade information from
// exchange APIs.
type APIScraper interface {
	io.Closer
	// ScrapePair returns a PairScraper that continuously scrapes trades for a
	// single pair from this APIScraper
	ScrapePair(pair dia.ExchangePair) (PairScraper, error)
	// FetchAvailablePairs returns a list with all trading pairs available on
	// the exchange associated to the APIScraper. The format is such that it can
	// be used by the corr. pairScraper in order to fetch trades.
	FetchAvailablePairs() ([]dia.ExchangePair, error)

	// FillSymbolData collects information associated to the symbol ticker of an
	// asset traded on the exchange associated to the APIScraper.
	// Ideally, data is returned as close to original (blockchain) notation as possible.
	// This is only needed for CEX. For DEX the trade can be filled.
	FillSymbolData(symbol string) (dia.Asset, error)

	NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error)
	// Channel returns a channel that can be used to receive trades
	Channel() chan *dia.Trade
}

// PairScraper receives trades for a single pc.ExchangePair from a single exchange.
type PairScraper interface {
	io.Closer
	// Error returns an error when the channel Channel() is closed
	// and nil otherwise
	Error() error

	// Pair returns the pair this scraper is subscribed to
	Pair() dia.ExchangePair
}

// NewAPIScraper returns an API scraper for @exchange. If scrape==true it actually does
// scraping. Otherwise can be used for pairdiscovery.
func NewAPIScraper(exchange string, scrape bool, key string, secret string, relDB *models.RelDB) APIScraper {
	switch exchange {
	case dia.BitstampExchange:
		return NewBitstampScraper(Exchanges[dia.BitstampExchange], scrape, relDB)
	case dia.BinanceExchange:
		return NewBinanceScraper(key, secret, Exchanges[dia.BinanceExchange], Exchanges[dia.BinanceExchange].Name, scrape, relDB)
	case dia.Binance2Exchange:
		return NewBinanceScraper(key, secret, Exchanges[dia.BinanceExchange], dia.Binance2Exchange, scrape, relDB)
	case dia.BinanceExchangeUS:
		return NewBinanceScraperUS(key, secret, Exchanges[dia.BinanceExchangeUS], scrape, relDB)
	case dia.BitfinexExchange:
		return NewBitfinexScraper(key, secret, Exchanges[dia.BitfinexExchange], scrape, relDB)
	case dia.CoinBaseExchange:
		return NewCoinBaseScraper(Exchanges[dia.CoinBaseExchange], scrape, relDB)
	case dia.KrakenExchange:
		return NewKrakenScraper(key, secret, Exchanges[dia.KrakenExchange], scrape, relDB)
	case dia.OKExExchange:
		return NewOKExScraper(Exchanges[dia.OKExExchange], scrape, relDB)
	case dia.CryptoDotComExchange:
		return NewCryptoDotComScraper(Exchanges[dia.CryptoDotComExchange], scrape, relDB)
	case dia.HuobiExchange:
		return NewHuobiScraper(Exchanges[dia.HuobiExchange], scrape, relDB)
	case dia.GateIOExchange:
		return NewGateIOScraper(Exchanges[dia.GateIOExchange], scrape, relDB)
	case dia.BancorExchange:
		return NewBancorScraper(Exchanges[dia.BancorExchange], scrape)
	case dia.UniswapExchange:
		return NewUniswapScraper(Exchanges[dia.UniswapExchange], scrape, relDB)
	case dia.UniswapExchangeBase:
		return NewUniswapScraper(Exchanges[dia.UniswapExchangeBase], scrape, relDB)
	case dia.PanCakeSwap:
		return NewUniswapScraper(Exchanges[dia.PanCakeSwap], scrape, relDB)
	case dia.PanCakeSwapExchangeV3:
		return NewUniswapV3Scraper(Exchanges[dia.PanCakeSwapExchangeV3], scrape, relDB)
	case dia.SushiSwapExchange:
		return NewUniswapScraper(Exchanges[dia.SushiSwapExchange], scrape, relDB)
	case dia.SushiSwapExchangePolygon:
		return NewUniswapScraper(Exchanges[dia.SushiSwapExchangePolygon], scrape, relDB)
	case dia.SushiSwapExchangeArbitrum:
		return NewUniswapScraper(Exchanges[dia.SushiSwapExchangeArbitrum], scrape, relDB)
	case dia.SushiSwapExchangeFantom:
		return NewUniswapScraper(Exchanges[dia.SushiSwapExchangeFantom], scrape, relDB)
	case dia.CamelotExchange:
		return NewUniswapScraper(Exchanges[dia.CamelotExchange], scrape, relDB)
	case dia.CamelotExchangeV3:
		return NewUniswapV3Scraper(Exchanges[dia.CamelotExchangeV3], scrape, relDB)
	case dia.CurveFIExchange:
		return NewCurveFIScraper(Exchanges[dia.CurveFIExchange], scrape, relDB)
	case dia.CurveFIExchangeFantom:
		return NewCurveFIScraper(Exchanges[dia.CurveFIExchangeFantom], scrape, relDB)
	case dia.CurveFIExchangeMoonbeam:
		return NewCurveFIScraper(Exchanges[dia.CurveFIExchangeMoonbeam], scrape, relDB)
	case dia.CurveFIExchangePolygon:
		return NewCurveFIScraper(Exchanges[dia.CurveFIExchangePolygon], scrape, relDB)
	case dia.CurveFIExchangeArbitrum:
		return NewCurveFIScraper(Exchanges[dia.CurveFIExchangeArbitrum], scrape, relDB)
	case dia.BalancerV2Exchange:
		return NewBalancerV2Scraper(Exchanges[dia.BalancerV2Exchange], scrape, relDB)
	case dia.BalancerV2ExchangeArbitrum:
		return NewBalancerV2Scraper(Exchanges[dia.BalancerV2ExchangeArbitrum], scrape, relDB)
	case dia.BalancerV2ExchangePolygon:
		return NewBalancerV2Scraper(Exchanges[dia.BalancerV2ExchangePolygon], scrape, relDB)
	case dia.BeetsExchange:
		return NewBalancerV2Scraper(Exchanges[dia.BeetsExchange], scrape, relDB)
	case dia.KuCoinExchange:
		return NewKuCoinScraper(key, secret, Exchanges[dia.KuCoinExchange], scrape, relDB)
	case dia.BitMartExchange:
		return NewBitMartScraper(Exchanges[dia.BitMartExchange], scrape, relDB)
	case dia.BitMaxExchange:
		return NewBitMaxScraper(Exchanges[dia.BitMaxExchange], scrape, relDB)
	case dia.MEXCExchange:
		return NewMEXCScraper(Exchanges[dia.MEXCExchange], scrape, relDB)
	case dia.BKEXExchange:
		return NewBKEXScraper(Exchanges[dia.BKEXExchange], dia.BKEXExchange, scrape, relDB)
	case dia.BKEX2Exchange:
		return NewBKEXScraper(Exchanges[dia.BKEXExchange], dia.BKEX2Exchange, scrape, relDB)
	case dia.UniswapExchangeV3:
		return NewUniswapV3Scraper(Exchanges[dia.UniswapExchangeV3], scrape, relDB)
	case dia.UniswapExchangeV3Base:
		return NewUniswapV3Scraper(Exchanges[dia.UniswapExchangeV3Base], scrape, relDB)
	case dia.UniswapExchangeV3Celo:
		return NewUniswapV3Scraper(Exchanges[dia.UniswapExchangeV3Celo], scrape, relDB)
	case dia.DfynNetwork:
		return NewUniswapScraper(Exchanges[dia.DfynNetwork], scrape, relDB)
	case dia.UbeswapExchange:
		return NewUniswapScraper(Exchanges[dia.UbeswapExchange], scrape, relDB)
	case dia.NileV1Exchange:
		return NewUniswapScraper(Exchanges[dia.NileV1Exchange], scrape, relDB)
	case dia.RamsesV1Exchange:
		return NewUniswapScraper(Exchanges[dia.RamsesV1Exchange], scrape, relDB)
	case dia.NileV2Exchange:
		return NewUniswapV3Scraper(Exchanges[dia.NileV2Exchange], scrape, relDB)
	case dia.RamsesV2Exchange:
		return NewUniswapV3Scraper(Exchanges[dia.RamsesV2Exchange], scrape, relDB)
	case dia.UniswapExchangeV3Polygon:
		return NewUniswapV3Scraper(Exchanges[dia.UniswapExchangeV3Polygon], scrape, relDB)
	case dia.UniswapExchangeV3Arbitrum:
		return NewUniswapV3Scraper(Exchanges[dia.UniswapExchangeV3Arbitrum], scrape, relDB)
	case dia.UniswapExchangeV4:
		return NewUniswapV4Scraper(Exchanges[dia.UniswapExchangeV4], scrape, relDB)

	case dia.HuckleberryExchange:
		return NewUniswapScraper(Exchanges[dia.HuckleberryExchange], scrape, relDB)
	case dia.TraderJoeExchange:
		return NewUniswapScraper(Exchanges[dia.TraderJoeExchange], scrape, relDB)
	case dia.PangolinExchange:
		return NewUniswapScraper(Exchanges[dia.PangolinExchange], scrape, relDB)
	case dia.PlatypusExchange:
		return NewPlatypusScraper(Exchanges[dia.PlatypusExchange], scrape, relDB)
	case dia.SpookyswapExchange:
		return NewUniswapScraper(Exchanges[dia.SpookyswapExchange], scrape, relDB)
	case dia.QuickswapExchange:
		return NewUniswapScraper(Exchanges[dia.QuickswapExchange], scrape, relDB)
	case dia.SpiritswapExchange:
		return NewUniswapScraper(Exchanges[dia.SpiritswapExchange], scrape, relDB)
	case dia.SolarbeamExchange:
		return NewUniswapScraper(Exchanges[dia.SolarbeamExchange], scrape, relDB)
	case dia.TrisolarisExchange:
		return NewUniswapScraper(Exchanges[dia.TrisolarisExchange], scrape, relDB)
	case dia.ByBitExchange:
		return NewByBitScraper(Exchanges[dia.ByBitExchange], scrape, relDB)
	case dia.OrcaExchange:
		return NewOrcaScraper(Exchanges[dia.OrcaExchange], scrape)
	case dia.AnyswapExchange:
		return NewAnyswapScraper(Exchanges[dia.AnyswapExchange], scrape, relDB)
	case dia.NetswapExchange:
		return NewUniswapScraper(Exchanges[dia.NetswapExchange], scrape, relDB)
	case dia.BitMexExchange:
		return NewBitMexScraper(Exchanges[dia.BitMexExchange], scrape, relDB)
	case dia.TethysExchange:
		return NewUniswapScraper(Exchanges[dia.TethysExchange], scrape, relDB)
	case dia.HermesExchange:
		return NewUniswapScraper(Exchanges[dia.HermesExchange], scrape, relDB)
	case dia.OmniDexExchange:
		return NewUniswapScraper(Exchanges[dia.OmniDexExchange], scrape, relDB)
	case dia.DiffusionExchange:
		return NewUniswapScraper(Exchanges[dia.DiffusionExchange], scrape, relDB)
	case dia.ApeswapExchange:
		return NewUniswapScraper(Exchanges[dia.ApeswapExchange], scrape, relDB)
	case dia.BiswapExchange:
		return NewUniswapScraper(Exchanges[dia.BiswapExchange], scrape, relDB)
	case dia.ArthswapExchange:
		return NewUniswapScraper(Exchanges[dia.ArthswapExchange], scrape, relDB)
	case dia.StellaswapExchange:
		return NewUniswapScraper(Exchanges[dia.StellaswapExchange], scrape, relDB)
	case dia.WanswapExchange:
		return NewUniswapScraper(Exchanges[dia.WanswapExchange], scrape, relDB)
	case dia.TraderJoeExchangeV2_1:
		return NewTraderJoeScraper(Exchanges[dia.TraderJoeExchangeV2_1], scrape, relDB)
	case dia.TraderJoeExchangeV2_1Arbitrum:
		return NewTraderJoeScraper(Exchanges[dia.TraderJoeExchangeV2_1Arbitrum], scrape, relDB)
	case dia.TraderJoeExchangeV2_1Avalanche:
		return NewTraderJoeScraper(Exchanges[dia.TraderJoeExchangeV2_1Avalanche], scrape, relDB)
	case dia.TraderJoeExchangeV2_1BNB:
		return NewTraderJoeScraper(Exchanges[dia.TraderJoeExchangeV2_1BNB], scrape, relDB)
	case dia.TraderJoeExchangeV2_2Avalanche:
		return NewTraderJoeScraper(Exchanges[dia.TraderJoeExchangeV2_2Avalanche], scrape, relDB)
	case dia.OsmosisExchange:
		return NewOsmosisScraper(Exchanges[dia.OsmosisExchange], scrape, relDB)
	case dia.ZenlinkswapExchange:
		return NewZenlinkScraper(Exchanges[dia.ZenlinkswapExchange], scrape)
	case dia.ZenlinkswapExchangeBifrostPolkadot:
		return NewZenlinkScraper(Exchanges[dia.ZenlinkswapExchangeBifrostPolkadot], scrape)
	case dia.VelodromeExchange:
		return NewVelodromeScraper(Exchanges[dia.VelodromeExchange], scrape, relDB)
	case dia.VelodromeExchangeSwellchain:
		return NewVelodromeScraper(Exchanges[dia.VelodromeExchangeSwellchain], scrape, relDB)
	case dia.AerodromeV1Exchange:
		return NewVelodromeScraper(Exchanges[dia.AerodromeV1Exchange], scrape, relDB)
	case dia.MaverickExchange:
		return NewMaverickScraper(Exchanges[dia.MaverickExchange], scrape, relDB)
	case dia.PearlfiExchangeTestnet:
		return NewUniswapV3Scraper(Exchanges[dia.PearlfiExchangeTestnet], scrape, relDB)
	case dia.PearlfiExchange:
		return NewUniswapV3Scraper(Exchanges[dia.PearlfiExchange], scrape, relDB)
	case dia.PearlfiStableswapExchange:
		return NewUniswapScraper(Exchanges[dia.PearlfiStableswapExchange], scrape, relDB)
	case dia.ThenaExchange:
		return NewUniswapScraper(Exchanges[dia.ThenaExchange], scrape, relDB)
	case dia.ThenaV3Exchange:
		return NewUniswapV3Scraper(Exchanges[dia.ThenaV3Exchange], scrape, relDB)
	// case dia.FinageForex:
	// 	return NewFinageForexScraper(Exchanges[dia.FinageForex], scrape, relDB, key, secret)
	case dia.BitgetExchange:
		return NewBitgetScraper(Exchanges[dia.BitgetExchange], scrape, relDB)
	case dia.AerodromeSlipstreamExchange:
		return NewUniswapV3Scraper(Exchanges[dia.AerodromeSlipstreamExchange], scrape, relDB)
	case dia.VelodromeSlipstreamExchange:
		return NewUniswapV3Scraper(Exchanges[dia.VelodromeSlipstreamExchange], scrape, relDB)
		// case dia.FinageForex:
		// 	return NewFinageForexScraper(Exchanges[dia.FinageForex], scrape, relDB, key, secret)

	case dia.MultiChain:
		return NewBridgeSwapScraper(Exchanges[dia.MultiChain], scrape, relDB)

	case "Influx":
		return NewInfluxScraper(scrape)

	case "UniswapHistory":
		return NewUniswapHistoryScraper(Exchanges[dia.UniswapExchange], scrape, relDB)
	case dia.AyinExchange:
		return NewAyinScraper(Exchanges[dia.AyinExchange], scrape, relDB)
	case dia.BitflowExchange:
		return NewBitflowScraper(Exchanges[dia.BitflowExchange], scrape, relDB)
	case dia.VelarExchange:
		return NewVelarScraper(Exchanges[dia.VelarExchange], scrape, relDB)
	case dia.BifrostExchange:
		return NewBifrostScraper(Exchanges[dia.BifrostExchange], scrape, relDB)
	case dia.HydrationExchange:
		return NewHydrationScraper(Exchanges[dia.HydrationExchange], scrape, relDB)
	default:
		return nil
	}

}
