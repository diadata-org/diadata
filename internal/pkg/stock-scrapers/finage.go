package stockscrapers

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"

	"time"

	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/gorilla/websocket"
)

const (
	msciWorldIndexTop10 = "AAPL,MSFT,AMZN,FB,GOOGL,GOOG,TSLA,NVDA,JPM,JNJ"
	subscribeMessage    = "{\"action\": \"subscribe\", \"symbols\":\"" + msciWorldIndexTop10 + "\"}"
	apiKey              = "api_finage"
)

type FinageScraper struct {
	stockScraper               StockScraper
	apiWsURL                   string
	timeResolutionMilliseconds int
}

func NewFinageScraper(db *models.DB) *FinageScraper {

	stockScraper := StockScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		errorLock:    new(sync.RWMutex),
		error:        nil,
		datastore:    db,
		chanStock:    make(chan models.StockQuotation),
		source:       "Finage",
	}
	s := &FinageScraper{
		stockScraper:               stockScraper,
		apiWsURL:                   getAPIKeyFromSecrets(),
		timeResolutionMilliseconds: 1000,
	}
	fmt.Println("scraper built. Start main loop.")
	go s.mainLoop()
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *FinageScraper) mainLoop() {

	// Either call FetchQuotes or implement quote fetching here and
	// leave FetchQuotes blank.

	defer close(scraper.GetStockQuotationChannel())

	c, _, err := websocket.DefaultDialer.Dial(scraper.apiWsURL, nil)
	if err != nil {
		log.Error("connecting to web socket: ", err)
		scraper.cleanup(err)
		return
	}
	defer c.Close()

	if subscribeErr := c.WriteMessage(websocket.TextMessage, []byte(subscribeMessage)); subscribeErr != nil {
		log.Error("creating subscription for the stock quotations: ", err)
		scraper.cleanup(err)
		return
	}

	for {
		select {
		case <-scraper.stockScraper.shutdown:
			scraper.cleanup(nil)
			return
		default:
			_, message, readErr := c.ReadMessage()
			if readErr != nil {
				log.Error("reading the response: ", readErr)
				scraper.cleanup(err)
				return
			}
			if quotation, unmarshalErr := scraper.unmarshalQuotationJSON(message); unmarshalErr != nil {
				log.Error("parsing quotation from response: ", unmarshalErr)
			} else {
				scraper.GetStockQuotationChannel() <- quotation
			}
		}
	}
}

// FetchQuotes fetches quotes from an API and feeds them into the channel.
func (scraper *FinageScraper) FetchQuotes() error {
	// ...
	// scraper.GetStockChannel() <- quote
	return nil
}

// GetDataChannel returns the scrapers data channel.
func (scraper *FinageScraper) GetStockQuotationChannel() chan models.StockQuotation {
	return scraper.stockScraper.chanStock
}

// closes all connected Scrapers. Must only be called from mainLoop
func (scraper *FinageScraper) cleanup(err error) {
	scraper.stockScraper.errorLock.Lock()
	defer scraper.stockScraper.errorLock.Unlock()
	if err != nil {
		scraper.stockScraper.error = err
	}
	scraper.stockScraper.closed = true
	close(scraper.stockScraper.shutdownDone)
}

// Close closes any existing API connections
func (scraper *FinageScraper) Close() error {
	if scraper.stockScraper.closed {
		return errors.New("scraper already closed")
	}
	close(scraper.stockScraper.shutdown)
	<-scraper.stockScraper.shutdownDone
	scraper.stockScraper.errorLock.RLock()
	defer scraper.stockScraper.errorLock.RUnlock()
	return scraper.stockScraper.error
}

// unmarshalQuotationJSON unmarshalls the JSON response into an auxiliary struct
// and converts it into models.StockQuotation
func (scraper *FinageScraper) unmarshalQuotationJSON(data []byte) (models.StockQuotation, error) {
	receivedQuotation := struct {
		Symbol   string  `json:"s"`
		PriceAsk float64 `json:"ap"`
		PriceBid float64 `json:"bp"`
		SizeAsk  float64 `json:"as"`
		SizeBid  float64 `json:"bs"`
		Time     int64   `json:"t"`
	}{}
	err := json.Unmarshal(data, &receivedQuotation)
	if err != nil {
		return models.StockQuotation{}, err
	}
	if len(receivedQuotation.Symbol) == 0 {
		return models.StockQuotation{}, errors.New("not a valid stock quotation")
	}
	name, isin := getCompanyDetails(receivedQuotation.Symbol)
	return models.StockQuotation{
		Symbol:   receivedQuotation.Symbol,
		Name:     name,
		PriceAsk: receivedQuotation.PriceAsk,
		PriceBid: receivedQuotation.PriceBid,
		SizeAsk:  receivedQuotation.SizeAsk,
		SizeBid:  receivedQuotation.SizeBid,
		Source:   scraper.stockScraper.source,
		Time:     time.Unix(0, receivedQuotation.Time*int64(time.Millisecond)),
		ISIN:     isin,
	}, nil
}

func getCompanyDetails(symbol string) (name string, isin string) {
	switch symbol {
	case "AAPL":
		return "APPLE", "US0378331005"
	case "MSFT":
		return "MICROSOFT CORP", "US5949181045"
	case "AMZN":
		return "AMAZON.COM", "US0231351067"
	case "FB":
		return "FACEBOOK A", "US30303M1027"
	case "GOOGL":
		return "ALPHABET A", "US02079K3059"
	case "GOOG":
		return "ALPHABET C", "US02079K1079"
	case "TSLA":
		return "TESLA", "US88160R1014"
	case "NVDA":
		return "NVIDIA", "US67066G1040"
	case "JPM":
		return "JPMORGAN CHASE & CO ", "US46625H1005"
	case "JNJ":
		return "JOHNSON & JOHNSON ", "US4781601046"
	}
	return "", ""
}

// getAPIKeyFromSecrets returns a github api key
func getAPIKeyFromSecrets() string {
	var lines []string
	executionMode := os.Getenv("EXEC_MODE")
	var file *os.File
	var err error
	if executionMode == "production" {
		file, err = os.Open("/run/secrets/" + apiKey)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		file, err = os.Open("../../secrets/" + apiKey)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if len(lines) != 1 {
		log.Fatal("Secrets file should have exactly one line")
	}
	return lines[0]
}
