package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// SalesOrderSalesOrder stores SalesOrderSalesOrder from exactonline
//
type SalesOrderSalesOrder struct {
	Timestamp                      types.Int64String `json:"Timestamp"`
	AmountDC                       float64           `json:"AmountDC"`
	AmountDiscount                 float64           `json:"AmountDiscount"`
	AmountDiscountExclVat          float64           `json:"AmountDiscountExclVat"`
	AmountFC                       float64           `json:"AmountFC"`
	AmountFCExclVat                float64           `json:"AmountFCExclVat"`
	ApprovalStatus                 int16             `json:"ApprovalStatus"`
	ApprovalStatusDescription      string            `json:"ApprovalStatusDescription"`
	Approved                       *types.Date       `json:"Approved"`
	Approver                       types.GUID        `json:"Approver"`
	ApproverFullName               string            `json:"ApproverFullName"`
	CostCenter                     string            `json:"CostCenter"`
	CostCenterDescription          string            `json:"CostCenterDescription"`
	CostPriceFC                    float64           `json:"CostPriceFC"`
	CostUnit                       string            `json:"CostUnit"`
	CostUnitDescription            string            `json:"CostUnitDescription"`
	Created                        *types.Date       `json:"Created"`
	Creator                        types.GUID        `json:"Creator"`
	CreatorFullName                string            `json:"CreatorFullName"`
	Currency                       string            `json:"Currency"`
	CustomerItemCode               string            `json:"CustomerItemCode"`
	DeliverTo                      types.GUID        `json:"DeliverTo"`
	DeliverToContactPerson         types.GUID        `json:"DeliverToContactPerson"`
	DeliverToContactPersonFullName string            `json:"DeliverToContactPersonFullName"`
	DeliverToName                  string            `json:"DeliverToName"`
	DeliveryAddress                types.GUID        `json:"DeliveryAddress"`
	DeliveryDate                   *types.Date       `json:"DeliveryDate"`
	DeliveryStatus                 int16             `json:"DeliveryStatus"`
	DeliveryStatusDescription      string            `json:"DeliveryStatusDescription"`
	Description                    string            `json:"Description"`
	Discount                       float64           `json:"Discount"`
	Division                       int32             `json:"Division"`
	Document                       types.GUID        `json:"Document"`
	DocumentNumber                 int32             `json:"DocumentNumber"`
	DocumentSubject                string            `json:"DocumentSubject"`
	ID                             types.GUID        `json:"ID"`
	InvoiceStatus                  int16             `json:"InvoiceStatus"`
	InvoiceStatusDescription       string            `json:"InvoiceStatusDescription"`
	InvoiceTo                      types.GUID        `json:"InvoiceTo"`
	InvoiceToContactPerson         types.GUID        `json:"InvoiceToContactPerson"`
	InvoiceToContactPersonFullName string            `json:"InvoiceToContactPersonFullName"`
	InvoiceToName                  string            `json:"InvoiceToName"`
	Item                           types.GUID        `json:"Item"`
	ItemCode                       string            `json:"ItemCode"`
	ItemDescription                string            `json:"ItemDescription"`
	ItemVersion                    types.GUID        `json:"ItemVersion"`
	ItemVersionDescription         string            `json:"ItemVersionDescription"`
	LineNumber                     int32             `json:"LineNumber"`
	Margin                         float64           `json:"Margin"`
	Modified                       *types.Date       `json:"Modified"`
	Modifier                       types.GUID        `json:"Modifier"`
	ModifierFullName               string            `json:"ModifierFullName"`
	NetPrice                       float64           `json:"NetPrice"`
	Notes                          string            `json:"Notes"`
	OrderDate                      *types.Date       `json:"OrderDate"`
	OrderedBy                      types.GUID        `json:"OrderedBy"`
	OrderedByContactPerson         types.GUID        `json:"OrderedByContactPerson"`
	OrderedByContactPersonFullName string            `json:"OrderedByContactPersonFullName"`
	OrderedByName                  string            `json:"OrderedByName"`
	OrderID                        types.GUID        `json:"OrderID"`
	OrderNumber                    int32             `json:"OrderNumber"`
	PaymentCondition               string            `json:"PaymentCondition"`
	PaymentConditionDescription    string            `json:"PaymentConditionDescription"`
	PaymentReference               string            `json:"PaymentReference"`
	Pricelist                      types.GUID        `json:"Pricelist"`
	PricelistDescription           string            `json:"PricelistDescription"`
	Project                        types.GUID        `json:"Project"`
	ProjectDescription             string            `json:"ProjectDescription"`
	PurchaseOrder                  types.GUID        `json:"PurchaseOrder"`
	PurchaseOrderLine              types.GUID        `json:"PurchaseOrderLine"`
	PurchaseOrderLineNumber        int32             `json:"PurchaseOrderLineNumber"`
	PurchaseOrderNumber            int32             `json:"PurchaseOrderNumber"`
	Quantity                       float64           `json:"Quantity"`
	QuantityDelivered              float64           `json:"QuantityDelivered"`
	QuantityInvoiced               float64           `json:"QuantityInvoiced"`
	Remarks                        string            `json:"Remarks"`
	SalesPerson                    types.GUID        `json:"Salesperson"`
	SalesPersonFullName            string            `json:"SalespersonFullName"`
	SelectionCode                  types.GUID        `json:"SelectionCode"`
	SelectionCodeCode              string            `json:"SelectionCodeCode"`
	SelectionCodeDescription       string            `json:"SelectionCodeDescription"`
	ShippingMethod                 types.GUID        `json:"ShippingMethod"`
	ShippingMethodDescription      string            `json:"ShippingMethodDescription"`
	ShopOrder                      types.GUID        `json:"ShopOrder"`
	Status                         int16             `json:"Status"`
	StatusDescription              string            `json:"StatusDescription"`
	TaxSchedule                    types.GUID        `json:"TaxSchedule"`
	TaxScheduleCode                string            `json:"TaxScheduleCode"`
	TaxScheduleDescription         string            `json:"TaxScheduleDescription"`
	UnitCode                       string            `json:"UnitCode"`
	UnitDescription                string            `json:"UnitDescription"`
	UnitPrice                      float64           `json:"UnitPrice"`
	UseDropShipment                byte              `json:"UseDropShipment"`
	VATAmount                      float64           `json:"VATAmount"`
	VATCode                        string            `json:"VATCode"`
	VATCodeDescription             string            `json:"VATCodeDescription"`
	VATPercentage                  float64           `json:"VATPercentage"`
	WarehouseCode                  string            `json:"WarehouseCode"`
	WarehouseDescription           string            `json:"WarehouseDescription"`
	WarehouseID                    types.GUID        `json:"WarehouseID"`
	YourRef                        string            `json:"YourRef"`
}

type SyncSalesOrderSalesOrdersCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewSyncSalesOrderSalesOrdersCall(timestamp *int64) *SyncSalesOrderSalesOrdersCall {
	selectFields := utilities.GetTaggedTagNames("json", SalesOrderSalesOrder{})
	url := service.url(fmt.Sprintf("SalesOrder/SalesOrders?$select=%s", selectFields))
	if timestamp != nil {
		url = fmt.Sprintf("%s&$filter=Timestamp gt %vL", url, *timestamp)
	}

	return &SyncSalesOrderSalesOrdersCall{
		service: service,
		urlNext: url,
	}
}

func (call *SyncSalesOrderSalesOrdersCall) Do() (*[]SalesOrderSalesOrder, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	salesOrderSalesOrders := []SalesOrderSalesOrder{}

	next, err := call.service.Get(call.urlNext, &salesOrderSalesOrders)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &salesOrderSalesOrders, nil
}
