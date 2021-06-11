package scrapers

import (
	datasource2 "github.com/diadata-org/diadata/pkg/datasource"
	"io"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
)

// The collector kills a scraper after @watchdogDelayXXX seconds of inactivity
const (
// TODO use this with test
// watchdogDelayShort = 10 * 60
// watchdogDelay     = 20 * 60
// watchdogDelayLong = 120 * 60
)

// empty type used for signaling
type nothing struct{}

var Exchanges map[string]dia.Exchange

func init() {
	source, err := datasource2.InitSource()
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
func NewAPIScraper(exchange string, scrape bool, key string, secret string, relDB *models.RelDB) APIScraper {
	switch exchange {
	case dia.BinanceExchange:
		return NewBinanceScraper(key, secret, Exchanges[dia.BinanceExchange], scrape, relDB)
	// case dia.BitBayExchange:
	// 	return NewBitBayScraper(Exchanges[dia.BitBayExchange], scrape, relDB)
	// case dia.BitfinexExchange:
	// 	return NewBitfinexScraper(key, secret, Exchanges[dia.BitfinexExchange], scrape, relDB)
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
	// case dia.SimexExchange:
	// 	return NewSimexScraper(Exchanges[dia.SimexExchange], scrape, relDB)
	// case dia.OKExExchange:
	// 	return NewOKExScraper(Exchanges[dia.OKExExchange], scrape, relDB)
	// case dia.HuobiExchange:
	// 	return NewHuobiScraper(Exchanges[dia.HuobiExchange], scrape, relDB)
	// case dia.LBankExchange:
	// 	return NewLBankScraper(Exchanges[dia.LBankExchange], scrape, relDB)
	// case dia.GateIOExchange:
	// 	return NewGateIOScraper(Exchanges[dia.GateIOExchange], scrape, relDB)
	// case dia.ZBExchange:
	// 	return NewZBScraper(Exchanges[dia.ZBExchange], scrape, relDB)
	// case dia.QuoineExchange:
	// 	return NewQuoineScraper(Exchanges[dia.QuoineExchange], scrape, relDB)
	// case dia.BancorExchange:
	// 	return NewBancorScraper(Exchanges[dia.BancorExchange], scrape)
	case dia.UniswapExchange:
		return NewUniswapScraper(Exchanges[dia.UniswapExchange], scrape)
	// case dia.PanCakeSwap:
	// 	return NewUniswapScraper(Exchanges[dia.PanCakeSwap], scrape)
	// case dia.SushiSwapExchange:
	// 	return NewUniswapScraper(Exchanges[dia.SushiSwapExchange], scrape)
	// case dia.LoopringExchange:
	// 	return NewLoopringScraper(Exchanges[dia.LoopringExchange], scrape)
	// case dia.CurveFIExchange:
	// 	return NewCurveFIScraper(Exchanges[dia.CurveFIExchange], scrape)
	// case dia.GnosisExchange:
	// 	return NewGnosisScraper(Exchanges[dia.GnosisExchange], scrape)
	// case dia.BalancerExchange:
	// 	return NewBalancerScraper(Exchanges[dia.BalancerExchange], scrape)
	// case dia.MakerExchange:
	// 	return NewMakerScraper(Exchanges[dia.MakerExchange], scrape)
	case dia.KuCoinExchange:
		return NewKuCoinScraper(key, secret, Exchanges[dia.KuCoinExchange], scrape, relDB)
	// case dia.DforceExchange:
	// 	return NewDforceScraper(Exchanges[dia.DforceExchange], scrape)
	// case dia.ZeroxExchange:
	// 	return NewZeroxScraper(Exchanges[dia.ZeroxExchange], scrape)
	// case dia.KyberExchange:
	// 	return NewKyberScraper(Exchanges[dia.KyberExchange], scrape)
	case dia.BitMaxExchange:
		return NewBitMaxScraper(Exchanges[dia.BitMaxExchange], scrape, relDB)
	case dia.STEXExchange:
		return NewSTEXScraper(Exchanges[dia.STEXExchange], scrape, relDB)

	default:
		return nil
	}

}
