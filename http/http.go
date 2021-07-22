package exactonline

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

const (
	APIName            string = "ExactOnline"
	APIURL             string = "https://start.exactonline.nl/api/v1"
	AuthURL            string = "https://start.exactonline.nl/api/oauth2/auth"
	TokenURL           string = "https://start.exactonline.nl/api/oauth2/token"
	TokenHttpMethod    string = http.MethodPost
	RedirectURL        string = "http://localhost:8080/oauth/redirect"
	LastModifiedFormat string = "2006-01-02T15:04:05"
)

// ExactOnline stores ExactOnline configuration
//
type Service struct {
	division                    int32
	oAuth2Service               *oauth2.Service
	xRateLimitMinutelyRemaining int
	xRateLimitMinutelyReset     int64
}

// methods
//
func NewService(division int32, oauth2Service *oauth2.Service) *Service {
	return &Service{
		division:      division,
		oAuth2Service: oauth2Service,
	}
}

func (service *Service) URL(path string) string {
	return fmt.Sprintf("%s/%v/%s", APIURL, service.division, path)
}

func (service *Service) LastModifiedFormat() string {
	return LastModifiedFormat
}

/*
func (service *Service) InitToken(scope string, accessType *string, prompt *string, state *string) *errortools.Error {
	return service.oAuth2Service.InitToken(scope, accessType, prompt, state)
}*/

// Response represents highest level of exactonline api response
//
type Response struct {
	Data Results `json:"d"`
}

// ResponseSingle represents highest level of exactonline api response
//
type ResponseSingle struct {
	Data json.RawMessage `json:"d"`
}

// Results represents second highest level of exactonline api response
//
type Results struct {
	Results json.RawMessage `json:"results"`
	Next    string          `json:"__next"`
	Count   string          `json:"__count"`
}

// wait assures the maximum of 300(?) api calls per minute dictated by exactonline's rate-limit
func (service *Service) wait() error {
	if service.xRateLimitMinutelyRemaining < 1 {
		reset := time.Unix(service.xRateLimitMinutelyReset/1000, 0)
		ms := reset.Sub(time.Now()).Milliseconds()

		if ms > 0 {
			fmt.Println("eo.xRateLimitMinutelyReset:", service.xRateLimitMinutelyReset)
			fmt.Println("reset:", reset)
			fmt.Println("waiting ms:", ms)
			time.Sleep(time.Duration(ms+1000) * time.Millisecond)
		}
	}

	return nil
}

func (service *Service) readRateLimitHeaders(res *http.Response) {
	if res == nil {
		return
	}
	remaining, errRem := strconv.Atoi(res.Header.Get("X-RateLimit-Minutely-Remaining"))
	reset, errRes := strconv.ParseInt(res.Header.Get("X-RateLimit-Minutely-Reset"), 10, 64)
	if errRem == nil && errRes == nil {
		service.xRateLimitMinutelyRemaining = remaining
		service.xRateLimitMinutelyReset = reset
	}
}

func (service *Service) getResponseSingle(url string) (*http.Request, *http.Response, *ResponseSingle, *errortools.Error) {
	exactOnlineError := ExactOnlineError{}
	responseSingle := ResponseSingle{}

	requestConfig := go_http.RequestConfig{
		URL:           url,
		ResponseModel: &responseSingle,
		ErrorModel:    &exactOnlineError,
	}

	request, response, e := service.oAuth2Service.Get(&requestConfig)
	if e != nil {
		e.SetRequest(request)
		e.SetResponse(response)
		if exactOnlineError.Err.Message.Value != "" {
			e.SetMessage(exactOnlineError.Err.Message.Value)
		}
		return request, response, nil, e
	}

	service.readRateLimitHeaders(response)

	return request, response, &responseSingle, nil
}

func (service *Service) getResponse(url string) (*http.Request, *http.Response, *Response, *errortools.Error) {
	exactOnlineError := ExactOnlineError{}
	_response := Response{}

	requestConfig := go_http.RequestConfig{
		URL:           url,
		ResponseModel: &_response,
		ErrorModel:    &exactOnlineError,
	}

	request, response, e := service.oAuth2Service.Get(&requestConfig)
	if e != nil {
		e.SetRequest(request)
		e.SetResponse(response)
		if exactOnlineError.Err.Message.Value != "" {
			e.SetMessage(exactOnlineError.Err.Message.Value)
		}
		return request, response, nil, e
	}

	service.readRateLimitHeaders(response)

	return request, response, &_response, nil
}

