package main

import (
	"fmt"
	basecamp "github.com/dmnemec/slackbot-go/basecamp"
	core "github.com/dmnemec/slackbot-go/core"
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	settings := core.LoadConfig("config.json")
	var timestamp time.Time
	var err error

	// Create initial timestamp
	timestamp, err = time.Parse(time.RFC1123, settings.Last_update)
	if err != nil {
		fmt.Println("Unable to read old time.\n")
		log.Fatal(err)
	}

	// Set a waitgroup
	wg.Add(1)

	// Run all routines
	go basecamp.BasecampHostingWatcher(&timestamp, &settings)

	// Waits for all routines to finish, which they won't
	wg.Wait()
}
