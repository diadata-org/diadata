package main

import (
	"fmt"
	filters "github.com/diadata-org/diadata/internal/pkg/filtersOptionService"
	log "github.com/sirupsen/logrus"
)

func main() {
	optionMetaNext, err := filters.GetNextTermOptionMeta("BTC")
	if err != nil {
		log.Error(err)
	}
	optionMetaNear, err := filters.GetNearTermOptionMeta("BTC", optionMetaNext[0].ExpirationTime)
	if err != nil {
		log.Error(err)
	}

	miyNear, err := filters.MinutesInYear(optionMetaNear[0].ExpirationTime.Year())
	if err != nil {
		log.Error(err)
	}
	miyNext, err := filters.MinutesInYear(optionMetaNext[0].ExpirationTime.Year())
	if err != nil {
		log.Error(err)
	}

	tNear := filters.TimeToMaturity(optionMetaNear[0]) / miyNear
	tNext := filters.TimeToMaturity(optionMetaNext[0]) / miyNext

	fmt.Println(optionMetaNear)
	fmt.Println(optionMetaNext)

	//fmt.Println(tNear)
	//fmt.Println(tNext)

	neartermFwdIndexLevel, err := filters.ForwardIndexLevel(optionMetaNear, 0.0054, tNear)
	if err != nil {
		log.Error(err)
	}
	nexttermFwdIndexLevel, err := filters.ForwardIndexLevel(optionMetaNext, 0.0054, tNext)
	if err != nil {
		log.Error(err)
	}
	//fmt.Println(neartermFwdIndexLevel)
	//fmt.Println(nexttermFwdIndexLevel)

	if optionMetaNear[0].GeneralizedInstrumentName != "" &&
	   optionMetaNext[0].GeneralizedInstrumentName != "" {
		omINear, err := filters.GetOptionMetaIndex("BTC", optionMetaNear[0].GeneralizedInstrumentName[4:11])
		if err != nil {
			log.Error(err)
		}

		fmt.Println("-----")
		fmt.Println(omINear)
		fmt.Println("-----")

		vindNear, err := filters.VarianceIndex(omINear, 0.0054, tNear, neartermFwdIndexLevel, 13000.0)
		if err != nil {
			log.Error(err)
		}
		fmt.Println("vindNear: ", vindNear)

		omINext, err := filters.GetOptionMetaIndex("BTC", optionMetaNext[0].GeneralizedInstrumentName[4:11])
		if err != nil {
			log.Error(err)
		}

		vindNext, err := filters.VarianceIndex(omINext, 0.0054, tNext, nexttermFwdIndexLevel, 13000.0)
		if err != nil {
			log.Error(err)
		}
		fmt.Println(vindNext)

		CVI, err := filters.CVI(vindNear, vindNext, tNear, tNext, 2020)
		if err != nil {
			log.Error(err)
		}
		fmt.Println(CVI)
		err = filters.CVIToDatastore(CVI)
		if err != nil {
			log.Error(err)
		}
	}

}
