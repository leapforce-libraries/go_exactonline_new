package exactonline

import (
	"encoding/json"
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// CRMAccount stores CRMAccount from exactonline
//
type CRMAccount struct {
	Timestamp                           types.Int64String `json:"Timestamp"`
	Accountant                          types.Guid        `json:"Accountant"`
	AccountManager                      types.Guid        `json:"AccountManager"`
	AccountManagerFullName              string            `json:"AccountManagerFullName"`
	AccountManagerHID                   int32             `json:"AccountManagerHID"`
	ActivitySector                      types.Guid        `json:"ActivitySector"`
	ActivitySubSector                   types.Guid        `json:"ActivitySubSector"`
	AddressLine1                        string            `json:"AddressLine1"`
	AddressLine2                        string            `json:"AddressLine2"`
	AddressLine3                        string            `json:"AddressLine3"`
	BankAccounts                        json.RawMessage   `json:"BankAccounts"` //to be implemented when needed
	Blocked                             bool              `json:"Blocked"`
	BRIN                                types.Guid        `json:"BRIN"`
	BSN                                 string            `json:"BSN"`
	BusinessType                        types.Guid        `json:"BusinessType"`
	CanDropShip                         bool              `json:"CanDropShip"`
	ChamberOfCommerce                   string            `json:"ChamberOfCommerce"`
	City                                string            `json:"City"`
	Classification                      string            `json:"Classification"`
	Classification1                     types.Guid        `json:"Classification1"`
	Classification2                     types.Guid        `json:"Classification2"`
	Classification3                     types.Guid        `json:"Classification3"`
	Classification4                     types.Guid        `json:"Classification4"`
	Classification5                     types.Guid        `json:"Classification5"`
	Classification6                     types.Guid        `json:"Classification6"`
	Classification7                     types.Guid        `json:"Classification7"`
	Classification8                     types.Guid        `json:"Classification8"`
	ClassificationDescription           string            `json:"ClassificationDescription"`
	Code                                string            `json:"Code"`
	CodeAtSupplier                      string            `json:"CodeAtSupplier"`
	CompanySize                         types.Guid        `json:"CompanySize"`
	ConsolidationScenario               byte              `json:"ConsolidationScenario"`
	ControlledDate                      *types.Date       `json:"ControlledDate"`
	Costcenter                          string            `json:"Costcenter"`
	CostcenterDescription               string            `json:"CostcenterDescription"`
	CostPaid                            byte              `json:"CostPaid"`
	Country                             string            `json:"Country"`
	CountryName                         string            `json:"CountryName"`
	Created                             *types.Date       `json:"Created"`
	Creator                             types.Guid        `json:"Creator"`
	CreatorFullName                     string            `json:"CreatorFullName"`
	CreditLinePurchase                  float64           `json:"CreditLinePurchase"`
	CreditLineSales                     float64           `json:"CreditLineSales"`
	Currency                            string            `json:"Currency"`
	CustomerSince                       *types.Date       `json:"CustomerSince"`
	DatevCreditorCode                   string            `json:"DatevCreditorCode"`
	DatevDebtorCode                     string            `json:"DatevDebtorCode"`
	DiscountPurchase                    float64           `json:"DiscountPurchase"`
	DiscountSales                       float64           `json:"DiscountSales"`
	Division                            int32             `json:"Division"`
	Document                            types.Guid        `json:"Document"`
	DunsNumber                          string            `json:"DunsNumber"`
	Email                               string            `json:"Email"`
	EndDate                             *types.Date       `json:"EndDate"`
	EstablishedDate                     *types.Date       `json:"EstablishedDate"`
	Fax                                 string            `json:"Fax"`
	GLAccountPurchase                   types.Guid        `json:"GLAccountPurchase"`
	GLAccountSales                      types.Guid        `json:"GLAccountSales"`
	GLAP                                types.Guid        `json:"GLAP"`
	GLAR                                types.Guid        `json:"GLAR"`
	GlnNumber                           string            `json:"GlnNumber"`
	HasWithholdingTaxSales              bool              `json:"HasWithholdingTaxSales"`
	ID                                  types.Guid        `json:"ID"`
	IgnoreDatevWarningMessage           bool              `json:"IgnoreDatevWarningMessage"`
	IntraStatArea                       string            `json:"IntraStatArea"`
	IntraStatDeliveryTerm               string            `json:"IntraStatDeliveryTerm"`
	IntraStatSystem                     string            `json:"IntraStatSystem"`
	IntraStatTransactionA               string            `json:"IntraStatTransactionA"`
	IntraStatTransactionB               string            `json:"IntraStatTransactionB"`
	IntraStatTransportMethod            string            `json:"IntraStatTransportMethod"`
	InvoiceAccount                      types.Guid        `json:"InvoiceAccount"`
	InvoiceAccountCode                  string            `json:"InvoiceAccountCode"`
	InvoiceAccountName                  string            `json:"InvoiceAccountName"`
	InvoiceAttachmentType               int32             `json:"InvoiceAttachmentType"`
	InvoicingMethod                     int32             `json:"InvoicingMethod"`
	IsAccountant                        byte              `json:"IsAccountant"`
	IsAgency                            byte              `json:"IsAgency"`
	IsAnonymised                        byte              `json:"IsAnonymised"`
	IsBank                              bool              `json:"IsBank"`
	IsCompetitor                        byte              `json:"IsCompetitor"`
	IsExtraDuty                         bool              `json:"IsExtraDuty"`
	IsMailing                           byte              `json:"IsMailing"`
	IsMember                            bool              `json:"IsMember"`
	IsPilot                             bool              `json:"IsPilot"`
	IsPurchase                          bool              `json:"IsPurchase"`
	IsReseller                          bool              `json:"IsReseller"`
	IsSales                             bool              `json:"IsSales"`
	IsSupplier                          bool              `json:"IsSupplier"`
	Language                            string            `json:"Language"`
	LanguageDescription                 string            `json:"LanguageDescription"`
	Latitude                            float64           `json:"Latitude"`
	LeadPurpose                         types.Guid        `json:"LeadPurpose"`
	LeadSource                          types.Guid        `json:"LeadSource"`
	Logo                                json.RawMessage   `json:"Logo"` //to be implemented when needed
	LogoFileName                        string            `json:"LogoFileName"`
	LogoThumbnailURL                    string            `json:"LogoThumbnailUrl"`
	LogoURL                             string            `json:"LogoUrl"`
	Longitude                           float64           `json:"Longitude"`
	MainContact                         *types.Guid       `json:"MainContact"`
	Modified                            *types.Date       `json:"Modified"`
	Modifier                            types.Guid        `json:"Modifier"`
	ModifierFullName                    string            `json:"ModifierFullName"`
	Name                                string            `json:"Name"`
	OINNumber                           string            `json:"OINNumber"`
	Parent                              types.Guid        `json:"Parent"`
	PayAsYouEarn                        string            `json:"PayAsYouEarn"`
	PaymentConditionPurchase            string            `json:"PaymentConditionPurchase"`
	PaymentConditionPurchaseDescription string            `json:"PaymentConditionPurchaseDescription"`
	PaymentConditionSales               string            `json:"PaymentConditionSales"`
	PaymentConditionSalesDescription    string            `json:"PaymentConditionSalesDescription"`
	Phone                               string            `json:"Phone"`
	PhoneExtension                      string            `json:"PhoneExtension"`
	Postcode                            string            `json:"Postcode"`
	PriceList                           types.Guid        `json:"PriceList"`
	PurchaseCurrency                    string            `json:"PurchaseCurrency"`
	PurchaseCurrencyDescription         string            `json:"PurchaseCurrencyDescription"`
	PurchaseLeadDays                    int32             `json:"PurchaseLeadDays"`
	PurchaseVATCode                     string            `json:"PurchaseVATCode"`
	PurchaseVATCodeDescription          string            `json:"PurchaseVATCodeDescription"`
	RecepientOfCommissions              bool              `json:"RecepientOfCommissions"`
	Remarks                             string            `json:"Remarks"`
	Reseller                            types.Guid        `json:"Reseller"`
	ResellerCode                        string            `json:"ResellerCode"`
	ResellerName                        string            `json:"ResellerName"`
	RSIN                                string            `json:"RSIN"`
	SalesCurrency                       string            `json:"SalesCurrency"`
	SalesCurrencyDescription            string            `json:"SalesCurrencyDescription"`
	SalesTaxSchedule                    types.Guid        `json:"SalesTaxSchedule"`
	SalesTaxScheduleCode                string            `json:"SalesTaxScheduleCode"`
	SalesTaxScheduleDescription         string            `json:"SalesTaxScheduleDescription"`
	SalesVATCode                        string            `json:"SalesVATCode"`
	SalesVATCodeDescription             string            `json:"SalesVATCodeDescription"`
	SearchCode                          string            `json:"SearchCode"`
	SecurityLevel                       int32             `json:"SecurityLevel"`
	SeparateInvoicePerSubscription      byte              `json:"SeparateInvPerSubscription"`
	ShippingLeadDays                    int32             `json:"ShippingLeadDays"`
	ShippingMethod                      types.Guid        `json:"ShippingMethod"`
	ShowRemarkForSales                  bool              `json:"ShowRemarkForSales"`
	StartDate                           *types.Date       `json:"StartDate"`
	State                               string            `json:"State"`
	StateName                           string            `json:"StateName"`
	Status                              string            `json:"Status"`
	StatusSince                         *types.Date       `json:"StatusSince"`
	TradeName                           string            `json:"TradeName"`
	Type                                string            `json:"Type"`
	UniqueTaxpayerReference             string            `json:"UniqueTaxpayerReference"`
	VATLiability                        string            `json:"VATLiability"`
	VATNumber                           string            `json:"VATNumber"`
	Website                             string            `json:"Website"`
}

type SyncCRMAccountsCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewSyncCRMAccountsCall(timestamp *int64) *SyncCRMAccountsCall {
	selectFields := utilities.GetTaggedTagNames("json", CRMAccount{})
	url := service.url(fmt.Sprintf("CRM/Accounts?$select=%s", selectFields))
	if timestamp != nil {
		url = fmt.Sprintf("%s&$filter=Timestamp gt %vL", url, *timestamp)
	}

	return &SyncCRMAccountsCall{
		service: service,
		urlNext: url,
	}
}

func (call *SyncCRMAccountsCall) Do() (*[]CRMAccount, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	crmAccounts := []CRMAccount{}

	next, err := call.service.Get(call.urlNext, &crmAccounts)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &crmAccounts, nil
}
