package utils

import (
	"encoding/json"
	signer "github.com/diadata-org/diadata/pkg/sign"
)

func StructToMap(data interface{}) (resp map[string]interface{}, err error) {

	var b []byte
	b, err = json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &resp)
	return
}

func SignData(responseMap map[string]interface{}) (signedData string, err error) {
	var (
		qByte      []byte
		dataSigner *signer.Signer
	)
	qByte, err = json.Marshal(responseMap)
	if err != nil {
		return "", err
	}
	dataSigner, err = signer.New()
	if err != nil {
		return "", err
	}
	signedData, err = dataSigner.Sign(qByte)
	if err != nil {
		return "", err
	}
	return signedData, nil

}
