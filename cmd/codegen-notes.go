/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"

	"github.com/preskton/arpctl/lib/music"
)

// codegenNotesCmd represents the notes command
var codegenNotesCmd = &cobra.Command{
	Use:   "notes",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		filename := cmd.Flag("filename").Value.String()

		generateCode(filename)
	},
}

func init() {
	codegenCmd.AddCommand(codegenNotesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// notesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// notesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	codegenNotesCmd.Flags().StringP("filename", "f", "", "Filename to use to generate structs")
	codegenNotesCmd.MarkFlagFilename("filename")
	codegenNotesCmd.MarkFlagRequired("filename")
}

func generateCode(filename string) {
	lines, err := readCsv(filename)

	if err != nil {
		log.WithError(err).WithField("filename", filename).Fatal("Failed to read csv")
		return
	}

	notes, err := csvLines2Notes(lines)

	if err != nil {
		log.WithError(err).WithField("filename", filename).Fatal("Failed to convert CSV lines to notes")
		return
	}

	fmt.Printf("%#v", notes)
}

func readCsv(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

func csvLines2Notes(lines [][]string) ([]music.Note, error) {
	notes := make([]music.Note, len(lines))

	for lineIndex, line := range lines {
		var err error

		note := music.Note{}

		note.AbsoluteIndex = uint(lineIndex)
		note.Name = line[0]
		note.Letter = line[1]
		// Skip index b/c we'll just use line number
		note.EnharmonicSharpName = line[3]
		note.EnharmonicFlatName = line[4]

		note.Frequency, err = strconv.ParseFloat(line[5], 32)
		if err != nil {
			log.WithError(err).WithField("coords", fmt.Sprintf("%d, %s", lineIndex, "line.5")).Warn("Couldn't parse a float")
		}

		note.Wavelength, err = strconv.ParseFloat(line[6], 32)
		if err != nil {
			log.WithError(err).WithField("coords", fmt.Sprintf("%d, %s", lineIndex, "line.6")).Warn("Couldn't parse a float")
		}

		castor64, err := strconv.ParseUint(line[7], 10, 16)
		if err != nil {
			log.WithError(err).WithField("coords", fmt.Sprintf("%d, %s", lineIndex, "line.7")).Warn("Couldn't parse a uint16")
		}
		note.Castor = uint16(castor64)

		pollux64, err := strconv.ParseUint(line[8], 10, 16)
		if err != nil {
			log.WithError(err).WithField("coords", fmt.Sprintf("%d, %s", lineIndex, "line.8")).Warn("Couldn't parse a uint16")
		}
		note.Pollux = uint16(pollux64)

		werkstatt64, err := strconv.ParseUint(line[9], 10, 16)
		if err != nil {
			log.WithError(err).WithField("coords", fmt.Sprintf("%d, %s", lineIndex, "line.9")).Warn("Couldn't parse a uint16")
		}
		note.Werkstatt = uint16(werkstatt64)

		notes[lineIndex] = note
	}

	return notes, nil
}
