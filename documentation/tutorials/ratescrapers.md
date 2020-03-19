# Write your own rate scraper

These instructions concern writing scrapers for single units characterised by a (floating point) number. For scrapers describing the relation between pairs of units, i.e. exchange rates see the instructions in exchangescrapers.md.

## Instructions for the addition of a rate scraper

<<<<<<< HEAD
In order to add your own scraper for a new data source, you must adhere to our format. We use Go modules for our scrapers, so that each data provider is living as an independent module.

### Practical advice
Let's assume you want to scrape a data source that provides floating point information. Create a new file in `exchange-scrapers/` and call it `MySourceScraper.go`. The main difference between scrapers is the Update() method, where the actual scraping is done.

```go
func (s *MyScraper) Update() error {
=======
In order to add your own scraper for a new data source, you must adhere to our format. Create the package file `UpdateMYRATE.go` in the package `/internal/pkg/ratescrapers`. The central method is ` UpdateMYRATE()`. This method acts on a RateScraper struct which is defined in RateScraper.go in the ratescrapers package. For instance, for the the Euro Short-Term Rate (ESTER) issued by the ECB, `UpdateESTER.go` would look like

```go
func (s *RateScraper) UpdateESTER() error {
>>>>>>> master
  // scraper code here
}
```

<<<<<<< HEAD
### Overview of architecture


```go
// PairScraper receives trades for a single pc.Pair from a single exchange.
type PairScraper interface {
  io.Closer

  // Channel returns a channel that can be used to receive trades
  Channel() chan *models.InterestRate

  // Error returns an error when the channel Channel() is closed
  // and nil otherwise
  Error() error

  // Pair returns the pair this scraper is subscribed to
  Pair() dia.Pair
=======
The scraped data has to be written into a struct of type InterestRate from `pkg/model/types.go`

```go
type InterestRate struct {
	Symbol string
	Value  float64
	Time   time.Time
	Source string
>>>>>>> master
}
```
and sent to the channel chanInterestRate of s. In order to write a new scraper, it is not imperative to understand the architecture of the pathway from top to bottom, but it might be helpful. For a first impression you can have a look at the following [diagram](github.com/diadata-org/diadata/documentation/tutorials/rate_scraper_diagram_down.pdf).
