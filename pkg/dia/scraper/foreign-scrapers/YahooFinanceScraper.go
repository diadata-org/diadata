package foreignscrapers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

type YahooFinScraper struct {
	ticker          *time.Ticker
	foreignScrapper ForeignScraper
	currenciesMap   map[string]string
}

type yahooFinWebCurrency struct {
	Symbol string
	Name   string
}

type yahooFinV6HttpQuoteResp struct {
	QuoteResponse struct {
		Result []struct {
			Symbol             string  `json:"symbol"`
			RegularMarketPrice float64 `json:"regularMarketPrice"`
			RegularMarketTime  int     `json:"regularMarketTime"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"quoteResponse"`
}

type yahooFinV8HttpChartRespResult struct {
	Meta struct {
		Currency             string  `json:"currency"`
		Symbol               string  `json:"symbol"`
		ExchangeName         string  `json:"exchangeName"`
		InstrumentType       string  `json:"instrumentType"`
		FirstTradeDate       int     `json:"firstTradeDate"`
		RegularMarketTime    int     `json:"regularMarketTime"`
		Gmtoffset            int     `json:"gmtoffset"`
		Timezone             string  `json:"timezone"`
		ExchangeTimezoneName string  `json:"exchangeTimezoneName"`
		RegularMarketPrice   float64 `json:"regularMarketPrice"`
		ChartPreviousClose   float64 `json:"chartPreviousClose"`
		PreviousClose        float64 `json:"previousClose"`
		Scale                int     `json:"scale"`
		PriceHint            int     `json:"priceHint"`
		CurrentTradingPeriod struct {
			Pre struct {
				Timezone  string `json:"timezone"`
				Start     int    `json:"start"`
				End       int    `json:"end"`
				Gmtoffset int    `json:"gmtoffset"`
			} `json:"pre"`
			Regular struct {
				Timezone  string `json:"timezone"`
				Start     int    `json:"start"`
				End       int    `json:"end"`
				Gmtoffset int    `json:"gmtoffset"`
			} `json:"regular"`
			Post struct {
				Timezone  string `json:"timezone"`
				Start     int    `json:"start"`
				End       int    `json:"end"`
				Gmtoffset int    `json:"gmtoffset"`
			} `json:"post"`
		} `json:"currentTradingPeriod"`
		TradingPeriods [][]struct {
			Timezone  string `json:"timezone"`
			Start     int    `json:"start"`
			End       int    `json:"end"`
			Gmtoffset int    `json:"gmtoffset"`
		} `json:"tradingPeriods"`
		DataGranularity string   `json:"dataGranularity"`
		Range           string   `json:"range"`
		ValidRanges     []string `json:"validRanges"`
	} `json:"meta"`
	Timestamp  []int `json:"timestamp"`
	Indicators struct {
		Quote []struct {
			Volume []float64 `json:"volume"`
			Open   []float64 `json:"open"`
			High   []float64 `json:"high"`
			Low    []float64 `json:"low"`
			Close  []float64 `json:"close"`
		} `json:"quote"`
	} `json:"indicators"`
}

type yahooFinV8HttpChartResp struct {
	Chart struct {
		Result []yahooFinV8HttpChartRespResult `json:"result"`
		Error  string                          `json:"error"`
	} `json:"chart"`
}

const (
	yahooFinSource            = "YahooFinance"
	yahooFinUpdateFreqDefault = 60 * 2 // Default update frequency (in seconds)
	yahooFinUpdateFreqEnv     = "YAHOOFIN_UPDATE_FREQ"
	yahooFinWebCurrencies     = "https://finance.yahoo.com/currencies"
	yahooFinHttpV10Host       = "https://query1.finance.yahoo.com"
	yahooFinHttpV11Host       = "https://query2.finance.yahoo.com"
	yahooFinTypeCurrency      = "CURRENCY"
	yahooFinV6HttpPathQuote   = "/v6/finance/quote"
	yahooFinV8HttpPathChart   = "/v8/finance/chart"
)

func NewYahooFinScraper(datastore models.Datastore) (s *YahooFinScraper) {
	foreignScrapper := ForeignScraper{
		shutdown:      make(chan nothing),
		error:         nil,
		datastore:     datastore,
		chanQuotation: make(chan *models.ForeignQuotation),
	}

	// Define the update frequency
	updateFreq := yahooFinUpdateFreqDefault * time.Second
	yahooFinUpdateFreqEnv, err := strconv.ParseInt(utils.Getenv(yahooFinUpdateFreqEnv, "0"), 10, 64)
	if err != nil {
		log.Errorf("parse fail to %v env variable: %v", yahooFinUpdateFreqEnv, err)
		return
	}
	if yahooFinUpdateFreqEnv != 0 {
		updateFreq = time.Duration(yahooFinUpdateFreqEnv) * time.Second
	}

	// Because Yahoo Finance don't have any public endpoint to discover available currency's symbols,
	// we need to scrape webpage to extract metadata. This map is the fallback in case the crawling process fails.
	// Also, some of the pairs are not contained in the webpage. These need to be added to the env var manually.
	// The data was extracted on Jan10 2023, maps the Yahoo Finance symbols to a pair of ISO 4217 friendly format.
	// Use the <YAHOO_SYMBOL>:<AAA-BBB> format separeted by a comma: EURGBP=X:EUR-GBP,CNY=X:USD-CNY
	currencyMapDefault := make(map[string]string)
	currenciesList := strings.Split(utils.Getenv("CURRENCIES_LIST_YAHOO", ""), ",")
	log.Infof("Meta info for %d currencies have been loaded into mapping\n", len(currenciesList))
	for _, c := range currenciesList {
		currency := strings.Split(c, ":")
		log.Debugf("- %s %s", currency[1], currency[0])
		if len(currency) != 2 {
			log.Fatal("currency must have 2 identifier: ", currency)
		}
		currencyMapDefault[currency[0]] = currency[1]
	}

	currencyMap := currencyMapDefault

	crawlStartTime := time.Now()
	data, err := yahooCrawlCurrencies()
	if err != nil {
		log.Warnf("Failed to crawl currencies, using default map: %s", err)
	} else {
		updateElapsedTime := time.Since(crawlStartTime)
		log.Infof("Meta information for %d currencies found in %f seconds", len(data), updateElapsedTime.Seconds())
		for _, currency := range data {
			currencyMap[currency.Symbol] = currency.Name
		}
		for symbol, name := range currencyMap {
			log.Debugf("- %s %s", name, symbol)
		}
		if len(currencyMap) < len(currencyMapDefault) {
			for symbol, name := range currencyMapDefault {
				if _, ok := currencyMap[symbol]; !ok {
					currencyMap[symbol] = name
				}
			}
		}
	}

	s = &YahooFinScraper{
		ticker:          time.NewTicker(updateFreq),
		foreignScrapper: foreignScrapper,
		currenciesMap:   currencyMap,
	}

	go s.mainLoop()

	return s
}

// Closes any existing connections
func (scraper *YahooFinScraper) Close() error {
	if scraper.foreignScrapper.closed {
		return errors.New("scraper already closed")
	}
	close(scraper.foreignScrapper.shutdown)
	<-scraper.foreignScrapper.shutdownDone
	scraper.foreignScrapper.errorLock.RLock()
	defer scraper.foreignScrapper.errorLock.RUnlock()
	return scraper.foreignScrapper.error
}

// Returns the channel to which new quotes are pushed
func (scraper *YahooFinScraper) GetQuoteChannel() chan *models.ForeignQuotation {
	return scraper.foreignScrapper.chanQuotation
}

// Retrieves new coin information from the Yahoo Finance API and stores it to influx
func (scraper *YahooFinScraper) UpdateQuotation() error {
	log.Infof("Updating quotes for %d currencies", len(scraper.currenciesMap))
	updateStartTime := time.Now()

	for k := range scraper.currenciesMap {
		chartDataRes, err := scraper.fetchChartData(k)
		if err != nil {
			log.Error("Error fetching chart data: ", err)
			return err
		}
		for _, result := range chartDataRes {
			quoteSymbol := scraper.currenciesMap[result.Meta.Symbol]
			if quoteSymbol == "" {
				log.Warnf("Warning, symbol %s not found in currencies map", result.Meta.Symbol)
			} else {
				for i, quoteUnixTime := range result.Timestamp {
					quoteDateTime := time.Unix(int64(quoteUnixTime), 0)
					quotePrice := result.Indicators.Quote[0].Close[i]
					if quotePrice > 0 {
						priceYesterday, err := scraper.foreignScrapper.datastore.GetForeignPriceYesterday(scraper.currenciesMap[quoteSymbol], yahooFinSource)
						if err != nil {
							priceYesterday = 0
						}
						quote := models.ForeignQuotation{
							Symbol:             scraper.currenciesMap[quoteSymbol],
							Name:               quoteSymbol,
							Price:              quotePrice,
							PriceYesterday:     priceYesterday,
							VolumeYesterdayUSD: 0.0, // Fetched volume data is always 0 (not available)
							Source:             yahooFinSource,
							Time:               quoteDateTime,
						}
						scraper.foreignScrapper.chanQuotation <- &quote
					}
				}
			}
		}
	}

	updateElapsedTime := time.Since(updateStartTime)
	log.Infof("Quotes updated in %f seconds", updateElapsedTime.Seconds())
	return nil
}

// Main loop runs in a goroutine until channel s is closed.
func (scraper *YahooFinScraper) mainLoop() {
	log.Infof("Initializing scraper with %d currencies loaded", len(scraper.currenciesMap))
	err := scraper.UpdateQuotation()
	if err != nil {
		log.Error(err)
	}

	log.Infof("Starting main loop")
	for {
		select {
		case <-scraper.ticker.C:
			err := scraper.UpdateQuotation()
			if err != nil {
				log.Error(err)
			}
		case <-scraper.foreignScrapper.shutdown: // user requested shutdown
			log.Printf("%s scraper shutting down", yahooFinSource)
			return
		}
	}
}

// Fetches chart data for a given symbol
func (scraper *YahooFinScraper) fetchChartData(symbol string) (chart []yahooFinV8HttpChartRespResult, err error) {

	// Prepare the request
	reqUrl := yahooFinHttpV10Host + yahooFinV8HttpPathChart + "/" + symbol
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Errorf("Error %s", err)
		return chart, err
	}

	// Add URL query parameters and encode them
	q := url.Values{}
	q.Add("interval", "1m")
	q.Add("range", "180m")
	req.URL.RawQuery = q.Encode()

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Error, cannot fetch %s", req.URL.String())
		return chart, err
	}

	// Reads the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Error, cannot read the chart data response")
		return chart, err
	}

	// Parse the response as json
	var chartReq yahooFinV8HttpChartResp
	err = json.Unmarshal(body, &chartReq)
	if err != nil {
		log.Errorf("Error, cannot unmarshall chart data")
		return chart, err
	}

	return chartReq.Chart.Result, nil
}

