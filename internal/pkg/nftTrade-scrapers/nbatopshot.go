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
	tradescraper    TradeScraper
	flowClient      *client.Client
	ticker          *time.Ticker
	lastBlockNumber uint64
	address         string
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
	log.Info("fetch trades...")
	if scraper.lastBlockNumber == 0 {
		nftclass := dia.NFTClass{
			Address:      TopshotAddress,
			Symbol:       "TS",
			Name:         "TopShot",
			Blockchain:   "Flow",
			ContractType: "non-fungible",
			Category:     "Collectibles",
		}

		// fetch latest block
		scraper.lastBlockNumber, err = scraper.tradescraper.datastore.GetLastBlockNFTTradeScraper(nftclass)
		if err != nil {
			// We couldn't find a last block number, fallback to NBATopshot first block number!
			scraper.lastBlockNumber = flowhelper.RootHeight1
		}
	}
	log.Info("current height: ", scraper.lastBlockNumber)

	scraper.flowClient, err = flowhelper.GetFlowClient(scraper.lastBlockNumber)
	if err != nil {
		return
	}

	// fetch block events of topshot Market.MomentPurchased events for the past 1000 blocks
	blockEvents, err := scraper.flowClient.GetEventsForHeightRange(context.Background(), client.EventRangeQuery{
		Type:        "A.c1e4f4f4c4257510.Market.MomentPurchased",
		StartHeight: scraper.lastBlockNumber,
		EndHeight:   scraper.lastBlockNumber + flowhelper.RequestLimit,
	})
	if err != nil {
		log.Error("flowClient.GetEventsForHeightRange: ", err)
	}

	for _, blockEvent := range blockEvents {
		for _, purchaseEvent := range blockEvent.Events {
			// loop through the Market.MomentPurchased events in this blockEvent
			e := MomentPurchasedEvent(purchaseEvent.Value)
			saleMoment, err := GetSaleMomentFromOwnerAtBlock(scraper.flowClient, blockEvent.Height-1, *e.Seller(), e.Id())
			if err != nil {
				log.Error("GetSaleMomentFromOwnerAtBlock: ", err)
			}

			metadata, err := scraper.GetMetadata(uint32(saleMoment.SetID()), uint32(saleMoment.PlayID()))
			if err != nil {
				return err
			}

			metadata["blocknumber"] = blockEvent.Height
			nft, err := scraper.tradescraper.datastore.GetNFT(scraper.address, dia.FLOW, strconv.Itoa(int(e.Id())))
			if err != nil {
				log.Error("fetch NFT: ", err)
				return err
			}
			// nft := dia.NFT{
			// 	NFTClass: dia.NFTClass{
			// 		Address:      TopshotAddress,
			// 		Symbol:       "TS",
			// 		Name:         "TopShot",
			// 		Blockchain:   "Flow",
			// 		ContractType: "non-fungible",
			// 		Category:     "Collectibles",
			// 	},
			// 	TokenID:        strconv.Itoa(int(e.Id())),
			// 	CreationTime:   blockEvent.BlockTimestamp,
			// 	CreatorAddress: "",
			// 	URI:            "not available",
			// 	Attributes:     metadata,
			// }
			trade := dia.NFTTrade{
				NFT:              nft,
				BlockNumber:      blockEvent.Height,
				FromAddress:      e.Seller().Hex(),
				ToAddress:        "",
				Exchange:         "NBATopshotMarket",
				TxHash:           purchaseEvent.TransactionID.Hex(),
				PriceUSD:         e.Price(),
				CurrencySymbol:   "USD",
				CurrencyDecimals: int32(18),
				CurrencyAddress:  "",
				Timestamp:        blockEvent.BlockTimestamp,
			}
			scraper.GetTradeChannel() <- trade

			log.Info(saleMoment)
			log.Infof("transactionID: %s, block height: %d\n",
				purchaseEvent.TransactionID.String(), blockEvent.Height)
			log.Info(" ")
		}
	}

	// Update the last lastBlockNumber value.
	scraper.lastBlockNumber = scraper.lastBlockNumber + flowhelper.RequestLimit
	return nil
}

// GetMetadata returns the metadata associated to the play with @playid in the set with @setid.
func (scraper *NBATopshotScraper) GetMetadata(setid uint32, playid uint32) (map[string]interface{}, error) {
	getPlaysScript := `
	import TopShot from 0x0b2a3299cc857e29

	pub struct MomentData  {
		pub var seriesId: UInt32
		pub var setId: UInt32
		pub var playId: UInt32
		
  
		pub var play: {String: String}	 
		pub var setName: String
		pub var numMoments: UInt32
	  
		init(playid: UInt32, setid: UInt32) {
		  self.seriesId = TopShot.getSetSeries(setID: setid)!
		  self.playId = playid
		  self.setId = setid
		   
		  self.play = TopShot.getPlayMetaData(playID: self.playId)!
		  self.setName = TopShot.getSetName(setID: self.setId)!
		  self.numMoments = TopShot.getNumMomentsInEdition(setID: self.setId, playID: self.playId)!
		  
		}  
	  }
	
	pub fun main(setid: UInt32, playid: UInt32): MomentData {
		var mom: MomentData = MomentData(playid: playid, setid: setid)		
		return mom
	}
	
`
	res, err := scraper.flowClient.ExecuteScriptAtLatestBlock(context.Background(), []byte(getPlaysScript), []cadence.Value{
		cadence.UInt32(setid),
		cadence.UInt32(playid),
	})
	if err != nil {
		return make(map[string]interface{}), fmt.Errorf("error fetching sale moment from flow: %w", err)
	}

	return cadenceMomentToMap(res.(cadence.Struct)), nil
}

// cadenceMomentToMap is a helper for GetMetadata and converts a moment to a map.
func cadenceMomentToMap(cadenceMoment cadence.Value) map[string]interface{} {
	castPlay := cadenceMoment.ToGoValue().([]interface{})

	numMoments := castPlay[5].(uint32)
	auxAttributes := castPlay[3].(map[interface{}]interface{})
	attributes := make(map[string]interface{})
	for key := range auxAttributes {
		attributes[key.(string)] = auxAttributes[key]
	}
	attributes["numMomentsInEdition"] = numMoments
	return attributes
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
// Get Trade
// ---------------------------------------------------------

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
