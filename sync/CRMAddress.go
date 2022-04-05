package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// CRMAddress stores CRMAddress from exactonline
//
type CRMAddress struct {
	Timestamp            types.Int64String `json:"Timestamp"`
	ID                   types.Guid        `json:"ID"`
	Account              types.Guid        `json:"Account"`
	AccountIsSupplier    bool              `json:"AccountIsSupplier"`
	AccountName          string            `json:"AccountName"`
	AddressLine1         string            `json:"AddressLine1"`
	AddressLine2         string            `json:"AddressLine2"`
	AddressLine3         string            `json:"AddressLine3"`
	City                 string            `json:"City"`
	Contact              types.Guid        `json:"Contact"`
	ContactName          string            `json:"ContactName"`
	Country              string            `json:"Country"`
	CountryName          string            `json:"CountryName"`
	Created              *types.Date       `json:"Created"`
	Creator              types.Guid        `json:"Creator"`
	CreatorFullName      string            `json:"CreatorFullName"`
	Division             int32             `json:"Division"`
	Fax                  string            `json:"Fax"`
	FreeBoolField01      bool              `json:"FreeBoolField_01"`
	FreeBoolField02      bool              `json:"FreeBoolField_02"`
	FreeBoolField03      bool              `json:"FreeBoolField_03"`
	FreeBoolField04      bool              `json:"FreeBoolField_04"`
	FreeBoolField05      bool              `json:"FreeBoolField_05"`
	FreeDateField01      *types.Date       `json:"FreeDateField_01"`
	FreeDateField02      *types.Date       `json:"FreeDateField_02"`
	FreeDateField03      *types.Date       `json:"FreeDateField_03"`
	FreeDateField04      *types.Date       `json:"FreeDateField_04"`
	FreeDateField05      *types.Date       `json:"FreeDateField_05"`
	FreeNumberField01    float64           `json:"FreeNumberField_01"`
	FreeNumberField02    float64           `json:"FreeNumberField_02"`
	FreeNumberField03    float64           `json:"FreeNumberField_03"`
	FreeNumberField04    float64           `json:"FreeNumberField_04"`
	FreeNumberField05    float64           `json:"FreeNumberField_05"`
	FreeTextField01      string            `json:"FreeTextField_01"`
	FreeTextField02      string            `json:"FreeTextField_02"`
	FreeTextField03      string            `json:"FreeTextField_03"`
	FreeTextField04      string            `json:"FreeTextField_04"`
	FreeTextField05      string            `json:"FreeTextField_05"`
	Mailbox              string            `json:"Mailbox"`
	Main                 bool              `json:"Main"`
	Modified             *types.Date       `json:"Modified"`
	Modifier             types.Guid        `json:"Modifier"`
	ModifierFullName     string            `json:"ModifierFullName"`
	NicNumber            string            `json:"NicNumber"`
	Notes                string            `json:"Notes"`
	Phone                string            `json:"Phone"`
	PhoneExtension       string            `json:"PhoneExtension"`
	Postcode             string            `json:"Postcode"`
	State                string            `json:"State"`
	StateDescription     string            `json:"StateDescription"`
	Type                 int16             `json:"Type"`
	Warehouse            types.Guid        `json:"Warehouse"`
	WarehouseCode        string            `json:"WarehouseCode"`
	WarehouseDescription string            `json:"WarehouseDescription"`
}

type SyncCRMAddresssCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewSyncCRMAddresssCall(timestamp *int64) *SyncCRMAddresssCall {
	selectFields := utilities.GetTaggedTagNames("json", CRMAddress{})
	url := service.url(fmt.Sprintf("CRM/Addresses?$select=%s", selectFields))
	if timestamp != nil {
		url = fmt.Sprintf("%s&$filter=Timestamp gt %vL", url, *timestamp)
	}

	return &SyncCRMAddresssCall{
		service: service,
		urlNext: url,
	}
}

func (call *SyncCRMAddresssCall) Do() (*[]CRMAddress, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	crmAddresss := []CRMAddress{}

	next, err := call.service.Get(call.urlNext, &crmAddresss)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &crmAddresss, nil
}
