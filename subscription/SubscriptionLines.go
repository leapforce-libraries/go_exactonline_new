package exactonline

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// SubscriptionLine stores SubscriptionLine from exactonline
//
type SubscriptionLine struct {
	ID                  types.GUID  `json:"ID"`
	AmountDC            float64     `json:"AmountDC"`
	AmountFC            float64     `json:"AmountFC"`
	Costcenter          string      `json:"Costcenter"`
	Costunit            string      `json:"Costunit"`
	Description         string      `json:"Description"`
	Discount            float64     `json:"Discount"`
	Division            int32       `json:"Division"`
	EntryID             types.GUID  `json:"EntryID"`
	FromDate            *types.Date `json:"FromDate"`
	Item                types.GUID  `json:"Item"`
	ItemDescription     string      `json:"ItemDescription"`
	LineNumber          int32       `json:"LineNumber"`
	LineType            int16       `json:"LineType"`
	LineTypeDescription string      `json:"LineTypeDescription"`
	Modified            *types.Date `json:"Modified"`
	NetPrice            float64     `json:"NetPrice"`
	Notes               string      `json:"Notes"`
	Quantity            float64     `json:"Quantity"`
	ToDate              *types.Date `json:"ToDate"`
	UnitCode            string      `json:"UnitCode"`
	UnitDescription     string      `json:"UnitDescription"`
	UnitPrice           float64     `json:"UnitPrice"`
	VATAmountFC         float64     `json:"VATAmountFC"`
	VATCode             string      `json:"VATCode"`
	VATCodeDescription  string      `json:"VATCodeDescription"`
}

// SubscriptionLineUpdate stores SubscriptionLine values for insert/update
//
type SubscriptionLineUpdate struct {
	AmountDC            *float64 `json:"AmountDC,omitempty"`
	AmountFC            *float64 `json:"AmountFC,omitempty"`
	Costcenter          *string  `json:"Costcenter,omitempty"`
	Costunit            *string  `json:"Costunit,omitempty"`
	Description         *string  `json:"Description,omitempty"`
	Discount            *float64 `json:"Discount,omitempty"`
	EntryID             *string  `json:"EntryID,omitempty"`
	FromDate            *string  `json:"FromDate,omitempty"`
	Item                *string  `json:"Item,omitempty"`
	LineNumber          *int32   `json:"LineNumber,omitempty"`
	LineType            *int16   `json:"LineType,omitempty"`
	LineTypeDescription *string  `json:"LineTypeDescription,omitempty"`
	NetPrice            *float64 `json:"NetPrice,omitempty"`
	Notes               *string  `json:"Notes,omitempty"`
	Quantity            *float64 `json:"Quantity,omitempty"`
	ToDate              *string  `json:"ToDate,omitempty"`
	UnitCode            *string  `json:"UnitCode,omitempty"`
	UnitDescription     *string  `json:"UnitDescription,omitempty"`
	UnitPrice           *float64 `json:"UnitPrice,omitempty"`
	VATAmountFC         *float64 `json:"VATAmountFC,omitempty"`
	VATCode             *string  `json:"VATCode,omitempty"`
	VATCodeDescription  *string  `json:"VATCodeDescription,omitempty"`
}

type GetSubscriptionLinesCall struct {
	urlNext string
	client  *Client
}

type GetSubscriptionLinesCallParams struct {
	EntryID       *types.GUID
	ModifiedAfter *time.Time
}

func (c *Client) NewGetSubscriptionLinesCall(params GetSubscriptionLinesCallParams) *GetSubscriptionLinesCall {
	call := GetSubscriptionLinesCall{}
	call.client = c

	selectFields := utilities.GetTaggedFieldNames("json", SubscriptionLine{})
	call.urlNext = fmt.Sprintf("%s/SubscriptionLines?$select=%s", c.BaseURL(), selectFields)

	filter := []string{}

	if params.EntryID != nil {
		filter = append(filter, fmt.Sprintf("EntryID eq guid'%s'", params.EntryID.String()))
	}
	if params.ModifiedAfter != nil {
		filter = append(filter, c.DateFilter("Modified", "gt", params.ModifiedAfter, true, "&"))
	}

	if len(filter) > 0 {
		call.urlNext = fmt.Sprintf("%s&$filter=%s", call.urlNext, strings.Join(filter, " AND "))
	}

	return &call
}

func (call *GetSubscriptionLinesCall) Do() (*[]SubscriptionLine, error) {
	if call.urlNext == "" {
		return nil, nil
	}

	subscriptionLines := []SubscriptionLine{}

	next, err := call.client.Get(call.urlNext, &subscriptionLines)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &subscriptionLines, nil
}

func (c *Client) CreateSubscriptionLine(subscriptionLine *SubscriptionLineUpdate) (*SubscriptionLine, error) {
	url := fmt.Sprintf("%s/SubscriptionLines", c.BaseURL())

	b, err := json.Marshal(subscriptionLine)
	if err != nil {
		return nil, err
	}

	subscriptionLineNew := SubscriptionLine{}

	err = c.Post(url, bytes.NewBuffer(b), &subscriptionLineNew)
	if err != nil {
		return nil, err
	}
	return &subscriptionLineNew, nil
}

func (c *Client) UpdateSubscriptionLine(id types.GUID, subscriptionLine *SubscriptionLineUpdate) error {
	url := fmt.Sprintf("%s/SubscriptionLines(guid'%s')", c.BaseURL(), id.String())

	b, err := json.Marshal(subscriptionLine)
	if err != nil {
		return err
	}

	err = c.Put(url, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteSubscriptionLine(id types.GUID) error {
	url := fmt.Sprintf("%s/SubscriptionLines(guid'%s')", c.BaseURL(), id.String())

	err := c.Delete(url)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetSubscriptionLinesCount(createdBefore *time.Time) (int64, error) {
	return c.GetCount("SubscriptionLines", createdBefore)
}
