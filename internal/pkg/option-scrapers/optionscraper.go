package optionscrapers

import "github.com/sirupsen/logrus"

var log = logrus.New()

type OptionScraper interface {
	Scrap()
	Fetchinstruments()
}