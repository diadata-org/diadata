package scrapers

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	orcaWhirlpoolIdlBind "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/orca/whirlpool"
	"github.com/diadata-org/diadata/pkg/utils"

	bin "github.com/gagliardetto/binary"
	tokenmetadata "github.com/gagliardetto/metaplex-go/clients/token-metadata"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/programs/tokenregistry"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"github.com/gorilla/websocket"
)

const (
	orcaSolanaHttpEndpoint           = "https://rpc.ankr.com/solana"
	orcaSolanaWsEndpoint             = rpc.MainNetBeta_WS
	OrcaProgWhirlpoolConfigAddr      = "2LecshUwdy9xi7meFgHtFJQNSKk4KdTrcpvaB56dP2NQ"
	OrcaProgWhirlpoolAddr            = "whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc"
	OrcaProgWhirlpoolAccountDataSize = 653
	OrcaMaxRetries                   = 5
	OrcaRetryDelay                   = 3 * time.Second
)

// The scraper object for Orca
type OrcaScraper struct {
	exchangeName string

	// state variables to signal events
	run          bool
	shutdown     chan nothing
	shutdownDone chan nothing

	errorLock sync.RWMutex
	error     error

	pairScrapers map[string]*OrcaPairScraper
	chanTrades   chan *dia.Trade

	RestClient *rpc.Client
	WsClient   *ws.Client
}

