package application

import (
	"fmt"
	"testing"

	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/usecase/searchcampaign"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/error/domainError"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/pkg/expect"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/tests/domain/campaign/application/fixtures"
)

func Test_GivenSearchCampaignExecuted_WhenPageInputPropertyZero_ThenEnsureError(t *testing.T) {
	//Arrange
	sut, _, _ := fixtures.NewSearchCampaignUseCaseSUT()
	inputPropertiesEmpty := &searchcampaign.Input{
		Page:    0,
		PerPage: 10,
	}

	//Act
	_, err := sut.Execute(inputPropertiesEmpty)

	//Assert
	expect.Equal(t, err.Error(), fmt.Sprintf(domainError.Err_PROPERTIES_NOT_LESS_THAN_OR_EQUAL_ZERO.Error(), "Page"))
}

func Test_GivenSearchCampaignExecuted_WhenPerPageInputPropertyZero_ThenEnsureError(t *testing.T) {
	//Arrange
	sut, _, _ := fixtures.NewSearchCampaignUseCaseSUT()
	inputPropertiesEmpty := &searchcampaign.Input{
		Page:    1,
		PerPage: 0,
	}

	//Act
	_, err := sut.Execute(inputPropertiesEmpty)

	//Assert
	expect.Equal(t, err.Error(), fmt.Sprintf(domainError.Err_PROPERTIES_NOT_LESS_THAN_OR_EQUAL_ZERO.Error(), "PerPage"))
}

func Test_GivenSearchCampaignExecuted_WhenInputValid_ThenEnsureGetErrorEmpty(t *testing.T) {
	//Arrange
	sut, _, repositorySpy := fixtures.NewSearchCampaignUseCaseSUT()
	repositorySpy.WithResultError()
	inputValid := fixtures.MakeSearchCampaignInput_Valid()

	//Act
	sut.Execute(inputValid)

	//Assert
	expect.False(t, inputValid.IsInvalid())
}

func Test_GivenSearchCampaignExecuted_WhenMapperToSearchCampaignModelIsCalled_ThenEnsureMapperCalledOnce(t *testing.T) {
	//Arrange
	inputValid := fixtures.MakeSearchCampaignInput_Valid()
	sut, mapperSpy, repositorySpy := fixtures.NewSearchCampaignUseCaseSUT()
	repositorySpy.WithResultError()

	//Act
	sut.Execute(inputValid)

	//Assert
	expect.Equal(t, mapperSpy.ToSearchCampaignModelCalledCount, 1)
}

func Test_GivenSearchCampaignExecuted_WhenMapperToSearchCampaignModelIsCalled_ThenEnsureMapperCalledWithCorrectParameters(t *testing.T) {
	//Arrange
	inputValid := fixtures.MakeSearchCampaignInput_Valid()
	sut, mapperSpy, repositorySpy := fixtures.NewSearchCampaignUseCaseSUT()
	repositorySpy.WithResultError()

	//Act
	sut.Execute(inputValid)

	//Assert
	expect.NotNil(t, mapperSpy.ParamSearchCampaignModel)
	expect.Equal(t, inputValid.Page, mapperSpy.ParamSearchCampaignModel.Page)
	expect.Equal(t, inputValid.PerPage, mapperSpy.ParamSearchCampaignModel.PerPage)
}

func Test_GivenSearchCampaignExecuted_WhenRepositorySearchIsCalled_ThenEnsureCalledOnce(t *testing.T) {
	//Arrange
	sut, _, repositorySpy := fixtures.NewSearchCampaignUseCaseSUT()
	repositorySpy.WithResultError()

	//Act
	sut.Execute(fixtures.MakeSearchCampaignInput_Valid())

	//Assert
	expect.Equal(t, repositorySpy.CalledCount, 1)
}

func Test_GivenSearchCampaignExecuted_WhenRepositorySearchIsCalled_ThenEnsureRepositoryCalledWithCorrectParameter(t *testing.T) {
	//Arrange
	sut, mapperSpy, repositorySpy := fixtures.NewSearchCampaignUseCaseSUT()
	repositorySpy.WithResultError()

	//Act
	sut.Execute(fixtures.MakeSearchCampaignInput_Valid())

	//Assert
	expect.StrictEqual(t, repositorySpy.ParamSearchCampaignModel, mapperSpy.ResultMapToSearchCampaignModel)
}

func Test_GivenSearchCampaignExecuted_WhenRepositorySearchIsError_ThenEnsureReturnError(t *testing.T) {
	//Arrange
	sut, _, repositorySpy := fixtures.NewSearchCampaignUseCaseSUT()
	repositorySpy.WithResultError()

	//Act
	result, err := sut.Execute(fixtures.MakeSearchCampaignInput_Valid())

	//Assert
	expect.Nil(t, result)
	expect.StrictEqual(t, err, repositorySpy.ResultError)
}

func Test_GivenSearchCampaignExecuted_WhenRepositorySearchIsSuccess_ThenEnsureReturnSliceOfCampaign(t *testing.T) {
	//Arrange
	validModel := fixtures.MakeSearchCampaignModel_Valid()
	sut, _, repositorySpy := fixtures.NewSearchCampaignUseCaseSUT()
	repositorySpy.WithSearchResultSuccess(validModel)

	//Act
	result, err := sut.Execute(fixtures.MakeSearchCampaignInput_Valid())

	//Assert
	expect.Nil(t, err)
	expect.NotNil(t, result)
	expect.Equal(t, validModel.PerPage, result.Total)
}
