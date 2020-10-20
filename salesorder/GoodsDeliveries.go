package exactonline

import (
	"encoding/json"
	"fmt"

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

func (c *Client) GetGoodsDeliveriesInternal(filter string) (*[]GoodsDelivery, error) {
	selectFields := utilities.GetTaggedFieldNames("json", GoodsDelivery{})
	urlStr := fmt.Sprintf("%s/salesorder/GoodsDeliveries?$select=%s", c.Http().BaseURL(), selectFields)
	if filter != "" {
		urlStr += fmt.Sprintf("&$filter=%s", filter)
	}
	//fmt.Println(urlStr)

	goodsDeliveries := []GoodsDelivery{}

	for urlStr != "" {
		ac := []GoodsDelivery{}

		next, err := c.Http().Get(urlStr, &ac)
		if err != nil {
			fmt.Println("ERROR in GetGoodsDeliveriesInternal:", err)
			fmt.Println("url:", urlStr)
			return nil, err
		}

		goodsDeliveries = append(goodsDeliveries, ac...)

		urlStr = next
	}

	return &goodsDeliveries, nil
}

func (c *Client) GetGoodsDeliveries() (*[]GoodsDelivery, error) {
	acc, err := c.GetGoodsDeliveriesInternal("")
	if err != nil {
		return nil, err
	}

	return acc, nil
}
