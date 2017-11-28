package main

import (
	"github.com/floriankammermann/vcloud-cli/cmd"
	"os"
	"fmt"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
