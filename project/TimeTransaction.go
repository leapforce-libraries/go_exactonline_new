package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// TimeTransaction stores TimeTransaction from exactonline
//
type TimeTransaction struct {
	ID                      types.Guid  `json:"ID"`
	Account                 types.Guid  `json:"Account"`
	AccountName             string      `json:"AccountName"`
	Activity                types.Guid  `json:"Activity"`
	ActivityDescription     string      `json:"ActivityDescription"`
	Amount                  float64     `json:"Amount"`
	AmountFC                float64     `json:"AmountFC"`
	Attachment              types.Guid  `json:"Attachment"`
	Created                 *types.Date `json:"Created,omitempty"`
	Creator                 types.Guid  `json:"Creator"`
	CreatorFullName         string      `json:"CreatorFullName"`
	Currency                string      `json:"Currency"`
	Date                    *types.Date `json:"Date,omitempty"`
	Division                int64       `json:"Division"`
	DivisionDescription     string      `json:"DivisionDescription"`
	Employee                types.Guid  `json:"Employee"`
	EndTime                 *types.Date `json:"EndTime,omitempty"`
	EntryNumber             int64       `json:"EntryNumber"`
	ErrorText               string      `json:"ErrorText"`
	HourStatus              int64       `json:"HourStatus"`
	Item                    types.Guid  `json:"Item"`
	ItemDescription         string      `json:"ItemDescription"`
	ItemDivisable           bool        `json:"ItemDivisable"`
	Modified                *types.Date `json:"Modified,omitempty"`
	Modifier                types.Guid  `json:"Modifier"`
	ModifierFullName        string      `json:"ModifierFullName"`
	Notes                   string      `json:"Notes"`
	Price                   float64     `json:"Price"`
	PriceFC                 float64     `json:"PriceFC"`
	Project                 types.Guid  `json:"Project"`
	ProjectAccount          types.Guid  `json:"ProjectAccount"`
	ProjectAccountCode      string      `json:"ProjectAccountCode"`
	ProjectAccountName      string      `json:"ProjectAccountName"`
	ProjectCode             string      `json:"ProjectCode"`
	ProjectDescription      string      `json:"ProjectDescription"`
	Quantity                float64     `json:"Quantity"`
	StartTime               *types.Date `json:"StartTime,omitempty"`
	Subscription            types.Guid  `json:"Subscription"`
	SubscriptionAccount     types.Guid  `json:"SubscriptionAccount"`
	SubscriptionAccountCode string      `json:"SubscriptionAccountCode"`
	SubscriptionAccountName string      `json:"SubscriptionAccountName"`
	SubscriptionDescription string      `json:"SubscriptionDescription"`
	SubscriptionNumber      int64       `json:"SubscriptionNumber"`
	Type                    int64       `json:"Type"`
}

type GetTimeTransactionsCall struct {
	modifiedAfter *time.Time
	urlNext       string
	service       *Service
}

func (service *Service) NewGetTimeTransactionsCall(modifiedAfter *time.Time) *GetTimeTransactionsCall {
	call := GetTimeTransactionsCall{}
	call.modifiedAfter = modifiedAfter
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", TimeTransaction{})
	call.urlNext = service.url(fmt.Sprintf("TimeTransactions?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetTimeTransactionsCall) Do() (*[]TimeTransaction, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	timeTransactions := []TimeTransaction{}

	next, err := call.service.Get(call.urlNext, &timeTransactions)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &timeTransactions, nil
}

func (service *Service) GetTimeTransactionsCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("TimeTransactions", createdBefore)
}
