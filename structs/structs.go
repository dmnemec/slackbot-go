package structs

//SetPurposeResponse is what comes back from the SetPurpose method
type SetPurposeResponse struct {
	Ok      bool   `json:"ok"`
	Purpose string `json:"purpose"`
}
