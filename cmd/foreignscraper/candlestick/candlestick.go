package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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
			log.Error("Binance scraper: ", err)
		}
	case "GateIO":
		log.Println("Gateio Scraper: Start scraping")
		err := scrapeGateio(assets, candleChan)
		if err != nil {
			log.Error("GateIO scraper: ", err)
		}
	case "Kucoin":
		log.Println("Kucoin Scraper: Start scraping")
		err := scrapeKucoin(assets, candleChan)
		if err != nil {
			log.Error("Kucoin scraper: ", err)
		}
	case "Huobi":
		log.Println("Huobi Scraper: Start scraping")
		err := scrapeHuobi(assets, candleChan)
		if err != nil {
			log.Error("Huobi scraper: ", err)
		}
	default:
		log.Errorf("Unknown scraper name %s", exchange)
	}
}

func scrapeBinance(assets string, candleChan chan candlestickMessage) error {
	log.Info("Entered Binance handler")
	wsBaseString := "wss://stream.binance.com:9443/stream?streams="
	wsAssetsString := ""

	for _, asset := range strings.Split(assets, ",") {
		wsAssetsString += strings.ToLower(asset) + "usdt@kline_1m" + "/"
	}
	// Remove trailing slash
	wsAssetsString = wsAssetsString[:len(wsAssetsString)-1]
	conn, _, err := ws.DefaultDialer.Dial(wsBaseString+wsAssetsString, nil)
	if err != nil {
		return err
	}
	defer conn.Close()

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
	wsBaseString := "wss://api.gateio.ws/ws/v4/"

	conn, _, err := ws.DefaultDialer.Dial(wsBaseString, nil)
	if err != nil {
		return err
	}
	defer conn.Close()
	for _, asset := range strings.Split(assets, ",") {
		msgToWrite := fmt.Sprintf("{\"time\":30,\"channel\":\"spot.candlesticks\",\"event\":\"subscribe\",\"payload\":[\"1m\",\"%s_USD\"]}", asset)
		conn.WriteMessage(ws.TextMessage, []byte(msgToWrite))
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Errorln("read:", err)
			return err
		}
		//log.Printf("recv GateIO: %s", message)
		messageMap := make(map[string]interface{})
		err = json.Unmarshal(message, &messageMap)
		if err != nil {
			return err
		}
		result := messageMap["result"].(map[string]interface{})

		// Check if we got a status msg
		if result["status"] != nil {
			continue
		}

		closingPriceString := result["c"].(string)
		closingPrice, err := strconv.ParseFloat(closingPriceString, 64)
		if err != nil {
			return err
		}
		volumeString := result["v"].(string)
		volume, err := strconv.ParseFloat(volumeString, 64)
		if err != nil {
			return err
		}
		timeUnixString := result["t"].(string)
		timeUnix, err := strconv.ParseFloat(timeUnixString, 64)
		if err != nil {
			return err
		}

		foreignNameString := result["n"].(string)
		foreignNameFiltered := strings.Split(foreignNameString, "_")[1]

		candleStickMessage := candlestickMessage{
			ForeignName:  strings.ToUpper(foreignNameFiltered) + "USDT",
			ClosingPrice: closingPrice,
			Volume:       volume,
			Timestamp:    time.Unix(int64(timeUnix), 0),
			Source:       "GateIO",
		}

		candleChan <- candleStickMessage
	}
	return nil
}

