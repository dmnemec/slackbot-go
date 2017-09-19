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
	//	hookGen          = "https://hooks.slack.com/services/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"
	hookGen          = "https://hooks.slack.com/services/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"
	jsonA            = "application/json"
	bcmpApi          = "https://basecamp.com/"
	bcmpApiVer       = "/api/v1/"
	bcmpApiEvnt      = "events.json?since="
	bcmpApiProj      = "projects"
	userAgent        = "Agent-Smith (devin.nemec@banno.com)"
	userAgentVersion = "0.1"
	uaString         = userAgent + userAgentVersion
)

func main() {
	settings := loadConfig("config.json")
	//var message string
	var timestamp time.Time
	var err error
	timestamp, err = time.Parse(time.RFC1123, settings.Last_update)
	/*
		timebytes := []byte(settings.Last_update)
		err := timestamp.UnmarshalText(timebytes)
	*/
	if err != nil {
		fmt.Println("Unable to read old time.\n")
		log.Fatal(err)
	}

	//if len(os.Args) > 1 {
	//		message = strings.Join(os.Args[1:], " ")
	//	}

	/*status := postGeneral(message)
	if status != 200 {
		fmt.Println("The message was not sent")
	}
	*/
	for {
		timestamp = getEvents(timestamp)
		settings.Last_update = timestamp.Format(string(time.RFC1123))
		if err != nil {
			fmt.Println("Unable to store old time as string.")
			log.Fatal(err)
		}
		err = updateConfig(settings, "config.json")
		if err != nil {
			log.Fatal(err)
			fmt.Println("There was a problem updating the config file.")
			os.Exit(1)
		}
	}
}

// Loads config file
func loadConfig(filename string) Config {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var out Config
	json.Unmarshal(raw, &out)
	return out
}

// Update config file
func updateConfig(new Config, filename string) error {
	//Turn JSON into bytes for writing
	bytes, err := json.Marshal(new)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	//Overwrite old file
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer file.Close()

	//Put JSON into new file
	_, err = file.Write(bytes)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	file.Sync()

	return nil
}

// Posts a regular message to the channel in Slack
func postGeneral(payload string) {
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

// Retreives all events from Basecamp
func getEvents(sinceT time.Time) time.Time {
	//add "since" field for URL
	var since string
	if sinceT.IsZero() {
		startTime := time.Now()
		startTime = startTime.Add(-30 * time.Minute)
		sinceB, err := startTime.MarshalText()
		if err != nil {
			log.Fatal(err)
		}
		since = string(sinceB[:])
		since = strings.Trim(since, `"`)
	} else {
		sinceB, err := sinceT.MarshalText()
		if err != nil {
			log.Fatal(err)
		}
		since = string(sinceB[:])
		since = strings.Trim(since, `"`)
	}

	//make new client
	client := &http.Client{}

	// build new request
	bcmpId := os.Getenv("BASECAMP_ID")
	bcmpApiUrl := bcmpApi + bcmpId + bcmpApiVer + bcmpApiEvnt
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

	// print all retrieved events (for testing)
	// fmt.Println("Basecamp Response Status Code: " + resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	events := make([]Event, 0)
	json.Unmarshal(body, &events)
	for i := len(events) - 1; i >= 0; i-- {
		fmt.Printf("\n\nRecord %v of %v\n", i+1, len(events))
		//for r, e := range events {
		if strings.Contains(events[i].Target, "Hosting Account Questions") || strings.Contains(events[i].Target, "Hosting Account Setup Questions") {
			//if (strings.Contains(e.Target, "Hosting Account Questions") || strings.Contains(e.Target, "Hosting Account Setup Questions")) && strings.Contains(e.Action, "commented on") {
			fmt.Printf("\n\nRecord %v of %v\n", i+1, len(events))
			events[i].Print()
			postGeneral(buildPost(events[i]))
		}
	}

	//return most recent timestamp
	if len(events) > 0 {
		sincer := []byte(events[0].Created_at)
		err = sinceT.UnmarshalText(sincer)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("No events found in the last 5 seconds.")
	}
	time.Sleep(5 * time.Second)
	return sinceT
}

// Creates the message post for Slack
func buildPost(e Event) string {
	projectName := getProjectName(e.Html_url[38:46])
	fmt.Println(e.Html_url[38:45])
	body := "*<" + e.Html_url + "|" + e.Creator.Name + " " + e.Action + " " + projectName + ">*\n_" + e.Excerpt + "_"
	return body
}

// Returns Basecamp project name
func getProjectName(p string) string {
	//make new client
	client := &http.Client{}

	// build new request
	bcmpId := os.Getenv("BASECAMP_ID")
	bcmpApiUrl := bcmpApi + bcmpId + bcmpApiVer + bcmpApiProj + "/"
	req, err := http.NewRequest("GET", bcmpApiUrl+p+".json", nil)
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

	// capture and return project name
	var proj Project
	//err = json.NewDecoder(resp.Body).Decode(proj)
	//if err != nil {
	//	log.Fatal(err)
	//}
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &proj)
	return proj.Name
}
