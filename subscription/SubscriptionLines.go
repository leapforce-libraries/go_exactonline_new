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
	AmountDC            *float64    `json:"AmountDC,omitempty"`
	AmountFC            *float64    `json:"AmountFC,omitempty"`
	Costcenter          *string     `json:"Costcenter,omitempty"`
	Costunit            *string     `json:"Costunit,omitempty"`
	Description         *string     `json:"Description,omitempty"`
	Discount            *float64    `json:"Discount,omitempty"`
	EntryID             *string     `json:"EntryID,omitempty"`
	FromDate            *types.Date `json:"FromDate,omitempty"`
	Item                *string     `json:"Item,omitempty"`
	LineNumber          *int32      `json:"LineNumber,omitempty"`
	LineType            *int16      `json:"LineType,omitempty"`
	LineTypeDescription *string     `json:"LineTypeDescription,omitempty"`
	NetPrice            *float64    `json:"NetPrice,omitempty"`
	Notes               *string     `json:"Notes,omitempty"`
	Quantity            *float64    `json:"Quantity,omitempty"`
	ToDate              *types.Date `json:"ToDate,omitempty"`
	UnitCode            *string     `json:"UnitCode,omitempty"`
	UnitDescription     *string     `json:"UnitDescription,omitempty"`
	UnitPrice           *float64    `json:"UnitPrice,omitempty"`
	VATAmountFC         *float64    `json:"VATAmountFC,omitempty"`
	VATCode             *string     `json:"VATCode,omitempty"`
	VATCodeDescription  *string     `json:"VATCodeDescription,omitempty"`
}

type GetSubscriptionLinesCall struct {
	urlNext string
	service *Service
}

type GetSubscriptionLinesCallParams struct {
	EntryID       *types.GUID
	ModifiedAfter *time.Time
}

func (service *Service) NewGetSubscriptionLinesCall(params *GetSubscriptionLinesCallParams) *GetSubscriptionLinesCall {
	call := GetSubscriptionLinesCall{}
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", SubscriptionLine{})
	call.urlNext = service.url(fmt.Sprintf("SubscriptionLines?$select=%s", selectFields))

	filter := []string{}

	if params != nil {
		if params.EntryID != nil {
			filter = append(filter, fmt.Sprintf("EntryID eq guid'%s'", params.EntryID.String()))
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

func (call *GetSubscriptionLinesCall) Do() (*[]SubscriptionLine, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	subscriptionLines := []SubscriptionLine{}

	next, err := call.service.Get(call.urlNext, &subscriptionLines)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &subscriptionLines, nil
}

func (call *GetSubscriptionLinesCall) DoAll() (*[]SubscriptionLine, *errortools.Error) {
	subscriptionLines := []SubscriptionLine{}

	for true {
		_subscriptionLines, e := call.Do()
		if e != nil {
			return nil, e
		}

		if _subscriptionLines == nil {
			break
		}

		if len(*_subscriptionLines) == 0 {
			break
		}

		subscriptionLines = append(subscriptionLines, *_subscriptionLines...)
	}

	return &subscriptionLines, nil
}

func (service *Service) GetSubscriptionLine(id types.GUID) (*SubscriptionLine, *errortools.Error) {
	url := service.url(fmt.Sprintf("SubscriptionLines(guid'%s')", id.String()))

	subscriptionLine := SubscriptionLine{}

	err := service.GetSingle(url, &subscriptionLine)
	if err != nil {
		return nil, err
	}
	return &subscriptionLine, nil
}

func (service *Service) CreateSubscriptionLine(subscriptionLine *SubscriptionLineUpdate) (*SubscriptionLine, *errortools.Error) {
	url := service.url("SubscriptionLines")

	subscriptionLineNew := SubscriptionLine{}

	e := service.Post(url, subscriptionLine, &subscriptionLineNew)
	if e != nil {
		return nil, e
	}
	return &subscriptionLineNew, nil
}

func (service *Service) UpdateSubscriptionLine(id types.GUID, subscriptionLine *SubscriptionLineUpdate, returnUpdated bool) (*SubscriptionLine, *errortools.Error) {
	requestConfig := go_http.RequestConfig{
		URL:       service.url(fmt.Sprintf("SubscriptionLines(guid'%s')", id.String())),
		BodyModel: subscriptionLine,
	}

	e := service.Put(&requestConfig)
	if e != nil {
		return nil, e
	}

	if !returnUpdated {
		return nil, nil
	}

	subscriptionLineUpdated, e := service.GetSubscriptionLine(id)
	if e != nil {
		return nil, e
	}

	return subscriptionLineUpdated, nil
}

func (service *Service) DeleteSubscriptionLine(id types.GUID) *errortools.Error {
	url := service.url(fmt.Sprintf("SubscriptionLines(guid'%s')", id.String()))

	err := service.Delete(url)
	if err != nil {
		return err
	}
	return nil
}

func (service *Service) GetSubscriptionLinesCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("SubscriptionLines", createdBefore)
}
