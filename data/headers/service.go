package headers

import (
	"net/http"
	"github.com/0xzer/snapper/types"
)

var SNAPCHAT_USER_AGENT = "SnapchatWeb/12.38.0 PROD (linux 0.0.0; chrome 113.0.0.0)"

// DO NOT REMOVE THE REFERERS/ORIGINS they are needed
func NewServiceHeaders(cookies *types.SnapCookies, tokens *types.SnapTokens, withSnapAgent bool, withDeviceInfo bool) http.Header {
	var h = http.Header{
		"accept": []string{"*/*"},
		"accept-language": []string{"en-US,en;q=0.9"},
		"connection": []string{"keep-alive"},
		"origin": []string{"https://web.snapchat.com"},
		"referer": []string{"https://web.snapchat.com/"},
		"sec-ch-ua": []string{"\"Google Chrome\";v=\"113\", \"Chromium\";v=\"113\", \"Not-A.Brand\";v=\"24\""},
		"sec-ch-ua-mobile": []string{"?0"},
		"sec-fetch-dest": []string{"empty"},
		"sec-fetch-mode": []string{"cors"},
		"sec-fetch-site": []string{"same-site"},
		"user-agent": []string{USER_AGENT},
		"x-grpc-web": []string{"1"},
		"content-type": []string{"application/grpc-web+proto"},
	}
	
	h.Add("Authorization", "Bearer "+tokens.SSO_TOKEN)
	h.Add("cookie", "sc-a-nonce="+cookies.HOST_SC_A_NONCE)

	if withSnapAgent {
		h.Add("x-snap-client-user-agent", SNAPCHAT_USER_AGENT)
	}
	return h
}