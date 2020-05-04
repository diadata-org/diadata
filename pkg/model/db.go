package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/go-redis/redis"
	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
	log "github.com/sirupsen/logrus"
)

type Datastore interface {
	SetVolume(symbol string, exchange string, volume float64, t time.Time) error
	GetVolume(symbol string) (*float64, error)
	SymbolsWithASupply() ([]string, error)
	SetPriceUSD(symbol string, price float64) error
	SetPriceEUR(symbol string, price float64) error
	GetPriceUSD(symbol string) (float64, error)
	GetQuotation(symbol string) (*Quotation, error)
	SetQuotation(quotation *Quotation) error
	SetQuotationEUR(quotation *Quotation) error
	GetSupply(symbol string) (*dia.Supply, error)
	SetSupply(supply *dia.Supply) error
	SetPriceZSET(symbol string, exchange string, price float64, t time.Time) error
	GetChartPoints7Days(symbol string) ([]Point, error)
	GetPairs(exchange string) ([]dia.Pair, error)
	GetSymbols(exchange string) ([]string, error)
	GetExchangesForSymbol(symbol string) ([]string, error)
	GetSymbolExchangeDetails(symbol string, exchange string) (*SymbolExchangeDetails, error)
	GetLastTradeTimeForExchange(symbol string, exchange string) (*time.Time, error)
	SetLastTradeTimeForExchange(symbol string, exchange string, t time.Time) error
	SaveTradeInflux(t *dia.Trade) error
	SaveFilterInflux(filter string, symbol string, exchange string, value float64, t time.Time) error
	GetLastTrades(symbol string, exchange string, maxTrades int) ([]dia.Trade, error)
	GetAllTrades(t time.Time, maxTrades int) ([]dia.Trade, error)
	Flush() error
	GetFilterPoints(filter string, exchange string, symbol string, scale string) (*Points, error)
	SetFilter(filterName string, symbol string, exchange string, value float64, t time.Time) error
	SetAvailablePairsForExchange(exchange string, pairs []dia.Pair) error
	GetAvailablePairsForExchange(exchange string) ([]dia.Pair, error)
	SetCurrencyChange(cc *Change) error
	GetCurrencyChange() (*Change, error)
	GetAllSymbols() []string
	GetCoins() (*Coins, error)
	GetSymbolDetails(symbol string) (*SymbolDetails, error)
	UpdateSymbolDetails(symbol string, rank int)
	GetConfigTogglePairDiscovery() (bool, error)
	GetExchanges() []string
	SetOptionMeta(optionMeta *dia.OptionMeta) error
	GetOptionMeta(baseCurrency string) ([]dia.OptionMeta, error)
	SaveCVIInflux(float64, time.Time) error
	GetCVIInflux(time.Time, time.Time) ([]dia.CviDataPoint, error)

	// Interest rates' methods
	SetInterestRate(ir *InterestRate) error
	GetInterestRate(symbol, date string) (*InterestRate, error)
	GetInterestRateRange(symbol, dateInit, dateFinal string) ([]*InterestRate, error)
	GetRatesMeta() (RatesMeta []InterestRateMeta, err error)
	GetCompoundedIndex(symbol string, date time.Time, daysPerYear int) (*InterestRate, error)
	GetCompoundedIndexRange(symbol string, dateInit, dateFinal time.Time, daysPerYear int) ([]*InterestRate, error)
	GetCompoundedAvg(symbol string, date time.Time, calDays, daysPerYear int) (*InterestRate, error)
}

const (
	influxMaxPointsInBatch = 5000
	timeOutRedisOneBlock   = 60 * 3 * time.Second
)

type DB struct {
	redisClient         *redis.Client
	influxClient        clientInfluxdb.Client
	influxBatchPoints   clientInfluxdb.BatchPoints
	influxPointsInBatch int
}

const (
	influxDbName         = "dia"
	influxDbTradesTable  = "trades"
	influxDbFiltersTable = "filters"
	influxDbOptionsTable = "options"
	influxDbCVITable     = "cvi"
)

