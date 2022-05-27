package main

import (
	"fmt"

	"periph.io/x/conn/v3/gpio/gpioreg"
	host "periph.io/x/host/v3"

	log "github.com/sirupsen/logrus"
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

	log.Info("Initializing host...")
	host.Init()

	log.Info("Listing all GPIO pins")

	printGpio(false, false)
}

func printGpio(invalid, showFunctions bool) {
	all := gpioreg.All()
	for _, p := range all {
		print(p.String() + "\n")
	}
}
