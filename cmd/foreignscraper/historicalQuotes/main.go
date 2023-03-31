package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("TO DO")

	// 1. Get timestamp of last recorded quotation.
	// 2. Fetch hourly ETH quotation(s) from last timestamp until now:
	//		Try to fetch from our DB -> GetAssetQuotation(asset dia.Asset, timestamp time.Time)
	// 		If not available in our DB, fetch from CMC API -> https://coinmarketcap.com/api/documentation/v1/#operation/getV2CryptocurrencyQuotesHistorical
	// 3. Store data in postgres table historicalquotes (also meaning first write a setter in pkg/model/quotation.go).
	// Repeat 1. to 3. every hour.
}