func (service *Service) DateFilter(field string, comparer string, time *time.Time, includeParameter bool, prefix string) string {
	filter := ""
	if time != nil {
		if includeParameter {
			filter = prefix + "$filter="
		}

		filter = fmt.Sprintf("%s%s %s DateTime'%s'", filter, field, comparer, time.Format(service.LastModifiedFormat()))

	}

	return filter
}

func (service *Service) GetCount(url string, createdBefore *time.Time) (int64, *errortools.Error) {
	urlStr := fmt.Sprintf("%s?$top=0&$inlinecount=allpages%s", url, service.DateFilter("Created", "lt", createdBefore, true, "&"))

	request, response, _response, e := service.getResponse(urlStr)
	if e != nil {
		return 0, e
	}

	count, err := strconv.ParseInt(_response.Data.Count, 10, 64)
	if err != nil {
		e := errortools.ErrorMessage(err)
		e.SetRequest(request)
		e.SetResponse(response)
		return 0, e
	}

	return count, nil
}

func (service *Service) GetSingle(url string, model interface{}) *errortools.Error {
	err := service.wait()
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	request, response, _response, e := service.getResponseSingle(url)
	if e != nil {
		return e
	}

	err = json.Unmarshal(_response.Data, &model)
	if err != nil {
		e := errortools.ErrorMessage(err)
		e.SetRequest(request)
		e.SetResponse(response)
		return e
	}

	return nil
}

func (service *Service) Get(url string, model interface{}) (string, *errortools.Error) {
	err := service.wait()
	if err != nil {
		return "", errortools.ErrorMessage(err)
	}

	request, response, _response, e := service.getResponse(url)
	if e != nil {
		return "", e
	}

	err = json.Unmarshal(_response.Data.Results, &model)
	if err != nil {
		e := errortools.ErrorMessage(err)
		e.SetRequest(request)
		e.SetResponse(response)
		return "", e
	}

	return _response.Data.Next, nil
}

func (service *Service) Put(requestConfig *go_http.RequestConfig) *errortools.Error {
	exactOnlineError := ExactOnlineError{}

	maxRetries := uint(0) // no retries to prevent errors like "stream error: stream ID 25; PROTOCOL_ERROR exactonline"
	_requestConfig := go_http.RequestConfig{
		URL:        requestConfig.URL,
		BodyModel:  requestConfig.BodyModel,
		ErrorModel: &exactOnlineError,
		MaxRetries: &maxRetries,
	}

	request, response, e := service.oAuth2Service.Put(&_requestConfig)
	if e != nil {
		e.SetRequest(request)
		e.SetResponse(response)
		if exactOnlineError.Err.Message.Value != "" {
			e.SetMessage(exactOnlineError.Err.Message.Value)
		}
		return e
	}

	service.readRateLimitHeaders(response)

	return nil
}

func (service *Service) PostValues(url string, values map[string]string, model interface{}) *errortools.Error {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(values)

	return service.Post(url, buf, model)
}

func (service *Service) PostBytes(url string, b []byte, model interface{}) *errortools.Error {
	return service.Post(url, bytes.NewBuffer(b), model)
}

func (service *Service) Post(url string, bodyModel interface{}, responseModel interface{}) *errortools.Error {
	exactOnlineError := ExactOnlineError{}
	responseSingle := ResponseSingle{}

	requestConfig := go_http.RequestConfig{
		URL:           url,
		BodyModel:     bodyModel,
		ResponseModel: &responseSingle,
		ErrorModel:    &exactOnlineError,
	}

	request, response, e := service.oAuth2Service.Post(&requestConfig)
	if e != nil {
		e.SetRequest(request)
		e.SetResponse(response)
		if exactOnlineError.Err.Message.Value != "" {
			e.SetMessage(exactOnlineError.Err.Message.Value)
		}
		return e
	}

	service.readRateLimitHeaders(response)

	err := json.Unmarshal(responseSingle.Data, responseModel)
	if err != nil {
		if e == nil {
			return errortools.ErrorMessage(err)
		}
		e.SetMessage(err)
		return e
	}

	return nil
}

func (service *Service) Delete(url string) *errortools.Error {
	exactOnlineError := ExactOnlineError{}

	requestConfig := go_http.RequestConfig{
		URL:        url,
		ErrorModel: &exactOnlineError,
	}

	request, response, e := service.oAuth2Service.Delete(&requestConfig)
	if e != nil {
		e.SetRequest(request)
		e.SetResponse(response)
		if exactOnlineError.Err.Message.Value != "" {
			e.SetMessage(exactOnlineError.Err.Message.Value)
		}
		return e
	}

	service.readRateLimitHeaders(response)

	return nil
}
