package v1

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/launchain/api"
)

// SMS ...
type SMS struct {
	uri string
}

// NewSMS ..
func NewSMS(c *api.Config) *SMS {
	c.Check()
	uri := c.URI()
	return &SMS{uri: uri}
}

// SMSCode ...
type SMSCode struct {
	ID        string    `json:"_id"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Phone     string    `json:"phone"`
	Type      int       `json:"type"`
	Code      string    `json:"code"`
	ExpiredAt time.Time `json:"expired_at"`
}

// FindCode ...
func (s *SMS) FindCode(phone, code string) (*SMSCode, error) {
	if phone == "" || code == "" {
		return nil, errors.New("参数错误")
	}

	sc := &SMSCode{}
	url := fmt.Sprintf("%s/v1/sms/%s/code/%s", s.uri, phone, code)
	err := api.Get(url, sc)
	if err != nil {
		return nil, err
	}
	return sc, nil
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
