package v1

import (
	"github.com/launchain/api"
	"net/url"
)

// OpenApiUser ...
type OpenapiUser struct {
	uri string
}

// NewOpenapiUser ...
func NewOpenapiUser(c *api.Config) *OpenapiUser {
	c.Check()
	uri := c.URI()
	return &OpenapiUser{uri: uri}
}

//OpenapiUserCheckPassReq ...
type OpenapiUserCheckPasswdReq struct {
	UserName string
	Passwd   string
}

//
type OpenapiUserCheckPasswdResData struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
}

type OpenapiUserCheckPasswdRes struct {
	Base
	Data OpenapiUserCheckPasswdResData `json:"data"`
}

//CheckPasswd ...
func (ou *OpenapiUser) CheckPasswd(req *OpenapiUserCheckPasswdReq) (OpenapiUserCheckPasswdRes, error) {
	data := make(url.Values)
	data["user_name"] = []string{req.UserName}
	data["passwd"] = []string{req.Passwd}

	url := ou.uri + "/v1/platform/user/passwd/check"
	out := OpenapiUserCheckPasswdRes{}
	err := api.PostForm(url, data, &out)
	if err != nil {
		return OpenapiUserCheckPasswdRes{}, err
	}

	return out, nil
}
