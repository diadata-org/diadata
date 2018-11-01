package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/go-redis/redis"
	clientInfluxdb "github.com/influxdata/influxdb/client/v2"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
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
}

const (
	influxMaxPointsInBatch = 500
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
	return NewDataStoreWithOptions(true)
}

func NewDataStoreWithOptions(red bool) (*DB, error) {
	var r *redis.Client
	if red {
		r = redis.NewClient(&redis.Options{
			Addr:     "redis:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		pong2, err := r.Ping().Result()
		if err != nil {
			log.Error("NewDataStore redis", err)
		}
		log.Debug("NewDB", pong2)
	}
	i, err := clientInfluxdb.NewHTTPClient(clientInfluxdb.HTTPConfig{
		Addr:     "http://influxdb:8086",
		Username: "",
		Password: "",
	})
	if err != nil {
		log.Error("NewDataStore influxdb", err)
	}

	bp, _ := createBatchInflux()

	_, err = queryInfluxDB(i, fmt.Sprintf("CREATE DATABASE %s", influxDbName))
	if err != nil {
		log.Errorln("queryInfluxDB CREATE DATABASE", err)
	}

	return &DB{r, i, bp, 0}, nil
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
		err = db.WriteBashInflux()
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

func (db *DB) WriteBashInflux() error {
	err := db.influxClient.Write(db.influxBatchPoints)
	if err != nil {
		log.Errorln("WriteBashInflux", err)
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
		db.WriteBashInflux()
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

	err := db.redisClient.ZAdd(key, redis.Z{
		Score:  float64(unixTime),
		Member: member,
	}).Err()
	log.Debug("SetZSETValue ", key, member, unixTime)
	if err != nil {
		log.Error("Error: %v on SetZSETValue %v\n", err, key)
	}
	// purging old values
	err = db.redisClient.ZRemRangeByScore(key, "-inf", "("+strconv.FormatInt(unixTime-maxWindow, 10)).Err()
	if err != nil {
		log.Error("Error: %v on SetZSETValue %v\n", err, key)
	}

	if err := db.redisClient.Expire(key, TimeOutRedis).Err(); err != nil {
		log.Error(err)
	} //TODO put two commands together ?
	return err
}

func (db *DB) getZSETValue(key string, atUnixTime int64) (float64, error) {

	result := 0.0
	max := strconv.FormatInt(atUnixTime, 10)
	vals, err := db.redisClient.ZRangeByScoreWithScores(key, redis.ZRangeBy{
		Min: "-inf",
		Max: max,
	}).Result()
	log.Debug(key, "vals: %v on getZSETValue maxScore: %v", vals, max)
	if err == nil {
		if len(vals) > 0 {
			fmt.Sscanf(vals[len(vals)-1].Member.(string), "%f", &result)
			log.Debug("returned value: %v", result)
		} else {
			err = errors.New("getZSETValue no value found")
		}
	}
	return result, err
}

func (db *DB) getZSETSum(key string) (*float64, error) {

	log.Debug("getZSETSum: %v \n", key)

	vals, err := db.redisClient.ZRange(key, 0, -1).Result()
	if err != nil {
		log.Error("Error: %v on getZSETSum %v\n", err, key)
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
			log.Debug("returned value: %v", value)
		} else {
			err = errors.New("getZSETLastValue no value found")
			log.Errorln("Error: on getZSETLastValue", err, key)
		}
	}
	return value, unixTime, err
}
