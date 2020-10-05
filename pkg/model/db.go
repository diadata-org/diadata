package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
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
	GetLatestSupply(string) (*dia.Supply, error)
	GetSupply(string, time.Time, time.Time) ([]dia.Supply, error)
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
	GetSymbolsByExchange(string) []string
	GetCoins() (*Coins, error)
	GetSymbolDetails(symbol string) (*SymbolDetails, error)
	UpdateSymbolDetails(symbol string, rank int)
	GetConfigTogglePairDiscovery() (bool, error)
	GetExchanges() []string
	SetOptionMeta(optionMeta *dia.OptionMeta) error
	GetOptionMeta(baseCurrency string) ([]dia.OptionMeta, error)
	SaveCVIInflux(float64, time.Time) error
	GetCVIInflux(time.Time, time.Time) ([]dia.CviDataPoint, error)
	GetSupplyInflux(string, time.Time, time.Time) ([]dia.Supply, error)
	GetVolumeInflux(string, time.Time, time.Time) (float64, error)
	Get24Volume(symbol string, exchange string) (float64, error)
	Get24VolumeExchange(exchange string) (float64, error)

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

	// Itin methods
	SetItinData(token dia.ItinToken) error
	GetItinBySymbol(symbol string) (dia.ItinToken, error)

	// Defi rates
	SetDefiProtocol(dia.DefiProtocol) error
	GetDefiProtocol(string) (dia.DefiProtocol, error)
	GetDefiProtocols() ([]dia.DefiProtocol, error)

	GetDefiRateInflux(time.Time, time.Time, string, string) ([]dia.DefiRate, error)
	SetDefiRateInflux(rate *dia.DefiRate) error

	GetDefiStateInflux(time.Time, time.Time, string) ([]dia.DefiProtocolState, error)
	SetDefiStateInflux(state *dia.DefiProtocolState) error

	// Foreign quotation methods
	SaveForeignQuotationInflux(fq ForeignQuotation) error
	GetForeignQuotationInflux(symbol, source string, timestamp time.Time) (ForeignQuotation, error)
	GetForeignPriceYesterday(symbol, source string) (float64, error)
	GetForeignSymbolsInflux(source string) (symbols []SymbolShort, err error)

	// Token methods
	// SaveTokenDetailInflux(tk Token) error
	// GetTokenDetailInflux(symbol, source string, timestamp time.Time) (Token, error)
	// GetCurentTotalSupply(symbol, source string) (float64, error)
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
	influxDbName           = "dia"
	influxDbTradesTable    = "trades"
	influxDbFiltersTable   = "filters"
	influxDbOptionsTable   = "options"
	influxDbCVITable       = "cvi"
	influxDbSupplyTable    = "supply"
	influxDbDefiRateTable  = "defiRate"
	influxDbDefiStateTable = "defiState"
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

// Get24Volume returns the volume in USD traded in the last 24 hours corresponding to quote token
// @symbol on exchange @exchange. It uses trades' volumes without filtering.
func (db *DB) Get24Volume(symbol string, exchange string) (float64, error) {
	q := fmt.Sprintf("SELECT estimatedUSDPrice, volume FROM %s WHERE symbol='%s' and exchange='%s' and time > now() - 1d", influxDbTradesTable, symbol, exchange)
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		log.Errorf("Get24HoursVolume for %s on %s: %v \n", symbol, exchange, err)
		return float64(0), err
	}
	var VolumeUSD float64
	if len(res) > 0 && len(res[0].Series) > 0 {
		for _, row := range res[0].Series[0].Values {
			USDPrice, err := strconv.ParseFloat(row[1].(string), 64)
			if err != nil {
				return float64(0), err
			}
			volume, err := strconv.ParseFloat(row[2].(string), 64)
			if err != nil {
				return 0, err
			}
			VolumeUSD += math.Abs(volume) * USDPrice
		}
		return VolumeUSD, nil
	}
	return float64(0), errors.New("no trades")
}

// Get24VolumeExchange returns the trade volume in USD traded on exchange @exchange.
// Uses trades' volumes without filtering.
func (db *DB) Get24VolumeExchange(exchange string) (float64, error) {
	allSymbols := db.GetSymbolsByExchange(exchange)
	var TVL float64
	for _, symbol := range allSymbols {
		volumeUSD, err := db.Get24Volume(symbol, exchange)
		if err != nil {
			log.Errorf("Error getting 24h trade volume of %s \n", symbol)
			continue
		}
		TVL += volumeUSD
	}
	return TVL, nil
}

