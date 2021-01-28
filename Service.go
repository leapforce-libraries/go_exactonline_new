package exactonline

import (
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	budget "github.com/leapforce-libraries/go_exactonline_new/budget"
	crm "github.com/leapforce-libraries/go_exactonline_new/crm"
	financialtransaction "github.com/leapforce-libraries/go_exactonline_new/financialtransaction"
	go_http "github.com/leapforce-libraries/go_exactonline_new/http"
	logistics "github.com/leapforce-libraries/go_exactonline_new/logistics"
	salesorder "github.com/leapforce-libraries/go_exactonline_new/salesorder"
	subscription "github.com/leapforce-libraries/go_exactonline_new/subscription"
	google "github.com/leapforce-libraries/go_google"
	bigquery "github.com/leapforce-libraries/go_google/bigquery"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

// Service stores GoogleService configuration
//
type Service struct {
	BudgetService               *budget.Service
	CRMService                  *crm.Service
	FinancialTransactionService *financialtransaction.Service
	LogisticsService            *logistics.Service
	SalesOrderService           *salesorder.Service
	SubscriptionService         *subscription.Service
	oAuth2                      *oauth2.OAuth2
}

type ServiceConfig struct {
	Division     int32
	ClientID     string
	ClientSecret string
}

const (
	APIName            string = "ExactOnline"
	APIURL             string = "https://start.exactonline.nl/api/v1"
	AuthURL            string = "https://start.exactonline.nl/api/oauth2/auth"
	TokenURL           string = "https://start.exactonline.nl/api/oauth2/token"
	TokenHTTPMethod    string = http.MethodPost
	RedirectURL        string = "http://localhost:8080/oauth/redirect"
	TableRefreshToken  string = "leapforce.oauth2"
	LastModifiedFormat string = "2006-01-02T15:04:05"
)

// methods
//
func NewService(serviceConfig ServiceConfig, bigQueryService *bigquery.Service) *Service {
	getTokenFunction := func() (*oauth2.Token, *errortools.Error) {
		return google.GetToken(APIName, serviceConfig.ClientID, bigQueryService)
	}

	saveTokenFunction := func(token *oauth2.Token) *errortools.Error {
		return google.SaveToken(APIName, serviceConfig.ClientID, token, bigQueryService)
	}

	maxRetries := uint(3)
	oauht2Config := oauth2.OAuth2Config{
		ClientID:          serviceConfig.ClientID,
		ClientSecret:      serviceConfig.ClientSecret,
		RedirectURL:       RedirectURL,
		AuthURL:           AuthURL,
		TokenURL:          TokenURL,
		TokenHTTPMethod:   TokenHTTPMethod,
		GetTokenFunction:  &getTokenFunction,
		SaveTokenFunction: &saveTokenFunction,
		MaxRetries:        &maxRetries,
	}
	oAuth2 := oauth2.NewOAuth(oauht2Config)
	httpService := go_http.NewService(serviceConfig.Division, oAuth2)

	return &Service{
		BudgetService:               budget.NewService(httpService),
		CRMService:                  crm.NewService(httpService),
		FinancialTransactionService: financialtransaction.NewService(httpService),
		LogisticsService:            logistics.NewService(httpService),
		SalesOrderService:           salesorder.NewService(httpService),
		SubscriptionService:         subscription.NewService(httpService),
		oAuth2:                      oAuth2,
	}
}

func (service *Service) InitToken() *errortools.Error {
	return service.oAuth2.InitToken()
}

func (service *Service) Get(requestConfig *oauth2.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	err := ErrorResponse{}
	request, response, e := service.oAuth2.Get(requestConfig)
	return request, response, service.captureError(e, &err)
}

func (service *Service) Post(requestConfig *oauth2.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	err := ErrorResponse{}
	request, response, e := service.oAuth2.Post(requestConfig)
	return request, response, service.captureError(e, &err)
}

func (service *Service) Put(requestConfig *oauth2.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	err := ErrorResponse{}
	request, response, e := service.oAuth2.Put(requestConfig)
	return request, response, service.captureError(e, &err)
}

func (service *Service) Patch(requestConfig *oauth2.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	err := ErrorResponse{}
	request, response, e := service.oAuth2.Patch(requestConfig)
	return request, response, service.captureError(e, &err)
}

func (service *Service) Delete(requestConfig *oauth2.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	err := ErrorResponse{}
	request, response, e := service.oAuth2.Delete(requestConfig)
	return request, response, service.captureError(e, &err)
}

func (service *Service) captureError(e *errortools.Error, err *ErrorResponse) *errortools.Error {
	if e == nil || err == nil {
		return nil
	}

	if err.Error.Message.Value != "" {
		e.SetMessage(err.Error.Message.Value)
	}

	return e
}

func (service *Service) ValidateToken() (*oauth2.Token, *errortools.Error) {
	return service.oAuth2.ValidateToken()
}
