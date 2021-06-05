package exactonline

import (
	"fmt"
	"strconv"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// ReportingBalanceByClassification stores ReportingBalanceByClassification from exactonline
//
type ReportingBalanceByClassification struct {
	ID                        string     `json:"ID"`
	Amount                    float64    `json:"Amount"`
	AmountCredit              float64    `json:"AmountCredit"`
	AmountDebit               float64    `json:"AmountDebit"`
	BalanceType               string     `json:"BalanceType"`
	ClassificationCode        string     `json:"ClassificationCode"`
	ClassificationDescription string     `json:"ClassificationDescription"`
	CostCenterCode            string     `json:"CostCenterCode"`
	CostCenterDescription     string     `json:"CostCenterDescription"`
	CostUnitCode              string     `json:"CostUnitCode"`
	CostUnitDescription       string     `json:"CostUnitDescription"`
	Count                     int32      `json:"Count"`
	Division                  int32      `json:"Division"`
	GLAccount                 types.GUID `json:"GLAccount"`
	GLAccountCode             string     `json:"GLAccountCode"`
	GLAccountDescription      string     `json:"GLAccountDescription"`
	GLScheme                  types.GUID `json:"GLScheme"`
	ReportingPeriod           int32      `json:"ReportingPeriod"`
	ReportingYear             int32      `json:"ReportingYear"`
	Status                    int32      `json:"Status"`
	Type                      int32      `json:"Type"`
}

type GetReportingBalanceByClassificationsCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewGetReportingBalanceByClassificationsCall(glScheme *GLScheme, reportingYear int) *GetReportingBalanceByClassificationsCall {
	call := GetReportingBalanceByClassificationsCall{}
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", ReportingBalanceByClassification{})

	call.urlNext = service.urlRead(fmt.Sprintf("ReportingBalanceByClassification?glScheme=guid'%s'&reportingYear=%s&$select=%s", glScheme.ID.String(), strconv.Itoa(reportingYear), selectFields))

	return &call
}

func (call *GetReportingBalanceByClassificationsCall) Do() (*[]ReportingBalanceByClassification, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	reportingBalanceByClassifications := []ReportingBalanceByClassification{}

	next, err := call.service.Get(call.urlNext, &reportingBalanceByClassifications)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &reportingBalanceByClassifications, nil
}
