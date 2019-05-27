package v1

import (
	"github.com/launchain/api"
	"net/url"
)

// RRDWxAppAuth ...
type RRDWxAppAuth struct {
	uri string
}

//GetSessionKeyRequest ...
type GetSessionKeyRequest struct {
	Code string `json:"code"`
}

type WxSessionKeyResp struct {
	OpenID     string  `json:"openid"`
	SessionKey string  `json:"session_key"`
	UnionID    string  `json:"unionid"`
	ErrCode    float64 `json:"errcode"`
	ErrMsg     string  `json:"errmsg"`
}

// GetSessionKeyResponse ...
type GetSessionKeyResponse struct {
	Base
	Data WxSessionKeyResp `json:"data"`
}

//DecryptDataRequest ...
type DecryptDataRequest struct {
	SessionKey    string `json:"session_key"`
	EncryptedData string `json:"encrypted_data"`
	Iv            string `json:"iv"`
	RawData       string `json:"raw_data"`
	Sign          string `json:"sign"`
}

//DecryptDataResponse ...
type DecryptDataResponse struct {
	Base
	Data string `json:"data"`
}

// NewRRDWxAppAuth ...
func NewRRDWxAppAuth(c *api.Config) *RRDWxAppAuth {
	c.Check()
	uri := c.URI()
	return &RRDWxAppAuth{uri: uri}
}

//GetSessionKeyByCode ...
func (a *RRDWxAppAuth) GetSessionKeyByCode(req *GetSessionKeyRequest) (*GetSessionKeyResponse, error) {
	out := &GetSessionKeyResponse{}
	uri := a.uri + "/v1/rrd-wx-app/auth/code"
	data := make(url.Values)
	data.Add("code", req.Code)
	return out, api.PostForm(uri, data, &out)
}

//DecryptData ...
func (a *RRDWxAppAuth) DecryptData(req *DecryptDataRequest) (*DecryptDataResponse, error) {
	out := &DecryptDataResponse{}
	uri := a.uri + "/v1/rrd-wx-app/auth/data/decryption"
	data := make(url.Values)
	data.Add("session_key", req.SessionKey)
	data.Add("encrypted_data", req.EncryptedData)
	data.Add("iv", req.Iv)
	data.Add("raw_data", req.RawData)
	data.Add("sign", req.Sign)
	return out, api.PostForm(uri, data, &out)
}
