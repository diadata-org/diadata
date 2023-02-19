package scrapers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/diadata-org/exchange-scrapers/pkg/dia"
)

const bitstampAPIURL = "https://www.bitstamp.net/api/v2"

type BitstampScraper struct {
	tradeChannel chan *dia.Trade
	stopChannel  chan struct{}
	doneChannel  chan struct{}
}

func NewBitstampScraper() *BitstampScraper {
	return &BitstampScraper{
		tradeChannel: make(chan *dia.Trade),
		stopChannel:  make(chan struct{}),
		doneChannel:  make(chan struct{}),
	}
}

func (s *BitstampScraper) Channel() chan *dia.Trade {
	return s.tradeChannel
}

func (s *BitstampScraper) Close() error {
	close(s.stopChannel)
	<-s.doneChannel
	return nil
}

func (s *BitstampScraper) FetchAvailablePairs() ([]dia.Pair, error) {
	resp, err := http.Get(bitstampAPIURL + "/trading-pairs-info/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	jsonObj, err := simplejson.NewJson(body)
	if err != nil {
		return nil, err
	}

	var pairs []dia.Pair

	for _, pairJSON := range jsonObj.MustArray() {
		pair := dia.Pair{
			Symbol: pairJSON.Get("name").MustString(),
		}
		pairs = append(pairs, pair)
	}

	return pairs, nil
}

func (s *BitstampScraper) ScrapePair(pair dia.Pair) (dia.PairScraper, error) {
	return &BitstampPairScraper{
		tradeChannel: s.tradeChannel,
		pair:         pair,
	}, nil
}

type BitstampPairScraper struct {
	tradeChannel chan *dia.Trade
	pair         dia.Pair
}

func (s *BitstampPairScraper) StreamTrades() error {
	resp, err := http.Get(fmt.Sprintf("%s/transactions/%s", bitstampAPIURL, s.pair.Symbol))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	for {
		select {
		case <-time.After(time.Minute):
			return nil
		default:
			trade := &dia.Trade{}
			if err := dec.Decode(trade); err != nil {
				return err
			}
			trade.Exchange = "Bitstamp"
			trade.Pair = s.pair
			trade.Timestamp = time.Unix(trade.TimestampUnix, 0)
			s.tradeChannel <- trade
		}
	}
}

func (s *BitstampPairScraper) Close() error {
	return nil
}
