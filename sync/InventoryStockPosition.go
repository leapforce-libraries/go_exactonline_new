package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// InventoryStockPosition stores InventoryStockPosition from exactonline
//
type InventoryStockPosition struct {
	Timestamp            int64      `json:"Timestamp"`
	CurrentStock         float64    `json:"CurrentStock"`
	Division             int32      `json:"Division"`
	FreeStock            float64    `json:"FreeStock"`
	ID                   types.Guid `json:"ID"`
	ItemCode             string     `json:"ItemCode"`
	ItemDescription      string     `json:"ItemDescription"`
	ItemId               types.Guid `json:"ItemId"`
	PlanningIn           float64    `json:"PlanningIn"`
	PlanningOut          float64    `json:"PlanningOut"`
	ProjectedStock       float64    `json:"ProjectedStock"`
	ReorderPoint         float64    `json:"ReorderPoint"`
	ReservedStock        float64    `json:"ReservedStock"`
	UnitCode             string     `json:"UnitCode"`
	UnitDescription      string     `json:"UnitDescription"`
	Warehouse            types.Guid `json:"Warehouse"`
	WarehouseCode        string     `json:"WarehouseCode"`
	WarehouseDescription string     `json:"WarehouseDescription"`
}

type SyncInventoryStockPositionsCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewSyncInventoryStockPositionsCall(timestamp *int64) *SyncInventoryStockPositionsCall {
	selectFields := utilities.GetTaggedTagNames("json", InventoryStockPosition{})
	url := service.url(fmt.Sprintf("Inventory/StockPositions?$select=%s", selectFields))
	if timestamp != nil {
		url = fmt.Sprintf("%s&$filter=Timestamp gt %vL", url, *timestamp)
	}

	return &SyncInventoryStockPositionsCall{
		service: service,
		urlNext: url,
	}
}

func (call *SyncInventoryStockPositionsCall) Do() (*[]InventoryStockPosition, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	logisticsItems := []InventoryStockPosition{}

	next, err := call.service.Get(call.urlNext, &logisticsItems)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &logisticsItems, nil
}
