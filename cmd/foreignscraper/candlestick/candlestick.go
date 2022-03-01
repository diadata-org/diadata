package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	_ "github.com/diadata-org/diadata/pkg/model"
	ws "github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type candlestickMessage struct {
	ForeignName string
	// ClosingPrice in USDT
	ClosingPrice float64
	Volume       float64
	Timestamp    time.Time
	Source       string
}

func main() {

	wg := sync.WaitGroup{}
	/*ds, err := models.NewDataStore()
	if err != nil {
		log.Fatal("datastore error: ", err)
	}*/

	assets := flag.String("assets", "BTC,ETH", "asset symbols to query (from BTC, ETH, SOL, GLMR, DOT")
	exchanges := flag.String("exchanges", "GateIO", "exchanges to query (from Binance, Kucoin, Coinbase, Huobi, Okex, Gateio")
	flag.Parse()
	cChan := make(chan candlestickMessage)

	for _, exchange := range strings.Split(*exchanges, ",") {
		wg.Add(1)
		go handleExchangeScraper(exchange, *assets, cChan, &wg)
	}
	defer wg.Wait()
	for {
		// TO DO: Handle data every n minutes.
		log.Info("data: ", <-cChan)
	}

}

func handleExchangeScraper(exchange string, assets string, candleChan chan candlestickMessage, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Printf("Entered Exchange handler for %s", exchange)
	switch exchange {
	case "Binance":
		log.Println("Binance Scraper: Start scraping")
		err := scrapeBinance(assets, candleChan)
		if err != nil {
			log.Error(err)
		}
	case "GateIO":
		log.Println("Gateio Scraper: Start scraping")
		err := scrapeGateio(assets, candleChan)
		if err != nil {
			log.Error(err)
		}
	case "Kucoin":
		log.Println("Kucoin Scraper: Start scraping")
	default:
		log.Errorf("Unknown scraper name %s", exchange)
	}
}

func scrapeBinance(assets string, candleChan chan candlestickMessage) error {
	log.Info("Entered Binance handler")
	wsBaseString := "wss://fstream.binance.com/stream?streams="
	wsAssetsString := ""

	for _, asset := range strings.Split(assets, ",") {
		wsAssetsString += strings.ToLower(asset) + "usdt@kline_1m" + "/"
	}
	conn, _, err := ws.DefaultDialer.Dial(wsBaseString+wsAssetsString, nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	//TODO: Send into channel to main handler
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Errorln("read:", err)
			return err
		}
		messageMap := make(map[string]interface{})
		err = json.Unmarshal(message, &messageMap)
		data := messageMap["data"].(map[string]interface{})["k"].(map[string]interface{})

		closingPriceString := data["c"].(string)
		closingPrice, err := strconv.ParseFloat(closingPriceString, 64)
		volumeString := data["V"].(string)
		volume, err := strconv.ParseFloat(volumeString, 64)
		timeUnix := data["t"].(float64)

		candleStickMessage := candlestickMessage{
			ForeignName:  data["s"].(string),
			ClosingPrice: closingPrice,
			Volume:       volume,
			Timestamp:    time.Unix(int64(timeUnix/1000), 0),
			Source:       "Binance",
		}

		candleChan <- candleStickMessage

	}
	return nil
}

func scrapeGateio(assets string, candleChan chan candlestickMessage) error {
	log.Info("Entered GateIO handler")
	wsBaseString := "wss://fx-ws-testnet.gateio.ws/v4/ws/"

	conn, _, err := ws.DefaultDialer.Dial(wsBaseString, nil)
	if err != nil {
		return err
	}
	defer conn.Close()
	for _, asset := range strings.Split(assets, ",") {
		msgToWrite := fmt.Sprintf("{\"time\":1,\"channel\":\"futures.candlesticks\",\"event\":\"subscribe\",\"payload\":[\"1m\",\"%s_USD\"]}", asset)
		log.Println(msgToWrite)
		conn.WriteMessage(ws.TextMessage, []byte(msgToWrite))
	}

	//TODO: Send into channel to main handler
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Errorln("read:", err)
			return err
		}
		log.Printf("recv GateIO: %s", message)
	}
	return nil
}

func makeCandle(candleData []candlestickMessage) {

}
