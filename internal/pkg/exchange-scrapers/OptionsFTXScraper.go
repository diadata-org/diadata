package scrapers

import (
	"fmt"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/nazariyv/go-ftx/auth"
	"github.com/nazariyv/go-ftx/rest"
	"github.com/nazariyv/go-ftx/rest/private/options"
	"github.com/nazariyv/go-ftx/types"
	log "github.com/sirupsen/logrus"
)

var exists nothing

// FTXOptionsScraper - acts as a container for all of the
// <code>FTXOptionScraper</code>s. Note the absence of character
// 's' in the name. They are two different struct types.
type FTXOptionsScraper struct {
	// REST related
	// createRequtestURL string
	// checkQuoteURL     string
	client *rest.Client
	// this contains key and secret
	cfg *dia.ConfigApi
	ds  *models.DB
	// store last error
	err error
	// a set of *FTXOptionScrapers
	scrapers sync.Map
	// each FTXOptionScraper will collect the OB data at this frequency. Alternatively, you might want to refactor this to allow per Scraper frequency granularity
	scrapeFrequency int64
	// channel to listen for a shutdown signal
	shutdown chan nothing
}

type activeRequests struct {
	buySide  int
	sellSide int
}

// FTXOptionScraper - gives us meta information about the running options scraper.
// Note its similarity to FTXRequest. The only difference is the
// Expiry field. Here it is the expiry of the option contract,
// whereas, in FTXRequest, it is the expiry of the quote we created.
type FTXOptionScraper struct {
	client *rest.Client
	// Request for quote
	Request FTXRequest
	// Active requests ids. There will be 2 at any point in time. One for Buys and one for Sells
	ActiveRequests activeRequests
	// flag to signal if the scraper is closed
	closed bool

	err error

	shutdown chan nothing
}

// FTXRequest describes the option for which we initiate the scraper
type FTXRequest struct {
	Underlying    string
	Type          dia.OptionType
	Strike        float64
	RequestExpiry time.Time // REQUEST EXPIRY epoch time
	OptionExpiry  time.Time // OPTION EXPIRY epoch time
	Size          float64
}

// NewFTXOptionsScraper - call this to get back the address of the instance of FTXOptionsScraper.
// Next step is to call Scrape on this instance
func NewFTXOptionsScraper(sf int64) (f *FTXOptionsScraper, err error) {
	// gets FTX api and secret
	var cfg *dia.ConfigApi
	cfg, err = dia.GetConfig(dia.FTXExchange)
	if err != nil {
		return nil, err
	}
	// influxdb instance
	var ds *models.DB
	ds, err = models.NewDataStore()
	if err != nil {
		return nil, err
	}
	f = &FTXOptionsScraper{
		client:          rest.New(auth.New(cfg.ApiKey, cfg.SecretKey)),
		cfg:             cfg,
		err:             nil,
		ds:              ds,
		scrapeFrequency: sf,
		shutdown:        make(chan nothing),
	}
	go f.mainLoop()
	return
}

