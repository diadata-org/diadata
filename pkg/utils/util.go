package utils

import "encoding/json"

func StructToMap(data interface{}) (resp map[string]interface{}, err error) {

 	var b []byte
	b, err = json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &resp)
	return
}
