package application

import (
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/contract"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/main/service"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/usecase/cancelcampaign"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/error/domainError"
)

type cancelCampaign struct {
	//dependencias
	repository contract.Repository
	mapper     service.Mapper
}

func NewCancelCampaignUseCase(repository contract.Repository, mapper service.Mapper) cancelcampaign.UseCase {
	return &cancelCampaign{
		repository: repository,
		mapper:     mapper,
	}
}

func (uc *cancelCampaign) Execute(input *cancelcampaign.Input) (*cancelcampaign.Output, error) {
	input.Validate()
	if input.IsInvalid() {
		return nil, input.GetErrors()
	}

	campaignModel, err := uc.repository.GetById(input.Id)
	if err != nil {
		return nil, domainError.Err_CAMPAIGN_NOT_FOUND
	}

	campaign := uc.mapper.ToEntity(campaignModel)
	campaign.Cancel()

	modelCanceled := uc.mapper.ToCampaignModel(campaign)

	uc.repository.Update(modelCanceled)

	return nil, nil
}
