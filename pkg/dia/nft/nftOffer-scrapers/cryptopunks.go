package nftofferscrapers

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
	CryptoPunksRefreshDelay = time.Minute * 60 * 10
	cryptoPunksFirstBlock   = 3918000
)

type CryptoPunksScraper struct {
	offerScraper    OfferScraper
	contractAddress common.Address
	ticker          *time.Ticker
	lastBlockNumber uint64
}

func NewCryptoPunksScraper(rdb *models.RelDB) *CryptoPunksScraper {
	connection, err := ethhelper.NewETHClient()
	if err != nil {
		log.Error("Error connecting Eth Client")
	}

	offerScraper := OfferScraper{
		shutdown:      make(chan nothing),
		shutdownDone:  make(chan nothing),
		error:         nil,
		ethConnection: connection,
		datastore:     rdb,
		chanOffer:     make(chan dia.NFTOffer),
	}
	s := &CryptoPunksScraper{
		contractAddress: common.HexToAddress("0xb47e3cd837dDF8e4c57F05d70Ab865de6e193BBB"),
		offerScraper:    offerScraper,
		ticker:          time.NewTicker(CryptokittiesRefreshDelay),
	}
	fmt.Println("scraper built. Start main loop.")
	go s.mainLoop()
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *CryptoPunksScraper) mainLoop() {
	err := scraper.FetchOffers()
	if err != nil {
		log.Fatal("fetching offers: ", err)
	}
	for {
		select {
		case <-scraper.ticker.C:
			err := scraper.FetchOffers()
			if err != nil {
				log.Fatal("fetching offers: ", err)
			}
		case <-scraper.offerScraper.shutdown: // user requested shutdown
			log.Printf("CryptoPunks scraper shutting down")
			err := scraper.Close()
			scraper.cleanup(err)
			return
		}
	}
}

func (scraper *CryptoPunksScraper) FetchOffers() error {
	log.Info("fetch offers...")

	var err error
	if scraper.lastBlockNumber == 0 {
		scraper.lastBlockNumber, err = scraper.offerScraper.datastore.GetLastBlockNFTTradeScraper(dia.NFTClass{
			Address:    scraper.contractAddress.Hex(),
			Blockchain: dia.ETHEREUM,
		})
		if err != nil {
			// We couldn't find a last block number, fallback to CryptoPunks first block number!
			scraper.lastBlockNumber = uint64(cryptoPunksFirstBlock)
		}
	}

	filterer, err := cryptopunk.NewCryptoPunksMarketFilterer(scraper.contractAddress, scraper.offerScraper.ethConnection)
	if err != nil {
		return err
	}

	// Get latest block number.
	header, err := scraper.offerScraper.ethConnection.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return err
	}

	// It's a good practise to stay a little behind the head.
	endBlockNumber := header.Number.Uint64() - blockDelayEthereum

	scraper.lastBlockNumber = uint64(3918000)

	nftclass, err := scraper.offerScraper.datastore.GetNFTClass(scraper.contractAddress.Hex(), dia.ETHEREUM)
	if err != nil {
		log.Error("fetching cryptopunks nft class: ", err)
	}

	for {
		iter, err := filterer.FilterPunkOffered(&bind.FilterOpts{
			Start: scraper.lastBlockNumber,
			End:   &endBlockNumber,
		}, nil, nil)
		if err != nil {
			if err.Error() == "query returned more than 10000 results" {
				fmt.Println("Got `query returned more than 10000 results` error, reduce the window size and try again...")
				endBlockNumber = scraper.lastBlockNumber + (endBlockNumber-scraper.lastBlockNumber)/2
				continue
			}
			fmt.Println("error filtering FilterPunkOffered: ", err)
			return err
		}

		for iter.Next() {

			header, err := scraper.offerScraper.ethConnection.HeaderByNumber(context.Background(), big.NewInt(int64(iter.Event.Raw.BlockNumber)))
			if err != nil {
				return err
			}

			tx, _, err := scraper.offerScraper.ethConnection.TransactionByHash(context.Background(), iter.Event.Raw.TxHash)
			if err != nil {
				return err
			}
			sender, err := scraper.offerScraper.ethConnection.TransactionSender(context.Background(), tx, iter.Event.Raw.BlockHash, iter.Event.Raw.Index)
			if err != nil {
				return err
			}

			offer := dia.NFTOffer{
				NFT: dia.NFT{
					NFTClass: nftclass,
					TokenID:  iter.Event.PunkIndex.String(),
				},
				StartValue:       iter.Event.MinValue,
				FromAddress:      sender.Hex(),
				AuctionType:      "OfferMinValue",
				CurrencySymbol:   "ETH",
				CurrencyAddress:  "0x0000000000000000000000000000000000000000",
				CurrencyDecimals: int32(18),
				BlockNumber:      iter.Event.Raw.BlockNumber,
				BlockPosition:    uint64(iter.Event.Raw.Index),
				Timestamp:        time.Unix(int64(header.Time), 0),
				TxHash:           iter.Event.Raw.TxHash.Hex(),
				Exchange:         "CrypoPunksMarket",
			}
			fmt.Printf("got offer at time %v\n", offer.Timestamp)
			fmt.Println("punk id: ", offer.NFT.TokenID)
			fmt.Println("startValue: ", offer.StartValue)
			fmt.Println("fromaddress: ", offer.FromAddress)

			scraper.GetOfferChannel() <- offer
		}
		break
	}

	// Update the last lastBlockNumber value.
	scraper.lastBlockNumber = endBlockNumber
	return nil
}

// GetDataChannel returns the scrapers data channel.
func (scraper *CryptoPunksScraper) GetOfferChannel() chan dia.NFTOffer {
	return scraper.offerScraper.chanOffer
}

// closes all connected Scrapers. Must only be called from mainLoop
func (scraper *CryptoPunksScraper) cleanup(err error) {
	scraper.offerScraper.errorLock.Lock()
	defer scraper.offerScraper.errorLock.Unlock()
	scraper.ticker.Stop()
	if err != nil {
		scraper.offerScraper.error = err
	}
	scraper.offerScraper.closed = true
	close(scraper.offerScraper.shutdownDone) // signal that shutdown is complete
}

// Close closes any existing API connections
func (scraper *CryptoPunksScraper) Close() error {
	if scraper.offerScraper.closed {
		return errors.New("scraper already closed")
	}
	close(scraper.offerScraper.shutdown)
	<-scraper.offerScraper.shutdownDone
	scraper.offerScraper.errorLock.RLock()
	defer scraper.offerScraper.errorLock.RUnlock()
	return scraper.offerScraper.error
}
