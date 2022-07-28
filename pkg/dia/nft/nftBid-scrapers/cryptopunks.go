package nftbidscrapers

// Please implement the scraping of coingecko quotations here.

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/diadata-org/diadata/config/nftContracts/cryptopunk"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	"github.com/diadata-org/diadata/pkg/utils"

	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	CryptoPunkRefreshDelay = time.Minute * 60 * 10
	cryptoPunksFirstBlock  = 3918000
)

type CryptoPunksScraper struct {
	bidScraper      BidScraper
	contractAddress common.Address
	ticker          *time.Ticker
	lastBlockNumber uint64
}

func NewCryptoPunksScraper(rdb *models.RelDB) *CryptoPunksScraper {

	connection, err := ethclient.Dial(utils.Getenv("ETH_URI_REST", ""))
	if err != nil {
		log.Error("Error connecting Eth Client")
	}

	bidScraper := BidScraper{
		shutdown:      make(chan nothing),
		shutdownDone:  make(chan nothing),
		error:         nil,
		ethConnection: connection,
		datastore:     rdb,
		chanBid:       make(chan dia.NFTBid),
	}
	s := &CryptoPunksScraper{
		contractAddress: common.HexToAddress("0xb47e3cd837dDF8e4c57F05d70Ab865de6e193BBB"),
		bidScraper:      bidScraper,
		ticker:          time.NewTicker(CryptoPunkRefreshDelay),
	}
	fmt.Println("scraper built. Start main loop.")
	go s.mainLoop()
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *CryptoPunksScraper) mainLoop() {
	err := scraper.FetchBids()
	if err != nil {
		log.Fatal("fetching bids: ", err)
	}
	for {
		select {
		case <-scraper.ticker.C:
			err := scraper.FetchBids()
			if err != nil {
				log.Fatal("fetching bids: ", err)
			}
		case <-scraper.bidScraper.shutdown: // user requested shutdown
			log.Printf("CryptoPunk scraper shutting down")
			err := scraper.Close()
			scraper.cleanup(err)
			return
		}
	}
}

func (scraper *CryptoPunksScraper) FetchBids() error {
	log.Info("fetch bids...")

	var err error
	if scraper.lastBlockNumber == 0 {
		scraper.lastBlockNumber, err = scraper.bidScraper.datastore.GetLastBlockNFTBid(dia.NFTClass{
			Address:    scraper.contractAddress.Hex(),
			Blockchain: dia.ETHEREUM,
		})
		if err != nil {
			// We couldn't find a last block number, fallback to CryptoPunks first block number!
			log.Warn("could not find last block in postgres: ", err)
			scraper.lastBlockNumber = uint64(cryptoPunksFirstBlock)
		}
	}
	filterer, err := cryptopunk.NewCryptoPunksMarketFilterer(scraper.contractAddress, scraper.bidScraper.ethConnection)
	if err != nil {
		return err
	}

	// Get latest block number.
	header, err := scraper.bidScraper.ethConnection.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return err
	}

	// It's a good practise to stay a little behind the head.
	endBlockNumber := header.Number.Uint64() - 18

	nftclass, err := scraper.bidScraper.datastore.GetNFTClass(scraper.contractAddress.Hex(), dia.ETHEREUM)
	if err != nil {
		log.Error("fetching cryptopunk nft class: ", err)
	}
	// Reduce the window size while there is a query limit error.
	for {

		iterBid, err := filterer.FilterPunkBidEntered(&bind.FilterOpts{
			Start: scraper.lastBlockNumber,
			End:   &endBlockNumber,
		}, nil, nil)
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

		for iterBid.Next() {

			// Get block time.
			timestamp, err := ethhelper.GetBlockTimeEth(int64(iterBid.Event.Raw.BlockNumber), scraper.bidScraper.datastore, scraper.bidScraper.ethConnection)
			if err != nil {
				log.Errorf("getting block time: %+v", err)
			}

			bid := dia.NFTBid{
				NFT: dia.NFT{
					NFTClass: nftclass,
					TokenID:  iterBid.Event.PunkIndex.String(),
				},
				Value:            iterBid.Event.Value,
				FromAddress:      iterBid.Event.FromAddress.Hex(),
				CurrencySymbol:   "ETH",
				CurrencyAddress:  "0x0000000000000000000000000000000000000000",
				CurrencyDecimals: int32(18),
				BlockNumber:      iterBid.Event.Raw.BlockNumber,
				BlockPosition:    uint64(iterBid.Event.Raw.Index),
				Timestamp:        timestamp,
				TxHash:           iterBid.Event.Raw.TxHash.Hex(),
				Exchange:         "CryptopunkMarket",
			}
			log.Infof("got bid at time %v: %v\n", bid.Timestamp, bid)
			scraper.GetBidChannel() <- bid
		}
		break
	}

	// Update the last lastBlockNumber value.
	scraper.lastBlockNumber = endBlockNumber
	return nil
}

// GetDataChannel returns the scrapers data channel.
func (scraper *CryptoPunksScraper) GetBidChannel() chan dia.NFTBid {
	return scraper.bidScraper.chanBid
}

// closes all connected Scrapers. Must only be called from mainLoop
func (scraper *CryptoPunksScraper) cleanup(err error) {
	scraper.bidScraper.errorLock.Lock()
	defer scraper.bidScraper.errorLock.Unlock()
	scraper.ticker.Stop()
	if err != nil {
		scraper.bidScraper.error = err
	}
	scraper.bidScraper.closed = true
	close(scraper.bidScraper.shutdownDone) // signal that shutdown is complete
}

// Close closes any existing API connections
func (scraper *CryptoPunksScraper) Close() error {
	if scraper.bidScraper.closed {
		return errors.New("scraper already closed")
	}
	close(scraper.bidScraper.shutdown)
	<-scraper.bidScraper.shutdownDone
	scraper.bidScraper.errorLock.RLock()
	defer scraper.bidScraper.errorLock.RUnlock()
	return scraper.bidScraper.error
}
