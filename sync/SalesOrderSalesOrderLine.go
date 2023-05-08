package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// SalesOrderSalesOrderLine stores SalesOrderSalesOrderLine from exactonline
type SalesOrderSalesOrderLine struct {
	Timestamp                 int64       `json:"Timestamp"`
	AmountDc                  float64     `json:"AmountDC"`
	AmountFc                  float64     `json:"AmountFC"`
	CostCenter                string      `json:"CostCenter"`
	CostCenterDescription     string      `json:"CostCenterDescription"`
	CostPriceFc               float64     `json:"CostPriceFC"`
	CostUnit                  string      `json:"CostUnit"`
	CostUnitDescription       string      `json:"CostUnitDescription"`
	Created                   *types.Date `json:"Created"`
	Creator                   types.Guid  `json:"Creator"`
	CreatorFullName           string      `json:"CreatorFullName"`
	CustomerItemCode          string      `json:"CustomerItemCode"`
	DeliveryDate              *types.Date `json:"DeliveryDate"`
	DeliveryStatus            int16       `json:"DeliveryStatus"`
	DeliveryStatusDescription string      `json:"DeliveryStatusDescription"`
	Description               string      `json:"Description"`
	Discount                  float64     `json:"Discount"`
	Division                  int32       `json:"Division"`
	Id                        types.Guid  `json:"ID"`
	InvoiceStatus             int16       `json:"InvoiceStatus"`
	InvoiceStatusDescription  string      `json:"InvoiceStatusDescription"`
	Item                      types.Guid  `json:"Item"`
	ItemCode                  string      `json:"ItemCode"`
	ItemDescription           string      `json:"ItemDescription"`
	ItemVersion               types.Guid  `json:"ItemVersion"`
	ItemVersionDescription    string      `json:"ItemVersionDescription"`
	LineNumber                int32       `json:"LineNumber"`
	Margin                    float64     `json:"Margin"`
	Modified                  *types.Date `json:"Modified"`
	Modifier                  types.Guid  `json:"Modifier"`
	ModifierFullName          string      `json:"ModifierFullName"`
	NetPrice                  float64     `json:"NetPrice"`
	Notes                     string      `json:"Notes"`
	OrderId                   types.Guid  `json:"OrderID"`
	OrderNumber               int32       `json:"OrderNumber"`
	Pricelist                 types.Guid  `json:"Pricelist"`
	PricelistDescription      string      `json:"PricelistDescription"`
	Project                   types.Guid  `json:"Project"`
	ProjectCode               string      `json:"ProjectCode"`
	ProjectDescription        string      `json:"ProjectDescription"`
	PurchaseOrder             types.Guid  `json:"PurchaseOrder"`
	PurchaseOrderLine         types.Guid  `json:"PurchaseOrderLine"`
	PurchaseOrderLineNumber   int32       `json:"PurchaseOrderLineNumber"`
	PurchaseOrderNumber       int32       `json:"PurchaseOrderNumber"`
	Quantity                  float64     `json:"Quantity"`
	QuantityDelivered         float64     `json:"QuantityDelivered"`
	QuantityInvoiced          float64     `json:"QuantityInvoiced"`
	ShopOrder                 types.Guid  `json:"ShopOrder"`
	ShopOrderNumber           int32       `json:"ShopOrderNumber"`
	Status                    int16       `json:"Status"`
	StatusDescription         string      `json:"StatusDescription"`
	UnitCode                  string      `json:"UnitCode"`
	UnitDescription           string      `json:"UnitDescription"`
	UnitPrice                 float64     `json:"UnitPrice"`
	UseDropShipment           byte        `json:"UseDropShipment"`
	VatAmount                 float64     `json:"VATAmount"`
	VatCode                   string      `json:"VATCode"`
	VatCodeDescription        string      `json:"VATCodeDescription"`
	VatPercentage             float64     `json:"VATPercentage"`
}

type SyncSalesOrderSalesOrderLinesCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewSyncSalesOrderSalesOrderLinesCall(timestamp *int64) *SyncSalesOrderSalesOrderLinesCall {
	selectFields := utilities.GetTaggedTagNames("json", SalesOrderSalesOrderLine{})
	url := service.url(fmt.Sprintf("SalesOrder/SalesOrderLines?$select=%s", selectFields))
	if timestamp != nil {
		url = fmt.Sprintf("%s&$filter=Timestamp gt %vL", url, *timestamp)
	}

	return &SyncSalesOrderSalesOrderLinesCall{
		service: service,
		urlNext: url,
	}
}

func (call *SyncSalesOrderSalesOrderLinesCall) Do() (*[]SalesOrderSalesOrderLine, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	salesOrderSalesOrderLines := []SalesOrderSalesOrderLine{}

	next, err := call.service.Get(call.urlNext, &salesOrderSalesOrderLines)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &salesOrderSalesOrderLines, nil
}
