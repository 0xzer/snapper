package snapper

import (
	"encoding/json"
	"github.com/0xzer/snapper/crypto"
	"github.com/0xzer/snapper/data/headers"
	"github.com/0xzer/snapper/data/paths"
	"github.com/0xzer/snapper/payload"
	"github.com/0xzer/snapper/protos"
	"github.com/0xzer/snapper/snaperror"
	"github.com/0xzer/snapper/types"
)

type Account struct {
	client *Client
}

func (a *Account) InitializeWebKey() (*types.FideliusKeys, error) {
	headers := headers.NewCoreHeaders(a.client.Session.SnapCookies, a.client.Session.SnapTokens, a.client.device)
	tentativeWebKey, pubKey, privKey, err := crypto.NewFideliusTentativeWebKey()
	if err != nil {
		return nil, err
	}
	
	payload := &protos.InitializeWebKeyRequest{
		TentativeWebKey: tentativeWebKey,
	}
	
	payloadBytes, err := protos.EncodeProtoMessage(payload)
	if err != nil {
		return nil, err
	}

	_, responseBody, err := a.client.Request.MakeRequest(paths.INITIALIZE_WEB_KEY, "POST", headers, payloadBytes, types.GRPC)
	if err != nil {
		return nil, err
	}


	var resData protos.InitializeWebKeyResponse
	err = protos.DecodeGRPCProtoMessage(responseBody, &resData)
	if err != nil {
		return nil, err
	}

	return &types.FideliusKeys{
		RWK: resData.Rwk,
		PubKey: pubKey,
		PrivKey: privKey,
		UncompressedPubkey: tentativeWebKey.UncompressedPubkey,
	}, nil
}

func (a *Account) Friends() (*FriendsResponse, error) {
	if a.client.Session.CurrentUser.ID == "" {
		return nil, &snaperror.GetFriendsError{ErrorMessage: "current user_id was empty in session"} 
	}

	payload := payload.GetFriends(a.client.Session.CurrentUser.ID)


	headers := headers.NewDefaultHeaders(a.client.Session.SnapCookies, a.client.Session.SnapTokens, true)
	headers.Add("content-type", "application/x-www-form-urlencoded; charset=utf-8")

	_, responseBody, requestErr := a.client.Request.MakeRequest(paths.GET_FRIENDS, "POST", headers, payload, types.NORMAL)
	if requestErr != nil {
		return nil, requestErr
	}

	friendsResponse := &FriendsResponse{}

	unmarshalErr := json.Unmarshal(responseBody, friendsResponse)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	
	return friendsResponse, nil
}

func (a *Account) Info() (*User, error) {
	headers := headers.NewDefaultHeaders(a.client.Session.SnapCookies, a.client.Session.SnapTokens, true)
	headers.Add("content-type", "application/json")

	_, responseBody, requestErr := a.client.Request.MakeRequest(paths.WEB_GRAPHQL_URL, "POST", headers, payload.USER_QUERY, types.NORMAL)
	if requestErr != nil {
		return nil, requestErr
	}

	accountResponse := &AccountResponse{}

	unmarshalErr := json.Unmarshal(responseBody, accountResponse)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	
	currentUser := &accountResponse.Data.User

	a.client.Session.CurrentUser = currentUser
	encodedUserId, encodeErr := crypto.EncodeUUID(currentUser.ID)
	if encodeErr != nil {
		return nil, encodeErr
	}
	a.client.Session.CurrentUser.EncodedUserId = &protos.UUID{
		EncodedId: encodedUserId,
	}

	return currentUser, nil
}

func (a *Account) Authenticate() error {
	if a.client.Session == nil {
		return &snaperror.AuthenticationError{ErrorMessage: "no session provided in client struct"}
	}

	headers := headers.NewDefaultHeaders(a.client.Session.SnapCookies, a.client.Session.SnapTokens, false)
	headers.Add("content-length", "0")
	_, responseBody, requestErr := a.client.Request.MakeRequest(paths.ACCOUNTS_SSO, "POST", headers, nil, types.NORMAL)
	if requestErr != nil {
		return requestErr
	}
	
	sso_token := string(responseBody)
	if len(sso_token) <= 400 {
		a.client.Session.SnapTokens.SSO_TOKEN = sso_token
		a.client.logger.Debug().Any("sso_token", string(sso_token)).Msg("SSO_TOKEN Response")
	} else {
		return &snaperror.AuthenticationError{ErrorMessage: "invalid snap cookies, try retrieving new ones"}
	}

	return nil
}