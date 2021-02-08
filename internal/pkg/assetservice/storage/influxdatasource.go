package assetstorage

import (
	"fmt"
	"github.com/diadata-org/diadata/pkg/dia"
	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	influxDbAsset = "assets"
)

type InfluxDataSource struct {
	client clientInfluxdb.Client
	batch  clientInfluxdb.BatchPoints
}

func NewInfluxDataSource() *InfluxDataSource {
	ci, err := clientInfluxdb.NewHTTPClient(clientInfluxdb.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: "",
		Password: "",
	})
	if err != nil {
		log.Error("NewDataStore influxdb", err)
	}

	_, err = queryInfluxDB(ci, fmt.Sprintf("CREATE DATABASE %s", influxDbAsset))
	if err != nil {
		log.Errorln("queryInfluxDB CREATE DATABASE", err)
	}

	bp, err := clientInfluxdb.NewBatchPoints(clientInfluxdb.BatchPointsConfig{
		Database:  "asset",
		Precision: "s",
	})
	if err != nil {
		log.Errorln("NewBatchPoints", err)
	}
	return &InfluxDataSource{client: ci, batch: bp}
}

func (ds *InfluxDataSource) Save(asset dia.Asset) {
	tags := map[string]string{"symbol": asset.Symbol, "name": asset.Name, "address": asset.Address}
	fields := map[string]interface{}{
		"asset": asset,
	}
	pt, err := clientInfluxdb.NewPoint(influxDbAsset, tags, fields, time.Now())
	if err != nil {
		log.Errorln("Save Asset:", err)
	} else {
		log.Println("Save Asset:", fields)

		ds.batch.AddPoint(pt)
	}
	ds.client.Write(ds.batch)

}

func (ds *InfluxDataSource) Get() *dia.Asset {
	return &dia.Asset{}

}

func queryInfluxDB(clnt clientInfluxdb.Client, cmd string) (res []clientInfluxdb.Result, err error) {
	q := clientInfluxdb.Query{
		Command:  cmd,
		Database: influxDbAsset,
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
