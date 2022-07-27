package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// PurchaseOrderPurchaseOrder stores PurchaseOrderPurchaseOrder from exactonline
//
type PurchaseOrderPurchaseOrder struct {
	Timestamp                     types.Int64String `json:"Timestamp"`
	AmountDC                      float64           `json:"AmountDC"`
	AmountFC                      float64           `json:"AmountFC"`
	CostCenter                    types.Guid        `json:"CostCenter"`
	CostCenterCode                string            `json:"CostCenterCode"`
	CostCenterDescription         string            `json:"CostCenterDescription"`
	CostUnit                      types.Guid        `json:"CostUnit"`
	CostUnitCode                  string            `json:"CostUnitCode"`
	CostUnitDescription           string            `json:"CostUnitDescription"`
	Created                       *types.Date       `json:"Created"`
	Creator                       types.Guid        `json:"Creator"`
	CreatorFullName               string            `json:"CreatorFullName"`
	Currency                      string            `json:"Currency"`
	DeliveryAccount               types.Guid        `json:"DeliveryAccount"`
	DeliveryAccountCode           string            `json:"DeliveryAccountCode"`
	DeliveryAccountName           string            `json:"DeliveryAccountName"`
	DeliveryAddress               types.Guid        `json:"DeliveryAddress"`
	DeliveryContact               types.Guid        `json:"DeliveryContact"`
	DeliveryContactPersonFullName string            `json:"DeliveryContactPersonFullName"`
	Description                   string            `json:"Description"`
	Discount                      float64           `json:"Discount"`
	Division                      int32             `json:"Division"`
	Document                      types.Guid        `json:"Document"`
	DocumentNumber                int32             `json:"DocumentNumber"`
	DocumentSubject               string            `json:"DocumentSubject"`
	DropShipment                  bool              `json:"DropShipment"`
	ExchangeRate                  float64           `json:"ExchangeRate"`
	Expense                       types.Guid        `json:"Expense"`
	ExpenseDescription            string            `json:"ExpenseDescription"`
	ID                            types.Guid        `json:"ID"`
	IncotermAddress               string            `json:"IncotermAddress"`
	IncotermCode                  string            `json:"IncotermCode"`
	IncotermVersion               int16             `json:"IncotermVersion"`
	InvoicedQuantity              float64           `json:"InvoicedQuantity"`
	InvoiceStatus                 int32             `json:"InvoiceStatus"`
	IsBatchNumberItem             byte              `json:"IsBatchNumberItem"`
	IsSerialNumberItem            byte              `json:"IsSerialNumberItem"`
	Item                          types.Guid        `json:"Item"`
	ItemBarcode                   string            `json:"ItemBarcode"`
	ItemCode                      string            `json:"ItemCode"`
	ItemDescription               string            `json:"ItemDescription"`
	ItemDivisable                 bool              `json:"ItemDivisable"`
	LineNumber                    int32             `json:"LineNumber"`
	Modified                      *types.Date       `json:"Modified"`
	Modifier                      types.Guid        `json:"Modifier"`
	ModifierFullName              string            `json:"ModifierFullName"`
	NetPrice                      float64           `json:"NetPrice"`
	Notes                         string            `json:"Notes"`
	OrderDate                     *types.Date       `json:"OrderDate"`
	OrderNumber                   int32             `json:"OrderNumber"`
	OrderStatus                   int32             `json:"OrderStatus"`
	PaymentCondition              string            `json:"PaymentCondition"`
	PaymentConditionDescription   string            `json:"PaymentConditionDescription"`
	Project                       types.Guid        `json:"Project"`
	ProjectCode                   string            `json:"ProjectCode"`
	ProjectDescription            string            `json:"ProjectDescription"`
	PurchaseAgent                 types.Guid        `json:"PurchaseAgent"`
	PurchaseAgentFullName         string            `json:"PurchaseAgentFullName"`
	PurchaseOrderID               types.Guid        `json:"PurchaseOrderID"`
	Quantity                      float64           `json:"Quantity"`
	QuantityInPurchaseUnits       float64           `json:"QuantityInPurchaseUnits"`
	Rebill                        bool              `json:"Rebill"`
	ReceiptDate                   *types.Date       `json:"ReceiptDate"`
	ReceiptStatus                 int32             `json:"ReceiptStatus"`
	ReceivedQuantity              float64           `json:"ReceivedQuantity"`
	Remarks                       string            `json:"Remarks"`
	SalesOrder                    types.Guid        `json:"SalesOrder"`
	SalesOrderLine                types.Guid        `json:"SalesOrderLine"`
	SalesOrderLineNumber          int32             `json:"SalesOrderLineNumber"`
	SalesOrderNumber              int32             `json:"SalesOrderNumber"`
	SelectionCode                 types.Guid        `json:"SelectionCode"`
	SelectionCodeCode             string            `json:"SelectionCodeCode"`
	SelectionCodeDescription      string            `json:"SelectionCodeDescription"`
	SendingMethod                 int32             `json:"SendingMethod"`
	ShippingMethod                types.Guid        `json:"ShippingMethod"`
	ShippingMethodCode            string            `json:"ShippingMethodCode"`
	ShippingMethodDescription     string            `json:"ShippingMethodDescription"`
	Source                        int16             `json:"Source"`
	Supplier                      types.Guid        `json:"Supplier"`
	SupplierCode                  string            `json:"SupplierCode"`
	SupplierContact               types.Guid        `json:"SupplierContact"`
	SupplierContactPersonFullName string            `json:"SupplierContactPersonFullName"`
	SupplierItemCode              string            `json:"SupplierItemCode"`
	SupplierItemCopyRemarks       byte              `json:"SupplierItemCopyRemarks"`
	SupplierName                  string            `json:"SupplierName"`
	Unit                          string            `json:"Unit"`
	UnitDescription               string            `json:"UnitDescription"`
	UnitPrice                     float64           `json:"UnitPrice"`
	VATAmount                     float64           `json:"VATAmount"`
	VATCode                       string            `json:"VATCode"`
	VATDescription                string            `json:"VATDescription"`
	VATPercentage                 float64           `json:"VATPercentage"`
	Warehouse                     types.Guid        `json:"Warehouse"`
	WarehouseCode                 string            `json:"WarehouseCode"`
	WarehouseDescription          string            `json:"WarehouseDescription"`
	YourRef                       string            `json:"YourRef"`
}

type SyncPurchaseOrderPurchaseOrdersCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewSyncPurchaseOrderPurchaseOrdersCall(timestamp *int64) *SyncPurchaseOrderPurchaseOrdersCall {
	selectFields := utilities.GetTaggedTagNames("json", PurchaseOrderPurchaseOrder{})
	url := service.url(fmt.Sprintf("PurchaseOrder/PurchaseOrders?$select=%s", selectFields))
	if timestamp != nil {
		url = fmt.Sprintf("%s&$filter=Timestamp gt %vL", url, *timestamp)
	}

	return &SyncPurchaseOrderPurchaseOrdersCall{
		service: service,
		urlNext: url,
	}
}

func (call *SyncPurchaseOrderPurchaseOrdersCall) Do() (*[]PurchaseOrderPurchaseOrder, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	purchaseOrderPurchaseOrders := []PurchaseOrderPurchaseOrder{}

	next, err := call.service.Get(call.urlNext, &purchaseOrderPurchaseOrders)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &purchaseOrderPurchaseOrders, nil
}
