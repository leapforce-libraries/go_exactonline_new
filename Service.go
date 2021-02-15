package exactonline

import (
	"net/http"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	budget "github.com/leapforce-libraries/go_exactonline_new/budget"
	crm "github.com/leapforce-libraries/go_exactonline_new/crm"
	financialtransaction "github.com/leapforce-libraries/go_exactonline_new/financialtransaction"
	eo_http "github.com/leapforce-libraries/go_exactonline_new/http"
	logistics "github.com/leapforce-libraries/go_exactonline_new/logistics"
	salesorder "github.com/leapforce-libraries/go_exactonline_new/salesorder"
	subscription "github.com/leapforce-libraries/go_exactonline_new/subscription"
	google "github.com/leapforce-libraries/go_google"
	bigquery "github.com/leapforce-libraries/go_google/bigquery"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

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

// Service stores Service configuration
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
	Division              int32
	ClientID              string
	ClientSecret          string
	MaxRetries            *uint
	SecondsBetweenRetries *uint32
}

// methods
//
func NewService(serviceConfig ServiceConfig, bigQueryService *bigquery.Service) (*Service, *errortools.Error) {
	if serviceConfig.ClientID == "" {
		return nil, errortools.ErrorMessage("ClientID not provided")
	}

	if serviceConfig.ClientSecret == "" {
		return nil, errortools.ErrorMessage("ClientSecret not provided")
	}

	getTokenFunction := func() (*oauth2.Token, *errortools.Error) {
		return google.GetToken(APIName, serviceConfig.ClientID, bigQueryService)
	}

	saveTokenFunction := func(token *oauth2.Token) *errortools.Error {
		return google.SaveToken(APIName, serviceConfig.ClientID, token, bigQueryService)
	}

	oauht2Config := oauth2.OAuth2Config{
		ClientID:              serviceConfig.ClientID,
		ClientSecret:          serviceConfig.ClientSecret,
		RedirectURL:           RedirectURL,
		AuthURL:               AuthURL,
		TokenURL:              TokenURL,
		TokenHTTPMethod:       TokenHTTPMethod,
		GetTokenFunction:      &getTokenFunction,
		SaveTokenFunction:     &saveTokenFunction,
		MaxRetries:            serviceConfig.MaxRetries,
		SecondsBetweenRetries: serviceConfig.SecondsBetweenRetries,
	}
	oAuth2 := oauth2.NewOAuth(oauht2Config)
	httpService := eo_http.NewService(serviceConfig.Division, oAuth2)

	return &Service{
		BudgetService:               budget.NewService(httpService),
		CRMService:                  crm.NewService(httpService),
		FinancialTransactionService: financialtransaction.NewService(httpService),
		LogisticsService:            logistics.NewService(httpService),
		SalesOrderService:           salesorder.NewService(httpService),
		SubscriptionService:         subscription.NewService(httpService),
		oAuth2:                      oAuth2,
	}, nil
}

func (service *Service) ValidateToken() (*oauth2.Token, *errortools.Error) {
	return service.oAuth2.ValidateToken()
}

func (service *Service) InitToken() *errortools.Error {
	return service.oAuth2.InitToken()
}

func ParseDateString(date string) *time.Time {
	if len(date) >= 19 {
		d, err := time.Parse("2006-01-02T15:04:05", date[:19])
		if err == nil {
			return &d
		}
	}

	return nil
}
