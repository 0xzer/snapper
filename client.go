package snapper

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"

	"github.com/0xzer/snapper/protos"
	"github.com/0xzer/snapper/types"
	"github.com/rs/zerolog"
)

type EventHandler func(evt interface{})
type Proxy func(*http.Request) (*url.URL, error)
type Client struct {
	Request *Request
	Messaging *Messaging
	Account *Account
	Users *Users

	logger zerolog.Logger

	device string
	Session *Session
	proxy Proxy
}

// pass in an empty logger if u don't want logging
func NewClient(session *Session, logger zerolog.Logger, proxy *string) *Client {
	cli := &Client{
		logger: logger,
	}

	cli.Request = &Request{client: cli, http: &http.Client{}}

	if proxy != nil {
		cli.SetProxy(*proxy)
	}

	cli.Messaging = &Messaging{client: cli}
	cli.Account = &Account{client: cli}
	cli.Users = &Users{client: cli}
	
	if session != nil {
		cli.Session = session
		if session.SnapTokens.SSO_TOKEN == "" {
			cli.logger.Info().Msg("Requesting new sso-token using cookies...")
			_, authErr := cli.RefreshSSOToken()
			if authErr != nil {
				cli.logger.Fatal().Err(authErr).Msg("failed to authenticate:")
			}
		}

		_, accErr := cli.Account.Info()
		if accErr != nil {
			cli.logger.Info().Msg("Failed to fetch user information, requesting new sso-token using current credentials...")
			_, authErr := cli.RefreshSSOToken()
			if authErr != nil {
				cli.logger.Fatal().Err(authErr).Msg("failed to authenticate:")
			}
		}


		if cli.Session.FideliusKeys == nil {
			fideliusKeys, err := cli.Account.InitializeWebKey()
			if err != nil {
				cli.logger.Fatal().Err(err).Msg("failed to initialize web key:")
			}
			cli.Session.FideliusKeys = fideliusKeys
		}
	}
	
	deviceStr := types.NewDevice()
	cli.device = deviceStr
	return cli
}

func (c *Client) RefreshSSOToken() (string, error) {
	err := c.Account.Authenticate()
	if err != nil {
		return "", err
	}

	return c.Session.SnapTokens.SSO_TOKEN, nil
}

func (c *Client) SaveSession(path string) error {
	jsonData, jsonErr := json.Marshal(c.Session)
	if jsonErr != nil {
		return jsonErr
	}

	writeErr := os.WriteFile("session.json", jsonData, os.ModePerm)
	if writeErr != nil {
		return writeErr
	}
	
	return nil
}

func (c *Client) SetProxy(proxy string) error {
	proxyParsed, err := url.Parse(proxy)
	if err != nil {
		c.logger.Fatal().Err(err).Msg("Failed to set proxy")
	}
	proxyUrl := http.ProxyURL(proxyParsed)
	c.Request.http.Transport = &http.Transport{
		Proxy: proxyUrl,
	}
	c.proxy = proxyUrl
	c.logger.Debug().Any("proxy", proxyParsed.Host).Msg("SetProxy")
	return nil
}

func (c *Client) HasUUID(slice []*protos.UUID, target *protos.UUID) bool {
	for _, user := range slice {
		if user == target {
			return true
		}
	}
	return false
}