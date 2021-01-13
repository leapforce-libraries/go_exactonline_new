package exactonline

import (
	"bytes"
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	http "github.com/leapforce-libraries/go_exactonline_new/http"
)

const path string = "subscription"

type Service struct {
	http *http.Http
}

func NewClient(http *http.Http) *Service {
	return &Service{http}
}

func (service *Service) GetSingle(url string, model interface{}) *errortools.Error {
	return service.http.GetSingle(url, model)
}

func (service *Service) Get(url string, model interface{}) (string, *errortools.Error) {
	return service.http.Get(url, model)
}

func (service *Service) Post(url string, buf *bytes.Buffer, model interface{}) *errortools.Error {
	return service.http.Post(url, buf, model)
}

func (service *Service) Put(url string, buf *bytes.Buffer) *errortools.Error {
	return service.http.Put(url, buf)
}

func (service *Service) Delete(url string) *errortools.Error {
	return service.http.Delete(url)
}

func (service *Service) GetCount(endpoint string, createdBefore *time.Time) (int64, *errortools.Error) {
	return service.http.GetCount(fmt.Sprintf("%s/%s", path, endpoint), createdBefore)
}

func (service *Service) BaseURL() string {
	return service.http.BaseURL(path)
}

func (service *Service) DateFilter(field string, comparer string, time *time.Time, includeParameter bool, prefix string) string {
	return service.http.DateFilter(field, comparer, time, includeParameter, prefix)
}
