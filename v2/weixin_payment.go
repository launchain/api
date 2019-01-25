package v2

import (
	"github.com/launchain/api/v1"
	"net/url"
	"github.com/launchain/api"
	"fmt"
)


// Query ...
func (v *v1.WeixinPayment) Prepay(req *v1.WeixinPaymentReq) (map[string]interface{}, error) {
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
