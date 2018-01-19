package v1

import (
	"errors"
	"fmt"

	"github.com/launchain/api"
)

// Captcha ...
type Captcha struct {
	uri string
}

// NewCaptcha ...
func NewCaptcha(c *api.Config) *Captcha {
	c.Check()
	uri := "http://" + c.Host + ":" + c.Port
	return &Captcha{uri: uri}
}

// FindOne ...
func (c *Captcha) FindOne(id, code string) error {
	if id == "" || code == "" {
		return errors.New("参数错误")
	}

	url := fmt.Sprintf("%s/v1/captcha/%s/code/%s", c.uri, id, code)
	return api.Get(url, nil)
}
