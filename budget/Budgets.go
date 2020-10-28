package exactonline

import (
	"fmt"

	types "github.com/Leapforce-nl/go_types"
	utilities "github.com/Leapforce-nl/go_utilities"
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

func (c *Client) GetBudgetsInternal(filter string) (*[]Budget, error) {
	selectFields := utilities.GetTaggedFieldNames("json", Budget{})
	urlStr := fmt.Sprintf("%s/budget/Budgets?$select=%s", c.Http().BaseURL(), selectFields)
	if filter != "" {
		urlStr += fmt.Sprintf("&$filter=%s", filter)
	}
	//fmt.Println(urlStr)

	budgets := []Budget{}

	for urlStr != "" {
		ac := []Budget{}

		next, err := c.Http().Get(urlStr, &ac)
		if err != nil {
			fmt.Println("ERROR in GetBudgetsInternal:", err)
			fmt.Println("url:", urlStr)
			return nil, err
		}

		budgets = append(budgets, ac...)

		urlStr = next
	}

	return &budgets, nil
}

func (c *Client) GetBudgets(filter string) (*[]Budget, error) {
	acc, err := c.GetBudgetsInternal(filter)
	if err != nil {
		return nil, err
	}

	return acc, nil
}
