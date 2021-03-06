package v1

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/launchain/api"
)

// ERC20 ...
type ERC20 struct {
	uri string
}

// ERC20Request ...
type ERC20Request struct {
	APIKey  string
	TokenID string
	To      string
	Value   string
}

// ERC20Response ...
type ERC20Response struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
	Logo    string `json:"logo"`
	Label   string `json:"label"`
	Total   string `json:"total"`
	Address string `json:"address"`
	Type    int    `json:"type"`
}

type ERC20PayRequest struct {
	FromAddress  string `json:"from_address"`
	ToAddress    string `json:"to_address"`
	FromPhrase   string `json:"from_phrase"`
	TokenID      string `json:"token_id"`
	Value        string `json:"value"`
	FromKeystore string `json:"from_keystore"`
}

type ERC20DeductRequest struct {
	APIKey       string `json:"api_key"`
	FromAddress  string `json:"from_address"`
	FromPhrase   string `json:"from_phrase"`
	TokenID      string `json:"token_id"`
	Value        string `json:"value"`
	FromKeystore string `json:"from_keystore"`
}

//AlliancePayRequest ...
type AlliancePayRequest struct {
	TracingBase
	FromAddress  string
	ToAddress    string
	FromPhrase   string
	Value        string
	FromKeystore string
	Rule         string
	Remark       string
	Device       string
	Nameless     bool
}

//AllianceGetBalanceRequest ...
type AllianceGetBalanceRequest struct {
	TracingBase
	WalletAddress string
}

//PayTokenDetail ...
type PayTokenDetail struct {
	Token   string `json:"token"`
	Value   string `json:"value"`
	TxnHash string `json:"txnhash"`
}

//PayDetail ...
type PayInfo struct {
	Value  string           `json:"value"`
	Detail []PayTokenDetail `json:"detail"`
}

//AlliancePayResponse ...
type AlliancePayResponse struct {
	Base
	Data PayInfo `json:"data"`
}

//TokenBalance ...
type TokenBalance struct {
	Token string `json:"token"`
	Value string `json:"value"`
}

//TotalBalance ...
type TotalBalance struct {
	Value  string         `json:"value"`
	Detail []TokenBalance `json:"detail"`
}

//AllianceGetBalanceResponse ...
type AllianceGetBalanceResponse struct {
	Base
	Data TotalBalance `json:"data"`
}

// NewERC20 ...
func NewERC20(c *api.Config) *ERC20 {
	c.Check()
	uri := c.URI()
	return &ERC20{uri: uri}
}

// Transfer ...
func (u *ERC20) Transfer(req ERC20Request) (map[string]string, error) {
	apiurl := u.uri + "/v1/token/transfer"
	out := make(map[string]string)
	data := make(url.Values)
	data.Add("api_key", req.APIKey)
	data.Add("token_id", req.TokenID)
	data.Add("to", req.To)
	data.Add("value", req.Value)
	err := api.PostForm(apiurl, data, &out)

	return out, err
}

// Balance ...
func (u *ERC20) Balance(tokenID, address string) (map[string]string, error) {
	apiurl := u.uri + "/v1/token/" + tokenID + "/balance" + "?address=" + address
	out := make(map[string]string)
	err := api.Get(apiurl, &out)

	return out, err
}

// FindOne ...
func (u *ERC20) FindOne(tokenID string) (*ERC20Response, error) {
	apiurl := u.uri + "/v1/token/" + tokenID
	out := &ERC20Response{}
	err := api.Get(apiurl, out)

	return out, err
}

// Pay  ...
func (u *ERC20) Pay(req *ERC20PayRequest) (map[string]interface{}, error) {
	apiURL := u.uri + "/v1/token/payment"
	out := make(map[string]interface{})
	data := make(url.Values)
	data.Add("from_address", req.FromAddress)
	data.Add("token_id", req.TokenID)
	data.Add("to_address", req.ToAddress)
	data.Add("value", req.Value)
	data.Add("from_keystore", req.FromKeystore)
	data.Add("from_phrase", req.FromPhrase)
	err := api.PostForm(apiURL, data, &out)

	return out, err
}

// Deduct  ...
func (u *ERC20) Deduct(req *ERC20DeductRequest) (map[string]interface{}, error) {
	apiURL := u.uri + "/v1/token/deduction"
	out := make(map[string]interface{})
	data := make(url.Values)
	data.Add("from_address", req.FromAddress)
	data.Add("token_id", req.TokenID)
	data.Add("api_key", req.APIKey)
	data.Add("value", req.Value)
	data.Add("from_keystore", req.FromKeystore)
	data.Add("from_phrase", req.FromPhrase)
	err := api.PostForm(apiURL, data, &out)

	return out, err
}

// GetAllTokens ...
func (u *ERC20) GetAllTokens() ([]ERC20Response, error) {
	apiurl := u.uri + "/v1/token"
	out := make([]ERC20Response, 0)
	err := api.Get(apiurl, &out)
	if err != nil {
		return out, err
	}
	return out, nil
}

// GetTokenByID ...
func (u *ERC20) GetTokenByID(id string) (ERC20Response, error) {
	apiurl := u.uri + fmt.Sprintf("/v1/token/pk/%s", id)
	var out ERC20Response
	err := api.Get(apiurl, &out)
	if err != nil {
		return out, err
	}
	return out, nil
}

//AllianceGetBalance ...
func (u *ERC20) AllianceGetBalance(request *AllianceGetBalanceRequest) (*AllianceGetBalanceResponse, error) {
	url := u.uri + "/v1/alliance/token/balance?address=" + request.WalletAddress
	out := &AllianceGetBalanceResponse{}
	err := api.GetAndTrace(request.SpanContext, url, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

//AlliancePayment ...
func (u *ERC20) AlliancePayment(request *AlliancePayRequest) (*AlliancePayResponse, error) {
	data := make(url.Values)
	data.Add("from_address", request.FromAddress)
	data.Add("to_address", request.ToAddress)
	data.Add("value", request.Value)
	data.Add("from_phrase", request.FromPhrase)
	data.Add("from_keystore", request.FromKeystore)
	data.Add("rule", request.Rule)
	data.Add("remark", request.Remark)
	data.Add("device", request.Device)
	data.Add("nameless", strconv.FormatBool(request.Nameless))

	out := &AlliancePayResponse{}
	url := u.uri + "/v1/alliance/token/payment"
	err := api.PostFormAndTrace(request.SpanContext, url, data, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
