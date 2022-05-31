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
	"github.com/preskton/arpctl/lib/music/scale"
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
	loadScaleconfig()

	log.Debug("Initializing host...")
	host.Init()

	cmd.Execute()
}

func unmarshalJson(filename string, target any) {
	fh, err := os.Open(filename)

	if err != nil {
		log.WithError(err).WithField("filename", filename).Error("Error while reading json from disk")
	}

	defer fh.Close()

	b, err := ioutil.ReadAll(fh)

	if err != nil {
		log.WithError(err).WithField("filename", filename).Error("Error while reading json from disk")
	}

	err = json.Unmarshal(b, target)

	if err != nil {
		log.WithError(err).WithField("filename", filename).Error("Error unmarhsaling json from file")
	}
}

func loadNoteconfig() {
	unmarshalJson("config/notes.json", &music.AllNotes)
}

func loadScaleconfig() {
	unmarshalJson("config/scales.json", &scale.AllScalePatterns)
}
