/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
)

// i2cListCmd represents the list command
var i2cListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		buses := i2creg.All()

		for _, bus := range buses {
			log.Infof("Bus %d: %s", bus.Number, bus.Name)

			busHandle, err := i2creg.Open(bus.Name)
			if err != nil {
				log.Fatal(err)
			}
			defer busHandle.Close()

			if p, ok := busHandle.(i2c.Pins); ok {
				log.Infof("SDA: %s", p.SDA())
				log.Infof("SCL: %s", p.SCL())
			}
		}
	},
}

func init() {
	i2cCmd.AddCommand(i2cListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
