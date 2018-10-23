package structs

//CloseResponse is returned from the Close method
type CloseResponse struct {
	Ok            bool   `json:"ok"`
	Error         string `json:"error"`
	NoOp          bool   `json:"no_op"`
	AlreadyClosed bool   `json:"already_closed"`
}

//CreateResponse is returned from the Create method
type CreateResponse struct {
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
		IsShared           bool          `json:"is_shared"`
		IsExtShared        bool          `json:"is_ext_shared"`
		IsOrgShared        bool          `json:"is_org_shared"`
		PendingShared      []interface{} `json:"pending_shared"`
		IsPendingExtShared bool          `json:"is_pending_ext_shared"`
		IsMember           bool          `json:"is_member"`
		IsPrivate          bool          `json:"is_private"`
		IsMpim             bool          `json:"is_mpim"`
		LastRead           string        `json:"last_read"`
		Latest             interface{}   `json:"latest"`
		UnreadCount        int           `json:"unread_count"`
		UnreadCountDisplay int           `json:"unread_count_display"`
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
		Priority      int           `json:"priority"`
	} `json:"channel"`
}

//DeleteResponse is returned by the Delete method
type DeleteResponse struct {
	Ok      bool   `json:"ok"`
	Error   string `json:"error"`
	Channel string `json:"channel"`
	Ts      string `json:"ts"`
}

//GetPermalinkResponse is returned from the GetPermalink method
type GetPermalinkResponse struct {
	Ok        bool   `json:"ok"`
	Channel   string `json:"channel"`
	Error     string `json:"error"`
	Permalink string `json:"permalink"`
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

//PostEphemeralResponse is returned from the PostEphemeral method
type PostEphemeralResponse struct {
	Ok        bool   `json:"ok"`
	Error     string `json:"error"`
	MessageTs string `json:"message_ts"`
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

//UnfurlResponse is returned from the Unfurl method
type UnfurlResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

//UpdateResponse is returned by the Update method
type UpdateResponse struct {
	Ok      bool   `json:"ok"`
	Error   string `json:"error"`
	Channel string `json:"channel"`
	Ts      string `json:"ts"`
	Text    string `json:"text"`
}

//UpdateUgResponse is returende by the UpdateUg method
type UpdateUgResponse struct {
	Ok        bool   `json:"ok"`
	Error     string `json:"error"`
	Usergroup struct {
		ID          string      `json:"id"`
		TeamID      string      `json:"team_id"`
		IsUsergroup bool        `json:"is_usergroup"`
		Name        string      `json:"name"`
		Description string      `json:"description"`
		Handle      string      `json:"handle"`
		IsExternal  bool        `json:"is_external"`
		DateCreate  int         `json:"date_create"`
		DateUpdate  int         `json:"date_update"`
		DateDelete  int         `json:"date_delete"`
		AutoType    interface{} `json:"auto_type"`
		CreatedBy   string      `json:"created_by"`
		UpdatedBy   string      `json:"updated_by"`
		DeletedBy   interface{} `json:"deleted_by"`
		Prefs       struct {
			Channels []interface{} `json:"channels"`
			Groups   []interface{} `json:"groups"`
		} `json:"prefs"`
		Users     []string `json:"users"`
		UserCount int      `json:"user_count"`
	} `json:"usergroup"`
}

//GetUgListResponse is returned by GetUgList method
type GetUgListResponse struct {
	Ok    bool     `json:"ok"`
	Error string   `json:"error"`
	Users []string `json:"users"`
}

type UserResponse struct {
	Ok   bool `json:"ok"`
	User struct {
		ID       string `json:"id"`
		TeamID   string `json:"team_id"`
		Name     string `json:"name"`
		Deleted  bool   `json:"deleted"`
		Color    string `json:"color"`
		RealName string `json:"real_name"`
		Tz       string `json:"tz"`
		TzLabel  string `json:"tz_label"`
		TzOffset int    `json:"tz_offset"`
		Profile  struct {
			AvatarHash            string `json:"avatar_hash"`
			StatusText            string `json:"status_text"`
			StatusEmoji           string `json:"status_emoji"`
			RealName              string `json:"real_name"`
			DisplayName           string `json:"display_name"`
			RealNameNormalized    string `json:"real_name_normalized"`
			DisplayNameNormalized string `json:"display_name_normalized"`
			Email                 string `json:"email"`
			Image24               string `json:"image_24"`
			Image32               string `json:"image_32"`
			Image48               string `json:"image_48"`
			Image72               string `json:"image_72"`
			Image192              string `json:"image_192"`
			Image512              string `json:"image_512"`
			Team                  string `json:"team"`
		} `json:"profile"`
		IsAdmin           bool `json:"is_admin"`
		IsOwner           bool `json:"is_owner"`
		IsPrimaryOwner    bool `json:"is_primary_owner"`
		IsRestricted      bool `json:"is_restricted"`
		IsUltraRestricted bool `json:"is_ultra_restricted"`
		IsBot             bool `json:"is_bot"`
		Updated           int  `json:"updated"`
		IsAppUser         bool `json:"is_app_user"`
		Has2Fa            bool `json:"has_2fa"`
	} `json:"user"`
}
