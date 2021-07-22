package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
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

type GetBankEntryLinesCall struct {
	modifiedAfter *time.Time
	urlNext       string
	service       *Service
}

func (service *Service) NewGetBankEntryLinesCall(modifiedAfter *time.Time) *GetBankEntryLinesCall {
	call := GetBankEntryLinesCall{}
	call.modifiedAfter = modifiedAfter
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", BankEntryLine{})
	call.urlNext = service.url(fmt.Sprintf("BankEntryLines?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetBankEntryLinesCall) Do() (*[]BankEntryLine, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	bankEntryLines := []BankEntryLine{}

	next, err := call.service.Get(call.urlNext, &bankEntryLines)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &bankEntryLines, nil
}

func (service *Service) GetBankEntryLinesCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("BankEntryLines", createdBefore)
}
