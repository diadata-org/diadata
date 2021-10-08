package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
	log "github.com/sirupsen/logrus"
)

/*const (
	IndexNormalization = float64(9507172.247746756)
)*/

// CryptoIndex is the container for API endpoint CryptoIndex
type CryptoIndex struct {
	Name              string
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
	Name              string
	Symbol            string
	Address           string
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

func (db *DB) GetCryptoIndex(starttime time.Time, endtime time.Time, name string) ([]CryptoIndex, error) {
	var retval []CryptoIndex
	q := fmt.Sprintf("SELECT constituents,\"name\",price,value,divisor from %s WHERE time > %d and time < %d and \"name\" = '%s' ORDER BY time DESC LIMIT 1", influxDbCryptoIndexTable, starttime.UnixNano(), endtime.UnixNano(), name)
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		return retval, err
	}

	if len(res) > 0 && len(res[0].Series) > 0 {
		for i := 0; i < len(res[0].Series[0].Values); i++ {
			// Get index fields
			currentIndex := CryptoIndex{}
			currentIndex.Name = res[0].Series[0].Values[i][2].(string)
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
			price1h, err := db.GetTradePrice1h(currentIndex.Name, "")
			if err != nil {
				log.Error("error index price 1h: ", err)
				currentIndex.Price1h = 0.0
			} else {
				currentIndex.Price1h = price1h.EstimatedUSDPrice
			}

			price24h, err := db.GetTradePrice24h(currentIndex.Name, "")
			if err != nil {
				log.Error("error index price 24h: ", err)
				currentIndex.Price24h = 0.0
			} else {
				currentIndex.Price24h = price24h.EstimatedUSDPrice
			}

			price7d, err := db.GetTradePrice7d(currentIndex.Name, "")
			if err != nil {
				log.Error("error index price 7d: ", err)
				currentIndex.Price7d = 0.0
			} else {
				currentIndex.Price7d = price7d.EstimatedUSDPrice
			}

			price14d, err := db.GetTradePrice14d(currentIndex.Name, "")
			if err != nil {
				log.Error("error index price 14d: ", err)
				currentIndex.Price14d = 0.0
			} else {
				currentIndex.Price14d = price14d.EstimatedUSDPrice
			}

			price30d, err := db.GetTradePrice30d(currentIndex.Name, "")
			if err != nil {
				log.Error("error index price 30d: ", err)
				currentIndex.Price30d = 0.0
			} else {
				currentIndex.Price30d = price30d.EstimatedUSDPrice
			}
			// TODO: Volume
			// Circulating supply
			diaSupply, err := db.GetLatestSupply(currentIndex.Name)
			if err != nil {
				log.Error(err)
				currentIndex.CirculatingSupply = 0
			} else {
				currentIndex.CirculatingSupply = diaSupply.CirculatingSupply
			}
			// Calculation time
			currentIndex.CalculationTime, err = time.Parse(time.RFC3339, res[0].Series[0].Values[i][0].(string))
			if err != nil {
				return retval, err
			}
			constituentsSerial := res[0].Series[0].Values[i][1].(string)

			// Get constituents
			var constituents []CryptoIndexConstituent
			for _, constituentSymbol := range strings.Split(constituentsSerial, ",") {
				if constituentSymbol == "" {
					log.Info("Skipping empty Symbol")
					continue
				}
				curr, err := db.GetCryptoIndexConstituents(currentIndex.CalculationTime.Add(-24 * time.Hour), endtime, constituentSymbol, name)
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

func (db *DB) SetCryptoIndex(index *CryptoIndex) error {
	constituentsSerial := ""
	for _, c := range index.Constituents {
		if constituentsSerial != "" {
			constituentsSerial += ","
		}
		constituentsSerial += c.Symbol
	}
	fields := map[string]interface{}{
		"price":        index.Price,
		"value":        index.Value,
		"constituents": constituentsSerial,
		"divisor":      index.Divisor,
	}
	tags := map[string]string{
		"name": index.Name,
	}
	pt, err := clientInfluxdb.NewPoint(influxDbCryptoIndexTable, tags, fields, index.CalculationTime)
	if err != nil {
		log.Error("Writing Crypto Index to Influx: ", err)
		return err
	}
	db.addPoint(pt)

	err = db.WriteBatchInflux()
	if err != nil {
		log.Error("Writing Crypto Index to Influx: ", err)
		return err
	}

	for _, constituent := range index.Constituents {
		err = db.SetCryptoIndexConstituent(&constituent, index.Name)
		if err != nil {
			return err
		}
	}
	return err
}

func (db *DB) GetCryptoIndexConstituentPrice(symbol string, date time.Time) (float64, error) {
	startdate := date.Add(-24 * time.Hour)
	q := fmt.Sprintf("SELECT price from %s where time > %d and time <= %d and symbol = '%s' ORDER BY time DESC LIMIT 1", influxDbCryptoIndexConstituentsTable, startdate.UnixNano(), date.UnixNano(), symbol)
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		return float64(0), err
	}
	var price float64
	if len(res) > 0 && len(res[0].Series) > 0 && len(res[0].Series[0].Values) > 0 {
		price, err = res[0].Series[0].Values[0][1].(json.Number).Float64()
	}
	return price, nil

}

func (db *DB) GetCryptoIndexConstituents(starttime time.Time, endtime time.Time, symbol string, indexSymbol string) ([]CryptoIndexConstituent, error) {
//func (db *DB) GetCryptoIndexConstituents(starttime time.Time, endtime time.Time, symbol string) ([]CryptoIndexConstituent, error) {
	var retval []CryptoIndexConstituent

	q := fmt.Sprintf("SELECT address,cappingfactor,circulatingsupply,\"name\",percentage,price,symbol,weight,numbasetokens from %s WHERE time > %d and time < %d and symbol = '%s' and cryptoindex = '%s' ORDER BY time DESC LIMIT 1", influxDbCryptoIndexConstituentsTable, starttime.UnixNano(), endtime.UnixNano(), symbol, indexSymbol)
	//q := fmt.Sprintf("SELECT address,cappingfactor,circulatingsupply,\"name\",percentage,price,symbol,weight,numbasetokens from %s WHERE time > %d and time < %d and symbol = '%s' ORDER BY time DESC LIMIT 1", influxDbCryptoIndexConstituentsTable, starttime.UnixNano(), endtime.UnixNano(), symbol)
	res, err := queryInfluxDB(db.influxClient, q)

	if err != nil {
		return retval, err
	}
	if len(res) > 0 && len(res[0].Series) > 0 && len(res[0].Series[0].Values) > 0 {

		currentConstituent := CryptoIndexConstituent{}
		//TODO: Do we need time?
		currentConstituent.Address = res[0].Series[0].Values[0][1].(string)
		currentConstituent.CappingFactor, err = res[0].Series[0].Values[0][2].(json.Number).Float64()
		if err != nil {
			return retval, err
		}
		currentConstituent.CirculatingSupply, err = res[0].Series[0].Values[0][3].(json.Number).Float64()
		if err != nil {
			return retval, err
		}
		currentConstituent.Name = res[0].Series[0].Values[0][4].(string)
		currentConstituent.Percentage, err = res[0].Series[0].Values[0][5].(json.Number).Float64()
		if err != nil {
			return retval, err
		}
		currentConstituent.Price, err = res[0].Series[0].Values[0][6].(json.Number).Float64()
		if err != nil {
			return retval, err
		}
		currentConstituent.Symbol = res[0].Series[0].Values[0][7].(string)
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
		// Get price yesterday
		// TO DO: Remove with asset service live
		priceYesterday, err := db.GetLastPriceBefore(currentConstituent.Symbol, "MAIR120", "", endtime.AddDate(0, 0, -1))
		if err != nil {
			currentConstituent.PriceYesterday = float64(0)
		} else {
			currentConstituent.PriceYesterday = priceYesterday.Price
		}
		// Get price yesterweek
		priceYesterweek, err := db.GetLastPriceBefore(currentConstituent.Symbol, "MAIR120", "", endtime.AddDate(0, 0, -7))
		if err != nil {
			currentConstituent.PriceYesterweek = float64(0)
		} else {
			currentConstituent.PriceYesterweek = priceYesterweek.Price
		}

		retval = append(retval, currentConstituent)

	}
	return retval, nil
}

func (db *DB) SetCryptoIndexConstituent(constituent *CryptoIndexConstituent, indexSymbol string) error {
	fields := map[string]interface{}{
		"percentage":        constituent.Percentage,
		"price":             constituent.Price,
		"circulatingsupply": constituent.CirculatingSupply,
		"weight":            constituent.Weight,
		"cappingfactor":     constituent.CappingFactor,
		"numbasetokens":     constituent.NumBaseTokens,
	}
	tags := map[string]string{
		"name":    constituent.Name,
		"symbol":  constituent.Symbol,
		"address": constituent.Address,
		"cryptoindex": indexSymbol,
	}
	pt, err := clientInfluxdb.NewPoint(influxDbCryptoIndexConstituentsTable, tags, fields, time.Now())
	if err != nil {
		log.Error("Adding Crypto Index Constituent point to Influx batch: ", err)
		return err
	} else {
		db.addPoint(pt)
	}

	err = db.WriteBatchInflux()
	if err != nil {
		log.Error("Writing Crypto Index Constituent to Influx: ", err)
	}
	return err
}

// WIP: Returns the amounts of constituents tokens needed to mint an index token
// For now we hard-code amounts. TO DO: Set and Get data to and from influx/config
func (db *DB) GetCryptoIndexMintAmounts(symbol string) ([]CryptoIndexMintAmount, error) {

	constituents := []string{"SUSHI", "REN", "KP3R", "UTK", "AXS", "Yf-DAI", "DIA", "STAKE", "POLS", "PICKLE", "EASY", "IDLE", "SPICE"}
	amounts := []uint64{102504643110709000, 907990711110561000, 206329281567188, 461546152853883000, 56696968122059100, 4185582958247, 26215696618443200, 3778532359289460, 38656197930994700, 972363917807713, 2038967220923070, 952603382004964, 16697065735724400}
	addresses := []string{
		"0x6b3595068778dd592e39a122f4f5a5cf09c90fe2",
		"0x408e41876cccdc0f92210600ef50372656052a38",
		"0x1ceb5cb57c4d4e2b2433641b95dd330a33185a44",
		"0xdc9Ac3C20D1ed0B540dF9b1feDC10039Df13F99c",
		"0xF5D669627376EBd411E34b98F19C868c8ABA5ADA",
		"0xf4CD3d3Fda8d7Fd6C5a500203e38640A70Bf9577",
		"0x84cA8bc7997272c7CfB4D0Cd3D55cd942B3c9419",
		"0x0Ae055097C6d159879521C384F1D2123D1f195e6",
		"0x83e6f1E41cdd28eAcEB20Cb649155049Fac3D5Aa",
		"0x429881672B9AE42b8EbA0E26cD9C73711b891Ca5",
		"0x913D8ADf7CE6986a8CbFee5A54725D9Eea4F0729",
		"0x875773784Af8135eA0ef43b5a374AaD105c5D39e",
		"0x1fdab294eda5112b7d066ed8f2e4e562d5bcc664"}
	var mintAmounts []CryptoIndexMintAmount
	for i, constituent := range constituents {
		var mintAmount CryptoIndexMintAmount
		mintAmount.Symbol = constituent
		mintAmount.Address = addresses[i]
		mintAmount.Amount = amounts[i]
		mintAmount.RebalanceTime = time.Unix(0, 1608403933000000000)
		mintAmounts = append(mintAmounts, mintAmount)
	}
	return mintAmounts, nil
}
