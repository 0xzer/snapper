package snapper

import (
	"encoding/json"
	"os"
	"github.com/0xzer/snapper/types"
)

type Session struct {
	CurrentUser *User `json:"currentUser,omitempty"`
	FideliusKeys *types.FideliusKeys `json:"fideliusKeys,omitempty"`
	SnapCookies *types.SnapCookies `json:"snapCookies,omitempty"`
	SnapTokens *types.SnapTokens `json:"snapTokens,omitempty"`
}

func (s *Session) GetCurrentUser() *User {
	return s.CurrentUser
}

func (s *Session) Save(path string) error {
	jsonData, jsonErr := json.Marshal(s)
	if jsonErr != nil {
		return jsonErr
	}

	writeErr := os.WriteFile(path, jsonData, os.ModePerm)
	if writeErr != nil {
		return writeErr
	}

	return nil
}

// FROM JSON FILE
func NewSessionFromFile(path string) (*Session, error) {
	session := &Session{}

	jsonBytes, readErr := os.ReadFile(path)
	if readErr != nil {
		return nil, readErr
	}

	unmarshalErr := json.Unmarshal(jsonBytes, session)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return session, nil
}

func NewSessionFromCookies(cookieStr string) (*Session, error) {
	snapCookies := types.NewCookiesFromString(cookieStr)

	return &Session{
		SnapCookies: snapCookies,
		SnapTokens: &types.SnapTokens{},
	}, nil
}