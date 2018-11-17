package v1

import (
	"github.com/launchain/api"
	"net/url"
)

// Erc20Payment ...
type Erc20Payment struct {
	uri string
}

// Erc20PaymentRequest ...
type Erc20PaymentRequest struct {
	FromKeystore string
	FromPhrase   string
	FromAddress  string
	ToAddress    string
	TokenID      string
	Value        string
}

// NewErc20Payment
func NewErc20Payment(c *api.Config) *Erc20Payment {
	c.Check()
	uri := c.URI()
	return &Erc20Payment{uri: uri}
}

// Payment ...
func (p *Erc20Payment) Payment(req *Erc20PaymentRequest) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	data := make(url.Values)
	data.Add("from_keystore", req.FromKeystore)
	data.Add("from_phrase", req.FromPhrase)
	data.Add("from_address", req.FromAddress)
	data.Add("to_address", req.ToAddress)
	data.Add("token_id", req.TokenID)
	data.Add("value", req.Value)
	url := p.uri + "/v1/token/payment"
	err := api.PostForm(url, data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
