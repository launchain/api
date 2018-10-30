package v1

import (
	"fmt"
	"github.com/launchain/api"
	"net/url"
)

// UserKeystore ...
type UserKeystore struct {
	uri string
}

// UserKeyStoreCreateRequest ...
type UserKeyStoreCreateRequest struct {
	UserID   string
	Address  string
	KeyStore string
	FileName string
	AesKey   string
}

// NewUserKeyStore ...
func NewUserKeyStorer(c *api.Config) *UserKeystore {
	c.Check()
	uri := c.URI()
	return &UserKeystore{uri: uri}
}

// UserKeyStoreCreate ...
func (u *UserKeystore) UserKeyStoreCreate(uk *UserKeyStoreCreateRequest) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	data := make(url.Values)
	data.Add("address", uk.Address)
	data.Add("keystore", uk.KeyStore)
	data.Add("filename", uk.FileName)
	url := fmt.Sprintf("%s/v1/user/%s/keystore/upload", u.uri, uk.UserID)
	return out, api.PostForm(url, data, &out)
}
