package testerclient

import (
	"github.com/gojek/heimdall/httpclient"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Options represents client configuration options
type Options struct {
	Host     string
	Port     int
	Secure   bool
	BasePath string
	TimeOut  time.Duration
}

// Client represents host and access token for accessing wallet api
type Client struct {
	opt    Options
	client *httpclient.Client
}

// NewClient returns a new client instance
func NewClient(opt Options) *Client {
	c := &Client{
		opt:    opt,
		client: httpclient.NewClient(httpclient.WithHTTPTimeout(opt.TimeOut)),
	}
	return c
}

// Do actually makes the http request
func (c *Client) Do(method, url string, body io.Reader, h http.Header, params url.Values) (*http.Response, error) {
	req, err := http.NewRequest(method, c.buildURL(url, params), body)
	if err != nil {
		return nil, err
	}
	if h != nil {
		req.Header = h
	}
	c.fillHeader(req)

	resp, err := c.client.Do(req)

	return resp, err
}

// DoRequest perform a request to client
func (c *Client) DoRequest(method, url string, body interface{}, h http.Header, params url.Values) (*http.Response, error) {
	buf, err := c.bodyBuffer(body)
	if err != nil {
		return nil, err
	}
	return c.Do(method, url, buf, h, params)
}
