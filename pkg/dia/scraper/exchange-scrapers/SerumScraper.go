package scrapers

import (
	"context"
	"errors"
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/mr-tron/base58"
	bin "github.com/streamingfast/binary"
	"github.com/streamingfast/solana-go"
	"github.com/streamingfast/solana-go/programs/serum"
	"github.com/streamingfast/solana-go/rpc"
)

const (
	rpcEndpointSolana         = ""
	dexProgramAddress         = "9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin" // refer - https://github.com/project-serum/serum-dex
	nameServiceProgramAddress = "namesLPneVptA9Z5rqUDD9tMTWEJwofgaYwp8cawRkX"
	dotTokenTLD               = "6NSu2tci4apRKQtt257bAVcvqYjB3zV2H1dWo56vgpa6"
	marketDataSize            = 388
	uint32Modulo              = 4294967296
)

type serumMarket struct {
	market     *serum.MarketV2
	baseAsset  tokenMeta
	quoteAsset tokenMeta
}

type SerumEventQueueHeader struct {
	SerumPadding [5]byte `json:"-"`
	AccountFlags serum.AccountFlag
	Head         uint32
	Zeros1       [4]byte
	Count        uint32
	Zeros2       [4]byte
	Seq          uint32
	Zeros3       [4]byte
}

type EventFlag uint8

const (
	Fill = EventFlag(1 << iota)
	Out
	Bid
	Maker
)

func (e *EventFlag) Is(flag EventFlag) bool { return *e&flag != 0 }

type SerumEvent struct {
	EventFlags             EventFlag
	OpenOrderSlots         uint8
	FeeTier                uint8
	Blob                   [5]byte `json:"-"`
	NativeQuantityReleased bin.Uint64
	NativeQuantityPaid     bin.Uint64
	NativeFeeOrRebate      bin.Uint64
	OrderId                bin.Uint128
	OpenOrders             solana.PublicKey
	ClientOrderId          bin.Uint64
}

type SerumScraper struct {
	solanaRpcClient *rpc.Client
	// signaling channels for session initialization and finishing
	run          bool
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	pairScrapers map[string]*SerumPairScraper
	exchangeName string
	chanTrades   chan *dia.Trade
	waitTime     int
}

func NewSerumScraper(exchange dia.Exchange, scrape bool) *SerumScraper {

	scraper := &SerumScraper{
		solanaRpcClient: rpc.NewClient(utils.Getenv("SOLANA_URI_REST", rpcEndpointSolana)),
		shutdown:        make(chan nothing),
		shutdownDone:    make(chan nothing),
		errorLock:       sync.RWMutex{},
		pairScrapers:    make(map[string]*SerumPairScraper),
		exchangeName:    exchange.Name,
		chanTrades:      make(chan *dia.Trade),
	}
	if scrape {
		go scraper.mainLoop()
	}
	return scraper
}

