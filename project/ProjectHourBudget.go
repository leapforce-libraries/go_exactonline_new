package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// ProjectHourBudget stores ProjectHourBudget from exactonline
//
type ProjectHourBudget struct {
	ID                 types.Guid  `json:"ID"`
	Budget             float64     `json:"Budget"`
	Created            *types.Date `json:"Created,omitempty"`
	Creator            types.Guid  `json:"Creator"`
	CreatorFullName    string      `json:"CreatorFullName"`
	Division           int64       `json:"Division"`
	Item               types.Guid  `json:"Item"`
	ItemCode           string      `json:"ItemCode"`
	ItemDescription    string      `json:"ItemDescription"`
	Modified           *types.Date `json:"Modified,omitempty"`
	Modifier           types.Guid  `json:"Modifier"`
	ModifierFullName   string      `json:"ModifierFullName"`
	Project            types.Guid  `json:"Project"`
	ProjectCode        string      `json:"ProjectCode"`
	ProjectDescription string      `json:"ProjectDescription"`
}

type GetProjectHourBudgetsCall struct {
	modifiedAfter *time.Time
	urlNext       string
	service       *Service
}

func (service *Service) NewGetProjectHourBudgetsCall(modifiedAfter *time.Time) *GetProjectHourBudgetsCall {
	call := GetProjectHourBudgetsCall{}
	call.modifiedAfter = modifiedAfter
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", ProjectHourBudget{})
	call.urlNext = service.url(fmt.Sprintf("ProjectHourBudgets?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetProjectHourBudgetsCall) Do() (*[]ProjectHourBudget, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	projectHourBudgets := []ProjectHourBudget{}

	next, err := call.service.Get(call.urlNext, &projectHourBudgets)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &projectHourBudgets, nil
}

func (service *Service) GetProjectHourBudgetsCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("ProjectHourBudgets", createdBefore)
}
