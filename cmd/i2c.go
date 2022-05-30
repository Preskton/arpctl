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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// i2cCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// i2cCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
