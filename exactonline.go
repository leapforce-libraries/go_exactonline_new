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

// ExactOnline stores ExactOnline configuration
//
type ExactOnline struct {
	BudgetClient               *budget.Client
	CRMClient                  *crm.Client
	FinancialTransactionClient *financialtransaction.Client
	LogisticsClient            *logistics.Client
	SalesOrderClient           *salesorder.Client
	SubscriptionClient         *subscription.Client
	http                       *http.Http
}

// methods
//
func NewExactOnline(division int32, clientID string, clientSecret string, bigQuery *google.BigQuery) (*ExactOnline, *errortools.Error) {
	eo := ExactOnline{}

	http, err := http.NewHttp(division, clientID, clientSecret, bigQuery)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}
	eo.http = http
	eo.BudgetClient = budget.NewClient(http)
	eo.CRMClient = crm.NewClient(http)
	eo.FinancialTransactionClient = financialtransaction.NewClient(http)
	eo.LogisticsClient = logistics.NewClient(http)
	eo.SalesOrderClient = salesorder.NewClient(http)
	eo.SubscriptionClient = subscription.NewClient(http)

	return &eo, nil
}

func (eo *ExactOnline) InitToken() *errortools.Error {
	return eo.http.InitToken()
}
