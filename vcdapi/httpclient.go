package vcdapi


import (
	"net/http"
	"fmt"
)


var vcdClient *VcdClientType

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