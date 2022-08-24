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
	"github.com/shopspring/decimal"
	"github.com/status-im/keycard-go/hexutils"
	"github.com/streamingfast/solana-go"
	"github.com/streamingfast/solana-go/rpc"
)

const (
	rpcEndpointSolana         = "https://solana-api.projectserum.com"
	MagicEdenV2ProgramAddress = "M2mx93ekt1fmXSVkTrUL9xVFHkmME8HTUi5Cyc5aF7K"
	SaleTxPrefix              = "f223c68952e1f2b6"
	MetadataFee               = 0.00561672
)

var (
	defMagicEdenConf = &MagicEdenScraperConfig{
		SolanaRestUri: SaleTxPrefix,
		ProgramAddr:   MagicEdenV2ProgramAddress,
		BatchSize:     1000,
		WaitPeriod:    30 * time.Second,
		MaxRetry:      5,
		SkipOnErr:     true,
	}
	defMagicEdenState   = &MagicEdenScraperState{}
	MagicEden           = ""
	assetCacheMagicEden = make(map[string]dia.Asset)
)

type SolanaNFTMetadata struct {
	sourceAccount   solana.PublicKey
	mintAccount     solana.PublicKey
	name            string
	symbol          string
	uri             string
	fee             int
	creators        []NFTCreator
	primarySaleDone int
	isMutable       int
}

type NFTCreator struct {
	account  solana.PublicKey
	verified int
	share    int
}

type MagicEdenScraper struct {
	solanaRpcClient *rpc.Client
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

func NewMagicEdenScraper(rdb *models.RelDB) *MagicEdenScraper {
	ctx := context.Background()
	scraper := &MagicEdenScraper{
		solanaRpcClient: rpc.NewClient(utils.Getenv("SOLANA_URI_REST", rpcEndpointSolana)),
		tradeScraper: TradeScraper{
			shutdown:     make(chan nothing),
			shutdownDone: make(chan nothing),
			datastore:    rdb,
			chanTrade:    make(chan dia.NFTTrade),
			source:       "MagicEden",
		},
	}
	MagicEden = utils.Getenv("SCRAPER_NAME_STATE", "")
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

		defState := *defMagicEdenState // copy
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

		if err := s.FetchHistoricalTrades(); err != nil {
			if errors.Is(err, errOpenSeaShutdownRequest) {
				stop = true
				continue
			}
		}

		log.Debugf("wait for %s", 5*time.Second)

		select {
		case <-time.After(5 * time.Second):
		case <-s.tradeScraper.shutdown:
			stop = true
		}
	}

}

