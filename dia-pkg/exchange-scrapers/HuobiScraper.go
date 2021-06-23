package scrapers

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	models "github.com/diadata-org/diadata/pkg/model"
	utils "github.com/diadata-org/diadata/pkg/utils"
	ws "github.com/gorilla/websocket"
)

var _HuobiSocketurl string = "wss://api.huobi.pro/ws"

type EventType struct {
	Sub  string `json:"sub,omitempty"`
	Id   string `json:"id,omitempty"`
	Pong int    `json:"pong,omitempty"`
}

type ResponseType struct {
	Id     int64       `json:"id,omitempty"`
	Status string      `json:"status,omitempty"`
	Subbed string      `json:"subbed,omitempty"`
	Ts     int64       `json:"ts,omitempty"`
	Ping   int         `json:"ping,omitempty"`
	Ch     string      `json:"ch,omitempty"`
	Tick   interface{} `json:"tick,omitempty"`
}

type HuobiCurrency struct {
	Code int `json:"code"`
	Data []struct {
		Currency  string `json:"currency"`
		AssetType int    `json:"assetType"`
		Chains    []struct {
			Chain                  string      `json:"chain"`
			DisplayName            string      `json:"displayName"`
			BaseChain              string      `json:"baseChain"`
			BaseChainProtocol      string      `json:"baseChainProtocol"`
			IsDynamic              bool        `json:"isDynamic"`
			NumOfConfirmations     int         `json:"numOfConfirmations"`
			NumOfFastConfirmations int         `json:"numOfFastConfirmations"`
			DepositStatus          string      `json:"depositStatus"`
			MinDepositAmt          string      `json:"minDepositAmt"`
			WithdrawStatus         string      `json:"withdrawStatus"`
			MinWithdrawAmt         string      `json:"minWithdrawAmt"`
			WithdrawPrecision      int         `json:"withdrawPrecision"`
			MaxWithdrawAmt         string      `json:"maxWithdrawAmt"`
			WithdrawQuotaPerDay    string      `json:"withdrawQuotaPerDay"`
			WithdrawQuotaPerYear   interface{} `json:"withdrawQuotaPerYear"`
			WithdrawQuotaTotal     interface{} `json:"withdrawQuotaTotal"`
			WithdrawFeeType        string      `json:"withdrawFeeType"`
			TransactFeeWithdraw    string      `json:"transactFeeWithdraw"`
			AddrWithTag            bool        `json:"addrWithTag"`
			AddrDepositTag         bool        `json:"addrDepositTag"`
		} `json:"chains"`
		InstStatus string `json:"instStatus"`
	} `json:"data"`
}

type HuobiScraper struct {
	wsClient *ws.Conn
	// signaling channels for session initialization and finishing
	//TODO: Channel not used. Consider removing or refactoring
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	pairScrapers map[string]*HuobiPairScraper
	exchangeName string
	chanTrades   chan *dia.Trade
	db           *models.RelDB
}

// NewHuobiScraper returns a new HuobiScraper for the given pair
func NewHuobiScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *HuobiScraper {

	s := &HuobiScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*HuobiPairScraper),
		exchangeName: exchange.Name,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
		db:           relDB,
	}

	var wsDialer ws.Dialer
	SwConn, _, err := wsDialer.Dial(_HuobiSocketurl, nil)
	if err != nil {
		println(err.Error())
	}
	s.wsClient = SwConn

	if scrape {
		go s.mainLoop()
	}
	return s
}

