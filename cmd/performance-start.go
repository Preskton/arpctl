/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/conn/v3/physic"

	log "github.com/sirupsen/logrus"

	"github.com/preskton/arpctl/lib/devices/adafruit/mcp4725"
	"github.com/preskton/arpctl/lib/music"
	"github.com/preskton/arpctl/lib/music/scale"

	"github.com/eiannone/keyboard"
)

var arpMode = false

type PerformanceParameters struct {
	BusName       string
	BusFrequency  physic.Frequency
	DeviceAddress uint16
	VoltageValue  int16
	TotalDuration time.Duration
	VoltageStep   int16
	StepDuration  time.Duration
	Scale         *scale.ScalePattern
	RootNote      *music.Note
	BPM           int
}

// performanceStartCmd represents the test command
var performanceStartCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		p, err := parsePerformanceParameters(cmd)
		if err != nil {
			log.WithError(err).Fatalf("Failed to parse test parameters from command line arguments")
			return
		}

		err = perform(p)
		if err != nil {
			log.WithError(err).Errorf("Something didn't work while running the demo")
		}

		return
	},
}

func init() {
	performanceCmd.AddCommand(performanceStartCmd)

	performanceStartCmd.Flags().StringP("bus", "b", "", "I2C bus with attached DAC, by default, use the system's default I2C bus")

	performanceStartCmd.Flags().StringP("address", "a", "", "Address of the DAC")
	performanceStartCmd.MarkFlagRequired("address")

	performanceStartCmd.Flags().IntP("startingVoltage", "v", 0, "Starting voltage value of the test")

	performanceStartCmd.Flags().StringP("duration", "d", "10s", "Total duration of the test")

	performanceStartCmd.Flags().IntP("step", "s", 250, "Size of each step during the test")

	performanceStartCmd.Flags().String("stepDuration", "500ms", "Duration of each note during the test")

	performanceStartCmd.Flags().String("scale", "Major thirds w/ octave", "Scale to use for arpin'")

	performanceStartCmd.Flags().StringP("root", "r", "A2", "Root note of the arp")

	performanceStartCmd.Flags().Int("bpm", 120, "Beats per minute, used to calc note length")

	// TODO note type
}

func parsePerformanceParameters(cmd *cobra.Command) (*PerformanceParameters, error) {
	p := &PerformanceParameters{}

	busName := cmd.Flag("bus").Value.String()
	p.BusName = busName

	startingVoltageText := cmd.Flag("startingVoltage").Value.String()

	startingVoltageBase := 10
	if strings.HasPrefix(startingVoltageText, "0x") {
		startingVoltageBase = 16
	}

	startingVoltage64, err := strconv.ParseUint(strings.Replace(startingVoltageText, "0x", "", 1), startingVoltageBase, 16)
	if err != nil {
		log.WithError(err).WithField("startingVoltage", startingVoltageText).Fatalf("Couldn't parse an int16 from the startingVoltage flag")
		return nil, fmt.Errorf("Couldn't parse startingVoltage flag: %w", err)
	}
	p.VoltageValue = int16(startingVoltage64)

	addressText := cmd.Flag("address").Value.String()
	address64, err := strconv.ParseUint(strings.Replace(addressText, "0x", "", 1), 16, 16)
	if err != nil {
		log.WithError(err).WithField("address", addressText).Fatalf("Couldn't parse an address from the address flag")
		return nil, fmt.Errorf("Couldn't parse value flag: %w", err)
	}
	p.DeviceAddress = uint16(address64)

	p.BusFrequency = 400 * physic.KiloHertz

	durationText := cmd.Flag("duration").Value.String()
	p.TotalDuration, err = time.ParseDuration(durationText)
	if err != nil {
		log.WithError(err).WithField("duration", durationText).Fatalf("Couldn't parse an address from the duration flag")
		return nil, fmt.Errorf("Couldn't parse duration flag: %w", err)
	}

	stepText := cmd.Flag("step").Value.String()
	step64, err := strconv.ParseInt(stepText, 10, 16)
	if err != nil {
		log.WithError(err).WithField("step", stepText).Fatalf("Couldn't parse an int16 from the step flag")
		return nil, fmt.Errorf("Couldn't parse step flag: %w", err)
	}
	p.VoltageStep = int16(step64)

	stepDurationText := cmd.Flag("stepDuration").Value.String()
	p.StepDuration, err = time.ParseDuration(stepDurationText)
	if err != nil {
		log.WithError(err).WithField("duration", stepDurationText).Fatalf("Couldn't parse an address from the stepDuration flag")
		return nil, fmt.Errorf("Couldn't parse duration flag: %w", err)
	}

	noteName := cmd.Flag("root").Value.String()
	note := music.GetNoteByName(noteName)
	if note == nil {
		log.WithField("noteName", noteName).Error("No note with matching name found")
	}
	p.RootNote = note

	scaleText := cmd.Flag("scale").Value.String()
	scale := scale.GetScaleByName(scaleText)
	if scale == nil {
		log.WithField("scaleName", scaleText).Errorf("No scale with matching name found")
	}
	p.Scale = scale

	log.Error(p.Scale.Name)

	return p, nil
}

