package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	ws "github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

const (
	tickerDurationDefault     = "60"
	outlierBasisPointsDefault = "30"
	bufferSize                = 10000
)

var (
	tickerDurationSeconds int64
	outlierBasisPoints    float64
)

type candlestickMessage struct {
	ForeignName string
	// ClosingPrice in USDT
	ClosingPrice float64
	Volume       float64
	Timestamp    time.Time
	ScrapeTime   time.Time
	Source       string
}

func getCandleStickMessageIdent(message candlestickMessage) string {
	return message.Source + "-" + message.ForeignName
}

func main() {

	wg := sync.WaitGroup{}
	ds, err := models.NewDataStore()
	if err != nil {
		log.Fatal("datastore error: ", err)
	}

	tickerDurationString := utils.Getenv("TICKER_DURATION", tickerDurationDefault)
	tickerDurationSeconds, err = strconv.ParseInt(tickerDurationString, 10, 64)
	if err != nil {
		log.Error("parse ticker duration from env variable: ", err)
	}
	ticker := time.NewTicker(time.Duration(tickerDurationSeconds) * time.Second)

	outlierBasisPointsString := utils.Getenv("BASIS_POINTS", outlierBasisPointsDefault)
	outlierBasisPoints, err = strconv.ParseFloat(outlierBasisPointsString, 64)
	if err != nil {
		log.Error("parse basis points from env variable: ", err)
	}

	assets := flag.String("assets", "BTC,ETH", "asset symbols to query (from BTC, ETH, SOL, GLMR, DOT")
	exchanges := flag.String("exchanges", "Binance,GateIO,Kucoin,Huobi", "exchanges to query (from Binance, Kucoin, Coinbase, Huobi, Okex, GateIO")
	flag.Parse()
	cChan := make(chan candlestickMessage, bufferSize)

	for _, exchange := range strings.Split(*exchanges, ",") {
		wg.Add(1)
		go handleExchangeScraper(exchange, *assets, cChan, &wg)
	}
	defer wg.Wait()
	for t := range ticker.C {
		log.Info("channel before takeout: ", len(cChan))
		channelData := getRecentDataFromChannel(cChan, t)
		log.Info("channel after takeout: ", len(cChan))
		pairData := getPairData(channelData)
		// log.Info("pair data: ", pairData)
		for key, values := range pairData {
			var currExchanges string
			for _, value := range values {
				currExchanges += value.Source + " , "
			}
			log.Infof("%s market on exchanges: %s", key, currExchanges)
		}
		vwap, err := makeVWAP(pairData, outlierBasisPoints)
		if err != nil {
			log.Error("makeVWAP: ", err)
		}
		log.Info("vwap: ", vwap)
		for key, value := range vwap {
			err = ds.SetVWAPFirefly(key, value, t)
			if err != nil {
				log.Error("set VWAP: ", err)
			}
		}
	}
}

func handleExchangeScraper(exchange string, assets string, candleChan chan candlestickMessage, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Printf("Entered Exchange handler for %s", exchange)
	for {
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
		case "OKEx":
			log.Println("OKEx Scraper: Start scraping")
			err := scrapeOkex(assets, candleChan)
			if err != nil {
				log.Error("OKEx scraper: ", err)
			}
		case "HitBTC":
			log.Println("HitBTC Scraper: Start scraping")
			err := scrapeHitbtc(assets, candleChan)
			if err != nil {
				log.Error("HitBTC scraper: ", err)
			}
		case "Kraken":
			log.Println("Kraken Scraper: Start scraping")
			err := scrapeKraken(assets, candleChan)
			if err != nil {
				log.Error("Kraken scraper: ", err)
			}
		default:
			log.Errorf("Unknown scraper name %s", exchange)
			return
		}
		log.Info("Sleeping 30sec for exchange ", exchange)
		time.Sleep(30 * time.Second)
	}
}

