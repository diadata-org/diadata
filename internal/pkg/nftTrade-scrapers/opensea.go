package nfttradescrapers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"math/big"
	"net/http"
	"time"

	"github.com/diadata-org/diadata/config/nftContracts/erc721"
	"github.com/diadata-org/diadata/config/nftContracts/opensea"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"
)

const (
	// opensea's exchange contract address on ethereum mainnet
	OpenSeaContractAddr = "0x7be8076f4ea4a4ad08075c2508e481d6c946d12b"

	// opensea's contract has been deployed on ...
	// will be used if there is no known last block info
	openSeaContractDeployedBlockNum = 5774644

	// indicates how many blocks we should stay behind head of the chain
	openSeaFollowingDistanceBlocks = 2

	// indicates the batch size during read the filtered events
	openSeaEvFetchBatchSize = 10000

	// wait for a while between batch retrieval of filtered events
	openSeaWaitPerBatch = time.Second * 10

	// it enables read contract data from the event's block height
	// instead of the last state
	openSeaUseArchiveNodeFeatures = false

	// we assume all of the NFTs traded on OpenSea are ERC721(1155 is an extension of it)
	openSeaNFTContractType = "ERC721"

	openSeaReadTotalSupply = true

	openSeaFetchAttrsTimeout     = 30 * time.Second
	openSeaFetchAttrsMaxSizeInKB = 50

	OpenSea = "OpenSea"
)

var (
	// 10^18
	tenToPowerOf18 = decimal.NewFromInt(10).Pow(decimal.NewFromInt(18))

	errOpenSeaShutdownRequest = errors.New("shutdown requested")
)

type OpenSeaScraper struct {
	contract     *opensea.Contract
	tradeScraper TradeScraper
	contractAddr common.Address
	ticker       *time.Ticker

	// since multiple OrderMatched events can be in a single block,
	// we need to store the last processed log index along with the block number.
	lastBlockNum *big.Int
	lastLogIndex *uint
}

type erc721Transfer struct {
	NFTAddress  common.Address
	Name        *string
	Symbol      *string
	TotalSupply *big.Int
	From        common.Address
	To          common.Address
	TokenID     *big.Int
	TokenURI    *string
	TokenAttrs  map[string]interface{}
}

func NewOpenSeaScraper(rdb *models.RelDB) *OpenSeaScraper {
	var err error

	scraper := &OpenSeaScraper{
		contractAddr: common.HexToAddress(OpenSeaContractAddr),
		ticker:       time.NewTicker(openSeaWaitPerBatch),
		tradeScraper: TradeScraper{
			shutdown:     make(chan nothing),
			shutdownDone: make(chan nothing),
			datastore:    rdb,
			chanTrade:    make(chan dia.NFTTrade),
			source:       OpenSea,
		},
	}

	scraper.tradeScraper.ethConnection, err = ethhelper.NewETHClient()
	if err != nil {
		log.Errorf("unable to get ethereum client: %s", err.Error())
		return scraper
	}

	scraper.contract, err = opensea.NewContract(common.HexToAddress(OpenSeaContractAddr), scraper.tradeScraper.ethConnection)
	if err != nil {
		log.Errorf("unable to create opensea contract instance: %s", err.Error())
		return scraper
	}

	go scraper.mainLoop()

	return scraper
}

func (s *OpenSeaScraper) loadScraperState() error {
	// TODO
	// I can't use GetLastBlockNFTTrade() due to it depends on a specific NFT.
	// OpenSea is an exchange, with many different NFTs trading on it.
	// We track its OrderMatched events, every single event can point out a different NFT.
	// So for OpenSea, we need to store the last processed event's block number as the last
	// known block number and we need to use it during the fresh start.

	// TODO: read the block number of last processed ordermatch event from rdb
	s.lastBlockNum = nil

	// TODO: read the log index of last processed ordermatch event from rdb
	s.lastLogIndex = nil

	// set the last block number as the same as the deployment of opensea's contract
	// if not processed any events before
	if s.lastBlockNum == nil || s.lastBlockNum.Cmp(&big.Int{}) == 0 {
		s.lastBlockNum = new(big.Int).SetInt64(openSeaContractDeployedBlockNum)
	}

	return nil
}

