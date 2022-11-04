package nfttradescrapers

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/jackc/pgx/v4"
	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/shopspring/decimal"
	"github.com/status-im/keycard-go/hexutils"
)

const (
	rpcEndpointSolana         = "https://solana-api.projectserum.com"
	MagicEdenV2ProgramAddress = "M2mx93ekt1fmXSVkTrUL9xVFHkmME8HTUi5Cyc5aF7K"
	SolTokenAddress           = "So11111111111111111111111111111111111111112"
	SaleTxPrefix              = "f223c68952e1f2b6"
	MetadataFee               = 0.00561672
	MagicEden                 = "MagicEden"
)

var (
	defMagicEdenConf = &MagicEdenScraperConfig{
		SolanaRestUri:    SaleTxPrefix,
		ProgramAddr:      MagicEdenV2ProgramAddress,
		BatchSize:        1000,
		WaitPeriod:       30 * time.Second,
		MaxRetry:         2,
		SkipOnErr:        true,
		ScrapeHistorical: false,
	}
	defMagicEdenState   = &MagicEdenScraperState{}
	assetCacheMagicEden = make(map[string]dia.Asset)
)

type SolanaNFTMetadata struct {
	sourceAccount     common.PublicKey
	mintAccount       common.PublicKey
	collectionAccount common.PublicKey
	name              string
	symbol            string
	uri               string
	fee               int
	creators          []NFTCreator
	primarySaleDone   int
	isMutable         int
	creationTime      int64
}

type NFTCreator struct {
	account  common.PublicKey
	verified int
	share    int
}

type MagicEdenScraper struct {
	solanaRpcClient *client.Client
	tradeScraper    TradeScraper
	mu              sync.Mutex
	conf            *MagicEdenScraperConfig
	state           *MagicEdenScraperState
}

type MagicEdenScraperConfig struct {
	SolanaRestUri string `json:"rest_uri"`

	ProgramAddr string `json:"program_addr"`

	BatchSize int `json:"batch_size"`

	WaitPeriod time.Duration `json:"wait_per_batch"`

	MaxRetry int `json:"max_retry"`

	SkipOnErr bool `json:"skip_on_error"`

	ScrapeHistorical bool `json:"scrape_historical"`
}

type MagicEdenScraperState struct {
	LastTx string `json:"last_tx"`

	LastTxHistorical string `json:"last_tx_historical"`

	LastErr string `json:"last_error"`

	LastErrHistorical string `json:"last_error_historical"`

	ErrCounter int `json:"count_of_error"`

	ErrCounterHistorical int `json:"count_of_error_historical"`
}

func (s *MagicEdenScraper) loadConfig(ctx context.Context) error {
	return s.tradeScraper.datastore.GetScraperConfig(ctx, MagicEden, s.conf)
}

func (s *MagicEdenScraper) loadState(ctx context.Context) error {
	return s.tradeScraper.datastore.GetScraperState(ctx, MagicEden, s.state)
}

func (s *MagicEdenScraper) storeState(ctx context.Context) error {
	return s.tradeScraper.datastore.SetScraperState(ctx, MagicEden, s.state)
}

func NewMagicEdenScraper(rdb *models.RelDB, exchange dia.NFTExchange) *MagicEdenScraper {
	ctx := context.Background()
	scraper := &MagicEdenScraper{
		solanaRpcClient: client.NewClient(utils.Getenv("SOLANA_URI_REST", rpcEndpointSolana)),
		conf:            &MagicEdenScraperConfig{},
		state:           &MagicEdenScraperState{},
		tradeScraper: TradeScraper{
			shutdown:     make(chan nothing),
			shutdownDone: make(chan nothing),
			datastore:    rdb,
			chanTrade:    make(chan dia.NFTTrade),
			source:       MagicEden,
		},
	}
	if err := scraper.initScraper(ctx); err != nil {
		log.Errorf("magiceden scraper could not be initialized: %s", err.Error())
		return nil
	}
	log.Infof("magiceden scraper starts from tx: %s", scraper.state.LastTx)
	go scraper.mainLoop()
	return scraper
}

