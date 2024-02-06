package datasource

import (
	"encoding/json"
	"go/build"
	"io/ioutil"
	"os"

	"github.com/diadata-org/diadata/pkg/dia"
)

type Source struct {
	Exchanges []dia.Exchange `json:"exchanges"`
}

// InitSource Read exchange.json file and put exchange metadata in Source struct
func InitSource() (source *Source, err error) {

	var (
		jsonFile  *os.File
		fileBytes []byte
	)
	executionMode := os.Getenv("EXEC_MODE")

	if executionMode == "production" {
		jsonFile, err = os.Open("/config/exchanges.json")
	} else {
		gopath := os.Getenv("GOPATH")
		if gopath == "" {
			jsonFile, err = os.Open(build.Default.GOPATH)
		} else {
			jsonFile, err = os.Open(os.Getenv("GOPATH") + "/src/github.com/diadata-org/diadata/config/exchanges.json")
		}
	}
	if err != nil {
		return
	}
	defer func() {
		cerr := jsonFile.Close()
		if err == nil {
			err = cerr
		}
	}()

	fileBytes, err = ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}
	err = json.Unmarshal(fileBytes, &source)
	if err != nil {
		return
	}
	return
}

// GetExchanges Get map of exchange name and exchange metadata
func (s *Source) GetExchanges() (exchanges map[string]dia.Exchange) {
	exchanges = make(map[string]dia.Exchange)
	for _, exchange := range s.Exchanges {
		exchanges[exchange.Name] = exchange
	}
	return
}
