package exactonline

import (
	"encoding/json"
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// GoodsReceiptLine stores GoodsReceiptLine from exactonline
//
type GoodsReceiptLine struct {
	ID                  types.Guid      `json:"ID"`
	BatchNumbers        json.RawMessage `json:"BatchNumbers"`
	Created             *types.Date     `json:"Created"`
	Creator             types.Guid      `json:"Creator"`
	CreatorFullName     string          `json:"CreatorFullName"`
	Description         string          `json:"Description"`
	Division            int32           `json:"Division"`
	GoodsReceiptID      types.Guid      `json:"GoodsReceiptID"`
	Item                types.Guid      `json:"Item"`
	ItemCode            string          `json:"ItemCode"`
	ItemDescription     string          `json:"ItemDescription"`
	ItemUnitCode        string          `json:"ItemUnitCode"`
	LineNumber          int32           `json:"LineNumber"`
	Location            types.Guid      `json:"Location"`
	LocationCode        string          `json:"LocationCode"`
	LocationDescription string          `json:"LocationDescription"`
	Modified            *types.Date     `json:"Modified"`
	Modifier            types.Guid      `json:"Modifier"`
	ModifierFullName    string          `json:"ModifierFullName"`
	Notes               string          `json:"Notes"`
	Project             types.Guid      `json:"Project"`
	ProjectCode         string          `json:"ProjectCode"`
	ProjectDescription  string          `json:"ProjectDescription"`
	PurchaseOrderID     types.Guid      `json:"PurchaseOrderID"`
	PurchaseOrderLineID types.Guid      `json:"PurchaseOrderLineID"`
	PurchaseOrderNumber int32           `json:"PurchaseOrderNumber"`
	QuantityOrdered     float64         `json:"QuantityOrdered"`
	QuantityReceived    float64         `json:"QuantityReceived"`
	SerialNumbers       json.RawMessage `json:"SerialNumbers"`
	SupplierItemCode    string          `json:"SupplierItemCode"`
}

type GetGoodsReceiptLinesCall struct {
	modifiedAfter *time.Time
	urlNext       string
	service       *Service
}

func (service *Service) NewGetGoodsReceiptLinesCall(modifiedAfter *time.Time) *GetGoodsReceiptLinesCall {
	call := GetGoodsReceiptLinesCall{}
	call.modifiedAfter = modifiedAfter
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", GoodsReceiptLine{})
	call.urlNext = service.url(fmt.Sprintf("GoodsReceiptLines?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetGoodsReceiptLinesCall) Do() (*[]GoodsReceiptLine, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	goodsReceiptLines := []GoodsReceiptLine{}

	next, err := call.service.Get(call.urlNext, &goodsReceiptLines)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &goodsReceiptLines, nil
}

func (service *Service) GetGoodsReceiptLinesCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("GoodsReceiptLines", createdBefore)
}
