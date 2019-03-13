package v1

import (
	"fmt"
	"github.com/launchain/api"
	"net/url"
	"time"
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

//GetTokenFlowReq ...
type GetTokenFlowReq struct {
	TokenID string
	Address string
	Page    int
	Limit   int
}

type TokenFlowSchema struct {
	TokenID   string    `json:"tokenid"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	Value     string    `json:"value"`
	TxnHash   string    `json:"txnhash"`
	Rule      string    `json:"rule"`
	Device    string    `json:"device"`
	Remark    string    `json:"remark"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TokenFlowsRes struct {
	Count int `json:"count"`
	Txn   []TokenFlowSchema
}

type GetTokenFlowRes struct {
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Data    TokenFlowsRes `json:"data"`
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

//GetTokenFlow ...
func (t *TokenFlow) GetTokenFlow(req *GetTokenFlowReq) (GetTokenFlowRes, error) {
	out := GetTokenFlowRes{}
	url := fmt.Sprintf("%s/v1/tokenflow/token/%s/address/%s?page=%d&limit=%d",
		t.uri, req.TokenID, req.Address, req.Page, req.Limit)
	err := api.Get(url, &out)
	if err != nil {
		return GetTokenFlowRes{}, err
	}
	return out, nil
}
