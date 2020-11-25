package exactonline

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// Account stores Account from exactonline
//
type Account struct {
	ID                                  types.GUID      `json:"ID"`
	Accountant                          types.GUID      `json:"Accountant"`
	AccountManager                      types.GUID      `json:"AccountManager"`
	AccountManagerFullName              string          `json:"AccountManagerFullName"`
	AccountManagerHID                   int32           `json:"AccountManagerHID"`
	ActivitySector                      types.GUID      `json:"ActivitySector"`
	ActivitySubSector                   types.GUID      `json:"ActivitySubSector"`
	AddressLine1                        string          `json:"AddressLine1"`
	AddressLine2                        string          `json:"AddressLine2"`
	AddressLine3                        string          `json:"AddressLine3"`
	BankAccounts                        json.RawMessage `json:"BankAccounts"` //to be implemented when needed
	Blocked                             bool            `json:"Blocked"`
	BRIN                                types.GUID      `json:"BRIN"`
	BSN                                 string          `json:"BSN"`
	BusinessType                        types.GUID      `json:"BusinessType"`
	CanDropShip                         bool            `json:"CanDropShip"`
	ChamberOfCommerce                   string          `json:"ChamberOfCommerce"`
	City                                string          `json:"City"`
	Classification                      string          `json:"Classification"`
	Classification1                     types.GUID      `json:"Classification1"`
	Classification2                     types.GUID      `json:"Classification2"`
	Classification3                     types.GUID      `json:"Classification3"`
	Classification4                     types.GUID      `json:"Classification4"`
	Classification5                     types.GUID      `json:"Classification5"`
	Classification6                     types.GUID      `json:"Classification6"`
	Classification7                     types.GUID      `json:"Classification7"`
	Classification8                     types.GUID      `json:"Classification8"`
	ClassificationDescription           string          `json:"ClassificationDescription"`
	Code                                string          `json:"Code"`
	CodeAtSupplier                      string          `json:"CodeAtSupplier"`
	CompanySize                         types.GUID      `json:"CompanySize"`
	ConsolidationScenario               byte            `json:"ConsolidationScenario"`
	ControlledDate                      *types.Date     `json:"ControlledDate"`
	Costcenter                          string          `json:"Costcenter"`
	CostcenterDescription               string          `json:"CostcenterDescription"`
	CostPaid                            byte            `json:"CostPaid"`
	Country                             string          `json:"Country"`
	CountryName                         string          `json:"CountryName"`
	Created                             *types.Date     `json:"Created"`
	Creator                             types.GUID      `json:"Creator"`
	CreatorFullName                     string          `json:"CreatorFullName"`
	CreditLinePurchase                  float64         `json:"CreditLinePurchase"`
	CreditLineSales                     float64         `json:"CreditLineSales"`
	Currency                            string          `json:"Currency"`
	CustomerSince                       *types.Date     `json:"CustomerSince"`
	DatevCreditorCode                   string          `json:"DatevCreditorCode"`
	DatevDebtorCode                     string          `json:"DatevDebtorCode"`
	DiscountPurchase                    float64         `json:"DiscountPurchase"`
	DiscountSales                       float64         `json:"DiscountSales"`
	Division                            int32           `json:"Division"`
	Document                            types.GUID      `json:"Document"`
	DunsNumber                          string          `json:"DunsNumber"`
	Email                               string          `json:"Email"`
	EndDate                             *types.Date     `json:"EndDate"`
	EstablishedDate                     *types.Date     `json:"EstablishedDate"`
	Fax                                 string          `json:"Fax"`
	GLAccountPurchase                   types.GUID      `json:"GLAccountPurchase"`
	GLAccountSales                      types.GUID      `json:"GLAccountSales"`
	GLAP                                types.GUID      `json:"GLAP"`
	GLAR                                types.GUID      `json:"GLAR"`
	GlnNumber                           string          `json:"GlnNumber"`
	HasWithholdingTaxSales              bool            `json:"HasWithholdingTaxSales"`
	IgnoreDatevWarningMessage           bool            `json:"IgnoreDatevWarningMessage"`
	IntraStatArea                       string          `json:"IntraStatArea"`
	IntraStatDeliveryTerm               string          `json:"IntraStatDeliveryTerm"`
	IntraStatSystem                     string          `json:"IntraStatSystem"`
	IntraStatTransactionA               string          `json:"IntraStatTransactionA"`
	IntraStatTransactionB               string          `json:"IntraStatTransactionB"`
	IntraStatTransportMethod            string          `json:"IntraStatTransportMethod"`
	InvoiceAccount                      types.GUID      `json:"InvoiceAccount"`
	InvoiceAccountCode                  string          `json:"InvoiceAccountCode"`
	InvoiceAccountName                  string          `json:"InvoiceAccountName"`
	InvoiceAttachmentType               int32           `json:"InvoiceAttachmentType"`
	InvoicingMethod                     int32           `json:"InvoicingMethod"`
	IsAccountant                        byte            `json:"IsAccountant"`
	IsAgency                            byte            `json:"IsAgency"`
	IsAnonymised                        byte            `json:"IsAnonymised"`
	IsBank                              bool            `json:"IsBank"`
	IsCompetitor                        byte            `json:"IsCompetitor"`
	IsExtraDuty                         bool            `json:"IsExtraDuty"`
	IsMailing                           byte            `json:"IsMailing"`
	IsMember                            bool            `json:"IsMember"`
	IsPilot                             bool            `json:"IsPilot"`
	IsPurchase                          bool            `json:"IsPurchase"`
	IsReseller                          bool            `json:"IsReseller"`
	IsSales                             bool            `json:"IsSales"`
	IsSupplier                          bool            `json:"IsSupplier"`
	Language                            string          `json:"Language"`
	LanguageDescription                 string          `json:"LanguageDescription"`
	Latitude                            float64         `json:"Latitude"`
	LeadPurpose                         types.GUID      `json:"LeadPurpose"`
	LeadSource                          types.GUID      `json:"LeadSource"`
	Logo                                json.RawMessage `json:"Logo"` //to be implemented when needed
	LogoFileName                        string          `json:"LogoFileName"`
	LogoThumbnailUrl                    string          `json:"LogoThumbnailUrl"`
	LogoUrl                             string          `json:"LogoUrl"`
	Longitude                           float64         `json:"Longitude"`
	MainContact                         types.GUID      `json:"MainContact"`
	Modified                            *types.Date     `json:"Modified"`
	Modifier                            types.GUID      `json:"Modifier"`
	ModifierFullName                    string          `json:"ModifierFullName"`
	Name                                string          `json:"Name"`
	OINNumber                           string          `json:"OINNumber"`
	Parent                              types.GUID      `json:"Parent"`
	PayAsYouEarn                        string          `json:"PayAsYouEarn"`
	PaymentConditionPurchase            string          `json:"PaymentConditionPurchase"`
	PaymentConditionPurchaseDescription string          `json:"PaymentConditionPurchaseDescription"`
	PaymentConditionSales               string          `json:"PaymentConditionSales"`
	PaymentConditionSalesDescription    string          `json:"PaymentConditionSalesDescription"`
	Phone                               string          `json:"Phone"`
	PhoneExtension                      string          `json:"PhoneExtension"`
	Postcode                            string          `json:"Postcode"`
	PriceList                           types.GUID      `json:"PriceList"`
	PurchaseCurrency                    string          `json:"PurchaseCurrency"`
	PurchaseCurrencyDescription         string          `json:"PurchaseCurrencyDescription"`
	PurchaseLeadDays                    int32           `json:"PurchaseLeadDays"`
	PurchaseVATCode                     string          `json:"PurchaseVATCode"`
	PurchaseVATCodeDescription          string          `json:"PurchaseVATCodeDescription"`
	RecepientOfCommissions              bool            `json:"RecepientOfCommissions"`
	Remarks                             string          `json:"Remarks"`
	Reseller                            types.GUID      `json:"Reseller"`
	ResellerCode                        string          `json:"ResellerCode"`
	ResellerName                        string          `json:"ResellerName"`
	RSIN                                string          `json:"RSIN"`
	SalesCurrency                       string          `json:"SalesCurrency"`
	SalesCurrencyDescription            string          `json:"SalesCurrencyDescription"`
	SalesTaxSchedule                    types.GUID      `json:"SalesTaxSchedule"`
	SalesTaxScheduleCode                string          `json:"SalesTaxScheduleCode"`
	SalesTaxScheduleDescription         string          `json:"SalesTaxScheduleDescription"`
	SalesVATCode                        string          `json:"SalesVATCode"`
	SalesVATCodeDescription             string          `json:"SalesVATCodeDescription"`
	SearchCode                          string          `json:"SearchCode"`
	SecurityLevel                       int32           `json:"SecurityLevel"`
	SeparateInvPerProject               byte            `json:"SeparateInvPerProject"`
	SeparateInvPerSubscription          byte            `json:"SeparateInvPerSubscription"`
	ShippingLeadDays                    int32           `json:"ShippingLeadDays"`
	ShippingMethod                      types.GUID      `json:"ShippingMethod"`
	StartDate                           *types.Date     `json:"StartDate"`
	State                               string          `json:"State"`
	StateName                           string          `json:"StateName"`
	Status                              string          `json:"Status"`
	StatusSince                         *types.Date     `json:"StatusSince"`
	TradeName                           string          `json:"TradeName"`
	Type                                string          `json:"Type"`
	UniqueTaxpayerReference             string          `json:"UniqueTaxpayerReference"`
	VATLiability                        string          `json:"VATLiability"`
	VATNumber                           string          `json:"VATNumber"`
	Website                             string          `json:"Website"`
}

