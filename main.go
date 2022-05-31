/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/preskton/arpctl/cmd"
	"github.com/preskton/arpctl/lib/music"
	"github.com/spf13/viper"

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
	log.Debug("Setting viper global defaults")
	// TODO figure out how to hook viper to logrus
	viper.SetConfigName("arpconfig")
	viper.AddConfigPath("config")
	viper.AddConfigPath("$HOME/.arpctl")

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.WithField("filename", e.Name).Infof("Config file reloaded from disk")
	})
	viper.WatchConfig()

	hideBanner := viper.GetBool("hideBanner")

	if !hideBanner {
		fmt.Print(banner)
		fmt.Print("\n")
	}

	loadNoteconfig()

	log.Debug("Initializing host...")
	host.Init()

	cmd.Execute()
}

func loadNoteconfig() {
	fh, err := os.Open("config/notes.json")

	if err != nil {
		log.WithError(err).Error("Error while opening notes.json")
	}

	defer fh.Close()

	b, err := ioutil.ReadAll(fh)

	if err != nil {
		log.WithError(err).Error("Error while reading notes.json")
	}

	json.Unmarshal(b, &music.AllNotes)
}
