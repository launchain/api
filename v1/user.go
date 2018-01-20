package v1

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/launchain/api"
)

// User ...
type User struct {
	uri string
}

// UserRequest ...
type UserRequest struct {
	Password       string
	Authentication int
}

// UserResponse ...
type UserResponse struct {
	ID             string    `json:"_id"`
	Authentication int       `json:"authentication"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Phone          string    `json:"phone"`
}

// NewUser ...
func NewUser(c *api.Config) *User {
	c.Check()
	uri := "http://" + c.Host + ":" + c.Port
	return &User{uri: uri}
}

// CheckPassword ...
func (u *User) CheckPassword(phone, password string) (*UserResponse, error) {
	data := make(url.Values)
	data["phone"] = []string{phone}
	data["password"] = []string{password}

	url := u.uri + "/v1/ps"
	out := &UserResponse{}
	err := api.PostForm(url, data, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// Create ...
func (u *User) Create(phone, password string) (*UserResponse, error) {
	data := make(url.Values)
	data["phone"] = []string{phone}
	data["password"] = []string{password}

	url := u.uri + "/v1/users"
	out := &UserResponse{}
	err := api.PostForm(url, data, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// FindID ...
func (u *User) FindID(id string) (*UserResponse, error) {
	if id == "" {
		return nil, errors.New("id不能为空")
	}

	url := u.uri + "/v1/users/" + id
	out := &UserResponse{}
	err := api.Get(url, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// UpdateID ...
func (u *User) UpdateID(id string, user *UserRequest) error {
	if id == "" {
		return errors.New("id不能为空")
	}

	data := make(url.Values)
	data["password"] = []string{user.Password}
	data["authentication"] = []string{fmt.Sprintf("%d", user.Authentication)}

	url := u.uri + "/v1/users/" + id
	return api.Patch(url, data, nil)
}
