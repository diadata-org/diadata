package nftdatascrapers

// Please implement the scraping of coingecko quotations here.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk/client"
	"google.golang.org/grpc"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
)

type NBATopshotScraper struct {
	nftscraper NFTScraper
	flowClient *client.Client
	ticker     *time.Ticker
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

	flowClient, err := client.New("access.mainnet.nodes.onflow.org:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	err = flowClient.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	nftScraper := NFTScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		error:        nil,
		relDB:        rdb,
		chanData:     make(chan dia.NFT),
	}
	s := &NBATopshotScraper{
		nftscraper: nftScraper,
		flowClient: flowClient,
		ticker:     time.NewTicker(refreshDelay),
	}

	amap, err := s.GetAttributeMap()
	if err != nil {
		fmt.Println("err here: ", err)
	}
	fmt.Println("amap: ", amap)

	go s.mainLoop()
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *NBATopshotScraper) mainLoop() {
	for {
		select {
		case <-scraper.ticker.C:
			err := scraper.UpdateNFT()
			if err != nil {
				log.Error("error updating NFT: ", err)
			}
		case <-scraper.nftscraper.shutdown: // user requested shutdown
			log.Info("NBA Topshot scraper shutting down")
			err := scraper.Close()
			scraper.cleanup(err)
			return
		}
	}
}

func (scraper *NBATopshotScraper) UpdateNFT() error {
	fmt.Println("fetch data...")
	nfts, err := scraper.FetchData()
	if err != nil {
		return err
	}
	for _, nft := range nfts {
		scraper.GetDataChannel() <- nft
	}
	return nil
}

func (scraper *NBATopshotScraper) FetchData() (nfts []dia.NFT, err error) {

	numSets, err := scraper.GetNumSets()
	if err != nil {
		return []dia.NFT{}, err
	}
	fmt.Println("number of sets: ", numSets)
	var nbaTopshotNFTs []dia.NFT
	allMoments, err := scraper.GetAllMoments()
	if err != nil {
		return []dia.NFT{}, err
	}
	// TO DO: Import attributes map and assign attributes to moments below.
	// TO DO: Get creation time = block time
	// TO DO: Can we get creator address from event?
	for _, moment := range allMoments {
		m := MomentMintedEvent(moment)
		metadata, err := scraper.GetMetadata(uint32(m.SetID()), uint32(m.PlayID()))
		if err != nil {
			log.Error(err)
		}
		nbaTopshotNFTs = append(nbaTopshotNFTs, dia.NFT{
			TokenID:    strconv.Itoa(int(m.ID())),
			Attributes: metadata,
		})
	}
	fmt.Println("results: ", nbaTopshotNFTs)

	return nbaTopshotNFTs, nil
}

// ---------------------------------------------------------
// Get Data
// ---------------------------------------------------------

// GetAllMoments returns all moments from genesis to the latest block by iterating through
// blocks and looking for MomentMinted events.
func (scraper *NBATopshotScraper) GetAllMoments() (mintedMoments []cadence.Event, err error) {
	latestBlock, err := scraper.flowClient.GetLatestBlock(context.Background(), false)
	if err != nil {
		log.Error(err)
	}
	// For some reason, for block heights smaller than 14mio, the rpc server returns an error
	for i := 0; i < 500; i++ {
		m, err := scraper.GetMintedMoments(latestBlock.Height-uint64((i+1)*249), latestBlock.Height-uint64(i*249))
		if err != nil {
			fmt.Println(err)
		}
		mintedMoments = append(mintedMoments, m...)
	}
	return
}

// GetMintedMoments returns all moments minted between blocks @startheight and @endheight.
// The difference @endheight-@starthight is limited to 250.
func (scraper *NBATopshotScraper) GetMintedMoments(startheight, endheight uint64) (mintedMoments []cadence.Event, err error) {

	blockEvents, err := scraper.flowClient.GetEventsForHeightRange(context.Background(), client.EventRangeQuery{
		Type:        "A.0b2a3299cc857e29.TopShot.MomentMinted",
		StartHeight: startheight,
		EndHeight:   endheight,
	})
	if err != nil {
		return
	}

	for _, blockEvent := range blockEvents {
		for _, playCreatedEvent := range blockEvent.Events {
			mintedMoments = append(mintedMoments, playCreatedEvent.Value)
		}
	}
	return

}

type MomentMintedEvent cadence.Event

func (mme MomentMintedEvent) ID() uint64 {
	return uint64(mme.Fields[0].(cadence.UInt64))
}

func (mme MomentMintedEvent) PlayID() uint32 {
	return uint32(mme.Fields[1].(cadence.UInt32))
}

func (mme MomentMintedEvent) SetID() uint32 {
	return uint32(mme.Fields[2].(cadence.UInt32))
}

func (mme MomentMintedEvent) SerialNumber() uint32 {
	return uint32(mme.Fields[3].(cadence.UInt32))
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
	// type Plays cadence.Struct
	// play := Plays(res.(cadence.Struct))
	// fmt.Println("play: ", play)

	return cadenceMomentToMap(res.(cadence.Struct)), nil
}

