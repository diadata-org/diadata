package nfttradescrapers

// Please implement the scraping of coingecko quotations here.

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/diadata-org/diadata/config/nftContracts/cryptokitties"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"

	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	CryptoKittiesRefreshDelay = time.Second * 60 * 10
	cryptoKittiesFirstBlock   = 4605169
)

var (
	assetCacheCryptokitties = make(map[string]dia.Asset)
)

type CryptoKittiesScraper struct {
	tradescraper    TradeScraper
	contractAddress common.Address
	ticker          *time.Ticker
	lastBlockNumber uint64
}

func NewCryptoKittiesScraper(rdb *models.RelDB) *CryptoKittiesScraper {
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
		case <-scraper.tradescraper.shutdown:
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
		scraper.lastBlockNumber, err = scraper.tradescraper.datastore.GetLastBlockNFTTrade(dia.NFTClass{
			Address:    scraper.contractAddress.Hex(),
			Blockchain: dia.ETHEREUM,
		})
		if err != nil {
			// We couldn't find a last block number, fallback to Cryptokitties first block number!
			scraper.lastBlockNumber = uint64(cryptoKittiesFirstBlock)
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

	// It's a good practise to stay a little behind the head.
	endBlockNumber := header.Number.Uint64() - blockDelayEthereum

	// Reduce the window size while there is an query limit error.
	for {
		fmt.Println("lastBlockNumber, endBlockNumber: ", scraper.lastBlockNumber, endBlockNumber)
		// We're interested in the FilterAuctionSuccessful events when actual auctions happened!
		iter, err := filterer.FilterAuctionSuccessful(&bind.FilterOpts{
			Start: scraper.lastBlockNumber,
			End:   &endBlockNumber,
		})
		if err != nil {
			if strings.Contains(err.Error(), "query returned more than 10000 results") || strings.Contains(err.Error(), "Log response size exceeded") {
				fmt.Println("Got `query returned more than 10000 results` error, reduce the window size and try again...")
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
			fmt.Println("error filtering FilterAuctionSuccessful: ", err)
			return err
		}

		// Iterate over FilterAuctionSuccessful events.
		for iter.Next() {
			currHeader, err := scraper.tradescraper.ethConnection.HeaderByNumber(context.Background(), big.NewInt(int64(iter.Event.Raw.BlockNumber)))
			if err != nil {
				log.Error("could not fetch current block header: ", err)
			}

			nft, err := scraper.tradescraper.datastore.GetNFT(scraper.contractAddress.Hex(), dia.ETHEREUM, iter.Event.TokenId.String())
			if err != nil {
				// TODO: should we continue if we failed to get NFT from the db or should we fail!
				// continue
				return err
			}

			lastOffer, err := scraper.tradescraper.datastore.GetLastNFTOffer(nft.NFTClass.Address, nft.NFTClass.Blockchain, nft.TokenID, iter.Event.Raw.BlockNumber, iter.Event.Raw.Index)
			if err != nil {
				return err
			}

			price := float64(iter.Event.TotalPrice.Uint64())

			trade := dia.NFTTrade{
				NFT:         nft,
				Price:       iter.Event.TotalPrice,
				FromAddress: lastOffer.FromAddress,
				ToAddress:   iter.Event.Winner.Hex(),
				BlockNumber: iter.Event.Raw.BlockNumber,
				Timestamp:   time.Unix(int64(currHeader.Time), 0),
				Exchange:    "CryptokittiesAuction",
				TxHash:      iter.Event.Raw.TxHash.Hex(),
			}
			if asset, ok := assetCacheCryptokitties[dia.ETHEREUM+"-"+"0x0000000000000000000000000000000000000000"]; ok {
				trade.Currency = asset
			} else {
				currency, err := scraper.tradescraper.datastore.GetAsset("0x0000000000000000000000000000000000000000", dia.ETHEREUM)
				if err != nil {
					log.Errorf("cannot fetch asset %s -- %s", dia.ETHEREUM, "0x0000000000000000000000000000000000000000")
				}
				trade.Currency = currency
				assetCacheCryptokitties[dia.ETHEREUM+"-"+"0x0000000000000000000000000000000000000000"] = currency
			}

			scraper.GetTradeChannel() <- trade

			log.Infof("got trade: ")
			log.Infof("iter: %v", iter)
			log.Info("price: ", price)
			log.Info("from address: ", trade.FromAddress)
			log.Info("to address: ", trade.ToAddress)
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
