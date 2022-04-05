package exactonline

import (
	"fmt"
	"strings"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// Item stores Item from exactonline
//
type Item struct {
	ID                      types.Guid  `json:"ID"`
	AverageCost             float64     `json:"AverageCost"`
	Barcode                 string      `json:"Barcode"`
	Class01                 string      `json:"Class_01"`
	Class02                 string      `json:"Class_02"`
	Class03                 string      `json:"Class_03"`
	Class04                 string      `json:"Class_04"`
	Class05                 string      `json:"Class_05"`
	Class06                 string      `json:"Class_06"`
	Class07                 string      `json:"Class_07"`
	Class08                 string      `json:"Class_08"`
	Class09                 string      `json:"Class_09"`
	Class10                 string      `json:"Class_10"`
	Code                    string      `json:"Code"`
	CopyRemarks             byte        `json:"CopyRemarks"`
	CostPriceCurrency       string      `json:"CostPriceCurrency"`
	CostPriceNew            float64     `json:"CostPriceNew"`
	CostPriceStandard       float64     `json:"CostPriceStandard"`
	Created                 *types.Date `json:"Created"`
	Creator                 types.Guid  `json:"Creator"`
	CreatorFullName         string      `json:"CreatorFullName"`
	Description             string      `json:"Description"`
	Division                int32       `json:"Division"`
	EndDate                 *types.Date `json:"EndDate"`
	ExtraDescription        string      `json:"ExtraDescription"`
	FreeBoolField01         bool        `json:"FreeBoolField_01"`
	FreeBoolField02         bool        `json:"FreeBoolField_02"`
	FreeBoolField03         bool        `json:"FreeBoolField_03"`
	FreeBoolField04         bool        `json:"FreeBoolField_04"`
	FreeBoolField05         bool        `json:"FreeBoolField_05"`
	FreeDateField01         *types.Date `json:"FreeDateField_01"`
	FreeDateField02         *types.Date `json:"FreeDateField_02"`
	FreeDateField03         *types.Date `json:"FreeDateField_03"`
	FreeDateField04         *types.Date `json:"FreeDateField_04"`
	FreeDateField05         *types.Date `json:"FreeDateField_05"`
	FreeNumberField01       float64     `json:"FreeNumberField_01"`
	FreeNumberField02       float64     `json:"FreeNumberField_02"`
	FreeNumberField03       float64     `json:"FreeNumberField_03"`
	FreeNumberField04       float64     `json:"FreeNumberField_04"`
	FreeNumberField05       float64     `json:"FreeNumberField_05"`
	FreeNumberField06       float64     `json:"FreeNumberField_06"`
	FreeNumberField07       float64     `json:"FreeNumberField_07"`
	FreeNumberField08       float64     `json:"FreeNumberField_08"`
	FreeTextField01         string      `json:"FreeTextField_01"`
	FreeTextField02         string      `json:"FreeTextField_02"`
	FreeTextField03         string      `json:"FreeTextField_03"`
	FreeTextField04         string      `json:"FreeTextField_04"`
	FreeTextField05         string      `json:"FreeTextField_05"`
	FreeTextField06         string      `json:"FreeTextField_06"`
	FreeTextField07         string      `json:"FreeTextField_07"`
	FreeTextField08         string      `json:"FreeTextField_08"`
	FreeTextField09         string      `json:"FreeTextField_09"`
	FreeTextField10         string      `json:"FreeTextField_10"`
	GLCosts                 types.Guid  `json:"GLCosts"`
	GLCostsCode             string      `json:"GLCostsCode"`
	GLCostsDescription      string      `json:"GLCostsDescription"`
	GLRevenue               types.Guid  `json:"GLRevenue"`
	GLRevenueCode           string      `json:"GLRevenueCode"`
	GLRevenueDescription    string      `json:"GLRevenueDescription"`
	GLStock                 types.Guid  `json:"GLStock"`
	GLStockCode             string      `json:"GLStockCode"`
	GLStockDescription      string      `json:"GLStockDescription"`
	GrossWeight             float64     `json:"GrossWeight"`
	IsBatchItem             byte        `json:"IsBatchItem"`
	IsFractionAllowedItem   bool        `json:"IsFractionAllowedItem"`
	IsMakeItem              byte        `json:"IsMakeItem"`
	IsNewContract           byte        `json:"IsNewContract"`
	IsOnDemandItem          byte        `json:"IsOnDemandItem"`
	IsPackageItem           bool        `json:"IsPackageItem"`
	IsPurchaseItem          bool        `json:"IsPurchaseItem"`
	IsSalesItem             bool        `json:"IsSalesItem"`
	IsSerialItem            bool        `json:"IsSerialItem"`
	IsStockItem             bool        `json:"IsStockItem"`
	IsSubcontractedItem     bool        `json:"IsSubcontractedItem"`
	IsTaxableItem           byte        `json:"IsTaxableItem"`
	IsTime                  byte        `json:"IsTime"`
	IsWebshopItem           byte        `json:"IsWebshopItem"`
	ItemGroup               types.Guid  `json:"ItemGroup"`
	ItemGroupCode           string      `json:"ItemGroupCode"`
	ItemGroupDescription    string      `json:"ItemGroupDescription"`
	Modified                *types.Date `json:"Modified"`
	Modifier                types.Guid  `json:"Modifier"`
	ModifierFullName        string      `json:"ModifierFullName"`
	NetWeight               float64     `json:"NetWeight"`
	NetWeightUnit           string      `json:"NetWeightUnit"`
	Notes                   string      `json:"Notes"`
	PictureName             string      `json:"PictureName"`
	PictureThumbnailURL     string      `json:"PictureThumbnailUrl"`
	PictureURL              string      `json:"PictureUrl"`
	SalesVatCode            string      `json:"SalesVatCode"`
	SalesVatCodeDescription string      `json:"SalesVatCodeDescription"`
	SearchCode              string      `json:"SearchCode"`
	SecurityLevel           int32       `json:"SecurityLevel"`
	StartDate               *types.Date `json:"StartDate"`
	Stock                   float64     `json:"Stock"`
	Unit                    string      `json:"Unit"`
	UnitDescription         string      `json:"UnitDescription"`
	UnitType                string      `json:"UnitType"`
}

type GetItemsCall struct {
	urlNext string
	service *Service
}

type GetItemsCallParams struct {
	ItemGroupCode *string
	ModifiedAfter *time.Time
}

func (service *Service) NewGetItemsCall(params *GetItemsCallParams) *GetItemsCall {
	call := GetItemsCall{}
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", Item{})
	call.urlNext = service.url(fmt.Sprintf("Items?$select=%s", selectFields))
	filter := []string{}

	if params != nil {
		if params.ItemGroupCode != nil {
			filter = append(filter, fmt.Sprintf("ItemGroupCode eq '%s'", *params.ItemGroupCode))
		}
		if params.ModifiedAfter != nil {
			call.urlNext += service.DateFilter("Modified", "gt", params.ModifiedAfter, true, "&")
		}
	}

	if len(filter) > 0 {
		call.urlNext = fmt.Sprintf("%s&$filter=%s", call.urlNext, strings.Join(filter, " and "))
	}

	return &call
}

func (call *GetItemsCall) Do() (*[]Item, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	items := []Item{}

	next, err := call.service.Get(call.urlNext, &items)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &items, nil
}

func (call *GetItemsCall) DoAll() (*[]Item, *errortools.Error) {
	items := []Item{}

	for true {
		_items, e := call.Do()
		if e != nil {
			return nil, e
		}

		if _items == nil {
			break
		}

		if len(*_items) == 0 {
			break
		}

		items = append(items, *_items...)
	}

	return &items, nil
}

func (service *Service) GetItemsCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("Items", createdBefore)
}
