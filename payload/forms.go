package payload

import (
	"net/url"

	"github.com/0xzer/snapper/data/methods"
)


func GetFriends(userId string) []byte {
	ts := methods.GetTimestampString()

	values := &url.Values{}
	values.Add("friends_request", "{}")
	values.Add("timestamp", ts)
	values.Add("snapchat_user_id", userId)

	return []byte(values.Encode())
}

func GetPublicUserInfo(userIds []string, source string) ([]byte, error) {
	idsString, stringErr := methods.SliceToString(userIds)
	if stringErr != nil {
		return nil, stringErr
	}

	values := &url.Values{}
	values.Add("user_ids", idsString)
	values.Add("source", source)

	return []byte(values.Encode()), nil
}