package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/launchain/api"
	"github.com/opentracing/opentracing-go"
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
	AppName             string
	DeviceName          string
	HeadImg             string
	Nickname            string
}

// UserResponse ...
type UserResponse struct {
	ID                  string    `json:"id"`
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
	Platform            int       `json:"platform"`
	RefreshToken        string    `json:"refresh_token"`
	UnionId             string    `json:"unionid"`
	OpenId              string    `json:"openid"`
	PassStatus          int       `json:"pass_status"`
	AppName             string    `json:"appname"`
	DeviceName          string    `json:"devicename"`
	HeadImg             string    `json:"head_img"`
	Nickname            string    `json:"nickname"`
}

// NewUser ...
func NewUser(c *api.Config) *User {
	c.Check()
	uri := c.URI()
	return &User{uri: uri}
}

// UserFindRequest ...
type UserFindRequest struct {
	TracingBase
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
	TracingBase
	Phone         string
	PassWord      string
	Platform      int
	WalletAddress string
	AppName       string
	DeviceName    string
	Email         string
	WechatInfo
}

// FindUserByAddressRequest ...
type FindUserByAddressRequest struct {
	TracingBase
	WalletAddress []string `json:"wallet_address"`
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

// CheckUserByEmailResponse ...
type CheckUserResponse struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Authentication int    `json:"authentication"`
	IsSetPassword  bool   `json:"is_set_password"`
}

// UserIDs ...
type UserIDs struct {
	UserIDs []string `json:"user_ids"`
}

