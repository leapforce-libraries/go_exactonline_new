package exactonline

import (
	"fmt"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// PurchaseEntryLine stores PurchaseEntryLine from exactonline
//
type PurchaseEntryLine struct {
	ID                         types.GUID  `json:"ID"`
	AmountDC                   float64     `json:"AmountDC"`
	AmountFC                   float64     `json:"AmountFC"`
	Asset                      types.GUID  `json:"Asset"`
	AssetDescription           string      `json:"AssetDescription"`
	CostCenter                 string      `json:"CostCenter"`
	CostCenterDescription      string      `json:"CostCenterDescription"`
	CostUnit                   string      `json:"CostUnit"`
	CostUnitDescription        string      `json:"CostUnitDescription"`
	Description                string      `json:"Description"`
	Division                   int32       `json:"Division"`
	EntryID                    types.GUID  `json:"EntryID"`
	From                       *types.Date `json:"From"`
	GLAccount                  types.GUID  `json:"GLAccount"`
	GLAccountCode              string      `json:"GLAccountCode"`
	GLAccountDescription       string      `json:"GLAccountDescription"`
	IntraStatArea              string      `json:"IntraStatArea"`
	IntraStatCountry           string      `json:"IntraStatCountry"`
	IntraStatDeliveryTerm      string      `json:"IntraStatDeliveryTerm"`
	IntraStatTransactionA      string      `json:"IntraStatTransactionA"`
	IntraStatTransactionB      string      `json:"IntraStatTransactionB"`
	IntraStatTransportMethod   string      `json:"IntraStatTransportMethod"`
	LineNumber                 int32       `json:"LineNumber"`
	Notes                      string      `json:"Notes"`
	PrivateUsePercentage       float64     `json:"PrivateUsePercentage"`
	Project                    types.GUID  `json:"Project"`
	ProjectDescription         string      `json:"ProjectDescription"`
	Quantity                   float64     `json:"Quantity"`
	SerialNumber               string      `json:"SerialNumber"`
	StatisticalNetWeight       float64     `json:"StatisticalNetWeight"`
	StatisticalNumber          string      `json:"StatisticalNumber"`
	StatisticalQuantity        float64     `json:"StatisticalQuantity"`
	StatisticalValue           float64     `json:"StatisticalValue"`
	Subscription               types.GUID  `json:"Subscription"`
	SubscriptionDescription    string      `json:"SubscriptionDescription"`
	To                         *types.Date `json:"To"`
	TrackingNumber             types.GUID  `json:"TrackingNumber"`
	TrackingNumberDescription  string      `json:"TrackingNumberDescription"`
	Type                       int32       `json:"Type"`
	VATAmountDC                float64     `json:"VATAmountDC"`
	VATAmountFC                float64     `json:"VATAmountFC"`
	VATBaseAmountDC            float64     `json:"VATBaseAmountDC"`
	VATBaseAmountFC            float64     `json:"VATBaseAmountFC"`
	VATCode                    string      `json:"VATCode"`
	VATCodeDescription         string      `json:"VATCodeDescription"`
	VATNonDeductiblePercentage float64     `json:"VATNonDeductiblePercentage"`
	VATPercentage              float64     `json:"VATPercentage"`
	WithholdingAmountDC        float64     `json:"WithholdingAmountDC"`
	WithholdingTax             string      `json:"WithholdingTax"`
}

type GetPurchaseEntryLinesCall struct {
	modifiedAfter *time.Time
	urlNext       string
	service       *Service
}

func (service *Service) NewGetPurchaseEntryLinesCall(modifiedAfter *time.Time) *GetPurchaseEntryLinesCall {
	call := GetPurchaseEntryLinesCall{}
	call.modifiedAfter = modifiedAfter
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", PurchaseEntryLine{})
	call.urlNext = service.url(fmt.Sprintf("PurchaseEntryLines?$select=%s", selectFields))
	if modifiedAfter != nil {
		call.urlNext += service.DateFilter("Modified", "gt", modifiedAfter, true, "&")
	}

	return &call
}

func (call *GetPurchaseEntryLinesCall) Do() (*[]PurchaseEntryLine, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	purchaseEntryLines := []PurchaseEntryLine{}

	next, err := call.service.Get(call.urlNext, &purchaseEntryLines)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &purchaseEntryLines, nil
}

func (service *Service) GetPurchaseEntryLinesCount(createdBefore *time.Time) (int64, *errortools.Error) {
	return service.GetCount("PurchaseEntryLines", createdBefore)
}
