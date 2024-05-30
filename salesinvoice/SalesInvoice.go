package exactonline

import (
	"fmt"
	"strings"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// SalesInvoice stores SalesInvoice from exactonline
type SalesInvoice struct {
	InvoiceID                      types.Guid  `json:"InvoiceID"`
	AmountDC                       float64     `json:"AmountDC"`
	AmountDiscount                 float64     `json:"AmountDiscount"`
	AmountDiscountExclVat          float64     `json:"AmountDiscountExclVat"`
	AmountFC                       float64     `json:"AmountFC"`
	AmountFCExclVat                float64     `json:"AmountFCExclVat"`
	Created                        *types.Date `json:"Created"`
	Creator                        types.Guid  `json:"Creator"`
	CreatorFullName                string      `json:"CreatorFullName"`
	Currency                       string      `json:"Currency"`
	DeliverTo                      types.Guid  `json:"DeliverTo"`
	DeliverToAddress               types.Guid  `json:"DeliverToAddress"`
	DeliverToContactPerson         types.Guid  `json:"DeliverToContactPerson"`
	DeliverToContactPersonFullName string      `json:"DeliverToContactPersonFullName"`
	DeliverToName                  string      `json:"DeliverToName"`
	Description                    string      `json:"Description"`
	Discount                       float64     `json:"Discount"`
	DiscountType                   int16       `json:"DiscountType"`
	Division                       int32       `json:"Division"`
	Document                       types.Guid  `json:"Document"`
	DocumentNumber                 int32       `json:"DocumentNumber"`
	DocumentSubject                string      `json:"DocumentSubject"`
	DueDate                        *types.Date `json:"DueDate"`
	ExtraDutyAmountFC              float64     `json:"ExtraDutyAmountFC"`
	GAccountAmountFC               float64     `json:"GAccountAmountFC"`
	IncotermAddress                string      `json:"IncotermAddress"`
	IncotermCode                   string      `json:"IncotermCode"`
	IncotermVersion                int16       `json:"IncotermVersion"`
	InvoiceDate                    *types.Date `json:"InvoiceDate"`
	InvoiceNumber                  int32       `json:"InvoiceNumber"`
	InvoiceTo                      types.Guid  `json:"InvoiceTo"`
	InvoiceToContactPerson         types.Guid  `json:"InvoiceToContactPerson"`
	InvoiceToContactPersonFullName string      `json:"InvoiceToContactPersonFullName"`
	InvoiceToName                  string      `json:"InvoiceToName"`
	IsExtraDuty                    bool        `json:"IsExtraDuty"`
	Journal                        string      `json:"Journal"`
	JournalDescription             string      `json:"JournalDescription"`
	Modified                       *types.Date `json:"Modified"`
	Modifier                       types.Guid  `json:"Modifier"`
	ModifierFullName               string      `json:"ModifierFullName"`
	OrderDate                      *types.Date `json:"OrderDate"`
	OrderedBy                      types.Guid  `json:"OrderedBy"`
	OrderedByContactPerson         types.Guid  `json:"OrderedByContactPerson"`
	OrderedByContactPersonFullName string      `json:"OrderedByContactPersonFullName"`
	OrderedByName                  string      `json:"OrderedByName"`
	OrderNumber                    int32       `json:"OrderNumber"`
	PaymentCondition               string      `json:"PaymentCondition"`
	PaymentConditionDescription    string      `json:"PaymentConditionDescription"`
	PaymentReference               string      `json:"PaymentReference"`
	Remarks                        string      `json:"Remarks"`
	SalesChannel                   types.Guid  `json:"SalesChannel"`
	SalesChannelCode               string      `json:"SalesChannelCode"`
	SalesChannelDescription        string      `json:"SalesChannelDescription"`
	//SalesInvoiceLines              []SalesInvoiceLine `json:"SalesInvoiceLines"`
	// SalesInvoiceOrderChargeLines  `json:"SalesInvoiceOrderChargeLines"`
	Salesperson                          types.Guid `json:"Salesperson"`
	SalespersonFullName                  string     `json:"SalespersonFullName"`
	SelectionCode                        types.Guid `json:"SelectionCode"`
	SelectionCodeCode                    string     `json:"SelectionCodeCode"`
	SelectionCodeDescription             string     `json:"SelectionCodeDescription"`
	ShippingMethod                       types.Guid `json:"ShippingMethod"`
	ShippingMethodCode                   string     `json:"ShippingMethodCode"`
	ShippingMethodDescription            string     `json:"ShippingMethodDescription"`
	StarterSalesInvoiceStatus            int16      `json:"StarterSalesInvoiceStatus"`
	StarterSalesInvoiceStatusDescription string     `json:"StarterSalesInvoiceStatusDescription"`
	Status                               int16      `json:"Status"`
	StatusDescription                    string     `json:"StatusDescription"`
	TaxSchedule                          types.Guid `json:"TaxSchedule"`
	TaxScheduleCode                      string     `json:"TaxScheduleCode"`
	TaxScheduleDescription               string     `json:"TaxScheduleDescription"`
	Type                                 int32      `json:"Type"`
	TypeDescription                      string     `json:"TypeDescription"`
	VATAmountDC                          float64    `json:"VATAmountDC"`
	VATAmountFC                          float64    `json:"VATAmountFC"`
	Warehouse                            types.Guid `json:"Warehouse"`
	WithholdingTaxAmountFC               float64    `json:"WithholdingTaxAmountFC"`
	WithholdingTaxBaseAmount             float64    `json:"WithholdingTaxBaseAmount"`
	WithholdingTaxPercentage             float64    `json:"WithholdingTaxPercentage"`
	YourRef                              string     `json:"YourRef"`
}

type GetSalesInvoicesCall struct {
	urlNext string
	service *Service
}

type GetSalesInvoicesCallParams struct {
	Fields        *[]string
	DeliverTo     *types.Guid
	ModifiedAfter *time.Time
}

func (service *Service) NewGetSalesInvoicesCall(params *GetSalesInvoicesCallParams) *GetSalesInvoicesCall {
	call := GetSalesInvoicesCall{}
	call.service = service

	selectFields := ""
	if params.Fields != nil {
		selectFields = strings.Join(*params.Fields, ",")
	} else {
		selectFields = utilities.GetTaggedTagNames("json", SalesInvoice{})
	}

	call.urlNext = service.url(fmt.Sprintf("SalesInvoices?$select=%s", selectFields))
	filter := []string{}

	if params != nil {
		if params.DeliverTo != nil {
			filter = append(filter, fmt.Sprintf("DeliverTo eq guid'%s'", params.DeliverTo.String()))
		}
		if params.ModifiedAfter != nil {
			filter = append(filter, service.DateFilter("Modified", "gt", params.ModifiedAfter, false, ""))
		}
	}

	if len(filter) > 0 {
		call.urlNext = fmt.Sprintf("%s&$filter=%s", call.urlNext, strings.Join(filter, " and "))
	}

	return &call
}

func (call *GetSalesInvoicesCall) Do() (*[]SalesInvoice, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	salesInvoices := []SalesInvoice{}

	next, err := call.service.Get(call.urlNext, &salesInvoices)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &salesInvoices, nil
}

func (call *GetSalesInvoicesCall) DoAll() (*[]SalesInvoice, *errortools.Error) {
	salesInvoices := []SalesInvoice{}

	for true {
		_salesInvoices, e := call.Do()
		if e != nil {
			return nil, e
		}

		if _salesInvoices == nil {
			break
		}

		if len(*_salesInvoices) == 0 {
			break
		}

		salesInvoices = append(salesInvoices, *_salesInvoices...)
	}

	return &salesInvoices, nil
}
