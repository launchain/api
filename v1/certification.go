package v1

import (
	"fmt"
	"github.com/launchain/api"
	"net/url"
)

// Certification ...
type Certification struct {
	uri string
}

// NewCertification ...
func NewCertification(c *api.Config) *Certification {
	c.Check()
	uri := c.URI()
	return &Certification{uri: uri}
}

// CreateChain ...
func (c *CertificationForChain) CreateCertification(userId string) (map[string]interface{}, error) {
	data := make(url.Values)
	Url := fmt.Sprintf("%s/v1/identity/certification/%s", c.uri, userId)
	out := make(map[string]interface{})
	return out, api.PostForm(Url, data, &out)
}
