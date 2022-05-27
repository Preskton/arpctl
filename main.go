package main

import (
	"fmt"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/conn/v3/pin/pinreg"
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

	p := gpioreg.ByName("16")
	t := time.NewTicker(500 * time.Millisecond)
	for l := gpio.High; ; l = !l {
		p.Out(l)
		<-t.C
	}
}

func printGpio(invalid, showFunctions bool) {
	all := gpioreg.All()
	for _, p := range all {
		fmt.Print(p.String() + " - " + p.Function())

		if pinreg.IsConnected(p) {
			fmt.Print(" (connected)")
		}

		fmt.Print("\n")
	}
}
