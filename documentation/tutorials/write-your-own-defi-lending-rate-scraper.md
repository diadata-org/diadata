# Write your own Defi Lending Rate Scraper

## Getting Started

To add a Defi Lending Rate Scraper you first have to create an entry for the defi protocol your rate is being scraped from, with this function call

```
SetDefiProtocol(protocol dia.DefiProtocol)
```

With dia.DefiProtocol being defined in pkg/dia/Messages.go:

```bash
type DefiProtocol struct {                                                         
  Name                 string                                                      
  Address              string                                                      
  UnderlyingBlockchain string                                                      
}
```

Once you added the metadata about the Protocol \(or made sure it already exists by querying it in the database\) you can proceed with scraping actual interest values.

### Scraping the lending rate value

For that, you need to get all data that is defined in this structure at pkg/dia/Messages.go

```bash
type DefiLendingRate type {                                                        
  Timestamp   time.Time                                                            
  LendingRate float64                                                              
  Asset       string                                                               
  Protocol    DefiProtocol                                                         
}
```

The timestamp should be the timestamp of observation \(i.e., now\), the LendingRate is the actual rate, the asset is the asset that is lent and the Protocol is the protocol object created above.

You can add the scraped value using this function in pkg/model/db.go:

```bash
SetDefiLendingRateInflux(rate *dia.DefiLendingRate) error
```

After that, you are done. To test your addition, retrieve your value using this function:

```bash
GetDefiLendingRateInflux(starttime time.Time, endtime time.Time, asset string, protocol string) ([]dia.DefiLendingRate, error) 
```

It will return a slice of scraped values between starttime and endtime for the asset on the selected protocol.

### Scraping additional protocol quantities

If possible, scrape data for the following structure as well \(structure at pkg/dia/Messages.go \).

```bash
type DefiProtocolState {
	TotalUSD  	 		float64
	TotalETH	   		float64
	TotalPerDay	    map[string]float64
	MostLockedAsset	string
}
```

Here, `TotalUSD` and `TotalETH` are the total amounts locked in the protocol. `TotalPerDay` is the total value locked in 24H in USD, Ether and Bitcoin. 