// queryInfluxDB convenience function to query the database
func queryInfluxDB(clnt clientInfluxdb.Client, cmd string) (res []clientInfluxdb.Result, err error) {
	q := clientInfluxdb.Query{
		Command:  cmd,
		Database: influxDbName,
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
	var ci clientInfluxdb.Client
	var bp clientInfluxdb.BatchPoints
	var r *redis.Client
	var err error
	// This environment variable is either set in docker-compose or empty
	executionMode := os.Getenv("EXEC_MODE")
	address := ""

	if withRedis {
		// Run localhost for testing and server for production
		if executionMode == "production" {
			address = "redis:6379"
		} else {
			address = "localhost:6379"
		}
		r = redis.NewClient(&redis.Options{
			Addr:     address,
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		pong2, err := r.Ping().Result()
		if err != nil {
			log.Error("NewDataStore redis", err)
		}
		log.Debug("NewDB", pong2)
	}
	if withInflux {
		if executionMode == "production" {
			address = "http://influxdb:8086"
		} else {
			address = "http://localhost:8086"
		}
		ci, err = clientInfluxdb.NewHTTPClient(clientInfluxdb.HTTPConfig{
			Addr:     address,
			Username: "",
			Password: "",
		})
		if err != nil {
			log.Error("NewDataStore influxdb", err)
		}
		bp, _ = createBatchInflux()
		_, err = queryInfluxDB(ci, fmt.Sprintf("CREATE DATABASE %s", influxDbName))
		if err != nil {
			log.Errorln("queryInfluxDB CREATE DATABASE", err)
		}
	}
	return &DB{r, ci, bp, 0}, nil
}

func createBatchInflux() (clientInfluxdb.BatchPoints, error) {
	bp, err := clientInfluxdb.NewBatchPoints(clientInfluxdb.BatchPointsConfig{
		Database:  influxDbName,
		Precision: "s",
	})
	if err != nil {
		log.Errorln("NewBatchPoints", err)
	}
	return bp, err
}

func (db *DB) Flush() error {
	var err error
	if db.influxBatchPoints != nil {
		err = db.WriteBatchInflux()
	}
	return err
}

func getKey(filter string, symbol string, exchange string) string {
	key := filter + "_" + symbol
	if exchange != "" {
		key = key + "_" + exchange
	}
	return key
}

func getKeyFilterZSET(key string) string {
	return "dia_" + key + "_ZSET"
}

func getKeyFilterSymbolAndExchangeZSET(filter string, symbol string, exchange string) string {
	if exchange == "" {
		return "dia_" + filter + "_" + symbol + "_ZSET"
	} else {
		return "dia_" + filter + "_" + symbol + "_" + exchange + "_ZSET"
	}
}

func (db *DB) WriteBatchInflux() error {
	err := db.influxClient.Write(db.influxBatchPoints)
	if err != nil {
		log.Errorln("WriteBatchInflux", err)
		db.influxBatchPoints, _ = createBatchInflux()
	} else {
		db.influxPointsInBatch = 0
	}
	return err
}

func (db *DB) addPoint(pt *clientInfluxdb.Point) {
	db.influxBatchPoints.AddPoint(pt)
	db.influxPointsInBatch++
	if db.influxPointsInBatch >= influxMaxPointsInBatch {
		log.Debug("AddPoint forcing write Bash")
		db.WriteBatchInflux()
	}
}

/*
select sum(value) from filters where filter='VOL120' and time > now() - 10m
select * from filters where  symbol='BTC' and filter='VOL120' and time > now() - 2m
select sum(value) from filters where  symbol='BTC' and filter='VOL120' and time > now()- 2m
*/

func (db *DB) Sum24HoursInflux(symbol string, exchange string, filter string) (*float64, error) {
	q := fmt.Sprintf("SELECT SUM(value) FROM %s WHERE symbol='%s' and exchange='%s' and filter='%s' and time > now() - 1d", influxDbFiltersTable, symbol, exchange, filter)
	var errorString string
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		log.Errorln("Sum24HoursInflux", err)
		return nil, err
	}
	if len(res) > 0 && len(res[0].Series) > 0 {
		for _, row := range res[0].Series[0].Values {

			var result float64
			v, o := row[1].(json.Number)
			if o {
				result, _ = v.Float64()
				return &result, nil
			} else {
				errorString = "error on parsing row 1"
				log.Errorln(errorString)
				return nil, errors.New(errorString)
			}
		}
	} else {
		errorString = "Empty res"
		log.Errorln(errorString)
		return nil, errors.New(errorString)
	}
	return nil, errors.New("couldnt sum in Sum24HoursInflux")
}

func (db *DB) SaveTradeInflux(t *dia.Trade) error {
	// Create a point and add to batch
	tags := map[string]string{"symbol": t.Symbol, "exchange": t.Source, "pair": t.Pair}
	fields := map[string]interface{}{
		"price":             t.Price,
		"volume":            t.Volume,
		"estimatedUSDPrice": t.EstimatedUSDPrice,
		"foreignTradeID":    t.ForeignTradeID,
	}

	pt, err := clientInfluxdb.NewPoint(influxDbTradesTable, tags, fields, t.Time)
	if err != nil {
		log.Errorln("NewTradeInflux:", err)
	} else {
		db.addPoint(pt)
	}
	return err
}

func (db *DB) SaveCVIInflux(cviValue float64, observationTime time.Time) error {
	fields := map[string]interface{}{
		"value": cviValue,
	}
	pt, err := clientInfluxdb.NewPoint(influxDbCVITable, nil, fields, observationTime)
	if err != nil {
		log.Errorln("NewOptionInflux:", err)
	} else {
		db.addPoint(pt)
	}

	err = db.WriteBatchInflux()
	if err != nil {
		log.Errorln("SaveOptionOrderbookDatumInflux", err)
	}

	return err
}

func (db *DB) GetCVIInflux(starttime time.Time, endtime time.Time) ([]dia.CviDataPoint, error) {
	retval := []dia.CviDataPoint{}
	q := fmt.Sprintf("SELECT * FROM %s WHERE time > %d and time < %d", influxDbCVITable, starttime.UnixNano(), endtime.UnixNano())
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		return retval, err
	}
	if len(res) > 0 && len(res[0].Series) > 0 {
		for i := 0; i < len(res[0].Series[0].Values); i++ {
			currentPoint := dia.CviDataPoint{}
			currentPoint.Timestamp, err = time.Parse(time.RFC3339, res[0].Series[0].Values[i][0].(string))
			if err != nil {
				return retval, err
			}
			currentPoint.Value, err = res[0].Series[0].Values[i][1].(json.Number).Float64()
			if err != nil {
				return retval, err
			}
			retval = append(retval, currentPoint)
		}
	} else {
		return retval, errors.New("Error parsing CVI value from Database")
	}
	return retval, nil
}

func (db *DB) SaveOptionOrderbookDatumInflux(t dia.OptionOrderbookDatum) error {
	tags := map[string]string{"instrumentName": t.InstrumentName}
	fields := map[string]interface{}{
		"askPrice": t.AskPrice,
		"bidPrice": t.BidPrice,
		"askSize":  t.AskSize,
		"bidSize":  t.BidSize,
	}
	pt, err := clientInfluxdb.NewPoint(influxDbOptionsTable, tags, fields, t.ObservationTime)
	if err != nil {
		log.Errorln("NewOptionInflux:", err)
	} else {
		db.addPoint(pt)
	}

	err = db.WriteBatchInflux()
	if err != nil {
		log.Errorln("SaveOptionOrderbookDatumInflux", err)
	}

	return err
}

func (db *DB) GetOptionOrderbookDataInflux(t dia.OptionMeta) (dia.OptionOrderbookDatum, error) {
	retval := dia.OptionOrderbookDatum{}
	q := fmt.Sprintf("SELECT LAST(askPrice), bidPrice, askSize, bidSize, observationTime FROM %s WHERE instrumentName ='%s'", influxDbOptionsTable, t.InstrumentName)
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		return retval, err
	}
	if len(res) > 0 && len(res[0].Series) > 0 {
		retval.InstrumentName = t.InstrumentName
		retval.ObservationTime, err = time.Parse(time.RFC3339, res[0].Series[0].Values[0][0].(string))
		if err != nil {
			return retval, err
		}
		retval.AskPrice, err = res[0].Series[0].Values[0][1].(json.Number).Float64()
		if err != nil {
			return retval, err
		}
		retval.BidPrice, err = res[0].Series[0].Values[0][2].(json.Number).Float64()
		if err != nil {
			return retval, err
		}
		retval.AskSize, err = res[0].Series[0].Values[0][3].(json.Number).Float64()
		if err != nil {
			return retval, err
		}
		retval.BidSize, err = res[0].Series[0].Values[0][4].(json.Number).Float64()
		if err != nil {
			return retval, err
		}
		return retval, nil
	}
	return retval, nil
}

