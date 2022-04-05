package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// SalesInvoiceLine stores SalesInvoiceLine from exactonline
//
type SalesInvoiceLine struct {
	ID                      types.Guid  `json:"ID"`
	AmountDC                float64     `json:"AmountDC"`
	AmountFC                float64     `json:"AmountFC"`
	CostCenter              string      `json:"CostCenter"`
	CostCenterDescription   string      `json:"CostCenterDescription"`
	CostUnit                string      `json:"CostUnit"`
	CostUnitDescription     string      `json:"CostUnitDescription"`
	CustomerItemCode        string      `json:"CustomerItemCode"`
	DeliveryDate            *types.Date `json:"DeliveryDate"`
	Description             string      `json:"Description"`
	Discount                float64     `json:"Discount"`
	Division                int32       `json:"Division"`
	Employee                types.Guid  `json:"Employee"`
	EmployeeFullName        string      `json:"EmployeeFullName"`
	EndTime                 *types.Date `json:"EndTime"`
	ExtraDutyAmountFC       float64     `json:"ExtraDutyAmountFC"`
	ExtraDutyPercentage     float64     `json:"ExtraDutyPercentage"`
	GLAccount               types.Guid  `json:"GLAccount"`
	GLAccountDescription    string      `json:"GLAccountDescription"`
	InvoiceID               types.Guid  `json:"InvoiceID"`
	Item                    types.Guid  `json:"Item"`
	ItemCode                string      `json:"ItemCode"`
	ItemDescription         string      `json:"ItemDescription"`
	LineNumber              int32       `json:"LineNumber"`
	NetPrice                float64     `json:"NetPrice"`
	Notes                   string      `json:"Notes"`
	Pricelist               types.Guid  `json:"Pricelist"`
	PricelistDescription    string      `json:"PricelistDescription"`
	Project                 types.Guid  `json:"Project"`
	ProjectDescription      string      `json:"ProjectDescription"`
	ProjectWBS              types.Guid  `json:"ProjectWBS"`
	ProjectWBSDescription   string      `json:"ProjectWBSDescription"`
	Quantity                float64     `json:"Quantity"`
	SalesOrder              types.Guid  `json:"SalesOrder"`
	SalesOrderLine          types.Guid  `json:"SalesOrderLine"`
	SalesOrderLineNumber    int32       `json:"SalesOrderLineNumber"`
	SalesOrderNumber        int32       `json:"SalesOrderNumber"`
	StartTime               *types.Date `json:"StartTime"`
	Subscription            types.Guid  `json:"Subscription"`
	SubscriptionDescription string      `json:"SubscriptionDescription"`
	TaxSchedule             types.Guid  `json:"TaxSchedule"`
	TaxScheduleCode         string      `json:"TaxScheduleCode"`
	TaxScheduleDescription  string      `json:"TaxScheduleDescription"`
	UnitCode                string      `json:"UnitCode"`
	UnitDescription         string      `json:"UnitDescription"`
	UnitPrice               float64     `json:"UnitPrice"`
	VATAmountDC             float64     `json:"VATAmountDC"`
	VATAmountFC             float64     `json:"VATAmountFC"`
	VATCode                 string      `json:"VATCode"`
	VATCodeDescription      string      `json:"VATCodeDescription"`
	VATPercentage           float64     `json:"VATPercentage"`
}

type GetSalesInvoiceLinesCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewGetSalesInvoiceLinesCall() *GetSalesInvoiceLinesCall {
	call := GetSalesInvoiceLinesCall{}
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", SalesInvoiceLine{})
	call.urlNext = service.url(fmt.Sprintf("SalesInvoiceLines?$select=%s", selectFields))

	return &call
}

func (call *GetSalesInvoiceLinesCall) Do() (*[]SalesInvoiceLine, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	salesOrderLines := []SalesInvoiceLine{}

	next, err := call.service.Get(call.urlNext, &salesOrderLines)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &salesOrderLines, nil
}
