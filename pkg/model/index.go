package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
)

/*const (
	IndexNormalization = float64(9507172.247746756)
)*/

var (
	CONSTITUENTS_SERIAL_SEPARATOR       = ","
	CONSTITUENTS_SERIAL_ASSET_SEPARATOR = "-"
)

// CryptoIndex is the container for API endpoint CryptoIndex
type CryptoIndex struct {
	// The index has a price, hence is traded, hence must(?) correspond to some underlying asset
	// In case there is no underlying token, just fill the field @Symbol.
	Asset             dia.Asset
	Value             float64
	Price             float64
	Price1h           float64
	Price24h          float64
	Price7d           float64
	Price14d          float64
	Price30d          float64
	Volume24hUSD      float64
	CirculatingSupply float64
	Divisor           float64
	CalculationTime   time.Time
	Constituents      []CryptoIndexConstituent
}

type CryptoIndexConstituent struct {
	Asset             dia.Asset
	Price             float64
	PriceYesterday    float64
	PriceYesterweek   float64
	CirculatingSupply float64
	Weight            float64
	Percentage        float64
	CappingFactor     float64
	NumBaseTokens     float64
}

type CryptoIndexMintAmount struct {
	Name          string
	Symbol        string
	Address       string
	Amount        uint64
	RebalanceTime time.Time
}

// MarshalBinary -
func (e *CryptoIndex) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *CryptoIndex) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

// GetCryptoIndexTime returns the latest recorded timestamp in the range [@starttime, @endtime].
func (datastore *DB) GetCryptoIndexTime(starttime, endtime time.Time, symbol string) (time.Time, error) {
	var timestamp time.Time
	q := fmt.Sprintf("SELECT value FROM %s WHERE time > %d AND time <= %d AND symbol = '%s' ORDER BY DESC LIMIT 1", influxDbCryptoIndexTable, starttime.UnixNano(), endtime.UnixNano(), symbol)
	res, err := queryInfluxDB(datastore.influxClient, q)
	if err != nil {
		return timestamp, err
	}

	if len(res) > 0 && len(res[0].Series) > 0 {
		// Calculation time
		timestamp, err = time.Parse(time.RFC3339, res[0].Series[0].Values[0][0].(string))
		if err != nil {
			return timestamp, err
		}
	} else {
		return timestamp, errors.New("no index in given time-range")
	}
	return timestamp, nil
}

