package models

import (
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
		"synthassetsymbol":       t.Asset.Symbol,
		"underlyingassetsymbol":  t.AssetUnderlying.Symbol,
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
