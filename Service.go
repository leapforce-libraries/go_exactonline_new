package exactonline

import (
	"net/http"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	assets "github.com/leapforce-libraries/go_exactonline_new/assets"
	budget "github.com/leapforce-libraries/go_exactonline_new/budget"
	cashflow "github.com/leapforce-libraries/go_exactonline_new/cashflow"
	crm "github.com/leapforce-libraries/go_exactonline_new/crm"
	financial "github.com/leapforce-libraries/go_exactonline_new/financial"
	financialtransaction "github.com/leapforce-libraries/go_exactonline_new/financialtransaction"
	eo_http "github.com/leapforce-libraries/go_exactonline_new/http"
	inventory "github.com/leapforce-libraries/go_exactonline_new/inventory"
	logistics "github.com/leapforce-libraries/go_exactonline_new/logistics"
	payroll "github.com/leapforce-libraries/go_exactonline_new/payroll"
	project "github.com/leapforce-libraries/go_exactonline_new/project"
	purchaseentry "github.com/leapforce-libraries/go_exactonline_new/purchaseentry"
	purchaseorder "github.com/leapforce-libraries/go_exactonline_new/purchaseorder"
	salesinvoice "github.com/leapforce-libraries/go_exactonline_new/salesinvoice"
	salesorder "github.com/leapforce-libraries/go_exactonline_new/salesorder"
	subscription "github.com/leapforce-libraries/go_exactonline_new/subscription"
	sync "github.com/leapforce-libraries/go_exactonline_new/sync"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
	tokensource "github.com/leapforce-libraries/go_oauth2/tokensource"
)

const (
	apiName            string = "ExactOnline"
	apiUrl             string = "https://start.exactonline.nl/api/v1"
	authUrl            string = "https://start.exactonline.nl/api/oauth2/auth"
	tokenUrl           string = "https://start.exactonline.nl/api/oauth2/token"
	tokenHTTPMethod    string = http.MethodPost
	defaultRedirectUrl string = "http://localhost:8080/oauth/redirect"
	// You can only request for a new access token after 570 seconds from the time you successfully received the previous access token.
	// see: https://support.exactonline.com/community/s/knowledge-base#All-All-DNO-Simulation-gen-apilimits
	refreshMargin time.Duration = 30 * time.Second
)

// Service stores Service configuration
//
type Service struct {
	clientID                    string
	AssetsService               *assets.Service
	BudgetService               *budget.Service
	CashflowService             *cashflow.Service
	CRMService                  *crm.Service
	FinancialService            *financial.Service
	FinancialTransactionService *financialtransaction.Service
	InventoryService            *inventory.Service
	LogisticsService            *logistics.Service
	PayrollService              *payroll.Service
	ProjectService              *project.Service
	PurchaseEntryService        *purchaseentry.Service
	PurchaseOrderService        *purchaseorder.Service
	SalesInvoiceService         *salesinvoice.Service
	SalesOrderService           *salesorder.Service
	SubscriptionService         *subscription.Service
	SyncService                 *sync.Service
	oAuth2Service               *oauth2.Service
}

type ServiceConfig struct {
	Division     int32
	ClientID     string
	ClientSecret string
	TokenSource  tokensource.TokenSource
	RedirectUrl  *string
}

// methods
//
func NewService(serviceConfig *ServiceConfig) (*Service, *errortools.Error) {
	if serviceConfig == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	if serviceConfig.ClientID == "" {
		return nil, errortools.ErrorMessage("ClientID not provided")
	}

	if serviceConfig.ClientSecret == "" {
		return nil, errortools.ErrorMessage("ClientSecret not provided")
	}

	redirectUrl := defaultRedirectUrl
	if serviceConfig.RedirectUrl != nil {
		redirectUrl = *serviceConfig.RedirectUrl
	}

	_refreshMargin := refreshMargin
	oauth2ServiceConfig := oauth2.ServiceConfig{
		ClientId:        serviceConfig.ClientID,
		ClientSecret:    serviceConfig.ClientSecret,
		RedirectUrl:     redirectUrl,
		AuthUrl:         authUrl,
		TokenUrl:        tokenUrl,
		RefreshMargin:   &_refreshMargin,
		TokenHttpMethod: tokenHTTPMethod,
		TokenSource:     serviceConfig.TokenSource,
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
		FinancialService:            financial.NewService(httpService),
		FinancialTransactionService: financialtransaction.NewService(httpService),
		InventoryService:            inventory.NewService(httpService),
		LogisticsService:            logistics.NewService(httpService),
		PayrollService:              payroll.NewService(httpService),
		ProjectService:              project.NewService(httpService),
		PurchaseEntryService:        purchaseentry.NewService(httpService),
		PurchaseOrderService:        purchaseorder.NewService(httpService),
		SalesInvoiceService:         salesinvoice.NewService(httpService),
		SalesOrderService:           salesorder.NewService(httpService),
		SubscriptionService:         subscription.NewService(httpService),
		SyncService:                 sync.NewService(httpService),
		oAuth2Service:               oAuth2Service,
	}, nil
}

func (service *Service) AuthorizeUrl(scope string, accessType *string, prompt *string, state *string) string {
	return service.oAuth2Service.AuthorizeUrl(scope, accessType, prompt, state)
}

/*func (service *Service) ValidateToken() (*oauth2.Token, *errortools.Error) {
	return service.oAuth2Service.ValidateToken()
}

func (service *Service) InitToken(scope string, accessType *string, prompt *string, state *string) *errortools.Error {
	return service.oAuth2Service.InitToken(scope, accessType, prompt, state)
}*/

func (service *Service) GetTokenFromCode(r *http.Request) *errortools.Error {
	return service.oAuth2Service.GetTokenFromCode(r, nil)
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

func (service Service) ApiName() string {
	return apiName
}

func (service Service) ApiKey() string {
	return service.clientID
}

func (service Service) ApiCallCount() int64 {
	return service.oAuth2Service.ApiCallCount()
}

func (service Service) ApiReset() {
	service.oAuth2Service.ApiReset()
}
