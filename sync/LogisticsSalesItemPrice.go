package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// LogisticsSalesItemPrice stores LogisticsSalesItemPrice from exactonline
//
type LogisticsSalesItemPrice struct {
	Timestamp                  types.Int64String `json:"Timestamp"`
	Account                    types.GUID        `json:"Account"`
	AccountName                string            `json:"AccountName"`
	Created                    *types.Date       `json:"Created"`
	Creator                    types.GUID        `json:"Creator"`
	CreatorFullName            string            `json:"CreatorFullName"`
	Currency                   string            `json:"Currency"`
	DefaultItemUnit            string            `json:"DefaultItemUnit"`
	DefaultItemUnitDescription string            `json:"DefaultItemUnitDescription"`
	Division                   int32             `json:"Division"`
	EndDate                    *types.Date       `json:"EndDate"`
	ID                         types.GUID        `json:"ID"`
	Item                       types.GUID        `json:"Item"`
	ItemCode                   string            `json:"ItemCode"`
	ItemDescription            string            `json:"ItemDescription"`
	Modified                   *types.Date       `json:"Modified"`
	Modifier                   types.GUID        `json:"Modifier"`
	ModifierFullName           string            `json:"ModifierFullName"`
	NumberOfItemsPerUnit       float64           `json:"NumberOfItemsPerUnit"`
	Price                      float64           `json:"Price"`
	Quantity                   float64           `json:"Quantity"`
	StartDate                  *types.Date       `json:"StartDate"`
	Unit                       string            `json:"Unit"`
	UnitDescription            string            `json:"UnitDescription"`
}

type SyncLogisticsSalesItemPricesCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewSyncLogisticsSalesItemPricesCall(timestamp *int64) *SyncLogisticsSalesItemPricesCall {
	selectFields := utilities.GetTaggedTagNames("json", LogisticsSalesItemPrice{})
	url := service.url(fmt.Sprintf("Logistics/SalesItemPrices?$select=%s", selectFields))
	if timestamp != nil {
		url = fmt.Sprintf("%s&$filter=Timestamp gt %vL", url, *timestamp)
	}

	return &SyncLogisticsSalesItemPricesCall{
		service: service,
		urlNext: url,
	}
}

func (call *SyncLogisticsSalesItemPricesCall) Do() (*[]LogisticsSalesItemPrice, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	logisticsSalesItemPrices := []LogisticsSalesItemPrice{}

	next, err := call.service.Get(call.urlNext, &logisticsSalesItemPrices)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &logisticsSalesItemPrices, nil
}
