package main

import (
	"encoding/json"
	"github.com/diadata-org/diadata/internal/pkg/exchange-scrapers"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
	"os"
	"os/signal"
	"os/user"
	"strings"
	"sync"
	"time"
)

var (
	db *redis.Client
)

func getConfig(exchange string) (*dia.ConfigApi, error) {
	var configApi dia.ConfigApi
	usr, _ := user.Current()
	dir := usr.HomeDir
	configFileApi := dir + "/config/secrets/api_" + strings.ToLower(exchange)
	err := gonfig.GetConf(configFileApi, &configApi)
	return &configApi, err
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
	for _, e := range dia.Exchanges() {
		c, err := getConfig(e)
		var s scrapers.APIScraper
		if err == nil {
			s = scrapers.NewAPIScraper(e, c.ApiKey, c.SecretKey)
		} else {
			log.Error("Error processing config for exchange:"+e+" error:", err.Error())
			s = scrapers.NewAPIScraper(e, "", "")
		}
		if s != nil {
			p, err := s.FetchAvailablePairs()
			if err == nil {
				err := addPairsToDb(e, p)
				if err == nil {
					log.Println("Exchange :" + e + " updated")
				} else {
					log.Error("Error adding pairs  to redis for exchange:"+e+" error:", err.Error())
				}
			} else {
				log.Error("Error fetching pairs for exchange:"+e+" error:", err.Error())
			}
		} else {
			log.Error("Error creating APIScraper forexchange:" + e)
		}
	}
	log.Println("Update complete.")
}

func addPairsToDb(exchange string, pairs []dia.Pair) error {
	jsonInfo, _ := json.Marshal(pairs)
	return db.Set(exchange, jsonInfo, 0).Err()
}

func initDb() (err error) {
	db = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong2, err := db.Ping().Result()
	if err != nil {
		log.Error("initDB error:", err)
		return
	}
	log.Debug("NewDB", pong2)
	return nil
}

func stopDb() {
	db.Close()
}

func main() {
	task := &Task{
		closed: make(chan struct{}),
		ticker: time.NewTicker(time.Second * 5),
	}
	err := initDb()
	if err != nil {
		panic("Can not initialize db error:" + err.Error())
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
