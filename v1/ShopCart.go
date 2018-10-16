package v1

import (
	"github.com/launchain/exchange-api"
	"time"
)

// ShopCart ...
type ShopCart struct {
	uri string
}

// ShopCartResponse ...
type ShopCartResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Apikey    string    `bson:"apikey" json:"apikey"`
	AssetId   string    `bson:"assetid" json:"assetid"`
	UserId    string    `bson:"userid" json:"userid"`
}

// ShopCartsResponse
type ShopCartsResponse struct {
	Count int                `json:"count"`
	Data  []ShopCartResponse `json:"data"`
}

// NewShopCart ..
func NewShopCart(c *api.Config) *ShopCart {
	c.Check()
	uri := c.URI()
	return &ShopCart{uri: uri}
}

// ShopCartList ...
func (s *ShopCart) ShopCartList(userid string) (*[]ShopCartResponse, error) {

	out := &ShopCartsResponse{}
	err := api.Get(s.uri+"/v1/shopcart/list/"+userid, out)
	if err != nil {
		return nil, err
	}

	return &(out.Data), nil
}
