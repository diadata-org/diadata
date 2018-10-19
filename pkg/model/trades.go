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

	if len(res) > 0 && len(res[0].Series) > 0 {
		for _, row := range res[0].Series[0].Values {

			if len(row) > 7 {

				t, err := time.Parse(time.RFC3339, row[0].(string))
				if err == nil {
					var estimatedUSDPrice float64
					v, o := row[1].(json.Number)
					if o {
						estimatedUSDPrice, _ = v.Float64()
					} else {
						log.Errorln("error on parsing row 1", row)
					}

					var price float64
					v, o = row[5].(json.Number)
					if o {
						price, _ = v.Float64()
					} else {
						log.Errorln("error on parsing row 5", row)
					}

					var volume float64
					v, o = row[7].(json.Number)
					if o {
						volume, _ = v.Float64()
					} else {
						log.Errorln("error on parsing row 6", row)
					}

					foreignTradeID, o := row[3].(string)
					if !o {
						log.Errorln("error on parsing row 3", row)
					}

					symbol, o := row[6].(string)
					if !o {
						log.Errorln("error on parsing row 6", row)
					}

					pair, o := row[4].(string)
					if !o {
						log.Errorln("error on parsing row 4", row)
					}

					source, o := row[2].(string)
					if !o {
						log.Errorln("error on parsing row 2", row)
					}

					trade := dia.Trade{
						Symbol:            symbol,
						Pair:              pair,
						Time:              t,
						Source:            source,
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

		}

		//log.Println(row, t, i)
		//val := row[1].(string)
		//log.Printf("[%2d] %s: %s\n", i, t.Format(time.Stamp), val)
	} else {
		log.Errorln("Empty res", len(res), res)
	}
	return r, nil
}
