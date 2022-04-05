package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// StockBatchNumber stores StockBatchNumber from exactonline
//
type StockBatchNumber struct {
	ID                         types.Guid  `json:"ID"`
	BatchNumber                string      `json:"BatchNumber"`
	BatchNumberID              types.Guid  `json:"BatchNumberID"`
	Created                    *types.Date `json:"Created"`
	Creator                    types.Guid  `json:"Creator"`
	CreatorFullName            string      `json:"CreatorFullName"`
	Division                   int32       `json:"Division"`
	DraftStockTransactionID    types.Guid  `json:"DraftStockTransactionID"`
	EndDate                    *types.Date `json:"EndDate"`
	IsBlocked                  byte        `json:"IsBlocked"`
	IsDraft                    byte        `json:"IsDraft"`
	Item                       types.Guid  `json:"Item"`
	ItemCode                   string      `json:"ItemCode"`
	ItemDescription            string      `json:"ItemDescription"`
	Modified                   *types.Date `json:"Modified"`
	Modifier                   types.Guid  `json:"Modifier"`
	ModifierFullName           string      `json:"ModifierFullName"`
	PickOrderLine              types.Guid  `json:"PickOrderLine"`
	Quantity                   float64     `json:"Quantity"`
	Remarks                    string      `json:"Remarks"`
	SalesReturnLine            types.Guid  `json:"SalesReturnLine"`
	StockCountLine             types.Guid  `json:"StockCountLine"`
	StockTransactionID         types.Guid  `json:"StockTransactionID"`
	StockTransactionType       int32       `json:"StockTransactionType"`
	StorageLocation            types.Guid  `json:"StorageLocation"`
	StorageLocationCode        string      `json:"StorageLocationCode"`
	StorageLocationDescription string      `json:"StorageLocationDescription"`
	Warehouse                  types.Guid  `json:"Warehouse"`
	WarehouseCode              string      `json:"WarehouseCode"`
	WarehouseDescription       string      `json:"WarehouseDescription"`
}

type GetStockBatchNumbersCall struct {
	modifiedAfter *time.Time
	urlNext       string
	service       *Service
}

func (service *Service) NewGetStockBatchNumbersCall(modifiedAfter *time.Time) *GetStockBatchNumbersCall {
	call := GetStockBatchNumbersCall{}
	call.modifiedAfter = modifiedAfter
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", StockBatchNumber{})
	call.urlNext = service.url(fmt.Sprintf("StockBatchNumbers?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetStockBatchNumbersCall) Do() (*[]StockBatchNumber, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	stockBatchNumbers := []StockBatchNumber{}

	next, err := call.service.Get(call.urlNext, &stockBatchNumbers)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &stockBatchNumbers, nil
}

func (service *Service) GetStockBatchNumbersCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("StockBatchNumbers", createdBefore)
}
