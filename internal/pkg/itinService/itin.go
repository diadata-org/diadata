package itinService

import (
	"encoding/json"

	dia "github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type ItinTokens []dia.ItinToken

func GetItins(itinUrl string) (ItinTokens, error) {
	var tokens ItinTokens
	body, _, err := utils.GetRequest(itinUrl)
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
