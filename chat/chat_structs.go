package chat

type deleteStruct struct {
	Token     string `json:"token"`
	Channel   string `json:"channel"`
	Timestamp string `json:"ts"`
	AsUser    bool   `json:"as_user"`
}

type postMessageStruct struct {
	Token   string `json:"token"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

type postEphemeralStruct struct {
	Token   string `json:"token"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
	User    string `json:"user"`
}

type unfurlStruct struct {
	Token     string `json:"token"`
	Channel   string `json:"channel"`
	Timestamp string `json:"ts"`
	Unfurls   string `json:"unfurls"`
}
