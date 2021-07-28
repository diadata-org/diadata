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

	}
	return nil
}

// GetAllTrades returns at most @maxTrades trades from influx with timestamp > @t. Only used by replayInflux option.
func (db *DB) GetAllTrades(t time.Time, maxTrades int) ([]dia.Trade, error) {
	r := []dia.Trade{}
	// TO DO: Substitute select * with precise statment select estimatedUSDPrice, source,...
	q := fmt.Sprintf("SELECTtime, estimatedUSDPrice, verified, foreignTradeID, pair, price,symbol, volume  FROM %s WHERE time > %d LIMIT %d", influxDbTradesTable, t.Unix()*1000000000, maxTrades)
	log.Debug(q)
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		log.Errorln("GetLastTrades", err)
		return r, err
	}
	if len(res) > 0 && len(res[0].Series) > 0 {
		for _, row := range res[0].Series[0].Values {
			t := parseTrade(row)
			log.Errorln("row trade parseTrade", row)
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
	q := fmt.Sprintf("SELECTtime, estimatedUSDPrice, verified, foreignTradeID, pair, price,symbol, volume  FROM %s WHERE exchange='%s' and symbol='%s' ORDER BY DESC LIMIT %d", influxDbTradesTable, exchange, symbol, maxTrades)
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

func (db *DB) GetLastTradesAllExchanges(symbol string, maxTrades int) ([]dia.Trade, error) {
	r := []dia.Trade{}
	q := fmt.Sprintf("SELECTtime, estimatedUSDPrice, verified, foreignTradeID, pair, price,symbol, volume  FROM %s WHERE symbol='%s' ORDER BY DESC LIMIT %d", influxDbTradesTable, symbol, maxTrades)
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
		log.Errorf("Empty response GetLastTradesAllExchanges for %s \n", symbol)
	}
	return r, nil
}

func (db *DB) GetTradesByExchanges(symbol string, exchanges []string, startTime, endTime time.Time, maxTrades int) ([]dia.Trade, error) {
	r := []dia.Trade{}
	subquery := ""
	query := ""
	if len(exchanges) > 0 {
		for count, exchange := range exchanges {
			if len(exchanges)-1 == count {
				subquery = subquery + fmt.Sprintf("exchange='%s'", exchange)
			} else {
				subquery = subquery + fmt.Sprintf("exchange='%s'", exchange) + " or "

			}

		}
		query = fmt.Sprintf("SELECT time, estimatedUSDPrice, verified, foreignTradeID, pair, price,symbol, volume  FROM %s WHERE symbol='%s' and %s and  time >= %d AND time <= %d ", influxDbTradesTable, symbol, subquery, startTime.UnixNano(), endTime.UnixNano())

	}
	query = fmt.Sprintf("SELECT time, estimatedUSDPrice, verified, foreignTradeID, pair, price,symbol, volume  FROM %s WHERE symbol='%s'  and  time >= %d AND time <= %d ", influxDbTradesTable, symbol, startTime.UnixNano(), endTime.UnixNano())

	log.Infoln("GetTradesByExchanges Query", query)
	res, err := queryInfluxDB(db.influxClient, query)

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
		log.Errorf("Empty response GetLastTradesAllExchanges for %s \n", symbol)
	}
	return r, nil
}

func (db *DB) GetTradesByExchange(symbol string, exchange string, startTime, endTime time.Time, maxTrades int) ([]dia.Trade, error) {
	r := []dia.Trade{}

	q := fmt.Sprintf("SELECTtime, estimatedUSDPrice, verified, foreignTradeID, pair, price,symbol, volume  FROM %s WHERE symbol='%s' and exchange='%s' and  time >= %d AND time <= %d ", influxDbTradesTable, symbol, exchange, startTime.UnixNano(), endTime.UnixNano())
	log.Infoln("GetTradesByExchange Query", q)
	res, err := queryInfluxDB(db.influxClient, q)
	log.Errorln("res", res)
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
		log.Errorf("Empty response GetLastTradesAllExchanges for %s \n", symbol)
	}
	return r, nil
}
