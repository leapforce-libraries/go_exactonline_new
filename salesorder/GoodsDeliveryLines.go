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
	ID                         types.GUID      `json:"ID"`
	BatchNumbers               json.RawMessage `json:"BatchNumbers"` //to be implemented when needed
	Created                    *types.Date     `json:"Created"`
	Creator                    types.GUID      `json:"Creator"`
	CreatorFullName            string          `json:"CreatorFullName"`
	CustomerItemCode           string          `json:"CustomerItemCode"`
	DeliveryDate               *types.Date     `json:"DeliveryDate"`
	Description                string          `json:"Description"`
	Division                   int32           `json:"Division"`
	EntryID                    types.GUID      `json:"EntryID"`
	Item                       types.GUID      `json:"Item"`
	ItemCode                   string          `json:"ItemCode"`
	ItemDescription            string          `json:"ItemDescription"`
	LineNumber                 int32           `json:"LineNumber"`
	Modified                   *types.Date     `json:"Modified"`
	Modifier                   types.GUID      `json:"Modifier"`
	ModifierFullName           string          `json:"ModifierFullName"`
	Notes                      string          `json:"Notes"`
	QuantityDelivered          float64         `json:"QuantityDelivered"`
	QuantityOrdered            float64         `json:"QuantityOrdered"`
	SalesOrderLineID           types.GUID      `json:"SalesOrderLineID"`
	SalesOrderLineNumber       int32           `json:"SalesOrderLineNumber"`
	SalesOrderNumber           int32           `json:"SalesOrderNumber"`
	SerialNumbers              json.RawMessage `json:"SerialNumbers"` //to be implemented when needed
	StorageLocation            types.GUID      `json:"StorageLocation"`
	StorageLocationCode        string          `json:"StorageLocationCode"`
	StorageLocationDescription string          `json:"StorageLocationDescription"`
	TrackingNumber             string          `json:"TrackingNumber"`
	Unitcode                   string          `json:"Unitcode"`
}

type GetGoodsDeliveryLinesCall struct {
	modifiedAfter *time.Time
	urlNext       string
	client        *Client
}

func (c *Client) NewGetGoodsDeliveryLinesCall(modifiedAfter *time.Time) *GetGoodsDeliveryLinesCall {
	call := GetGoodsDeliveryLinesCall{}
	call.modifiedAfter = modifiedAfter
	call.client = c

	selectFields := utilities.GetTaggedFieldNames("json", GoodsDeliveryLine{})
	call.urlNext = fmt.Sprintf("%s/GoodsDeliveryLines?$select=%s", c.BaseURL(), selectFields)
	if modifiedAfter != nil {
		call.urlNext += c.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetGoodsDeliveryLinesCall) Do() (*[]GoodsDeliveryLine, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	goodsDeliveryLines := []GoodsDeliveryLine{}

	next, err := call.client.Get(call.urlNext, &goodsDeliveryLines)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &goodsDeliveryLines, nil
}
func (c *Client) GetGoodsDeliveryLinesCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return c.GetCount("GoodsDeliveryLines", createdBefore)
}
