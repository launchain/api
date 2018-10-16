package v1

import (
	"fmt"
	"github.com/launchain/api"
)

// Asset ...
type Asset struct {
	uri string
}

// NewComment ...
func NewAsset(config *api.Config) *Asset {
	uri := config.URI()
	return &Asset{uri: uri}
}

// CommentResponse ...
type AssetResponse struct {
	Owner string `json:"owner"`
}

//GetCommentList ...
func (a *Asset) GetAssetOwner(assetid, apikey string) (AssetResponse, error) {
	var out AssetResponse
	if assetid == "" || apikey == "" {
		return out, nil
	}
	url := fmt.Sprintf("%s/v1/asset/owner/%s/%s", a.uri, assetid, apikey)
	err := api.Get(url, &out)
	if err != nil {
		return out, err
	}
	return out, nil
}
