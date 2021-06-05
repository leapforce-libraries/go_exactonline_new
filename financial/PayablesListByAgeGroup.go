package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// PayablesListByAgeGroup stores PayablesListByAgeGroup from exactonline
//
type PayablesListByAgeGroup struct {
	HID                string      `json:"HID"`
	AccountCode        string      `json:"AccountCode"`
	AccountID          types.GUID  `json:"AccountId"`
	AccountName        string      `json:"AccountName"`
	Amount             float64     `json:"Amount"`
	AmountInTransit    float64     `json:"AmountInTransit"`
	ApprovalStatus     int16       `json:"ApprovalStatus"`
	CurrencyCode       string      `json:"CurrencyCode"`
	Description        string      `json:"Description"`
	DueDate            *types.Date `json:"DueDate"`
	EntryNumber        int32       `json:"EntryNumber"`
	ID                 types.GUID  `json:"Id"`
	InvoiceDate        *types.Date `json:"InvoiceDate"`
	InvoiceNumber      int32       `json:"InvoiceNumber"`
	JournalCode        string      `json:"JournalCode"`
	JournalDescription string      `json:"JournalDescription"`
	YourRef            string      `json:"YourRef"`
}

type GetPayablesListByAgeGroupsCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewGetPayablesListByAgeGroupsCall() *GetPayablesListByAgeGroupsCall {
	call := GetPayablesListByAgeGroupsCall{}
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", PayablesListByAgeGroup{})
	call.urlNext = service.urlRead(fmt.Sprintf("PayablesListByAgeGroup?$select=%s", selectFields))

	return &call
}

func (call *GetPayablesListByAgeGroupsCall) Do() (*[]PayablesListByAgeGroup, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	transactionLines := []PayablesListByAgeGroup{}

	next, err := call.service.Get(call.urlNext, &transactionLines)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &transactionLines, nil
}