func (datastore *DB) GetCryptoIndex(starttime time.Time, endtime time.Time, symbol string, maxResults int) ([]CryptoIndex, error) {
	var retval []CryptoIndex
	var q string
	if maxResults > 0 {
		q = fmt.Sprintf("SELECT constituents,symbol,price,value,divisor,\"name\",address,blockchain FROM %s WHERE time > %d AND time <= %d AND \"symbol\" = '%s' ORDER BY DESC LIMIT %d", influxDbCryptoIndexTable, starttime.UnixNano(), endtime.UnixNano(), symbol, maxResults)
	} else {
		q = fmt.Sprintf("SELECT constituents,symbol,price,value,divisor,\"name\",address,blockchain FROM %s WHERE time > %d AND time <= %d AND \"symbol\" = '%s' ORDER BY DESC", influxDbCryptoIndexTable, starttime.UnixNano(), endtime.UnixNano(), symbol)
	}
	res, err := queryInfluxDB(datastore.influxClient, q)
	if err != nil {
		return retval, err
	}

	if len(res) > 0 && len(res[0].Series) > 0 {
		for i := 0; i < len(res[0].Series[0].Values); i++ {
			// Get index fields
			currentIndex := CryptoIndex{}
			// Symbol
			if res[0].Series[0].Values[i][2] != nil {
				currentIndex.Asset.Symbol = res[0].Series[0].Values[i][2].(string)
			}

			// Divisor and Value
			divisor := 9507172.247746756
			// Check if divisor exists in DB, otherwise use the default value
			if res[0].Series[0].Values[i][5] != nil {
				divisor, err = res[0].Series[0].Values[i][5].(json.Number).Float64()
				if err != nil {
					return retval, err
				}
			}
			currentIndex.Divisor = divisor
			tmp, err := res[0].Series[0].Values[i][4].(json.Number).Float64()
			if err != nil {
				return retval, err
			}
			tmp /= currentIndex.Divisor
			currentIndex.Value = tmp
			// Prices (actual, 1h, 24h, ...)
			// TO DO: instead of log.Error return error as soon as we have long enough data trail
			currentPrice, err := res[0].Series[0].Values[i][3].(json.Number).Float64()
			if err != nil {
				log.Error("error current index price: ", err)
				currentIndex.Price = 0
			} else {
				currentIndex.Price = currentPrice
			}
			log.Infof("get trade price 1h for %s..", currentIndex.Asset.Symbol)
			price1h, err := datastore.GetTradePrice1h(currentIndex.Asset, "")
			if err != nil {
				log.Error("error index price 1h: ", err)
				currentIndex.Price1h = 0.0
			} else {
				currentIndex.Price1h = price1h.EstimatedUSDPrice
			}

			log.Infof("get trade price 24h for %s..", currentIndex.Asset.Symbol)
			price24h, err := datastore.GetTradePrice24h(currentIndex.Asset, "")
			if err != nil {
				log.Error("error index price 24h: ", err)
				currentIndex.Price24h = 0.0
			} else {
				currentIndex.Price24h = price24h.EstimatedUSDPrice
			}

			log.Infof("get trade price 7d for %s..", currentIndex.Asset.Symbol)
			price7d, err := datastore.GetTradePrice7d(currentIndex.Asset, "")
			if err != nil {
				log.Error("error index price 7d: ", err)
				currentIndex.Price7d = 0.0
			} else {
				currentIndex.Price7d = price7d.EstimatedUSDPrice
			}

			log.Infof("get trade price 14d for %s..", currentIndex.Asset.Symbol)
			price14d, err := datastore.GetTradePrice14d(currentIndex.Asset, "")
			if err != nil {
				log.Error("error index price 14d: ", err)
				currentIndex.Price14d = 0.0
			} else {
				currentIndex.Price14d = price14d.EstimatedUSDPrice
			}

			log.Infof("get trade price 30d for %s..", currentIndex.Asset.Symbol)
			price30d, err := datastore.GetTradePrice30d(currentIndex.Asset, "")
			if err != nil {
				log.Error("error index price 30d: ", err)
				currentIndex.Price30d = 0.0
			} else {
				currentIndex.Price30d = price30d.EstimatedUSDPrice
			}
			// Circulating supply
			log.Infof("get supply for %s..", currentIndex.Asset.Symbol)
			diaSupply, err := datastore.GetSupplyInflux(currentIndex.Asset, time.Time{}, time.Time{})
			if err != nil || len(diaSupply) < 1 {
				log.Error(err)
				currentIndex.CirculatingSupply = 0
			} else {
				currentIndex.CirculatingSupply = diaSupply[0].CirculatingSupply
			}
			// Calculation time
			currentIndex.CalculationTime, err = time.Parse(time.RFC3339, res[0].Series[0].Values[i][0].(string))
			if err != nil {
				return retval, err
			}
			constituentsSerial := res[0].Series[0].Values[i][1].(string)
			// Name
			if res[0].Series[0].Values[i][6] != nil {
				currentIndex.Asset.Name = res[0].Series[0].Values[i][6].(string)
			}
			// Address
			if res[0].Series[0].Values[i][7] != nil {
				currentIndex.Asset.Address = res[0].Series[0].Values[i][7].(string)
			}
			// Blockchain
			if res[0].Series[0].Values[i][8] != nil {
				currentIndex.Asset.Blockchain = res[0].Series[0].Values[i][8].(string)
			}

			// Get constituents
			var constituents []CryptoIndexConstituent
			for _, constituent := range strings.Split(constituentsSerial, CONSTITUENTS_SERIAL_SEPARATOR) {
				if constituent == "" {
					log.Info("Skipping empty Address")
					continue
				}
				// Address and blockchain is sufficient information for constituent getter.
				constituentAddress := strings.Split(constituent, CONSTITUENTS_SERIAL_ASSET_SEPARATOR)[0]
				constituentBlockchain := strings.Split(constituent, CONSTITUENTS_SERIAL_ASSET_SEPARATOR)[1]
				constituentAsset := dia.Asset{
					Address:    constituentAddress,
					Blockchain: constituentBlockchain,
				}
				log.Infof("get constituent asset %s..", constituentAsset.Symbol)
				curr, err := datastore.GetCryptoIndexConstituents(currentIndex.CalculationTime.Add(-24*time.Hour), endtime, constituentAsset, symbol)
				if err != nil {
					return retval, err
				}
				if len(curr) > 0 {
					constituents = append(constituents, curr[0])
				}
			}
			currentIndex.Constituents = constituents
			retval = append(retval, currentIndex)
		}
	}
	return retval, nil
}

