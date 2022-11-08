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
	solanav2 "github.com/gagliardetto/solana-go"
	solanawsclient "github.com/gagliardetto/solana-go/rpc/ws"

	"github.com/mr-tron/base58"
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

type SerumScraper struct {
	solanaRpcClient *rpc.Client
	solanaWSClient  *solanawsclient.Client
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

	wsclient, err := solanawsclient.Connect(context.Background(), utils.Getenv("SOLANA_URI_WS", rpcEndpointSolana))
	if err != nil {
		log.Errorln("error connecting wsclient", err)
	}

	scraper := &SerumScraper{
		solanaRpcClient: rpc.NewClient(utils.Getenv("SOLANA_URI_REST", rpcEndpointSolana)),
		solanaWSClient:  wsclient,
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
	// lastSeqNo := make(map[string]uint32)
	// scraperInitialized := make(map[string]bool)
	wg := sync.WaitGroup{}
	markets := make(map[string]serumMarket)
	s.run = true
	for s.run {
		// if len(s.pairScrapers) == 0 {
		// 	s.error = errors.New("no pairs to scrape provided")
		// 	log.Error(s.error.Error())
		// 	break
		// }
		timestart := time.Now()
		serumMarkets, err := s.getMarkets()
		if err != nil {
			log.Error("get markets: ", err)
			return
		}

		timeend := time.Now().Unix() - timestart.Unix()
		log.Infoln("time spend to get markets: ", timeend)
		log.Infoln("Total serumMarkets: ", len(serumMarkets))

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
		for pair, _ := range markets {
			if marketForPair, ok := markets[pair]; ok {

				wg.Add(1)

				go func(k string) {
					{
						log.Infoln("subscribing ", k)

						accountID := solanav2.MustPublicKeyFromBase58(marketForPair.market.EventQueue.String())
						ac, err := s.solanaWSClient.AccountSubscribe(accountID, "")
						if err != nil {
							log.Errorln("error on AccountSubscribe", err)
						}
						for {
							result, err := ac.Recv()
							if err != nil {
								log.Errorf("error on recv %v for pair %s  unsubscribing", err, k)
								continue
							}

							databytes := result.Value.Account.Data.GetBinary()

							eq := serum.EventQueue{}
							eq.Decode(databytes)

							log.Infof("got event for %s  total events %d and account flag is %s  head %d sequence number %d and accountID %v", k, len(eq.Events), eq.AccountFlags.String(), eq.Head, eq.SeqNum, accountID)

							for _, event := range eq.Events {
								// log.Infoln("event", event)

								if event.Flag.IsFill() && event.Flag.IsBid() {
									volume, price := parseEvent(event, math.Pow10(int(marketForPair.baseAsset.decimals)), math.Pow10(int(marketForPair.quoteAsset.decimals)))
									// Remark: base and quote token is used the other way around by serum dex than we do at DIA.
									trade := dia.Trade{
										Symbol: marketForPair.baseAsset.symbol,
										Pair:   k,
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
										ForeignTradeID: event.OrderID.HexString(false),
										Source:         s.exchangeName,
										VerifiedPair:   true,
									}

									log.Errorf("got trade -- timestamp: %v, symbol: %s, pair: %s, price %v, volume: %v, tradeID: %s", trade.Time, trade.Symbol, trade.Pair, trade.Price, trade.Volume, trade.ForeignTradeID)
									s.chanTrades <- &trade

								}
							}

						}

					}
				}(pair)
			}
		}
		wg.Wait()
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
	acctInfo, err := s.solanaRpcClient.GetAccountInfo(eventQueueAddr)
	if err != nil {
		return nil, fmt.Errorf("unable to get events:%w", err)
	}
	return acctInfo.Value.Data, nil
}

func parseEvent(e *serum.Event, baseMultiplier, quoteMultiplier float64) (volume, price float64) {
	var priceBeforeFees float64
	if e.Flag.IsBid() {
		if e.Flag.IsMaker() {
			priceBeforeFees = float64(e.NativeQtyPaid + e.NativeFeeOrRebate)
		} else {
			priceBeforeFees = float64(e.NativeQtyPaid - e.NativeFeeOrRebate)
		}
		if quoteMultiplier*float64(e.NativeQtyReleased) != 0 {
			price = (priceBeforeFees * baseMultiplier) / (quoteMultiplier * float64(e.NativeQtyReleased))
		}
		volume = float64(e.NativeQtyReleased) / baseMultiplier
	} else {
		if e.Flag.IsMaker() {
			priceBeforeFees = float64(e.NativeQtyReleased - e.NativeFeeOrRebate)
		} else {
			priceBeforeFees = float64(e.NativeQtyReleased + e.NativeFeeOrRebate)
		}
		if quoteMultiplier*float64(e.NativeQtyPaid) != 0 {
			price = (priceBeforeFees * baseMultiplier) / (quoteMultiplier * float64(e.NativeQtyPaid))
		}
		volume = float64(e.NativeQtyPaid) / baseMultiplier
	}
	return
}

func (s *SerumScraper) getMarkets() ([]*serum.MarketV2, error) {
	resp, err := s.solanaRpcClient.GetProgramAccounts(
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
	// tldPublicKey := solana.MustPublicKeyFromBase58(dotTokenTLD)
	// resp, err := s.solanaRpcClient.GetProgramAccounts(
	// 	solana.MustPublicKeyFromBase58(nameServiceProgramAddress),
	// 	&rpc.GetProgramAccountsOpts{
	// 		Filters: []rpc.RPCFilter{
	// 			{
	// 				Memcmp: &rpc.RPCFilterMemcmp{
	// 					Bytes: tldPublicKey[:],
	// 				},
	// 			},
	// 		},
	// 	},
	// )
	// if err != nil {
	// 	return nil, err
	// }
	// for _, keyedAcct := range resp {
	// 	if t, ok := extractTokenMetaFromData(keyedAcct.Account.Data[96:]); ok {
	// 		names[t.mint] = t
	// 	}
	// }

	names["mSoLzYCxHdYgdzU16g5QSh3i5K3z3KZK7ytfqcJm7So"] = tokenMeta{name: "Marinade staked SOL (mSOL)", symbol: "mSOL", decimals: 9, mint: "mSoLzYCxHdYgdzU16g5QSh3i5K3z3KZK7ytfqcJm7So"}
	names["So11111111111111111111111111111111111111112"] = tokenMeta{name: "Wrapped SOL", symbol: "SOL", decimals: 9, mint: "So11111111111111111111111111111111111111112"}
	names["4k3Dyjzvzp8eMZWUXbBCjEvwSkkk59S5iCNLY3QrkX6R"] = tokenMeta{name: "Raydium", symbol: "RAY", decimals: 6, mint: "4k3Dyjzvzp8eMZWUXbBCjEvwSkkk59S5iCNLY3QrkX6R"}
	names["8Yv9Jz4z7BUHP68dz8E8m3tMe6NKgpMUKn8KVqrPA6Fr"] = tokenMeta{name: "Wrapped USDC (Allbridge from Avalanche)", symbol: "aaUSDC", decimals: 9, mint: "8Yv9Jz4z7BUHP68dz8E8m3tMe6NKgpMUKn8KVqrPA6Fr"}
	names["Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB"] = tokenMeta{name: "USDT", symbol: "USDT", decimals: 6, mint: "Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB"}
	names["EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v"] = tokenMeta{name: "USD Coin", symbol: "USDC", decimals: 6, mint: "8Yv9Jz4z7BUHP68dz8E8m3tMe6NKgpMUKn8KVqrPA6Fr"}
	names["DubwWZNWiNGMMeeQHPnMATNj77YZPZSAz2WVR5WjLJqz"] = tokenMeta{name: "CropperFinance", symbol: "CRP", decimals: 9, mint: "DubwWZNWiNGMMeeQHPnMATNj77YZPZSAz2WVR5WjLJqz"}
	names["7xKXtg2CW87d97TXJSDpbD5jBkheTqA83TZRuJosgAsU"] = tokenMeta{name: "Samoyed Coin", symbol: "SAMO", decimals: 9, mint: "7xKXtg2CW87d97TXJSDpbD5jBkheTqA83TZRuJosgAsU"}
	names["kinXdEcpDQeHPEuQnqmUgtYykqKGVFq6CeVX5iAHJq6"] = tokenMeta{name: "KIN", symbol: "KIN", decimals: 5, mint: "kinXdEcpDQeHPEuQnqmUgtYykqKGVFq6CeVX5iAHJq6"}
	names["DUSTawucrTsGU8hcqRdHDCbuYhCPADMLM2VcCb8VnFnQ"] = tokenMeta{name: "DUST Protocol", symbol: "DUST", decimals: 9, mint: "DUSTawucrTsGU8hcqRdHDCbuYhCPADMLM2VcCb8VnFnQ"}
	names["USDH1SM1ojwWUga67PGrgFWUHibbjqMvuMaDkRJTgkX"] = tokenMeta{name: "USDH Hubble Stablecoin", symbol: "USDH", decimals: 6, mint: "USDH1SM1ojwWUga67PGrgFWUHibbjqMvuMaDkRJTgkX"}
	names["5goWRao6a3yNC4d6UjMdQxonkCMvKBwdpubU3qhfcdf1"] = tokenMeta{name: "Tether USD (Portal from Polygon)", symbol: "USDTpo", decimals: 6, mint: "5goWRao6a3yNC4d6UjMdQxonkCMvKBwdpubU3qhfcdf1"}
	names["2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk"] = tokenMeta{name: "Wrapped Ethereum (Sollet)", symbol: "soETH", decimals: 6, mint: "2FPyTwcZLUg1MDrwsyoP4D6s1tM7hAkHYRjkNb5w6Pxk"}
	names["BQcdHdAQW1hczDbBi9hiegXAR7A98Q9jx3X3iBBBDiq4"] = tokenMeta{name: "Wrapped USDT (Sollet)", symbol: "soUSDT", decimals: 6, mint: "BQcdHdAQW1hczDbBi9hiegXAR7A98Q9jx3X3iBBBDiq4"}
	names["SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt"] = tokenMeta{name: "Serum", symbol: "SRM", decimals: 6, mint: "SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt"}

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
