package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	http "github.com/leapforce-libraries/go_exactonline_new/http"
	go_http "github.com/leapforce-libraries/go_http"
)

const servicePath string = "crm"

type Service struct {
	httpService *http.Service
}

func NewService(httpService *http.Service) *Service {
	return &Service{httpService}
}

func (service *Service) GetSingle(url string, responseModel interface{}) *errortools.Error {
	return service.httpService.GetSingle(url, responseModel)
}

func (service *Service) Get(url string, responseModel interface{}) (string, *errortools.Error) {
	return service.httpService.Get(url, responseModel)
}

func (service *Service) Post(url string, bodyModel interface{}, responseModel interface{}) *errortools.Error {
	return service.httpService.Post(url, bodyModel, responseModel)
}

func (service *Service) Put(requestConfig *go_http.RequestConfig) *errortools.Error {
	return service.httpService.Put(requestConfig)
}

func (service *Service) Delete(url string) *errortools.Error {
	return service.httpService.Delete(url)
}

func (service *Service) GetCount(endpoint string, createdBefore *time.Time) (int64, *errortools.Error) {
	return service.httpService.GetCount(service.url(endpoint), createdBefore)
}

func (service *Service) url(path string) string {
	return service.httpService.Url(fmt.Sprintf("%s/%s", servicePath, path))
}

func (service *Service) DateFilter(field string, comparer string, time *time.Time, includeParameter bool, prefix string) string {
	return service.httpService.DateFilter(field, comparer, time, includeParameter, prefix)
}
