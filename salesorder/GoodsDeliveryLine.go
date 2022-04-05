package exactonline

import (
	"encoding/json"
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// GoodsDeliveryLine stores GoodsDeliveryLine from exactonline
//
type GoodsDeliveryLine struct {
	ID                         types.Guid      `json:"ID"`
	BatchNumbers               json.RawMessage `json:"BatchNumbers"` //to be implemented when needed
	Created                    *types.Date     `json:"Created"`
	Creator                    types.Guid      `json:"Creator"`
	CreatorFullName            string          `json:"CreatorFullName"`
	CustomerItemCode           string          `json:"CustomerItemCode"`
	DeliveryDate               *types.Date     `json:"DeliveryDate"`
	Description                string          `json:"Description"`
	Division                   int32           `json:"Division"`
	EntryID                    types.Guid      `json:"EntryID"`
	Item                       types.Guid      `json:"Item"`
	ItemCode                   string          `json:"ItemCode"`
	ItemDescription            string          `json:"ItemDescription"`
	LineNumber                 int32           `json:"LineNumber"`
	Modified                   *types.Date     `json:"Modified"`
	Modifier                   types.Guid      `json:"Modifier"`
	ModifierFullName           string          `json:"ModifierFullName"`
	Notes                      string          `json:"Notes"`
	QuantityDelivered          float64         `json:"QuantityDelivered"`
	QuantityOrdered            float64         `json:"QuantityOrdered"`
	SalesOrderLineID           types.Guid      `json:"SalesOrderLineID"`
	SalesOrderLineNumber       int32           `json:"SalesOrderLineNumber"`
	SalesOrderNumber           int32           `json:"SalesOrderNumber"`
	SerialNumbers              json.RawMessage `json:"SerialNumbers"` //to be implemented when needed
	StorageLocation            types.Guid      `json:"StorageLocation"`
	StorageLocationCode        string          `json:"StorageLocationCode"`
	StorageLocationDescription string          `json:"StorageLocationDescription"`
	TrackingNumber             string          `json:"TrackingNumber"`
	Unitcode                   string          `json:"Unitcode"`
}

type GetGoodsDeliveryLinesCall struct {
	modifiedAfter *time.Time
	urlNext       string
	service       *Service
}

func (service *Service) NewGetGoodsDeliveryLinesCall(modifiedAfter *time.Time) *GetGoodsDeliveryLinesCall {
	call := GetGoodsDeliveryLinesCall{}
	call.modifiedAfter = modifiedAfter
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", GoodsDeliveryLine{})
	call.urlNext = service.url(fmt.Sprintf("GoodsDeliveryLines?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetGoodsDeliveryLinesCall) Do() (*[]GoodsDeliveryLine, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	goodsDeliveryLines := []GoodsDeliveryLine{}

	next, err := call.service.Get(call.urlNext, &goodsDeliveryLines)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &goodsDeliveryLines, nil
}
func (service *Service) GetGoodsDeliveryLinesCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("GoodsDeliveryLines", createdBefore)
}
