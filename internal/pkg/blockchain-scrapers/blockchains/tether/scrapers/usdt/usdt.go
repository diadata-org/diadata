package main

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	// ! nazariyv: after the refactor this is gone. there is a built-in function now
	//"github.com/chromedp/chromedp/runner" ! https://github.com/chromedp/chromedp/issues/318
	"github.com/diadata-org/diadata/pkg/dia"
)

const symbol = "USDT"

func firstLoad(supply *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(`https://wallet.tether.to/transparency`),
		chromedp.Sleep(10 * time.Second),
		chromedp.WaitVisible(`#transparency`, chromedp.ByID),
		chromedp.Text(`#transparency > div:nth-child(2) > table > tbody > tr:nth-child(10) > td`, supply, chromedp.NodeVisible, chromedp.ByID),
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

	// create chrome instance
	var opts = []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-first-run", true),
		chromedp.Flag("no-default-browser-check", true),
		chromedp.Flag("no-sandbox", true),
	}
	allocContext, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(allocContext, chromedp.WithLogf(noLog))
	defer cancel()
	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var data string
	err := chromedp.Run(ctx, firstLoad(&data))
	if err != nil {
		log.Fatal(err)
	}

	for {
		err = chromedp.Run(ctx, reload(&data))
		if err != nil {
			log.Fatal(err)
		}
		result, err := strconv.ParseFloat(strings.Replace(strings.Replace(data, ",", "", -1), "$", "", -1), 64)
		if err == nil {
			log.Printf("Symbol: %s ; circulatingSupply: %f\n", symbol, result)
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
