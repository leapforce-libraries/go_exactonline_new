package exactonline

import (
	"fmt"
	"strings"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// Address stores Address from exactonline
//
type Address struct {
	ID                   types.GUID  `json:"ID"`
	Account              types.GUID  `json:"Account"`
	AccountIsSupplier    bool        `json:"AccountIsSupplier"`
	AccountName          string      `json:"AccountName"`
	AddressLine1         string      `json:"AddressLine1"`
	AddressLine2         string      `json:"AddressLine2"`
	AddressLine3         string      `json:"AddressLine3"`
	City                 string      `json:"City"`
	Contact              types.GUID  `json:"Contact"`
	ContactName          string      `json:"ContactName"`
	Country              string      `json:"Country"`
	CountryName          string      `json:"CountryName"`
	Created              *types.Date `json:"Created"`
	Creator              types.GUID  `json:"Creator"`
	CreatorFullName      string      `json:"CreatorFullName"`
	Division             int32       `json:"Division"`
	Fax                  string      `json:"Fax"`
	FreeBoolField01      bool        `json:"FreeBoolField_01"`
	FreeBoolField02      bool        `json:"FreeBoolField_02"`
	FreeBoolField03      bool        `json:"FreeBoolField_03"`
	FreeBoolField04      bool        `json:"FreeBoolField_04"`
	FreeBoolField05      bool        `json:"FreeBoolField_05"`
	FreeDateField01      *types.Date `json:"FreeDateField_01"`
	FreeDateField02      *types.Date `json:"FreeDateField_02"`
	FreeDateField03      *types.Date `json:"FreeDateField_03"`
	FreeDateField04      *types.Date `json:"FreeDateField_04"`
	FreeDateField05      *types.Date `json:"FreeDateField_05"`
	FreeNumberField01    float64     `json:"FreeNumberField_01"`
	FreeNumberField02    float64     `json:"FreeNumberField_02"`
	FreeNumberField03    float64     `json:"FreeNumberField_03"`
	FreeNumberField04    float64     `json:"FreeNumberField_04"`
	FreeNumberField05    float64     `json:"FreeNumberField_05"`
	FreeTextField01      string      `json:"FreeTextField_01"`
	FreeTextField02      string      `json:"FreeTextField_02"`
	FreeTextField03      string      `json:"FreeTextField_03"`
	FreeTextField04      string      `json:"FreeTextField_04"`
	FreeTextField05      string      `json:"FreeTextField_05"`
	Mailbox              string      `json:"Mailbox"`
	Main                 bool        `json:"Main"`
	Modified             *types.Date `json:"Modified"`
	Modifier             types.GUID  `json:"Modifier"`
	ModifierFullName     string      `json:"ModifierFullName"`
	NicNumber            string      `json:"NicNumber"`
	Notes                string      `json:"Notes"`
	Phone                string      `json:"Phone"`
	PhoneExtension       string      `json:"PhoneExtension"`
	Postcode             string      `json:"Postcode"`
	State                string      `json:"State"`
	StateDescription     string      `json:"StateDescription"`
	Type                 int16       `json:"Type"`
	Warehouse            types.GUID  `json:"Warehouse"`
	WarehouseCode        string      `json:"WarehouseCode"`
	WarehouseDescription string      `json:"WarehouseDescription"`
}

type GetAddressesCall struct {
	urlNext string
	service *Service
}

type GetAddressesCallParams struct {
	ModifiedAfter *time.Time
}

func (service *Service) NewGetAddressesCall(params *GetAddressesCallParams) *GetAddressesCall {
	call := GetAddressesCall{}
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", Address{})
	call.urlNext = fmt.Sprintf("%s/Addresses?$select=%s", service.BaseURL(), selectFields)
	filter := []string{}

	if params != nil {
		if params.ModifiedAfter != nil {
			filter = append(filter, service.DateFilter("Modified", "gt", params.ModifiedAfter, false, ""))
		}
	}

	if len(filter) > 0 {
		call.urlNext = fmt.Sprintf("%s&$filter=%s", call.urlNext, strings.Join(filter, " and "))
	}

	return &call
}

func (call *GetAddressesCall) Do() (*[]Address, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	addresses := []Address{}

	next, err := call.service.Get(call.urlNext, &addresses)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &addresses, nil
}

func (call *GetAddressesCall) DoAll() (*[]Address, *errortools.Error) {
	addresses := []Address{}

	for true {
		_addresses, e := call.Do()
		if e != nil {
			return nil, e
		}

		if _addresses == nil {
			break
		}

		if len(*_addresses) == 0 {
			break
		}

		addresses = append(addresses, *_addresses...)
	}

	return &addresses, nil
}

func (service *Service) GetAddress(id types.GUID) (*Address, *errortools.Error) {
	url := fmt.Sprintf("%s/Addresses(guid'%s')", service.BaseURL(), id.String())

	addressNew := Address{}

	_, e := service.Get(url, &addressNew)
	if e != nil {
		return nil, e
	}
	return &addressNew, nil
}

func (service *Service) GetAddressesCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("Addresses", createdBefore)
}
