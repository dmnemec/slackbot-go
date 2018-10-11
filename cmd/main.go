package main

import (
	//Local
	"github.com/dmnemec/slackbot-go/client"
)

func main() {
	slackAPI := client.New()

	res, err := slackAPI.List()
}
