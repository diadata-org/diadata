package models

import (
	"encoding/json"
	"fmt"
	"github.com/diadata-org/diadata/pkg/dia"
	log "github.com/sirupsen/logrus"
	"time"
)

func (db *DB) GetLastTrades(symbol string, exchange string, maxTrades int) ([]dia.Trade, error) {
	r := []dia.Trade{}
	log.Info("GetLastTrades")
	q := fmt.Sprintf("SELECT * FROM %s WHERE exchange='%s' and symbol='%s' ORDER BY DESC LIMIT %d", influxDbTradesTable, exchange, symbol, maxTrades)
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		log.Errorln("GetLastTrades", err)
		return r, err
	}

	//	http_restserver.1.wvd8ndffl6ft@d2.diadata.org    | time="2018-10-11T11:27:51Z" level=info msg="[2018-10-10T13:45:47Z 6516.237007925503 OKEx BTC_USDT 6608.4357 BTC 0.01] 2018-10-10 13:45:47 +0000 UTC 1"
	/*
	   type Trade struct {
	   	Symbol            string
	   	Pair              string
	   	Price             float64
	   	Volume            float64 // negative if result of Market order Sell
	   	Time              time.Time
	   	ForeignTradeID    string
	   	EstimatedUSDPrice float64 // will be filled by the TradeBlock Service
	   	Source            string
	   }
	*/

	if len(res) > 0 && len(res[0].Series) > 0 {
		for _, row := range res[0].Series[0].Values {

			t, err := time.Parse(time.RFC3339, row[0].(string))
			if err == nil {
				log.Errorln("GetLastTrades", err)
				estimatedUSDPrice, _ := row[1].(json.Number).Float64()
				price, _ := row[4].(json.Number).Float64()
				volume, _ := row[6].(json.Number).Float64()
				foreignTradeID := ""
				if len(row) > 7 {
					foreignTradeID = row[7].(string)
				}
				trade := dia.Trade{
					Symbol:            row[5].(string),
					Pair:              row[3].(string),
					Time:              t,
					Source:            row[2].(string),
					EstimatedUSDPrice: estimatedUSDPrice,
					Price:             price,
					Volume:            volume,
					ForeignTradeID:    foreignTradeID,
				}
				r = append(r, trade)
			} else {
				log.Errorln("Parsing", t)
			}
			log.Println(row, t)
		}

		//log.Println(row, t, i)
		//val := row[1].(string)
		//log.Printf("[%2d] %s: %s\n", i, t.Format(time.Stamp), val)
	} else {
		log.Errorln("Empty res", len(res), res)
	}
	return r, nil
}
