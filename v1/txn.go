package v1

import (
	"errors"
	"net/url"
	"time"

	"github.com/launchain/api"
)

// TXN ...
type TXN struct {
	uri string
}

// NewTXN ..
func NewTXN(c *api.Config) *TXN {
	c.Check()
	uri := c.URI()
	return &TXN{uri: uri}
}

// TXNCreateRequest ...
type TXNCreateRequest struct {
	Hash      string    `json:"hash"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	Value     string    `json:"value"`
	Gas       string    `json:"gas"`
	GasPrice  string    `json:"gasPrice"`
	Timestamp time.Time `json:"timestamp"`
}

// Create ...
func (t *TXN) Create(req *TXNCreateRequest) (map[string]interface{}, error) {
	if req.Hash == "" || req.From == "" || req.To == "" {
		return nil, errors.New("参数错误")
	}

	if req.Timestamp.IsZero() {
		req.Timestamp = time.Now().UTC()
	}

	data := make(url.Values)
	data["hash"] = []string{req.Hash}
	data["from"] = []string{req.From}
	data["to"] = []string{req.To}
	data["value"] = []string{req.Value}
	data["gas"] = []string{req.Gas}
	data["gas_price"] = []string{req.GasPrice}
	data["timestamp"] = []string{req.Timestamp.Format(time.RFC3339)}
	url := t.uri + "/v1/txn"
	out := make(map[string]interface{})
	err := api.PostForm(url, data, &out)
	if err != nil {
		return nil, err
	}
	return out, err
}
