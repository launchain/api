package v1

import (
	"fmt"
	"github.com/launchain/api"
	"time"
)

// WxAppCar ...
type WxAppCar struct {
	uri string
}

// GetManualResponse ...
type GetManualResponse struct {
	Code    string     `json:"code"`
	Message string     `json:"message"`
	Data    ManualInfo `json:"data"`
}

//ManualInfo ...
type ManualInfo struct {
	ID        string    `json:"id"`
	Brand     string    `json:"brand"`
	Series    string    `json:"series"`
	Year      string    `json:"year"`
	Type      int       `json:"type"`
	FirstMenu string    `json:"first_menu"`
	FirstPage string    `json:"first_page"`
	TotalPage int       `json:"total_page"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewWxAppCar ...
func NewWxAppCar(c *api.Config) *WxAppCar {
	c.Check()
	uri := c.URI()
	return &WxAppCar{uri: uri}
}

//GetManualByID ...
func (u *WxAppCar) GetManualByID(id string) (*GetManualResponse, error) {
	out := &GetManualResponse{}
	url := fmt.Sprintf("%s/v1/rrd-wx-app/car/info/%s", u.uri, id)
	err := api.Get(url, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
