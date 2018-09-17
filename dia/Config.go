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
	CoinBaseExchange = "CoinBase"
	HitBTCExchange   = "HitBTC"
	SimexExchange    = "Simex"
	OKExExchange     = "OKEx"
	HuobiExchange    = "Huobi"
	LBankExchange    = "LBank"
	GateIOExchange   = "GateIO"
	UnknownExchange  = "Unknown"
)

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
		OKExExchange}
}

type ConfigApi struct {
	ApiKey    string
	SecretKey string
}

type ConfigPair struct {
	Symbol      string
	ForeignName string
	Exchange    string
}

type ConfigConnector struct {
	Coins []ConfigPair
}

func GetConfig(exchange string) (*ConfigApi, error) {
	var configApi ConfigApi
	usr, _ := user.Current()
	dir := usr.HomeDir
	configFileApi := dir + "/config/secrets/api_" + strings.ToLower(exchange)
	err := gonfig.GetConf(configFileApi, &configApi)
	return &configApi, err
}
