package db

import (
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
	"time"
)

type IndexTimeStamp struct {
	Date time.Time
	Value string
}

func getInfluxConnection() *models.DB {
	database, err := models.NewInfluxDataStore()
	if err != nil {
		log.Error("error connecting to influx", err)
	}
	return database
}

func CreateIndexTimestampList(data [][]string) []IndexTimeStamp {
	var indexTimeStamps []IndexTimeStamp
	for i, line := range data {
		if i > 0 { // omit header line
			var rec IndexTimeStamp
			save := true
			for j, field := range line {
				if j == 0 {
					date, errDateConv := time.Parse("02/01/2006", field)
					if errDateConv != nil {
						log.Error("could not parse date ", field, errDateConv, " will not save data")
						save = false
						continue
					}
					rec.Date = date
				} else if j == 1 {
					rec.Value = field
				}
			}
			if save {
				indexTimeStamps = append(indexTimeStamps, rec)
			}
		}
	}
	return indexTimeStamps
}

func SaveIndexEngineIntoInflux(indexTimeStamp IndexTimeStamp, name string) {
	influxDb := getInfluxConnection()
	// Create a point and add to batch
	tags := map[string]string{
		"symbol" : name,
		//"basetokenblockchain":  t.BaseToken.Blockchain,
	}
	fields := map[string]interface{}{
		"name":  name,
		"value": indexTimeStamp.Value,
	}
	err := influxDb.SaveIndexEngineTimeInflux(tags, fields, indexTimeStamp.Date)
	if err != nil {
		log.Errorln("Error saving Index into influx:", err)
	}
}