package source

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/configCollectors"
)

type Assets struct {
	Assets []dia.Asset `json:"Assets"`
}

type jsonReader struct {
	path         string
	filename     string
	assetChannel chan dia.Asset
	closed       chan bool
}

func NewJSONReader(path string, filename string) *jsonReader {
	var jr jsonReader
	jr.path = path
	jr.filename = filename
	jr.assetChannel = make(chan dia.Asset)
	jr.closed = make(chan bool)

	go func() {
		jr.fetchAssets()
	}()
	return &jr
}

func (jr *jsonReader) Asset() chan dia.Asset {
	return jr.assetChannel
}

func (jr *jsonReader) Close() chan bool {
	return jr.closed
}

// fetchAssets fetches all assets from the json file and sends them into the assetChannel.
func (jr *jsonReader) fetchAssets() {
	data, err := readJSONFromConfig(jr.path + "/" + jr.filename)
	if err != nil {
		log.Error(err)
	}
	var assets Assets
	json.Unmarshal(data, &assets)
	for _, asset := range assets.Assets {
		log.Info("got asset: ", asset)
		jr.assetChannel <- asset
	}

	// Gracefully close the asset channel after iterating through all assets.
	jr.closed <- true
}

// readJSONFromConfig reads a json file from the config folder and returns the slice of items.
func readJSONFromConfig(filename string) (content []byte, err error) {
	var (
		jsonFile *os.File
	)
	path := configCollectors.ConfigFileConnectors(filename, ".json")
	jsonFile, err = os.Open(path)
	if err != nil {
		return
	}
	defer jsonFile.Close()
	content, err = ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}
	return
}
