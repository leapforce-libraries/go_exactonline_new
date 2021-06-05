package exactonline

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
	utilities "github.com/leapforce-libraries/go_utilities"
)

// GLAccountClassificationMapping stores GLAccountClassificationMapping from exactonline
//
type GLAccountClassificationMapping struct {
	ID                        types.GUID `json:"ID"`
	Classification            types.GUID `json:"Classification"`
	ClassificationCode        string     `json:"ClassificationCode"`
	ClassificationDescription string     `json:"ClassificationDescription"`
	Division                  int64      `json:"Division"`
	GLAccount                 types.GUID `json:"GLAccount"`
	GLAccountCode             string     `json:"GLAccountCode"`
	GLAccountDescription      string     `json:"GLAccountDescription"`
	GLSchemeCode              string     `json:"GLSchemeCode"`
	GLSchemeDescription       string     `json:"GLSchemeDescription"`
	GLSchemeID                types.GUID `json:"GLSchemeID"`
}

type GetGLAccountClassificationMappingsCall struct {
	urlNext string
	service *Service
}

func (service *Service) NewGetGLAccountClassificationMappingsCall() *GetGLAccountClassificationMappingsCall {
	call := GetGLAccountClassificationMappingsCall{}
	call.service = service

	selectFields := utilities.GetTaggedTagNames("json", GLAccountClassificationMapping{})
	call.urlNext = service.url(fmt.Sprintf("GLAccountClassificationMappings?$select=%s", selectFields))

	return &call
}

func (call *GetGLAccountClassificationMappingsCall) Do() (*[]GLAccountClassificationMapping, *errortools.Error) {
	if call.urlNext == "" {
		return nil, nil
	}

	transactionLines := []GLAccountClassificationMapping{}

	next, err := call.service.Get(call.urlNext, &transactionLines)
	if err != nil {
		return nil, err
	}

	call.urlNext = next

	return &transactionLines, nil
}
