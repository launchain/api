package v1

import (
	"fmt"
	url2 "net/url"

	"github.com/launchain/api"
)

type WxPrePayData struct {
	AppId      string `json:"app_id"`
	MchId      string `json:"mch_id"`
	NonceStr   string `json:"nonce_str"`
	PrepayId   string `json:"prepay_id"`
	PackageStr string `json:"package_str"`
	Timestamp  string `json:"timestamp"`
	Sign       string `json:"sign"`
}

// RRPointsWxPay ...
type RRPointsWxPay struct {
	uri string
}

// GetPrePayHeadRequest ...
type GetPrePayHeadRequest struct {
	TracingBase
	UId   string
	Ip    string
	Desc  string
	Price string
}

// GetPrePayHeadResponse ...
type GetPrePayHeadResponse struct {
	WxPrePayData
}

// GetOrderQueryRequest ...
type GetOrderQueryRequest struct {
	TracingBase
	UId     string
	TransId string
}

// GetOrderQueryResponse ...
type GetOrderQueryResponse struct {
	UId     string `json:"uid"`
	TransId string `json:"trans_id"`
	Paid    bool   `json:"paid"`
}

// NewRRPointsWxPay ...
func NewRRPointsWxPay(c *api.Config) *RRPointsWxPay {
	c.Check()
	uri := c.URI()
	return &RRPointsWxPay{uri: uri}
}

// GetPrePayHead ...
func (o *RRPointsWxPay) GetPrePayHead(request *GetPrePayHeadRequest) (*GetPrePayHeadResponse, error) {
	out := GetPrePayHeadResponse{}
	params := url2.Values{}
	params.Set("uid", request.UId)
	params.Set("desc", request.Desc)
	params.Set("price", request.Price)
	params.Set("user_ip", request.Ip)
	url := fmt.Sprintf("%s/v1/rrpoints-saas/wxpay/prepayheader?%v", o.uri, params.Encode())
	err := api.GetAndTrace(request.SpanContext, url, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// GetOrderQuery ...
func (o *RRPointsWxPay) GetOrderQuery(request *GetOrderQueryRequest) (*GetOrderQueryResponse, error) {
	out := GetOrderQueryResponse{}
	params := url2.Values{}
	params.Set("uid", request.UId)
	params.Set("trans_id", request.TransId)
	url := fmt.Sprintf("%s/v1/rrpoints-saas/wxpay/verify-pay-result?%v", o.uri, params.Encode())
	err := api.GetAndTrace(request.SpanContext, url, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
