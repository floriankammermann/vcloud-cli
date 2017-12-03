package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "vcloud-cli",
	Short: "a command line interface for the vcloud director api",
	Long: "a command line interface for the vcloud director api",
}


// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	RootCmd.PersistentFlags().String("url", "", "url of vcloud director api")
	RootCmd.PersistentFlags().String("user", "", "Port to run Application server on")
	RootCmd.PersistentFlags().String( "password", "", "password of vcloud director api")
	RootCmd.PersistentFlags().String("org", "", "org of vcloud director api")
	RootCmd.PersistentFlags().String("verbose", "", "verbose output")
	viper.BindPFlag("url", RootCmd.PersistentFlags().Lookup("url"))
	viper.BindPFlag("user", RootCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("password", RootCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("org", RootCmd.PersistentFlags().Lookup("org"))
	viper.BindPFlag("verbose", RootCmd.PersistentFlags().Lookup("verbose"))

	viper.SetEnvPrefix("vcd") // will be uppercased automatically
	viper.AutomaticEnv()

	url := viper.GetString("url")
	if len(url) == 0 {
		fmt.Println("url has to be set, either as env var VCD_URL or as flag url")
		os.Exit(1)
	}
	user := viper.GetString("user")
	if len(user) == 0 {
		fmt.Println("user has to be set, either as env var VCD_USER or as flag user")
		os.Exit(1)
	}
	password := viper.GetString("password")
	if len(password) == 0 {
		fmt.Println("password has to be set, either as env var VCD_PASSWORD or as flag password")
		os.Exit(1)
	}
	org := viper.GetString("org")
	if len(org) == 0 {
		fmt.Println("user has to be set, either as env var VCD_ORG or as flag org")
		os.Exit(1)
	}

	fmt.Printf("VCD_URL: [%s]\n", url)
	fmt.Printf("VCD_USER: [%s]\n", user)
	fmt.Printf("VCD_PASSWORD: [%s]\n", "***************")
	fmt.Printf("VCD_ORG: [%s]\n", org)
}