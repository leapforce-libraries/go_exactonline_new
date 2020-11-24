package exactonline

import (
	"encoding/json"
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// GoodsDelivery stores GoodsDelivery from exactonline
//
type GoodsDelivery struct {
	EntryID                       types.GUID      `json:"EntryID"`
	Created                       *types.Date     `json:"Created"`
	Creator                       types.GUID      `json:"Creator"`
	CreatorFullName               string          `json:"CreatorFullName"`
	DeliveryAccount               types.GUID      `json:"DeliveryAccount"`
	DeliveryAccountCode           string          `json:"DeliveryAccountCode"`
	DeliveryAccountName           string          `json:"DeliveryAccountName"`
	DeliveryAddress               types.GUID      `json:"DeliveryAddress"`
	DeliveryContact               types.GUID      `json:"DeliveryContact"`
	DeliveryContactPersonFullName string          `json:"DeliveryContactPersonFullName"`
	DeliveryDate                  *types.Date     `json:"DeliveryDate"`
	DeliveryNumber                int32           `json:"DeliveryNumber"`
	Description                   string          `json:"Description"`
	Division                      int32           `json:"Division"`
	Document                      types.GUID      `json:"Document"`
	DocumentSubject               string          `json:"DocumentSubject"`
	EntryNumber                   int32           `json:"EntryNumber"`
	GoodsDeliveryLines            json.RawMessage `json:"GoodsDeliveryLines"` //to be implemented when needed
	Modified                      *types.Date     `json:"Modified"`
	Modifier                      types.GUID      `json:"Modifier"`
	ModifierFullName              string          `json:"ModifierFullName"`
	Remarks                       string          `json:"Remarks"`
	ShippingMethod                types.GUID      `json:"ShippingMethod"`
	ShippingMethodCode            string          `json:"ShippingMethodCode"`
	ShippingMethodDescription     string          `json:"ShippingMethodDescription"`
	TrackingNumber                string          `json:"TrackingNumber"`
	Warehouse                     types.GUID      `json:"Warehouse"`
	WarehouseCode                 string          `json:"WarehouseCode"`
	WarehouseDescription          string          `json:"WarehouseDescription"`
}

type GetGoodsDeliveriesCall struct {
	modifiedAfter *time.Time
	urlNext       string
	client        *Client
}

func (c *Client) NewGetGoodsDeliveriesCall(modifiedAfter *time.Time) *GetGoodsDeliveriesCall {
	call := GetGoodsDeliveriesCall{}
	call.modifiedAfter = modifiedAfter
	call.client = c

	selectFields := utilities.GetTaggedFieldNames("json", GoodsDelivery{})
	call.urlNext = fmt.Sprintf("%s/GoodsDeliveries?$select=%s", c.BaseURL(), selectFields)
	if modifiedAfter != nil {
		call.urlNext += c.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetGoodsDeliveriesCall) Do() (*[]GoodsDelivery, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	goodsDeliveries := []GoodsDelivery{}

	next, err := call.client.Get(call.urlNext, &goodsDeliveries)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &goodsDeliveries, nil
}

func (c *Client) GetGoodsDeliveriesCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return c.GetCount("GoodsDeliveries", createdBefore)
}
