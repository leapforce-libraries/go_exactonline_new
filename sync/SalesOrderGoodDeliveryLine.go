package exactonline

import (
	"encoding/json"
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// SalesOrderGoodDeliveryLine stores SalesOrderGoodDeliveryLine from exactonline
//
type SalesOrderGoodDeliveryLine struct {
	Timestamp                  types.Int64String `json:"Timestamp"`
	ID                         types.GUID        `json:"ID"`
	BatchNumbers               json.RawMessage   `json:"BatchNumbers"` //to be implemented when needed
	Created                    *types.Date       `json:"Created"`
	Creator                    types.GUID        `json:"Creator"`
	CreatorFullName            string            `json:"CreatorFullName"`
	CustomerItemCode           string            `json:"CustomerItemCode"`
	DeliveryDate               *types.Date       `json:"DeliveryDate"`
	Description                string            `json:"Description"`
	Division                   int32             `json:"Division"`
	EntryID                    types.GUID        `json:"EntryID"`
	Item                       types.GUID        `json:"Item"`
	ItemCode                   string            `json:"ItemCode"`
	ItemDescription            string            `json:"ItemDescription"`
	LineNumber                 int32             `json:"LineNumber"`
	Modified                   *types.Date       `json:"Modified"`
	Modifier                   types.GUID        `json:"Modifier"`
	ModifierFullName           string            `json:"ModifierFullName"`
	Notes                      string            `json:"Notes"`
	QuantityDelivered          float64           `json:"QuantityDelivered"`
	QuantityOrdered            float64           `json:"QuantityOrdered"`
	SalesOrderLineID           types.GUID        `json:"SalesOrderLineID"`
	SalesOrderLineNumber       int32             `json:"SalesOrderLineNumber"`
	SalesOrderNumber           int32             `json:"SalesOrderNumber"`
	SerialNumbers              json.RawMessage   `json:"SerialNumbers"` //to be implemented when needed
	StorageLocation            types.GUID        `json:"StorageLocation"`
	StorageLocationCode        string            `json:"StorageLocationCode"`
	StorageLocationDescription string            `json:"StorageLocationDescription"`
	TrackingNumber             string            `json:"TrackingNumber"`
	Unitcode                   string            `json:"Unitcode"`
}

type SyncSalesOrderGoodDeliveryLinesCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewSyncSalesOrderGoodDeliveryLinesCall(timestamp *int64) *SyncSalesOrderGoodDeliveryLinesCall {
	selectFields := utilities.GetTaggedTagNames("json", SalesOrderGoodDeliveryLine{})
	url := service.url(fmt.Sprintf("SalesOrder/GoodDeliveryLines?$select=%s", selectFields))
	if timestamp != nil {
		url = fmt.Sprintf("%s&$filter=Timestamp gt %vL", url, *timestamp)
	}

	return &SyncSalesOrderGoodDeliveryLinesCall{
		service: service,
		urlNext: url,
	}
}

func (call *SyncSalesOrderGoodDeliveryLinesCall) Do() (*[]SalesOrderGoodDeliveryLine, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	salesOrderGoodDeliveryLines := []SalesOrderGoodDeliveryLine{}

	next, err := call.service.Get(call.urlNext, &salesOrderGoodDeliveryLines)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &salesOrderGoodDeliveryLines, nil
}
