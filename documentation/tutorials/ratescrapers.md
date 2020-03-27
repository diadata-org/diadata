# Write your own rate scraper

These instructions concern writing scrapers for single units characterised by a (floating point) number. For scrapers describing the relation between pairs of units, i.e. exchange rates see the instructions in exchangescrapers.md.

## Instructions for the addition of a rate scraper

In order to add your own scraper for a new data source, you must adhere to our format. Create the package file `UpdateMYRATE.go` in the package `/internal/pkg/ratescrapers`. The central method is ` UpdateMYRATE()`. This method acts on a RateScraper struct which is defined in RateScraper.go in the ratescrapers package. For instance, for the the Euro Short-Term Rate (ESTER) issued by the ECB, `UpdateESTER.go` would look like

### Practical advice
Let's assume you want to scrape a data source that provides floating point information. Create a new file in `exchange-scrapers/` and call it `MySourceScraper.go`. The main difference between scrapers is the Update() method, where the actual scraping is done.

```go
func (s *RateScraper) UpdateESTER() error {
  // scraper code here
}
```

The scraped data has to be written into a struct of type InterestRate from `pkg/model/types.go`

```go
type InterestRate struct {
	Symbol string
	Value  float64
	Time   time.Time
	Source string
}
```
and sent to the channel chanInterestRate of s. In order to write a new scraper, it is not imperative to understand the architecture of the pathway from top to bottom, but it might be helpful. For a first impression you can have a look at the following [diagram](github.com/diadata-org/diadata/documentation/tutorials/rate_scraper_diagram_down.pdf).
