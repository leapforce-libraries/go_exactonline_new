package exactonline

import (
	"fmt"
	"strings"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// ItemAssortmentProperty stores ItemAssortmentProperty from exactonline
//
type ItemAssortmentProperty struct {
	ID                 types.GUID `json:"ID "`
	Code               string     `json:"Code"`
	Description        string     `json:"Description"`
	Division           int32      `json:"Division"`
	ItemAssortmentCode int32      `json:"ItemAssortmentCode"`
}

type GetItemAssortmentPropertiesCall struct {
	urlNext string
	service *Service
}

type GetItemAssortmentPropertiesCallParams struct {
	ModifiedAfter *time.Time
}

func (service *Service) NewGetItemAssortmentPropertiesCall(params *GetItemAssortmentPropertiesCallParams) *GetItemAssortmentPropertiesCall {
	call := GetItemAssortmentPropertiesCall{}
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", ItemAssortmentProperty{})
	call.urlNext = service.url(fmt.Sprintf("ItemAssortmentProperty?$select=%s", selectFields))
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

func (call *GetItemAssortmentPropertiesCall) Do() (*[]ItemAssortmentProperty, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	items := []ItemAssortmentProperty{}

	next, err := call.service.Get(call.urlNext, &items)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &items, nil
}

func (call *GetItemAssortmentPropertiesCall) DoAll() (*[]ItemAssortmentProperty, *errortools.Error) {
	items := []ItemAssortmentProperty{}

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

func (service *Service) GetItemAssortmentPropertiesCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("ItemAssortmentProperties", createdBefore)
}
