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
	GetVolume(asset dia.Asset) (*float64, error)

	// Deprecating
	SetPriceUSD(symbol string, price float64) error
	SetPriceEUR(symbol string, price float64) error
	GetPriceUSD(symbol string) (float64, error)
	GetQuotation(symbol string) (*Quotation, error)
	SetQuotation(quotation *Quotation) error
	SetQuotationEUR(quotation *Quotation) error
	SetBatchFiatPriceInflux(fqs []*FiatQuotation) error
	SetSingleFiatPriceRedis(fiatQuotation *FiatQuotation) error

	GetLatestSupply(string) (*dia.Supply, error)
	GetSupply(string, time.Time, time.Time) ([]dia.Supply, error)
	SetSupply(supply *dia.Supply) error
	GetSupplyInflux(dia.Asset, time.Time, time.Time) ([]dia.Supply, error)
	// Deprecating: GetPairs(exchange string) ([]dia.ExchangePair, error)
	GetSymbols(exchange string) ([]string, error)
	// Deprecating: GetExchangesForSymbol(symbol string) ([]string, error)
	// Deprecating: GetSymbolExchangeDetails(symbol string, exchange string) (*SymbolExchangeDetails, error)
	GetLastTradeTimeForExchange(asset dia.Asset, exchange string) (*time.Time, error)
	SetLastTradeTimeForExchange(asset dia.Asset, exchange string, t time.Time) error
	SaveTradeInflux(t *dia.Trade) error
	GetTradeInflux(dia.Asset, string, time.Time) (*dia.Trade, error)
	SaveFilterInflux(filter string, asset dia.Asset, exchange string, value float64, t time.Time) error
	GetLastTrades(asset dia.Asset, exchange string, maxTrades int) ([]dia.Trade, error)
	GetAllTrades(t time.Time, maxTrades int) ([]dia.Trade, error)
	Flush() error
	GetFilterPoints(filter string, exchange string, symbol string, scale string, starttime time.Time, endtime time.Time) (*Points, error)
	SetFilter(filterName string, asset dia.Asset, exchange string, value float64, t time.Time) error
	GetLastPriceBefore(asset dia.Asset, filter string, exchange string, timestamp time.Time) (Price, error)
	SetAvailablePairs(exchange string, pairs []dia.ExchangePair) error
	GetAvailablePairs(exchange string) ([]dia.ExchangePair, error)
	SetCurrencyChange(cc *Change) error
	GetCurrencyChange() (*Change, error)

	GetExchanges() []string
	SetOptionMeta(optionMeta *dia.OptionMeta) error
	GetOptionMeta(baseCurrency string) ([]dia.OptionMeta, error)
	SaveCVIInflux(float64, time.Time) error
	GetCVIInflux(time.Time, time.Time, string) ([]dia.CviDataPoint, error)
	GetVolumeInflux(dia.Asset, time.Time, time.Time) (float64, error)
	// Get24Volume(symbol string, exchange string) (float64, error)
	// Get24VolumeExchange(exchange string) (float64, error)
	Sum24HoursInflux(asset dia.Asset, exchange string, filter string) (*float64, error)
	Sum24HoursExchange(exchange string) (float64, error)

	// New Asset pricing methods: 23/02/2021
	SetAssetPriceUSD(asset dia.Asset, price float64, timestamp time.Time) error
	GetAssetPriceUSD(asset dia.Asset) (float64, error)
	SetAssetQuotation(quotation *AssetQuotation) error
	GetAssetQuotation(asset dia.Asset) (*AssetQuotation, error)
	GetSortedAssetQuotations(assets []dia.Asset) ([]AssetQuotation, error)
	AddAssetQuotationsToBatch(quotations []*AssetQuotation) error
	SetAssetQuotationCache(quotation *AssetQuotation) (bool, error)
	GetAssetQuotationCache(asset dia.Asset) (*AssetQuotation, error)
	GetAssetPriceUSDCache(asset dia.Asset) (price float64, err error)
	GetTopAsset(symbol string, relDB *RelDB) (dia.Asset, error)

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

	// Pool  methods
	SetFarmingPool(pr *FarmingPool) error
	GetFarmingPoolData(starttime, endtime time.Time, protocol, poolID string) ([]FarmingPool, error)
	GetFarmingPools() ([]FarmingPoolType, error)

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

	// Gold token methods
	GetPaxgQuotationOunces() (*Quotation, error)
	GetPaxgQuotationGrams() (*Quotation, error)
	// Crypto Index methods
	GetCryptoIndex(time.Time, time.Time, string) ([]CryptoIndex, error)
	SetCryptoIndex(index *CryptoIndex) error
	GetCryptoIndexConstituents(time.Time, time.Time, dia.Asset, string) ([]CryptoIndexConstituent, error)
	SetCryptoIndexConstituent(*CryptoIndexConstituent, dia.Asset) error
	GetCryptoIndexConstituentPrice(symbol string, date time.Time) (float64, error)
	GetIndexPrice(asset dia.Asset, time time.Time) (*dia.Trade, error)
	// Token methods
	// SaveTokenDetailInflux(tk Token) error
	// GetTokenDetailInflux(symbol, source string, timestamp time.Time) (Token, error)
	// GetCurentTotalSupply(symbol, source string) (float64, error)

	// Github methods
	SetCommit(commit GithubCommit) error
	GetCommitByDate(user, repository string, date time.Time) (GithubCommit, error)
	GetCommitByHash(user, repository, hash string) (GithubCommit, error)
	GetLatestCommit(user, repository string) (GithubCommit, error)
}

