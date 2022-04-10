/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package pgk

import (
	weatherHandler "CliTaskManager/cmd/WeatherApp/handler"
	"github.com/spf13/cobra"
	"strings"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "hava",
	Short: "Istenilen yerin sicakligini verir ",
	Run: func(cmd *cobra.Command, args []string) {
		weatherHandler.GetCelciusFromLocation(strings.Join(args, " "))
	},
}

func init() {
	RootCmd.AddCommand(configCmd)
}
