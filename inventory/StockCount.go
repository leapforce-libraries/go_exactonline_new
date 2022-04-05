package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// StockCount stores StockCount from exactonline
//
type StockCount struct {
	StockCountID                 types.Guid  `json:"StockCountID"`
	Created                      *types.Date `json:"Created"`
	Creator                      types.Guid  `json:"Creator"`
	CreatorFullName              string      `json:"CreatorFullName"`
	Description                  string      `json:"Description"`
	Division                     int32       `json:"Division"`
	EntryNumber                  int32       `json:"EntryNumber"`
	Modified                     *types.Date `json:"Modified"`
	Modifier                     types.Guid  `json:"Modifier"`
	ModifierFullName             string      `json:"ModifierFullName"`
	OffsetGLInventory            types.Guid  `json:"OffsetGLInventory"`
	OffsetGLInventoryCode        string      `json:"OffsetGLInventoryCode"`
	OffsetGLInventoryDescription string      `json:"OffsetGLInventoryDescription"`
	Source                       int16       `json:"Source"`
	Status                       int16       `json:"Status"`
	StockCountDate               *types.Date `json:"StockCountDate"`
	StockCountNumber             int32       `json:"StockCountNumber"`
	Warehouse                    types.Guid  `json:"Warehouse"`
	WarehouseCode                string      `json:"WarehouseCode"`
	WarehouseDescription         string      `json:"WarehouseDescription"`
}

type GetStockCountsCall struct {
	modifiedAfter *time.Time
	urlNext       string
	service       *Service
}

func (service *Service) NewGetStockCountsCall(modifiedAfter *time.Time) *GetStockCountsCall {
	call := GetStockCountsCall{}
	call.modifiedAfter = modifiedAfter
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", StockCount{})
	call.urlNext = service.url(fmt.Sprintf("StockCounts?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetStockCountsCall) Do() (*[]StockCount, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	stockCounts := []StockCount{}

	next, err := call.service.Get(call.urlNext, &stockCounts)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &stockCounts, nil
}

func (service *Service) GetStockCountsCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("StockCounts", createdBefore)
}
