package dia

import (
	"errors"
	"os/user"
	"strings"

	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/tkanos/gonfig"
)

const (
	BalancerExchange  = "Balancer"
	GnosisExchange    = "Gnosis"
	KrakenExchange    = "Kraken"
	BitfinexExchange  = "Bitfinex"
	BinanceExchange   = "Binance"
	BitBayExchange    = "BitBay"
	BittrexExchange   = "Bittrex"
	CoinBaseExchange  = "CoinBase"
	HitBTCExchange    = "HitBTC"
	SimexExchange     = "Simex"
	OKExExchange      = "OKEx"
	HuobiExchange     = "Huobi"
	LBankExchange     = "LBank"
	GateIOExchange    = "GateIO"
	ZBExchange        = "ZB"
	QuoineExchange    = "Quoine"
	UnknownExchange   = "Unknown"
	BlockSizeSeconds  = 120
	FilterKing        = "MA120"
	BancorExchange    = "Bancor"
	UniswapExchange   = "Uniswap"
	LoopringExchange  = "Loopring"
	CurveFIExchange   = "Curvefi"
	MakerExchange     = "Maker"
	KuCoinExchange    = "KuCoin"
	SushiSwapExchange = "SushiSwap"
	PanCakeSwap       = "PanCakeSwap"
	DforceExchange    = "Dforce"
	ZeroxExchange     = "0x"
	KyberExchange     = "Kyber"
	BitMaxExchange    = "Bitmax"
	CREX24Exchange    = "CREX24"
	STEXExchange      = "STEX"
)

func Exchanges() []string {
	return []string{
		BitBayExchange,
		BittrexExchange,
		CoinBaseExchange,
		KrakenExchange,
		KuCoinExchange,

		BitfinexExchange,
		GateIOExchange,
		HitBTCExchange,
		OKExExchange,
		SimexExchange,

		BinanceExchange,
		CREX24Exchange,
		HuobiExchange,
		LBankExchange,
		QuoineExchange,
		STEXExchange,

		BancorExchange,
		CurveFIExchange,
		DforceExchange,
		GnosisExchange,
		KyberExchange,
		LoopringExchange,
		MakerExchange,
		PanCakeSwap,
		SushiSwapExchange,
		UniswapExchange,
		ZBExchange,
		ZeroxExchange,
		BalancerExchange,
		BitMaxExchange,
		UnknownExchange,
	}
}

type ConfigApi struct {
	ApiKey    string
	SecretKey string
}

type ConfigConnector struct {
	Coins []ExchangePair
}

func GetConfig(exchange string) (*ConfigApi, error) {
	if utils.Getenv("USE_ENV", "false") == "true" {
		return GetConfigFromEnv(exchange)
	}
	var configApi ConfigApi
	usr, _ := user.Current()
	dir := usr.HomeDir
	configFileApi := dir + "/config/secrets/api_" + strings.ToLower(exchange)
	err := gonfig.GetConf(configFileApi, &configApi)
	return &configApi, err
}

func GetConfigFromEnv(exchange string) (*ConfigApi, error) {
	if utils.Getenv("USE_ENV", "false") != "true" {
		return nil, errors.New("use of config by env without env activation ")
	}

	configApi := ConfigApi{
		ApiKey:    utils.Getenv("API_"+strings.ToUpper(exchange)+"_APIKEY", ""),
		SecretKey: utils.Getenv("API_"+strings.ToUpper(exchange)+"_SECRETKEY", ""),
	}
	return &configApi, nil
}
