package main

import (
	//Global
	"fmt"
	"os"

	//Local
	"github.com/dmnemec/slackbot-go/client"
)

//Outputs the user ids of the people in the listed channel
//To Run: go run test_conversations.go <channel_id>
func main() {
	slackAPI := client.New(os.Getenv("WORKSPACE_ACCESS_TOKEN"))

	res, err := slackAPI.PostMessage(os.Args[1], "test message")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Response: %v", res.Ok)
}
