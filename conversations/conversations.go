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
[ ] Set Purpose
[ ] Set Topic
[ ] Unarchive

*/

import (
	"errors"
	"log"
	"net/http"
	"regexp"
	"strings"
)

const (
	url = "https://slack.com/api/conversations."
)

// Creates a new client with an access token
func newClient(t string) Client {
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

// Archive
// Archives a conversation
// https://api.slack.com/methods/conversations.archive
func (c *Client) Archive() {
}

// Close
// Closes a direct message or multi-person direct message
// https://api.slack.com/methods/conversations.close
func (c *Client) Close(name string) (r *http.Response, e error) {
	//Validate name string
	if !validChannel(name) {
		return nil, errors.New("invalid channel name")
	}
	//Build request
	h := http.Client()
	p := url.Values{}
	p.Add("token", c.token)
	p.Add("channel", name)
	req, err := http.PostForm(url+"close", p)
	check(err)
	//Send Request
	res, e := h.Do(req)
	check(err)
	//Return Response
	return res, nil
}

// Create
// Initiates a public channel-based conversation
// https://api.slack.com/methods/conversations.create
func (c *Client) Create(name string) (r *http.Response, e error) {
	//Validate name string
	if !validChannel(name) {
		return nil, errors.New("invalid channel name")
	}
	//Build request
	h := http.Client()
	p := url.Values{}
	p.Add("token", c.token)
	p.Add("name", name)
	req, err := http.PostForm(url+"create", p)
	check(err)
	//Send Request
	res, e := h.Do(req)
	check(err)
	//Return Response
	return res, nil
}

// CreatePrivate
// Initiates a private channel-based conversation
// https://api.slack.com/methods/conversations.create
func (c *Client) CreatePrivate(name string) (r *http.Response, e error) {
	//Validate name string
	if !validChannel(name) {
		return nil, errors.New("invalid channel name")
	}
	//Build request
	h := http.Client()
	p := url.Values{}
	p.Add("token", c.token)
	p.Add("name", name)
	p.Add("is_private", "true")
	req, err := http.PostForm(url+"create", p)
	check(err)
	//Send Request
	res, e := h.Do(req)
	check(err)
	//Return Response
	return res, nil
}

// History
// Fetches a conversation's history of messages and events
// https://api.slack.com/methods/conversations.history
func (c *Client) History() {
}

// Info
// Retrieve information about a conversation
// https://api.slack.com/methods/conversations.info
func (c *Client) Info() {
}

// Invite
// Invites users to a channel.
// https://api.slack.com/methods/conversations.invite
func (c *Client) Invite(name string, users ...string) (r *http.Response, e error) {
	//Validate name string
	if !validChannel(name) {
		return nil, errors.New("Invalid channel name.")
	}
	//Build request
	h := http.Client()
	p := url.Values{}
	p.Add("token", c.token)
	p.Add("channel", name)
	p.Add("users", strings.Join(users, ","))
	req, err := http.PostForm(url+"create", p)
	check(err)
	//Send Request
	res, e := h.Do(req)
	check(err)
	//Return Response
	return res, nil
}

// Join
// Joins an existing conversation.
// https://api.slack.com/methods/conversations.join
func (c *Client) Join() {
}

// Kick
// Removes a user from a conversation.
// https://api.slack.com/methods/conversations.kick
func (c *Client) Kick() {
}

// Leave
// Leaves a conversation.
// https://api.slack.com/methods/conversations.leave
func (c *Client) Leave() {
}

// List
// Lists all channels in a Slack team.
// https://api.slack.com/methods/conversations.list
func (c *Client) List() {
}

// Members
// Retrieve members of a conversation.
// https://api.slack.com/methods/conversations.members
func (c *Client) Members(channelID string) (*http.Response, error) {
	var res = new(http.Response)
	//Validate name string
	if !validChannel(channelID) {
		return res, errors.New("Invalid channel name")
	}
	//Build request
	h := http.Client
	p := url.Values{}
	p.Add("token", c.token)
	p.Add("channel", channelID)
	req, err := http.Get(url + "members?" + p.Encode())
	check(err)
	//Send Request
	res, e := h.Do(req)
	check(err)
	//Return Response
	return res, nil
}

// Open
// Opens or resumes a direct message or multi-person direct message.
// https://api.slack.com/methods/conversations.open
func (c *Client) Open() {
}

// Rename
// Renames a conversation.
// https://api.slack.com/methods/conversations.rename
func (c *Client) Rename() {
}

// Replies
// Retrieve a thread of messages posted to a conversation
// https://api.slack.com/methods/conversations.replies
func (c *Client) Replies() {
}

// SetPurpose
// Sets the purpose for a conversation.
// https://api.slack.com/methods/conversations.setPurpose
func (c *Client) SetPurpose() {
}

// SetTopic
// Sets the topic for a conversation.
// https://api.slack.com/methods/conversations.setTopic
func (c *Client) SetTopic() {
}

// Unarchive
// Reverses conversation archival.
// https://api.slack.com/methods/conversations.unarchive
func (c *Client) Unarchive() {
}

// Space-saving function to check errors
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Checks if name is over 21 characters, and only contains
// lower-case letters, numbers, hyphens, and underscores
func validChannel(n string) bool {
	return regexp.MatchString("^(a-z0-9-_){1,21}$", n), _
}