// AccountUpdate stores Account value to insert/update
//
type AccountUpdate struct {
	Accountant                          *string  `json:"Accountant,omitempty"`
	AccountManager                      *string  `json:"AccountManager,omitempty"`
	ActivitySector                      *string  `json:"ActivitySector,omitempty"`
	ActivitySubSector                   *string  `json:"ActivitySubSector,omitempty"`
	AddressLine1                        *string  `json:"AddressLine1,omitempty"`
	AddressLine2                        *string  `json:"AddressLine2,omitempty"`
	AddressLine3                        **string `json:"AddressLine3,omitempty"`
	Blocked                             *bool    `json:"Blocked,omitempty"`
	BRIN                                *string  `json:"BRIN,omitempty"`
	BSN                                 *string  `json:"BSN,omitempty"`
	BusinessType                        *string  `json:"BusinessType,omitempty"`
	CanDropShip                         *bool    `json:"CanDropShip,omitempty"`
	ChamberOfCommerce                   *string  `json:"ChamberOfCommerce,omitempty"`
	City                                *string  `json:"City,omitempty"`
	Classification                      *string  `json:"Classification,omitempty"`
	Classification1                     *string  `json:"Classification1,omitempty"`
	Classification2                     *string  `json:"Classification2,omitempty"`
	Classification3                     *string  `json:"Classification3,omitempty"`
	Classification4                     *string  `json:"Classification4,omitempty"`
	Classification5                     *string  `json:"Classification5,omitempty"`
	Classification6                     *string  `json:"Classification6,omitempty"`
	Classification7                     *string  `json:"Classification7,omitempty"`
	Classification8                     *string  `json:"Classification8,omitempty"`
	ClassificationDescription           *string  `json:"ClassificationDescription,omitempty"`
	Code                                *string  `json:"Code,omitempty"`
	CodeAtSupplier                      *string  `json:"CodeAtSupplier,omitempty"`
	CompanySize                         *string  `json:"CompanySize,omitempty"`
	ConsolidationScenario               *byte    `json:"ConsolidationScenario,omitempty"`
	ControlledDate                      *string  `json:"ControlledDate,omitempty"`
	Costcenter                          *string  `json:"Costcenter,omitempty"`
	CostcenterDescription               *string  `json:"CostcenterDescription,omitempty"`
	CostPaid                            *byte    `json:"CostPaid,omitempty"`
	Country                             *string  `json:"Country,omitempty"`
	CountryName                         *string  `json:"CountryName,omitempty"`
	Created                             *string  `json:"Created,omitempty"`
	CreditLinePurchase                  *float64 `json:"CreditLinePurchase,omitempty"`
	CreditLineSales                     *float64 `json:"CreditLineSales,omitempty"`
	Currency                            *string  `json:"Currency,omitempty"`
	CustomerSince                       *string  `json:"CustomerSince,omitempty"`
	DatevCreditorCode                   *string  `json:"DatevCreditorCode,omitempty"`
	DatevDebtorCode                     *string  `json:"DatevDebtorCode,omitempty"`
	DiscountPurchase                    *float64 `json:"DiscountPurchase,omitempty"`
	DiscountSales                       *float64 `json:"DiscountSales,omitempty"`
	Division                            *int32   `json:"Division,omitempty"`
	Document                            *string  `json:"Document,omitempty"`
	DunsNumber                          *string  `json:"DunsNumber,omitempty"`
	Email                               *string  `json:"Email,omitempty"`
	EndDate                             *string  `json:"EndDate,omitempty"`
	EstablishedDate                     *string  `json:"EstablishedDate,omitempty"`
	Fax                                 *string  `json:"Fax,omitempty"`
	GLAccountPurchase                   *string  `json:"GLAccountPurchase,omitempty"`
	GLAccountSales                      *string  `json:"GLAccountSales,omitempty"`
	GLAP                                *string  `json:"GLAP,omitempty"`
	GLAR                                *string  `json:"GLAR,omitempty"`
	GlnNumber                           *string  `json:"GlnNumber,omitempty"`
	HasWithholdingTaxSales              *bool    `json:"HasWithholdingTaxSales,omitempty"`
	IgnoreDatevWarningMessage           *bool    `json:"IgnoreDatevWarningMessage,omitempty"`
	IntraStatArea                       *string  `json:"IntraStatArea,omitempty"`
	IntraStatDeliveryTerm               *string  `json:"IntraStatDeliveryTerm,omitempty"`
	IntraStatSystem                     *string  `json:"IntraStatSystem,omitempty"`
	IntraStatTransactionA               *string  `json:"IntraStatTransactionA,omitempty"`
	IntraStatTransactionB               *string  `json:"IntraStatTransactionB,omitempty"`
	IntraStatTransportMethod            *string  `json:"IntraStatTransportMethod,omitempty"`
	InvoiceAccount                      *string  `json:"InvoiceAccount,omitempty"`
	InvoiceAttachmentType               *int32   `json:"InvoiceAttachmentType,omitempty"`
	InvoicingMethod                     *int32   `json:"InvoicingMethod,omitempty"`
	IsAccountant                        *byte    `json:"IsAccountant,omitempty"`
	IsAgency                            *byte    `json:"IsAgency,omitempty"`
	IsAnonymised                        *byte    `json:"IsAnonymised,omitempty"`
	IsBank                              *bool    `json:"IsBank,omitempty"`
	IsCompetitor                        *byte    `json:"IsCompetitor,omitempty"`
	IsExtraDuty                         *bool    `json:"IsExtraDuty,omitempty"`
	IsMailing                           *byte    `json:"IsMailing,omitempty"`
	IsMember                            *bool    `json:"IsMember,omitempty"`
	IsPilot                             *bool    `json:"IsPilot,omitempty"`
	IsPurchase                          *bool    `json:"IsPurchase,omitempty"`
	IsReseller                          *bool    `json:"IsReseller,omitempty"`
	IsSales                             *bool    `json:"IsSales,omitempty"`
	IsSupplier                          *bool    `json:"IsSupplier,omitempty"`
	Language                            *string  `json:"Language,omitempty"`
	LanguageDescription                 *string  `json:"LanguageDescription,omitempty"`
	Latitude                            *float64 `json:"Latitude,omitempty"`
	LeadPurpose                         *string  `json:"LeadPurpose,omitempty"`
	LeadSource                          *string  `json:"LeadSource,omitempty"`
	LogoFileName                        *string  `json:"LogoFileName,omitempty"`
	LogoThumbnailUrl                    *string  `json:"LogoThumbnailUrl,omitempty"`
	LogoUrl                             *string  `json:"LogoUrl,omitempty"`
	Longitude                           *float64 `json:"Longitude,omitempty"`
	MainContact                         *string  `json:"MainContact,omitempty"`
	Modified                            *string  `json:"Modified,omitempty"`
	Name                                *string  `json:"Name,omitempty"`
	OINNumber                           *string  `json:"OINNumber,omitempty"`
	Parent                              *string  `json:"Parent,omitempty"`
	PayAsYouEarn                        *string  `json:"PayAsYouEarn,omitempty"`
	PaymentConditionPurchase            *string  `json:"PaymentConditionPurchase,omitempty"`
	PaymentConditionPurchaseDescription *string  `json:"PaymentConditionPurchaseDescription,omitempty"`
	PaymentConditionSales               *string  `json:"PaymentConditionSales,omitempty"`
	PaymentConditionSalesDescription    *string  `json:"PaymentConditionSalesDescription,omitempty"`
	Phone                               *string  `json:"Phone,omitempty"`
	PhoneExtension                      *string  `json:"PhoneExtension,omitempty"`
	Postcode                            *string  `json:"Postcode,omitempty"`
	PriceList                           *string  `json:"PriceList,omitempty"`
	PurchaseCurrency                    *string  `json:"PurchaseCurrency,omitempty"`
	PurchaseCurrencyDescription         *string  `json:"PurchaseCurrencyDescription,omitempty"`
	PurchaseLeadDays                    *int32   `json:"PurchaseLeadDays,omitempty"`
	PurchaseVATCode                     *string  `json:"PurchaseVATCode,omitempty"`
	PurchaseVATCodeDescription          *string  `json:"PurchaseVATCodeDescription,omitempty"`
	RecepientOfCommissions              *bool    `json:"RecepientOfCommissions,omitempty"`
	Remarks                             *string  `json:"Remarks,omitempty"`
	Reseller                            *string  `json:"Reseller,omitempty"`
	RSIN                                *string  `json:"RSIN,omitempty"`
	SalesCurrency                       *string  `json:"SalesCurrency,omitempty"`
	SalesCurrencyDescription            *string  `json:"SalesCurrencyDescription,omitempty"`
	SalesTaxSchedule                    *string  `json:"SalesTaxSchedule,omitempty"`
	SalesVATCode                        *string  `json:"SalesVATCode,omitempty"`
	SalesVATCodeDescription             *string  `json:"SalesVATCodeDescription,omitempty"`
	SearchCode                          *string  `json:"SearchCode,omitempty"`
	SecurityLevel                       *int32   `json:"SecurityLevel,omitempty"`
	SeparateInvPerProject               *byte    `json:"SeparateInvPerProject,omitempty"`
	SeparateInvPerSubscription          *byte    `json:"SeparateInvPerSubscription,omitempty"`
	ShippingLeadDays                    *int32   `json:"ShippingLeadDays,omitempty"`
	ShippingMethod                      *string  `json:"ShippingMethod,omitempty"`
	StartDate                           *string  `json:"StartDate,omitempty"`
	State                               *string  `json:"State,omitempty"`
	StateName                           *string  `json:"StateName,omitempty"`
	Status                              *string  `json:"Status,omitempty"`
	StatusSince                         *string  `json:"StatusSince,omitempty"`
	TradeName                           *string  `json:"TradeName,omitempty"`
	Type                                *string  `json:"Type,omitempty"`
	UniqueTaxpayerReference             *string  `json:"UniqueTaxpayerReference,omitempty"`
	VATLiability                        *string  `json:"VATLiability,omitempty"`
	VATNumber                           *string  `json:"VATNumber,omitempty"`
	Website                             *string  `json:"Website,omitempty"`
}

