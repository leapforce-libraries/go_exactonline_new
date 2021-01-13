package exactonline

import (
	errortools "github.com/leapforce-libraries/go_errortools"
	budget "github.com/leapforce-libraries/go_exactonline_new/budget"
	crm "github.com/leapforce-libraries/go_exactonline_new/crm"
	financialtransaction "github.com/leapforce-libraries/go_exactonline_new/financialtransaction"
	http "github.com/leapforce-libraries/go_exactonline_new/http"
	logistics "github.com/leapforce-libraries/go_exactonline_new/logistics"
	salesorder "github.com/leapforce-libraries/go_exactonline_new/salesorder"
	subscription "github.com/leapforce-libraries/go_exactonline_new/subscription"
	google "github.com/leapforce-libraries/go_google"
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
	http                        *http.Http
}

// methods
//
func NewService(division int32, clientID string, clientSecret string, bigQuery *google.BigQuery) (*Service, *errortools.Error) {
	eo := Service{}

	http, err := http.NewHttp(division, clientID, clientSecret, bigQuery)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}
	eo.http = http
	eo.BudgetService = budget.NewService(http)
	eo.CRMService = crm.NewService(http)
	eo.FinancialTransactionService = financialtransaction.NewService(http)
	eo.LogisticsService = logistics.NewService(http)
	eo.SalesOrderService = salesorder.NewService(http)
	eo.SubscriptionService = subscription.NewService(http)

	return &eo, nil
}

func (service *Service) InitToken() *errortools.Error {
	return service.http.InitToken()
}