func (s *MagicEdenScraper) FetchTrades() error {
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

	txToProcess := make([]*rpc.TransactionSignature, 0)
	lastFetchedTx := ""
	for {
		txList, err := s.solanaRpcClient.GetSignaturesForAddress(ctx, solana.MustPublicKeyFromBase58(MagicEdenV2ProgramAddress),
			&rpc.GetSignaturesForAddressOpts{
				Before: lastFetchedTx,
				Until:  s.state.LastTx,
				Limit:  uint64(s.conf.BatchSize),
			})
		if err != nil {
			log.Warnf("unable to retrieve confirmed transaction signatures for account: %s", err.Error())
			return err
		}

		for _, tx := range txList {
			txToProcess = append(txToProcess, tx)
			lastFetchedTx = tx.Signature
		}

		if len(txList) == 0 || s.state.LastTx == "" {
			break
		}

	}

	log.Infof("processing magiceden %d transactions", len(txToProcess))

	numTrades := 0

	for i := len(txToProcess) - 1; i >= 0; i-- {
		tx := txToProcess[i]
		s.state.LastErr = ""
		log.Info("current state.ErrCounter: ", s.state.ErrCounter)

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

	txList, err := s.solanaRpcClient.GetSignaturesForAddress(ctx, solana.MustPublicKeyFromBase58(MagicEdenV2ProgramAddress),
		&rpc.GetSignaturesForAddressOpts{
			Before: s.state.LastTxHistorical,
			Limit:  uint64(s.conf.BatchSize),
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
		log.Info("done with tx")
	}

	if err := s.storeState(ctx); err != nil {
		log.Warnf("unable to store scraper state: %s", err.Error())
		return err
	}

	log.Infof("processed %d historical trades", numTrades)

	return nil

}

func (s *MagicEdenScraper) processTx(ctx context.Context, tx *rpc.TransactionSignature) (bool, error) {
	confirmedTx, err := s.solanaRpcClient.GetConfirmedTransaction(ctx, tx.Signature)
	if err != nil || confirmedTx.Meta == nil ||
		confirmedTx.Transaction == nil || confirmedTx.Transaction.Message.AccountKeys == nil {
		log.Errorf("unable to get confirmed transaction with signature %q: %v", tx.Signature, err)
		return false, err
	} else if confirmedTx.Meta.Err != nil {
		return true, err
	}
	instData := hexutils.BytesToHex(confirmedTx.Transaction.Message.Instructions[0].Data)
	instDataLowerCase := strings.ToLower(instData)

	if strings.HasPrefix(instDataLowerCase, SaleTxPrefix) && len(confirmedTx.Transaction.Message.Instructions) > 2 {
		nftAddrIndex := confirmedTx.Transaction.Message.Instructions[1].Accounts[2]
		nftAddr := confirmedTx.Transaction.Message.AccountKeys[nftAddrIndex]
		toIndex := confirmedTx.Transaction.Message.Instructions[0].Accounts[0]
		to := confirmedTx.Transaction.Message.AccountKeys[toIndex]
		fromIndex := confirmedTx.Transaction.Message.Instructions[2].Accounts[1]
		from := confirmedTx.Transaction.Message.AccountKeys[fromIndex]

		price := big.NewInt(int64(float64(confirmedTx.Meta.PreBalances[0]) - float64(confirmedTx.Meta.PostBalances[0])))
		normPrice := decimal.NewFromBigInt(price, 0).Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt(9)))
		usdPrice, err := s.calcUSDPrice(normPrice)
		if err != nil {
			return false, err
		}

		metadata, err := s.fetchNFTMetadata(ctx, nftAddr)
		if err != nil {
			return false, err
		}

		if err := s.notifyTrade(tx, nftAddr, from, to, metadata, price, usdPrice); err != nil {
			return false, err
		}
		return false, nil
	} else {
		return true, nil
	}
}

