package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
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

			source, o := row[2].(string)
			if !o {
				log.Errorln("error on parsing row 2", row)
			}

			foreignTradeID, o := row[3].(string)
			if !o {
				log.Errorln("error on parsing row 3", row)
			}

			pair, o := row[4].(string)
			if !o {
				log.Errorln("error on parsing row 4", row)
			}

			var price float64
			v, o = row[5].(json.Number)
			if o {
				price, _ = v.Float64()
			} else {
				log.Errorln("error on parsing row 5", row)
			}

			symbol, o := row[6].(string)
			if !o {
				log.Errorln("error on parsing row 6", row)
			}

			var volume float64
			v, o = row[7].(json.Number)
			if o {
				volume, _ = v.Float64()
			} else {
				log.Errorln("error on parsing row 6", row)
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

// GetAllTrades returns at most @maxTrades trades from influx with timestamp > @t. Only used by replayInflux option.
func (db *DB) GetAllTrades(t time.Time, maxTrades int) ([]dia.Trade, error) {
	r := []dia.Trade{}
	// TO DO: Substitute select * with precise statment select estimatedUSDPrice, source,...
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

// GetLastTrades returns the last @maxTrades of @asset on @exchange.
// If exchange is empty string it returns trades from all exchanges.
func (db *DB) GetLastTrades(asset dia.Asset, exchange string, maxTrades int) ([]dia.Trade, error) {
	r := []dia.Trade{}
	var queryString string
	var q string
	if exchange == "" {
		queryString = "SELECT estimatedUSDPrice,\"exchange\",foreignTradeID,\"pair\",price,\"symbol\",volume" +
			" FROM %s WHERE quotetokenaddress='%s' and quotetokenblockchain='%s' ORDER BY DESC LIMIT %d"
		q = fmt.Sprintf(queryString, influxDbTradesTable, asset.Address, asset.Blockchain.Name, maxTrades)
	} else {
		queryString = "SELECT estimatedUSDPrice,\"exchange\",foreignTradeID,\"pair\",price,\"symbol\",volume" +
			" FROM %s WHERE exchange='%s' and quotetokenaddress='%s' and quotetokenblockchain='%s' ORDER BY DESC LIMIT %d"
		q = fmt.Sprintf(queryString, influxDbTradesTable, exchange, asset.Address, asset.Blockchain.Name, maxTrades)
	}
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
		log.Errorf("Empty response GetLastTrades for %s on %s \n", asset.Symbol, exchange)
	}
	return r, nil
}
