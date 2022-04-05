package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// Payment stores Payment from exactonline
//
type Payment struct {
	ID                           types.Guid  `json:"ID"`
	Account                      types.Guid  `json:"Account"`
	AccountBankAccountID         types.Guid  `json:"AccountBankAccountID"`
	AccountBankAccountNumber     string      `json:"AccountBankAccountNumber"`
	AccountCode                  string      `json:"AccountCode"`
	AccountContact               types.Guid  `json:"AccountContact"`
	AccountContactName           string      `json:"AccountContactName"`
	AccountName                  string      `json:"AccountName"`
	AmountDC                     float64     `json:"AmountDC"`
	AmountDiscountDC             float64     `json:"AmountDiscountDC"`
	AmountDiscountFC             float64     `json:"AmountDiscountFC"`
	AmountFC                     float64     `json:"AmountFC"`
	BankAccountID                types.Guid  `json:"BankAccountID"`
	BankAccountNumber            string      `json:"BankAccountNumber"`
	CashflowTransactionBatchCode string      `json:"CashflowTransactionBatchCode"`
	Created                      *types.Date `json:"Created"`
	Creator                      types.Guid  `json:"Creator"`
	CreatorFullName              string      `json:"CreatorFullName"`
	Currency                     string      `json:"Currency"`
	Description                  string      `json:"Description"`
	DiscountDueDate              *types.Date `json:"DiscountDueDate"`
	Division                     int32       `json:"Division"`
	Document                     types.Guid  `json:"Document"`
	DocumentNumber               int32       `json:"DocumentNumber"`
	DocumentSubject              string      `json:"DocumentSubject"`
	DueDate                      *types.Date `json:"DueDate"`
	EndDate                      *types.Date `json:"EndDate"`
	EndPeriod                    int16       `json:"EndPeriod"`
	EndYear                      int16       `json:"EndYear"`
	EntryDate                    *types.Date `json:"EntryDate"`
	EntryID                      types.Guid  `json:"EntryID"`
	EntryNumber                  int32       `json:"EntryNumber"`
	GLAccount                    types.Guid  `json:"GLAccount"`
	GLAccountCode                string      `json:"GLAccountCode"`
	GLAccountDescription         string      `json:"GLAccountDescription"`
	InvoiceDate                  *types.Date `json:"InvoiceDate"`
	InvoiceNumber                int32       `json:"InvoiceNumber"`
	IsBatchBooking               byte        `json:"IsBatchBooking"`
	Journal                      string      `json:"Journal"`
	JournalDescription           string      `json:"JournalDescription"`
	Modified                     *types.Date `json:"Modified"`
	Modifier                     types.Guid  `json:"Modifier"`
	ModifierFullName             string      `json:"ModifierFullName"`
	PaymentBatchNumber           int32       `json:"PaymentBatchNumber"`
	PaymentCondition             string      `json:"PaymentCondition"`
	PaymentConditionDescription  string      `json:"PaymentConditionDescription"`
	PaymentDays                  int32       `json:"PaymentDays"`
	PaymentDaysDiscount          int32       `json:"PaymentDaysDiscount"`
	PaymentDiscountPercentage    float64     `json:"PaymentDiscountPercentage"`
	PaymentMethod                string      `json:"PaymentMethod"`
	PaymentReference             string      `json:"PaymentReference"`
	PaymentSelected              *types.Date `json:"PaymentSelected"`
	PaymentSelector              types.Guid  `json:"PaymentSelector"`
	PaymentSelectorFullName      string      `json:"PaymentSelectorFullName"`
	RateFC                       float64     `json:"RateFC"`
	Source                       int32       `json:"Source"`
	Status                       int16       `json:"Status"`
	TransactionAmountDC          float64     `json:"TransactionAmountDC"`
	TransactionAmountFC          float64     `json:"TransactionAmountFC"`
	TransactionDueDate           *types.Date `json:"TransactionDueDate"`
	TransactionEntryID           types.Guid  `json:"TransactionEntryID"`
	TransactionID                types.Guid  `json:"TransactionID"`
	TransactionIsReversal        bool        `json:"TransactionIsReversal"`
	TransactionReportingPeriod   int16       `json:"TransactionReportingPeriod"`
	TransactionReportingYear     int16       `json:"TransactionReportingYear"`
	TransactionStatus            int16       `json:"TransactionStatus"`
	TransactionType              int32       `json:"TransactionType"`
	YourRef                      string      `json:"YourRef"`
}

type GetPaymentsCall struct {
	modifiedAfter *time.Time
	urlNext       string
	service       *Service
}

func (service *Service) NewGetPaymentsCall(modifiedAfter *time.Time) *GetPaymentsCall {
	call := GetPaymentsCall{}
	call.modifiedAfter = modifiedAfter
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", Payment{})
	call.urlNext = service.url(fmt.Sprintf("Payments?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetPaymentsCall) Do() (*[]Payment, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	banks := []Payment{}

	next, err := call.service.Get(call.urlNext, &banks)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &banks, nil
}

func (service *Service) GetPaymentsCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("Payments", createdBefore)
}
