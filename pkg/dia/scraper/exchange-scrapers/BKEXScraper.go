package scrapers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	ws "github.com/gorilla/websocket"
)

type BKEXScraper struct {
	wsClient map[int]*ws.Conn
	// signaling channels for session initialization and finishing
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	pairScrapers map[string]*BKEXPairScraper
	exchangeName string
	chanTrades   chan *dia.Trade
	db           *models.RelDB
}

func NewBKEXScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *BKEXScraper {
	s := &BKEXScraper{
		wsClient:     make(map[int]*ws.Conn),
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*BKEXPairScraper),
		exchangeName: exchange.Name,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
		db:           relDB,
	}

	if scrape {
		go s.mainLoop()
	}

	return s
}

type BKEXTradeRecord struct {
	Symbol    string  `json:"symbol"`
	Price     string  `json:"price"`
	Volume    float64 `json:"volume"`
	Direction string  `json:"direction"`
	Ts        int64   `json:"ts"`
}

type BKEXTradeResponse struct {
	quotationAllDeal string
	records          []BKEXTradeRecord
}

func chunkSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for {
		if len(slice) == 0 {
			break
		}

		// necessary check to avoid slicing beyond
		// slice capacity
		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}

		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}

	return chunks
}

func (s *BKEXScraper) connect(i int) *ws.Conn {
	var wsDialer ws.Dialer
	SwConn, _, err := wsDialer.Dial("wss://api.bkex.com/socket.io/?EIO=3&transport=websocket", nil)

	if err != nil {
		println(err.Error())
		return nil
	}
	s.wsClient[i] = SwConn

	// Two time read message
	messageType, p, err := SwConn.ReadMessage()

	if err != nil {
		println(err.Error())
	}

	log.Info("Connected ", messageType, "-", string(p))

	messageType, p, err = SwConn.ReadMessage()

	if err != nil {
		println(err.Error())
	}

	log.Info("Connected ", messageType, "-", string(p))
	// Connect Finished

	// Send 40/quotation and receive it
	SwConn.WriteMessage(ws.TextMessage, []byte("40/quotation"))
	messageType, p, err = SwConn.ReadMessage()
	log.Info("Connected ", messageType, "-", string(p))

	return SwConn
}

func (s *BKEXScraper) subLoop(wsClient *ws.Conn, pairs string) {
	pingTimer := time.NewTicker(25 * time.Second)
	go func() {
		for range pingTimer.C {
			go s.ping(wsClient)
		}
	}()

	message := `42/quotation,["quotationDealConnect",{"symbol": "` + pairs + `","number": 50}]`

	if err := wsClient.WriteMessage(ws.TextMessage, []byte(message)); err != nil {
		log.Error("write pair sub: ", err.Error())
	}
	log.Info("Subscribed to get trades for ", pairs)

	for {
		messageType, p, err := wsClient.ReadMessage()
		if err != nil {
			log.Fatal("read message: ", err.Error())
		}
		if messageType != ws.TextMessage {
			log.Fatal("unknow type ", messageType)
		}

		c := string(p)

		if c == "3" {
			continue
		}

		if len(strings.Split(c, "42/quotation,")) < 2 {
			continue
		}
		d := strings.Split(c, "42/quotation,")[1]

		var r BKEXTradeResponse
		tmp := []interface{}{&r.quotationAllDeal, &r.records}

		json.Unmarshal([]byte(d), &tmp)

		if e := len(tmp); e != 2 {
			log.Fatal("unknow length ", e)
		}

		records := tmp[1].(*[]BKEXTradeRecord)

		for _, trade := range *records {
			var exchangePair dia.ExchangePair
			priceFloat, _ := strconv.ParseFloat(trade.Price, 64)

			exchangePair, err = s.db.GetExchangePairCache(s.exchangeName, trade.Symbol)
			if err != nil {
				log.Error("Get Exchange Pair  ", trade.Symbol)
			}
			volume := trade.Volume
			if trade.Direction == "S" {
				volume *= -1
			}

			t := &dia.Trade{
				Symbol:       strings.Split(trade.Symbol, "_")[0],
				Pair:         trade.Symbol,
				Price:        priceFloat,
				Volume:       volume,
				Time:         time.Unix(0, trade.Ts*int64(time.Millisecond)),
				Source:       s.exchangeName,
				VerifiedPair: exchangePair.Verified,
				BaseToken:    exchangePair.UnderlyingPair.BaseToken,
				QuoteToken:   exchangePair.UnderlyingPair.QuoteToken,
			}

			if exchangePair.Verified {
				log.Infoln("Got verified trade", t)
			}
			s.chanTrades <- t
		}
	}
}

