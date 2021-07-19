package filters

import (
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"math"
	"time"
)


//Retursn Option b/w 23 to 37 days of expirationss
func GetOptionComponents(baseCurrency string) ([]dia.OptionMeta, error) {
	nextTerm := []dia.OptionMeta{}
	ds, err := models.NewDataStore()
	if err != nil {
		return nextTerm, err
	}

	optionsMeta, err := ds.GetOptionMeta(baseCurrency)

	log.Errorln("optionsMeta",optionsMeta)
	if err != nil {
		return nil, err
	}

	// Determine date of NextTerm
	for _, optionMeta := range optionsMeta {
		//more than 23 days and less than 37 days to expiration
 		if optionMeta.ExpirationTime.Sub(time.Now()) < 37*24*time.Hour && optionMeta.ExpirationTime.Sub(time.Now()) > 8*24*time.Hour {
			nextTerm = append(nextTerm, optionMeta)
		}

	}

	return nextTerm, nil
}

//Return Option expiring in 25 days
func GetNear(optionsMeta []dia.OptionMeta) ([]dia.OptionMeta, error) {
	options := []dia.OptionMeta{}


	// Determine date of NextTerm
	for _, optionMeta := range optionsMeta {
		//more than 23 days and less than 37 days to expiration
		if optionMeta.ExpirationTime.Sub(time.Now()) < 30*24*time.Hour  {
			options = append(options, optionMeta)
		}

	}

	return options, nil
}

// Expiring in greater than 10 days
func GetNext(optionsMeta []dia.OptionMeta) ([]dia.OptionMeta, error) {
	options := []dia.OptionMeta{}


	// Determine date of NextTerm
	for _, optionMeta := range optionsMeta {
		//more than 23 days and less than 37 days to expiration
		if optionMeta.ExpirationTime.Sub(time.Now()) > 14*24*time.Hour  && optionMeta.ExpirationTime.Sub(time.Now()) < 25*24*time.Hour{
			options = append(options, optionMeta)
		}

	}

	return options, nil
}

func CalculateT(settlementMinutes float64,tmod float64) float64 {
	//  T = (Minutes Left in Current Day + Minutes in Settlement Day + Minutes in Other Days)/Minutes in Year
	tcd := remainingMinutesForToday()
 	f := tcd + settlementMinutes + tmod
	return f/525600
}

func remainingMinutesForToday() float64 {
	currentTime := time.Now()
	todayEndTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	t := currentTime.Sub(todayEndTime)
	return t.Minutes()
}
func remainingMinutesForUpto(settlementTime time.Time) float64 {
	currentTime := time.Now()
	todayEndTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	t := currentTime.Sub(todayEndTime)
	return t.Minutes()
}

func CalculateForwardIndex(strikePrice,roi,time,callprice,putprice float64) float64{

 	return strikePrice + math.Exp(roi*time)*(callprice-putprice)



}



