package core

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

const (
	hookGen = "https://hooks.slack.com/services/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"
	jsonA   = "application/json"
)

// Posts a regular message to the channel in Slack
func PostGeneral(payload string) {
	var body = []byte(`{"text":"` + payload + `"}`)

	fmt.Println(string(body[:]))
	req, err := http.Post(hookGen, jsonA, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()

	//	return req.StatusCode
}

//TODO Finish This
/*
// Posts a fancy message to the channel
func postFancy(fallback, summary, excerpt string) {
	var body = []byte(`{"attachements":[{` +
		`"fallback":"` + fallback + `",` +
		`"text":"` + summary + `",` +
		`"color":"#36a64f",` +
		`"fields": [` +
		`{` +

		`}`)

	fmt.Println(string(body[:]))
	req, err := http.Post(hookGen, jsonA, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()

	//	return req.StatusCode
}
*/
