package v1

import (
	"errors"
	"fmt"
	"time"

	"github.com/launchain/api"
)

// Email ...
type Email struct {
	uri string
}

// Code ...
type Code struct {
	ID        string    `json:"_id"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Type      int       `json:"type"`
	Code      string    `json:"code"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewEmail ..
func NewEmail(c *api.Config) *Email {
	c.Check()
	uri := c.URI()
	return &Email{uri: uri}
}

// FindCode ...
func (e *Email) FindCode(email, code string) (*Code, error) {
	if email == "" || code == "" {
		return nil, errors.New("参数错误")
	}

	url := fmt.Sprintf("%s/v1/emailcode/%s/email/%s", e.uri, code, email)
	c := &Code{}
	err := api.Get(url, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
