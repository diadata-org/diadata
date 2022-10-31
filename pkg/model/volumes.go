package models

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/jackc/pgx/v4"
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
// Will deprecate GetAggVolumesByExchange below.
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

// SetAggregatedVolume sets the aggregated volume @aggVol in postgres.
func (rdb *RelDB) SetAggregatedVolume(aggVol dia.AggregatedVolume) error {
	quotetokenQuery := fmt.Sprintf("(SELECT asset_id FROM %s WHERE blockchain=$1 and address=$2)", assetTable)
	basetokenQuery := fmt.Sprintf("(SELECT asset_id FROM %s WHERE blockchain=$3 and address=$4)", assetTable)
	query := fmt.Sprintf("INSERT INTO %s (quotetoken_id,basetoken_id,volume,exchange,time_range_seconds,compute_time) VALUES(%s,%s,$5,$6,$7,$8);", aggregatedVolumeTable, quotetokenQuery, basetokenQuery)

	_, err := rdb.postgresClient.Exec(context.Background(), query,
		aggVol.Pair.QuoteToken.Blockchain,
		aggVol.Pair.QuoteToken.Address,
		aggVol.Pair.BaseToken.Blockchain,
		aggVol.Pair.BaseToken.Address,
		aggVol.Volume,
		aggVol.Exchange,
		aggVol.TimeRangeSeconds,
		aggVol.Timestamp,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetAggregatedVolume returns all aggregated volumes in the time-range @starttime - @endtime.
func (rdb *RelDB) GetAggregatedVolumes(asset dia.Asset, starttime time.Time, endtime time.Time) (aggVolumes []dia.AggregatedVolume, err error) {
	valuesQuery := "a.volume,a.exchange,a.time_range_seconds,a.compute_time"
	pairQuery := "b.address,b.blockchain,b.name,b.symbol,b.decimals,c.address,c.blockchain,c.name,c.symbol,c.decimals"
	query := fmt.Sprintf(`
	SELECT %s,%s 
	FROM %s a 
	INNER JOIN %s b 
	ON a.quotetoken_id=b.asset_id 
	INNER JOIN %s c
	ON a.basetoken_id=c.asset_id 
	WHERE b.address=$1 
	AND b.blockchain=$2 
	AND compute_time>$3 
	AND compute_time<=$4 
	ORDER BY compute_time DESC`,
		valuesQuery,
		pairQuery,
		aggregatedVolumeTable,
		assetTable,
		assetTable,
	)

	var rows pgx.Rows
	rows, err = rdb.postgresClient.Query(context.Background(), query, asset.Address, asset.Blockchain, starttime, endtime)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var timestamp sql.NullTime
		var quotetokenSymbol sql.NullString
		var basetokenSymbol sql.NullString
		var quotetokenName sql.NullString
		var basetokenName sql.NullString
		var decimalsQuotetokenString string
		var decimalsBasetokenString string
		var decimalsQuotetoken int
		var decimalsBasetoken int

		var aggVolume dia.AggregatedVolume
		err = rows.Scan(
			&aggVolume.Volume,
			&aggVolume.Exchange,
			&aggVolume.TimeRangeSeconds,
			&timestamp,
			&aggVolume.Pair.QuoteToken.Address,
			&aggVolume.Pair.QuoteToken.Blockchain,
			&quotetokenName,
			&quotetokenSymbol,
			&decimalsQuotetokenString,
			&aggVolume.Pair.BaseToken.Address,
			&aggVolume.Pair.BaseToken.Blockchain,
			&basetokenName,
			&basetokenSymbol,
			&decimalsBasetokenString,
		)
		if err != nil {
			return
		}

		if timestamp.Valid {
			aggVolume.Timestamp = timestamp.Time
		}
		if quotetokenSymbol.Valid {
			aggVolume.Pair.QuoteToken.Symbol = quotetokenSymbol.String
		}
		if basetokenSymbol.Valid {
			aggVolume.Pair.BaseToken.Symbol = basetokenSymbol.String
		}
		if quotetokenName.Valid {
			aggVolume.Pair.QuoteToken.Name = quotetokenName.String
		}
		if basetokenName.Valid {
			aggVolume.Pair.BaseToken.Name = basetokenName.String
		}

		decimalsQuotetoken, err = strconv.Atoi(decimalsQuotetokenString)
		if err != nil {
			return
		}
		aggVolume.Pair.QuoteToken.Decimals = uint8(decimalsQuotetoken)
		decimalsBasetoken, err = strconv.Atoi(decimalsBasetokenString)
		if err != nil {
			return
		}
		aggVolume.Pair.BaseToken.Decimals = uint8(decimalsBasetoken)
		aggVolumes = append(aggVolumes, aggVolume)
	}
	return
}

// GetAggVolumesByExchange returns all aggregated volumes in the time-range @starttime - @endtime, sorted by exchanges.
func (rdb *RelDB) GetAggVolumesByExchange(asset dia.Asset, starttime time.Time, endtime time.Time) (exchVolumes []dia.ExchangeVolumesList, err error) {
	valuesQuery := "SUM(volume),exchange,compute_time"
	query := fmt.Sprintf("SELECT %s FROM %s WHERE quotetoken_id=(SELECT asset_id FROM asset WHERE address=$1 AND blockchain=$2) AND compute_time>$3 AND compute_time<=$4 GROUP BY exchange,compute_time ORDER BY compute_time DESC",
		valuesQuery,
		aggregatedVolumeTable,
	)

	var rows pgx.Rows
	rows, err = rdb.postgresClient.Query(context.Background(), query, asset.Address, asset.Blockchain, starttime, endtime)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var volume float64
		var exchange string
		var timestamp sql.NullTime
		var t time.Time

		var exchVolume dia.ExchangeVolumesList
		err = rows.Scan(
			&volume,
			&exchange,
			&timestamp,
		)
		if err != nil {
			return
		}

		if timestamp.Valid {
			t = timestamp.Time
		}

		// New timestamp, hence new entry in exchVolumes slice.
		if len(exchVolumes) == 0 || t.Before(exchVolumes[len(exchVolumes)-1].Timestamp) {
			exchVolume.Volumes = []dia.ExchangeVolume{{Volume: volume, Exchange: exchange}}
			exchVolume.Timestamp = t
			exchVolumes = append(exchVolumes, exchVolume)
			continue
		}

		// Existing timestamp, append to the inner exchangeVolume slice.
		if !t.Before(exchVolumes[len(exchVolumes)-1].Timestamp) {
			exchVolumes[len(exchVolumes)-1].Volumes = append(exchVolumes[len(exchVolumes)-1].Volumes, dia.ExchangeVolume{Volume: volume, Exchange: exchange})
		}

	}
	return
}

