package vcdapi

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/xml"
	"github.com/spf13/viper"
)

var vcdClient *VcdClientType
var	Verbose string // verbose output

type VcdClientType struct {
	VAToken string // vCloud Air authorization token
}

func GetAuthToken(url string, user string, password string, org string) {
	req, err := http.NewRequest("POST", url+"/api/sessions", nil)
	req.Header.Set("Accept", "application/*+xml;version=5.5")
	req.SetBasicAuth(user+"@"+org, password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	auth := resp.Header.Get("x-vcloud-authorization")
	vcdClient = &VcdClientType{
		VAToken: auth,
	}
	fmt.Printf("authorization: [%s]\n", auth)
}

func ExecRequest(url string, path string, queryRes interface{}) {

	if viper.GetString("verbose") == "true" {
		fmt.Printf("url: [%s], path [%s]", url, path)
	}

	req, err := http.NewRequest("GET", url+path, nil)
	req.Header.Set("x-vcloud-authorization", vcdClient.VAToken)
	req.Header.Set("Accept", "application/*+xml;version=5.5")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	decodeBody(resp, queryRes)
}

// decodeBody is used to XML decode a response body
func decodeBody(resp *http.Response, out interface{}) error {

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if viper.GetString("verbose") == "true" {
		fmt.Println("response Body:", string(body))
	}

	// Unmarshal the XML.
	if err = xml.Unmarshal(body, &out); err != nil {
		return err
	}

	return nil
}