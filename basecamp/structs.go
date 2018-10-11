package basecamp

import (
	"fmt"
)

type Event struct {
	Id        int        `json:"id"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
	Private   bool       `json:"private"`
	Action    string     `json:"action"`
	Target    string     `json:"target"`
	Eventable EventableO `json:"eventable"`
	Creator   CreatorO   `json:"creator"`
	//	Attachements Interface  `json:"attachements"`
	Excerpt    string `json:"excerpt"`
	RawExcerpt string `json:"raw_excerpt"`
	Summary    string `json:"summary"`
	Url        string `json:"url"`
	HtmlURL    string `json:"html_url"`
}

type EventableO struct {
	Id     int    `json:"id"`
	Type   string `json:"type"`
	Url    string `json:"url"`
	AppURL string `json:"app_url"`
}

type CreatorO struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	AvatarURL         string `json:"avatar_url"`
	FullsizeAvatarURL string `json:"fullsize_avatar_url"`
}

func (e *Event) Print() {
	fmt.Printf("Id = %v\n"+
		"CreatedAt = %v\n"+
		"UpdatedAt = %v\n"+
		"Private = %v\n"+
		"Action = %v\n"+
		"Target = %v\n"+
		"Eventable:\n",
		e.Id, e.CreatedAt, e.UpdatedAt, e.Private, e.Action,
		e.Target)
	e.Eventable.Print()
	fmt.Print("Creator:\n")
	e.Creator.Print()
	if e.Excerpt != "" {
		fmt.Printf("Excerpt = %v\n"+
			"RawExcerpt = %v\n"+
			"Summary = %v\n"+
			"Url = %v\n"+
			"HtmlURL = %v\n",
			e.Excerpt, e.RawExcerpt, e.Summary, e.Url, e.HtmlURL)
	} else {
		fmt.Printf("Summary = %v\n"+
			"Url = %v\n"+
			"HtmlURL = %v\n",
			e.Summary, e.Url, e.HtmlURL)
	}
}

func (e *EventableO) Print() {
	fmt.Printf("   Id = %v\n"+
		"   Type = %v\n"+
		"   Url = %v\n"+
		"   AppURL = %v\n",
		e.Id, e.Type, e.Url, e.AppURL)
}

func (e *CreatorO) Print() {
	fmt.Printf("   Id = %v\n"+
		"   Name = %v\n"+
		"   AvatarURL = %v\n"+
		"   FullsizeAvatarURL = %v\n",
		e.Id, e.Name, e.AvatarURL, e.FullsizeAvatarURL)
}

type Project struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Archived        bool   `json:"archived"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	Template        bool   `json:"template"`
	Starred         bool   `json:"starred"`
	Trashed         bool   `json:"trashed"`
	Draft           bool   `json:"draft"`
	IsClientProject bool   `json:"is_client_project"`
	Color           string `json:"color"`
	Creator         struct {
		ID                int    `json:"id"`
		Name              string `json:"name"`
		AvatarURL         string `json:"avatar_url"`
		FullsizeAvatarURL string `json:"fullsize_avatar_url"`
	} `json:"creator"`
	Accesses struct {
		Count     int    `json:"count"`
		UpdatedAt string `json:"updated_at"`
		URL       string `json:"url"`
		AppURL    string `json:"app_url"`
	} `json:"accesses"`
	Attachments struct {
		Count     int         `json:"count"`
		UpdatedAt interface{} `json:"updated_at"`
		URL       string      `json:"url"`
		AppURL    string      `json:"app_url"`
	} `json:"attachments"`
	CalendarEvents struct {
		Count     int    `json:"count"`
		UpdatedAt string `json:"updated_at"`
		URL       string `json:"url"`
		AppURL    string `json:"app_url"`
	} `json:"calendar_events"`
	Documents struct {
		Count     int         `json:"count"`
		UpdatedAt interface{} `json:"updated_at"`
		URL       string      `json:"url"`
		AppURL    string      `json:"app_url"`
	} `json:"documents"`
	Topics struct {
		Count     int    `json:"count"`
		UpdatedAt string `json:"updated_at"`
		URL       string `json:"url"`
		AppURL    string `json:"app_url"`
	} `json:"topics"`
	Todolists struct {
		RemainingCount int    `json:"remaining_count"`
		CompletedCount int    `json:"completed_count"`
		UpdatedAt      string `json:"updated_at"`
		URL            string `json:"url"`
		AppURL         string `json:"app_url"`
	} `json:"todolists"`
}
