package testerclient

import (
	"fmt"
	"net/http"
)

type StatusResp struct {
	ID     int  `json:"id"`
	Online bool `json:"online"`
}

func (c *Client) GetStatusByID(id string) (*StatusResp, error) {
	r, err := c.DoRequest(http.MethodGet, "/objects/"+id, nil, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("connection error: %v", err)
	}

	if r.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", r.StatusCode)
	}

	resp := &StatusResp{}
	if err := c.parseBody(r, resp); err != nil {
		return nil, fmt.Errorf("unable to parse body: %v", err)
	}

	return resp, nil
}
