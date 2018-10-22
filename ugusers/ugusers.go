package ugusers

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
	ugURL = "https://slack.com/api/usergroups.users."
)

// NewUgClient creates a new client with an access token
func NewUgClient(t string) *UgClient {
	c := new(UgClient)
	c.setToken(t)
	return c
}

//UgClient is the way to interface with Slack's chat methods
type UgClient struct {
	token string
}

// Change the access token for some reason
func (c *UgClient) setToken(t string) {
	c.token = t
}

// Update updates the list of users that belong to a User Group
// https://api.slack.com/methods/usergroups.users.update
func (c *UgClient) Update(usergroup string, users []string) (res structs.UpdateUgResponse, err error) {

	//Build request
	reqBod := updateStruct{
		Token:     c.token,
		Usergroup: usergroup,
		Users:     strings.Join(users, ","),
	}
	err = jsonRequest(ugURL, "update", c.token, reqBod, &res)
	check(err)
	//Return Response
	return
}

// GetUgList retrieves a list of users to a message in Slack
// https://api.slack.com/methods/usergroups.users.list
func (c *UgClient) GetUgList(usergroup string) (res structs.GetUgListResponse, err error) {
	//Build Request
	p := url.Values{}
	p.Add("token", c.token)
	p.Add("usergroup", usergroup)
	//Return response
	err = urlEncodedClient(ugURL, "list", c.token, p, &res)
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
