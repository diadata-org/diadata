package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/diadata-org/api-golang/dia"
)

const (
	HOST     = "dashd"
	PORT     = 9998
	SYMBOL   = "DASH"
	USERNAME = "dashrpc"
	PASSWORD = "dash_rpc_521d43b"
	PROTOCOL = "http"
)

type GetTxOutResponse struct {
	Result struct {
		Height          int     `json:"height"`
		Bestblock       string  `json:"bestblock"`
		Transactions    int     `json:"transactions"`
		Txouts          int     `json:"txouts"`
		HashSerialized2 string  `json:"hash_serialized_2"`
		DiskSize        int     `json:"disk_size"`
		TotalAmount     float64 `json:"total_amount"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    string      `json:"id"`
}

func main() {

	config := dia.GetConfigApi()
	if config == nil {
		panic("Couldnt load config")
	}
	client := dia.NewClient(config)
	if client == nil {
		panic("Couldnt load client")
	}

	for {
		info, err := Call("gettxoutsetinfo")
		if err != nil {
			log.Fatalln(err)
		}

		err = client.SendSupply(&dia.Supply{
			Symbol:            SYMBOL,
			CirculatingSupply: info.Result.TotalAmount,
		})
		if err != nil {
			log.Println("Err communicating with api:", err)
		}

		time.Sleep(10 * time.Second)
	}
}

func Call(address string) (*GetTxOutResponse, error) {

	body := map[string]interface{}{
		"id":      "blockchain_scraper",
		"method":  address,
		"jsonrpc": "1.0",
		"params":  []string{},
	}
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s://%s:%d", PROTOCOL, HOST, PORT), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(USERNAME, PASSWORD)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Length", "0")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b := GetTxOutResponse{}

	err = json.NewDecoder(resp.Body).Decode(&b)
	if err != nil {
		return nil, err
	}

	return &b, nil
}
