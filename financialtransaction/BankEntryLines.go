package exactonline

import (
	"fmt"

	types "github.com/Leapforce-nl/go_types"
	utilities "github.com/Leapforce-nl/go_utilities"
)

// BankEntryLine stores BankEntryLine from exactonline
//
type BankEntryLine struct {
	ID                    types.GUID  `json:"ID"`
	Account               types.GUID  `json:"Account"`
	AccountCode           string      `json:"AccountCode"`
	AccountName           string      `json:"AccountName"`
	AmountDC              float64     `json:"AmountDC"`
	AmountFC              float64     `json:"AmountFC"`
	AmountVATFC           float64     `json:"AmountVATFC"`
	Asset                 types.GUID  `json:"Asset"`
	AssetCode             string      `json:"AssetCode"`
	AssetDescription      string      `json:"AssetDescription"`
	CostCenter            string      `json:"CostCenter"`
	CostCenterDescription string      `json:"CostCenterDescription"`
	CostUnit              string      `json:"CostUnit"`
	CostUnitDescription   string      `json:"CostUnitDescription"`
	Created               *types.Date `json:"Created"`
	Creator               types.GUID  `json:"Creator"`
	CreatorFullName       string      `json:"CreatorFullName"`
	Date                  *types.Date `json:"Date"`
	Description           string      `json:"Description"`
	Division              int32       `json:"Division"`
	Document              types.GUID  `json:"Document"`
	DocumentNumber        int32       `json:"DocumentNumber"`
	DocumentSubject       string      `json:"DocumentSubject"`
	EntryID               types.GUID  `json:"EntryID"`
	EntryNumber           int32       `json:"EntryNumber"`
	ExchangeRate          float64     `json:"ExchangeRate"`
	GLAccount             types.GUID  `json:"GLAccount"`
	GLAccountCode         string      `json:"GLAccountCode"`
	GLAccountDescription  string      `json:"GLAccountDescription"`
	LineNumber            int32       `json:"LineNumber"`
	Modified              *types.Date `json:"Modified"`
	Modifier              types.GUID  `json:"Modifier"`
	ModifierFullName      string      `json:"ModifierFullName"`
	Notes                 string      `json:"Notes"`
	OffsetID              types.GUID  `json:"OffsetID"`
	OurRef                int32       `json:"OurRef"`
	Project               types.GUID  `json:"Project"`
	ProjectCode           string      `json:"ProjectCode"`
	ProjectDescription    string      `json:"ProjectDescription"`
	Quantity              float64     `json:"Quantity"`
	VATCode               string      `json:"VATCode"`
	VATCodeDescription    string      `json:"VATCodeDescription"`
	VATPercentage         float64     `json:"VATPercentage"`
	VATType               string      `json:"VATType"`
}

func (c *Client) GetBankEntryLinesInternal(filter string) (*[]BankEntryLine, error) {
	selectFields := utilities.GetTaggedFieldNames("json", BankEntryLine{})
	urlStr := fmt.Sprintf("%s/financialtransaction/BankEntryLines?$select=%s", c.Http().BaseURL(), selectFields)
	if filter != "" {
		urlStr += fmt.Sprintf("&$filter=%s", filter)
	}
	fmt.Println(urlStr)

	bankEntryLines := []BankEntryLine{}

	for urlStr != "" {
		ac := []BankEntryLine{}

		next, err := c.Http().Get(urlStr, &ac)
		if err != nil {
			fmt.Println("ERROR in GetBankEntryLinesInternal:", err)
			fmt.Println("url:", urlStr)
			return nil, err
		}

		bankEntryLines = append(bankEntryLines, ac...)

		urlStr = next
	}

	return &bankEntryLines, nil
}

func (c *Client) GetBankEntryLines() (*[]BankEntryLine, error) {
	acc, err := c.GetBankEntryLinesInternal("")
	if err != nil {
		return nil, err
	}

	return acc, nil
}