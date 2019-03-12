package v2

import (
	"fmt"
	"github.com/launchain/api"
	"github.com/launchain/api/v1"
)

// WeixinAuth ...
type WeixinAuth struct {
	uri string
}

// NewWeixinAuth ...
func NewWeixinAuth(c *api.Config) *WeixinAuth {
	c.Check()
	uri := c.URI()
	return &WeixinAuth{uri: uri}
}

//GetUserInfoByCode ...
func (a *WeixinAuth) GetUserInfoByCode(req v1.WxAuthGetUserInfoByCodeReq) (v1.WxAuthGetUserInfoByCodeRes,error) {
	out := v1.WxAuthGetUserInfoByCodeRes{}
	url := fmt.Sprintf("%s/v2/weixin/auth?code=%s", a.uri, req.Code)
	err := api.Get(url, &out)
	if err != nil {
		return v1.WxAuthGetUserInfoByCodeRes{}, err
	}
	return out, nil
}
