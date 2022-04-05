package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// FinancialGLAccount stores FinancialGLAccount from exactonline
//
type FinancialGLAccount struct {
	Timestamp                      types.Int64String `json:"Timestamp"`
	AssimilatedVATBox              int16             `json:"AssimilatedVATBox"`
	BalanceSide                    string            `json:"BalanceSide"`
	BalanceType                    string            `json:"BalanceType"`
	BelcotaxType                   int32             `json:"BelcotaxType"`
	Code                           string            `json:"Code"`
	Compress                       bool              `json:"Compress"`
	Costcenter                     string            `json:"Costcenter"`
	CostcenterDescription          string            `json:"CostcenterDescription"`
	Costunit                       string            `json:"Costunit"`
	CostunitDescription            string            `json:"CostunitDescription"`
	Created                        *types.Date       `json:"Created"`
	Creator                        types.Guid        `json:"Creator"`
	CreatorFullName                string            `json:"CreatorFullName"`
	Description                    string            `json:"Description"`
	Division                       int32             `json:"Division"`
	ExcludeVATListing              byte              `json:"ExcludeVATListing"`
	ExpenseNonDeductiblePercentage float64           `json:"ExpenseNonDeductiblePercentage"`
	ID                             types.Guid        `json:"ID"`
	IsBlocked                      bool              `json:"IsBlocked"`
	Matching                       bool              `json:"Matching"`
	Modified                       *types.Date       `json:"Modified"`
	Modifier                       types.Guid        `json:"Modifier"`
	ModifierFullName               string            `json:"ModifierFullName"`
	PrivateGLAccount               types.Guid        `json:"PrivateGLAccount"`
	PrivatePercentage              float64           `json:"PrivatePercentage"`
	ReportingCode                  string            `json:"ReportingCode"`
	RevalueCurrency                bool              `json:"RevalueCurrency"`
	SearchCode                     string            `json:"SearchCode"`
	Type                           int32             `json:"Type"`
	TypeDescription                string            `json:"TypeDescription"`
	UseCostcenter                  byte              `json:"UseCostcenter"`
	UseCostunit                    byte              `json:"UseCostunit"`
	VATCode                        string            `json:"VATCode"`
	VATDescription                 string            `json:"VATDescription"`
	VATGLAccountType               string            `json:"VATGLAccountType"`
	VATNonDeductibleGLAccount      types.Guid        `json:"VATNonDeductibleGLAccount"`
	VATNonDeductiblePercentage     float64           `json:"VATNonDeductiblePercentage"`
	VATSystem                      string            `json:"VATSystem"`
	YearEndCostGLAccount           types.Guid        `json:"YearEndCostGLAccount"`
	YearEndReflectionGLAccount     types.Guid        `json:"YearEndReflectionGLAccount"`
}

type SyncFinancialGLAccountsCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewSyncFinancialGLAccountsCall(timestamp *int64) *SyncFinancialGLAccountsCall {
	selectFields := utilities.GetTaggedTagNames("json", FinancialGLAccount{})
	url := service.url(fmt.Sprintf("Financial/GLAccounts?$select=%s", selectFields))
	if timestamp != nil {
		url = fmt.Sprintf("%s&$filter=Timestamp gt %vL", url, *timestamp)
	}

	return &SyncFinancialGLAccountsCall{
		service: service,
		urlNext: url,
	}
}

func (call *SyncFinancialGLAccountsCall) Do() (*[]FinancialGLAccount, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	financialGLAccounts := []FinancialGLAccount{}

	next, err := call.service.Get(call.urlNext, &financialGLAccounts)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &financialGLAccounts, nil
}