// GetAggVolumesByPair returns all aggregated volumes in the time-range @starttime - @endtime, sorted by pairs.
// Asset should be given fully.
func (rdb *RelDB) GetAggVolumesByPair(asset dia.Asset, starttime time.Time, endtime time.Time) (allPairVolumes []dia.PairVolumesList, err error) {
	valuesQuery := "compute_time,a.address,a.blockchain,a.name,a.symbol,a.decimals"

	query := fmt.Sprintf("SELECT SUM(volume),%s FROM %s INNER JOIN asset a ON basetoken_id=a.asset_id WHERE quotetoken_id=(SELECT asset_id FROM asset WHERE address=$1 AND blockchain=$2) AND compute_time>$3 AND compute_time<=$4 GROUP BY %s ORDER BY compute_time DESC",
		valuesQuery,
		aggregatedVolumeTable,
		valuesQuery,
	)

	var rows pgx.Rows
	rows, err = rdb.postgresClient.Query(context.Background(), query, asset.Address, asset.Blockchain, starttime, endtime)
	if err != nil {
		return
	}
	defer rows.Close()

	var volume float64
	var timestamp sql.NullTime
	var address string
	var blockchain string
	var currTime time.Time
	var symbol sql.NullString
	var name sql.NullString
	var decimalsString string
	var decimals int
	// var count int

	for rows.Next() {

		err = rows.Scan(
			&volume,
			&timestamp,
			&address,
			&blockchain,
			&name,
			&symbol,
			&decimalsString,
		)
		if err != nil {
			return
		}

		if timestamp.Valid {
			currTime = timestamp.Time
			// exchangeVolume.Timestamp = timestamp.Time
		}
		var pair dia.Pair
		pair.QuoteToken = asset
		pair.BaseToken.Address = address
		pair.BaseToken.Blockchain = blockchain
		if symbol.Valid {
			pair.BaseToken.Symbol = symbol.String
		}
		if name.Valid {
			pair.BaseToken.Name = name.String
		}
		decimals, err = strconv.Atoi(decimalsString)
		if err != nil {
			return
		}
		pair.BaseToken.Decimals = uint8(decimals)

		var pairVolume dia.PairVolume
		// new timestamp, hence new entry in allPairVolumes slice.
		if len(allPairVolumes) == 0 || currTime.Before(allPairVolumes[len(allPairVolumes)-1].Timestamp) {
			pairVolume.Pair = pair
			pairVolume.Volume = volume
			allPairVolumes = append(allPairVolumes, dia.PairVolumesList{Volumes: []dia.PairVolume{pairVolume}, Timestamp: currTime})
			continue
		}

		// Existing timestamp, append to the inner PairVolume slice.
		if !currTime.Before(allPairVolumes[len(allPairVolumes)-1].Timestamp) {
			// If @pair already exists for this timestamp, add volume...
			var pairCheck bool
			for i, vol := range allPairVolumes[len(allPairVolumes)-1].Volumes {
				if pair == vol.Pair {
					allPairVolumes[len(allPairVolumes)-1].Volumes[i].Volume += volume
					pairCheck = true
				}
			}
			// ...otherwise add new entry for @pair.
			if !pairCheck {
				allPairVolumes[len(allPairVolumes)-1].Volumes = append(allPairVolumes[len(allPairVolumes)-1].Volumes, dia.PairVolume{Pair: pair, Volume: volume})
			}
		}

	}
	return
}

