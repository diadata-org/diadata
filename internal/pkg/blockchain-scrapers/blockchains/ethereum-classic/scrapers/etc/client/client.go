package ethclassic

import (
	"encoding/json"
	"gopkg.in/resty.v1"
	"math"
	"strconv"
)

//EthClient to retrieve bocks from etc
type EthClient struct {
	URL        string
	LastBlock  int64
	LastSupply float64
}

//EthBlock To estimate total supply we just need block ID and uncles
type EthBlock struct {
	ID     string   `json:"number"`
	Uncles []string `json:"uncles"`
}

//Constants used to estimate total circulation
const eraBase = 5000000
const uncleReward = 0.03125

func (e *EthClient) post(content interface{}) ([]byte, error) {
	url := e.URL
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(content).
		Post(url)
	if err != nil {
		return resp.Body(), err
	}
	return resp.Body(), nil
}

//GetTotalSupplyAtBlock recursive function to compute total ETH
// this function needs to be called periodically since it depends on all previousblocks
func (e *EthClient) GetTotalSupplyAtBlock(number int64) (float64, error) {
	if number == e.LastBlock {
		return e.LastSupply, nil
	}
	if r, err := e.GetBlockReward(number); err == nil {
		if s, err := e.GetTotalSupplyAtBlock(number - 1); err == nil {
			e.LastBlock = number
			e.LastSupply = s + r
			return e.LastSupply, nil
		} else {
			return 0, err
		}
	} else {
		return r, err
	}
}

//GetBlockReward Compute block (N) reward using
//E=(N-1)/5000000
//r=5*0.8^E*(1+0.03125*nUncles )
func (e *EthClient) GetBlockReward(number int64) (float64, error) {
	if b, err := e.getBlock(number); err == nil {
		E := (number - 1) / eraBase
		baseReward := 5 * math.Pow(0.8, float64(E))
		var uR float64
		for i := 0; i < len(b.Uncles); i++ {
			if r, err := e.getUncleRewardByIDAndPosition(number, i); err != nil {
				return 0, err
			} else {
				uR += r
			}
		}
		return baseReward*(1.+uncleReward*float64(len(b.Uncles))) + uR, nil
	} else {
		return 0, err
	}
}

//GetLatestBlockNumber this block will return 0 until the
// node is fully initialized if --fast option is provided to geth
func (e *EthClient) GetLatestBlockNumber() (int64, error) {
	if b, err := e.getBlock(-1); err == nil {
		return strconv.ParseInt(b.ID, 0, 64)
	} else {
		return 0, err
	}
}

//will return latest block if number = 1
func (e *EthClient) getBlock(number int64) (EthBlock, error) {
	// var block EthBlock
	type msg struct {
		Jsonrpc string
		Method  string
		Params  []interface{}
		Id      int
	}
	type response struct {
		Jsonrpc string
		Id      int
		Result  EthBlock
	}
	var r response

	content := msg{
		Jsonrpc: "2.0",
		Method:  "eth_getBlockByNumber",
		Params:  make([]interface{}, 2),
		Id:      1,
	}
	// retrieve latest block
	if number == -1 {
		content.Params[0] = "latest"
	} else {
		content.Params[0] = "0x" + strconv.FormatInt(number, 16)
	}
	content.Params[1] = true
	if data, err := e.post(content); err == nil {
		err = json.Unmarshal(data, &r)
		return r.Result, err
	} else {
		return r.Result, err
	}
}
func (e *EthClient) getUncleRewardByIDAndPosition(number int64, postion int) (float64, error) {
	if b, err := e.getUncleBlockByIDAndPosition(number, postion); err == nil {
		E := (number - 1) / eraBase
		baseReward := 5. * math.Pow(0.8, float64(E))
		id, _ := strconv.ParseInt(b.ID, 0, 64)
		return float64(int(id)+8-int(number)) * baseReward / 8., nil
	} else {
		return 0, err
	}
}
func (e *EthClient) getUncleBlockByIDAndPosition(number int64, postion int) (EthBlock, error) {
	// var block EthBlock
	type msg struct {
		Jsonrpc string
		Method  string
		Params  []interface{}
		Id      int
	}
	type response struct {
		Jsonrpc string
		Id      int
		Result  EthBlock
	}
	var r response

	content := msg{
		Jsonrpc: "2.0",
		Method:  "eth_getUncleByBlockNumberAndIndex",
		Params:  make([]interface{}, 2),
		Id:      1,
	}
	// retrieve latest block
	if number == -1 {
		content.Params[0] = "latest"
	} else {
		content.Params[0] = "0x" + strconv.FormatInt(number, 16)
	}
	content.Params[1] = postion
	if data, err := e.post(content); err == nil {
		err = json.Unmarshal(data, &r)
		return r.Result, err
	} else {
		return r.Result, err
	}

}