func (s *SerumScraper) mainLoop() {
	lastSeqNo := make(map[string]uint32)
	scraperInitialized := make(map[string]bool)
	markets := make(map[string]serumMarket)
	s.run = true
	for s.run {
		if len(s.pairScrapers) == 0 {
			s.error = errors.New("no pairs to scrape provided")
			log.Error(s.error.Error())
			break
		}
		serumMarkets, err := s.getMarkets()
		if err != nil {
			log.Error("get markets: ", err)
			return
		}
		tokenNameRegistry, err := s.getTokenNames()
		if err != nil {
			log.Error("get token names: ", err)
			return
		}
		for _, market := range serumMarkets {
			baseToken, baseTokenValid := tokenNameRegistry[market.BaseMint.String()]
			quoteToken, quoteTokenValid := tokenNameRegistry[market.QuoteMint.String()]
			if baseTokenValid && quoteTokenValid {
				marketName := baseToken.symbol + "/" + quoteToken.symbol
				markets[marketName] = serumMarket{
					market:     market,
					baseAsset:  baseToken,
					quoteAsset: quoteToken,
				}
			}
		}
		for pair := range s.pairScrapers {
			if marketForPair, ok := markets[pair]; ok {
				eventQueue, err := s.getEvents(marketForPair.market.EventQueue)
				if err != nil {
					continue
				}
				eventStart := 37
				eventSize := 88
				headerData := eventQueue[:eventStart]
				var header SerumEventQueueHeader
				err = bin.NewDecoder(headerData).Decode(&header)
				if err != nil {
					log.Error("unable to parse event header: ", err)
					continue
				}
				eventCount := (len(eventQueue) - eventStart - 7) / eventSize
				currentEventNo := 0
				startIndex := 1
				endIndex := eventCount
				if lastSeqNo[pair] != 0 {
					missedEvents := (int(header.Seq-lastSeqNo[pair]) + uint32Modulo) % uint32Modulo
					endIndex = int(header.Head+header.Count) % eventCount
					startIndex = (endIndex - missedEvents + eventCount) % eventCount
					if startIndex < 0 {
						startIndex = 1
					}
				}
				lastSeqNo[pair] = header.Seq
				if !scraperInitialized[pair] {
					scraperInitialized[pair] = true
					continue
				}
				for i := eventStart; i < len(eventQueue); {
					eventEnd := i + eventSize
					if eventEnd >= len(eventQueue) {
						break
					}
					eventData := eventQueue[i:eventEnd]
					i = eventEnd
					currentEventNo++
					if currentEventNo > endIndex {
						break
					}
					if currentEventNo >= startIndex {
						var event SerumEvent
						err = bin.NewDecoder(eventData).Decode(&event)
						if err != nil {
							log.Error("unable to parse event data: ", err)
							continue
						}
						if event.EventFlags.Is(Fill) && event.EventFlags.Is(Bid) {
							volume, price := parseEvent(&event, math.Pow10(int(marketForPair.baseAsset.decimals)), math.Pow10(int(marketForPair.quoteAsset.decimals)))
							// Remark: base and quote token is used the other way around by serum dex than we do at DIA.
							trade := dia.Trade{
								Symbol: marketForPair.baseAsset.symbol,
								Pair:   pair,
								BaseToken: dia.Asset{
									Symbol:     marketForPair.quoteAsset.symbol,
									Name:       marketForPair.quoteAsset.name,
									Address:    marketForPair.quoteAsset.mint,
									Decimals:   marketForPair.quoteAsset.decimals,
									Blockchain: dia.SOLANA,
								},
								QuoteToken: dia.Asset{
									Symbol:     marketForPair.baseAsset.symbol,
									Name:       marketForPair.baseAsset.name,
									Address:    marketForPair.baseAsset.mint,
									Decimals:   marketForPair.baseAsset.decimals,
									Blockchain: dia.SOLANA,
								},
								Price:          price,
								Volume:         volume,
								Time:           time.Now(),
								ForeignTradeID: event.OrderId.DecimalString(),
								Source:         s.exchangeName,
								VerifiedPair:   true,
							}
							log.Infof("got trade -- timestamp: %v, symbol: %s, pair: %s, price %v, volume: %v, tradeID: %s", trade.Time, trade.Symbol, trade.Pair, trade.Price, trade.Volume, trade.ForeignTradeID)
							s.chanTrades <- &trade
						}
					}
				}
			}
		}
		time.Sleep(time.Duration(60) * time.Second)
	}
	if s.error == nil {
		s.error = errors.New("main loop terminated by Close()")
	}
	s.cleanup(nil)
}

func (s *SerumScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	serumMarkets, err := s.getMarkets()
	if err != nil {
		log.Error(err)
		return
	}
	tokenNameRegistry, err := s.getTokenNames()
	if err != nil {
		log.Error(err)
		return
	}
	for _, market := range serumMarkets {
		var quoteToken, baseToken dia.Asset
		var quoteTokenValid, baseTokenValid bool
		var pairName string
		if tokenInfo, ok := tokenNameRegistry[market.BaseMint.String()]; ok {
			baseToken = dia.Asset{
				Symbol:     tokenInfo.symbol,
				Name:       tokenInfo.name,
				Address:    tokenInfo.mint,
				Decimals:   tokenInfo.decimals,
				Blockchain: Exchanges[s.exchangeName].BlockChain.Name,
			}
			baseTokenValid = true
			pairName = pairName + tokenInfo.symbol
		}
		if tokenInfo, ok := tokenNameRegistry[market.QuoteMint.String()]; ok {
			quoteToken = dia.Asset{
				Symbol:     tokenInfo.symbol,
				Name:       tokenInfo.name,
				Address:    tokenInfo.mint,
				Decimals:   tokenInfo.decimals,
				Blockchain: Exchanges[s.exchangeName].BlockChain.Name,
			}
			quoteTokenValid = true
			pairName = pairName + "/" + tokenInfo.symbol
		}
		if quoteTokenValid && baseTokenValid {
			pair := dia.ExchangePair{
				Symbol:         baseToken.Name,
				ForeignName:    pairName,
				Exchange:       "Serum",
				Verified:       true,
				UnderlyingPair: dia.Pair{BaseToken: baseToken, QuoteToken: quoteToken},
			}
			pairs = append(pairs, pair)
		}
	}
	return
}

func (s *SerumScraper) getEvents(eventQueueAddr solana.PublicKey) (eventQueue solana.Data, err error) {
	acctInfo, err := s.solanaRpcClient.GetAccountInfo(context.TODO(), eventQueueAddr)
	if err != nil {
		return nil, fmt.Errorf("unable to get events:%w", err)
	}
	return acctInfo.Value.Data, nil
}