/*
select sum(value) from filters where filter='VOL120' and time > now() - 10m
select * from filters where  symbol='BTC' and filter='VOL120' and time > now() - 2m
select sum(value) from filters where  symbol='BTC' and filter='VOL120' and time > now()- 2m
*/

// Sum24HoursInflux returns the 24h  volume of @symbol on @exchange using the filter @filter.
func (db *DB) Sum24HoursInflux(symbol string, exchange string, filter string) (*float64, error) {
	q := fmt.Sprintf("SELECT SUM(value) FROM %s WHERE symbol='%s' and exchange='%s' and filter='%s' and time > now() - 1d", influxDbFiltersTable, symbol, exchange, filter)
	var errorString string
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		log.Errorln("Sum24HoursInflux ", err)
		return nil, err
	}
	if len(res) > 0 && len(res[0].Series) > 0 {
		for _, row := range res[0].Series[0].Values {
			var result float64
			v, o := row[1].(json.Number)
			if o {
				result, _ = v.Float64()
				return &result, nil
			}
			errorString = "error on parsing row 1"
			log.Errorln(errorString)
			return nil, errors.New(errorString)

		}
	} else {
		errorString = "Empty response in Sum24HoursInflux"
		log.Errorln(errorString)
		return nil, errors.New(errorString)
	}
	return nil, errors.New("couldn't sum in Sum24HoursInflux")
}

// GetVolumeInflux returns the trade volume of @symbol in the time range @starttime - @endtime.
// It uses the VOL filter from the filter services.
func (db *DB) GetVolumeInflux(symbol string, starttime time.Time, endtime time.Time) (float64, error) {
	var retval float64
	var q string
	filter := "VOL120"
	if starttime.IsZero() || endtime.IsZero() {
		q = fmt.Sprintf("SELECT SUM(value) FROM %s WHERE symbol='%s' and filter='%s' and time > now() - 1d", influxDbFiltersTable, symbol, filter)
	} else {
		q = fmt.Sprintf("SELECT SUM(value) FROM %s WHERE symbol='%s' and filter='%s' and time > %d and time < %d", influxDbFiltersTable, symbol, filter, starttime.UnixNano(), endtime.UnixNano())
	}
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		return retval, err
	}
	if len(res) > 0 && len(res[0].Series) > 0 {
		for _, row := range res[0].Series[0].Values {
			v, o := row[1].(json.Number)
			if o {
				retval, _ = v.Float64()
				return retval, nil
			} else {
				errorString := "error on parsing row 1"
				log.Errorln(errorString)
				return retval, errors.New(errorString)
			}
		}
	} else {
		return retval, errors.New("Error parsing Volume value from Database")
	}
	return retval, nil
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

func (db *DB) SaveSupplyInflux(supply *dia.Supply) error {
	fields := map[string]interface{}{
		"supply": supply.CirculatingSupply,
		"source": supply.Source,
	}
	tags := map[string]string{
		"symbol": supply.Symbol,
		"name":   supply.Name,
	}
	pt, err := clientInfluxdb.NewPoint(influxDbSupplyTable, tags, fields, supply.Time)
	if err != nil {
		log.Errorln("NewSupplyInflux:", err)
	} else {
		db.addPoint(pt)
	}

	err = db.WriteBatchInflux()
	if err != nil {
		log.Errorln("SaveSupplyInflux", err)
	}

	return err
}

func (db *DB) SetDefiRateInflux(rate *dia.DefiRate) error {
	fields := map[string]interface{}{
		"lendingRate": rate.LendingRate,
		"borrowRate":  rate.BorrowingRate,
	}
	tags := map[string]string{
		"asset":    rate.Asset,
		"protocol": rate.Protocol,
	}
	pt, err := clientInfluxdb.NewPoint(influxDbDefiRateTable, tags, fields, rate.Timestamp)
	if err != nil {
		log.Errorln("SetDefiRateInflux:", err)
	} else {
		db.addPoint(pt)
	}

	err = db.WriteBatchInflux()
	if err != nil {
		log.Errorln("SetDefiRateInflux", err)
	}

	return err
}

