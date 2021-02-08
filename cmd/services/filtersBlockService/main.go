package main

import (
	"context"
	"flag"
	"sync"
	"time"

	filters "github.com/diadata-org/diadata/internal/pkg/filtersBlockService"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/kafkaHelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

var (
	replayInflux = flag.Bool("replayInflux", false, "replayInflux ?")
)

func init() {
	flag.Parse()
	log.Println("replayInflux=", *replayInflux)
}

func handler(channel chan *dia.FiltersBlock, wg *sync.WaitGroup, w *kafka.Writer) {
	var block int
	for {
		t, ok := <-channel
		if !ok {
			log.Printf("handler: finishing channel")
			wg.Done()
			return
		}
		block++
		log.Infoln("kafka: generated ", block, " blocks")
		err := kafkaHelper.WriteMessage(w, t)
		if err != nil {
			log.Errorln("kafka: handleBlocks", err)
		}
	}
}

func loadFilterPointsFromPreviousBlock() []dia.FilterPoint {
	// load the previous block points so that we have a value even if
	// there is no trades
	lastFilterPoints := []dia.FilterPoint{}
	lastFilterBlock, err := kafkaHelper.GetLastElement(kafkaHelper.TopicFiltersBlock)
	if err == nil {
		lastFilterPoints = lastFilterBlock.(dia.FiltersBlock).FiltersBlockData.FilterPoints
	}
	return lastFilterPoints
}

//  docker exec -it <cointainer> filtersBlockService -replayInflux

func createTradeBlockFromInflux(d models.Datastore, f *filters.FiltersBlockService) {
	//now := time.Now()
	//then := now.AddDate(0, -1, 0)
	then := time.Unix(1539475200, 0)
	//"1405544146"

	log.Info("createTradeBlockFromInflux")
	var currentBlock int64
	trades := []dia.Trade{}
	for {
		log.Info("sleeping")
		time.Sleep(1 * time.Second)
		r, err := d.GetAllTrades(then, 1000)
		if err != nil {
			log.Errorln("createTradeBlockFromInflux", r)
			continue
		}
		if len(r) == 0 {
			log.Info("no new trades...")
			break
		} else {
			then = r[len(r)-1].Time
			log.Infoln("x got", len(r), "trades", then)
			for _, v := range r {
				if v.Source == "Simex" {
					continue
				}
				block := (v.Time.Unix() / dia.BlockSizeSeconds)
				if block != currentBlock {
					var t1 time.Time
					var t2 time.Time
					currentBlock = block
					if len(trades) > 0 {
						t1 = trades[0].Time
						t2 = trades[len(trades)-1].Time
					}
					b := &dia.TradesBlock{
						TradesBlockData: dia.TradesBlockData{
							Trades:    trades,
							BeginTime: t1,
							EndTime:   t2,
						},
					}
					if len(trades) > 5 {
						log.Infoln("calling ProcessTradesBlock", len(trades), "trades blocknumber:", currentBlock, t1, t2)
						f.ProcessTradesBlock(b)
						log.Infoln("bang", currentBlock)
					} else {
						log.Info("not enough trades in block ignoring...", len(trades), currentBlock, t1, t2)
					}
					trades = []dia.Trade{}
				} else {
					trades = append(trades, v)
				}
			}
		}
	}
}

func main() {

	if *replayInflux {
		s, err := models.NewInfluxDataStore()
		if err != nil {
			log.Errorln("NewDataStore", err)
		}
		f := filters.NewFiltersBlockService(nil, s, nil)
		createTradeBlockFromInflux(s, f)
	} else {
		s, err := models.NewDataStore()
		if err != nil {
			log.Errorln("NewDataStore", err)
		}
		channel := make(chan *dia.FiltersBlock)

		f := filters.NewFiltersBlockService(loadFilterPointsFromPreviousBlock(), s, channel)

		w := kafkaHelper.NewSyncWriter(kafkaHelper.TopicFiltersBlock)

		defer w.Close()

		wg := sync.WaitGroup{}

		go handler(channel, &wg, w)

		r := kafkaHelper.NewReaderNextMessage(kafkaHelper.TopicTradesBlock)
		defer r.Close()

		for {
			m, err := r.ReadMessage(context.Background())
			if err != nil {
				log.Printf(err.Error())
			} else {
				log.Info("get block from tradesBlock")
				var tb dia.TradesBlock
				err := tb.UnmarshalBinary(m.Value)
				if err != nil {
					log.Error("error unmarshalling trades block")
				}
				if err == nil {
					f.ProcessTradesBlock(&tb)
				}
			}
		}
	}
}
