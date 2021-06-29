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
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"

	// "github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

const (
	CryptoPunkRefreshDelay = time.Second * 60
)

type CryptoPunkScraper struct {
	tradescraper    TradeScraper
	contractAddress common.Address
	ticker          *time.Ticker
	lastBlockNumber uint64
}

func NewCryptoPunkScraper(rdb *models.RelDB) *CryptoPunkScraper {
	connection, err := ethhelper.NewETHClient()
	if err != nil {
		log.Error("Error connecting Eth Client")
	}
	// connection, err := ethclient.Dial("node url")
	// if err != nil {
	// 	log.Error("Error connecting Eth Client")
	// }

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
		// TODO: what is the required value to the GetLastBlockNFTTrade method?
		scraper.lastBlockNumber, err = scraper.tradescraper.datastore.GetLastBlockNFTTrade(dia.NFT{})
		if err != nil {
			// We couldn't find a last block number, fallback to CryptoPunks first block number!
			scraper.lastBlockNumber = uint64(3919706)
		}
	}

	scraper.lastBlockNumber = uint64(12453867)
	filterer, err := cryptopunk.NewCryptoPunksMarketFilterer(scraper.contractAddress, scraper.tradescraper.ethConnection)
	if err != nil {
		return err
	}

	// Get latest block number.
	header, err := scraper.tradescraper.ethConnection.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return err
	}

	// TODO: It's a good practise to stay a little behind the head.
	endBlockNumber := header.Number.Uint64() - 18

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
			if err.Error() == "query returned more than 10000 results" {
				fmt.Println("Got `query returned more than 10000 results` error, reduce the window size and try again...")
				endBlockNumber = scraper.lastBlockNumber + (endBlockNumber-scraper.lastBlockNumber)/2
				continue
			}
			fmt.Println("error filtering FilterPunkBought: ", err)
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
					transferEvent.To = common.BytesToAddress(vLog.Topics[2].Bytes())
					break
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
				fmt.Println("value is zero. Fetch from bids..")
				fmt.Println(".. value of bid: ", bid.Value)
				fmt.Println("block of bid: ", bid.BlockNumber)
				fmt.Println("position in block: ", bid.BlockPosition)
				fmt.Println("from address of bid: ", bid.FromAddress)
				fmt.Println("----------")
				fmt.Println("block of trade: ", iter.Event.Raw.BlockNumber)
				fmt.Println("to address of trade: ", transferEvent.To.Hex())
				if transferEvent.To.Hex() == bid.FromAddress {
					price = bid.Value
				} else {
					log.Warn("fromAddress of bid does not coincide with toAddress of trade: .")
				}
			}

			trade := dia.NFTTrade{
				NFT:         nft,
				BlockNumber: iter.Event.Raw.BlockNumber,
				// TODO: Event.Value is in ETH value, how we can convert it to a USD value using
				// a internal function?
				FromAddress:      iter.Event.FromAddress,
				ToAddress:        transferEvent.To,
				Exchange:         "CryptopunkMarket",
				TxHash:           iter.Event.Raw.TxHash,
				Price:            price,
				CurrencySymbol:   "WETH",
				CurrencyDecimals: int32(18),
				CurrencyAddress:  common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
				Timestamp:        time.Unix(int64(currHeader.Time), 0),
			}
			scraper.GetTradeChannel() <- trade

			log.Infof("got trade: ")
			log.Infof("iter: %v", iter)
			log.Info("price: ", price)
			log.Info("from address: ", iter.Event.FromAddress)
			log.Info("to address: ", iter.Event.ToAddress)
			log.Info("to address(from the tx Transfer event): ", transferEvent.To)
			log.Info("tx: ", iter.Event.Raw.TxHash)
			log.Info("blockNumber: ", big.NewInt(int64(iter.Event.Raw.BlockNumber)))
			log.Info("id: ", iter.Event.PunkIndex.String())
			log.Info("input data: ", iter.Event.Raw.Data)
			log.Info("input data string: ", iter.Event.Raw.Data)
			log.Info("-----------------------------------------------")
			log.Info(" ")
			log.Info("-----------------------------------------------")
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
