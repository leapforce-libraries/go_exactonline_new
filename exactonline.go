package exactonline

import (
	bigquerytools "github.com/leapforce-libraries/go_bigquerytools"
	budget "github.com/leapforce-libraries/go_exactonline_new/budget"
	crm "github.com/leapforce-libraries/go_exactonline_new/crm"
	financialtransaction "github.com/leapforce-libraries/go_exactonline_new/financialtransaction"
	http "github.com/leapforce-libraries/go_exactonline_new/http"
	logistics "github.com/leapforce-libraries/go_exactonline_new/logistics"
	salesorder "github.com/leapforce-libraries/go_exactonline_new/salesorder"
	subscription "github.com/leapforce-libraries/go_exactonline_new/subscription"
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
}

// methods
//
func NewExactOnline(division int32, clientID string, clientSecret string, bigQuery *bigquerytools.BigQuery, isLive bool) (*ExactOnline, error) {
	eo := ExactOnline{}

	http, err := http.NewHttp(division, clientID, clientSecret, bigQuery, isLive)
	if err != nil {
		return nil, err
	}
	eo.BudgetClient = budget.NewClient(http)
	eo.CRMClient = crm.NewClient(http)
	eo.FinancialTransactionClient = financialtransaction.NewClient(http)
	eo.LogisticsClient = logistics.NewClient(http)
	eo.SalesOrderClient = salesorder.NewClient(http)
	eo.SubscriptionClient = subscription.NewClient(http)

	return &eo, nil
}
