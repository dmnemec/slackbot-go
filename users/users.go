package users

import (
	//Local

	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/dmnemec/slackbot-go/structs"
)

const (
	ugURL = "https://slack.com/api/users."
)

// NewUsersClient creates a new client with an access token
func NewUsersClient(t string) *UsersClient {
	c := new(UsersClient)
	c.setToken(t)
	return c
}

//UsersClient is the way to interface with Slack's chat methods
type UsersClient struct {
	token string
}

// Change the access token for some reason
func (c *UsersClient) setToken(t string) {
	c.token = t
}

// GetUserByEmail retrieves a user, given an email.
// https://api.slack.com/methods/usergroups.users.list
func (c *UsersClient) GetByEmail(email string) (res structs.UserResponse, err error) {
	//Build Request
	p := url.Values{}
	p.Add("token", c.token)
	p.Add("email", email)
	//Return response
	err = urlEncodedClient(ugURL, "lookupByEmail", c.token, p, &res)
	return res, err
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
