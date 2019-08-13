package v1

import (
	"github.com/launchain/api"
	"net/url"
)

// RegisterReward ...
type RegisterReward struct {
	uri string
}

// CreateTokenFlowRequest ...
type RegisterRewardRequest struct {
	UserID   string `json:"user_id"`
	OpenID   string `json:"openid"`
	Credits  string `json:"credits"`
	WallAddr string `json:"wallet_addr"`
}

func NewRegisterReward(c *api.Config) *RegisterReward {
	c.Check()
	uri := c.URI()
	return &RegisterReward{uri: uri}
}

func (t *RegisterReward) AddRegisterReward(req *RegisterRewardRequest) (map[string]interface{}, error) {
	uri := t.uri + "/v1/renren-points/register-reword"
	data := make(url.Values)
	data.Add("user_id", req.UserID)
	data.Add("openid", req.OpenID)
	data.Add("credits", req.Credits)
	data.Add("wallet_addr", req.WallAddr)
	out := make(map[string]interface{})
	return out, api.PostForm(uri, data, &out)
}
