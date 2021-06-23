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

	// "github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	CryptoPunkRefreshDelay = time.Second * 60
)

type CryptoPunkScraper struct {
	tradescraper    TradeScraper
	contractAddress common.Address
	ticker          *time.Ticker
	lastBlockNumber *big.Int
}

func NewCryptoPunkScraper(rdb *models.RelDB) *CryptoPunkScraper {
	// connection, err := ethhelper.NewETHClient()
	// if err != nil {
	// 	log.Error("Error connecting Eth Client")
	// }
	connection, err := ethclient.Dial("node url")
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
	for true {
		select {
		case <-scraper.ticker.C:
			scraper.UpdateTrades()
		case <-scraper.tradescraper.shutdown: // user requested shutdown
			log.Printf("CryptoPunk scraper shutting down")
			err := scraper.Close()
			scraper.cleanup(err)
			return
		}
	}
}

func (scraper *CryptoPunkScraper) UpdateTrades() error {
	trades, err := scraper.FetchTrades()
	if err != nil {
		return err
	}
	for _, trade := range trades {
		scraper.GetTradeChannel() <- trade
	}
	return nil
}

func (scraper *CryptoPunkScraper) FetchTrades() (trades []dia.NFTTrade, err error) {
	log.Info("fetch trades...")
	if scraper.lastBlockNumber == nil || scraper.lastBlockNumber.Uint64() == 0 {
		// TODO: what is the required value to the GetLastBlockNFTTrade method?
		scraper.lastBlockNumber, err = scraper.tradescraper.datastore.GetLastBlockNFTTrade(dia.NFT{})
		if err != nil {
			// We couldn't find a last block number, fallback to CryptoPunks first block number!
			scraper.lastBlockNumber = big.NewInt(3919706)
		}
	}
	scraper.lastBlockNumber = big.NewInt(12653867)
	filterer, err := cryptopunk.NewCryptoPunksMarketFilterer(scraper.contractAddress, scraper.tradescraper.ethConnection)
	if err != nil {
		return nil, err
	}

	// Get latest block number.
	header, err := scraper.tradescraper.ethConnection.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	// TODO: It's a good practise to stay a litlle behind the head.
	endBlockNumber := header.Number.Uint64() - 18
	endBlockNumber = 12653869
	// We need the cryptopunk abi to unpack the transfer event.
	abi, err := abi.JSON(strings.NewReader(string(cryptopunk.CryptoPunksMarketABI)))
	if err != nil {
		return nil, err
	}
	trades = make([]dia.NFTTrade, 0)
	// Reduce the window size while there is an query limit error.
	for {
		fmt.Println("lastBlockNumber, endBlockNumber: ", scraper.lastBlockNumber.Uint64(), endBlockNumber)
		// We're interested in the FilterPunkBought events when actual trades happened!
		iter, err := filterer.FilterPunkBought(&bind.FilterOpts{
			Start: scraper.lastBlockNumber.Uint64(),
			End:   &endBlockNumber,
		}, nil, nil, nil)

	for iter.Next() {
		fmt.Println("iter ")
		// TODO: What value should i use for the blockchain argument?
		nft, err := scraper.tradescraper.datastore.GetNFT(scraper.contractAddress.Hex(), dia.ETHEREUM, iter.Event.PunkIndex.String())

		if err != nil {
			if err.Error() == "query returned more than 10000 results" {
				fmt.Println("Got `query returned more than 10000 results` error, reduce the window size and try again...")
				endBlockNumber = scraper.lastBlockNumber.Uint64() + (endBlockNumber-scraper.lastBlockNumber.Uint64())/2
				continue
			}
			fmt.Println("error filtering FilterPunkBought: ", err)
			return nil, err
		}

		log.Infof("iter: %v", iter)
		// Iter over FilterPunkBought events.
		for iter.Next() {
			// TODO: What value should i use for the blockchain argument?
			// nft, err := scraper.tradescraper.datastore.GetNFT(scraper.contractAddress, "ethereum", iter.Event.PunkIndex.String())
			// if err != nil {
			// 	// TODO: should we continue if we failed to get NFT from the db or should we fail!
			// 	// continue
			// 	return nil, err
			// }

			// QUICK HACK: this is a workaround to a Cryptopunks contract bug that leads to an empty ToAddress.
			tx, err := scraper.tradescraper.ethConnection.TransactionReceipt(context.TODO(), iter.Event.Raw.TxHash)
			if err != nil {
				// TODO: should we continue if we failed to get tx or should we fail!
				// continue
				return nil, err
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

					log.Info("event: ", transferEvent)
					transferEvent.To = common.BytesToAddress(vLog.Topics[2].Bytes())
					break
				}
			}

			// trades = append(trades, dia.NFTTrade{
			// 	NFT:         nft,
			// 	BlockNumber: big.NewInt(int64(iter.Event.Raw.BlockNumber)),
			// 	// TODO: Event.Value is in ETH value, how we can convert it to a USD value using
			// 	// a internal function?
			// 	PriceUSD:    float64(iter.Event.Value.Uint64()),
			// 	FromAddress: iter.Event.FromAddress,
			// 	ToAddress:   iter.Event.ToAddress,
			// 	// TODO: exchange name? any Idea?
			// 	Exchange: "",
			// })

			log.Infof("got trade: ")
			log.Infof("iter: %v", iter)
			log.Info("price: ", float64(iter.Event.Value.Uint64()))
			log.Info("from address: ", iter.Event.FromAddress)
			log.Info("to address: ", iter.Event.ToAddress)
			log.Info("to address(from the tx Transfer event): ", transferEvent.To)
			log.Info("tx: ", iter.Event.Raw.TxHash)
			log.Info("blockNumber: ", big.NewInt(int64(iter.Event.Raw.BlockNumber)))
			log.Info("id: ", iter.Event.PunkIndex.String())
			log.Info("-----------------------------------------------")
			log.Info(" ")
			log.Info("-----------------------------------------------")
		}
		break
	}

	// Update the last lastBlockNumber value.
	scraper.lastBlockNumber = new(big.Int).SetUint64(endBlockNumber)
	return trades, nil
}

// GetTotalSupply returns the total supply of the NFT from on-chain.
func (scraper *CryptoPunkScraper) GetTotalSupply() (*big.Int, error) {
	contract, err := cryptopunk.NewCryptoPunksMarketCaller(scraper.contractAddress, scraper.tradescraper.ethConnection)
	if err != nil {
		fmt.Println("error getting contract: ", err)
	}
	return contract.TotalSupply(&bind.CallOpts{})
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
