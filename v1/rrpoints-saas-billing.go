package v1

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/launchain/api"
)

type BillingState = int
type BillingType = int

const (
	BillingStateInit BillingState = iota + 1
	BillingStatePayFinish
	BillingStateSendGoodsFinish
	BillingStateClosed
)

const (
	// BillingTypeWxPay 微信支付
	BillingTypeWxPay BillingType = iota + 1
	// BillingTypePurchase 购买服务包
	BillingTypePurchase
)

// Billing ...
type Billing struct {
	ID          string    `gorm:"column:id" json:"id"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
	State       int       `gorm:"column:state" json:"state"`
	Type        int       `gorm:"column:type" json:"type"`
	UserID      string    `gorm:"column:user_id" json:"user_id"`
	StoreID     string    `gorm:"column:store_id" json:"store_id"`
	Price       string    `gorm:"column:price" json:"price"`
	TokenId     string    `gorm:"column:token_id" json:"token_id"`
	TokenAmount string    `gorm:"column:token_amount" json:"token_amount"`
	TokenSymbol string    `gorm:"column:token_symbol" json:"token_symbol"`
	SrcAddr     string    `gorm:"column:src_addr" json:"src_addr"`
	DstAddr     string    `gorm:"column:dst_addr" json:"dst_addr"`
	Remark      string    `gorm:"column:remark" json:"remark"`
	Phone       string    `gorm:"column:phone" json:"phone"`
}

// RRPointsBilling ...
type RRPointsBilling struct {
	uri string
}

// GetAllBillingsRequest ...
type GetAllBillingsRequest struct {
	TracingBase
	Page     int
	PageSize int
	Order    string
}

// GetAllBillingsResponse ...
type GetAllBillingsResponse struct {
	Billings []Billing `json:"billings"`
}

// GetBillingRequest ...
type GetBillingRequest struct {
	TracingBase
	Id string
}

// GetBillingResponse ...
type GetBillingResponse struct {
	Billing
}

// AddBillingRequest ...
type AddBillingRequest struct {
	TracingBase
	Data Billing
}

// AddBillingResponse ...
type AddBillingResponse struct {
	Billing
}

// UpdateBillingRequest ...
type UpdateBillingRequest struct {
	TracingBase
	Data Billing
}

// UpdateBillingResponse ...
type UpdateBillingResponse struct {
	Billing
}

// DeleteBillingRequest ...
type DeleteBillingRequest struct {
	TracingBase
	Id string
}

// GetStoreBillingRequest ...
type GetStoreBillingRequest struct {
	TracingBase
	StoreID  string
	Page     int
	PageSize int
	Order    string
	DateFrom string
	DateTo   string
}

// GetStoreBillingResponse ...
type GetStoreBillingResponse struct {
	Billings []Billing `json:"billings"`
	Page     int       `json:"page"`
	PageSize int       `json:"page_size"`
	Pages    int       `json:"pages"`
	Total    string    `json:"total"`
}

// GetUserBillingRequest ...
type GetUserBillingRequest struct {
	TracingBase
	UserID   string
	Page     int
	PageSize int
	Order    string
	DateFrom string
	DateTo   string
}

// GetUserBillingResponse ...
type GetUserBillingResponse struct {
	Billings []Billing `json:"billings"`
	Page     int       `json:"page"`
	PageSize int       `json:"page_size"`
	Pages    int       `json:"pages"`
	Total    string    `json:"total"`
}

// NewRRPointsBilling ...
func NewRRPointsBilling(c *api.Config) *RRPointsBilling {
	c.Check()
	uri := c.URI()
	return &RRPointsBilling{uri: uri}
}

// GetAllBillings ...
func (o *RRPointsBilling) GetAllBillings(request *GetAllBillingsRequest) (*GetAllBillingsResponse, error) {
	out := GetAllBillingsResponse{}
	url := fmt.Sprintf("%s/v1/rrpoints-saas/billings/billing?page=%v&pagesize=%v&order=%v",
		o.uri, request.Page, request.PageSize, request.Order)
	err := api.GetAndTrace(request.SpanContext, url, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// GetBilling ...
func (o *RRPointsBilling) GetBilling(request *GetBillingRequest) (*GetBillingResponse, error) {
	out := GetBillingResponse{}
	url := fmt.Sprintf("%s/v1/rrpoints-saas/billings/billing/%v", o.uri, request.Id)
	err := api.GetAndTrace(request.SpanContext, url, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// AddBilling ...
func (o *RRPointsBilling) AddBilling(request *AddBillingRequest) (*AddBillingResponse, error) {
	out := AddBillingResponse{}
	url := fmt.Sprintf("%s/v1/rrpoints-saas/billings/billing", o.uri)
	params, _ := json.Marshal(request.Data)

	err := api.PostJsonAndTrace(request.SpanContext, url, string(params), &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateBilling ...
func (o *RRPointsBilling) UpdateBilling(request *UpdateBillingRequest) (*UpdateBillingResponse, error) {
	out := UpdateBillingResponse{}
	url := fmt.Sprintf("%s/v1/rrpoints-saas/billings/billing/%v", o.uri, request.Data.ID)
	params, _ := json.Marshal(request.Data)

	err := api.PutJsonAndTrace(request.SpanContext, url, string(params), &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteBilling ...
func (o *RRPointsBilling) DeleteBilling(request *DeleteBillingRequest) error {
	url := fmt.Sprintf("%s/v1/rrpoints-saas/billings/billing/%v", o.uri, request.Id)

	err := api.DeleteAndTrace(request.SpanContext, url)
	if err != nil {
		return err
	}
	return nil
}

// GetStoreBilling ...
func (o *RRPointsBilling) GetStoreBilling(request *GetStoreBillingRequest) (*GetStoreBillingResponse, error) {
	out := GetStoreBillingResponse{}
	url := fmt.Sprintf("%s/v1/rrpoints-saas/billings/store/%v?page=%v&pagesize=%v&order=%v&date_from=%v&date_to=%v",
		o.uri, request.StoreID, request.Page, request.PageSize, request.Order, request.DateFrom, request.DateTo)
	err := api.GetAndTrace(request.SpanContext, url, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// GetUserBilling ...
func (o *RRPointsBilling) GetUserBilling(request *GetUserBillingRequest) (*GetUserBillingResponse, error) {
	out := GetUserBillingResponse{}
	url := fmt.Sprintf("%s/v1/rrpoints-saas/billings/user/%v?page=%v&pagesize=%v&order=%v&date_from=%v&date_to=%v",
		o.uri, request.UserID, request.Page, request.PageSize, request.Order, request.DateFrom, request.DateTo)
	err := api.GetAndTrace(request.SpanContext, url, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