// SetTradesDistribution sets the trades distribution parameter for @asset in postgres.
func (rdb *RelDB) SetTradesDistribution(tradesDist dia.TradesDistribution) error {
	assetQuery := fmt.Sprintf("(SELECT asset_id FROM %s WHERE blockchain=$1 and address=$2)", assetTable)
	query := fmt.Sprintf("INSERT INTO %s (asset_id,num_trades_total,num_low_bins,threshold,size_bin_seconds,avg_num_per_bin,std_deviation,time_range_seconds,compute_time) VALUES(%s,$3,$4,$5,$6,$7,$8,$9,$10);", tradesDistributionTable, assetQuery)

	_, err := rdb.postgresClient.Exec(context.Background(), query,
		tradesDist.Asset.Blockchain,
		tradesDist.Asset.Address,
		tradesDist.NumTradesTotal,
		tradesDist.NumLowBins,
		tradesDist.Threshold,
		tradesDist.SizeBinSeconds,
		tradesDist.AvgNumPerBin,
		tradesDist.StdDeviation,
		tradesDist.TimeRangeSeconds,
		tradesDist.Timestamp,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetTradesDistribution returns all trades distribution parameters in the time-range @starttime - @endtime.
func (rdb *RelDB) GetTradesDistribution(asset dia.Asset, starttime time.Time, endtime time.Time) (tradesDistributions []dia.TradesDistribution, err error) {
	valuesQuery := "a.num_trades_total,a.num_low_bins,a.threshold,a.size_bin_seconds,a.avg_num_per_bin,a.std_deviation,a.time_range_seconds,a.compute_time"
	pairQuery := "b.address,b.blockchain,b.name,b.symbol,b.decimals"
	query := fmt.Sprintf("SELECT %s,%s FROM %s a INNER JOIN %s b ON a.asset_id=b.asset_id WHERE b.address=$1 AND b.blockchain=$2 AND compute_time>$3 and compute_time<=$4 oRDER BY compute_time DESC",
		valuesQuery,
		pairQuery,
		tradesDistributionTable,
		assetTable,
	)

	var rows pgx.Rows
	rows, err = rdb.postgresClient.Query(context.Background(), query, asset.Address, asset.Blockchain, starttime, endtime)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var timestamp sql.NullTime
		var symbol sql.NullString
		var name sql.NullString
		var decimalsString string
		var decimals int
		//a.time_range_seconds,a.compute_time
		var tradesDist dia.TradesDistribution
		err = rows.Scan(
			&tradesDist.NumTradesTotal,
			&tradesDist.NumLowBins,
			&tradesDist.Threshold,
			&tradesDist.SizeBinSeconds,
			&tradesDist.AvgNumPerBin,
			&tradesDist.StdDeviation,
			&tradesDist.TimeRangeSeconds,
			&timestamp,
			&tradesDist.Asset.Address,
			&tradesDist.Asset.Blockchain,
			&name,
			&symbol,
			&decimalsString,
		)
		if err != nil {
			return
		}

		if timestamp.Valid {
			tradesDist.Timestamp = timestamp.Time
		}
		if symbol.Valid {
			tradesDist.Asset.Symbol = symbol.String
		}
		if name.Valid {
			tradesDist.Asset.Name = name.String
		}
		decimals, err = strconv.Atoi(decimalsString)
		if err != nil {
			return
		}
		tradesDist.Asset.Decimals = uint8(decimals)

		tradesDistributions = append(tradesDistributions, tradesDist)
	}
	return
}
