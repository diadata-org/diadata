package scrapers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	writers "github.com/diadata-org/diadata/internal/pkg/scraper-writers"
)

// DeribitOptionsScraper - used to maintain the order book and save it every x seconds
type DeribitOptionsScraper struct {
	deribitScraper   DeribitScraper
	collectMetaEvery int16 // minutes, polls the rest api to see if there are any new options and saves any new options to the same file
	optionsWaitGroup *sync.WaitGroup
}

type deribitInstrument struct {
	InstrumentName      string  `json:"instrument_name"`
	Kind                string  `json:"kind"`
	TickSize            float64 `json:"tick_size"`
	TakerCommission     float64 `json:"taker_commission"`
	Strike              float64 `json:"strike"`
	SettlementPeriod    string  `json:"settlement_period"`
	QuoteCurrency       string  `json:"quote_currency"`
	OptionType          string  `json:"option_type"`
	MinTradeAmount      float64 `json:"min_trade_amount"`
	MakerCommission     float64 `json:"maker_commission"`
	IsActive            bool    `json:"is_active"`
	ExpirationTimestamp int64   `json:"expiration_timestamp"`
	CreationTimestamp   int64   `json:"creation_timestamp"`
	ContractSize        float64 `json:"contract_size"`
	BaseCurrency        string  `json:"base_currency"`
}

type deribitOrderbookDatum struct {
	Params struct {
		Data struct {
			Timestamp      int64         `json:"timestamp"`
			InstrumentName string        `json:"instrument_name"`
			ChangeID       int64         `json:"change_id"`
			Bids           [1][2]float64 `json:"bids"`
			Asks           [1][2]float64 `json:"asks"`
		} `json:"data"`
	} `json:"params"`
}

type deribitInstruments struct {
	Result []deribitInstrument `json:"result"`
}

const deribitOptionsMetaFilename string = "deribit-options-meta.txt"

// NewDeribitOptionsScraper - returns an instance of an options scraper.
func NewDeribitOptionsScraper(markets []string, accessKey string, accessSecret string) OptionsScraper {
	wg := sync.WaitGroup{}
	writer := writers.FileWriter{}
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	var scraper DeribitScraper = DeribitScraper{
		WaitGroup: &wg,
		Markets:   markets, // e.g. []string{"BTC-20DEC19-5750-C", "BTC-20DEC19-7500-P"}
		Writer:    &writer,
		Logger:    logger,

		AccessKey:    accessKey,
		AccessSecret: accessSecret,

		// expiry is 900 seconds
		RefreshTokenEvery: 800,
		MarketKind:        DeribitOption, // DO NOT change this.
	}

	owg := sync.WaitGroup{}

	var optionsScraper DeribitOptionsScraper = DeribitOptionsScraper{
		deribitScraper:   scraper,
		collectMetaEvery: 6, // hours
		optionsWaitGroup: &owg,
	}

	return &optionsScraper
}

// Authenticate - authenticates
func (s *DeribitOptionsScraper) Authenticate(market string, websocketConnection interface{}) error {
	return s.deribitScraper.Authenticate(market, websocketConnection)
}

// ScraperClose - responsible for closing out the scraper for a market
func (s *DeribitOptionsScraper) ScraperClose(market string, websocketConnection interface{}) error {
	return s.deribitScraper.ScraperClose(market, websocketConnection)
}

// Scrape - scrapes the options markets (meta + order book data)
func (s *DeribitOptionsScraper) Scrape(market string) {
	// calling the futures scraper to collect the options order book data
	go s.deribitScraper.Scrape(market) // this will work forever and it will close the scraper inside of it
	tick := time.NewTicker(time.Duration(s.collectMetaEvery) * time.Hour)
	defer tick.Stop()
	s.GetAndStoreOptionsMeta("BTC")
	s.GetAndStoreOptionsMeta("ETH")
	go func() {
		for {
			select {
			case <-tick.C:
				s.GetAndStoreOptionsMeta("BTC")
				s.GetAndStoreOptionsMeta("ETH")
			}
		}
	}()
}

