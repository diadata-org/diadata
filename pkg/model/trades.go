package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
)

// parseTrade parses a trade as retreived from influx. If fullAsset=true blockchain and address of
// the corresponding asset is returned as well.
func parseTrade(row []interface{}, fullBasetoken bool) *dia.Trade {
	if len(row) > 10 {
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
				log.Errorln("error on parsing row 7", row)
			}
			var verified bool
			ver, ok := row[8].(string)
			if ok {
				if ver == "true" {
					verified = true
				}
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
				VerifiedPair:      verified,
			}

			if fullBasetoken {
				basetokenblockchain, o := row[9].(string)
				if !o {
					log.Errorln("error on parsing row 9", row)
				}
				basetokenaddress, o := row[10].(string)
				if !o {
					log.Errorln("error on parsing row 10", row)
				}
				trade.BaseToken.Blockchain = basetokenblockchain
				trade.BaseToken.Address = basetokenaddress
			}

			return &trade
		}

	}
	return nil
}

func (datastore *DB) GetTradesByExchanges(asset dia.Asset, exchanges []string, startTime, endTime time.Time) ([]dia.Trade, error) {
	return datastore.GetTradesByExchangesFull(asset, exchanges, false, startTime, endTime)
}
func (datastore *DB) GetTradesByExchangesAndBaseAssets(asset dia.Asset, exchanges []string, startTime, endTime time.Time) ([]dia.Trade, error) {
	return datastore.GetTradesByExchangesFull(asset, exchanges, false, startTime, endTime)
}

func (datastore *DB) GetTradesByExchangesFull(asset dia.Asset, baseassets []dia.Asset, exchanges []string, returnBasetoken bool, startTime, endTime time.Time) ([]dia.Trade, error) {
	var r []dia.Trade
	subQuery := ""
	subQueryBase := ""
	if len(exchanges) > 0 {
		for _, exchange := range exchanges {
			subQuery = subQuery + fmt.Sprintf("%s|", exchange)
		}
		subQuery = "and exchange =~ /" + strings.TrimRight(subQuery, "|") + "/"

		if len(baseassets) > 0 {
			for i, baseasset := range baseassets {
				if i == 0 {
					subQueryBase = subQueryBase + fmt.Sprintf(`(basetokenaddress='%s' and basetokenblockchain='%s')`, baseasset.Address, baseasset.Blockchain)

				} else {
					subQueryBase = subQueryBase + fmt.Sprintf(`or (basetokenaddress='%s' and basetokenblockchain='%s')`, baseasset.Address, baseasset.Blockchain)
				}
			}
			//(basetokenaddress='0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48' and basetokenblockchain='Ethereum')
		}
		log.Errorln("subQueryBase", subQueryBase)
	}
	query := fmt.Sprintf("SELECT time,estimatedUSDPrice,exchange,foreignTradeID,pair,price,symbol,volume,verified,basetokenblockchain,basetokenaddress FROM %s WHERE (quotetokenaddress='%s' and quotetokenblockchain='%s') %s %s AND estimatedUSDPrice > 0 AND time >= %d AND time <= %d ", influxDbTradesTable, asset.Address, asset.Blockchain, subQuery, subQueryBase, startTime.UnixNano(), endTime.UnixNano())
	log.Errorln("Query", query)
	res, err := queryInfluxDB(datastore.influxClient, query)
	if err != nil {
		return r, err
	}

	if len(res) > 0 && len(res[0].Series) > 0 {
		for _, row := range res[0].Series[0].Values {
			t := parseTrade(row, returnBasetoken)
			if t != nil {
				r = append(r, *t)
			}
		}
	} else {
		return nil, fmt.Errorf("no trades found")
	}
	return r, nil
}

// GetTradesByExchangesBatched executes multiple select queries on the trades table in one batch.
// The time ranges of the queries are given by the intervals [startTimes[i], endTimes[i]].
func (datastore *DB) GetTradesByExchangesBatched(quoteasset dia.Asset, baseassets []dia.Asset, exchanges []string, startTimes, endTimes []time.Time) ([]dia.Trade, error) {
	return datastore.GetTradesByExchangesBatchedFull(quoteasset, baseassets, exchanges, false, startTimes, endTimes)
}

