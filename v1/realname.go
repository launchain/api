package v1

import (
	"fmt"
	"github.com/launchain/api"
	"net/url"
)

// RealNameAuth ...
type RealNameAuth struct {
	uri string
}

// RealNameAuthResponse ...
type RealNameAuthResponse struct {
	Name   string `json:"name"`
	IDCard string `json:"idcard"`
	Status int    `json:"status"` //校验状态 10=一致 20=不一致
}

// NewRealNameAuth ...
func NewRealNameAuth(c *api.Config) *RealNameAuth {
	c.Check()
	uri := c.URI()
	return &RealNameAuth{uri: uri}
}

// CheckRealNameAuth ...
func (r *RealNameAuth) CheckRealNameAuth(name, idcard string) (*RealNameAuthResponse, error) {
	data := make(url.Values)
	data.Add("name", name)
	data.Add("idcard", idcard)
	Url := fmt.Sprintf("%s%s", r.uri, "/v1/realname-auth")
	out := RealNameAuthResponse{}
	err := api.PostForm(Url, data, &out)
	if err != nil {
		return &out, err
	}
	return &out, nil
}