func (s *MagicEdenScraper) initScraper(ctx context.Context) error {
	if err := s.loadConfig(ctx); err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			log.Errorf("unable to read scraper config from rdb: %s", err.Error())
			return err
		}

		defConf := *defMagicEdenConf
		s.conf = &defConf
		if err := s.tradeScraper.datastore.SetScraperConfig(ctx, "MagicEden", s.conf); err != nil {
			log.Errorf("unable to store scraper config on rdb: %s", err.Error())
			return err
		}

		defState := *defMagicEdenState
		s.state = &defState
		if err := s.tradeScraper.datastore.SetScraperState(ctx, OpenSea, s.state); err != nil {
			log.Errorf("unable to store scraper state on rdb: %s", err.Error())
			return err
		}

		return nil
	}

	return s.loadState(ctx)
}

func (s *MagicEdenScraper) mainLoop() {
	defer func() {
		s.tradeScraper.closed = true

		close(s.tradeScraper.chanTrade)
		close(s.tradeScraper.shutdownDone)
	}()

	log.Infof("magiceden scraper has been started (batch: %d, period: %s)", s.conf.BatchSize, s.conf.WaitPeriod.String())

	for stop := false; !stop; {
		if err := s.FetchTrades(); err != nil {
			if errors.Is(err, errOpenSeaShutdownRequest) {
				stop = true
				continue
			}
		}

		if s.conf.ScrapeHistorical {
			if err := s.FetchHistoricalTrades(); err != nil {
				if errors.Is(err, errOpenSeaShutdownRequest) {
					stop = true
					continue
				}
			}
		}

		log.Debugf("wait for %s", 1*time.Second)

		select {
		case <-time.After(2 * time.Second):
		case <-s.tradeScraper.shutdown:
			stop = true
		}
	}

}

func (s *MagicEdenScraper) FetchTrades() error {
	log.Info("start fetch trades...")
	var err error
	ctx := context.Background()

	// it must be run once at a time
	s.mu.Lock()
	defer s.mu.Unlock()

	// read config
	if err = s.loadConfig(ctx); err != nil {
		log.Warnf("unable to load scraper config: %s", err.Error())
		return err
	}

	// read state
	if err = s.loadState(ctx); err != nil {
		log.Warnf("unable to load scraper state: %s", err.Error())
		return err
	}

	txToProcess := make([]rpc.SignatureWithStatus, 0)
	lastFetchedTx := ""
	for {
		txList, err := s.solanaRpcClient.GetSignaturesForAddressWithConfig(ctx, MagicEdenV2ProgramAddress,
			rpc.GetSignaturesForAddressConfig{
				Before: lastFetchedTx,
				Until:  s.state.LastTx,
				Limit:  s.conf.BatchSize,
			})
		if err != nil {
			log.Warnf("unable to retrieve confirmed transaction signatures for account: %s", err.Error())
			return err
		}

		for _, tx := range txList {
			txToProcess = append(txToProcess, tx)
			if tx.Signature != "" {
				lastFetchedTx = tx.Signature
			}
		}

		if len(txList) < s.conf.BatchSize || s.state.LastTx == "" {
			break
		}

	}

	log.Infof("processing magiceden %d transactions", len(txToProcess))

	if s.state.LastTxHistorical == "" {
		s.state.LastTxHistorical = txToProcess[len(txToProcess)-1].Signature
	}

	numTrades := 0

	for i := len(txToProcess) - 1; i >= 0; i-- {
		tx := txToProcess[i]
		s.state.LastErr = ""
		if s.state.ErrCounter != 0 {
			log.Info("current state.ErrCounter: ", s.state.ErrCounter)
		}

		skipped, err := s.processTx(ctx, tx)

		if err != nil {
			s.state.ErrCounter++

			if s.state.ErrCounter <= s.conf.MaxRetry {
				s.state.LastErr = fmt.Sprintf("unable to process trade transaction(%s): %s", tx.Signature, err.Error())
				log.Error(s.state.LastErr)
				// store state
				if err := s.storeState(ctx); err != nil {
					log.Warnf("unable to store scraper state: %s", err.Error())
					return err
				}
				return err
			}

		}

		if !skipped {
			numTrades++
		}

		// reset consecutive error counter
		s.state.ErrCounter = 0

		// move next
		s.state.LastTx = tx.Signature

		// store state
		if err := s.storeState(ctx); err != nil {
			log.Warnf("unable to store scraper state: %s", err.Error())
			return err
		}
	}

	if err := s.storeState(ctx); err != nil {
		log.Warnf("unable to store scraper state: %s", err.Error())
		return err
	}

	log.Infof("processed %d trades", numTrades)

	return nil

}

