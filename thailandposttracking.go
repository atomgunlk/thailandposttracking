package thailandposttracking

import (
	"time"

	"github.com/atomgunlk/golang-common/pkg/request"
)

type (
	Client struct {
		secretToken     string
		apiToken        Token
		webhookToken    Token
		webhookCallback func()
		httpClient      request.Client
	}
	Token struct {
		Expire time.Time
		Token  string
	}
)

func isTokenExpired(token *Token) bool {
	return time.Now().After(token.Expire)
}

func (c *Client) SetApiToken(token Token) {
	c.apiToken = token
}

// New()
// You can get yourSecretToken from
// https://track.thailandpost.co.th/dashboard
// at Menu "<>สำหรับนักพัฒนา"
func New(yourSecretToken string) *Client {
	httpClient := request.NewClientWithDebug(
		true,
		request.WithRetryMax(3),
		request.WithTimeout(10*time.Second),
	)
	return &Client{
		secretToken: yourSecretToken,
		httpClient:  httpClient,
	}
}
