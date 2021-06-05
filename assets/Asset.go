package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// Asset stores Asset from exactonline
//
type Asset struct {
	ID                            types.GUID  `json:"ID"`
	AlreadyDepreciated            byte        `json:"AlreadyDepreciated"`
	AssetFrom                     types.GUID  `json:"AssetFrom"`
	AssetFromDescription          string      `json:"AssetFromDescription"`
	AssetGroup                    types.GUID  `json:"AssetGroup"`
	AssetGroupCode                string      `json:"AssetGroupCode"`
	AssetGroupDescription         string      `json:"AssetGroupDescription"`
	CatalogueValue                float64     `json:"CatalogueValue"`
	Code                          string      `json:"Code"`
	Costcenter                    string      `json:"Costcenter"`
	CostcenterDescription         string      `json:"CostcenterDescription"`
	Costunit                      string      `json:"Costunit"`
	CostunitDescription           string      `json:"CostunitDescription"`
	Created                       *types.Date `json:"Created"`
	Creator                       types.GUID  `json:"Creator"`
	CreatorFullName               string      `json:"CreatorFullName"`
	DeductionPercentage           float64     `json:"DeductionPercentage"`
	DepreciatedAmount             float64     `json:"DepreciatedAmount"`
	DepreciatedPeriods            int32       `json:"DepreciatedPeriods"`
	DepreciatedStartDate          *types.Date `json:"DepreciatedStartDate"`
	Description                   string      `json:"Description"`
	Division                      int32       `json:"Division"`
	EndDate                       *types.Date `json:"EndDate"`
	EngineEmission                int16       `json:"EngineEmission"`
	EngineType                    int16       `json:"EngineType"`
	GLTransactionLine             types.GUID  `json:"GLTransactionLine"`
	GLTransactionLineDescription  string      `json:"GLTransactionLineDescription"`
	InvestmentAccount             types.GUID  `json:"InvestmentAccount"`
	InvestmentAccountCode         string      `json:"InvestmentAccountCode"`
	InvestmentAccountName         string      `json:"InvestmentAccountName"`
	InvestmentAmountDC            float64     `json:"InvestmentAmountDC"`
	InvestmentAmountFC            float64     `json:"InvestmentAmountFC"`
	InvestmentCurrency            string      `json:"InvestmentCurrency"`
	InvestmentCurrencyDescription string      `json:"InvestmentCurrencyDescription"`
	InvestmentDate                *types.Date `json:"InvestmentDate"`
	InvestmentDeduction           int16       `json:"InvestmentDeduction"`
	Modified                      *types.Date `json:"Modified"`
	Modifier                      types.GUID  `json:"Modifier"`
	ModifierFullName              string      `json:"ModifierFullName"`
	Notes                         string      `json:"Notes"`
	Parent                        types.GUID  `json:"Parent"`
	ParentCode                    string      `json:"ParentCode"`
	ParentDescription             string      `json:"ParentDescription"`
	PictureFileName               string      `json:"PictureFileName"`
	PrimaryMethod                 types.GUID  `json:"PrimaryMethod"`
	PrimaryMethodCode             string      `json:"PrimaryMethodCode"`
	PrimaryMethodDescription      string      `json:"PrimaryMethodDescription"`
	ResidualValue                 float64     `json:"ResidualValue"`
	StartDate                     *types.Date `json:"StartDate"`
	Status                        int16       `json:"Status"`
	TransactionEntryID            types.GUID  `json:"TransactionEntryID"`
	TransactionEntryNo            int32       `json:"TransactionEntryNo"`
}

type GetAssetsCall struct {
	modifiedAfter *time.Time
	urlNext       string
	service       *Service
}

func (service *Service) NewGetAssetsCall(modifiedAfter *time.Time) *GetAssetsCall {
	call := GetAssetsCall{}
	call.modifiedAfter = modifiedAfter
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", Asset{})
	call.urlNext = service.url(fmt.Sprintf("Assets?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetAssetsCall) Do() (*[]Asset, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	assets := []Asset{}

	next, err := call.service.Get(call.urlNext, &assets)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &assets, nil
}

func (service *Service) GetAssetsCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("Assets", createdBefore)
}
