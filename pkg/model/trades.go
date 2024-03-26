package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/go-redis/redis"
	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
)

const (
	// TO DO: Switch Redis Cache constants to env vars?
	TRADES_TIMEOUT_REDIS_24H        = 60 * 60 * 24
	TRADES_TIMEOUT_BUFFER           = 60 * 10
	TRADED_ASSETS_KEY               = "tradedAssets"
	TRADE_EXCHANGEPAIR_PREFIX_REDIS = "exchangepairKeys_"
	TRADE_PREFIX_REDIS              = "trades_"
)

// SaveTradeInflux stores a trade in influx. Flushed when more than maxPoints in batch.
// Wrapper around SaveTradeInfluxToTable.
func (datastore *DB) SaveTradeInflux(t *dia.Trade) error {
	return datastore.SaveTradeInfluxToTable(t, influxDbTradesTable)
}

// SaveTradeInfluxToTable stores a trade in influx into @table.
// Flushed when more than maxPoints in batch.
func (datastore *DB) SaveTradeInfluxToTable(t *dia.Trade, table string) error {

	// Create a point and add to batch
	tags := map[string]string{
		"symbol":               EscapeReplacer.Replace(t.Symbol),
		"pair":                 t.Pair,
		"exchange":             t.Source,
		"verified":             strconv.FormatBool(t.VerifiedPair),
		"quotetokenaddress":    t.QuoteToken.Address,
		"basetokenaddress":     t.BaseToken.Address,
		"quotetokenblockchain": t.QuoteToken.Blockchain,
		"basetokenblockchain":  t.BaseToken.Blockchain,
		"pooladdress":          t.PoolAddress,
	}
	fields := map[string]interface{}{
		"price":             t.Price,
		"volume":            t.Volume,
		"estimatedUSDPrice": t.EstimatedUSDPrice,
		"foreignTradeID":    t.ForeignTradeID,
	}

	pt, err := clientInfluxdb.NewPoint(table, tags, fields, t.Time)
	if err != nil {
		log.Errorln("NewTradeInflux:", err)
	} else {
		datastore.addPoint(pt)
	}

	return err
}

// GetTradeInflux returns the latest trade of @asset on @exchange before @timestamp in the time-range [endtime-window, endtime].
func (datastore *DB) GetTradeInflux(asset dia.Asset, exchange string, endtime time.Time, window time.Duration) (*dia.Trade, error) {
	starttime := endtime.Add(-window)
	retval := dia.Trade{}
	var q string
	if exchange != "" {
		queryString := "SELECT estimatedUSDPrice,\"exchange\",foreignTradeID,\"pair\",price,\"symbol\",volume FROM %s WHERE quotetokenaddress='%s' AND quotetokenblockchain='%s' AND exchange='%s' AND time >= %d AND time < %d ORDER BY DESC LIMIT 1"
		q = fmt.Sprintf(queryString, influxDbTradesTable, asset.Address, asset.Blockchain, exchange, starttime.UnixNano(), endtime.UnixNano())
	} else {
		queryString := "SELECT estimatedUSDPrice,\"exchange\",foreignTradeID,\"pair\",price,\"symbol\",volume FROM %s WHERE quotetokenaddress='%s' AND quotetokenblockchain='%s' AND time >= %d AND time < %d ORDER BY DESC LIMIT 1"
		q = fmt.Sprintf(queryString, influxDbTradesTable, asset.Address, asset.Blockchain, starttime.UnixNano(), endtime.UnixNano())
	}

	/// TODO
	res, err := queryInfluxDB(datastore.influxClient, q)
	if err != nil {
		return &retval, err
	}
	if len(res) > 0 && len(res[0].Series) > 0 {
		for i := 0; i < len(res[0].Series[0].Values); i++ {
			retval.Time, err = time.Parse(time.RFC3339, res[0].Series[0].Values[i][0].(string))
			if err != nil {
				return &retval, err
			}
			retval.EstimatedUSDPrice, err = res[0].Series[0].Values[i][1].(json.Number).Float64()
			if err != nil {
				return &retval, err
			}
			retval.Source = res[0].Series[0].Values[i][2].(string)
			retval.ForeignTradeID = res[0].Series[0].Values[i][3].(string)
			retval.Pair = res[0].Series[0].Values[i][4].(string)
			retval.Price, err = res[0].Series[0].Values[i][5].(json.Number).Float64()
			if err != nil {
				return &retval, err
			}
			retval.Symbol = res[0].Series[0].Values[i][6].(string)
			retval.Volume, err = res[0].Series[0].Values[i][7].(json.Number).Float64()
			if err != nil {
				return &retval, err
			}
		}
	} else {
		return &retval, errors.New("parsing trade from database")
	}
	return &retval, nil
}

