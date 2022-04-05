package exactonline

import (
	"encoding/json"
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// SalesOrderGoodsDeliveryLine stores SalesOrderGoodsDeliveryLine from exactonline
//
type SalesOrderGoodsDeliveryLine struct {
	Timestamp                  types.Int64String `json:"Timestamp"`
	BatchNumbers               json.RawMessage   `json:"BatchNumbers"` //to be implemented when needed
	Created                    *types.Date       `json:"Created"`
	Creator                    types.Guid        `json:"Creator"`
	CreatorFullName            string            `json:"CreatorFullName"`
	CustomerItemCode           string            `json:"CustomerItemCode"`
	DeliveryDate               *types.Date       `json:"DeliveryDate"`
	Description                string            `json:"Description"`
	Division                   int32             `json:"Division"`
	EntryID                    types.Guid        `json:"EntryID"`
	ID                         types.Guid        `json:"ID"`
	Item                       types.Guid        `json:"Item"`
	ItemCode                   string            `json:"ItemCode"`
	ItemDescription            string            `json:"ItemDescription"`
	LineNumber                 int32             `json:"LineNumber"`
	Modified                   *types.Date       `json:"Modified"`
	Modifier                   types.Guid        `json:"Modifier"`
	ModifierFullName           string            `json:"ModifierFullName"`
	Notes                      string            `json:"Notes"`
	QuantityDelivered          float64           `json:"QuantityDelivered"`
	QuantityOrdered            float64           `json:"QuantityOrdered"`
	SalesOrderLineID           types.Guid        `json:"SalesOrderLineID"`
	SalesOrderLineNumber       int32             `json:"SalesOrderLineNumber"`
	SalesOrderNumber           int32             `json:"SalesOrderNumber"`
	SerialNumbers              json.RawMessage   `json:"SerialNumbers"` //to be implemented when needed
	StorageLocation            types.Guid        `json:"StorageLocation"`
	StorageLocationCode        string            `json:"StorageLocationCode"`
	StorageLocationDescription string            `json:"StorageLocationDescription"`
	TrackingNumber             string            `json:"TrackingNumber"`
	Unitcode                   string            `json:"Unitcode"`
}

type SyncSalesOrderGoodsDeliveryLinesCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewSyncSalesOrderGoodsDeliveryLinesCall(timestamp *int64) *SyncSalesOrderGoodsDeliveryLinesCall {
	selectFields := utilities.GetTaggedTagNames("json", SalesOrderGoodsDeliveryLine{})
	url := service.url(fmt.Sprintf("SalesOrder/GoodsDeliveryLines?$select=%s", selectFields))
	if timestamp != nil {
		url = fmt.Sprintf("%s&$filter=Timestamp gt %vL", url, *timestamp)
	}

	return &SyncSalesOrderGoodsDeliveryLinesCall{
		service: service,
		urlNext: url,
	}
}

func (call *SyncSalesOrderGoodsDeliveryLinesCall) Do() (*[]SalesOrderGoodsDeliveryLine, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	salesOrderGoodsDeliveryLines := []SalesOrderGoodsDeliveryLine{}

	next, err := call.service.Get(call.urlNext, &salesOrderGoodsDeliveryLines)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &salesOrderGoodsDeliveryLines, nil
}
