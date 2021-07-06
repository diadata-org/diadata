package nftbidscrapers

// Please implement the scraping of coingecko quotations here.

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/diadata-org/diadata/config/nftContracts/cryptopunk"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"

	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

const (
	CryptoPunkRefreshDelay = time.Minute * 60 * 2
)

type CryptoPunksScraper struct {
	bidScraper      BidScraper
	contractAddress common.Address
	ticker          *time.Ticker
	lastBlockNumber uint64
}

func NewCryptoPunksScraper(rdb *models.RelDB) *CryptoPunksScraper {
	connection, err := ethhelper.NewETHClient()
	if err != nil {
		log.Error("Error connecting Eth Client")
	}
	// connection, err := ethclient.Dial("node url")
	// if err != nil {
	// 	log.Error("Error connecting Eth Client")
	// }

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
		// TODO: what is the required value to the GetLastBlockNFTTrade method?
		scraper.lastBlockNumber, err = scraper.bidScraper.datastore.GetLastBlockNFTBid(dia.NFTClass{
			Address:    scraper.contractAddress.Hex(),
			Blockchain: dia.ETHEREUM,
		})
		if err != nil {
			// We couldn't find a last block number, fallback to CryptoPunks first block number!
			scraper.lastBlockNumber = uint64(3918000)
		}
	}
	scraper.lastBlockNumber = uint64(12453867)
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
	// endBlockNumber = 12653869

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
			if err.Error() == "query returned more than 10000 results" {
				fmt.Println("Got `query returned more than 10000 results` error, reduce the window size and try again...")
				endBlockNumber = scraper.lastBlockNumber + (endBlockNumber-scraper.lastBlockNumber)/2
				continue
			}
			fmt.Println("error filtering FilterPunkBought: ", err)
			return err
		}

		for iterBid.Next() {
			// Get block time.
			header, err := scraper.bidScraper.ethConnection.HeaderByNumber(context.Background(), big.NewInt(int64(iterBid.Event.Raw.BlockNumber)))
			if err != nil {
				return err
			}
			bid := dia.NFTBid{
				NFT: dia.NFT{
					NFTClass: nftclass,
					TokenID:  iterBid.Event.PunkIndex.String(),
				},
				Value:       iterBid.Event.Value,
				FromAddress: iterBid.Event.FromAddress.Hex(),
				// TO DO: Switch to asset once deployed on IBM
				CurrencySymbol:   "WETH",
				CurrencyAddress:  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
				CurrencyDecimals: int32(18),

				BlockNumber:   iterBid.Event.Raw.BlockNumber,
				BlockPosition: uint64(iterBid.Event.Raw.Index),
				Timestamp:     time.Unix(int64(header.Time), 0),
				TxHash:        iterBid.Event.Raw.TxHash.Hex(),
				Exchange:      "CryptopunkMarket",
			}
			fmt.Printf("got bid at time %v: %v\n", bid.Timestamp, bid)
			scraper.GetBidChannel() <- bid
		}
		// ---------------------------------------------------------------------------------------------
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
