package exactonline

import (
	"encoding/json"
	"fmt"
	"time"

	types "github.com/Leapforce-nl/go_types"
	utilities "github.com/Leapforce-nl/go_utilities"
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

func (c *Client) GetAccountsInternal(filter string) (*[]Account, error) {
	selectFields := utilities.GetTaggedFieldNames("json", Account{})
	urlStr := fmt.Sprintf("%s/crm/Accounts?$select=%s", c.Http().BaseURL(), selectFields)
	if filter != "" {
		urlStr += fmt.Sprintf("&$filter=%s", filter)
	}
	//fmt.Println(urlStr)

	accounts := []Account{}

	for urlStr != "" {
		ac := []Account{}

		next, err := c.Http().Get(urlStr, &ac)
		if err != nil {
			fmt.Println("ERROR in GetAccountsInternal:", err)
			fmt.Println("url:", urlStr)
			return nil, err
		}

		accounts = append(accounts, ac...)

		urlStr = next
	}

	return &accounts, nil
}

func (c *Client) GetAccounts(filter string) (*[]Account, error) {
	acc, err := c.GetAccountsInternal(filter)
	if err != nil {
		return nil, err
	}

	return acc, nil
}

func (c *Client) GetAccountsCount(modifiedBefore *time.Time) (int64, error) {
	return c.Http().GetCount("crm/Accounts", modifiedBefore)
}
