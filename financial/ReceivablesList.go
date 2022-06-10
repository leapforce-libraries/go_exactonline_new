package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// ReceivablesList stores ReceivablesList from exactonline
//
type ReceivablesList struct {
	HID                int64       `json:"HID"`
	AccountCode        string      `json:"AccountCode"`
	AccountId          types.Guid  `json:"AccountId"`
	AccountName        string      `json:"AccountName"`
	Amount             float64     `json:"Amount"`
	AmountInTransit    float64     `json:"AmountInTransit"`
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

type GetReceivablesListsCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewGetReceivablesListsCall() *GetReceivablesListsCall {
	call := GetReceivablesListsCall{}
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", ReceivablesList{})
	call.urlNext = service.urlRead(fmt.Sprintf("ReceivablesList?$select=%s", selectFields))

	return &call
}

func (call *GetReceivablesListsCall) Do() (*[]ReceivablesList, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	receivablesList := []ReceivablesList{}

	next, err := call.service.Get(call.urlNext, &receivablesList)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &receivablesList, nil
}
