package optionscrapers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"golang.org/x/time/rate"
	"strconv"
	"sync"
	"time"
)

var logger = logrus.New()

type OKExInstrumentDetails []OKExInstrumentDetail

type OKExInstrumentDetail struct {
	Category           string    `json:"category"`
	ContractVal        string    `json:"contract_val"`
	Delivery           time.Time `json:"delivery"`
	InstrumentID       string    `json:"instrument_id"`
	Listing            time.Time `json:"listing"`
	LotSize            string    `json:"lot_size"`
	OptionType         string    `json:"option_type"`
	SettlementCurrency string    `json:"settlement_currency"`
	State              string    `json:"state"`
	Strike             string    `json:"strike"`
	TickSize           string    `json:"tick_size"`
	Timestamp          time.Time `json:"timestamp"`
	TradingStartTime   time.Time `json:"trading_start_time"`
	Underlying         string    `json:"underlying"`
}

type OKExOptionsScraper struct {
	Markets []string

	// OKEx Options endpoint is a REST one and you are limited to 20 requests per 2 seconds. So you have
	// to throttle your polling frequency. This is in seconds. So you can have at most 10
	// markets running at the same time on the same ip address
	PollFrequency      int8
	ScraperIsRunning   bool
	ScraperIsRunningMu sync.Mutex
	optionsWaitGroup   *sync.WaitGroup
	DataStore          *models.DB
	chanOrderBook      chan *dia.OptionOrderbookDatum
	Ratelimiter        *rate.Limiter
}

type AllOKExOptionsScrapers struct {
	Scrapers []*OKExOptionsScraper
	Markets  []string
	ds       *models.DB
	owg      *sync.WaitGroup
}

type rawOKExOBDatum struct {
	Asks            [][]string `json:"asks"`
	Bids            [][]string `json:"bids"`
	ObservationTime string     `json:"timestamp"`
}

type OKExOptionState int

const (
	PreOpen OKExOptionState = iota + 1
	Live
	Suspended
	Settlement
)

type OKExInstrument struct {
	InstrumentName        string `json:"instrument_id"`
	Underlying            string `json:"underlying"`
	SettlementCurrency    string `json:"settlement_currency"`
	MinTradeAmount        string `json:"contract_val"`
	OptionType            string `json:"option_type"`
	Strike                string `json:"strike"`
	TickSize              string `json:"tick_size"`
	LotSize               string `json:"lot_size"`
	Listing               string `json:"listing"`
	Expiration            string `json:"delivery"`
	State                 string `json:"state"`
	TradingStartTimestamp string `json:"trading_start_time"`
}

type OKExInstruments struct {
	Result []OKExInstrument `json:"result"`
}

//func NewAllOKExOptionsScrapers(owg *sync.WaitGroup, markets []string) (al AllOKExOptionsScrapers) {
//	all := AllOKExOptionsScrapers{}
//	ds, err := models.NewDataStore()
//	if err != nil {
//		logger.WithFields(logrus.Fields{"prefix": "OKEx"}).Error(err)
//		return
//	}
//	for _, market := range markets {
//		newScraper := NewOKExOptionsScraper(ds, market, 1) // ! pollFreq hardcoded to 1 by default
//		all.Scrapers = append(all.Scrapers, &newScraper)
//	}
//	all.ds = ds
//	all.owg = owg
//	return all
//}

func NewOKExOptionsScraper(pollFreq int8) *OKExOptionsScraper {
	rl := rate.NewLimiter(rate.Every(2*time.Second), 10) // 10 request every 2 seconds
	optionsScraper := &OKExOptionsScraper{
		PollFrequency: pollFreq,
		chanOrderBook:  make(chan *dia.OptionOrderbookDatum),
		Ratelimiter: rl,// if pollFreq = 1 second. can have 10 goroutines at the same time
	}
	return optionsScraper
}

func (s *OKExOptionsScraper) parseObDatum(datum *rawOKExOBDatum, market string) (resolvedDatum dia.OptionOrderbookDatum, err error) {
	resolvedDatum = dia.OptionOrderbookDatum{}

	resolvedTime, err := time.Parse(time.RFC3339, datum.ObservationTime)
	if err != nil {
		logger.WithFields(logrus.Fields{"prefix": "OKEx", "market": market}).Error(err)
		return
	}
	var resolvedAskPX float64
	resolvedAskPX, err = strconv.ParseFloat(datum.Asks[0][0], 64)
	if err != nil {
		logger.WithFields(logrus.Fields{"prefix": "OKEx", "market": market}).Error(err)
		return
	}
	var resolvedBidPX float64
if len(datum.Bids)>0 {
	resolvedBidPX, err = strconv.ParseFloat(datum.Bids[0][0], 64)
	if err != nil {
		logger.WithFields(logrus.Fields{"prefix": "OKEx", "market": market}).Error(err)
		return
	}
}
	var resolvedAskSize float64
	resolvedAskSize, err = strconv.ParseFloat(datum.Asks[0][1], 64)
	if err != nil {
		logger.WithFields(logrus.Fields{"prefix": "OKEx", "market": market}).Error(err)
		return
	}
	var resolvedBidSize float64
	if len(datum.Bids)>0 {

		resolvedBidSize, err = strconv.ParseFloat(datum.Bids[0][1], 64)
		if err != nil {
			logger.WithFields(logrus.Fields{"prefix": "OKEx", "market": market}).Error(err)
			return
		}
	}

	resolvedDatum.InstrumentName = market
	resolvedDatum.ObservationTime = resolvedTime
	resolvedDatum.AskPrice = resolvedAskPX
	resolvedDatum.BidPrice = resolvedBidPX
	resolvedDatum.AskSize = resolvedAskSize
	resolvedDatum.BidSize = resolvedBidSize

	return resolvedDatum, nil
}

