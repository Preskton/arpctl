/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// i2cCmd represents the i2c command
var i2cCmd = &cobra.Command{
	Use:   "i2c",
	Short: "Manage I²C devices",
	Long:  `TODO write this`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("i2c called")
	},
}

func init() {
	rootCmd.AddCommand(i2cCmd)
}
