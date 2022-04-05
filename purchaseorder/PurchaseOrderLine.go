package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// PurchaseOrderLine stores PurchaseOrderLine from exactonline
//
type PurchaseOrderLine struct {
	ID                      types.Guid  `json:"ID"`
	AmountDC                float64     `json:"AmountDC"`
	AmountFC                float64     `json:"AmountFC"`
	CostCenter              string      `json:"CostCenter"`
	CostCenterDescription   string      `json:"CostCenterDescription"`
	CostUnit                string      `json:"CostUnit"`
	CostUnitDescription     string      `json:"CostUnitDescription"`
	Created                 *types.Date `json:"Created"`
	Creator                 types.Guid  `json:"Creator"`
	CreatorFullName         string      `json:"CreatorFullName"`
	Description             string      `json:"Description"`
	Discount                float64     `json:"Discount"`
	Division                int32       `json:"Division"`
	Expense                 types.Guid  `json:"Expense"`
	ExpenseDescription      string      `json:"ExpenseDescription"`
	InStock                 float64     `json:"InStock"`
	InvoicedQuantity        float64     `json:"InvoicedQuantity"`
	Item                    types.Guid  `json:"Item"`
	ItemCode                string      `json:"ItemCode"`
	ItemDescription         string      `json:"ItemDescription"`
	ItemDivisable           bool        `json:"ItemDivisable"`
	LineNumber              int32       `json:"LineNumber"`
	Modified                *types.Date `json:"Modified"`
	Modifier                types.Guid  `json:"Modifier"`
	ModifierFullName        string      `json:"ModifierFullName"`
	NetPrice                float64     `json:"NetPrice"`
	Notes                   string      `json:"Notes"`
	Project                 types.Guid  `json:"Project"`
	ProjectCode             string      `json:"ProjectCode"`
	ProjectDescription      string      `json:"ProjectDescription"`
	ProjectedStock          float64     `json:"ProjectedStock"`
	PurchaseOrderID         types.Guid  `json:"PurchaseOrderID"`
	Quantity                float64     `json:"Quantity"`
	QuantityInPurchaseUnits float64     `json:"QuantityInPurchaseUnits"`
	Rebill                  bool        `json:"Rebill"`
	ReceiptDate             *types.Date `json:"ReceiptDate"`
	ReceivedQuantity        float64     `json:"ReceivedQuantity"`
	SalesOrder              types.Guid  `json:"SalesOrder"`
	SalesOrderLine          types.Guid  `json:"SalesOrderLine"`
	SalesOrderLineNumber    int32       `json:"SalesOrderLineNumber"`
	SalesOrderNumber        int32       `json:"SalesOrderNumber"`
	SupplierItemCode        string      `json:"SupplierItemCode"`
	SupplierItemCopyRemarks byte        `json:"SupplierItemCopyRemarks"`
	Unit                    string      `json:"Unit"`
	UnitDescription         string      `json:"UnitDescription"`
	UnitPrice               float64     `json:"UnitPrice"`
	VATAmount               float64     `json:"VATAmount"`
	VATCode                 string      `json:"VATCode"`
	VATDescription          string      `json:"VATDescription"`
	VATPercentage           float64     `json:"VATPercentage"`
}

type GetPurchaseOrderLinesCall struct {
	modifiedAfter *time.Time
	urlNext       string
	service       *Service
}

func (service *Service) NewGetPurchaseOrderLinesCall(modifiedAfter *time.Time) *GetPurchaseOrderLinesCall {
	call := GetPurchaseOrderLinesCall{}
	call.modifiedAfter = modifiedAfter
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", PurchaseOrderLine{})
	call.urlNext = service.url(fmt.Sprintf("PurchaseOrderLines?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetPurchaseOrderLinesCall) Do() (*[]PurchaseOrderLine, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	purchaseOrderLines := []PurchaseOrderLine{}

	next, err := call.service.Get(call.urlNext, &purchaseOrderLines)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &purchaseOrderLines, nil
}

func (service *Service) GetPurchaseOrderLinesCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("PurchaseOrderLines", createdBefore)
}
