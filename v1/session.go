package v1

import (
	"fmt"
	"net/url"
	"github.com/launchain/api"
)

// Session ...
type Session struct {
	uri string
}

// SessionResponse ...
type SessionResponse struct {
	ID             string `json:"id"`
	UserID         string `json:"user_id"`
	Token          string `json:"token"`
	PassStatus     int    `json:"pass_status"`
	Authentication int    `json:"authentication"`
	SessionId      string `json:"session_id"`
	OpenID         string `json:"openid"`
	UnionID        string `json:"unionid"`
	RefreshToken   string `json:"refresh_token"`
	Phone          string `json:"phone"`
	WalletAddress  string `json:"wallet_address"`
}

// NewSession ...
func NewSession(c *api.Config) *Session {
	c.Check()
	uri := c.URI()
	return &Session{uri: uri}
}

// SignIn ...
func (s *Session) SignIn(phone, password, deviceID string, platform int) (*SessionResponse, error) {
	data := make(url.Values)
	data["phone"] = []string{phone}
	data["password"] = []string{password}
	data["device_id"] = []string{deviceID}
	data["platform"] = []string{fmt.Sprintf("%d", platform)}

	url := s.uri + "/v1/sessions"
	out := &SessionResponse{}
	err := api.PostForm(url, data, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

//SignWithPhoneInGolo ...
func (s *Session) SignWithPhoneInGolo(phone, code, deviceID string, platform int) (*SessionResponse, error) {
	data := make(url.Values)
	data["phone"] = []string{phone}
	data["code"] = []string{code}
	data["type"] = []string{"1"}
	data["device_id"] = []string{deviceID}

	url := s.uri + "/v1/sessions/phone"
	out := &SessionResponse{}
	err := api.PostForm(url, data, out)
	if err != nil {
		return nil, err
	}

	return out, nil

}
