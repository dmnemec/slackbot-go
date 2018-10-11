package main

import (
	//Global
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	//Local
	"github.com/dmnemec/slackbot-go/client"
)

//Outputs the user ids of the people in the listed channel
//To Run: go run test_conversations.go <channel_id>
func main() {
	slackAPI := client.New(os.Getenv("WORKSPACE_ACCESS_TOKEN"))

	res, err := slackAPI.Members(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Response:\n%v", strings.Replace(string(body), ",", "\n", -1))
}
