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
)

const (
	Bitcoin  = "Bitcoin"
	Ethereum = "Ethereum"
)

func Exchanges() []string {
	return []string{
		BalancerExchange,
		BancorExchange,
		BinanceExchange,
		BitBayExchange,
		BitfinexExchange,
		BitMaxExchange,
		BittrexExchange,
		CoinBaseExchange,
		CurveFIExchange,
		CREX24Exchange,
		DforceExchange,
		GateIOExchange,
		GnosisExchange,
		HitBTCExchange,
		HuobiExchange,
		KrakenExchange,
		KuCoinExchange,
		KyberExchange,
		LBankExchange,
		LoopringExchange,
		MakerExchange,
		OKExExchange,
		PanCakeSwap,
		QuoineExchange,
		SimexExchange,
		STEXExchange,
		SushiSwapExchange,
		UniswapExchange,
		UniswapExchangeV3,
		ZBExchange,
		ZeroxExchange,
		UnknownExchange,
		DfynNetwork,
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
