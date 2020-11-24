package exactonline

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
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
	CancellationDate            *string                   `json:"CancellationDate,omitempty"`
	Classification              *string                   `json:"Classification,omitempty"`
	Currency                    *string                   `json:"Currency,omitempty"`
	CustomerPONumber            *string                   `json:"CustomerPONumber,omitempty"`
	Description                 *string                   `json:"Description,omitempty"`
	EndDate                     *string                   `json:"EndDate,omitempty"`
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
	StartDate                   *string                   `json:"StartDate,omitempty"`
	SubscriptionLines           *[]SubscriptionLineUpdate `json:"SubscriptionLines,omitempty"`
	SubscriptionType            *string                   `json:"SubscriptionType,omitempty"`
}

type GetSubscriptionsCall struct {
	urlNext string
	client  *Client
}

type GetSubscriptionsCallParams struct {
	OrderedBy     *types.GUID
	ModifiedAfter *time.Time
}

func (c *Client) NewGetSubscriptionsCall(params GetSubscriptionsCallParams) *GetSubscriptionsCall {
	call := GetSubscriptionsCall{}
	call.client = c

	selectFields := utilities.GetTaggedFieldNames("json", Subscription{})
	call.urlNext = fmt.Sprintf("%s/Subscriptions?$select=%s", c.BaseURL(), selectFields)
	filter := []string{}

	if params.OrderedBy != nil {
		filter = append(filter, fmt.Sprintf("OrderedBy eq guid'%s'", params.OrderedBy.String()))
	}
	if params.ModifiedAfter != nil {
		filter = append(filter, c.DateFilter("Modified", "gt", params.ModifiedAfter, true, "&"))
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

	next, err := call.client.Get(call.urlNext, &subscriptions)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &subscriptions, nil
}

func (c *Client) GetSubscription(entryID types.GUID) (*Subscription, *errortools.Error) {
	url := fmt.Sprintf("%s/Subscriptions(guid'%s')", c.BaseURL(), entryID.String())

	subscription := Subscription{}

	err := c.GetSingle(url, &subscription)
	if err != nil {
		return nil, err
	}
	return &subscription, nil
}

func (c *Client) CreateSubscription(subscription *SubscriptionUpdate) (*Subscription, *errortools.Error) {
	url := fmt.Sprintf("%s/Subscriptions", c.BaseURL())

	b, err := json.Marshal(subscription)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	subscriptionNew := Subscription{}

	e := c.Post(url, bytes.NewBuffer(b), &subscriptionNew)
	if e != nil {
		return nil, e
	}
	return &subscriptionNew, nil
}

func (c *Client) UpdateSubscription(entryID types.GUID, subscription *SubscriptionUpdate) *errortools.Error {
	url := fmt.Sprintf("%s/Subscriptions(guid'%s')", c.BaseURL(), entryID.String())

	b, err := json.Marshal(subscription)
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	e := c.Put(url, bytes.NewBuffer(b))
	if e != nil {
		return e
	}
	return nil
}

func (c *Client) DeleteSubscription(entryID types.GUID) *errortools.Error {
	url := fmt.Sprintf("%s/Subscriptions(guid'%s')", c.BaseURL(), entryID.String())

	err := c.Delete(url)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetSubscriptionsCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return c.GetCount("Subscriptions", createdBefore)
}
