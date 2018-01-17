package v1

import (
	"net/url"
	"time"

	"github.com/launchain/api"
)

// User ...
type User struct {
	uri string
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
