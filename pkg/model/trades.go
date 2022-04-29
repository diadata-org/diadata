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
	var r []dia.Trade
	subQuery := ""
	if len(exchanges) > 0 {
		for _, exchange := range exchanges {
			subQuery = subQuery + fmt.Sprintf("%s|", exchange)
		}
		subQuery = "and exchange =~ /" + strings.TrimRight(subQuery, "|") + "/"
	}
	query := fmt.Sprintf("SELECT time,estimatedUSDPrice,verified,foreignTradeID,pair,price,symbol,volume,verified,basetokenblockchain,basetokenaddress FROM %s WHERE quotetokenaddress='%s' and quotetokenblockchain='%s' %s AND estimatedUSDPrice > 0 AND time >= %d AND time <= %d ", influxDbTradesTable, asset.Address, asset.Blockchain, subQuery, startTime.UnixNano(), endTime.UnixNano())

	log.Infoln("GetTradesByExchanges Query", query)
	res, err := queryInfluxDB(datastore.influxClient, query)
	if err != nil {
		return r, err
	}

	if len(res) > 0 && len(res[0].Series) > 0 {
		for _, row := range res[0].Series[0].Values {
			t := parseTrade(row, false)
			if t != nil {
				r = append(r, *t)
			}
		}
	} else {
		log.Errorf("Empty response GetTradesByExchanges for %s \n", asset.Symbol)
		return nil, fmt.Errorf("no trades found")
	}
	return r, nil
}

// GetTradesByExchangesBatched executes multiple select queries on the trades table in one batch.
// The time ranges of the queries are given by the intervals [startTimes[i], endTimes[i]].
func (datastore *DB) GetTradesByExchangesBatched(asset dia.Asset, exchanges []string, startTimes, endTimes []time.Time) ([]dia.Trade, error) {
	return datastore.GetTradesByExchangesBatchedFull(asset, exchanges, false, startTimes, endTimes)
}

// GetTradesByExchangesBatchedFull executes multiple select queries on the trades table in one batch.
// The time ranges of the queries are given by the intervals [startTimes[i], endTimes[i]].
func (datastore *DB) GetTradesByExchangesBatchedFull(asset dia.Asset, exchanges []string, returnBasetoken bool, startTimes, endTimes []time.Time) ([]dia.Trade, error) {
	var r []dia.Trade
	if len(startTimes) != len(endTimes) {
		return []dia.Trade{}, errors.New("number of start times must equal number of end times.")
	}
	var query string
	for i := range startTimes {
		subQuery := ""
		if len(exchanges) > 0 {
			for _, exchange := range exchanges {
				subQuery = subQuery + fmt.Sprintf("%s|", exchange)
			}
			subQuery = "and exchange =~ /" + strings.TrimRight(subQuery, "|") + "/"
		}
		query = query + fmt.Sprintf("SELECT time,estimatedUSDPrice,exchange,foreignTradeID,pair,price,symbol,volume,verified,basetokenblockchain,basetokenaddress FROM %s WHERE quotetokenaddress='%s' AND quotetokenblockchain='%s' %s AND estimatedUSDPrice > 0 AND time > %d AND time <= %d ;", influxDbTradesTable, asset.Address, asset.Blockchain, subQuery, startTimes[i].UnixNano(), endTimes[i].UnixNano())
	}

	res, err := queryInfluxDB(datastore.influxClient, query)
	if err != nil {
		return r, err
	}

	if len(res) > 0 {
		for i := range res {
			if len(res[i].Series) > 0 {
				for _, row := range res[i].Series[0].Values {
					t := parseTrade(row, returnBasetoken)
					if t != nil {
						r = append(r, *t)
					}
				}
			}
		}
	} else {
		log.Errorf("Empty response GetTradesByExchangesBatched for %s \n", asset.Symbol)
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
