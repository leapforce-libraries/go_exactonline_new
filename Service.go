package exactonline

import (
	"net/http"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	assets "github.com/leapforce-libraries/go_exactonline_new/assets"
	budget "github.com/leapforce-libraries/go_exactonline_new/budget"
	cashflow "github.com/leapforce-libraries/go_exactonline_new/cashflow"
	crm "github.com/leapforce-libraries/go_exactonline_new/crm"
	financialtransaction "github.com/leapforce-libraries/go_exactonline_new/financialtransaction"
	eo_http "github.com/leapforce-libraries/go_exactonline_new/http"
	logistics "github.com/leapforce-libraries/go_exactonline_new/logistics"
	purchaseentry "github.com/leapforce-libraries/go_exactonline_new/purchaseentry"
	purchaseorder "github.com/leapforce-libraries/go_exactonline_new/purchaseorder"
	salesorder "github.com/leapforce-libraries/go_exactonline_new/salesorder"
	subscription "github.com/leapforce-libraries/go_exactonline_new/subscription"
	sync "github.com/leapforce-libraries/go_exactonline_new/sync"
	google "github.com/leapforce-libraries/go_google"
	bigquery "github.com/leapforce-libraries/go_google/bigquery"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

const (
	apiName         string = "ExactOnline"
	apiURL          string = "https://start.exactonline.nl/api/v1"
	authURL         string = "https://start.exactonline.nl/api/oauth2/auth"
	tokenURL        string = "https://start.exactonline.nl/api/oauth2/token"
	tokenHTTPMethod string = http.MethodPost
	redirectURL     string = "http://localhost:8080/oauth/redirect"
)

// Service stores Service configuration
//
type Service struct {
	clientID                    string
	AssetsService               *assets.Service
	BudgetService               *budget.Service
	CashflowService             *cashflow.Service
	CRMService                  *crm.Service
	FinancialTransactionService *financialtransaction.Service
	LogisticsService            *logistics.Service
	PurchaseEntryService        *purchaseentry.Service
	PurchaseOrderService        *purchaseorder.Service
	SalesOrderService           *salesorder.Service
	SubscriptionService         *subscription.Service
	SyncService                 *sync.Service
	oAuth2Service               *oauth2.Service
}

type ServiceConfig struct {
	Division     int32
	ClientID     string
	ClientSecret string
}

// methods
//
func NewService(serviceConfig *ServiceConfig, bigQueryService *bigquery.Service) (*Service, *errortools.Error) {
	if serviceConfig == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	if serviceConfig.ClientID == "" {
		return nil, errortools.ErrorMessage("ClientID not provided")
	}

	if serviceConfig.ClientSecret == "" {
		return nil, errortools.ErrorMessage("ClientSecret not provided")
	}

	getTokenFunction := func() (*oauth2.Token, *errortools.Error) {
		return google.GetToken(apiName, serviceConfig.ClientID, bigQueryService)
	}

	saveTokenFunction := func(token *oauth2.Token) *errortools.Error {
		return google.SaveToken(apiName, serviceConfig.ClientID, token, bigQueryService)
	}

	oauth2ServiceConfig := oauth2.ServiceConfig{
		ClientID:          serviceConfig.ClientID,
		ClientSecret:      serviceConfig.ClientSecret,
		RedirectURL:       redirectURL,
		AuthURL:           authURL,
		TokenURL:          tokenURL,
		TokenHTTPMethod:   tokenHTTPMethod,
		GetTokenFunction:  &getTokenFunction,
		SaveTokenFunction: &saveTokenFunction,
	}
	oAuth2Service, e := oauth2.NewService(&oauth2ServiceConfig)
	if e != nil {
		return nil, e
	}
	httpService := eo_http.NewService(serviceConfig.Division, oAuth2Service)

	return &Service{
		clientID:                    serviceConfig.ClientID,
		AssetsService:               assets.NewService(httpService),
		BudgetService:               budget.NewService(httpService),
		CashflowService:             cashflow.NewService(httpService),
		CRMService:                  crm.NewService(httpService),
		FinancialTransactionService: financialtransaction.NewService(httpService),
		LogisticsService:            logistics.NewService(httpService),
		PurchaseEntryService:        purchaseentry.NewService(httpService),
		PurchaseOrderService:        purchaseorder.NewService(httpService),
		SalesOrderService:           salesorder.NewService(httpService),
		SubscriptionService:         subscription.NewService(httpService),
		SyncService:                 sync.NewService(httpService),
		oAuth2Service:               oAuth2Service,
	}, nil
}

func (service *Service) ValidateToken() (*oauth2.Token, *errortools.Error) {
	return service.oAuth2Service.ValidateToken()
}

func (service *Service) InitToken(scope string) *errortools.Error {
	return service.oAuth2Service.InitToken(scope)
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

func (service Service) APIName() string {
	return apiName
}

func (service Service) APIKey() string {
	return service.clientID
}

func (service Service) APICallCount() int64 {
	return service.oAuth2Service.APICallCount()
}

func (service Service) APIReset() {
	service.oAuth2Service.APIReset()
}
