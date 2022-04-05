package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// StockSerialNumber stores StockSerialNumber from exactonline
//
type StockSerialNumber struct {
	ID                         types.Guid  `json:"ID"`
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
	Remarks                    string      `json:"Remarks"`
	SalesReturnLine            types.Guid  `json:"SalesReturnLine"`
	SerialNumber               string      `json:"SerialNumber"`
	SerialNumberID             types.Guid  `json:"SerialNumberID"`
	StartDate                  *types.Date `json:"StartDate"`
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

type GetStockSerialNumbersCall struct {
	modifiedAfter *time.Time
	urlNext       string
	service       *Service
}

func (service *Service) NewGetStockSerialNumbersCall(modifiedAfter *time.Time) *GetStockSerialNumbersCall {
	call := GetStockSerialNumbersCall{}
	call.modifiedAfter = modifiedAfter
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", StockSerialNumber{})
	call.urlNext = service.url(fmt.Sprintf("StockSerialNumbers?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetStockSerialNumbersCall) Do() (*[]StockSerialNumber, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	stockSerialNumbers := []StockSerialNumber{}

	next, err := call.service.Get(call.urlNext, &stockSerialNumbers)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &stockSerialNumbers, nil
}

func (service *Service) GetStockSerialNumbersCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("StockSerialNumbers", createdBefore)
}