func scrapeKraken(assets string, candleChan chan candlestickMessage) error {
	log.Info("Entered Kraken handler")
	wsBaseString := "wss://ws.kraken.com"

	conn, _, err := ws.DefaultDialer.Dial(wsBaseString, nil)
	if err != nil {
		return err
	}
	defer conn.Close()
	for _, asset := range strings.Split(assets, ",") {
		assetName := asset
		if strings.ToUpper(asset) == "BTC" {
			assetName = "XBT"
		}
		baseAsset := "USDT"
		if strings.ToUpper(asset) == "GLMR" {
			baseAsset = "USD"
		}
		msgToWrite := fmt.Sprintf("{\"event\":\"subscribe\",\"pair\":[\"%s/%s\"],\"subscription\":{\"interval\":1,\"name\":\"ohlc\"}}", strings.ToUpper(assetName), strings.ToUpper(baseAsset))
		conn.WriteMessage(ws.TextMessage, []byte(msgToWrite))
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Errorln("read:", err)
			return err
		}
		// log.Printf("recv Kraken: %s", message)
		messageMap := make(map[string]interface{})
		messageList := make([]interface{}, 3)
		err = json.Unmarshal(message, &messageMap)
		if err != nil {
			// We have a real result here
			err = json.Unmarshal(message, &messageList)
			if err != nil {
				return err
			}
		}

		// Check if we got a status msg
		if messageMap["event"] == "systemStatus" {
			continue
		}
		if messageMap["event"] == "heartbeat" {
			continue
		}
		if messageMap["event"] == "subscriptionStatus" {
			continue
		}

		content := messageList[1].([]interface{})

		closingPriceString := content[5].(string)
		closingPrice, err := strconv.ParseFloat(closingPriceString, 64)
		if err != nil {
			return err
		}
		volumeString := content[7].(string)
		volume, err := strconv.ParseFloat(volumeString, 64)
		if err != nil {
			return err
		}

		timeUnixString := content[1].(string)
		timeUnix, err := strconv.ParseFloat(timeUnixString, 64)
		if err != nil {
			return err
		}

		pairString := messageList[3].(string)
		assetString := strings.Split(pairString, "/")[0]
		if assetString == "XBT" {
			assetString = "BTC"
		}

		candleStickMessage := candlestickMessage{
			ForeignName:  assetString + strings.Split(pairString, "/")[1],
			ClosingPrice: closingPrice,
			Volume:       volume,
			Timestamp:    time.Unix(int64(timeUnix), 0),
			ScrapeTime:   time.Now(),
			Source:       "Kraken",
		}

		candleChan <- candleStickMessage
	}
}

func scrapeHitbtc(assets string, candleChan chan candlestickMessage) error {
	log.Info("Entered HitBTC handler")
	wsBaseString := "wss://api.hitbtc.com/api/3/ws/public"

	conn, _, err := ws.DefaultDialer.Dial(wsBaseString, nil)
	if err != nil {
		return err
	}
	defer conn.Close()
	for _, asset := range strings.Split(assets, ",") {
		msgToWrite := fmt.Sprintf("{\"method\":\"subscribe\",\"ch\":\"candles/M1\",\"params\":{\"symbols\":[\"%sUSDT\"],\"limit\":10},\"id\": 1}", strings.ToUpper(asset))
		conn.WriteMessage(ws.TextMessage, []byte(msgToWrite))
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Errorln("read:", err)
			return err
		}
		//log.Printf("recv HitBTC: %s", message)
		messageMap := make(map[string]interface{})
		err = json.Unmarshal(message, &messageMap)
		if err != nil {
			return err
		}

		// Check if we got a status msg
		if messageMap["result"] != nil {
			continue
		}
		// or the initial snapshot
		if messageMap["snapshot"] != nil {
			continue
		}

		update := messageMap["update"].(map[string]interface{})
		assetKey := ""
		for k := range update {
			assetKey = k
			break
		}

		content := update[assetKey].([]interface{})
		result := content[0].(map[string]interface{})

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
		timeUnix := result["t"].(float64)

		foreignNameString := assetKey

		candleStickMessage := candlestickMessage{
			ForeignName:  strings.ToUpper(foreignNameString),
			ClosingPrice: closingPrice,
			Volume:       volume,
			Timestamp:    time.Unix(int64(timeUnix)/1000, 0),
			ScrapeTime:   time.Now(),
			Source:       "HitBTC",
		}

		candleChan <- candleStickMessage
	}
}

