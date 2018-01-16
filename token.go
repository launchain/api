package api

import (
	"net/url"
)

// Token ...
type Token struct {
	uri string
}

// NewToken ...
func NewToken(c *Config) *Token {
	uri := "http://" + c.Host + ":" + c.Port
	return &Token{uri: uri}
}

// Generate ...
func (t *Token) Generate(userID, deviceID string) (map[string]string, error) {
	data := make(url.Values)
	data["user_id"] = []string{userID}
	data["device_id"] = []string{deviceID}

	url := t.uri + "/v1/tokens"
	out := make(map[string]string)
	err := postForm(url, data, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
