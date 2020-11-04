package defiscrapers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type BitfinexTickerIndex uint

// Bitfinex response message indices
const (
	FRR BitfinexTickerIndex = iota
	BID
	BID_PERIOD
	BID_SIZE
	ASK
	ASK_PERIOD
	ASK_SIZE
	DAILY_CHANGE
	DAILY_CHANGE_RELATIVE
	LAST_PRICE
	VOLUME
	HIGH
	LOW
	_PLACE_HOLDER1
	_PLACE_HOLDER2
	FRR_AMOUNT_AVAILABLE
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 1024
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type BitfinexProtocol struct {
	scraper        *DefiScraper
	protocol       dia.DefiProtocol
	WsConnections  []*BitfinexWSSConnection
	fundingSymbols []string
}

type BitfinexWSSConnection struct {
	proto   *BitfinexProtocol
	restart chan bool
	// The websocket connection.
	conn   *websocket.Conn
	symbol string
}

func NewBitfinex(scraper *DefiScraper, protocol dia.DefiProtocol) *BitfinexProtocol {
	proto := &BitfinexProtocol{scraper: scraper, protocol: protocol}
	// Manually fetch funding symbols
	fundingSymbols, err := fetchFundingSymbols()
	if err != nil {
		panic(err)
	}
	proto.fundingSymbols = fundingSymbols
	// Add a websocket connection for every funding symbol
	proto.WsConnections = make([]*BitfinexWSSConnection, 0)
	for _, sym := range proto.fundingSymbols {
		proto.WsConnections = append(proto.WsConnections,
			&BitfinexWSSConnection{
				proto:   proto,
				symbol:  sym,
				restart: make(chan bool),
			})
	}
	go func() {
		for {
			for _, wsConn := range proto.WsConnections {
				select {
				// Reconnect websocket connection on restart
				case <-wsConn.restart:
					// Try to close the connection first
					if wsConn.conn != nil {
						wsConn.conn.Close()
					}
					go wsConn.connectWebsocket()
				}
			}
		}
	}()
	// Start all connections
	for _, wsConn := range proto.WsConnections {
		wsConn.restart <- true
	}
	return proto
}

// Fetch all funding symbols from Bitfinex api in bulk
func fetchFundingSymbols() (fundingSymbols []string, err error) {
	fundingSymbols = make([]string, 0)
	jsondata, err := utils.GetRequest("https://api-pub.bitfinex.com/v2/tickers?symbols=ALL")
	log.Debugf("%s", jsondata)
	resp := make([][]interface{}, 0)
	err = json.Unmarshal(jsondata, &resp)
	if err != nil {
		log.Debug("While getting all symbols")
		return nil, err
	}
	// Filter funding symbols
	for _, sym := range resp {
		symName, ok := sym[0].(string)
		if ok && strings.HasPrefix(symName, "f") {
			fundingSymbols = append(fundingSymbols, symName[1:])
		}
	}
	return fundingSymbols, nil
}

func (ws *BitfinexWSSConnection) pingWebsocket() {
	w, err := ws.conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return
	}
	w.Write([]byte(`{ "event": "subscribe", "channel": "ticker", "symbol": "f` + ws.symbol + `" }`))
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		ws.restart <- true
	}()
	for {
		select {
		case <-ticker.C:
			ws.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (ws *BitfinexWSSConnection) fetchingForever() {
	defer func() {
		ws.restart <- true

	}()
	ws.conn.SetReadLimit(maxMessageSize)
	ws.conn.SetReadDeadline(time.Now().Add(pongWait))
	ws.conn.SetPongHandler(func(string) error { ws.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := ws.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
				ws.restart <- true
				return
			}
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		resp := make([]interface{}, 0)
		if err := json.Unmarshal(message, &resp); err != nil {
			continue
		}
		event, ok := resp[1].([]interface{})
		if !ok || len(event) == 0 {
			// Skipping on empty message
			continue
		}
		log.Printf("Updating DEFI Rate for %+v\n ", ws.proto.protocol.Name)
		borrowingRate, ok := event[FRR].(float64)
		if !ok {
			log.Errorf("BitfinexProtocol: While parsing borrowing rate: %v", err)
			continue
		}
		// Estimate annual percentage yield from daily percent
		borrowingRate = 100 * 365 * borrowingRate
		// LenderRate affects by Bitfinex fees
		// See https://www.bitfinex.com/fees/#margin-funding
		// Also compare to values given on https://www.bitfinex.com/funding
		lendingRate := borrowingRate * 0.85

		asset := &dia.DefiRate{
			Timestamp:     time.Now(),
			Asset:         ws.symbol,
			Protocol:      ws.proto.protocol.Name,
			LendingRate:   lendingRate,
			BorrowingRate: borrowingRate,
		}
		log.Printf("writing DEFI rate for  %#v in %v\n", asset, ws.proto.scraper.RateChannel())
		ws.proto.scraper.RateChannel() <- asset
		log.Info("Update complete")
	}
}
func (ws *BitfinexWSSConnection) connectWebsocket() {
	var err error
	u := url.URL{Scheme: "wss", Host: "api-pub.bitfinex.com", Path: "/ws/2"}
	ws.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	log.Printf("connecting to %s", u.String())
	if err != nil {
		log.Errorf("dial: %v", err)
		ws.restart <- true
		return
	}
	go ws.pingWebsocket()
	go ws.fetchingForever()

}
func (proto *BitfinexProtocol) UpdateRate() error {
	//NO OP
	return nil
}

func fetchTotalLocked(symbols []string) (amount float64, err error) {
	// Fetch market prices in bulk, so we don't be affected by Bitfinex api rate limits
	// See https://docs.bitfinex.com/reference#rest-public-tickers
	jsondata, err := utils.GetRequest("https://api-pub.bitfinex.com/v2/tickers?symbols=ALL")
	log.Debugf("%s", jsondata)
	priceList := make([][]interface{}, 0)
	err = json.Unmarshal(jsondata, &priceList)
	if err != nil {
		return 0, err
	}
	for _, symbol := range symbols {
		// We need to make an api call for every funding symbol
		jsondata, err := utils.GetRequest(fmt.Sprintf("https://api-pub.bitfinex.com/v2/stats1/funding.size:1m:f%s/last", symbol))
		resp := make([]float64, 0)
		err = json.Unmarshal(jsondata, &resp)
		if err != nil {
			return 0, err
		}
		if len(resp) != 2 {
			return 0, fmt.Errorf("Unexpected response from bitfinex stats api")
		}
		symTotal := resp[1]
		// Find price for funding symbols
		for _, priceInfo := range priceList {
			symName, ok := priceInfo[0].(string)
			if ok && symName == fmt.Sprintf("t%sUSD", symbol) {
				// Check if we found a valid funding symbol
				price, ok := priceInfo[1].(float64)
				if ok {
					amount += symTotal * price
				}
			}
		}
	}
	return
}
func (proto *BitfinexProtocol) UpdateState() error {
	log.Print("Updating DEFI state for %+v\n ", proto.protocol)
	totalSupplyUSD, err := fetchTotalLocked(proto.fundingSymbols)
	if err != nil {
		log.Errorf("BitfinexProtocol: While parsing totalSupplyUSD: %v", err)
	}
	priceETH, err := utils.GetCoinPrice("ETH")
	if err != nil {
		log.Errorf("BitfinexProtocol: While parsing priceETH: %v", err)
	}
	totalSupplyETH := totalSupplyUSD / priceETH
	defistate := &dia.DefiProtocolState{
		TotalUSD:  totalSupplyUSD,
		TotalETH:  totalSupplyETH,
		Protocol:  proto.protocol,
		Timestamp: time.Now(),
	}
	proto.scraper.StateChannel() <- defistate
	log.Printf("writing DEFI state for  %#v in %v\n", defistate, proto.scraper.StateChannel())

	log.Info("Update State complete")
	return nil
}
