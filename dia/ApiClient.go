package dia

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	config *ConfigApi
	token  string
}

const baseURL string = "https://api.diadata.org/"

type response struct {
	Token string
}

func (c *Client) refresh() error {

	url := baseURL + "auth/refresh_token"

	req, err := http.NewRequest("GET", url, nil)

	bytes, err := c.doRequest(req, false)
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
	url := baseURL + "login"

	jsonStr, err := json.Marshal(&login{
		Username: c.config.ApiKey,
		Password: c.config.SecretKey,
	})

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var r response
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil
	}
	c.token = r.Token
	return nil
}

func NewClient(config *ConfigApi) *Client {
	c := &Client{
		config: config,
		token:  "",
	}
	err := c.login()
	if err != nil {
		return nil
	}
	return c
}

func (c *Client) doRequest(req *http.Request, refresh bool) ([]byte, error) {

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

	if 200 != resp.StatusCode {

		if refresh {
			if resp.StatusCode == 401 {
				err = c.refresh()
				if err != nil {
					err = c.login()
					if err == nil {
						return c.doRequest(req, true)
					}
				}
			}
		}
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

func (c *Client) SendSupply(s *Supply) error {

	log.Println(s)
	jsonStr, err := json.Marshal(s)
	if err != nil {
		return err
	}

	url := baseURL + "v1/supply"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	_, err = c.doRequest(req, true)
	if err != nil {
		return err
	}
	return nil
}
