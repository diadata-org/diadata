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

var Exchanges = make(map[string]dia.Exchange)

// var blockchains map[string]dia.BlockChain

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

	// blockchains = make(map[string]dia.BlockChain)
	// blockchains[dia.BITCOIN] = dia.BlockChain{Name: dia.BITCOIN, NativeToken: dia.Asset{Symbol: "BTC"}, VerificationMechanism: dia.PROOF_OF_WORK}
	// blockchains[dia.ETHEREUM] = dia.BlockChain{Name: dia.ETHEREUM, NativeToken: dia.Asset{Symbol: "ETH"}, VerificationMechanism: dia.PROOF_OF_WORK}
	// blockchains[dia.BINANCESMARTCHAIN] = dia.BlockChain{Name: dia.BINANCESMARTCHAIN, NativeToken: dia.Asset{Symbol: "BNB"}, VerificationMechanism: dia.PROOF_OF_STAKE}
	// blockchains[dia.POLYGON] = dia.BlockChain{Name: dia.POLYGON, NativeToken: dia.Asset{Symbol: "MATIC"}, VerificationMechanism: dia.PROOF_OF_STAKE}
	// blockchains[dia.CELO] = dia.BlockChain{Name: dia.CELO, NativeToken: dia.Asset{Symbol: "CELO"}, VerificationMechanism: dia.PROOF_OF_STAKE}
	// blockchains[dia.FANTOM] = dia.BlockChain{Name: dia.FANTOM, NativeToken: dia.Asset{Symbol: "FTM"}, VerificationMechanism: dia.PROOF_OF_STAKE}
	// blockchains[dia.NEAR] = dia.BlockChain{Name: dia.NEAR, NativeToken: dia.Asset{Symbol: "NEAR"}, VerificationMechanism: dia.PROOF_OF_STAKE}
	// blockchains[dia.AURORA] = dia.BlockChain{Name: dia.AURORA, NativeToken: dia.Asset{Symbol: "AURORA"}, VerificationMechanism: dia.PROOF_OF_STAKE}
	// blockchains[dia.SOLANA] = dia.BlockChain{Name: dia.SOLANA, NativeToken: dia.Asset{Symbol: "SOL"}, VerificationMechanism: dia.PROOF_OF_STAKE}
	// blockchains[dia.FLOW] = dia.BlockChain{Name: dia.FLOW, NativeToken: dia.Asset{Symbol: "FLOW"}, VerificationMechanism: dia.PROOF_OF_STAKE}
	// blockchains[dia.MOONRIVER] = dia.BlockChain{Name: dia.MOONRIVER, NativeToken: dia.Asset{Symbol: "MOVR"}, VerificationMechanism: dia.PROOF_OF_STAKE}
	// blockchains[dia.MOONBEAM] = dia.BlockChain{Name: dia.MOONBEAM, NativeToken: dia.Asset{Symbol: "GLMR"}, VerificationMechanism: dia.PROOF_OF_STAKE}
	// blockchains[dia.AVALANCHE] = dia.BlockChain{Name: dia.AVALANCHE, NativeToken: dia.Asset{Symbol: "AVAX"}, VerificationMechanism: dia.PROOF_OF_STAKE}
	// blockchains[dia.ASTAR] = dia.BlockChain{Name: dia.ASTAR, NativeToken: dia.Asset{Symbol: "ASTR"}, VerificationMechanism: dia.PROOF_OF_STAKE}
	// blockchains[dia.SHIDEN] = dia.BlockChain{Name: dia.SHIDEN, NativeToken: dia.Asset{Symbol: "SDN"}, VerificationMechanism: dia.PROOF_OF_STAKE}
	// blockchains[dia.ARBITRUM] = dia.BlockChain{Name: dia.ARBITRUM, NativeToken: dia.Asset{Symbol: "ARB"}, VerificationMechanism: dia.PROOF_OF_STAKE}
	// blockchains[dia.METIS] = dia.BlockChain{Name: dia.METIS, NativeToken: dia.Asset{Symbol: "Metis"}, VerificationMechanism: dia.PROOF_OF_STAKE}
	// blockchains[dia.KILT] = dia.BlockChain{Name: dia.KILT, NativeToken: dia.Asset{Symbol: "KILT"}, VerificationMechanism: dia.PROOF_OF_STAKE}
	// blockchains[dia.FETCH] = dia.BlockChain{Name: dia.FETCH, NativeToken: dia.Asset{Symbol: "FET"}, VerificationMechanism: dia.PROOF_OF_STAKE}
	// blockchains[dia.FUSE] = dia.BlockChain{Name: dia.FUSE, NativeToken: dia.Asset{Symbol: "FUSE"}, VerificationMechanism: dia.PROOF_OF_STAKE}
	// blockchains[dia.TELOS] = dia.BlockChain{Name: dia.TELOS, NativeToken: dia.Asset{Symbol: "TLOS"}, VerificationMechanism: dia.PROOF_OF_STAKE}
	// blockchains[dia.EVMOS] = dia.BlockChain{Name: dia.EVMOS, NativeToken: dia.Asset{Symbol: "EVMOS"}, VerificationMechanism: dia.PROOF_OF_STAKE}
	// blockchains[dia.KUSAMA] = dia.BlockChain{Name: dia.KUSAMA, NativeToken: dia.Asset{Symbol: "KSM"}, VerificationMechanism: dia.PROOF_OF_STAKE}
	// blockchains[dia.ACALA] = dia.BlockChain{Name: dia.ACALA, NativeToken: dia.Asset{Symbol: "ACL"}, VerificationMechanism: dia.PROOF_OF_STAKE}
	// blockchains[dia.POLKADOT] = dia.BlockChain{Name: dia.POLKADOT, NativeToken: dia.Asset{Symbol: "DOT"}, VerificationMechanism: dia.PROOF_OF_STAKE}
	// blockchains[dia.FIAT] = dia.BlockChain{Name: dia.FIAT}

	// Exchanges = make(map[string]dia.Exchange)
	// Exchanges[dia.BalancerExchange] = dia.Exchange{Name: dia.BalancerExchange, Centralized: false, Contract: "0x9424B1412450D0f8Fc2255FAf6046b98213B76Bd", WatchdogDelay: watchdogDelay}
	// Exchanges[dia.BalancerV2Exchange] = dia.Exchange{Name: dia.BalancerV2Exchange, Centralized: false, BlockChain: blockchains[dia.ETHEREUM], Contract: "0xBA12222222228d8Ba445958a75a0704d566BF2C8", WatchdogDelay: watchdogDelayLong}
	// Exchanges[dia.BeetsExchange] = dia.Exchange{Name: dia.BeetsExchange, Centralized: false, BlockChain: blockchains[dia.FANTOM], Contract: "0x20dd72Ed959b6147912C2e529F0a0C651c33c9ce", WatchdogDelay: watchdogDelayLong}
	// Exchanges[dia.BinanceExchange] = dia.Exchange{Name: dia.BinanceExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	// Exchanges[dia.BinanceExchangeUS] = dia.Exchange{Name: dia.BinanceExchangeUS, Centralized: true, WatchdogDelay: watchdogDelay}
	// Exchanges[dia.KrakenExchange] = dia.Exchange{Name: dia.KrakenExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	// Exchanges[dia.CREX24Exchange] = dia.Exchange{Name: dia.CREX24Exchange, Centralized: true, WatchdogDelay: watchdogDelay}
	// Exchanges[dia.BitfinexExchange] = dia.Exchange{Name: dia.BitfinexExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	// Exchanges[dia.BitBayExchange] = dia.Exchange{Name: dia.BitBayExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	// Exchanges[dia.BittrexExchange] = dia.Exchange{Name: dia.BittrexExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	// Exchanges[dia.BitforexExchange] = dia.Exchange{Name: dia.BitforexExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	// Exchanges[dia.CoinBaseExchange] = dia.Exchange{Name: dia.CoinBaseExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	// Exchanges[dia.HitBTCExchange] = dia.Exchange{Name: dia.HitBTCExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	// Exchanges[dia.SimexExchange] = dia.Exchange{Name: dia.SimexExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	// Exchanges[dia.OKExExchange] = dia.Exchange{Name: dia.OKExExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	// Exchanges[dia.FTXExchange] = dia.Exchange{Name: dia.FTXExchange, Centralized: true, WatchdogDelay: watchdogDelay3Mins}
	// Exchanges[dia.CryptoDotComExchange] = dia.Exchange{Name: dia.CryptoDotComExchange, Centralized: true, WatchdogDelay: watchdogDelay3Mins}
	// Exchanges[dia.HuobiExchange] = dia.Exchange{Name: dia.HuobiExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	// Exchanges[dia.LBankExchange] = dia.Exchange{Name: dia.LBankExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	// Exchanges[dia.GateIOExchange] = dia.Exchange{Name: dia.GateIOExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	// Exchanges[dia.ZBExchange] = dia.Exchange{Name: dia.ZBExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	// Exchanges[dia.QuoineExchange] = dia.Exchange{Name: dia.QuoineExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	// Exchanges[dia.UnknownExchange] = dia.Exchange{Name: dia.UnknownExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	// Exchanges[dia.BancorExchange] = dia.Exchange{Name: dia.BancorExchange, Centralized: false, BlockChain: blockchains[dia.ETHEREUM], WatchdogDelay: watchdogDelayLong} //API is used instead of contracts
	// Exchanges[dia.UniswapExchange] = dia.Exchange{Name: dia.UniswapExchange, Centralized: false, BlockChain: blockchains[dia.ETHEREUM], Contract: "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f", WatchdogDelay: watchdogDelay}
	// Exchanges[dia.UniswapExchangeV3] = dia.Exchange{Name: dia.UniswapExchangeV3, Centralized: false, BlockChain: blockchains[dia.ETHEREUM], Contract: "0x1F98431c8aD98523631AE4a59f267346ea31F984", WatchdogDelay: watchdogDelay}
	// Exchanges[dia.LoopringExchange] = dia.Exchange{Name: dia.LoopringExchange, Centralized: true, BlockChain: blockchains[dia.ETHEREUM], WatchdogDelay: watchdogDelay} //API is used instead of contracts
	// Exchanges[dia.CurveFIExchange] = dia.Exchange{Name: dia.CurveFIExchange, Centralized: false, BlockChain: blockchains[dia.ETHEREUM], Contract: "0x90E00ACe148ca3b23Ac1bC8C240C2a7Dd9c2d7f5", WatchdogDelay: watchdogDelayLong}
	// Exchanges[dia.MakerExchange] = dia.Exchange{Name: dia.MakerExchange, Centralized: false, BlockChain: blockchains[dia.ETHEREUM], WatchdogDelay: watchdogDelay} //API is used instead of contracts
	// Exchanges[dia.KuCoinExchange] = dia.Exchange{Name: dia.KuCoinExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	// Exchanges[dia.SushiSwapExchange] = dia.Exchange{Name: dia.SushiSwapExchange, Centralized: false, BlockChain: blockchains[dia.ETHEREUM], Contract: "0xc0aee478e3658e2610c5f7a4a2e1777ce9e4f2ac", WatchdogDelay: watchdogDelay}
	// Exchanges[dia.PanCakeSwap] = dia.Exchange{Name: dia.PanCakeSwap, Centralized: false, BlockChain: blockchains[dia.BINANCESMARTCHAIN], Contract: "0xcA143Ce32Fe78f1f7019d7d551a6402fC5350c73", WatchdogDelay: watchdogDelayLong}
	// Exchanges[dia.ApeswapExchange] = dia.Exchange{Name: dia.ApeswapExchange, Centralized: false, BlockChain: blockchains[dia.BINANCESMARTCHAIN], Contract: "0x0841BD0B734E4F5853f0dD8d7Ea041c241fb0Da6", WatchdogDelay: watchdogDelayLong}
	// Exchanges[dia.BiswapExchange] = dia.Exchange{Name: dia.BiswapExchange, Centralized: false, BlockChain: blockchains[dia.BINANCESMARTCHAIN], Contract: "0x858E3312ed3A876947EA49d572A7C42DE08af7EE", WatchdogDelay: watchdogDelayLong}
	// Exchanges[dia.DforceExchange] = dia.Exchange{Name: dia.DforceExchange, Centralized: false, BlockChain: blockchains[dia.ETHEREUM], Contract: "0x03eF3f37856bD08eb47E2dE7ABc4Ddd2c19B60F2", WatchdogDelay: watchdogDelayLong}
	// Exchanges[dia.ZeroxExchange] = dia.Exchange{Name: dia.ZeroxExchange, Centralized: true, WatchdogDelay: watchdogDelayLong}
	// Exchanges[dia.KyberExchange] = dia.Exchange{Name: dia.KyberExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	// Exchanges[dia.BitMaxExchange] = dia.Exchange{Name: dia.BitMaxExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	// Exchanges[dia.STEXExchange] = dia.Exchange{Name: dia.STEXExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	// Exchanges[dia.DfynNetwork] = dia.Exchange{Name: dia.DfynNetwork, Centralized: false, BlockChain: blockchains[dia.POLYGON], Contract: "0xe7fb3e833efe5f9c441105eb65ef8b261266423b", WatchdogDelay: watchdogDelayLong}
	// Exchanges[dia.QuickswapExchange] = dia.Exchange{Name: dia.QuickswapExchange, Centralized: false, BlockChain: blockchains[dia.POLYGON], Contract: "0x5757371414417b8C6CAad45bAeF941aBc7d3Ab32", WatchdogDelay: watchdogDelayLong}
	// Exchanges[dia.UbeswapExchange] = dia.Exchange{Name: dia.UbeswapExchange, Centralized: false, BlockChain: blockchains[dia.CELO], Contract: "0x62d5b84bE28a183aBB507E125B384122D2C25fAE", WatchdogDelay: watchdogDelayLong}
	// Exchanges[dia.SpookyswapExchange] = dia.Exchange{Name: dia.SpookyswapExchange, Centralized: false, BlockChain: blockchains[dia.FANTOM], Contract: "0x152ee697f2e276fa89e96742e9bb9ab1f2e61be3", WatchdogDelay: watchdogDelayLong}
	// Exchanges[dia.SpiritswapExchange] = dia.Exchange{Name: dia.SpiritswapExchange, Centralized: false, BlockChain: blockchains[dia.FANTOM], Contract: "0xef45d134b73241eda7703fa787148d9c9f4950b0", WatchdogDelay: watchdogDelayLong}
	// Exchanges[dia.SerumExchange] = dia.Exchange{Name: dia.SerumExchange, Centralized: false, BlockChain: blockchains[dia.SOLANA], Contract: "", WatchdogDelay: watchdogDelayLong}
	// Exchanges[dia.SolarbeamExchange] = dia.Exchange{Name: dia.SolarbeamExchange, Centralized: false, BlockChain: blockchains[dia.MOONRIVER], Contract: "0x049581aEB6Fe262727f290165C29BDAB065a1B68", WatchdogDelay: watchdogDelay3Mins}
	// Exchanges[dia.HuckleberryExchange] = dia.Exchange{Name: dia.HuckleberryExchange, Centralized: false, BlockChain: blockchains[dia.MOONRIVER], Contract: "0x017603C8f29F7f6394737628a93c57ffBA1b7256", WatchdogDelay: watchdogDelay}
	// Exchanges[dia.TraderJoeExchange] = dia.Exchange{Name: dia.TraderJoeExchange, Centralized: false, BlockChain: blockchains[dia.AVALANCHE], Contract: "0x9Ad6C38BE94206cA50bb0d90783181662f0Cfa10", WatchdogDelay: watchdogDelay}
	// Exchanges[dia.PangolinExchange] = dia.Exchange{Name: dia.PangolinExchange, Centralized: false, BlockChain: blockchains[dia.AVALANCHE], Contract: "0xefa94DE7a4656D787667C749f7E1223D71E9FD88", WatchdogDelay: watchdogDelay}
	// Exchanges[dia.TrisolarisExchange] = dia.Exchange{Name: dia.TrisolarisExchange, Centralized: false, BlockChain: blockchains[dia.AURORA], Contract: "0xc66F594268041dB60507F00703b152492fb176E7", WatchdogDelay: watchdogDelay}
	// Exchanges[dia.ByBitExchange] = dia.Exchange{Name: dia.ByBitExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	// Exchanges[dia.AnyswapExchange] = dia.Exchange{Name: dia.AnyswapExchange, Centralized: false, BlockChain: blockchains[dia.ETHEREUM], Contract: "0x6b7a87899490EcE95443e979cA9485CBE7E71522", WatchdogDelay: watchdogDelayLong}
	// Exchanges[dia.NetswapExchange] = dia.Exchange{Name: dia.NetswapExchange, Centralized: false, BlockChain: blockchains[dia.METIS], Contract: "0x70f51d68D16e8f9e418441280342BD43AC9Dff9f", WatchdogDelay: watchdogDelayLong}
	// Exchanges[dia.TethysExchange] = dia.Exchange{Name: dia.TethysExchange, Centralized: false, BlockChain: blockchains[dia.METIS], Contract: "0x2CdFB20205701FF01689461610C9F321D1d00F80", WatchdogDelay: watchdogDelayLong}
	// Exchanges[dia.HermesExchange] = dia.Exchange{Name: dia.HermesExchange, Centralized: false, BlockChain: blockchains[dia.METIS], Contract: "0x633a093C9e94f64500FC8fCBB48e90dd52F6668F", WatchdogDelay: watchdogDelayLong}
	// Exchanges[dia.OmniDexExchange] = dia.Exchange{Name: dia.OmniDexExchange, Centralized: false, BlockChain: blockchains[dia.TELOS], Contract: "0x7a2A35706f5d1CeE2faa8A254dd6F6D7d7Becc25", WatchdogDelay: watchdogDelayLong}
	// Exchanges[dia.DiffusionExchange] = dia.Exchange{Name: dia.DiffusionExchange, Centralized: false, BlockChain: blockchains[dia.EVMOS], Contract: "0x6aBdDa34Fb225be4610a2d153845e09429523Cd2", WatchdogDelay: watchdogDelayLong}
	// Exchanges[dia.ArthswapExchange] = dia.Exchange{Name: dia.ArthswapExchange, Centralized: false, BlockChain: blockchains[dia.ASTAR], Contract: "0xA9473608514457b4bF083f9045fA63ae5810A03E", WatchdogDelay: watchdogDelayLong}
	// Exchanges[dia.BitMexExchange] = dia.Exchange{Name: dia.BitMexExchange, Centralized: true, WatchdogDelay: watchdogDelay}

	// Exchanges[dia.SushiSwapExchangePolygon] = dia.Exchange{Name: dia.SushiSwapExchangePolygon, Centralized: false, BlockChain: blockchains[dia.POLYGON], Contract: "0xc35dadb65012ec5796536bd9864ed8773abc74c4", WatchdogDelay: watchdogDelay}
	// Exchanges[dia.UniswapExchangeV3Polygon] = dia.Exchange{Name: dia.UniswapExchangeV3Polygon, Centralized: false, BlockChain: blockchains[dia.POLYGON], Contract: "0x1F98431c8aD98523631AE4a59f267346ea31F984", WatchdogDelay: watchdogDelayLong}

	// Exchanges[dia.SushiSwapExchangeFantom] = dia.Exchange{Name: dia.SushiSwapExchangeFantom, Centralized: false, BlockChain: blockchains[dia.FANTOM], Contract: "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506", WatchdogDelay: watchdogDelay}
	// Exchanges[dia.SushiSwapExchangeArbitrum] = dia.Exchange{Name: dia.SushiSwapExchangeArbitrum, Centralized: false, BlockChain: blockchains[dia.ARBITRUM], Contract: "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506", WatchdogDelay: watchdogDelay}

	// // Exchanges[dia.FinageForex] = dia.Exchange{Name: dia.FinageForex, Centralized: true, BlockChain: blockchains[dia.FIAT], WatchdogDelay: watchdogDelay}
	// Exchanges["Influx"] = dia.Exchange{Name: "Influx", WatchdogDelay: 360000}
	// Exchanges["UniswapHistory"] = dia.Exchange{Name: "UniswapHistory", Centralized: false, BlockChain: blockchains[dia.ETHEREUM], Contract: "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f", WatchdogDelay: 3600}
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
	case dia.BinanceExchange:
		return NewBinanceScraper(key, secret, Exchanges[dia.BinanceExchange], scrape, relDB)
	case dia.BinanceExchangeUS:
		return NewBinanceScraperUS(key, secret, Exchanges[dia.BinanceExchangeUS], scrape, relDB)
	case dia.BitBayExchange:
		return NewBitBayScraper(Exchanges[dia.BitBayExchange], scrape, relDB)
	case dia.BitfinexExchange:
		return NewBitfinexScraper(key, secret, Exchanges[dia.BitfinexExchange], scrape, relDB)
	case dia.BitforexExchange:
		return NewBitforexScraper(Exchanges[dia.BitforexExchange], scrape, relDB)
	case dia.BittrexExchange:
		return NewBittrexScraper(Exchanges[dia.BittrexExchange], scrape, relDB)
	case dia.CoinBaseExchange:
		return NewCoinBaseScraper(Exchanges[dia.CoinBaseExchange], scrape, relDB)
	case dia.CREX24Exchange:
		return NewCREX24Scraper(Exchanges[dia.CREX24Exchange], relDB)
	case dia.KrakenExchange:
		return NewKrakenScraper(key, secret, Exchanges[dia.KrakenExchange], scrape, relDB)
	case dia.HitBTCExchange:
		return NewHitBTCScraper(Exchanges[dia.HitBTCExchange], scrape, relDB)
	case dia.SimexExchange:
		return NewSimexScraper(Exchanges[dia.SimexExchange], scrape, relDB)
	case dia.OKExExchange:
		return NewOKExScraper(Exchanges[dia.OKExExchange], scrape, relDB)
	case dia.CryptoDotComExchange:
		return NewCryptoDotComScraper(Exchanges[dia.CryptoDotComExchange], scrape, relDB)
	case dia.FTXExchange:
		return NewFTXScraper(Exchanges[dia.FTXExchange], scrape, relDB)
	case dia.HuobiExchange:
		return NewHuobiScraper(Exchanges[dia.HuobiExchange], scrape, relDB)
	case dia.LBankExchange:
		return NewLBankScraper(Exchanges[dia.LBankExchange], scrape, relDB)
	case dia.GateIOExchange:
		return NewGateIOScraper(Exchanges[dia.GateIOExchange], scrape, relDB)
	case dia.ZBExchange:
		return NewZBScraper(Exchanges[dia.ZBExchange], scrape, relDB)
	case dia.QuoineExchange:
		return NewQuoineScraper(Exchanges[dia.QuoineExchange], scrape, relDB)
	case dia.BancorExchange:
		return NewBancorScraper(Exchanges[dia.BancorExchange], scrape)
	case dia.UniswapExchange:
		return NewUniswapScraper(Exchanges[dia.UniswapExchange], scrape)
	case dia.PanCakeSwap:
		return NewUniswapScraper(Exchanges[dia.PanCakeSwap], scrape)
	case dia.SushiSwapExchange:
		return NewUniswapScraper(Exchanges[dia.SushiSwapExchange], scrape)
	case dia.LoopringExchange:
		return NewLoopringScraper(Exchanges[dia.LoopringExchange], scrape, relDB)
	case dia.CurveFIExchange:
		return NewCurveFIScraper(Exchanges[dia.CurveFIExchange], scrape)
	case dia.BalancerExchange:
		return NewBalancerScraper(Exchanges[dia.BalancerExchange], scrape)
	case dia.BalancerV2Exchange:
		return NewBalancerV2Scraper(Exchanges[dia.BalancerV2Exchange], scrape)
	case dia.BeetsExchange:
		return NewBalancerV2Scraper(Exchanges[dia.BeetsExchange], scrape)
	case dia.MakerExchange:
		return NewMakerScraper(Exchanges[dia.MakerExchange], scrape, relDB)
	case dia.KuCoinExchange:
		return NewKuCoinScraper(key, secret, Exchanges[dia.KuCoinExchange], scrape, relDB)
	case dia.DforceExchange:
		return NewDforceScraper(Exchanges[dia.DforceExchange], scrape)
	case dia.ZeroxExchange:
		return NewZeroxScraper(Exchanges[dia.ZeroxExchange], scrape)
	case dia.KyberExchange:
		return NewKyberScraper(Exchanges[dia.KyberExchange], scrape)
	case dia.BitMaxExchange:
		return NewBitMaxScraper(Exchanges[dia.BitMaxExchange], scrape, relDB)
	case dia.STEXExchange:
		return NewSTEXScraper(Exchanges[dia.STEXExchange], scrape, relDB)
	case dia.UniswapExchangeV3:
		return NewUniswapV3Scraper(Exchanges[dia.UniswapExchangeV3], scrape)
	case dia.DfynNetwork:
		return NewUniswapScraper(Exchanges[dia.DfynNetwork], scrape)
	case dia.UbeswapExchange:
		return NewUniswapScraper(Exchanges[dia.UbeswapExchange], scrape)
	case dia.SushiSwapExchangePolygon:
		return NewUniswapScraper(Exchanges[dia.SushiSwapExchangePolygon], scrape)

	case dia.UniswapExchangeV3Polygon:
		return NewUniswapV3Scraper(Exchanges[dia.UniswapExchangeV3Polygon], scrape)
	case dia.HuckleberryExchange:
		return NewUniswapScraper(Exchanges[dia.HuckleberryExchange], scrape)
	case dia.TraderJoeExchange:
		return NewUniswapScraper(Exchanges[dia.TraderJoeExchange], scrape)
	case dia.PangolinExchange:
		return NewUniswapScraper(Exchanges[dia.PangolinExchange], scrape)
	case dia.SpookyswapExchange:
		return NewUniswapScraper(Exchanges[dia.SpookyswapExchange], scrape)
	case dia.QuickswapExchange:
		return NewUniswapScraper(Exchanges[dia.QuickswapExchange], scrape)
	case dia.SpiritswapExchange:
		return NewUniswapScraper(Exchanges[dia.SpiritswapExchange], scrape)
	case dia.SolarbeamExchange:
		return NewUniswapScraper(Exchanges[dia.SolarbeamExchange], scrape)
	case dia.TrisolarisExchange:
		return NewUniswapScraper(Exchanges[dia.TrisolarisExchange], scrape)
	case dia.ByBitExchange:
		return NewByBitScraper(Exchanges[dia.ByBitExchange], scrape, relDB)
	case dia.SerumExchange:
		return NewSerumScraper(Exchanges[dia.SerumExchange], scrape)
	case dia.AnyswapExchange:
		return NewAnyswapScraper(Exchanges[dia.AnyswapExchange], scrape, relDB)
	case dia.NetswapExchange:
		return NewUniswapScraper(Exchanges[dia.NetswapExchange], scrape)
	case dia.BitMexExchange:
		return NewBitMexScraper(Exchanges[dia.BitMexExchange], scrape, relDB)
	case dia.TethysExchange:
		return NewUniswapScraper(Exchanges[dia.TethysExchange], scrape)
	case dia.HermesExchange:
		return NewUniswapScraper(Exchanges[dia.HermesExchange], scrape)
	case dia.OmniDexExchange:
		return NewUniswapScraper(Exchanges[dia.OmniDexExchange], scrape)
	case dia.DiffusionExchange:
		return NewUniswapScraper(Exchanges[dia.DiffusionExchange], scrape)
	case dia.ApeswapExchange:
		return NewUniswapScraper(Exchanges[dia.ApeswapExchange], scrape)
	case dia.BiswapExchange:
		return NewUniswapScraper(Exchanges[dia.BiswapExchange], scrape)
	case dia.ArthswapExchange:
		return NewUniswapScraper(Exchanges[dia.ArthswapExchange], scrape)
	// case dia.FinageForex:
	// 	return NewFinageForexScraper(Exchanges[dia.FinageForex], scrape, relDB, key, secret)

	case "Influx":
		return NewInfluxScraper(scrape)

	case "UniswapHistory":
		return NewUniswapHistoryScraper(Exchanges[dia.UniswapExchange], scrape, relDB)

	default:
		return nil
	}

}
