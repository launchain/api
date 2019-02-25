package v1

import (
	"fmt"
	"net/url"

	"github.com/launchain/api"
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
	Phrase   string
}

// NewUserKeyStorer ...
func NewUserKeyStorer(c *api.Config) *UserKeystore {
	c.Check()
	uri := c.URI()
	return &UserKeystore{uri: uri}
}

//UserDefalutAddrResquest ...
type UserDefalutAddrResquest struct {
	Code    string
	Message string
	Data    UserDefalutAddr
}

// UserDefalutKeystoreAllResponse ...
type UserDefalutKeystoreAllResponse struct {
	Base
	Data UserDefalutKeystoreAll `json:"data"`
}

// UserDefalutKeystoreAll ...
type UserDefalutKeystoreAll struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Address  string `json:"address"`
	KeyStore string `json:"keystore"`
	Filename string `json:"filename"`
}

//UserDefalutAddr ...
type UserDefalutAddr struct {
	UserID  string
	Address string
}

// UserKeyStoreCreate ...
func (u *UserKeystore) UserKeyStoreCreate(uk *UserKeyStoreCreateRequest) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	data := make(url.Values)
	data.Add("address", uk.Address)
	data.Add("keystore", uk.KeyStore)
	data.Add("filename", uk.FileName)
	data.Add("phrase", uk.Phrase)
	url := fmt.Sprintf("%s/v1/user/%s/keystore/upload", u.uri, uk.UserID)
	return out, api.PostForm(url, data, &out)
}

// GetUserDefaultKeystore ...
func (u *UserKeystore) GetUserDefaultKeystore(userId string) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	url := fmt.Sprintf("%s/v1/user/%s/keystore/default", u.uri, userId)
	return out, api.Get(url, &out)
}

// GetUserDefaultAddress ...
func (u *UserKeystore) GetUserDefaultAddress(userId string) (*UserDefalutAddrResquest, error) {
	out := &UserDefalutAddrResquest{}
	url := fmt.Sprintf("%s/v1/user/%s/address/default", u.uri, userId)
	err := api.Get(url, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetUserDefaultKeystoreAll ...
func (u *UserKeystore) GetUserDefaultKeystoreAll(userId string) (*UserDefalutKeystoreAllResponse, error) {
	out := &UserDefalutKeystoreAllResponse{}
	url := fmt.Sprintf("%s/v1/user/%s/keystore/default-all", u.uri, userId)
	return out, api.Get(url, &out)
}

// GetUserDefaultKeystoreAll ...
func (u *UserKeystore) GetKeystoreByAddress(address string) (*UserDefalutKeystoreAllResponse, error) {
	out := &UserDefalutKeystoreAllResponse{}
	url := fmt.Sprintf("%s/v1/user-keystore/address/%s", u.uri, address)
	err := api.Get(url, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetUserinfoByAddress ...
func (u *UserKeystore) GetUserinfoByAddress(address string) (*UserDefalutKeystoreAllResponse, error) {
	out := &UserDefalutKeystoreAllResponse{}
	url := fmt.Sprintf("%s/v1/user-keystore/userinfo/%s", u.uri, address)
	err := api.Get(url, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
