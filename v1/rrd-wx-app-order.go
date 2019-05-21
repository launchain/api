package v1

import (
	"fmt"
	"github.com/launchain/api"
)

// RRDWxAppOrder ...
type RRDWxAppOrder struct {
	uri string
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
func (o *RRDWxAppOrder) GetPermissionByUserID(userid string) (*GetPermissionByUserIDResponse, error) {
	out := &GetPermissionByUserIDResponse{}
	url := fmt.Sprintf("%s/v1/rrd-wx-app/order/permission/%s", o.uri, userid)
	err := api.Get(url, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