// runs in a goroutine until s is closed
func (s *HuobiScraper) mainLoop() {
	relDB, err := models.NewRelDataStore()
	if err != nil {
		panic("Couldn't initialize relDB, error: " + err.Error())
	}
	for {
		message := &ResponseType{}
		_, testRead, err := s.wsClient.NextReader()

		if err != nil {
			// Conn errors are non-recoverable.
			// Terminate the routine if theres any error
			fmt.Println(err.Error())
			break
		} else {

			//It has to gzip response data
			reader, _ := gzip.NewReader(testRead)
			jsonBase := json.NewDecoder(reader)
			err := jsonBase.Decode(message)
			if err != nil {
				log.Error(err)
			}

			// If msg is ping type, it needs to resend a pong msg to ws.
			// for avoid to disconnect it
			if message.Ping > 0 {

				a := &EventType{
					Pong: message.Ping,
				}

				if err := s.wsClient.WriteJSON(a); err != nil {
					// Conn errors are non-recoverable.
					// Terminate the routine if theres any error
					fmt.Println(err.Error())
					break
				}
			} else {

				if message.Status == "" {

					var splitString = strings.Split(message.Ch, ".")
					var forName = strings.ToUpper(splitString[1])
					ps, ok := s.pairScrapers[forName]

					if ok {

						md := message.Tick.(map[string]interface{})
						md_data := md["data"].([]interface{})

						for _, value := range md_data {

							md_element := value.(map[string]interface{})
							f64Price := md_element["price"].(float64)
							f64Volume := md_element["amount"].(float64)
							timeStamp := time.Now().UTC()

							if md_element["direction"] == "sell" {
								f64Volume = -f64Volume
							}

							exchangepair, err := relDB.GetExchangePairCache(s.exchangeName, md["symbol"].(string))
							if err != nil {
								log.Error(err)
							}
							// element id is more than int64/uint64 in size
							// leave the id in float64 format
							t := &dia.Trade{
								Symbol:         ps.pair.Symbol,
								Pair:           forName,
								Price:          f64Price,
								Volume:         f64Volume,
								Time:           timeStamp,
								ForeignTradeID: strconv.FormatFloat(md_element["id"].(float64), 'E', -1, 64),
								Source:         s.exchangeName,
								VerifiedPair:   exchangepair.Verified,
								BaseToken:      exchangepair.UnderlyingPair.BaseToken,
								QuoteToken:     exchangepair.UnderlyingPair.QuoteToken,
							}
							ps.parent.chanTrades <- t
							if exchangepair.Verified {
								log.Infoln("Got verified trade", t)
							}
						}
					} else {
						log.Printf("Unknown Pair %v", forName)
					}
				}
			}
		}
	}
	s.cleanup(nil)
}

// FillSymbolData collects all available information on an asset traded on huobi
func (s *HuobiScraper) FillSymbolData(symbol string) (asset dia.Asset, err error) {
	var response HuobiCurrency
	data, _, err := utils.GetRequest("https://api.huobi.pro/v2/reference/currencies?currency=" + symbol)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return
	}

	// Loop through chain if ETH is available put ETH chain details
	// TO DO: This has to be extended. So far, we only have symbol, which we already had before.
	asset.Symbol = response.Data[0].Currency
	asset.Name = response.Data[0].Currency
	return asset, nil
}

func (s *HuobiScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()

	if err != nil {
		s.error = err
	}
	s.closed = true

	close(s.shutdownDone)
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *HuobiScraper) Close() error {

	if s.closed {
		return errors.New("HuobiScraper: Already closed")
	}
	err := s.wsClient.Close()
	if err != nil {
		return err
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *HuobiScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("HuobiScraper: Call ScrapePair on closed scraper")
	}

	ps := &HuobiPairScraper{
		parent: s,
		pair:   pair,
	}
	s.pairScrapers[pair.ForeignName] = ps
	a := &EventType{
		Sub: "market." + strings.ToLower(pair.ForeignName) + ".trade.detail",
		Id:  "id1",
	}
	if err := s.wsClient.WriteJSON(a); err != nil {
		fmt.Println(err.Error())
	}
	return ps, nil
}

func (s *HuobiScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	symbol := strings.ToUpper(pair.Symbol)
	pair.Symbol = symbol

	if helpers.NameForSymbol(symbol) == symbol {
		if !helpers.SymbolIsName(symbol) {
			if pair.Symbol == "IOTA" {
				pair.Symbol = "MIOTA"
			}
			if pair.Symbol == "PROPY" {
				pair.Symbol = "PRO"
			}
			return pair, errors.New("Foreign name can not be normalized:" + pair.ForeignName + " symbol:" + symbol)
		}
	}
	if helpers.SymbolIsBlackListed(symbol) {
		return pair, errors.New("Symbol is black listed:" + symbol)
	}
	return pair, nil

}

// FetchAvailablePairs returns a list with all available trade pairs
func (s *HuobiScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	type DataT struct {
		Id           string `json:"symbol"`
		BaseCurrency string `json:"base-currency"`
	}
	type APIResponse struct {
		Data []DataT `json:"data"`
	}

	data, _, err := utils.GetRequest("http://api.huobi.pro/v1/common/symbols")

	if err != nil {
		return
	}

	var ar APIResponse
	err = json.Unmarshal(data, &ar)
	if err == nil {
		for _, p := range ar.Data {
			pairToNormalize := dia.ExchangePair{
				Symbol:      p.BaseCurrency,
				ForeignName: p.Id,
				Exchange:    s.exchangeName,
			}
			pair, serr := s.NormalizePair(pairToNormalize)
			if serr == nil {
				pairs = append(pairs, pair)
			} else {
				log.Error(serr)
			}
		}
	}
	return
}

// HuobiPairScraper implements PairScraper for Huobi exchange
type HuobiPairScraper struct {
	parent *HuobiScraper
	pair   dia.ExchangePair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *HuobiPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Channel returns a channel that can be used to receive trades
func (ps *HuobiScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *HuobiPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *HuobiPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
