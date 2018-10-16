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
	ugcURL = "https://slack.com/api/usergroups.users."
)

// NewUGCClient creates a new client with an access token
func NewUGCClient(t string) *UGCClient {
	c := new(UGCClient)
	c.setToken(t)
	return c
}

//UGCClient is the way to interface with Slack's chat methods
type UGCClient struct {
	token string
}

// Change the access token for some reason
func (c *UGCClient) setToken(t string) {
	c.token = t
}

// Delete removes a message from a channel
// https://api.slack.com/methods/chat.delete
func (c *UGCClient) Delete(channelID, ts string) (res structs.DeleteResponse, err error) {
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

// GetPermalink retrieves a permanent link to a message in Slack
// https://api.slack.com/methods/chat.getPermalink
func (c *UGCClient) GetPermalink(channelID, mts string) (res structs.GetPermalinkResponse, err error) {
	//Build Request
	p := url.Values{}
	p.Add("token", c.token)
	p.Add("channel", channelID)
	p.Add("message_ts", mts)
	//Return response
	err = urlEncodedClient(chatURL, "getPermalink", c.token, p, &res)
	return
}

// check is a space-saving function to check errors
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Creates a urlencoded request with the appropriate headers on the http request
func urlEncodedClient(url, endpoint, token string, vals url.Values, output interface{}) (err error) {
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
	check(err)
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
