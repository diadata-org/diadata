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
	updateRange     string
	updateInterval  string
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
	yahooFinUpdateRangeEnv    = "YAHOOFIN_UPDATE_RANGE"
	yahooFinUpdateIntervalEnv = "YAHOOFIN_UPDATE_INTERVAL"
	yahooFinCurrenciesMapEnv  = "YAHOOFIN_CURRENCIES_MAP"
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

	// Define the defaults
	updateRange := "120m"
	log.Infof("Default update range set to %s\n", updateRange)
	updateInterval := "1m"
	log.Infof("Default update interval set to %s\n", updateInterval)
	updateFreq := yahooFinUpdateFreqDefault * time.Second
	log.Infof("Default update frequency set to %d seconds\n", yahooFinUpdateFreqDefault)
	currencyMap := make(map[string]string)

	// Read env variables and override defaults if needed
	// TODO: validate range and interval formats (1m, 1h, 1d, 1w, 1mo, 1y)
	yahooFinUpdateRange := utils.Getenv(yahooFinUpdateRangeEnv, "")
	if yahooFinUpdateRange != "" {
		updateRange = yahooFinUpdateRange
		log.Infof("Config set update range to %s\n", updateRange)
	}
	yahooFinUpdateInterval := utils.Getenv(yahooFinUpdateIntervalEnv, "")
	if yahooFinUpdateInterval != "" {
		updateInterval = yahooFinUpdateInterval
		log.Infof("Config set update interval to %s\n", updateInterval)
	}
	yahooFinUpdateFreq, err := strconv.ParseInt(utils.Getenv(yahooFinUpdateFreqEnv, "0"), 10, 64)
	if err != nil {
		log.Errorf("fail to parse %v env variable: %v", yahooFinUpdateFreqEnv, err)
		return
	}
	if yahooFinUpdateFreq != 0 {
		updateFreq = time.Duration(yahooFinUpdateFreq) * time.Second
		log.Infof("Config set update frequency to %f seconds\n", updateFreq.Seconds())
	}
	// Because Yahoo Finance don't have any public endpoint to discover available currency's symbols,
	// we need to scrape webpage to extract metadata. This map is the fallback in case the crawling process fails.
	// Also, some of the pairs are not contained in the webpage. These need to be added to the env var manually.
	// The data was extracted on Jan10 2023, maps the Yahoo Finance symbols to a pair of ISO 4217 friendly format.
	// Use the <YAHOO_SYMBOL>:<AAA-BBB> format separeted by a comma: EURGBP=X:EUR-GBP,CNY=X:USD-CNY
	currenciesList := utils.Getenv(yahooFinCurrenciesMapEnv, "")
	if currenciesList != "" {
		currenciesListSplit := strings.Split(currenciesList, ",")
		if len(currenciesListSplit) > 1 {
			currencyMapDefault := make(map[string]string)
			log.Infof("Config set meta-info for %d currencies\n", len(currenciesListSplit))
			for _, c := range currenciesListSplit {
				currency := strings.Split(c, ":")
				if len(currency) != 2 {
					log.Fatal("currency must have 2 identifier: ", currency)
				}
				symbol := currency[0]
				if symbol[len(symbol)-2:] == "=X" {
					if len(symbol) == 5 {
						symbol = "USD" + symbol
					}
				}
				currencyMapDefault[symbol] = currency[1]
				log.Infof("- %s %s", currency[1], symbol)

			}
			currencyMap = currencyMapDefault
		}
	}

	// Crawl currencies webpage to extract metadata
	data, err := yahooCrawlCurrencies()
	if err != nil {
		log.Warnf("Failed to crawl currencies, using default map: %s", err)
	} else {
		log.Infof("Meta information for %d currencies found", len(data))
		for _, currency := range data {
			if _, ok := currencyMap[currency.Symbol]; !ok {
				currencyMap[currency.Symbol] = currency.Name
			}
		}
	}

	// Create the scraper
	s = &YahooFinScraper{
		ticker:          time.NewTicker(updateFreq),
		updateRange:     updateRange,
		updateInterval:  updateInterval,
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
	updateStartTime := time.Now()

	for k := range scraper.currenciesMap {
		chartDataRes, err := scraper.fetchChartData(k)
		if err != nil {
			log.Error("Error fetching chart data: ", err)
			return err
		}
		for _, result := range chartDataRes {
			symbol := result.Meta.Symbol
			if (len(symbol) == 8 || len(symbol) == 5) && symbol[len(symbol)-2:] == "=X" {
				if _, ok := scraper.currenciesMap[symbol]; !ok {
					if len(symbol) == 5 {
						symbol = "USD" + symbol[len(symbol)-5:]
					}
				}

				quoteSymbol := scraper.currenciesMap[symbol]

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
			} else {
				log.Warnf("Warning, the received symbol %s was not parsed, ignoring it", result.Meta.Symbol)
			}
		}
	}

	updateElapsedTime := time.Since(updateStartTime)
	log.Infof("Quotes updated in %f seconds", updateElapsedTime.Seconds())
	return nil
}

// Main loop runs in a goroutine until channel s is closed.
func (scraper *YahooFinScraper) mainLoop() {

	// Update quotes on startup
	log.Infof("Initializing scraper with %d currencies", len(scraper.currenciesMap))
	err := scraper.UpdateQuotation()
	if err != nil {
		log.Error(err)
	}

	// Start main loop with ticker
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
	q.Add("interval", scraper.updateInterval)
	q.Add("range", scraper.updateRange)
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
			if (len(symbol) == 5 || len(symbol) == 8) && symbol[len(symbol)-2:] == "=X" {
				if len(symbol) == 5 {
					symbol = "USD" + symbol
				}
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
			} else {
				log.Warnf("Warning, cannot parse symbol %s", symbol)
			}
		})
	})

	// Visit the currencies webpage
	log.Printf("Crawling currencies metadata")
	err = c.Visit(yahooFinWebCurrencies)
	if err != nil {
		return currencies, err
	}

	return currencies, nil
}
