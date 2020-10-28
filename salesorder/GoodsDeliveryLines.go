package exactonline

import (
	"encoding/json"
	"fmt"

	types "github.com/Leapforce-nl/go_types"
	utilities "github.com/Leapforce-nl/go_utilities"
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

func (c *Client) GetGoodsDeliveryLinesInternal(filter string) (*[]GoodsDeliveryLine, error) {
	selectFields := utilities.GetTaggedFieldNames("json", GoodsDeliveryLine{})
	urlStr := fmt.Sprintf("%s/salesorder/GoodsDeliveryLines?$select=%s", c.Http().BaseURL(), selectFields)
	if filter != "" {
		urlStr += fmt.Sprintf("&$filter=%s", filter)
	}
	//fmt.Println(urlStr)

	goodsDeliveryLines := []GoodsDeliveryLine{}

	for urlStr != "" {
		ac := []GoodsDeliveryLine{}

		next, err := c.Http().Get(urlStr, &ac)
		if err != nil {
			fmt.Println("ERROR in GetGoodsDeliveryLinesInternal:", err)
			fmt.Println("url:", urlStr)
			return nil, err
		}

		goodsDeliveryLines = append(goodsDeliveryLines, ac...)

		urlStr = next
	}

	return &goodsDeliveryLines, nil
}

func (c *Client) GetGoodsDeliveryLines(filter string) (*[]GoodsDeliveryLine, error) {
	acc, err := c.GetGoodsDeliveryLinesInternal(filter)
	if err != nil {
		return nil, err
	}

	return acc, nil
}