func (s *OpenSeaScraper) storeScraperState() error {
	// TODO: save the block number and log index value for the last processed event to the rdb

	log.Debugf("store the scraper state(lastBlock: %d, lastLog: %v)", s.lastBlockNum, s.lastLogIndex)

	return nil
}

func (s *OpenSeaScraper) mainLoop() {
	if err := s.loadScraperState(); err != nil {
		log.Errorf("unable to load scraper state: %s", err.Error())
		return
	}

	defer func() {
		s.tradeScraper.closed = true

		close(s.tradeScraper.chanTrade)
		close(s.tradeScraper.shutdownDone)
	}()

	log.Infof("opensea scraper has been started (batch: %d, period: %s)", openSeaEvFetchBatchSize, openSeaWaitPerBatch.String())

	for stop := false; !stop; {
		select {
		case <-s.ticker.C:
			if err := s.FetchTrades(); err != nil {
				if errors.Is(err, errOpenSeaShutdownRequest) {
					stop = true
				}

				continue
			}

			if err := s.storeScraperState(); err != nil {
				log.Errorf("unable to store the scraper state: %s", err.Error())
				continue
			}

		case <-s.tradeScraper.shutdown:
			stop = true
		}
	}
}

func (s *OpenSeaScraper) FetchTrades() error {
	ctx := context.Background()

	startBlockNum := s.lastBlockNum.Uint64()
	endBlockNum := startBlockNum + openSeaEvFetchBatchSize

	if sync, err := s.tradeScraper.ethConnection.SyncProgress(ctx); err != nil {
		log.Errorf("unable to get synchronization info: %s", err.Error())
		return err

	} else if sync != nil {
		log.Warnf("this cycle has been skipped because the connected blockchain node is not yet synchronized(head: %d)", sync.CurrentBlock)
		return nil
	}

	head, err := s.tradeScraper.ethConnection.BlockNumber(ctx)
	if err != nil {
		log.Errorf("unable to get block number of the head: %s", err.Error())
		return err
	}

	// stay behind of the head if needed
	if endBlockNum > head-openSeaFollowingDistanceBlocks {
		endBlockNum = head - openSeaFollowingDistanceBlocks
	}

	if startBlockNum > endBlockNum {
		log.Debugf("scraper has been synchronised(head: %d)", endBlockNum)
		return nil
	}

	log.Debugf("fetching opensea trades between %d and %d", startBlockNum, endBlockNum)

	filter, err := s.contract.FilterOrdersMatched(&bind.FilterOpts{Start: startBlockNum, End: &endBlockNum}, nil, nil, nil)
	if err != nil {
		log.Errorf("unable to create ordersmatched filter: %s", err.Error())
		return err
	}

	for filter.Next() {
		if err := filter.Error(); err != nil {
			log.Errorf("unable to parse filtered event: %s", err.Error())
			return err
		}

		ev := filter.Event

		currBlockNum := new(big.Int).SetUint64(ev.Raw.BlockNumber)

		if s.lastBlockNum.Cmp(currBlockNum) != 0 {
			s.lastLogIndex = nil
		}

		s.lastBlockNum = currBlockNum

		// skip if the event has already been processed
		if s.lastLogIndex != nil && ev.Raw.Index <= *s.lastLogIndex {
			log.Tracef("event(block: %d, log: %d, tx: %s) skipped because it has already been processed", ev.Raw.BlockNumber, ev.Raw.Index, ev.Raw.TxHash.Hex())
			continue
		}

		// skip if the event removed due to chain reorganization
		if ev.Raw.Removed {
			log.Tracef("event(block: %d, log: %d, tx: %s) skipped due to chain reorganization", ev.Raw.BlockNumber, ev.Raw.Index, ev.Raw.TxHash.Hex())
			continue
		}

		transfers, err := s.findERC721Transfers(ctx, ev.Raw.TxHash)
		if err != nil {
			log.Errorf("unable to find transfers of the event(block: %d, log: %d, tx: %s): %s", ev.Raw.BlockNumber, ev.Raw.Index, ev.Raw.TxHash.Hex(), err.Error())
			return err
		}

		// skip if the event has no transfer
		if len(transfers) == 0 {
			log.Tracef("event(block: %d, log: %d, tx: %s) skipped due to it has no erc721 transfer log", ev.Raw.BlockNumber, ev.Raw.Index, ev.Raw.TxHash.Hex())
			s.lastLogIndex = &ev.Raw.Index
			continue
		}

		// skip if the event has multiple transfers due to we can't calculate the price of trade
		if len(transfers) > 1 {
			log.Tracef("event(block: %d, log: %d, tx: %s) skipped due to it has multiple erc721 transfer logs", ev.Raw.BlockNumber, ev.Raw.Index, ev.Raw.TxHash.Hex())
			s.lastLogIndex = &ev.Raw.Index
			continue
		}

		if err := s.notifyTrade(ev, transfers[0]); err != nil {
			if !errors.Is(err, errOpenSeaShutdownRequest) {
				log.Warnf("event(block: %d, log: %d, tx: %s) couldn't processed: %s", ev.Raw.BlockNumber, ev.Raw.Index, ev.Raw.TxHash.Hex(), err.Error())
			}

			return err
		}

		// store the log index of last processed event
		s.lastLogIndex = &ev.Raw.Index
	}

	s.lastBlockNum = new(big.Int).SetUint64(endBlockNum + 1)
	s.lastLogIndex = nil

	return nil
}

