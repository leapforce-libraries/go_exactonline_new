package exactonline

import (
	"encoding/json"
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// SalesOrderGoodDelivery stores SalesOrderGoodDelivery from exactonline
//
type SalesOrderGoodDelivery struct {
	Timestamp                     types.Int64String `json:"Timestamp"`
	EntryID                       types.GUID        `json:"EntryID"`
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
	EntryNumber                   int32             `json:"EntryNumber"`
	GoodsDeliveryLines            json.RawMessage   `json:"GoodsDeliveryLines"` //to be implemented when needed
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

type SyncSalesOrderGoodDeliveriesCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewSyncSalesOrderGoodDeliveriesCall(timestamp *int64) *SyncSalesOrderGoodDeliveriesCall {
	selectFields := utilities.GetTaggedTagNames("json", SalesOrderGoodDelivery{})
	url := service.url(fmt.Sprintf("SalesOrder/GoodDeliveries?$select=%s", selectFields))
	if timestamp != nil {
		url = fmt.Sprintf("%s&$filter=Timestamp gt %vL", url, *timestamp)
	}

	return &SyncSalesOrderGoodDeliveriesCall{
		service: service,
		urlNext: url,
	}
}

func (call *SyncSalesOrderGoodDeliveriesCall) Do() (*[]SalesOrderGoodDelivery, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	salesOrderGoodDeliveries := []SalesOrderGoodDelivery{}

	next, err := call.service.Get(call.urlNext, &salesOrderGoodDeliveries)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &salesOrderGoodDeliveries, nil
}
