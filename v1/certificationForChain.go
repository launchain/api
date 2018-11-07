package v1

import (
	"fmt"
	"github.com/launchain/api"
	"net/url"
)

// CertificationForChain ...
type CertificationForChain struct {
	uri string
}

// CertificationRequest ..
type CertificationForChainRequest struct {
	Name    string
	Vin     string
	Address string
}

// NewCertificationForChain ...
func NewCertificationForChain(c *api.Config) *CertificationForChain {
	c.Check()
	uri := c.URI()
	return &CertificationForChain{uri: uri}
}

// CreateChain ...
func (c *CertificationForChain) CreateChain(req *CertificationForChainRequest) (map[string]interface{}, error) {
	data := make(url.Values)
	data.Add("name", req.Name)
	data.Add("address", req.Address)
	data.Add("vin", req.Vin)
	Url := fmt.Sprintf("%s%s", c.uri, "/v1/car/certification")
	out := make(map[string]interface{})
	return out, api.PostForm(Url, data, &out)
}
