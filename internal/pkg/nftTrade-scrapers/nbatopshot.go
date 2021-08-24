package nfttradescrapers

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/flowhelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/client"
	"google.golang.org/grpc"
)

const (
	TopshotAddress    = "0x0b2a3299cc857e29"
	refreshDelayTrade = time.Second * 60 * 10
)

type NBATopshotScraper struct {
	tradescraper TradeScraper
	flowClient   *client.Client
	ticker       *time.Ticker
	address      string
}

func NewNBATopshotScraper(rdb *models.RelDB) *NBATopshotScraper {
	flowClient, err := client.New(flowhelper.FlowAPICurrent, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	err = flowClient.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	tradeScraper := TradeScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		error:        nil,
		datastore:    rdb,
		chanTrade:    make(chan dia.NFTTrade),
	}

	s := &NBATopshotScraper{
		tradescraper: tradeScraper,
		flowClient:   flowClient,
		ticker:       time.NewTicker(refreshDelayTrade),
		address:      TopshotAddress,
	}

	go s.mainLoop()
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *NBATopshotScraper) mainLoop() {
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
			log.Printf("NBATopshot scraper shutting down")
			err := scraper.Close()
			scraper.cleanup(err)
			return
		}
	}
}

func (scraper *NBATopshotScraper) FetchTrades() (err error) {
	var lastBlock uint64
	lastBlock, err = scraper.tradescraper.datastore.GetLastBlockheightTopshot(time.Now())
	if err != nil {
		log.Error("fetch last topshot block: ", err)
	}
	if lastBlock == uint64(0) {
		// No last block in db. Start from genesis block.
		lastBlock = flowhelper.RootHeight1
	}
	log.Info("current height: ", lastBlock)

	// get all recipients from block
	recipientsMap := make(map[string]string)
	allDepositMoments, blocknumbers, err := scraper.GetAllDepositMoments(lastBlock)
	if err != nil {
		return err
	}

	for i, moment := range allDepositMoments {
		e := DepositEvent(moment)
		// GetNFT from token Id
		tokenId := strconv.Itoa(int(e.Id()))
		recipientAddress := e.Recipient().Hex()
		BlockNumber := strconv.Itoa(int(blocknumbers[i]))
		recipientsMap[tokenId+","+BlockNumber] = recipientAddress
	}

	// get all purchased moments from block, add to trade db
	allMomentsPurchased, timestamps, blocknumbers, err := scraper.GetAllMomentsPurchased(lastBlock)
	if err != nil {
		return err
	}

	for i, moment := range allMomentsPurchased {
		e := MomentPurchasedEvent(moment.Value)

		nft, err := scraper.tradescraper.datastore.GetNFT(scraper.address, dia.FLOW, strconv.Itoa(int(e.Id())))
		if err != nil {
			log.Error("fetch NFT: ", err)
			return err
		}

		// GetNFT from token Id
		tokenId := strconv.Itoa(int(e.Id()))
		BlockNumber := strconv.Itoa(int(blocknumbers[i]))
		recipientAddress := recipientsMap[tokenId+","+BlockNumber]

		trade := dia.NFTTrade{
			NFT:              nft,
			BlockNumber:      blocknumbers[i],
			FromAddress:      e.Seller().Hex(),
			ToAddress:        recipientAddress,
			Exchange:         "NBATopshotMarket",
			TxHash:           moment.TransactionID.Hex(),
			PriceUSD:         e.Price(),
			CurrencySymbol:   "USD",
			CurrencyDecimals: int32(18),
			CurrencyAddress:  "",
			Timestamp:        timestamps[i],
		}
		scraper.GetTradeChannel() <- trade

	}

	return nil
}

// ---------------------------------------------------------
// Get Data
// ---------------------------------------------------------

