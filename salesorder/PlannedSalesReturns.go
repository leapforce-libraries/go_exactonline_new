package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// PlannedSalesReturn stores PlannedSalesReturn from exactonline
//
type PlannedSalesReturn struct {
	PlannedSalesReturnID             types.GUID               `json:"PlannedSalesReturnID"`
	Created                          *types.Date              `json:"Created"`
	Creator                          types.GUID               `json:"Creator"`
	CreatorFullName                  string                   `json:"CreatorFullName"`
	DeliveredTo                      types.GUID               `json:"DeliveredTo"`
	DeliveredToContactPerson         types.GUID               `json:"DeliveredToContactPerson"`
	DeliveredToContactPersonFullName string                   `json:"DeliveredToContactPersonFullName"`
	DeliveredToName                  string                   `json:"DeliveredToName"`
	DeliveryAddress                  types.GUID               `json:"DeliveryAddress"`
	Description                      string                   `json:"Description"`
	Division                         int32                    `json:"Division"`
	Document                         types.GUID               `json:"Document"`
	DocumentSubject                  string                   `json:"DocumentSubject"`
	Modified                         *types.Date              `json:"Modified"`
	Modifier                         types.GUID               `json:"Modifier"`
	ModifierFullName                 string                   `json:"ModifierFullName"`
	PlannedSalesReturnLines          []PlannedSalesReturnLine `json:"PlannedSalesReturnLines"`
	Remarks                          string                   `json:"Remarks"`
	ReturnDate                       *types.Date              `json:"ReturnDate"`
	ReturnNumber                     int32                    `json:"ReturnNumber"`
	Source                           int16                    `json:"Source"`
	Status                           int16                    `json:"Status"`
	Warehouse                        types.GUID               `json:"Warehouse"`
	WarehouseCode                    string                   `json:"WarehouseCode"`
	WarehouseDescription             string                   `json:"WarehouseDescription"`
}

type GetPlannedSalesReturnsCall struct {
	modifiedAfter *time.Time
	urlNext       string
	service       *Service
}

func (service *Service) NewGetPlannedSalesReturnsCall(modifiedAfter *time.Time) *GetPlannedSalesReturnsCall {
	call := GetPlannedSalesReturnsCall{}
	call.modifiedAfter = modifiedAfter
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", PlannedSalesReturn{})
	call.urlNext = service.url(fmt.Sprintf("PlannedSalesReturns?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetPlannedSalesReturnsCall) Do() (*[]PlannedSalesReturn, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	plannedSalesReturns := []PlannedSalesReturn{}

	next, err := call.service.Get(call.urlNext, &plannedSalesReturns)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &plannedSalesReturns, nil
}
func (service *Service) GetPlannedSalesReturnsCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("PlannedSalesReturns", createdBefore)
}
