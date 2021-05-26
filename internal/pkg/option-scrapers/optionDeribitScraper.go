package optionscrapers

import (
	"encoding/json"
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	ws "github.com/gorilla/websocket"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

const clientID = "2fSbOqqC"
const clientSecret = "78gpYOitaZkn3y39XboRAWgDrtocdtHVWOz1_IlaN74"


type DeribitInstrumentsResponse struct {
	Jsonrpc string              `json:"jsonrpc"`
	Result  []DeribitInstrument `json:"result"`
	UsIn    int64               `json:"usIn"`
	UsOut   int64               `json:"usOut"`
	UsDiff  int                 `json:"usDiff"`
	Testnet bool                `json:"testnet"`
}

type DeribitInstrument struct {
	TickSize                 float64 `json:"tick_size"`
	TakerCommission          float64 `json:"taker_commission"`
	Strike                   float64 `json:"strike,omitempty"`
	SettlementPeriod         string  `json:"settlement_period"`
	QuoteCurrency            string  `json:"quote_currency"`
	OptionType               string  `json:"option_type,omitempty"`
	MinTradeAmount           int     `json:"min_trade_amount"`
	MakerCommission          float64 `json:"maker_commission"`
	Kind                     string  `json:"kind"`
	IsActive                 bool    `json:"is_active"`
	InstrumentName           string  `json:"instrument_name"`
	ExpirationTimestamp      int64   `json:"expiration_timestamp"`
	CreationTimestamp        int64   `json:"creation_timestamp"`
	ContractSize             float64 `json:"contract_size"`
	BlockTradeCommission     float64 `json:"block_trade_commission"`
	BaseCurrency             string  `json:"base_currency"`
	MaxLiquidationCommission float64 `json:"max_liquidation_commission,omitempty"`
	MaxLeverage              int     `json:"max_leverage,omitempty"`
}

type DeribitETHOptionScraper struct {
	markets            []string
	wsClient           *ws.Conn
	PollFrequency      int8
	ScraperIsRunning   bool
	ScraperIsRunningMu sync.Mutex
	optionsWaitGroup   *sync.WaitGroup
	DataStore          *models.DB
	chanOrderBook      chan *dia.OptionOrderbookDatum
	Ratelimiter        *rate.Limiter
	refreshToken   string
}

type DeribitRequest struct {
	Jsonrpc string       `json:"jsonrpc"`
	ID      int          `json:"id"`
	Method  string       `json:"method"`
	Params  interface{} `json:"params"`
}

type DeribitParam struct {
	Channels []string `json:"channels"`
}

func NewDeribitETHOptionScraper() *DeribitETHOptionScraper {

	deribitAPI := "wss://www.deribit.com/ws/api/v2/"
	ds, err := models.NewDataStore()
	if err != nil {
		log.Errorln("error getting modelstore")
	}

	s := &DeribitETHOptionScraper{
		chanOrderBook: make(chan *dia.OptionOrderbookDatum),
		DataStore:     ds,
	}
	s.GetAndStoreOptionsMeta()

	var wsDialer ws.Dialer
	SwConn, _, err := wsDialer.Dial(deribitAPI, nil)

	if err != nil {
		println(err.Error())
	}
	s.wsClient = SwConn

	return s

}



func (scraper *DeribitETHOptionScraper) subscribe() {

	id := 0
	request := &DeribitRequest{
		Jsonrpc: "2.0",
		ID:      id,
		Method:  "/public/subscribe",
	}
	params := DeribitParam{}
	for _, market := range scraper.markets {
		params.Channels = append(params.Channels, "book."+market+".none.1.100ms")

	}
	//params.Channels = append(params.Channels, "book.ETH-25JUN21-2240-P.none.1.100ms")

	request.Params = params
	log.Info("Subscribing to Instrument ", request)
	scraper.wsClient.WriteJSON(request)
	id++

}

type DeribitOptionResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Channel string `json:"channel"`
		Data    struct {
			Asks           [][]interface{} `json:"asks,omitempty"`
			Bids           [][]interface{} `json:"bids,omitempty"`
			ChangeID       int64           `json:"change_id"`
			InstrumentName string          `json:"instrument_name"`
			PrevChangeID   int64           `json:"prev_change_id"`
			Timestamp      int64           `json:"timestamp"`
			Type           string          `json:"type"`
		} `json:"data"`
	} `json:"params"`
}

