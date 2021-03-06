package v1

import (
	"github.com/launchain/api"
	"net/url"
)

// Vin ...
type Vin struct {
	uri string
}

// VinRequest  ...
type VinRequest struct {
	Img    string
	UserId string
}

//VinInfo ...
type VinInfo struct {
	Name          string `json:"name"`
	CarVin        string `json:"car_vin"`
	LicenseNo     string `json:"license_no"`
	BrandModelNum string `json:"brand_model_num"`
	EngineNo      string `json:"engine_no"`
	Url           string `json:"url"`
}

//VinResponse ...
type VinResponse struct {
	Base
	Data VinInfo `json:"data"`
}

// NewSession ...
func NewVin(c *api.Config) *Vin {
	c.Check()
	uri := c.URI()
	return &Vin{uri: uri}
}

//CheckVin ...
func (v *Vin) CheckVin(req *VinRequest) (*VinResponse, error) {
	data := make(url.Values)
	data.Add("img", req.Img)
	data.Add("user_id", req.UserId)
	url := v.uri + "/v1/launchain/ocr/vin"
	out := &VinResponse{}
	err := api.PostForm(url, data, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
