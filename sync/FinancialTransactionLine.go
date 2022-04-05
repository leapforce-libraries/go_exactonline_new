package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// FinancialTransactionLine stores FinancialTransactionLine from exactonline
//
type FinancialTransactionLine struct {
	Timestamp                 types.Int64String `json:"Timestamp"`
	Account                   types.Guid        `json:"Account"`
	AccountCode               string            `json:"AccountCode"`
	AccountName               string            `json:"AccountName"`
	AmountDC                  float64           `json:"AmountDC"`
	AmountFC                  float64           `json:"AmountFC"`
	AmountVATBaseFC           float64           `json:"AmountVATBaseFC"`
	AmountVATFC               float64           `json:"AmountVATFC"`
	Asset                     types.Guid        `json:"Asset"`
	AssetCode                 string            `json:"AssetCode"`
	AssetDescription          string            `json:"AssetDescription"`
	CostCenter                string            `json:"CostCenter"`
	CostCenterDescription     string            `json:"CostCenterDescription"`
	CostUnit                  string            `json:"CostUnit"`
	CostUnitDescription       string            `json:"CostUnitDescription"`
	Created                   *types.Date       `json:"Created"`
	Creator                   types.Guid        `json:"Creator"`
	CreatorFullName           string            `json:"CreatorFullName"`
	Currency                  string            `json:"Currency"`
	Date                      *types.Date       `json:"Date"`
	Description               string            `json:"Description"`
	Division                  int64             `json:"Division"`
	Document                  types.Guid        `json:"Document"`
	DocumentNumber            int64             `json:"DocumentNumber"`
	DocumentSubject           string            `json:"DocumentSubject"`
	DueDate                   *types.Date       `json:"DueDate"`
	EntryID                   types.Guid        `json:"EntryID"`
	EntryNumber               int64             `json:"EntryNumber"`
	ExchangeRate              float64           `json:"ExchangeRate"`
	ExtraDutyAmountFC         float64           `json:"ExtraDutyAmountFC"`
	ExtraDutyPercentage       float64           `json:"ExtraDutyPercentage"`
	FinancialPeriod           int64             `json:"FinancialPeriod"`
	FinancialYear             int64             `json:"FinancialYear"`
	GLAccount                 types.Guid        `json:"GLAccount"`
	GLAccountCode             string            `json:"GLAccountCode"`
	GLAccountDescription      string            `json:"GLAccountDescription"`
	ID                        types.Guid        `json:"ID"`
	InvoiceNumber             int64             `json:"InvoiceNumber"`
	Item                      types.Guid        `json:"Item"`
	ItemCode                  string            `json:"ItemCode"`
	ItemDescription           string            `json:"ItemDescription"`
	JournalCode               string            `json:"JournalCode"`
	JournalDescription        string            `json:"JournalDescription"`
	LineNumber                int64             `json:"LineNumber"`
	LineType                  int64             `json:"LineType"`
	Modified                  *types.Date       `json:"Modified"`
	Modifier                  types.Guid        `json:"Modifier"`
	ModifierFullName          string            `json:"ModifierFullName"`
	Notes                     string            `json:"Notes"`
	OffsetID                  types.Guid        `json:"OffsetID"`
	OrderNumber               int64             `json:"OrderNumber"`
	PaymentDiscountAmount     float64           `json:"PaymentDiscountAmount"`
	PaymentReference          string            `json:"PaymentReference"`
	Project                   types.Guid        `json:"Project"`
	ProjectCode               string            `json:"ProjectCode"`
	ProjectDescription        string            `json:"ProjectDescription"`
	Quantity                  float64           `json:"Quantity"`
	SerialNumber              string            `json:"SerialNumber"`
	Status                    int64             `json:"Status"`
	Subscription              types.Guid        `json:"Subscription"`
	SubscriptionDescription   string            `json:"SubscriptionDescription"`
	TrackingNumber            string            `json:"TrackingNumber"`
	TrackingNumberDescription string            `json:"TrackingNumberDescription"`
	Type                      int64             `json:"Type"`
	VATCode                   string            `json:"VATCode"`
	VATCodeDescription        string            `json:"VATCodeDescription"`
	VATPercentage             float64           `json:"VATPercentage"`
	VATType                   string            `json:"VATType"`
	YourRef                   string            `json:"YourRef"`
}

type SyncFinancialTransactionLinesCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewSyncFinancialTransactionLinesCall(timestamp *int64) *SyncFinancialTransactionLinesCall {
	selectFields := utilities.GetTaggedTagNames("json", FinancialTransactionLine{})
	url := service.url(fmt.Sprintf("Financial/TransactionLines?$select=%s", selectFields))
	if timestamp != nil {
		url = fmt.Sprintf("%s&$filter=Timestamp gt %vL", url, *timestamp)
	}

	return &SyncFinancialTransactionLinesCall{
		service: service,
		urlNext: url,
	}
}

func (call *SyncFinancialTransactionLinesCall) Do() (*[]FinancialTransactionLine, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	financialTransactionLines := []FinancialTransactionLine{}

	next, err := call.service.Get(call.urlNext, &financialTransactionLines)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &financialTransactionLines, nil
}
