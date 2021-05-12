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
	refreshDelay = time.Second * 60 * 30
)

type SorareScraper struct {
	tradescraper    TradeScraper
	contractAddress common.Address
	ticker          *time.Ticker
	lastBlockNumber *big.Int
}

func NewSorareScraper(rdb *models.RelDB) *SorareScraper {
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
	}
	s := &SorareScraper{
		contractAddress: common.HexToAddress("0x629A673A8242c2AC4B7B8C5D8735fbeac21A6205"),
		tradescraper:    tradeScraper,
		ticker:          time.NewTicker(refreshDelay),
	}

	go s.mainLoop()
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *SorareScraper) mainLoop() {
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

func (scraper *SorareScraper) UpdateTrades() error {
	trades, err := scraper.FetchTrades()
	if err != nil {
		return err
	}
	for _, trade := range trades {
		scraper.GetTradeChannel() <- &trade
	}
	return nil
}

func (scraper *SorareScraper) FetchTrades() (trades []dia.NFTTrade, err error) {
	// TO DO: Fetch all trades/transactions since scraper.lastBlockNumber from on-chain.
	// If scraper.lastBlockNumber == 0 fetch last blockNumber from postgres using:
	// scraper.tradescraper.datastore.GetLastBlockNFTTrade(nft dia.NFT)

	// The field nftTrade.NFT can be fetched from postgres through scraper.tradescraper.datastore.GetNFT()
	//
	return []dia.NFTTrade{}, nil
}

// GetTotalSupply returns the total supply of the NFT from on-chain.
func (scraper *SorareScraper) GetTotalSupply() (*big.Int, error) {
	contract, err := sorare.NewSorareTokensCaller(scraper.contractAddress, scraper.tradescraper.ethConnection)
	if err != nil {
		fmt.Println("error getting contract: ", err)
	}
	return contract.TotalSupply(&bind.CallOpts{})
}

// GetDataChannel returns the scrapers data channel.
func (scraper *SorareScraper) GetTradeChannel() chan *dia.NFTTrade {
	return scraper.tradescraper.chanTrade
}

// closes all connected Scrapers. Must only be called from mainLoop
func (scraper *SorareScraper) cleanup(err error) {
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
func (scraper *SorareScraper) Close() error {
	if scraper.tradescraper.closed {
		return errors.New("scraper already closed")
	}
	close(scraper.tradescraper.shutdown)
	<-scraper.tradescraper.shutdownDone
	scraper.tradescraper.errorLock.RLock()
	defer scraper.tradescraper.errorLock.RUnlock()
	return scraper.tradescraper.error
}
