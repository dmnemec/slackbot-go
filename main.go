package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const hookGen = "https://hooks.slack.com/services/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"
const json = "application/json"
const basecampEventsURL = ""

func main() {
	var message string
	if len(os.Args) > 1 {
		message = strings.Join(os.Args[1:], " ")
	}

	status := postGeneral(message)
	if status != 200 {
		fmt.Println("The message was not sent")
	}

}

// Posts a regular message to the General channel in Slack
func postGeneral(payload string) int {
	var body = []byte(`{"text":"` + payload + `"}`)

	fmt.Println(string(body[:]))
	req, err := http.Post(hookGen, json, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()

	return req.StatusCode
}

// Retreives all events from Basecamp
func getEvents() {
	//TODO add "since" field for URL

	//make new client
	client := &http.Client{}

	//don't thnk I need this line
	//resp, err := http.Get(basecampEventsURL)

	// build new request
	req, err := http.NewRequest("GET", basecampEventsURL, nil)

	// add additional headers
	req.Header.Add("User-Agent", "Agent-Smith (devin.nemec@banno.com)")

	// send request
	resp, err := client.Do(req)

	// close request
	defer resp.Body.Close()

	fmt.Printf(resp.Body)
}
