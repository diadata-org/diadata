package scrapers

import (
	"io"

	"github.com/ethereum/go-ethereum/common"

	"github.com/diadata-org/diadata/pkg/dia"
)

// empty type used for signaling
type nothing struct{}

var exchanges map[string]dia.Exchange
var blockchains map[string]dia.BlockChain

func init() {

	blockchains = make(map[string]dia.BlockChain)
	blockchains[dia.Bitcoin] = dia.BlockChain{Name: dia.BinanceExchange, NativeToken: "BTC", VerificationMechanism: dia.PROOF_OF_STAKE}
	blockchains[dia.Ethereum] = dia.BlockChain{Name: dia.BinanceExchange, NativeToken: "ETH", VerificationMechanism: dia.PROOF_OF_STAKE}

	exchanges = make(map[string]dia.Exchange)
	exchanges[dia.BinanceExchange] = dia.Exchange{Name: dia.BinanceExchange, Centralized: true}
	exchanges[dia.GnosisExchange] = dia.Exchange{Name: dia.GnosisExchange, Centralized: false, Contract: common.HexToAddress("0x6F400810b62df8E13fded51bE75fF5393eaa841F"), BlockChain: blockchains[dia.Ethereum]}
	exchanges[dia.KrakenExchange] = dia.Exchange{Name: dia.KrakenExchange, Centralized: true}
	exchanges[dia.CREX24Exchange] = dia.Exchange{Name: dia.CREX24Exchange, Centralized: true}
	exchanges[dia.BitfinexExchange] = dia.Exchange{Name: dia.BitfinexExchange, Centralized: true}
	exchanges[dia.BitBayExchange] = dia.Exchange{Name: dia.BitBayExchange, Centralized: true}
	exchanges[dia.BittrexExchange] = dia.Exchange{Name: dia.BittrexExchange, Centralized: true}
	exchanges[dia.CoinBaseExchange] = dia.Exchange{Name: dia.CoinBaseExchange, Centralized: true}
	exchanges[dia.HitBTCExchange] = dia.Exchange{Name: dia.HitBTCExchange, Centralized: true}
	exchanges[dia.SimexExchange] = dia.Exchange{Name: dia.SimexExchange, Centralized: true}
	exchanges[dia.OKExExchange] = dia.Exchange{Name: dia.OKExExchange, Centralized: true}
	exchanges[dia.HuobiExchange] = dia.Exchange{Name: dia.HuobiExchange, Centralized: true}
	exchanges[dia.LBankExchange] = dia.Exchange{Name: dia.LBankExchange, Centralized: true}
	exchanges[dia.GateIOExchange] = dia.Exchange{Name: dia.GateIOExchange, Centralized: true}
	exchanges[dia.ZBExchange] = dia.Exchange{Name: dia.ZBExchange, Centralized: true}
	exchanges[dia.QuoineExchange] = dia.Exchange{Name: dia.QuoineExchange, Centralized: true}
	exchanges[dia.UnknownExchange] = dia.Exchange{Name: dia.UnknownExchange, Centralized: true}
	exchanges[dia.FilterKing] = dia.Exchange{Name: dia.FilterKing, Centralized: true}
	exchanges[dia.BancorExchange] = dia.Exchange{Name: dia.BancorExchange, Centralized: false, BlockChain: blockchains[dia.Ethereum]} //API is used instead of contracts
	exchanges[dia.UniswapExchange] = dia.Exchange{Name: dia.UniswapExchange, Centralized: false, BlockChain: blockchains[dia.Ethereum], Contract: common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")}
	exchanges[dia.LoopringExchange] = dia.Exchange{Name: dia.LoopringExchange, Centralized: false, BlockChain: blockchains[dia.Ethereum]} //API is used instead of contracts
	exchanges[dia.CurveFIExchange] = dia.Exchange{Name: dia.CurveFIExchange, Centralized: false, BlockChain: blockchains[dia.Ethereum], Contract: common.HexToAddress("0x7002B727Ef8F5571Cb5F9D70D13DBEEb4dFAe9d1")}
	exchanges[dia.MakerExchange] = dia.Exchange{Name: dia.MakerExchange, Centralized: false, BlockChain: blockchains[dia.Ethereum]} //API is used instead of contracts
	exchanges[dia.KuCoinExchange] = dia.Exchange{Name: dia.KuCoinExchange, Centralized: true}
	exchanges[dia.SushiSwapExchange] = dia.Exchange{Name: dia.SushiSwapExchange, Centralized: false, BlockChain: blockchains[dia.Ethereum], Contract: common.HexToAddress("0xc0aee478e3658e2610c5f7a4a2e1777ce9e4f2ac")}
	exchanges[dia.PanCakeSwap] = dia.Exchange{Name: dia.PanCakeSwap, Centralized: false, BlockChain: blockchains[dia.Ethereum], Contract: common.HexToAddress("0xbcfccbde45ce874adcb698cc183debcf17952812")}
	exchanges[dia.DforceExchange] = dia.Exchange{Name: dia.DforceExchange, Centralized: false, BlockChain: blockchains[dia.Ethereum], Contract: common.HexToAddress("0x03eF3f37856bD08eb47E2dE7ABc4Ddd2c19B60F2")}
	exchanges[dia.ZeroxExchange] = dia.Exchange{Name: dia.ZeroxExchange, Centralized: true}
	exchanges[dia.KyberExchange] = dia.Exchange{Name: dia.KyberExchange, Centralized: true}
	exchanges[dia.BitMaxExchange] = dia.Exchange{Name: dia.BitMaxExchange, Centralized: true}
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
		return NewBinanceScraper(key, secret, exchanges[dia.BinanceExchange])
	case dia.BitBayExchange:
		return NewBitBayScraper(exchanges[dia.BitBayExchange])
	case dia.BitfinexExchange:
		return NewBitfinexScraper(key, secret, exchanges[dia.BitfinexExchange])
	case dia.BittrexExchange:
		return NewBittrexScraper(exchanges[dia.BittrexExchange])
	case dia.CoinBaseExchange:
		return NewCoinBaseScraper(exchanges[dia.CoinBaseExchange])
	case dia.CREX24Exchange:
		return NewCREX24Scraper(exchanges[dia.CREX24Exchange])
	case dia.KrakenExchange:
		return NewKrakenScraper(key, secret, exchanges[dia.KrakenExchange])
	case dia.HitBTCExchange:
		return NewHitBTCScraper(exchanges[dia.HitBTCExchange])
	case dia.SimexExchange:
		return NewSimexScraper(exchanges[dia.SimexExchange])
	case dia.OKExExchange:
		return NewOKExScraper(exchanges[dia.OKExExchange])
	case dia.HuobiExchange:
		return NewHuobiScraper(exchanges[dia.HuobiExchange])
	case dia.LBankExchange:
		return NewLBankScraper(exchanges[dia.LBankExchange])
	case dia.GateIOExchange:
		return NewGateIOScraper(exchanges[dia.GateIOExchange])
	case dia.ZBExchange:
		return NewZBScraper(exchanges[dia.ZBExchange])
	case dia.QuoineExchange:
		return NewQuoineScraper(exchanges[dia.QuoineExchange])
	case dia.BancorExchange:
		return NewBancorScraper(exchanges[dia.BancorExchange])
	case dia.UniswapExchange:
		return NewUniswapScraper(exchanges[dia.UniswapExchange])
	case dia.PanCakeSwap:
		return NewUniswapScraper(exchanges[dia.PanCakeSwap])
	case dia.SushiSwapExchange:
		return NewUniswapScraper(exchanges[dia.SushiSwapExchange])
	case dia.LoopringExchange:
		return NewLoopringScraper(exchanges[dia.LoopringExchange])
	case dia.CurveFIExchange:
		return NewCurveFIScraper(exchanges[dia.CurveFIExchange])
	case dia.GnosisExchange:
		return NewGnosisScraper(exchanges[dia.GnosisExchange])
	case dia.BalancerExchange:
		return NewBalancerScraper(exchanges[dia.BalancerExchange])
	case dia.MakerExchange:
		return NewMakerScraper(exchanges[dia.MakerExchange])
	case dia.KuCoinExchange:
		return NewKuCoinScraper(key, secret, exchanges[dia.KuCoinExchange])
	case dia.DforceExchange:
		return NewDforceScraper(exchanges[dia.DforceExchange])
	case dia.ZeroxExchange:
		return NewZeroxScraper(exchanges[dia.ZeroxExchange])
	case dia.KyberExchange:
		return NewKyberScraper(exchanges[dia.KyberExchange])
	case dia.BitMaxExchange:
		return NewBitMaxScraper(exchanges[dia.BitMaxExchange])

	default:
		return nil
	}

}
