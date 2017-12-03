package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/floriankammermann/vcloud-cli/vcdapi"
	"github.com/spf13/viper"
)

var path string

var requestCmd = &cobra.Command{
	Use:   "request",
	Short: "execute requests",
	Long: "execute requests",
	Run: func(cmd *cobra.Command, args []string) {

		if len(path) > 0 {
			url := viper.GetString("url")
			user := viper.GetString("user")
			password := viper.GetString("password")
			org := viper.GetString("org")
			vcdapi.GetAuthToken(url, user, password, org)
			vcdapi.ExecRequest(url, path, nil)
		} else {
			fmt.Println("you have to provide the path")
		}
	},
}

func init() {
	requestCmd.Flags().StringVarP(&path, "path", "n", "", "url to request")
	RootCmd.AddCommand(requestCmd)
}