// GetOldTradesFromInflux returns all recorded trades from @table done on @exchange between @timeInit and @timeFinal
// where the time interval is closed on the left and open on the right side.
// If @exchange is empty, trades across all exchanges are returned.
// If @verified is true, address and blockchain are also parsed for both assets.
func (datastore *DB) GetOldTradesFromInflux(table string, exchange string, verified bool, timeInit, timeFinal time.Time) ([]dia.Trade, error) {
	allTrades := []dia.Trade{}
	var queryString, query, addQueryString string
	if verified {
		addQueryString = ",\"quotetokenaddress\",\"basetokenaddress\",\"quotetokenblockchain\",\"basetokenblockchain\",\"verified\""
	}
	if exchange == "" {
		queryString = "SELECT estimatedUSDPrice,\"exchange\",foreignTradeID,\"pair\",price,\"symbol\",volume" +
			addQueryString +
			" FROM %s WHERE time>=%d and time<%d order by asc"
		query = fmt.Sprintf(queryString, table, timeInit.UnixNano(), timeFinal.UnixNano())
	} else {
		queryString = "SELECT estimatedUSDPrice,\"exchange\",foreignTradeID,\"pair\",price,\"symbol\",volume" +
			addQueryString +
			" FROM %s WHERE exchange='%s' and time>=%d and time<%d order by asc"
		query = fmt.Sprintf(queryString, table, exchange, timeInit.UnixNano(), timeFinal.UnixNano())
	}
	res, err := queryInfluxDB(datastore.influxClient, query)
	if err != nil {
		log.Error("influx query: ", err)
		return allTrades, err
	}

	log.Info("query: ", query)

	if len(res) > 0 && len(res[0].Series) > 0 {
		for i := 0; i < len(res[0].Series[0].Values); i++ {
			var trade dia.Trade
			trade.Time, err = time.Parse(time.RFC3339, res[0].Series[0].Values[i][0].(string))
			if err != nil {
				return allTrades, err
			}
			trade.EstimatedUSDPrice, err = res[0].Series[0].Values[i][1].(json.Number).Float64()
			if err != nil {
				return allTrades, err
			}
			if res[0].Series[0].Values[i][2] != nil {
				trade.Source = res[0].Series[0].Values[i][2].(string)
			}
			if res[0].Series[0].Values[i][3] != nil {
				trade.ForeignTradeID = res[0].Series[0].Values[i][3].(string)
			}
			if res[0].Series[0].Values[i][4] != nil {
				trade.Pair = res[0].Series[0].Values[i][4].(string)
			}
			trade.Price, err = res[0].Series[0].Values[i][5].(json.Number).Float64()
			if err != nil {
				return allTrades, err
			}
			if res[0].Series[0].Values[i][6] == nil {
				continue
			}
			if res[0].Series[0].Values[i][6] != nil {
				trade.Symbol = res[0].Series[0].Values[i][6].(string)
			}
			trade.Volume, err = res[0].Series[0].Values[i][7].(json.Number).Float64()
			if err != nil {
				return allTrades, err
			}
			if verified {
				if res[0].Series[0].Values[i][8] != nil {
					trade.QuoteToken.Address = res[0].Series[0].Values[i][8].(string)
				}
				if res[0].Series[0].Values[i][9] != nil {
					trade.BaseToken.Address = res[0].Series[0].Values[i][9].(string)
				}
				if res[0].Series[0].Values[i][10] != nil {
					trade.QuoteToken.Blockchain = res[0].Series[0].Values[i][10].(string)
				}
				if res[0].Series[0].Values[i][11] != nil {
					trade.BaseToken.Blockchain = res[0].Series[0].Values[i][11].(string)
				}
				verifiedPair, ok := res[0].Series[0].Values[i][12].(string)
				if ok {
					trade.VerifiedPair, err = strconv.ParseBool(verifiedPair)
					if err != nil {
						log.Error("parse verified pair boolean: ", err)
					}
				}
			}
			allTrades = append(allTrades, trade)
		}
	} else {
		return allTrades, errors.New("no trades in time range")
	}
	return allTrades, nil
}

// parseFullTrade parses a trade as retreived from influx. If fullAsset=true blockchain and address of
// the corresponding asset is returned as well.
func parseFullTrade(row []interface{}) *dia.Trade {
	if len(row) > 13 {
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
			basetokenblockchain, o := row[9].(string)
			if !o {
				log.Errorln("error on parsing row 9", row)
			}
			basetokenaddress, o := row[10].(string)
			if !o {
				log.Errorln("error on parsing row 10", row)
			}
			quotetokenblockchain, o := row[11].(string)
			if !o {
				log.Errorln("error on parsing row 11", row)
			}
			quotetokenaddress, o := row[12].(string)
			if !o {
				log.Errorln("error on parsing row 12", row)
			}
			pooladdress, _ := row[13].(string)

			trade := dia.Trade{
				Symbol:            symbol,
				Pair:              pair,
				QuoteToken:        dia.Asset{Address: quotetokenaddress, Blockchain: quotetokenblockchain},
				BaseToken:         dia.Asset{Address: basetokenaddress, Blockchain: basetokenblockchain},
				PoolAddress:       pooladdress,
				Time:              t,
				Source:            source,
				EstimatedUSDPrice: estimatedUSDPrice,
				Price:             price,
				Volume:            volume,
				ForeignTradeID:    foreignTradeID,
				VerifiedPair:      verified,
			}

			return &trade
		}

	}
	return nil
}

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

func (datastore *DB) GetTradesByExchangesAndBaseAssets(asset dia.Asset, baseassets []dia.Asset, exchanges []string, startTime, endTime time.Time, maxTrades int) ([]dia.Trade, error) {
	return datastore.GetTradesByExchangesFull(asset, baseassets, exchanges, false, startTime, endTime, maxTrades)
}