func (db *DB) GetDefiRateInflux(starttime time.Time, endtime time.Time, asset string, protocol string) ([]dia.DefiRate, error) {
	retval := []dia.DefiRate{}
	q := fmt.Sprintf("SELECT * FROM %s WHERE time > %d and time < %d and asset = '%s' and protocol = '%s'", influxDbDefiRateTable, starttime.UnixNano(), endtime.UnixNano(), asset, protocol)
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		return retval, err
	}
	if len(res) > 0 && len(res[0].Series) > 0 {
		for i := 0; i < len(res[0].Series[0].Values); i++ {
			currentRate := dia.DefiRate{}
			currentRate.Timestamp, err = time.Parse(time.RFC3339, res[0].Series[0].Values[i][0].(string))
			if err != nil {
				return retval, err
			}
			currentRate.Asset = res[0].Series[0].Values[i][1].(string)
			if err != nil {
				return retval, err
			}
			currentRate.BorrowingRate, err = res[0].Series[0].Values[i][2].(json.Number).Float64()
			if err != nil {
				return retval, err
			}
			currentRate.LendingRate, err = res[0].Series[0].Values[i][3].(json.Number).Float64()
			if err != nil {
				return retval, err
			}
			currentRate.Protocol = protocol
			retval = append(retval, currentRate)
		}
	} else {
		return retval, errors.New("Error parsing Defi Lending Rate from Database")
	}
	return retval, nil
}

func (db *DB) SetDefiStateInflux(state *dia.DefiProtocolState) error {
	fields := map[string]interface{}{
		"totalUSD": state.TotalUSD,
		"totalETH": state.TotalETH,
	}
	tags := map[string]string{
		"protocol": state.Protocol.Name,
	}
	pt, err := clientInfluxdb.NewPoint(influxDbDefiStateTable, tags, fields, state.Timestamp)
	if err != nil {
		log.Errorln("SetDefiStateInflux:", err)
	} else {
		db.addPoint(pt)
	}

	err = db.WriteBatchInflux()
	if err != nil {
		log.Errorln("SetDefiStateInflux", err)
	}

	return err
}

func (db *DB) GetDefiStateInflux(starttime time.Time, endtime time.Time, protocol string) (retval []dia.DefiProtocolState, err error) {
	q := fmt.Sprintf("SELECT * FROM %s WHERE time > %d and time < %d and protocol = '%s'", influxDbDefiStateTable, starttime.UnixNano(), endtime.UnixNano(), protocol)
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		return retval, err
	}
	if len(res) > 0 && len(res[0].Series) > 0 {
		for i := 0; i < len(res[0].Series[0].Values); i++ {
			defiState := dia.DefiProtocolState{}
			defiState.Timestamp, err = time.Parse(time.RFC3339, res[0].Series[0].Values[i][0].(string))
			if err != nil {
				return
			}
			// defiState.Protocol.Name = res[0].Series[0].Values[i][1].(string)
			// if err != nil {
			// 	return
			// }
			defiState.TotalETH, err = res[0].Series[0].Values[i][2].(json.Number).Float64()
			if err != nil {
				return
			}
			defiState.TotalUSD, err = res[0].Series[0].Values[i][3].(json.Number).Float64()
			if err != nil {
				return
			}
			defiState.Protocol, err = db.GetDefiProtocol(protocol)
			if err != nil {
				return
			}
			retval = append(retval, defiState)
		}
	} else {
		err = errors.New("Error parsing Defi Lending Rate from Database")
		return
	}
	return
}

func (db *DB) GetSupplyInflux(symbol string, starttime time.Time, endtime time.Time) ([]dia.Supply, error) {
	retval := []dia.Supply{}
	var q string
	if starttime.IsZero() || endtime.IsZero() {
		q = fmt.Sprintf("SELECT supply,source FROM %s WHERE symbol = '%s' ORDER BY time DESC LIMIT 1", influxDbSupplyTable, symbol)
	} else {
		q = fmt.Sprintf("SELECT supply,source FROM %s WHERE time > %d and time < %d and symbol = '%s'", influxDbSupplyTable, starttime.UnixNano(), endtime.UnixNano(), symbol)
	}
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		return retval, err
	}
	if len(res) > 0 && len(res[0].Series) > 0 {
		for i := 0; i < len(res[0].Series[0].Values); i++ {
			currentSupply := dia.Supply{}
			currentSupply.Time, err = time.Parse(time.RFC3339, res[0].Series[0].Values[i][0].(string))
			if err != nil {
				return retval, err
			}
			currentSupply.CirculatingSupply, err = res[0].Series[0].Values[i][1].(json.Number).Float64()
			if err != nil {
				return retval, err
			}
			currentSupply.Source = res[0].Series[0].Values[i][2].(string)
			if err != nil {
				return retval, err
			}
			currentSupply.Symbol = symbol
			currentSupply.Name = helpers.NameForSymbol(symbol)
			retval = append(retval, currentSupply)
		}
	} else {
		return retval, errors.New("Error parsing Supply value from Database")
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

	err := db.redisClient.ZAdd(key, redis.Z{
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
	vals, err := db.redisClient.ZRangeByScoreWithScores(key, redis.ZRangeBy{
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
