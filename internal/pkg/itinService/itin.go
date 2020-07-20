package itinService

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	log "github.com/sirupsen/logrus"
	dia "github.com/diadata-org/diadata/pkg/dia"
)

type ItinTokens []dia.ItinToken

func GetItins(itinUrl string) (ItinTokens, error) {
	var tokens ItinTokens
	resp, err := http.Get(itinUrl)
	if err != nil {
		log.Error(err)
		return tokens, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return tokens, err
	}

	err = json.Unmarshal(body, &tokens)
	if err != nil {
		log.Error(err)
		return tokens, err
	}
	return tokens, nil
}
