package v1

import (
	"fmt"
	"net/url"

	"github.com/launchain/api"
)

// WeixinAuth ...
type WeixinAuth struct {
	uri string
}

type WxAuthGetUserInfoByCodeReq struct {
	TracingBase
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

type WxAuthSendMsgReq struct {
	TracingBase
	OpenId string
	Text   string
}

//WxAuthSendMsgRes ...
type WxAuthSendMsgRes struct {
	Code    string
	Message string
}

// NewWeixinAuth ...
func NewWeixinAuth(c *api.Config) *WeixinAuth {
	c.Check()
	uri := c.URI()
	return &WeixinAuth{uri: uri}
}

// Deprecated: use TracingGetUserInfoByCode
func (a *WeixinAuth) GetUserInfoByCode(req WxAuthGetUserInfoByCodeReq) (WxAuthGetUserInfoByCodeRes, error) {
	out := WxAuthGetUserInfoByCodeRes{}
	dstUrl := fmt.Sprintf("%s/v1/weixin/auth?code=%s", a.uri, req.Code)
	err := api.Get(dstUrl, &out)
	if err != nil {
		return WxAuthGetUserInfoByCodeRes{}, err
	}
	return out, nil
}

// Deprecated: use TracingGetRRDUserInfoByCode
func (a *WeixinAuth) GetRRDUserInfoByCode(req WxAuthGetUserInfoByCodeReq) (WxAuthGetUserInfoByCodeRes, error) {
	out := WxAuthGetUserInfoByCodeRes{}
	dstUrl := fmt.Sprintf("%s/v1/weixin/rrd/auth?code=%s", a.uri, req.Code)
	err := api.Get(dstUrl, &out)
	if err != nil {
		return WxAuthGetUserInfoByCodeRes{}, err
	}
	return out, nil
}

// TracingGetUserInfoByCode ...
func (a *WeixinAuth) TracingGetUserInfoByCode(req WxAuthGetUserInfoByCodeReq) (WxAuthGetUserInfoByCodeRes, error) {
	out := WxAuthGetUserInfoByCodeRes{}
	dstUrl := fmt.Sprintf("%s/v1/weixin/auth?code=%s", a.uri, req.Code)
	err := api.GetAndTrace(req.SpanContext, dstUrl, &out)
	if err != nil {
		return WxAuthGetUserInfoByCodeRes{}, err
	}
	return out, nil
}

// TracingGetRRDUserInfoByCode ...
func (a *WeixinAuth) TracingGetRRDUserInfoByCode(req WxAuthGetUserInfoByCodeReq) (WxAuthGetUserInfoByCodeRes, error) {
	out := WxAuthGetUserInfoByCodeRes{}
	dstUrl := fmt.Sprintf("%s/v1/weixin/rrd/auth?code=%s", a.uri, req.Code)
	err := api.GetAndTrace(req.SpanContext, dstUrl, &out)
	if err != nil {
		return WxAuthGetUserInfoByCodeRes{}, err
	}
	return out, nil
}

// TracingSendMsgByOpenIdAndText ...
func (a *WeixinAuth) TracingSendMsgByOpenIdAndText(req WxAuthSendMsgReq) (WxAuthSendMsgRes, error) {
	out := WxAuthSendMsgRes{}

	data := make(url.Values)
	data.Set("open_id", req.OpenId)
	data.Set("text", req.Text)
	dstUrl := fmt.Sprintf("%s/v1/weixin/rrd/auth/send-msg", a.uri)
	err := api.PostFormAndTrace(req.SpanContext, dstUrl, data, &out)
	if err != nil {
		return WxAuthSendMsgRes{}, err
	}
	return out, nil
}
