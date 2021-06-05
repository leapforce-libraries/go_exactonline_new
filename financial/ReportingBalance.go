package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// ReportingBalance stores ReportingBalance from exactonline
//
type ReportingBalance struct {
	ID                    string     `json:"ID"`
	Amount                float64    `json:"Amount"`
	AmountCredit          float64    `json:"AmountCredit"`
	AmountDebit           float64    `json:"AmountDebit"`
	BalanceType           string     `json:"BalanceType"`
	CostCenterCode        string     `json:"CostCenterCode"`
	CostCenterDescription string     `json:"CostCenterDescription"`
	CostUnitCode          string     `json:"CostUnitCode"`
	CostUnitDescription   string     `json:"CostUnitDescription"`
	Count                 int32      `json:"Count"`
	Division              int32      `json:"Division"`
	GLAccount             types.GUID `json:"GLAccount"`
	GLAccountCode         string     `json:"GLAccountCode"`
	GLAccountDescription  string     `json:"GLAccountDescription"`
	ReportingPeriod       int32      `json:"ReportingPeriod"`
	ReportingYear         int32      `json:"ReportingYear"`
	Status                int32      `json:"Status"`
	Type                  int32      `json:"Type"`
}

type GetReportingBalancesCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewGetReportingBalancesCall() *GetReportingBalancesCall {
	call := GetReportingBalancesCall{}
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", ReportingBalance{})
	call.urlNext = service.url(fmt.Sprintf("ReportingBalance?$select=%s", selectFields))

	return &call
}

func (call *GetReportingBalancesCall) Do() (*[]ReportingBalance, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	transactionLines := []ReportingBalance{}

	next, err := call.service.Get(call.urlNext, &transactionLines)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &transactionLines, nil
}
