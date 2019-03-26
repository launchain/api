package v1

import (
	"fmt"
	"github.com/launchain/api"
)

// WeixinAuth ...
type WeixinAuth struct {
	uri string
}

type WxAuthGetUserInfoByCodeReq struct {
	Code string
}

//WeixinAuthGetUserInfoRes ...
type WxAuthGetUserInfoByCodeRes struct {
	Code    string
	Message string
	Data    WeixinAuthUserInfoRes
}

//WeixinAuthUserInfoRes ...
type WeixinAuthUserInfoRes struct {
	OpenID     string   `json:"open_id"`
	NickName   string   `json:"nick_name"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgURL string   `json:"head_imgurl"`
	Privilege  []string `json:"privilege"`
	UnionID    string   `json:"union_id"`
}

// NewWeixinAuth ...
func NewWeixinAuth(c *api.Config) *WeixinAuth {
	c.Check()
	uri := c.URI()
	return &WeixinAuth{uri: uri}
}

//GetUserInfoByCode ...
func (a *WeixinAuth) GetUserInfoByCode(req WxAuthGetUserInfoByCodeReq) (WxAuthGetUserInfoByCodeRes,error) {
	out := WxAuthGetUserInfoByCodeRes{}
	url := fmt.Sprintf("%s/v1/weixin/auth?code=%s", a.uri, req.Code)
	err := api.Get(url, &out)
	if err != nil {
		return WxAuthGetUserInfoByCodeRes{}, err
	}
	return out, nil
}

//GetRRDUserInfoByCode ...
func (a *WeixinAuth) GetRRDUserInfoByCode(req WxAuthGetUserInfoByCodeReq) (WxAuthGetUserInfoByCodeRes,error) {
	out := WxAuthGetUserInfoByCodeRes{}
	url := fmt.Sprintf("%s/v1/weixin/rrd/auth?code=%s", a.uri, req.Code)
	err := api.Get(url, &out)
	if err != nil {
		return WxAuthGetUserInfoByCodeRes{}, err
	}
	return out, nil
}