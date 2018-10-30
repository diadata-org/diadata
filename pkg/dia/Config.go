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
	QuoineExchange   = "Quoine"
	UnknownExchange  = "Unknown"
	BlockSizeSeconds = 120
	FilterKing       = "MA120"
)

func Exchanges() []string {
	return []string{
		BinanceExchange,
		BitfinexExchange,
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
