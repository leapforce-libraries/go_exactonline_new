package exactonline

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	google "github.com/leapforce-libraries/go_google"
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
type Http struct {
	division                    int32
	oAuth2                      *oauth2.OAuth2
	xRateLimitMinutelyRemaining int
	xRateLimitMinutelyReset     int64
}

// methods
//
func NewHttp(division int32, clientID string, clientSecret string, bigQuery *google.BigQuery) (*Http, *errortools.Error) {
	getTokenFunction := func() (*oauth2.Token, *errortools.Error) {
		return google.GetToken(APIName, clientID, bigQuery)
	}

	saveTokenFunction := func(token *oauth2.Token) *errortools.Error {
		return google.SaveToken(APIName, clientID, token, bigQuery)
	}

	config := oauth2.OAuth2Config{
		ClientID:          clientID,
		ClientSecret:      clientSecret,
		RedirectURL:       RedirectURL,
		AuthURL:           AuthURL,
		TokenURL:          TokenURL,
		TokenHTTPMethod:   TokenHttpMethod,
		GetTokenFunction:  &getTokenFunction,
		SaveTokenFunction: &saveTokenFunction,
	}

	return &Http{
		division: division,
		oAuth2:   oauth2.NewOAuth(config),
	}, nil
}

func (h *Http) BaseURL(path string) string {
	return fmt.Sprintf("%s/%v/%s", APIURL, h.division, path)
}

func (h *Http) LastModifiedFormat() string {
	return LastModifiedFormat
}

func (h *Http) InitToken() *errortools.Error {
	return h.oAuth2.InitToken()
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
func (h *Http) wait() error {
	if h.xRateLimitMinutelyRemaining < 1 {
		reset := time.Unix(h.xRateLimitMinutelyReset/1000, 0)
		ms := reset.Sub(time.Now()).Milliseconds()

		if ms > 0 {
			fmt.Println("eo.xRateLimitMinutelyReset:", h.xRateLimitMinutelyReset)
			fmt.Println("reset:", reset)
			fmt.Println("waiting ms:", ms)
			time.Sleep(time.Duration(ms+1000) * time.Millisecond)
		}
	}

	return nil
}

// generic methods
//

func (h *Http) readRateLimitHeaders(res *http.Response) {
	if res == nil {
		return
	}
	remaining, errRem := strconv.Atoi(res.Header.Get("X-RateLimit-Minutely-Remaining"))
	reset, errRes := strconv.ParseInt(res.Header.Get("X-RateLimit-Minutely-Reset"), 10, 64)
	if errRem == nil && errRes == nil {
		h.xRateLimitMinutelyRemaining = remaining
		h.xRateLimitMinutelyReset = reset
	}
}

func (h *Http) getResponseSingle(url string) (*ResponseSingle, *errortools.Error) {
	ee := ExactOnlineError{}
	response := ResponseSingle{}
	_, res, e := h.oAuth2.Get(url, &response, &ee)

	if e != nil {
		if ee.Err.Message.Value != "" {
			e.SetMessage(ee.Err.Message.Value)
		}
		return nil, e
	}

	h.readRateLimitHeaders(res)

	return &response, nil
}

func (h *Http) getResponse(url string) (*Response, *errortools.Error) {
	ee := ExactOnlineError{}
	response := Response{}
	_, res, e := h.oAuth2.Get(url, &response, &ee)
	if e != nil {
		if ee.Err.Message.Value != "" {
			e.SetMessage(ee.Err.Message.Value)
		}
		return nil, e
	}

	h.readRateLimitHeaders(res)

	return &response, nil
}

func (h *Http) DateFilter(field string, comparer string, time *time.Time, includeParameter bool, prefix string) string {
	filter := ""
	if time != nil {
		if includeParameter {
			filter = prefix + "$filter="
		}

		filter = fmt.Sprintf("%s%s %s DateTime'%s'", filter, field, comparer, time.Format(h.LastModifiedFormat()))

	}

	return filter
}

func (h *Http) GetCount(path string, createdBefore *time.Time) (int64, *errortools.Error) {
	urlStr := fmt.Sprintf("%s?$top=0&$inlinecount=allpages%s", h.BaseURL(path), h.DateFilter("Created", "lt", createdBefore, true, "&"))

	response, e := h.getResponse(urlStr)
	if e != nil {
		return 0, e
	}

	count, err := strconv.ParseInt(response.Data.Count, 10, 64)
	if err != nil {
		return 0, errortools.ErrorMessage(err)
	}

	return count, nil
}

func (h *Http) GetSingle(url string, model interface{}) *errortools.Error {
	err := h.wait()
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	response, e := h.getResponseSingle(url)
	if e != nil {
		return e
	}

	err = json.Unmarshal(response.Data, &model)
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	return nil
}

func (h *Http) Get(url string, model interface{}) (string, *errortools.Error) {
	err := h.wait()
	if err != nil {
		return "", errortools.ErrorMessage(err)
	}

	response, e := h.getResponse(url)
	if e != nil {
		return "", e
	}

	err = json.Unmarshal(response.Data.Results, &model)
	if err != nil {
		return "", errortools.ErrorMessage(err)
	}

	return response.Data.Next, nil
}

func (h *Http) PutValues(url string, values map[string]string) *errortools.Error {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(values)

	return h.Put(url, buf)
}

func (h *Http) PutBytes(url string, b []byte) *errortools.Error {
	return h.Put(url, bytes.NewBuffer(b))
}

func (h *Http) Put(url string, buf *bytes.Buffer) *errortools.Error {
	ee := ExactOnlineError{}
	_, res, e := h.oAuth2.Put(url, buf, nil, &ee)
	if e != nil {
		if ee.Err.Message.Value != "" {
			e.SetMessage(ee.Err.Message.Value)
		}
		return e
	}

	h.readRateLimitHeaders(res)

	return nil
}

func (h *Http) PostValues(url string, values map[string]string, model interface{}) *errortools.Error {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(values)

	return h.Post(url, buf, model)
}

func (h *Http) PostBytes(url string, b []byte, model interface{}) *errortools.Error {
	return h.Post(url, bytes.NewBuffer(b), model)
}

func (h *Http) Post(url string, buf *bytes.Buffer, model interface{}) *errortools.Error {
	ee := ExactOnlineError{}
	response := ResponseSingle{}
	_, res, e := h.oAuth2.Post(url, buf, &response, &ee)
	if e != nil {
		if ee.Err.Message.Value != "" {
			e.SetMessage(ee.Err.Message.Value)
		}
		return e
	}

	h.readRateLimitHeaders(res)

	defer res.Body.Close()

	err := json.Unmarshal(response.Data, &model)
	if err != nil {
		e.SetMessage(err)
		return e
	}

	return nil
}

func (h *Http) Delete(url string) *errortools.Error {
	ee := ExactOnlineError{}
	_, res, e := h.oAuth2.Delete(url, nil, nil, &ee)
	if e != nil {
		if ee.Err.Message.Value != "" {
			e.SetMessage(ee.Err.Message.Value)
		}
		return e
	}

	h.readRateLimitHeaders(res)

	return nil
}
