package dia

import (
	"github.com/tkanos/gonfig"
	"os/user"
	"strings"
)

const (
	KrakenExchange   = "Kraken"
	BitfinexExchange = "Bitfinex"
	BinanceExchange  = "Binance"
	CoinBaseExchange = "CoinBase"
	HitBTCExchange   = "HitBTC"
	SimexExchange    = "Simex"
	OKExExchange     = "OKEx"
	HuobiExchange    = "Huobi"
	LBankExchange    = "LBank"
	GateIOExchange   = "GateIO"
	ZBExchange       = "ZB"
	UnknownExchange  = "Unknown"
	BlockSizeSeconds = 120
	FilterKing       = "MA120"
)

func SymbolsFrontPage() []string {
	return []string{
		"BCH",
		"BTC",
		"ADA",
		"DASH",
		"DOGE",
		"EOS",
		"ETH",
		"BNB",
		"LTC",
		"XMR",
		"NEO",
		"MIOTA",
		"XRP",
		"XLM",
		"TRX",
		"USDT",
		"ETC",
		"XEM",
		"XTZ",
		"VET",
		"ZEC",
		"OMG",
		"MKR",
		"BTG",
		"ONT",
		"ZRX",
		"LSK",
		"DCR",
		"QTUM",
		"BCN",
		"DGB",
		"ICX",
		"VET",
		"PAX",
		"AE",
		"BTS",
		"BCD",
		"ZIL",
		"NANO",
		"ICX",
		"SC",
		"STEEM",
		"XVG",
		"NPXS",
		"WAVES",
		"ETN",
		"BTM",
		"ETP",
		"BAT",
		"GNT",
		"STRAT",
		"HOT",
		"REP",
	}
}

func Exchanges() []string {
	return []string{
		KrakenExchange,
		BitfinexExchange,
		BinanceExchange,
		CoinBaseExchange,
		HitBTCExchange,
		SimexExchange,
		HuobiExchange,
		LBankExchange,
		GateIOExchange,
		OKExExchange,
		ZBExchange,
	}
}

type ConfigApi struct {
	ApiKey    string
	SecretKey string
}

type ConfigConnector struct {
	Coins []Pair
}

func GetConfig(exchange string) (*ConfigApi, error) {
	var configApi ConfigApi
	usr, _ := user.Current()
	dir := usr.HomeDir
	configFileApi := dir + "/config/secrets/api_" + strings.ToLower(exchange)
	err := gonfig.GetConf(configFileApi, &configApi)
	return &configApi, err
}
