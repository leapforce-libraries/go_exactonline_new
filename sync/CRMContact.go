package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// CRMContact stores CRMContact from exactonline
//
type CRMContact struct {
	Timestamp                 types.Int64String `json:"Timestamp"`
	ID                        types.GUID        `json:"ID"`
	Account                   types.GUID        `json:"Account"`
	AccountIsCustomer         bool              `json:"AccountIsCustomer"`
	AccountIsSupplier         bool              `json:"AccountIsSupplier"`
	AccountMainContact        types.GUID        `json:"AccountMainContact"`
	AccountName               string            `json:"AccountName"`
	AddressLine2              string            `json:"AddressLine2"`
	AddressStreet             string            `json:"AddressStreet"`
	AddressStreetNumber       string            `json:"AddressStreetNumber"`
	AddressStreetNumberSuffix string            `json:"AddressStreetNumberSuffix"`
	AllowMailing              int32             `json:"AllowMailing"`
	BirthDate                 *types.Date       `json:"BirthDate"`
	BirthName                 string            `json:"BirthName"`
	BirthNamePrefix           string            `json:"BirthNamePrefix"`
	BirthPlace                string            `json:"BirthPlace"`
	BusinessEmail             string            `json:"BusinessEmail"`
	BusinessFax               string            `json:"BusinessFax"`
	BusinessMobile            string            `json:"BusinessMobile"`
	BusinessPhone             string            `json:"BusinessPhone"`
	BusinessPhoneExtension    string            `json:"BusinessPhoneExtension"`
	City                      string            `json:"City"`
	Code                      string            `json:"Code"`
	Country                   string            `json:"Country"`
	Created                   *types.Date       `json:"Created"`
	Creator                   types.GUID        `json:"Creator"`
	CreatorFullName           string            `json:"CreatorFullName"`
	Division                  int32             `json:"Division"`
	Email                     string            `json:"Email"`
	EndDate                   *types.Date       `json:"EndDate"`
	FirstName                 string            `json:"FirstName"`
	FullName                  string            `json:"FullName"`
	Gender                    string            `json:"Gender"`
	HID                       int32             `json:"HID"`
	IdentificationDate        *types.Date       `json:"IdentificationDate"`
	IdentificationDocument    types.GUID        `json:"IdentificationDocument"`
	IdentificationUser        types.GUID        `json:"IdentificationUser"`
	Initials                  string            `json:"Initials"`
	IsAnonymised              byte              `json:"IsAnonymised"`
	IsMailingExcluded         bool              `json:"IsMailingExcluded"`
	IsMainContact             bool              `json:"IsMainContact"`
	JobTitleDescription       string            `json:"JobTitleDescription"`
	Language                  string            `json:"Language"`
	LastName                  string            `json:"LastName"`
	LeadPurpose               types.GUID        `json:"LeadPurpose"`
	LeadSource                types.GUID        `json:"LeadSource"`
	MarketingNotes            string            `json:"MarketingNotes"`
	MiddleName                string            `json:"MiddleName"`
	Mobile                    string            `json:"Mobile"`
	Modified                  *types.Date       `json:"Modified"`
	Modifier                  types.GUID        `json:"Modifier"`
	ModifierFullName          string            `json:"ModifierFullName"`
	Nationality               string            `json:"Nationality"`
	Notes                     string            `json:"Notes"`
	PartnerName               string            `json:"PartnerName"`
	PartnerNamePrefix         string            `json:"PartnerNamePrefix"`
	Person                    types.GUID        `json:"Person"`
	Phone                     string            `json:"Phone"`
	PhoneExtension            string            `json:"PhoneExtension"`
	PictureName               string            `json:"PictureName"`
	PictureThumbnailUrl       string            `json:"PictureThumbnailUrl"`
	PictureUrl                string            `json:"PictureUrl"`
	Postcode                  string            `json:"Postcode"`
	SocialSecurityNumber      string            `json:"SocialSecurityNumber"`
	StartDate                 *types.Date       `json:"StartDate"`
	State                     string            `json:"State"`
	Title                     string            `json:"Title"`
}

type SyncCRMContactsCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewSyncCRMContactsCall(timestamp *int64) *SyncCRMContactsCall {
	selectFields := utilities.GetTaggedTagNames("json", CRMContact{})
	url := service.url(fmt.Sprintf("CRM/Contacts?$select=%s", selectFields))
	if timestamp != nil {
		url = fmt.Sprintf("%s&$filter=Timestamp gt %vL", url, *timestamp)
	}

	return &SyncCRMContactsCall{
		service: service,
		urlNext: url,
	}
}

func (call *SyncCRMContactsCall) Do() (*[]CRMContact, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	crmContacts := []CRMContact{}

	next, err := call.service.Get(call.urlNext, &crmContacts)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &crmContacts, nil
}