func parseEvent(e *SerumEvent, baseMultiplier, quoteMultiplier float64) (volume, price float64) {
	var priceBeforeFees float64
	if e.EventFlags.Is(Bid) {
		if e.EventFlags.Is(Maker) {
			priceBeforeFees = float64(e.NativeQuantityPaid + e.NativeFeeOrRebate)
		} else {
			priceBeforeFees = float64(e.NativeQuantityPaid - e.NativeFeeOrRebate)
		}
		if quoteMultiplier*float64(e.NativeQuantityReleased) != 0 {
			price = (priceBeforeFees * baseMultiplier) / (quoteMultiplier * float64(e.NativeQuantityReleased))
		}
		volume = float64(e.NativeQuantityReleased) / baseMultiplier
	} else {
		if e.EventFlags.Is(Maker) {
			priceBeforeFees = float64(e.NativeQuantityReleased - e.NativeFeeOrRebate)
		} else {
			priceBeforeFees = float64(e.NativeQuantityReleased + e.NativeFeeOrRebate)
		}
		if quoteMultiplier*float64(e.NativeQuantityPaid) != 0 {
			price = (priceBeforeFees * baseMultiplier) / (quoteMultiplier * float64(e.NativeQuantityPaid))
		}
		volume = float64(e.NativeQuantityPaid) / baseMultiplier
	}
	return
}

func (s *SerumScraper) getMarkets() ([]*serum.MarketV2, error) {
	resp, err := s.solanaRpcClient.GetProgramAccounts(
		context.TODO(),
		solana.MustPublicKeyFromBase58(dexProgramAddress),
		&rpc.GetProgramAccountsOpts{
			Filters: []rpc.RPCFilter{
				{
					DataSize: marketDataSize,
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}
	out := make([]*serum.MarketV2, 0)
	for _, keyedAcct := range resp {
		acct := keyedAcct.Account
		marketV2 := &serum.MarketV2{}
		if err := marketV2.Decode(acct.Data); err != nil {
			return nil, fmt.Errorf("decoding market v2: %w", err)
		}
		out = append(out, marketV2)
	}
	return out, nil
}

func (s *SerumScraper) getTokenNames() (map[string]tokenMeta, error) {
	names := make(map[string]tokenMeta)
	tldPublicKey := solana.MustPublicKeyFromBase58(dotTokenTLD)
	resp, err := s.solanaRpcClient.GetProgramAccounts(
		context.TODO(),
		solana.MustPublicKeyFromBase58(nameServiceProgramAddress),
		&rpc.GetProgramAccountsOpts{
			Filters: []rpc.RPCFilter{
				{
					Memcmp: &rpc.RPCFilterMemcmp{
						Bytes: tldPublicKey[:],
					},
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}
	for _, keyedAcct := range resp {
		if t, ok := extractTokenMetaFromData(keyedAcct.Account.Data[96:]); ok {
			names[t.mint] = t
		}
	}
	return names, nil
}

type tokenMeta struct {
	name     string
	symbol   string
	mint     string
	decimals uint8
}

func extractTokenMetaFromData(data []byte) (tokenMeta, bool) {
	var t tokenMeta
	if len(data) > 0 {
		nameSize := int(data[0])
		nameStart := 4
		nameEnd := nameStart + nameSize
		if len(data) > nameEnd {
			t.name = string(data[nameStart:nameEnd])
			symbolSize := int(data[nameEnd])
			symbolStart := 4 + nameEnd
			symbolEnd := symbolStart + symbolSize
			if len(data) > symbolEnd {
				t.symbol = string(data[symbolStart:symbolEnd])
				mintSize := 32
				mintStart := symbolEnd
				mintEnd := mintStart + mintSize
				if len(data) > mintEnd {
					t.mint = base58.Encode(data[mintStart:mintEnd])
					t.decimals = data[mintEnd]
					return t, true
				}
			}
		}
	}
	return tokenMeta{}, false
}

func (s *SerumScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{}, nil
}

func (s *SerumScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

func (s *SerumScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

func (s *SerumScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()
	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.shutdownDone)
}

func (s *SerumScraper) Close() error {
	// close the pair scraper channels
	s.run = false
	for _, pairScraper := range s.pairScrapers {
		pairScraper.closed = true
	}

	close(s.shutdown)
	<-s.shutdownDone
	return nil
}

func (s *SerumScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("SerumScraper: Call ScrapePair on closed scraper")
	}
	ps := &SerumPairScraper{
		parent: s,
		pair:   pair,
	}
	s.pairScrapers[pair.ForeignName] = ps
	return ps, nil
}

// SerumPairScraper implements PairScraper for Serum
type SerumPairScraper struct {
	parent *SerumScraper
	pair   dia.ExchangePair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (sps *SerumPairScraper) Close() error {
	sps.parent.errorLock.RLock()
	defer sps.parent.errorLock.RUnlock()
	sps.closed = true
	return nil
}

// Channel returns a channel that can be used to receive trades
func (sps *SerumPairScraper) Channel() chan *dia.Trade {
	return sps.parent.Channel()
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (sps *SerumPairScraper) Error() error {
	s := sps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (sps *SerumPairScraper) Pair() dia.ExchangePair {
	return sps.pair
}
