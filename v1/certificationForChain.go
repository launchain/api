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

// VinForChainRequest ...
type VinForChainRequest struct {
	Name    string
	Vin     string
	Address string
}

// IDCardForChainRequest ..
type IDCardForChainRequest struct {
	Name    string
	IDCard  string
	Address string
}

// RoleForChainRequest ...
type RoleForChainRequest struct {
	RoleID   string
	RoleName string
	Address  string
}

// CertForChainRequest ...
type CertForChainRequest struct {
	CertNo   string
	CertName string
	Address  string
}

// RecordForChainRequest ...
type RecordForChainRequest struct {
	RecordID   string
	RecordHash string
	Address    string
}

// NewCertificationForChain ...
func NewCertificationForChain(c *api.Config) *CertificationForChain {
	c.Check()
	uri := c.URI()
	return &CertificationForChain{uri: uri}
}

// CreateVinForChain ...
func (c *CertificationForChain) CreateVinForChain(req *VinForChainRequest) (map[string]interface{}, error) {
	data := make(url.Values)
	data.Add("name", req.Name)
	data.Add("address", req.Address)
	data.Add("vin", req.Vin)
	Url := fmt.Sprintf("%s%s", c.uri, "/v1/blockchain/vin-certification")
	out := make(map[string]interface{})
	return out, api.PostForm(Url, data, &out)
}

// CreateIDCardChain ...
func (c *CertificationForChain) CreateIDCardChain(req *IDCardForChainRequest) (map[string]interface{}, error) {
	data := make(url.Values)
	data.Add("name", req.Name)
	data.Add("address", req.Address)
	data.Add("id_card", req.IDCard)
	Url := fmt.Sprintf("%s%s", c.uri, "/v1/blockchain/id-certification")
	out := make(map[string]interface{})
	return out, api.PostForm(Url, data, &out)
}

// CreateRoleChain ...
func (c *CertificationForChain) CreateRoleChain(req *RoleForChainRequest) (map[string]interface{}, error) {
	data := make(url.Values)
	data.Add("role_id", req.RoleID)
	data.Add("role_name", req.RoleName)
	data.Add("address", req.Address)
	Url := fmt.Sprintf("%s%s", c.uri, "/v1/blockchain/role-certification")
	out := make(map[string]interface{})
	return out, api.PostForm(Url, data, &out)
}

// CreateRoleChain ...
func (c *CertificationForChain) CreateCertChain(req *CertForChainRequest) (map[string]interface{}, error) {
	data := make(url.Values)
	data.Add("cert_no", req.CertNo)
	data.Add("cert_name", req.CertName)
	data.Add("address", req.Address)
	Url := fmt.Sprintf("%s%s", c.uri, "/v1/blockchain/cert-certification")
	out := make(map[string]interface{})
	err := api.PostForm(Url, data, &out)
	if err != nil {
		return out, err
	}
	return out, nil
}

// CreateRoleChain ...
func (c *CertificationForChain) CreateRecordChain(req *RecordForChainRequest) (map[string]interface{}, error) {
	data := make(url.Values)
	data.Add("record_id", req.RecordID)
	data.Add("record_hash", req.RecordHash)
	data.Add("address", req.Address)
	Url := fmt.Sprintf("%s%s", c.uri, "/v1/blockchain/record-certification")
	out := make(map[string]interface{})
	err := api.PostForm(Url, data, &out)
	if err != nil {
		return out, err
	}
	return out, nil
}
