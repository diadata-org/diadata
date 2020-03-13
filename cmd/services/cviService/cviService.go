package main

import (
	"fmt"
	_ "github.com/diadata-org/diadata/pkg/dia"
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

	fmt.Println(tNear)
	fmt.Println(tNext)

	neartermFwdIndexLevel, err := filters.ForwardIndexLevel(optionMetaNear, 0.0054, tNear)
	if err != nil {
		log.Error(err)
	}
	nexttermFwdIndexLevel, err := filters.ForwardIndexLevel(optionMetaNext, 0.0054, tNext)
	if err != nil {
		log.Error(err)
	}
	fmt.Println(neartermFwdIndexLevel)
	fmt.Println(nexttermFwdIndexLevel)

	//vindNear, err := VarianceIndex(optionsMeta []dia.OptionMetaIndex, 0.0054, tNear, neartermFwdIndexLevel, k0 float64)
}
