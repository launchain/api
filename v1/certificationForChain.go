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
	Name        string
	VinOrIDCard string
	Address     string
}

// NewCertificationForChain ...
func NewCertificationForChain(c *api.Config) *CertificationForChain {
	c.Check()
	uri := c.URI()
	return &CertificationForChain{uri: uri}
}

// CreateCarChain ...
func (c *CertificationForChain) CreateCarChain(req *CertificationForChainRequest) (map[string]interface{}, error) {
	data := make(url.Values)
	data.Add("name", req.Name)
	data.Add("address", req.Address)
	data.Add("vin", req.VinOrIDCard)
	Url := fmt.Sprintf("%s%s", c.uri, "/v1/blockchain/vin-certification")
	out := make(map[string]interface{})
	return out, api.PostForm(Url, data, &out)
}

// CreateIDCardChain ...
func (c *CertificationForChain) CreateIDCardChain(req *CertificationForChainRequest) (map[string]interface{}, error) {
	data := make(url.Values)
	data.Add("name", req.Name)
	data.Add("address", req.Address)
	data.Add("id_card", req.VinOrIDCard)
	Url := fmt.Sprintf("%s%s", c.uri, "/v1/blockchain/id-certification")
	out := make(map[string]interface{})
	return out, api.PostForm(Url, data, &out)
}
