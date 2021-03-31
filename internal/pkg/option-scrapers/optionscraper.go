package optionscrapers

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

var Exchanges map[string]dia.Exchange
var blockchains map[string]dia.BlockChain

func init() {

	blockchains = make(map[string]dia.BlockChain)
	blockchains[dia.Bitcoin] = dia.BlockChain{Name: dia.BinanceExchange, NativeToken: "BTC", VerificationMechanism: dia.PROOF_OF_WORK}
	blockchains[dia.Ethereum] = dia.BlockChain{Name: dia.BinanceExchange, NativeToken: "ETH", VerificationMechanism: dia.PROOF_OF_WORK}
	// TODO move all this to single json
	Exchanges = make(map[string]dia.Exchange)
	Exchanges[dia.OKExExchange] = dia.Exchange{Name: dia.OKExExchange, Centralized: true}

}

/* OptionsScraper provides common methods needed to get Option orderBook information from
exchange APIs.*/
type OptionsScraper interface {
	//io.Closer
	FetchInstruments()
	Scrape()
	// Channel returns a channel that can be used to receive trades
	Channel() chan *dia.OptionOrderbookDatum
}

func New(exchange string, key string, secret string) OptionsScraper {
	switch exchange {
	case dia.OKExExchange:
		return NewOKExOptionsScraper(int8(30))

	default:
		return nil
	}

}