// cadenceMomentToMap converts a moment to a map.
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

// GetPlaysBySet returns all plays contained in a set.
func (scraper *NBATopshotScraper) GetPlaysBySet(setid uint32) ([]cadence.Value, error) {
	getPlaysScript := `
	import TopShot from 0x0b2a3299cc857e29

	pub struct MomentData  {
		pub var seriesId: UInt32
		pub var setId: UInt32
		pub var playId: UInt32
  
		pub var play: {String: String}	 
		pub var setName: String
	  
		init(playid: UInt32, setid: UInt32) {
		  self.seriesId = TopShot.getSetSeries(setID: setid)!
		  self.playId = playid
		  self.setId = setid
		   
		  self.play = TopShot.getPlayMetaData(playID: self.playId)!
		  self.setName = TopShot.getSetName(setID: self.setId)!
		  
		}  
	  }
	
	pub fun main(setid: UInt32): [MomentData] {
		var moments: [MomentData] = []
		var allplayids: [UInt32] = TopShot.getPlaysInSet(setID: setid)!

		for playid in allplayids {
			var mom: MomentData = MomentData(playid: playid, setid: setid)
			moments.append(mom)
		}
		
		return moments
	}
	
`
	res, err := scraper.flowClient.ExecuteScriptAtLatestBlock(context.Background(), []byte(getPlaysScript), []cadence.Value{
		cadence.UInt32(setid),
	})
	if err != nil {
		return []cadence.Value{}, fmt.Errorf("error fetching sale moment from flow: %w", err)
	}
	type Plays cadence.Array

	setID := Plays(res.(cadence.Array))
	fmt.Println("number of plays: ", len(setID.Values))
	return setID.Values, nil
}

// GetNumSets returns the number of available sets.
func (scraper *NBATopshotScraper) GetNumSets() (uint32, error) {
	getSetIDScript := `
	import TopShot from 0x0b2a3299cc857e29
	pub fun main(): UInt32 {
		return TopShot.nextSetID
	}
	
`
	res, err := scraper.flowClient.ExecuteScriptAtLatestBlock(context.Background(), []byte(getSetIDScript), []cadence.Value{})
	if err != nil {
		return 0, fmt.Errorf("error fetching set id from flow: %w", err)
	}
	type SetID cadence.UInt32
	setID := SetID(res.(cadence.UInt32))

	return uint32(setID), nil
}

type identifier struct {
	SetID  uint32
	PlayID uint32
}

// GetAttributesMap returns a map that uniquely maps setID and playID onto attributes.
func (scraper *NBATopshotScraper) GetAttributeMap() (map[identifier]map[string]interface{}, error) {
	attrMap := make(map[identifier]map[string]interface{})
	numSets, err := scraper.GetNumSets()
	if err != nil {
		return attrMap, err
	}
	for i := 1; i < int(numSets); i++ {

		values, err := scraper.GetPlaysBySet(uint32(i))
		if err != nil {
			fmt.Println("getting setID: ", err)
		}
		for _, val := range values {
			play := cadenceToPlay(val)
			idfier := identifier{
				SetID:  play.SetID,
				PlayID: play.PlayID,
			}
			attributes, err := scraper.GetMetadata(idfier.SetID, idfier.PlayID)
			if err != nil {
				log.Errorf("fetching attributes for setID %v and playID %v: %v", idfier.SetID, idfier.PlayID, err)
			}
			attributes["seriesID"] = play.SeriesID
			attributes["setName"] = play.SetName
			attrMap[idfier] = attributes
			fmt.Println("attributes: ", attributes)
		}

	}
	return attrMap, nil
}

func cadenceToPlay(cadencePlay cadence.Value) (play Play) {
	castPlay := cadencePlay.ToGoValue().([]interface{})
	play.SeriesID = castPlay[0].(uint32)
	play.SetID = castPlay[1].(uint32)
	play.PlayID = castPlay[2].(uint32)
	play.SetName = castPlay[4].(string)

	auxAttributes := castPlay[3].(map[interface{}]interface{})
	attributes := make(map[string]interface{})
	for key := range auxAttributes {
		attributes[key.(string)] = auxAttributes[key]
	}
	play.Attributes = attributes

	return play
}

// GetDataChannel returns the scrapers data channel.
func (scraper *NBATopshotScraper) GetDataChannel() chan dia.NFT {
	return scraper.nftscraper.chanData
}

// closes all connected Scrapers. Must only be called from mainLoop
func (scraper *NBATopshotScraper) cleanup(err error) {
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
func (scraper *NBATopshotScraper) Close() error {
	if scraper.nftscraper.closed {
		return errors.New("scraper already closed")
	}
	close(scraper.nftscraper.shutdown)
	<-scraper.nftscraper.shutdownDone
	scraper.nftscraper.errorLock.RLock()
	defer scraper.nftscraper.errorLock.RUnlock()
	return scraper.nftscraper.error
}