func (db *DB) SaveFilterInflux(filter string, symbol string, exchange string, value float64, t time.Time) error {
	// Create a point and add to batch
	tags := map[string]string{"filter": filter, "symbol": symbol, "exchange": exchange}
	fields := map[string]interface{}{
		"value":        value,
		"ignore":       false,
		"allExchanges": exchange == "",
	}
	pt, err := clientInfluxdb.NewPoint(influxDbFiltersTable, tags, fields, t)
	if err != nil {
		log.Errorln("newPoint:", err)
	} else {
		db.addPoint(pt)
	}
	return err
}

func (db *DB) setZSETValue(key string, value float64, unixTime int64, maxWindow int64) error {

	if db.redisClient == nil {
		return nil
	}

	member := strconv.FormatFloat(value, 'f', -1, 64) + " " + strconv.FormatInt(unixTime, 10)

	err := db.redisClient.ZAdd(key, &redis.Z{
		Score:  float64(unixTime),
		Member: member,
	}).Err()
	log.Debug("SetZSETValue ", key, member, unixTime)
	if err != nil {
		log.Errorf("Error: %v on SetZSETValue %v\n", err, key)
	}
	// purging old values
	err = db.redisClient.ZRemRangeByScore(key, "-inf", "("+strconv.FormatInt(unixTime-maxWindow, 10)).Err()
	if err != nil {
		log.Errorf("Error: %v on SetZSETValue %v\n", err, key)
	}

	if err := db.redisClient.Expire(key, TimeOutRedis).Err(); err != nil {
		log.Error(err)
	} //TODO put two commands together ?
	return err
}

