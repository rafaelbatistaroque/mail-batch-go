package handler

import (
	"net/http"
	"testing"

	"github.com/rafaelbatistaroque/mail-batch-go/internal/pkg/expect"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/tests/domain/campaign/ports/handler/fixtures"
)

var (
	FAKE_ID string = "idfake"
)

func Test_GivenGetCampaignById_WhenParameterExtractError_ThenEnsureReturnBadRequestWithError(t *testing.T) {
	// Arrange
	sut, _, res, req := fixtures.BuildGetCampaignByIdHandlerSUT("")

	// Act
	result, code, err := sut.GetCampaignById(res, req)

	// Assert
	expect.Nil(t, result)
	expect.Equal(t, code, http.StatusBadRequest)
	expect.NotNil(t, err.Error())
}

func Test_GivenGetCampaignById_WhenParameterInvalid_ThenEnsureReturnBadRequestWithError(t *testing.T) {
	// Arrange
	sut, _, res, req := fixtures.BuildGetCampaignByIdHandlerSUT("-")

	// Act
	result, code, err := sut.GetCampaignById(res, req)

	// Assert
	expect.Nil(t, result)
	expect.Equal(t, code, http.StatusBadRequest)
	expect.NotNil(t, err.Error())
}

func Test_GivenGetCampaignById_WhenUseCaseInvoked_ThenEnsureCalledOnce(t *testing.T) {
	// Arrange
	sut, usecaseSpy, res, req := fixtures.BuildGetCampaignByIdHandlerSUT(FAKE_ID)
	usecaseSpy.WithResultError()

	// Act
	sut.GetCampaignById(res, req)

	// Assert
	expect.Equal(t, usecaseSpy.CallsCount, 1)
}

func Test_GivenGetCampaignById_WhenUseCaseReturnError_ThenEnsureReturnCorrectResult(t *testing.T) {
	// Arrange
	sut, usecaseSpy, res, req := fixtures.BuildGetCampaignByIdHandlerSUT(FAKE_ID)
	usecaseSpy.WithResultError()

	// Act
	result, code, err := sut.GetCampaignById(res, req)

	// Assert
	expect.Nil(t, result)
	expect.Equal(t, code, http.StatusInternalServerError)
	expect.NotNil(t, err.Error())
}

func Test_GivenGetCampaignById_WhenUseCaseReturnSuccess_ThenEnsureReturnCorrectResult(t *testing.T) {
	// Arrange
	sut, usecaseSpy, res, req := fixtures.BuildGetCampaignByIdHandlerSUT(FAKE_ID)
	usecaseSpy.WithResultSuccess(FAKE_ID)

	// Act
	result, code, err := sut.GetCampaignById(res, req)

	// Assert
	expect.Nil(t, err)
	expect.Equal(t, code, http.StatusOK)
	expect.NotNil(t, result)
}
