package usergroups

import (
	//Local
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/dmnemec/slackbot-go/structs"
)

const (
	usergroupsURL = "https://slack.com/api/usergroups."
)

// NewUsergroupsClient creates a new client with an access token
func NewUsergroupsClient(t string) *UsergroupsClient {
	c := new(UsergroupsClient)
	c.setToken(t)
	return c
}

//UsergroupsClient is the way to interface with Slack's chat methods
type UsergroupsClient struct {
	token string
}

// Change the access token for some reason
func (c *UsergroupsClient) setToken(t string) {
	c.token = t
}

// CreateUsergroups creates a User Group
// https://api.slack.com/methods/usergroups.create
func (c *UsergroupsClient) UpdateUsergroups(name string) (res structs.UpdateUsergroupsResponse, err error) {

	//Build request
	reqBod := updateStruct{
		Token: c.token,
		Name:  string,
	}
	err = jsonRequest(usergroupsURL, "create", c.token, reqBod, &res)
	check(err)
	//Return Response
	return
}

// https://api.slack.com/methods/usergroups.enable
func (c *UsergroupsClient) EnableUsergroups(name string) (res structs.UpdateUsergroupsResponse, err error) {

	//Build request
	reqBod := updateStruct{
		Token: c.token,
		Name:  string,
	}
	err = jsonRequest(usergroupsURL, "enable", c.token, reqBod, &res)
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
