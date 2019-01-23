package v2

import (
	"net/url"
	"github.com/launchain/api"
)

// Violation ...
type Violation struct {
	uri string
}

type ViolationQueryReq struct {
	Hphm     string
	Hpzl     string
	EngineNo string
	ClassNo  string
	UserId   string
}

// NewViolation ...
func NewViolation(c *api.Config) *Violation {
	c.Check()
	uri := c.URI()
	return &Violation{uri: uri}
}

// Query ...
func (v *Violation) Query(req *ViolationQueryReq) (map[string]interface{}, error) {
	data := make(url.Values)
	data["userid"] = []string{req.UserId}
	data["hphm"] = []string{req.Hphm}
	data["hpzl"] = []string{req.Hpzl}
	data["engineno"] = []string{req.EngineNo}
	data["classno"] = []string{req.ClassNo}

	url := v.uri + "/v2/golo/violation/query"
	out := make(map[string]interface{})
	err := api.PostForm(url, data, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
