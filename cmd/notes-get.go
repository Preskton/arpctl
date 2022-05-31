/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/preskton/arpctl/lib/music"
)

// notesGetCmd represents the get command
var notesGetCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%#v", music.GetNoteByGeneralSearch(cmd.Flag("name").Value.String()))
	},
}

func init() {
	notesCmd.AddCommand(notesGetCmd)

	notesGetCmd.Flags().StringP("name", "n", "", "Name of the note to get")
	notesGetCmd.MarkFlagRequired("name")
}
