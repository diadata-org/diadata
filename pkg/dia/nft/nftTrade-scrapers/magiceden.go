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
	defMagicEdenState = &MagicEdenScraperState{}
)

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

	LastTxHistorical uint `json:"last_tx_historical"`

	LastErr string `json:"last_error"`

	ErrCounter int `json:"count_of_error"`
}

func (s *MagicEdenScraper) loadConfig(ctx context.Context) error {
	return s.tradeScraper.datastore.GetScraperConfig(ctx, OpenSea, s.conf)
}

func (s *MagicEdenScraper) loadState(ctx context.Context) error {
	return s.tradeScraper.datastore.GetScraperState(ctx, OpenSea, s.state)
}

func (s *MagicEdenScraper) storeState(ctx context.Context) error {
	return s.tradeScraper.datastore.SetScraperState(ctx, OpenSea, s.state)
}

func NewMagicEdenScraper(rdb *models.RelDB) *MagicEdenScraper {
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

		log.Debugf("wait for %s", 60*time.Second)

		select {
		case <-time.After(60 * time.Second):
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
	for {
		txList, err := s.solanaRpcClient.GetSignaturesForAddress(ctx, solana.MustPublicKeyFromBase58(MagicEdenV2ProgramAddress),
			&rpc.GetSignaturesForAddressOpts{
				Until: s.state.LastTx,
			})
		if err != nil {
			log.Warnf("unable to retrieve confirmed transaction signatures for account: %s", err.Error())
			return err
		}

		if len(txList) == 0 {
			break
		}

		for _, tx := range txList {
			txToProcess = append(txToProcess, tx)
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

func (s *MagicEdenScraper) processTx(ctx context.Context, tx *rpc.TransactionSignature) (bool, error) {
	confirmedTx, err := s.solanaRpcClient.GetConfirmedTransaction(ctx, tx.Signature)
	if err != nil || confirmedTx.Meta == nil || confirmedTx.Meta.Err != nil ||
		confirmedTx.Transaction == nil || confirmedTx.Transaction.Message.AccountKeys == nil {
		log.Errorf("unable to get confirmed transaction with signature %q: %w", tx.Signature, err)
		return true, err
	}
	instData := hexutils.BytesToHex(confirmedTx.Transaction.Message.Instructions[0].Data)
	instDataLowerCase := strings.ToLower(instData)

	if strings.HasPrefix(instDataLowerCase, SaleTxPrefix) {
		if err := s.notifyTrade(tx, confirmedTx); err != nil {
			return false, err
		}
		return false, nil
	} else {
		return true, nil
	}
}

func (s *MagicEdenScraper) notifyTrade(tx *rpc.TransactionSignature, txWithMeta rpc.TransactionWithMeta) error {
	//nftClass, err := s.createOrReadNFTClass(transfer)
	//if err != nil {
	//	return err
	//}
	//
	//nft, err := s.createOrReadNFT(nftClass, transfer)
	//if err != nil {
	//	return err
	//}
	//
	//// Get block time.
	//timestamp, err := ethhelper.GetBlockTimeEth(int64(ev.Raw.BlockNumber), s.tradeScraper.datastore, s.tradeScraper.ethConnection)
	//if err != nil {
	//	log.Errorf("getting block time: %+v", err)
	//}

	nftAddrIndex := txWithMeta.Transaction.Message.Instructions[1].Accounts[2]
	nftAddr := txWithMeta.Transaction.Message.AccountKeys[nftAddrIndex]
	fromIndex := txWithMeta.Transaction.Message.Instructions[0].Accounts[0]
	from := txWithMeta.Transaction.Message.AccountKeys[fromIndex]
	toIndex := txWithMeta.Transaction.Message.Instructions[2].Accounts[1]
	to := txWithMeta.Transaction.Message.AccountKeys[toIndex]

	price := big.NewInt(int64(float64(txWithMeta.Meta.PostBalances[0]) - float64(txWithMeta.Meta.PreBalances[0])))
	normPrice := decimal.NewFromBigInt(price, 0).Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt(9)))
	usdPrice, err := s.calcUSDPrice(normPrice)
	if err != nil {
		return err
	}

	trade := dia.NFTTrade{
		NFT: dia.NFT{
			NFTClass: dia.NFTClass{Address: nftAddr.String()},
		},
		Price:       price,
		PriceUSD:    usdPrice,
		FromAddress: from.String(),
		ToAddress:   to.String(),
		BlockNumber: uint64(tx.Slot),
		Timestamp:   time.Now(),
		TxHash:      tx.Signature,
		Exchange:    "MagicEden",
	}

	if asset, ok := assetCacheOpensea[dia.SOLANA+"-"+tx.Signature]; ok {
		trade.Currency = asset
	} else {
		currency, err := s.tradeScraper.datastore.GetAsset(tx.Signature, dia.SOLANA)
		if err != nil {
			log.Errorf("cannot fetch asset %s -- %s", dia.SOLANA, tx.Signature)
		}
		trade.Currency = currency
		assetCacheOpensea[dia.SOLANA+"-"+tx.Signature] = currency
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

//func (s *MagicEdenScraper) createOrReadNFTClass(transfer *erc721Transfer) (*dia.NFTClass, error) {
//	nftClass, err := s.tradeScraper.datastore.GetNFTClass(transfer.NFTAddress.Hex(), dia.ETHEREUM)
//	if err != nil {
//		if !errors.Is(err, pgx.ErrNoRows) {
//			log.Warnf("unable to read nftclass from reldb: %s", err.Error())
//			return nil, err
//		}
//
//		nftClass = dia.NFTClass{
//			Address:      transfer.NFTAddress.Hex(),
//			Blockchain:   dia.ETHEREUM,
//			ContractType: openSeaNFTContractType,
//		}
//
//		if transfer.Name != nil {
//			nftClass.Name = *transfer.Name
//		}
//
//		if transfer.Symbol != nil {
//			nftClass.Symbol = *transfer.Symbol
//		}
//
//		if err = s.tradeScraper.datastore.SetNFTClass(nftClass); err != nil {
//			log.Warnf("unable to create nftclass on reldb: %s", err.Error())
//			return nil, err
//		}
//	}
//
//	return &nftClass, nil
//}
//
//func (s *MagicEdenScraper) createOrReadNFT(nftClass *dia.NFTClass, transfer *erc721Transfer) (*dia.NFT, error) {
//	nft, err := s.tradeScraper.datastore.GetNFT(nftClass.Address, dia.ETHEREUM, transfer.TokenID.String())
//	if err != nil {
//		if !errors.Is(err, pgx.ErrNoRows) {
//			log.Warnf("unable to read nft from reldb: %s", err.Error())
//			return nil, err
//		}
//
//		createdBy, createdAt, err := s.findContractCreationInfo(context.Background(), common.HexToAddress(nftClass.Address))
//		if err != nil {
//			log.Warnf("unable to find the creation info for the nft contract(%s): %s", nftClass.Address, err.Error())
//			return nil, err
//		}
//
//		nft = dia.NFT{
//			NFTClass:       *nftClass,
//			TokenID:        transfer.TokenID.String(),
//			CreationTime:   createdAt,
//			CreatorAddress: createdBy.Hex(),
//			Attributes:     transfer.TokenAttrs,
//		}
//
//		if transfer.TokenURI != nil {
//			nft.URI = *transfer.TokenURI
//		}
//
//		if err = s.tradeScraper.datastore.SetNFT(nft); err != nil {
//			log.Warnf("unable to create nft on reldb: %s", err.Error())
//			return nil, err
//		}
//	}
//
//	return &nft, nil
//}

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