func scrapeKucoin(assets string, candleChan chan candlestickMessage) error {
	log.Info("Entered Kucoin handler")
	wsBaseString := "wss://ws-api.kucoin.com/endpoint?token="

	// Get token
	tokenUrl := "https://api.kucoin.com/api/v1/bullet-public"
	resp, err := http.Post(tokenUrl, "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	bodyMap := make(map[string]interface{})
	err = json.Unmarshal(body, &bodyMap)
	if err != nil {
		return err
	}
	token := bodyMap["data"].(map[string]interface{})["token"].(string)

	conn, _, err := ws.DefaultDialer.Dial(wsBaseString + token, nil)
	if err != nil {
		return err
	}
	defer conn.Close()
	for _, asset := range strings.Split(assets, ",") {
		msgToWrite := fmt.Sprintf("{\"id\":1,\"type\":\"subscribe\",\"topic\": \"/market/candles:%s-USDT_1min\",\"response\": true}", strings.ToUpper(asset))
		conn.WriteMessage(ws.TextMessage, []byte(msgToWrite))
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Errorln("read:", err)
			return err
		}
		//log.Printf("recv Kucoin: %s", message)
		messageMap := make(map[string]interface{})
		err = json.Unmarshal(message, &messageMap)
		if err != nil {
			return err
		}
		// Check if we got a status msg
		if messageMap["type"] != "message" {
			continue
		}
		result := messageMap["data"].(map[string]interface{})
		candles := result["candles"].([]interface{})

		closingPriceString := candles[2].(string)
		closingPrice, err := strconv.ParseFloat(closingPriceString, 64)
		if err != nil {
			return err
		}
		volumeString := candles[5].(string)
		volume, err := strconv.ParseFloat(volumeString, 64)
		if err != nil {
			return err
		}
		timeUnix := result["time"].(float64)
		timeUnix /= 1e9

		foreignNameString := result["symbol"].(string)
		foreignNameFiltered := strings.Split(foreignNameString, "-")[0]

		candleStickMessage := candlestickMessage{
			ForeignName:  foreignNameFiltered + "USDT",
			ClosingPrice: closingPrice,
			Volume:       volume,
			Timestamp:    time.Unix(int64(timeUnix), 0),
			Source:       "Kucoin",
		}

		candleChan <- candleStickMessage
	}
	return nil
}

func scrapeHuobi(assets string, candleChan chan candlestickMessage) error {
	log.Info("Entered Huobi handler")
	wsBaseString := "wss://api.huobi.pro/ws"
	
	conn, _, err := ws.DefaultDialer.Dial(wsBaseString, nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	for _, asset := range strings.Split(assets, ",") {
		msgToWrite := fmt.Sprintf("{\"sub\":\"market.%susdt.kline.1min\",\"id\":\"id1\"}", strings.ToLower(asset))
		conn.WriteMessage(ws.TextMessage, []byte(msgToWrite))
	}

	for {
		_, zippedMessage, err := conn.ReadMessage()
		bytesReader := bytes.NewReader([]byte(zippedMessage))
		gzreader, err := gzip.NewReader(bytesReader);
		if err != nil {
    	return err
		}

		message, err := ioutil.ReadAll(gzreader);
		if err != nil {
    	return err
		}

		if err != nil {
			log.Errorln("read:", err)
			return err
		}
		log.Printf("recv Huobi: %s", message)
		messageMap := make(map[string]interface{})
		err = json.Unmarshal(message, &messageMap)
		if err != nil {
			return err
		}
		// Check if we got a status msg
		if messageMap["type"] != "message" {
			continue
		}
		result := messageMap["data"].(map[string]interface{})
		candles := result["candles"].([]interface{})

		closingPriceString := candles[2].(string)
		closingPrice, err := strconv.ParseFloat(closingPriceString, 64)
		if err != nil {
			return err
		}
		volumeString := candles[5].(string)
		volume, err := strconv.ParseFloat(volumeString, 64)
		if err != nil {
			return err
		}
		timeUnix := result["time"].(float64)
		timeUnix /= 1e9

		foreignNameString := result["symbol"].(string)
		foreignNameFiltered := strings.Split(foreignNameString, "-")[0]

		candleStickMessage := candlestickMessage{
			ForeignName:  foreignNameFiltered + "USDT",
			ClosingPrice: closingPrice,
			Volume:       volume,
			Timestamp:    time.Unix(int64(timeUnix), 0),
			Source:       "Kucoin",
		}

		candleChan <- candleStickMessage
	}
	return nil
}

func makeCandle(candleData []candlestickMessage) {

}
