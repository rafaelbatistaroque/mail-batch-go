package application

import (
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/contract"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/main/service"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/usecase/searchcampaign"
)

type searchCampaign struct {
	//dependencias
	repository contract.Repository
	mapper     service.Mapper
}

func NewSearchCampaignUseCase(repository contract.Repository, mapper service.Mapper) searchcampaign.UseCase {
	return &searchCampaign{
		repository: repository,
		mapper:     mapper,
	}
}

func (uc *searchCampaign) Execute(input *searchcampaign.Input) (*searchcampaign.Output, error) {
	input.Validate()
	if input.IsInvalid() {
		return nil, input.GetErrors()
	}

	model := uc.mapper.ToSearchCampaignModel(input)

	result, err := uc.repository.Search(model)
	if err != nil {
		return nil, err
	}

	output := uc.mapper.ToSearchCampaignOutput(result)

	return output, nil
}
