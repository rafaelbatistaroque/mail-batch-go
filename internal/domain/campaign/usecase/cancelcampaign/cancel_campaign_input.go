package cancelcampaign

import (
	"github.com/rafaelbatistaroque/mail-batch-go/internal/error/domainError"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/pkg/helpers/validation"
)

type Input = cancelCampaignInput

type cancelCampaignInput struct {
	validation.InputValidation
	Id string
}

func (i *cancelCampaignInput) Validate() {
	if validation.IsNilOrEmpty(i.Id) {
		i.AppendError("Id", domainError.Err_PARAMETER_NOT_EMPTY.Error())
	}
}
