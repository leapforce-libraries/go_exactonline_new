package exactonline

import (
	"encoding/json"
	"fmt"
	"time"

	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// BankEntry stores BankEntry from exactonline
//
type BankEntry struct {
	EntryID                      types.GUID      `json:"EntryID"`
	BankEntryLines               json.RawMessage `json:"BankEntryLines"` //to be implemented when needed
	BankStatementDocument        types.GUID      `json:"BankStatementDocument"`
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
	client        *Client
}

func (c *Client) NewGetBankEntriesCall(modifiedAfter *time.Time) *GetBankEntriesCall {
	call := GetBankEntriesCall{}
	call.modifiedAfter = modifiedAfter
	call.client = c

	selectFields := utilities.GetTaggedFieldNames("json", BankEntry{})
	call.urlNext = fmt.Sprintf("%s/BankEntries?$select=%s", c.BaseURL(), selectFields)
	if modifiedAfter != nil {
		call.urlNext += c.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetBankEntriesCall) Do() (*[]BankEntry, error) {
	if call.urlNext == "" {
		return nil, nil
	}

	bankEntries := []BankEntry{}

	next, err := call.client.Get(call.urlNext, &bankEntries)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &bankEntries, nil
}

func (c *Client) GetBankEntriesCount(createdBefore *time.Time) (int64, error) {
	return c.GetCount("BankEntries", createdBefore)
}
