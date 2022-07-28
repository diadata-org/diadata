package nfttradescrapers

// Please implement the scraping of coingecko quotations here.

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/diadata-org/diadata/config/nftContracts/cryptopunk"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"

	// "github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	CryptoPunkRefreshDelay = time.Second * 60 * 10
	cryptoPunksFirstBlock  = 3918000
)

var (
	assetCacheCryptopunk = make(map[string]dia.Asset)
)

type CryptoPunkScraper struct {
	tradescraper    TradeScraper
	contractAddress common.Address
	ticker          *time.Ticker
	lastBlockNumber uint64
}

func NewCryptoPunkScraper(rdb *models.RelDB) *CryptoPunkScraper {
	connection, err := ethclient.Dial(utils.Getenv("ETH_URI_REST", ""))
	if err != nil {
		log.Error("Error connecting Eth Client")
	}

	tradeScraper := TradeScraper{
		shutdown:      make(chan nothing),
		shutdownDone:  make(chan nothing),
		error:         nil,
		ethConnection: connection,
		datastore:     rdb,
		chanTrade:     make(chan dia.NFTTrade),
	}
	s := &CryptoPunkScraper{
		contractAddress: common.HexToAddress("0xb47e3cd837dDF8e4c57F05d70Ab865de6e193BBB"),
		tradescraper:    tradeScraper,
		ticker:          time.NewTicker(CryptoPunkRefreshDelay),
	}
	fmt.Println("scraper built. Start main loop.")
	go s.mainLoop()
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *CryptoPunkScraper) mainLoop() {
	err := scraper.FetchTrades()
	if err != nil {
		log.Error("fetching nft trades: ", err)
	}
	for {
		select {
		case <-scraper.ticker.C:
			err := scraper.FetchTrades()
			if err != nil {
				log.Error("fetching nft trades: ", err)
			}
		case <-scraper.tradescraper.shutdown: // user requested shutdown
			log.Printf("CryptoPunk scraper shutting down")
			err := scraper.Close()
			scraper.cleanup(err)
			return
		}
	}
}