func (s *OpenSeaScraper) createOrReadNFTClass(transfer *erc721Transfer) (*dia.NFTClass, error) {
	nftClass, err := s.tradeScraper.datastore.GetNFTClass(transfer.NFTAddress.Hex(), dia.ETHEREUM)
	if err != nil {
		if err != pgx.ErrNoRows {
			log.Warnf("unable to read nftclass from reldb: %s", err.Error())
			return nil, err
		}

		nftClass = dia.NFTClass{
			Address:      transfer.NFTAddress.Hex(),
			Blockchain:   dia.ETHEREUM,
			ContractType: openSeaNFTContractType,
		}

		if transfer.Name != nil {
			nftClass.Name = *transfer.Name
		}

		if transfer.Symbol != nil {
			nftClass.Symbol = *transfer.Symbol
		}

		if err = s.tradeScraper.datastore.SetNFTClass(nftClass); err != nil {
			log.Warnf("unable to create nftclass on reldb: %s", err.Error())
			return nil, err
		}
	}

	return &nftClass, nil
}

func (s *OpenSeaScraper) createOrReadNFT(nftClass *dia.NFTClass, transfer *erc721Transfer) (*dia.NFT, error) {
	nft, err := s.tradeScraper.datastore.GetNFT(nftClass.Address, dia.ETHEREUM, transfer.TokenID.String())
	if err != nil {
		if err != pgx.ErrNoRows {
			log.Warnf("unable to read nft from reldb: %s", err.Error())
			return nil, err
		}

		createdBy, createdAt, err := s.findContractCreationInfo(context.Background(), common.HexToAddress(nftClass.Address))
		if err != nil {
			log.Warnf("unable to find the creation info for the nft contract(%s): %s", nftClass.Address, err.Error())
			return nil, err
		}

		nft = dia.NFT{
			NFTClass:       *nftClass,
			TokenID:        transfer.TokenID.String(),
			CreationTime:   createdAt,
			CreatorAddress: createdBy.Hex(),
			Attributes:     transfer.TokenAttrs,
		}

		if transfer.TokenURI != nil {
			nft.URI = *transfer.TokenURI
		}

		if err = s.tradeScraper.datastore.SetNFT(nft); err != nil {
			log.Warnf("unable to create nft on reldb: %s", err.Error())
			return nil, err
		}
	}

	return &nft, nil
}