// Scrape - responsible for scraping bid-ask spread
func (f *FTXOptionsScraper) Scrape(fr FTXRequest) {
	// you may want to tweak this frequency depending on how many instances of FTXOptionScraper are running at the same time
	tick := time.NewTicker(time.Duration(f.scrapeFrequency) * time.Second)
	defer tick.Stop()
	// we instantiate the FTXOptionScraper and it to f.scrapers
	fo := &FTXOptionScraper{
		client:   rest.New(auth.New(f.cfg.ApiKey, f.cfg.SecretKey)),
		Request:  fr,
		closed:   false,
		err:      nil,
		shutdown: make(chan nothing),
	}
	f.scrapers.Store(fo, exists)
	// saving option meta to the db
	optionMeta := dia.OptionMeta{
		InstrumentName: fr.Underlying,
		BaseCurrency:   fr.Underlying,
		ExpirationTime: fr.OptionExpiry,
		StrikePrice:    fr.Strike,
		OptionType:     fr.Type,
	}
	f.ds.SetOptionMeta(&optionMeta)
	// containers for best buy and sell quotes
	var bestBuyQuote options.Quote
	var bestSellQuote options.Quote
	// container for writing to the db
	var orderbookEntry dia.OptionOrderbookDatum

	// and now, every f.scrapeFrequency seconds we either
	// (a) create a RFQ (request for quote) and scrape bid-ask spread and store to f.ds
	// (b) scrape the bid-ask spread from an existing RFQ
	// we create RFQ when we have (i) not created it before, (ii) it has expired
	for {
		select {
		case <-fo.shutdown:
			logWith(fmt.Sprintf("closing %+v scraper. last error: %+v", fo.Request, fo.err), "error")
			return
		case <-tick.C:
			now := time.Now()
			// if the option expiration is past now, there is nothing to scrape
			if now.After(fo.Request.OptionExpiry) {
				fo.close(fmt.Errorf("option %+v expired, closing the scraper", fo.Request))
			}
			// if the request expiry is past now, we need to create a new request
			// this condition will be met infrequently, sine most of the time
			// our request will be active for a very long time. At least for as long
			// as the option expiration. This variable is set by us, so to avoid
			// spamming FTX with a lot of create request requests, we will just create
			// one with an expiration time very far out in the future
			if now.After(fo.Request.RequestExpiry) {
				b, e := fo.createRequest(types.BUY)
				fo.close(e) // only closes if the error is non-nil
				logWith(fmt.Sprintf("new request for %+v\nresponse: %+v", fo.Request, b), "info")
				s, e := fo.createRequest(types.SELL)
				fo.close(e)
				logWith(fmt.Sprintf("new request for %+v\nresponse: %+v", fo.Request, s), "info")
				fo.Request.RequestExpiry = b.RequestExpiry                        // don't use s here, because its expiry will be LATER and so you might run into an issue where you are either using an old quote for the buy side, or you are simply not getting a quote for a fraction of a second
				fo.ActiveRequests = activeRequests{buySide: b.ID, sellSide: s.ID} // these expire in 5 minutes all the time
				continue
			}

			// otherwise get the quote and save it to the db
			if fo.ActiveRequests.buySide == 0 || fo.ActiveRequests.sellSide == 0 {
				logWith(fmt.Sprintf("no quotes for %v or %v", fo.ActiveRequests.buySide, fo.ActiveRequests.sellSide), "info")
				continue
			}
			b, e := fo.client.MyOpQuoteRequest(&options.RequestForMyOpQuoteRequest{RequestID: fo.ActiveRequests.buySide})
			fo.close(e)
			logWith(fmt.Sprintf("buy quote is %+v", b), "debug")
			s, e := fo.client.MyOpQuoteRequest(&options.RequestForMyOpQuoteRequest{RequestID: fo.ActiveRequests.sellSide})
			fo.close(e)
			logWith(fmt.Sprintf("sell quote is %+v", s), "debug")

			if len(*b) < 1 || len(*s) < 1 {
				continue
			}

			// quotes are returned in descending order. this is fine for sellQuote, since we want the highest price for our inventory
			// but **NOT** OK for buyQuote, since we want to buy at cheapest price, thus we access slice at the last idx :)
			bestBuyQuote = (*b)[len(*b)-1]
			bestSellQuote = (*s)[0]

			// store to the db
			orderbookEntry = dia.OptionOrderbookDatum{
				InstrumentName:  fo.Request.Underlying,
				ObservationTime: time.Now().UTC(),
				// ! CAUTION: the above prices are only effective if you post collateral defined in those quotes
				// ! we are not saving that into the db here
				AskPrice: bestBuyQuote.Price,  // we want to buy, that means that the quote is on the ask side
				BidPrice: bestSellQuote.Price, // we want to sell, that means that the quote is on the bid side
				AskSize:  fo.Request.Size,
				BidSize:  fo.Request.Size,
			}
			logWith(fmt.Sprintf("writing to the db: %+v", orderbookEntry), "info")
			e = f.ds.SaveOptionOrderbookDatumInflux(orderbookEntry)
			fo.close(e)
		}
	}
}

