package testerclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// buildURL returns the full url
func (c *Client) buildURL(path string, params url.Values) string {
	u := url.URL{
		Scheme: "http",
		Host:   c.opt.Host,
		Path:   c.opt.BasePath + path,
	}
	if c.opt.Secure {
		u.Scheme += "s"
	}
	if c.opt.Port != 0 {
		u.Host += fmt.Sprintf(":%d", c.opt.Port)
	}
	if params != nil {
		u.RawQuery = params.Encode()
	}
	return u.String()
}

func (c *Client) bodyBuffer(body interface{}) (io.Reader, error) {
	jsonStr, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer([]byte(jsonStr)), nil
}

func (c *Client) fillHeader(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
}

func (c *Client) captureBody(r io.Reader) (string, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func (c *Client) parseBody(r *http.Response, body interface{}) error {
	return json.NewDecoder(r.Body).Decode(body)
}