func (s *MagicEdenScraper) notifyTrade(tx *rpc.TransactionSignature, addr, from, to solana.PublicKey, metadata SolanaNFTMetadata, price *big.Int, usdPrice float64) error {
	nft, err := s.createOrReadNFT(tx, addr, metadata)
	if err != nil {
		return err
	}

	// Get block time.
	block, err := s.solanaRpcClient.GetConfirmedBlock(context.TODO(), uint64(tx.Slot), "")
	if err != nil && block != nil {
		log.Errorf("getting block time: %+v", err)
	}

	trade := dia.NFTTrade{
		NFT:         *nft,
		Price:       price,
		PriceUSD:    usdPrice,
		FromAddress: from.String(),
		ToAddress:   to.String(),
		BlockNumber: uint64(tx.Slot),
		Timestamp:   time.Unix(int64(block.BlockTime), 0),
		TxHash:      tx.Signature,
		Exchange:    "MagicEden",
	}

	if asset, ok := assetCacheMagicEden[dia.SOLANA+"-"+tx.Signature]; ok {
		trade.Currency = asset
	} else {
		currency, err := s.tradeScraper.datastore.GetAsset(tx.Signature, dia.SOLANA)
		if err != nil {
			log.Errorf("cannot fetch asset %s -- %s", dia.SOLANA, tx.Signature)
		}
		trade.Currency = currency
		assetCacheMagicEden[dia.SOLANA+"-"+tx.Signature] = currency
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

func (s *MagicEdenScraper) createOrReadNFT(tx *rpc.TransactionSignature, addr solana.PublicKey, metadata SolanaNFTMetadata) (*dia.NFT, error) {
	nftClass, err := s.tradeScraper.datastore.GetNFTClass(tx.Signature, dia.SOLANA)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			log.Warnf("unable to read nftclass from reldb: %s", err.Error())
			return nil, err
		}

		nftClass = dia.NFTClass{
			Address:    addr.String(),
			Blockchain: dia.SOLANA,
			Name:       metadata.name,
			Symbol:     metadata.symbol,
		}

		if err = s.tradeScraper.datastore.SetNFTClass(nftClass); err != nil {
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
			NFTClass: nftClass,
			TokenID:  "1",
			URI:      metadata.uri,
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

func (s *MagicEdenScraper) fetchNFTMetadata(ctx context.Context, nftAddr solana.PublicKey) (SolanaNFTMetadata, error) {
	metadata := SolanaNFTMetadata{}

	lastTxFetched := ""
	for {
		txList, err := s.solanaRpcClient.GetSignaturesForAddress(ctx, nftAddr,
			&rpc.GetSignaturesForAddressOpts{
				Before: lastTxFetched,
			})
		if err != nil {
			log.Warnf("unable to retrieve transaction signatures for account: %s", err.Error())
			return metadata, err
		}

		if len(txList) == 0 {
			break
		}

		lastTxFetched = txList[len(txList)-1].Signature
	}

	var addr solana.PublicKey
	if lastTxFetched != "" {
		confirmedTx, err := s.solanaRpcClient.GetConfirmedTransaction(ctx, lastTxFetched)
		if err != nil || confirmedTx.Meta == nil || confirmedTx.Meta.Err != nil ||
			confirmedTx.Transaction == nil || confirmedTx.Transaction.Message.AccountKeys == nil {
			log.Errorf("unable to get confirmed transaction with signature %q: %v", lastTxFetched, err)
			return metadata, err
		}
		for i, postBalance := range confirmedTx.Meta.PostBalances {
			normPrice := decimal.NewFromInt(int64(postBalance)).Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt(9)))
			if normPrice.Equals(decimal.NewFromFloat(MetadataFee)) {
				addr = confirmedTx.Transaction.Message.AccountKeys[i]
				break
			}
		}
	} else {
		return metadata, errors.New("unable to fetch create tx for nft")
	}

	if out, err := s.solanaRpcClient.GetAccountInfo(ctx, addr); err != nil {
		return metadata, err
	} else {
		if out != nil && out.Value != nil {
			data := out.Value.Data
			if len(data) > 0 {
				i := 1
				source := data[i : i+32]
				metadata.sourceAccount = solana.PublicKeyFromBytes(source)
				i += 32

				mint := data[i : i+32]
				metadata.mintAccount = solana.PublicKeyFromBytes(mint)
				i += 32

				nameSize := int(data[i])
				i += 4
				metadata.name = string(data[i : i+nameSize])
				i += nameSize

				symbolSize := int(data[i])
				i += 4
				metadata.symbol = string(data[i : i+symbolSize])
				i += symbolSize

				uriSize := int(data[i])
				i += 4
				metadata.uri = string(data[i : i+uriSize])
				i += uriSize

				metadata.fee = int(data[i])
				i += 2

				hasCreator := int(data[i])
				i += 1
				if hasCreator == 1 {
					creatorCount := int(data[i])
					metadata.creators = make([]NFTCreator, creatorCount)
					i += 4
					for j := 0; j < creatorCount; j++ {
						nftCreator := NFTCreator{}
						account := data[i : i+32]
						nftCreator.account = solana.PublicKeyFromBytes(account)
						i += 32

						nftCreator.verified = int(data[i])
						i += 1

						nftCreator.share = int(data[i])
						i += 1

						metadata.creators = append(metadata.creators, nftCreator)
					}
				}

				metadata.primarySaleDone = int(data[i])
				i += 1

				metadata.isMutable = int(data[i])

				return metadata, err

			} else {
				return metadata, errors.New("no data found in account info response")
			}
		} else {
			return metadata, errors.New("no value found in account info response")
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
