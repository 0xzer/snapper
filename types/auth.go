package types

import (
	"encoding/json"
	"os"
	"strings"
)

type FideliusKeys struct {
	RWK []byte `json:"rwk"` // hmac_sha256_key used to sign the keypairid
	PubKey []byte `json:"pubKey"` // ecdsa pubkey
	PrivKey []byte `json:"privKey"` // ecdsa privkey
	UncompressedPubkey []byte `json:"uncompressedKey"` // uncompressed pubkey
}

type SnapTokens struct {
	SSO_TOKEN string `json:"sso_token,omitempty"`
}

type SnapCookies struct {
	HOST_SC_A_SESSION string `json:"__Host-sc-a-session,omitempty"`
	HOST_X_SNAP_CLIENT_COOKIE string `json:"__Host-X-Snap-Client-Cookie,omitempty"`
	HOST_SC_A_NONCE string `json:"__Host-sc-a-nonce,omitempty"`
}

// FROM JSON FILE.
func NewCookiesFromFile(path string) (*SnapCookies, error) {
	jsonBytes, jsonErr := os.ReadFile(path)
	if jsonErr != nil {
		return nil, jsonErr
	}

	session := &SnapCookies{}

	marshalErr := json.Unmarshal(jsonBytes, session)
	if marshalErr != nil {
		return nil, marshalErr
	}

	return session, nil
}


func NewCookiesFromString(cookieStr string) *SnapCookies {
	sc_nonce := extractCookieValue(cookieStr, "__Host-sc-a-nonce")
	sc_session := extractCookieValue(cookieStr, "__Host-sc-a-session")
	sc_snap_client_cookie := extractCookieValue(cookieStr, "__Host-X-Snap-Client-Cookie")

	return &SnapCookies{
		HOST_SC_A_SESSION: sc_session,
		HOST_X_SNAP_CLIENT_COOKIE: sc_snap_client_cookie,
		HOST_SC_A_NONCE: sc_nonce,
	}
}

func extractCookieValue(cookieString, key string) string {
	startIndex := strings.Index(cookieString, key)
	if startIndex == -1 {
		return ""
	}

	startIndex += len(key) + 1
	endIndex := strings.IndexByte(cookieString[startIndex:], ';')
	if endIndex == -1 {
		return cookieString[startIndex:]
	}

	return cookieString[startIndex : startIndex+endIndex]
}