func (datastore *DB) SetCryptoIndex(index *CryptoIndex) error {
	constituentsSerial := ""
	for _, c := range index.Constituents {
		if constituentsSerial != "" {
			constituentsSerial += CONSTITUENTS_SERIAL_SEPARATOR
		}
		constituentsSerial += c.Asset.Address + CONSTITUENTS_SERIAL_ASSET_SEPARATOR + c.Asset.Blockchain
	}
	fields := map[string]interface{}{
		"price":        index.Price,
		"value":        index.Value,
		"constituents": constituentsSerial,
		"divisor":      index.Divisor,
	}
	tags := map[string]string{
		"symbol":     index.Asset.Symbol,
		"name":       index.Asset.Name,
		"address":    index.Asset.Address,
		"blockchain": index.Asset.Blockchain,
	}
	pt, err := clientInfluxdb.NewPoint(influxDbCryptoIndexTable, tags, fields, index.CalculationTime)
	if err != nil {
		log.Error("Writing Crypto Index to Influx: ", err)
		return err
	}
	datastore.addPoint(pt)

	err = datastore.WriteBatchInflux()
	if err != nil {
		log.Error("Writing Crypto Index to Influx: ", err)
		return err
	}

	for i := range index.Constituents {
		err = datastore.SetCryptoIndexConstituent(&index.Constituents[i], index.Asset, index.CalculationTime)
		if err != nil {
			return err
		}
	}
	return err
}

// GetCryptoIndexConstituentPrice returns the price of cryptoindexconstituent by @symbol. Not used at the moment.
func (datastore *DB) GetCryptoIndexConstituentPrice(symbol string, date time.Time) (float64, error) {
	startdate := date.Add(-24 * time.Hour)
	q := fmt.Sprintf("SELECT price from %s where time > %d and time <= %d and symbol = '%s' ORDER BY time DESC LIMIT 1", influxDbCryptoIndexConstituentsTable, startdate.UnixNano(), date.UnixNano(), symbol)
	res, err := queryInfluxDB(datastore.influxClient, q)
	if err != nil {
		return float64(0), err
	}
	var price float64
	if len(res) > 0 && len(res[0].Series) > 0 && len(res[0].Series[0].Values) > 0 {
		price, err = res[0].Series[0].Values[0][1].(json.Number).Float64()
		if err != nil {
			return 0, err
		}
	}
	return price, nil

}

