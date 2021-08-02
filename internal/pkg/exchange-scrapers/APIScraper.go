package scrapers

import (
	"io"

	"github.com/ethereum/go-ethereum/common"

	"github.com/diadata-org/diadata/pkg/dia"
)

// The collector kills a scraper after @watchdogDelayXXX seconds of inactivity
const (
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
	blockchains[dia.Bitcoin] = dia.BlockChain{Name: dia.BinanceExchange, NativeToken: "BTC", VerificationMechanism: dia.PROOF_OF_WORK}
	blockchains[dia.Ethereum] = dia.BlockChain{Name: dia.BinanceExchange, NativeToken: "ETH", VerificationMechanism: dia.PROOF_OF_WORK}

	Exchanges = make(map[string]dia.Exchange)
	Exchanges[dia.BalancerExchange] = dia.Exchange{Name: dia.BalancerExchange, Centralized: false, Contract: common.HexToAddress("0x9424B1412450D0f8Fc2255FAf6046b98213B76Bd"), WatchdogDelay: watchdogDelay}
	Exchanges[dia.BinanceExchange] = dia.Exchange{Name: dia.BinanceExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	Exchanges[dia.GnosisExchange] = dia.Exchange{Name: dia.GnosisExchange, Centralized: false, Contract: common.HexToAddress("0x6F400810b62df8E13fded51bE75fF5393eaa841F"), BlockChain: blockchains[dia.Ethereum], WatchdogDelay: watchdogDelayLong}
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
	Exchanges[dia.BancorExchange] = dia.Exchange{Name: dia.BancorExchange, Centralized: false, BlockChain: blockchains[dia.Ethereum], WatchdogDelay: watchdogDelayLong} //API is used instead of contracts
	Exchanges[dia.UniswapExchange] = dia.Exchange{Name: dia.UniswapExchange, Centralized: false, BlockChain: blockchains[dia.Ethereum], Contract: common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"), WatchdogDelay: watchdogDelay}
	Exchanges[dia.UniswapExchangeV3] = dia.Exchange{Name: dia.UniswapExchangeV3, Centralized: false, BlockChain: blockchains[dia.Ethereum], Contract: common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"), WatchdogDelay: watchdogDelay}
	Exchanges[dia.LoopringExchange] = dia.Exchange{Name: dia.LoopringExchange, Centralized: false, BlockChain: blockchains[dia.Ethereum], WatchdogDelay: watchdogDelayLong} //API is used instead of contracts
	Exchanges[dia.CurveFIExchange] = dia.Exchange{Name: dia.CurveFIExchange, Centralized: false, BlockChain: blockchains[dia.Ethereum], Contract: common.HexToAddress("0x7002B727Ef8F5571Cb5F9D70D13DBEEb4dFAe9d1"), WatchdogDelay: watchdogDelay}
	Exchanges[dia.MakerExchange] = dia.Exchange{Name: dia.MakerExchange, Centralized: false, BlockChain: blockchains[dia.Ethereum], WatchdogDelay: watchdogDelay} //API is used instead of contracts
	Exchanges[dia.KuCoinExchange] = dia.Exchange{Name: dia.KuCoinExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	Exchanges[dia.SushiSwapExchange] = dia.Exchange{Name: dia.SushiSwapExchange, Centralized: false, BlockChain: blockchains[dia.Ethereum], Contract: common.HexToAddress("0xc0aee478e3658e2610c5f7a4a2e1777ce9e4f2ac"), WatchdogDelay: watchdogDelay}
	Exchanges[dia.PanCakeSwap] = dia.Exchange{Name: dia.PanCakeSwap, Centralized: false, BlockChain: blockchains[dia.Ethereum], Contract: common.HexToAddress("0xbcfccbde45ce874adcb698cc183debcf17952812"), WatchdogDelay: watchdogDelayLong}
	Exchanges[dia.DforceExchange] = dia.Exchange{Name: dia.DforceExchange, Centralized: false, BlockChain: blockchains[dia.Ethereum], Contract: common.HexToAddress("0x03eF3f37856bD08eb47E2dE7ABc4Ddd2c19B60F2"), WatchdogDelay: watchdogDelayLong}
	Exchanges[dia.ZeroxExchange] = dia.Exchange{Name: dia.ZeroxExchange, Centralized: true, WatchdogDelay: watchdogDelayLong}
	Exchanges[dia.KyberExchange] = dia.Exchange{Name: dia.KyberExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	Exchanges[dia.BitMaxExchange] = dia.Exchange{Name: dia.BitMaxExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	Exchanges[dia.STEXExchange] = dia.Exchange{Name: dia.STEXExchange, Centralized: true, WatchdogDelay: watchdogDelay}
	Exchanges[dia.DfynNetwork] = dia.Exchange{Name: dia.DfynNetwork, Centralized: false, BlockChain: blockchains[dia.Ethereum], Contract: common.HexToAddress("0xe7fb3e833efe5f9c441105eb65ef8b261266423b"), WatchdogDelay: watchdogDelay}
}

// APIScraper provides common methods needed to get Trade information from
// exchange APIs.
type APIScraper interface {
	io.Closer
	// ScrapePair returns a PairScraper that continuously scrapes trades for a
	// single pair from this APIScraper
	ScrapePair(pair dia.Pair) (PairScraper, error)
	// FetchAvailablePairs returns a list with all available trade pairs
	FetchAvailablePairs() (pairs []dia.Pair, err error)
	NormalizePair(pair dia.Pair) (dia.Pair, error)
	// Channel returns a channel that can be used to receive trades
	Channel() chan *dia.Trade
}

// PairScraper receives trades for a single pc.Pair from a single exchange.
type PairScraper interface {
	io.Closer
	// Error returns an error when the channel Channel() is closed
	// and nil otherwise
	Error() error

	// Pair returns the pair this scraper is subscribed to
	Pair() dia.Pair
}

func NewAPIScraper(exchange string, key string, secret string) APIScraper {
	switch exchange {
	case dia.BinanceExchange:
		return NewBinanceScraper(key, secret, Exchanges[dia.BinanceExchange])
	case dia.BitBayExchange:
		return NewBitBayScraper(Exchanges[dia.BitBayExchange])
	case dia.BitfinexExchange:
		return NewBitfinexScraper(key, secret, Exchanges[dia.BitfinexExchange])
	case dia.BittrexExchange:
		return NewBittrexScraper(Exchanges[dia.BittrexExchange])
	case dia.CoinBaseExchange:
		return NewCoinBaseScraper(Exchanges[dia.CoinBaseExchange])
	case dia.CREX24Exchange:
		return NewCREX24Scraper(Exchanges[dia.CREX24Exchange])
	case dia.KrakenExchange:
		return NewKrakenScraper(key, secret, Exchanges[dia.KrakenExchange])
	case dia.HitBTCExchange:
		return NewHitBTCScraper(Exchanges[dia.HitBTCExchange])
	case dia.SimexExchange:
		return NewSimexScraper(Exchanges[dia.SimexExchange])
	case dia.OKExExchange:
		return NewOKExScraper(Exchanges[dia.OKExExchange])
	case dia.HuobiExchange:
		return NewHuobiScraper(Exchanges[dia.HuobiExchange])
	case dia.LBankExchange:
		return NewLBankScraper(Exchanges[dia.LBankExchange])
	case dia.GateIOExchange:
		return NewGateIOScraper(Exchanges[dia.GateIOExchange])
	case dia.ZBExchange:
		return NewZBScraper(Exchanges[dia.ZBExchange])
	case dia.QuoineExchange:
		return NewQuoineScraper(Exchanges[dia.QuoineExchange])
	case dia.BancorExchange:
		return NewBancorScraper(Exchanges[dia.BancorExchange])
	case dia.UniswapExchange:
		return NewUniswapScraper(Exchanges[dia.UniswapExchange])
	case dia.PanCakeSwap:
		return NewUniswapScraper(Exchanges[dia.PanCakeSwap])
	case dia.SushiSwapExchange:
		return NewUniswapScraper(Exchanges[dia.SushiSwapExchange])
	case dia.LoopringExchange:
		return NewLoopringScraper(Exchanges[dia.LoopringExchange])
	case dia.CurveFIExchange:
		return NewCurveFIScraper(Exchanges[dia.CurveFIExchange])
	case dia.GnosisExchange:
		return NewGnosisScraper(Exchanges[dia.GnosisExchange])
	case dia.BalancerExchange:
		return NewBalancerScraper(Exchanges[dia.BalancerExchange])
	case dia.MakerExchange:
		return NewMakerScraper(Exchanges[dia.MakerExchange])
	case dia.KuCoinExchange:
		return NewKuCoinScraper(key, secret, Exchanges[dia.KuCoinExchange])
	case dia.DforceExchange:
		return NewDforceScraper(Exchanges[dia.DforceExchange])
	case dia.ZeroxExchange:
		return NewZeroxScraper(Exchanges[dia.ZeroxExchange])
	case dia.KyberExchange:
		return NewKyberScraper(Exchanges[dia.KyberExchange])
	case dia.BitMaxExchange:
		return NewBitMaxScraper(Exchanges[dia.BitMaxExchange])
	case dia.STEXExchange:
		return NewSTEXScraper(Exchanges[dia.STEXExchange])
	case dia.UniswapExchangeV3:
		return NewUniswapV3Scraper(Exchanges[dia.UniswapExchangeV3])
	case dia.DfynNetwork:
		return NewUniswapScraper(Exchanges[dia.DfynNetwork])

	default:
		return nil
	}

}
