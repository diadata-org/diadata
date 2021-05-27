package dia

import (
	"os/user"
	"strings"

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
		BinanceExchange,
		KuCoinExchange,
		// BalancerExchange,
		// BancorExchange,
		// BitBayExchange,
		// BitfinexExchange,
		BitMaxExchange,
		BittrexExchange,
		CoinBaseExchange,
		CREX24Exchange,
		// CurveFIExchange,
		// DforceExchange,
		// GateIOExchange,
		// GnosisExchange,
		HitBTCExchange,
		// HuobiExchange,
		KrakenExchange,
		// KyberExchange,
		// LBankExchange,
		// LoopringExchange,
		// MakerExchange,
		// OKExExchange,
		// PanCakeSwap,
		// QuoineExchange,
		// SimexExchange,
		STEXExchange,
		// SushiSwapExchange,
		UniswapExchange,
		// ZBExchange,
		// ZeroxExchange,
		// UnknownExchange,
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
	var configApi ConfigApi
	usr, _ := user.Current()
	dir := usr.HomeDir
	configFileApi := dir + "/config/secrets/api_" + strings.ToLower(exchange)
	err := gonfig.GetConf(configFileApi, &configApi)
	return &configApi, err
}
