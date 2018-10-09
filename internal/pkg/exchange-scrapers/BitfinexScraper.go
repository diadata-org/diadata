package scrapers

import (
	"context"
	"errors"
	bitfinex "github.com/bitfinexcom/bitfinex-api-go/v2"
	"github.com/bitfinexcom/bitfinex-api-go/v2/rest"
	"github.com/bitfinexcom/bitfinex-api-go/v2/websocket"
	"github.com/diadata-org/api-golang/pkg/dia"
	"log"
	"strconv"
	"sync"
	"time"
)

type pairScraperSet map[*BitfinexPairScraper]nothing

// BitfinexScraper is a Scraper for collecting trades from the Bitfinex websocket API
type BitfinexScraper struct {
	// the websocket connection to the Bitfinex API
	wsClient   *websocket.Client
	restClient *rest.Client
	// signaling channels for session initialization and finishing
	initDone     chan nothing
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	// use sync.Maps to concurrently handle multiple pairs
	pairScrapers      sync.Map          // dia.Pair -> pairScraperSet
	pairSubscriptions sync.Map          // dia.Pair -> string (subscription ID)
	pairLocks         sync.Map          // dia.Pair -> sync.Mutex
	symbols           map[string]string // pair to symbol mapping
}

// NewBitfinexScraper returns a new BitfinexScraper for the given pair
func NewBitfinexScraper(key string, secret string) *BitfinexScraper {
	// we want to ensure there are no gaps in our stream
	// -> close the returned channel on disconnect, forcing the caller to handle
	// possible gaps
	params := websocket.NewDefaultParameters()
	//TODO: Set to false again because now we can have holes in our data stream
	params.AutoReconnect = true
	// params.HeartbeatTimeout = 5 * time.Second // used for testing

	s := &BitfinexScraper{
		wsClient:     websocket.NewWithParams(params),
		restClient:   rest.NewClient().Credentials(key, secret),
		initDone:     make(chan nothing),
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		symbols:      make(map[string]string),
		error:        nil,
	}

	// establish connection in the background
	go s.mainLoop()
	return s
}

// runs in a goroutine until s is closed
func (s *BitfinexScraper) mainLoop() {
	err := s.wsClient.Connect()
	listener := s.wsClient.Listen()
	close(s.initDone)
	if err != nil {
		s.cleanup(err)
		return
	}
	for {
		select {
		case msg, ok := <-listener:
			if ok {
				//	log.Printf("MSG RECV: %#v\n", msg)
				// find out message type
				switch m := msg.(type) {
				case *bitfinex.Trade:
					volume := m.Amount
					if m.Side != bitfinex.Bid {
						volume = -volume
					}

					// parse trade data structure
					t := &dia.Trade{
						Symbol:         s.symbols[m.Pair],
						Pair:           m.Pair,
						Price:          m.Price,
						Volume:         volume,
						Time:           time.Unix(m.MTS/1000, (m.MTS%1000)*int64(time.Millisecond)),
						ForeignTradeID: strconv.FormatInt(m.ID, 16),
						Source:         dia.BitfinexExchange,
					}
					// get lock for pair
					pairLock, ok := s.pairLocks.Load(m.Pair)
					if !ok {
						log.Println("t.Pair:", m.Pair)
						s.cleanup(errors.New("BitfinexScraper: Received trade for pair not subscribed to"))
						return
					}
					func() { // wrap in func, allowing to unlock using defer
						pairLock.(*sync.Mutex).Lock()
						defer pairLock.(*sync.Mutex).Unlock()
						// send trade to all registered pairscrapers
						if pairScrapers, ok := s.pairScrapers.Load(t.Pair); ok {
							for ps := range pairScrapers.(pairScraperSet) {
								ps.chanTrades <- t
							}
						}
					}()
				case error:
					s.cleanup(m)
					return
				}
			} else {
				s.cleanup(errors.New("BitfinexScraper: Listener channel was closed unexpectedly"))
				return
			}
		case <-s.shutdown: // user requested shutdown
			log.Println("BitfinexScraper shutting down")
			s.cleanup(nil)
			return
		}
	}
}

