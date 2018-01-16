package api

import (
	"net/url"

	"gitlab.com/launchain/session/config"
)

// Token ...
type Token struct {
	uri string
}

var t *Token

func init() {
	conf := config.NewConfig()
	uri := "http://" + conf.Token.Host + ":" + conf.Token.Port
	t = &Token{uri: uri}
}

// NewToken ...
func NewToken() *Token {
	return t
}

// Generate ...
func (t *Token) Generate(userID, deviceID string) (map[string]string, error) {
	data := make(url.Values)
	data["user_id"] = []string{userID}
	data["device_id"] = []string{deviceID}

	res, err := postForm(t.uri, data)
	if err != nil {
		return nil, err
	}

	m, _ := res.(map[string]string)
	return m, nil
}
