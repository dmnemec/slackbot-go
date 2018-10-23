[![Go Report Card](https://goreportcard.com/badge/github.com/dmnemec/slackbot-go)](https://goreportcard.com/report/github.com/dmnemec/slackbot-go)
# slackbot-go 
A framework in Go for building apps (currently only internal implementations) for Slack.

Package Descriptions
* client - this is the primary thing you need, it contains all the other package clients inside it
  * `import "github.com/dmnemec/slackbot-go/slack"`
  * `client := slack.NewClient("slack_token_string")`
  * `client.AnyWebAPIFunction(Input)` to use the desired WebAPI function
  * Functions can be found by looking through the packages. I'll eventually put the full list here (probably)
* core - housekeeping functions for a slackbot
  * Configuration file management
  * Post messages using Incoming Webhooks
* basecamp - for integration with Basecamp
  * Pull events from basecamp
  * Perform actions on events
  * Pull project names
* converstaions - implementing the Slack Conversations API
  * 7/19 Functions
