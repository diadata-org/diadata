package nftdatascrapers

// Please implement the scraping of coingecko quotations here.

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/diadata-org/diadata/config/nftContracts/sorare"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

const (
	source       = "Sorare"
	refreshDelay = time.Second * 60 * 5
)

type nothing struct{}

type SorareScraper struct {
	nftscraper    NFTScraper
	address       common.Address
	apiURLOpensea string
	ticker        *time.Ticker
}

func NewSorareScraper(rdb *models.RelDB) *SorareScraper {
	connection, err := ethhelper.NewETHClient()
	if err != nil {
		log.Error("Error connecting Eth Client")
	}

	nftScraper := NFTScraper{
		shutdown:      make(chan nothing),
		shutdownDone:  make(chan nothing),
		error:         nil,
		ethConnection: connection,
		relDB:         rdb,
		chanData:      make(chan *dia.NFT),
	}
	s := &SorareScraper{
		address:       common.HexToAddress("0x629A673A8242c2AC4B7B8C5D8735fbeac21A6205"),
		apiURLOpensea: "https://api.opensea.io/api/v1/",
		nftscraper:    nftScraper,
		ticker:        time.NewTicker(refreshDelay),
	}

	go s.mainLoop()
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *SorareScraper) mainLoop() {
	for true {
		select {
		case <-scraper.ticker.C:
			scraper.UpdateNFT()
		case <-scraper.nftscraper.shutdown: // user requested shutdown
			log.Printf("Sorare scraper shutting down")
			err := scraper.Close()
			scraper.cleanup(err)
			return
		}
	}
}

func (scraper *SorareScraper) UpdateNFT() error {
	nfts, err := scraper.FetchData()
	if err != nil {
		return err
	}
	for _, nft := range nfts {
		scraper.GetDataChannel() <- &nft
	}
	return nil
}

func (scraper *SorareScraper) FetchData() (nfts []dia.NFT, err error) {
	// TO DO: iterate over all NFT IDs and fill []dia.NFT
	totalSupply, err := scraper.GetTotalSupply()
	if err != nil {
		return
	}
	for i := 0; i < int(totalSupply.Int64()); i++ {
		// 1. fetch data from onchain
		// 2. fetch data from offchain
		// 3. combine both in order to fill dia.NFT
	}
	return []dia.NFT{}, nil
}

// GetTotalSupply returns the total supply of the NFT from on-chain.
func (scraper *SorareScraper) GetTotalSupply() (*big.Int, error) {
	contract, err := sorare.NewSorareTokensCaller(scraper.address, scraper.nftscraper.ethConnection)
	if err != nil {
		fmt.Println("error getting contract: ", err)
	}
	return contract.TotalSupply(&bind.CallOpts{})
}

// GetDataChannel returns the scrapers data channel.
func (scraper *SorareScraper) GetDataChannel() chan *dia.NFT {
	return scraper.nftscraper.chanData
}

// closes all connected Scrapers. Must only be called from mainLoop
func (scraper *SorareScraper) cleanup(err error) {
	scraper.nftscraper.errorLock.Lock()
	defer scraper.nftscraper.errorLock.Unlock()
	scraper.ticker.Stop()
	if err != nil {
		scraper.nftscraper.error = err
	}
	scraper.nftscraper.closed = true
	close(scraper.nftscraper.shutdownDone) // signal that shutdown is complete
}

// Close closes any existing API connections
func (scraper *SorareScraper) Close() error {
	if scraper.nftscraper.closed {
		return errors.New("scraper already closed")
	}
	close(scraper.nftscraper.shutdown)
	<-scraper.nftscraper.shutdownDone
	scraper.nftscraper.errorLock.RLock()
	defer scraper.nftscraper.errorLock.RUnlock()
	return scraper.nftscraper.error
}
