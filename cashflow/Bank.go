package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// Bank stores Bank from exactonline
//
type Bank struct {
	ID              types.Guid  `json:"ID"`
	BankName        string      `json:"BankName"`
	BICCode         string      `json:"BICCode"`
	Country         string      `json:"Country"`
	Created         *types.Date `json:"Created"`
	Description     string      `json:"Description"`
	Format          string      `json:"Format"`
	HomePageAddress string      `json:"HomePageAddress"`
	Modified        *types.Date `json:"Modified"`
	Status          string      `json:"Status"`
}

type GetBanksCall struct {
	modifiedAfter *time.Time
	urlNext       string
	service       *Service
}

func (service *Service) NewGetBanksCall(modifiedAfter *time.Time) *GetBanksCall {
	call := GetBanksCall{}
	call.modifiedAfter = modifiedAfter
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", Bank{})
	call.urlNext = service.url(fmt.Sprintf("Banks?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetBanksCall) Do() (*[]Bank, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	banks := []Bank{}

	next, err := call.service.Get(call.urlNext, &banks)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &banks, nil
}

func (service *Service) GetBanksCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("Banks", createdBefore)
}
