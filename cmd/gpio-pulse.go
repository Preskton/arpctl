/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"time"

	"github.com/spf13/cobra"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"

	log "github.com/sirupsen/logrus"
)

// gpioPulseCmd represents the pulse command
var gpioPulseCmd = &cobra.Command{
	Use:   "pulse",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		pinName := cmd.Flag("pin").Value.String()
		duration, _ := time.ParseDuration(cmd.Flag("duration").Value.String())

		pulse(pinName, duration)
	},
}

func init() {
	gpioCmd.AddCommand(gpioPulseCmd)

	gpioPulseCmd.Flags().StringP("pin", "p", "", "Pin to pulse between high & low")
	gpioPulseCmd.MarkFlagRequired("pin")

	gpioPulseCmd.Flags().StringP("duration", "d", "2s", "Number of miliseconds to pulse the pin")
}

func pulse(pinName string, duration time.Duration) error {
	log.Info("Pulsing GPIO pin " + pinName)
	pin := gpioreg.ByName(pinName)

	pulseTicker := time.NewTicker(500 * time.Millisecond)
	durationTicker := time.NewTimer(duration)

	for pinLevel := gpio.High; ; pinLevel = !pinLevel {
		err := pin.Out(pinLevel)

		if err != nil {
			log.Warn("Failed to signal pin: " + err.Error())
		}

		select {
		case <-pulseTicker.C:
			log.Infof("Setting pin %s to %d", pinName, pinLevel)
			continue
		case <-durationTicker.C:
			log.Infof("Reached total duration for pulse after %s", duration)
			pin.Out(gpio.Low)
			return nil
		}
	}
}
