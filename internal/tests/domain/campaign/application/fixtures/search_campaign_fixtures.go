package fixtures

import (
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/application"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/model"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/usecase/searchcampaign"
)

func MakeSearchCampaignInput_Valid() *searchcampaign.Input {
	return &searchcampaign.Input{
		Page:    1,
		PerPage: 10,
	}
}

func MakeSearchCampaignModel_Valid() *model.SearchCampaignModel {
	return &model.SearchCampaignModel{
		Page:    1,
		PerPage: 10,
	}
}

func NewSearchCampaignUseCaseSUT() (searchcampaign.UseCase, *MapperSpy, *RepositorySpy) {
	repositorySpy := MakeRepositorySpy()
	mapperSpy := MakeMapperSpy()
	sut := application.NewSearchCampaignUseCase(repositorySpy, mapperSpy)
	return sut, mapperSpy, repositorySpy
}
