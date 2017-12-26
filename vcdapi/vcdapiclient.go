package vcdapi

import (
	"fmt"
	"github.com/floriankammermann/vcloud-cli/types"
	"errors"
	"strconv"
	"text/tabwriter"
	"os"
)

func GetAllVdcorg(url string) {

	path := "/api/query?type=orgVdc&fields=name&pageSize=512"
	queryRes := new(types.QueryResultRecordsType)
	ExecRequest(url, path, queryRes)

	for _, vdc := range queryRes.OrgVdcRecord {
		fmt.Printf("orgVdc Name [%s]\n", vdc.Name)
	}

}

func GetAllOrgNetworks(url string) {

	path := "/api/query?type=orgVdcNetwork&fields=name&pageSize=512"
	queryRes := new(types.QueryResultRecordsType)
	ExecRequest(url, path, queryRes)

	// create a new tabwriter
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, '\t', tabwriter.Debug|tabwriter.AlignRight)
	fmt.Fprintln(w, "Name\tHref\t")
	for _, orgnetwork := range queryRes.OrgVdcNetworkRecord {
		fmt.Fprintf(w, "%s\t%s\t\n", orgnetwork.Name, orgnetwork.HREF)
	}
	fmt.Fprintln(w)
	w.Flush()


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

	for _, net := range queryRes.OrgVdcNetworkRecord {
		fmt.Printf("org network name [%s]\n", net.Name)
	}

	if len(queryRes.OrgVdcNetworkRecord) > 1 {
		return errors.New("found more than one org network for name: " + networkname)
	}
	if len(queryRes.OrgVdcNetworkRecord) == 0 {
		return errors.New("found no org network for name: " + networkname)
	}
	getAllocatedIpsForNetworkHref(queryRes.OrgVdcNetworkRecord[0].HREF)
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

type RenderEdgegatewayResults func(edgegateway string) error

func GetEdgeGatweway(url string, edgegatewayname string, renderer RenderEdgegatewayResults) error {

	if len(edgegatewayname) == 0 {
		return errors.New("networkname is empty")
	}

	path := "/api/query?type=edgeGateway&fields=name&filter=name=="+edgegatewayname
	queryRes := new(types.QueryResultRecordsType)
	ExecRequest(url, path, queryRes)

	for _, net := range queryRes.OrgVdcNetworkRecord {
		fmt.Printf("org network name [%s]\n", net.Name)
	}

	if len(queryRes.EdgeGatewayRecord) > 1 {
		return errors.New("found more than one org network for name: " + edgegatewayname)
	}
	if len(queryRes.EdgeGatewayRecord) == 0 {
		return errors.New("found no org network for name: " + edgegatewayname)
	}
	renderer(queryRes.EdgeGatewayRecord[0].HREF)
	return nil
}

func RenderNATRulesForEdgegateway(edgegatewayhref string) error {
	if len(edgegatewayhref) == 0 {
		return errors.New(" edgegatewayref is empty")
	}

	fmt.Printf("the edgegateway href: [%s]\n", edgegatewayhref)

	queryRes := new(types.EdgeGateway)
	ExecRequest(edgegatewayhref, "", queryRes)

	fmt.Println("[id]    [type] [enabled] [interface]    [originalIp]   [originalPort]    [TranslatedIp]     [TranslatedPort]   [protocol]")
	for _, natRule := range queryRes.Configuration.EdgeGatewayServiceConfiguration.NatService.NatRule {

		fmt.Printf("[%s] [%s] [%t]   [%s] [%s] [%s] [%s] [%s] [%s]", natRule.ID, natRule.RuleType, natRule.IsEnabled,
																	 natRule.GatewayNatRule.Interface.Name, natRule.GatewayNatRule.OriginalIP,
																	 natRule.GatewayNatRule.OriginalPort, natRule.GatewayNatRule.TranslatedIP,
																	 natRule.GatewayNatRule.TranslatedPort, natRule.Description)
		fmt.Print("\n")
	}

	return nil
}

func RenderFirewallRulesForEdgegateway(edgegatewayhref string) error {
	if len(edgegatewayhref) == 0 {
		return errors.New(" edgegatewayref is empty")
	}

	fmt.Printf("the edgegateway href: [%s]\n", edgegatewayhref)

	queryRes := new(types.EdgeGateway)
	ExecRequest(edgegatewayhref, "", queryRes)

	fmt.Println("[id]    [source] [destination] [protocol]    [policy]   [description]")
	for _, fwRule := range queryRes.Configuration.EdgeGatewayServiceConfiguration.FirewallService.FirewallRule {

		var protocol string
		if fwRule.Protocols.Any == true {
			protocol = "any"
		} else if fwRule.Protocols.ICMP == true {
			protocol = "icmp"
		} else if fwRule.Protocols.TCP == true {
			protocol = "tcp"
		} else if fwRule.Protocols.UDP == true {
			protocol = "udp"
		}

		source := fwRule.SourceIP+":"+strconv.Itoa(fwRule.SourcePort)
		destination := fwRule.DestinationIP+":"+fwRule.DestinationPortRange
		fmt.Printf("[%s] [%s] [%s]   [%s] [%s] [%s]", fwRule.ID, source, destination, protocol, fwRule.Policy, fwRule.Description)
		fmt.Print("\n")
	}

	return nil
}