func (s *OpenSeaScraper) notifyTrade(ev *opensea.ContractOrdersMatched, transfer *erc721Transfer) error {
	nftClass, err := s.createOrReadNFTClass(transfer)
	if err != nil {
		return err
	}

	nft, err := s.createOrReadNFT(nftClass, transfer)
	if err != nil {
		return err
	}

	trade := dia.NFTTrade{
		NFT:         *nft,
		BlockNumber: new(big.Int).SetUint64(ev.Raw.BlockNumber),
		// different NFTs can use different asset instead of eth for trades,
		// it could be found the used asset token by checking more log records
		// for now we assume eth is used for trades
		PriceUSD:    s.calcUSDFromWEI(ev.Price, ev.Raw.BlockNumber),
		FromAddress: transfer.From,
		ToAddress:   transfer.To,
		Exchange:    OpenSea,
		TxHash:      ev.Raw.TxHash,
	}

	// handle close request too if the chanTrade not consumed immediately
	select {
	case s.tradeScraper.chanTrade <- trade:
	case <-s.tradeScraper.shutdown:
		return errOpenSeaShutdownRequest
	}

	return nil
}

func (s *OpenSeaScraper) calcUSDFromWEI(wei *big.Int, blockNum uint64) float64 {
	ethPrice, err := s.ethPrice(blockNum)
	if err != nil {
		return 0
	}

	eth := decimal.NewFromBigInt(wei, 0).Div(tenToPowerOf18)

	price := eth.Mul(ethPrice)

	// using float type is not a good idea to handle prices
	// we ignore if the price cannot be presentable as float64
	f, _ := price.Float64()

	return f
}

func (s *OpenSeaScraper) ethPrice(blockNum uint64) (decimal.Decimal, error) {
	// TODO: find the eth price in usd for the given block number
	return decimal.NewFromString("1")
}

// GetDataChannel returns the scrapers data channel.
func (s *OpenSeaScraper) GetTradeChannel() chan dia.NFTTrade {
	return s.tradeScraper.chanTrade
}

func (s *OpenSeaScraper) Close() error {
	if s.tradeScraper.closed {
		return errors.New("scraper already closed")
	}

	close(s.tradeScraper.shutdown)

	return nil
}

// it searches the creation transaction for the given contract address using binary search in complexity of o(log n)
func (s *OpenSeaScraper) findContractCreationInfo(ctx context.Context, contractAddr common.Address) (createdBy common.Address, createdAt time.Time, err error) {
	var (
		lo, hi, blockNum uint64
		code             []byte
		receipt          *types.Receipt
		chainID          *big.Int
		block            *types.Block
	)

	hi, err = s.tradeScraper.ethConnection.BlockNumber(ctx)
	if err != nil {
		return
	}

	for lo <= hi {
		blockNum = (lo + hi) / 2

		code, err = s.tradeScraper.ethConnection.CodeAt(ctx, contractAddr, new(big.Int).SetUint64(blockNum))
		if err != nil {
			return
		}

		if len(code) == 0 {
			lo = blockNum
		} else {
			hi = blockNum
		}

		if hi == lo+1 {
			blockNum = hi
			break
		}
	}

	block, err = s.tradeScraper.ethConnection.BlockByNumber(ctx, new(big.Int).SetUint64(blockNum))
	if err != nil {
		return
	}

	chainID, err = s.tradeScraper.ethConnection.NetworkID(ctx)
	if err != nil {
		return
	}

	signer := types.NewEIP155Signer(chainID)

	for _, trx := range block.Transactions() {
		// recipient must be nill for contract creation transactions
		if trx.To() != nil {
			continue
		}

		receipt, err = s.tradeScraper.ethConnection.TransactionReceipt(ctx, trx.Hash())
		if err != nil {
			return
		}

		// note that if the nft was created by another smart contract
		// we can't find its creation info with this method
		if bytes.Equal(receipt.ContractAddress.Bytes(), contractAddr.Bytes()) {
			createdAt = time.Unix(int64(block.Time()), 0).UTC()
			createdBy, err = types.Sender(signer, trx)
			if err != nil {
				return
			}

			return
		}
	}

	return
}

