package nfttradescrapers

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
)

const (
	source       = "Opensea"
	refreshDelay = time.Second * 60 * 30
)

type OpenseaScraper struct {
	tradescraper             TradeScraper
	openseaStorefrontAddress common.Address
	ticker                   *time.Ticker
	lastBlockNumber          *big.Int
}

func NewOpenseaScraper(rdb *models.RelDB) *OpenseaScraper {
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
		chanTrade:     make(chan *dia.NFTTrade),
		exchange:      source,
	}
	s := &OpenseaScraper{
		openseaStorefrontAddress: common.HexToAddress("0x495f947276749Ce646f68AC8c248420045cb7b5e"),
		tradescraper:             tradeScraper,
		ticker:                   time.NewTicker(refreshDelay),
	}

	go s.mainLoop()
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *OpenseaScraper) mainLoop() {
	for true {
		select {
		case <-scraper.ticker.C:
			scraper.UpdateTrades()
		case <-scraper.tradescraper.shutdown: // user requested shutdown
			log.Printf("Sorare scraper shutting down")
			err := scraper.Close()
			scraper.cleanup(err)
			return
		}
	}
}

func (scraper *OpenseaScraper) UpdateTrades() error {
	trades, err := scraper.FetchTrades()
	if err != nil {
		return err
	}
	for _, trade := range trades {
		scraper.GetTradeChannel() <- &trade
	}
	return nil
}

func (scraper *OpenseaScraper) FetchTrades() (trades []dia.NFTTrade, err error) {
	// TO DO: Fetch all trades/transactions since scraper.lastBlockNumber from on-chain
	// The field nftTrade.NFT has to be fetched from postgres through scraper.tradescraper.datastore.GetNFT()
	return []dia.NFTTrade{}, nil
}

// GetTotalSupply returns the total supply of the NFT from on-chain.
func (scraper *OpenseaScraper) GetTotalSupply() (*big.Int, error) {
	contract, err := sorare.NewSorareTokensCaller(scraper.openseaStorefrontAddress, scraper.tradescraper.ethConnection)
	if err != nil {
		fmt.Println("error getting contract: ", err)
	}
	return contract.TotalSupply(&bind.CallOpts{})
}

// GetDataChannel returns the scrapers data channel.
func (scraper *OpenseaScraper) GetTradeChannel() chan *dia.NFTTrade {
	return scraper.tradescraper.chanTrade
}

// closes all connected Scrapers. Must only be called from mainLoop
func (scraper *OpenseaScraper) cleanup(err error) {
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
func (scraper *OpenseaScraper) Close() error {
	if scraper.tradescraper.closed {
		return errors.New("scraper already closed")
	}
	close(scraper.tradescraper.shutdown)
	<-scraper.tradescraper.shutdownDone
	scraper.tradescraper.errorLock.RLock()
	defer scraper.tradescraper.errorLock.RUnlock()
	return scraper.tradescraper.error
}
