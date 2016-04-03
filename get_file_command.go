package main

import (
	"bufio"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"strings"
)

const NO_FILE_ARGUMENT_MESSAGE = "Filename not specified as param in get command"

type GetFileCommand struct {
	Filename          string
	ServerURL         string
	DefaultPathToSave string
}

func NewGetFileCommand(filename string, config *Config) GetFileCommand {
	return GetFileCommand{
		Filename:          filename,
		ServerURL:         fmt.Sprintf("http://%s:%s", config.Host, config.Port),
		DefaultPathToSave: config.DefaultPathToSave,
	}
}

func (getFileCommand GetFileCommand) Do() string {
	subURL := "v1/file/" + getFileCommand.Filename
	if !strings.HasSuffix(getFileCommand.ServerURL, "/") {
		subURL = "/" + subURL
	}

	response, err := doRequest("GET", getFileCommand.ServerURL, subURL)
	if err != nil {
		log.WithFields(log.Fields{
			"serverURL": getFileCommand.ServerURL,
			"subURL":    subURL,
			"error":     err.Error(),
		}).Error("Get file do request error")
		return "Error command execution"
	}

	if response != nil && response.Body != nil {
		defer response.Body.Close()

		switch response.StatusCode {
		case http.StatusOK:
			reader := bufio.NewReader(os.Stdin)
			customPathToSave := ""
			isCorrectPathToSave := false

			for !isCorrectPathToSave {
				fmt.Printf("Enter path to folder for save, or press Enter to save in default folder (%s) -> ", getFileCommand.DefaultPathToSave)

				input, err := reader.ReadString('\n')
				if err != nil {
					log.WithField("error", err.Error()).Error("Error read custom path to save")
				} else {
					input = strings.Trim(input, "\n")
					input = strings.TrimSpace(input)
					if input == "" {
						break
					}
					if _, err := os.Stat(input); os.IsNotExist(err) {
						log.Error("Specified directory to downloading files is not exists, enter correct custom path or use default")
					} else {
						customPathToSave = input
						isCorrectPathToSave = true
					}
				}
			}

			pathToSave := getFileCommand.DefaultPathToSave

			if isCorrectPathToSave && customPathToSave != "" {
				pathToSave = customPathToSave
			}

			pathToFile := fmt.Sprintf("%s/%s", pathToSave, getFileCommand.Filename)
			out, err := os.Create(pathToFile)
			if err != nil {
				log.WithField("error", err.Error()).Error("Error create file " + pathToFile)
				return "Error command execution"
			}

			if out != nil {
				defer out.Close()
				_, err := io.Copy(out, response.Body)

				if err != nil {
					log.WithField("error", err.Error()).Error("Error save file")
					return "Can't save file " + getFileCommand.Filename
				}
				return "File downloaded to " + pathToFile
			}
		case http.StatusNotFound:
			return "File " + getFileCommand.Filename + " not exists"

		default:
			return "Can't download file " + getFileCommand.Filename
		}
	}
	return "Can't download file " + getFileCommand.Filename
}

func (getFileCommand GetFileCommand) isExitCommand() bool {
	return false
}
