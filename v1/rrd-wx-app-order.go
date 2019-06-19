package v1

import (
	"fmt"
	"github.com/launchain/api"
)

// RRDWxAppOrder ...
type RRDWxAppOrder struct {
	uri string
}

//GetPermissionByUserIDRequest ...
type GetPermissionByUserIDRequest struct {
	TracingBase
	UserID string
}

// GetPermissionByUserIDResponse ...
type GetPermissionByUserIDResponse struct {
	Code    string          `json:"code"`
	Message string          `json:"message"`
	Data    map[string]bool `json:"data"`
}

// NewRRDWxAppOrder ...
func NewRRDWxAppOrder(c *api.Config) *RRDWxAppOrder {
	c.Check()
	uri := c.URI()
	return &RRDWxAppOrder{uri: uri}
}

//GetPermissionByManualID ...
func (o *RRDWxAppOrder) GetPermissionByUserID(request *GetPermissionByUserIDRequest) (*GetPermissionByUserIDResponse, error) {
	out := &GetPermissionByUserIDResponse{}
	url := fmt.Sprintf("%s/v1/rrd-wx-app/order/permission/%s", o.uri, request.UserID)
	err := api.GetAndTrace(request.SpanContext, url, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
