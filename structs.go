package main

import (
	"fmt"
)

type Event struct {
	Id         int        `json:"id"`
	Created_at string     `json:"created_at"`
	Updated_at string     `json:"updated_at"`
	Private    bool       `json:"private"`
	Action     string     `json:"action"`
	Target     string     `json:"target"`
	Eventable  EventableO `json:"eventable"`
	Creator    CreatorO   `json:"creator"`
	//	Attachements Interface  `json:"attachements"`
	Excerpt     string `json:"excerpt"`
	Raw_excerpt string `json:"raw_excerpt"`
	Summary     string `json:"summary"`
	Url         string `json:"url"`
	Html_url    string `json:"html_url"`
}

type EventableO struct {
	Id      int    `json:"id"`
	Type    string `json:"type"`
	Url     string `json:"url"`
	App_url string `json:"app_url"`
}

type CreatorO struct {
	Id                  int    `json:"id"`
	Name                string `json:"name"`
	Avatar_url          string `json:"avatar_url"`
	Fullsize_avatar_url string `json:"fullsize_avatar_url"`
}

func (e *Event) Print() {
	fmt.Printf("Id = %v\n"+
		"Created_at = %v\n"+
		"Updated_at = %v\n"+
		"Private = %v\n"+
		"Action = %v\n"+
		"Target = %v\n"+
		"Eventable:\n",
		e.Id, e.Created_at, e.Updated_at, e.Private, e.Action,
		e.Target)
	e.Eventable.Print()
	fmt.Print("Creator:\n")
	e.Creator.Print()
	if e.Excerpt != "" {
		fmt.Printf("Excerpt = %v\n"+
			"Raw_excerpt = %v\n"+
			"Summary = %v\n"+
			"Url = %v\n"+
			"Html_url = %v\n",
			e.Excerpt, e.Raw_excerpt, e.Summary, e.Url, e.Html_url)
	} else {
		fmt.Printf("Summary = %v\n"+
			"Url = %v\n"+
			"Html_url = %v\n",
			e.Summary, e.Url, e.Html_url)
	}
}

func (e *EventableO) Print() {
	fmt.Printf("   Id = %v\n"+
		"   Type = %v\n"+
		"   Url = %v\n"+
		"   App_url = %v\n",
		e.Id, e.Type, e.Url, e.App_url)
}

func (e *CreatorO) Print() {
	fmt.Printf("   Id = %v\n"+
		"   Name = %v\n"+
		"   Avatar_url = %v\n"+
		"   Fullsize_avatar_url = %v\n",
		e.Id, e.Name, e.Avatar_url, e.Fullsize_avatar_url)
}
