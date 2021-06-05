package exactonline

import (
	"fmt"
	"strings"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// ItemGroup stores ItemGroup from exactonline
//
type ItemGroup struct {
	ID                             types.GUID  `json:"ID"`
	Code                           string      `json:"Code"`
	Created                        *types.Date `json:"Created"`
	Creator                        types.GUID  `json:"Creator"`
	CreatorFullName                string      `json:"CreatorFullName"`
	Description                    string      `json:"Description"`
	Division                       int32       `json:"Division"`
	GLCosts                        types.GUID  `json:"GLCosts"`
	GLCostsCode                    string      `json:"GLCostsCode"`
	GLCostsDescription             string      `json:"GLCostsDescription"`
	GLPurchaseAccount              types.GUID  `json:"GLPurchaseAccount"`
	GLPurchaseAccountCode          string      `json:"GLPurchaseAccountCode"`
	GLPurchaseAccountDescription   string      `json:"GLPurchaseAccountDescription"`
	GLPurchasePriceDifference      types.GUID  `json:"GLPurchasePriceDifference"`
	GLPurchasePriceDifferenceCode  string      `json:"GLPurchasePriceDifferenceCode"`
	GLPurchasePriceDifferenceDescr string      `json:"GLPurchasePriceDifferenceDescr"`
	GLRevenue                      types.GUID  `json:"GLRevenue"`
	GLRevenueCode                  string      `json:"GLRevenueCode"`
	GLRevenueDescription           string      `json:"GLRevenueDescription"`
	GLStock                        types.GUID  `json:"GLStock"`
	GLStockCode                    string      `json:"GLStockCode"`
	GLStockDescription             string      `json:"GLStockDescription"`
	GLStockVariance                types.GUID  `json:"GLStockVariance"`
	GLStockVarianceCode            string      `json:"GLStockVarianceCode"`
	GLStockVarianceDescription     string      `json:"GLStockVarianceDescription"`
	IsDefault                      byte        `json:"IsDefault"`
	Modified                       *types.Date `json:"Modified"`
	Modifier                       types.GUID  `json:"Modifier"`
	ModifierFullName               string      `json:"ModifierFullName"`
	Notes                          string      `json:"Notes"`
}

type GetItemGroupsCall struct {
	urlNext string
	service *Service
}

type GetItemGroupsCallParams struct {
	ModifiedAfter *time.Time
}

func (service *Service) NewGetItemGroupsCall(params *GetItemGroupsCallParams) *GetItemGroupsCall {
	call := GetItemGroupsCall{}
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", ItemGroup{})
	call.urlNext = service.url(fmt.Sprintf("ItemGroups?$select=%s", selectFields))
	filter := []string{}

	if params != nil {
		if params.ModifiedAfter != nil {
			call.urlNext += service.DateFilter("Modified", "gt", params.ModifiedAfter, true, "&")
		}
	}

	if len(filter) > 0 {
		call.urlNext = fmt.Sprintf("%s&$filter=%s", call.urlNext, strings.Join(filter, " and "))
	}

	return &call
}

func (call *GetItemGroupsCall) Do() (*[]ItemGroup, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	items := []ItemGroup{}

	next, err := call.service.Get(call.urlNext, &items)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &items, nil
}

func (call *GetItemGroupsCall) DoAll() (*[]ItemGroup, *errortools.Error) {
	items := []ItemGroup{}

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

func (service *Service) GetItemGroupsCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("ItemGroups", createdBefore)
}
