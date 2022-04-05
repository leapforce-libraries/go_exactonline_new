package exactonline

import (
	"fmt"
	"strings"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// Contact stores Contact from exactonline
//
type Contact struct {
	ID                        types.Guid  `json:"ID"`
	Account                   types.Guid  `json:"Account"`
	AccountIsCustomer         bool        `json:"AccountIsCustomer"`
	AccountIsSupplier         bool        `json:"AccountIsSupplier"`
	AccountMainContact        types.Guid  `json:"AccountMainContact"`
	AccountName               string      `json:"AccountName"`
	AddressLine2              string      `json:"AddressLine2"`
	AddressStreet             string      `json:"AddressStreet"`
	AddressStreetNumber       string      `json:"AddressStreetNumber"`
	AddressStreetNumberSuffix string      `json:"AddressStreetNumberSuffix"`
	AllowMailing              int32       `json:"AllowMailing"`
	BirthDate                 *types.Date `json:"BirthDate"`
	BirthName                 string      `json:"BirthName"`
	BirthNamePrefix           string      `json:"BirthNamePrefix"`
	BirthPlace                string      `json:"BirthPlace"`
	BusinessEmail             string      `json:"BusinessEmail"`
	BusinessFax               string      `json:"BusinessFax"`
	BusinessMobile            string      `json:"BusinessMobile"`
	BusinessPhone             string      `json:"BusinessPhone"`
	BusinessPhoneExtension    string      `json:"BusinessPhoneExtension"`
	City                      string      `json:"City"`
	Code                      string      `json:"Code"`
	Country                   string      `json:"Country"`
	Created                   *types.Date `json:"Created"`
	Creator                   types.Guid  `json:"Creator"`
	CreatorFullName           string      `json:"CreatorFullName"`
	Division                  int32       `json:"Division"`
	Email                     string      `json:"Email"`
	EndDate                   *types.Date `json:"EndDate"`
	FirstName                 string      `json:"FirstName"`
	FullName                  string      `json:"FullName"`
	Gender                    string      `json:"Gender"`
	HID                       int32       `json:"HID"`
	IdentificationDate        *types.Date `json:"IdentificationDate"`
	IdentificationDocument    types.Guid  `json:"IdentificationDocument"`
	IdentificationUser        types.Guid  `json:"IdentificationUser"`
	Initials                  string      `json:"Initials"`
	IsAnonymised              byte        `json:"IsAnonymised"`
	IsMailingExcluded         bool        `json:"IsMailingExcluded"`
	IsMainContact             bool        `json:"IsMainContact"`
	JobTitleDescription       string      `json:"JobTitleDescription"`
	Language                  string      `json:"Language"`
	LastName                  string      `json:"LastName"`
	LeadPurpose               types.Guid  `json:"LeadPurpose"`
	LeadSource                types.Guid  `json:"LeadSource"`
	MarketingNotes            string      `json:"MarketingNotes"`
	MiddleName                string      `json:"MiddleName"`
	Mobile                    string      `json:"Mobile"`
	Modified                  *types.Date `json:"Modified"`
	Modifier                  types.Guid  `json:"Modifier"`
	ModifierFullName          string      `json:"ModifierFullName"`
	Nationality               string      `json:"Nationality"`
	Notes                     string      `json:"Notes"`
	PartnerName               string      `json:"PartnerName"`
	PartnerNamePrefix         string      `json:"PartnerNamePrefix"`
	Person                    types.Guid  `json:"Person"`
	Phone                     string      `json:"Phone"`
	PhoneExtension            string      `json:"PhoneExtension"`
	PictureName               string      `json:"PictureName"`
	PictureThumbnailUrl       string      `json:"PictureThumbnailUrl"`
	PictureUrl                string      `json:"PictureUrl"`
	Postcode                  string      `json:"Postcode"`
	SocialSecurityNumber      string      `json:"SocialSecurityNumber"`
	StartDate                 *types.Date `json:"StartDate"`
	State                     string      `json:"State"`
	Title                     string      `json:"Title"`
}

// ContactUpdate stores Contact value to insert/update
//
type ContactUpdate struct {
	Account                   *string `json:"Account,omitempty"`
	AccountMainContact        *string `json:"AccountMainContact,omitempty"`
	AddressLine2              *string `json:"AddressLine2,omitempty"`
	AddressStreet             *string `json:"AddressStreet,omitempty"`
	AddressStreetNumber       *string `json:"AddressStreetNumber,omitempty"`
	AddressStreetNumberSuffix *string `json:"AddressStreetNumberSuffix,omitempty"`
	AllowMailing              *int32  `json:"AllowMailing,omitempty"`
	BirthDate                 *string `json:"BirthDate,omitempty"`
	BirthName                 *string `json:"BirthName,omitempty"`
	BirthNamePrefix           *string `json:"BirthNamePrefix,omitempty"`
	BirthPlace                *string `json:"BirthPlace,omitempty"`
	BusinessEmail             *string `json:"BusinessEmail,omitempty"`
	BusinessFax               *string `json:"BusinessFax,omitempty"`
	BusinessMobile            *string `json:"BusinessMobile,omitempty"`
	BusinessPhone             *string `json:"BusinessPhone,omitempty"`
	BusinessPhoneExtension    *string `json:"BusinessPhoneExtension,omitempty"`
	City                      *string `json:"City,omitempty"`
	Code                      *string `json:"Code,omitempty"`
	Country                   *string `json:"Country,omitempty"`
	Email                     *string `json:"Email,omitempty"`
	EndDate                   *string `json:"EndDate,omitempty"`
	FirstName                 *string `json:"FirstName,omitempty"`
	FullName                  *string `json:"FullName,omitempty"`
	Gender                    *string `json:"Gender,omitempty"`
	HID                       *int32  `json:"HID,omitempty"`
	IdentificationDate        *string `json:"IdentificationDate,omitempty"`
	IdentificationDocument    *string `json:"IdentificationDocument,omitempty"`
	IdentificationUser        *string `json:"IdentificationUser,omitempty"`
	Initials                  *string `json:"Initials,omitempty"`
	IsAnonymised              *byte   `json:"IsAnonymised,omitempty"`
	IsMailingExcluded         *bool   `json:"IsMailingExcluded,omitempty"`
	IsMainContact             *bool   `json:"IsMainContact,omitempty"`
	JobTitleDescription       *string `json:"JobTitleDescription,omitempty"`
	Language                  *string `json:"Language,omitempty"`
	LastName                  *string `json:"LastName,omitempty"`
	LeadPurpose               *string `json:"LeadPurpose,omitempty"`
	LeadSource                *string `json:"LeadSource,omitempty"`
	MarketingNotes            *string `json:"MarketingNotes,omitempty"`
	MiddleName                *string `json:"MiddleName,omitempty"`
	Mobile                    *string `json:"Mobile,omitempty"`
	Nationality               *string `json:"Nationality,omitempty"`
	Notes                     *string `json:"Notes,omitempty"`
	PartnerName               *string `json:"PartnerName,omitempty"`
	PartnerNamePrefix         *string `json:"PartnerNamePrefix,omitempty"`
	Person                    *string `json:"Person,omitempty"`
	Phone                     *string `json:"Phone,omitempty"`
	PhoneExtension            *string `json:"PhoneExtension,omitempty"`
	PictureName               *string `json:"PictureName,omitempty"`
	PictureThumbnailUrl       *string `json:"PictureThumbnailUrl,omitempty"`
	PictureUrl                *string `json:"PictureUrl,omitempty"`
	Postcode                  *string `json:"Postcode,omitempty"`
	SocialSecurityNumber      *string `json:"SocialSecurityNumber,omitempty"`
	StartDate                 *string `json:"StartDate,omitempty"`
	State                     *string `json:"State,omitempty"`
	Title                     *string `json:"Title,omitempty"`
}

type GetContactsCall struct {
	urlNext string
	service *Service
}

type GetContactsCallParams struct {
	Account       *types.Guid
	Email         *string
	FullName      *string
	ModifiedAfter *time.Time
}

func (service *Service) NewGetContactsCall(params *GetContactsCallParams) *GetContactsCall {
	call := GetContactsCall{}
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", Contact{})
	call.urlNext = service.url(fmt.Sprintf("Contacts?$select=%s", selectFields))
	filter := []string{}

	if params != nil {
		if params.Account != nil {
			filter = append(filter, fmt.Sprintf("Account eq guid'%s'", (*params.Account).String()))
		}
		if params.Email != nil {
			filter = append(filter, fmt.Sprintf("Email eq '%s'", *params.Email))
		}
		if params.FullName != nil {
			filter = append(filter, fmt.Sprintf("FullName eq '%s'", *params.FullName))
		}
		if params.ModifiedAfter != nil {
			filter = append(filter, service.DateFilter("Modified", "gt", params.ModifiedAfter, false, ""))
		}
	}

	if len(filter) > 0 {
		call.urlNext = fmt.Sprintf("%s&$filter=%s", call.urlNext, strings.Join(filter, " and "))
	}

	return &call
}

func (call *GetContactsCall) Do() (*[]Contact, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	contacts := []Contact{}

	next, err := call.service.Get(call.urlNext, &contacts)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &contacts, nil
}

func (call *GetContactsCall) DoAll() (*[]Contact, *errortools.Error) {
	contacts := []Contact{}

	for true {
		_contacts, e := call.Do()
		if e != nil {
			return nil, e
		}

		if _contacts == nil {
			break
		}

		if len(*_contacts) == 0 {
			break
		}

		contacts = append(contacts, *_contacts...)
	}

	return &contacts, nil
}

func (service *Service) GetContact(id types.Guid) (*Contact, *errortools.Error) {
	url := service.url(fmt.Sprintf("Contacts(guid'%s')", id.String()))

	contactNew := Contact{}

	e := service.GetSingle(url, &contactNew)
	if e != nil {
		return nil, e
	}
	return &contactNew, nil
}

func (service *Service) CreateContact(contact *ContactUpdate) (*Contact, *errortools.Error) {
	url := service.url("Contacts")

	contactNew := Contact{}

	e := service.Post(url, contact, &contactNew)
	if e != nil {
		return nil, e
	}
	return &contactNew, nil
}

func (service *Service) UpdateContact(id types.Guid, contact *ContactUpdate, returnUpdated bool) (*Contact, *errortools.Error) {
	requestConfig := go_http.RequestConfig{
		Url:       service.url(fmt.Sprintf("Contacts(guid'%s')", id.String())),
		BodyModel: contact,
	}

	e := service.Put(&requestConfig)
	if e != nil {
		return nil, e
	}

	if !returnUpdated {
		return nil, nil
	}

	contactUpdated, e := service.GetContact(id)
	if e != nil {
		return nil, e
	}

	return contactUpdated, nil
}

func (service *Service) DeleteContact(id types.Guid) *errortools.Error {
	url := service.url(fmt.Sprintf("Contacts(guid'%s')", id.String()))

	err := service.Delete(url)
	if err != nil {
		return err
	}
	return nil
}

func (service *Service) GetContactsCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("Contacts", createdBefore)
}