func (s *MagicEdenScraper) FetchHistoricalTrades() error {
	log.Info("start fetch historical trades...")
	var err error
	ctx := context.Background()

	// it must be run once at a time
	s.mu.Lock()
	defer s.mu.Unlock()

	// read config
	if err = s.loadConfig(ctx); err != nil {
		log.Warnf("unable to load scraper config: %s", err.Error())
		return err
	}

	// read state
	if err = s.loadState(ctx); err != nil {
		log.Warnf("unable to load scraper state: %s", err.Error())
		return err
	}
	log.Info("last tx: ", s.state.LastTxHistorical)

	if s.state.LastTxHistorical == "" {
		return nil
	}

	txList, err := s.solanaRpcClient.GetSignaturesForAddressWithConfig(ctx, MagicEdenV2ProgramAddress,
		rpc.GetSignaturesForAddressConfig{
			Before: s.state.LastTxHistorical,
			Limit:  s.conf.BatchSize,
		})

	if err != nil {
		log.Warnf("unable to retrieve confirmed transaction signatures for account: %s", err.Error())
		return err
	}

	log.Infof("processing magiceden %d historical transactions", len(txList))

	numTrades := 0

	for i := 0; i < len(txList); i++ {
		tx := txList[i]
		s.state.LastErrHistorical = ""

		skipped, err := s.processTx(ctx, tx)

		if err != nil {
			s.state.ErrCounterHistorical++

			if s.state.ErrCounterHistorical <= s.conf.MaxRetry {
				s.state.LastErrHistorical = fmt.Sprintf("unable to process trade transaction(%s): %s", tx.Signature, err.Error())
				log.Error(s.state.LastErrHistorical)
				// store state
				if err := s.storeState(ctx); err != nil {
					log.Warnf("unable to store scraper state: %s", err.Error())
					return err
				}
				return err
			}
		}

		if !skipped {
			numTrades++
		}

		// reset consecutive error counter
		s.state.ErrCounterHistorical = 0

		// move next
		s.state.LastTxHistorical = tx.Signature

		// store state
		if err := s.storeState(ctx); err != nil {
			log.Warnf("unable to store scraper state: %s", err.Error())
			return err
		}
	}

	if err := s.storeState(ctx); err != nil {
		log.Warnf("unable to store scraper state: %s", err.Error())
		return err
	}

	log.Infof("processed %d historical trades", numTrades)

	return nil

}

