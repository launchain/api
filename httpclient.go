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

// PostForm ...
func PostForm(url string, data url.Values, out interface{}) error {
	resp, err := http.PostForm(url, data)
	if err != nil {
		return err
	}

	if resp != nil {
		defer resp.Body.Close()
	}
	resp.Request.Close = true

	return parseResp(resp, out)
}

func parseResp(resp *http.Response, out interface{}) error {
	switch resp.StatusCode {
	case 200:
		if out == nil {
			return nil
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(b, &out)
		if err != nil {
			return err
		}

		return nil
	case 201, 202, 204:
		return nil
	case 401, 404, 500:
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		e := &MyError{}
		err = json.Unmarshal(b, e)
		if err != nil {
			return err
		}

		return errors.New(e.Message)
	default:
		return fmt.Errorf("不支持的状态码[%d]", resp.StatusCode)
	}
}
