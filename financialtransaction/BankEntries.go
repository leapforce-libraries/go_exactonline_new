package exactonline

import (
	"encoding/json"
	"fmt"
	"time"

	types "github.com/Leapforce-nl/go_types"
	utilities "github.com/Leapforce-nl/go_utilities"
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

func (c *Client) GetBankEntriesInternal(filter string) (*[]BankEntry, error) {
	selectFields := utilities.GetTaggedFieldNames("json", BankEntry{})
	urlStr := fmt.Sprintf("%s/financialtransaction/BankEntries?$select=%s", c.Http().BaseURL(), selectFields)
	if filter != "" {
		urlStr += fmt.Sprintf("&$filter=%s", filter)
	}
	//fmt.Println(urlStr)

	bankEntries := []BankEntry{}

	for urlStr != "" {
		ac := []BankEntry{}

		next, err := c.Http().Get(urlStr, &ac)
		if err != nil {
			fmt.Println("ERROR in GetBankEntriesInternal:", err)
			fmt.Println("url:", urlStr)
			return nil, err
		}

		bankEntries = append(bankEntries, ac...)

		urlStr = next
	}

	return &bankEntries, nil
}

func (c *Client) GetBankEntries(modifiedAfter *time.Time) (*[]BankEntry, error) {
	acc, err := c.GetBankEntriesInternal(c.Http().DateFilter("Modified", "gt", modifiedAfter, false, ""))
	if err != nil {
		return nil, err
	}

	return acc, nil
}

func (c *Client) GetBankEntriesCount(createdBefore *time.Time) (int64, error) {
	return c.Http().GetCount("financialtransaction/BankEntries", createdBefore)
}
