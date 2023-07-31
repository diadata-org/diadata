package dscrapers

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	ws "github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

const (
	WS_ENDPOINT_BASE = "wss://stream.binance.com:9443/stream?streams="
)

func init() {
	log.Info("TO DO")
}

type binanceResponse struct {
	StreamType string       `json:"stream"`
	Data       binanceTrade `json:"data"`
}

type binanceTrade struct {
	Event          string `json:"e"`
	ForeignName    string `json:"s"`
	ForeignTradeID int    `json:"t"`
	Price          string `json:"p"`
	Volume         string `json:"q"`
	TimeTrade      int    `json:"T"`
	MarketMaker    bool   `json:"m"`
	TimeEvent      int    `json:"E"`
}

type BinanceScraper struct {
	tradeChannel chan *dia.DTrade
	pairs        []string
}

func NewBinanceScraper() BinanceScraper {
	scraper := BinanceScraper{
		tradeChannel: make(chan *dia.DTrade),
	}
	go scraper.mainLoop()
	return scraper
}

func (scraper *BinanceScraper) mainLoop() {
	// TO DO: Fetch trades from ws API and send into tradeChannel.
	var pairsString string
	id := ""
	scraper.ReadPairs(id)
	for i, pair := range scraper.pairs {
		if i == 0 {
			pairsString += pair
		} else {
			pairsString += "/" + pair
		}
	}
	conn, _, err := ws.DefaultDialer.Dial(WS_ENDPOINT_BASE+pairsString, nil)
	if err != nil {
		log.Error("connect to websocket server: ", err)
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Error("read message:", err)
		}

		var response binanceResponse
		err = json.Unmarshal(message, &response)
		if err != nil {
			log.Error("unmarshal binance trade: ", err)
		}

		dTrade, err := parseBinanceResponse(response.Data)
		if err != nil {
			log.Error("parse binance response: ", err)
		} else {
			scraper.TradeChannel() <- &dTrade
		}
	}
}

func (scraper *BinanceScraper) ReadPairs(id string) {
	// TO DO: Read pairs from IPFS using id (as CID?).
	scraper.pairs = []string{"btcusdt@trade"}
}

func (scraper *BinanceScraper) TradeChannel() chan *dia.DTrade {
	return scraper.tradeChannel
}

func (scraper *BinanceScraper) FetchClock() (time.Time, uint32, error) {
	// TO DO: Fetch clock from central control unit.
	return time.Now(), 120, nil
}

func parseBinanceResponse(bt binanceTrade) (trade dia.DTrade, err error) {
	trade.ForeignTradeID = strconv.Itoa(bt.ForeignTradeID)
	trade.Price, err = strconv.ParseFloat(bt.Price, 64)
	if err != nil {
		return
	}
	trade.Volume, err = strconv.ParseFloat(bt.Volume, 64)
	if err != nil {
		return
	}
	trade.Time = time.Unix(0, int64(bt.TimeTrade)*1e6)
	return
}
