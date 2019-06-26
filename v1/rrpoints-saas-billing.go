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
	ID        string    `gorm:"column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	State     int       `gorm:"column:state" json:"state"`
	Type      int       `gorm:"column:type" json:"type"`
	UserID    string    `gorm:"column:user_id" json:"user_id"`
	StoreID   string    `gorm:"column:store_id" json:"store_id"`
	Price     string    `gorm:"column:price" json:"price"`
	SrcAddr   string    `gorm:"column:src_addr" json:"src_addr"`
	DstAddr   string    `gorm:"column:dst_addr" json:"dst_addr"`
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
	Code    string    `json:"code"`
	Message string    `json:"message"`
	Data    []Billing `json:"data"`
}

// GetAllBillingRequest ...
type GetAllBillingRequest struct {
	TracingBase
	Id string
}

// GetAllBillingsResponse ...
type GetAllBillingResponse struct {
	Code    string  `json:"code"`
	Message string  `json:"message"`
	Data    Billing `json:"data"`
}

// AddBillingRequest ...
type AddBillingRequest struct {
	TracingBase
	Data Billing
}

// AddBillingsResponse ...
type AddBillingResponse struct {
	Code    string  `json:"code"`
	Message string  `json:"message"`
	Data    Billing `json:"data"`
}

// UpdateBillingRequest ...
type UpdateBillingRequest struct {
	TracingBase
	Data Billing
}

// UpdateBillingsResponse ...
type UpdateBillingResponse struct {
	Code    string  `json:"code"`
	Message string  `json:"message"`
	Data    Billing `json:"data"`
}

// DeleteBillingRequest ...
type DeleteBillingRequest struct {
	TracingBase
	Id string
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
	url := fmt.Sprintf("%s/v1/rrpoints-saas/billings?page=%v&pagesize=%v&order=%v",
		o.uri, request.Page, request.PageSize, request.Order)
	err := api.GetAndTrace(request.SpanContext, url, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// GetAllBilling ...
func (o *RRPointsBilling) GetAllBilling(request *GetAllBillingRequest) (*GetAllBillingResponse, error) {
	out := GetAllBillingResponse{}
	url := fmt.Sprintf("%s/v1/rrpoints-saas/billings/%v", o.uri, request.Id)
	err := api.GetAndTrace(request.SpanContext, url, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// AddBilling ...
func (o *RRPointsBilling) AddBilling(request *AddBillingRequest) (*AddBillingResponse, error) {
	out := AddBillingResponse{}
	url := fmt.Sprintf("%s/v1/rrpoints-saas/billings", o.uri)
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
	url := fmt.Sprintf("%s/v1/rrpoints-saas/billings/%v", o.uri, request.Data.ID)
	params, _ := json.Marshal(request.Data)

	err := api.PutJsonAndTrace(request.SpanContext, url, string(params), &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// DeleteBilling ...
func (o *RRPointsBilling) DeleteBilling(request *DeleteBillingRequest) error {
	url := fmt.Sprintf("%s/v1/rrpoints-saas/billings/%v", o.uri, request.Id)

	err := api.DeleteAndTrace(request.SpanContext, url)
	if err != nil {
		return err
	}
	return nil
}
