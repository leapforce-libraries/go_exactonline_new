package exactonline

import (
	"fmt"
	"strings"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// Subscription stores Subscription from exactonline
//
type Subscription struct {
	EntryID                        types.GUID  `json:"EntryID"`
	BlockEntry                     bool        `json:"BlockEntry"`
	CancellationDate               *types.Date `json:"CancellationDate"`
	Classification                 types.GUID  `json:"Classification"`
	ClassificationCode             string      `json:"ClassificationCode"`
	ClassificationDescription      string      `json:"ClassificationDescription"`
	Created                        *types.Date `json:"Created"`
	Creator                        types.GUID  `json:"Creator"`
	CreatorFullName                string      `json:"CreatorFullName"`
	Currency                       string      `json:"Currency"`
	CustomerPONumber               string      `json:"CustomerPONumber"`
	Description                    string      `json:"Description"`
	Division                       int32       `json:"Division"`
	EndDate                        *types.Date `json:"EndDate"`
	InvoiceDay                     byte        `json:"InvoiceDay"`
	InvoicedTo                     *types.Date `json:"InvoicedTo"`
	InvoiceTo                      types.GUID  `json:"InvoiceTo"`
	InvoiceToContactPerson         types.GUID  `json:"InvoiceToContactPerson"`
	InvoiceToContactPersonFullName string      `json:"InvoiceToContactPersonFullName"`
	InvoiceToName                  string      `json:"InvoiceToName"`
	InvoicingStartDate             *types.Date `json:"InvoicingStartDate"`
	Modified                       *types.Date `json:"Modified"`
	Modifier                       types.GUID  `json:"Modifier"`
	ModifierFullName               string      `json:"ModifierFullName"`
	Notes                          string      `json:"Notes"`
	Number                         int32       `json:"Number"`
	OrderedBy                      types.GUID  `json:"OrderedBy"`
	OrderedByContactPerson         types.GUID  `json:"OrderedByContactPerson"`
	OrderedByContactPersonFullName string      `json:"OrderedByContactPersonFullName"`
	OrderedByName                  string      `json:"OrderedByName"`
	PaymentCondition               string      `json:"PaymentCondition"`
	PaymentConditionDescription    string      `json:"PaymentConditionDescription"`
	Printed                        bool        `json:"Printed"`
	ReasonCancelled                types.GUID  `json:"ReasonCancelled"`
	ReasonCancelledCode            string      `json:"ReasonCancelledCode"`
	ReasonCancelledDescription     string      `json:"ReasonCancelledDescription"`
	StartDate                      *types.Date `json:"StartDate"`
	SubscriptionType               types.GUID  `json:"SubscriptionType"`
	SubscriptionTypeCode           string      `json:"SubscriptionTypeCode"`
	SubscriptionTypeDescription    string      `json:"SubscriptionTypeDescription"`
}

// SubscriptionUpdate stores Subscription value to insert/update
//
type SubscriptionUpdate struct {
	BlockEntry                  *bool                     `json:"BlockEntry,omitempty"`
	CancellationDate            *types.Date               `json:"CancellationDate,omitempty"`
	Classification              *string                   `json:"Classification,omitempty"`
	Currency                    *string                   `json:"Currency,omitempty"`
	CustomerPONumber            *string                   `json:"CustomerPONumber,omitempty"`
	Description                 *string                   `json:"Description,omitempty"`
	EndDate                     *types.Date               `json:"EndDate,omitempty"`
	InvoiceDay                  *byte                     `json:"InvoiceDay,omitempty"`
	InvoicedTo                  *string                   `json:"InvoicedTo,omitempty"`
	InvoiceTo                   *string                   `json:"InvoiceTo,omitempty"`
	InvoiceToContactPerson      *string                   `json:"InvoiceToContactPerson,omitempty"`
	InvoicingStartDate          *string                   `json:"InvoicingStartDate,omitempty"`
	Notes                       *string                   `json:"Notes,omitempty"`
	Number                      *int32                    `json:"Number,omitempty"`
	OrderedBy                   *string                   `json:"OrderedBy,omitempty"`
	OrderedByContactPerson      *string                   `json:"OrderedByContactPerson,omitempty"`
	PaymentCondition            *string                   `json:"PaymentCondition,omitempty"`
	PaymentConditionDescription *string                   `json:"PaymentConditionDescription,omitempty"`
	Printed                     *bool                     `json:"Printed,omitempty"`
	ReasonCancelled             *string                   `json:"ReasonCancelled,omitempty"`
	StartDate                   *types.Date               `json:"StartDate,omitempty"`
	SubscriptionLines           *[]SubscriptionLineUpdate `json:"SubscriptionLines,omitempty"`
	SubscriptionType            *string                   `json:"SubscriptionType,omitempty"`
}

type GetSubscriptionsCall struct {
	urlNext string
	service *Service
}

type GetSubscriptionsCallParams struct {
	OrderedBy     *types.GUID
	ModifiedAfter *time.Time
}

func (service *Service) NewGetSubscriptionsCall(params *GetSubscriptionsCallParams) *GetSubscriptionsCall {
	call := GetSubscriptionsCall{}
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", Subscription{})
	call.urlNext = service.url(fmt.Sprintf("Subscriptions?$select=%s", selectFields))
	filter := []string{}

	if params != nil {
		if params.OrderedBy != nil {
			filter = append(filter, fmt.Sprintf("OrderedBy eq guid'%s'", params.OrderedBy.String()))
		}
		if params.ModifiedAfter != nil {
			filter = append(filter, service.DateFilter("Modified", "gt", params.ModifiedAfter, true, "&"))
		}
	}

	if len(filter) > 0 {
		call.urlNext = fmt.Sprintf("%s&$filter=%s", call.urlNext, strings.Join(filter, " AND "))
	}

	return &call
}

func (call *GetSubscriptionsCall) Do() (*[]Subscription, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	subscriptions := []Subscription{}

	next, err := call.service.Get(call.urlNext, &subscriptions)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &subscriptions, nil
}

func (call *GetSubscriptionsCall) DoAll() (*[]Subscription, *errortools.Error) {
	subscriptions := []Subscription{}

	for true {
		_subscriptions, e := call.Do()
		if e != nil {
			return nil, e
		}

		if _subscriptions == nil {
			break
		}

		if len(*_subscriptions) == 0 {
			break
		}

		subscriptions = append(subscriptions, *_subscriptions...)
	}

	return &subscriptions, nil
}

func (service *Service) GetSubscription(entryID types.GUID) (*Subscription, *errortools.Error) {
	url := service.url(fmt.Sprintf("Subscriptions(guid'%s')", entryID.String()))

	subscription := Subscription{}

	err := service.GetSingle(url, &subscription)
	if err != nil {
		return nil, err
	}
	return &subscription, nil
}

func (service *Service) CreateSubscription(subscription *SubscriptionUpdate) (*Subscription, *errortools.Error) {
	url := service.url("Subscriptions")

	subscriptionNew := Subscription{}

	e := service.Post(url, subscription, &subscriptionNew)
	if e != nil {
		return nil, e
	}
	return &subscriptionNew, nil
}

func (service *Service) UpdateSubscription(entryID types.GUID, subscription *SubscriptionUpdate, returnUpdated bool) (*Subscription, *errortools.Error) {
	requestConfig := go_http.RequestConfig{
		URL:       service.url(fmt.Sprintf("Subscriptions(guid'%s')", entryID.String())),
		BodyModel: subscription,
	}

	e := service.Put(&requestConfig)
	if e != nil {
		return nil, e
	}

	if !returnUpdated {
		return nil, nil
	}

	subscriptionUpdated, e := service.GetSubscription(entryID)
	if e != nil {
		return nil, e
	}

	return subscriptionUpdated, nil
}

func (service *Service) DeleteSubscription(entryID types.GUID) *errortools.Error {
	url := service.url(fmt.Sprintf("Subscriptions(guid'%s')", entryID.String()))

	err := service.Delete(url)
	if err != nil {
		return err
	}
	return nil
}

func (service *Service) GetSubscriptionsCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("Subscriptions", createdBefore)
}
