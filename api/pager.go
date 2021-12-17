package api

import (
	"net/url"
	"strconv"
)

// pager represents api pagination object
type pager struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	TotalPage   int `json:"total_page"`
	Total       int `json:"total"`
}

func (p *pager) offset() int {
	return (p.CurrentPage - 1) * p.PerPage
}

func (p *pager) limit() int {
	return p.PerPage
}

func newPager(url *url.URL, maxLimit int) *pager {
	if maxLimit < 1 {
		maxLimit = 1
	}

	pgr := pager{
		CurrentPage: 1,
		PerPage:     maxLimit,
	}
	if s := url.Query().Get("page"); s != "" {
		if n, err := strconv.Atoi(s); err == nil && n > 0 {
			pgr.CurrentPage = n
		}
	}
	if s := url.Query().Get("per_page"); s != "" {
		if n, err := strconv.Atoi(s); err == nil && n > 0 && n <= maxLimit {
			pgr.PerPage = n
		}
	}
	return &pgr
}
