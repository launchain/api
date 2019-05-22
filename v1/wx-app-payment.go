package v1

import (
	"fmt"
	"github.com/launchain/api"
	"net/url"
)

// RRDWxAppPayment ...
type RRDWxAppPayment struct {
	uri string
}

//RRDWxAppPaymentReq ...
type RRDWxAppPaymentReq struct {
	OrderId     string
	Platform    string
	GoodsDesc   string
	TotalAmount string
	NotifyUrl   string
	TradeType   string
	OpenId      string
}

//RRDWxAppWithdrawReq ...
type RRDWxAppWithdrawReq struct {
	Openid string
	Amount int
}

// NewRRDWxAppPayment ...
func NewRRDWxAppPayment(c *api.Config) *RRDWxAppPayment {
	c.Check()
	uri := c.URI()
	return &RRDWxAppPayment{uri: uri}
}

// RRDWxAppPrepay ...
func (v *RRDWxAppPayment) RRDWxAppPrepay(req *RRDWxAppPaymentReq) (map[string]interface{}, error) {
	data := make(url.Values)
	data["order_id"] = []string{req.OrderId}
	data["platform"] = []string{req.Platform}
	data["goods_desc"] = []string{req.GoodsDesc}
	data["total_amount"] = []string{req.TotalAmount}
	data["notify_url"] = []string{req.NotifyUrl}
	data["trade_type"] = []string{req.TradeType}
	data["open_id"] = []string{req.OpenId}
	url := v.uri + "/v1/rrd-wx-app/payment/prepay"
	out := make(map[string]interface{})
	err := api.PostForm(url, data, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

//RRDWxAppWithdraw ...
func (v *RRDWxAppPayment) RRDWxAppWithdraw(req *RRDWxAppWithdrawReq) (*Base, error) {
	data := make(url.Values)
	data["openid"] = []string{req.Openid}
	data["amount"] = []string{fmt.Sprintf("%d", req.Amount)}
	url := v.uri + "/v1/rrd-wx-app/payment/withdraw"
	out := &Base{}
	err := api.PostForm(url, data, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
