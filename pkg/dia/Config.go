package dia

import (
	"errors"
	"os/user"
	"strings"

	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/tkanos/gonfig"
)

const (
	ArthswapExchange          = "Arthswap"
	DiffusionExchange         = "Diffusion"
	OmniDexExchange           = "OmniDex"
	NetswapExchange           = "Netswap"
	TethysExchange            = "Tethys"
	HermesExchange            = "Hermes"
	AnyswapExchange           = "Anyswap"
	BalancerExchange          = "Balancer"
	BalancerV2Exchange        = "BalancerV2"
	BeetsExchange             = "Beets"
	KrakenExchange            = "Kraken"
	BitfinexExchange          = "Bitfinex"
	BitforexExchange          = "Bitforex"
	BinanceExchange           = "Binance"
	BinanceExchangeUS         = "BinanceUS"
	CryptoDotComExchange      = "Crypto.com"
	FTXExchange               = "FTX"
	Opyn                      = "OPYN"
	Premia                    = "Premia"
	BitBayExchange            = "BitBay"
	BittrexExchange           = "Bittrex"
	CoinBaseExchange          = "CoinBase"
	HitBTCExchange            = "HitBTC"
	HuckleberryExchange       = "Huckleberry"
	TraderJoeExchange         = "TraderJoe"
	PangolinExchange          = "Pangolin"
	SimexExchange             = "Simex"
	OKExExchange              = "OKEx"
	HuobiExchange             = "Huobi"
	LBankExchange             = "LBank"
	GateIOExchange            = "GateIO"
	ZBExchange                = "ZB"
	QuoineExchange            = "Quoine"
	UnknownExchange           = "Unknown"
	BlockSizeSeconds          = 120
	FilterKing                = "MAIR120"
	BancorExchange            = "Bancor"
	UniswapExchange           = "Uniswap"
	UniswapExchangeV3         = "UniswapV3"
	UniswapExchangeV3Polygon  = "UniswapV3-polygon"
	LoopringExchange          = "Loopring"
	CurveFIExchange           = "Curvefi"
	MakerExchange             = "Maker"
	KuCoinExchange            = "KuCoin"
	SushiSwapExchange         = "SushiSwap"
	SushiSwapExchangeArbitrum = "SushiSwap-arbitrum"
	SushiSwapExchangePolygon  = "SushiSwap-polygon"
	SushiSwapExchangeFantom   = "SushiSwap-fantom"
	PanCakeSwap               = "PanCakeSwap"
	ApeswapExchange           = "Apeswap"
	BiswapExchange            = "Biswap"
	DforceExchange            = "Dforce"
	ZeroxExchange             = "0x"
	KyberExchange             = "Kyber"
	BitMaxExchange            = "Bitmax"
	CREX24Exchange            = "CREX24"
	STEXExchange              = "STEX"
	Deribit                   = "Deribit"
	DfynNetwork               = "DFYN"
	UbeswapExchange           = "Ubeswap"
	SpookyswapExchange        = "Spookyswap"
	SpiritswapExchange        = "Spiritswap"
	QuickswapExchange         = "Quickswap"
	SerumExchange             = "Serum"
	SolarbeamExchange         = "Solarbeam"
	TrisolarisExchange        = "Trisolaris"
	ByBitExchange             = "ByBit"
	BitMexExchange            = "BitMex"
	MultiChain                = "MultiChain"

	// FinageForex        = "FinageForex"
)

// func Exchanges() []string {
// 	return []string{

// 		BitfinexExchange,
// 		BitMexExchange,
// 		HuobiExchange,
// 		CoinBaseExchange,
// 		GateIOExchange,
// 		HitBTCExchange,
// 		OKExExchange,
// 		BittrexExchange,
// 		KrakenExchange,
// 		KuCoinExchange,
// 		BitBayExchange,
// 		LoopringExchange,
// 		BitforexExchange,
// 		SimexExchange,

// 		BinanceExchange,
// 		BinanceExchangeUS,
// 		LBankExchange,
// 		QuoineExchange,

// 		// FinageForex,
// 		ByBitExchange,
// 		BitMaxExchange,
// 		CryptoDotComExchange,
// 		ZBExchange,
// 		FTXExchange,
// 		CREX24Exchange,
// 		STEXExchange,
// 		UniswapExchangeV3Polygon,

// 		DiffusionExchange,
// 		ArthswapExchange,
// 		ApeswapExchange,
// 		BiswapExchange,
// 		OmniDexExchange,
// 		HermesExchange,
// 		TethysExchange,
// 		TraderJoeExchange,
// 		PangolinExchange,
// 		HuckleberryExchange,
// 		NetswapExchange,
// 		DfynNetwork,
// 		UbeswapExchange,
// 		SpookyswapExchange,
// 		SpiritswapExchange,
// 		SerumExchange,
// 		SolarbeamExchange,
// 		TrisolarisExchange,
// 		AnyswapExchange,

// 		SushiSwapExchange,
// 		SushiSwapExchangeArbitrum,
// 		SushiSwapExchangePolygon,
// 		SushiSwapExchangeFantom,
// 		BeetsExchange,
// 		UniswapExchange,
// 		UniswapExchangeV3,
// 		QuickswapExchange,
// 		PanCakeSwap,

// 		CurveFIExchange,
// 		DforceExchange,
// 		KyberExchange,
// 		MakerExchange,
// 		ZeroxExchange,
// 		BalancerExchange,
// 		BalancerV2Exchange,
// 		BancorExchange,
// 		UnknownExchange,
// 	}
// }

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