type GetAccountsCall struct {
	urlNext string
	client  *Client
}

type GetAccountsCallParams struct {
	ChamberOfCommerce *string
	ModifiedAfter     *time.Time
}

func (c *Client) NewGetAccountsCall(params GetAccountsCallParams) *GetAccountsCall {
	call := GetAccountsCall{}
	call.client = c

	selectFields := utilities.GetTaggedFieldNames("json", Account{})
	call.urlNext = fmt.Sprintf("%s/Accounts?$select=%s", c.BaseURL(), selectFields)
	filter := []string{}

	if params.ChamberOfCommerce != nil {
		filter = append(filter, fmt.Sprintf("ChamberOfCommerce eq '%s'", *params.ChamberOfCommerce))
	}
	if params.ModifiedAfter != nil {
		filter = append(filter, c.DateFilter("Modified", "gt", params.ModifiedAfter, true, "&"))
	}

	if len(filter) > 0 {
		call.urlNext = fmt.Sprintf("%s&$filter=%s", call.urlNext, strings.Join(filter, " and "))
	}

	return &call
}

func (call *GetAccountsCall) Do() (*[]Account, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	accounts := []Account{}

	next, err := call.client.Get(call.urlNext, &accounts)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &accounts, nil
}

func (c *Client) CreateAccount(account *AccountUpdate) (*Account, *errortools.Error) {
	url := fmt.Sprintf("%s/Accounts", c.BaseURL())

	b, err := json.Marshal(account)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	accountNew := Account{}

	e := c.Post(url, bytes.NewBuffer(b), &accountNew)
	if e != nil {
		return nil, e
	}
	return &accountNew, nil
}

func (c *Client) UpdateAccount(id types.GUID, account *AccountUpdate) *errortools.Error {
	url := fmt.Sprintf("%s/Accounts(guid'%s')", c.BaseURL(), id.String())

	b, err := json.Marshal(account)
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	e := c.Put(url, bytes.NewBuffer(b))
	if e != nil {
		return e
	}
	return nil
}

func (c *Client) DeleteAccount(id types.GUID) *errortools.Error {
	url := fmt.Sprintf("%s/Accounts(guid'%s')", c.BaseURL(), id.String())

	err := c.Delete(url)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetAccountsCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return c.GetCount("Accounts", createdBefore)
}
