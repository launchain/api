package v2

import (
	"github.com/launchain/api"
	"net/url"
)

// BuriedPoint ...
type BuriedPoint struct {
	uri string
}

//InsertRecordReq ...
type InsertRecordReq struct {
	FirstLevel  string
	SecondLevel string
	ApiID       string
	RecordType  string
}

// NewBuriedPoint ...
func NewBuriedPoint(c *api.Config) *BuriedPoint {
	c.Check()
	uri := c.URI()
	return &BuriedPoint{uri: uri}
}

//InsertRecord ...
func (p *BuriedPoint) InsertRecord(req *InsertRecordReq) (map[string]interface{}, error) {
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
