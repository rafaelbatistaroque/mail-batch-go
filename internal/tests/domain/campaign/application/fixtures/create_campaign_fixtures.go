package fixtures

import (
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/application"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/usecase/createcampaign"
)

func MakeCreateCampaignInput_Valid() *createcampaign.Input {
	return &createcampaign.Input{
		Name:     "fake_name",
		Content:  "fake_content",
		Contacts: []string{"t@t.com"},
	}
}

func NewCreateCampaignUseCaseSUT() (createcampaign.UseCase, *MapperSpy, *RepositorySpy) {
	mapperSpy := MakeMapperSpy()
	repositorySpy := MakeRepositorySpy()
	sut := application.NewCreateCampaignUseCase(repositorySpy, mapperSpy)

	return sut, mapperSpy, repositorySpy
}
