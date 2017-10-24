package basecamp

import (
	"encoding/json"
	"fmt"
	core "github.com/dmnemec/slackbot-go/core"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	bcmpApi          = "https://basecamp.com/"
	bcmpApiVer       = "/api/v1/"
	bcmpApiEvnt      = "events.json?since="
	bcmpApiProj      = "projects"
	jsonA            = "application/json"
	userAgent        = "Slackbot-Go (github.com/dmnemec/slackbot-go)"
	userAgentVersion = " v0.2"
	uaString         = userAgent + userAgentVersion
)

// Pulls events from Basecamp and do something with them
// This is an example function showing a way to implement these packages
// It will continuously pull events from Basecamp and perform a function on them
func BasecampWatcher(timestamp *time.Time, settings *core.Config) {
	var err error
	for {
		timestamp = ProcessEvents(timestamp, func(e []Event) {
			for _, i := range e {
				i.Print()
			}
		})
		settings.Last_update = timestamp.Format(string(time.RFC1123))
		if err != nil {
			fmt.Println("Unable to store old time as string.")
			log.Fatal(err)
		}
		err = core.UpdateConfig(settings, os.Args[1])
		if err != nil {
			log.Fatal(err)
			fmt.Println("There was a problem updating the config file.")
			os.Exit(1)
		}
	}
}

// Retreives all events from Basecamp and performs a function on them
func ProcessEvents(sinceT *time.Time, p func([]Event)) *time.Time {
	// create events slice
	events := make([]Event, 0)
	var err error

	// add "since" field for URL
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

	// Get array of events
	GetBasecampList(since, &events)

	// Process events
	p(events)

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

// Return a list of events
func GetBasecampList(since string, events *[]Event) error {
	bcmpId := os.Getenv("BASECAMP_ID")
	userBcmp := os.Getenv("BASECAMP_USER")
	passBcmp := os.Getenv("BASECAMP_PASS")
	page := 1
	rem := 50
	pull := make([]Event, 0)

	// make new client
	client := &http.Client{}

	//events := make([]Event, 0)
	for rem == 50 {
		// build new request
		bcmpApiUrl := bcmpApi + bcmpId + bcmpApiVer + bcmpApiEvnt + since
		if page > 1 {
			pageString := strconv.Itoa(page)
			bcmpApiUrl = bcmpApiUrl + "&page=" + pageString
		}

		req, err := http.NewRequest("GET", bcmpApiUrl, nil)
		if err != nil {
			log.Fatal(err)
		}
		// add basic auth
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
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(body, &pull)
		*events = append(*events, pull...)
		rem = len(pull)
		page = page + 1
	}

	return nil
}

// Returns Basecamp project name
func GetProjectName(p string) string {
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
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &proj)
	return proj.Name
}

// Retreives all events from Basecamp since a time
func GetNewEvents(sinceT *time.Time) (*time.Time, []Event) {
	// create events slice
	events := make([]Event, 0)
	var err error

	// add "since" field for URL
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

	// Return list of events
	GetBasecampList(since, &events)

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
	return sinceT, events
}
