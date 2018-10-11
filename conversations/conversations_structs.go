package conversations

type createChannelStruct struct {
	Token     string   `json:"token"`
	Name      string   `json:"name"`
	IsPrivate bool     `json:"is_private"`
	UserIds   []string `json:"user_ids"`
}

type postMessageStruct struct {
	Token   string `json:"token"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

type setPurposeStruct struct {
	Token   string `json:"token"`
	Channel string `json:"channel"`
	Purpose string `json:"purpose"`
}

type inviteStruct struct {
	Token   string   `json:"token"`
	Channel string   `json:"channel"`
	Users   []string `json:"users"`
}
