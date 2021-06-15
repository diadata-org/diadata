package ratescrapers

import (
	"errors"
	"github.com/anaskhan96/soup"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"strconv"
	"strings"
	"time"
)

var logger = logrus.New()

type SONIA struct {
	RefDate time.Time
	RefRate float64
}

func (s *SONIA) reset() {
	s.RefRate = 0
	s.RefDate = time.Time{}
}

type SONIAHistorical []SONIA

func (s *RateScraper) UpdateSonia() (err error) {
	err = nil
	logger.Formatter = new(prefixed.TextFormatter)
	logger.Level = logrus.InfoLevel

	// ! if the data format changes, uncomment this to see where the problem is
	//soup.SetDebug(true)

	resp, err := soup.Get("https://www.bankofengland.co.uk/boeapps/database/fromshowcolumns.asp?Travel=NIxSUx&FromSeries=1&ToSeries=50&DAT=ALL&FNY=&CSVF=TT&html.x=132&html.y=32&C=5JK&Filter=N")
	if err != nil {
		return
	}
	doc := soup.HTMLParse(resp)
	tds := doc.Find("table", "id", "stats-table").Find("tbody").FindAll("td")

	nilDt := time.Time{}
	var historicalRates SONIAHistorical
	var sonia = SONIA{}
	var refDt time.Time
	var refRt float64

	for ix, refDtOrRt := range tds {
		switch ix % 2 {
		case 0:
			refDt, err = time.Parse("2 Jan 06", strings.TrimSpace(refDtOrRt.Text()))
			if err != nil {
				logger.WithFields(logrus.Fields{"prefix": "SONIA"}).Error(err)
				return
			}
			if sonia.RefDate != nilDt || sonia.RefRate != 0 {
				err = errors.New("sonia should be empty but is not. looks like the data format changed")
				logger.WithFields(logrus.Fields{"prefix": "SONIA"}).Error(err)
				return
			}
			sonia.RefDate = refDt
		case 1:
			refRt, err = strconv.ParseFloat(strings.TrimSpace(refDtOrRt.Text()), 64)
			if err != nil {
				logger.WithFields(logrus.Fields{"prefix": "SONIA"}).Error(err)
				return
			}
			if sonia.RefDate == nilDt || sonia.RefRate != 0 {
				err = errors.New("sonia should should have empty date and zero rate. looks like the data format changed")
				logger.WithFields(logrus.Fields{"prefix": "SONIA"}).Error(err)
				return
			}
			sonia.RefRate = refRt
			historicalRates = append(historicalRates, sonia)
			sonia.reset()
		}
	}

	logger.WithFields(logrus.Fields{"prefix": "SONIA"}).Debug(historicalRates)

	for _, sonia := range historicalRates {
		t := &models.InterestRate{
			Symbol:          "SONIA",
			Value:           sonia.RefRate,
			PublicationTime: time.Now(),
			EffectiveDate:   sonia.RefDate,
			Source:          "BOE",
		}

		// Send new data through channel chanInterestRate
		logger.WithFields(logrus.Fields{"prefix": "SONIA"}).Debugf("writing interestRate %#v in %v\n", t, s.chanInterestRate)
		s.chanInterestRate <- t
	}

	logger.WithFields(logrus.Fields{"prefix": "SONIA"}).Info("update complete")

	return
}