// closes all connected PairScrapers
// must only be called from mainLoop
func (s *BitfinexScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()
	// close all channels of PairScraper children
	s.pairScrapers.Range(func(k, v interface{}) bool {
		for ps := range v.(pairScraperSet) {
			close(ps.chanTrades)
			ps.closed = true
		}
		s.pairScrapers.Delete(k)
		return true
	})
	if s.wsClient.IsConnected() {
		s.wsClient.Close()
	}
	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.shutdownDone) // signal that shutdown is complete
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *BitfinexScraper) Close() error {
	if s.closed {
		return errors.New("BitfinexScraper: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *BitfinexScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {
	<-s.initDone // wait until wsClient is connected
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("BitfinexScraper: Call ScrapePair on closed scraper")
	}
	ps := &BitfinexPairScraper{
		parent:     s,
		pair:       pair,
		chanTrades: make(chan *dia.Trade),
	}

	s.symbols[pair.ForeignName] = pair.Symbol

	// get lock for pair
	pairLock, _ := s.pairLocks.LoadOrStore(pair.ForeignName, &sync.Mutex{})
	pairLock.(*sync.Mutex).Lock()
	defer pairLock.(*sync.Mutex).Unlock()
	// initialize pairScraperSet for pair if not already done
	pairScrapers, _ := s.pairScrapers.LoadOrStore(pair.ForeignName, pairScraperSet{})
	// register ps
	pairScrapers.(pairScraperSet)[ps] = nothing{}
	// subscribe to trading pair if we are the first scraper for this pair
	if _, ok := s.pairSubscriptions.Load(pair.ForeignName); !ok {
		ctx1, ctx1cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer ctx1cancel()
		id, err := s.wsClient.SubscribeTrades(ctx1, pair.ForeignName)
		if err != nil {
			// well that didn't work -> cleanup and return error
			delete(pairScrapers.(pairScraperSet), ps)
			return nil, err
		}
		s.pairSubscriptions.Store(pair.ForeignName, id)
	}
	return ps, nil
}

// BitfinexPairScraper implements PairScraper for Bitfinex
type BitfinexPairScraper struct {
	parent     *BitfinexScraper
	pair       dia.Pair
	chanTrades chan *dia.Trade
	closed     bool
}

// Close stops listening for trades of the pair associated with s
func (ps *BitfinexPairScraper) Close() error {
	var err error
	s := ps.parent
	// if parent already errored, return early
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return s.error
	}
	if ps.closed {
		return errors.New("BitfinexPairScraper: Already closed")
	}
	// get lock for pair
	pairLock, _ := s.pairLocks.Load(ps.pair.Symbol)
	pairLock.(*sync.Mutex).Lock()
	defer pairLock.(*sync.Mutex).Unlock()
	pairScrapers, ok := s.pairScrapers.Load(ps.pair.Symbol)
	if !ok { // should never happen
		panic("BitfinexPairScraper: pairScraperSet not found")
	}
	// deregister and close channel
	delete(pairScrapers.(pairScraperSet), ps)
	close(ps.chanTrades)
	// if we're the last one for this pair -> unsubscribe
	if len(pairScrapers.(pairScraperSet)) == 0 {
		id, ok := s.pairSubscriptions.Load(ps.pair.Symbol)
		if !ok { // should never happen
			panic("BitfinexPairScraper: Subscription ID not found")
		}
		ctx1, ctx1cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer ctx1cancel()
		err = s.wsClient.Unsubscribe(ctx1, id.(string))
	}
	ps.closed = true
	return err
}

// Channel returns a channel that can be used to receive trades
func (ps *BitfinexPairScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *BitfinexPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *BitfinexPairScraper) Pair() dia.Pair {
	return ps.pair
}
