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

// Contact stores Contact from exactonline
//
type Contact struct {
	ID                        types.GUID      `json:"ID"`
	Account                   types.GUID      `json:"Account"`
	AccountIsCustomer         bool            `json:"AccountIsCustomer"`
	AccountIsSupplier         bool            `json:"AccountIsSupplier"`
	AccountMainContact        types.GUID      `json:"AccountMainContact"`
	AccountName               string          `json:"AccountName"`
	AddressLine2              string          `json:"AddressLine2"`
	AddressStreet             string          `json:"AddressStreet"`
	AddressStreetNumber       string          `json:"AddressStreetNumber"`
	AddressStreetNumberSuffix string          `json:"AddressStreetNumberSuffix"`
	AllowMailing              int32           `json:"AllowMailing"`
	BirthDate                 *types.Date     `json:"BirthDate"`
	BirthName                 string          `json:"BirthName"`
	BirthNamePrefix           string          `json:"BirthNamePrefix"`
	BirthPlace                string          `json:"BirthPlace"`
	BusinessEmail             string          `json:"BusinessEmail"`
	BusinessFax               string          `json:"BusinessFax"`
	BusinessMobile            string          `json:"BusinessMobile"`
	BusinessPhone             string          `json:"BusinessPhone"`
	BusinessPhoneExtension    string          `json:"BusinessPhoneExtension"`
	City                      string          `json:"City"`
	Code                      string          `json:"Code"`
	Country                   string          `json:"Country"`
	Created                   *types.Date     `json:"Created"`
	Creator                   types.GUID      `json:"Creator"`
	CreatorFullName           string          `json:"CreatorFullName"`
	Division                  int32           `json:"Division"`
	Email                     string          `json:"Email"`
	EndDate                   *types.Date     `json:"EndDate"`
	FirstName                 string          `json:"FirstName"`
	FullName                  string          `json:"FullName"`
	Gender                    string          `json:"Gender"`
	HID                       int32           `json:"HID"`
	IdentificationDate        *types.Date     `json:"IdentificationDate"`
	IdentificationDocument    types.GUID      `json:"IdentificationDocument"`
	IdentificationUser        types.GUID      `json:"IdentificationUser"`
	Initials                  string          `json:"Initials"`
	IsAnonymised              byte            `json:"IsAnonymised"`
	IsMailingExcluded         bool            `json:"IsMailingExcluded"`
	IsMainContact             bool            `json:"IsMainContact"`
	JobTitleDescription       string          `json:"JobTitleDescription"`
	Language                  string          `json:"Language"`
	LastName                  string          `json:"LastName"`
	LeadPurpose               types.GUID      `json:"LeadPurpose"`
	LeadSource                types.GUID      `json:"LeadSource"`
	MarketingNotes            string          `json:"MarketingNotes"`
	MiddleName                string          `json:"MiddleName"`
	Mobile                    string          `json:"Mobile"`
	Modified                  *types.Date     `json:"Modified"`
	Modifier                  types.GUID      `json:"Modifier"`
	ModifierFullName          string          `json:"ModifierFullName"`
	Nationality               string          `json:"Nationality"`
	Notes                     string          `json:"Notes"`
	PartnerName               string          `json:"PartnerName"`
	PartnerNamePrefix         string          `json:"PartnerNamePrefix"`
	Person                    types.GUID      `json:"Person"`
	Phone                     string          `json:"Phone"`
	PhoneExtension            string          `json:"PhoneExtension"`
	Picture                   json.RawMessage `json:"Picture"`
	PictureName               string          `json:"PictureName"`
	PictureThumbnailUrl       string          `json:"PictureThumbnailUrl"`
	PictureUrl                string          `json:"PictureUrl"`
	Postcode                  string          `json:"Postcode"`
	SocialSecurityNumber      string          `json:"SocialSecurityNumber"`
	StartDate                 *types.Date     `json:"StartDate"`
	State                     string          `json:"State"`
	Title                     string          `json:"Title"`
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
	client  *Client
}

type GetContactsCallParams struct {
	Account       *types.GUID
	Email         *string
	FullName      *string
	ModifiedAfter *time.Time
}

func (c *Client) NewGetContactsCall(params GetContactsCallParams) *GetContactsCall {
	call := GetContactsCall{}
	call.client = c

	selectFields := utilities.GetTaggedFieldNames("json", Contact{})
	call.urlNext = fmt.Sprintf("%s/Contacts?$select=%s", c.BaseURL(), selectFields)
	filter := []string{}

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
		filter = append(filter, c.DateFilter("Modified", "gt", params.ModifiedAfter, true, "&"))
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

	next, err := call.client.Get(call.urlNext, &contacts)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &contacts, nil
}

func (c *Client) CreateContact(contact *ContactUpdate) (*Contact, *errortools.Error) {
	url := fmt.Sprintf("%s/Contacts", c.BaseURL())

	b, err := json.Marshal(contact)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	contactNew := Contact{}

	e := c.Post(url, bytes.NewBuffer(b), &contactNew)
	if e != nil {
		return nil, e
	}
	return &contactNew, nil
}

func (c *Client) UpdateContact(id types.GUID, contact *ContactUpdate) *errortools.Error {
	url := fmt.Sprintf("%s/Contacts(guid'%s')", c.BaseURL(), id.String())

	b, err := json.Marshal(contact)
	if err != nil {
		return errortools.ErrorMessage(err)
	}

	e := c.Put(url, bytes.NewBuffer(b))
	if e != nil {
		return e
	}
	return nil
}

func (c *Client) DeleteContact(id types.GUID) *errortools.Error {
	url := fmt.Sprintf("%s/Contacts(guid'%s')", c.BaseURL(), id.String())

	err := c.Delete(url)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetContactsCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return c.GetCount("Contacts", createdBefore)
}