// GetCryptoIndexConstituents returns the constituent corresponding to @asset along with underlying information.
// Necessary and sufficient information is asset's address and blockchain.
func (datastore *DB) GetCryptoIndexConstituents(starttime time.Time, endtime time.Time, asset dia.Asset, indexSymbol string) ([]CryptoIndexConstituent, error) {
	//func (db *DB) GetCryptoIndexConstituents(starttime time.Time, endtime time.Time, symbol string) ([]CryptoIndexConstituent, error) {
	var retval []CryptoIndexConstituent
	queryString := "SELECT \"address\",cappingfactor,circulatingsupply,\"name\",percentage,price,\"symbol\",weight,numbasetokens,\"blockchain\"" +
		" from %s WHERE time > %d and time < %d and address = '%s' and blockchain = '%s' and cryptoindex = '%s' ORDER BY time DESC LIMIT 1"
	q := fmt.Sprintf(queryString, influxDbCryptoIndexConstituentsTable, starttime.UnixNano(), endtime.UnixNano(), asset.Address, asset.Blockchain, indexSymbol)
	//q := fmt.Sprintf("SELECT address,cappingfactor,circulatingsupply,\"name\",percentage,price,symbol,weight,numbasetokens from %s WHERE time > %d and time < %d and symbol = '%s' ORDER BY time DESC LIMIT 1", influxDbCryptoIndexConstituentsTable, starttime.UnixNano(), endtime.UnixNano(), symbol)
	res, err := queryInfluxDB(datastore.influxClient, q)

	if err != nil {
		return retval, err
	}
	if len(res) > 0 && len(res[0].Series) > 0 && len(res[0].Series[0].Values) > 0 {

		currentConstituent := CryptoIndexConstituent{}
		//TODO: Do we need time?
		currentConstituent.Asset.Address = res[0].Series[0].Values[0][1].(string)
		currentConstituent.CappingFactor, err = res[0].Series[0].Values[0][2].(json.Number).Float64()
		if err != nil {
			return retval, err
		}
		currentConstituent.CirculatingSupply, err = res[0].Series[0].Values[0][3].(json.Number).Float64()
		if err != nil {
			return retval, err
		}
		currentConstituent.Asset.Name = res[0].Series[0].Values[0][4].(string)
		currentConstituent.Percentage, err = res[0].Series[0].Values[0][5].(json.Number).Float64()
		if err != nil {
			return retval, err
		}
		currentConstituent.Price, err = res[0].Series[0].Values[0][6].(json.Number).Float64()
		if err != nil {
			return retval, err
		}
		currentConstituent.Asset.Symbol = res[0].Series[0].Values[0][7].(string)
		currentConstituent.Weight, err = res[0].Series[0].Values[0][8].(json.Number).Float64()
		if err != nil {
			return retval, err
		}
		if res[0].Series[0].Values[0][9] != nil {
			currentConstituent.NumBaseTokens, err = res[0].Series[0].Values[0][9].(json.Number).Float64()
			if err != nil {
				return retval, err
			}
		} else {
			currentConstituent.NumBaseTokens = 0
		}
		currentConstituent.Asset.Blockchain = res[0].Series[0].Values[0][10].(string)

		// Get price yesterday
		priceYesterday, err := datastore.GetLastPriceBefore(currentConstituent.Asset, "MAIR120", "", endtime.AddDate(0, 0, -1))
		if err != nil {
			currentConstituent.PriceYesterday = float64(0)
		} else {
			currentConstituent.PriceYesterday = priceYesterday.Price
		}
		// Get price yesterweek
		priceYesterweek, err := datastore.GetLastPriceBefore(currentConstituent.Asset, "MAIR120", "", endtime.AddDate(0, 0, -7))
		if err != nil {
			currentConstituent.PriceYesterweek = float64(0)
		} else {
			currentConstituent.PriceYesterweek = priceYesterweek.Price
		}

		retval = append(retval, currentConstituent)

	}
	return retval, nil
}

func (datastore *DB) SetCryptoIndexConstituent(constituent *CryptoIndexConstituent, index dia.Asset, timestamp time.Time) error {
	fields := map[string]interface{}{
		"percentage":        constituent.Percentage,
		"price":             constituent.Price,
		"circulatingsupply": constituent.CirculatingSupply,
		"weight":            constituent.Weight,
		"cappingfactor":     constituent.CappingFactor,
		"numbasetokens":     constituent.NumBaseTokens,
	}
	tags := map[string]string{
		"name":        constituent.Asset.Name,
		"symbol":      constituent.Asset.Symbol,
		"address":     constituent.Asset.Address,
		"blockchain":  constituent.Asset.Blockchain,
		"cryptoindex": index.Symbol,
	}
	pt, err := clientInfluxdb.NewPoint(influxDbCryptoIndexConstituentsTable, tags, fields, timestamp)
	if err != nil {
		log.Error("Adding Crypto Index Constituent point to Influx batch: ", err)
		return err
	}
	datastore.addPoint(pt)

	err = datastore.WriteBatchInflux()
	if err != nil {
		log.Error("Writing Crypto Index Constituent to Influx: ", err)
	}
	return err
}

// GetIndexPrice returns the last price of index represented by @asset with respect to the
// time-range [time-window, time].
// If @asset only consists of a symbol, a different method for price retrieval has to be implemented.
func (datastore *DB) GetIndexPrice(asset dia.Asset, time time.Time, window time.Duration) (trade *dia.Trade, err error) {
	if asset.Address != "" && asset.Blockchain != "" {
		trade, err = datastore.GetTradeInflux(asset, "", time, window)
		return
	}
	err = errors.New("asset's address or blockchain missing")
	return
}

