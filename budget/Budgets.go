package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// Budget stores Budget from exactonline
//
type Budget struct {
	ID                        types.GUID  `json:"ID"`
	AmountDC                  float64     `json:"AmountDC"`
	BudgetScenario            types.GUID  `json:"BudgetScenario"`
	BudgetScenarioCode        string      `json:"BudgetScenarioCode"`
	BudgetScenarioDescription string      `json:"BudgetScenarioDescription"`
	Costcenter                string      `json:"Costcenter"`
	CostcenterDescription     string      `json:"CostcenterDescription"`
	Costunit                  string      `json:"Costunit"`
	CostunitDescription       string      `json:"CostunitDescription"`
	Created                   *types.Date `json:"Created"`
	Creator                   types.GUID  `json:"Creator"`
	CreatorFullName           string      `json:"CreatorFullName"`
	Division                  int32       `json:"Division"`
	GLAccount                 types.GUID  `json:"GLAccount"`
	GLAccountCode             string      `json:"GLAccountCode"`
	GLAccountDescription      string      `json:"GLAccountDescription"`
	HID                       string      `json:"HID"`
	Item                      types.GUID  `json:"Item"`
	ItemCode                  string      `json:"ItemCode"`
	ItemDescription           string      `json:"ItemDescription"`
	Modified                  *types.Date `json:"Modified"`
	Modifier                  types.GUID  `json:"Modifier"`
	ModifierFullName          string      `json:"ModifierFullName"`
	ReportingPeriod           int16       `json:"ReportingPeriod"`
	ReportingYear             int16       `json:"ReportingYear"`
}

type GetBudgetsCall struct {
	modifiedAfter *time.Time
	urlNext       string
	client        *Client
}

func (c *Client) NewGetBudgetsCall(modifiedAfter *time.Time) *GetBudgetsCall {
	call := GetBudgetsCall{}
	call.modifiedAfter = modifiedAfter
	call.client = c

	selectFields := utilities.GetTaggedTagNames("json", Budget{})
	call.urlNext = fmt.Sprintf("%s/Budgets?$select=%s", c.BaseURL(), selectFields)
	if modifiedAfter != nil {
		call.urlNext += c.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetBudgetsCall) Do() (*[]Budget, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	budgets := []Budget{}

	next, err := call.client.Get(call.urlNext, &budgets)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &budgets, nil
}

func (c *Client) GetBudgetsCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return c.GetCount("Budgets", createdBefore)
}
