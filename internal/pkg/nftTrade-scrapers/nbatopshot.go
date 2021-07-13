package nfttradescrapers

// Credits to Eric Ren, author of the medium post:
// https://medium.com/@eric.ren_51534/polling-nba-top-shot-p2p-market-purchase-events-from-flow-blockchain-using-flow-go-sdk-3ec80119e75f
// and the underlying github repository:
// https://github.com/rrrkren/topshot-sales

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/client"
	"google.golang.org/grpc"

	"github.com/diadata-org/diadata/pkg/dia"
	flowhelper "github.com/diadata-org/diadata/pkg/dia/helpers/flowhelper"
	models "github.com/diadata-org/diadata/pkg/model"
)

const (
	TopshotAddress = "0x0b2a3299cc857e29"
)

type MomentPurchased cadence.Event

func (evt MomentPurchased) Id() uint64 {
	return uint64(evt.Fields[0].(cadence.UInt64))
}

func (evt MomentPurchased) Price() float64 {
	return float64(evt.Fields[1].(cadence.UFix64).ToGoValue().(uint64)) / 1e8 // ufixed 64 have 8 digits of precision
}

func (evt MomentPurchased) Seller() *flow.Address {
	optionalAddress := (evt.Fields[2]).(cadence.Optional)
	if cadenceAddress, ok := optionalAddress.Value.(cadence.Address); ok {
		sellerAddress := flow.BytesToAddress(cadenceAddress.Bytes())
		return &sellerAddress
	}
	return nil
}

func (evt MomentPurchased) String() string {
	return fmt.Sprintf("moment purchased: momentid: %d, price: %f, seller: %s",
		evt.Id(), evt.Price(), evt.Seller())
}

func GetSaleMomentFromOwnerAtBlock(flowClient *client.Client, blockHeight uint64, ownerAddress flow.Address, momentFlowID uint64) (*SaleMoment, error) {
	getSaleMomentScript := `
		import TopShot from 0x0b2a3299cc857e29
        import Market from 0xc1e4f4f4c4257510
        pub struct SaleMoment {
          pub var id: UInt64
          pub var playId: UInt32
          pub var play: {String: String}
          pub var setId: UInt32
          pub var setName: String
          pub var serialNumber: UInt32
          pub var price: UFix64
          init(moment: &TopShot.NFT, price: UFix64) {
            self.id = moment.id
            self.playId = moment.data.playID
            self.play = TopShot.getPlayMetaData(playID: self.playId)!
            self.setId = moment.data.setID
            self.setName = TopShot.getSetName(setID: self.setId)!
            self.serialNumber = moment.data.serialNumber
            self.price = price
          }
        }
		pub fun main(owner:Address, momentID:UInt64): SaleMoment {
			let acct = getAccount(owner)
            let collectionRef = acct.getCapability(/public/topshotSaleCollection)!.borrow<&{Market.SalePublic}>() ?? panic("Could not borrow capability from public collection")
			return SaleMoment(moment: collectionRef.borrowMoment(id: momentID)!,price: collectionRef.getPrice(tokenID: momentID)!)
		}
`
	res, err := flowClient.ExecuteScriptAtBlockHeight(context.Background(), blockHeight, []byte(getSaleMomentScript), []cadence.Value{
		cadence.BytesToAddress(ownerAddress.Bytes()),
		cadence.UInt64(momentFlowID),
	})
	if err != nil {
		return nil, fmt.Errorf("error fetching sale moment from flow: %w", err)
	}
	saleMoment := SaleMoment(res.(cadence.Struct))
	return &saleMoment, nil
}

type SaleMoment cadence.Struct

func (s SaleMoment) ID() uint64 {
	return uint64(s.Fields[0].(cadence.UInt64))
}

func (s SaleMoment) PlayID() uint32 {
	return uint32(s.Fields[1].(cadence.UInt32))
}

func (s SaleMoment) SetName() string {
	return string(s.Fields[4].(cadence.String))
}

func (s SaleMoment) SetID() uint32 {
	return uint32(s.Fields[3].(cadence.UInt32))
}

func (s SaleMoment) Play() map[string]string {
	dict := s.Fields[2].(cadence.Dictionary)
	res := map[string]string{}
	for _, kv := range dict.Pairs {
		res[string(kv.Key.(cadence.String))] = string(kv.Value.(cadence.String))
	}
	return res
}

