package v1

import (
	"errors"
	"fmt"

	"github.com/launchain/api"
)

// SMS ...
type SMS struct {
	uri string
}

// NewSMS ..
func NewSMS(c *api.Config) *SMS {
	c.Check()
	uri := "http://" + c.Host + ":" + c.Port
	return &SMS{uri: uri}
}

// FindCode ...
func (s *SMS) FindCode(phone, code string) error {
	if phone == "" || code == "" {
		return errors.New("参数错误")
	}

	url := fmt.Sprintf("%s/v1/sms/%s/code/%s", s.uri, phone, code)
	return api.Get(url, nil)
}
