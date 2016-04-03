package main

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"strings"
)

type ListFilesCommand struct {
	ServerURL string
}

func NewListFilesCommand(config *Config) ListFilesCommand {
	return ListFilesCommand{
		ServerURL: fmt.Sprintf("http://%s:%s", config.Host, config.Port),
	}
}

type ListFilesResult struct {
	Files []string `json:"files"`
}

func (listFilesCommand ListFilesCommand) Do() string {
	subURL := "v1/files"
	if !strings.HasSuffix(listFilesCommand.ServerURL, "/") {
		subURL = "/" + subURL
	}

	response, err := doRequest("GET", listFilesCommand.ServerURL, subURL)
	if err != nil {
		log.WithFields(log.Fields{
			"serverURL": listFilesCommand.ServerURL,
			"subURL":    subURL,
			"error":     err.Error(),
		}).Error("List files do request error")
		return "Error command execution"
	}

	result := new(ListFilesResult)

	if response != nil && response.Body != nil {
		defer response.Body.Close()

		if err := json.NewDecoder(response.Body).Decode(result); err != nil {
			log.WithField("error", err.Error()).Error("Encode body error")
			return "Error command execution"
		}
		prettyList := ""

		for index, fileName := range result.Files {
			prettyList += fileName
			if index != len(result.Files)-1 {
				prettyList += "\n"
			}
		}
		return prettyList
	} else {
		log.Error("Response is nil")
	}

	return ""

}

func (listFilesCommand ListFilesCommand) isExitCommand() bool {
	return false
}
