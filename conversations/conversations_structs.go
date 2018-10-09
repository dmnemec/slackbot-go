package conversations

type createChannelStruct struct {
	Token     string   `json:"token"`
	Name      string   `json:"name"`
	IsPrivate bool     `json:"is_private"`
	UserIds   []string `json:"user_ids"`
}

type setPurposeStruct struct {
	Token   string `json:"token"`
	Channel string `json:"channel"`
	Purpose string `json:"purpose"`
}