const (
	influxMaxPointsInBatch = 5000
	// timeOutRedisOneBlock   = 60 * 3 * time.Second
)

type DB struct {
	redisClient         *redis.Client
	influxClient        clientInfluxdb.Client
	influxBatchPoints   clientInfluxdb.BatchPoints
	influxPointsInBatch int
}

const (
	influxDbName                         = "dia"
	influxDbTradesTable                  = "trades"
	influxDbFiltersTable                 = "filters"
	influxDbFiatQuotationsTable          = "fiat"
	influxDbOptionsTable                 = "options"
	influxDbCVITable                     = "cvi"
	influxDbETHCVITable                  = "cviETH"
	influxDbSupplyTable                  = "supplies"
	influxDbDefiRateTable                = "defiRate"
	influxDbDefiStateTable               = "defiState"
	influxDbPoolTable                    = "defiPools"
	influxDbCryptoIndexTable             = "cryptoindex"
	influxDbCryptoIndexConstituentsTable = "cryptoindexconstituents"
	influxDbGithubCommitTable            = "githubcommits"
	influxDBAssetQuotationsTable         = "assetQuotations"
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
	var influxClient clientInfluxdb.Client
	var influxBatchPoints clientInfluxdb.BatchPoints
	var redisClient *redis.Client

	if withRedis {
		redisClient = db.GetRedisClient()
	}
	if withInflux {
		var err error
		influxClient = db.GetInfluxClient()
		influxBatchPoints = createBatchInflux()
		_, err = queryInfluxDB(influxClient, fmt.Sprintf("CREATE DATABASE %s", influxDbName))
		if err != nil {
			log.Errorln("queryInfluxDB CREATE DATABASE", err)
		}
	}
	return &DB{redisClient, influxClient, influxBatchPoints, 0}, nil
}

func createBatchInflux() clientInfluxdb.BatchPoints {
	bp, err := clientInfluxdb.NewBatchPoints(clientInfluxdb.BatchPointsConfig{
		Database:  influxDbName,
		Precision: "s",
	})
	if err != nil {
		log.Errorln("NewBatchPoints", err)
	}
	return bp
}

