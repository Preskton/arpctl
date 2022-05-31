/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"

	"github.com/preskton/arpctl/cmd"

	log "github.com/sirupsen/logrus"
	host "periph.io/x/host/v3"
)

var banner = `
                                         █████    ████ 
                                        ░░███    ░░███ 
  ██████   ████████  ████████   ██████  ███████   ░███ 
 ░░░░░███ ░░███░░███░░███░░███ ███░░███░░░███░    ░███ 
  ███████  ░███ ░░░  ░███ ░███░███ ░░░   ░███     ░███ 
 ███░░███  ░███      ░███ ░███░███  ███  ░███ ███ ░███ 
░░████████ █████     ░███████ ░░██████   ░░█████  █████
 ░░░░░░░░ ░░░░░      ░███░░░   ░░░░░░     ░░░░░  ░░░░░ 
                     ░███                              
                     █████                             
                    ░░░░░                           
`

func main() {
	fmt.Print(banner)
	fmt.Print("\n")

	log.Debug("Initializing host...")
	host.Init()

	cmd.Execute()
}
