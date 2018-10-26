package main

import (
	"github.com/DefinitelyNotAGoat/goTezos"
	"github.com/diadata-org/diadata/pkg/dia"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

const rpcURL = "http://tezosnode:8732"
const xtzAddress = "tz1RCFbB9GpALpsZtu6J58sb74dm8qe6XBzv"
const totalSupply = 763306930.69
const decimals = 1e8
const tezosSymbol = "xtz"

/// Retrieve every minute according to block frequency
const blockFrequency = 60

var client *dia.Client

func retrieveXTZSupply() {
	b, eg := goTezos.GetChainHead()
	if eg != nil {
		log.Error("Error getting block:" + eg.Error())
	} else {
		log.Println("Block hash:" + b.Hash)
		v, e := goTezos.GetAccountBalanceAtBlock(xtzAddress, b.Hash)
		if e != nil {
			log.Error("Error retrieving balance for:" + xtzAddress + " error:" + e.Error())
		} else {
			balance := float64(v / decimals)
			supply := totalSupply - balance
			log.Println("Balance:" + strconv.FormatFloat(balance, 'f', 4, 64) + " Supply:" + strconv.FormatFloat(supply, 'f', 4, 64))
			client.SendSupply(&dia.Supply{
				Symbol:            tezosSymbol,
				CirculatingSupply: supply,
			})

		}
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
			retrieveXTZSupply()
		}
	}
}

func (t *Task) stop() {
	log.Println("Stoping tezos supply scrapper update thread...")
	close(t.closed)
	t.wg.Wait()
	log.Println("Done")
}

func main() {
	goTezos.SetRPCURL(rpcURL)
	task := &Task{
		closed: make(chan struct{}),
		ticker: time.NewTicker(time.Second * blockFrequency),
	}

	config := dia.GetConfigApi()
	if config == nil {
		panic("Could not load config")
	}
	client = dia.NewClient(config)
	if client == nil {
		panic("Could not load client")
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
