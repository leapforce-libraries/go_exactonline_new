package exactonline

import (
	"encoding/json"
	"fmt"
	"time"

	types "github.com/Leapforce-nl/go_types"
	utilities "github.com/Leapforce-nl/go_utilities"
)

// GoodsDelivery stores GoodsDelivery from exactonline
//
type GoodsDelivery struct {
	EntryID                       types.GUID      `json:"EntryID"`
	Created                       *types.Date     `json:"Created"`
	Creator                       types.GUID      `json:"Creator"`
	CreatorFullName               string          `json:"CreatorFullName"`
	DeliveryGoodsDelivery         types.GUID      `json:"DeliveryGoodsDelivery"`
	DeliveryGoodsDeliveryCode     string          `json:"DeliveryGoodsDeliveryCode"`
	DeliveryGoodsDeliveryName     string          `json:"DeliveryGoodsDeliveryName"`
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
	call.urlNext = fmt.Sprintf("%s/salesorder/GoodsDeliveries?$select=%s", c.Http().BaseURL(), selectFields)
	if modifiedAfter != nil {
		call.urlNext += c.Http().DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetGoodsDeliveriesCall) Do() (*[]GoodsDelivery, error) {
	if call.urlNext == "" {
		return nil, nil
	}

	goodsDeliveries := []GoodsDelivery{}

	next, err := call.client.Http().Get(call.urlNext, &goodsDeliveries)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &goodsDeliveries, nil
}

func (c *Client) GetGoodsDeliveriesCount(createdBefore *time.Time) (int64, error) {
	return c.Http().GetCount("salesorder/GoodsDeliveries", createdBefore)
}