func (datastore *DB) GetTradesByExchangesFull(
	asset dia.Asset,
	baseassets []dia.Asset,
	exchanges []string,
	returnBasetoken bool,
	startTime time.Time,
	endTime time.Time,
	maxTrades int,
) ([]dia.Trade, error) {
	var r []dia.Trade
	subQuery := ""
	subQueryBase := ""
	if len(exchanges) > 0 {
		for _, exchange := range exchanges {
			subQuery = subQuery + fmt.Sprintf("%s|", exchange)
		}
		subQuery = "AND exchange =~ /" + strings.TrimRight(subQuery, "|") + "/"

		if len(baseassets) > 0 {
			for i, baseasset := range baseassets {
				if i == 0 {
					subQueryBase = subQueryBase + fmt.Sprintf(` AND ((basetokenaddress='%s' AND basetokenblockchain='%s')`, baseasset.Address, baseasset.Blockchain)

				} else {
					subQueryBase = subQueryBase + fmt.Sprintf(` OR (basetokenaddress='%s' AND basetokenblockchain='%s')`, baseasset.Address, baseasset.Blockchain)
				}
			}
			subQueryBase = subQueryBase + ") "
		}
	}
	query := fmt.Sprintf(`
	SELECT time,estimatedUSDPrice,exchange,foreignTradeID,pair,price,symbol,volume,verified,basetokenblockchain,basetokenaddress 
	FROM %s 
	WHERE (quotetokenaddress='%s' and quotetokenblockchain='%s') %s %s 
	AND estimatedUSDPrice > 0 
	AND time > %d AND time <= %d `,
		influxDbTradesTable, asset.Address, asset.Blockchain, subQuery, subQueryBase, startTime.UnixNano(), endTime.UnixNano())
	if maxTrades > 0 {
		query += fmt.Sprintf("ORDER BY DESC LIMIT %d ", maxTrades)
	}
	log.Info("query: ", query)
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
func (datastore *DB) GetTradesByExchangesBatched(
	quoteasset dia.Asset,
	baseassets []dia.Asset,
	exchanges []string,
	startTimes []time.Time,
	endTimes []time.Time,
	maxTrades int,
) ([]dia.Trade, error) {
	return datastore.GetTradesByExchangesBatchedFull(quoteasset, baseassets, exchanges, false, startTimes, endTimes, maxTrades)
}

// GetTradesByExchangesBatchedFull executes multiple select queries on the trades table in one batch.
// The time ranges of the queries are given by the intervals [startTimes[i], endTimes[i]].
func (datastore *DB) GetTradesByExchangesBatchedFull(
	quoteasset dia.Asset,
	baseassets []dia.Asset,
	exchanges []string,
	returnBasetoken bool,
	startTimes []time.Time,
	endTimes []time.Time,
	maxTrades int,
) ([]dia.Trade, error) {

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
					subQueryBase = subQueryBase + fmt.Sprintf(` and ((basetokenaddress='%s' and basetokenblockchain='%s')`, baseasset.Address, baseasset.Blockchain)

				} else {
					subQueryBase = subQueryBase + fmt.Sprintf(` or (basetokenaddress='%s' and basetokenblockchain='%s')`, baseasset.Address, baseasset.Blockchain)
				}

			}
			subQueryBase = subQueryBase + ") "

		}
		query = query + fmt.Sprintf(`
		SELECT time,estimatedUSDPrice,exchange,foreignTradeID,pair,price,symbol,volume,verified,basetokenblockchain,basetokenaddress 
		FROM %s 
		WHERE (quotetokenaddress='%s' AND quotetokenblockchain='%s') %s %s 
		AND estimatedUSDPrice > 0 
		AND time > %d AND time <= %d ; `,
			influxDbTradesTable, quoteasset.Address, quoteasset.Blockchain, subQuery, subQueryBase, startTimes[i].UnixNano(), endTimes[i].UnixNano())
	}
	log.Info("query: ", query)
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

// GetxcTradesByExchangesBatched executes multiple select queries on the trades table in one batch.
// The time ranges of the queries are given by the intervals [startTimes[i], endTimes[i]].
func (datastore *DB) GetxcTradesByExchangesBatched(
	quoteassets []dia.Asset,
	exchanges []string,
	startTimes []time.Time,
	endTimes []time.Time,
) ([]dia.Trade, error) {

	var r []dia.Trade
	if len(startTimes) != len(endTimes) {
		return []dia.Trade{}, errors.New("number of start times must equal number of end times.")
	}
	var query string
	for i := range startTimes {
		subQueryExchanges := ""
		subQueryAssets := ""
		if len(exchanges) > 0 {
			for _, exchange := range exchanges {
				subQueryExchanges = subQueryExchanges + fmt.Sprintf("%s|", exchange)
			}
			subQueryExchanges = "AND exchange =~ /" + strings.TrimRight(subQueryExchanges, "|") + "/"
		}

		if len(quoteassets) > 0 {
			for i, quoteasset := range quoteassets {
				if i == 0 {
					subQueryAssets = subQueryAssets + fmt.Sprintf(` AND ((quotetokenaddress='%s' AND quotetokenblockchain='%s')`, quoteasset.Address, quoteasset.Blockchain)

				} else {
					subQueryAssets = subQueryAssets + fmt.Sprintf(` OR (quotetokenaddress='%s' AND quotetokenblockchain='%s')`, quoteasset.Address, quoteasset.Blockchain)
				}

			}
			subQueryAssets = subQueryAssets + ") "

		}
		query = query + fmt.Sprintf(`
		SELECT time,estimatedUSDPrice,exchange,foreignTradeID,pair,price,symbol,volume,verified,basetokenblockchain,basetokenaddress
		FROM %s 
		WHERE estimatedUSDPrice > 0 
		AND time > %d AND time <= %d 
		%s %s ;`,
			influxDbTradesTable, startTimes[i].UnixNano(), endTimes[i].UnixNano(), subQueryExchanges, subQueryAssets)
	}
	res, err := queryInfluxDB(datastore.influxClient, query)
	if err != nil {
		return r, err
	}

	if len(res) > 0 {
		for i := range res {
			if len(res[i].Series) > 0 {
				log.Infof("parse %v trades...", len(res[i].Series[0].Values))
				for _, row := range res[i].Series[0].Values {
					t := parseTrade(row, false)
					if t != nil {
						r = append(r, *t)
					}
				}
				log.Info("...done parsing.")
			}
		}
	} else {
		log.Error("Empty response GetxcTradesByExchangesBatched")
		return nil, fmt.Errorf("no trades found")
	}

	return r, nil
}

