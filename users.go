package snapper

import (
	"encoding/json"
	"net/http"

	"github.com/0xzer/snapper/data/headers"
	"github.com/0xzer/snapper/data/paths"
	"github.com/0xzer/snapper/payload"
	"github.com/0xzer/snapper/types"
)


type Users struct {
	client *Client
}

func (u *Users) GetPublicInfo(userIds []string) ([]Snapchatter, error) {
	payload, payloadErr := payload.GetPublicUserInfo(userIds, "CHAT")
	if payloadErr != nil {
		return nil, payloadErr
	}

	headers := headers.NewDefaultHeaders(u.client.Session.SnapCookies, u.client.Session.SnapTokens, true)
	headers.Add("content-type", "application/x-www-form-urlencoded; charset=utf-8")

	_, responseBody, requestErr := u.client.Request.MakeRequest(paths.GET_USER_PUBLIC_INFO, "POST", headers, payload, types.NORMAL)
	if requestErr != nil {
		return nil, requestErr
	}

	usersInfoResponse := &UsersPublicInfoResponse{}

	unmarshalErr := json.Unmarshal(responseBody, usersInfoResponse)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return usersInfoResponse.Snapchatters, nil
}

func (u *Users) GetBitmojiImage(bitmojiSelfieId string, bitmojiAvatarId string, scale string) ([]byte, error) {
	urlPath := u.GetBitmojiURL(bitmojiSelfieId, bitmojiAvatarId, scale)
	_, responseBody, requestErr := u.client.Request.MakeRequest(urlPath, "GET", http.Header{}, nil, types.NORMAL)
	if requestErr != nil {
		return nil, requestErr
	}

	return responseBody, requestErr
}

// scale is a number, 0-2 but in string for the url formatting
func (u *Users) GetBitmojiURL(bitmojiSelfieId string, bitmojiAvatarId string, scale string) string {
	return "https://sdk.bitmoji.com/render/panel/"+bitmojiSelfieId+"-"+bitmojiAvatarId+"-v1.png?transparent=1&scale="+scale
}