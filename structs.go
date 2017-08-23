package main

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
	Exerpt     string `json:"exerpt"`
	Raw_exerpt string `json:"raw_exerpt"`
	Summary    string `json:"summary"`
	Url        string `json:"url"`
	Html_url   string `json:"html_url"`
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
