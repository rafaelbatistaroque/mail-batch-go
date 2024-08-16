package createcampaign

import (
	"github.com/rafaelbatistaroque/mail-batch-go/internal/error/domainError"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/pkg/helpers/validation"
)

type Input = createCampaignInput

type createCampaignInput struct {
	validation.InputValidation
	Name     string
	Content  string
	Contacts []string
}

func (i *createCampaignInput) Validate() {
	if validation.IsNilOrEmpty(i.Name) {
		i.AppendError("Name", domainError.Err_PARAMETER_NOT_EMPTY.Error())
	}

	if validation.IsNilOrEmpty(i.Content) {
		i.AppendError("Content", domainError.Err_PARAMETER_NOT_EMPTY.Error())
	}

	if validation.IsNilOrEmpty(i.Contacts) {
		i.AppendError("Contacts", domainError.Err_PARAMETER_NOT_EMPTY.Error())
	}
}
