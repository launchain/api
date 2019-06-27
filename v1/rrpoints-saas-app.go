package v1

import (
	"fmt"
	"net/url"

	"github.com/launchain/api"
)

// RRPointsApp ...
type RRPointsApp struct {
	uri string
}

// NewRRPointsApp ...
func NewRRPointsApp(c *api.Config) *RRPointsApp {
	c.Check()
	uri := c.URI()
	return &RRPointsApp{uri: uri}
}

type AppWxNoticeRequest struct {
	TracingBase
	BillingId string `json:"billing_id"`
	TransId   string `json:"trans_id"`
}

type AppWxNoticeResponse struct {
	Base
}

// PostWxNotice ...
func (o *RRPointsApp) PostWxNotice(request *AppWxNoticeRequest) (*AppWxNoticeResponse, error) {
	out := AppWxNoticeResponse{}
	params := url.Values{}
	params.Set("billing_id", request.BillingId)
	params.Set("trans_id", request.TransId)
	apiUrl := fmt.Sprintf("%s/v1/rrpoints-saas/app/wx-notice", o.uri)
	err := api.PostFormAndTrace(request.SpanContext, apiUrl, params, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
