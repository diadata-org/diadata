package dia

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
	"io/ioutil"
	"net/http"
	"os/user"
	"time"
)

type Client struct {
	config                *ConfigApi
	token                 string
	lastSupplyUpdateTime  time.Time
	lastSupplyUpdateValue float64
	url                   string
}

type Symbols struct {
	Symbols []string
}

const BaseUrl string = "https://api.diadata.org/"

type response struct {
	Token string
}

func (c *Client) refresh() error {

	url := c.url + "auth/refresh_token"

	req, err := http.NewRequest("GET", url, nil)

	bytes, err := c.DoRequest(req, false)
	if err != nil {
		return err
	}
	var r response
	err = json.Unmarshal(bytes, &r)
	if err != nil {
		fmt.Println("error:", err)
		return err
	}
	c.token = r.Token
	return nil
}

func (c *Client) login() error {

	client := &http.Client{}

	type login struct {
		Username string
		Password string
	}
	url := c.url + "login"

	jsonStr, err := json.Marshal(&login{
		Username: c.config.ApiKey,
		Password: c.config.SecretKey,
	})

	if err != nil {
		log.Println(err)
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var r response
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Println(err)
		return err
	}

	c.token = r.Token

	return nil
}

func GetSupply(symbol string) (*Supply, error) {
	url := BaseUrl + "/v1/supply/" + symbol
	log.Println("Checking supply for", symbol, "on", url)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()
		if 200 != response.StatusCode {
			return nil, fmt.Errorf("error on %v -> %v", url, response.StatusCode)
		}
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		log.Debug("%s\n", string(contents))
		var b Supply
		err = b.UnmarshalBinary(contents)
		if err == nil {
			log.Debug("got", b)
			return &b, nil
		}
		return nil, err
	}
}

// TODO remove URL
func GetSymbolsList(url string) ([]string, error) {
	log.Println("getSymbolList")
	response, err := http.Get(url + "/v1/symbols")
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()
		if 200 != response.StatusCode {
			return nil, fmt.Errorf("Error getSymbolList")
		}
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		log.Debug("%s\n", string(contents))
		var b Symbols

		err = json.Unmarshal(contents, &b)

		if err == nil {
			return b.Symbols, nil
		}
		return nil, err
	}
}

func GetConfigApi() *ConfigApi {
	var c ConfigApi
	configFile := "/run/secrets/api_diadata"
	err := gonfig.GetConf(configFile, &c)
	if err != nil {
		log.Errorln("GetConfigApi", err)
		usr, _ := user.Current()
		dir := usr.HomeDir
		configFile = dir + "/secrets/api_diadata.json"
		err = gonfig.GetConf(configFile, &c)
	}
	if err != nil {
		log.Println(err)
		return nil
	} else {
		log.Println("Loaded secret in", configFile)
	}
	return &c
}

func NewClientWithUrl(config *ConfigApi, url string) *Client {
	c := &Client{
		config: config,
		token:  "",
		url:    url,
	}
	err := c.login()
	if err != nil {
		log.Println(err)
		return nil
	}
	return c
}

func NewClient(config *ConfigApi) *Client {
	c := &Client{
		config: config,
		token:  "",
		url:    BaseUrl,
	}
	err := c.login()
	if err != nil {
		log.Println(err)
		return nil
	}
	return c
}

func (c *Client) DoRequest(req *http.Request, refresh bool) ([]byte, error) {

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Debug("StatusCode", resp.StatusCode)

	if 200 != resp.StatusCode {

		if refresh {
			if resp.StatusCode == 401 {
				err = c.refresh()
				if err != nil {
					err = c.login()
					if err == nil {
						return c.DoRequest(req, true)
					}
				}
			}
		}
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

func (c *Client) SendSupplyWithForceOption(s *Supply, force bool) error {
	lastUpdate := time.Since(c.lastSupplyUpdateTime)
	if lastUpdate.Hours() >= 1.0 || c.lastSupplyUpdateValue != s.CirculatingSupply || force {
		c.lastSupplyUpdateTime = time.Now()
		c.lastSupplyUpdateValue = s.CirculatingSupply
		return c.sendSupply(s)
	} else {
		log.Println("Skipping sending to API", s, "last update:", lastUpdate)
		return nil
	}
}

func (c *Client) SendSupply(s *Supply) error {
	return c.SendSupplyWithForceOption(s, false)
}

func (c *Client) sendSupply(s *Supply) error {

	jsonStr, err := json.Marshal(s)
	if err != nil {
		return err
	}

	url := c.url + "v1/supply"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	_, err = c.DoRequest(req, true)
	if err != nil {
		log.Println("Error: SendSupply", err)
		return err
	}

	return nil
}
