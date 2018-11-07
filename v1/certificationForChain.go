package v1

import (
	"fmt"
	"github.com/launchain/api"
	"net/url"
)

// Certification ...
type Certification struct {
	uri string
}

// CertificationRequest ..
type CertificationRequest struct {
	Name    string
	Vin     string
	Address string
}

// NewCertification ...
func NewCertification(c *api.Config) *Certification {
	c.Check()
	uri := c.URI()
	return &Certification{uri: uri}
}

// CreateChain ...
func (c *Certification) CreateChain(req *CertificationRequest) (map[string]interface{}, error) {
	data := make(url.Values)
	data.Add("name", req.Name)
	data.Add("address", req.Address)
	data.Add("vin", req.Vin)
	Url := fmt.Sprintf("%s%s", c.uri, "/v1/car/certification")
	out := make(map[string]interface{})
	return out, api.PostForm(Url, data, &out)
}
