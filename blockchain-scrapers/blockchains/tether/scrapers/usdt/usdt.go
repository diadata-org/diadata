package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/runner"
	"github.com/diadata-org/api-golang/pkg/dia"
)

const symbol = "USDT"

func firstLoad(supply *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(`https://wallet.tether.to/transparency`),
	}
}

func reload(supply *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Reload(),
		chromedp.Sleep(10 * time.Second),
		chromedp.WaitVisible(`#transparency`, chromedp.ByID),
		chromedp.Text(`#transparency > div:nth-child(2) > table > tbody > tr:nth-child(10) > td`, supply, chromedp.NodeVisible, chromedp.ByID),
	}
}

func noLog(string, ...interface{}) {

}

func main() {
	config := dia.GetConfigApi()
	if config == nil {
		panic("Couldnt load config")
	}
	client := dia.NewClient(config)
	if client == nil {
		panic("Couldnt load client")
	}
	prevResult := 0.0

	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	var options chromedp.Option
	options = chromedp.WithRunnerOptions(
		runner.Flag("headless", true),
		runner.Flag("disable-gpu", true),
		runner.Flag("no-first-run", true),
		runner.Flag("no-default-browser-check", true),
		runner.Flag("no-sandbox", true),
	)

	c, err := chromedp.New(ctxt, options, chromedp.WithLog(noLog))
	if err != nil {
		log.Fatal(err)
	}

	var data string
	err = c.Run(ctxt, firstLoad(&data))
	if err != nil {
		log.Fatal(err)
	}

	for {
		err = c.Run(ctxt, reload(&data))
		if err != nil {
			log.Fatal(err)
		}
		result, err := strconv.ParseFloat(strings.Replace(strings.Replace(data, ",", "", -1), "$", "", -1), 64)
		if err == nil {
			fmt.Printf("Symbol: %s ; circulatingSupply: %f\n", symbol, result)
			if prevResult != result {
				client.SendSupply(&dia.Supply{
					Symbol:            symbol,
					CirculatingSupply: result,
				})
				prevResult = result
			}
		} else {
			log.Println("Float conversion error:", err)
		}
	}
}
