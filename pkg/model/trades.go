package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
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
		"symbol":               t.Symbol,
		"pair":                 t.Pair,
		"exchange":             t.Source,
		"verified":             strconv.FormatBool(t.VerifiedPair),
		"quotetokenaddress":    t.QuoteToken.Address,
		"basetokenaddress":     t.BaseToken.Address,
		"quotetokenblockchain": t.QuoteToken.Blockchain,
		"basetokenblockchain":  t.BaseToken.Blockchain,
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

func (datastore *DB) GetTradesByExchanges(asset dia.Asset, baseassets []dia.Asset, exchanges []string, startTime, endTime time.Time) ([]dia.Trade, error) {
	return datastore.GetTradesByExchangesFull(asset, baseassets, exchanges, false, startTime, endTime)
}
func (datastore *DB) GetTradesByExchangesAndBaseAssets(asset dia.Asset, baseassets []dia.Asset, exchanges []string, startTime, endTime time.Time) ([]dia.Trade, error) {
	return datastore.GetTradesByExchangesFull(asset, baseassets, exchanges, false, startTime, endTime)
}

func (datastore *DB) GetTradesByExchangesFull(asset dia.Asset, baseassets []dia.Asset, exchanges []string, returnBasetoken bool, startTime, endTime time.Time) ([]dia.Trade, error) {
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
	query := fmt.Sprintf("SELECT time,estimatedUSDPrice,exchange,foreignTradeID,pair,price,symbol,volume,verified,basetokenblockchain,basetokenaddress FROM %s WHERE (quotetokenaddress='%s' and quotetokenblockchain='%s') %s %s AND estimatedUSDPrice > 0 AND time >= %d AND time <= %d ", influxDbTradesTable, asset.Address, asset.Blockchain, subQuery, subQueryBase, startTime.UnixNano(), endTime.UnixNano())
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
					subQueryBase = subQueryBase + fmt.Sprintf(` and ((basetokenaddress='%s' and basetokenblockchain='%s')`, baseasset.Address, baseasset.Blockchain)

				} else {
					subQueryBase = subQueryBase + fmt.Sprintf(` or (basetokenaddress='%s' and basetokenblockchain='%s')`, baseasset.Address, baseasset.Blockchain)
				}

			}
			subQueryBase = subQueryBase + ") "

		}
		log.Errorln("subQueryBase", subQueryBase)
		query = query + fmt.Sprintf("SELECT time,estimatedUSDPrice,exchange,foreignTradeID,pair,price,symbol,volume,verified,basetokenblockchain,basetokenaddress FROM %s WHERE (quotetokenaddress='%s' AND quotetokenblockchain='%s') %s %s AND estimatedUSDPrice > 0 AND time > %d AND time <= %d ;", influxDbTradesTable, quoteasset.Address, quoteasset.Blockchain, subQuery, subQueryBase, startTimes[i].UnixNano(), endTimes[i].UnixNano())
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

// GetLastTrades returns the last @maxTrades of @asset on @exchange.
// If exchange is empty string it returns trades from all exchanges.
// If fullAsset=true, blockchain and address of both involved assets is returned as well
func (datastore *DB) GetLastTrades(asset dia.Asset, exchange string, maxTrades int, fullAsset bool) ([]dia.Trade, error) {
	var (
		r           []dia.Trade
		queryString string
		q           string
	)

	if exchange == "" {
		queryString = `
		SELECT estimatedUSDPrice,"exchange",foreignTradeID,"pair",price,"symbol",volume,"verified","basetokenblockchain","basetokenaddress" 
		FROM %s 
		WHERE time<now() 
		AND time>now()-10d 
		AND quotetokenaddress='%s' 
		AND quotetokenblockchain='%s' 
		AND estimatedUSDPrice>0 
		ORDER BY DESC LIMIT %d
		`
		q = fmt.Sprintf(queryString, influxDbTradesTable, asset.Address, asset.Blockchain, maxTrades)
	} else if (dia.Asset{}) == asset {
		queryString = `
		SELECT estimatedUSDPrice,"exchange",foreignTradeID,"pair",price,"symbol",volume,"verified","basetokenblockchain","basetokenaddress" 
		FROM %s 
		WHERE time<now() 
		AND time>now()-10d 
		AND exchange='%s' 
		AND estimatedUSDPrice>0 
		ORDER BY DESC LIMIT %d
		`
		q = fmt.Sprintf(queryString, influxDbTradesTable, exchange, maxTrades)
	} else {
		queryString = `
		SELECT estimatedUSDPrice,"exchange",foreignTradeID,"pair",price,"symbol",volume,"verified","basetokenblockchain","basetokenaddress" 
		FROM %s 
		WHERE time<now() 
		AND time>now()-10d 
		AND exchange='%s' 
		AND quotetokenaddress='%s' 
		AND quotetokenblockchain='%s' 
		AND estimatedUSDPrice>0 
		ORDER BY DESC LIMIT %d
		`
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
	exchange string,
	pair string,
	starttime time.Time,
	endtime time.Time,
	grouping string,
	quotetoken dia.Asset,
	basetoken dia.Asset,
) (numTrades []int, err error) {
	var query string
	if pair != "" {
		queryString := `SELECT COUNT(price) 
		FROM %s WHERE exchange='%s' 
		AND pair='%s' 
		AND time<=%d 
		AND time>%d 
		GROUP BY time('%s') 
		ORDER BY ASC`
		query = fmt.Sprintf(
			queryString,
			influxDbTradesTable,
			exchange,
			pair,
			endtime.UnixNano(),
			starttime.UnixNano(),
			grouping,
		)
	} else {
		queryString := `SELECT COUNT(price) FROM %s 
		WHERE exchange='%s' 
		AND quotetokenaddress='%s' AND quotetokenblockchain='%s' 
		AND basetokenaddress='%s' AND basetokenblockchain='%s' 
		AND time<=%d 
		AND time>%d 
		GROUP BY time('%s') 
		ORDER BY ASC`
		query = fmt.Sprintf(
			queryString,
			influxDbTradesTable,
			exchange,
			quotetoken.Address,
			quotetoken.Blockchain,
			basetoken.Address,
			basetoken.Blockchain,
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
				numTrades = append(numTrades, int(aux))
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
