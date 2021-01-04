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

type GetSalesItemPricesCall struct {
	modifiedAfter *time.Time
	urlNext       string
	client        *Client
}

func (c *Client) NewGetSalesItemPricesCall(modifiedAfter *time.Time) *GetSalesItemPricesCall {
	call := GetSalesItemPricesCall{}
	call.modifiedAfter = modifiedAfter
	call.client = c

	selectFields := utilities.GetTaggedTagNames("json", SalesItemPrice{})
	call.urlNext = fmt.Sprintf("%s/SalesItemPrices?$select=%s", c.BaseURL(), selectFields)
	if modifiedAfter != nil {
		call.urlNext += c.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetSalesItemPricesCall) Do() (*[]SalesItemPrice, *errortools.Error) {
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

func (call *GetSalesItemPricesCall) DoAll() (*[]SalesItemPrice, *errortools.Error) {
	salesItemPrices := []SalesItemPrice{}

	for true {
		_salesItemPrices, e := call.Do()
		if e != nil {
			return nil, e
		}

		if _salesItemPrices == nil {
			break
		}

		if len(*_salesItemPrices) == 0 {
			break
		}

		salesItemPrices = append(salesItemPrices, *_salesItemPrices...)
	}

	return &salesItemPrices, nil
}

func (c *Client) GetSalesItemPricesCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return c.GetCount("SalesItemPrices", createdBefore)
}