func (datastore *DB) GetCurrentIndexCompositionForIndex(index dia.Asset) []CryptoIndexConstituent {
	var constituents []CryptoIndexConstituent
	log.Infof("get crypto index %s from influx..", index.Symbol)
	cryptoIndex, err := datastore.GetCryptoIndex(time.Now().Add(-5*time.Hour), time.Now(), index.Symbol, 1)
	if err != nil {
		log.Error("get crypto index: ", err)
		return constituents
	}
	for _, constituent := range cryptoIndex[0].Constituents {
		curr, err := datastore.GetCryptoIndexConstituents(time.Now().Add(-5*time.Hour), time.Now(), constituent.Asset, index.Symbol)
		if err != nil {
			log.Error("get crypto index constituents: ", err)
			return constituents
		}
		if len(curr) > 0 {
			constituents = append(constituents, curr[0])
		}
	}
	return constituents
}

func (datastore *DB) IndexValueCalculation(currentConstituents []CryptoIndexConstituent, indexAsset dia.Asset, indexValue float64) CryptoIndex {

	var price float64
	tradeObject, err := datastore.GetIndexPrice(indexAsset, time.Now(), time.Duration(7*24*time.Hour))
	if err == nil {
		// Quotation does exist
		price = tradeObject.EstimatedUSDPrice
	}
	var circSupply float64
	supplyObject, err := datastore.GetSupplyInflux(indexAsset, time.Time{}, time.Time{})
	if err == nil && len(supplyObject) > 0 {
		// Supply does exist
		circSupply = supplyObject[0].CirculatingSupply
	}

	currCryptoIndex, err := datastore.GetCryptoIndex(time.Now().Add(-5*time.Hour), time.Now(), indexAsset.Symbol, 1)
	if err != nil {
		log.Error(err)
	}
	index := CryptoIndex{
		Asset:             indexAsset,
		Price:             price,
		CirculatingSupply: circSupply,
		Value:             indexValue,
		CalculationTime:   time.Now(),
		Constituents:      currentConstituents,
		Divisor:           currCryptoIndex[0].Divisor,
	}
	log.Info("Index: ", index)
	return index
}

func (datastore *DB) UpdateConstituentsMarketData(index string, currentConstituents *[]CryptoIndexConstituent) error {

	for i, c := range *currentConstituents {
		currSupply, err := datastore.GetSupplyInflux(c.Asset, time.Time{}, time.Time{})
		if err != nil {
			log.Error("Error when retrieveing supply for ", c.Asset.Symbol)
			return err
		}
		currLastTrade, err := datastore.GetLastTrades(c.Asset, "", 1, false)
		if err != nil {
			log.Error("Error when retrieveing last trades for ", c.Asset.Symbol)
			return err
		}
		(*currentConstituents)[i].Price = currLastTrade[0].EstimatedUSDPrice
		(*currentConstituents)[i].CirculatingSupply = currSupply[0].CirculatingSupply
	}

	// Calculate current percentages: 1. get index value 2. Determine percentage of each asset
	currIndexValue := GetIndexValue(index, *currentConstituents)
	if index == "SCIFI" {
		for i := range *currentConstituents {
			currPercentage := ((*currentConstituents)[i].Price * (*currentConstituents)[i].CirculatingSupply * (*currentConstituents)[i].CappingFactor) / currIndexValue
			(*currentConstituents)[i].Percentage = currPercentage
		}
	} else {
		// GBI
		for i := range *currentConstituents {
			currPercentage := ((*currentConstituents)[i].Price * (*currentConstituents)[i].NumBaseTokens * 1e-16) / currIndexValue //1e-16 because index value is 100 at start
			(*currentConstituents)[i].Percentage = currPercentage
		}
	}
	return nil
}

func GetIndexValue(indexSymbol string, currentConstituents []CryptoIndexConstituent) float64 {
	indexValue := 0.0
	if indexSymbol == "SCIFI" {
		for _, constituent := range currentConstituents {
			indexValue += constituent.Price * constituent.CirculatingSupply * constituent.CappingFactor
		}
	} else {
		// GBI etc
		for _, constituent := range currentConstituents {
			indexValue += constituent.Price * constituent.NumBaseTokens * 1e-16 //1e-16 because index value is 100 at start
		}
	}
	return indexValue
}
