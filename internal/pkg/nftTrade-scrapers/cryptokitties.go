package nfttradescrapers

// Please implement the scraping of coingecko quotations here.

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/diadata-org/diadata/config/nftContracts/cryptokitties"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"

	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

const (
	CryptoKittiesRefreshDelay = time.Second * 60
)

type CryptoKittiesScraper struct {
	tradescraper    TradeScraper
	contractAddress common.Address
	ticker          *time.Ticker
	lastBlockNumber uint64
}

func NewCryptoKittiesScraper(rdb *models.RelDB) *CryptoKittiesScraper {
	connection, err := ethhelper.NewETHClient()
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
	s := &CryptoKittiesScraper{
		contractAddress: common.HexToAddress("0xb1690C08E213a35Ed9bAb7B318DE14420FB57d8C"),
		tradescraper:    tradeScraper,
		ticker:          time.NewTicker(CryptoKittiesRefreshDelay),
	}
	fmt.Println("scraper built. Start main loop.")
	go s.mainLoop()
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *CryptoKittiesScraper) mainLoop() {
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
			log.Printf("Cryptokitties scraper shutting down")
			err := scraper.Close()
			scraper.cleanup(err)
			return
		}
	}
}

func (scraper *CryptoKittiesScraper) FetchTrades() error {
	log.Info("fetch trades...")
	var err error
	if scraper.lastBlockNumber == 0 {
		scraper.lastBlockNumber, err = scraper.tradescraper.datastore.GetLastBlockNFTTradeScraper(dia.NFTClass{
			Address:    scraper.contractAddress.Hex(),
			Blockchain: dia.ETHEREUM,
		})
		if err != nil {
			// We couldn't find a last block number, fallback to Cryptokitties first block number!
			scraper.lastBlockNumber = uint64(4605169)
		}
	}

	filterer, err := cryptokitties.NewSaleClockAuctionFilterer(scraper.contractAddress, scraper.tradescraper.ethConnection)
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

	// Reduce the window size while there is an query limit error.
	for {
		fmt.Println("lastBlockNumber, endBlockNumber: ", scraper.lastBlockNumber, endBlockNumber)
		// We're interested in the FilterAuctionSuccessful events when actual auctions happened!
		iter, err := filterer.FilterAuctionSuccessful(&bind.FilterOpts{
			Start: scraper.lastBlockNumber,
			End:   &endBlockNumber,
		})
		if err != nil {
			if err.Error() == "query returned more than 10000 results" {
				fmt.Println("Got `query returned more than 10000 results` error, reduce the window size and try again...")
				endBlockNumber = scraper.lastBlockNumber + (endBlockNumber-scraper.lastBlockNumber)/2
				continue
			}
			fmt.Println("error filtering FilterAuctionSuccessful: ", err)
			return err
		}

		log.Infof("iter: %v", iter)
		// Iter over FilterAuctionSuccessful events.

		for iter.Next() {
			fmt.Println("iter ")
			if err != nil {
				return err
			}

			time.Sleep(1 * time.Second)

			nft, err := scraper.tradescraper.datastore.GetNFT(scraper.contractAddress.Hex(), dia.ETHEREUM, iter.Event.TokenId.String())
			if err != nil {
				// TODO: should we continue if we failed to get NFT from the db or should we fail!
				// continue
				return err
			}

			price := float64(iter.Event.TotalPrice.Uint64())

			trade := dia.NFTTrade{
				NFT:         nft,
				BlockNumber: iter.Event.Raw.BlockNumber,
				// TO DO: Fix FromAddress using data from CryptoKitties Offers Scraper (WIP)
				FromAddress:      common.HexToAddress("0xb1690C08E213a35Ed9bAb7B318DE14420FB57d8C").Hex(),
				ToAddress:        iter.Event.Winner.Hex(),
				Exchange:         "CryptokittiesAuction",
				TxHash:           iter.Event.Raw.TxHash.Hex(),
				Price:            iter.Event.TotalPrice,
				CurrencySymbol:   "WETH",
				CurrencyDecimals: int32(18),
				CurrencyAddress:  common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2").Hex(),
			}
			scraper.GetTradeChannel() <- trade

			log.Infof("got trade: ")
			log.Infof("iter: %v", iter)
			log.Info("price: ", price)
			log.Info("from address: ", common.HexToAddress("0xb1690C08E213a35Ed9bAb7B318DE14420FB57d8C"))
			log.Info("to address: ", iter.Event.Winner)
			log.Info("tx: ", iter.Event.Raw.TxHash)
			log.Info("blockNumber: ", iter.Event.Raw.BlockNumber)
			log.Info("id: ", iter.Event.TokenId.String())
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
func (scraper *CryptoKittiesScraper) GetTradeChannel() chan dia.NFTTrade {
	return scraper.tradescraper.chanTrade
}

// closes all connected Scrapers. Must only be called from mainLoop
func (scraper *CryptoKittiesScraper) cleanup(err error) {
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
func (scraper *CryptoKittiesScraper) Close() error {
	if scraper.tradescraper.closed {
		return errors.New("scraper already closed")
	}
	close(scraper.tradescraper.shutdown)
	<-scraper.tradescraper.shutdownDone
	scraper.tradescraper.errorLock.RLock()
	defer scraper.tradescraper.errorLock.RUnlock()
	return scraper.tradescraper.error
}
