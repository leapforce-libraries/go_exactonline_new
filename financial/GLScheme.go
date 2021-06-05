package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// GLScheme stores GLScheme from exactonline
//
type GLScheme struct {
	ID               types.GUID  `json:"ID"`
	Code             string      `json:"Code"`
	Created          *types.Date `json:"Created"`
	Creator          types.GUID  `json:"Creator"`
	CreatorFullName  string      `json:"CreatorFullName"`
	Description      string      `json:"Description"`
	Division         int32       `json:"Division"`
	Main             byte        `json:"Main"`
	Modified         *types.Date `json:"Modified"`
	Modifier         types.GUID  `json:"Modifier"`
	ModifierFullName string      `json:"ModifierFullName"`
	TargetNamespace  string      `json:"TargetNamespace"`
}

type GetGLSchemesCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewGetGLSchemesCall(modifiedAfter *time.Time) *GetGLSchemesCall {
	call := GetGLSchemesCall{}
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", GLScheme{})
	call.urlNext = service.url(fmt.Sprintf("GLScheme?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetGLSchemesCall) Do() (*[]GLScheme, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	transactionLines := []GLScheme{}

	next, err := call.service.Get(call.urlNext, &transactionLines)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &transactionLines, nil
}

func (service *Service) GetGLSchemesCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("GLSchemes", createdBefore)
}
