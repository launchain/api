package v2

import (
	"net/url"
	"github.com/launchain/api"
	"fmt"
)

// Violation ...
type WeixinPayment struct {
	uri string
}

type WeixinPaymentReq struct {
	OrderId     string
	Platform    string
	GoodsDesc   string
	TotalAmount string
	NotifyUrl   string
	TradeType   string
	OpenId      string
}

// NewWeixinPayment ...
func NewWeixinPayment(c *api.Config) *WeixinPayment {
	c.Check()
	uri := c.URI()
	return &WeixinPayment{uri: uri}
}

// Query ...
func (v *WeixinPayment) Prepay(req *WeixinPaymentReq) (map[string]interface{}, error) {
	data := make(url.Values)
	data["order_id"] = []string{req.OrderId}
	data["platform"] = []string{req.Platform}
	data["goods_desc"] = []string{req.GoodsDesc}
	data["total_amount"] = []string{req.TotalAmount}
	data["notify_url"] = []string{req.NotifyUrl}
	data["trade_type"] = []string{req.TradeType}
	data["open_id"] = []string{req.OpenId}
	fmt.Println("xx", req.GoodsDesc, req.Platform)
	url := v.uri + "/v2/weixin/payment/prepay"
	out := make(map[string]interface{})
	err := api.PostForm(url, data, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
