package snapper

type FriendsResponse struct {
	Friends                      []Friend                    `json:"friends"`
	FriendsSyncToken             string                      `json:"friends_sync_token"`
	FriendsSyncType              string                      `json:"friends_sync_type"`
	AddedFriends                 []AddedFriend               `json:"added_friends"`
	Bests                        []any                       `json:"bests"`
	ExtraFriendmojiMutableDict   ExtraFriendmojiMutableDict  `json:"extra_friendmoji_mutable_dict"`
	ExtraFriendmojiReadOnlyDict  ExtraFriendmojiReadOnlyDict `json:"extra_friendmoji_read_only_dict"`
	AddedFriendsSyncToken        string                      `json:"added_friends_sync_token"`
	AddedFriendsSyncType         string                      `json:"added_friends_sync_type"`
	BestsUserIds                 []any                       `json:"bests_user_ids"`
	IsResponseWithPartialColumns bool                        `json:"is_response_with_partial_columns"`
	InvitedUsers                 []any                       `json:"invited_users"`
}

type Devices struct {
	OutBeta string `json:"out_beta"`
	Version int    `json:"version"`
}

type FideliusInfo struct {
	Devices []Devices `json:"devices"`
}

type Friend struct {
	Username                            string       `json:"name"`
	ID                              	string       `json:"user_id"`
	Type                                int          `json:"type"`
	DisplayName                         string       `json:"display"`
	Birthday                            string       `json:"birthday,omitempty"`
	Ts                                  int64        `json:"ts,omitempty"`
	ReverseTs                           int64        `json:"reverse_ts,omitempty"`
	Direction                           string       `json:"direction"`
	CanSeeCustomStories                 bool         `json:"can_see_custom_stories"`
	Expiration                          int          `json:"expiration"`
	FriendmojiString                    string       `json:"friendmoji_string"`
	Friendmojis                         []any        `json:"friendmojis"`
	SnapStreakCount                     int          `json:"snap_streak_count"`
	FideliusInfo                        FideliusInfo `json:"fidelius_info,omitempty"`
	IsPopular                           bool         `json:"is_popular"`
	IsStoryMuted                        bool         `json:"is_story_muted"`
	MutableUsername                     string       `json:"mutable_username"`
	IsCameosSharingSupported            bool         `json:"is_cameos_sharing_supported,omitempty"`
	CameosSharingPolicy                 int          `json:"cameos_sharing_policy"`
	PlusBadgeVisibility                 int          `json:"plus_badge_visibility"`
	BitmojiAvatarID                     string       `json:"bitmoji_avatar_id,omitempty"`
	BitmojiSelfieID                     string       `json:"bitmoji_selfie_id,omitempty"`
	BitmojiSceneID                      string       `json:"bitmoji_scene_id,omitempty"`
	BitmojiBackgroundID                 string       `json:"bitmoji_background_id,omitempty"`
	IsBitmojiFriendmojiSharingSupported bool         `json:"is_bitmoji_friendmoji_sharing_supported,omitempty"`
	IgnoredLink                         bool         `json:"ignored_link,omitempty"`
}

type AddedFriend struct {
	Name                          string `json:"name"`
	UserID                        string `json:"user_id"`
	Type                          int    `json:"type"`
	Display                       string `json:"display"`
	Ts                            int64  `json:"ts"`
	ReverseTs                     int    `json:"reverse_ts"`
	Direction                     string `json:"direction"`
	Expiration                    int    `json:"expiration"`
	AddSource                     string `json:"add_source"`
	AddSourceType                 string `json:"add_source_type"`
	IsIncomingFriendRequestViewed bool   `json:"is_incoming_friend_request_viewed"`
	MutableUsername               string `json:"mutable_username"`
	SnapshotMetadata              string `json:"snapshot_metadata"`
	CameosSharingPolicy           int    `json:"cameos_sharing_policy"`
}

type ExtraFriendmojiMutableDict struct {}

type ExtraFriendmojiReadOnlyDict struct {}