/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/conn/v3/pin/pinreg"

	log "github.com/sirupsen/logrus"
)

// gpioListCmd represents the list command
var gpioListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Listing all GPIO pins")
		printGpio(false, false)
	},
}

func init() {
	gpioCmd.AddCommand(gpioListCmd)
}

func printGpio(invalid, showFunctions bool) {
	all := gpioreg.All()
	for _, p := range all {
		fmt.Print(p.String() + " - " + p.Function())

		if pinreg.IsConnected(p) {
			fmt.Print(" (connected)")
		}

		fmt.Print("\n")
	}
}
