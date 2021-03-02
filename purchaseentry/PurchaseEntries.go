package exactonline

import (
	"encoding/json"
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// PurchaseEntry stores PurchaseEntry from exactonline
//
type PurchaseEntry struct {
	EntryID                     types.GUID      `json:"EntryID"`
	AmountDC                    float64         `json:"AmountDC"`
	AmountFC                    float64         `json:"AmountFC"`
	BatchNumber                 int32           `json:"BatchNumber"`
	Created                     *types.Date     `json:"Created"`
	Creator                     types.GUID      `json:"Creator"`
	CreatorFullName             string          `json:"CreatorFullName"`
	Currency                    string          `json:"Currency"`
	Description                 string          `json:"Description"`
	Division                    int32           `json:"Division"`
	Document                    types.GUID      `json:"Document"`
	DocumentNumber              int32           `json:"DocumentNumber"`
	DocumentSubject             string          `json:"DocumentSubject"`
	DueDate                     *types.Date     `json:"DueDate"`
	EntryDate                   *types.Date     `json:"EntryDate"`
	EntryNumber                 int32           `json:"EntryNumber"`
	ExternalLinkDescription     string          `json:"ExternalLinkDescription"`
	ExternalLinkReference       string          `json:"ExternalLinkReference"`
	GAccountAmountFC            float64         `json:"GAccountAmountFC"`
	InvoiceNumber               int32           `json:"InvoiceNumber"`
	Journal                     string          `json:"Journal"`
	JournalDescription          string          `json:"JournalDescription"`
	Modified                    *types.Date     `json:"Modified"`
	Modifier                    types.GUID      `json:"Modifier"`
	ModifierFullName            string          `json:"ModifierFullName"`
	OrderNumber                 int32           `json:"OrderNumber"`
	PaymentCondition            string          `json:"PaymentCondition"`
	PaymentConditionDescription string          `json:"PaymentConditionDescription"`
	PaymentReference            string          `json:"PaymentReference"`
	ProcessNumber               int32           `json:"ProcessNumber"`
	PurchaseEntryLines          json.RawMessage `json:"PurchaseEntryLines"`
	Rate                        float64         `json:"Rate"`
	ReportingPeriod             int16           `json:"ReportingPeriod"`
	ReportingYear               int16           `json:"ReportingYear"`
	Reversal                    bool            `json:"Reversal"`
	Status                      int16           `json:"Status"`
	StatusDescription           string          `json:"StatusDescription"`
	Supplier                    types.GUID      `json:"Supplier"`
	SupplierName                string          `json:"SupplierName"`
	Type                        int32           `json:"Type"`
	TypeDescription             string          `json:"TypeDescription"`
	VATAmountDC                 float64         `json:"VATAmountDC"`
	VATAmountFC                 float64         `json:"VATAmountFC"`
	YourRef                     string          `json:"YourRef"`
}

type GetPurchaseEntriesCall struct {
	modifiedAfter *time.Time
	urlNext       string
	service       *Service
}

func (service *Service) NewGetPurchaseEntriesCall(modifiedAfter *time.Time) *GetPurchaseEntriesCall {
	call := GetPurchaseEntriesCall{}
	call.modifiedAfter = modifiedAfter
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", PurchaseEntry{})
	call.urlNext = service.url(fmt.Sprintf("PurchaseEntries?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetPurchaseEntriesCall) Do() (*[]PurchaseEntry, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	purchaseEntries := []PurchaseEntry{}

	next, err := call.service.Get(call.urlNext, &purchaseEntries)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &purchaseEntries, nil
}

func (service *Service) GetPurchaseEntriesCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("PurchaseEntries", createdBefore)
}
