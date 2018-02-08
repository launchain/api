package v1

import (
	"errors"
	"net/url"

	"github.com/launchain/api"
)

// TXN ...
type TXN struct {
	uri string
}

// NewTXN ..
func NewTXN(c *api.Config) *TXN {
	c.Check()
	uri := "http://" + c.Host + ":" + c.Port
	return &TXN{uri: uri}
}

// TXNCreateRequest ...
type TXNCreateRequest struct {
	Address  string
	From     string
	To       string
	Value    string
	Gas      string
	GasPrice string
}

// Create ...
func (t *TXN) Create(req *TXNCreateRequest) (map[string]string, error) {
	if req.From == "" || req.To == "" {
		return nil, errors.New("参数错误")
	}

	data := make(url.Values)
	data["address"] = []string{req.Address}
	data["from"] = []string{req.From}
	data["to"] = []string{req.To}
	data["value"] = []string{req.Value}
	data["gas"] = []string{req.Gas}
	data["gas_price"] = []string{req.GasPrice}
	url := t.uri + "/v1/txn"
	out := make(map[string]string)
	err := api.PostForm(url, data, &out)
	if err != nil {
		return nil, err
	}
	return out, err
}
