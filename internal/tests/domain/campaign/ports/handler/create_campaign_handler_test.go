package handler

import (
	"net/http"
	"testing"

	"github.com/rafaelbatistaroque/mail-batch-go/internal/pkg/expect"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/tests/domain/campaign/ports/handler/fixtures"
)

func Test_GivenCreateCampaign_WhenDecodeRequestBodyError_ThenEnsureReturnBadRequestWithError(t *testing.T) {
	// Arrange
	sut, _, res, req := fixtures.BuildCreateCampaignHandlerSUT([]byte("request_error"))

	// Act
	result, code, err := sut.CreateCampaign(res, req)

	// Assert
	expect.Nil(t, result)
	expect.Equal(t, code, http.StatusBadRequest)
	expect.NotNil(t, err.Error())
}

func Test_GivenCreateCampaign_WhenUseCaseInvoked_ThenEnsureCalledOnce(t *testing.T) {
	// Arrange
	sut, usecaseSpy, res, req := fixtures.BuildCreateCampaignHandlerSUT([]byte(fixtures.MakeCreateCampaignInput_Valid()))
	usecaseSpy.WithResultError()

	// Act
	sut.CreateCampaign(res, req)

	// Assert
	expect.Equal(t, usecaseSpy.CallsCount, 1)
}

func Test_GivenCreateCampaign_WhenUseCaseReturnError_ThenEnsureReturnCorrectResult(t *testing.T) {
	// Arrange
	sut, usecaseSpy, res, req := fixtures.BuildCreateCampaignHandlerSUT([]byte(fixtures.MakeCreateCampaignInput_Valid()))
	usecaseSpy.WithResultError()

	// Act
	result, code, err := sut.CreateCampaign(res, req)

	// Assert
	expect.Nil(t, result)
	expect.Equal(t, code, http.StatusInternalServerError)
	expect.NotNil(t, err.Error())
}

func Test_GivenCreateCampaign_WhenUseCaseReturnSuccess_ThenEnsureReturnCorrectResult(t *testing.T) {
	// Arrange
	sut, usecaseSpy, res, req := fixtures.BuildCreateCampaignHandlerSUT([]byte(fixtures.MakeCreateCampaignInput_Valid()))
	usecaseSpy.WithResultSuccess()

	// Act
	result, code, err := sut.CreateCampaign(res, req)

	// Assert
	expect.Nil(t, err)
	expect.Equal(t, code, http.StatusCreated)
	expect.NotNil(t, result)
}
