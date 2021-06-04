package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// FinancialGLClassification stores FinancialGLClassification from exactonline
//
type FinancialGLClassification struct {
	Timestamp                    types.Int64String `json:"Timestamp"`
	ID                           types.GUID        `json:"ID"`
	Abstract                     bool              `json:"Abstract"`
	Balance                      string            `json:"Balance"`
	Code                         string            `json:"Code"`
	Created                      *types.Date       `json:"Created"`
	Creator                      types.GUID        `json:"Creator"`
	CreatorFullName              string            `json:"CreatorFullName"`
	Description                  string            `json:"Description"`
	Division                     int32             `json:"Division"`
	IsTupleSubElement            bool              `json:"IsTupleSubElement"`
	Modified                     *types.Date       `json:"Modified"`
	Modifier                     types.GUID        `json:"Modifier"`
	ModifierFullName             string            `json:"ModifierFullName"`
	Name                         string            `json:"Name"`
	Nillable                     bool              `json:"Nillable"`
	Parent                       types.GUID        `json:"Parent"`
	PeriodType                   string            `json:"PeriodType"`
	SubstitutionGroup            string            `json:"SubstitutionGroup"`
	TaxonomyNamespace            types.GUID        `json:"TaxonomyNamespace"`
	TaxonomyNamespaceDescription string            `json:"TaxonomyNamespaceDescription"`
	Type                         types.GUID        `json:"Type"`
}

type SyncFinancialGLClassificationsCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewSyncFinancialGLClassificationsCall(timestamp *int64) *SyncFinancialGLClassificationsCall {
	selectFields := utilities.GetTaggedTagNames("json", FinancialGLClassification{})
	url := service.url(fmt.Sprintf("Financial/GLClassifications?$select=%s", selectFields))
	if timestamp != nil {
		url = fmt.Sprintf("%s&$filter=Timestamp gt %vL", url, *timestamp)
	}

	return &SyncFinancialGLClassificationsCall{
		service: service,
		urlNext: url,
	}
}

func (call *SyncFinancialGLClassificationsCall) Do() (*[]FinancialGLClassification, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	financialGLClassifications := []FinancialGLClassification{}

	next, err := call.service.Get(call.urlNext, &financialGLClassifications)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &financialGLClassifications, nil
}
