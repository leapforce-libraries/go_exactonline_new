package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// Item stores Item from exactonline
//
type Item struct {
	ID                      types.GUID  `json:"ID "`
	AverageCost             float64     `json:"AverageCost"`
	Barcode                 string      `json:"Barcode"`
	Class_01                string      `json:"Class_01"`
	Class_02                string      `json:"Class_02"`
	Class_03                string      `json:"Class_03"`
	Class_04                string      `json:"Class_04"`
	Class_05                string      `json:"Class_05"`
	Class_06                string      `json:"Class_06"`
	Class_07                string      `json:"Class_07"`
	Class_08                string      `json:"Class_08"`
	Class_09                string      `json:"Class_09"`
	Class_10                string      `json:"Class_10"`
	Code                    string      `json:"Code"`
	CopyRemarks             byte        `json:"CopyRemarks"`
	CostPriceCurrency       string      `json:"CostPriceCurrency"`
	CostPriceNew            float64     `json:"CostPriceNew"`
	CostPriceStandard       float64     `json:"CostPriceStandard"`
	Created                 *types.Date `json:"Created"`
	Creator                 types.GUID  `json:"Creator"`
	CreatorFullName         string      `json:"CreatorFullName"`
	Description             string      `json:"Description"`
	Division                int32       `json:"Division"`
	EndDate                 *types.Date `json:"EndDate"`
	ExtraDescription        string      `json:"ExtraDescription"`
	FreeBoolField_01        bool        `json:"FreeBoolField_01"`
	FreeBoolField_02        bool        `json:"FreeBoolField_02"`
	FreeBoolField_03        bool        `json:"FreeBoolField_03"`
	FreeBoolField_04        bool        `json:"FreeBoolField_04"`
	FreeBoolField_05        bool        `json:"FreeBoolField_05"`
	FreeDateField_01        *types.Date `json:"FreeDateField_01"`
	FreeDateField_02        *types.Date `json:"FreeDateField_02"`
	FreeDateField_03        *types.Date `json:"FreeDateField_03"`
	FreeDateField_04        *types.Date `json:"FreeDateField_04"`
	FreeDateField_05        *types.Date `json:"FreeDateField_05"`
	FreeNumberField_01      float64     `json:"FreeNumberField_01"`
	FreeNumberField_02      float64     `json:"FreeNumberField_02"`
	FreeNumberField_03      float64     `json:"FreeNumberField_03"`
	FreeNumberField_04      float64     `json:"FreeNumberField_04"`
	FreeNumberField_05      float64     `json:"FreeNumberField_05"`
	FreeNumberField_06      float64     `json:"FreeNumberField_06"`
	FreeNumberField_07      float64     `json:"FreeNumberField_07"`
	FreeNumberField_08      float64     `json:"FreeNumberField_08"`
	FreeTextField_01        string      `json:"FreeTextField_01"`
	FreeTextField_02        string      `json:"FreeTextField_02"`
	FreeTextField_03        string      `json:"FreeTextField_03"`
	FreeTextField_04        string      `json:"FreeTextField_04"`
	FreeTextField_05        string      `json:"FreeTextField_05"`
	FreeTextField_06        string      `json:"FreeTextField_06"`
	FreeTextField_07        string      `json:"FreeTextField_07"`
	FreeTextField_08        string      `json:"FreeTextField_08"`
	FreeTextField_09        string      `json:"FreeTextField_09"`
	FreeTextField_10        string      `json:"FreeTextField_10"`
	GLCosts                 types.GUID  `json:"GLCosts"`
	GLCostsCode             string      `json:"GLCostsCode"`
	GLCostsDescription      string      `json:"GLCostsDescription"`
	GLRevenue               types.GUID  `json:"GLRevenue"`
	GLRevenueCode           string      `json:"GLRevenueCode"`
	GLRevenueDescription    string      `json:"GLRevenueDescription"`
	GLStock                 types.GUID  `json:"GLStock"`
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
	ItemGroup               types.GUID  `json:"ItemGroup"`
	ItemGroupCode           string      `json:"ItemGroupCode"`
	ItemGroupDescription    string      `json:"ItemGroupDescription"`
	Modified                *types.Date `json:"Modified"`
	Modifier                types.GUID  `json:"Modifier"`
	ModifierFullName        string      `json:"ModifierFullName"`
	NetWeight               float64     `json:"NetWeight"`
	NetWeightUnit           string      `json:"NetWeightUnit"`
	Notes                   string      `json:"Notes"`
	PictureName             string      `json:"PictureName"`
	PictureThumbnailUrl     string      `json:"PictureThumbnailUrl"`
	PictureUrl              string      `json:"PictureUrl"`
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
	modifiedAfter *time.Time
	urlNext       string
	client        *Client
}

func (c *Client) NewGetItemsCall(modifiedAfter *time.Time) *GetItemsCall {
	call := GetItemsCall{}
	call.modifiedAfter = modifiedAfter
	call.client = c

	selectFields := utilities.GetTaggedTagNames("json", Item{})
	call.urlNext = fmt.Sprintf("%s/Items?$select=%s", c.BaseURL(), selectFields)
	if modifiedAfter != nil {
		call.urlNext += c.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetItemsCall) Do() (*[]Item, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	items := []Item{}

	next, err := call.client.Get(call.urlNext, &items)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &items, nil
}

func (c *Client) GetItemsCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return c.GetCount("Items", createdBefore)
}
