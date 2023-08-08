package snapper

import "github.com/0xzer/snapper/protos"

type AccountResponse struct {
	Data       UserData   `json:"data,omitempty"`
}

type UserData struct {
	User User `json:"user,omitempty"`
}

type User struct {
	ID                       string `json:"id"`
	EncodedUserId 			 *protos.UUID `json:"encodedUserId,omitempty"`
	BitmojiAvatarID          string `json:"bitmojiAvatarId"`
	BitmojiSelfieID          string `json:"bitmojiSelfieId"`
	BitmojiBackgroundID      string `json:"bitmojiBackgroundId"`
	BitmojiSceneID           string `json:"bitmojiSceneId"`
	IsEmployee               bool   `json:"isEmployee"`
	Username                 string `json:"username"`
	DisplayName              string `json:"displayName"`
	SnapPrivacy              int    `json:"snapPrivacy"`
	HasUserSeenDWeb          bool   `json:"hasUserSeenDWeb"`
	HasUserAcceptedMerlinJIT bool   `json:"hasUserAcceptedMerlinJIT"`
}