// GetTradesByFeedSelection returns all trades with restrictions given by the struct @feedselection.
func (datastore *DB) GetTradesByFeedSelection(
	feedselection []dia.FeedSelection,
	starttimes []time.Time,
	endtimes []time.Time,
	limit int,
) ([]dia.Trade, error) {
	var (
		query string
		r     []dia.Trade
	)

	if len(starttimes) != len(endtimes) {
		return []dia.Trade{}, errors.New("number of start times must equal number of end times.")
	}

	for i := range starttimes {
		query += fmt.Sprintf(`
		SELECT time,estimatedUSDPrice,exchange,foreignTradeID,pair,price,symbol,volume,verified,basetokenblockchain,basetokenaddress,quotetokenblockchain,quotetokenaddress,pooladdress  
		FROM %s 
		WHERE ( `,
			influxDbTradesTable,
		)

		// --------------------- Iterate over assets. ---------------------
		for i, item := range feedselection {
			if i > 0 {
				query += " OR "
			}
			//  ---------------------Iterate over exchanges. ---------------------
			var exchangeQuery string
			for j, exchangepairs := range item.Exchangepairs {
				if j == 0 {
					exchangeQuery += " AND ("
				} else {
					exchangeQuery += " OR  "
				}

				// --------------------- Iterate over pairs/pools. ---------------------
				var pairsQuery string
				if exchangepairs.Exchange.Centralized {
					for k, pair := range exchangepairs.Pairs {
						if k == 0 {
							pairsQuery += " AND ("
						} else {
							pairsQuery += " OR "
						}
						pairsQuery += fmt.Sprintf(`
							( quotetokenaddress='%s' AND quotetokenblockchain='%s' AND basetokenaddress='%s' and basetokenblockchain='%s')
							`,
							pair.QuoteToken.Address,
							pair.QuoteToken.Blockchain,
							pair.BaseToken.Address,
							pair.BaseToken.Blockchain,
						)
					}
				} else {
					for k, pool := range exchangepairs.Pools {
						if k == 0 {
							pairsQuery += " AND ("
						} else {
							pairsQuery += " OR "
						}
						pairsQuery += fmt.Sprintf(" pooladdress='%s' ", pool.Address)
					}
				}
				if len(exchangepairs.Pairs) > 0 || len(exchangepairs.Pools) > 0 {
					pairsQuery += " ) "
				}

				if exchangepairs.Exchange.Name != "" {
					exchangeQuery += fmt.Sprintf(`(exchange='%s' %s)`, exchangepairs.Exchange.Name, pairsQuery)
				} else {
					// Take into account trades on all exchanges.
					exchangeQuery += fmt.Sprintf(`exchange=~/./ %s`, pairsQuery)
				}
			}
			if len(item.Exchangepairs) > 0 {
				exchangeQuery += " ) "
			}

			// Main query for trades by asset.
			query += fmt.Sprintf(`
			( (quotetokenaddress='%s' AND quotetokenblockchain='%s') %s ) 
			`,
				item.Asset.Address,
				item.Asset.Blockchain,
				exchangeQuery,
			)
		}

		// The bracket closes the main statement from the first WHERE clause.
		var limitQuery string
		if len(starttimes) == 1 && limit > 0 {
			limitQuery = fmt.Sprintf(" ORDER BY DESC LIMIT %v", limit)
		}
		query += fmt.Sprintf(`
		 )	
		AND estimatedUSDPrice > 0
		AND time > %d
		AND time < %d %s;`,
			starttimes[i].UnixNano(),
			endtimes[i].UnixNano(),
			limitQuery,
		)
	}

	res, err := queryInfluxDB(datastore.influxClient, query)
	if err != nil {
		return r, err
	}

	if len(res) > 0 {
		for i := range res {
			if len(res[i].Series) > 0 {
				log.Infof("parse %v trades...", len(res[i].Series[0].Values))
				for _, row := range res[i].Series[0].Values {
					t := parseFullTrade(row)
					if t != nil {
						r = append(r, *t)
					}
				}
				log.Info("...done parsing.")
			}
		}
	} else {
		return nil, fmt.Errorf("No trades found.")
	}

	return r, nil
}

