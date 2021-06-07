package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// Project stores Project from exactonline
//
type Project struct {
	ID                        types.GUID  `json:"ID"`
	Account                   types.GUID  `json:"Account"`
	AccountCode               string      `json:"AccountCode"`
	AccountContact            types.GUID  `json:"AccountContact"`
	AccountName               string      `json:"AccountName"`
	AllowAdditionalInvoicing  bool        `json:"AllowAdditionalInvoicing"`
	BlockEntry                bool        `json:"BlockEntry"`
	BlockRebilling            bool        `json:"BlockRebilling"`
	BudgetedAmount            float64     `json:"BudgetedAmount"`
	BudgetedCosts             float64     `json:"BudgetedCosts"`
	BudgetedRevenue           float64     `json:"BudgetedRevenue"`
	BudgetOverrunHours        byte        `json:"BudgetOverrunHours"`
	BudgetType                int64       `json:"BudgetType"`
	BudgetTypeDescription     string      `json:"BudgetTypeDescription"`
	Classification            types.GUID  `json:"Classification"`
	ClassificationDescription string      `json:"ClassificationDescription"`
	Code                      string      `json:"Code"`
	CostsAmountFC             float64     `json:"CostsAmountFC"`
	Created                   *types.Date `json:"Created,omitempty"`
	Creator                   types.GUID  `json:"Creator"`
	CreatorFullName           string      `json:"CreatorFullName"`
	CustomerPONumber          string      `json:"CustomerPOnumber"`
	Description               string      `json:"Description"`
	Division                  int64       `json:"Division"`
	DivisionName              string      `json:"DivisionName"`
	EndDate                   *types.Date `json:"EndDate,omitempty"`
	FixedPriceItem            types.GUID  `json:"FixedPriceItem"`
	FixedPriceItemDescription string      `json:"FixedPriceItemDescription"`
	HasWBSLines               bool        `json:"HasWBSLines"`
	InternalNotes             string      `json:"InternalNotes"`
	InvoiceAsQuoted           bool        `json:"InvoiceAsQuoted"`
	Manager                   types.GUID  `json:"Manager"`
	ManagerFullname           string      `json:"ManagerFullname"`
	MarkupPercentage          float64     `json:"MarkupPercentage"`
	Modified                  *types.Date `json:"Modified,omitempty"`
	Modifier                  types.GUID  `json:"Modifier"`
	ModifierFullName          string      `json:"ModifierFullName"`
	Notes                     string      `json:"Notes"`
	PrepaidItem               types.GUID  `json:"PrepaidItem"`
	PrepaidItemDescription    string      `json:"PrepaidItemDescription"`
	PrepaidType               int64       `json:"PrepaidType"`
	PrepaidTypeDescription    string      `json:"PrepaidTypeDescription"`
	SalesTimeQuantity         float64     `json:"SalesTimeQuantity"`
	SourceQuotation           types.GUID  `json:"SourceQuotation"`
	StartDate                 *types.Date `json:"StartDate,omitempty"`
	TimeQuantityToAlert       float64     `json:"TimeQuantityToAlert"`
	Type                      int64       `json:"Type"`
	TypeDescription           string      `json:"TypeDescription"`
	UseBillingMilestones      bool        `json:"UseBillingMilestones"`
}

type GetProjectsCall struct {
	modifiedAfter *time.Time
	urlNext       string
	service       *Service
}

func (service *Service) NewGetProjectsCall(modifiedAfter *time.Time) *GetProjectsCall {
	call := GetProjectsCall{}
	call.modifiedAfter = modifiedAfter
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", Project{})
	call.urlNext = service.url(fmt.Sprintf("Projects?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetProjectsCall) Do() (*[]Project, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	timeTransactions := []Project{}

	next, err := call.service.Get(call.urlNext, &timeTransactions)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &timeTransactions, nil
}

func (service *Service) GetProjectsCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("Projects", createdBefore)
}
