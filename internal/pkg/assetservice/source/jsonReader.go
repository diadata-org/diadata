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
}

func NewJSONReader(path string, filename string) *jsonReader {
	var jr jsonReader
	var assetChannel = make(chan dia.Asset)
	jr.assetChannel = assetChannel
	jr.path = path
	jr.filename = filename

	go func() {
		jr.fetchAssets()
	}()
	return &jr
}

func (jr *jsonReader) Asset() chan dia.Asset {
	return jr.assetChannel
}

// fetchAssets fetches all assets from the json file and sends them into the assetChannel.
func (jr *jsonReader) fetchAssets() {
	data, err := readJSONFromConfig(jr.path + "/" + jr.filename)
	if err != nil {
		log.Error(err)
	}
	var assets Assets
	err = json.Unmarshal(data, &assets)
	if err != nil {
		log.Error(err)
	}
	for _, asset := range assets.Assets {
		log.Info("got asset: ", asset)
		jr.assetChannel <- asset
	}
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
