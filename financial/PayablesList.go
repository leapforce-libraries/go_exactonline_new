package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// PayablesList stores PayablesList from exactonline
//
type PayablesList struct {
	HID                string      `json:"HID"`
	AccountCode        string      `json:"AccountCode"`
	AccountId          types.Guid  `json:"AccountId"`
	AccountName        string      `json:"AccountName"`
	Amount             float64     `json:"Amount"`
	AmountInTransit    float64     `json:"AmountInTransit"`
	ApprovalStatus     int16       `json:"ApprovalStatus"`
	CurrencyCode       string      `json:"CurrencyCode"`
	Description        string      `json:"Description"`
	DueDate            *types.Date `json:"DueDate"`
	EntryNumber        int32       `json:"EntryNumber"`
	Id                 types.Guid  `json:"Id"`
	InvoiceDate        *types.Date `json:"InvoiceDate"`
	InvoiceNumber      int32       `json:"InvoiceNumber"`
	JournalCode        string      `json:"JournalCode"`
	JournalDescription string      `json:"JournalDescription"`
	YourRef            string      `json:"YourRef"`
}

type GetPayablesListsCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewGetPayablesListsCall() *GetPayablesListsCall {
	call := GetPayablesListsCall{}
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", PayablesList{})
	call.urlNext = service.urlRead(fmt.Sprintf("PayablesList?$select=%s", selectFields))

	return &call
}

func (call *GetPayablesListsCall) Do() (*[]PayablesList, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	payablesList := []PayablesList{}

	next, err := call.service.Get(call.urlNext, &payablesList)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &payablesList, nil
}
