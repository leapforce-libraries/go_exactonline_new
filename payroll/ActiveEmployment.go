package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// ActiveEmployment stores ActiveEmployment from exactonline
//
type ActiveEmployment struct {
	ID                       types.GUID  `json:"ID"`
	AverageDaysPerWeek       float64     `json:"AverageDaysPerWeek"`
	AverageHoursPerWeek      float64     `json:"AverageHoursPerWeek"`
	Contract                 types.GUID  `json:"Contract"`
	ContractDocument         types.GUID  `json:"ContractDocument"`
	ContractEndDate          *types.Date `json:"ContractEndDate,omitempty"`
	ContractProbationEndDate *types.Date `json:"ContractProbationEndDate,omitempty"`
	ContractProbationPeriod  int64       `json:"ContractProbationPeriod"`
	ContractStartDate        *types.Date `json:"ContractStartDate,omitempty"`
	ContractType             int64       `json:"ContractType"`
	ContractTypeDescription  string      `json:"ContractTypeDescription"`
	Created                  *types.Date `json:"Created,omitempty"`
	Creator                  types.GUID  `json:"Creator"`
	CreatorFullName          string      `json:"CreatorFullName"`
	Department               types.GUID  `json:"Department"`
	DepartmentCode           string      `json:"DepartmentCode"`
	DepartmentDescription    string      `json:"DepartmentDescription"`
	Division                 int64       `json:"Division"`
	Employee                 types.GUID  `json:"Employee"`
	EmployeeFullName         string      `json:"EmployeeFullName"`
	EmployeeHID              int64       `json:"EmployeeHID"`
	EmploymentOrganization   types.GUID  `json:"EmploymentOrganization"`
	EndDate                  *types.Date `json:"EndDate,omitempty"`
	HID                      int64       `json:"HID"`
	HourlyWage               float64     `json:"HourlyWage"`
	InternalRate             float64     `json:"InternalRate"`
	Jobtitle                 types.GUID  `json:"Jobtitle"`
	JobtitleDescription      string      `json:"JobtitleDescription"`
	Modified                 *types.Date `json:"Modified,omitempty"`
	Modifier                 types.GUID  `json:"Modifier"`
	ModifierFullName         string      `json:"ModifierFullName"`
	ReasonEnd                int64       `json:"ReasonEnd"`
	ReasonEndDescription     string      `json:"ReasonEndDescription"`
	ReasonEndFlex            int64       `json:"ReasonEndFlex"`
	ReasonEndFlexDescription string      `json:"ReasonEndFlexDescription"`
	Salary                   types.GUID  `json:"Salary"`
	Schedule                 types.GUID  `json:"Schedule"`
	ScheduleAverageHours     float64     `json:"ScheduleAverageHours"`
	ScheduleCode             string      `json:"ScheduleCode"`
	ScheduleDays             float64     `json:"ScheduleDays"`
	ScheduleDescription      string      `json:"ScheduleDescription"`
	ScheduleHours            float64     `json:"ScheduleHours"`
	StartDate                *types.Date `json:"StartDate,omitempty"`
	StartDateOrganization    *types.Date `json:"StartDateOrganization,omitempty"`
}

type GetActiveEmploymentsCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewGetActiveEmploymentsCall(modifiedAfter *time.Time) *GetActiveEmploymentsCall {
	call := GetActiveEmploymentsCall{}
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", ActiveEmployment{})
	call.urlNext = service.url(fmt.Sprintf("ActiveEmployments?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetActiveEmploymentsCall) Do() (*[]ActiveEmployment, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	activeEmployments := []ActiveEmployment{}

	next, err := call.service.Get(call.urlNext, &activeEmployments)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &activeEmployments, nil
}

func (service *Service) GetActiveEmploymentsCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("ActiveEmployments", createdBefore)
}
