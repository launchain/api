package v1

import (
	"fmt"
	"github.com/launchain/api"
	"net/url"
)

// Session ...
type Session struct {
	uri string
}

type SignWithPhoneReq struct {
	Phone    string
	Code     string
	Type     int
	DeviceID string
	Platform int
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
func (s *Session) SignIn(phone, password, deviceID string, platform int) (map[string]interface{}, error) {
	data := make(url.Values)
	data["phone"] = []string{phone}
	data["password"] = []string{password}
	data["device_id"] = []string{deviceID}
	data["platform"] = []string{fmt.Sprintf("%d", platform)}

	url := s.uri + "/v1/sessions"
	out := make(map[string]interface{})
	err := api.PostForm(url, data, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

//SignWithPhone ...
func (s *Session) SignWithPhone(req *SignWithPhoneReq) (*SessionResponse, error) {
	data := make(url.Values)
	data["phone"] = []string{req.Phone}
	data["code"] = []string{req.Code}
	data["type"] = []string{fmt.Sprintf("%d", req.Type)}
	data["device_id"] = []string{req.DeviceID}
	data["platform"] = []string{fmt.Sprintf("%d", req.Platform)}

	url := s.uri + "/v1/sessions/phone"
	out := &SessionResponse{}
	err := api.PostForm(url, data, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// SignIn ...
func (s *Session) SignInWithEmail(email, password, deviceID string, platform int) (map[string]interface{}, error) {
	data := make(url.Values)
	data["email"] = []string{email}
	data["password"] = []string{password}
	data["device_id"] = []string{deviceID}
	data["platform"] = []string{fmt.Sprintf("%d", platform)}

	url := s.uri + "/v1/sessions/email"
	out := make(map[string]interface{})
	err := api.PostForm(url, data, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

//SignWithPhoneInternal ...
func (s *Session) SignWithPhoneInternal(req *SignWithPhoneReq) (*SessionResponse, error) {
	data := make(url.Values)
	data["phone"] = []string{req.Phone}
	data["type"] = []string{fmt.Sprintf("%d", req.Type)}
	data["device_id"] = []string{req.DeviceID}
	data["platform"] = []string{fmt.Sprintf("%d", req.Platform)}

	url := s.uri + "/v1/sessions/phone/internal"
	out := &SessionResponse{}
	err := api.PostForm(url, data, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
