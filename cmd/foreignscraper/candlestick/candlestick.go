package main

import (
	"flag"
	"fmt"
	"strings"
	"sync"

	_ "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
	ws "github.com/gorilla/websocket"
)

func main() {

	wg := sync.WaitGroup{}
	/*ds, err := models.NewDataStore()
	if err != nil {
		log.Fatal("datastore error: ", err)
	}*/

	assets := flag.String("assets", "BTC,ETH", "asset symbols to query (from BTC, ETH, SOL, GLMR, DOT")
	exchanges := flag.String("exchanges", "Binance,Kucoin", "exchanges to query (from Binance, Kucoin, Coinbase, Huobi, Okex, Gateio")
	flag.Parse()

	for _, exchange := range strings.Split(*exchanges, ",") {
		wg.Add(1)
		go handleExchangeScraper(exchange, *assets, &wg)
	}

	defer wg.Wait()
}

func handleExchangeScraper(exchange string, assets string, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("Entered Exchange handler for %s", exchange)
	switch exchange {
	case "Binance":
		log.Println("Binance Scraper: Start scraping")
		err := scrapeBinance(assets)
		if err != nil {
			log.Error(err)
		}
	case "Gateio":
		log.Println("Gateio Scraper: Start scraping")
		err := scrapeGateio(assets)
		if err != nil {
			log.Error(err)
		}
	case "Kucoin":
		log.Println("Kucoin Scraper: Start scraping")
	default:
		log.Errorf("Unknown scraper name %s", exchange)
	}

}

func scrapeBinance(assets string) error {
	log.Info("Entered Binance handler")
	wsBaseString := "wss://fstream.binance.com/stream?streams="
	wsAssetsString := ""

	for _, asset := range strings.Split(assets, ",") {
		wsAssetsString += strings.ToLower(asset) + "usdt@kline_1m" + "/"
	}
	conn, _, err := ws.DefaultDialer.Dial(wsBaseString + wsAssetsString, nil)
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
		log.Printf("recv: %s", message)
	}
	return nil
}

func scrapeGateio(assets string) error {
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
		log.Printf("recv: %s", message)
	}
	return nil
}
