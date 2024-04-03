package scrapers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
)

const (
	refreshDelay = time.Second * 20 * 60
	ecbRSSURL    = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"
)

type (
	XMLHistoricalEnvelope struct {
		XMLName xml.Name `xml:"GenericData"`
		Obs     []XMLObs `xml:"DataSet>Series>Obs"`
	}

	XMLObs struct {
		XMLName   xml.Name        `xml:"Obs"`
		Timestamp XMLObsDimension `xml:"ObsDimension"`
		Price     XMLObsValue     `xml:"ObsValue"`
	}

	XMLObsDimension struct {
		XMLName xml.Name `xml:"ObsDimension"`
		Value   string   `xml:"value,attr"`
	}

	XMLObsValue struct {
		XMLName xml.Name `xml:"ObsValue"`
		Value   string   `xml:"value,attr"`
	}
)

type (
	// rssDocument defines the fields associated with the rss document.

	XMLCube struct {
		XMLName  xml.Name `xml:"Cube"`
		Currency string   `xml:"currency,attr"`
		Rate     string   `xml:"rate,attr"`
	}

	XMLCubeTime struct {
		XMLName xml.Name  `xml:"Cube"`
		Time    string    `xml:"time,attr"`
		Cube    []XMLCube `xml:"Cube"`
	}

	XMLEnvelope struct {
		XMLName  xml.Name      `xml:"Envelope"`
		CubeTime []XMLCubeTime `xml:"Cube>Cube"`
	}
)

type ECBScraper struct {
	// signaling channels
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock    sync.RWMutex
	error        error
	closed       bool
	pairScrapers map[string]*ECBPairScraper // dia.ExchangePair -> pairScraperSet
	ticker       *time.Ticker
	datastore    models.Datastore
	chanTrades   chan *dia.Trade
}

// SpawnECBScraper returns a new ECBScraper initialized with default values.
// The instance is asynchronously scraping as soon as it is created.
func SpawnECBScraper(datastore models.Datastore) *ECBScraper {
	s := &ECBScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*ECBPairScraper),
		error:        nil,
		ticker:       time.NewTicker(refreshDelay),
		datastore:    datastore,
		chanTrades:   make(chan *dia.Trade),
	}

	log.Info("Scraper is built and initiated")
	go s.mainLoop()
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (s *ECBScraper) mainLoop() {
	err := s.Update()
	if err != nil {
		log.Error(err)
	}
	for {
		select {
		case <-s.ticker.C:
			err := s.Update()
			if err != nil {
				log.Error(err)
			}
		case <-s.shutdown: // user requested shutdown
			log.Println("ECBScraper shutting down")
			s.cleanup(nil)
			return
		}
	}
}

// Update performs a HTTP Get request for the rss feed and decodes the results.
func (s *ECBScraper) Update() error {

	log.Printf("Executing ECBScraper update")

	// Retrieve the rss feed document from the web.
	resp, err := http.Get(ecbRSSURL) //nolint:noctx,gosec
	if err != nil {
		return err
	}

	// Close the response once we return from the function.
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	// Check the status code for a 200 so we know we have received a
	// proper response.
	if resp.StatusCode != 200 {
		return fmt.Errorf("HTTP Response Error %d", resp.StatusCode)
	}

	// Decode the rss feed document into our struct type.
	// We don't need to check for errors, the caller can do this.
	var document XMLEnvelope
	err = xml.NewDecoder(resp.Body).Decode(&document)
	if err != nil {
		fmt.Println(err)
	}

	for _, valueCubeTime := range document.CubeTime {
		change := &models.Change{
			USD: []models.CurrencyChange{},
		}

		euroDollar := 1.0
		for _, valueCube := range valueCubeTime.Cube {
			if valueCube.Currency == "USD" {
				euroDollar, err = strconv.ParseFloat(valueCube.Rate, 64)
				if err != nil {
					return fmt.Errorf("error parsing rate %s: %w", valueCube.Rate, err)
				}
			}
		}

		for _, valueCube := range valueCubeTime.Cube {
			pair := string("EUR" + valueCube.Currency)
			ps := s.pairScrapers[pair]
			if ps != nil {
				var rate float64
				var timestamp time.Time
				rate, err = strconv.ParseFloat(valueCube.Rate, 64)
				if err != nil {
					return fmt.Errorf("error parsing rate %s: %w", valueCube.Rate, err)
				}
				timestamp, err = time.Parse("2006-01-02", valueCubeTime.Time)
				if err != nil {
					return fmt.Errorf("error parsing time %s: %w", valueCubeTime.Time, err)
				}

				t := &dia.Trade{
					Pair:   pair,
					Symbol: pair,
					Price:  rate,
					Volume: 0,
					Time:   timestamp,
					Source: "ECB",
				}

				log.Printf("writing trade %#v ", t.Pair)

				s.chanTrades <- t
				c := valueCube.Currency
				if c == "USD" {
					change.USD = append(change.USD, models.CurrencyChange{
						Symbol:        "EUR",
						Rate:          1.0 / euroDollar,
						RateYesterday: 1.0 / euroDollar, // TOFIX
					})
				} else {
					// list for coinhub
					if (c == "JPY") || c == "GBP" || c == "SEK" || c == "CHF" || c == "NOK" || c == "AUD" || c == "CAD" || c == "CNY" || c == "KRW" {
						change.USD = append(change.USD, models.CurrencyChange{
							Symbol:        c,
							Rate:          rate / euroDollar,
							RateYesterday: rate / euroDollar, // TOFIX
						})
					}
				}
			}
		}
		err = s.datastore.SetCurrencyChange(change)
		if err != nil {
			return err
		}
	}
	err = s.datastore.ExecuteRedisPipe()
	if err != nil {
		log.Error("execute redis pipe: ", err)
	}

	s.datastore.FlushRedisPipe()
	if err != nil {
		log.Error("flush redis pipe: ", err)
	}
	log.Info("Update done")
	return err
}

