package structs

//CloseResponse is returned from the Close method
type CloseResponse struct {
	Ok            bool   `json:"ok"`
	Error         string `json:"error"`
	NoOp          bool   `json:"no_op"`
	AlreadyClosed bool   `json:"already_closed"`
}

//InviteResponse is returned from the Invite method
type InviteResponse struct {
	Ok      bool   `json:"ok"`
	Error   string `json:"error"`
	Channel struct {
		ID                 string        `json:"id"`
		Name               string        `json:"name"`
		IsChannel          bool          `json:"is_channel"`
		IsGroup            bool          `json:"is_group"`
		IsIm               bool          `json:"is_im"`
		Created            int           `json:"created"`
		Creator            string        `json:"creator"`
		IsArchived         bool          `json:"is_archived"`
		IsGeneral          bool          `json:"is_general"`
		Unlinked           int           `json:"unlinked"`
		NameNormalized     string        `json:"name_normalized"`
		IsReadOnly         bool          `json:"is_read_only"`
		IsShared           bool          `json:"is_shared"`
		IsExtShared        bool          `json:"is_ext_shared"`
		IsOrgShared        bool          `json:"is_org_shared"`
		PendingShared      []interface{} `json:"pending_shared"`
		IsPendingExtShared bool          `json:"is_pending_ext_shared"`
		IsMember           bool          `json:"is_member"`
		IsPrivate          bool          `json:"is_private"`
		IsMpim             bool          `json:"is_mpim"`
		LastRead           string        `json:"last_read"`
		Topic              struct {
			Value   string `json:"value"`
			Creator string `json:"creator"`
			LastSet int    `json:"last_set"`
		} `json:"topic"`
		Purpose struct {
			Value   string `json:"value"`
			Creator string `json:"creator"`
			LastSet int    `json:"last_set"`
		} `json:"purpose"`
		PreviousNames []string `json:"previous_names"`
		NumMembers    int      `json:"num_members"`
		Locale        string   `json:"locale"`
	} `json:"channel"`
}

//ListResponse is returned from the List method
type ListResponse struct {
	Ok       bool   `json:"ok"`
	Error    string `json:"error"`
	Channels []struct {
		ID                 string        `json:"id"`
		Name               string        `json:"name"`
		IsChannel          bool          `json:"is_channel"`
		IsGroup            bool          `json:"is_group"`
		IsIm               bool          `json:"is_im"`
		Created            int           `json:"created"`
		Creator            string        `json:"creator"`
		IsArchived         bool          `json:"is_archived"`
		IsGeneral          bool          `json:"is_general"`
		Unlinked           int           `json:"unlinked"`
		NameNormalized     string        `json:"name_normalized"`
		IsShared           bool          `json:"is_shared"`
		IsExtShared        bool          `json:"is_ext_shared"`
		IsOrgShared        bool          `json:"is_org_shared"`
		PendingShared      []interface{} `json:"pending_shared"`
		IsPendingExtShared bool          `json:"is_pending_ext_shared"`
		IsMember           bool          `json:"is_member"`
		IsPrivate          bool          `json:"is_private"`
		IsMpim             bool          `json:"is_mpim"`
		Topic              struct {
			Value   string `json:"value"`
			Creator string `json:"creator"`
			LastSet int    `json:"last_set"`
		} `json:"topic"`
		Purpose struct {
			Value   string `json:"value"`
			Creator string `json:"creator"`
			LastSet int    `json:"last_set"`
		} `json:"purpose"`
		PreviousNames []interface{} `json:"previous_names"`
		NumMembers    int           `json:"num_members"`
	} `json:"channels"`
	ResponseMetadata struct {
		NextCursor string `json:"next_cursor"`
	} `json:"response_metadata"`
}

//MembersResponse is returned from the Members method
type MembersResponse struct {
	Ok               bool     `json:"ok"`
	Error            string   `json:"error"`
	Members          []string `json:"members"`
	ResponseMetadata struct {
		NextCursor string `json:"next_cursor"`
	} `json:"response_metadata"`
}

//PostMessageResponse is returned from the PostMessage method
type PostMessageResponse struct {
	Ok      bool   `json:"ok"`
	Error   string `json:"error"`
	Channel string `json:"channel"`
	Ts      string `json:"ts"`
	Message struct {
		Text        string `json:"text"`
		Username    string `json:"username"`
		BotID       string `json:"bot_id"`
		Attachments []struct {
			Text     string `json:"text"`
			ID       int    `json:"id"`
			Fallback string `json:"fallback"`
		} `json:"attachments"`
		Type    string `json:"type"`
		Subtype string `json:"subtype"`
		Ts      string `json:"ts"`
	} `json:"message"`
}

//SetPurposeResponse is returned from the SetPurpose method
type SetPurposeResponse struct {
	Ok      bool   `json:"ok"`
	Purpose string `json:"purpose"`
	Error   string `json:"error"`
}
