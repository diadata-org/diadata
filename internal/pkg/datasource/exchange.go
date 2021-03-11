package datasource

import (
	"encoding/json"
	"github.com/diadata-org/diadata/pkg/dia"
	"go/build"
	"io/ioutil"
	"os"
)

type Source struct {
	Exchanges []dia.Exchange `json:"exchanges"`
}

//InitSource Read exchange.json file and put exchange metadata in Source struct
func InitSource() (source *Source, err error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	var (
		jsonFile  *os.File
		fileBytes []byte
	)
	jsonFile, err = os.Open(gopath + "/src/github.com/diadata-org/diadata/internal/pkg/datasource/exchange.json")
	if err != nil {
		return
	}
	defer jsonFile.Close()

	fileBytes, err = ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}
	json.Unmarshal(fileBytes, &source)
	return
}

//GetExchanges Get map of exchane name and exchange metadata
func (s *Source) GetExchanges() (exchanges map[string]dia.Exchange) {
	exchanges = make(map[string]dia.Exchange)
	for _, exchange := range s.Exchanges {
		exchanges[exchange.Name] = exchange
	}
	return
}
