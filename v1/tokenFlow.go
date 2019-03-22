package v1

import (
	"fmt"
	"net/url"
	"time"

	"github.com/launchain/api"
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

// GetTokenFlowReq ...
type GetTokenFlowReq struct {
	TokenID string
	Address string
	Page    int
	Limit   int
}

// TokenFlowSchema ...
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

// TokenFlowsRes ...
type TokenFlowsRes struct {
	Count int `json:"count"`
	Txn   []TokenFlowSchema
}

// GetTokenFlowRes ...
type GetTokenFlowRes struct {
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Data    TokenFlowsRes `json:"data"`
}

// GetFlowCountByTimeReq ...
type GetFlowCountByTimeReq struct {
	TokenID   string
	StartTime string
	EndTime   string
}

// GetFlowsByTimeReq ...
type GetFlowsByTimeReq struct {
	GetFlowCountByTimeReq
	Count int
	Skip  int
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

// GetTokenFlow ...
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

// GetFlowCountByTime ...
func (t *TokenFlow) GetFlowCountByTime(req *GetFlowCountByTimeReq) (*GetTokenFlowRes, error) {
	out := &GetTokenFlowRes{}
	url := fmt.Sprintf("%s/v1/tokenflow/token/%s/count?start_time=%s&end_time=%s",
		t.uri, req.TokenID, req.StartTime, req.EndTime)

	err := api.Get(url, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetFlowsByTime ...
func (t *TokenFlow) GetFlowsByTime(req *GetFlowsByTimeReq) (*GetTokenFlowRes, error) {
	out := &GetTokenFlowRes{}
	url := fmt.Sprintf("%s/v1/tokenflow/token/%s/txns?start_time=%s&end_time=%s&skip=%d&count=%d",
		t.uri, req.TokenID, req.StartTime, req.EndTime, req.Skip, req.Count)
	err := api.Get(url, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
