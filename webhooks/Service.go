package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	http "github.com/leapforce-libraries/go_exactonline_new/http"
)

const servicePath string = "webhooks"

type Service struct {
	httpService *http.Service
}

func NewService(httpService *http.Service) *Service {
	return &Service{httpService}
}

func (service *Service) Get(url string, responseModel interface{}) (string, *errortools.Error) {
	return service.httpService.Get(url, responseModel)
}

func (service *Service) url(path string) string {
	return service.httpService.URL(fmt.Sprintf("%s/%s", servicePath, path))
}