func (db *DB) getZSETValue(key string, atUnixTime int64) (float64, error) {

	result := 0.0
	max := strconv.FormatInt(atUnixTime, 10)
	vals, err := db.redisClient.ZRangeByScoreWithScores(key, &redis.ZRangeBy{
		Min: "-inf",
		Max: max,
	}).Result()
	log.Debug(key, "vals: %v on getZSETValue maxScore: %v", vals, max)
	if err == nil {
		if len(vals) > 0 {
			fmt.Sscanf(vals[len(vals)-1].Member.(string), "%f", &result)
			log.Debugf("returned value: %v", result)
		} else {
			err = errors.New("getZSETValue no value found")
		}
	}
	return result, err
}

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

func (db *DB) getZSETLastValue(key string) (float64, int64, error) {
	value := 0.0
	var unixTime int64
	vals, err := db.redisClient.ZRange(key, -1, -1).Result()
	log.Debug(key, "on getZSETLastValue:", vals)
	if err == nil {
		if len(vals) == 1 {
			fmt.Sscanf(vals[0], "%f %d", &value, &unixTime)
			log.Debugf("returned value: %v", value)
		} else {
			err = errors.New("getZSETLastValue no value found")
			log.Errorln("Error: on getZSETLastValue", err, key)
		}
	}
	return value, unixTime, err
}