func (scraper *DeribitETHOptionScraper) handleWSMessage() {
	var response DeribitOptionResponse
	var err error

 	err = scraper.wsClient.ReadJSON(&response)
	if err != nil {
		log.Errorln("Error reading wsclient", err)
		return

	}

	var (
		resolvedAskPX, resolvedAskSize, resolvedBidSize, resolvedBidPX float64
	)

	everything := "new"

	b, _ := json.Marshal(response)
	log.Infoln("Message Received", string(b[:]))
	for _, askOut := range response.Params.Data.Asks {
		log.Infoln("askOut", askOut)
		switch everything {
		case "delete":
			log.Infoln("Delete")
		case "new":
			log.Infoln("new")
			resolvedAskSize = askOut[1].(float64)
			resolvedAskPX = askOut[0].(float64)
		default:
			log.Infoln(" askOut[0].(string)", askOut[0].(string))

		}

	}

	for _, bidOut := range response.Params.Data.Bids {
		log.Infoln("bidOut", bidOut)
		switch everything {
		case "delete":
			log.Infoln("Delete")
		case "new":
			log.Infoln("new")
			resolvedBidSize = bidOut[1].(float64)
			resolvedBidPX = bidOut[0].(float64)
		default:
			log.Infoln(" bidOut[0].(string)", bidOut[0].(string))

		}

	}

	var o = dia.OptionOrderbookDatum{
		InstrumentName:  response.Params.Data.InstrumentName,
		ObservationTime: time.Unix(response.Params.Data.Timestamp/1e3, 0),
		AskSize:         resolvedAskSize,
		AskPrice:        resolvedAskPX,
		BidPrice:        resolvedBidPX,
		BidSize:         resolvedBidSize,
	}
	log.Infoln("Got trade", o)

	scraper.chanOrderBook <- &o

	//optionMeta := dia.OptionMeta{
	//	InstrumentName: response.Params.Data.InstrumentName,
	//	BaseCurrency:   instrument.SettlementCurrency,
	//	ExpirationTime: expTime,
	//	StrikePrice:    strikePrice,
	//	OptionType:     optionType,
	//}
	//
	//scrapper.DataStore.SetOptionMeta(&optionMeta)

}
func (scraper *DeribitETHOptionScraper) Scrape() {


 	go scraper.heartBeat()
	scraper.subscribe()
	go func() {
		for {
			scraper.handleWSMessage()

		}

	}()

}

func (scraper *DeribitETHOptionScraper) heartBeat() {
	for {
		var params map[string]int

		id := 5600000
		request := &DeribitRequest{
			Jsonrpc: "2.0",
			ID:      id,
			Method:  "/public/set_heartbeat",
		}
		params = make(map[string]int)
		params["interval"] = 30

		request.Params = params
		log.Info("set_heartbeat ", request)
		scraper.wsClient.WriteJSON(request)
		id++
		time.Sleep(1 * time.Minute)
	}
}

func (s *DeribitETHOptionScraper) send(message *map[string]interface{}, websocketConn *ws.Conn) error {
	err := websocketConn.WriteJSON(*message)
	if err != nil {
		return err
	}
	return nil
}

func (s *DeribitETHOptionScraper) handleRefreshToken(previousToken string, websocketConn *ws.Conn) (bool, error) {
	if previousToken == "" {
		return false, nil
	}
	// refresh
	err := s.send(&map[string]interface{}{
		"method": "public/auth",
		"params": &map[string]string{
			"grant_type":    "refresh_token",
			"refresh_token": previousToken,
		},
		"jsonrpc": "2.0",
	}, websocketConn)
	if err != nil {
		return false, err
	}
	return true, nil
}


type DeribitOptionsResponse struct {
	Jsonrpc string              `json:"jsonrpc"`
	Result  []DeribitInstrument `json:"result"`
	UsIn    int64               `json:"usIn"`
	UsOut   int64               `json:"usOut"`
	UsDiff  int                 `json:"usDiff"`
	Testnet bool                `json:"testnet"`
}

func (scraper *DeribitETHOptionScraper) FetchInstruments() {

	var response DeribitOptionsResponse
	rawResponse, err := utils.GetRequest("https://www.deribit.com/api/v2/public/get_instruments?currency=ETH&expired=false&kind=option")
	if err != nil {
		log.Errorln("Error Getting markets", err)

	}

	json.Unmarshal(rawResponse, &response)

	for _, instrument := range response.Result {
		scraper.markets = append(scraper.markets, instrument.InstrumentName)
	}

}

func (scraper *DeribitETHOptionScraper) MetaOnOptionIsAvailable(option DeribitInstrument) (available bool, err error) {
	available = false
	err = nil

	// TODO: can make this faster by specifying BaseCurrency/QuoteCurrency instead
	optionMetas, err := scraper.DataStore.GetOptionMeta(option.BaseCurrency)
	if err != nil {
		return
	}

	for _, optionMeta := range optionMetas {
		if optionMeta.InstrumentName == option.InstrumentName {
			return true, nil
		}
	}

	return
}

func (scraper *DeribitETHOptionScraper) GetAndStoreOptionsMeta() (err error) {
	body, err := utils.GetRequest("https://www.deribit.com/api/v2/public/get_instruments?currency=ETH&expired=false&kind=option")
	if err != nil {
		return
	}

	var decodedMsg DeribitOptionsResponse
	err = json.Unmarshal(body, &decodedMsg)

	if err != nil {
		return
	}

	for _, instrument := range decodedMsg.Result {
		var available bool
		available, err = scraper.MetaOnOptionIsAvailable(instrument)

		if err != nil {
			return
		}

		if !available {

			optionType := dia.CallOption
			if instrument.OptionType == "put" {
				optionType = dia.PutOption
			}

			var expTime time.Time
			expTime = time.Unix(instrument.ExpirationTimestamp/1e3, 0)

			optionMeta := dia.OptionMeta{
				InstrumentName: instrument.InstrumentName,
				BaseCurrency:   instrument.BaseCurrency,
				ExpirationTime: expTime,
				StrikePrice:    instrument.Strike,
				OptionType:     optionType,
			}

			scraper.DataStore.SetOptionMeta(&optionMeta)

		}
	}

	return nil
}

func (scraper *DeribitETHOptionScraper) Channel() chan *dia.OptionOrderbookDatum {
	return scraper.chanOrderBook
}
