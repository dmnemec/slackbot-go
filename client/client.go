package client

import (
	//Local
	"github.com/dmnemec/slackbot-go/chat"
	"github.com/dmnemec/slackbot-go/conversations"
)

// Client is a total Slack API client
type Client struct {
	*conversations.ConvoClient
	*chat.ChatClient
	Token string
}

func (c *Client) new(token string) *Client {
	c = &Client{}
	c.ConvoClient = conversations.NewConvoClient(token)
	c.ChatClient = chat.NewChatClient(token)
	c.Token = token
	return c
}

// New creates a new Slack API Client struct with an OAuth token
func New(token string) (client *Client) {
	return client.new(token)
}
