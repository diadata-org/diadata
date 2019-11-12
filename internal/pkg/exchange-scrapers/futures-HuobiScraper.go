package scrapers

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

// --------------------------------- Config --------------------------------------------------
// ScrapeMarkets - all the markets that you want to scrape
// var scrapeMarkets = []string{"BTC_CQ", "ETH_CQ", "XRP_CQ"}
// or for all markets
var scrapeMarkets = allFuturesMarkets()

// Huobi API configuration
const (
	scrapeDataSaveLocation string = "" // location in which to save. ensure that the folder exists.
	// files will be created automatically

	// API Endpoints
	retryIn   uint8  = 5 // how long to wait in seconds before restarting a failed websocket
	marketURL string = "https://api.hbdm.com"
	wsURL     string = "wss://www.hbdm.com/ws"
)

var allowedMarkets = []string{"BTC", "ETC", "ETH", "EOS", "LTC", "BCH", "XRP", "TRX", "BSV"}
var allowedFrequencies = []string{"CW", "NW", "CQ"}

// --------------------------------------------------------------------------------------------

var buffer bytes.Buffer

// NewHuobiFuturesScraper starts the scraping of all of the futures markets
func NewHuobiFuturesScraper() {
	mainLoop()
}

func allFuturesMarkets() []string {
	allContracts := []string{}
	for _, market := range allowedMarkets {
		for _, maturity := range allowedFrequencies {
			allContracts = append(allContracts, market+"_"+maturity)
		}
	}
	return allContracts
}

// Ensures that the provided market to scrape is supported by Huobi
func validateMarket(market string) {
	parts := strings.Split(market, "_")
	if len(parts) != 2 {
		panic("Incorrect market provided. Should be of the form symbol_frequency")
	}
	tradeSymbol := parts[0]
	tradeFrequency := parts[1]
	containsSymbol := contains(allowedMarkets, tradeSymbol)
	if !containsSymbol {
		panic("Provided trade symbol is not supported")
	}
	containsFreq := contains(allowedFrequencies, tradeFrequency)
	if !containsFreq {
		panic("Provided frequency is not supported")
	}
}

func parseGzip(data []byte) ([]byte, error) {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, data)
	r, err := gzip.NewReader(b)
	if err != nil {
		fmt.Println("[ParseGzip] NewReader error:, maybe data is ungzip:", err, string(data))
		fmt.Println("[ParseGzip] NewReader error:, maybe data is ungzip:", err, string(b.Bytes()))
		return nil, err
	}
	defer r.Close()
	undatas, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println("[ParseGzip] ioutil.ReadAll error: :", err, string(data))
		return nil, err
	}
	return undatas, nil
}

func send(message []byte, market string, ws *websocket.Conn) (int, error) {
	n, err := ws.Write(message)
	fmt.Printf("[DEBUG] [%s] Send: %s\n", market, message)
	return n, err
}

// Pongs back the websocket client to keep the connection alive
func pong(time string, market string, ws *websocket.Conn) (int, error) {
	n, err := send([]byte(fmt.Sprintf("{\"pong\":%s}", time)), market, ws)
	return n, err
}

func appendToFile(filename string, text string) (int, error) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		return 0, err
	}
	n, err := f.WriteString(text)
	if err != nil {
		return n, err
	}
	return n, nil
}

func contains(s []string, str string) bool {
	for _, a := range s {
		if a == str {
			return true
		}
	}
	return false
}

// cleans up at end of WSRun execution
func wsClose(ws *websocket.Conn, wg *sync.WaitGroup, market *string) {
	message := []byte("{\"Unsub\":\"market." + *market + ".trade.detail\"}")
	send(message, *market, ws)
	ws.Close()
	time.Sleep(time.Duration(retryIn) * time.Second)
}

// creates a file name like yyyy-mm-dd-market.txt
func getWriteFileName(market string) string {
	now := time.Now()
	return fmt.Sprintf("%v-%v-%v-%v.txt", now.Year(), int(now.Month()), now.Day(), market)
}

// Scrape starts a websocket scraper for market
func scrape(market string, wg *sync.WaitGroup) {
	validateMarket(market)
	for {
		// IIFE for easy cleanup with defer
		func() {
			ws, err := websocket.Dial(wsURL, "", "http://www.google.com")
			// defer inside of the function will cleanup before the next run
			defer wsClose(ws, wg, &market)
			if err != nil {
				// an error opening is fatal. let this kill the programme
				fmt.Print(fmt.Sprintf("[ERROR] [%s] %v\n", market, err))
				return
			}
			// subscribe to the trade detail channel
			message := []byte("{\"Sub\":\"market." + market + ".trade.detail\"}")
			_, err = send(message, market, ws)
			if err != nil {
				fmt.Print(fmt.Sprintf("Problem subscriping to the [%s] trade channel. %v\n", market, err))
				return
			}
			// create the conduit for the received messages
			var msg = make([]byte, 512)
			for {
				m, err := ws.Read(msg)
				if err != nil {
					fmt.Print(fmt.Sprintf("[ERROR] [%s] %v\n", market, err))
					// an error reading means we may have lost the connection
					// return out and just try again
					return
				}
				newmsg := msg[:m]
				unzipmsg, err := parseGzip(newmsg)
				if err != nil {
					fmt.Print(fmt.Sprintf("[ERROR] [%s] %v\n", market, err))
					return
				}
				fmt.Print(fmt.Sprintf("[DEBUG] [%s] byteLen:%d, unzipLen:%d %s\n", market, m, len(unzipmsg), unzipmsg))
				if len(unzipmsg) == 22 {
					if "ping" == string(unzipmsg[2:6]) {
						_, err := pong(string(unzipmsg[8:21]), market, ws)
						if err != nil {
							fmt.Print(fmt.Sprintf("[ERROR] [%s] Problem ponging the websocket server. %v\n", market, err))
							return
						}
					}
				} else {
					// dynamic writeFileName to change at midnight. "|" so that you can split on it.
					filename := getWriteFileName(market)
					// ensure that scrapeDataSaveLocation exists
					_, err := appendToFile(scrapeDataSaveLocation+filename, string(unzipmsg)+"|")
					if err != nil {
						fmt.Print(fmt.Sprintf("[ERROR] [%s] problem saving to %s. %v\n", market, filename, err))
						return
					}
				}
			}
		}()
	}
	wg.Done()
}

// your entry point
func mainLoop() {
	wg := sync.WaitGroup{}
	for _, market := range scrapeMarkets {
		wg.Add(1)
		go scrape(market, &wg)
	}
	wg.Wait()
}