// GetAllMomentsPurchased returns all moments from genesis to the latest block by iterating through
// blocks and looking for MomentPurchased events.
func (scraper *NBATopshotScraper) GetAllMomentsPurchased(startheight uint64) (purchasedMoments []flow.Event, timestamps []time.Time, blocknumbers []uint64, err error) {
	log.Info("Getting purchased moments...")
	latestBlock, err := scraper.flowClient.GetLatestBlock(context.Background(), false)
	if err != nil {
		log.Error(err)
	}

	// Get first interval.
	var currentIndex int
	if startheight > flowhelper.RootHeights[len(flowhelper.RootHeights)-1] {
		currentIndex = len(flowhelper.RootHeights)
	} else {
		for i, root := range flowhelper.RootHeights {
			if startheight < root {
				currentIndex = i
				break
			}
		}
	}

	log.Infof("make flow client at startheight %v: ", startheight)
	log.Infof("currentIndex: %v\n", currentIndex)

	flowClient, err := flowhelper.GetFlowClient(startheight)
	if err != nil {
		return
	}

	for startheight < latestBlock.Height {

		if currentIndex == len(flowhelper.RootHeights) || startheight+flowhelper.RequestLimit < flowhelper.RootHeights[currentIndex] {
			// all blocks within the range of given client.
			m, t, b, err := GetPurchasedMoments(startheight, startheight+flowhelper.RequestLimit, flowClient)
			if err != nil {
				log.Error("getting purchased moments: ", err)
			}
			purchasedMoments = append(purchasedMoments, m...)
			timestamps = append(timestamps, t...)
			blocknumbers = append(blocknumbers, b...)
			startheight += flowhelper.RequestLimit
			fmt.Println("current startheight: ", startheight)
		} else {
			// Reached new block range and thus need new client.
			fmt.Println("reached new block range")
			m, t, b, err := GetPurchasedMoments(startheight, flowhelper.RootHeights[currentIndex]-1, flowClient)
			if err != nil {
				log.Error(err)
			}
			purchasedMoments = append(purchasedMoments, m...)
			timestamps = append(timestamps, t...)
			blocknumbers = append(blocknumbers, b...)

			startheight = flowhelper.RootHeights[currentIndex]
			currentIndex += 1
			flowClient, err = flowhelper.GetFlowClient(startheight)
			if err != nil {
				log.Error(err)
			}
		}
	}
	log.Info("... done getting purchased moments.")
	return
}

// GetPurchasedMoments returns all moments minted between blocks @startheight and @endheight.
// The difference @endheight-@starthight is limited to 250.
// The range @startheight, @endheight must not be spread over more than the given @flowClient.
// https://docs.onflow.org/node-operation/past-sporks/
func GetPurchasedMoments(startheight, endheight uint64, flowClient *client.Client) (purchasedMoments []flow.Event, timestamps []time.Time, blockNumbers []uint64, err error) {

	blockEvents, err := flowClient.GetEventsForHeightRange(context.Background(), client.EventRangeQuery{
		Type:        "A.c1e4f4f4c4257510.Market.MomentPurchased",
		StartHeight: startheight,
		EndHeight:   endheight,
	})
	if err != nil {
		return
	}
	for _, blockEvent := range blockEvents {
		timestamp := blockEvent.BlockTimestamp
		for _, momentPurchasedEvent := range blockEvent.Events {
			fmt.Printf("got purchased moment %v at time %v: \n", momentPurchasedEvent.Value, timestamp)
			timestamps = append(timestamps, timestamp)
			blockNumbers = append(blockNumbers, blockEvent.Height)
			purchasedMoments = append(purchasedMoments, momentPurchasedEvent)
		}
	}
	return
}

// GetAllDepositMoments returns all moments from genesis to the latest block by iterating through
// blocks and looking for Deposit events.
func (scraper *NBATopshotScraper) GetAllDepositMoments(startheight uint64) (depositMoments []cadence.Event, blocknumbers []uint64, err error) {
	log.Info("Getting Deposit moments...")
	latestBlock, err := scraper.flowClient.GetLatestBlock(context.Background(), false)
	if err != nil {
		log.Error(err)
	}

	// Get first interval.
	var currentIndex int
	if startheight > flowhelper.RootHeights[len(flowhelper.RootHeights)-1] {
		currentIndex = len(flowhelper.RootHeights)
	} else {
		for i, root := range flowhelper.RootHeights {
			if startheight < root {
				currentIndex = i
				break
			}
		}
	}

	log.Infof("make flow client at startheight %v: ", startheight)
	log.Infof("currentIndex: %v\n", currentIndex)

	flowClient, err := flowhelper.GetFlowClient(startheight)
	if err != nil {
		return
	}

	for startheight < latestBlock.Height {

		if currentIndex == len(flowhelper.RootHeights) || startheight+flowhelper.RequestLimit < flowhelper.RootHeights[currentIndex] {
			// all blocks within the range of given client.
			m, b, err := GetDepositMoments(startheight, startheight+flowhelper.RequestLimit, flowClient)
			if err != nil {
				log.Error("getting deposit moments: ", err)
			}
			depositMoments = append(depositMoments, m...)
			blocknumbers = append(blocknumbers, b...)
			startheight += flowhelper.RequestLimit
			fmt.Println("current startheight: ", startheight)
		} else {
			// Reached new block range and thus need new client.
			fmt.Println("reached new block range")
			m, b, err := GetDepositMoments(startheight, flowhelper.RootHeights[currentIndex]-1, flowClient)
			if err != nil {
				log.Error(err)
			}
			depositMoments = append(depositMoments, m...)
			blocknumbers = append(blocknumbers, b...)

			startheight = flowhelper.RootHeights[currentIndex]
			currentIndex += 1
			flowClient, err = flowhelper.GetFlowClient(startheight)
			if err != nil {
				log.Error(err)
			}
		}
	}
	log.Info("... done getting deposit moments.")
	return
}