func (s *MagicEdenScraper) processTx(ctx context.Context, tx rpc.SignatureWithStatus) (bool, error) {
	confirmedTx, err := s.solanaRpcClient.GetTransaction(ctx, tx.Signature)
	if confirmedTx == nil {
		err = errors.New("confirmedTx == nil")
		log.Error(err)
		return false, err
	}
	if err != nil || confirmedTx.Meta == nil || confirmedTx.Transaction.Message.Accounts == nil {
		log.Errorf("unable to get confirmed transaction with signature %q: %v", tx.Signature, err)
		return false, err
	} else if confirmedTx.Meta.Err != nil {
		return true, err
	}
	instDataEncoded := confirmedTx.Transaction.Message.Instructions[0].Data

	instDataStr := hexutils.BytesToHex(instDataEncoded)
	instDataLowerCase := strings.ToLower(instDataStr)

	if strings.HasPrefix(instDataLowerCase, SaleTxPrefix) && len(confirmedTx.Transaction.Message.Instructions) > 2 {
		nftAddrIndex := confirmedTx.Transaction.Message.Instructions[1].Accounts[2]
		nftAddr := confirmedTx.Transaction.Message.Accounts[nftAddrIndex]
		toIndex := confirmedTx.Transaction.Message.Instructions[0].Accounts[0]
		to := confirmedTx.Transaction.Message.Accounts[toIndex]
		fromIndex := confirmedTx.Transaction.Message.Instructions[2].Accounts[1]
		from := confirmedTx.Transaction.Message.Accounts[fromIndex]

		price := big.NewInt(int64(float64(confirmedTx.Meta.PreBalances[0]) - float64(confirmedTx.Meta.PostBalances[0])))
		normPrice := decimal.NewFromBigInt(price, 0).Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt(9)))
		usdPrice, err := s.calcUSDPrice(normPrice)
		if err != nil {
			return false, err
		}

		metadata, collectionFound, err := s.fetchNFTMetadata(ctx, nftAddr.String())
		if err != nil {
			return false, err
		}

		classMetadata := SolanaNFTMetadata{}
		if collectionFound {
			classMetadata, _, err = s.fetchNFTMetadata(ctx, metadata.collectionAccount.String())
			if err != nil {
				return false, err
			}
		}

		if err := s.notifyTrade(tx, nftAddr, from, to, metadata, classMetadata, price, usdPrice); err != nil {
			return false, err
		}
		return false, nil
	} else {
		return true, nil
	}
}

func (s *MagicEdenScraper) notifyTrade(tx rpc.SignatureWithStatus, addr, from, to common.PublicKey,
	metadata SolanaNFTMetadata, classMetadata SolanaNFTMetadata, price *big.Int, usdPrice float64) error {
	nft, err := s.createOrReadNFT(addr, metadata, classMetadata)
	if err != nil {
		return err
	}

	trade := dia.NFTTrade{
		NFT:         *nft,
		Price:       price,
		PriceUSD:    usdPrice,
		FromAddress: from.String(),
		ToAddress:   to.String(),
		BlockNumber: tx.Slot,
		Timestamp:   time.Unix(*tx.BlockTime, 0),
		TxHash:      tx.Signature,
		Exchange:    MagicEden,
	}

	if asset, ok := assetCacheMagicEden[dia.SOLANA+"-"+SolTokenAddress]; ok {
		trade.Currency = asset
	} else {
		currency, err := s.tradeScraper.datastore.GetAsset(SolTokenAddress, dia.SOLANA)
		if err != nil {
			log.Errorf("cannot fetch asset %s -- %s", dia.SOLANA, SolTokenAddress)
		}
		trade.Currency = currency
		assetCacheMagicEden[dia.SOLANA+"-"+SolTokenAddress] = currency
	}

	fmt.Println("found trade: ", trade)

	// handle close request if the chanTrade not consumed immediately
	select {
	case s.tradeScraper.chanTrade <- trade:
	case <-s.tradeScraper.shutdown:
		return errOpenSeaShutdownRequest
	}

	return nil
}

