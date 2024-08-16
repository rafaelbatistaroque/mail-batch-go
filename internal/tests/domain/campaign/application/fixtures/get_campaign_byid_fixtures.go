package fixtures

import (
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/application"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/usecase/getcampaignbyid"
)

func NewGetCampaignByIdUseCaseSUT() (getcampaignbyid.UseCase, *RepositorySpy) {
	repositorySpy := MakeRepositorySpy()
	sut := application.NewGetCampaignByIdUseCase(repositorySpy)

	return sut, repositorySpy
}