// GetAggregatedFeedSelection returns aggregated quantities with restrictions given by the struct @feedselection.
func (datastore *DB) GetAggregatedFeedSelection(
	feedselection []dia.FeedSelection,
	starttime time.Time,
	endtime time.Time,
	tradeVolumeThreshold float64,
) ([]dia.FeedSelectionAggregated, error) {
	var (
		query                   string
		feedSelectionAggregated []dia.FeedSelectionAggregated
	)

	if starttime.After(endtime) {
		return feedSelectionAggregated, errors.New("starttime is after endtime.")
	}

	query += fmt.Sprintf(`
		SELECT SUM(abs),COUNT(multiplication),LAST(multiplication)
		FROM (
			SELECT ABS(estimatedUSDPrice*volume),estimatedUSDPrice AS multiplication
			FROM %s 
			WHERE ( `,
		influxDbTradesTable,
	)

	// --------------------- Iterate over assets. ---------------------
	for i, item := range feedselection {
		if i > 0 {
			query += " OR "
		}
		//  ---------------------Iterate over exchanges. ---------------------
		var exchangeQuery string
		for j, exchangepairs := range item.Exchangepairs {
			if j == 0 {
				exchangeQuery += " AND ("
			} else {
				exchangeQuery += " OR  "
			}

			// --------------------- Iterate over pairs/pools. ---------------------
			var pairsQuery string
			if exchangepairs.Exchange.Centralized {
				for k, pair := range exchangepairs.Pairs {
					if k == 0 {
						pairsQuery += " AND ("
					} else {
						pairsQuery += " OR "
					}
					pairsQuery += fmt.Sprintf(`
							( quotetokenaddress='%s' AND quotetokenblockchain='%s' AND basetokenaddress='%s' and basetokenblockchain='%s')
							`,
						pair.QuoteToken.Address,
						pair.QuoteToken.Blockchain,
						pair.BaseToken.Address,
						pair.BaseToken.Blockchain,
					)
				}
			} else {
				for k, pool := range exchangepairs.Pools {
					if k == 0 {
						pairsQuery += " AND ("
					} else {
						pairsQuery += " OR "
					}
					pairsQuery += fmt.Sprintf(" pooladdress='%s' ", pool.Address)
				}
			}
			if len(exchangepairs.Pairs) > 0 || len(exchangepairs.Pools) > 0 {
				pairsQuery += " ) "
			}

			if exchangepairs.Exchange.Name != "" {
				exchangeQuery += fmt.Sprintf(`(exchange='%s' %s)`, exchangepairs.Exchange.Name, pairsQuery)
			} else {
				// Take into account trades on all exchanges.
				exchangeQuery += fmt.Sprintf(`exchange=~/./ %s`, pairsQuery)
			}
		}
		if len(item.Exchangepairs) > 0 {
			exchangeQuery += " ) "
		}

		// Main query for trades by asset.
		query += fmt.Sprintf(`
			( (quotetokenaddress='%s' AND quotetokenblockchain='%s') %s ) 
			`,
			item.Asset.Address,
			item.Asset.Blockchain,
			exchangeQuery,
		)
	}

	query += fmt.Sprintf(`	
		)
		AND estimatedUSDPrice > 0
		AND time >= %d
		AND time < %d)
		WHERE multiplication>%v `,
		starttime.UnixNano(),
		endtime.UnixNano(),
		tradeVolumeThreshold,
	)
	query += ` GROUP BY "exchange","quotetokenaddress","quotetokenblockchain","basetokenaddress","basetokenblockchain","pooladdress","symbol","pair"`

	res, err := queryInfluxDB(datastore.influxClient, query)
	if err != nil {
		return feedSelectionAggregated, err
	}

	if len(res) > 0 && len(res[0].Series) > 0 {
		for _, row := range res[0].Series {
			if len(row.Values[0]) > 1 {
				var fsa dia.FeedSelectionAggregated
				fsa.Exchange = row.Tags["exchange"]
				fsa.Quotetoken.Address = row.Tags["quotetokenaddress"]
				fsa.Quotetoken.Blockchain = row.Tags["quotetokenblockchain"]
				fsa.Basetoken.Address = row.Tags["basetokenaddress"]
				fsa.Basetoken.Blockchain = row.Tags["basetokenblockchain"]
				// Parse symbol of basetoken
				pairSymbols, err := dia.GetPairSymbols(dia.ExchangePair{
					Symbol:      row.Tags["symbol"],
					ForeignName: row.Tags["pair"],
					Exchange:    row.Tags["exchange"],
				})
				if err != nil {
					log.Error("Get pair symbols: ", err)
				} else if len(pairSymbols) > 1 {
					fsa.Quotetoken.Symbol = strings.ToUpper(pairSymbols[0])
					fsa.Basetoken.Symbol = strings.ToUpper(pairSymbols[1])
				}
				fsa.Pooladdress = row.Tags["pooladdress"]
				fsa.Volume, err = row.Values[0][1].(json.Number).Float64()
				if err != nil {
					log.Error("cast float64: ", err)
				}
				tradescount, err := row.Values[0][2].(json.Number).Int64()
				if err != nil {
					log.Error("cast int64: ", err)
				}
				fsa.TradesCount = int32(tradescount)
				fsa.LastPrice, err = row.Values[0][3].(json.Number).Float64()
				if err != nil {
					log.Error("cast float64: ", err)
				}
				fsa.Starttime = starttime
				fsa.Endtime = endtime
				feedSelectionAggregated = append(feedSelectionAggregated, fsa)
			}
		}
	} else {
		return feedSelectionAggregated, fmt.Errorf("No trades found.")
	}

	// Sort response by volume.
	sort.Slice(feedSelectionAggregated, func(m, n int) bool {
		return feedSelectionAggregated[m].Volume > feedSelectionAggregated[n].Volume
	})

	return feedSelectionAggregated, nil
}

