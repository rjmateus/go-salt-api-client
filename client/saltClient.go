package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/rjmateus/go-salt-api-client/calls"
	"github.com/rjmateus/go-salt-api-client/target"
	"io/ioutil"
	"log"
	"net/http"
)

type SaltClient struct {
	BaseUrl    string
	Username   string
	Password   string
	Eauth      string
	HttpClient *http.Client
}

type RunResult struct {
	Return []interface{} `json:"return"`
}

func NewClient(url, username, password, eauth string) SaltClient {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore SSL certificates
	}
	httpClient := &http.Client{Transport: transCfg}

	return SaltClient{
		BaseUrl:    url,
		Username:   username,
		Password:   password,
		Eauth:      eauth,
		HttpClient: httpClient,
	}
}

func (c *SaltClient) Run(clilentType ClientType, fun string, tgt target.SaltTarget,
	args []interface{}, kwargs map[string]interface{}) (data *RunResult, err error) {

	props := make(map[string]interface{})
	props["username"] = c.Username
	props["password"] = c.Password
	props["eauth"] = c.Eauth

	props["client"] = clilentType

	props["fun"] = fun
	if args != nil {
		props["arg"] = args
	}
	if kwargs != nil {
		props["kwarg"] = kwargs
	}

	for k, v := range tgt.GetPros() {
		props[k] = v
	}

	postBody, _ := json.Marshal([]map[string]interface{}{props})
	resp, err := c.HttpClient.Post(c.BaseUrl+"/run", "application/json", bytes.NewBuffer(postBody))

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("net/http: return error %q", resp.StatusCode)
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Raw data: >%s<\n", body)

	var decodedData = &RunResult{}
	err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&decodedData)
	if err != nil {
		log.Fatalln(err)
	}

	return decodedData, nil
}

//call(Call<?> call, Client client, Optional<Target<?>> target,
//Map<String, Object> custom, TypeToken<R> type, AuthMethod auth)

func (c *SaltClient) Call(call calls.Call, clientType ClientType, tgt *target.SaltTarget) (data []byte, err error) {

	props := make(map[string]interface{})
	props["username"] = c.Username
	props["password"] = c.Password
	props["eauth"] = c.Eauth

	props["client"] = clientType

	if call.GetPayload() != nil {
		for k, v := range call.GetPayload() {
			props[k] = v
		}
	}

	if tgt != nil {
		for k, v := range tgt.GetPros() {
			props[k] = v
		}
	}

	postBody, _ := json.Marshal([]map[string]interface{}{props})
	resp, err := c.HttpClient.Post(c.BaseUrl+"/run", "application/json", bytes.NewBuffer(postBody))

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("net/http: return error %q", resp.StatusCode)
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Raw data: >%s<\n", body)
	return body, nil
}
