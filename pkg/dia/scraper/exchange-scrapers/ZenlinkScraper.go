package scrapers

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	ws "github.com/gorilla/websocket"
)

var (
	NodeScriptPathBifrostKusama   = utils.Getenv("PATH_TO_NODE_SCRIPT", "scripts/bifrost/main.js")
	NodeScriptPathBifrostPolkadot = utils.Getenv("PATH_TO_NODE_SCRIPT", "scripts/bifrost/zenlink-bifrost-polkadot.js")
)

type ZenlinkPairResponse struct {
	Code int           `json:"code"`
	Data []ZenlinkPair `json:"data"`
}

type ZenlinkPair struct {
	Symbol                string `json:"symbol"`
	DisplayName           string `json:"displayName"`
	BaseAsset             string `json:"baseAsset"`
	QuoteAsset            string `json:"quoteAsset"`
	Status                string `json:"status"`
	MinNotional           string `json:"minNotional"`
	MaxNotional           string `json:"maxNotional"`
	MarginTradable        bool   `json:"marginTradable"`
	CommissionType        string `json:"commissionType"`
	CommissionReserveRate string `json:"commissionReserveRate"`
	TickSize              string `json:"tickSize"`
	LotSize               string `json:"lotSize"`
}

type ZenlinkScraper struct {
	exchangeName string

	// channels to signal events
	initDone     chan nothing
	shutdown     chan nothing
	shutdownDone chan nothing

	errorLock sync.RWMutex
	error     error
	closed    bool

	pairScrapers map[string]*ZenlinkPairScraper
	// productPairIds map[string]int
	chanTrades chan *dia.Trade
	wsClient   *ws.Conn
}

func NewZenlinkScraper(exchange dia.Exchange, scrape bool) *ZenlinkScraper {
	var (
		socketURL      string
		nodeScriptPath string
	)

	switch exchange.Name {
	case dia.ZenlinkswapExchange:
		socketURL = "wss://bifrost-rpc.liebi.com/ws"
		nodeScriptPath = NodeScriptPathBifrostKusama
	case dia.ZenlinkswapExchangeBifrostPolkadot:
		socketURL = "wss://hk.p.bifrost-rpc.liebi.com/ws"
		nodeScriptPath = NodeScriptPathBifrostPolkadot
	}
	// establish connection in the background
	var wsDialer ws.Dialer
	wsClient, _, err := wsDialer.Dial(socketURL, nil)
	if err != nil {
		log.Fatal("connect to websocket server: ", err)
	}

	scraper := &ZenlinkScraper{
		exchangeName: exchange.Name,
		wsClient:     wsClient,
		initDone:     make(chan nothing),
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		// productPairIds: make(map[string]int),
		pairScrapers: make(map[string]*ZenlinkPairScraper),
		chanTrades:   make(chan *dia.Trade),
	}

	if scrape {
		go scraper.mainLoop(nodeScriptPath)
	}
	return scraper
}

func (s *ZenlinkScraper) receive(nodeScriptPath string) {
	trades := make(chan string)

	go func() {
		cmd := exec.Command("node", nodeScriptPath)
		stdout, _ := cmd.StdoutPipe()
		err := cmd.Start()
		if err != nil {
			log.Error("start main.js: ", err)
		} else {
			log.Info("started main.js")
		}
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			if strings.HasPrefix(scanner.Text(), "Trade:") {
				trades <- scanner.Text()
			}
			if strings.HasPrefix(scanner.Text(), "blockHeight:") {
				fmt.Println(scanner.Text())
			}
		}
		// Wait for the script to finish
		cmd.Wait()
	}()

	for trade := range trades {
		after := strings.Split(trade, "Trade:")[1]
		fields := strings.Split(after, " ")
		if len(fields) < 8 {
			continue
		}
		FromAmount, err := strconv.ParseFloat(fields[4], 64)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		toAmount, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		price := FromAmount / toAmount
		basetoken := dia.Asset{
			Symbol:     fields[2],
			Blockchain: dia.BIFROST,
			Address:    fields[6],
		}
		quotetoken := dia.Asset{
			Symbol:     fields[1],
			Blockchain: dia.BIFROST,
			Address:    fields[5],
		}
		trade := &dia.Trade{
			Symbol:         fields[1],
			Pair:           fields[0],
			Price:          price,
			Volume:         toAmount,
			Time:           time.Now(),
			ForeignTradeID: fields[7],
			Source:         s.exchangeName,
			BaseToken:      basetoken,
			QuoteToken:     quotetoken,
			VerifiedPair:   true,
		}

		log.Info("Got Trade: ", trade)
		s.chanTrades <- trade
	}
}

func (s *ZenlinkScraper) mainLoop(nodeScriptPath string) {
	for {
		select {
		case <-s.shutdown:
			log.Warn("Shutting down ZenlinkScraper")
			s.cleanup(nil)
			return
		default:
		}
		s.receive(nodeScriptPath)
	}
}

func (scraper *ZenlinkScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	scraper.errorLock.RLock()
	defer scraper.errorLock.RUnlock()

	if scraper.error != nil {
		return nil, scraper.error
	}

	if scraper.closed {
		return nil, errors.New("ZenlinkScraper is closed")
	}

	pairScraper := &ZenlinkPairScraper{
		parent: scraper,
		pair:   pair,
	}

	scraper.pairScrapers[pair.ForeignName] = pairScraper
	return pairScraper, nil
}

func (s *ZenlinkScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	var zenlinkResponse ZenlinkPairResponse
	body := `{
    "code": 0,
    "data": [
        {
			symbol: "vKSM/KSM",
			displayName: "vKSM/KSM",
			baseAsset: "vKSM",
			quoteAsset: "KSM"
    ]
}`
	err = json.Unmarshal([]byte(body), &zenlinkResponse)
	if err != nil {
		log.Error("unmarshal symbols: ", err)
	}

	for _, p := range zenlinkResponse.Data {
		pairToNormalize := dia.ExchangePair{
			Symbol:      strings.Split(p.Symbol, "/")[0],
			ForeignName: p.Symbol,
			Exchange:    s.exchangeName,
		}
		pairs = append(pairs, pairToNormalize)
	}
	return
}

func (s *ZenlinkScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return dia.ExchangePair{}, nil
}

func (s *ZenlinkScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

func (s *ZenlinkScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()
	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.shutdownDone)
}

func (s *ZenlinkScraper) Close() error {
	if s.closed {
		return errors.New("ZenlinkScraper: Already closed")
	}
	if err := s.wsClient.Close(); err != nil {
		log.Error("Error closing Zenlink.wsClient", err)
	}
	close(s.shutdown)
	<-s.shutdownDone
	defer s.errorLock.RUnlock()
	return s.error
}

type ZenlinkPairScraper struct {
	parent *ZenlinkScraper
	pair   dia.ExchangePair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *ZenlinkPairScraper) Close() error {
	var err error
	s := ps.parent
	// if parent already errored, return early
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return s.error
	}
	if ps.closed {
		return errors.New("ZenlinkPairScraper: Already closed")
	}
	// TODO stop collection for the pair

	ps.closed = true
	return err
}

// Channel returns a channel that can be used to receive trades
func (ps *ZenlinkScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *ZenlinkPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *ZenlinkPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
