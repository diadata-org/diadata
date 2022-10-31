package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/diadata-org/diadata/pkg/dia/helpers/db"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/go-redis/redis"
	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
)

type Datastore interface {
	SetInfluxClient(url string)
	SetBatchFiatPriceInflux(fqs []*FiatQuotation) error
	SetSingleFiatPriceRedis(fiatQuotation *FiatQuotation) error

	GetLatestSupply(string, *RelDB) (*dia.Supply, error)
	GetSupplyCache(asset dia.Asset) (dia.Supply, error)
	GetSupply(string, time.Time, time.Time, *RelDB) ([]dia.Supply, error)
	SetSupply(supply *dia.Supply) error
	GetSupplyInflux(dia.Asset, time.Time, time.Time) ([]dia.Supply, error)
	SaveSynthSupplyInfluxToTable(*dia.SynthAssetSupply, string) error
	SaveSynthSupplyInflux(*dia.SynthAssetSupply) error
	GetSynthSupplyInflux(string, string, string, int, time.Time, time.Time) ([]dia.SynthAssetSupply, error)
	GetSynthAssets(string, string) ([]string, error)

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
	GetActiveExchangesAndPairs(address string, blockchain string, starttime time.Time, endtime time.Time) (map[string][]dia.Pair, error)
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
	GetVolumesAllExchanges(asset dia.Asset, starttime time.Time, endtime time.Time) (exchVolumes dia.ExchangeVolumesList, err error)

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
	influxDbSynthSupplyTable          = "synthsupply"

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

func (datastore *DB) ExecuteRedisPipe() (err error) {
	// TO DO: Handle first return value for read requests.
	_, err = datastore.redisPipe.Exec()
	return
}

func (datastore *DB) FlushRedisPipe() error {
	return datastore.redisPipe.Discard()
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
