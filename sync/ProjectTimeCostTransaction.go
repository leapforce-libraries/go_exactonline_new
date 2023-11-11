package exactonline

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// ProjectTimeCostTransaction stores ProjectTimeCostTransaction from exactonline
type ProjectTimeCostTransaction struct {
	Timestamp               types.Int64String `json:"Timestamp"`
	Account                 types.Guid        `json:"Account"`
	AccountName             string            `json:"AccountName"`
	AmountFC                float64           `json:"AmountFC"`
	Attachment              types.Guid        `json:"Attachment"`
	Created                 *types.Date       `json:"Created"`
	Creator                 types.Guid        `json:"Creator"`
	CreatorFullName         string            `json:"CreatorFullName"`
	Currency                string            `json:"Currency"`
	CustomField             string            `json:"CustomField"`
	Date                    *types.Date       `json:"Date"`
	Division                int32             `json:"Division"`
	DivisionDescription     string            `json:"DivisionDescription"`
	Employee                types.Guid        `json:"Employee"`
	EndTime                 *types.Date       `json:"EndTime"`
	EntryNumber             int32             `json:"EntryNumber"`
	ErrorText               string            `json:"ErrorText"`
	HourStatus              int16             `json:"HourStatus"`
	ID                      types.Guid        `json:"ID"`
	Item                    types.Guid        `json:"Item"`
	ItemDescription         string            `json:"ItemDescription"`
	ItemDivisable           bool              `json:"ItemDivisable"`
	Modified                *types.Date       `json:"Modified"`
	Modifier                types.Guid        `json:"Modifier"`
	ModifierFullName        string            `json:"ModifierFullName"`
	Notes                   string            `json:"Notes"`
	PriceFC                 float64           `json:"PriceFC"`
	Project                 types.Guid        `json:"Project"`
	ProjectAccount          types.Guid        `json:"ProjectAccount"`
	ProjectAccountCode      string            `json:"ProjectAccountCode"`
	ProjectAccountName      string            `json:"ProjectAccountName"`
	ProjectCode             string            `json:"ProjectCode"`
	ProjectDescription      string            `json:"ProjectDescription"`
	Quantity                float64           `json:"Quantity"`
	StartTime               *types.Date       `json:"StartTime"`
	Subscription            types.Guid        `json:"Subscription"`
	SubscriptionAccount     types.Guid        `json:"SubscriptionAccount"`
	SubscriptionAccountCode string            `json:"SubscriptionAccountCode"`
	SubscriptionAccountName string            `json:"SubscriptionAccountName"`
	SubscriptionDescription string            `json:"SubscriptionDescription"`
	SubscriptionNumber      int32             `json:"SubscriptionNumber"`
	Type                    int16             `json:"Type"`
	WBS                     types.Guid        `json:"WBS"`
	WBSDescription          string            `json:"WBSDescription"`
}

type SyncProjectTimeCostTransactionsCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewSyncProjectTimeCostTransactionsCall(timestamp *int64) *SyncProjectTimeCostTransactionsCall {
	selectFields := utilities.GetTaggedTagNames("json", ProjectTimeCostTransaction{})
	url := service.url(fmt.Sprintf("Project/TimeCostTransactions?$select=%s", selectFields))
	if timestamp != nil {
		url = fmt.Sprintf("%s&$filter=Timestamp gt %vL", url, *timestamp)
	}

	return &SyncProjectTimeCostTransactionsCall{
		service: service,
		urlNext: url,
	}
}

func (call *SyncProjectTimeCostTransactionsCall) Do() (*[]ProjectTimeCostTransaction, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	projectTimeCostTransactions := []ProjectTimeCostTransaction{}

	next, err := call.service.Get(call.urlNext, &projectTimeCostTransactions)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &projectTimeCostTransactions, nil
}
