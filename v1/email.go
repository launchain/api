package v1

import (
	"errors"
	"fmt"

	"github.com/launchain/api"
)

// Email ...
type Email struct {
	uri string
}

// NewEmail ..
func NewEmail(c *api.Config) *Email {
	c.Check()
	uri := "http://" + c.Host + ":" + c.Port
	return &Email{uri: uri}
}

// FindCode ...
func (e *Email) FindCode(email, code string) error {
	if email == "" || code == "" {
		return errors.New("参数错误")
	}

	url := fmt.Sprintf("%s/v1/emailcode/%s/email/%s", e.uri, code, email)
	return api.Get(url, nil)
}
