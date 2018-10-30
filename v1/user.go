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
	Password            string
	Email               string
	Authentication      int
	IDCard              string
	RealName            string
	Portrait            string
	Residence           string
	Degree              string
	Career              string
	Phone               string
	CarInfo             string
	DriverLicense       string
	Alipay              string
	RoleName            string
	RoleImg             string
	WalletAddress       string
	CarInfoStatus       int
	RoleStatus          int
	DriverLicenseStatus int
	RefreshToken        string
	UnionId             string
	OpenId              string
}

// UserResponse ...
type UserResponse struct {
	ID                  string    `json:"_id"`
	Authentication      int       `json:"authentication"`
	Email               string    `json:"email"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	Phone               string    `json:"phone"`
	RealName            string    `json:"realname"`
	IDCard              string    `json:"idcard"`
	Residence           string    `json:"residence"`
	Degree              string    `json:"degree"`
	Career              string    `json:"career"`
	CarInfo             string    `json:"carinfo"`
	DriverLicense       string    `json:"driverlicense"`
	Alipay              string    `json:"alipay"`
	RoleName            string    `json:"rolename"`
	RoleImg             string    `json:"roleimg"`
	WalletAddress       string    `json:"wallet_address"`
	CarInfoStatus       int       `json:"carinfo_status"`
	RoleStatus          int       `json:"role_status"`
	DriverLicenseStatus int       `json:"driverlicense_status"`
	RefreshToken        string    `json:"refresh_token"`
	UnionId             string    `json:"unionid"`
	OpenId              string    `json:"openid"`
}

// NewUser ...
func NewUser(c *api.Config) *User {
	c.Check()
	uri := c.URI()
	return &User{uri: uri}
}

// UserFindRequest ...
type UserFindRequest struct {
	Page          int
	Limit         int
	Auth          int
	Email         string
	Phone         string
	WalletAddress string
	OpenId        string
}

//UserCreatRequest ...
type UserCreateRequest struct {
	Phone         string
	PassWord      string
	Platform      int
	WalletAddress string
	WechatInfo
}

//WechatInfo ...
type WechatInfo struct {
	UnionId      string
	OpenId       string
	RefreshToken string
}

// UserFindResponse ...
type UserFindResponse struct {
	Users []*UserResponse `json:"users"`
	Count int             `json:"count"`
}

// Find ...
func (u *User) Find(fr *UserFindRequest) (*UserFindResponse, error) {
	//	url := u.uri + "/v1/users?"

	data := make(url.Values)
	if fr.Email != "" {
		data.Add("email", fr.Email)
	}
	if fr.Auth != 0 {
		data.Add("auth", fmt.Sprintf("%d", fr.Auth))
	}
	if fr.Page != 0 {
		data.Add("page", fmt.Sprintf("%d", fr.Page))
	}
	if fr.Limit != 0 {
		data.Add("limit", fmt.Sprintf("%d", fr.Limit))
	}
	if fr.Phone != "" {
		data.Add("phone", fr.Phone)
	}
	if fr.WalletAddress != "" {
		data.Add("wallet_address", fr.WalletAddress)
	}
	if fr.OpenId != "" {
		data.Add("openid", fr.OpenId)
	}
	out := &UserFindResponse{}
	err := api.Get(u.uri+"/v1/users?"+data.Encode(), out)
	if err != nil {
		return nil, err
	}

	return out, nil
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
func (u *User) Create(user UserCreateRequest) (*UserResponse, error) {
	data := make(url.Values)
	data["phone"] = []string{user.Phone}
	data["password"] = []string{user.PassWord}
	data["platform"] = []string{fmt.Sprintf("%d", user.Platform)}

	url := u.uri + "/v1/users"
	out := &UserResponse{}
	err := api.PostForm(url, data, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// AutoCreate ...
func (u *User) AutoCreate(user UserCreateRequest) (*UserResponse, error) {
	data := make(url.Values)
	data["phone"] = []string{user.Phone}
	if user.UnionId != "" && user.OpenId != "" && user.RefreshToken != "" {
		data["unionid"] = []string{user.UnionId}
		data["openid"] = []string{user.OpenId}
		data["refresh_token"] = []string{user.RefreshToken}
	}
	data["wallet_address"] = []string{user.WalletAddress}
	data["platform"] = []string{fmt.Sprintf("%d", user.Platform)}
	url := u.uri + "/v1/users/phone"
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
	data.Add("password", user.Password)
	data.Add("authentication", fmt.Sprintf("%d", user.Authentication))
	data.Add("phone", user.Phone)
	data.Add("email", user.Email)
	data.Add("idcard", user.IDCard)
	data.Add("realname", user.RealName)
	data.Add("portrait", user.Portrait)
	data.Add("residence", user.Residence)
	data.Add("degree", user.Degree)
	data.Add("career", user.Career)
	data.Add("wallet_address", user.WalletAddress)
	data.Add("carinfo", user.CarInfo)
	data.Add("alipay", user.Alipay)
	data.Add("driverlicense", user.DriverLicense)
	data.Add("wallet_address", user.WalletAddress)
	data.Add("rolename", user.RoleName)
	data.Add("roleimg", user.RoleImg)
	data.Add("phone", user.Phone)
	data.Add("driverlicense_status", fmt.Sprintf("%d", user.DriverLicenseStatus))
	data.Add("carinfo_status", fmt.Sprintf("%d", user.CarInfoStatus))
	data.Add("role_status", fmt.Sprintf("%d", user.RoleStatus))
	data.Add("unionid", user.UnionId)
	data.Add("openid", user.OpenId)
	data.Add("refresh_token", user.RefreshToken)
	url := u.uri + "/v1/users/" + id
	return api.Patch(url, data, nil)
}

// SensitiveData ...
func (u *User) SensitiveData(id string) (*UserResponse, error) {
	if id == "" {
		return nil, errors.New("id不能为空")
	}

	url := fmt.Sprintf("%s/v1/users/%s/sensitivedata", u.uri, id)
	out := &UserResponse{}
	err := api.Get(url, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
