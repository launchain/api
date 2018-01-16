package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// MyError ...
type MyError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func postForm(url string, data url.Values) (interface{}, error) {
	resp, err := http.PostForm(url, data)
	if err != nil {
		return nil, err
	}

	if resp != nil {
		defer resp.Body.Close()
	}
	resp.Request.Close = true

	return parseResp(resp)
}

func parseResp(resp *http.Response) (interface{}, error) {
	switch resp.StatusCode {
	case 200:
		return ioutil.ReadAll(resp.Body)
	case 201, 202, 204:
		return nil, nil
	case 401, 404, 500:
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var e MyError
		err = json.Unmarshal(b, &e)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(e.Message)
	default:
		return nil, fmt.Errorf("不支持的状态码[%d]", resp.StatusCode)
	}
}