func (s *BKEXScraper) mainLoop() {

	log.Info("Wait 5s untill subscribe all Pairs")
	time.Sleep(5 * time.Second)

	keys := make([]string, 0, len(s.pairScrapers))
	for k := range s.pairScrapers {
		keys = append(keys, k)
	}

	miniPairs := chunkSlice(keys, 10)

	for i, v := range miniPairs {
		log.Info("Connect Websocket ...", i, v)
		time.Sleep(5 * time.Second)
		conn := s.connect(i)
		if conn != nil {
			log.Info("Connect Done Websocket", i, v)
			go s.subLoop(conn, strings.Join(v, ","))
		} else {
			log.Error("Connection Failed !!!", i)
			return
		}
	}
}

func (s *BKEXScraper) ping(conn *ws.Conn) {
	conn.WriteMessage(ws.TextMessage, []byte("2"))
}

// FillSymbolData from MEXCScraper
// @todo more update
func (s *BKEXScraper) FillSymbolData(symbol string) (asset dia.Asset, err error) {
	asset.Symbol = symbol
	return
}

func (s *BKEXScraper) Close() error {
	if s.closed {
		return errors.New("BKEXScraper: Already closed")
	}
	close(s.shutdown)
	for _, c := range s.wsClient {
		err := c.Close()
		if err != nil {
			return err
		}
	}

	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

func (s *BKEXScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return dia.ExchangePair{}, nil
}

func (s *BKEXScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()

	if s.error != nil {
		return nil, s.error
	}

	if s.closed {
		return nil, errors.New("BKEXScraper: Call ScrapePair on closed scraper")
	}

	ps := &BKEXPairScraper{
		parent: s,
		pair:   pair,
	}

	// message := `42/quotation,["quotationDealConnect",{"symbol": "` + pair.ForeignName + `","number": 50}]`
	// // message := `42/quotation,["quotationDealConnect",{"symbol": "BTC_USDT","number": 50}]`
	// log.Info(message)

	// if err := s.wsClient.WriteMessage(ws.TextMessage, []byte(message)); err != nil {
	// 	log.Error("write pair sub: ", err.Error())
	// }
	log.Info("Add to get trades for ", pair.ForeignName)
	s.pairScrapers[pair.ForeignName] = ps
	return ps, nil
}

type BKEXExchangeSymbol struct {
	MinimumOrderSize   float64 `json:"minimumOrderSize"`
	MinimumTradeVolume float64 `json:"minimumTradeVolume"`
	PricePrecision     int     `json:"pricePrecision"`
	SupportTrade       bool    `json:"supportTrade"`
	Symbol             string  `json:"symbol"`
	VolumePrecision    int     `json:"volumePrecision"`
}

type BKEXExchangeInfo struct {
	Code string               `json:"code"`
	Data []BKEXExchangeSymbol `json:"data"`
}

func (s *BKEXScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	var bkexExchangeInfo BKEXExchangeInfo
	response, err := http.Get("https://api.bkex.com/v2/common/symbols")
	if err != nil {
		log.Error("get symbols: ", err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Error("read symbols: ", err)
	}

	err = json.Unmarshal(body, &bkexExchangeInfo)

	if err != nil {
		log.Error("unmarshal symbols: ", err)
	}

	for _, p := range bkexExchangeInfo.Data {
		pairToNormalized := dia.ExchangePair{
			Symbol:      strings.Split(p.Symbol, "_")[0],
			ForeignName: p.Symbol,
			Exchange:    s.exchangeName,
		}
		pairs = append(pairs, pairToNormalized)
	}
	return
}

// Channel returns a channel that can be used to receive trades
func (s *BKEXScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

// BKEXPairScraper implements PairScraper for BKEX
type BKEXPairScraper struct {
	parent *BKEXScraper
	pair   dia.ExchangePair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *BKEXPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *BKEXPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *BKEXPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
