package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia/helpers/db"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/go-redis/redis"
	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
)

type Datastore interface {
	SetInfluxClient(url string)

	// Deprecating
	SetPriceUSD(symbol string, price float64) error
	SetPriceEUR(symbol string, price float64) error
	GetPriceUSD(symbol string) (float64, error)
	GetQuotation(symbol string) (*Quotation, error)
	SetQuotation(quotation *Quotation) error
	SetQuotationEUR(quotation *Quotation) error
	SetBatchFiatPriceInflux(fqs []*FiatQuotation) error
	SetSingleFiatPriceRedis(fiatQuotation *FiatQuotation) error

	GetLatestSupply(string, *RelDB) (*dia.Supply, error)
	GetSupplyCache(asset dia.Asset) (dia.Supply, error)
	GetSupply(string, time.Time, time.Time, *RelDB) ([]dia.Supply, error)
	SetSupply(supply *dia.Supply) error
	GetSupplyInflux(dia.Asset, time.Time, time.Time) ([]dia.Supply, error)

	SetDiaTotalSupply(totalSupply float64) error
	GetDiaTotalSupply() (float64, error)
	SetDiaCirculatingSupply(circulatingSupply float64) error
	GetDiaCirculatingSupply() (float64, error)

	GetSymbols(exchange string) ([]string, error)
	GetLastTradeTimeForExchange(asset dia.Asset, exchange string) (*time.Time, error)
	SetLastTradeTimeForExchange(asset dia.Asset, exchange string, t time.Time) error
	GetFirstTradeDate(table string) (time.Time, error)
	SaveTradeInflux(t *dia.Trade) error
	SaveTradeInfluxToTable(t *dia.Trade, table string) error
	GetTradeInflux(dia.Asset, string, time.Time, time.Duration) (*dia.Trade, error)
	SaveFilterInflux(filter string, asset dia.Asset, exchange string, value float64, t time.Time) error
	GetLastTrades(asset dia.Asset, exchange string, maxTrades int, fullAsset bool) ([]dia.Trade, error)
	GetAllTrades(t time.Time, maxTrades int) ([]dia.Trade, error)
	GetTradesByExchanges(asset dia.Asset, baseAssets []dia.Asset, exchange []string, startTime, endTime time.Time) ([]dia.Trade, error)
	GetTradesByExchangesFull(asset dia.Asset, baseAssets []dia.Asset, exchanges []string, returnBasetoken bool, startTime, endTime time.Time) ([]dia.Trade, error)
	GetTradesByExchangesBatched(asset dia.Asset, baseAssets []dia.Asset, exchanges []string, startTimes, endTimes []time.Time) ([]dia.Trade, error)
	GetTradesByExchangesBatchedFull(asset dia.Asset, baseAssets []dia.Asset, exchanges []string, returnBasetoken bool, startTimes, endTimes []time.Time) ([]dia.Trade, error)
	GetActiveExchangesAndPairs(address string, blockchain string, starttime time.Time, endtime time.Time) (map[string][]string, error)
	GetOldTradesFromInflux(table string, exchange string, verified bool, timeInit, timeFinal time.Time) ([]dia.Trade, error)
	CopyInfluxMeasurements(dbOrigin string, dbDestination string, tableOrigin string, tableDestination string, timeInit time.Time, timeFinal time.Time) (int64, error)

	Flush() error
	ExecuteRedisPipe() error
	FlushRedisPipe() error
	GetFilterPoints(filter string, exchange string, symbol string, scale string, starttime time.Time, endtime time.Time) (*Points, error)
	GetFilterPointsAsset(filter string, exchange string, address string, blockchain string, starttime time.Time, endtime time.Time) (*Points, error)
	SetFilter(filterName string, asset dia.Asset, exchange string, value float64, t time.Time) error
	GetLastPriceBefore(asset dia.Asset, filter string, exchange string, timestamp time.Time) (Price, error)
	SetAvailablePairs(exchange string, pairs []dia.ExchangePair) error
	GetAvailablePairs(exchange string) ([]dia.ExchangePair, error)
	SetCurrencyChange(cc *Change) error
	GetCurrencyChange() (*Change, error)

	// Volume methods
	GetVolumeInflux(asset dia.Asset, exchange string, starttime time.Time, endtime time.Time) (*float64, error)
	Get24HoursAssetVolume(asset dia.Asset) (*float64, error)
	Get24HoursExchangeVolume(exchange string) (*float64, error)
	GetNumTradesExchange24H(exchange string) (int64, error)
	GetNumTrades(exchange string, address string, blockchain string, starttime time.Time, endtime time.Time) (int64, error)
	GetNumTradesSeries(
		exchange string,
		pair string,
		starttime time.Time,
		endtime time.Time,
		grouping string,
		quotetoken dia.Asset,
		basetoken dia.Asset,
	) ([]int, error)

	// New Asset pricing methods: 23/02/2021
	SetAssetPriceUSD(asset dia.Asset, price float64, timestamp time.Time) error
	GetAssetPriceUSD(asset dia.Asset, timestamp time.Time) (float64, error)
	GetAssetPriceUSDLatest(asset dia.Asset) (price float64, err error)
	SetAssetQuotation(quotation *AssetQuotation) error
	GetAssetQuotation(asset dia.Asset, timestamp time.Time) (*AssetQuotation, error)
	GetAssetQuotations(asset dia.Asset, starttime time.Time, endtime time.Time) ([]AssetQuotation, error)
	GetAssetQuotationLatest(asset dia.Asset) (*AssetQuotation, error)
	GetSortedAssetQuotations(assets []dia.Asset) ([]AssetQuotation, error)
	AddAssetQuotationsToBatch(quotations []*AssetQuotation) error
	SetAssetQuotationCache(quotation *AssetQuotation, check bool) (bool, error)
	GetAssetQuotationCache(asset dia.Asset) (*AssetQuotation, error)
	GetAssetPriceUSDCache(asset dia.Asset) (price float64, err error)
	GetTopAssetByMcap(symbol string, relDB *RelDB) (dia.Asset, error)
	GetTopAssetByVolume(symbol string, relDB *RelDB) (topAsset dia.Asset, err error)
	GetAssetsWithVOLInflux(timeInit time.Time) ([]dia.Asset, error)

	// DEX Pool  methods
	SavePoolInflux(p dia.Pool) error
	GetPoolInflux(poolAddress string, starttime time.Time, endtime time.Time) ([]dia.Pool, error)

	// Market Measures
	GetAssetsMarketCap(asset dia.Asset) (float64, error)

	// Interest rates' methods
	SetInterestRate(ir *InterestRate) error
	GetInterestRate(symbol, date string) (*InterestRate, error)
	GetInterestRateRange(symbol, dateInit, dateFinal string) ([]*InterestRate, error)
	GetRatesMeta() (RatesMeta []InterestRateMeta, err error)
	GetCompoundedIndex(symbol string, date time.Time, daysPerYear int, rounding int) (*InterestRate, error)
	GetCompoundedIndexRange(symbol string, dateInit, dateFinal time.Time, daysPerYear int, rounding int) ([]*InterestRate, error)
	GetCompoundedAvg(symbol string, date time.Time, calDays, daysPerYear int, rounding int) (*InterestRate, error)
	GetCompoundedAvgRange(symbol string, dateInit, dateFinal time.Time, calDays, daysPerYear int, rounding int) ([]*InterestRate, error)
	GetCompoundedAvgDIARange(symbol string, dateInit, dateFinal time.Time, calDays, daysPerYear int, rounding int) ([]*InterestRate, error)

	// Foreign quotation methods
	SaveForeignQuotationInflux(fq ForeignQuotation) error
	GetForeignQuotationInflux(symbol, source string, timestamp time.Time) (ForeignQuotation, error)
	GetForeignPriceYesterday(symbol, source string) (float64, error)
	GetForeignSymbolsInflux(source string) ([]string, error)

	SetVWAPFirefly(foreignName string, value float64, timestamp time.Time) error
	GetVWAPFirefly(foreignName string, starttime time.Time, endtime time.Time) ([]float64, []time.Time, error)

	// Gold token methods
	GetPaxgQuotationOunces() (*Quotation, error)
	GetPaxgQuotationGrams() (*Quotation, error)

	SaveIndexEngineTimeInflux(map[string]string, map[string]interface{}, time.Time) error
	GetBenchmarkedIndexValuesInflux(string, time.Time, time.Time) (BenchmarkedIndex, error)
	// Token methods
	// SaveTokenDetailInflux(tk Token) error
	// GetTokenDetailInflux(symbol, source string, timestamp time.Time) (Token, error)
	// GetCurentTotalSupply(symbol, source string) (float64, error)

	// Stock methods
	SetStockQuotation(sq StockQuotation) error
	GetStockQuotation(source string, symbol string, timeInit time.Time, timeFinal time.Time) ([]StockQuotation, error)
	GetStockSymbols() (map[Stock]string, error)
}