func (s *MagicEdenScraper) createOrReadNFT(addr common.PublicKey, metadata SolanaNFTMetadata, classMetadata SolanaNFTMetadata) (*dia.NFT, error) {
	collectionAddr := addr.String()
	if classMetadata.name != "" {
		collectionAddr = classMetadata.mintAccount.String()
	}
	nftClass, err := s.tradeScraper.datastore.GetNFTClass(collectionAddr, dia.SOLANA)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			log.Warnf("unable to read nftclass from reldb: %s", err.Error())
			return nil, err
		}

		if classMetadata.name == "" {
			nftClass = dia.NFTClass{
				Address:      collectionAddr,
				Blockchain:   dia.SOLANA,
				Name:         strings.ReplaceAll(metadata.name, "\x00", ""),
				Symbol:       strings.ReplaceAll(metadata.symbol, "\x00", ""),
				ContractType: MagicEden,
			}
		} else {
			nftClass = dia.NFTClass{
				Address:      collectionAddr,
				Blockchain:   dia.SOLANA,
				Name:         strings.ReplaceAll(classMetadata.name, "\x00", ""),
				Symbol:       strings.ReplaceAll(classMetadata.symbol, "\x00", ""),
				ContractType: MagicEden,
			}
		}

		//// name can be empty - refer https://explorer.solana.com/address/F2CQrLYDRyVnpqTvAt6AAXdnZLUhit3i4EbeHAyDMZe3
		if len(nftClass.Name) == 0 {
			nftClass.Name = addr.String()
		}

		if err = s.tradeScraper.datastore.SetNFTClass(nftClass); err != nil {
			log.Printf("%v", nftClass)
			log.Warnf("unable to create nftclass on reldb: %s", err.Error())
			return nil, err
		}
	}

	nft, err := s.tradeScraper.datastore.GetNFT(nftClass.Address, dia.SOLANA, "1")
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			log.Warnf("unable to read nft from reldb: %s", err.Error())
			return nil, err
		}

		nft = dia.NFT{
			NFTClass:     nftClass,
			TokenID:      addr.String(),
			URI:          strings.ReplaceAll(metadata.uri, "\x00", ""),
			CreationTime: time.Unix(metadata.creationTime, 0),
		}

		maxShare := -1
		for _, c := range metadata.creators {
			if c.share > maxShare {
				maxShare = c.share
				nft.CreatorAddress = c.account.String()
			}
		}

		if err = s.tradeScraper.datastore.SetNFT(nft); err != nil {
			log.Warnf("unable to create nft on reldb: %s", err.Error())
			return nil, err
		}
	}

	return &nft, nil
}

