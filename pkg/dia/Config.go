package dia

import (
	"os/user"
	"strings"

	"github.com/tkanos/gonfig"
)

const (
	KrakenExchange   = "Kraken"
	BitfinexExchange = "Bitfinex"
	BinanceExchange  = "Binance"
	BitBayExchange   = "BitBay"
	BittrexExchange  = "Bittrex"
	CoinBaseExchange = "CoinBase"
	HitBTCExchange   = "HitBTC"
	SimexExchange    = "Simex"
	OKExExchange     = "OKEx"
	HuobiExchange    = "Huobi"
	LBankExchange    = "LBank"
	GateIOExchange   = "GateIO"
	ZBExchange       = "ZB"
	QuoineExchange   = "Quoine"
	UnknownExchange  = "Unknown"
	BlockSizeSeconds = 120
	FilterKing       = "MA120"
	BancorExchange   = "Bancor"
)

func Exchanges() []string {
	return []string{
		BinanceExchange,
		BitfinexExchange,
		BittrexExchange,
		CoinBaseExchange,
		GateIOExchange,
		HitBTCExchange,
		HuobiExchange,
		KrakenExchange,
		LBankExchange,
		OKExExchange,
		QuoineExchange,
		SimexExchange,
		ZBExchange,
		BancorExchange,
		UnknownExchange,
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
