package v1

import (
	"github.com/launchain/api"
	"net/url"
)

// TokenFlow ...
type TokenFlow struct {
	uri string
}

// CreateTokenFlowRequest ...
type CreateTokenFlowRequest struct {
	TokenID string `json:"tokenid"`
	From    string `json:"from"`
	To      string `json:"to"`
	Value   string `json:"value"`
	TxnHash string `json:"txnhash"`
	Rule    string `json:"rule"`
	Device  string `json:"device"`
	Remark  string `json:"remark"`
}

// NewTokenFlow ...
func NewTokenFlow(c *api.Config) *TokenFlow {
	c.Check()
	uri := c.URI()
	return &TokenFlow{uri: uri}
}

// CreteTokenFlow ...
func (t *TokenFlow) CreteTokenFlow(req *CreateTokenFlowRequest) (map[string]interface{}, error) {
	uri := t.uri + "/v1/tokenflow"
	data := make(url.Values)
	data.Add("tokenid", req.TokenID)
	data.Add("from", req.From)
	data.Add("to", req.To)
	data.Add("value", req.Value)
	data.Add("txnhash", req.TxnHash)
	data.Add("rule", req.Rule)
	data.Add("device", req.Device)
	data.Add("remark", req.Remark)
	out := make(map[string]interface{})
	return out, api.PostForm(uri, data, &out)
}
