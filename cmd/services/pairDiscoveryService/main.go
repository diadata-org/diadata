package main

import (
	"github.com/diadata-org/diadata/internal/pkg/exchange-scrapers"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

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
			updateExchangePairs()
		}
	}
}

func (t *Task) stop() {
	log.Println("Stoping exchange pair update thread...")
	close(t.closed)
	t.wg.Wait()
	log.Println("Thread stopped, cleaning...")
	// Clean if required
	stopDb()
	log.Println("Done")
}

func updateExchangePairs() {
	log.Println("Updating exchanges pairs...")
	sg := scrapers.NewGateIOScraper()
	p := sg.FetchAvailablePairs()
	addPairsToDb("gateio", p)
	log.Println("Update complete.")
}

func addPairsToDb(exchange string, pairs []string) {
	log.Println("Updating:", exchange)
	for _, p := range pairs {
		log.Println(p)
	}
}

func initDb() {

}

func stopDb() {

}

func main() {
	task := &Task{
		closed: make(chan struct{}),
		ticker: time.NewTicker(time.Second * 5),
	}
	initDb()
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
