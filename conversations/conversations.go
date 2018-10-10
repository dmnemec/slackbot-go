package conversations

/*
This package is a Go/Golang implementation of the Conversations API for Slack.
https://api.slack.com/docs/conversations-api

Completion List
[ ] Archive
[x] Close
[x] Create
[x] Create Private
[ ] History
[ ] Info
[x] Invite
[ ] Join
[ ] Kick
[ ] Leave
[ ] List
[ ] Members
[ ] Open
[ ] Rename
[ ] Replies
[x] Set Purpose
[ ] Set Topic
[ ] Unarchive

*/

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

const (
	convURL = "https://slack.com/api/conversations."
)

// NewClient creates a new client with an access token
func NewClient(t string) *Client {
	c := new(Client)
	c.setToken(t)
	return c
}

// The Client struct
type Client struct {
	token string
}

// Change the access token for some reason
func (c *Client) setToken(t string) {
	c.token = t
}

// Archive a conversation
// https://api.slack.com/methods/conversations.archive
func (c *Client) Archive() {
}

// Close a direct message or multi-person direct message
// https://api.slack.com/methods/conversations.close
func (c *Client) Close(name string) (res *http.Response, err error) {
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
	//Send Request
	res, err = http.PostForm(convURL+"close", p)
	check(err)
	//Return Response
	return res, nil
}

// Create initiates a public channel-based conversation
// https://api.slack.com/methods/conversations.create
func (c *Client) Create(name string, private bool, members ...string) (res *http.Response, err error) {
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
	bod, err := json.Marshal(reqBod)
	check(err)
	req, err := http.NewRequest("POST", convURL+"create", bytes.NewBuffer(bod))
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")
	//Send Request
	client := &http.Client{}
	res, err = client.Do(req)
	check(err)
	//Return Response
	return res, nil
}

// History fetches a conversation's history of messages and events
// https://api.slack.com/methods/conversations.history
func (c *Client) History() {
}

// Info retrieve information about a conversation
// https://api.slack.com/methods/conversations.info
func (c *Client) Info() {
}

// Invite users to a channel.
// https://api.slack.com/methods/conversations.invite
func (c *Client) Invite(name string, users ...string) (res *http.Response, err error) {
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
	p.Add("users", strings.Join(users, ","))
	//Send Request
	res, err = http.PostForm(convURL+"create", p)
	check(err)
	//Return Response
	return res, nil
}

// Join an existing conversation.
// https://api.slack.com/methods/conversations.join
func (c *Client) Join() {
}

// Kick a user from a conversation.
// https://api.slack.com/methods/conversations.kick
func (c *Client) Kick() {
}

// Leave a conversation.
// https://api.slack.com/methods/conversations.leave
func (c *Client) Leave() {
}

// List all channels in a Slack team.
// https://api.slack.com/methods/conversations.list
func (c *Client) List() {
}

// Members retrieves members of a conversation.
// https://api.slack.com/methods/conversations.members
func (c *Client) Members(channelID string) (res *http.Response, err error) {
	//Validate name string
	valid, err := validChannel(channelID)
	check(err)
	if !valid {
		return res, errors.New("invalid channel name")
	}
	//Build request
	client := &http.Client{}
	p := url.Values{}
	p.Add("token", c.token)
	p.Add("channel", channelID)
	req, err := http.NewRequest("GET", convURL+"members?"+p.Encode(), nil)
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//Send Request
	res, err = client.Do(req)
	check(err)
	//Return Response
	return res, nil
}

// Open or resumes a direct message or multi-person direct message.
// https://api.slack.com/methods/conversations.open
func (c *Client) Open() {
}

// Rename a conversation.
// https://api.slack.com/methods/conversations.rename
func (c *Client) Rename() {
}

// Replies - retrieves a thread of messages posted to a conversation
// https://api.slack.com/methods/conversations.replies
func (c *Client) Replies() {
}

// SetPurpose sets the purpose for a conversation.
// https://api.slack.com/methods/conversations.setPurpose
func (c *Client) SetPurpose(name, purpose string) (res *http.Response, err error) {
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
	bod, err := json.Marshal(reqBod)
	check(err)
	req, err := http.NewRequest("POST", convURL+"setpurpose", bytes.NewBuffer(bod))
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")
	//Send Request
	client := &http.Client{}
	res, err = client.Do(req)
	check(err)
	//Return Response
	return res, nil
}

// SetTopic sets the topic for a conversation.
// https://api.slack.com/methods/conversations.setTopic
func (c *Client) SetTopic() {
}

// Unarchive reverses conversation archival.
// https://api.slack.com/methods/conversations.unarchive
func (c *Client) Unarchive() {
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
	return regexp.MatchString("^([a-z0-9-_]){1,21}$", n)
}
