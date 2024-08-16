package application

import (
	"fmt"
	"testing"

	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/usecase/createcampaign"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/error/domainError"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/pkg/expect"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/tests/domain/campaign/application/fixtures"
)

func Test_GivenCreateCampaignExecuted_WhenInputNameInvalid_ThenEnsureError(t *testing.T) {
	//Arrange
	sut, _, _ := fixtures.NewCreateCampaignUseCaseSUT()
	inputWithInvalidName := &createcampaign.Input{
		Name:     "",
		Content:  "fake_content",
		Contacts: []string{"t@t.com"},
	}

	//Act
	_, err := sut.Execute(inputWithInvalidName)

	//Assert
	expect.Equal(t, err.Error(), fmt.Sprintf(domainError.Err_PARAMETER_NOT_EMPTY.Error(), "Name"))
}

func Test_GivenCreateCampaignExecuted_WhenInputContentInvalid_ThenEnsureError(t *testing.T) {
	//Arrange
	sut, _, _ := fixtures.NewCreateCampaignUseCaseSUT()
	inputWithInvalidContent := &createcampaign.Input{
		Name:     "fake_name",
		Content:  "",
		Contacts: []string{"t@t.com"},
	}

	//Act
	_, err := sut.Execute(inputWithInvalidContent)

	//Assert
	expect.Equal(t, err.Error(), fmt.Sprintf(domainError.Err_PARAMETER_NOT_EMPTY.Error(), "Content"))
}

func Test_GivenCreateCampaignExecuted_WhenInputContactsInvalid_ThenEnsureError(t *testing.T) {
	//Arrange
	sut, _, _ := fixtures.NewCreateCampaignUseCaseSUT()
	inputWithInvalidContacts := &createcampaign.Input{
		Name:     "fake_name",
		Content:  "fake_content",
		Contacts: []string{},
	}

	//Act
	_, err := sut.Execute(inputWithInvalidContacts)

	//Assert
	expect.Equal(t, err.Error(), fmt.Sprintf(domainError.Err_PARAMETER_NOT_EMPTY.Error(), "Contacts"))
}

func Test_GivenCreateCampaignExecuted_WhenMapperToCampaignModelIsCalled_ThenEnsureMapperCalledOnce(t *testing.T) {
	//Arrange
	sut, mapperSpy, _ := fixtures.NewCreateCampaignUseCaseSUT()

	//Act
	sut.Execute(fixtures.MakeCreateCampaignInput_Valid())

	//Assert
	expect.Equal(t, mapperSpy.ToCampaignModelCalledCount, 1)
}

func Test_GivenCreateCampaignExecuted_WhenMapperToModelIsCalled_ThenEnsureMapperCalledWithCorrectParameters(t *testing.T) {
	//Arrange
	inputValid := fixtures.MakeCreateCampaignInput_Valid()
	sut, mapperSpy, _ := fixtures.NewCreateCampaignUseCaseSUT()

	//Act
	sut.Execute(inputValid)

	//Assert
	expect.NotNil(t, mapperSpy.ParamEntity.GetId())
	expect.False(t, mapperSpy.ParamEntity.GetCreatedOn().IsZero())
	expect.Equal(t, mapperSpy.ParamEntity.Name, inputValid.Name)
	expect.Equal(t, mapperSpy.ParamEntity.Content, inputValid.Content)
	expect.ContainsAll(t, mapperSpy.ParamEntity.GetContactsString(), inputValid.Contacts)
}

func Test_GivenCreateCampaignExecuted_WhenRepositorySaveIsCalled_ThenEnsureCalledOnce(t *testing.T) {
	//Arrange
	sut, _, repositorySpy := fixtures.NewCreateCampaignUseCaseSUT()

	//Act
	sut.Execute(fixtures.MakeCreateCampaignInput_Valid())

	//Assert
	expect.Equal(t, repositorySpy.CalledCount, 1)
}

func Test_GivenCreateCampaignExecuted_WhenRepositorySaveIsCalled_ThenEnsureRepositoryCalledWithCorrectParameter(t *testing.T) {
	//Arrange
	sut, mapperSpy, repositorySpy := fixtures.NewCreateCampaignUseCaseSUT()
	mapperSpy.WithResultModel(nil)

	//Act
	sut.Execute(fixtures.MakeCreateCampaignInput_Valid())

	//Assert
	expect.StrictEqual(t, repositorySpy.ParamModel, mapperSpy.ResultMapToModel)
}

func Test_GivenCreateCampaignExecuted_WhenRepositorySaveIsError_ThenEnsureReturnError(t *testing.T) {
	//Arrange
	sut, _, repositorySpy := fixtures.NewCreateCampaignUseCaseSUT()
	repositorySpy.WithResultError()

	//Act
	result, err := sut.Execute(fixtures.MakeCreateCampaignInput_Valid())

	//Assert
	expect.Nil(t, result)
	expect.StrictEqual(t, err, repositorySpy.ResultError)
}

func Test_GivenCreateCampaignExecuted_WhenRepositorySaveIsSuccess_ThenEnsureReturnCreatedCampaingData(t *testing.T) {
	//Arrange
	sut, mapperSpy, _ := fixtures.NewCreateCampaignUseCaseSUT()
	mapperSpy.WithResultModel(nil)

	//Act
	result, err := sut.Execute(fixtures.MakeCreateCampaignInput_Valid())

	//Assert
	expect.Nil(t, err)
	expect.Equal(t, result.Id, mapperSpy.ResultMapToModel.Id)
	expect.Equal(t, result.Name, mapperSpy.ResultMapToModel.Name)
	expect.Equal(t, result.Content, mapperSpy.ResultMapToModel.Content)
	expect.ContainsAll(t, result.Contacts, mapperSpy.ResultMapToModel.Contacts)
}
