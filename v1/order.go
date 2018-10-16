package v1

import (
	"fmt"
	"github.com/launchain/exchange-api"
	"time"
)

// Order ...
type Order struct {
	uri string
}

// NewOrder ...
func NewOrder(config *api.Config) *Order {
	uri := config.URI()
	return &Order{uri: uri}
}

// OrderResponse ...
type OrderResponse struct {
	OrderNum    string    `json:"orderNum"`
	OrderStatus int       `json:"orderStatus"`
	UserId      string    `json:"userid"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Remark      string    `json:"remark"`
}

// GetOneOrder ...
func (o *Order) GetOneOrder(orderId string) (OrderResponse, error) {
	var out OrderResponse
	if orderId == "" {
		return out, nil
	}

	url := fmt.Sprintf("%s/v1/order/detail/%s", o.uri, orderId)
	err := api.Get(url, &out)
	if err != nil {
		return out, err
	}
	return out, nil
}
