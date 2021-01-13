package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	http "github.com/leapforce-libraries/go_exactonline_new/http"
)

const path string = "budget"

type Service struct {
	http *http.Http
}

func NewService(http *http.Http) *Service {
	return &Service{http}
}

func (service *Service) Get(url string, model interface{}) (string, *errortools.Error) {
	return service.http.Get(url, model)
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
