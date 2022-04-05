package exactonline

import (
	"encoding/json"
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// BankEntry stores BankEntry from exactonline
//
type BankEntry struct {
	EntryID                      types.Guid      `json:"EntryID"`
	BankEntryLines               json.RawMessage `json:"BankEntryLines"` //to be implemented when needed
	BankStatementDocument        types.Guid      `json:"BankStatementDocument"`
	BankStatementDocumentNumber  int32           `json:"BankStatementDocumentNumber"`
	BankStatementDocumentSubject string          `json:"BankStatementDocumentSubject"`
	ClosingBalanceFC             float64         `json:"ClosingBalanceFC"`
	Created                      *types.Date     `json:"Created"`
	Currency                     string          `json:"Currency"`
	Division                     int32           `json:"Division"`
	EntryNumber                  int32           `json:"EntryNumber"`
	FinancialPeriod              int16           `json:"FinancialPeriod"`
	FinancialYear                int16           `json:"FinancialYear"`
	JournalCode                  string          `json:"JournalCode"`
	JournalDescription           string          `json:"JournalDescription"`
	Modified                     *types.Date     `json:"Modified"`
	OpeningBalanceFC             float64         `json:"OpeningBalanceFC"`
	Status                       int16           `json:"Status"`
	StatusDescription            string          `json:"StatusDescription"`
}

type GetBankEntriesCall struct {
	modifiedAfter *time.Time
	urlNext       string
	service       *Service
}

func (service *Service) NewGetBankEntriesCall(modifiedAfter *time.Time) *GetBankEntriesCall {
	call := GetBankEntriesCall{}
	call.modifiedAfter = modifiedAfter
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", BankEntry{})
	call.urlNext = service.url(fmt.Sprintf("BankEntries?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetBankEntriesCall) Do() (*[]BankEntry, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	bankEntries := []BankEntry{}

	next, err := call.service.Get(call.urlNext, &bankEntries)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &bankEntries, nil
}

func (service *Service) GetBankEntriesCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("BankEntries", createdBefore)
}
