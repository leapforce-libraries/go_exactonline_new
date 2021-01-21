package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	http "github.com/leapforce-libraries/go_exactonline_new/http"
)

const path string = "subscription"

type Service struct {
	http *http.Http
}

func NewService(http *http.Http) *Service {
	return &Service{http}
}

func (service *Service) GetSingle(url string, responseModel interface{}) *errortools.Error {
	return service.http.GetSingle(url, responseModel)
}

func (service *Service) Get(url string, responseModel interface{}) (string, *errortools.Error) {
	return service.http.Get(url, responseModel)
}

func (service *Service) Post(url string, bodyModel interface{}, responseModel interface{}) *errortools.Error {
	return service.http.Post(url, bodyModel, responseModel)
}

func (service *Service) Put(url string, bodyModel interface{}) *errortools.Error {
	return service.http.Put(url, bodyModel)
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
