package v1

import (
	"net/url"

	"github.com/launchain/api"
)

// ERC20 ...
type ERC20 struct {
	uri string
}

// REC20Request ...
type REC20Request struct {
	APIKey  string
	TokenID string
	To      string
	Value   string
}

// NewERC20 ...
func NewERC20(c *api.Config) *ERC20 {
	c.Check()
	uri := c.URI()
	return &ERC20{uri: uri}
}

// Transfer ...
func (u *ERC20) Transfer(req REC20Request) (map[string]string, error) {
	apiurl := u.uri + "/v1/token/transfer"
	out := make(map[string]string)
	data := make(url.Values)
	data.Add("api_key", req.APIKey)
	data.Add("token_id", req.TokenID)
	data.Add("to", req.To)
	data.Add("value", req.Value)
	err := api.PostForm(apiurl, data, &out)

	return out, err
}

// Balance ...
func (u *ERC20) Balance(tokenID, address string) (map[string]string, error) {
	apiurl := u.uri + "/v1/token/" + tokenID + "?address=" + address
	out := make(map[string]string)
	err := api.Get(apiurl, &out)

	return out, err
}
