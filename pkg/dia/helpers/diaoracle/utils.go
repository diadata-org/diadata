package diaoracle

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
)

type Wallet struct {
	Key        string
	PassPhrase string
}

func NewWallet(filePath string) (*Wallet, error) {
	var lines []string
	file, err := os.Open(filePath) // Read in key information
	if err != nil {
		return nil, err
	}
	defer func() {
		err = file.Close()
		if err != nil {
			return
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	if len(lines) != 2 {
		return nil, errors.New("secrets file should have exactly two lines")
	}
	return &Wallet{Key: lines[0], PassPhrase: lines[1]}, nil

}

func getQuotationFromDia(symbol string) (*models.Quotation, error) {
	contents, statusCode, err := utils.GetRequest(dia.BaseUrl + "/v1/quotation/" + strings.ToUpper(symbol))
	if err != nil {
		return nil, err
	}
	if statusCode != 200 {
		return nil, fmt.Errorf("error on dia api with return code %d", statusCode)
	}

	var quotation models.Quotation
	err = quotation.UnmarshalBinary(contents)
	if err != nil {
		return nil, err
	}
	return &quotation, nil
}

func getSupplyFromDia(symbol string) (*dia.Supply, error) {
	contents, statusCode, err := utils.GetRequest(dia.BaseUrl + "/v1/supply/" + symbol)
	if err != nil {
		return nil, err
	}
	if statusCode != 200 {
		return nil, fmt.Errorf("error on dia api with return code %d", statusCode)
	}

	var supply dia.Supply
	err = supply.UnmarshalBinary(contents)
	if err != nil {
		return nil, err
	}
	return &supply, nil
}

func getDefiRatesFromDia(protocol string, symbol string) (*dia.DefiRate, error) {
	response, err := http.Get(dia.BaseUrl + "/v1/defiLendingRate/" + strings.ToUpper(protocol) + "/" + strings.ToUpper(symbol) + "/" + strconv.FormatInt(time.Now().Unix(), 10))
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()

		if 200 != response.StatusCode {
			return nil, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
		}

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		var b dia.DefiRate
		err = b.UnmarshalBinary(contents)
		if err == nil {
			return &b, nil
		}
		return nil, err
	}
}

func getDefiStateFromDia(protocol string) (*dia.DefiProtocolState, error) {
	response, err := http.Get(dia.BaseUrl + "/v1/defiLendingState/" + strings.ToUpper(protocol))
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()

		if 200 != response.StatusCode {
			return nil, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
		}

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		var b dia.DefiProtocolState
		err = b.UnmarshalBinary(contents)
		if err == nil {
			return &b, nil
		}
		return nil, err
	}
}

func getDEXFromDia(dexname string, symbol string) (*models.Points, error) {
	response, err := http.Get(dia.BaseUrl + "/v1/chartPoints/MAIR120/" + dexname + "/" + strings.ToUpper(symbol))
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()

		if 200 != response.StatusCode {
			return nil, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
		}

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		var b models.Points
		err = b.UnmarshalBinary(contents)
		if err == nil {
			return &b, nil
		}
		return nil, err
	}
}

func getForeignQuotationFromDia(source, symbol string) (*models.ForeignQuotation, error) {
	response, err := http.Get(dia.BaseUrl + "/v1/foreignQuotation/" + strings.Title(strings.ToLower(source)) + "/" + strings.ToUpper(symbol))
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if 200 != response.StatusCode {
		return nil, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var quotation models.ForeignQuotation
	err = quotation.UnmarshalBinary(contents)
	if err != nil {
		return nil, err
	}
	return &quotation, nil
}

func getFarmingPoolFromDia(protocol string, poolID string) (*models.FarmingPool, error) {
	response, err := http.Get(dia.BaseUrl + "v1/FarmingPoolData/" + strings.ToUpper(protocol) + "/" + poolID)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if 200 != response.StatusCode {
		return nil, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var fp models.FarmingPool
	err = fp.UnmarshalBinary(contents)
	if err == nil {
		return &fp, nil
	}
	return nil, err
}
func getECBRatesFromDia(symbol string) (*models.CurrencyChange, error) {
	response, err := http.Get(dia.BaseUrl + "/v1/coins")
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()

		if 200 != response.StatusCode {
			return nil, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
		}

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		var b models.Coins
		err = b.UnmarshalBinary(contents)
		if err != nil {
			return nil, err
		}

		for _, change := range b.Change.USD {
			if strings.ToUpper(change.Symbol) == strings.ToUpper(symbol) {
				return &change, nil
			}
		}
		return nil, nil
	}
}

func getIndexValueFromDia(symbol string) (*models.CryptoIndex, error) {
	response, err := http.Get(dia.BaseUrl + "/v1/index/" + symbol + "?starttime=" + strconv.FormatInt(time.Now().Add(-200*time.Second).Unix(), 10) + "&endtime=" + strconv.FormatInt(time.Now().Unix(), 10))
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if 200 != response.StatusCode {
		return nil, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var indices []models.CryptoIndex
	err = json.Unmarshal([]byte(contents), &indices)
	if err != nil {
		return nil, err
	}
	return &indices[0], nil
}
