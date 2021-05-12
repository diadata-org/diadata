package dia

import (
	"os/user"
	"strings"
	"time"

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

const (
	Bitcoin  = "Bitcoin"
	Ethereum = "Ethereum"
)

func Exchanges() []string {
	return []string{
		KuCoinExchange,
		UniswapExchange,
		BalancerExchange,
		MakerExchange,
		GnosisExchange,
		CurveFIExchange,
		BinanceExchange,
		BitBayExchange,
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
		LoopringExchange,
		SushiSwapExchange,
		DforceExchange,
		ZeroxExchange,
		KyberExchange,
		BitMaxExchange,
		PanCakeSwap,
		CREX24Exchange,
		STEXExchange,
	}
}

type ConfigApi struct {
	ApiKey    string
	SecretKey string
}

type ConfigConnector struct {
	Coins []Pair
}

type BlockChain struct {
	Name                  string
	GenesisDate           time.Time
	NativeToken           string
	VerificationMechanism VerificationMechanism
}

func GetConfig(exchange string) (*ConfigApi, error) {
	var configApi ConfigApi
	usr, _ := user.Current()
	dir := usr.HomeDir
	configFileApi := dir + "/config/secrets/api_" + strings.ToLower(exchange)
	err := gonfig.GetConf(configFileApi, &configApi)
	return &configApi, err
}
