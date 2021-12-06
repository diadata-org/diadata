package scrapers

import (
	"io"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/common"
)

// The collector kills a scraper after @watchdogDelayXXX seconds of inactivity
const (
	// TODO use this with test
	watchdogDelay3Mins = 3 * 60
	watchdogDelayShort = 10 * 60
	watchdogDelay      = 20 * 60
	watchdogDelayLong  = 120 * 60
)

// empty type used for signaling
type nothing struct{}

var Exchanges map[string]dia.Exchange
var blockchains map[string]dia.BlockChain

func init() {

	blockchains = make(map[string]dia.BlockChain)
	blockchains[dia.BITCOIN] = dia.BlockChain{Name: dia.BITCOIN, NativeToken: "BTC", VerificationMechanism: dia.PROOF_OF_WORK}
	blockchains[dia.ETHEREUM] = dia.BlockChain{Name: dia.ETHEREUM, NativeToken: "ETH", VerificationMechanism: dia.PROOF_OF_WORK}
	blockchains[dia.BINANCESMARTCHAIN] = dia.BlockChain{Name: dia.BINANCESMARTCHAIN, NativeToken: "", VerificationMechanism: dia.PROOF_OF_STAKE}
	blockchains[dia.POLYGON] = dia.BlockChain{Name: dia.POLYGON, NativeToken: "MATIC", VerificationMechanism: dia.PROOF_OF_STAKE}
	blockchains[dia.CELO] = dia.BlockChain{Name: dia.CELO, NativeToken: "CELO", VerificationMechanism: dia.PROOF_OF_STAKE}
	blockchains[dia.FANTOM] = dia.BlockChain{Name: dia.FANTOM, NativeToken: "FTM", VerificationMechanism: dia.PROOF_OF_STAKE}
	blockchains[dia.NEAR] = dia.BlockChain{Name: dia.NEAR, NativeToken: "NEAR", VerificationMechanism: dia.PROOF_OF_STAKE}
	blockchains[dia.AURORA] = dia.BlockChain{Name: dia.AURORA, NativeToken: "AURORA", VerificationMechanism: dia.PROOF_OF_STAKE}
	blockchains[dia.SOLANA] = dia.BlockChain{Name: dia.SOLANA, NativeToken: "SOL", VerificationMechanism: dia.PROOF_OF_STAKE}
	blockchains[dia.FLOW] = dia.BlockChain{Name: dia.FLOW, NativeToken: "FLOW", VerificationMechanism: dia.PROOF_OF_STAKE}
	blockchains[dia.MOONRIVER] = dia.BlockChain{Name: dia.MOONRIVER, NativeToken: "MOVR", VerificationMechanism: dia.PROOF_OF_STAKE}
	blockchains[dia.FIAT] = dia.BlockChain{Name: dia.FIAT}

	Exchanges = make(map[string]dia.Exchange)
	Exchanges[dia.BalancerExchange] = dia.Exchange{Name: dia.BalancerExchange, Centralized: false, Contract: common.HexToAddress("0x9424B1412450D0f8Fc2255FAf6046b98213B76Bd"), WatchdogDelay: watchdogDelay}
	Exchanges[dia.BinanceExchange] = dia.Exchange{Name: dia.BinanceExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	Exchanges[dia.GnosisExchange] = dia.Exchange{Name: dia.GnosisExchange, Centralized: false, Contract: common.HexToAddress("0x6F400810b62df8E13fded51bE75fF5393eaa841F"), BlockChain: blockchains[dia.ETHEREUM], WatchdogDelay: watchdogDelayLong}
	Exchanges[dia.KrakenExchange] = dia.Exchange{Name: dia.KrakenExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	Exchanges[dia.CREX24Exchange] = dia.Exchange{Name: dia.CREX24Exchange, Centralized: true, WatchdogDelay: watchdogDelay}
	Exchanges[dia.BitfinexExchange] = dia.Exchange{Name: dia.BitfinexExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	Exchanges[dia.BitBayExchange] = dia.Exchange{Name: dia.BitBayExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	Exchanges[dia.BittrexExchange] = dia.Exchange{Name: dia.BittrexExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	Exchanges[dia.CoinBaseExchange] = dia.Exchange{Name: dia.CoinBaseExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	Exchanges[dia.HitBTCExchange] = dia.Exchange{Name: dia.HitBTCExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	Exchanges[dia.SimexExchange] = dia.Exchange{Name: dia.SimexExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	Exchanges[dia.OKExExchange] = dia.Exchange{Name: dia.OKExExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	Exchanges[dia.HuobiExchange] = dia.Exchange{Name: dia.HuobiExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	Exchanges[dia.LBankExchange] = dia.Exchange{Name: dia.LBankExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	Exchanges[dia.GateIOExchange] = dia.Exchange{Name: dia.GateIOExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	Exchanges[dia.ZBExchange] = dia.Exchange{Name: dia.ZBExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	Exchanges[dia.QuoineExchange] = dia.Exchange{Name: dia.QuoineExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	Exchanges[dia.UnknownExchange] = dia.Exchange{Name: dia.UnknownExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	Exchanges[dia.FilterKing] = dia.Exchange{Name: dia.FilterKing, Centralized: true, WatchdogDelay: watchdogDelay}
	Exchanges[dia.BancorExchange] = dia.Exchange{Name: dia.BancorExchange, Centralized: false, BlockChain: blockchains[dia.ETHEREUM], WatchdogDelay: watchdogDelayLong} //API is used instead of contracts
	Exchanges[dia.UniswapExchange] = dia.Exchange{Name: dia.UniswapExchange, Centralized: false, BlockChain: blockchains[dia.ETHEREUM], Contract: common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"), WatchdogDelay: watchdogDelay}
	Exchanges[dia.UniswapExchangeV3] = dia.Exchange{Name: dia.UniswapExchangeV3, Centralized: false, BlockChain: blockchains[dia.ETHEREUM], Contract: common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"), WatchdogDelay: watchdogDelay}
	Exchanges[dia.LoopringExchange] = dia.Exchange{Name: dia.LoopringExchange, Centralized: true, BlockChain: blockchains[dia.ETHEREUM], WatchdogDelay: watchdogDelay} //API is used instead of contracts
	Exchanges[dia.CurveFIExchange] = dia.Exchange{Name: dia.CurveFIExchange, Centralized: false, BlockChain: blockchains[dia.ETHEREUM], Contract: common.HexToAddress("0x7002B727Ef8F5571Cb5F9D70D13DBEEb4dFAe9d1"), WatchdogDelay: watchdogDelay}
	Exchanges[dia.MakerExchange] = dia.Exchange{Name: dia.MakerExchange, Centralized: false, BlockChain: blockchains[dia.ETHEREUM], WatchdogDelay: watchdogDelay} //API is used instead of contracts
	Exchanges[dia.KuCoinExchange] = dia.Exchange{Name: dia.KuCoinExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	Exchanges[dia.SushiSwapExchange] = dia.Exchange{Name: dia.SushiSwapExchange, Centralized: false, BlockChain: blockchains[dia.ETHEREUM], Contract: common.HexToAddress("0xc0aee478e3658e2610c5f7a4a2e1777ce9e4f2ac"), WatchdogDelay: watchdogDelay}
	Exchanges[dia.PanCakeSwap] = dia.Exchange{Name: dia.PanCakeSwap, Centralized: false, BlockChain: blockchains[dia.BINANCESMARTCHAIN], Contract: common.HexToAddress("0xbcfccbde45ce874adcb698cc183debcf17952812"), WatchdogDelay: watchdogDelayLong}
	Exchanges[dia.DforceExchange] = dia.Exchange{Name: dia.DforceExchange, Centralized: false, BlockChain: blockchains[dia.ETHEREUM], Contract: common.HexToAddress("0x03eF3f37856bD08eb47E2dE7ABc4Ddd2c19B60F2"), WatchdogDelay: watchdogDelayLong}
	Exchanges[dia.ZeroxExchange] = dia.Exchange{Name: dia.ZeroxExchange, Centralized: true, WatchdogDelay: watchdogDelayLong}
	Exchanges[dia.KyberExchange] = dia.Exchange{Name: dia.KyberExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	Exchanges[dia.BitMaxExchange] = dia.Exchange{Name: dia.BitMaxExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	Exchanges[dia.STEXExchange] = dia.Exchange{Name: dia.STEXExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	Exchanges[dia.DfynNetwork] = dia.Exchange{Name: dia.DfynNetwork, Centralized: false, BlockChain: blockchains[dia.POLYGON], Contract: common.HexToAddress("0xe7fb3e833efe5f9c441105eb65ef8b261266423b"), WatchdogDelay: watchdogDelayLong}
	Exchanges[dia.QuickswapExchange] = dia.Exchange{Name: dia.QuickswapExchange, Centralized: false, BlockChain: blockchains[dia.POLYGON], Contract: common.HexToAddress("0x5757371414417b8C6CAad45bAeF941aBc7d3Ab32"), WatchdogDelay: watchdogDelayLong}
	Exchanges[dia.UbeswapExchange] = dia.Exchange{Name: dia.UbeswapExchange, Centralized: false, BlockChain: blockchains[dia.CELO], Contract: common.HexToAddress("0x62d5b84bE28a183aBB507E125B384122D2C25fAE"), WatchdogDelay: watchdogDelayLong}
	Exchanges[dia.SpookyswapExchange] = dia.Exchange{Name: dia.SpookyswapExchange, Centralized: false, BlockChain: blockchains[dia.FANTOM], Contract: common.HexToAddress("0x152ee697f2e276fa89e96742e9bb9ab1f2e61be3"), WatchdogDelay: watchdogDelayLong}
	Exchanges[dia.SpiritswapExchange] = dia.Exchange{Name: dia.SpiritswapExchange, Centralized: false, BlockChain: blockchains[dia.FANTOM], Contract: common.HexToAddress("0xef45d134b73241eda7703fa787148d9c9f4950b0"), WatchdogDelay: watchdogDelayLong}
	Exchanges[dia.SerumExchange] = dia.Exchange{Name: dia.SerumExchange, Centralized: false, BlockChain: blockchains[dia.SOLANA], Contract: common.HexToAddress(""), WatchdogDelay: watchdogDelayLong}
	Exchanges[dia.SolarbeamExchange] = dia.Exchange{Name: dia.SolarbeamExchange, Centralized: false, BlockChain: blockchains[dia.MOONRIVER], Contract: common.HexToAddress("0x049581aEB6Fe262727f290165C29BDAB065a1B68"), WatchdogDelay: watchdogDelay3Mins}
	Exchanges[dia.TrisolarisExchange] = dia.Exchange{Name: dia.TrisolarisExchange, Centralized: false, BlockChain: blockchains[dia.AURORA], Contract: common.HexToAddress("0xc66F594268041dB60507F00703b152492fb176E7"), WatchdogDelay: watchdogDelay}

	// Exchanges[dia.FinageForex] = dia.Exchange{Name: dia.FinageForex, Centralized: true, BlockChain: blockchains[dia.FIAT], WatchdogDelay: watchdogDelay}
	Exchanges["Influx"] = dia.Exchange{Name: "Influx", WatchdogDelay: 360000}
	Exchanges["UniswapHistory"] = dia.Exchange{Name: "UniswapHistory", Centralized: false, BlockChain: blockchains[dia.ETHEREUM], Contract: common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"), WatchdogDelay: 3600}
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
	case dia.BitBayExchange:
		return NewBitBayScraper(Exchanges[dia.BitBayExchange], scrape, relDB)
	case dia.BitfinexExchange:
		return NewBitfinexScraper(key, secret, Exchanges[dia.BitfinexExchange], scrape, relDB)
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
	case dia.GnosisExchange:
		return NewGnosisScraper(Exchanges[dia.GnosisExchange], scrape)
	case dia.BalancerExchange:
		return NewBalancerScraper(Exchanges[dia.BalancerExchange], scrape)
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
	// case dia.SerumExchange:
	// 	return NewSerumScraper(Exchanges[dia.SerumExchange], scrape)
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
