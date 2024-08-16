package application

import (
	"fmt"
	"testing"

	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/entity"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/usecase/cancelcampaign"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/error/domainError"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/pkg/expect"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/tests/domain/campaign/application/fixtures"
)

func Test_GivenCancelCampaignExecuted_WhenInputInvalid_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut, _, _ := fixtures.NewCancelCampaignSUT()
	inputInvalid := &cancelcampaign.Input{
		Id: "",
	}

	// Act
	_, err := sut.Execute(inputInvalid)

	// Assert
	expect.Equal(t, err.Error(), fmt.Sprintf(domainError.Err_PARAMETER_NOT_EMPTY.Error(), "Id"))
}

func Test_GivenCancelCampaignExecuted_WhenRepositoryGetByIdIsCalled_ThenEnsureRepositoryCalledOnce(t *testing.T) {
	// Arrange
	sut, mapperSpy, repositorySpy := fixtures.NewCancelCampaignSUT()
	mapperSpy.WithResultEntity(nil)

	// Act
	sut.Execute(fixtures.MakeCancelCampaignInput_Valid())

	// Assert
	expect.Equal(t, repositorySpy.CalledCount, 1)
}

func Test_GivenCancelCampaignExecuted_WhenRepositoryGetByIdIsCalled_ThenEnsureRepositoryCalledWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, mapperSpy, repositorySpy := fixtures.NewCancelCampaignSUT()
	mapperSpy.WithResultEntity(nil)
	inputValid := fixtures.MakeCancelCampaignInput_Valid()

	// Act
	sut.Execute(inputValid)

	// Assert
	expect.Equal(t, repositorySpy.ParamGetCampaignById, inputValid.Id)
}

func Test_GivenCancelCampaignExecuted_WhenCampaignNotFound_ThenEnsureError(t *testing.T) {
	// Arrange
	sut, _, repositorySpy := fixtures.NewCancelCampaignSUT()
	repositorySpy.WithResultError()

	// Act
	_, err := sut.Execute(fixtures.MakeCancelCampaignInput_Valid())

	// Assert
	expect.Equal(t, err.Error(), domainError.Err_CAMPAIGN_NOT_FOUND.Error())
}

func Test_GivenCancelCampaignExecuted_WhenCampaignFound_ThenEnsureMapperToEntityCalledOnce(t *testing.T) {
	// Arrange
	sut, mapperSpy, _ := fixtures.NewCancelCampaignSUT()
	mapperSpy.WithResultEntity(nil)

	// Act
	sut.Execute(fixtures.MakeCancelCampaignInput_Valid())

	// Assert
	expect.Equal(t, mapperSpy.ToEntityCalledCount, 1)
}

func Test_GivenCancelCampaignExecuted_WhenCampaignFound_ThenEnsureMapperToEntityCalledWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, mapperSpy, repositorySpy := fixtures.NewCancelCampaignSUT()
	mapperSpy.WithResultEntity(nil)
	resultModel := fixtures.MakeCampaignModel_Fake()
	repositorySpy.WithResultModel(resultModel)

	// Act
	sut.Execute(fixtures.MakeCancelCampaignInput_Valid())

	// Assert
	expect.Equal(t, mapperSpy.ParamModel.Id, resultModel.Id)
	expect.Equal(t, mapperSpy.ParamModel.Content, resultModel.Content)
	expect.Equal(t, mapperSpy.ParamModel.CreatedOn, resultModel.CreatedOn)
	expect.ContainsAll(t, mapperSpy.ParamModel.Contacts, resultModel.Contacts)
}

func Test_GivenCancelCampaignExecuted_WhenCampaingCanceled_ThenEnsureToCampaignModelCalledOnce(t *testing.T) {
	// Arrange
	sut, mapperSpy, _ := fixtures.NewCancelCampaignSUT()
	mapperSpy.WithResultEntity(nil)

	// Act
	sut.Execute(fixtures.MakeCancelCampaignInput_Valid())

	// Assert
	expect.Equal(t, mapperSpy.ToCampaignModelCalledCount, 1)
}

func Test_GivenCancelCampaignExecuted_WhenToCampaignModelCalled_ThenEnsureToCampaignModelCalledWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, mapperSpy, _ := fixtures.NewCancelCampaignSUT()
	campaign := fixtures.MakeCampaignEntity_Fake()
	mapperSpy.WithResultEntity(campaign)

	// Act
	sut.Execute(fixtures.MakeCancelCampaignInput_Valid())

	// Assert
	expect.Equal(t, mapperSpy.ParamEntity.GetId(), campaign.GetId())
	expect.Equal(t, mapperSpy.ParamEntity.Name, campaign.Name)
	expect.Equal(t, mapperSpy.ParamEntity.Content, campaign.Content)
	expect.Equal(t, mapperSpy.ParamEntity.GetCreatedOn(), campaign.GetCreatedOn())
	expect.Equal(t, mapperSpy.ParamEntity.Status, entity.CANCELED)
	expect.ContainsAll(t, mapperSpy.ParamEntity.GetContactsString(), campaign.GetContactsString())
}

func Test_GivenCancelCampaignExecuted_WhenModelMappedToEntity_ThenEnsureRepositoryUpdateCalledOnce(t *testing.T) {
	// Arrange
	sut, mapperSpy, repositorySpy := fixtures.NewCancelCampaignSUT()
	mapperSpy.WithResultEntity(nil)

	// Act
	sut.Execute(fixtures.MakeCancelCampaignInput_Valid())

	// Assert
	expect.Equal(t, repositorySpy.UpdateCalledCount, 1)
}

func Test_GivenCancelCampaignExecuted_WhenMOdelMappedToEntity_ThenEnsureRepositoryUpdateCalledWithCorrectParameter(t *testing.T) {
	// Arrange
	sut, mapperSpy, repositorySpy := fixtures.NewCancelCampaignSUT()
	modelCanceled := fixtures.MakeCampaignModel_Fake()
	modelCanceled.Status = entity.CANCELED
	mapperSpy.WithResultEntity(fixtures.MakeCampaignEntity_Fake())
	mapperSpy.WithResultModel(modelCanceled)

	// Act
	sut.Execute(fixtures.MakeCancelCampaignInput_Valid())

	// Assert
	expect.Equal(t, repositorySpy.ParamModel.Id, modelCanceled.Id)
	expect.Equal(t, repositorySpy.ParamModel.Name, modelCanceled.Name)
	expect.Equal(t, repositorySpy.ParamModel.Content, modelCanceled.Content)
	expect.Equal(t, repositorySpy.ParamModel.CreatedOn, modelCanceled.CreatedOn)
	expect.Equal(t, repositorySpy.ParamModel.Status, entity.CANCELED)
	expect.ContainsAll(t, repositorySpy.ParamModel.Contacts, modelCanceled.Contacts)
}