func (s SaleMoment) SerialNumber() uint32 {
	return uint32(s.Fields[5].(cadence.UInt32))
}

func (s SaleMoment) String() string {
	playData := s.Play()
	return fmt.Sprintf("saleMoment: serialNumber: %d, setID: %d, setName: %s, playID: %d, playerName: %s",
		s.SerialNumber(), s.SetID(), s.SetName(), s.PlayID(), playData["FullName"])
}

type NBATopshotScraper struct {
	tradescraper    TradeScraper
	contractAddress common.Address
	flowClient      *client.Client
	ticker          *time.Ticker
	lastBlockNumber uint64
}

type Play struct {
	SeriesID   uint32
	SetID      uint32
	PlayID     uint32
	SetName    string
	Attributes map[string]interface{}
}

type Moment struct {
	ID           uint64
	SetID        uint32
	PlayID       uint32
	SerialNumber uint32
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

	tradescraper := TradeScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		error:        nil,
		datastore:    rdb,
		chanTrade:    make(chan dia.NFTTrade),
	}

	s := &NBATopshotScraper{
		contractAddress: common.HexToAddress(TopshotAddress),
		tradescraper:    tradescraper,
		flowClient:      flowClient,
		ticker:          time.NewTicker(refreshDelay),
	}

	go s.mainLoop()
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *NBATopshotScraper) mainLoop() {
	err := scraper.FetchTrades()
	if err != nil {
		log.Error("error updating NFT: ", err)
	}
	for {
		select {
		case <-scraper.ticker.C:
			err := scraper.FetchTrades()
			if err != nil {
				log.Error("error updating NFT: ", err)
			}
		case <-scraper.tradescraper.shutdown: // user requested shutdown
			log.Info("NBA Topshot scraper shutting down")
			err := scraper.Close()
			scraper.cleanup(err)
			return
		}
	}
}

// FetchData returns a slice of all NFTs fetched.
func (scraper *NBATopshotScraper) FetchTrades() (err error) {
	log.Info("fetch trades...")
	latestBlock, err := scraper.flowClient.GetLatestBlock(context.Background(), false)
	fmt.Println("Current Height: ", latestBlock.Height)
	scraper.lastBlockNumber = latestBlock.Height
	if err != nil {
		return err
	}

	blockEvents, err := scraper.flowClient.GetEventsForHeightRange(context.Background(), client.EventRangeQuery{
		Type:        "A.c1e4f4f4c4257510.Market.MomentPurchased",
		StartHeight: latestBlock.Height - 50,
		EndHeight:   latestBlock.Height,
	})

	if err != nil {
		return err
	}

	for _, blockEvent := range blockEvents {
		for _, purchaseEvent := range blockEvent.Events {
			// loop through the Market.MomentPurchased events in this blockEvent
			e := MomentPurchased(purchaseEvent.Value)

			nft, err := scraper.tradescraper.datastore.GetNFT(scraper.contractAddress.Hex(), dia.FLOW, fmt.Sprintf("%d", e.Id()))
			if err != nil {
				// TODO: should we continue if we failed to get NFT from the db or should we fail!
				// continue
				return err
			}

			fmt.Println(e.Price(), e.Seller().Hex(), blockEvent.Height)
			trade := dia.NFTTrade{
				NFT:         nft,
				PriceUSD:    e.Price(),
				FromAddress: e.Seller().Hex(),
				BlockNumber: blockEvent.Height,
				Timestamp:   blockEvent.BlockTimestamp,
				Exchange:    "TopshotMarket",
				TxHash:      purchaseEvent.TransactionID.Hex(),
			}
			scraper.GetTradeChannel() <- trade
			log.Info("Topshot purchase: ")
			log.Info("Price in USD: ", trade.PriceUSD)
			log.Info("From address: ", trade.FromAddress)
			log.Info("Block Number: ", trade.BlockNumber)
			log.Info("---------------------------------------------")
		}
	}

	return nil
}

func (scraper *NBATopshotScraper) GetTradeChannel() chan dia.NFTTrade {
	return scraper.tradescraper.chanTrade
}

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