// GetTradesByExchangesBatchedFull executes multiple select queries on the trades table in one batch.
// The time ranges of the queries are given by the intervals [startTimes[i], endTimes[i]].
func (datastore *DB) GetTradesByExchangesBatchedFull(quoteasset dia.Asset, baseassets []dia.Asset, exchanges []string, returnBasetoken bool, startTimes, endTimes []time.Time) ([]dia.Trade, error) {
	log.Errorln("GetTradesByExchangesBatchedFull baseassets", baseassets)

	var r []dia.Trade
	if len(startTimes) != len(endTimes) {
		return []dia.Trade{}, errors.New("number of start times must equal number of end times.")
	}
	var query string
	for i := range startTimes {
		subQuery := ""
		subQueryBase := ""
		if len(exchanges) > 0 {
			for _, exchange := range exchanges {
				subQuery = subQuery + fmt.Sprintf("%s|", exchange)
			}
			subQuery = "and exchange =~ /" + strings.TrimRight(subQuery, "|") + "/"
		}

		if len(baseassets) > 0 {
			for i, baseasset := range baseassets {
				if i == 0 {
					subQueryBase = subQueryBase + fmt.Sprintf(`(basetokenaddress='%s' and basetokenblockchain='%s')`, baseasset.Address, baseasset.Blockchain)

				} else {
					subQueryBase = subQueryBase + fmt.Sprintf(`or (basetokenaddress='%s' and basetokenblockchain='%s')`, baseasset.Address, baseasset.Blockchain)
				}
			}
		}
		log.Errorln("subQueryBase", subQueryBase)
		query = query + fmt.Sprintf("SELECT time,estimatedUSDPrice,exchange,foreignTradeID,pair,price,symbol,volume,verified,basetokenblockchain,basetokenaddress FROM %s WHERE quotetokenaddress='%s' AND quotetokenblockchain='%s' %s %s AND estimatedUSDPrice > 0 AND time > %d AND time <= %d ;", influxDbTradesTable, quoteasset.Address, quoteasset.Blockchain, subQuery, subQueryBase, startTimes[i].UnixNano(), endTimes[i].UnixNano())
	}
	log.Errorln("query", query)
	res, err := queryInfluxDB(datastore.influxClient, query)
	if err != nil {
		return r, err
	}

	if len(res) > 0 {
		for i := range res {
			if len(res[i].Series) > 0 {
				log.Infof("parse %v trades...", len(res[i].Series[0].Values))
				for _, row := range res[i].Series[0].Values {
					t := parseTrade(row, returnBasetoken)
					if t != nil {
						r = append(r, *t)
					}
				}
				log.Info("...done parsing.")
			}
		}
	} else {
		log.Errorf("Empty response GetTradesByExchangesBatched for %s \n", quoteasset.Symbol)
		return nil, fmt.Errorf("no trades found")
	}

	return r, nil
}

// GetAllTrades returns at most @maxTrades trades from influx with timestamp > @t. Only used by replayInflux option.
func (datastore *DB) GetAllTrades(t time.Time, maxTrades int) ([]dia.Trade, error) {
	var r []dia.Trade
	// TO DO: Substitute select * with precise statment select estimatedUSDPrice, source,...
	q := fmt.Sprintf("SELECT time, estimatedUSDPrice, exchange, foreignTradeID, pair, price,symbol, volume,verified,basetokenblockchain,basetokenaddress  FROM %s WHERE time > %d LIMIT %d", influxDbTradesTable, t.Unix()*1000000000, maxTrades)
	log.Debug(q)
	res, err := queryInfluxDB(datastore.influxClient, q)
	if err != nil {
		log.Errorln("GetLastTrades", err)
		return r, err
	}
	if len(res) > 0 && len(res[0].Series) > 0 {
		for _, row := range res[0].Series[0].Values {
			t := parseTrade(row, false)
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

// GetLastTrades returns the last @maxTrades of @asset on @exchange.
// If exchange is empty string it returns trades from all exchanges.
// If fullAsset=true, blockchain and address of both involved assets is returned as well
func (datastore *DB) GetLastTrades(asset dia.Asset, exchange string, maxTrades int, fullAsset bool) ([]dia.Trade, error) {
	var r []dia.Trade
	var queryString string
	var q string
	if exchange == "" {
		queryString = "SELECT estimatedUSDPrice,\"exchange\",foreignTradeID,\"pair\",price,\"symbol\",volume,\"verified\"," +
			"\"basetokenblockchain\",\"basetokenaddress\"" +
			" FROM %s WHERE time<now() AND time>now()-30d AND quotetokenaddress='%s' AND quotetokenblockchain='%s' AND estimatedUSDPrice>0 ORDER BY DESC LIMIT %d"
		q = fmt.Sprintf(queryString, influxDbTradesTable, asset.Address, asset.Blockchain, maxTrades)
	} else {
		queryString = "SELECT estimatedUSDPrice,\"exchange\",foreignTradeID,\"pair\",price,\"symbol\",volume,\"verified\"," +
			"\"basetokenblockchain\",\"basetokenaddress\"" +
			" FROM %s WHERE time<now() AND time>now()-30d AND exchange='%s' AND quotetokenaddress='%s' AND quotetokenblockchain='%s' AND estimatedUSDPrice>0 ORDER BY DESC LIMIT %d"
		q = fmt.Sprintf(queryString, influxDbTradesTable, exchange, asset.Address, asset.Blockchain, maxTrades)
	}
	res, err := queryInfluxDB(datastore.influxClient, q)
	if err != nil {
		log.Errorln("GetLastTrades", err)
		return r, err
	}

	if len(res) > 0 && len(res[0].Series) > 0 {
		for _, row := range res[0].Series[0].Values {
			t := parseTrade(row, fullAsset)
			if t != nil {
				t.QuoteToken = asset
				r = append(r, *t)
			}
		}
	} else {
		err = fmt.Errorf("Empty response for %s on %s", asset.Symbol, exchange)
		log.Error(err)
		return r, err
	}
	return r, nil
}

func (datastore *DB) GetNumTrades(exchange string) (numTrades int64, err error) {
	queryString := "SELECT COUNT(*) FROM %s WHERE exchange='%s' AND time > now() - 24h AND time<now()"
	q := fmt.Sprintf(queryString, influxDbTradesTable, exchange)

	res, err := queryInfluxDB(datastore.influxClient, q)
	if err != nil {
		log.Errorln("Get24HVolumePerExchange ", err)
		return
	}

	if len(res) > 0 && len(res[0].Series) > 0 {
		vol, ok := res[0].Series[0].Values[0][1].(json.Number)
		if ok {
			numTrades, err = vol.Int64()
			if err != nil {
				return numTrades, err
			}
		}
	}
	return
}
