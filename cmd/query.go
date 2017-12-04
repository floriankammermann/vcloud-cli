package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/floriankammermann/vcloud-cli/vcdapi"
	"github.com/spf13/viper"
)

var networkname string
var edgegatewayname string

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "execute queries",
	Long: "execute queries",
}

var allocatedipCmd = &cobra.Command{
	Use:   "allocatedip",
	Short: "allocatedip for an org network",
	Long: "get all allocated ips of an org network",
	Run: func(cmd *cobra.Command, args []string) {

		if len(networkname) > 0 {
			url := viper.GetString("url")
			user := viper.GetString("user")
			password := viper.GetString("password")
			org := viper.GetString("org")
			vcdapi.GetAuthToken(url, user, password, org)
			vcdapi.GetAllocatedIpsForNetworkName(url, networkname)
		} else {
			fmt.Println("you have to provide the networkname")
		}
	},
}

var natruleCmd = &cobra.Command{
	Use:   "natrule",
	Short: "natrules for an edge gateway",
	Long: "get all nat rules for an edge gateway",
	Run: func(cmd *cobra.Command, args []string) {

		if len(edgegatewayname) > 0 {
			url := viper.GetString("url")
			user := viper.GetString("user")
			password := viper.GetString("password")
			org := viper.GetString("org")
			vcdapi.GetAuthToken(url, user, password, org)
			vcdapi.GetNATRulesForEdgeGatweway(url, edgegatewayname)
		} else {
			fmt.Println("you have to provide the edgegatewayname")
		}
	},
}

var vmCommand = &cobra.Command{
	Use:   "vapp",
	Short: "vApps of org",
	Long: "show all vApps of the org",
	Run: func(cmd *cobra.Command, args []string) {
		url := viper.GetString("url")
		user := viper.GetString("user")
		password := viper.GetString("password")
		org := viper.GetString("org")
		vcdapi.GetAuthToken(url, user, password, org)
		vcdapi.GetAllVApp(url)
	},
}

var orgvdcCmd = &cobra.Command{
	Use:   "orgvdc",
	Short: "query orgvdc requests",
	Long: "query orgvdc requests",
	Run: func(cmd *cobra.Command, args []string) {
		url := viper.GetString("url")
		user := viper.GetString("user")
		password := viper.GetString("password")
		org := viper.GetString("org")
		vcdapi.GetAuthToken(url, user, password, org)
		vcdapi.GetAllVdcorg(url)
	},
}
func init() {
	queryCmd.AddCommand(allocatedipCmd)
	allocatedipCmd.Flags().StringVarP(&networkname, "network", "n", "", "network name to search allocated ips on")
	queryCmd.AddCommand(natruleCmd)
	natruleCmd.Flags().StringVarP(&edgegatewayname, "edgegateway", "e", "", "edgegateway name to search nat rules on")
	queryCmd.AddCommand(vmCommand)
	queryCmd.AddCommand(orgvdcCmd)
	RootCmd.AddCommand(queryCmd)
}
