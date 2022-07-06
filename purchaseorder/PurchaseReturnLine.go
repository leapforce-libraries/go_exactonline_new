package exactonline

import (
	"encoding/json"
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// PurchaseReturnLine stores PurchaseReturnLine from exactonline
//
type PurchaseReturnLine struct {
	ID                          types.Guid      `json:"ID"`
	BatchNumbers                json.RawMessage `json:"BatchNumbers"`
	CreateCredit                bool            `json:"CreateCredit"`
	Created                     *types.Date     `json:"Created"`
	Creator                     types.Guid      `json:"Creator"`
	CreatorFullName             string          `json:"CreatorFullName"`
	Division                    int32           `json:"Division"`
	EntryID                     types.Guid      `json:"EntryID"`
	GoodsReceiptLineID          types.Guid      `json:"GoodsReceiptLineID"`
	Item                        types.Guid      `json:"Item"`
	ItemCode                    string          `json:"ItemCode"`
	ItemDescription             string          `json:"ItemDescription"`
	LineNumber                  int32           `json:"LineNumber"`
	Location                    types.Guid      `json:"Location"`
	LocationCode                string          `json:"LocationCode"`
	LocationDescription         string          `json:"LocationDescription"`
	Modified                    *types.Date     `json:"Modified"`
	Modifier                    types.Guid      `json:"Modifier"`
	ModifierFullName            string          `json:"ModifierFullName"`
	Notes                       string          `json:"Notes"`
	PurchaseOrderLineID         types.Guid      `json:"PurchaseOrderLineID"`
	PurchaseOrderNumber         int32           `json:"PurchaseOrderNumber"`
	ReceiptNumber               int32           `json:"ReceiptNumber"`
	ReceivedQuantity            float64         `json:"ReceivedQuantity"`
	ReturnQuantity              float64         `json:"ReturnQuantity"`
	ReturnReasonCodeDescription string          `json:"ReturnReasonCodeDescription"`
	ReturnReasonCodeID          types.Guid      `json:"ReturnReasonCodeID"`
	SerialNumbers               json.RawMessage `json:"SerialNumbers"`
	SupplierItemCode            string          `json:"SupplierItemCode"`
	UnitCode                    string          `json:"UnitCode"`
}

type GetPurchaseReturnLinesCall struct {
	modifiedAfter *time.Time
	urlNext       string
	service       *Service
}

func (service *Service) NewGetPurchaseReturnLinesCall(modifiedAfter *time.Time) *GetPurchaseReturnLinesCall {
	call := GetPurchaseReturnLinesCall{}
	call.modifiedAfter = modifiedAfter
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", PurchaseReturnLine{})
	call.urlNext = service.url(fmt.Sprintf("PurchaseReturnLines?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetPurchaseReturnLinesCall) Do() (*[]PurchaseReturnLine, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	purchaseReturnLines := []PurchaseReturnLine{}

	next, err := call.service.Get(call.urlNext, &purchaseReturnLines)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &purchaseReturnLines, nil
}

func (service *Service) GetPurchaseReturnLinesCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("PurchaseReturnLines", createdBefore)
}
