package conversations

/*
This package is a Go/Golang implementation of the Conversations API for Slack.
https://api.slack.com/docs/conversations-api

Completion List
[ ] Archive
[x] Close
[x] Create
[ ] History
[ ] Info
[x] Invite
[ ] Join
[ ] Kick
[ ] Leave
[ ] List
[x] Members
[ ] Open
[x] PostMessage
[ ] Rename
[ ] Replies
[x] Set Purpose
[ ] Set Topic
[ ] Unarchive

*/

import (
	//Global
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"regexp"

	//Local
	"github.com/dmnemec/slackbot-go/structs"
)

const (
	convURL = "https://slack.com/api/conversations."
)

// NewConvoClient creates a new client with an access token
func NewConvoClient(t string) *ConvoClient {
	c := new(ConvoClient)
	c.setToken(t)
	return c
}

// The ConvoClient struct
type ConvoClient struct {
	token string
}

// Change the access token for some reason
func (c *ConvoClient) setToken(t string) {
	c.token = t
}

// Archive a conversation
// https://api.slack.com/methods/conversations.archive
func (c *ConvoClient) Archive() {
}

// Close a direct message or multi-person direct message
// https://api.slack.com/methods/conversations.close
func (c *ConvoClient) Close(name string) (res structs.CloseResponse, err error) {
	//Validate name string
	valid, err := validChannel(name)
	check(err)
	if !valid {
		return res, errors.New("invalid channel name")
	}
	//Build request
	p := url.Values{}
	p.Add("token", c.token)
	p.Add("channel", name)
	err = urlEncodedConvoClient(convURL, "close", c.token, p, &res)
	return
}

// Create initiates a public channel-based conversation
// https://api.slack.com/methods/conversations.create
func (c *ConvoClient) Create(name string, private bool, members ...string) (res structs.CreateResponse, err error) {
	//Validate name string
	valid, err := validChannel(name)
	check(err)
	if !valid {
		return res, errors.New("invalid channel name")
	}
	//Build request
	reqBod := createChannelStruct{
		Token:     c.token,
		Name:      name,
		IsPrivate: private,
		UserIds:   members,
	}
	err = jsonRequest(convURL, "create", c.token, reqBod, &res)
	check(err)
	//Return Response
	return
}

// History fetches a conversation's history of messages and events
// https://api.slack.com/methods/conversations.history
func (c *ConvoClient) History() {
}

// Info retrieve information about a conversation
// https://api.slack.com/methods/conversations.info
func (c *ConvoClient) Info(name string) (res structs.InfoResponse, err error) {
	valid, err := validChannel(name)
	check(err)
	if !valid {
		return res, errors.New("Invalid channel name")
	}
	reqBod := getInfoRequest{
		Token:   c.token,
		Channel: name,
	}
	err = jsonRequest(convURL, "info", c.token, reqBod, &res)
	check(err)
	return res, nil
}

// Invite users to a channel.
// https://api.slack.com/methods/conversations.invite
func (c *ConvoClient) Invite(name string, users ...string) (res structs.InviteResponse, err error) {
	//Validate name string
	valid, err := validChannel(name)
	check(err)
	if !valid {
		return res, errors.New("invalid channel name")
	}
	//Build request
	reqBod := inviteStruct{
		Token:   c.token,
		Channel: name,
		Users:   users,
	}
	err = jsonRequest(convURL, "invite", c.token, reqBod, &res)
	check(err)
	return res, nil
}

// Join an existing conversation.
// https://api.slack.com/methods/conversations.join
func (c *ConvoClient) Join() {
}

// Kick a user from a conversation.
// https://api.slack.com/methods/conversations.kick
func (c *ConvoClient) Kick() {
}

// Leave a conversation.
// https://api.slack.com/methods/conversations.leave
func (c *ConvoClient) Leave() {
}

// List all channels in a Slack team.
// https://api.slack.com/methods/conversations.list
func (c *ConvoClient) List() (res structs.ListResponse, err error) {
	//Build request
	p := url.Values{}
	p.Add("token", c.token)
	p.Add("exclude_archived", "true")
	p.Add("types", "public_channel")
	err = urlEncodedConvoClient(convURL, "list", c.token, p, &res)
	return
}

// Members retrieves members of a conversation.
// https://api.slack.com/methods/conversations.members
func (c *ConvoClient) Members(channelID string) (res structs.MembersResponse, err error) {
	//Build request
	p := url.Values{}
	p.Add("token", c.token)
	p.Add("channel", channelID)
	err = urlEncodedConvoClient(convURL, "members", c.token, p, &res)
	check(err)
	return
}

// Open or resumes a direct message or multi-person direct message.
// https://api.slack.com/methods/conversations.open
func (c *ConvoClient) Open() {
}

// Rename a conversation.
// https://api.slack.com/methods/conversations.rename
func (c *ConvoClient) Rename() {
}

// Replies - retrieves a thread of messages posted to a conversation
// https://api.slack.com/methods/conversations.replies
func (c *ConvoClient) Replies() {
}

// SetPurpose sets the purpose for a conversation.
// https://api.slack.com/methods/conversations.setPurpose
func (c *ConvoClient) SetPurpose(name, purpose string) (res structs.SetPurposeResponse, err error) {
	//Validate name string
	valid, err := validChannel(name)
	check(err)
	if !valid {
		return res, errors.New("invalid channel name")
	}
	//Build request
	reqBod := setPurposeStruct{
		Token:   c.token,
		Channel: name,
		Purpose: purpose,
	}
	err = jsonRequest(convURL, "setPurpose", c.token, reqBod, &res)
	check(err)
	//Return Response
	return
}

// SetTopic sets the topic for a conversation.
// https://api.slack.com/methods/conversations.setTopic
func (c *ConvoClient) SetTopic(name string, topic string) (res structs.SetTopicResponse, err error) {
	valid, err := validChannel(name)
	check(err)
	if !valid {
		return res, errors.New("Invalid Channel Name")
	}
	reqBod := setTopicStruct{
		Token:   c.token,
		Channel: name,
		Topic:   topic,
	}
	err = jsonRequest(convURL, "setTopic", c.token, reqBod, &res)
	check(err)
	return res, err
}

// Unarchive reverses conversation archival.
// https://api.slack.com/methods/conversations.unarchive
func (c *ConvoClient) Unarchive() {
}

// check is a space-saving function to check errors
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Checks if name is over 21 characters, and only contains
// lower-case letters, numbers, hyphens, and underscores
func validChannel(n string) (bool, error) {
	return regexp.MatchString("^([A-Za-z0-9-_]){1,21}$", n)
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
