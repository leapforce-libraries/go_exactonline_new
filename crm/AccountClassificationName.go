package exactonline

import (
	"fmt"
	types "github.com/leapforce-libraries/go_types"
	"strings"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// AccountClassificationName stores AccountClassificationName from exactonline
type AccountClassificationName struct {
	ID               types.Guid  `json:"ID"`
	Created          *types.Date `json:"Created"`
	Creator          types.Guid  `json:"Creator"`
	CreatorFullName  string      `json:"CreatorFullName"`
	Description      string      `json:"Description"`
	Division         int32       `json:"Division"`
	Modified         *types.Date `json:"Modified"`
	Modifier         types.Guid  `json:"Modifier"`
	ModifierFullName string      `json:"ModifierFullName"`
	SequenceNumber   int32       `json:"SequenceNumber"`
}

type GetAccountClassificationNamesCall struct {
	urlNext string
	service *Service
}

type GetAccountClassificationNamesCallParams struct {
	ModifiedAfter *time.Time
}

func (service *Service) NewGetAccountClassificationNamesCall(params *GetAccountClassificationNamesCallParams) *GetAccountClassificationNamesCall {
	call := GetAccountClassificationNamesCall{}
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", AccountClassificationName{})
	call.urlNext = service.url(fmt.Sprintf("AccountClassificationNames?$select=%s", selectFields))
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

func (call *GetAccountClassificationNamesCall) Do() (*[]AccountClassificationName, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	accountClassificationNames := []AccountClassificationName{}

	next, err := call.service.Get(call.urlNext, &accountClassificationNames)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &accountClassificationNames, nil
}

func (call *GetAccountClassificationNamesCall) DoAll() (*[]AccountClassificationName, *errortools.Error) {
	accountClassificationNames := []AccountClassificationName{}

	for true {
		_accountClassificationNames, e := call.Do()
		if e != nil {
			return nil, e
		}

		if _accountClassificationNames == nil {
			break
		}

		if len(*_accountClassificationNames) == 0 {
			break
		}

		accountClassificationNames = append(accountClassificationNames, *_accountClassificationNames...)
	}

	return &accountClassificationNames, nil
}

func (service *Service) GetAccountClassificationName(id types.Guid) (*AccountClassificationName, *errortools.Error) {
	url := service.url(fmt.Sprintf("AccountClassificationNames(guid'%s')", id.String()))

	accountClassificationNameNew := AccountClassificationName{}

	e := service.GetSingle(url, &accountClassificationNameNew)
	if e != nil {
		return nil, e
	}
	return &accountClassificationNameNew, nil
}

func (service *Service) GetAccountClassificationNamesCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("AccountClassificationNames", createdBefore)
}
