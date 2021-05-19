package itinService

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	dia "github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type ItinTokens []dia.ItinToken

func GetItins(itinUrl string) (ItinTokens, error) {
	var tokens ItinTokens
	resp, err := http.Get(itinUrl)
	if err != nil {
		log.Error(err)
		return tokens, err
	}
	defer utils.CloseHTTPResp(resp)
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