// PayPassRequest ...
type PayPassRequest struct {
	TracingBase
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

// Deprecated: use TracingFind
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

// Deprecated: use TracingCheckPassword
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

// Deprecated: use TracingCreate
func (u *User) Create(user UserCreateRequest) (*UserResponse, error) {
	data := make(url.Values)
	data["phone"] = []string{user.Phone}
	data["password"] = []string{user.PassWord}
	data["platform"] = []string{fmt.Sprintf("%d", user.Platform)}
	data["app_name"] = []string{user.AppName}
	data["device_name"] = []string{user.DeviceName}

	url := u.uri + "/v1/users"
	out := &UserResponse{}
	err := api.PostForm(url, data, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// Deprecated: use TracingAutoCreate
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
	data["app_name"] = []string{user.AppName}
	data["device_name"] = []string{user.DeviceName}
	data["email"] = []string{user.Email}

	url := u.uri + "/v1/users/phone"
	out := &UserResponse{}
	err := api.PostForm(url, data, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// Deprecated: use TracingFindID
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

// Deprecated: use TracingUpdateID
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
	data.Add("app_name", user.AppName)
	data.Add("device_name", user.DeviceName)
	data.Add("head_img", user.HeadImg)
	data.Add("nickname", user.Nickname)
	url := u.uri + "/v1/users/" + id
	return api.Patch(url, data, nil)
}

// Deprecated: use TracingSensitiveData
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

// Deprecated: use TracingCheckPasswordWithEmail
func (u *User) CheckPasswordWithEmail(email, password string) (*UserResponse, error) {
	data := make(url.Values)
	data["email"] = []string{email}
	data["password"] = []string{password}

	url := u.uri + "/v1/ps/email"
	out := &UserResponse{}
	err := api.PostForm(url, data, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// Deprecated: use TracingCheckUser
func (u *User) CheckUser(email, phone, userID string) (*CheckUserResponse, error) {
	// 三者传其一
	url := u.uri + fmt.Sprintf("/v1/users/check?email=%s&phone=%s&user_id=%s", email, phone, userID)
	out := &CheckUserResponse{}
	err := api.Get(url, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// Deprecated: use TracingUpdateByEmail
func (u *User) UpdateByEmail(email string, user *UserRequest) error {
	if email == "" {
		return errors.New("email不能为空")
	}

	data := make(url.Values)
	data.Add("password", user.Password)
	data.Add("authentication", fmt.Sprintf("%d", user.Authentication))
	data.Add("phone", user.Phone)
	//data.Add("email", user.Email)
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
	data.Add("app_name", user.AppName)
	data.Add("device_name", user.DeviceName)
	data.Add("head_img", user.HeadImg)
	data.Add("nickname", user.Nickname)
	url := u.uri + fmt.Sprintf("/v1/users/email/%s", email)
	return api.Patch(url, data, nil)
}

// Deprecated: use TracingFindAllUsersByAddress
func (u *User) FindAllUsersByAddress(req *FindUserByAddressRequest) (*UserFindResponse, error) {
	out := &UserFindResponse{}
	params, _ := json.Marshal(req)
	err := api.PostJson(u.uri+"/v1/users/address", string(params), out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// Deprecated: use TracingFindAllUsersByUserIDs
func (u *User) FindAllUsersByUserIDs(userIDs UserIDs) (*UserFindResponse, error) {
	out := &UserFindResponse{}
	params, _ := json.Marshal(userIDs)
	err := api.PostJson(u.uri+"/v1/users/multiply", string(params), out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// Deprecated: use TracingCreatePayPassword
func (u *User) CreatePayPassword(req *PayPassRequest) error {
	var out interface{}
	data := make(url.Values)
	data.Add("user_id", req.UserID)
	data.Add("password", req.Password)
	err := api.PostForm(u.uri+"/v1/users/paypassword", data, out)
	if err != nil {
		return err
	}

	return nil
}

// Deprecated: use TracingFindPayPassword
func (u *User) FindPayPassword(req *PayPassRequest) (string, error) {
	out := make(map[string]string)
	url := u.uri + fmt.Sprintf("/v1/users/%s/paypassword", req.UserID)
	err := api.Get(url, &out)
	if err != nil {
		return "", err
	}

	return out["password"], nil
}

// Deprecated: use TracingCheckPayPassword
func (u *User) CheckPayPassword(req *PayPassRequest) (bool, error) {
	out := make(map[string]bool)
	url := u.uri + fmt.Sprintf("/v1/users/%s/paypassword/%s", req.UserID, req.Password)
	err := api.Get(url, &out)
	if err != nil {
		return false, err
	}

	if !out["result"] {
		return false, nil
	}

	return true, nil
}

// TracingFind ...
func (u *User) TracingFind(fr *UserFindRequest) (*UserFindResponse, error) {
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
	err := api.GetAndTrace(fr.SpanContext, u.uri+"/v1/users?"+data.Encode(), out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// TracingCheckPassword ...
func (u *User) TracingCheckPassword(spanCtx opentracing.SpanContext, phone, password string) (*UserResponse, error) {
	data := make(url.Values)
	data["phone"] = []string{phone}
	data["password"] = []string{password}

	url := u.uri + "/v1/ps"
	out := &UserResponse{}
	err := api.PostFormAndTrace(spanCtx, url, data, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// TracingCreate ...
func (u *User) TracingCreate(user UserCreateRequest) (*UserResponse, error) {
	data := make(url.Values)
	data["phone"] = []string{user.Phone}
	data["password"] = []string{user.PassWord}
	data["platform"] = []string{fmt.Sprintf("%d", user.Platform)}
	data["app_name"] = []string{user.AppName}
	data["device_name"] = []string{user.DeviceName}

	url := u.uri + "/v1/users"
	out := &UserResponse{}
	err := api.PostFormAndTrace(user.SpanContext, url, data, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// TracingAutoCreate ...
func (u *User) TracingAutoCreate(user UserCreateRequest) (*UserResponse, error) {
	data := make(url.Values)
	data["phone"] = []string{user.Phone}
	if user.UnionId != "" && user.OpenId != "" && user.RefreshToken != "" {
		data["unionid"] = []string{user.UnionId}
		data["openid"] = []string{user.OpenId}
		data["refresh_token"] = []string{user.RefreshToken}
	}
	data["wallet_address"] = []string{user.WalletAddress}
	data["platform"] = []string{fmt.Sprintf("%d", user.Platform)}
	data["app_name"] = []string{user.AppName}
	data["device_name"] = []string{user.DeviceName}
	data["email"] = []string{user.Email}

	url := u.uri + "/v1/users/phone"
	out := &UserResponse{}
	err := api.PostFormAndTrace(user.SpanContext, url, data, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// TracingFindID ...
func (u *User) TracingFindID(spanCtx opentracing.SpanContext, id string) (*UserResponse, error) {
	if id == "" {
		return nil, errors.New("id不能为空")
	}

	url := u.uri + "/v1/users/" + id
	out := &UserResponse{}
	err := api.GetAndTrace(spanCtx, url, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// TracingUpdateID ...
func (u *User) TracingUpdateID(spanCtx opentracing.SpanContext, id string, user *UserRequest) error {
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
	data.Add("app_name", user.AppName)
	data.Add("device_name", user.DeviceName)
	data.Add("head_img", user.HeadImg)
	data.Add("nickname", user.Nickname)
	url := u.uri + "/v1/users/" + id
	return api.PatchAndTrace(spanCtx, url, data, nil)
}

// TracingSensitiveData ...
func (u *User) TracingSensitiveData(spanCtx opentracing.SpanContext, id string) (*UserResponse, error) {
	if id == "" {
		return nil, errors.New("id不能为空")
	}

	url := fmt.Sprintf("%s/v1/users/%s/sensitivedata", u.uri, id)
	out := &UserResponse{}
	err := api.GetAndTrace(spanCtx, url, out)

	if err != nil {
		return nil, err
	}
	return out, nil
}

// TracingCheckPasswordWithEmail ...
func (u *User) TracingCheckPasswordWithEmail(spanCtx opentracing.SpanContext, email, password string) (
	*UserResponse, error) {
	data := make(url.Values)
	data["email"] = []string{email}
	data["password"] = []string{password}

	url := u.uri + "/v1/ps/email"
	out := &UserResponse{}
	err := api.PostFormAndTrace(spanCtx, url, data, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// TracingCheckUser ...
func (u *User) TracingCheckUser(spanCtx opentracing.SpanContext, email, phone, userID string) (
	*CheckUserResponse, error) {

	// 三者传其一
	url := u.uri + fmt.Sprintf("/v1/users/check?email=%s&phone=%s&user_id=%s", email, phone, userID)
	out := &CheckUserResponse{}
	err := api.GetAndTrace(spanCtx, url, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// TracingUpdateByEmail ...
func (u *User) TracingUpdateByEmail(spanCtx opentracing.SpanContext, email string, user *UserRequest) error {
	if email == "" {
		return errors.New("email不能为空")
	}

	data := make(url.Values)
	data.Add("password", user.Password)
	data.Add("authentication", fmt.Sprintf("%d", user.Authentication))
	data.Add("phone", user.Phone)
	//data.Add("email", user.Email)
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
	data.Add("app_name", user.AppName)
	data.Add("device_name", user.DeviceName)
	data.Add("head_img", user.HeadImg)
	data.Add("nickname", user.Nickname)
	url := u.uri + fmt.Sprintf("/v1/users/email/%s", email)
	return api.PatchAndTrace(spanCtx, url, data, nil)
}

// TracingFindAllUsersByAddress ...
func (u *User) TracingFindAllUsersByAddress(req *FindUserByAddressRequest) (*UserFindResponse, error) {
	out := &UserFindResponse{}
	params, _ := json.Marshal(req)
	err := api.PostJsonAndTrace(req.SpanContext, u.uri+"/v1/users/address", string(params), out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// TracingFindAllUsersByUserIDs ...
func (u *User) TracingFindAllUsersByUserIDs(spanCtx opentracing.SpanContext, userIDs UserIDs) (*UserFindResponse, error) {
	out := &UserFindResponse{}
	params, _ := json.Marshal(userIDs)
	err := api.PostJsonAndTrace(spanCtx, u.uri+"/v1/users/multiply", string(params), out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// TracingCreatePayPassword ...
func (u *User) TracingCreatePayPassword(req *PayPassRequest) error {
	var out interface{}
	data := make(url.Values)
	data.Add("user_id", req.UserID)
	data.Add("password", req.Password)
	err := api.PostFormAndTrace(req.SpanContext, u.uri+"/v1/users/paypassword", data, out)
	if err != nil {
		return err
	}

	return nil
}

// TracingFindPayPassword ...
func (u *User) TracingFindPayPassword(req *PayPassRequest) (string, error) {
	out := make(map[string]string)
	url := u.uri + fmt.Sprintf("/v1/users/%s/paypassword", req.UserID)
	err := api.GetAndTrace(req.SpanContext, url, &out)
	if err != nil {
		return "", err
	}

	return out["password"], nil
}

// TracingCheckPayPassword ...
func (u *User) TracingCheckPayPassword(req *PayPassRequest) (bool, error) {
	out := make(map[string]bool)
	url := u.uri + fmt.Sprintf("/v1/users/%s/paypassword/%s", req.UserID, req.Password)
	err := api.GetAndTrace(req.SpanContext, url, &out)
	if err != nil {
		return false, err
	}

	if !out["result"] {
		return false, nil
	}

	return true, nil
}