func scrapeOkex(assets string, candleChan chan candlestickMessage) error {
	log.Info("Entered OKExhandler")
	wsBaseString := "wss://ws.okx.com:8443/ws/v5/public"

	conn, _, err := ws.DefaultDialer.Dial(wsBaseString, nil)
	if err != nil {
		return err
	}
	defer conn.Close()
	for _, asset := range strings.Split(assets, ",") {
		//{"op": "subscribe","args": [{"channel": "candle1m","instId": "GLMR-USDT"}]}
		msgToWrite := fmt.Sprintf("{\"op\":\"subscribe\",\"args\":[{\"channel\": \"candle1m\",\"instId\":\"%s-USDT\"}]}", strings.ToUpper(asset))
		conn.WriteMessage(ws.TextMessage, []byte(msgToWrite))
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			return err
		}

		//log.Printf("recv OKEx: %s", message)
		messageMap := make(map[string]interface{})
		err = json.Unmarshal(message, &messageMap)
		if err != nil {
			return err
		}

		// Check if we got a status msg
		if messageMap["event"] == "subscribe" {
			continue
		}
		data := messageMap["data"].([]interface{})
		result := data[0].([]interface{})

		arg := messageMap["arg"].(map[string]interface{})

		closingPriceString := result[4].(string)
		closingPrice, err := strconv.ParseFloat(closingPriceString, 64)
		if err != nil {
			return err
		}
		volumeString := result[5].(string)
		volume, err := strconv.ParseFloat(volumeString, 64)
		if err != nil {
			return err
		}

		timeUnixString := result[0].(string)
		timeUnix, err := strconv.ParseFloat(timeUnixString, 64)
		if err != nil {
			return err
		}

		foreignNameString := arg["instId"].(string)
		foreignNameFiltered := strings.Split(foreignNameString, "-")[0]

		candleStickMessage := candlestickMessage{
			ForeignName:  foreignNameFiltered + "USDT",
			ClosingPrice: closingPrice,
			Volume:       volume,
			Timestamp:    time.Unix(int64(timeUnix/1000), 0).Add(1 * time.Minute),
			ScrapeTime:   time.Now(),
			Source:       "OKEx",
		}

		candleChan <- candleStickMessage
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
		//log.Printf("recv Binance: %s", message)
		messageMap := make(map[string]interface{})
		err = json.Unmarshal(message, &messageMap)
		data := messageMap["data"].(map[string]interface{})["k"].(map[string]interface{})
		timeUnix := messageMap["data"].(map[string]interface{})["E"].(float64)

		closingPriceString := data["c"].(string)
		closingPrice, err := strconv.ParseFloat(closingPriceString, 64)
		volumeString := data["V"].(string)
		volume, err := strconv.ParseFloat(volumeString, 64)

		candleStickMessage := candlestickMessage{
			ForeignName:  data["s"].(string),
			ClosingPrice: closingPrice,
			Volume:       volume,
			Timestamp:    time.Unix(int64(timeUnix/1000), 0),
			ScrapeTime:   time.Now(),
			Source:       "Binance",
		}

		candleChan <- candleStickMessage

	}
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
		msgToWrite := fmt.Sprintf("{\"time\":5,\"channel\":\"spot.candlesticks\",\"event\":\"subscribe\",\"payload\":[\"1m\",\"%s_USDT\"]}", asset)
		conn.WriteMessage(ws.TextMessage, []byte(msgToWrite))
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Errorln("read:", err)
			return err
		}
		// log.Printf("recv GateIO: %s", message)
		messageMap := make(map[string]interface{})
		err = json.Unmarshal(message, &messageMap)
		if err != nil {
			return err
		}
		if messageMap["error"] != nil {
			log.Errorf("error GateIO: %s", message)
			continue
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
			ScrapeTime:   time.Now(),
			Source:       "GateIO",
		}

		candleChan <- candleStickMessage
	}
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
	body, err := ioutil.ReadAll(resp.Body)
	bodyMap := make(map[string]interface{})
	err = json.Unmarshal(body, &bodyMap)
	if err != nil {
		return err
	}
	token := bodyMap["data"].(map[string]interface{})["token"].(string)

	conn, _, err := ws.DefaultDialer.Dial(wsBaseString+token, nil)
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
			ScrapeTime:   time.Now(),
			Source:       "Kucoin",
		}

		// send a ping for every msg TODO: set to timer
		msgToWrite := fmt.Sprintf("{\"id\":%d,\"type\":\"ping\"}", 1)
		conn.WriteMessage(ws.TextMessage, []byte(msgToWrite))

		candleChan <- candleStickMessage

	}
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
		gzreader, err := gzip.NewReader(bytesReader)
		if err != nil {
			return err
		}

		message, err := ioutil.ReadAll(gzreader)
		if err != nil {
			log.Errorln("read:", err)
			return err
		}
		//log.Printf("recv Huobi: %s", message)

		messageMap := make(map[string]interface{})
		err = json.Unmarshal(message, &messageMap)
		if err != nil {
			return err
		}

		// Check if we got a ping
		if messageMap["ping"] != nil {
			pingNumber := int(messageMap["ping"].(float64))
			msgToWrite := fmt.Sprintf("{\"pong\":%d}", pingNumber)
			conn.WriteMessage(ws.TextMessage, []byte(msgToWrite))
			continue
		}

		if messageMap["tick"] == nil {
			continue
		}

		result := messageMap["tick"].(map[string]interface{})

		closingPrice := result["close"].(float64)
		volume := result["amount"].(float64)
		timeUnix := messageMap["ts"].(float64)
		timeUnix /= 1e3

		foreignNameString := messageMap["ch"].(string)
		foreignNameFiltered := strings.Split(foreignNameString, ".")[1]

		candleStickMessage := candlestickMessage{
			ForeignName:  strings.ToUpper(foreignNameFiltered),
			ClosingPrice: closingPrice,
			Volume:       volume,
			Timestamp:    time.Unix(int64(timeUnix), 0),
			ScrapeTime:   time.Now(),
			Source:       "Huobi",
		}

		candleChan <- candleStickMessage
	}
}

