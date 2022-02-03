package dia

import (
	"errors"
	"os/user"
	"strings"

	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/tkanos/gonfig"
)

const (
	NetswapExchange      = "Netswap"
	AnyswapExchange      = "Anyswap"
	BalancerExchange     = "Balancer"
	BalancerV2Exchange   = "BalancerV2"
	GnosisExchange       = "Gnosis"
	KrakenExchange       = "Kraken"
	BitfinexExchange     = "Bitfinex"
	BitforexExchange     = "Bitforex"
	BinanceExchange      = "Binance"
	CryptoDotComExchange = "Crypto.com"
	FTXExchange          = "FTX"
	Opyn                 = "OPYN"
	Premia               = "Premia"
	BitBayExchange       = "BitBay"
	BittrexExchange      = "Bittrex"
	CoinBaseExchange     = "CoinBase"
	HitBTCExchange       = "HitBTC"
	SimexExchange        = "Simex"
	OKExExchange         = "OKEx"
	HuobiExchange        = "Huobi"
	LBankExchange        = "LBank"
	GateIOExchange       = "GateIO"
	ZBExchange           = "ZB"
	QuoineExchange       = "Quoine"
	UnknownExchange      = "Unknown"
	BlockSizeSeconds     = 120
	FilterKing           = "MAIR120"
	BancorExchange       = "Bancor"
	UniswapExchange      = "Uniswap"
	UniswapExchangeV3    = "UniswapV3"
	LoopringExchange     = "Loopring"
	CurveFIExchange      = "Curvefi"
	MakerExchange        = "Maker"
	KuCoinExchange       = "KuCoin"
	SushiSwapExchange    = "SushiSwap"
	PanCakeSwap          = "PanCakeSwap"
	DforceExchange       = "Dforce"
	ZeroxExchange        = "0x"
	KyberExchange        = "Kyber"
	BitMaxExchange       = "Bitmax"
	CREX24Exchange       = "CREX24"
	STEXExchange         = "STEX"
	Deribit              = "Deribit"
	DfynNetwork          = "DFYN"
	UbeswapExchange      = "Ubeswap"
	SpookyswapExchange   = "Spookyswap"
	SpiritswapExchange   = "Spiritswap"
	QuickswapExchange    = "Quickswap"
	SerumExchange        = "Serum"
	SolarbeamExchange    = "Solarbeam"
	TrisolarisExchange   = "Trisolaris"
	ByBitExchange        = "ByBit"
	BitMexExchange       = "BitMex"
	// FinageForex        = "FinageForex"
)

func Exchanges() []string {
	return []string{

		LoopringExchange,
		BitfinexExchange,
		BitforexExchange,
		SimexExchange,

		BinanceExchange,
		LBankExchange,
		QuoineExchange,
		CREX24Exchange,
		STEXExchange,

		// FinageForex,
		ByBitExchange,
		BitMaxExchange,
		CryptoDotComExchange,
		ZBExchange,
		FTXExchange,
		HuobiExchange,
		BitBayExchange,
		BittrexExchange,
		CoinBaseExchange,
		KrakenExchange,
		KuCoinExchange,
		GateIOExchange,
		HitBTCExchange,
		OKExExchange,

		NetswapExchange,
		DfynNetwork,
		UbeswapExchange,
		SpookyswapExchange,
		SpiritswapExchange,
		SerumExchange,
		SolarbeamExchange,
		TrisolarisExchange,
		AnyswapExchange,

		SushiSwapExchange,
		UniswapExchange,
		UniswapExchangeV3,
		QuickswapExchange,
		PanCakeSwap,

		CurveFIExchange,
		GnosisExchange,
		DforceExchange,
		KyberExchange,
		MakerExchange,
		ZeroxExchange,
		BalancerExchange,
		BalancerV2Exchange,
		BancorExchange,
		UnknownExchange,
		BitMexExchange,
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
