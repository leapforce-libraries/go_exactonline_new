package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	http "github.com/leapforce-libraries/go_exactonline_new/http"
)

const path string = "budget"

type Client struct {
	http *http.Http
}

func NewClient(http *http.Http) *Client {
	return &Client{http}
}

func (c *Client) Get(url string, model interface{}) (string, *errortools.Error) {
	return c.http.Get(url, model)
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
