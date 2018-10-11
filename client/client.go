package client

import (
	//Global
	"context"
	"fmt"
	"log"
	"os"

	//Local
	"github.com/DMNemec/slackbot-go/conversations"

	//Remote
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/slack"
)

// Client is a total Slack API client
type Client struct {
	*conversations.Client
	Token string
}

func (c *Client) new(token string) *Client {
	c.Client = conversations.NewClient(token)
	c.Token = token
	return c
}

// New creates a new Slack API Client struct with an OAuth token
func New() (client *Client) {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     os.Getenv("workspace_ID"),
		ClientSecret: os.Getenv("workspace_token"),
		Scopes:       []string{"conversations:read", "conversations:write"},
		Endpoint:     slack.Endpoint,
	}

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	return client.new(tok.AccessToken)
}
