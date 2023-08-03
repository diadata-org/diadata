package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
)

// SaveSynthSupplyInflux stores a synth supply in influx. Flushed when more than maxPoints in batch.
// Wrapper around SaveSynthSupplyInfluxToTable.
func (datastore *DB) SaveSynthSupplyInflux(t *dia.SynthAssetSupply) error {
	return datastore.SaveSynthSupplyInfluxToTable(t, influxDbSynthSupplyTable)
}

// SaveSynthSupplyInfluxToTable stores a synth supply in influx into @table.
// Flushed when more than maxPoints in batch.
func (datastore *DB) SaveSynthSupplyInfluxToTable(t *dia.SynthAssetSupply, table string) error {

	// Create a point and add to batch
	tags := map[string]string{
		"synthassetsymbol":       EscapeReplacer.Replace(t.Asset.Symbol),
		"underlyingassetsymbol":  EscapeReplacer.Replace(t.AssetUnderlying.Symbol),
		"synthtokenaddress":      t.Asset.Address,
		"underlyingtokenaddress": t.AssetUnderlying.Address,
		"blockchain":             t.Asset.Blockchain,
		"protocol":               t.Protocol,
	}
	fields := map[string]interface{}{
		"supply":           t.Supply,
		"underlyinglocked": t.LockedUnderlying,
		"collateralRatio":  t.ColleteralRatio,
		"blocknumber":      int64(t.BlockNumber),
		"totaldebt":        int64(t.TotalDebt),
	}

	pt, err := clientInfluxdb.NewPoint(table, tags, fields, t.Time)
	if err != nil {
		log.Errorln("SaveSynthSupplyInfluxToTable:", err)
	} else {
		datastore.addPoint(pt)
	}

	return err
}

func (datastore *DB) GetSynthLastSupplyInflux(blockchain, protocol, address string) {

}

// GetSynthSupplyInflux Get distinct syntasset per protocol
func (datastore *DB) GetSynthAssets(blockchain, protocol string) (r []string, err error) {
	queryString := ` 
	SHOW TAG VALUES FROM %s 
	WITH KEY="synthtokenaddress" 
	WHERE blockchain='%s'
	and protocol='%s'`

	q := fmt.Sprintf(queryString, influxDbSynthSupplyTable, blockchain, protocol)

	log.Info("query: ", q)
	res, err := queryInfluxDB(datastore.influxClient, q)
	if err != nil {
		log.Errorln("GetSynthAssets", err)
		return r, err
	}

	if len(res) > 0 && len(res[0].Series) > 0 {
		for _, row := range res[0].Series[0].Values {
			value, o := row[1].(string)
			if !o {
				log.Errorln("error on parsing value", row)
			} else {
				r = append(r, value)
			}
		}
	} else {
		err = fmt.Errorf("Empty response for")
		log.Error(err)
		return r, err
	}
	return r, nil

}

// GetSynthSupplyInflux
func (datastore *DB) GetSynthSupplyInflux(blockchain, protocol, address string, limit int, starttime, endtime time.Time) ([]dia.SynthAssetSupply, error) {
	var r []dia.SynthAssetSupply

	queryString := ` 
	SELECT  
	time,blockchain, blocknumber, collateralRatio,
	protocol, supply,   synthassetsymbol, synthtokenaddress,
	totaldebt, underlyingassetsymbol, underlyinglocked, underlyingtokenaddress  
	FROM %s 
	WHERE blockchain='%s'
	`

	if protocol != "" {
		queryString = queryString + `AND protocol='` + protocol + `'`
	}
	if address != "" && address != "0x0000000000000000000000000000000000000000" {
		queryString = queryString + `AND underlyingtokenaddress='` + address + `'`
		queryString = queryString + ` OR synthtokenaddress='` + address + `'`

	}

	queryString = queryString + ` AND time > %d AND time<= %d  `

	queryString = queryString + " ORDER BY time DESC"

	if limit == 1 {
		queryString = queryString + " LIMIT 1"
	}
	queryString = queryString + " ;"
	q := fmt.Sprintf(queryString, influxDbSynthSupplyTable, blockchain, starttime.UnixNano(), endtime.UnixNano())

	log.Info("query: ", q)
	res, err := queryInfluxDB(datastore.influxClient, q)
	if err != nil {
		log.Errorln("GetSynthSupplyInflux", err)
		return r, err
	}

	if len(res) > 0 && len(res[0].Series) > 0 {
		for _, row := range res[0].Series[0].Values {
			t := parseSynthAsset(row)
			if t != nil {
				r = append(r, *t)
			}
		}
	} else {
		err = fmt.Errorf("Empty response for")
		log.Error(err)
		return r, err
	}
	return r, nil

}

func parseSynthAsset(row []interface{}) *dia.SynthAssetSupply {
	if len(row) > 10 {
		t, err := time.Parse(time.RFC3339, row[0].(string))
		if err == nil {

			blockchain, o := row[1].(string)
			if !o {
				log.Errorln("error on parsing row 1", row)
			}

			var blocknumber int64
			v, o := row[2].(json.Number)
			if o {
				blocknumber, _ = v.Int64()
			} else {
				log.Errorln("error on parsing row 2", row)
			}

			var collateralRatio float64
			v, o = row[3].(json.Number)
			if o {
				collateralRatio, _ = v.Float64()
			} else {
				log.Errorln("error on parsing row 3", row)
			}

			protocol, o := row[4].(string)
			if !o {
				log.Errorln("error on parsing row 4", row)
			}

			var supply float64
			v, o = row[5].(json.Number)
			if o {
				supply, _ = v.Float64()
			} else {
				log.Errorln("error on parsing row 5", row)
			}

			synthassetsymbol, o := row[6].(string)
			if !o {
				log.Errorln("error on parsing row 6", row)
			}

			synthtokenaddress, o := row[7].(string)
			if !o {
				log.Errorln("error on parsing row 7", row)
			}

			var totaldebt float64
			v, o = row[8].(json.Number)
			if o {
				totaldebt, _ = v.Float64()
			} else {
				log.Errorln("error on parsing row 8", row)
			}

			underlyingassetsymbol, o := row[9].(string)
			if !o {
				log.Errorln("error on parsing row 9", row)
			}

			var underlyinglocked float64
			v, o = row[10].(json.Number)
			if o {
				underlyinglocked, _ = v.Float64()
			} else {
				log.Errorln("error on parsing row 10", row)
			}

			underlyingtokenaddress, o := row[11].(string)
			if !o {
				log.Errorln("error on parsing row 10", row)
			}

			synthassetsupply := dia.SynthAssetSupply{
				BlockNumber:      uint64(blocknumber),
				ColleteralRatio:  collateralRatio,
				Time:             t,
				Protocol:         protocol,
				Supply:           supply,
				Asset:            dia.Asset{Symbol: synthassetsymbol, Blockchain: blockchain, Address: synthtokenaddress},
				AssetUnderlying:  dia.Asset{Symbol: underlyingassetsymbol, Blockchain: blockchain, Address: underlyingtokenaddress},
				LockedUnderlying: underlyinglocked,
				TotalDebt:        totaldebt,
			}

			return &synthassetsupply
		}

	}
	return nil
}
