package main

import (
	"github.com/diadata-org/diadata/internal/pkg/blockchain-scrapers/blockchains/ethereum-classic/scrapers/etc/client"
	"github.com/diadata-org/diadata/pkg/dia"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

/// Retrieve every minute according to block frequency
const blockFrequency = 10
const dataPersistenceFreq = 1000
const rpcURL = "http://etcnode:8545"
const etcSymbol = "etc"

var db *dia.Client
var e *ethclassic.EthClient

func retrieveETCSupply() {
	if bn, err := e.GetLatestBlockNumber(); err == nil {
		log.Println("Latest block is: " + strconv.Itoa(int(bn)) + " currently at : " + strconv.Itoa(int(e.LastBlock)))
		b, err := e.GetTotalSupplyAtBlock(bn)
		if err != nil {
			log.Error("Error retrieveing supply at latest block reason: " + err.Error())
			return
		}
		if err = db.SendSupply(&dia.Supply{
			Symbol:            etcSymbol,
			CirculatingSupply: b,
			Block:             bn,
		}); err != nil {
			log.Error("Error sending supply reason: " + err.Error())
		}
		log.Println("Supply at block: " + strconv.Itoa(int(bn)) + " is : " + strconv.FormatFloat(b, 'G', -1, 64))
	} else {
		log.Error("Error retrieveing latest block reason: " + err.Error())
	}
}

type Task struct {
	closed chan struct{}
	wg     sync.WaitGroup
	ticker *time.Ticker
}

func (t *Task) run() {
	for {
		select {
		case <-t.closed:
			return
		case <-t.ticker.C:
			retrieveETCSupply()
		}
	}
}

func getInitialSupply() {
	//block until close to the latest block
	for {
		if bn, err := e.GetLatestBlockNumber(); err == nil {
			log.Println("Latest block is: " + strconv.Itoa(int(bn)) + " currently at : " + strconv.Itoa(int(e.LastBlock)))
			if int(bn-e.LastBlock) < dataPersistenceFreq {
				/// we are close enough to completion let regular routine complete
				log.Println("Terminating init with Latest: " + strconv.Itoa(int(bn)) + " currently at : " + strconv.Itoa(int(e.LastBlock)))
				break
			}
			for i := 0; i < int(bn-e.LastBlock)/dataPersistenceFreq; i++ {
				blockN := e.LastBlock + int64(dataPersistenceFreq)
				b, err := e.GetTotalSupplyAtBlock(blockN)
				if err != nil {
					log.Fatal(err.Error())
				}
				log.Println("Supply at block: " + strconv.Itoa(int(blockN)) + " is : " + strconv.FormatFloat(b, 'G', -1, 64))
				if err = db.SendSupply(&dia.Supply{
					Symbol:            etcSymbol,
					CirculatingSupply: b,
					Block:             blockN,
				}); err != nil {
					log.Error("Error sending supply reason: " + err.Error())
				}
			}
			b, err := e.GetTotalSupplyAtBlock(bn)
			if err != nil {
				log.Fatal(err.Error())
			}
			log.Println("Supply at block: " + strconv.Itoa(int(bn)) + " is : " + strconv.FormatFloat(b, 'G', -1, 64))
		} else {
			log.Warning("Error retrieving latest block number. ETC node not found. This is normal if the node have not started yet.")
			time.Sleep(time.Second)

		}
	}
}

func initialize() {
	e = &ethclassic.EthClient{
		URL:        rpcURL,
		LastBlock:  -1,
		LastSupply: 72009990.50, // Genesis (60M Crowdsale 12M Other)
	}
	config := dia.GetConfigApi()
	if config == nil {
		panic("Could not load config")
	}
	db = dia.NewClient(config)
	if db == nil {
		panic("Could not load client")
	}
	if s, err := dia.GetSupply(etcSymbol); err == nil {
		e.LastBlock = s.Block
		e.LastSupply = s.CirculatingSupply
		log.Println("Supply read from API is : " + strconv.FormatFloat(s.CirculatingSupply, 'G', -1, 64) + " at block: " + strconv.Itoa(int(s.Block)))
	} else {
		log.Error("Error retrieveing initial supply from API. Restaring from block 0. This will take a while. Reason: " + err.Error())
	}
	getInitialSupply()
}
func (t *Task) stop() {
	log.Println("Stoping ethereum classic supply scrapper update thread...")
	close(t.closed)
	t.wg.Wait()
	log.Println("Done")
}

func main() {
	log.Println("Starting ethereum classic scrapper")
	initialize()
	log.Println("Initializing complete")
	task := &Task{
		closed: make(chan struct{}),
		ticker: time.NewTicker(time.Second * blockFrequency),
	}
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	task.wg.Add(1)
	go func() { defer task.wg.Done(); task.run() }()
	select {
	case <-c:
		log.Println("Got signal.")
		task.stop()
	}
}
