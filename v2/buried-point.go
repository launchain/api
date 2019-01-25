package v2

import (
	"github.com/launchain/api"
	"net/url"
	"github.com/launchain/api/v1"
)

//InsertRecord ...
func (p *v1.BuriedPoint) InsertRecord(req *v1.InsertRecordReq) (map[string]interface{}, error) {
	data := make(url.Values)
	data["first_level"] = []string{req.FirstLevel}
	data["second_level"] = []string{req.SecondLevel}
	data["api_id"] = []string{req.ApiID}
	data["record_type"] = []string{req.RecordType}

	url := p.uri + "/v2/golo-buried-point-record"
	out := make(map[string]interface{})
	err := api.PostForm(url, data, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