// getRecentDataFromChannel returns the most recent data for each identifier
// Source-ForeignName sitting in the channel @candleChan. It reads data from the channel
// until the first data point hits @endtime.
func getRecentDataFromChannel(candleChan chan candlestickMessage, endtime time.Time) map[string]candlestickMessage {
	lastCandleData := make(map[string]candlestickMessage)
	for message := range candleChan {
		// log.Info("message: ", message)
		// Channels are passed by reference. As the channel is continuously written to,
		// we need to stop fetching from it as soon as endtime is passed.
		if message.ScrapeTime.After(endtime) {
			return lastCandleData
		}
		if endtime.Sub(message.Timestamp) > time.Duration(tickerDurationSeconds)*time.Second || message.Timestamp.Sub(endtime) > 0 {
			continue
		}
		messageIdent := getCandleStickMessageIdent(message)
		if _, ok := lastCandleData[messageIdent]; !ok {
			lastCandleData[messageIdent] = message
		} else {
			if message.Timestamp.After(lastCandleData[messageIdent].Timestamp) {
				lastCandleData[messageIdent] = message
			}
		}
	}
	return lastCandleData
}

// makeCandle returns the VWAP of a trading pair in USD.
// It discards everything above 30 basis points from the median.
func getPairData(candleData map[string]candlestickMessage) map[string][]candlestickMessage {
	pairData := make(map[string][]candlestickMessage)
	for key, value := range candleData {
		pair := strings.Split(key, "-")[1]
		if _, ok := pairData[pair]; !ok {
			pairData[pair] = []candlestickMessage{value}
		} else {
			pairData[pair] = append(pairData[pair], value)
		}
	}
	return pairData
}

// makeVWAP takes a map with foreign names as keys and []candlestickMessage as values,
// containing all values of the underlying pair across sources, i.e. (at most) one value per source.
func makeVWAP(pairData map[string][]candlestickMessage, basispoints float64) (map[string]float64, error) {
	vwapMap := make(map[string]float64)
	// key is a market, such as BTCUSDT, and value is a slice of data on this market from various exchanges.
	for key, value := range pairData {
		cleanedPrices, cleanedVolumes, err := discardOutliers(getPrices(value), getVolumes(value), basispoints, value)
		if err != nil {
			return vwapMap, err
		}
		vwapMap[key], err = vwap(cleanedPrices, cleanedVolumes)
		if err != nil {
			return vwapMap, err
		}
	}
	return vwapMap, nil
}

// vwap returns the volume weighted average price for the slices @prices and @volumes.
func vwap(prices []float64, volumes []float64) (float64, error) {
	//log.Info("prices, volumes: ", prices, volumes)
	if len(prices) != len(volumes) {
		return 0, errors.New("number of prices does not equal number of volumes ")
	}
	avg := float64(0)
	totalVolume := float64(0)
	for i := 0; i < len(prices); i++ {
		avg += prices[i] * math.Abs(volumes[i])
		totalVolume += math.Abs(volumes[i])
	}
	if totalVolume > 0 {
		return avg / totalVolume, nil
	} else {
		return 0, nil
	}
}

// discardOutliers discards every data point from @prices and @volumes that deviates from
// the price median by more than @basispoints basis points.
func discardOutliers(prices []float64, volumes []float64, basispoints float64, values []candlestickMessage) (newPrices []float64, newVolumes []float64, err error) {
	if len(prices) != len(volumes) {
		err = errors.New("number of prices does not equal number of volumes ")
		return
	}
	median := computeMedian(prices)
	threshold := basispoints * float64(0.0001) * median
	for i := 0; i < len(prices); i++ {
		if math.Abs(prices[i]-median) < threshold {
			newPrices = append(newPrices, prices[i])
			newVolumes = append(newVolumes, volumes[i])
		} else {
			log.Warnf("discard %s on %s", values[i].ForeignName, values[i].Source)
		}
	}
	return
}

// computeMedian returns the median of @samples.
func computeMedian(samples []float64) (median float64) {
	var length = len(samples)
	if length > 0 {
		sort.Float64s(samples)
		if length%2 == 0 {
			median = (samples[length/2-1] + samples[length/2]) / 2
		} else {
			median = samples[(length+1)/2-1]
		}
	}
	return
}

func getPrices(messages []candlestickMessage) (prices []float64) {
	for _, message := range messages {
		prices = append(prices, message.ClosingPrice)
	}
	return
}

func getVolumes(messages []candlestickMessage) (volumes []float64) {
	for _, message := range messages {
		volumes = append(volumes, message.Volume)
	}
	return
}