func (db *DB) Flush() error {
	var err error
	if db.influxBatchPoints != nil {
		err = db.WriteBatchInflux()
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

func (db *DB) WriteBatchInflux() error {
	err := db.influxClient.Write(db.influxBatchPoints)
	if err != nil {
		log.Errorln("WriteBatchInflux", err)
		db.influxBatchPoints = createBatchInflux()
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
		err := db.WriteBatchInflux()
		if err != nil {
			log.Error("add point to influx batch: ", err)
		}
	}
}

// // Get24Volume returns the volume in USD traded in the last 24 hours corresponding to quote token
// // @symbol on exchange @exchange. It uses trades' volumes without filtering.
// func (db *DB) Get24Volume(symbol string, exchange string) (float64, error) {
// 	q := fmt.Sprintf("SELECT estimatedUSDPrice, volume FROM %s WHERE symbol='%s' and exchange='%s' and time > now() - 1d", influxDbTradesTable, symbol, exchange)
// 	res, err := queryInfluxDB(db.influxClient, q)
// 	if err != nil {
// 		log.Errorf("Get24HoursVolume for %s on %s: %v \n", symbol, exchange, err)
// 		return float64(0), err
// 	}
// 	var VolumeUSD float64
// 	if len(res) > 0 && len(res[0].Series) > 0 {
// 		for _, row := range res[0].Series[0].Values {
// 			USDPrice, err := strconv.ParseFloat(row[1].(string), 64)
// 			if err != nil {
// 				return float64(0), err
// 			}
// 			volume, err := strconv.ParseFloat(row[2].(string), 64)
// 			if err != nil {
// 				return 0, err
// 			}
// 			VolumeUSD += math.Abs(volume) * USDPrice
// 		}
// 		return VolumeUSD, nil
// 	}
// 	return float64(0), errors.New("no trades")
// }

// // Get24VolumeExchange returns the trade volume in USD traded on exchange @exchange.
// // Uses trades' volumes without filtering.
// func (db *DB) Get24VolumeExchange(exchange string) (float64, error) {
// 	allSymbols := db.GetSymbolsByExchange(exchange)
// 	var TVL float64
// 	for _, symbol := range allSymbols {
// 		volumeUSD, err := db.Get24Volume(symbol, exchange)
// 		if err != nil {
// 			log.Errorf("Error getting 24h trade volume of %s: %v \n", symbol, err)
// 			continue
// 		}
// 		TVL += volumeUSD
// 	}
// 	return TVL, nil
// }

/*
select sum(value) from filters where filter='VOL120' and time > now() - 10m
select * from filters where  symbol='BTC' and filter='VOL120' and time > now() - 2m
select sum(value) from filters where  symbol='BTC' and filter='VOL120' and time > now()- 2m
*/

// Sum24HoursInflux returns the 24h  volume of @asset on @exchange using the filter @filter.
func (db *DB) Sum24HoursInflux(asset dia.Asset, exchange string, filter string) (*float64, error) {
	queryString := "SELECT SUM(value) FROM %s WHERE address='%s' and blockchain='%s' and exchange='%s' and filter='%s' and time > now() - 1d"
	q := fmt.Sprintf(queryString, influxDbFiltersTable, asset.Address, asset.Blockchain, exchange, filter)
	var errorString string
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		log.Errorln("Sum24HoursInflux ", err)
		return nil, err
	}

	if len(res) > 0 && len(res[0].Series) > 0 {

		var result float64
		v, o := res[0].Series[0].Values[0][1].(json.Number)
		if o {
			result, err = v.Float64()
			if err != nil {
				log.Error(err)
				return nil, err
			}
			return &result, nil
		}
		errorString = "error on parsing row 1"
		log.Errorln(errorString)
		return nil, errors.New(errorString)

	} else {
		errorString = "empty response in Sum24HoursInflux"
		log.Errorln(errorString)
		return nil, errors.New(errorString)
	}

}

// Sum24HoursExchange returns 24h trade volumes summed up for all assets on @exchange,
// using VOL120 filtered data from influx.
// TO DO: Rewrite for assets
func (db *DB) Sum24HoursExchange(exchange string) (float64, error) {
	// allSymbols := db.GetSymbolsByExchange(exchange)
	// filter := "VOL120"
	// var TVL float64
	// for _, symbol := range allSymbols {
	// 	volumeUSD, err := db.Sum24HoursInflux(symbol, exchange, filter)
	// 	if err != nil {
	// 		log.Errorf("Error getting 24h trade volume of %s: %v \n", symbol, err)
	// 		continue
	// 	}
	// 	TVL += *volumeUSD
	// }
	var TVL float64
	return TVL, nil

}

// GetVolumeInflux returns the trade volume of @asset in the time range @starttime - @endtime.
// It uses the VOL filter from the filter services.
func (db *DB) GetVolumeInflux(asset dia.Asset, starttime time.Time, endtime time.Time) (float64, error) {
	var retval float64
	var q string
	filter := "VOL120"
	if starttime.IsZero() || endtime.IsZero() {
		queryString := "SELECT SUM(value) FROM %s WHERE address='%s' and blockchain='%s' and filter='%s' and time > now() - 1d"
		q = fmt.Sprintf(queryString, influxDbFiltersTable, asset.Address, asset.Blockchain, filter)
	} else {
		queryString := "SELECT SUM(value) FROM %s WHERE address='%s' and blockchain='%s' and filter='%s' and time > %d and time < %d"
		q = fmt.Sprintf(queryString, influxDbFiltersTable, asset.Address, asset.Blockchain, filter, starttime.UnixNano(), endtime.UnixNano())
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
		return retval, errors.New("parsing volume value from database")
	}
	return retval, nil
}

// SaveTradeInflux stores a trade in influx. Flushed when more than maxPoints in batch.
func (db *DB) SaveTradeInflux(t *dia.Trade) error {
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

	pt, err := clientInfluxdb.NewPoint(influxDbTradesTable, tags, fields, t.Time)
	if err != nil {
		log.Errorln("NewTradeInflux:", err)
	} else {
		db.addPoint(pt)
	}
	return err
}

// GetTradeInflux returns
func (db *DB) GetTradeInflux(asset dia.Asset, exchange string, timestamp time.Time) (*dia.Trade, error) {
	retval := dia.Trade{}
	var q string
	if exchange != "" {
		queryString := "SELECT estimatedUSDPrice,\"exchange\",foreignTradeID,\"pair\",price,\"symbol\",volume FROM %s WHERE quotetokenaddress='%s' and quotetokenblockchain='%s' and echange='%s' and time < %d order by desc limit 1"
		q = fmt.Sprintf(queryString, influxDbTradesTable, asset.Address, asset.Blockchain, exchange, timestamp.UnixNano())
	} else {
		queryString := "SELECT estimatedUSDPrice,\"exchange\",foreignTradeID,\"pair\",price,\"symbol\",volume FROM %s WHERE quotetokenaddress='%s' and quotetokenblockchain='%s' and time < %d order by desc limit 1"
		q = fmt.Sprintf(queryString, influxDbTradesTable, asset.Address, asset.Blockchain, timestamp.UnixNano())
	}

	/// TODO
	res, err := queryInfluxDB(db.influxClient, q)
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

func (db *DB) SaveETHCVIInflux(cviValue float64, observationTime time.Time) error {
	fields := map[string]interface{}{
		"value": cviValue,
	}
	pt, err := clientInfluxdb.NewPoint(influxDbETHCVITable, nil, fields, observationTime)
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

func (db *DB) GetCVIInflux(starttime time.Time, endtime time.Time, symbol string) ([]dia.CviDataPoint, error) {
	retval := []dia.CviDataPoint{}
	var q string
	if symbol == "ETH" {
		q = fmt.Sprintf("SELECT * FROM %s WHERE time > %d and time < %d", influxDbETHCVITable, starttime.UnixNano(), endtime.UnixNano())

	} else {
		q = fmt.Sprintf("SELECT * FROM %s WHERE time > %d and time < %d", influxDbCVITable, starttime.UnixNano(), endtime.UnixNano())
	}

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
		return retval, errors.New("parsing CVI value from database")
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

func (db *DB) SetFarmingPool(pool *FarmingPool) error {
	fields := map[string]interface{}{
		"rate":        pool.Rate,
		"balance":     pool.Balance,
		"blockNumber": pool.BlockNumber,
	}
	inputAssetBytes, err := json.Marshal(pool.InputAsset)
	if err != nil {
		return err
	}
	outputAssetBytes, err := json.Marshal(pool.OutputAsset)
	if err != nil {
		return err
	}
	tags := map[string]string{
		"inputAssets":  string(inputAssetBytes),
		"outputAssets": string(outputAssetBytes),
		"protocol":     pool.ProtocolName,
		"poolID":       pool.PoolID,
	}
	pt, err := clientInfluxdb.NewPoint(influxDbPoolTable, tags, fields, pool.TimeStamp)
	if err != nil {
		log.Errorln("SetPoolInfo:", err)
	} else {
		db.addPoint(pt)
	}

	err = db.WriteBatchInflux()
	if err != nil {
		log.Errorln("SetPoolInfo", err)
	}

	return err
}

// GetFarmingPools returns all farming pool states in the given time range
// time, balance, blocknumber, inputAssets, outputAssets, poolID, protocol, rate
func (db *DB) GetFarmingPools() ([]FarmingPoolType, error) {
	var retval []FarmingPoolType
	// First get all protocols
	qProtocol := fmt.Sprintf("SHOW TAG VALUES FROM %s with key=%s", influxDbPoolTable, "protocol")
	fmt.Println("protocol query: ", qProtocol)
	resProtocols, err := queryInfluxDB(db.influxClient, qProtocol)
	if err != nil {
		return retval, err
	}

	if len(resProtocols) > 0 && len(resProtocols[0].Series) > 0 {
		for i := 0; i < len(resProtocols[0].Series[0].Values); i++ {
			protocolName := resProtocols[0].Series[0].Values[i][1].(string)
			// For each protocol, get available pools by ID
			qPoolIDs := fmt.Sprintf("SHOW TAG VALUES FROM %s with key=%s where protocol='%s'", influxDbPoolTable, "poolID", protocolName)
			resPoolIDs, err := queryInfluxDB(db.influxClient, qPoolIDs)
			if err != nil {
				return retval, err
			}
			if len(resPoolIDs) > 0 && len(resPoolIDs[0].Series) > 0 {
				for k := 0; k < len(resPoolIDs[0].Series[0].Values); k++ {
					poolType := FarmingPoolType{}
					poolType.ProtocolName = protocolName
					poolType.PoolID = resPoolIDs[0].Series[0].Values[k][1].(string)
					// Get input assets of pool
					qAssets := fmt.Sprintf("SHOW TAG VALUES FROM %s with key=%s where protocol='%s' and poolID='%s'", influxDbPoolTable, "inputAssets", protocolName, poolType.PoolID)
					resAssets, err := queryInfluxDB(db.influxClient, qAssets)
					if err != nil {
						return retval, err
					}
					inputAssets := []string{}
					err = json.Unmarshal([]byte(resAssets[0].Series[0].Values[0][1].(string)), &inputAssets)
					if err != nil {
						return retval, err
					}
					poolType.InputAsset = inputAssets

					retval = append(retval, poolType)
				}
			}
		}
	} else {
		return retval, errors.New("parsing farming pool from Database")
	}
	return retval, nil

}

// GetFarmingPoolData returns all farming pool states in the given time range
// time, balance, blocknumber, inputAssets, outputAssets, poolID, protocol, rate
func (db *DB) GetFarmingPoolData(starttime, endtime time.Time, protocol, poolID string) ([]FarmingPool, error) {
	retval := []FarmingPool{}
	influxQuery := "SELECT balance,blockNumber,\"inputAssets\",\"outputAssets\",\"poolID\",\"protocol\",rate FROM %s WHERE time > %d and time <= %d and protocol = '%s' and poolID='%s' order by desc"
	q := fmt.Sprintf(influxQuery, influxDbPoolTable, starttime.UnixNano(), endtime.UnixNano(), protocol, poolID)
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		return retval, err
	}
	if len(res) > 0 && len(res[0].Series) > 0 {
		for i := 0; i < len(res[0].Series[0].Values); i++ {
			pool := FarmingPool{}
			pool.TimeStamp, err = time.Parse(time.RFC3339, res[0].Series[0].Values[i][0].(string))
			if err != nil {
				return retval, err
			}
			pool.Balance, err = res[0].Series[0].Values[i][1].(json.Number).Float64()
			if err != nil {
				return retval, err
			}
			pool.BlockNumber, err = res[0].Series[0].Values[i][2].(json.Number).Int64()
			if err != nil {
				return retval, err
			}
			inputAssets := []string{}
			err = json.Unmarshal([]byte(res[0].Series[0].Values[i][3].(string)), &inputAssets)
			if err != nil {
				return retval, err
			}
			pool.InputAsset = inputAssets
			outputAssets := []string{}
			err = json.Unmarshal([]byte(res[0].Series[0].Values[i][4].(string)), &outputAssets)
			if err != nil {
				return retval, err
			}
			pool.OutputAsset = outputAssets
			pool.PoolID = res[0].Series[0].Values[i][5].(string)
			pool.ProtocolName = res[0].Series[0].Values[i][6].(string)
			pool.Rate, err = res[0].Series[0].Values[i][7].(json.Number).Float64()
			if err != nil {
				return retval, err
			}

			retval = append(retval, pool)
		}
	} else {
		return retval, errors.New("parsing farming pool from database")
	}
	return retval, nil

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
	influxQuery := "SELECT \"asset\",borrowRate,lendingRate,\"protocol\" FROM %s WHERE time > %d and time < %d and asset = '%s' and protocol = '%s'"
	q := fmt.Sprintf(influxQuery, influxDbDefiRateTable, starttime.UnixNano(), endtime.UnixNano(), asset, protocol)
	fmt.Println("influx query: ", q)
	res, err := queryInfluxDB(db.influxClient, q)
	fmt.Println("res, err: ", res, err)
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
		return retval, errors.New("parsing defi lending rate from database")
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
	influxQuery := "SELECT totalETH,totalUSD FROM %s WHERE time > %d and time < %d and protocol = '%s'"
	q := fmt.Sprintf(influxQuery, influxDbDefiStateTable, starttime.UnixNano(), endtime.UnixNano(), protocol)
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
			defiState.TotalETH, err = res[0].Series[0].Values[i][1].(json.Number).Float64()
			if err != nil {
				return
			}
			defiState.TotalUSD, err = res[0].Series[0].Values[i][2].(json.Number).Float64()
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
		err = errors.New("parsing defi lending state from database")
		return
	}
	return
}

func (db *DB) SaveSupplyInflux(supply *dia.Supply) error {
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
		db.addPoint(pt)
	}

	err = db.WriteBatchInflux()
	if err != nil {
		log.Errorln("SaveSupplyInflux", err)
	}

	return err
}

// GetSupplyInflux returns supply and circulating supply of @asset. Needs asset.Address and asset.Blockchain.
// If no time range is given it returns the latest supply.
func (db *DB) GetSupplyInflux(asset dia.Asset, starttime time.Time, endtime time.Time) ([]dia.Supply, error) {
	retval := []dia.Supply{}
	var q string
	if starttime.IsZero() || endtime.IsZero() {
		queryString := "SELECT supply,circulatingsupply,source,\"name\",\"symbol\" FROM %s WHERE \"address\" = '%s' and \"blockchain\"='%s' ORDER BY time DESC LIMIT 1"
		q = fmt.Sprintf(queryString, influxDbSupplyTable, asset.Address, asset.Blockchain)
	} else {
		queryString := "SELECT supply,circulatingsupply,source,\"name\",\"symbol\" FROM %s WHERE time > %d and time < %d and \"address\" = '%s' and \"blockchain\"='%s'"
		q = fmt.Sprintf(queryString, influxDbSupplyTable, starttime.UnixNano(), endtime.UnixNano(), asset.Address, asset.Blockchain)
	}
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		return retval, err
	}
	if len(res) > 0 && len(res[0].Series) > 0 {
		for i := 0; i < len(res[0].Series[0].Values); i++ {
			currentSupply := dia.Supply{Asset: asset}
			currentSupply.Time, err = time.Parse(time.RFC3339, res[0].Series[0].Values[i][0].(string))
			if err != nil {
				return retval, err
			}
			currentSupply.Supply, err = res[0].Series[0].Values[i][1].(json.Number).Float64()
			if err != nil {
				return retval, err
			}
			currentSupply.CirculatingSupply, err = res[0].Series[0].Values[i][2].(json.Number).Float64()
			if err != nil {
				return retval, err
			}
			currentSupply.Source = res[0].Series[0].Values[i][3].(string)
			if err != nil {
				return retval, err
			}
			currentSupply.Asset.Name = res[0].Series[0].Values[i][4].(string)
			if err != nil {
				log.Error("error getting symbol name from influx: ", err)
			}
			currentSupply.Asset.Symbol = res[0].Series[0].Values[i][5].(string)
			if err != nil {
				log.Error("error getting symbol name from influx: ", err)
			}
			retval = append(retval, currentSupply)
		}
	} else {
		return retval, errors.New("parsing supply value from database")
	}
	return retval, nil
}

// SaveFilterInflux stores a filter point in influx.
func (db *DB) SaveFilterInflux(filter string, asset dia.Asset, exchange string, value float64, t time.Time) error {
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
	if err = db.redisClient.Expire(key, TimeOutRedis).Err(); err != nil {
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

func (db *DB) getZSETLastValue(key string) (float64, int64, error) {
	value := 0.0
	var unixTime int64
	vals, err := db.redisClient.ZRange(key, -1, -1).Result()
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
