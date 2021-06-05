package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// RevenueList stores RevenueList from exactonline
//
type RevenueList struct {
	Period int32   `json:"Period"`
	Year   int32   `json:"Year"`
	Amount float64 `json:"Amount"`
}

type GetRevenueListsCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewGetRevenueListsCall() *GetRevenueListsCall {
	call := GetRevenueListsCall{}
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", RevenueList{})
	call.urlNext = service.urlRead(fmt.Sprintf("RevenueList?$select=%s", selectFields))

	return &call
}

func (call *GetRevenueListsCall) Do() (*[]RevenueList, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	transactionLines := []RevenueList{}

	next, err := call.service.Get(call.urlNext, &transactionLines)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &transactionLines, nil
}