// Crawl currencies webpage and return a slice of currency metadata
func yahooCrawlCurrencies() (currencies []yahooFinWebCurrency, err error) {

	// Instantiate default cralwer collector
	c := colly.NewCollector()

	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode == 200 {
			log.Debugf("%d: %s", r.StatusCode, r.Request.URL)
		} else {
			log.Debugf("%d: %s", r.StatusCode, r.Request.URL)
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Errorln(err)
	})

	// When a HTML element is found with the given selector
	c.OnHTML("#list-res-table > div > table > tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			symbol := el.ChildText("td:nth-child(1)")
			name := el.ChildText("td:nth-child(2)")
			if nameSplit := strings.Split(name, "/"); len(nameSplit) != 2 {
				log.Errorf("Cannot parse name %s", name)
				return
			}
			currency := yahooFinWebCurrency{
				Symbol: symbol,
				Name:   strings.Split(name, "/")[0] + "-" + strings.Split(name, "/")[1],
			}
			log.Debugf("- %s: %s", currency.Name, currency.Symbol)
			currencies = append(currencies, currency)
		})
	})

	// Start scraping
	log.Printf("Crawling currencies metadata from %s", yahooFinWebCurrencies)
	err = c.Visit(yahooFinWebCurrencies)
	if err != nil {
		return currencies, err
	}

	return currencies, nil
}