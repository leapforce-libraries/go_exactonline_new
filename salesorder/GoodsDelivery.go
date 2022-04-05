package exactonline

import (
	"encoding/json"
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// GoodsDelivery stores GoodsDelivery from exactonline
//
type GoodsDelivery struct {
	EntryID                       types.Guid      `json:"EntryID"`
	Created                       *types.Date     `json:"Created"`
	Creator                       types.Guid      `json:"Creator"`
	CreatorFullName               string          `json:"CreatorFullName"`
	DeliveryAccount               types.Guid      `json:"DeliveryAccount"`
	DeliveryAccountCode           string          `json:"DeliveryAccountCode"`
	DeliveryAccountName           string          `json:"DeliveryAccountName"`
	DeliveryAddress               types.Guid      `json:"DeliveryAddress"`
	DeliveryContact               types.Guid      `json:"DeliveryContact"`
	DeliveryContactPersonFullName string          `json:"DeliveryContactPersonFullName"`
	DeliveryDate                  *types.Date     `json:"DeliveryDate"`
	DeliveryNumber                int32           `json:"DeliveryNumber"`
	Description                   string          `json:"Description"`
	Division                      int32           `json:"Division"`
	Document                      types.Guid      `json:"Document"`
	DocumentSubject               string          `json:"DocumentSubject"`
	EntryNumber                   int32           `json:"EntryNumber"`
	GoodsDeliveryLines            json.RawMessage `json:"GoodsDeliveryLines"` //to be implemented when needed
	Modified                      *types.Date     `json:"Modified"`
	Modifier                      types.Guid      `json:"Modifier"`
	ModifierFullName              string          `json:"ModifierFullName"`
	Remarks                       string          `json:"Remarks"`
	ShippingMethod                types.Guid      `json:"ShippingMethod"`
	ShippingMethodCode            string          `json:"ShippingMethodCode"`
	ShippingMethodDescription     string          `json:"ShippingMethodDescription"`
	TrackingNumber                string          `json:"TrackingNumber"`
	Warehouse                     types.Guid      `json:"Warehouse"`
	WarehouseCode                 string          `json:"WarehouseCode"`
	WarehouseDescription          string          `json:"WarehouseDescription"`
}

type GetGoodsDeliveriesCall struct {
	modifiedAfter *time.Time
	urlNext       string
	service       *Service
}

func (service *Service) NewGetGoodsDeliveriesCall(modifiedAfter *time.Time) *GetGoodsDeliveriesCall {
	call := GetGoodsDeliveriesCall{}
	call.modifiedAfter = modifiedAfter
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", GoodsDelivery{})
	call.urlNext = service.url(fmt.Sprintf("GoodsDeliveries?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetGoodsDeliveriesCall) Do() (*[]GoodsDelivery, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	goodsDeliveries := []GoodsDelivery{}

	next, err := call.service.Get(call.urlNext, &goodsDeliveries)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &goodsDeliveries, nil
}

func (service *Service) GetGoodsDeliveriesCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("GoodsDeliveries", createdBefore)
}
