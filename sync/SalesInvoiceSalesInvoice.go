package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// SalesInvoiceSalesInvoice stores SalesInvoiceSalesInvoice from exactonline
//
type SalesInvoiceSalesInvoice struct {
	Timestamp                            types.Int64String `json:"Timestamp"`
	ID                                   types.GUID        `json:"ID"`
	AmountDC                             float64           `json:"AmountDC"`
	AmountDiscount                       float64           `json:"AmountDiscount"`
	AmountDiscountExclVat                float64           `json:"AmountDiscountExclVat"`
	AmountFC                             float64           `json:"AmountFC"`
	AmountFCExclVat                      float64           `json:"AmountFCExclVat"`
	CostCenter                           types.GUID        `json:"CostCenter"`
	CostCenterDescription                string            `json:"CostCenterDescription"`
	CostUnit                             string            `json:"CostUnit"`
	CostUnitDescription                  string            `json:"CostUnitDescription"`
	Created                              *types.Date       `json:"Created"`
	Creator                              types.GUID        `json:"Creator"`
	CreatorFullName                      string            `json:"CreatorFullName"`
	Currency                             string            `json:"Currency"`
	CustomerItemCode                     string            `json:"CustomerItemCode"`
	DeliverTo                            types.GUID        `json:"DeliverTo"`
	DeliverToAddress                     types.GUID        `json:"DeliverToAddress"`
	DeliverToContactPerson               types.GUID        `json:"DeliverToContactPerson"`
	DeliverToContactPersonFullName       string            `json:"DeliverToContactPersonFullName"`
	DeliveryDate                         string            `json:"DeliveryDate"`
	Description                          string            `json:"Description"`
	Discount                             float64           `json:"Discount"`
	DiscountType                         int16             `json:"DiscountType"`
	Division                             int32             `json:"Division"`
	Document                             types.GUID        `json:"Document"`
	DocumentNumber                       int32             `json:"DocumentNumber"`
	DocumentSubject                      string            `json:"DocumentSubject"`
	DueDate                              *types.Date       `json:"DueDate"`
	Employee                             types.GUID        `json:"Employee"`
	EmployeeFullName                     string            `json:"EmployeeFullName"`
	EndTime                              *types.Date       `json:"EndTime"`
	ExtraDutyAmountFC                    float64           `json:"ExtraDutyAmountFC"`
	ExtraDutyPercentage                  float64           `json:"ExtraDutyPercentage"`
	GAccountAmountFC                     float64           `json:"GAccountAmountFC"`
	GLAccount                            types.GUID        `json:"GLAccount"`
	GLAccountDescription                 string            `json:"GLAccountDescription"`
	InvoiceDate                          *types.Date       `json:"InvoiceDate"`
	InvoiceID                            types.GUID        `json:"InvoiceID"`
	InvoiceNumber                        int32             `json:"InvoiceNumber"`
	InvoiceTo                            types.GUID        `json:"InvoiceTo"`
	InvoiceToContactPerson               types.GUID        `json:"InvoiceToContactPerson"`
	InvoiceToContactPersonFullName       string            `json:"InvoiceToContactPersonFullName"`
	InvoiceToName                        string            `json:"InvoiceToName"`
	IsExtraDuty                          bool              `json:"IsExtraDuty"`
	Item                                 types.GUID        `json:"Item"`
	ItemCode                             string            `json:"ItemCode"`
	ItemDescription                      string            `json:"ItemDescription"`
	Journal                              string            `json:"Journal"`
	JournalDescription                   string            `json:"JournalDescription"`
	LineNumber                           int32             `json:"LineNumber"`
	Modified                             *types.Date       `json:"Modified"`
	Modifier                             types.GUID        `json:"Modifier"`
	ModifierFullName                     string            `json:"ModifierFullName"`
	NetPrice                             float64           `json:"NetPrice"`
	OrderDate                            *types.Date       `json:"OrderDate"`
	OrderedBy                            types.GUID        `json:"OrderedBy"`
	OrderedByContactPerson               types.GUID        `json:"OrderedByContactPerson"`
	OrderedByContactPersonFullName       string            `json:"OrderedByContactPersonFullName"`
	OrderedByName                        string            `json:"OrderedByName"`
	OrderNumber                          int32             `json:"OrderNumber"`
	PaymentCondition                     string            `json:"PaymentCondition"`
	PaymentConditionDescription          string            `json:"PaymentConditionDescription"`
	PaymentReference                     string            `json:"PaymentReference"`
	Pricelist                            types.GUID        `json:"Pricelist"`
	PricelistDescription                 string            `json:"PricelistDescription"`
	Project                              types.GUID        `json:"Project"`
	ProjectDescription                   string            `json:"ProjectDescription"`
	ProjectWBS                           types.GUID        `json:"ProjectWBS"`
	ProjectWBSDescription                string            `json:"ProjectWBSDescription"`
	Quantity                             float64           `json:"Quantity"`
	Remarks                              string            `json:"Remarks"`
	SalesOrder                           types.GUID        `json:"SalesOrder"`
	SalesOrderLine                       types.GUID        `json:"SalesOrderLine"`
	SalesOrderLineNumber                 int32             `json:"SalesOrderLineNumber"`
	SalesOrderNumber                     int32             `json:"SalesOrderNumber"`
	SalesPerson                          types.GUID        `json:"Salesperson"`
	SalesPersonFullName                  string            `json:"SalespersonFullName"`
	StarterSalesInvoiceStatus            int16             `json:"StarterSalesInvoiceStatus"`
	StarterSalesInvoiceStatusDescription string            `json:"StarterSalesInvoiceStatusDescription"`
	Status                               int16             `json:"Status"`
	StatusDescription                    string            `json:"StatusDescription"`
	Subscription                         types.GUID        `json:"Subscription"`
	SubscriptionDescription              string            `json:"SubscriptionDescription"`
	TaxSchedule                          types.GUID        `json:"TaxSchedule"`
	TaxScheduleCode                      string            `json:"TaxScheduleCode"`
	TaxScheduleDescription               string            `json:"TaxScheduleDescription"`
	Type                                 int32             `json:"Type"`
	TypeDescription                      string            `json:"TypeDescription"`
	UnitCode                             string            `json:"UnitCode"`
	UnitDescription                      string            `json:"UnitDescription"`
	UnitPrice                            float64           `json:"UnitPrice"`
	VATAmountDC                          float64           `json:"VATAmountDC"`
	VATAmountFC                          float64           `json:"VATAmountFC"`
	VATCode                              string            `json:"VATCode"`
	VATCodeDescription                   string            `json:"VATCodeDescription"`
	VATPercentage                        float64           `json:"VATPercentage"`
	Warehouse                            types.GUID        `json:"Warehouse"`
	WithholdingTaxAmountFC               float64           `json:"WithholdingTaxAmountFC"`
	WithholdingTaxBaseAmount             float64           `json:"WithholdingTaxBaseAmount"`
	WithholdingTaxPercentage             float64           `json:"WithholdingTaxPercentage"`
	YourRef                              string            `json:"YourRef"`
}

type SyncSalesInvoiceSalesInvoicesCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewSyncSalesInvoiceSalesInvoicesCall(timestamp *int64) *SyncSalesInvoiceSalesInvoicesCall {
	selectFields := utilities.GetTaggedTagNames("json", SalesInvoiceSalesInvoice{})
	url := service.url(fmt.Sprintf("SalesInvoice/SalesInvoices?$select=%s", selectFields))
	if timestamp != nil {
		url = fmt.Sprintf("%s&$filter=Timestamp gt %vL", url, *timestamp)
	}

	return &SyncSalesInvoiceSalesInvoicesCall{
		service: service,
		urlNext: url,
	}
}

func (call *SyncSalesInvoiceSalesInvoicesCall) Do() (*[]SalesInvoiceSalesInvoice, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	salesInvoiceSalesInvoices := []SalesInvoiceSalesInvoice{}

	next, err := call.service.Get(call.urlNext, &salesInvoiceSalesInvoices)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &salesInvoiceSalesInvoices, nil
}
