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
	uri := c.URI()
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

// Remove ...
func (c *Captcha) Remove(id string) error {
	if id == "" {
		return errors.New("参数错误")
	}

	url := c.uri + "/v1/captcha/" + id
	return api.Delete(url)
}
