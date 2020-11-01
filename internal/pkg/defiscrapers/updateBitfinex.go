package defiscrapers

import (
	"bytes"
	"encoding/json"
	"math/big"
	"net/url"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type BitfinexTickerIndex uint

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
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type BitfinexProtocol struct {
	scraper  *DefiScraper
	protocol dia.DefiProtocol
	restart  chan bool
	// The websocket connection.
	conn *websocket.Conn
}

type BitfinexEvent struct {
	Data []interface{}          `json:"-"`
	X    map[string]interface{} `json:"-"` // Rest of the fields should go here.
}

func NewBitfinex(scraper *DefiScraper, protocol dia.DefiProtocol) *BitfinexProtocol {
	log.Info("Hello mothefucker")
	proto := &BitfinexProtocol{scraper: scraper, protocol: protocol}
	proto.restart = make(chan bool)
	go func() {
		for {
			select {
			case <-proto.restart:
				go proto.connectWebsocket()
			}
		}
	}()
	proto.restart <- true
	return proto
}

func (proto *BitfinexProtocol) pingWebsocket() {
	w, err := proto.conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return
	}
	w.Write([]byte(`{ "event": "subscribe", "channel": "ticker", "symbol": "fUSD"}`))
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		proto.conn.Close()
	}()
	for {
		select {
		case <-ticker.C:
			proto.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := proto.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (proto *BitfinexProtocol) fetchingForever() {
	defer func() {
		proto.conn.Close()
		// added by me?
		proto.restart <- true
	}()
	proto.conn.SetReadLimit(maxMessageSize)
	proto.conn.SetReadDeadline(time.Now().Add(pongWait))
	proto.conn.SetPongHandler(func(string) error { proto.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := proto.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
				proto.restart <- true
				return
			}
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		resp := make([]interface{}, 0)
		log.Printf("Arrived: %s", string(message))
		if err := json.Unmarshal(message, &resp); err != nil {
			log.Errorf("BitfinexProtocol: While receive from websocket: %v", err)
			continue
		}
		event := resp[1].([]interface{})
		if len(event) == 0 {
			continue
		}
		log.Printf("Updating DEFI Rate for %+v\n ", proto.protocol.Name)
		borrowingRate, ok := event[FRR].(float64)
		if !ok {
			log.Errorf("BitfinexProtocol: While parsing borrowing rate: %v", err)
			continue
		}
		borrowingRateFloat := big.NewFloat(borrowingRate)
		lendingRate := new(big.Float).Sub(borrowingRateFloat, new(big.Float).Mul(borrowingRateFloat, big.NewFloat(0.15)))

		lendingRateF64, _ := lendingRate.Float64()
		asset := &dia.DefiRate{
			Timestamp:     time.Now(),
			Asset:         "USD",
			Protocol:      proto.protocol.Name,
			LendingRate:   lendingRateF64,
			BorrowingRate: borrowingRate,
		}
		log.Printf("writing DEFI rate for  %#v in %v\n", asset, proto.scraper.RateChannel())
		proto.scraper.RateChannel() <- asset
		log.Info("Update complete")
		log.Print("Updating DEFI state for %+v\n ", proto.protocol)
		totalSupplyUSD, ok := event[VOLUME].(float64)
		if !ok {
			log.Errorf("BitfinexProtocol: While parsing totalSupplyUSD: %v", err)
			continue
		}
		priceETH, err := utils.GetCoinPrice("ETH")
		if err != nil {
			log.Errorf("BitfinexProtocol: While parsing priceETH: %v", err)
			continue
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
	}
}
func (proto *BitfinexProtocol) connectWebsocket() {
	var err error
	u := url.URL{Scheme: "wss", Host: "api-pub.bitfinex.com", Path: "/ws/2"}
	proto.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	log.Printf("connecting to %s", u.String())
	if err != nil {
		log.Fatal("dial:", err)
		proto.restart <- true
	}
	go proto.pingWebsocket()
	go proto.fetchingForever()

}

func (proto *BitfinexProtocol) UpdateRate() error {
	//NO OP
	return nil
}

func (proto *BitfinexProtocol) UpdateState() error {
	//NO OP
	return nil
}
