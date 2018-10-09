package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/diadata-org/api-golang/pkg/dia"
)

type RequestBody struct {
	ID         int64         `json:"id"`
	Method     string        `json:"method"`
	Parameters []interface{} `json:"params"`
	Version    string        `json:"jsonrpc"`
}

type ResponseBody struct {
	ID      int     `json:"id"`
	JSONRPC string  `json:"jsonrpc"`
	Result  Account `json:"result"`
}

type CountResponseBody struct {
	ID      int    `json:"id"`
	JSONRPC string `json:"jsonrpc"`
	Result  int    `json:"result"`
}

type Account struct {
	Balances []Balance `json:"balances"`
}

type Balance struct {
	Asset string `json:"asset"`
	Value string `json:"value"`
}

const (
	apiVersion = "2.0"
	symbol     = "NEO"
	// URI of neoclient running in docker (port: 10332 is for JSON-RPC via HTTP)
	NodeURI       = "http://neonode:10332"
	TotalSupply   = 100000000
	TargetAddress = "AQVh2pG732YvtNaxEGkQUei3YA4cvo7d2i"
	// AssetId referes to NEO (not GAS or other Coins)
	TargetAssetId = "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b"

	scrapInterval = time.Second * 10
	syncInterval  = time.Second * 30
)

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
		ok := Ping(NodeURI)
		if ok {

			Sync(NodeURI)

			neo, err := GetNEO(TargetAddress, NodeURI)

			if err != nil {
				fmt.Println("Error Fetching NEO ", err)
			} else {
				circulatingSupply := TotalSupply - neo

				client.SendSupply(&dia.Supply{
					Symbol:            "NEO",
					CirculatingSupply: float64(circulatingSupply),
				})
				fmt.Println("New CirculatingSupply of NEO : ", circulatingSupply)
			}

		} else {
			fmt.Println("Error connecting to NEO node, Check if the docker service is running")
		}

		time.Sleep(scrapInterval)
	}
}

func Sync(Node string) bool {
	for {
		globalBlockCount, err := GetGlobalBlockCount()

		if err != nil {
			fmt.Println("Error Fetching Global Block Count ", err)
		} else {
			localBlockCount, err := GetBlockCount(Node)

			if err != nil {
				fmt.Println("Error Fetching Block Count", err)
			} else {
				if localBlockCount >= globalBlockCount {
					return true
				} else {
					fmt.Println("Syncing - Local Block Count: ", localBlockCount, " Global Block Count:", globalBlockCount)
				}
			}
		}

		time.Sleep(scrapInterval)
	}
	return true
}

func GetGlobalBlockCount() (int, error) {
	seedNodes := [16]string{
		"http://seed1.aphelion-neo.com:10332",
		"http://seed2.aphelion-neo.com:10332",
		"http://seed3.aphelion-neo.com:10332",
		"http://seed2.neo.org:10332",
		"http://seed3.neo.org:10332",
		"http://seed4.neo.org:10332",
		"http://seed5.neo.org:10332",
		"http://seed2.ngd.network:10332",
		"http://seed3.ngd.network:10332",
		"http://seed4.ngd.network:10332",
		"http://seed5.ngd.network:10332",
		"http://seed6.ngd.network:10332",
		"http://seed7.ngd.network:10332",
		"http://seed8.ngd.network:10332",
		"http://seed9.ngd.network:10332",
		"http://seed10.ngd.network:10332"}

	for _, seedNode := range seedNodes {
		globalBlockCount, err := GetBlockCount(seedNode)

		if err == nil {
			return globalBlockCount, nil
		}
	}

	return 0, fmt.Errorf("Seed Nodes are not Responding")
}

// Return NEO at address
func GetNEO(Address string, Node string) (int, error) {
	account, err := GetAccountState(TargetAddress, Node)

	if err != nil {
		return 0, err
	} else {
		value := 0

		for _, balance := range account.Balances {
			if balance.Asset == TargetAssetId {
				value, err = strconv.Atoi(balance.Value)

				if err != nil {
					return 0, err
				} else {
					return value, nil
				}
			}
		}
		return 0, nil
	}
}

// Returns Account State at the address
func GetAccountState(Address string, Node string) (*Account, error) {
	requestBodyParams := []interface{}{
		Address,
	}
	var response ResponseBody

	err := executeRequest("getaccountstate", requestBodyParams, Node, &response)
	if err != nil {
		return nil, err
	}

	return &response.Result, nil
}

// Returns Block count at node
func GetBlockCount(Node string) (int, error) {
	var response CountResponseBody

	err := executeRequest("getblockcount", []interface{}{}, Node, &response)
	if err != nil {
		return 0, err
	}

	return response.Result, nil
}

// Ping checks if the node is online.
func Ping(node string) bool {
	parsedURI, err := url.Parse(node)
	if err != nil {
		return false
	}

	conn, err := net.Dial("tcp", parsedURI.Host)
	if err != nil {
		return false
	}

	_ = conn.Close()

	return true
}

// Makes a POST request to the node with given method, parameters
func executeRequest(method string, bodyParameters []interface{}, nodeURI string, model interface{}) error {
	var body []byte
	var err error

	body, err = json.Marshal(RequestBody{
		ID:         1234,
		Method:     method,
		Parameters: bodyParameters,
		Version:    apiVersion,
	})
	if err != nil {
		return err
	}

	ioBody := bytes.NewReader(body)

	request, err := http.NewRequest("POST", nodeURI, ioBody)
	if err != nil {
		return err
	}

	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode > 200 {
		return fmt.Errorf(
			"Non-200 status code returned from NEO node, got: '%d'",
			response.StatusCode,
		)
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &model)
	if err != nil {
		return err
	}

	return nil
}
