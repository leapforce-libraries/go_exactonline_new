package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// AgingReceivablesList stores AgingReceivablesList from exactonline
//
type AgingReceivablesList struct {
	AccountID            types.GUID `json:"AccountId"`
	AccountCode          string     `json:"AccountCode"`
	AccountName          string     `json:"AccountName"`
	AgeGroup1            int32      `json:"AgeGroup1"`
	AgeGroup1Amount      float64    `json:"AgeGroup1Amount"`
	AgeGroup1Description string     `json:"AgeGroup1Description"`
	AgeGroup2            int32      `json:"AgeGroup2"`
	AgeGroup2Amount      float64    `json:"AgeGroup2Amount"`
	AgeGroup2Description string     `json:"AgeGroup2Description"`
	AgeGroup3            int32      `json:"AgeGroup3"`
	AgeGroup3Amount      float64    `json:"AgeGroup3Amount"`
	AgeGroup3Description string     `json:"AgeGroup3Description"`
	AgeGroup4            int32      `json:"AgeGroup4"`
	AgeGroup4Amount      float64    `json:"AgeGroup4Amount"`
	AgeGroup4Description string     `json:"AgeGroup4Description"`
	CurrencyCode         string     `json:"CurrencyCode"`
	TotalAmount          float64    `json:"TotalAmount"`
}

type GetAgingReceivablesListsCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewGetAgingReceivablesListsCall() *GetAgingReceivablesListsCall {
	call := GetAgingReceivablesListsCall{}
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", AgingReceivablesList{})
	call.urlNext = service.urlRead(fmt.Sprintf("AgingReceivablesLists?$select=%s", selectFields))

	return &call
}

func (call *GetAgingReceivablesListsCall) Do() (*[]AgingReceivablesList, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	transactionLines := []AgingReceivablesList{}

	next, err := call.service.Get(call.urlNext, &transactionLines)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &transactionLines, nil
}
