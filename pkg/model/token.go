package models

const influxDbTokenTable = "tokendetail"

// SaveTokenDetailInflux stores a quotation which is not from DIA to an influx batch
// func (db *DB) SaveTokenDetailInflux(tk Token) error {
// 	tags := map[string]string{"source": tk.Source}
// 	fields := map[string]interface{}{
// 		"symbol":      tk.Symbol,
// 		"name":        tk.Name,
// 		"totalSupply": tk.TotalSupply,
// 		"decimals":    tk.Decimals,
// 	}
// 	pt, err := clientInfluxdb.NewPoint(influxDbTokenTable, tags, fields, tk.Time)
// 	if err != nil {
// 		log.Errorln("NewTokenInflux:", err)
// 	} else {
// 		db.addPoint(pt)
// 	}
// 	err = db.WriteBatchInflux()
// 	if err != nil {
// 		log.Errorln("Write influx batch: ", err)
// 	}

// 	return err
// }

// // GetTokenDetialInflux returns the last quotation of @symbol before @timestamp
// func (db *DB) GetTokenDetailInflux(symbol, source string, timestamp time.Time) (Token, error) {
// 	retval := Token{}

// 	unixtime := timestamp.UnixNano()
// 	q := fmt.Sprintf("SELECT * FROM %s WHERE source='%s' and symbol='%s' and time<%d order by time desc limit 1", influxDbTokenTable, source, symbol, unixtime)
// 	fmt.Println("query: ", q)
// 	res, err := queryInfluxDB(db.influxClient, q)
// 	if err != nil {
// 		fmt.Println("Error querying influx")
// 		return retval, err
// 	}

// 	fmt.Printf("Return data %#v", res)

// 	if len(res) > 0 && len(res[0].Series) > 0 {
// 		layout := "2006-01-02T15:04:05Z"
// 		vals := res[0].Series[0].Values[0]

// 		retval.Time, err = time.Parse(layout, vals[0].(string))
// 		if err != nil {
// 			log.Error(err)
// 		}
// 		val64, err := vals[1].(json.Number).Int64()
// 		if err != nil {
// 			log.Error(err)
// 		}
// 		retval.Decimals = int(val64)
// 		retval.Name = vals[2].(string)
// 		retval.Symbol = vals[4].(string)
// 		retval.Source = source
// 		retval.TotalSupply, err = vals[5].(json.Number).Float64()
// 		if err != nil {
// 			log.Error(err)
// 		}

// 		return retval, nil
// 	}
// 	return retval, err
// }

// // GetCurentTotalSupply returns the average price of @symbol on @source from yesterday
// func (db *DB) GetCurentTotalSupply(symbol, source string) (float64, error) {

// 	// Get time

// 	timestamp := time.Now()
// 	unixtime := timestamp.UnixNano()

// 	// Make corresponding influx query
// 	q := fmt.Sprintf("SELECT * FROM %s WHERE source='%s' and symbol='%s' and time<%d order by time desc limit 1", influxDbTokenTable, source, symbol, unixtime)
// 	res, err := queryInfluxDB(db.influxClient, q)
// 	if err != nil {
// 		fmt.Println("Error querying influx")
// 		return 0, err
// 	}

// 	if len(res) > 0 && len(res[0].Series) > 0 {
// 		vals := res[0].Series[0].Values[0]
// 		totalSupply, err := vals[5].(json.Number).Float64()
// 		if err != nil {
// 			log.Error(err)
// 			return 0, errors.New("No data available")
// 		}

// 		return totalSupply, nil

// 	}
// 	return 0, errors.New("No data available from yesterday")
// }