// ScrapeMarkets - scrapes all the optiosn markets
func (s *DeribitOptionsScraper) ScrapeMarkets() {
	for _, market := range s.deribitScraper.Markets {
		s.optionsWaitGroup.Add(1)
		go s.Scrape(market)
	}
	s.optionsWaitGroup.Wait()
}

// note, this function requires meta to be stored in a file
func metaOnOptionIsAvailable(option deribitInstrument) (bool, error) {
	// O(n) complexity at worst, where n is the number of rows in the file
	// ^ complexity in terms of the file, not the input. O(1) from input perspective.
	// 1. read the file
	file, err := os.Open(deribitOptionsMetaFilename)
	if err != nil {
		return false, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	// 2. check each line whether option is in there
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		optionMeta := &deribitInstrument{}
		err = json.Unmarshal(line, optionMeta)
		if err != nil {
			return false, err
		}
		if *optionMeta == option {
			return true, nil
		}
	}
	if err != nil {
		return false, err
	}
	// 3. return the boolean that signifies its presence or lack of it (true / available)
	return false, nil
}

// GetAndStoreOptionsMeta - will make a REST API call to obtain options meta, it then checks if we already have this data. If we do not, we save the new data.
func (s *DeribitOptionsScraper) GetAndStoreOptionsMeta(market string) error {
	if market != "BTC" && market != "ETH" {
		panic("unsupported deribit market. only btc and eth are supported")
	}
	resp, err := http.Get("https://www.deribit.com/api/v2/public/get_instruments?currency=" + market)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	decodedMsg := deribitInstruments{}
	err = json.Unmarshal(body, &decodedMsg)
	if err != nil {
		return err
	}
	for _, instrument := range decodedMsg.Result {
		if instrument.Kind != "option" {
			continue
		}
		available, err := metaOnOptionIsAvailable(instrument)
		if err != nil {
			switch os.IsNotExist(err) {
			case true: // if the file does not exist, then we clearly have no meta
			case false:
				return err
			}
		}
		if !available {
			line, err := json.Marshal(instrument)
			if err != nil {
				return err
			}
			s.deribitScraper.Logger.Printf("[DEBUG] new option, writing to meta file. option: %s", line)
			_, err = s.deribitScraper.Writer.Write(string(line)+"\n", deribitOptionsMetaFilename)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// ParseOrderbookFile reads the order book data file and converts the order book into an array of OptionOrderbookDatum
func ParseOrderbookFile(path string) ([]OptionOrderbookDatum, error) {
	file, err := os.Open(path)
	if err != nil {
		return []OptionOrderbookDatum{}, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		optionDatum := &deribitOrderbookDatum{}
		err = json.Unmarshal(line, optionDatum)
		if optionDatum.Params.Data.InstrumentName != "" {
			tm := time.Unix(optionDatum.Params.Data.Timestamp/1000, 0) // their unix timestamp is in milliseconds, therefore divide by 1000
			optionOrderbookDatum := OptionOrderbookDatum{
				InstrumentName:  optionDatum.Params.Data.InstrumentName,
				ObservationTime: tm,
				AskPrice:        optionDatum.Params.Data.Asks[0][0],
				BidPrice:        optionDatum.Params.Data.Bids[0][0],
				AskSize:         optionDatum.Params.Data.Asks[0][1],
				BidSize:         optionDatum.Params.Data.Bids[0][1],
			}
			fmt.Println(optionOrderbookDatum)
		}
		if err != nil {
			return []OptionOrderbookDatum{}, err
		}
	}
	if err != nil {
		return []OptionOrderbookDatum{}, err
	}
	return []OptionOrderbookDatum{}, nil
}

// usage example
// func main() {
// 	optionsScraper := scrapers.NewDeribitOptionsScraper([]string{"BTC-27DEC19-7750-C", "BTC-31JAN20-8000-C", "BTC-26JUN20-8000-C", "BTC-27DEC19-12000-C"}, "accessKey", "secretKey")
// 	optionsScraper.ScrapeMarkets()
// }
