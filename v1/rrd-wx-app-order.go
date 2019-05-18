package v1

import (
	"fmt"
	"github.com/launchain/api"
)

// RRDWxAppOrder ...
type RRDWxAppOrder struct {
	uri string
}

// GetPermissionByManualIDResponse ...
type GetPermissionByManualIDResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    bool   `json:"data"`
}

// NewRRDWxAppOrder ...
func NewRRDWxAppOrder(c *api.Config) *RRDWxAppOrder {
	c.Check()
	uri := c.URI()
	return &RRDWxAppOrder{uri: uri}
}

//GetPermissionByManualID ...
func (o *RRDWxAppOrder) GetPermissionByManualID(manualid, userid string) (*GetPermissionByManualIDResponse, error) {
	out := &GetPermissionByManualIDResponse{}
	url := fmt.Sprintf("%s/v1/rrd-wx-app/order/permission/%s/%s", o.uri, manualid, userid)
	err := api.Get(url, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
