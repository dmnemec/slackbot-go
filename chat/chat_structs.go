package chat

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
