package exactonline

import (
	"fmt"
	types "github.com/leapforce-libraries/go_types"
	"strings"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// AccountClassification stores AccountClassification from exactonline
type AccountClassification struct {
	ID                                   types.Guid  `json:"ID"`
	AccountClassificationName            types.Guid  `json:"AccountClassificationName"`
	AccountClassificationNameDescription string      `json:"AccountClassificationNameDescription"`
	Code                                 string      `json:"Code"`
	Created                              *types.Date `json:"Created"`
	Creator                              types.Guid  `json:"Creator"`
	CreatorFullName                      string      `json:"CreatorFullName"`
	Description                          string      `json:"Description"`
	Division                             int32       `json:"Division"`
	Modified                             *types.Date `json:"Modified"`
	Modifier                             types.Guid  `json:"Modifier"`
	ModifierFullName                     string      `json:"ModifierFullName"`
}

type GetAccountClassificationsCall struct {
	urlNext string
	service *Service
}

type GetAccountClassificationsCallParams struct {
	ModifiedAfter *time.Time
}

func (service *Service) NewGetAccountClassificationsCall(params *GetAccountClassificationsCallParams) *GetAccountClassificationsCall {
	call := GetAccountClassificationsCall{}
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", AccountClassification{})
	call.urlNext = service.url(fmt.Sprintf("AccountClassifications?$select=%s", selectFields))
	filter := []string{}

	if params != nil {
		if params.ModifiedAfter != nil {
			filter = append(filter, service.DateFilter("Modified", "gt", params.ModifiedAfter, false, ""))
		}
	}

	if len(filter) > 0 {
		call.urlNext = fmt.Sprintf("%s&$filter=%s", call.urlNext, strings.Join(filter, " and "))
	}

	return &call
}

func (call *GetAccountClassificationsCall) Do() (*[]AccountClassification, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	accountClassifications := []AccountClassification{}

	next, err := call.service.Get(call.urlNext, &accountClassifications)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &accountClassifications, nil
}

func (call *GetAccountClassificationsCall) DoAll() (*[]AccountClassification, *errortools.Error) {
	accountClassifications := []AccountClassification{}

	for true {
		_accountClassifications, e := call.Do()
		if e != nil {
			return nil, e
		}

		if _accountClassifications == nil {
			break
		}

		if len(*_accountClassifications) == 0 {
			break
		}

		accountClassifications = append(accountClassifications, *_accountClassifications...)
	}

	return &accountClassifications, nil
}

func (service *Service) GetAccountClassification(id types.Guid) (*AccountClassification, *errortools.Error) {
	url := service.url(fmt.Sprintf("AccountClassifications(guid'%s')", id.String()))

	accountClassificationNew := AccountClassification{}

	e := service.GetSingle(url, &accountClassificationNew)
	if e != nil {
		return nil, e
	}
	return &accountClassificationNew, nil
}

func (service *Service) GetAccountClassificationsCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("AccountClassifications", createdBefore)
}
