package v1

import (
	"github.com/launchain/api"
	"net/url"
)

// Token ...
type Announce struct {
	uri string
}

// NewToken ...
func NewAnnounce(c *api.Config) *Announce {
	c.Check()
	uri := c.URI()
	return &Announce{uri: uri}
}

// CheckCode ...
func (t *Announce) CheckCode(code string) (map[string]string, error) {

	url := t.uri + "/v1/weixin-mp/rrpoints-saas/checkcode?code=" + code
	out := make(map[string]string)
	err := api.Get(url, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// 发送交易通知
func (t *Announce) AnnounceTxInfo(val, txid, openid string) (map[string]string, error) {
	data := make(url.Values)
	data["val"] = []string{val}
	data["txid"] = []string{txid}
	data["openid"] = []string{openid}

	murl := t.uri + "/v1/weixin-mp/rrpoints-saas/sendtxinfo"
	out := make(map[string]string)
	err := api.PostForm(murl, data, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
