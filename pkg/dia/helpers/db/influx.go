package db

import (
	"github.com/diadata-org/diadata/pkg/utils"
	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
	log "github.com/sirupsen/logrus"
)

func GetInfluxClient() clientInfluxdb.Client {
	var influxClient clientInfluxdb.Client
	var err error

	address := utils.Getenv("INFLUXURL", "http://localhost:8086")
	username := utils.Getenv( "INFLUXUSER", "")
	password := utils.Getenv("INFLUXPASSWORD", "")

	influxClient, err = clientInfluxdb.NewHTTPClient(clientInfluxdb.HTTPConfig{
		Addr:     address,
		Username: username,
		Password: password,
	})

	if err != nil {
		log.Error("NewDataStore influxdb", err)
	}

	return influxClient
}
