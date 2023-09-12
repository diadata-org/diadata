package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
)

const (
	WindowVolume = 60 * 60 * 24
)

var (
	volumeKey = "VOL" + strconv.Itoa(dia.BlockSizeSeconds)
)

// GetVolumeInflux returns the volume of @asset on @exchange using the VOL120 filter in the given time-range.
// Both, @asset and @exchange may be empty.
// If @starttime,@endtime are empty, the last 24h are taken into account.
func (datastore *DB) GetVolumeInflux(asset dia.Asset, exchange string, starttime time.Time, endtime time.Time) (*float64, error) {

	if endtime.IsZero() {
		endtime = time.Now()
		starttime = endtime.AddDate(0, 0, -1)
	}

	var q string

	if asset == (dia.Asset{}) {
		queryString := `
		SELECT SUM(value) 
		FROM %s 
		WHERE exchange='%s' 
		AND filter='%s' 
		AND time > %d AND time<= %d
		`
		q = fmt.Sprintf(queryString, influxDbFiltersTable, exchange, volumeKey, starttime.UnixNano(), endtime.UnixNano())
	} else if exchange == "" {
		queryString := `
		SELECT SUM(value) 
		FROM %s 
		WHERE address='%s' AND blockchain='%s' 
		AND exchange=''
		AND filter='%s' 
		AND time > %d AND time<= %d
		`
		q = fmt.Sprintf(queryString, influxDbFiltersTable, asset.Address, asset.Blockchain, volumeKey, starttime.UnixNano(), endtime.UnixNano())
	} else {
		queryString := `
		SELECT SUM(value) 
		FROM %s 
		WHERE address='%s' AND blockchain='%s' 
		AND exchange='%s' 
		AND filter='%s' 
		AND time > %d AND time<= %d
		`
		q = fmt.Sprintf(queryString, influxDbFiltersTable, asset.Address, asset.Blockchain, exchange, volumeKey, starttime.UnixNano(), endtime.UnixNano())
	}

	var errorString string
	res, err := queryInfluxDB(datastore.influxClient, q)
	if err != nil {
		log.Errorln("GetVolumeInflux ", err)
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
		return nil, errors.New(errorString)

	} else {
		volume := float64(0)
		log.Warnf("no volume on %s in influx filter table", exchange)
		return &volume, nil
	}
}

// Get24HoursAssetVolume returns the 24h trading volume of @asset across exchanges.
func (datastore *DB) Get24HoursAssetVolume(asset dia.Asset) (*float64, error) {
	endtime := time.Now()
	return datastore.GetVolumeInflux(asset, "", endtime.AddDate(0, 0, -1), endtime)
}

// Get24HoursExchangeVolume returns 24h trade volume on @exchange using VOL120 filtered data from influx.
func (datastore *DB) Get24HoursExchangeVolume(exchange string) (*float64, error) {
	endtime := time.Now()
	return datastore.GetVolumeInflux(dia.Asset{}, exchange, endtime.AddDate(0, 0, -1), endtime)
}

// Returns aggregated volumes from filters measurement on all exchanges.
func (datastore *DB) GetVolumesAllExchanges(asset dia.Asset, starttime time.Time, endtime time.Time) (exchVolumes dia.ExchangeVolumesList, err error) {
	q := fmt.Sprintf(
		`SELECT SUM(value) 
		FROM %s 
		WHERE filter='VOL120' 
		AND address='%s' 
		AND blockchain='%s'
		AND exchange!='' 
		AND time>%d 
		AND time<=%d 
		GROUP BY exchange`,
		influxDbFiltersTable,
		asset.Address,
		asset.Blockchain,
		starttime.UnixNano(),
		endtime.UnixNano(),
	)

	res, err := queryInfluxDB(datastore.influxClient, q)
	if err != nil {
		log.Errorln("GetLastTrades", err)
		return
	}

	if len(res) > 0 && len(res[0].Series) > 0 {
		for i, group := range res[0].Series {
			var exchangevolume dia.ExchangeVolume
			if val, ok := group.Tags["exchange"]; ok {
				exchangevolume.Exchange = val
			}
			if len(group.Values) > 0 && len(group.Values[0]) > 1 {
				exchangevolume.Volume, err = group.Values[0][1].(json.Number).Float64()
				if err != nil {
					return
				}
			}
			exchVolumes.Volumes = append(exchVolumes.Volumes, exchangevolume)
			if i == 0 {
				exchVolumes.Timestamp, err = time.Parse(time.RFC3339, res[0].Series[0].Values[0][0].(string))
				if err != nil {
					return
				}
			}
		}

	} else {
		log.Error("Empty response GetVolumesAllExchanges")
	}
	return
}

func (datastore *DB) GetExchangePairVolumes(asset dia.Asset, starttime time.Time, endtime time.Time, threshold float64) (map[string][]dia.PairVolume, error) {
	volumeMap := make(map[string][]dia.PairVolume)

	query := fmt.Sprintf(
		`
		SELECT SUM(multiplication)
		FROM (
			SELECT ABS(estimatedUSDPrice*volume)
			AS multiplication
			FROM %s
			WHERE quotetokenaddress='%s'
			AND quotetokenblockchain='%s'
			AND basetokenaddress!=''
			AND time>%d
			AND time<=%d
			)
		GROUP BY "exchange","basetokenaddress","basetokenblockchain","pooladdress"
		`,
		influxDbTradesTable,
		asset.Address,
		asset.Blockchain,
		starttime.UnixNano(),
		endtime.UnixNano(),
	)

	res, err := queryInfluxDB(datastore.influxClient, query)
	if err != nil {
		return volumeMap, err
	}

	if len(res) > 0 && len(res[0].Series) > 0 {
		for _, row := range res[0].Series {
			if len(row.Values[0]) > 1 {
				var (
					pairvolume dia.PairVolume
					err        error
				)

				exchange := row.Tags["exchange"]
				pairvolume.Volume, err = row.Values[0][1].(json.Number).Float64()
				if err != nil {
					log.Warn("parse volume: ", err)
				}
				if !(pairvolume.Volume >= threshold) {
					continue
				}
				pairvolume.Pair = dia.Pair{
					QuoteToken: dia.Asset{Blockchain: asset.Blockchain, Address: asset.Address},
					BaseToken:  dia.Asset{Blockchain: row.Tags["basetokenblockchain"], Address: row.Tags["basetokenaddress"]},
				}
				pairvolume.PoolAddress = row.Tags["pooladdress"]
				volumeMap[exchange] = append(volumeMap[exchange], pairvolume)
			}
		}
	}
	return volumeMap, nil
}