func perform(p *PerformanceParameters) error {
	log.Infof("Setting device on bus %s, at address 0x%x, to value %d", p.BusName, p.DeviceAddress, p.VoltageValue)
	// TODO handle this as part of MCP4725, consumer shouldn't know
	log.WithField("busName", p.BusName).Infof("Opening I2C bus")
	busHandle, err := i2creg.Open(p.BusName)
	if err != nil {
		log.WithError(err).Fatalf("Failed to open I2C bus with error")
		return fmt.Errorf("Failed to open IC2 bus: %w", err)
	}
	defer busHandle.Close()

	log.WithField("address", fmt.Sprintf("0x%x", p.DeviceAddress)).Debug("Creating new MCP4725 device")
	dac := &mcp4725.Mcp4725{Address: uint16(p.DeviceAddress), Bus: busHandle}

	keypressChannel, err := keyboard.GetKeys(2)
	if err != nil {
		log.WithError(err).Fatalf("Failed to open keyboard handler")
		return fmt.Errorf("Failed to open keyboard handler: %w", err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	pulseTicker := time.NewTicker(p.StepDuration)
	durationTicker := time.NewTimer(p.TotalDuration)

	currentVoltageStep := p.VoltageValue
	direction := int16(1)

	pc := music.PatternContext{
		RootNote:     p.RootNote,
		NextNote:     p.RootNote,
		Scale:        p.Scale,
		PatternIndex: 0,
	}

	for {
		select {
		case event := <-keypressChannel:
			if event.Err != nil {
				log.WithError(event.Err).Error("Unexpected error while handling keypress")
			}
			log.WithFields(log.Fields{"rune": fmt.Sprintf("%q", event.Rune), "key": fmt.Sprintf("0x%x", event.Key)}).Debug("Keypress detected")
			if event.Key == keyboard.KeyEsc {
				return nil
			} else if event.Key == keyboard.KeySpace {
				pulseTicker.Stop()
				durationTicker.Stop()
				log.Info("Entering experimentation mode - all tickers stopped")
			} else if event.Key == keyboard.KeyPgdn {
				currentVoltageStep = mcp4725.MinRawVoltage
				setDacVoltage(dac, currentVoltageStep, p.BusFrequency)
			} else if event.Key == keyboard.KeyPgup {
				currentVoltageStep = mcp4725.MaxRawVoltage
				setDacVoltage(dac, currentVoltageStep, p.BusFrequency)
			} else if string(event.Rune) == "w" {
				// major adjustment up
				currentVoltageStep, direction = getNextVoltage(currentVoltageStep, p.VoltageStep, 1)
				setDacVoltage(dac, currentVoltageStep, p.BusFrequency)
			} else if string(event.Rune) == "s" {
				// major adjustment down
				currentVoltageStep, direction = getNextVoltage(currentVoltageStep, p.VoltageStep, -1)
				setDacVoltage(dac, currentVoltageStep, p.BusFrequency)
			} else if string(event.Rune) == "a" {
				// minor adjustment down
				currentVoltageStep, direction = getNextVoltage(currentVoltageStep, -1, 1)
				setDacVoltage(dac, currentVoltageStep, p.BusFrequency)
			} else if string(event.Rune) == "d" {
				// minor adjustment up
				currentVoltageStep, direction = getNextVoltage(currentVoltageStep, 1, 1)
				setDacVoltage(dac, currentVoltageStep, p.BusFrequency)
			} else if string(event.Rune) == "q" {
				// minor adjustment down
				currentVoltageStep, direction = getNextVoltage(currentVoltageStep, -10, 1)
				setDacVoltage(dac, currentVoltageStep, p.BusFrequency)
			} else if string(event.Rune) == "e" {
				// minor adjustment up
				currentVoltageStep, direction = getNextVoltage(currentVoltageStep, 10, 1)
				setDacVoltage(dac, currentVoltageStep, p.BusFrequency)
			} else if string(event.Rune) == "r" {
				// arp mode
				arpMode = !arpMode

				if arpMode {
					log.Info("Starting arp mode")
					log.Infof("%#v", pc)
				} else {
					log.Info("Leaving arp mode")
				}
			}
		case <-pulseTicker.C:
			if arpMode {
				setDacVoltage(dac, int16(pc.NextNote.Castor), p.BusFrequency)
				pc.Advance()
			} else {
				setDacVoltage(dac, currentVoltageStep, p.BusFrequency)
				currentVoltageStep, direction = getNextVoltage(currentVoltageStep, p.VoltageStep, direction)
			}

			continue
		case <-durationTicker.C:
			log.WithField("duration", p.TotalDuration).Infof("Test complete!")
			return nil
		}
	}
}

func getNextVoltage(currentVoltage int16, voltageStep int16, direction int16) (int16, int16) {
	nextVoltageStep := currentVoltage + (voltageStep * direction)

	if nextVoltageStep > mcp4725.MaxRawVoltage {
		log.WithField("nextStep", nextVoltageStep).Warn("Next voltage step would exceed max voltage, flipping direction -")
		direction = -1
		nextVoltageStep = currentVoltage + (voltageStep * direction)
	} else if nextVoltageStep < mcp4725.MinRawVoltage {
		log.WithField("nextStep", nextVoltageStep).Warn("Next voltage step would go below min voltage, flipping direction +")
		direction = 1
		nextVoltageStep = currentVoltage + (voltageStep * direction)
	}
	return nextVoltageStep, direction
}

func setDacVoltage(d *mcp4725.Mcp4725, v int16, f physic.Frequency) {
	log.WithField("voltage", v).Info("Setting voltage")
	err := d.SetVoltage(v, false, f)
	if err != nil {
		log.WithError(err).Error("Failed to set voltage on DAC")
	}
}
