package nftbidscrapers

// Please implement the scraping of coingecko quotations here.

import (
	"context"
	"errors"
	"fmt"
	"math"
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
	CryptoPunkRefreshDelay = time.Second * 60
)

type CryptoPunkScraper struct {
	bidScraper      BidScraper
	contractAddress common.Address
	ticker          *time.Ticker
	lastBlockNumber *big.Int
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

	bidScraper := BidScraper{
		shutdown:      make(chan nothing),
		shutdownDone:  make(chan nothing),
		error:         nil,
		ethConnection: connection,
		datastore:     rdb,
		chanBid:       make(chan dia.NFTBid),
	}
	s := &CryptoPunkScraper{
		contractAddress: common.HexToAddress("0xb47e3cd837dDF8e4c57F05d70Ab865de6e193BBB"),
		bidScraper:      bidScraper,
		ticker:          time.NewTicker(CryptoPunkRefreshDelay),
	}
	fmt.Println("scraper built. Start main loop.")
	go s.mainLoop()
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *CryptoPunkScraper) mainLoop() {
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

func (scraper *CryptoPunkScraper) FetchBids() error {
	log.Info("fetch bids...")

	var err error
	if scraper.lastBlockNumber == nil || scraper.lastBlockNumber.Uint64() == 0 {
		// TODO: what is the required value to the GetLastBlockNFTTrade method?
		scraper.lastBlockNumber, err = scraper.bidScraper.datastore.GetLastBlockNFTTrade(dia.NFT{})
		if err != nil {
			// We couldn't find a last block number, fallback to CryptoPunks first block number!
			scraper.lastBlockNumber = big.NewInt(3919706)
		}
	}
	// scraper.lastBlockNumber = big.NewInt(12653867)
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
			Start: scraper.lastBlockNumber.Uint64(),
			End:   &endBlockNumber,
		}, nil, nil)
		if err != nil {
			if err.Error() == "query returned more than 10000 results" {
				fmt.Println("Got `query returned more than 10000 results` error, reduce the window size and try again...")
				endBlockNumber = scraper.lastBlockNumber.Uint64() + (endBlockNumber-scraper.lastBlockNumber.Uint64())/2
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
			value, _ := new(big.Float).Quo(new(big.Float).SetInt(iterBid.Event.Value), big.NewFloat(math.Pow10(18))).Float64()
			bid := dia.NFTBid{
				NFT: dia.NFT{
					NFTClass: nftclass,
					TokenID:  iterBid.Event.PunkIndex.String(),
				},
				Value: value,
				// TO DO: Switch to asset once deployed on IBM
				Currency:      "ETH",
				FromAddress:   iterBid.Event.FromAddress,
				TxHash:        iterBid.Event.Raw.TxHash,
				BlockNumber:   iterBid.Event.Raw.BlockNumber,
				BlockPosition: iterBid.Event.Raw.Index,
				Time:          time.Unix(int64(header.Time), 0),
				Exchange:      "CryptopunkMarket",
			}
			fmt.Println("got bid: ", bid)
			scraper.GetBidChannel() <- bid
		}
		// ---------------------------------------------------------------------------------------------
		break
	}

	// Update the last lastBlockNumber value.
	scraper.lastBlockNumber = new(big.Int).SetUint64(endBlockNumber)
	return nil
}

// GetDataChannel returns the scrapers data channel.
func (scraper *CryptoPunkScraper) GetBidChannel() chan dia.NFTBid {
	return scraper.bidScraper.chanBid
}

// closes all connected Scrapers. Must only be called from mainLoop
func (scraper *CryptoPunkScraper) cleanup(err error) {
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
func (scraper *CryptoPunkScraper) Close() error {
	if scraper.bidScraper.closed {
		return errors.New("scraper already closed")
	}
	close(scraper.bidScraper.shutdown)
	<-scraper.bidScraper.shutdownDone
	scraper.bidScraper.errorLock.RLock()
	defer scraper.bidScraper.errorLock.RUnlock()
	return scraper.bidScraper.error
}
