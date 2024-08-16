package application

import (
	"testing"

	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/model"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/error/domainError"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/pkg/expect"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/tests/domain/campaign/application/fixtures"
)

var (
	FAKE_ID string = "id_fake"
)

func Test_GivenGetCampaignByIdExecuted_WhenIdIsNullOrEmpty_ThenEnsureError(t *testing.T) {
	//Arrange
	sut, _ := fixtures.NewGetCampaignByIdUseCaseSUT()

	//Act
	_, err := sut.Execute("")

	//Assert
	expect.Equal(t, err.Error(), domainError.Err_PARAMETER_ID_NOT_EMPTY.Error())
}

func Test_GivenGetCampaignByIdExecuted_WhenRepositoryGetByIdIsCalled_ThenEnsureRepositoryCalledOnce(t *testing.T) {
	//Arrange
	sut, repositorySpy := fixtures.NewGetCampaignByIdUseCaseSUT()
	repositorySpy.WithResultError()

	//Act
	sut.Execute(FAKE_ID)

	//Assert
	expect.Equal(t, repositorySpy.CalledCount, 1)
}

func Test_GivenGetCampaignByIdExecuted_WhenRepositoryGetByIdIsCalled_ThenEnsureRepositoryCalledWithCorrectParameter(t *testing.T) {
	//Arrange
	sut, repositorySpy := fixtures.NewGetCampaignByIdUseCaseSUT()
	repositorySpy.WithResultError()

	//Act
	sut.Execute(FAKE_ID)

	//Assert
	expect.Equal(t, repositorySpy.ParamGetCampaignById, FAKE_ID)
}

func Test_GivenGetCampaignByIdExecuted_WhenRepositoryGetByIdReturnError_ThenEnsureReturnCorrectResultError(t *testing.T) {
	//Arrange
	sut, repositorySpy := fixtures.NewGetCampaignByIdUseCaseSUT()
	repositorySpy.WithResultError()

	//Act
	result, err := sut.Execute(FAKE_ID)

	//Assert
	expect.Nil(t, result)
	expect.Equal(t, err.Error(), repositorySpy.ResultError.Error())
}

func Test_GivenGetCampaignByIdExecuted_WhenRepositoryGetByIdReturnSuccess_ThenEnsureReturnCorrectResultSuccess(t *testing.T) {
	//Arrange
	sut, repositorySpy := fixtures.NewGetCampaignByIdUseCaseSUT()
	repositorySpy.WithResultModel(&model.CampaignModel{
		Id: FAKE_ID,
	})

	//Act
	result, err := sut.Execute(FAKE_ID)

	//Assert
	expect.Nil(t, err)
	expect.Equal(t, result.Id, repositorySpy.ResultCampaignModel.Id)
	expect.Equal(t, result.Id, FAKE_ID)
}