// GetDepositMoments returns all moments minted between blocks @startheight and @endheight.
// The difference @endheight-@starthight is limited to 250.
// The range @startheight, @endheight must not be spread over more than the given @flowClient.
// https://docs.onflow.org/node-operation/past-sporks/
func GetDepositMoments(startheight, endheight uint64, flowClient *client.Client) (depositMoments []cadence.Event, blockNumbers []uint64, err error) {

	blockEvents, err := flowClient.GetEventsForHeightRange(context.Background(), client.EventRangeQuery{
		Type:        "A.0b2a3299cc857e29.TopShot.Deposit",
		StartHeight: startheight,
		EndHeight:   endheight,
	})
	if err != nil {
		return
	}
	for _, blockEvent := range blockEvents {
		timestamp := blockEvent.BlockTimestamp
		for _, momentDepositEvent := range blockEvent.Events {
			fmt.Printf("got deposit moment %v at time %v: \n", momentDepositEvent.Value, timestamp)
			depositMoments = append(depositMoments, momentDepositEvent.Value)
			blockNumbers = append(blockNumbers, blockEvent.Height)
		}
	}
	return
}

// GetDataChannel returns the scrapers data channel.
func (scraper *NBATopshotScraper) GetTradeChannel() chan dia.NFTTrade {
	return scraper.tradescraper.chanTrade
}

// closes all connected Scrapers. Must only be called from mainLoop
func (scraper *NBATopshotScraper) cleanup(err error) {
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
func (scraper *NBATopshotScraper) Close() error {
	if scraper.tradescraper.closed {
		return errors.New("scraper already closed")
	}
	close(scraper.tradescraper.shutdown)
	<-scraper.tradescraper.shutdownDone
	scraper.tradescraper.errorLock.RLock()
	defer scraper.tradescraper.errorLock.RUnlock()
	return scraper.tradescraper.error
}

// ---------------------------------------------------------
// Get Trade Event
// ---------------------------------------------------------

//  pub event Deposit(id: UInt64, to: Address?)
type DepositEvent cadence.Event

// Token Id
func (evt DepositEvent) Id() uint64 {
	return uint64(evt.Fields[0].(cadence.UInt64))
}

// Recipient address
func (evt DepositEvent) Recipient() *flow.Address {
	optionalAddress := (evt.Fields[1]).(cadence.Optional)
	if cadenceAddress, ok := optionalAddress.Value.(cadence.Address); ok {
		recipientAddress := flow.BytesToAddress(cadenceAddress.Bytes())
		return &recipientAddress
	}
	return nil
}

func (evt DepositEvent) String() string {
	return fmt.Sprintf("deposit: momentid: %d, recipient: %s",
		evt.Id(), evt.Recipient())
}

// pub event MomentPurchased(id: UInt64, price: UFix64, seller: Address?)
type MomentPurchasedEvent cadence.Event

func (evt MomentPurchasedEvent) Id() uint64 {
	return uint64(evt.Fields[0].(cadence.UInt64))
}

func (evt MomentPurchasedEvent) Price() float64 {
	return float64(evt.Fields[1].(cadence.UFix64).ToGoValue().(uint64)) / 1e8 // ufixed 64 have 8 digits of precision
}

func (evt MomentPurchasedEvent) Seller() *flow.Address {
	optionalAddress := (evt.Fields[2]).(cadence.Optional)
	if cadenceAddress, ok := optionalAddress.Value.(cadence.Address); ok {
		sellerAddress := flow.BytesToAddress(cadenceAddress.Bytes())
		return &sellerAddress
	}
	return nil
}

func (evt MomentPurchasedEvent) String() string {
	return fmt.Sprintf("moment purchased: momentid: %d, price: %f, seller: %s",
		evt.Id(), evt.Price(), evt.Seller())
}
