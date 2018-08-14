package dia

const (
	KrakenExchange   = "Kraken"
	BitfinexExchange = "Bitfinex"
	BinanceExchange  = "Binance"
	CoinBaseExchange = "CoinBase"
	UnknownExchange  = "Unknown"
)

func (c *ConfigConnector) Exchanges() []string {
	return []string{BinanceExchange, BitfinexExchange, CoinBaseExchange, KrakenExchange, UnknownExchange}
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
