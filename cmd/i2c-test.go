/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"periph.io/x/conn/v3/i2c/i2creg"

	log "github.com/sirupsen/logrus"

	"github.com/preskton/arpctl/lib/devices/adafruit/mcp4725"
)

// i2cTestCmd represents the test command
var i2cTestCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		busName := cmd.Flag("bus").Value.String()
		addressText := cmd.Flag("address").Value.String()
		value, err := strconv.Atoi(cmd.Flag("value").Value.String())

		if err != nil {
			log.WithError(err).WithField("value", cmd.Flag("value").Value.String()).Fatalf("Couldn't parse an int from the `value` flag")
		}

		address64, err := strconv.ParseUint(addressText, 16, 16)

		if err != nil {
			log.WithError(err).WithField("address", addressText).Fatalf("Couldn't parse an address from the address flag")
			return
		}

		address := uint16(address64)

		log.Infof("Setting device on bus %s, at address %s, to value %d", busName, addressText, value)

		log.WithField("busName", busName).Infof("Opening I2C bus")

		busHandle, err := i2creg.Open(busName)
		if err != nil {
			log.WithError(err).Fatalf("Failed to open I2C bus with error")
			return
		}
		defer busHandle.Close()

		log.WithField("address", address).Debug("Creating new MCP4725 device")

		dac := &mcp4725.Mcp4725{Address: address, Bus: busHandle}
		log.Debug("Preparing to call SetVoltage")
		dac.SetVoltage(2000, false, 400000)
	},
}

func init() {
	i2cCmd.AddCommand(i2cTestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	i2cTestCmd.Flags().StringP("bus", "b", "", "I2C bus with attached DAC, by default, use the system's default I2C bus")
	i2cTestCmd.Flags().StringP("address", "a", "", "Address of the DAC")
	i2cTestCmd.MarkFlagRequired("address")
	i2cTestCmd.Flags().IntP("value", "v", 0, "Value to send")
	i2cTestCmd.MarkFlagRequired("value")
}
