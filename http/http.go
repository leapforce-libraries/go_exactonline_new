package exactonline

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
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
	oAuth2                      *oauth2.OAuth2
	xRateLimitMinutelyRemaining int
	xRateLimitMinutelyReset     int64
}

// methods
//
func NewService(division int32, oauth2 *oauth2.OAuth2) *Service {
	return &Service{
		division: division,
		oAuth2:   oauth2,
	}
}

func (service *Service) URL(path string) string {
	return fmt.Sprintf("%s/%v/%s", APIURL, service.division, path)
}

func (service *Service) LastModifiedFormat() string {
	return LastModifiedFormat
}

func (service *Service) InitToken() *errortools.Error {
	return service.oAuth2.InitToken()
}

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

func (service *Service) getResponseSingle(url string) (*ResponseSingle, *errortools.Error) {
	exactOnlineError := ExactOnlineError{}
	response := ResponseSingle{}

	requestConfig := oauth2.RequestConfig{
		URL:           url,
		ResponseModel: &response,
		ErrorModel:    &exactOnlineError,
	}

	_, res, e := service.oAuth2.Get(&requestConfig)
	if e != nil {
		if exactOnlineError.Err.Message.Value != "" {
			e.SetMessage(exactOnlineError.Err.Message.Value)
		}
		return nil, e
	}

	service.readRateLimitHeaders(res)

	return &response, nil
}

func (service *Service) getResponse(url string) (*Response, *errortools.Error) {
	exactOnlineError := ExactOnlineError{}
	response := Response{}

	requestConfig := oauth2.RequestConfig{
		URL:           url,
		ResponseModel: &response,
		ErrorModel:    &exactOnlineError,
	}

	_, res, e := service.oAuth2.Get(&requestConfig)
	if e != nil {
		if exactOnlineError.Err.Message.Value != "" {
			e.SetMessage(exactOnlineError.Err.Message.Value)
		}
		return nil, e
	}

	service.readRateLimitHeaders(res)

	return &response, nil
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

	response, e := service.getResponse(urlStr)
	if e != nil {
		return 0, e
	}

	count, err := strconv.ParseInt(response.Data.Count, 10, 64)
	if err != nil {
		return 0, errortools.ErrorMessage(err)
	}

	return count, nil
}

func (service *Service) GetSingle(url string, model interface{}) *errortools.Error {
	err := service.wait()
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	response, e := service.getResponseSingle(url)
	if e != nil {
		return e
	}

	err = json.Unmarshal(response.Data, &model)
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	return nil
}

func (service *Service) Get(url string, model interface{}) (string, *errortools.Error) {
	err := service.wait()
	if err != nil {
		return "", errortools.ErrorMessage(err)
	}

	response, e := service.getResponse(url)
	if e != nil {
		return "", e
	}

	err = json.Unmarshal(response.Data.Results, &model)
	if err != nil {
		return "", errortools.ErrorMessage(err)
	}

	return response.Data.Next, nil
}

func (service *Service) Put(url string, bodyModel interface{}) *errortools.Error {
	exactOnlineError := ExactOnlineError{}

	requestConfig := oauth2.RequestConfig{
		URL:        url,
		BodyModel:  bodyModel,
		ErrorModel: &exactOnlineError,
	}

	_, res, e := service.oAuth2.Put(&requestConfig)
	if e != nil {
		if exactOnlineError.Err.Message.Value != "" {
			e.SetMessage(exactOnlineError.Err.Message.Value)
		}
		return e
	}

	service.readRateLimitHeaders(res)

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
	response := ResponseSingle{}

	requestConfig := oauth2.RequestConfig{
		URL:        url,
		BodyModel:  bodyModel,
		ErrorModel: &exactOnlineError,
	}

	_, res, e := service.oAuth2.Post(&requestConfig)
	if e != nil {
		if exactOnlineError.Err.Message.Value != "" {
			e.SetMessage(exactOnlineError.Err.Message.Value)
		}
		return e
	}

	service.readRateLimitHeaders(res)

	defer res.Body.Close()

	err := json.Unmarshal(response.Data, responseModel)
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

	requestConfig := oauth2.RequestConfig{
		URL:        url,
		ErrorModel: &exactOnlineError,
	}

	_, res, e := service.oAuth2.Delete(&requestConfig)
	if e != nil {
		if exactOnlineError.Err.Message.Value != "" {
			e.SetMessage(exactOnlineError.Err.Message.Value)
		}
		return e
	}

	service.readRateLimitHeaders(res)

	return nil
}
