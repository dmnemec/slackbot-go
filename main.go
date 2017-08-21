package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	const hookGen = "https://hooks.slack.com/services/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"
	const json = "application/json"
	var payload = "test message"

	if len(os.Args) > 1 {
		payload = strings.Join(os.Args[1:], " ")
	}
	var body = []byte(`{"text":"` + payload + `"}`)
	fmt.Printf(string(body[:]))
	req, err := http.Post(hookGen, json, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()
	fmt.Printf("Response Code:", req.StatusCode)
}
