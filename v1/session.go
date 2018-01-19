package v1

import (
	"net/url"
	"time"

	"github.com/launchain/api"
)

// Session ...
type Session struct {
	uri string
}

// SessionResponse ...
type SessionResponse struct {
	ID        string    `json:"_id"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    string    `json:"user_id"`
	DeviceID  string    `json:"device_id"`
	Platform  int       `json:"platform"`
	Token     string    `json:"token"`
}

// NewSession ...
func NewSession(c *api.Config) *Session {
	c.Check()
	uri := "http://" + c.Host + ":" + c.Port
	return &Session{uri: uri}
}

// SignIn ...
func (s *Session) SignIn(phone, deviceID, platform string) (*SessionResponse, error) {
	data := make(url.Values)
	data["phone"] = []string{phone}
	data["device_id"] = []string{deviceID}
	data["platform"] = []string{platform}

	url := s.uri + "/v1/sessions"
	out := &SessionResponse{}
	err := api.PostForm(url, data, out)
	if err != nil {
		return nil, err
	}

	return out, nil

}
