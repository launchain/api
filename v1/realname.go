package v1

import (
	"fmt"
	"github.com/launchain/api"
	"net/url"
)

// RealNamAuth ...
type RealNamAuth struct {
	uri string
}

// RealNamAuthResponse ...
type RealNamAuthResponse struct {
	Name   string `json:"name"`
	IDCard string `json:"idcard"`
	Status int    `json:"status"` //校验状态 10=一致 20=不一致
}

// NewRealNamAuth ...
func NewRealNamAuth(c *api.Config) *RealNamAuth {
	c.Check()
	uri := c.URI()
	return &RealNamAuth{uri: uri}
}

// RealNameAuth ...
func (r *RealNamAuth) CheckRealNameAuth(name, idcard string) (*RealNamAuthResponse, error) {
	data := make(url.Values)
	data.Add("name", name)
	data.Add("idcard", idcard)
	Url := fmt.Sprintf("%s%s", r.uri, "/v1/car/certification")
	out := RealNamAuthResponse{}
	err := api.PostForm(Url, data, &out)
	if err != nil {
		return &out, err
	}
	return &out, nil
}
