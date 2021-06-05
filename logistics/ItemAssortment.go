package exactonline

import (
	"fmt"
	"strings"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// ItemAssortment stores ItemAssortment from exactonline
//
type ItemAssortment struct {
	ID          types.GUID `json:"ID"`
	Code        int32      `json:"Code"`
	Description string     `json:"Description"`
	Division    int32      `json:"Division"`
	//Properties  []ItemAssortmentProperty `json:"Properties"`
}

type GetItemAssortmentsCall struct {
	urlNext string
	service *Service
}

type GetItemAssortmentsCallParams struct {
	ModifiedAfter *time.Time
}

func (service *Service) NewGetItemAssortmentsCall(params *GetItemAssortmentsCallParams) *GetItemAssortmentsCall {
	call := GetItemAssortmentsCall{}
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", ItemAssortment{})
	call.urlNext = service.url(fmt.Sprintf("ItemAssortments?$select=%s", selectFields))
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

func (call *GetItemAssortmentsCall) Do() (*[]ItemAssortment, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	items := []ItemAssortment{}

	next, err := call.service.Get(call.urlNext, &items)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &items, nil
}

func (call *GetItemAssortmentsCall) DoAll() (*[]ItemAssortment, *errortools.Error) {
	items := []ItemAssortment{}

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

func (service *Service) GetItemAssortmentsCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("ItemAssortments", createdBefore)
}
