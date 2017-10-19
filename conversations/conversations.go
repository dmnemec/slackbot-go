package conversations

import (
	"errors"
	"net/http"
	"net/url"
)

const (
	url = "https://slack.com/api/conversations."
)

// Creates a new client with an access token
func newClient(t string) Client {
	c := new(Client)
	c.setToke(t)
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
func (c *Client) Close() {
}

// Create
// Initiates a public channel-based conversation
// https://api.slack.com/methods/conversations.create
func (c *Client) Create(name string) (r *http.Response, e error) {
	//Validate name string
	if !validChannel(name) {
		return errors.New("Invalid channel name.")
	}
	//Build request
	h := http.Client()
	p := url.Values{}
	p.Add("token", t)
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
		return errors.New("Invalid channel name.")
	}
	//Build request
	h := http.Client()
	p := url.Values{}
	p.Add("token", t)
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
//
//
func (c *Client) Invite() {
}

// Join
//
//
func (c *Client) Join() {
}

// Kick
//
//
func (c *Client) Kick() {
}

// Leave
//
//
func (c *Client) Leave() {
}

// List
//
//
func (c *Client) List() {
}

// Members
//
//
func (c *Client) Members() {
}

// Open
//
//
func (c *Client) Open() {
}

// Rename
//
//
func (c *Client) Rename() {
}

// Replies
//
//
func (c *Client) Replies() {
}

// SetPurpose
//
//
func (c *Client) SetPurpose() {
}

// SetTopic
//
//
func (c *Client) SetTopic() {
}

// Unarchive
//
//
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
func validName(n string) bool {
	l := len(n)

	if l > 21 || l < 1 {
		return false
	} else if m, _ := regexp.MatchString("^(a-z0-9-_)+$", n); !m {
		return false
	}
	return true
}
