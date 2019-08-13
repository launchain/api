package v1

import (
	"fmt"
	"github.com/launchain/api"
)

// Comment ...
type Comment struct {
	uri string
}

// NewComment ...
func NewComment(config *api.Config) *Comment {
	uri := config.URI()
	return &Comment{uri: uri}
}

// CommentResponse ...
type CommentResponse struct {
	OrderId []string `json:"order_id"`
}

//GetCommentList ...
func (c *Comment) GetCommentList(userId string) (CommentResponse, error) {
	var out CommentResponse
	if userId == "" {
		return out, nil
	}
	url := fmt.Sprintf("%s/v1/asset-comment/comment/list/%s", c.uri, userId)
	err := api.Get(url, &out)
	if err != nil {
		return out, err
	}
	return out, nil
}
