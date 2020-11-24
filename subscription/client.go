package exactonline

import (
	"bytes"
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	http "github.com/leapforce-libraries/go_exactonline_new/http"
)

const path string = "subscription"

type Client struct {
	http *http.Http
}

func NewClient(http *http.Http) *Client {
	return &Client{http}
}

func (c *Client) GetSingle(url string, model interface{}) *errortools.Error {
	return c.http.GetSingle(url, model)
}

func (c *Client) Get(url string, model interface{}) (string, *errortools.Error) {
	return c.http.Get(url, model)
}

func (c *Client) Post(url string, buf *bytes.Buffer, model interface{}) *errortools.Error {
	return c.http.Post(url, buf, model)
}

func (c *Client) Put(url string, buf *bytes.Buffer) *errortools.Error {
	return c.http.Put(url, buf)
}

func (c *Client) Delete(url string) *errortools.Error {
	return c.http.Delete(url)
}

func (c *Client) GetCount(endpoint string, createdBefore *time.Time) (int64, *errortools.Error) {
	return c.http.GetCount(fmt.Sprintf("%s/%s", path, endpoint), createdBefore)
}

func (c *Client) BaseURL() string {
	return c.http.BaseURL(path)
}

func (c *Client) DateFilter(field string, comparer string, time *time.Time, includeParameter bool, prefix string) string {
	return c.http.DateFilter(field, comparer, time, includeParameter, prefix)
}
