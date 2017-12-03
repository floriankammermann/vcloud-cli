package vcdapi

import (
	"net/http"
	"fmt"
	"github.com/floriankammermann/vcloud-cli/types"
	"errors"
)

func GetAllVApp(url string) {
	req, err := http.NewRequest("GET", url+"/api/query?type=vApp&fields=name&pageSize=512", nil)
	req.Header.Set("x-vcloud-authorization", vcdClient.VAToken)
	req.Header.Set("Accept", "application/*+xml;version=5.5")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	queryRes := new(types.QueryResultRecordsType)
	decodeBody(resp, queryRes)

	for _, vapp := range queryRes.VAppRecord {
		fmt.Printf("vApp Name [%s]\n", vapp.Name)
	}

}

func GetAllocatedIpsForNetworkName(url string, networkname string) error {

	if len(networkname) == 0 {
		return errors.New("networkname is empty")
	}

	path := "/api/query?type=orgNetwork&fields=name&filter=name=="+networkname
	queryRes := new(types.QueryResultRecordsType)
	ExecRequest(url, path, queryRes)

	for _, net := range queryRes.OrgNetworkRecord {
		fmt.Printf("org network name [%s]\n", net.Name)
	}

	if len(queryRes.OrgNetworkRecord) > 1 {
		return errors.New("found more than one org network for name: " + networkname)
	}
	if len(queryRes.OrgNetworkRecord) == 0 {
		return errors.New("found no org network for name: " + networkname)
	}
	getAllocatedIpsForNetworkHref(queryRes.OrgNetworkRecord[0].HREF)
	return nil
}

func getAllocatedIpsForNetworkHref(networkref string) error {
	if len(networkref) == 0 {
		return errors.New("networkref is empty")
	}

	fmt.Printf("the network href: [%s]\n", networkref)

	queryRes := new(types.AllocatedIpAddressesType)
	ExecRequest(networkref, "/allocatedAddresses", queryRes)

	for _, ipAddress := range queryRes.IpAddress {
		fmt.Printf("ip [%s] ", ipAddress.IpAddress)
		for _, link := range ipAddress.Link {
			if "down" == link.Rel && "application/vnd.vmware.vcloud.vApp+xml" == link.Type {
				fmt.Printf("vApp: [%s] ", link.Name)
			}
			if "down" == link.Rel && "application/vnd.vmware.vcloud.vm+xml" == link.Type {
				fmt.Printf("vm: [%s] ", link.Name)
			}
		}
		fmt.Print("\n")
	}

	return nil
}