// Returns a new exchange scraper
func NewOrcaScraper(exchange dia.Exchange, scrape bool) *OrcaScraper {

	log.Infof("init rest and ws client for %s", exchange.BlockChain.Name)
	restClient := rpc.New(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", orcaSolanaHttpEndpoint))
	wsClient, err := ws.Connect(context.Background(), utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_WS", orcaSolanaWsEndpoint))
	if err != nil {
		log.Fatal("init ws client: ", err)
	}

	scraper := &OrcaScraper{
		exchangeName: exchange.Name,
		RestClient:   restClient,
		WsClient:     wsClient,
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*OrcaPairScraper),
		chanTrades:   make(chan *dia.Trade),
	}

	_, err = scraper.loadMarketsMetadata()
	if err != nil {
		log.Fatal("load metadata: %s", err)
	}

	if scrape {
		go scraper.mainLoop()
	}
	return scraper
}

// Closes any existing API connections, as well as channels of
// pairScrapers from calls to ScrapePair
func (s *OrcaScraper) Close() error {
	s.run = false
	for _, pairScraper := range s.pairScrapers {
		pairScraper.closed = true
	}
	s.WsClient.Close()
	s.RestClient.Close()

	close(s.shutdown)
	<-s.shutdownDone
	return nil
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from the scraper
func (s *OrcaScraper) ScrapePair(pair dia.ExchangePair) (ps PairScraper, err error) {
	return
}

// Returns the list of all available trade pairs in order to pairDiscoveryService service work
func (s *OrcaScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	pairs, err = s.loadMarketsMetadata()
	return
}

// Channel returns a channel that can be used to receive trades
func (s *OrcaScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

// FillSymbolData adds the name to the asset underlying @symbol on scraper
func (s *OrcaScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

// NormalizePair accounts for the pair
func (s *OrcaScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

type OrcaPairScraper struct {
	parent *OrcaScraper
	pair   dia.ExchangePair
	closed bool
}

// Close stops listening for trades of the pair associated with the scraper
func (ps *OrcaPairScraper) Close() error {
	ps.parent.errorLock.RLock()
	defer ps.parent.errorLock.RUnlock()
	ps.closed = true
	return nil
}

// Error returns an error when the channel Channel() is closed and nil otherwise
func (ps *OrcaPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *OrcaPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}

func (s *OrcaScraper) mainLoop() {
	wg := sync.WaitGroup{}

	s.run = true
	for s.run {

		ic := make(chan map[string]nothing)

		wg.Add(2)

		go func() {
			sub, err := s.WsClient.LogsSubscribeMentions(
				solana.MustPublicKeyFromBase58(OrcaProgWhirlpoolAddr),
				rpc.CommitmentFinalized,
			)
			if err != nil {
				log.Fatal("sub: %s", err)
			}
			defer sub.Unsubscribe()
			defer s.WsClient.Close()
			log.Infof("start subscription routine")
			defer log.Warnf("subscription routine end\n")
			lastSlot := uint64(0)
			pendingTxs := make(map[string]nothing)
			for {
				got, err := sub.Recv()
				if err != nil {
					if !websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
						log.Warnf("recv expected error: %s", err)
					} else {
						log.Warnf("recv error: %s", err)
					}
					retries := 0
					for retries <= OrcaMaxRetries {
						retries++
						log.Warnf("Recovering websocket connection %d/%d ...", retries, OrcaMaxRetries)
						s.WsClient, err = ws.Connect(context.Background(), orcaSolanaWsEndpoint)
						time.Sleep(OrcaRetryDelay)
						if err != nil {
							log.Warnf("retry failed: ", err)
						} else {
							sub, err = s.WsClient.LogsSubscribeMentions(
								solana.MustPublicKeyFromBase58(OrcaProgWhirlpoolAddr),
								rpc.CommitmentFinalized,
							)
							if err != nil {
								log.Warnf("re-sub failed: ", err)
							} else {
								log.Infof("Websocket connection recovered with success.")
								retries = 0
								break
							}
						}
					}
				} else {
					if got.Value.Err == nil {
						if got.Context.Slot > lastSlot {
							ic <- pendingTxs
							pendingTxs = make(map[string]nothing)
							lastSlot = got.Context.Slot
							pendingTxs[got.Value.Signature.String()] = struct{}{}
						} else if got.Context.Slot == lastSlot {
							pendingTxs[got.Value.Signature.String()] = struct{}{}
						} else {
							log.Fatalf("invalid order for subscription: %s (curr %d, last %d)\n", got.Value.Signature.String(), got.Context.Slot, lastSlot)
						}

					}
				}
			}
		}()

		go func() {
			log.Infof("start processing routine")
			defer log.Warnf("processing routine end\n")
			for v := range ic {
				for len(v) > 0 {
					for k := range v {

						max := uint64(0)
						resp, err := s.RestClient.GetTransaction(
							context.TODO(),
							solana.MustSignatureFromBase58(k),
							&rpc.GetTransactionOpts{
								MaxSupportedTransactionVersion: &max,
								Commitment:                     rpc.CommitmentFinalized,
								Encoding:                       solana.EncodingBase64,
							},
						)
						if err != nil {
							continue
						}
						if resp != nil {
							tx := resp.Transaction
							if tx != nil {
								txParsed, err := tx.GetTransaction()
								if err == nil && txParsed != nil {
									if txParsed.Message.GetVersion() == 0 && !txParsed.Message.IsVersioned() {
										accMetaList, err := txParsed.AccountMetaList()
										if err == nil && accMetaList != nil {
											for i, inst := range txParsed.Message.Instructions {
												if txParsed.Message.AccountKeys[inst.ProgramIDIndex].String() == OrcaProgWhirlpoolAddr {
													instDecoded, err := orcaWhirlpoolIdlBind.DecodeInstruction(accMetaList, inst.Data)
													if err != nil {
														log.Warnf("  %88s cannot-decode\n", k)
													}
													if instDecoded.TypeID == orcaWhirlpoolIdlBind.Instruction_Swap {
														pair := s.pairScrapers[txParsed.Message.AccountKeys[inst.Accounts[2]].String()].pair.UnderlyingPair
														baseToken := pair.BaseToken
														quoteToken := pair.QuoteToken
														for _, inner := range resp.Meta.InnerInstructions {
															if inner.Index == uint16(i) {
																var amountA, amountB, volume, price float64
																for j, innerInsts := range inner.Instructions {
																	innerInstData := bin.NewBorshDecoder(innerInsts.Data)
																	typeIdEnconding, err := innerInstData.ReadUint8()
																	if err != nil {
																		log.Fatalf("cannot read inner inst param: %s", err)
																	}
																	innerInstParam, err := innerInstData.ReadUint64(bin.LE)
																	if err != nil {
																		log.Fatalf("cannot read inner inst param: %s", err)
																	}
																	if bin.TypeIDEncoding(uint32(typeIdEnconding)) == bin.AnchorTypeIDEncoding {
																		if j == 0 {
																			if *instDecoded.Impl.(*orcaWhirlpoolIdlBind.Swap).AToB {
																				amountA = FormatUint64Decimals(innerInstParam, int(baseToken.Decimals))
																			} else {
																				amountB = FormatUint64Decimals(innerInstParam, int(quoteToken.Decimals))
																			}
																		} else {
																			if *instDecoded.Impl.(*orcaWhirlpoolIdlBind.Swap).AToB {
																				amountB = FormatUint64Decimals(innerInstParam, int(quoteToken.Decimals))
																			} else {
																				amountA = FormatUint64Decimals(innerInstParam, int(baseToken.Decimals))
																			}
																		}
																	}
																}
																if *instDecoded.Impl.(*orcaWhirlpoolIdlBind.Swap).AToB {
																	price = amountA / amountB
																	volume = -amountB
																} else {
																	price = amountA / amountB
																	volume = amountB
																}
																trade := &dia.Trade{
																	Symbol:         quoteToken.Symbol,
																	Pair:           quoteToken.Symbol + "-" + baseToken.Symbol,
																	BaseToken:      baseToken,
																	QuoteToken:     quoteToken,
																	Price:          price,
																	Volume:         volume,
																	Time:           resp.BlockTime.Time(),
																	ForeignTradeID: k,
																	Source:         s.exchangeName,
																	VerifiedPair:   true,
																}
																log.Infof("pair -- price -- volume: %s -- %v -- %v", trade.Pair, trade.Price, trade.Volume)
																log.Info("tx hash: ", k)
																s.chanTrades <- trade
															}
														}
													}
												}
											}
										}
									}
								}
							}
							delete(v, k)
						}
					}
				}
			}
		}()

		wg.Wait()

	}
}

// Load markets and tokens metadata
func (s *OrcaScraper) loadMarketsMetadata() (pairs []dia.ExchangePair, err error) {
	log.Infof("loading initial data from pools ...")
	start := time.Now()
	poolPairs, err := s.loadMarketPools()
	if err != nil {
		return
	}
	pairs = append(pairs, poolPairs...)
	log.Infof("loaded %d pairs from legacy pools in %.1fs", len(poolPairs), time.Since(start).Seconds())
	start = time.Now()
	whirlpoolPairs, err := s.loadMarketWhirlpools()
	if err != nil {
		return
	}
	pairs = append(pairs, whirlpoolPairs...)
	log.Infof("loaded %d pairs from whirlpools in %.1fs", len(whirlpoolPairs), time.Since(start).Seconds())
	return
}

// Get Orca market legacy pools
func (s *OrcaScraper) loadMarketPools() (pairs []dia.ExchangePair, err error) {
	return
}

// Get Orca market whirlpools
func (s *OrcaScraper) loadMarketWhirlpools() (pairs []dia.ExchangePair, err error) {
	hardcodedTokenMeta := GetOrcaTokensMetadata()
	resp, err := s.RestClient.GetProgramAccountsWithOpts(
		context.TODO(),
		solana.MustPublicKeyFromBase58(OrcaProgWhirlpoolAddr),
		&rpc.GetProgramAccountsOpts{
			Filters: []rpc.RPCFilter{
				{
					DataSize: OrcaProgWhirlpoolAccountDataSize,
				},
			},
		},
	)
	if err != nil {
		return
	}
	if resp == nil {
		return nil, fmt.Errorf("program account not found")
	}
	log.Infof("discovered %d accounts in whirlpool program, retrieving metadata ...", len(resp))
	for _, progAcc := range resp {
		acct := progAcc.Account
		pubKey := progAcc.Pubkey.String()
		if acct.Owner.String() == OrcaProgWhirlpoolAddr {
			d := bin.NewBorshDecoder(acct.Data.GetBinary())
			var w orcaWhirlpoolIdlBind.Whirlpool
			d.Decode(&w)
			// Blacklist XXX/USDC, ATLAS/USDC, SHIB/USDC
			if pubKey == "FfBeru58Q7hjqHq9T2Trw1BeyjE1YwHsx9MivKUwoTLQ" || pubKey == "9vqFu6v9CcVDaSx2oRD3jo8H5gqkE2urYQgpT16V1BTa" || pubKey == "DahhciLA89UkZoqrqVWL2nojwPLmSVkXQGTiEhAtkaFa" {
				continue
			}
			if w.WhirlpoolsConfig.String() == OrcaProgWhirlpoolConfigAddr {
				var tokenA, tokenB dia.Asset

				// Get token A mint data and metadata
				if mintData, err := s.getTokenMintData(w.TokenMintA.String()); err == nil {
					if mintData.IsInitialized {
						tokenA.Decimals = mintData.Decimals
					}
				} else {
					return nil, err
				}
				if metadata, err := s.getTokenMetadata(w.TokenMintA.String()); err != nil {
					if v, ok := hardcodedTokenMeta[w.TokenMintA.String()]; ok {
						tokenA.Symbol = v.(OrcaTokenMetadata).GetSymbol()
						tokenA.Name = v.(OrcaTokenMetadata).GetName()
					} else {
						log.Warnf("token metadata not found for %s: %s", w.TokenMintA.String(), err)
						continue
					}
				} else {
					tokenA.Symbol = strings.TrimRight(metadata.Data.Symbol, "\x00")
					tokenA.Name = metadata.Data.Name
				}
				tokenA.Address = w.TokenMintA.String()
				tokenA.Blockchain = "Solana"

				// Get token B mint data and metadata
				if mintData, err := s.getTokenMintData(w.TokenMintB.String()); err == nil {
					tokenB.Decimals = mintData.Decimals
				} else {
					return nil, err
				}
				if metadata, err := s.getTokenMetadata(w.TokenMintB.String()); err != nil {
					if v, ok := hardcodedTokenMeta[w.TokenMintB.String()]; ok {
						tokenB.Symbol = v.(OrcaTokenMetadata).GetSymbol()
						tokenB.Name = v.(OrcaTokenMetadata).GetName()
					} else {
						log.Warnf("token metadata not found for %s: %s", w.TokenMintB.String(), err)
						continue
					}
				} else {
					tokenB.Symbol = strings.TrimRight(metadata.Data.Symbol, "\x00")
					tokenB.Name = metadata.Data.Name
				}
				tokenB.Address = w.TokenMintB.String()
				tokenB.Blockchain = "Solana"

				log.Infof("whirlpool loaded: %44s (%10s / %10s)\n", pubKey, tokenA.Symbol, tokenB.Symbol)
				pair := dia.ExchangePair{
					Symbol:      tokenA.Symbol,
					Verified:    true,
					ForeignName: pubKey,
					Exchange:    s.exchangeName,
					UnderlyingPair: dia.Pair{
						BaseToken:  tokenA,
						QuoteToken: tokenB,
					},
				}
				pairs = append(pairs, pair)
				s.pairScrapers[pubKey] = &OrcaPairScraper{
					parent: s,
					pair:   pair,
					closed: false,
				}
			}
		}
	}
	return
}

// Get Solana token mint data
func (s *OrcaScraper) getTokenMintData(account string) (mint token.Mint, err error) {
	resp, err := s.RestClient.GetAccountInfoWithOpts(
		context.TODO(),
		solana.MustPublicKeyFromBase58(account),
		&rpc.GetAccountInfoOpts{},
	)
	if err != nil {
		return
	}
	d := bin.NewBorshDecoder(resp.Value.Data.GetBinary())
	err = d.Decode(&mint)
	if err != nil {
		return
	}
	return
}

// Get Solana token metadata
func (s *OrcaScraper) getTokenMetadata(account string) (metadata tokenmetadata.Metadata, err error) {
	accMint := solana.MustPublicKeyFromBase58(account)
	tMeta, err := tokenregistry.GetTokenRegistryEntry(context.TODO(), s.RestClient, accMint)
	if err != nil {
		metaAddress, _, err := solana.FindTokenMetadataAddress(accMint)
		if err != nil {
			return metadata, err
		}
		resp, err := s.RestClient.GetAccountInfo(
			context.TODO(),
			metaAddress,
		)
		if err != nil {
			return metadata, err
		}
		d := bin.NewBorshDecoder(resp.Value.Data.GetBinary())
		err = d.Decode(&metadata)
		if err != nil {
			return metadata, err
		}
		return metadata, nil
	}
	return tokenmetadata.Metadata{Data: tokenmetadata.Data{Symbol: tMeta.Symbol.String()}}, nil
}

type OrcaTokenMetadata interface {
	GetName() string
	GetSymbol() string
}

func (t *orcaTokenMetadata) GetName() string {
	return t.Name
}

func (t *orcaTokenMetadata) GetSymbol() string {
	return t.Symbol
}

type orcaTokenMetadata struct {
	Name   string
	Symbol string
}

func GetOrcaTokensMetadata() map[string]interface{} {
	tokenMetadata := make(map[string]interface{})
	tokenMetadata["zebeczgi5fSEtbpfQKVZKCJ3WgYXxjkMUkNNx7fLKAF"] = &orcaTokenMetadata{Name: "ZEBEC", Symbol: "ZBC"}
	tokenMetadata["GEJpt3Wjmr628FqXxTgxMce1pLntcPV4uFi8ksxMyPQh"] = &orcaTokenMetadata{Name: "daoSOL Token", Symbol: "daoSOL"}
	tokenMetadata["CT1iZ7MJzm8Riy6MTgVht2PowGetEWrnq1SfmUjKvz8c"] = &orcaTokenMetadata{Name: "Balloonsville Solvent Droplet", Symbol: "svtBV"}
	tokenMetadata["7Q2afV64in6N6SeZsAAB81TJzwDoD6zpqmHkzi9Dcavn"] = &orcaTokenMetadata{Name: "JPOOL Solana Token", Symbol: "JSOL"}
	tokenMetadata["USDH1SM1ojwWUga67PGrgFWUHibbjqMvuMaDkRJTgkX"] = &orcaTokenMetadata{Name: "USDH Hubble Stablecoin", Symbol: "USDH"}
	tokenMetadata["6naWDMGNWwqffJnnXFLBCLaYu1y5U9Rohe5wwJPHvf1p"] = &orcaTokenMetadata{Name: "SCRAP", Symbol: "SCRAP"}
	tokenMetadata["SLNDpmoWTVADgEdndyvWzroNL7zSi1dF9PC3xHGtPwp"] = &orcaTokenMetadata{Name: "Solend", Symbol: "SLND"}
	tokenMetadata["EiasWmzy9MrkyekABHLfFRkGhRakaWNvmQ8h5DV86zyn"] = &orcaTokenMetadata{Name: "Visionary Studios Solvent Droplet", Symbol: "svtVSNRY"}
	tokenMetadata["DUSTawucrTsGU8hcqRdHDCbuYhCPADMLM2VcCb8VnFnQ"] = &orcaTokenMetadata{Name: "DUST Protocol", Symbol: "DUST"}
	tokenMetadata["BoeDfSFRyaeuaLP97dhxkHnsn7hhhes3w3X8GgQj5obK"] = &orcaTokenMetadata{Name: "Famous Fox Federation Solvent Droplet", Symbol: "svtFFF"}
	tokenMetadata["ANAxByE6G2WjFp7A4NqtWYXb3mgruyzZYg3spfxe6Lbo"] = &orcaTokenMetadata{Name: "ANA", Symbol: "ANA"}
	tokenMetadata["9iLH8T7zoWhY7sBmj1WK9ENbWdS1nL8n9wAxaeRitTa6"] = &orcaTokenMetadata{Name: "Hedge USD", Symbol: "USH"}
	tokenMetadata["HBB111SCo9jkCejsZfz8Ec8nH7T6THF8KEKSnvwT6XK6"] = &orcaTokenMetadata{Name: "Hubble Protocol Token", Symbol: "HBB"}
	tokenMetadata["52GzcLDMfBveMRnWXKX7U3Pa5Lf7QLkWWvsJRDjWDBSk"] = &orcaTokenMetadata{Name: "NGN Coin", Symbol: "NGNC"}
	tokenMetadata["5PmpMzWjraf3kSsGEKtqdUsCoLhptg4yriZ17LKKdBBy"] = &orcaTokenMetadata{Name: "Hedge Token", Symbol: "HDG"}
	tokenMetadata["9tzZzEHsKnwFL1A3DyFJwj36KnZj3gZ7g4srWp9YTEoh"] = &orcaTokenMetadata{Name: "ARB Protocol", Symbol: "ARB"}
	tokenMetadata["AG5j4hhrd1ReYi7d1JsZL8ZpcoHdjXvc8sdpWF74RaQh"] = &orcaTokenMetadata{Name: "Okay Bears Solvent Droplet", Symbol: "svtOKAY"}
	tokenMetadata["7kbnvuGBxxj8AG9qp8Scn56muWGaRaFqxg1FsRp3PaFT"] = &orcaTokenMetadata{Name: "UXD Stablecoin", Symbol: "UXD"}
	tokenMetadata["SHDWyBxihqiCj6YekG2GUr7wqKLeLAMK1gHZck9pL6y"] = &orcaTokenMetadata{Name: "Shadow Token", Symbol: "SHDW"}
	tokenMetadata["GENEtH5amGSi8kHAtQoezp1XEXwZJ8vcuePYnXdKrMYz"] = &orcaTokenMetadata{Name: "Genopets", Symbol: "GENE"}
	tokenMetadata["EsPKhGTMf3bGoy4Qm7pCv3UCcWqAmbC1UGHBTDxRjjD4"] = &orcaTokenMetadata{Name: "FTM (Allbridge from Fantom)", Symbol: "FTM"}
	tokenMetadata["bSo13r4TkiE4KumL71LsHTPpL2euBYLFx6h9HP3piy1"] = &orcaTokenMetadata{Name: "BlazeStake Staked SOL (bSOL)", Symbol: "bSOL"}
	tokenMetadata["SuperbZyz7TsSdSoFAZ6RYHfAWe9NmjXBLVQpS8hqdx"] = &orcaTokenMetadata{Name: "SuperBonds Token", Symbol: "SB"}
	tokenMetadata["8W2ZFYag9zTdnVpiyR4sqDXszQfx2jAZoMcvPtCSQc7D"] = &orcaTokenMetadata{Name: "The Catalina Whale Mixer Solvent Droplet", Symbol: "svtCWM"}
	tokenMetadata["PsyFiqqjiv41G7o5SMRzDJCu4psptThNR2GtfeGHfSq"] = &orcaTokenMetadata{Name: "PsyOptions", Symbol: "PSY"}
	tokenMetadata["GePFQaZKHcWE5vpxHfviQtH5jgxokSs51Y5Q4zgBiMDs"] = &orcaTokenMetadata{Name: "Jungle DeFi", Symbol: "JFI"}
	tokenMetadata["METAmTMXwdb8gYzyCPfXXFmZZw4rUsXX58PNsDg7zjL"] = &orcaTokenMetadata{Name: "Solice", Symbol: "SLC"}
	tokenMetadata["USDrbBQwQbQ2oWHUPfA8QBHcyVxKUq1xHyXsSLKdUq2"] = &orcaTokenMetadata{Name: "Ratio stable Token", Symbol: "USDr"}
	tokenMetadata["4MSMKZwGnkT8qxK8LsdH28Uu8UfKRT2aNaGTU8TEMuHz"] = &orcaTokenMetadata{Name: "Genopets Genesis - Solvent Droplet", Symbol: "svtGENE"}
	tokenMetadata["F3nefJBcejYbtdREjui1T9DPh5dBgpkKq7u2GAAMXs5B"] = &orcaTokenMetadata{Name: "ALL ART", Symbol: "AART"}
	tokenMetadata["BKipkearSqAUdNKa1WDstvcMjoPsSKBuNyvKDQDDu9WE"] = &orcaTokenMetadata{Name: "Hawksight", Symbol: "HAWK"}
	tokenMetadata["CowKesoLUaHSbAMaUxJUj7eodHHsaLsS65cy8NFyRDGP"] = &orcaTokenMetadata{Name: "Cash Cow", Symbol: "COW"}
	tokenMetadata["Ez2zVjw85tZan1ycnJ5PywNNxR6Gm4jbXQtZKyQNu3Lv"] = &orcaTokenMetadata{Name: "Fluid USDC", Symbol: "fUSDC"}
	tokenMetadata["AFbX8oGjGpmVFywbVouvhQSRmiW2aR1mohfahi4Y2AdB"] = &orcaTokenMetadata{Name: "GST", Symbol: "GST"}
	tokenMetadata["svtMpL5eQzdmB3uqK9NXaQkq8prGZoKQFNVJghdWCkV"] = &orcaTokenMetadata{Name: "Solvent", Symbol: "SVT"}
	tokenMetadata["6F5A4ZAtQfhvi3ZxNex9E1UN5TK7VM2enDCYG1sx1AXT"] = &orcaTokenMetadata{Name: "Degenerate Ape Academy Solvent Droplet", Symbol: "svtDAPE"}
	tokenMetadata["UXPhBoR3qG4UCiGNJfV7MqhHyFqKN68g45GoYvAeL2M"] = &orcaTokenMetadata{Name: "UXP Governance Token", Symbol: "UXP"}
	tokenMetadata["FANTafPFBAt93BNJVpdu25pGPmca3RfwdsDsRrT3LX1r"] = &orcaTokenMetadata{Name: "Phantasia", Symbol: "FANT"}
	tokenMetadata["Bp6k6xacSc4KJ5Bmk9D5xfbw8nN42ZHtPAswEPkNze6U"] = &orcaTokenMetadata{Name: "Pesky Penguins Solvent Droplet", Symbol: "svtPSK"}
	tokenMetadata["2zzC22UBgJGCYPdFyo7GDwz7YHq5SozJc1nnBqLU8oZb"] = &orcaTokenMetadata{Name: "1SPACE", Symbol: "1SP"}
	tokenMetadata["EmLJ8cNEsUtboiV2eiD6VgaEscSJ6zu3ELhqixUP4J56"] = &orcaTokenMetadata{Name: "Thugbirdz - Solvent Droplet", Symbol: "svtTHUGZ"}
	tokenMetadata["9acdc5M9F9WVM4nVZ2gPtVvkeYiWenmzLW9EsTkKdsUJ"] = &orcaTokenMetadata{Name: "Gooney Toons Solvent Droplet", Symbol: "svtGOON"}
	tokenMetadata["GNCjk3FmPPgZTkbQRSxr6nCvLtYMbXKMnRxg8BgJs62e"] = &orcaTokenMetadata{Name: "CELO (Allbridge from Celo)", Symbol: "CELO"}
	tokenMetadata["DXA9itWDGmGgqqUoHnBhw6CjvJKMUmTMKB17hBuoYkfQ"] = &orcaTokenMetadata{Name: "Honey Genesis Bee Solvent Droplet", Symbol: "svtHNYG"}
	tokenMetadata["HYtdDGdMFqBrtyUe5z74bKCtH2WUHZiWRicjNVaHSfkg"] = &orcaTokenMetadata{Name: "Aurory - Solvent Droplet", Symbol: "svtAURY"}
	tokenMetadata["G9tt98aYSznRk7jWsfuz9FnTdokxS6Brohdo9hSmjTRB"] = &orcaTokenMetadata{Name: "PUFF", Symbol: "PUFF"}
	tokenMetadata["8vkTew1mT8w5NapTqpAoNUNHW2MSnAGVNeu8QPmumSJM"] = &orcaTokenMetadata{Name: "Playground Waves Solvent Droplet", Symbol: "svtWAVE"}
	tokenMetadata["PRSMNsEPqhGVCH1TtWiJqPjJyh2cKrLostPZTNy1o5x"] = &orcaTokenMetadata{Name: "PRISM", Symbol: "PRISM"}
	tokenMetadata["seedEDBqu63tJ7PFqvcbwvThrYUkQeqT6NLf81kLibs"] = &orcaTokenMetadata{Name: "Seeded Network", Symbol: "SEEDED"}
	tokenMetadata["FoXyMu5xwXre7zEoSvzViRk3nGawHUp9kUh97y2NDhcq"] = &orcaTokenMetadata{Name: "Famous Fox Federation", Symbol: "FOXY"}
	tokenMetadata["BDrL8huis6S5tpmozaAaT5zhE5A7ZBAB2jMMvpKEeF8A"] = &orcaTokenMetadata{Name: "NOVA FINANCE", Symbol: "NOVA"}
	tokenMetadata["3GQqCi9cuGhAH4VwkmWD32gFHHJhxujurzkRCQsjxLCT"] = &orcaTokenMetadata{Name: "Galactic Geckos Space Garage Solvent Droplet", Symbol: "svtGGSG"}
	tokenMetadata["DCgRa2RR7fCsD63M3NgHnoQedMtwH1jJCwZYXQqk9x3v"] = &orcaTokenMetadata{Name: "DeGods Solvent Droplet", Symbol: "svtDGOD"}
	tokenMetadata["F8Wh3zT1ydxPYfQ3p1oo9SCJbjedqDsaC1WaBwh64NHA"] = &orcaTokenMetadata{Name: "Serum Surfers Droplet", Symbol: "SSURF"}
	tokenMetadata["Fm9rHUTF5v3hwMLbStjZXqNBBoZyGriQaFM6sTFz3K8A"] = &orcaTokenMetadata{Name: "MonkeyBucks", Symbol: "MBS"}
	tokenMetadata["4h41QKUkQPd2pCAFXNNgZUyGUxQ6E7fMexaZZHziCvhh"] = &orcaTokenMetadata{Name: "The Suites Token", Symbol: "SUITE"}
	tokenMetadata["7i5KKsX2weiTkry7jA4ZwSuXGhs5eJBEjY8vVxR4pfRx"] = &orcaTokenMetadata{Name: "GMT", Symbol: "GMT"}
	tokenMetadata["ratioMVg27rSZbSvBopUvsdrGUzeALUfFma61mpxc8J"] = &orcaTokenMetadata{Name: "Ratio Governance Token", Symbol: "RATIO"}
	tokenMetadata["3b9wtU4VP6qSUDL6NidwXxK6pMvYLFUTBR1QHWCtYKTS"] = &orcaTokenMetadata{Name: "Playground Epochs Solvent Droplet", Symbol: "svtEPOCH"}
	tokenMetadata["FoRGERiW7odcCBGU1bztZi16osPBHjxharvDathL5eds"] = &orcaTokenMetadata{Name: "FORGE", Symbol: "FORGE"}
	tokenMetadata["4wGimtLPQhbRT1cmKFJ7P7jDTgBqDnRBWsFXEhLoUep2"] = &orcaTokenMetadata{Name: "Lifinity Flares Solvent Droplet", Symbol: "svtFLARE"}
	tokenMetadata["SNSNkV9zfG5ZKWQs6x4hxvBRV6s8SqMfSGCtECDvdMd"] = &orcaTokenMetadata{Name: "SynesisOne", Symbol: "SNS"}
	tokenMetadata["9WMwGcY6TcbSfy9XPpQymY3qNEsvEaYL3wivdwPG2fpp"] = &orcaTokenMetadata{Name: "Jelly", Symbol: "JELLY"}
	tokenMetadata["Ca5eaXbfQQ6gjZ5zPVfybtDpqWndNdACtKVtxxNHsgcz"] = &orcaTokenMetadata{Name: "Solana Monkey Business Solvent Droplet", Symbol: "svtSMB"}
	tokenMetadata["5Wsd311hY8NXQhkt9cWHwTnqafk7BGEbLu8Py3DSnPAr"] = &orcaTokenMetadata{Name: "Compendium Finance", Symbol: "CMFI"}
	tokenMetadata["GWsZd8k85q2ie9SNycVSLeKkX7HLZfSsgx6Jdat9cjY1"] = &orcaTokenMetadata{Name: "Pollen Coin", Symbol: "PCN"}
	return tokenMetadata
}

// Format a uint64 to a float64 with the given number of decimals
func FormatUint64Decimals(value uint64, decimals int) (valueFormatted float64) {
	balance, _ := new(big.Float).Quo(big.NewFloat(0).SetUint64(value), big.NewFloat(math.Pow10(decimals))).Float64()
	return balance
}
