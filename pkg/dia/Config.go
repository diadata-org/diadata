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
	FTX               = "FTX"
	Opyn              = "OPYN"
	Premia            = "Premia"
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
	UniswapExchangeV3 = "UniswapV3"
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
	Deribit           = "Deribit"
	DfynNetwork       = "DFYN"
	UbeswapExchange   = "Ubeswap"
)

func Exchanges() []string {
	return []string{

		BitMaxExchange,
		ZBExchange,
		HuobiExchange,
		BitfinexExchange,
		BitBayExchange,
		BittrexExchange,
		CoinBaseExchange,
		KrakenExchange,
		KuCoinExchange,
		GateIOExchange,
		HitBTCExchange,
		OKExExchange,
		BinanceExchange,
		LBankExchange,
		QuoineExchange,
		LoopringExchange,
		SimexExchange,
		CREX24Exchange,
		STEXExchange,

		BancorExchange,
		CurveFIExchange,
		GnosisExchange,
		DforceExchange,
		KyberExchange,
		MakerExchange,
		PanCakeSwap,
		SushiSwapExchange,
		UniswapExchange,
		UniswapExchangeV3,
		UbeswapExchange,

		ZeroxExchange,
		BalancerExchange,
		DfynNetwork,
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
