package fixtures

import (
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/application"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/usecase/cancelcampaign"
)

func NewCancelCampaignSUT() (cancelcampaign.UseCase, *MapperSpy, *RepositorySpy) {
	mapperSpy := MakeMapperSpy()
	repositorySpy := MakeRepositorySpy()
	sut := application.NewCancelCampaignUseCase(repositorySpy, mapperSpy)

	return sut, mapperSpy, repositorySpy
}

func MakeCancelCampaignInput_Valid() *cancelcampaign.Input {
	return &cancelcampaign.Input{
		Id: "fake_name",
	}
}