func (s *OKExOptionsScraper) FetchInstruments() {

	// Get underlying pairs https://www.okex.com/api/option/v3/underlying
	b, err := utils.GetRequest("https://www.okex.com/api/option/v3/underlying")
	if err != nil {
		log.Errorln("Error gettinb underlying assets", err)
	}

	var underlying []string
	json.Unmarshal(b, &underlying)

	for _, pair := range underlying {
		var instruments OKExInstrumentDetails
		log.Println(pair)

		b, err := utils.GetRequest("https://www.okex.com/api/option/v3/instruments/" + pair)
		if err != nil {
			log.Errorln("Error getting instrumentId")
		}

		json.Unmarshal(b, &instruments)

		for _, v := range instruments {
			s.Markets = append(s.Markets, v.InstrumentID)
		}

	}
	return
}
func (s *OKExOptionsScraper) Scrape() {

	for _, market := range s.Markets {
		go s.ScrapeInstrument(market)

	}

}

func (s *OKExOptionsScraper) ScrapeInstrument(market string) {

 	ctx := context.Background()
	err := s.Ratelimiter.Wait(ctx) // This is a blocking call. Honors the rate limit
	if err != nil {
		log.Errorln("Error on ratelimit",err)
		return
	}


	logger.Formatter = new(prefixed.TextFormatter)
	logger.Level = logrus.InfoLevel

	var rawOB = rawOKExOBDatum{}
	url := fmt.Sprintf("https://www.okex.com/api/option/v3/instruments/%s/book?size=1", market)
	log.Infoln("Requesting Url ", url)
	// * change size query param to larger number for greater depth. the largest you can go to is 200
	body, err := utils.GetRequest(url)
	if err != nil {
		logger.WithFields(logrus.Fields{"prefix": "OKEx", "market": market}).Error(err)
		return
	}

	err = json.Unmarshal(body, &rawOB)
	if err != nil {
		logger.WithFields(logrus.Fields{"prefix": "OKEx", "market": market}).Error(err)
		return
	}

	log.Println("rawOB", rawOB)

	var obEntry dia.OptionOrderbookDatum
	obEntry, err = s.parseObDatum(&rawOB, market)
	if err != nil {
		logger.WithFields(logrus.Fields{"prefix": "OKEx", "market": market}).Error(err)
		return
	}

	s.chanOrderBook <- &obEntry

}



func (s *AllOKExOptionsScrapers) MetaOnOptionIsAvailable(option OKExInstrument) (available bool, err error) {
	available = false
	err = nil

	// TODO: can make this faster by specifying BaseCurrency/QuoteCurrency instead
	optionMetas, err := s.ds.GetOptionMeta(option.SettlementCurrency)
	if err != nil {
		return
	}

	for _, optionMeta := range optionMetas {
		if optionMeta.InstrumentName == option.InstrumentName {
			return true, nil
		}
	}

	return
}

func (s *AllOKExOptionsScrapers) GetAndStoreOptionsMeta() (err error) {
	body, err := utils.GetRequest("https://www.okex.com/api/option/v3/instruments/BTC-USD")
	if err != nil {
		return
	}

	var decodedMsg []OKExInstrument
	err = json.Unmarshal(body, &decodedMsg)

	if err != nil {
		return
	}

	for _, instrument := range decodedMsg {
		var available bool
		available, err = s.MetaOnOptionIsAvailable(instrument)

		if err != nil {
			return
		}

		if !available {

			optionType := dia.CallOption
			if instrument.OptionType == "P" {
				optionType = dia.PutOption
			}

			var expTime time.Time
			expTime, err = time.Parse(time.RFC3339, instrument.Expiration)
			if err != nil {
				return
			}

			var strikePrice float64
			strikePrice, err = strconv.ParseFloat(instrument.Strike, 64)
			if err != nil {
				return
			}

			optionMeta := dia.OptionMeta{
				InstrumentName: instrument.InstrumentName,
				BaseCurrency:   instrument.SettlementCurrency,
				ExpirationTime: expTime,
				StrikePrice:    strikePrice,
				OptionType:     optionType,
			}

			s.ds.SetOptionMeta(&optionMeta)
		}
	}

	return nil
}

func (s *OKExOptionsScraper) Channel() chan *dia.OptionOrderbookDatum {
	return s.chanOrderBook
}