// GetTradesByExchangepairs returns all trades where either of the following is fulfilled.
// 1. The exchange is a key of @exchangepairMap AND the pair is in the corresponding slice @[]dia.Pair.
// 2. The exchange is a key of @exchangepoolMap AND the pool is in the corresponding slice @[]string.
func (datastore *DB) GetTradesByExchangepairs(exchangepairMap map[string][]dia.Pair, exchangepoolMap map[string][]string, starttime time.Time, endtime time.Time) ([]dia.Trade, error) {
	var (
		query string
		r     []dia.Trade
	)

	query = fmt.Sprintf(`
		SELECT time,estimatedUSDPrice,exchange,foreignTradeID,pair,price,symbol,volume,verified,basetokenblockchain,basetokenaddress,quotetokenblockchain,quotetokenaddress  
		FROM %s 
		WHERE ( `,
		influxDbTradesTable,
	)

	// Iterate over centralized exchanges.
	var CEXCount int
	for exchange := range exchangepairMap {
		if CEXCount != 0 {
			query += " OR "
		}

		// If, in addition to exchanges, pairs are also given, make pairs subquery for each exchange.
		var pairsQuery string
		if len(exchangepairMap) > 0 {
			pairsQuery += " AND ( "
			for i, pair := range exchangepairMap[exchange] {
				if i != 0 {
					pairsQuery += " OR "
				}
				pairsQuery += fmt.Sprintf(`
				( quotetokenaddress='%s' AND quotetokenblockchain='%s' AND basetokenaddress='%s' and basetokenblockchain='%s')
				`,
					pair.QuoteToken.Address,
					pair.QuoteToken.Blockchain,
					pair.BaseToken.Address,
					pair.BaseToken.Blockchain,
				)
			}
			pairsQuery += " ) "
		}

		// Main query for trades by exchange.
		query += fmt.Sprintf(" ( exchange='%s' %s ) ", exchange, pairsQuery)
		CEXCount++
	}

	// Iterate over decentralized exchanges.
	var DEXCount int
	for exchange := range exchangepoolMap {
		if DEXCount != 0 || len(exchangepairMap) > 0 {
			query += " OR "
		}

		// If, in addition to exchanges, pools are also given, make pool subquery for each exchange.
		var poolsQuery string
		if len(exchangepairMap) > 0 {
			poolsQuery += " AND ( "
			for i, pooladdress := range exchangepoolMap[exchange] {
				if i != 0 {
					poolsQuery += " OR "
				}
				poolsQuery += fmt.Sprintf("( pooladdress='%s' )", pooladdress)
			}
			poolsQuery += " ) "
		}

		// Main query for trades by exchange.
		query += fmt.Sprintf(" ( exchange='%s' %s ) ", exchange, poolsQuery)
		DEXCount++
	}

	// The bracket closes the main statement from the first WHERE clause.
	query += fmt.Sprintf(`
		 )	
		AND estimatedUSDPrice > 0
		AND time > %d
		AND time < %d`,
		starttime.UnixNano(),
		endtime.UnixNano(),
	)

	log.Info("query: ", query)
	res, err := queryInfluxDB(datastore.influxClient, query)
	if err != nil {
		return r, err
	}

	if len(res) > 0 {
		for i := range res {
			if len(res[i].Series) > 0 {
				log.Infof("parse %v trades...", len(res[i].Series[0].Values))
				for _, row := range res[i].Series[0].Values {
					t := parseFullTrade(row)
					if t != nil {
						r = append(r, *t)
					}
				}
				log.Info("...done parsing.")
			}
		}
	} else {
		log.Error("Empty response GetxcTradesByExchangesBatched")
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
		log.Errorln("GetAllTrades", err)
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

// GetLastTrades returns the last @maxTrades of @asset on @exchange before @timestamp.
// If exchange is empty string it returns trades from all exchanges.
// If fullAsset=true, blockchain and address of both involved assets is returned as well
func (datastore *DB) GetLastTrades(asset dia.Asset, exchange string, timestamp time.Time, maxTrades int, fullAsset bool) ([]dia.Trade, error) {
	var (
		r           []dia.Trade
		queryString string
		q           string
	)

	if exchange == "" {
		queryString = `
		SELECT estimatedUSDPrice,"exchange",foreignTradeID,"pair",price,"symbol",volume,"verified","basetokenblockchain","basetokenaddress" 
		FROM %s 
		WHERE time<%d 
		AND time>%d-10d 
		AND quotetokenaddress='%s' 
		AND quotetokenblockchain='%s' 
		AND estimatedUSDPrice>0 
		ORDER BY DESC LIMIT %d
		`
		q = fmt.Sprintf(queryString, influxDbTradesTable, timestamp.UnixNano(), timestamp.UnixNano(), asset.Address, asset.Blockchain, maxTrades)
	} else if (dia.Asset{}) == asset {
		queryString = `
		SELECT estimatedUSDPrice,"exchange",foreignTradeID,"pair",price,"symbol",volume,"verified","basetokenblockchain","basetokenaddress" 
		FROM %s 
		WHERE time<%d 
		AND time>%d-10d 
		AND exchange='%s' 
		AND estimatedUSDPrice>0 
		ORDER BY DESC LIMIT %d
		`
		q = fmt.Sprintf(queryString, influxDbTradesTable, timestamp.UnixNano(), timestamp.UnixNano(), exchange, maxTrades)
	} else {
		queryString = `
		SELECT estimatedUSDPrice,"exchange",foreignTradeID,"pair",price,"symbol",volume,"verified","basetokenblockchain","basetokenaddress" 
		FROM %s 
		WHERE time<%d
		AND time>%d-10d 
		AND exchange='%s' 
		AND quotetokenaddress='%s' 
		AND quotetokenblockchain='%s' 
		AND estimatedUSDPrice>0 
		ORDER BY DESC LIMIT %d
		`
		q = fmt.Sprintf(queryString, influxDbTradesTable, timestamp.UnixNano(), timestamp.UnixNano(), exchange, asset.Address, asset.Blockchain, maxTrades)
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

// GetNumTradesExchange24H returns the number of trades on @exchange in the last 24 hours.
func (datastore *DB) GetNumTradesExchange24H(exchange string) (numTrades int64, err error) {
	endtime := time.Now()
	return datastore.GetNumTrades(exchange, "", "", endtime.AddDate(0, 0, -1), endtime)
}

// GetNumTrades returns the number of trades on @exchange for asset with @address and @blockchain in the given time-range.
// If @address and @blockchain are empty, it returns all trades on @exchange in the given-time range.
func (datastore *DB) GetNumTrades(exchange string, address string, blockchain string, starttime time.Time, endtime time.Time) (numTrades int64, err error) {
	var q string

	if address != "" && blockchain != "" {
		queryString := `
	SELECT COUNT(*) 
	FROM %s 
	WHERE exchange='%s' 
	AND quotetokenaddress='%s' AND quotetokenblockchain='%s' 
	AND time > %d AND time<= %d
	`
		q = fmt.Sprintf(queryString, influxDbTradesTable, exchange, address, blockchain, starttime.UnixNano(), endtime.UnixNano())
	} else {
		queryString := `
	SELECT COUNT(*) 
	FROM %s 
	WHERE exchange='%s' 
	AND time > %d AND time<= %d
	`
		q = fmt.Sprintf(queryString, influxDbTradesTable, exchange, starttime.UnixNano(), endtime.UnixNano())
	}

	res, err := queryInfluxDB(datastore.influxClient, q)
	if err != nil {
		log.Errorln("GetNumTrades ", err)
		return
	}

	if len(res) > 0 && len(res[0].Series) > 0 {
		num, ok := res[0].Series[0].Values[0][1].(json.Number)
		if ok {
			numTrades, err = num.Int64()
			if err != nil {
				return numTrades, err
			}
		}
	}

	return
}

// GetNumTradesSeries returns a time-series of number of trades in the respective time-ranges.
// If pair is the empty string, trades are identified by address/blockchain.
// @grouping defines the time-ranges in the notation of influx such as 30s, 40m, 2h,...
func (datastore *DB) GetNumTradesSeries(
	asset dia.Asset,
	exchange string,
	starttime time.Time,
	endtime time.Time,
	grouping string,
) (numTrades []int64, err error) {
	var query string
	selectQuery := "SELECT COUNT(price) FROM %s "
	midQuery := "WHERE quotetokenaddress='%s' AND quotetokenblockchain='%s' AND time<=%d AND time>%d "
	endQuery := "GROUP BY time(%s) ORDER BY ASC"
	exchangeQuery := "AND exchange='%s' "
	if exchange != "" {
		query = fmt.Sprintf(selectQuery+midQuery+exchangeQuery+endQuery,
			influxDbTradesTable,
			exchange,
			asset.Address,
			asset.Blockchain,
			endtime.UnixNano(),
			starttime.UnixNano(),
			grouping,
		)
	} else {
		query = fmt.Sprintf(selectQuery+midQuery+endQuery,
			influxDbTradesTable,
			asset.Address,
			asset.Blockchain,
			endtime.UnixNano(),
			starttime.UnixNano(),
			grouping,
		)
	}

	res, err := queryInfluxDB(datastore.influxClient, query)
	if err != nil {
		return
	}
	if len(res) > 0 && len(res[0].Series) > 0 {
		for _, val := range res[0].Series[0].Values {
			num, ok := val[1].(json.Number)
			if ok {
				aux, err := num.Int64()
				if err != nil {
					return numTrades, err
				}
				numTrades = append(numTrades, aux)
			}
		}
	}
	return
}

func (datastore *DB) GetFirstTradeDate(table string) (time.Time, error) {
	var query string
	queryString := "SELECT \"exchange\",price FROM %s  where time<now() order by asc limit 1"
	query = fmt.Sprintf(queryString, table)

	res, err := queryInfluxDB(datastore.influxClient, query)
	if err != nil {
		return time.Time{}, err
	}
	if len(res) > 0 && len(res[0].Series) > 0 {
		return time.Parse(time.RFC3339, res[0].Series[0].Values[0][0].(string))
	}
	return time.Time{}, errors.New("no trade found")

}

func getKeyLastTradeTimeForExchange(asset dia.Asset, exchange string) string {
	if exchange == "" {
		return "dia_TLT_" + asset.Blockchain + "_" + asset.Address

	} else {
		return "dia_TLT_" + asset.Blockchain + "_" + asset.Address + "_" + exchange
	}
}

func (datastore *DB) GetLastTradeTimeForExchange(asset dia.Asset, exchange string) (*time.Time, error) {
	key := getKeyLastTradeTimeForExchange(asset, exchange)
	t, err := datastore.redisClient.Get(key).Result()
	if err != nil {
		log.Errorln("Error: on GetLastTradeTimeForExchange", err, key)
		return nil, err
	}
	i64, err := strconv.ParseInt(t, 10, 64)
	if err == nil {
		t2 := time.Unix(i64, 0)
		return &t2, nil
	} else {
		return nil, err
	}
}

func (datastore *DB) SetLastTradeTimeForExchange(asset dia.Asset, exchange string, t time.Time) error {
	if datastore.redisClient == nil {
		return nil
	}
	key := getKeyLastTradeTimeForExchange(asset, exchange)
	log.Debug("setting ", key, t)
	err := datastore.redisPipe.Set(key, t.Unix(), TimeOutRedis).Err()
	if err != nil {
		log.Printf("Error: %v on SetLastTradeTimeForExchange %v\n", err, asset.Symbol)
	}
	return err
}

// -------------------------------------------------------------------------------------------------------
// REDIS TRADES CACHE
// -------------------------------------------------------------------------------------------------------

func (datastore *DB) SaveTradeRedis(t dia.Trade) error {

	tradeMarshaled, err := t.MarshalBinary()
	if err != nil {
		return err
	}

	member := redis.Z{
		Score:  float64(t.Time.UnixMicro()),
		Member: tradeMarshaled,
	}
	err = datastore.redisPipe.ZAdd(GetTradesKeyRedis(t), member).Err()
	if err != nil {
		return err
	}

	//Set key as value in redis keystore with key Blockchain_Address.
	return datastore.SaveTradeSetElementRedis(t)
}

// SaveTradeSetElementRedis has blockchain-address strings as keys. The elements of the
// values set are all keys from trade_... zsets which store the corresponding trades as members.
func (datastore *DB) SaveTradeSetElementRedis(t dia.Trade) error {
	// Store trades keys (i.e. exchangepairs) as values with asset identifier as key.
	err := datastore.redisPipe.SAdd(getAssetKeyRedis(t.QuoteToken), GetTradesKeyRedis(t)).Err()
	if err != nil {
		return err
	}
	if err = datastore.redisPipe.Expire(getAssetKeyRedis(t.QuoteToken), time.Duration((TRADES_TIMEOUT_REDIS_24H+TRADES_TIMEOUT_BUFFER)*time.Second)).Err(); err != nil {
		log.Errorf("expire %s: %v", getAssetKeyRedis(t.QuoteToken), err)
	}

	// For easier retrieval, save all assets that were traded at least once in a set.
	return datastore.redisPipe.SAdd(TRADED_ASSETS_KEY, getAssetKeyRedis(t.QuoteToken)).Err()
}

func (datastore *DB) GetExchangepairKeysRedis(asset dia.Asset) (exchangepairKeys []string, err error) {
	err = datastore.redisClient.SMembers(getAssetKeyRedis(asset)).ScanSlice(&exchangepairKeys)
	return
}

// GetTradesRedis returns all trades from redis cache given by @key in the specified time range.
func (datastore *DB) GetTradesRedis(key string, startTime time.Time, endTime time.Time, offset int64, limit int64) (trades []dia.Trade, err error) {
	vals, err := datastore.redisClient.ZRangeByScoreWithScores(key, redis.ZRangeBy{
		Min:    strconv.FormatInt(startTime.UnixMicro(), 10),
		Max:    strconv.FormatInt(endTime.UnixMicro(), 10),
		Offset: offset,
		Count:  limit,
	}).Result()
	if err != nil {
		return
	}

	for _, v := range vals {
		var t dia.Trade
		err = t.UnmarshalBinary([]byte(v.Member.(string)))
		if err != nil {
			return
		}
		trades = append(trades, t)
	}
	// TO DO: Add possibility of sorting as var?
	// // Sort by time in descending order.
	// sort.Slice(trades, func(i, j int) bool {
	// 	return trades[i].Time.UnixMicro() > trades[j].Time.UnixMicro()
	// })
	return
}

// GetTradesByFeedselectionRedis returns all trades fulfilling conditions in @feedselection.
// Trades are returned sorted by time in ascending order.
// If @limit > 0, only @limit number of trades are returned.
func (datastore *DB) GetTradesByFeedselectionRedis(feedselection []dia.FeedSelection, startTimes []time.Time, endTimes []time.Time, limit int64, desc bool) ([]dia.Trade, error) {
	var trades []dia.Trade
	if len(startTimes) != len(endTimes) {
		return []dia.Trade{}, errors.New("number of start times must equal number of end times")
	}

	for i := range startTimes {
		for _, fs := range feedselection {
			keys, err := datastore.GetRedisKeysByFeedselection(fs)
			if err != nil {
				return trades, err
			}

			for _, key := range keys {
				ts, err := datastore.GetTradesRedis(key, startTimes[i], endTimes[i], 0, limit)
				if err != nil {
					return []dia.Trade{}, err
				}
				trades = append(trades, ts...)

				if limit > 0 && int64(len(trades)) >= limit {
					return trades, err
				}
			}
		}
	}

	// Sort by time.
	if desc {
		// Sort in descending order.
		sort.Slice(trades, func(i, j int) bool {
			return trades[i].Time.UnixMicro() > trades[j].Time.UnixMicro()
		})
	} else {
		sort.Slice(trades, func(i, j int) bool {
			return trades[i].Time.UnixMicro() < trades[j].Time.UnixMicro()
		})
	}
	if limit > 0 {
		indMax, err := utils.Min([]int64{limit, int64(len(trades))})
		if err != nil {
			return trades, err
		}
		return trades[:indMax], nil
	}
	return trades, nil
}

// GetRedisKeysByFeedselection returns a list of keys which are needed in order to retrieve
// all trades from the redis cache which fulfill the conditions given by @feedselection.
func (datastore *DB) GetRedisKeysByFeedselection(fs dia.FeedSelection) ([]string, error) {
	if len(fs.Exchangepairs) == 0 {
		// Take into account all available exchangepairs for the asset given by feedselection entry.
		return datastore.GetExchangepairKeysRedis(fs.Asset)
	}

	var keys []string
	exchangepairKeys, err := datastore.GetExchangepairKeysRedis(fs.Asset)
	if err != nil {
		return []string{}, err
	}

	for _, ep := range fs.Exchangepairs {
		exchange := ep.Exchange
		if len(ep.Pairs)+len(ep.Pools) == 0 {
			// Take into account all pairs/pools on the given exchange.
			for _, ep := range exchangepairKeys {
				if strings.Contains(ep, exchange.Name) {
					keys = append(keys, ep)
				}
			}
			continue
		}

		if exchange.Centralized {
			for _, pair := range ep.Pairs {
				key := GetTradesKeyRedis(dia.Trade{QuoteToken: pair.QuoteToken, BaseToken: pair.BaseToken, Source: exchange.Name})
				keys = append(keys, key)
			}
		} else {
			for _, pool := range ep.Pools {
				// Get exchangepair key in redis which contains the pool address (should be at most one).
				for _, ep := range exchangepairKeys {
					if strings.Contains(ep, pool.Address) {
						keys = append(keys, ep)
						break
					}
				}
			}
		}
	}
	return keys, nil
}

// PurgeTradesAndKeysRedis deletes all trades older than @timestampLatest.
func (datastore *DB) PurgeTradesAndKeysRedis(timestampLatest time.Time, expire time.Duration) {

	// Get all (quote) assets traded.
	allTradedAssetKeys := datastore.redisClient.SMembers(TRADED_ASSETS_KEY).Val()
	for _, assetKey := range allTradedAssetKeys {
		// Get all exchangepair keys of a given quote asset.
		exchangepairKeys := datastore.redisClient.SMembers(assetKey).Val()
		for _, exchangepairKey := range exchangepairKeys {
			// Purging old values.
			err := datastore.redisClient.ZRemRangeByScore(exchangepairKey, "-inf", strconv.FormatInt(timestampLatest.UnixMicro(), 10)).Err()
			if err != nil {
				log.Errorf("ZRemRangeByScore on %s: %v", err, exchangepairKey)
			}
			// Remove set member if no trades are left.
			err = datastore.redisClient.SRem(assetKey, exchangepairKey).Err()
			if err != nil {
				log.Errorf("SRem assetKey -- exchangepairKey: %s -- %s", assetKey, exchangepairKey)
			}
		}
	}

}

func getAssetKeyRedis(asset dia.Asset) string {
	return TRADE_EXCHANGEPAIR_PREFIX_REDIS + asset.Identifier()
}

func GetTradesKeyRedis(t dia.Trade) string {
	return TRADE_PREFIX_REDIS + t.TradeIdentifierShort()
}
