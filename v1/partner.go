package v1

import (
	"errors"
	"time"

	"github.com/launchain/api"
)

// Partner ...
type Partner struct {
	uri string
}

// PartnerRequest ...
type PartnerRequest struct {
	ID string
}

// PartnerResponse ...
type PartnerResponse struct {
	ID          string    `json:"_id"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Description string    `json:"desc"`
	Address     string    `json:"address"`
	PassPhrase  string    `json:"passphrase"`
	Keystore    string    `json:"keystore"`
}

// NewPartner ...
func NewPartner(c *api.Config) *Partner {
	c.Check()
	uri := c.URI()
	return &Partner{uri: uri}
}

// FindID ...
func (u *Partner) FindID(id string) (*PartnerResponse, error) {
	if id == "" {
		return nil, errors.New("id不能为空")
	}

	url := u.uri + "/v1/partner/" + id
	out := &PartnerResponse{}
	err := api.Get(url, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
