package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

type EntityType int64

const (
	EntityTypeTransactionLines     EntityType = 1
	EntityTypeAccounts             EntityType = 2
	EntityTypeAddresses            EntityType = 3
	EntityTypeAttachments          EntityType = 4
	EntityTypeContacts             EntityType = 5
	EntityTypeDocuments            EntityType = 6
	EntityTypeGLAccounts           EntityType = 7
	EntityTypeSalesItemPrices      EntityType = 8
	EntityTypeItems                EntityType = 9
	EntityTypePaymentTerms         EntityType = 10
	EntityTypeQuotations           EntityType = 11
	EntityTypeSalesOrders          EntityType = 12
	EntityTypeSalesInvoices        EntityType = 13
	EntityTypeTimeCostTransactions EntityType = 14
	EntityTypeStockPositions       EntityType = 15
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

type SyncDeletedCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewSyncDeletedCall(timestamp *int64) *SyncDeletedCall {
	selectFields := utilities.GetTaggedTagNames("json", Deleted{})
	url := service.url(fmt.Sprintf("Deleted?$select=%s", selectFields))
	if timestamp != nil {
		url = fmt.Sprintf("%s&$filter=Timestamp gt %vL", url, *timestamp)
	}

	return &SyncDeletedCall{
		service: service,
		urlNext: url,
	}
}

func (call *SyncDeletedCall) Do() (*[]Deleted, *errortools.Error) {
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
