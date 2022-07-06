package exactonline

import (
	"encoding/json"
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// GoodsReceipt stores GoodsReceipt from exactonline
//
type GoodsReceipt struct {
	ID                      types.Guid      `json:"ID"`
	Created                 *types.Date     `json:"Created"`
	Creator                 types.Guid      `json:"Creator"`
	CreatorFullName         string          `json:"CreatorFullName"`
	Description             string          `json:"Description"`
	Division                int32           `json:"Division"`
	Document                types.Guid      `json:"Document"`
	DocumentSubject         string          `json:"DocumentSubject"`
	EntryNumber             int32           `json:"EntryNumber"`
	GoodsReceiptLineCount   int32           `json:"GoodsReceiptLineCount"`
	GoodsReceiptLines       json.RawMessage `json:"GoodsReceiptLines"`
	Modified                *types.Date     `json:"Modified"`
	Modifier                types.Guid      `json:"Modifier"`
	ModifierFullName        string          `json:"ModifierFullName"`
	ReceiptDate             *types.Date     `json:"ReceiptDate"`
	ReceiptNumber           int32           `json:"ReceiptNumber"`
	Remarks                 string          `json:"Remarks"`
	Supplier                types.Guid      `json:"Supplier"`
	SupplierCode            string          `json:"SupplierCode"`
	SupplierContact         types.Guid      `json:"SupplierContact"`
	SupplierContactFullName string          `json:"SupplierContactFullName"`
	SupplierName            string          `json:"SupplierName"`
	Warehouse               types.Guid      `json:"Warehouse"`
	WarehouseCode           string          `json:"WarehouseCode"`
	WarehouseDescription    string          `json:"WarehouseDescription"`
	YourRef                 string          `json:"YourRef"`
}

type GetGoodsReceiptsCall struct {
	modifiedAfter *time.Time
	urlNext       string
	service       *Service
}

func (service *Service) NewGetGoodsReceiptsCall(modifiedAfter *time.Time) *GetGoodsReceiptsCall {
	call := GetGoodsReceiptsCall{}
	call.modifiedAfter = modifiedAfter
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", GoodsReceipt{})
	call.urlNext = service.url(fmt.Sprintf("GoodsReceipts?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetGoodsReceiptsCall) Do() (*[]GoodsReceipt, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	goodsReceipts := []GoodsReceipt{}

	next, err := call.service.Get(call.urlNext, &goodsReceipts)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &goodsReceipts, nil
}

func (service *Service) GetGoodsReceiptsCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("GoodsReceipts", createdBefore)
}
