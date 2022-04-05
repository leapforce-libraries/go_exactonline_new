package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// PlannedSalesReturnLine stores PlannedSalesReturnLine from exactonline
//
type PlannedSalesReturnLine struct {
	ID types.Guid `json:"ID"`
	//BatchNumbers               []inventory.StockBatchNumber  `json:"BatchNumbers"`
	CreateCredit          byte        `json:"CreateCredit"`
	Created               *types.Date `json:"Created"`
	Creator               types.Guid  `json:"Creator"`
	CreatorFullName       string      `json:"CreatorFullName"`
	Division              int32       `json:"Division"`
	GoodDeliveryLineID    types.Guid  `json:"GoodDeliveryLineID"`
	Item                  types.Guid  `json:"Item"`
	ItemCode              string      `json:"ItemCode"`
	ItemDescription       string      `json:"ItemDescription"`
	LineNumber            int32       `json:"LineNumber"`
	Modified              *types.Date `json:"Modified"`
	Modifier              types.Guid  `json:"Modifier"`
	ModifierFullName      string      `json:"ModifierFullName"`
	Notes                 string      `json:"Notes"`
	PlannedReturnQuantity float64     `json:"PlannedReturnQuantity"`
	PlannedSalesReturnID  types.Guid  `json:"PlannedSalesReturnID"`
	ReceivedQuantity      float64     `json:"ReceivedQuantity"`
	SalesOrderLineID      types.Guid  `json:"SalesOrderLineID"`
	SalesOrderNumber      int32       `json:"SalesOrderNumber"`
	//SerialNumbers              []inventory.StockSerialNumber `json:"SerialNumbers"`
	StockTransactionEntryID    types.Guid `json:"StockTransactionEntryID"`
	StorageLocation            types.Guid `json:"StorageLocation"`
	StorageLocationCode        string     `json:"StorageLocationCode"`
	StorageLocationDescription string     `json:"StorageLocationDescription"`
	UnitCode                   string     `json:"UnitCode"`
	UnitDescription            string     `json:"UnitDescription"`
}

type GetPlannedSalesReturnLinesCall struct {
	modifiedAfter *time.Time
	urlNext       string
	service       *Service
}

func (service *Service) NewGetPlannedSalesReturnLinesCall(modifiedAfter *time.Time) *GetPlannedSalesReturnLinesCall {
	call := GetPlannedSalesReturnLinesCall{}
	call.modifiedAfter = modifiedAfter
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", PlannedSalesReturnLine{})
	call.urlNext = service.url(fmt.Sprintf("PlannedSalesReturnLines?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetPlannedSalesReturnLinesCall) Do() (*[]PlannedSalesReturnLine, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	plannedSalesReturnLines := []PlannedSalesReturnLine{}

	next, err := call.service.Get(call.urlNext, &plannedSalesReturnLines)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &plannedSalesReturnLines, nil
}
func (service *Service) GetPlannedSalesReturnLinesCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("PlannedSalesReturnLines", createdBefore)
}
