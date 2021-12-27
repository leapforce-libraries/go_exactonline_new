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
	division      int32
	oAuth2Service *oauth2.Service
	xRateLimit    *struct {
		Remaining int
		Reset     int64
	}
	xRateLimitMinutely *struct {
		Remaining int
		Reset     int64
	}
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
func (service *Service) wait() *errortools.Error {
	reset := int64(0)

	if service.xRateLimit == nil {
		if service.xRateLimit.Remaining < 1 {
			reset = service.xRateLimit.Reset
		}
	}

	if service.xRateLimitMinutely == nil {
		if service.xRateLimitMinutely.Remaining < 1 {
			if service.xRateLimitMinutely.Reset > reset {
				reset = service.xRateLimitMinutely.Reset
			}
		}
	}

	if reset > 0 {
		resetTime := time.Unix(reset/1000, 0)
		ms := time.Until(resetTime).Milliseconds()
		maxWait := int64(10 * 60 * 1000) // 10 minutes

		if ms > 0 {
			if ms > maxWait {
				errortools.SetContext("rate_limit_reset", resetTime.Format(time.RFC3339))
				return errortools.ErrorMessage("Rate limit waiting time exceeds maximum waiting time")
			}

			fmt.Println("xRateLimitReset:", service.xRateLimit.Reset)
			fmt.Println("xRateLimitMinutelyReset:", service.xRateLimitMinutely.Reset)
			fmt.Println("resetTime:", reset)
			fmt.Println("waiting ms:", ms)
			time.Sleep(time.Duration(ms+1000) * time.Millisecond)
		}
	}

	return nil
}

func (service *Service) readRateLimitHeaders(res *http.Response) bool {
	if res == nil {
		return false
	}

	init := false

	remaining, errRem := strconv.Atoi(res.Header.Get("X-RateLimit-Remaining"))
	reset, errRes := strconv.ParseInt(res.Header.Get("X-RateLimit-Reset"), 10, 64)
	if errRem == nil && errRes == nil {
		if service.xRateLimit == nil {
			init = true

			service.xRateLimit = &struct {
				Remaining int
				Reset     int64
			}{0, 0}
		}

		service.xRateLimit.Remaining = remaining
		service.xRateLimit.Reset = reset
	}

	remaining, errRem = strconv.Atoi(res.Header.Get("X-Ratelimit-Minutely-Remaining"))
	reset, errRes = strconv.ParseInt(res.Header.Get("X-Ratelimit-Minutely-Reset"), 10, 64)
	if errRem == nil && errRes == nil {
		if service.xRateLimitMinutely == nil {
			init = true

			service.xRateLimitMinutely = &struct {
				Remaining int
				Reset     int64
			}{0, 0}
		}

		service.xRateLimitMinutely.Remaining = remaining
		service.xRateLimitMinutely.Reset = reset
	}

	return init
}

func (service *Service) httpRequest(requestConfig *go_http.RequestConfig) *errortools.Error {
retry:
	err := service.wait()
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	exactOnlineError := ExactOnlineError{}
	requestConfig.ErrorModel = &exactOnlineError

	request, response, e := service.oAuth2Service.HTTPRequest(requestConfig)
	init := service.readRateLimitHeaders(response)
	if response != nil {
		if response.StatusCode == http.StatusTooManyRequests && init {
			// retry because rate limit headers were not yet known
			goto retry
		}
	}
	if e != nil {
		e.SetRequest(request)
		e.SetResponse(response)
		if exactOnlineError.Err.Message.Value != "" {
			e.SetMessage(exactOnlineError.Err.Message.Value)
		}
	}

	return e
}

func (service *Service) getResponseSingle(url string) (*ResponseSingle, *errortools.Error) {
	responseSingle := ResponseSingle{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		URL:           url,
		ResponseModel: &responseSingle,
	}

	e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &responseSingle, nil
}

func (service *Service) getResponse(url string) (*Response, *errortools.Error) {
	_response := Response{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		URL:           url,
		ResponseModel: &_response,
	}

	e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &_response, nil
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

	_response, e := service.getResponse(urlStr)
	if e != nil {
		return 0, e
	}

	count, err := strconv.ParseInt(_response.Data.Count, 10, 64)
	if err != nil {
		return 0, errortools.ErrorMessage(err)
	}

	return count, nil
}

func (service *Service) GetSingle(url string, model interface{}) *errortools.Error {
	_response, e := service.getResponseSingle(url)
	if e != nil {
		return e
	}

	err := json.Unmarshal(_response.Data, &model)
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	return nil
}

func (service *Service) Get(url string, model interface{}) (string, *errortools.Error) {
	_response, e := service.getResponse(url)
	if e != nil {
		return "", e
	}

	err := json.Unmarshal(_response.Data.Results, &model)
	if err != nil {
		return "", errortools.ErrorMessage(err)
	}

	return _response.Data.Next, nil
}

func (service *Service) Put(requestConfig *go_http.RequestConfig) *errortools.Error {
	maxRetries := uint(0) // no retries to prevent errors like "stream error: stream ID 25; PROTOCOL_ERROR exactonline"
	_requestConfig := go_http.RequestConfig{
		Method:     http.MethodPut,
		URL:        requestConfig.URL,
		BodyModel:  requestConfig.BodyModel,
		MaxRetries: &maxRetries,
	}

	return service.httpRequest(&_requestConfig)
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
	responseSingle := ResponseSingle{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		URL:           url,
		BodyModel:     bodyModel,
		ResponseModel: &responseSingle,
	}

	e := service.httpRequest(&requestConfig)
	if e != nil {
		return e
	}

	err := json.Unmarshal(responseSingle.Data, responseModel)
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	return nil
}

func (service *Service) Delete(url string) *errortools.Error {
	requestConfig := go_http.RequestConfig{
		Method: http.MethodDelete,
		URL:    url,
	}

	return service.httpRequest(&requestConfig)
}