func (s *MagicEdenScraper) fetchNFTMetadata(ctx context.Context, nftAddr string) (SolanaNFTMetadata, bool, error) {
	metadata := SolanaNFTMetadata{}

	lastTxFetched := ""
	for {
		txList, err := s.solanaRpcClient.GetSignaturesForAddressWithConfig(ctx, nftAddr,
			rpc.GetSignaturesForAddressConfig{
				Before: lastTxFetched,
			})
		if err != nil {
			log.Warnf("unable to retrieve transaction signatures for account: %s", err.Error())
			return metadata, false, err
		}

		if len(txList) == 0 {
			break
		}

		lastTxFetched = txList[len(txList)-1].Signature
	}

	var addr common.PublicKey
	if lastTxFetched != "" {
		confirmedTx, err := s.solanaRpcClient.GetTransaction(ctx, lastTxFetched)
		if confirmedTx == nil {
			log.Error("confirmedTX == nil")
			return metadata, false, err
		}
		if err != nil || confirmedTx.Meta == nil ||
			confirmedTx.Meta.Err != nil || confirmedTx.Transaction.Message.Accounts == nil {
			log.Errorf("unable to get confirmed transaction with signature %q: %v", lastTxFetched, err)
			return metadata, false, err
		}
		for i, postBalance := range confirmedTx.Meta.PostBalances {
			normPrice := decimal.NewFromInt(postBalance).Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt(9)))
			if normPrice.Equals(decimal.NewFromFloat(MetadataFee)) {
				addr = confirmedTx.Transaction.Message.Accounts[i]
				break
			}
		}
		metadata.creationTime = *confirmedTx.BlockTime
	} else {
		return metadata, false, errors.New("unable to fetch create tx for nft")
	}

	if out, err := s.solanaRpcClient.GetAccountInfo(ctx, addr.String()); err != nil {
		return metadata, false, err
	} else {
		if out.Data != nil {
			data := out.Data
			if len(data) > 0 {
				i := 1
				if len(data) >= i+32 {
					source := data[i : i+32]
					metadata.sourceAccount = common.PublicKeyFromBytes(source)
				} else {
					return metadata, false, err
				}
				i += 32

				if len(data) >= i+32 {
					mint := data[i : i+32]
					metadata.mintAccount = common.PublicKeyFromBytes(mint)
				} else {
					return metadata, false, err
				}
				i += 32

				nameSize := int(data[i])
				i += 4
				if len(data) >= i+nameSize {
					metadata.name = string(data[i : i+nameSize])
				} else {
					return metadata, false, err
				}
				i += nameSize

				symbolSize := int(data[i])
				i += 4
				if len(data) >= i+symbolSize {
					metadata.symbol = string(data[i : i+symbolSize])
				} else {
					return metadata, false, err
				}
				i += symbolSize

				uriSize := int(data[i])
				i += 4
				if len(data) >= i+uriSize {
					metadata.uri = string(data[i : i+uriSize])
				} else {
					return metadata, false, err
				}
				i += uriSize

				if len(data) > i {
					metadata.fee = int(data[i])
				} else {
					return metadata, false, err
				}
				i += 2

				hasCreator := 0
				if len(data) > i {
					hasCreator = int(data[i])
				} else {
					return metadata, false, err
				}
				i += 1
				if hasCreator == 1 {
					creatorCount := int(data[i])
					metadata.creators = make([]NFTCreator, creatorCount)
					i += 4
					for j := 0; j < creatorCount; j++ {
						nftCreator := NFTCreator{}
						if len(data) >= i+32 {
							account := data[i : i+32]
							nftCreator.account = common.PublicKeyFromBytes(account)
						} else {
							return metadata, false, err
						}
						i += 32

						if len(data) > i {
							nftCreator.verified = int(data[i])
						} else {
							return metadata, false, err
						}
						i += 1

						if len(data) > i {
							nftCreator.share = int(data[i])
						} else {
							return metadata, false, err
						}
						i += 1

						metadata.creators[j] = nftCreator
					}
				}

				if len(data) > i {
					metadata.primarySaleDone = int(data[i])
				} else {
					return metadata, false, err
				}
				i += 1

				if len(data) > i {
					metadata.isMutable = int(data[i])
				} else {
					return metadata, false, err
				}
				i += 1

				// skip edition nonce
				if len(data) > i {
					if int(data[i]) == 1 {
						i += 2
					} else {
						i += 1
					}
				} else {
					return metadata, false, err
				}

				// skip token standard
				if len(data) > i {
					if int(data[i]) == 1 {
						i += 2
					} else {
						i += 1
					}
				} else {
					return metadata, false, err
				}

				collectionFound := false
				if len(data) > i {
					if int(data[i]) == 1 {
						i += 2
						if len(data) >= i+32 {
							collection := data[i : i+32]
							metadata.collectionAccount = common.PublicKeyFromBytes(collection)
							collectionFound = true
						} else {
							return metadata, false, err
						}
					}
				} else {
					return metadata, false, err
				}

				return metadata, collectionFound, err

			} else {
				return metadata, false, errors.New("no data found in account info response")
			}
		} else {
			return metadata, false, errors.New("no value found in account info response")
		}
	}
}

func (s *MagicEdenScraper) calcUSDPrice(price decimal.Decimal) (float64, error) {
	tokenPrice, err := decimal.NewFromString("100")
	if err != nil {
		return 0, err
	}

	// TODO: find the token price in usd for the given block number
	usdPrice := price.Mul(tokenPrice)

	// using float type is not a good idea to handle prices
	// we ignore if the price cannot be presentable as float64
	f, _ := usdPrice.Float64()

	return f, nil
}

func (s *MagicEdenScraper) GetTradeChannel() chan dia.NFTTrade {
	return s.tradeScraper.chanTrade
}
