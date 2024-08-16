package application

import (
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/contract"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/usecase/getcampaignbyid"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/error/domainError"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/pkg/helpers/validation"
)

type getCampaignById struct {
	//dependencias
	repository contract.Repository
}

func NewGetCampaignByIdUseCase(repository contract.Repository) getcampaignbyid.UseCase {
	return &getCampaignById{
		repository: repository,
	}
}

func (uc *getCampaignById) Execute(id string) (*getcampaignbyid.Output, error) {
	if validation.IsNilOrEmpty(id) {
		return nil, domainError.Err_PARAMETER_ID_NOT_EMPTY
	}

	model, err := uc.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	return &getcampaignbyid.Output{
		Id:       model.Id,
		Name:     model.Name,
		Content:  model.Content,
		Status:   model.Status,
		Contacts: model.Contacts,
	}, nil
}