func (s *OpenSeaScraper) readNFTAttr(ctx context.Context, uri string, timeout time.Duration, maxKB int) (map[string]interface{}, error) {
	attrs := make(map[string]interface{})

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New("unable to read token attributes: " + resp.Status)
	}

	if err := json.NewDecoder(io.LimitReader(resp.Body, int64(maxKB*1024))).Decode(&attrs); err != nil {
		return nil, err
	}

	return attrs, nil
}

// it finds the transfer events of ERC721 in the given transaction
func (s *OpenSeaScraper) findERC721Transfers(ctx context.Context, txHash common.Hash) ([]*erc721Transfer, error) {
	receipt, err := s.tradeScraper.ethConnection.TransactionReceipt(ctx, txHash)
	if err != nil {
		log.Warnf("unable to fetch the receipt of transaction(%s): %s", txHash.Hex(), err.Error())
		return nil, err
	}

	transfers := make([]*erc721Transfer, 0, 1)

	callOpts := &bind.CallOpts{
		Context: ctx,
	}

	if openSeaUseArchiveNodeFeatures {
		callOpts.BlockNumber = receipt.BlockNumber
	}

	for _, txLog := range receipt.Logs {
		nftContract, err := erc721.NewERC721(txLog.Address, s.tradeScraper.ethConnection)
		if err != nil {
			log.Warnf("unable to bind erc721 contract at address %s: %s", txLog.Address.Hex(), err.Error())
			continue
		}

		transferLog, err := nftContract.ParseTransfer(*txLog)
		if err != nil {
			// it means this log data not comply to erc721's transfer event
			//
			// some old erc721 contracts have unindexed tokenid parameter
			// so it is not compliant with the eip-721. we ignore this kind of trades
			log.Tracef("the event cannot comply to erc721's transfer: %s", err)
			continue
		}

		transfer := &erc721Transfer{
			NFTAddress: txLog.Address,
			From:       transferLog.From,
			To:         transferLog.To,
			TokenID:    transferLog.TokenId,
			TokenAttrs: make(map[string]interface{}),
		}

		if md, err := erc721.NewERC721Metadata(txLog.Address, s.tradeScraper.ethConnection); err != nil {
			log.Warnf("unable to bind erc721 metadata contract at address %s: %s", txLog.Address.Hex(), err.Error())
		} else {
			if nftName, err := md.Name(callOpts); err != nil {
				log.Warnf("unable to read nft name from metadata interface of nft(addr: %s): %s", txLog.Address.Hex(), err.Error())
			} else {
				transfer.Name = &nftName
			}

			if nftSymbol, err := md.Symbol(callOpts); err != nil {
				log.Warnf("unable to read nft symbol from metadata interface of nft(addr: %s): %s", txLog.Address.Hex(), err.Error())
			} else {
				transfer.Symbol = &nftSymbol
			}

			if tokenURI, err := md.TokenURI(callOpts, transfer.TokenID); err != nil {
				log.Warnf("unable to find token(%s) uri: %s", transfer.TokenID.String(), err.Error())
			} else if attrs, err := s.readNFTAttr(ctx, tokenURI, openSeaFetchAttrsTimeout, openSeaFetchAttrsMaxSizeInKB); err != nil {
				log.Warnf("unable to read token(%s) attributes: %s", transfer.TokenID.String(), err.Error())
			} else {
				transfer.TokenURI = &tokenURI
				transfer.TokenAttrs = attrs
			}
		}

		if openSeaReadTotalSupply {
			if enum, err := erc721.NewERC721Enumerable(txLog.Address, s.tradeScraper.ethConnection); err != nil {
				log.Warnf("unable to bind erc721 enumerable contract at address %s: %s", txLog.Address.Hex(), err.Error())
			} else {
				if totalSupply, err := enum.TotalSupply(callOpts); err != nil {
					log.Warnf("unable to read nft total supply from enumerable interface of nft(addr: %s): %s", txLog.Address.Hex(), err.Error())
				} else {
					transfer.TotalSupply = totalSupply
				}
			}
		}

		transfers = append(transfers, transfer)
	}

	return transfers, nil
}
