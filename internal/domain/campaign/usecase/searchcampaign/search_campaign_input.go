package searchcampaign

import (
	"github.com/rafaelbatistaroque/mail-batch-go/internal/error/domainError"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/pkg/helpers/validation"
)

type Input = searchCampaignInput

type searchCampaignInput struct {
	validation.InputValidation
	Page    int
	PerPage int
}

func (i *searchCampaignInput) Validate() {
	if i.Page <= 0 {
		i.AppendError("Page", domainError.Err_PROPERTIES_NOT_LESS_THAN_OR_EQUAL_ZERO.Error())
	}

	if i.PerPage <= 0 {
		i.AppendError("PerPage", domainError.Err_PROPERTIES_NOT_LESS_THAN_OR_EQUAL_ZERO.Error())
	}
}
