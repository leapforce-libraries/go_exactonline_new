package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// SalesOrderLine stores SalesOrderLine from exactonline
//
type SalesOrderLine struct {
	ID                      types.Guid  `json:"ID"`
	AmountDC                float64     `json:"AmountDC"`
	AmountFC                float64     `json:"AmountFC"`
	CostCenter              string      `json:"CostCenter"`
	CostCenterDescription   string      `json:"CostCenterDescription"`
	CostPriceFC             float64     `json:"CostPriceFC"`
	CostUnit                string      `json:"CostUnit"`
	CostUnitDescription     string      `json:"CostUnitDescription"`
	CustomerItemCode        string      `json:"CustomerItemCode"`
	DeliveryDate            *types.Date `json:"DeliveryDate"`
	Description             string      `json:"Description"`
	Discount                float64     `json:"Discount"`
	Division                int32       `json:"Division"`
	Item                    types.Guid  `json:"Item"`
	ItemCode                string      `json:"ItemCode"`
	ItemDescription         string      `json:"ItemDescription"`
	ItemVersion             types.Guid  `json:"ItemVersion"`
	ItemVersionDescription  string      `json:"ItemVersionDescription"`
	LineNumber              int32       `json:"LineNumber"`
	NetPrice                float64     `json:"NetPrice"`
	Notes                   string      `json:"Notes"`
	OrderID                 types.Guid  `json:"OrderID"`
	OrderNumber             int32       `json:"OrderNumber"`
	Pricelist               types.Guid  `json:"Pricelist"`
	PricelistDescription    string      `json:"PricelistDescription"`
	Project                 types.Guid  `json:"Project"`
	ProjectDescription      string      `json:"ProjectDescription"`
	PurchaseOrder           types.Guid  `json:"PurchaseOrder"`
	PurchaseOrderLine       types.Guid  `json:"PurchaseOrderLine"`
	PurchaseOrderLineNumber int32       `json:"PurchaseOrderLineNumber"`
	PurchaseOrderNumber     int32       `json:"PurchaseOrderNumber"`
	Quantity                float64     `json:"Quantity"`
	ShopOrder               types.Guid  `json:"ShopOrder"`
	UnitCode                string      `json:"UnitCode"`
	UnitDescription         string      `json:"UnitDescription"`
	UnitPrice               float64     `json:"UnitPrice"`
	UseDropShipment         byte        `json:"UseDropShipment"`
	VATAmount               float64     `json:"VATAmount"`
	VATCode                 string      `json:"VATCode"`
	VATCodeDescription      string      `json:"VATCodeDescription"`
	VATPercentage           float64     `json:"VATPercentage"`
}

type GetSalesOrderLinesCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewGetSalesOrderLinesCall() *GetSalesOrderLinesCall {
	call := GetSalesOrderLinesCall{}
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", SalesOrderLine{})
	call.urlNext = service.url(fmt.Sprintf("SalesOrderLines?$select=%s", selectFields))

	return &call
}

func (call *GetSalesOrderLinesCall) Do() (*[]SalesOrderLine, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	salesOrderLines := []SalesOrderLine{}

	next, err := call.service.Get(call.urlNext, &salesOrderLines)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &salesOrderLines, nil
}
