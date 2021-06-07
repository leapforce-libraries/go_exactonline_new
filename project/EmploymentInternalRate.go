package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// EmploymentInternalRate stores EmploymentInternalRate from exactonline
//
type EmploymentInternalRate struct {
	ID               types.GUID  `json:"ID"`
	Created          *types.Date `json:"Created,omitempty"`
	Creator          types.GUID  `json:"Creator"`
	CreatorFullName  string      `json:"CreatorFullName"`
	Division         int64       `json:"Division"`
	Employee         types.GUID  `json:"Employee"`
	EmployeeFullName string      `json:"EmployeeFullName"`
	EmployeeHID      int64       `json:"EmployeeHID"`
	Employment       types.GUID  `json:"Employment"`
	EmploymentHID    int64       `json:"EmploymentHID"`
	EndDate          *types.Date `json:"EndDate,omitempty"`
	InternalRate     float64     `json:"InternalRate"`
	Modified         *types.Date `json:"Modified,omitempty"`
	Modifier         types.GUID  `json:"Modifier"`
	ModifierFullName string      `json:"ModifierFullName"`
	StartDate        *types.Date `json:"StartDate,omitempty"`
}

type GetEmploymentInternalRatesCall struct {
	modifiedAfter *time.Time
	urlNext       string
	service       *Service
}

func (service *Service) NewGetEmploymentInternalRatesCall(modifiedAfter *time.Time) *GetEmploymentInternalRatesCall {
	call := GetEmploymentInternalRatesCall{}
	call.modifiedAfter = modifiedAfter
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", EmploymentInternalRate{})
	call.urlNext = service.url(fmt.Sprintf("EmploymentInternalRates?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetEmploymentInternalRatesCall) Do() (*[]EmploymentInternalRate, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	timeTransactions := []EmploymentInternalRate{}

	next, err := call.service.Get(call.urlNext, &timeTransactions)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &timeTransactions, nil
}

func (service *Service) GetEmploymentInternalRatesCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("EmploymentInternalRates", createdBefore)
}
