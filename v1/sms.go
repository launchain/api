package v1

import (
	"errors"
	"fmt"
	"net/url"

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

// NewWarning ...
func (s *SMS) NewWarning(parter string, level int) error {
	if parter == "" {
		return errors.New("参数错误")
	}

	data := make(url.Values)
	data["partner_desc"] = []string{parter}
	data["level"] = []string{fmt.Sprintf("%d", level)}
	url := s.uri + "/v1/sms/warning"
	return api.PostForm(url, data, nil)
}