// Populate fetches historical daily datas from 1999 until today and saves them on the database
func Populate(datastore *models.DB, rdb *models.RelDB, pairs []string) {
	// Start with USD to have conversion reference
	xmlEurusd := populateCurrency(datastore, rdb, "USD", nil)

	// Populate every other currency
	for _, p := range pairs {
		currency := p[3:]
		if currency != "USD" {
			populateCurrency(datastore, rdb, currency, xmlEurusd)
		}
	}
}

func populateCurrency(datastore *models.DB, rdb *models.RelDB, currency string, xmlEurusd *XMLHistoricalEnvelope) *XMLHistoricalEnvelope {
	var asset dia.Asset
	var err error
	if currency == "USD" {
		// TO DO: fiat assets have yet to be filled into the asset table
		// by adding an asset source for fiat currencies in the asset service.
		asset, err = rdb.GetFiatAssetBySymbol("EUR")
		if err != nil {
			log.Errorf("fetching fiat asset %s: %v", "EUR", err)
			return &XMLHistoricalEnvelope{}
		}
	} else {
		asset, err = rdb.GetFiatAssetBySymbol(currency)
		if err != nil {
			log.Errorf("fetching fiat asset %s: %v", currency, err)
			return &XMLHistoricalEnvelope{}
		}
	}
	log.Printf("Historical prices population starting for %s\n", asset.Symbol)
	time.Sleep(5 * time.Second)

	// Fetch URL
	resp, err := http.Get(fmt.Sprintf("https://sdw-wsrest.ecb.europa.eu/service/data/EXR/D.%s.EUR.SP00.A", currency)) //nolint:noctx,gosec

	if err != nil {
		log.Errorf("error fetching url %v\n", err)
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	// Parse XML in response
	var xmlSheet XMLHistoricalEnvelope
	err = xml.NewDecoder(resp.Body).Decode(&xmlSheet)
	if err != nil {
		log.Errorf("error parsing xml %v\n", err)
	}

	// Format each value as a fiatQuotation struct and put them into the fqs slice
	var quotations []*models.AssetQuotation
	for _, o := range xmlSheet.Obs {
		if o.Price.Value == "NaN" {
			continue
		}
		var timestamp time.Time
		var price float64
		timestamp, err = time.Parse("2006-01-02", o.Timestamp.Value)
		if err != nil {
			log.Errorf("error formating timestamp %v\n", err)
		}
		price, err = strconv.ParseFloat(o.Price.Value, 64)
		if err != nil {
			log.Errorf("error parsing price %v %v", o.Price.Value, err)
		}

		if currency != "USD" {
			// If other than USD, conversion from EUR as a quote currency to USD as base currency is made
			var usdFor1Euro float64
			for _, eurusdObs := range xmlEurusd.Obs {
				if eurusdObs.Timestamp.Value == o.Timestamp.Value {
					usdFor1Euro, err = strconv.ParseFloat(eurusdObs.Price.Value, 64)
					if err != nil {
						log.Errorf("error parsing price %v %v", eurusdObs.Price.Value, err)
					}
				}
			}
			if usdFor1Euro == 0 {
				continue
			}
			price = usdFor1Euro / price
		}

		assetquotation := models.AssetQuotation{
			Asset:  asset,
			Price:  price,
			Source: dia.Diadata,
			Time:   timestamp,
		}
		quotations = append(quotations, &assetquotation)
	}
	err = datastore.AddAssetQuotationsToBatch(quotations)
	if err != nil {
		log.Errorf("add quotation to batch: %v", err)
	}
	// Write quotations on influxdb
	err = datastore.WriteBatchInflux()
	if err != nil {
		log.Errorf("asset quotations batch write: %v", err)
	} else {
		log.Printf("historical prices for %s successfully populated\n", currency)
	}

	return &xmlSheet
}

// closes all connected PairScrapers
// must only be called from mainLoop
func (s *ECBScraper) cleanup(err error) {

	s.errorLock.Lock()
	defer s.errorLock.Unlock()

	s.ticker.Stop()

	if err != nil {
		s.error = err
	}
	s.closed = true

	close(s.shutdownDone) // signal that shutdown is complete
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *ECBScraper) Close() error {
	if s.closed {
		return errors.New("ECBScraper: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// ECBPairScraper implements PairScraper for ECB
type ECBPairScraper struct {
	parent *ECBScraper
	pair   dia.ExchangePair
	closed bool
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *ECBScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {

	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("ECBScraper: Call ScrapePair on closed scraper")
	}
	ps := &ECBPairScraper{
		parent: s,
		pair:   pair,
	}

	s.pairScrapers[pair.Symbol] = ps

	return ps, nil
}

// Channel returns a channel that can be used to receive trades/pricing information
func (ps *ECBScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

func (ps *ECBPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *ECBPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *ECBPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
