package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	log "github.com/sirupsen/logrus"
)

func parseTrade(row []interface{}) *dia.Trade {
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
			return &trade
		}
		log.Errorln("Parsing ", t)

	}
	return nil
}

func (db *DB) GetAllTrades(t time.Time, maxTrades int) ([]dia.Trade, error) {
	r := []dia.Trade{}
	q := fmt.Sprintf("SELECT * FROM %s WHERE time > %d LIMIT %d", influxDbTradesTable, t.Unix()*1000000000, maxTrades)
	log.Debug(q)
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		log.Errorln("GetLastTrades", err)
		return r, err
	}
	if len(res) > 0 && len(res[0].Series) > 0 {
		for _, row := range res[0].Series[0].Values {
			t := parseTrade(row)
			if t != nil {
				r = append(r, *t)
			}
		}
	} else {
		log.Error("Empty response GetAllTrades")
	}
	return r, nil
}

func (db *DB) GetLastTrades(symbol string, exchange string, maxTrades int) ([]dia.Trade, error) {
	r := []dia.Trade{}
	q := fmt.Sprintf("SELECT * FROM %s WHERE exchange='%s' and symbol='%s' ORDER BY DESC LIMIT %d", influxDbTradesTable, exchange, symbol, maxTrades)
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		log.Errorln("GetLastTrades", err)
		return r, err
	}

	if len(res) > 0 && len(res[0].Series) > 0 {
		for _, row := range res[0].Series[0].Values {
			t := parseTrade(row)
			if t != nil {
				r = append(r, *t)
			}
		}
	} else {
		log.Errorf("Empty response GetLastTrades for %s on %s \n", symbol, exchange)
	}
	return r, nil
}
