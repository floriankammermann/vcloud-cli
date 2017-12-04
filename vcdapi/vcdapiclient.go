package vcdapi

import (
	"fmt"
	"github.com/floriankammermann/vcloud-cli/types"
	"errors"
)

func GetAllVdcorg(url string) {

	path := "/api/query?type=orgVdc&fields=name&pageSize=512"
	queryRes := new(types.QueryResultRecordsType)
	ExecRequest(url, path, queryRes)

	for _, vapp := range queryRes.OrgVdcRecord {
		fmt.Printf("orgVdc Name [%s]\n", vapp.Name)
	}

}

func GetAllVApp(url string) {

	path := "/api/query?type=vApp&fields=name&pageSize=512"
	queryRes := new(types.QueryResultRecordsType)
	ExecRequest(url, path, queryRes)

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

func GetNATRulesForEdgeGatweway(url string, edgegatewayname string) error {

	if len(edgegatewayname) == 0 {
		return errors.New("networkname is empty")
	}

	path := "/api/query?type=edgeGateway&fields=name&filter=name=="+edgegatewayname
	queryRes := new(types.QueryResultRecordsType)
	ExecRequest(url, path, queryRes)

	for _, net := range queryRes.OrgNetworkRecord {
		fmt.Printf("org network name [%s]\n", net.Name)
	}

	if len(queryRes.EdgeGatewayRecord) > 1 {
		return errors.New("found more than one org network for name: " + edgegatewayname)
	}
	if len(queryRes.EdgeGatewayRecord) == 0 {
		return errors.New("found no org network for name: " + edgegatewayname)
	}
	getNATRulesForEdgeGatewayHref(queryRes.EdgeGatewayRecord[0].HREF)
	return nil
}

func getNATRulesForEdgeGatewayHref(edgegatewayref string) error {
	if len(edgegatewayref) == 0 {
		return errors.New(" edgegatewayref is empty")
	}

	fmt.Printf("the edgegateway href: [%s]\n", edgegatewayref)

	queryRes := new(types.EdgeGateway)
	ExecRequest(edgegatewayref, "", queryRes)

	fmt.Println("[id]    [type] [enabled] [interface]    [originalIp]   [originalPort]    [TranslatedIp]     [TranslatedPort]   [protocol]")
	for _, natRule := range queryRes.Configuration.EdgeGatewayServiceConfiguration.NatService.NatRule {

		fmt.Printf("[%s] [%s] [%t]   [%s] [%s] [%s] [%s] [%s]", natRule.ID, natRule.RuleType, natRule.IsEnabled,
																	 natRule.GatewayNatRule.Interface.Name, natRule.GatewayNatRule.OriginalIP,
																	 natRule.GatewayNatRule.OriginalPort, natRule.GatewayNatRule.TranslatedIP,
																	 natRule.GatewayNatRule.TranslatedPort)
		fmt.Print("\n")
	}

	return nil
}
