package handler

import (
	"net/http"
	"testing"

	"github.com/rafaelbatistaroque/mail-batch-go/internal/pkg/expect"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/tests/domain/campaign/ports/handler/fixtures"
)

func Test_GivenSearchCampaign_WhenDecodeRequestBodyError_ThenEnsureReturnBadRequestWithError(t *testing.T) {
	// Arrange
	sut, _, res, req := fixtures.BuildSearchCampaignHandlerSUT([]byte("request_error"))

	// Act
	result, code, err := sut.SearchCampaign(res, req)

	// Assert
	expect.Nil(t, result)
	expect.Equal(t, code, http.StatusBadRequest)
	expect.NotNil(t, err.Error())
}

func Test_GivenSearchCampaign_WhenUseCaseInvoked_ThenEnsureCalledOnce(t *testing.T) {
	// Arrange
	sut, usecaseSpy, res, req := fixtures.BuildSearchCampaignHandlerSUT([]byte(fixtures.MakeSearchCampaignInput_Valid()))
	usecaseSpy.WithResultError()

	// Act
	sut.SearchCampaign(res, req)

	// Assert
	expect.Equal(t, usecaseSpy.CallsCount, 1)
}

func Test_GivenSearchCampaign_WhenUseCaseReturnError_ThenEnsureReturnCorrectResult(t *testing.T) {
	// Arrange
	sut, usecaseSpy, res, req := fixtures.BuildSearchCampaignHandlerSUT([]byte(fixtures.MakeSearchCampaignInput_Valid()))
	usecaseSpy.WithResultError()

	// Act
	result, code, err := sut.SearchCampaign(res, req)

	// Assert
	expect.Nil(t, result)
	expect.Equal(t, code, http.StatusInternalServerError)
	expect.NotNil(t, err.Error())
}

func Test_GivenSearchCampaign_WhenUseCaseReturnSuccess_ThenEnsureReturnCorrectResult(t *testing.T) {
	// Arrange
	sut, usecaseSpy, res, req := fixtures.BuildSearchCampaignHandlerSUT([]byte(fixtures.MakeSearchCampaignInput_Valid()))
	usecaseSpy.WithResultSuccess()

	// Act
	result, code, err := sut.SearchCampaign(res, req)

	// Assert
	expect.Nil(t, err)
	expect.Equal(t, code, http.StatusOK)
	expect.NotNil(t, result)
}
