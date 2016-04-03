package main

import (
	"bufio"
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
	"strings"
)

func main() {

	pathToConfig := flag.String("config", "", "Config not found")
	flag.Parse()

	if *pathToConfig == "" {
		log.Fatal("Config not specified")
	}

	config := ReadConfig(*pathToConfig)

	if config != nil {
		if _, err := os.Stat(config.DefaultPathToSave); os.IsNotExist(err) {
			log.Fatal("Specified directory to downloading files is not exists, change default-path-to-save in " + *pathToConfig)
		}

		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("-> ")
			input, err := reader.ReadString('\n')
			if err != nil {
				log.WithField("error", err.Error()).Error("Read command error")
			} else {
				input = strings.Trim(input, "\n")

				command := makeCommand(input, config)
				fmt.Println(command.Do())

				if command.isExitCommand() {
					os.Exit(0)
				}
			}
		}
	} else {
		log.Fatal("Handling config error!")
	}
}
