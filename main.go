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
