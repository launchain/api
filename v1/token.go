package v1

import (
	"net/url"

	"github.com/launchain/api"
)

// Token ...
type Token struct {
	uri string
}

// NewToken ...
func NewToken(c *api.Config) *Token {
	c.Check()
	uri := c.URI()
	return &Token{uri: uri}
}

// Generate ...
func (t *Token) Generate(userID, deviceID string) (map[string]string, error) {
	data := make(url.Values)
	data["user_id"] = []string{userID}
	data["device_id"] = []string{deviceID}

	url := t.uri + "/v1/tokens"
	out := make(map[string]string)
	err := api.PostForm(url, data, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// Parse ...
func (t *Token) Parse(token string) (out map[string]interface{}, err error) {
	url := t.uri + "/v1/tokens/" + token

	err = api.Get(url, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Remove ...
func (t *Token) Remove(token string) error {
	url := t.uri + "/v1/tokens/" + token
	return api.Delete(url)
}
