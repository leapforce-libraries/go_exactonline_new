package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// SalesOrderSalesOrderHeader stores SalesOrderSalesOrderHeader from exactonline
type SalesOrderSalesOrderHeader struct {
	Timestamp                      int64       `json:"Timestamp"`
	AmountDc                       float64     `json:"AmountDC"`
	AmountDiscount                 float64     `json:"AmountDiscount"`
	AmountDiscountExclVat          float64     `json:"AmountDiscountExclVat"`
	AmountFc                       float64     `json:"AmountFC"`
	AmountFcExclVat                float64     `json:"AmountFCExclVat"`
	ApprovalStatus                 int16       `json:"ApprovalStatus"`
	ApprovalStatusDescription      string      `json:"ApprovalStatusDescription"`
	Approved                       *types.Date `json:"Approved"`
	Approver                       types.Guid  `json:"Approver"`
	ApproverFullName               string      `json:"ApproverFullName"`
	Created                        *types.Date `json:"Created"`
	Creator                        types.Guid  `json:"Creator"`
	CreatorFullName                string      `json:"CreatorFullName"`
	Currency                       string      `json:"Currency"`
	DeliverTo                      types.Guid  `json:"DeliverTo"`
	DeliverToContactPerson         types.Guid  `json:"DeliverToContactPerson"`
	DeliverToContactPersonFullName string      `json:"DeliverToContactPersonFullName"`
	DeliverToName                  string      `json:"DeliverToName"`
	DeliveryAddress                types.Guid  `json:"DeliveryAddress"`
	DeliveryDate                   *types.Date `json:"DeliveryDate"`
	DeliveryStatus                 int16       `json:"DeliveryStatus"`
	DeliveryStatusDescription      string      `json:"DeliveryStatusDescription"`
	Description                    string      `json:"Description"`
	Discount                       float64     `json:"Discount"`
	Division                       int32       `json:"Division"`
	Document                       types.Guid  `json:"Document"`
	DocumentNumber                 int32       `json:"DocumentNumber"`
	DocumentSubject                string      `json:"DocumentSubject"`
	Id                             types.Guid  `json:"ID"`
	IncotermAddress                string      `json:"IncotermAddress"`
	IncotermCode                   string      `json:"IncotermCode"`
	IncotermVersion                int16       `json:"IncotermVersion"`
	InvoiceStatus                  int16       `json:"InvoiceStatus"`
	InvoiceStatusDescription       string      `json:"InvoiceStatusDescription"`
	InvoiceTo                      types.Guid  `json:"InvoiceTo"`
	InvoiceToContactPerson         types.Guid  `json:"InvoiceToContactPerson"`
	InvoiceToContactPersonFullName string      `json:"InvoiceToContactPersonFullName"`
	InvoiceToName                  string      `json:"InvoiceToName"`
	Modified                       *types.Date `json:"Modified"`
	Modifier                       types.Guid  `json:"Modifier"`
	ModifierFullName               string      `json:"ModifierFullName"`
	Notes                          string      `json:"Notes"`
	OrderDate                      *types.Date `json:"OrderDate"`
	OrderedBy                      types.Guid  `json:"OrderedBy"`
	OrderedByContactPerson         types.Guid  `json:"OrderedByContactPerson"`
	OrderedByContactPersonFullName string      `json:"OrderedByContactPersonFullName"`
	OrderedByName                  string      `json:"OrderedByName"`
	OrderId                        types.Guid  `json:"OrderID"`
	OrderNumber                    int32       `json:"OrderNumber"`
	PaymentCondition               string      `json:"PaymentCondition"`
	PaymentConditionDescription    string      `json:"PaymentConditionDescription"`
	PaymentReference               string      `json:"PaymentReference"`
	Project                        types.Guid  `json:"Project"`
	ProjectCode                    string      `json:"ProjectCode"`
	ProjectDescription             string      `json:"ProjectDescription"`
	Remarks                        string      `json:"Remarks"`
	SalesChannel                   types.Guid  `json:"SalesChannel"`
	SalesChannelCode               string      `json:"SalesChannelCode"`
	SalesChannelDescription        string      `json:"SalesChannelDescription"`
	Salesperson                    types.Guid  `json:"Salesperson"`
	SalespersonFullName            string      `json:"SalespersonFullName"`
	SelectionCode                  types.Guid  `json:"SelectionCode"`
	SelectionCodeCode              string      `json:"SelectionCodeCode"`
	SelectionCodeDescription       string      `json:"SelectionCodeDescription"`
	ShippingMethod                 types.Guid  `json:"ShippingMethod"`
	ShippingMethodCode             string      `json:"ShippingMethodCode"`
	ShippingMethodDescription      string      `json:"ShippingMethodDescription"`
	Status                         int16       `json:"Status"`
	StatusDescription              string      `json:"StatusDescription"`
	VatAmount                      float64     `json:"VATAmount"`
	VatCode                        string      `json:"VATCode"`
	VatCodeDescription             string      `json:"VATCodeDescription"`
	WarehouseCode                  string      `json:"WarehouseCode"`
	WarehouseDescription           string      `json:"WarehouseDescription"`
	WarehouseId                    types.Guid  `json:"WarehouseID"`
	YourRef                        string      `json:"YourRef"`
}

type SyncSalesOrderSalesOrderHeadersCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewSyncSalesOrderSalesOrderHeadersCall(timestamp *int64) *SyncSalesOrderSalesOrderHeadersCall {
	selectFields := utilities.GetTaggedTagNames("json", SalesOrderSalesOrderHeader{})
	url := service.url(fmt.Sprintf("SalesOrder/SalesOrderHeaders?$select=%s", selectFields))
	if timestamp != nil {
		url = fmt.Sprintf("%s&$filter=Timestamp gt %vL", url, *timestamp)
	}

	return &SyncSalesOrderSalesOrderHeadersCall{
		service: service,
		urlNext: url,
	}
}

func (call *SyncSalesOrderSalesOrderHeadersCall) Do() (*[]SalesOrderSalesOrderHeader, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	salesOrderSalesOrderHeaders := []SalesOrderSalesOrderHeader{}

	next, err := call.service.Get(call.urlNext, &salesOrderSalesOrderHeaders)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &salesOrderSalesOrderHeaders, nil
}
