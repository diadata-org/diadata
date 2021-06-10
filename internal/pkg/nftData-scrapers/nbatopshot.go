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

	for i := 0; i < int(numSets)-1; i++ {

		values, err := scraper.GetPlaysBySet(uint32(i))
		if err != nil {
			fmt.Println("getting setID: ", err)
		}
		fmt.Println("number of values: ", len(values))
		val := values[0]
		testplay := cadenceToPlay(val)
		fmt.Println(testplay)

		nbaTopshotNFTs = append(nbaTopshotNFTs, dia.NFT{
			TokenID: strconv.Itoa(i),
			// NFTClass:       cryptopunkNFTClass,
			// CreationTime:   creationTime,
			// CreatorAddress: creatorAddress,
			// Attributes:     result,
			URI: "to do",
		})
	}
	return nbaTopshotNFTs, nil
}

// ---------------------------------------------------------
// Get Data
// ---------------------------------------------------------

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
	fmt.Println("Play: ", setID.Values[10])
	return setID.Values, nil
}

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

// // GetCryptopunkCreationTime returns the creation time from Opensea
// func GetCryptopunkCreationTime(punkResp []byte) (time.Time, error) {
// 	var resp OpenSeaCryptopunkResponse
// 	var t time.Time
// 	if err := json.Unmarshal(punkResp, &resp); err != nil {
// 		return t, err
// 	}

// 	layout := "2006-01-02T15:04:05"
// 	t, err := time.Parse(layout, resp.AssetContract.CreatedDate)

// 	if err != nil {
// 		return t, err
// 	}

// 	return t, nil
// }

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
