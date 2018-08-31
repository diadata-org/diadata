package dia

const (
	KrakenExchange   = "Kraken"
	BitfinexExchange = "Bitfinex"
	BinanceExchange  = "Binance"
	CoinBaseExchange = "CoinBase"
	HitBTCExchange   = "HitBTC"
	SimexExchange    = "Simex"
	UnknownExchange  = "Unknown"
)

func Exchanges() []string {
	return []string{
		KrakenExchange,
		BitfinexExchange,
		BinanceExchange,
		CoinBaseExchange,
		HitBTCExchange,
		SimexExchange}
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
