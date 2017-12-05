package core

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

const (
	jsonA = "application/json"
)

// Posts a regular message to the channel in Slack
func PostChannel(payload, name string, config Config) {
	//var body = []byte(`{"text":"` + payload + `"}`)
	var body = []byte(`{"link_names":1,"parse":"full","text":"` + payload + `"}`)

	fmt.Println(string(body[:]))
	url := GetHook(config, name)
	req, err := http.Post(url, jsonA, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()
}

func PostReply(payload, responseUrl string) {
	var body = []byte(`{"link_names":1,"parse":"full","text":"` + payload + `"}`)

	fmt.Println(string(body[:]))
	req, err := http.Post(responseUrl, jsonA, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	defer req.Body.Close()

	//	return req.StatusCode
}

/*
//TODO Finish This
// Posts a fancy message to the channel
func PostFancy(summary, excerpt, color string) {
	color = "#36a64f"
	message := AttachmentMessage{}
	var message.Attachments [1]AttachmentObject
	message.Attachments[0].Fallback = excerpt
	message.Attachments[0].Text = excerpt
	message.Attachments[0].Color = color //TODO create custom color maker
	message.Attachments[0].Pretext = "Definately not self-aware"
	message.Attachments[0].Title = summary
	message.Attachments[0].Title_Link = ""

	body, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}
	/*
		var body = []byte(`{"attachements":[{` +
			`"fallback":"` + fallback + `",` +
			`"text":"` + summary + `",` +
			`"color":"#36a64f",` +
			`"fields": [` +
			`{` +

			`}`)
*/
/*
	fmt.Println(string(body[:]))
	req, err := http.Post(hookGen, jsonA, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()

	//	return req.StatusCode
}*/
