package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// MyError ...
type MyError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Get ...
func Get(url string, out interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	resp.Request.Close = true
	return parseResp(resp, out)
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

// Patch ...
func Patch(url string, data url.Values, out interface{}) error {
	req, err := http.NewRequest("PATCH", url, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	if resp != nil {
		defer resp.Body.Close()
		resp.Request.Close = true
	}

	return parseResp(resp, out)
}

// Delete ...
func Delete(url string) error {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	resp.Request.Close = true

	if resp.StatusCode == 204 {
		return nil
	}

	return fmt.Errorf("未知错误[%d]", resp.StatusCode)
}

func parseResp(resp *http.Response, out interface{}) error {
	if resp == nil {
		return errors.New("resp is nil")
	}

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
	case 400, 401, 404, 500:
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		e := &MyError{}
		err = json.Unmarshal(b, e)
		if err != nil {
			return err
		}

		return errors.New(e.Code)
	default:
		return fmt.Errorf("不支持的状态码[%d]", resp.StatusCode)
	}
}
