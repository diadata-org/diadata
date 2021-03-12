package scrapers

import (
	"io"

	"github.com/diadata-org/diadata/internal/pkg/datasource"

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

func init() {
	source, err := datasource.InitSource()
	if err != nil {
		panic(err)
	}
	Exchanges = source.GetExchanges()
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
func NewAPIScraper(exchange string, scrape bool, key string, secret string) APIScraper {
	switch exchange {
	// case dia.BinanceExchange:
	// 	return NewBinanceScraper(key, secret, Exchanges[dia.BinanceExchange])
	// case dia.BitBayExchange:
	// 	return NewBitBayScraper(Exchanges[dia.BitBayExchange])
	// case dia.BitfinexExchange:
	// 	return NewBitfinexScraper(key, secret, Exchanges[dia.BitfinexExchange])
	// case dia.BittrexExchange:
	// 	return NewBittrexScraper(Exchanges[dia.BittrexExchange], scrape)
	case dia.CoinBaseExchange:
		return NewCoinBaseScraper(Exchanges[dia.CoinBaseExchange], scrape)
	case dia.CREX24Exchange:
		return NewCREX24Scraper(Exchanges[dia.CREX24Exchange])
	// case dia.KrakenExchange:
	// 	return NewKrakenScraper(key, secret, Exchanges[dia.KrakenExchange])
	case dia.HitBTCExchange:
		return NewHitBTCScraper(Exchanges[dia.HitBTCExchange], scrape)
	// case dia.SimexExchange:
	// 	return NewSimexScraper(Exchanges[dia.SimexExchange])
	// case dia.OKExExchange:
	// 	return NewOKExScraper(Exchanges[dia.OKExExchange])
	// case dia.HuobiExchange:
	// 	return NewHuobiScraper(Exchanges[dia.HuobiExchange])
	// case dia.LBankExchange:
	// 	return NewLBankScraper(Exchanges[dia.LBankExchange])
	// case dia.GateIOExchange:
	// 	return NewGateIOScraper(Exchanges[dia.GateIOExchange])
	// case dia.ZBExchange:
	// 	return NewZBScraper(Exchanges[dia.ZBExchange])
	// case dia.QuoineExchange:
	// 	return NewQuoineScraper(Exchanges[dia.QuoineExchange])
	// case dia.BancorExchange:
	// 	return NewBancorScraper(Exchanges[dia.BancorExchange])
	case dia.UniswapExchange:
		return NewUniswapScraper(Exchanges[dia.UniswapExchange], scrape)
	// case dia.PanCakeSwap:
	// 	return NewUniswapScraper(Exchanges[dia.PanCakeSwap])
	// case dia.SushiSwapExchange:
	// 	return NewUniswapScraper(Exchanges[dia.SushiSwapExchange])
	// case dia.LoopringExchange:
	// 	return NewLoopringScraper(Exchanges[dia.LoopringExchange])
	// case dia.CurveFIExchange:
	// 	return NewCurveFIScraper(Exchanges[dia.CurveFIExchange])
	// case dia.GnosisExchange:
	// 	return NewGnosisScraper(Exchanges[dia.GnosisExchange])
	// case dia.BalancerExchange:
	// 	return NewBalancerScraper(Exchanges[dia.BalancerExchange])
	// case dia.MakerExchange:
	// 	return NewMakerScraper(Exchanges[dia.MakerExchange])
	// case dia.KuCoinExchange:
	// 	return NewKuCoinScraper(key, secret, Exchanges[dia.KuCoinExchange])
	// case dia.DforceExchange:
	// 	return NewDforceScraper(Exchanges[dia.DforceExchange])
	// case dia.ZeroxExchange:
	// 	return NewZeroxScraper(Exchanges[dia.ZeroxExchange])
	// case dia.KyberExchange:
	// 	return NewKyberScraper(Exchanges[dia.KyberExchange])
	case dia.BitMaxExchange:
		return NewBitMaxScraper(Exchanges[dia.BitMaxExchange], scrape)
	// case dia.STEXExchange:
	// 	return NewSTEXScraper(Exchanges[dia.STEXExchange])

	default:
		return nil
	}

}
