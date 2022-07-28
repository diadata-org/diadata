package db

import (
	"github.com/diadata-org/diadata/pkg/utils"
	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
}

// GetInfluxClient returns an influx client connecting through the
// URL given in the environment variable INFLUXURL.
// If INFLUXURL is not set, it connects to @url per default.
func GetInfluxClient(url string) clientInfluxdb.Client {
	var influxClient clientInfluxdb.Client
	var err error

	address := utils.Getenv("INFLUXURL", url)
	log.Info("INFLUXURL: ", address)
	username := utils.Getenv("INFLUXUSER", "")
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
