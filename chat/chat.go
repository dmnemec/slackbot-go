package chat

import (
	//Local
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/dmnemec/slackbot-go/structs"
)

const (
	chatURL = "https://slack.com/api/chat."
)

/*
GetPermalink
Update
*/

// NewChatClient creates a new client with an access token
func NewChatClient(t string) *ChatClient {
	c := new(ChatClient)
	c.setToken(t)
	return c
}

//ChatClient is the way to interface with Slack's chat methods
type ChatClient struct {
	token string
}

// Change the access token for some reason
func (c *ChatClient) setToken(t string) {
	c.token = t
}

// Delete removes a message from a channel
// https://api.slack.com/methods/chat.delete
func (c *ChatClient) Delete(channelID, ts string) (res structs.DeleteResponse, err error) {
	//Build request
	reqBod := deleteStruct{
		Token:     c.token,
		Channel:   channelID,
		Timestamp: ts,
	}
	err = jsonRequest(chatURL, "delete", c.token, reqBod, &res)
	check(err)
	//Return Response
	return
}

// PostEphemeral sends a message to a channel.
// https://api.slack.com/methods/chat.postEphemeral
func (c *ChatClient) PostEphemeral(channelID, user, text string) (res structs.PostEphemeralResponse, err error) {
	//Build request
	reqBod := postEphemeralStruct{
		Token:   c.token,
		Channel: channelID,
		User:    user,
		Text:    text,
	}
	err = jsonRequest(chatURL, "postEphemeral", c.token, reqBod, &res)
	check(err)
	//Return Response
	return
}

// PostMessage sends a message to a channel.
// https://api.slack.com/methods/chat.postMessage
func (c *ChatClient) PostMessage(channelID, text string) (res structs.PostMessageResponse, err error) {
	//Build request
	reqBod := postMessageStruct{
		Token:   c.token,
		Channel: channelID,
		Text:    text,
	}
	err = jsonRequest(chatURL, "postMessage", c.token, reqBod, &res)
	check(err)
	//Return Response
	return
}

// Unfurl adds custom unfurling action to an object
// https://api.slack.com/methods/chat.unfurl
func (c *ChatClient) Unfurl(channelID, ts string, unfurls url.Values) (res structs.UnfurlResponse, err error) {
	//Build request
	reqBod := unfurlStruct{
		Token:     c.token,
		Channel:   channelID,
		Timestamp: ts,
		Unfurls:   unfurls.Encode(),
	}
	err = jsonRequest(chatURL, "unfurl", c.token, reqBod, &res)
	check(err)
	//Return Response
	return
}

// Update changes a message in Slack (ie. edit)
// https://api.slack.com/methods/chat.unupdate
func (c *ChatClient) Update(channelID, ts, text string) (res structs.UpdateResponse, err error) {
	//Build request
	reqBod := updateStruct{
		Token:     c.token,
		Channel:   channelID,
		Timestamp: ts,
		Text:      text,
	}
	err = jsonRequest(chatURL, "update", c.token, reqBod, &res)
	check(err)
	//Return Response
	return
}

// check is a space-saving function to check errors
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Creates a urlencoded request with the appropriate headers on the http request
func urlEncodedConvoClient(url, endpoint, token string, vals url.Values, output interface{}) (err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url+endpoint+"?"+vals.Encode(), nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	check(err)
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&output)
	check(err)
	return
}

func jsonRequest(url, endpoint, token string, input, output interface{}) error {
	bod, err := json.Marshal(input)
	check(err)
	req, err := http.NewRequest("POST", url+endpoint, bytes.NewBuffer(bod))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	//Send Request
	client := &http.Client{}
	res, err := client.Do(req)
	check(err)
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&output)
	check(err)
	return err
}
