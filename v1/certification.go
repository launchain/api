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

// CreateCertificationRequest ...
type CreateCertificationRequest struct {
	UserID        string
	Name          string
	WalletAddress string
	Idcard        string
}

//BindVinRequest ...
type BindVinRequest struct {
	UserID string
	Name   string
	Vin    string
}

// NewCertification ...
func NewCertification(c *api.Config) *Certification {
	c.Check()
	uri := c.URI()
	return &Certification{uri: uri}
}

// CreateCertification ...
func (c *Certification) CreateCertification(req *CreateCertificationRequest) (map[string]interface{}, error) {
	data := make(url.Values)
	data.Add("name", req.Name)
	data.Add("wallet_address", req.WalletAddress)
	data.Add("idcard", req.Idcard)
	Url := fmt.Sprintf("%s/v1/identity/certification/%s", c.uri, req.UserID)
	out := make(map[string]interface{})
	return out, api.PostForm(Url, data, &out)
}

//BindVin ...
func (c *Certification) BindVin(req *BindVinRequest) (map[string]interface{}, error) {
	data := make(url.Values)
	data.Add("name", req.Name)
	data.Add("vin", req.Vin)
	Url := fmt.Sprintf("%s/v1/car/vin/%s", c.uri, req.UserID)
	out := make(map[string]interface{})
	return out, api.PostForm(Url, data, &out)
}
