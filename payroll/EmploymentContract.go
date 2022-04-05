package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// EmploymentContract stores EmploymentContract from exactonline
//
type EmploymentContract struct {
	ID                           types.Guid  `json:"ID"`
	ContractFlexPhase            int32       `json:"ContractFlexPhase"`
	ContractFlexPhaseDescription string      `json:"ContractFlexPhaseDescription"`
	Created                      *types.Date `json:"Created"`
	Creator                      types.Guid  `json:"Creator"`
	CreatorFullName              string      `json:"CreatorFullName"`
	Division                     int32       `json:"Division"`
	Document                     types.Guid  `json:"Document"`
	Employee                     types.Guid  `json:"Employee"`
	EmployeeFullName             string      `json:"EmployeeFullName"`
	EmployeeHID                  int32       `json:"EmployeeHID"`
	EmployeeType                 int32       `json:"EmployeeType"`
	EmployeeTypeDescription      string      `json:"EmployeeTypeDescription"`
	Employment                   types.Guid  `json:"Employment"`
	EmploymentHID                int32       `json:"EmploymentHID"`
	EndDate                      *types.Date `json:"EndDate"`
	Modified                     *types.Date `json:"Modified"`
	Modifier                     types.Guid  `json:"Modifier"`
	ModifierFullName             string      `json:"ModifierFullName"`
	Notes                        string      `json:"Notes"`
	ProbationEndDate             *types.Date `json:"ProbationEndDate"`
	ProbationPeriod              int32       `json:"ProbationPeriod"`
	ReasonContract               int32       `json:"ReasonContract"`
	ReasonContractDescription    string      `json:"ReasonContractDescription"`
	Sequence                     int32       `json:"Sequence"`
	StartDate                    *types.Date `json:"StartDate"`
	Type                         int32       `json:"Type"`
	TypeDescription              string      `json:"TypeDescription"`
}

type GetEmploymentContractsCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewGetEmploymentContractsCall(modifiedAfter *time.Time) *GetEmploymentContractsCall {
	call := GetEmploymentContractsCall{}
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", EmploymentContract{})
	call.urlNext = service.url(fmt.Sprintf("EmploymentContracts?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetEmploymentContractsCall) Do() (*[]EmploymentContract, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	employmentContracts := []EmploymentContract{}

	next, err := call.service.Get(call.urlNext, &employmentContracts)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &employmentContracts, nil
}

func (service *Service) GetEmploymentContractsCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("EmploymentContracts", createdBefore)
}