func (f *FTXOptionsScraper) mainLoop() {
	for {
		select {
		case <-f.shutdown:
			logWith("closing scrapers", "warn")
			f.cleanup(nil)
			return
		}
	}
}

// cleanup - loops through all of the scrapers in f.scrapers and closes them
func (f *FTXOptionsScraper) cleanup(err error) (errors []error) {
	// if cleanup called with an actual error, add it to errors
	if err != nil {
		errors = append(errors, err)
	}
	// close each of the individual options scrapers. we are using sync.Map to guard against concurrent access to the set of pointers to `FTXOptionScraper`s
	f.scrapers.Range(func(k interface{}, v interface{}) bool {
		return closeScraper(k.(*FTXOptionScraper), f, &errors)
	})
	// if there were any errors during cleanup, return all of them
	if len(errors) > 0 {
		logWith("unsuccessful cleanup", "warn")
		return
	}
	logWith("successful cleanup", "info")
	return nil // explicitly indicating that at this point there are no errors
}

// Close will close all of the individual scrapers
func (f *FTXOptionsScraper) Close() {
	logWith("closing all individual scrapers", "info")
	close(f.shutdown)
}

// Close will close a particular scraper
func (fo *FTXOptionScraper) Close() (err error) {
	logWith(fmt.Sprintf("closing %+v", fo.Request), "info")
	// no need to close the active requests, since they automatically expire in 5 minutes
	// if however, you need to close them, then you can simply loop through them and issue
	// a close request using github.com/nazariyv/go-ftx lib
	if fo.closed {
		return
	}
	fo.close(err)
	return fo.err
}

// functions that improve the readability of the code

func (fo *FTXOptionScraper) createRequest(s string) (r *options.ResponseForCreateOpQuoteRequest, e error) {
	r, e = fo.client.CreateOpQuoteRequest(&options.RequestForCreateOpQuoteRequest{
		Underlying: fo.Request.Underlying,
		Type:       optionType(fo.Request.Type),
		Strike:     fo.Request.Strike,
		Expiry:     fo.Request.OptionExpiry.Unix(),
		// RequestExpiry: re, // we are setting request expiry fairly into the future so that we don't have to bombard FTX with unnecessary requests
		Side: s,
		Size: fo.Request.Size,
	})
	return
}

func closeScraper(s *FTXOptionScraper, f *FTXOptionsScraper, errors *[]error) bool {
	err := s.Close()
	if err != nil {
		*errors = append(*errors, err)
	}
	f.scrapers.Delete(s)
	return true
}

func optionType(ot dia.OptionType) string {
	if ot == 1 {
		return "call"
	}
	return "put"
}

func logWith(msg string, level string) {
	flds := log.WithFields(log.Fields{
		"exchange": dia.FTXExchange,
		"type":     "options",
	})
	switch level {
	case "debug":
		flds.Debugln(msg)
	case "info":
		flds.Infoln(msg)
	case "warn":
		flds.Warnln(msg)
	case "error":
		flds.Errorln(msg)
	default:
		return
	}
}

func (fo *FTXOptionScraper) close(e error) {
	if e == nil {
		return
	}
	defer close(fo.shutdown)
	fo.err = e
	fo.closed = true
	return
}

// * have not tested this, but it might be working
// func (f *FTXOptionsScraper) FetchAvailable() string {
// 	market, err := f.client.Markets(&markets.RequestForMarkets{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// products List
// fmt.Printf("%+v\n", strings.Join(market.List(), "\n"))
// product ranking by USD
// fmt.Printf("%+v\n", strings.Join(market.Ranking(markets.ALL), "\n"))
// 	return strings.Join(market.Ranking(markets.ALL), "\n")
// }
