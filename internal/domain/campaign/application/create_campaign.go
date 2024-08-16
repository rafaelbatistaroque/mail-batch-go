package application

import (
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/contract"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/entity"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/main/service"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/usecase/createcampaign"
)

type createCampaign struct {
	//dependencias
	repository contract.Repository
	mapper     service.Mapper
}

func NewCreateCampaignUseCase(repository contract.Repository, mapper service.Mapper) createcampaign.UseCase {
	return &createCampaign{
		repository: repository,
		mapper:     mapper,
	}
}

func (uc *createCampaign) Execute(input *createcampaign.Input) (*createcampaign.Output, error) {
	input.Validate()
	if input.IsInvalid() {
		return nil, input.GetErrors()
	}

	campaign, _ := entity.MakeCampaign(input.Name, input.Content, input.Contacts)

	model := uc.mapper.ToCampaignModel(campaign)

	err := uc.repository.Save(model)
	if err != nil {
		return nil, err
	}

	return &createcampaign.Output{
		Id:       model.Id,
		Name:     model.Name,
		Content:  model.Content,
		Contacts: model.Contacts,
	}, nil
}
