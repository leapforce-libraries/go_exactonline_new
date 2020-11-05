package exactonline

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	bigquerytools "github.com/leapforce-libraries/go_bigquerytools"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
	types "github.com/leapforce-libraries/go_types"
)

const (
	apiName            string = "ExactOnline"
	apiURL             string = "https://start.exactonline.nl/api/v1"
	authURL            string = "https://start.exactonline.nl/api/oauth2/auth"
	tokenURL           string = "https://start.exactonline.nl/api/oauth2/token"
	tokenHttpMethod    string = http.MethodPost
	redirectURL        string = "http://localhost:8080/oauth/redirect"
	lastModifiedFormat string = "2006-01-02T15:04:05"
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
func NewHttp(division int32, clientID string, clientSecret string, bigQuery *bigquerytools.BigQuery, isLive bool) (*Http, error) {
	h := Http{}
	h.division = division

	config := oauth2.OAuth2Config{
		ApiName:         apiName,
		ClientID:        clientID,
		ClientSecret:    clientSecret,
		RedirectURL:     redirectURL,
		AuthURL:         authURL,
		TokenURL:        tokenURL,
		TokenHTTPMethod: tokenHttpMethod,
	}
	h.oAuth2 = oauth2.NewOAuth(config, bigQuery, isLive)
	return &h, nil
}

func (h *Http) BaseURL() string {
	return fmt.Sprintf("%s/%v", apiURL, h.division)
}

func (h *Http) LastModifiedFormat() string {
	return lastModifiedFormat
}

func (h *Http) InitToken() error {
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
	remaining, errRem := strconv.Atoi(res.Header.Get("X-RateLimit-Minutely-Remaining"))
	reset, errRes := strconv.ParseInt(res.Header.Get("X-RateLimit-Minutely-Reset"), 10, 64)
	if errRem == nil && errRes == nil {
		h.xRateLimitMinutelyRemaining = remaining
		h.xRateLimitMinutelyReset = reset
	}
}

func (h *Http) printError(res *http.Response) error {
	fmt.Println("Status", res.Status)

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		//fmt.Println("errUnmarshal1")
		return err
	}

	ee := ExactOnlineError{}

	err = json.Unmarshal(b, &ee)
	if err != nil {
		//fmt.Println("errUnmarshal1")
		return err
	}

	//fmt.Println(ee.Err.Message.Value)
	message := fmt.Sprintf("Server returned statuscode %v, error:%s", res.StatusCode, ee.Err.Message.Value)
	return &types.ErrorString{message}
}

func (h *Http) GetResponse(url string) (*Response, error) {
	response := Response{}
	res, err := h.oAuth2.Get(url, &response)
	if err != nil {
		if res != nil {
			return nil, h.printError(res)
		} else {
			return nil, err
		}

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

func (h *Http) GetCount(endpoint string, createdBefore *time.Time) (int64, error) {
	urlStr := fmt.Sprintf("%s/%s?$top=0&$inlinecount=allpages%s", h.BaseURL(), endpoint, h.DateFilter("Created", "lt", createdBefore, true, "&"))

	response, err := h.GetResponse(urlStr)
	if err != nil {
		return 0, err
	}

	count, err := strconv.ParseInt(response.Data.Count, 10, 64)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (h *Http) Get(url string, model interface{}) (string, error) {
	err := h.wait()
	if err != nil {
		return "", err
	}

	response, err := h.GetResponse(url)
	if err != nil {
		return "", err
	}
	/*
		response := Response{}
		res, err := h.oAuth2.Get(url, &response)
		if err != nil {
			if res != nil {
				return "", h.printError(res)
			} else {
				return "", err
			}

		}

		h.readRateLimitHeaders(res)*/

	err = json.Unmarshal(response.Data.Results, &model)
	if err != nil {
		return "", err
	}

	return response.Data.Next, nil
}

func (h *Http) PutValues(url string, values map[string]string) error {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(values)

	return h.Put(url, buf)
}

func (h *Http) PutBytes(url string, b []byte) error {
	return h.Put(url, bytes.NewBuffer(b))
}

func (h *Http) Put(url string, buf *bytes.Buffer) error {
	res, err := h.oAuth2.Put(url, buf, nil)
	if err != nil {
		if res != nil {
			return h.printError(res)
		} else {
			return err
		}
	}

	h.readRateLimitHeaders(res)

	return nil
}

func (h *Http) PostValues(url string, values map[string]string, model interface{}) error {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(values)

	return h.Post(url, buf, model)
}

func (h *Http) PostBytes(url string, b []byte, model interface{}) error {
	return h.Post(url, bytes.NewBuffer(b), model)
}

func (h *Http) Post(url string, buf *bytes.Buffer, model interface{}) error {
	response := ResponseSingle{}
	res, err := h.oAuth2.Post(url, buf, &response)
	if err != nil {
		if res != nil {
			return h.printError(res)
		} else {
			return err
		}
	}

	h.readRateLimitHeaders(res)

	defer res.Body.Close()

	err = json.Unmarshal(response.Data, &model)
	if err != nil {
		return err
	}

	return nil
}

func (h *Http) Delete(url string) error {
	res, err := h.oAuth2.Delete(url, nil, nil)
	if err != nil {
		if res != nil {
			return h.printError(res)
		} else {
			return err
		}
	}

	h.readRateLimitHeaders(res)

	return nil
}
