package exactonline

import (
	http "github.com/leapforce-libraries/go_exactonline_new/http"
)

type Client struct {
	http *http.Http
}

func NewClient(http *http.Http) *Client {
	return &Client{http}
}

func (c *Client) Http() *http.Http {
	return c.http
}
