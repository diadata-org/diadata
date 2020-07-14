package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	scrapers "github.com/diadata-org/diadata/internal/pkg/exchange-scrapers"
	"github.com/diadata-org/diadata/pkg/dia"
	log "github.com/sirupsen/logrus"
)

func init() {
	// log.SetReportCaller(true)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

// ? Not working? Probably because option already expired. i.e. change OptionExpiry below
func scrapeOptions() (ftx *scrapers.FTXOptionsScraper, err error) {
	var scrapeFrequency int64 = 10
	ftx, err = scrapers.NewFTXOptionsScraper(scrapeFrequency)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req := scrapers.FTXRequest{
		Underlying: "BTC",
		Type:       dia.CallOption,
		Strike:     10000,
		// ! option expiry MUST be at 03:00 UTC whatever date of your picking in the future
		OptionExpiry: time.Date(2020, 6, 26, 3, 0, 0, 0, time.UTC),
		Size:         5,
	}
	go ftx.Scrape(req)
	return
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	ftx, _ := scrapeOptions()
	time.Sleep(time.Duration(time.Minute))
	ftx.Close()
	time.Sleep(time.Duration(time.Second * 5))
}
