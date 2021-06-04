package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

type EntityType int64

const (
	EntityTypeTransactionLines     int64 = 1
	EntityTypeAccounts             int64 = 2
	EntityTypeAddresses            int64 = 3
	EntityTypeAttachments          int64 = 4
	EntityTypeContacts             int64 = 5
	EntityTypeDocuments            int64 = 6
	EntityTypeGLAccounts           int64 = 7
	EntityTypeSalesItemPrices      int64 = 8
	EntityTypeItems                int64 = 9
	EntityTypePaymentTerms         int64 = 10
	EntityTypeQuotations           int64 = 11
	EntityTypeSalesOrders          int64 = 12
	EntityTypeSalesInvoices        int64 = 13
	EntityTypeTimeCostTransactions int64 = 14
	EntityTypeStockPositions       int64 = 15
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
