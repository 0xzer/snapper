package snapper

type UsersPublicInfoResponse struct {
	Snapchatters []Snapchatter `json:"snapchatters"`
}
type Snapchatter struct {
	ID              	string `json:"user_id"`
	Username            string `json:"username"`
	DisplayName         string `json:"display_name"`
	BitmojiAvatarID     string `json:"bitmoji_avatar_id"`
	BitmojiSelfieID     string `json:"bitmoji_selfie_id"`
	MutableUsername     string `json:"mutable_username"`
	BitmojiBackgroundID string `json:"bitmoji_background_id,omitempty"`
}