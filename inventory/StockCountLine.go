package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// StockCountLine stores StockCountLine from exactonline
//
type StockCountLine struct {
	ID                         types.GUID  `json:"ID"`
	CostPrice                  float64     `json:"CostPrice"`
	Created                    *types.Date `json:"Created"`
	Creator                    types.GUID  `json:"Creator"`
	CreatorFullName            string      `json:"CreatorFullName"`
	Division                   int32       `json:"Division"`
	Item                       types.GUID  `json:"Item"`
	ItemCode                   string      `json:"ItemCode"`
	ItemCostPrice              float64     `json:"ItemCostPrice"`
	ItemDescription            string      `json:"ItemDescription"`
	ItemDivisable              bool        `json:"ItemDivisable"`
	LineNumber                 int32       `json:"LineNumber"`
	Modified                   *types.Date `json:"Modified"`
	Modifier                   types.GUID  `json:"Modifier"`
	ModifierFullName           string      `json:"ModifierFullName"`
	QuantityDifference         float64     `json:"QuantityDifference"`
	QuantityInStock            float64     `json:"QuantityInStock"`
	QuantityNew                float64     `json:"QuantityNew"`
	StockCountID               types.GUID  `json:"StockCountID"`
	StockKeepingUnit           string      `json:"StockKeepingUnit"`
	StorageLocation            types.GUID  `json:"StorageLocation"`
	StorageLocationCode        string      `json:"StorageLocationCode"`
	StorageLocationDescription string      `json:"StorageLocationDescription"`
}

type GetStockCountLinesCall struct {
	modifiedAfter *time.Time
	urlNext       string
	service       *Service
}

func (service *Service) NewGetStockCountLinesCall(modifiedAfter *time.Time) *GetStockCountLinesCall {
	call := GetStockCountLinesCall{}
	call.modifiedAfter = modifiedAfter
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", StockCountLine{})
	call.urlNext = service.url(fmt.Sprintf("StockCountLines?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetStockCountLinesCall) Do() (*[]StockCountLine, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	stockCountLines := []StockCountLine{}

	next, err := call.service.Get(call.urlNext, &stockCountLines)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &stockCountLines, nil
}

func (service *Service) GetStockCountLinesCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("StockCountLines", createdBefore)
}
