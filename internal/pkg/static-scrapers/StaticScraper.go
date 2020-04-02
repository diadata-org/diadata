package staticscrapers

import (
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

// LoadHistoricRate downloads xml files for historic rate data
func LoadHistoricRate(rateType string) error {

	switch rateType {
	case "ESTER":
		err := GetHistoricESTER()
		if err != nil {
			log.Errorln("Error on download of ESTER data: ", err)
			return err
		}
	case "PRE-ESTER":
		err := GetHistoricPreESTER()
		if err != nil {
			log.Errorln("Error on download of Pre-ESTER data: ", err)
			return err
		}
	}
	return nil

}

// WriteHistoricRate writes the historic rate data into the redis database.
func WriteHistoricRate(ds models.Datastore, rateType string) error {

	switch rateType {
	case "ESTER":
		err := WriteHistoricESTER(ds)
		if err != nil {
			log.Errorln("Error on writing historic ESTER data: ", err)
			return err
		}
	case "SOFR":
		err := WriteHistoricSOFR(ds)
		if err != nil {
			log.Errorln("Error on writing historic SOFR data: ", err)
			return err
		}
	case "SOFR-AVG":
		err := WriteHistoricSOFRAvg(ds)
		if err != nil {
			log.Errorln("Error on writing historic SOFR average data: ", err)
			return err
		}
	case "PRE-ESTER":
		err := WriteHistoricPreESTER(ds)
		if err != nil {
			log.Errorln("Error on writing historic Pre-ESTER data: ", err)
			return err
		}
	}
	return nil

}
