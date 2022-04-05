package exactonline

import (
	"encoding/json"
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// PurchaseOrder stores PurchaseOrder from exactonline
//
type PurchaseOrder struct {
	PurchaseOrderID               types.Guid      `json:"PurchaseOrderID"`
	AmountDC                      float64         `json:"AmountDC"`
	AmountFC                      float64         `json:"AmountFC"`
	Created                       *types.Date     `json:"Created"`
	Creator                       types.Guid      `json:"Creator"`
	CreatorFullName               string          `json:"CreatorFullName"`
	Currency                      string          `json:"Currency"`
	DeliveryAccount               types.Guid      `json:"DeliveryAccount"`
	DeliveryAccountCode           string          `json:"DeliveryAccountCode"`
	DeliveryAccountName           string          `json:"DeliveryAccountName"`
	DeliveryAddress               types.Guid      `json:"DeliveryAddress"`
	DeliveryContact               types.Guid      `json:"DeliveryContact"`
	DeliveryContactPersonFullName string          `json:"DeliveryContactPersonFullName"`
	Description                   string          `json:"Description"`
	Division                      int32           `json:"Division"`
	Document                      types.Guid      `json:"Document"`
	DocumentSubject               string          `json:"DocumentSubject"`
	DropShipment                  bool            `json:"DropShipment"`
	ExchangeRate                  float64         `json:"ExchangeRate"`
	InvoiceStatus                 int32           `json:"InvoiceStatus"`
	Modified                      *types.Date     `json:"Modified"`
	Modifier                      types.Guid      `json:"Modifier"`
	ModifierFullName              string          `json:"ModifierFullName"`
	OrderDate                     *types.Date     `json:"OrderDate"`
	OrderNumber                   int32           `json:"OrderNumber"`
	OrderStatus                   int32           `json:"OrderStatus"`
	PaymentCondition              string          `json:"PaymentCondition"`
	PaymentConditionDescription   string          `json:"PaymentConditionDescription"`
	PurchaseAgent                 types.Guid      `json:"PurchaseAgent"`
	PurchaseAgentFullName         string          `json:"PurchaseAgentFullName"`
	PurchaseOrderLines            json.RawMessage `json:"PurchaseOrderLines"`
	ReceiptDate                   *types.Date     `json:"ReceiptDate"`
	ReceiptStatus                 int32           `json:"ReceiptStatus"`
	Remarks                       string          `json:"Remarks"`
	SalesOrder                    types.Guid      `json:"SalesOrder"`
	SalesOrderNumber              int32           `json:"SalesOrderNumber"`
	SelectionCode                 types.Guid      `json:"SelectionCode"`
	SelectionCodeCode             string          `json:"SelectionCodeCode"`
	SelectionCodeDescription      string          `json:"SelectionCodeDescription"`
	ShippingMethod                types.Guid      `json:"ShippingMethod"`
	ShippingMethodDescription     string          `json:"ShippingMethodDescription"`
	Source                        int16           `json:"Source"`
	Supplier                      types.Guid      `json:"Supplier"`
	SupplierCode                  string          `json:"SupplierCode"`
	SupplierContact               types.Guid      `json:"SupplierContact"`
	SupplierContactPersonFullName string          `json:"SupplierContactPersonFullName"`
	SupplierName                  string          `json:"SupplierName"`
	VATAmount                     float64         `json:"VATAmount"`
	Warehouse                     types.Guid      `json:"Warehouse"`
	WarehouseCode                 string          `json:"WarehouseCode"`
	WarehouseDescription          string          `json:"WarehouseDescription"`
	YourRef                       string          `json:"YourRef"`
}

type GetPurchaseOrdersCall struct {
	modifiedAfter *time.Time
	urlNext       string
	service       *Service
}

func (service *Service) NewGetPurchaseOrdersCall(modifiedAfter *time.Time) *GetPurchaseOrdersCall {
	call := GetPurchaseOrdersCall{}
	call.modifiedAfter = modifiedAfter
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", PurchaseOrder{})
	call.urlNext = service.url(fmt.Sprintf("PurchaseOrders?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetPurchaseOrdersCall) Do() (*[]PurchaseOrder, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	purchaseOrders := []PurchaseOrder{}

	next, err := call.service.Get(call.urlNext, &purchaseOrders)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &purchaseOrders, nil
}

func (service *Service) GetPurchaseOrdersCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("PurchaseOrders", createdBefore)
}
