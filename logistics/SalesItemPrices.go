package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// SalesItemPrice stores SalesItemPrice from exactonline
//
type SalesItemPrice struct {
	ID                         types.GUID  `json:"ID"`
	Account                    types.GUID  `json:"Account"`
	AccountName                string      `json:"AccountName"`
	Created                    *types.Date `json:"Created"`
	Creator                    types.GUID  `json:"Creator"`
	CreatorFullName            string      `json:"CreatorFullName"`
	Currency                   string      `json:"Currency"`
	DefaultItemUnit            string      `json:"DefaultItemUnit"`
	DefaultItemUnitDescription string      `json:"DefaultItemUnitDescription"`
	Division                   int32       `json:"Division"`
	EndDate                    *types.Date `json:"EndDate"`
	Item                       types.GUID  `json:"Item"`
	ItemCode                   string      `json:"ItemCode"`
	ItemDescription            string      `json:"ItemDescription"`
	Modified                   *types.Date `json:"Modified"`
	Modifier                   types.GUID  `json:"Modifier"`
	ModifierFullName           string      `json:"ModifierFullName"`
	NumberOfItemsPerUnit       float64     `json:"NumberOfItemsPerUnit"`
	Price                      float64     `json:"Price"`
	Quantity                   float64     `json:"Quantity"`
	StartDate                  *types.Date `json:"StartDate"`
	Unit                       string      `json:"Unit"`
	UnitDescription            string      `json:"UnitDescription"`
}

type GetSalesSalesItemPricesCall struct {
	modifiedAfter *time.Time
	urlNext       string
	client        *Client
}

func (c *Client) NewGetSalesSalesItemPricesCall(modifiedAfter *time.Time) *GetSalesSalesItemPricesCall {
	call := GetSalesSalesItemPricesCall{}
	call.modifiedAfter = modifiedAfter
	call.client = c

	selectFields := utilities.GetTaggedFieldNames("json", SalesItemPrice{})
	call.urlNext = fmt.Sprintf("%s/SalesSalesItemPrices?$select=%s", c.BaseURL(), selectFields)
	if modifiedAfter != nil {
		call.urlNext += c.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetSalesSalesItemPricesCall) Do() (*[]SalesItemPrice, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	salesItemPrices := []SalesItemPrice{}

	next, err := call.client.Get(call.urlNext, &salesItemPrices)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &salesItemPrices, nil
}

func (c *Client) GetSalesSalesItemPricesCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return c.GetCount("SalesSalesItemPrices", createdBefore)
}