const (
	influxMaxPointsInBatch = 5000
	// timeOutRedisOneBlock   = 60 * 3 * time.Second
)

type DB struct {
	redisClient         *redis.Client
	redisPipe           redis.Pipeliner
	influxClient        clientInfluxdb.Client
	influxBatchPoints   clientInfluxdb.BatchPoints
	influxPointsInBatch int
}

const (
	influxDbName                      = "dia"
	influxDbTradesTable               = "trades"
	influxDbFiltersTable              = "filters"
	influxDbFiatQuotationsTable       = "fiat"
	influxDbSupplyTable               = "supplies"
	influxDbDEXPoolTable              = "DEXPools"
	influxDbStockQuotationsTable      = "stockquotations"
	influxDBAssetQuotationsTable      = "assetQuotations"
	influxDbBenchmarkedIndexTableName = "benchmarkedIndexValues"
	influxDbVwapFireflyTable          = "vwapFirefly"

	influxDBDefaultURL = "http://influxdb:8086"
)

// queryInfluxDB convenience function to query the database.
func queryInfluxDB(clnt clientInfluxdb.Client, cmd string) (res []clientInfluxdb.Result, err error) {
	res, err = queryInfluxDBName(clnt, influxDbName, cmd)
	return
}

// queryInfluxDBName is a wrapper for queryInfluxDB that allows for queries on the database with name @dbName.
func queryInfluxDBName(clnt clientInfluxdb.Client, dbName string, cmd string) (res []clientInfluxdb.Result, err error) {
	q := clientInfluxdb.Query{
		Command:  cmd,
		Database: dbName,
	}
	if response, err := clnt.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

func NewDataStore() (*DB, error) {
	return NewDataStoreWithOptions(true, true)
}
func NewInfluxDataStore() (*DB, error) {
	return NewDataStoreWithOptions(false, true)
}

func NewRedisDataStore() (*DB, error) {
	return NewDataStoreWithOptions(true, false)
}

func NewDataStoreWithoutInflux() (*DB, error) {
	return NewDataStoreWithOptions(true, false)
}

func NewDataStoreWithoutRedis() (*DB, error) {
	return NewDataStoreWithOptions(false, true)
}

func NewDataStoreWithOptions(withRedis bool, withInflux bool) (*DB, error) {
	var influxClient clientInfluxdb.Client
	var influxBatchPoints clientInfluxdb.BatchPoints
	var redisClient *redis.Client
	var redisPipe redis.Pipeliner

	if withRedis {
		redisClient = db.GetRedisClient()
		redisPipe = redisClient.TxPipeline()
	}
	if withInflux {
		var err error
		influxClient = db.GetInfluxClient(influxDBDefaultURL)
		influxBatchPoints = createBatchInflux()
		_, err = queryInfluxDB(influxClient, fmt.Sprintf("CREATE DATABASE %s", influxDbName))
		if err != nil {
			log.Errorln("queryInfluxDB CREATE DATABASE", err)
		}
	}
	return &DB{redisClient, redisPipe, influxClient, influxBatchPoints, 0}, nil
}

// SetInfluxClient resets influx's client url to @url.
func (datastore *DB) SetInfluxClient(url string) {
	datastore.influxClient = db.GetInfluxClient(url)
}

func createBatchInflux() clientInfluxdb.BatchPoints {
	bp, err := clientInfluxdb.NewBatchPoints(clientInfluxdb.BatchPointsConfig{
		Database:  influxDbName,
		Precision: "ns",
	})
	if err != nil {
		log.Errorln("NewBatchPoints", err)
	}
	return bp
}

func (datastore *DB) Flush() error {
	var err error
	if datastore.influxBatchPoints != nil {
		err = datastore.WriteBatchInflux()
	}
	return err
}

func getKey(filter string, asset dia.Asset, exchange string) string {
	key := filter + "_" + asset.Blockchain + "_" + asset.Address
	if exchange != "" {
		key = key + "_" + exchange
	}
	return key
}

func getKeyFilterZSET(key string) string {
	return "dia_" + key + "_ZSET"
}

func getKeyFilterSymbolAndExchangeZSET(filter string, asset dia.Asset, exchange string) string {
	if exchange == "" {
		return "dia_" + filter + "_" + asset.Blockchain + "_" + asset.Address + "_ZSET"
	} else {
		return "dia_" + filter + "_" + asset.Blockchain + "_" + asset.Address + "_ZSET"
	}
}

func (datastore *DB) WriteBatchInflux() (err error) {
	err = datastore.influxClient.Write(datastore.influxBatchPoints)
	if err != nil {
		log.Errorln("WriteBatchInflux", err)
		return
	}
	datastore.influxPointsInBatch = 0
	datastore.influxBatchPoints = createBatchInflux()
	return
}

func (datastore *DB) addPoint(pt *clientInfluxdb.Point) {
	datastore.influxBatchPoints.AddPoint(pt)
	datastore.influxPointsInBatch++

	if datastore.influxPointsInBatch >= influxMaxPointsInBatch {
		err := datastore.WriteBatchInflux()
		if err != nil {
			log.Error("write influx batch: ", err)
		}
	}
}

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

// CopyInfluxMeasurements copies entries from measurement @tableOrigin in database @dbOrigin into @tableDestination in database @dbDestination.
// It takes into account all data ranging from @timeInit until @timeFinal.
func (datastore *DB) CopyInfluxMeasurements(dbOrigin string, dbDestination string, tableOrigin string, tableDestination string, timeInit time.Time, timeFinal time.Time) (numCopiedRows int64, err error) {
	queryString := "select * into %s..%s from %s..%s where time>%d and time<=%d group by *"
	query := fmt.Sprintf(queryString, dbDestination, tableDestination, dbOrigin, tableOrigin, timeInit.UnixNano(), timeFinal.UnixNano())
	res, err := queryInfluxDB(datastore.influxClient, query)
	if err != nil {
		return
	}
	if len(res) > 0 && len(res[0].Series) > 0 {
		numCopiedRows, err = res[0].Series[0].Values[0][1].(json.Number).Int64()
		if err != nil {
			return
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

func (datastore *DB) SaveSupplyInflux(supply *dia.Supply) error {
	fields := map[string]interface{}{
		"supply":            supply.Supply,
		"circulatingsupply": supply.CirculatingSupply,
		"source":            supply.Source,
	}
	tags := map[string]string{
		"symbol":     supply.Asset.Symbol,
		"name":       supply.Asset.Name,
		"address":    supply.Asset.Address,
		"blockchain": supply.Asset.Blockchain,
	}
	pt, err := clientInfluxdb.NewPoint(influxDbSupplyTable, tags, fields, supply.Time)
	if err != nil {
		log.Errorln("NewSupplyInflux:", err)
	} else {
		datastore.addPoint(pt)
	}

	err = datastore.WriteBatchInflux()
	if err != nil {
		log.Errorln("SaveSupplyInflux", err)
	}

	return err
}

// GetSupplyInflux returns supply and circulating supply of @asset. Needs asset.Address and asset.Blockchain.
// If no time range is given it returns the latest supply.
func (datastore *DB) GetSupplyInflux(asset dia.Asset, starttime time.Time, endtime time.Time) ([]dia.Supply, error) {
	retval := []dia.Supply{}
	var q string
	if starttime.IsZero() || endtime.IsZero() {
		queryString := "SELECT supply,circulatingsupply,source,\"name\",\"symbol\" FROM %s WHERE \"address\" = '%s' AND \"blockchain\"='%s' AND time<now() ORDER BY DESC LIMIT 1"
		q = fmt.Sprintf(queryString, influxDbSupplyTable, asset.Address, asset.Blockchain)
	} else {
		queryString := "SELECT supply,circulatingsupply,source,\"name\",\"symbol\" FROM %s WHERE time > %d AND time < %d AND \"address\" = '%s' AND \"blockchain\"='%s' ORDER BY DESC"
		q = fmt.Sprintf(queryString, influxDbSupplyTable, starttime.UnixNano(), endtime.UnixNano(), asset.Address, asset.Blockchain)
	}
	res, err := queryInfluxDB(datastore.influxClient, q)
	if err != nil {
		return retval, err
	}
	if len(res) > 0 && len(res[0].Series) > 0 {
		for i := 0; i < len(res[0].Series[0].Values); i++ {
			currentSupply := dia.Supply{Asset: asset}
			if res[0].Series[0].Values[i][0] != nil {
				currentSupply.Time, err = time.Parse(time.RFC3339, res[0].Series[0].Values[i][0].(string))
				if err != nil {
					return retval, err
				}
			}
			currentSupply.Supply, err = res[0].Series[0].Values[i][1].(json.Number).Float64()
			if err != nil {
				return retval, err
			}
			currentSupply.CirculatingSupply, err = res[0].Series[0].Values[i][2].(json.Number).Float64()
			if err != nil {
				return retval, err
			}
			if res[0].Series[0].Values[i][3] != nil {
				currentSupply.Source = res[0].Series[0].Values[i][3].(string)
				if err != nil {
					return retval, err
				}
			}
			if res[0].Series[0].Values[i][4] != nil {
				currentSupply.Asset.Name = res[0].Series[0].Values[i][4].(string)
				if err != nil {
					log.Error("error getting symbol name from influx: ", err)
				}
			}
			if res[0].Series[0].Values[i][5] != nil {
				currentSupply.Asset.Symbol = res[0].Series[0].Values[i][5].(string)
				if err != nil {
					log.Error("error getting symbol name from influx: ", err)
				}
			}
			retval = append(retval, currentSupply)
		}
	} else {
		return retval, errors.New("parsing supply value from database")
	}
	return retval, nil
}

// SaveFilterInflux stores a filter point in influx.
func (datastore *DB) SaveFilterInflux(filter string, asset dia.Asset, exchange string, value float64, t time.Time) error {
	// Create a point and add to batch
	tags := map[string]string{
		"filter":     filter,
		"symbol":     asset.Symbol,
		"address":    asset.Address,
		"blockchain": asset.Blockchain,
		"exchange":   exchange,
	}
	fields := map[string]interface{}{
		"value":        value,
		"allExchanges": exchange == "",
	}
	pt, err := clientInfluxdb.NewPoint(influxDbFiltersTable, tags, fields, t)
	if err != nil {
		log.Errorln("new filter influx:", err)
	} else {
		datastore.addPoint(pt)
	}
	return err
}

func (datastore *DB) setZSETValue(key string, value float64, unixTime int64, maxWindow int64) error {
	if datastore.redisClient == nil {
		return nil
	}
	member := strconv.FormatFloat(value, 'f', -1, 64) + " " + strconv.FormatInt(unixTime, 10)

	err := datastore.redisPipe.ZAdd(key, redis.Z{
		Score:  float64(unixTime),
		Member: member,
	}).Err()
	log.Debug("SetZSETValue ", key, member, unixTime)
	if err != nil {
		log.Errorf("Error: %v on SetZSETValue %v\n", err, key)
	}
	// purging old values
	err = datastore.redisPipe.ZRemRangeByScore(key, "-inf", "("+strconv.FormatInt(unixTime-maxWindow, 10)).Err()
	if err != nil {
		log.Errorf("Error: %v on SetZSETValue %v\n", err, key)
	}
	if err = datastore.redisPipe.Expire(key, TimeOutRedis).Err(); err != nil {
		log.Error(err)
	} //TODO put two commands together ?
	return err
}

func (datastore *DB) getZSETValue(key string, atUnixTime int64) (float64, error) {

	result := 0.0
	max := strconv.FormatInt(atUnixTime, 10)
	vals, err := datastore.redisClient.ZRangeByScoreWithScores(key, redis.ZRangeBy{
		Min: "-inf",
		Max: max,
	}).Result()
	log.Debug(key, "vals: %v on getZSETValue maxScore: %v", vals, max)
	if err == nil {
		if len(vals) > 0 {
			_, err = fmt.Sscanf(vals[len(vals)-1].Member.(string), "%f", &result)
			if err != nil {
				log.Error(err)
			}
			log.Debugf("returned value: %v", result)
		} else {
			err = errors.New("getZSETValue no value found")
		}
	}
	return result, err
}

func (datastore *DB) ExecuteRedisPipe() (err error) {
	// TO DO: Handle first return value for read requests.
	_, err = datastore.redisPipe.Exec()
	return
}

func (datastore *DB) FlushRedisPipe() error {
	return datastore.redisPipe.Discard()
}

/*
func (db *DB) getZSETSum(key string) (*float64, error) {

	log.Debugf("getZSETSum: %v \n", key)

	vals, err := db.redisClient.ZRange(key, 0, -1).Result()
	if err != nil {
		log.Errorf("Error: %v on getZSETSum %v\n", err, key)
		return nil, err
	} else {
		result := 0.0
		for _, v := range vals {
			f := 0.0
			fmt.Sscanf(v, "%f", &f)
			result += f
		}
		return &result, err
	}
}
*/

func (datastore *DB) getZSETLastValue(key string) (float64, int64, error) {
	value := 0.0
	var unixTime int64
	vals, err := datastore.redisClient.ZRange(key, -1, -1).Result()
	log.Debug(key, "on getZSETLastValue:", vals)
	if err == nil {
		if len(vals) == 1 {
			_, err = fmt.Sscanf(vals[0], "%f %d", &value, &unixTime)
			if err != nil {
				log.Error(err)
			}
			log.Debugf("returned value: %v", value)
		} else {
			err = errors.New("getZSETLastValue no value found")
			log.Errorln("Error: on getZSETLastValue", err, key)
		}
	}
	return value, unixTime, err
}

func (datastore *DB) SetVWAPFirefly(foreignName string, value float64, timestamp time.Time) error {
	tags := map[string]string{
		"foreignName": foreignName,
	}
	fields := map[string]interface{}{
		"value": value,
	}
	pt, err := clientInfluxdb.NewPoint(influxDbVwapFireflyTable, tags, fields, timestamp)
	if err != nil {
		log.Errorln("new filter influx:", err)
	} else {
		datastore.addPoint(pt)
	}

	err = datastore.WriteBatchInflux()
	if err != nil {
		log.Errorln("Write influx batch: ", err)
	}

	return err
}

func (datastore *DB) GetVWAPFirefly(foreignName string, starttime time.Time, endtime time.Time) (values []float64, timestamps []time.Time, err error) {

	influxQuery := "SELECT value FROM %s WHERE time > %d AND time <= %d AND foreignName = '%s' ORDER BY DESC"
	q := fmt.Sprintf(influxQuery, influxDbVwapFireflyTable, starttime.UnixNano(), endtime.UnixNano(), foreignName)
	res, err := queryInfluxDB(datastore.influxClient, q)
	if err != nil {
		return
	}
	if len(res) > 0 && len(res[0].Series) > 0 {
		for i := 0; i < len(res[0].Series[0].Values); i++ {
			var value float64
			var timestamp time.Time
			timestamp, err = time.Parse(time.RFC3339, res[0].Series[0].Values[i][0].(string))
			if err != nil {
				return
			}
			value, err = res[0].Series[0].Values[i][1].(json.Number).Float64()
			if err != nil {
				return
			}

			values = append(values, value)
			timestamps = append(timestamps, timestamp)
		}
	} else {
		err = errors.New("no data available in given time range")
		return
	}
	return
}