func (scraper *CryptoPunkScraper) FetchTrades() error {
	log.Info("fetch trades...")
	var err error
	if scraper.lastBlockNumber == 0 {
		scraper.lastBlockNumber, err = scraper.tradescraper.datastore.GetLastBlockNFTTrade(dia.NFTClass{
			Address:    scraper.contractAddress.Hex(),
			Blockchain: dia.ETHEREUM,
		})
		if err != nil {
			// We couldn't find a last block number, fallback to CryptoPunks first block number!
			scraper.lastBlockNumber = uint64(cryptoPunksFirstBlock)
		}
	}

	// scraper.lastBlockNumber = uint64(12453867)
	filterer, err := cryptopunk.NewCryptoPunksMarketFilterer(scraper.contractAddress, scraper.tradescraper.ethConnection)
	if err != nil {
		return err
	}

	// Get latest block number.
	header, err := scraper.tradescraper.ethConnection.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return err
	}

	// It's a good practise to stay a little behind the head.
	endBlockNumber := header.Number.Uint64() - blockDelayEthereum

	// We need the cryptopunk abi to unpack the transfer event.
	abi, err := abi.JSON(strings.NewReader(string(cryptopunk.CryptoPunksMarketABI)))
	if err != nil {
		return err
	}
	var iter *cryptopunk.CryptoPunksMarketPunkBoughtIterator
	// Reduce the window size while there is an query limit error.
	for {
		fmt.Println("lastBlockNumber, endBlockNumber: ", scraper.lastBlockNumber, endBlockNumber)
		// We're interested in the FilterPunkBought events when actual trades happened!
		iter, err = filterer.FilterPunkBought(&bind.FilterOpts{
			Start: scraper.lastBlockNumber,
			End:   &endBlockNumber,
		}, nil, nil, nil)
		if err != nil {
			if strings.Contains(err.Error(), "query returned more than 10000 results") || strings.Contains(err.Error(), "Log response size exceeded") {
				log.Warn("Got `query returned more than 10000 results` error, reduce the window size and try again...")
				endBlockNumber = scraper.lastBlockNumber + (endBlockNumber-scraper.lastBlockNumber)/2
				continue
			}
			if strings.Contains(err.Error(), "502 Bad Gateway") {
				log.Info("Got `502 Bad Gateway` error, reduce the window size and try again...")
				endBlockNumber = scraper.lastBlockNumber + (endBlockNumber-scraper.lastBlockNumber)/2
				continue
			}
			if strings.Contains(err.Error(), "504 Gateway Timeout") {
				log.Info("Got `504 Gateway Timeout` error, reduce the window size and try again...")
				endBlockNumber = scraper.lastBlockNumber + (endBlockNumber-scraper.lastBlockNumber)/2
				continue
			}
			log.Error("error filtering FilterPunkBought: ", err)
			return err
		}

		// Iter over FilterPunkBought events.
		for iter.Next() {
			time.Sleep(1 * time.Second)
			currHeader, err := scraper.tradescraper.ethConnection.HeaderByNumber(context.Background(), big.NewInt(int64(iter.Event.Raw.BlockNumber)))
			if err != nil {
				log.Error("could not fetch current block header: ", err)
			}
			nft, err := scraper.tradescraper.datastore.GetNFT(scraper.contractAddress.Hex(), dia.ETHEREUM, iter.Event.PunkIndex.String())
			if err != nil {
				// TODO: should we continue if we failed to get NFT from the db or should we fail!
				continue
				// 	return nil, err
			}

			// This is a workaround to a Cryptopunks contract bug that leads to an empty ToAddress.
			tx, err := scraper.tradescraper.ethConnection.TransactionReceipt(context.TODO(), iter.Event.Raw.TxHash)
			if err != nil {
				// TODO: should we continue if we failed to get tx or should we fail!
				// continue
				return err
			}
			log.Info("tx: ", tx.TxHash.Hex())

			var transferEvent struct {
				From  common.Address
				To    common.Address
				Value *big.Int
			}
			for _, vLog := range tx.Logs {
				log.Info("log", vLog)
				err := abi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
				if err == nil {
					log.Info("found a Transfer event")
					if len(vLog.Topics) > 2 {
						transferEvent.To = common.BytesToAddress(vLog.Topics[2].Bytes())
						break
					}
				}
			}

			price := iter.Event.Value
			// If acceptBidForPunk is called, get the bid value from the bidding history.
			// TO DO: Check that transaction input is acceptBidForPunk.
			if price.Cmp(big.NewInt(0)) == 0 {
				bid, err := scraper.tradescraper.datastore.GetLastNFTBid(scraper.contractAddress.Hex(), dia.ETHEREUM, iter.Event.PunkIndex.String(), uint64(iter.Event.Raw.BlockNumber), iter.Event.Raw.Index)
				if err != nil {
					log.Error("could not find last bid: ", err)
				}
				log.Info("value is zero. Fetch from bids..")
				log.Info(".. value of bid: ", bid.Value)
				log.Info("block of bid: ", bid.BlockNumber)
				log.Info("position in block: ", bid.BlockPosition)
				log.Info("from address of bid: ", bid.FromAddress)
				log.Info("----------")
				log.Info("block of trade: ", iter.Event.Raw.BlockNumber)
				log.Info("to address of trade: ", transferEvent.To.Hex())
				if transferEvent.To.Hex() == bid.FromAddress {
					price = bid.Value
				} else {
					log.Warn("fromAddress of bid does not coincide with toAddress of trade: .")
				}
			}

			trade := dia.NFTTrade{
				NFT:         nft,
				BlockNumber: iter.Event.Raw.BlockNumber,
				FromAddress: iter.Event.FromAddress.Hex(),
				ToAddress:   transferEvent.To.Hex(),
				Exchange:    "CryptopunkMarket",
				TxHash:      iter.Event.Raw.TxHash.Hex(),
				Price:       price,
				Timestamp:   time.Unix(int64(currHeader.Time), 0),
			}

			if asset, ok := assetCacheCryptopunk[dia.ETHEREUM+"-"+"0x0000000000000000000000000000000000000000"]; ok {
				trade.Currency = asset
			} else {
				currency, err := scraper.tradescraper.datastore.GetAsset("0x0000000000000000000000000000000000000000", dia.ETHEREUM)
				if err != nil {
					log.Errorf("cannot fetch asset %s -- %s", dia.ETHEREUM, "0x0000000000000000000000000000000000000000")
				}
				trade.Currency = currency
				assetCacheCryptopunk[dia.ETHEREUM+"-"+"0x0000000000000000000000000000000000000000"] = currency
			}

			scraper.GetTradeChannel() <- trade

			log.Info("got trade: ")
			log.Info("price: ", price)
			log.Info("from address: ", iter.Event.FromAddress)
			log.Info("to address: ", iter.Event.ToAddress)
			log.Info("tx: ", iter.Event.Raw.TxHash)
			log.Info("id: ", iter.Event.PunkIndex.String())
		}
		break
	}

	// Update the last lastBlockNumber value.
	scraper.lastBlockNumber = endBlockNumber
	return nil
}

// GetDataChannel returns the scrapers data channel.
func (scraper *CryptoPunkScraper) GetTradeChannel() chan dia.NFTTrade {
	return scraper.tradescraper.chanTrade
}

// closes all connected Scrapers. Must only be called from mainLoop
func (scraper *CryptoPunkScraper) cleanup(err error) {
	scraper.tradescraper.errorLock.Lock()
	defer scraper.tradescraper.errorLock.Unlock()
	scraper.ticker.Stop()
	if err != nil {
		scraper.tradescraper.error = err
	}
	scraper.tradescraper.closed = true
	close(scraper.tradescraper.shutdownDone) // signal that shutdown is complete
}

// Close closes any existing API connections
func (scraper *CryptoPunkScraper) Close() error {
	if scraper.tradescraper.closed {
		return errors.New("scraper already closed")
	}
	close(scraper.tradescraper.shutdown)
	<-scraper.tradescraper.shutdownDone
	scraper.tradescraper.errorLock.RLock()
	defer scraper.tradescraper.errorLock.RUnlock()
	return scraper.tradescraper.error
}
