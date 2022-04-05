package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// Receivable stores Receivable from exactonline
//
type Receivable struct {
	ID                            types.Guid  `json:"ID"`
	Account                       types.Guid  `json:"Account"`
	AccountBankAccountID          types.Guid  `json:"AccountBankAccountID"`
	AccountBankAccountNumber      string      `json:"AccountBankAccountNumber"`
	AccountCode                   string      `json:"AccountCode"`
	AccountContact                types.Guid  `json:"AccountContact"`
	AccountContactName            string      `json:"AccountContactName"`
	AccountCountry                string      `json:"AccountCountry"`
	AccountName                   string      `json:"AccountName"`
	AmountDC                      float64     `json:"AmountDC"`
	AmountDiscountDC              float64     `json:"AmountDiscountDC"`
	AmountDiscountFC              float64     `json:"AmountDiscountFC"`
	AmountFC                      float64     `json:"AmountFC"`
	BankAccountID                 types.Guid  `json:"BankAccountID"`
	BankAccountNumber             string      `json:"BankAccountNumber"`
	CashflowTransactionBatchCode  string      `json:"CashflowTransactionBatchCode"`
	Created                       *types.Date `json:"Created"`
	Creator                       types.Guid  `json:"Creator"`
	CreatorFullName               string      `json:"CreatorFullName"`
	Currency                      string      `json:"Currency"`
	Description                   string      `json:"Description"`
	DirectDebitMandate            types.Guid  `json:"DirectDebitMandate"`
	DirectDebitMandateDescription string      `json:"DirectDebitMandateDescription"`
	DirectDebitMandatePaymentType int16       `json:"DirectDebitMandatePaymentType"`
	DirectDebitMandateReference   string      `json:"DirectDebitMandateReference"`
	DirectDebitMandateType        int16       `json:"DirectDebitMandateType"`
	DiscountDueDate               *types.Date `json:"DiscountDueDate"`
	Division                      int32       `json:"Division"`
	Document                      types.Guid  `json:"Document"`
	DocumentNumber                int32       `json:"DocumentNumber"`
	DocumentSubject               string      `json:"DocumentSubject"`
	DueDate                       *types.Date `json:"DueDate"`
	EndDate                       *types.Date `json:"EndDate"`
	EndPeriod                     int16       `json:"EndPeriod"`
	EndToEndID                    string      `json:"EndToEndID"`
	EndYear                       int16       `json:"EndYear"`
	EntryDate                     *types.Date `json:"EntryDate"`
	EntryID                       types.Guid  `json:"EntryID"`
	EntryNumber                   int32       `json:"EntryNumber"`
	GLAccount                     types.Guid  `json:"GLAccount"`
	GLAccountCode                 string      `json:"GLAccountCode"`
	GLAccountDescription          string      `json:"GLAccountDescription"`
	InvoiceDate                   *types.Date `json:"InvoiceDate"`
	InvoiceNumber                 int32       `json:"InvoiceNumber"`
	IsBatchBooking                byte        `json:"IsBatchBooking"`
	IsFullyPaid                   bool        `json:"IsFullyPaid"`
	Journal                       string      `json:"Journal"`
	JournalDescription            string      `json:"JournalDescription"`
	LastPaymentDate               *types.Date `json:"LastPaymentDate"`
	Modified                      *types.Date `json:"Modified"`
	Modifier                      types.Guid  `json:"Modifier"`
	ModifierFullName              string      `json:"ModifierFullName"`
	PaymentCondition              string      `json:"PaymentCondition"`
	PaymentConditionDescription   string      `json:"PaymentConditionDescription"`
	PaymentDays                   int32       `json:"PaymentDays"`
	PaymentDaysDiscount           int32       `json:"PaymentDaysDiscount"`
	PaymentDiscountPercentage     float64     `json:"PaymentDiscountPercentage"`
	PaymentInformationID          string      `json:"PaymentInformationID"`
	PaymentMethod                 string      `json:"PaymentMethod"`
	PaymentReference              string      `json:"PaymentReference"`
	RateFC                        float64     `json:"RateFC"`
	ReceivableBatchNumber         int32       `json:"ReceivableBatchNumber"`
	ReceivableSelected            *types.Date `json:"ReceivableSelected"`
	ReceivableSelector            types.Guid  `json:"ReceivableSelector"`
	ReceivableSelectorFullName    string      `json:"ReceivableSelectorFullName"`
	Source                        int32       `json:"Source"`
	Status                        int16       `json:"Status"`
	TransactionAmountDC           float64     `json:"TransactionAmountDC"`
	TransactionAmountFC           float64     `json:"TransactionAmountFC"`
	TransactionDueDate            *types.Date `json:"TransactionDueDate"`
	TransactionEntryID            types.Guid  `json:"TransactionEntryID"`
	TransactionID                 types.Guid  `json:"TransactionID"`
	TransactionIsReversal         bool        `json:"TransactionIsReversal"`
	TransactionReportingPeriod    int16       `json:"TransactionReportingPeriod"`
	TransactionReportingYear      int16       `json:"TransactionReportingYear"`
	TransactionStatus             int16       `json:"TransactionStatus"`
	TransactionType               int32       `json:"TransactionType"`
	YourRef                       string      `json:"YourRef"`
}

type GetReceivablesCall struct {
	modifiedAfter *time.Time
	urlNext       string
	service       *Service
}

func (service *Service) NewGetReceivablesCall(modifiedAfter *time.Time) *GetReceivablesCall {
	call := GetReceivablesCall{}
	call.modifiedAfter = modifiedAfter
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", Receivable{})
	call.urlNext = service.url(fmt.Sprintf("Receivables?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetReceivablesCall) Do() (*[]Receivable, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	banks := []Receivable{}

	next, err := call.service.Get(call.urlNext, &banks)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &banks, nil
}

func (service *Service) GetReceivablesCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("Receivables", createdBefore)
}
