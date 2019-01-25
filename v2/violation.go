package v2

import (
	"github.com/launchain/api/v1"
	"net/url"
	"github.com/launchain/api"
)

// Query ...
func (v *v1.Violation) Query(req *v1.ViolationQueryReq) (map[string]interface{}, error) {
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
