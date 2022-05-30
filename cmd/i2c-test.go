/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

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
		valueText := cmd.Flag("value").Value.String()

		valueBase := 10
		if strings.HasPrefix(valueText, "0x") {
			valueBase = 16
		}

		value64, err := strconv.ParseUint(strings.Replace(valueText, "0x", "", 1), valueBase, 16)
		if err != nil {
			log.WithError(err).WithField("value", valueText).Fatalf("Couldn't parse an int16 from the `value` flag")
		}
		value := int16(value64)

		address64, err := strconv.ParseUint(strings.Replace(addressText, "0x", "", 1), 16, 16)
		if err != nil {
			log.WithError(err).WithField("address", addressText).Fatalf("Couldn't parse an address from the address flag")
			return
		}
		address := uint16(address64)

		log.Infof("Setting device on bus %s, at address %x, to value %d", busName, fmt.Sprintf("0x%x", address), value)

		log.WithField("busName", busName).Infof("Opening I2C bus")
		busHandle, err := i2creg.Open(busName)
		if err != nil {
			log.WithError(err).Fatalf("Failed to open I2C bus with error")
			return
		}
		defer busHandle.Close()

		log.WithField("address", fmt.Sprintf("0x%x", address)).Debug("Creating new MCP4725 device")
		dac := &mcp4725.Mcp4725{Address: address, Bus: busHandle}

		log.Debug("Preparing to call SetVoltage")
		dac.SetVoltage(value, false, 400000)
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

	i2cTestCmd.Flags().StringP("duration", "d", "10s", "Total duration of the test")

	i2cTestCmd.Flags().IntP("step", "s", 250, "Size of each step during the test")

	i2cTestCmd.Flags().StringP("stepDuration", "sd", "500", "Duration of each note during the test")
}
