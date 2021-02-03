package exactonline

import (
	"fmt"
	"strings"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// SubscriptionType stores SubscriptionType from exactonline
//
type SubscriptionType struct {
	ID               types.GUID  `json:"ID"`
	Code             string      `json:"Code"`
	Created          *types.Date `json:"Created"`
	Creator          types.GUID  `json:"Creator"`
	CreatorFullName  string      `json:"CreatorFullName"`
	Description      string      `json:"Description"`
	Division         int32       `json:"Division"`
	Modified         *types.Date `json:"Modified"`
	Modifier         types.GUID  `json:"Modifier"`
	ModifierFullName string      `json:"ModifierFullName"`
}

// SubscriptionTypeUpdate stores SubscriptionType values for insert/update
//
type SubscriptionTypeUpdate struct {
	Code        *string `json:"Code,omitempty"`
	Description *string `json:"Description,omitempty"`
}

type GetSubscriptionTypesCall struct {
	urlNext string
	service *Service
}

type GetSubscriptionTypesCallParams struct {
	ModifiedAfter *time.Time
}

func (service *Service) NewGetSubscriptionTypesCall(params *GetSubscriptionTypesCallParams) *GetSubscriptionTypesCall {
	call := GetSubscriptionTypesCall{}
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", SubscriptionType{})
	call.urlNext = service.url(fmt.Sprintf("SubscriptionTypes?$select=%s", selectFields))
	filter := []string{}

	if params != nil {
		if params.ModifiedAfter != nil {
			filter = append(filter, service.DateFilter("Modified", "gt", params.ModifiedAfter, true, "&"))
		}
	}

	if len(filter) > 0 {
		call.urlNext = fmt.Sprintf("%s&$filter=%s", call.urlNext, strings.Join(filter, " AND "))
	}

	return &call
}

func (call *GetSubscriptionTypesCall) Do() (*[]SubscriptionType, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	subscriptionTypes := []SubscriptionType{}

	next, err := call.service.Get(call.urlNext, &subscriptionTypes)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &subscriptionTypes, nil
}

func (call *GetSubscriptionTypesCall) DoAll() (*[]SubscriptionType, *errortools.Error) {
	subscriptionTypes := []SubscriptionType{}

	for true {
		_subscriptionTypes, e := call.Do()
		if e != nil {
			return nil, e
		}

		if _subscriptionTypes == nil {
			break
		}

		if len(*_subscriptionTypes) == 0 {
			break
		}

		subscriptionTypes = append(subscriptionTypes, *_subscriptionTypes...)
	}

	return &subscriptionTypes, nil
}

func (service *Service) CreateSubscriptionType(subscriptionType *SubscriptionTypeUpdate) (*SubscriptionType, *errortools.Error) {
	url := service.url("SubscriptionTypes")

	subscriptionTypeNew := SubscriptionType{}

	e := service.Post(url, subscriptionType, &subscriptionTypeNew)
	if e != nil {
		return nil, e
	}
	return &subscriptionTypeNew, nil
}

func (service *Service) UpdateSubscriptionType(id types.GUID, subscriptionType *SubscriptionTypeUpdate) *errortools.Error {
	url := service.url(fmt.Sprintf("SubscriptionTypes(guid'%s')", id.String()))

	e := service.Put(url, subscriptionType)
	if e != nil {
		return e
	}
	return nil
}

func (service *Service) DeleteSubscriptionType(id types.GUID) *errortools.Error {
	url := service.url(fmt.Sprintf("SubscriptionTypes(guid'%s')", id.String()))

	err := service.Delete(url)
	if err != nil {
		return err
	}
	return nil
}

func (service *Service) GetSubscriptionTypesCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("SubscriptionTypes", createdBefore)
}