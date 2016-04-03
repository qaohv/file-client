package main

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
)

func doRequest(method, url, subURL string) (*http.Response, error) {
	request, err := http.NewRequest(method, url+subURL, nil)
	if err != nil {
		log.WithFields(log.Fields{
			"method": method,
			"url":    url,
			"subURL": subURL,
			"error":  err.Error(),
		}).Error("New request error")
		return nil, err
	}

	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		log.WithFields(log.Fields{
			"method": method,
			"url":    url,
			"subURL": subURL,
			"error":  err.Error(),
		}).Error("Do request error")
		return nil, err
	}

	return response, nil
}
