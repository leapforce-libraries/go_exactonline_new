package exactonline

import (
	"encoding/json"
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// PurchaseReturn stores PurchaseReturn from exactonline
//
type PurchaseReturn struct {
	ID                      types.Guid      `json:"ID"`
	Created                 *types.Date     `json:"Created"`
	Creator                 types.Guid      `json:"Creator"`
	CreatorFullName         string          `json:"CreatorFullName"`
	Description             string          `json:"Description"`
	Division                int32           `json:"Division"`
	Document                types.Guid      `json:"Document"`
	Modified                *types.Date     `json:"Modified"`
	Modifier                types.Guid      `json:"Modifier"`
	ModifierFullName        string          `json:"ModifierFullName"`
	PurchaseReturnLines     json.RawMessage `json:"PurchaseReturnLines"`
	Remarks                 string          `json:"Remarks"`
	ReturnDate              *types.Date     `json:"ReturnDate"`
	ReturnNumber            int32           `json:"ReturnNumber"`
	Status                  int16           `json:"Status"`
	Supplier                types.Guid      `json:"Supplier"`
	SupplierAddress         types.Guid      `json:"SupplierAddress"`
	SupplierContact         types.Guid      `json:"SupplierContact"`
	SupplierContactFullName string          `json:"SupplierContactFullName"`
	TrackingNumber          string          `json:"TrackingNumber"`
	Warehouse               types.Guid      `json:"Warehouse"`
	WarehouseCode           string          `json:"WarehouseCode"`
	WarehouseDescription    string          `json:"WarehouseDescription"`
	YourRef                 string          `json:"YourRef"`
}

type GetPurchaseReturnsCall struct {
	modifiedAfter *time.Time
	urlNext       string
	service       *Service
}

func (service *Service) NewGetPurchaseReturnsCall(modifiedAfter *time.Time) *GetPurchaseReturnsCall {
	call := GetPurchaseReturnsCall{}
	call.modifiedAfter = modifiedAfter
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", PurchaseReturn{})
	call.urlNext = service.url(fmt.Sprintf("PurchaseReturns?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetPurchaseReturnsCall) Do() (*[]PurchaseReturn, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	purchaseReturns := []PurchaseReturn{}

	next, err := call.service.Get(call.urlNext, &purchaseReturns)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &purchaseReturns, nil
}

func (service *Service) GetPurchaseReturnsCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("PurchaseReturns", createdBefore)
}
