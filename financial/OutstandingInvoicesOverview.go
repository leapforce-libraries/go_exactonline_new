package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// OutstandingInvoicesOverview stores OutstandingInvoicesOverview from exactonline
//
type OutstandingInvoicesOverview struct {
	CurrencyCode                       string  `json:"CurrencyCode"`
	OutstandingPayableInvoiceAmount    float64 `json:"OutstandingPayableInvoiceAmount"`
	OutstandingPayableInvoiceCount     float64 `json:"OutstandingPayableInvoiceCount"`
	OutstandingReceivableInvoiceAmount float64 `json:"OutstandingReceivableInvoiceAmount"`
	OutstandingReceivableInvoiceCount  float64 `json:"OutstandingReceivableInvoiceCount"`
	OverduePayableInvoiceAmount        float64 `json:"OverduePayableInvoiceAmount"`
	OverduePayableInvoiceCount         float64 `json:"OverduePayableInvoiceCount"`
	OverdueReceivableInvoiceAmount     float64 `json:"OverdueReceivableInvoiceAmount"`
	OverdueReceivableInvoiceCount      float64 `json:"OverdueReceivableInvoiceCount"`
}

type GetOutstandingInvoicesOverviewsCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewGetOutstandingInvoicesOverviewsCall() *GetOutstandingInvoicesOverviewsCall {
	call := GetOutstandingInvoicesOverviewsCall{}
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", OutstandingInvoicesOverview{})
	call.urlNext = service.urlRead(fmt.Sprintf("OutstandingInvoicesOverview?$select=%s", selectFields))

	return &call
}

func (call *GetOutstandingInvoicesOverviewsCall) Do() (*[]OutstandingInvoicesOverview, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	transactionLines := []OutstandingInvoicesOverview{}

	next, err := call.service.Get(call.urlNext, &transactionLines)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &transactionLines, nil
}
