package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	hookGen          = "https://hooks.slack.com/services/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"
	jsonA            = "application/json"
	bcmpApi          = "https://basecamp.com/"
	bcmpApiVer       = "/api/v1/"
	bcmpApiEvnt      = "events.json?since="
	userAgent        = "Agent-Smith (devin.nemec@banno.com)"
	userAgentVersion = "0.1"
	uaString         = userAgent + userAgentVersion
)

func main() {
	var message string
	var timestamp time.Time
	if len(os.Args) > 1 {
		message = strings.Join(os.Args[1:], " ")
	}

	status := postGeneral(message)
	if status != 200 {
		fmt.Println("The message was not sent")
	}

	getEvents(timestamp)

}

// Posts a regular message to the General channel in Slack
func postGeneral(payload string) int {
	var body = []byte(`{"text":"` + payload + `"}`)

	fmt.Println(string(body[:]))
	req, err := http.Post(hookGen, jsonA, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()

	return req.StatusCode
}

// Retreives all events from Basecamp
func getEvents(sinceT time.Time) {
	//TODO add "since" field for URL
	var since string
	if sinceT.IsZero() {
		startTime := time.Now()
		//2012-03-24T11:00:00-06:00 RFC3339?
		//2006-01-02T15:04:05Z07:00 RFC3339
		//hour := time.Duration(1) * time.Hour
		startTime = startTime.Add(-1 * time.Hour)
		sinceB, err := startTime.MarshalText()
		if err != nil {
			log.Fatal(err)
		}
		since = string(sinceB[:])
		since = strings.Trim(since, `"`)
	}

	//since = "2017-08-22T13:59:13.000-06:00"
	//since = "2017-08-22T13:34:20.245234-06:00"

	//make new client
	client := &http.Client{}

	//don't thnk I need this line
	//resp, err := http.Get(basecampEventsURL)
	bcmpId := os.Getenv("BASECAMP_ID")
	bcmpApiUrl := bcmpApi + bcmpId + bcmpApiVer + bcmpApiEvnt

	fmt.Println(bcmpApiUrl + since)

	// build new request
	req, err := http.NewRequest("GET", bcmpApiUrl+since, nil)
	if err != nil {
		log.Fatal(err)
	}

	// add basic auth
	userBcmp := os.Getenv("BASECAMP_USER")
	passBcmp := os.Getenv("BASECAMP_PASS")
	req.SetBasicAuth(userBcmp, passBcmp)

	// add additional headers
	req.Header.Add("User-Agent", uaString)
	req.Header.Add("Content-Type", jsonA)

	// send request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// close request
	defer resp.Body.Close()

	// print the first event (for testing)
	fmt.Println("Basecamp Response Status Code: " + resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	events := make([]Event, 0)
	json.Unmarshal(body, &events)
	for r, e := range events {
		fmt.Printf("\n\nRecord %v of %v\n", r+1, len(events))
		e.Print()

	}
}
