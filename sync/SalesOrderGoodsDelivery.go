package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// SalesOrderGoodsDelivery stores SalesOrderGoodsDelivery from exactonline
//
type SalesOrderGoodsDelivery struct {
	Timestamp                     types.Int64String `json:"Timestamp"`
	Created                       *types.Date       `json:"Created"`
	Creator                       types.GUID        `json:"Creator"`
	CreatorFullName               string            `json:"CreatorFullName"`
	DeliveryAccount               types.GUID        `json:"DeliveryAccount"`
	DeliveryAccountCode           string            `json:"DeliveryAccountCode"`
	DeliveryAccountName           string            `json:"DeliveryAccountName"`
	DeliveryAddress               types.GUID        `json:"DeliveryAddress"`
	DeliveryContact               types.GUID        `json:"DeliveryContact"`
	DeliveryContactPersonFullName string            `json:"DeliveryContactPersonFullName"`
	DeliveryDate                  *types.Date       `json:"DeliveryDate"`
	DeliveryNumber                int32             `json:"DeliveryNumber"`
	Description                   string            `json:"Description"`
	Division                      int32             `json:"Division"`
	Document                      types.GUID        `json:"Document"`
	DocumentSubject               string            `json:"DocumentSubject"`
	EntryID                       types.GUID        `json:"EntryID"`
	EntryNumber                   int32             `json:"EntryNumber"`
	Modified                      *types.Date       `json:"Modified"`
	Modifier                      types.GUID        `json:"Modifier"`
	ModifierFullName              string            `json:"ModifierFullName"`
	Remarks                       string            `json:"Remarks"`
	ShippingMethod                types.GUID        `json:"ShippingMethod"`
	ShippingMethodCode            string            `json:"ShippingMethodCode"`
	ShippingMethodDescription     string            `json:"ShippingMethodDescription"`
	TrackingNumber                string            `json:"TrackingNumber"`
	Warehouse                     types.GUID        `json:"Warehouse"`
	WarehouseCode                 string            `json:"WarehouseCode"`
	WarehouseDescription          string            `json:"WarehouseDescription"`
}

type SyncSalesOrderGoodsDeliveriesCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewSyncSalesOrderGoodsDeliveriesCall(timestamp *int64) *SyncSalesOrderGoodsDeliveriesCall {
	selectFields := utilities.GetTaggedTagNames("json", SalesOrderGoodsDelivery{})
	url := service.url(fmt.Sprintf("SalesOrder/GoodsDeliveries?$select=%s", selectFields))
	if timestamp != nil {
		url = fmt.Sprintf("%s&$filter=Timestamp gt %vL", url, *timestamp)
	}

	return &SyncSalesOrderGoodsDeliveriesCall{
		service: service,
		urlNext: url,
	}
}

func (call *SyncSalesOrderGoodsDeliveriesCall) Do() (*[]SalesOrderGoodsDelivery, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	salesOrderGoodsDeliveries := []SalesOrderGoodsDelivery{}

	next, err := call.service.Get(call.urlNext, &salesOrderGoodsDeliveries)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &salesOrderGoodsDeliveries, nil
}
