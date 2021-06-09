package nfttradescrapers

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
	CryptoPunkRefreshDelay = time.Second * 60 * 30
)

type CryptoPunkScraper struct {
	tradescraper    TradeScraper
	contractAddress common.Address
	ticker          *time.Ticker
	lastBlockNumber *big.Int
}

func NewCryptoPunkScraper(rdb *models.RelDB) *CryptoPunkScraper {
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
	s := &CryptoPunkScraper{
		contractAddress: common.HexToAddress("0xb47e3cd837dDF8e4c57F05d70Ab865de6e193BBB"),
		tradescraper:    tradeScraper,
		ticker:          time.NewTicker(CryptoPunkRefreshDelay),
	}

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
	if scraper.lastBlockNumber == nil || scraper.lastBlockNumber.Uint64() == 0 {
		// TODO: what is the required value to the GetLastBlockNFTTrade method?
		scraper.lastBlockNumber, err = scraper.tradescraper.datastore.GetLastBlockNFTTrade(dia.NFT{})
		if err != nil {
			// We couldn't find a last block number, fallback to CryptoPunks first block number!
			scraper.lastBlockNumber = big.NewInt(3919706)
		}
	}
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

	// We're interested in the FilterPunkBought events when actual trades happened!
	iter, err := filterer.FilterPunkBought(&bind.FilterOpts{
		Start: scraper.lastBlockNumber.Uint64(),
		End:   &endBlockNumber,
	}, nil, nil, nil)
	if err != nil {
		fmt.Println("error filtering FilterPunkBought: ", err)
		return nil, err
	}

	// Iter over FilterPunkBought events.
	trades = make([]dia.NFTTrade, 0)
	for iter.Next() {
		// TODO: What value should i use for the blockchain argument?
		nft, err := scraper.tradescraper.datastore.GetNFT(scraper.contractAddress, "ethereum", iter.Event.PunkIndex.String())
		if err != nil {
			// TODO: should we continue if we failed to get NFT from the db or should we fail!
			// continue
			return nil, err
		}
		trades = append(trades, dia.NFTTrade{
			NFT:         nft,
			BlockNumber: big.NewInt(int64(iter.Event.Raw.BlockNumber)),
			// TODO: Event.Value is in ETH value, how we can convert it to a USD value using
			// a internal function?
			PriceUSD:    float64(iter.Event.Value.Uint64()),
			FromAddress: iter.Event.FromAddress,
			ToAddress:   iter.Event.ToAddress,
			// TODO: exchange name? any Idea?
			Exchange: "",
		})

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
