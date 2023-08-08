package headers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/0xzer/snapper/types"
)

var USER_AGENT = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"

// DO NOT REMOVE THE REFERERS/ORIGINS they are needed
func NewDefaultHeaders(cookies *types.SnapCookies, tokens *types.SnapTokens, withAuth bool) http.Header {
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
	}

	if (withAuth) {
		token := "Bearer " + tokens.SSO_TOKEN
		h.Add("Authorization", token)
		c := "sc-a-nonce=" + cookies.HOST_SC_A_NONCE
		h.Add("Cookie", c)
	} else {
		h.Add("Cookie", fmt.Sprintf("__Host-X-Snap-Client-Cookie=%s; __Host-sc-a-session=%s; __Host-sc-a-nonce=%s", cookies.HOST_X_SNAP_CLIENT_COOKIE, cookies.HOST_SC_A_SESSION, cookies.HOST_SC_A_NONCE))
	}
	return h
}

func ConvertAuthSessionToCookieString(authSession *types.SnapCookies) string {
	var sb strings.Builder

	addCookie := func(name, value string) {
		sb.WriteString(name)
		sb.WriteString("=")
		sb.WriteString(value)
		sb.WriteString("; ")
	}

	addCookie("__Host-sc-a-nonce", authSession.HOST_SC_A_NONCE)

	cookieString := sb.String()
	cookieString = cookieString[:len(cookieString)-2]

	return cookieString
}