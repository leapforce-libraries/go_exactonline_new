package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// Deleted stores Deleted from exactonline
//
type Deleted struct {
	Timestamp  types.Int64String `json:"Timestamp"`
	Division   int32             `json:"Division"`
	EntityKey  types.GUID        `json:"EntityKey"`
	EntityType int32             `json:"EntityType"`
	ID         types.GUID        `json:"ID"`
}

type SyncDeletedsCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewSyncDeletedsCall(timestamp *int64) *SyncDeletedsCall {
	call := SyncDeletedsCall{}
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", Deleted{})
	url := service.url(fmt.Sprintf("Deleted?$select=%s", selectFields))
	if timestamp != nil {
		url = fmt.Sprintf("%s&$filter=Timestamp eq %vL", url, *timestamp)
	}
	call.urlNext = url

	return &call
}

func (call *SyncDeletedsCall) Do() (*[]Deleted, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	deleteds := []Deleted{}

	next, err := call.service.Get(call.urlNext, &deleteds)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &deleteds, nil
}
